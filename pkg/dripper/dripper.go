// Copyright © 2018 Mark Spicer
// Made available under the MIT license.

// Package dripper provides an interface to interact with the underlying cold
// brew dripper hardware.
package dripper

import (
	"log"
	"sync"
	"time"

	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

const (
	// Number of seconds per minute which is used to convert between minutes
	// and seconds.
	secondsPerMin = 60

	// Number of milliseconds per second which is used to convert between
	// milliseconds and seconds.
	millisecondsPerSec = 1000

	// RUN is used for internal state tracking to represent the pump fully
	// on, which is useful for blooming, priming, and draining the pump.
	RUN = "run"

	// DRIP is used for internal state tracking to represent the pump in the
	// dripping state.
	DRIP = "drip"

	// OFF is used for internal state tracking to represent the pump fully
	// stopped.
	OFF = "off"
)

// Dripper is the base object used to implement methods to control the cold brew
// coffee dripper.
type Dripper struct {
	// The motor driver for the peristaltic pump.
	pump MotorController

	// The motor number on the Adafruit Motor HAT indexed on zero.
	motorNum int

	// A wait group used to ensure the dipper goroutine has been stopped
	// successfully.
	dripperWG sync.WaitGroup

	// A channel used to send a stop signal to the dripper goroutine.
	stopDripper chan bool

	// The drips per minute to set the dripper to.
	dripsPerMin float64

	// dripsPerMinMutex is used to update the dripsPerMinute across multiple
	// goroutines.
	dripsPerMinMutex sync.Mutex

	// State is used internally to track the state of the dripper hardware.
	state string

	// stateMutex is used to modify the dripper state across multiple
	// goroutines.
	stateMutex sync.Mutex

	// Settings is a dripper configuration object used to set values for the
	// dripper.
	Settings Settings
}

// MotorController is an interface to allow the pump to be mocked out in tests.
type MotorController interface {
	Start() error
	RunDCMotor(int, i2c.AdafruitDirection) error
	SetDCMotorSpeed(int, int32) error
}

// New creates a new dripper instance and initializes the motor HAT driver.
func New(config Settings) (*Dripper, error) {
	r := raspi.NewAdaptor()

	d := Dripper{
		motorNum:    2,
		pump:        i2c.NewAdafruitMotorHatDriver(r),
		state:       OFF,
		stopDripper: make(chan bool),
		Settings:    config,
	}

	// TODO: The start method attempts to initialize both the servo and motor
	// drivers on the HAT. This is problematic because this package only uses
	// the motor interface. In order to account for this, edits have been made
	// to the vendored dependency, but this should really be a pull request:
	// https://github.com/hybridgroup/gobot/blob/master/drivers/i2c/adafruit_driver.go#L201
	err := d.pump.Start()
	if err != nil {
		return &d, err
	}

	return &d, nil
}

// Drip starts the dripper at the desired drip rate.
func (d *Dripper) Drip(dripsPerMin float64) error {
	err := d.setSpeed(d.Settings.DripSpeed)
	if err != nil {
		return err
	}

	// We are setting dripsPerMin before the sanity check in order to update
	// regardless.
	d.SetDripsPerMinute(dripsPerMin)

	// This is a sanity check to ensure the dripper is not already in the drip
	// state.
	if d.GetState() == DRIP {
		return nil
	}

	d.setState(DRIP)

	d.dripperWG.Add(1)
	go d.runDrip()

	return nil
}

// setState is used as a setter to set the dripper state concurrently.
func (d *Dripper) setState(state string) {
	d.stateMutex.Lock()
	d.state = state
	d.stateMutex.Unlock()
}

// Run turns on the dripper at the maximum pump speed. This is useful for
// blooming the batch of coffee, priming the pump with water, and draining the
// pump when finished brewing.
func (d *Dripper) Run() error {
	// This is a sanity check to ensure the drip goroutine is stopped before
	// trying to control the pump. This prevents the weird state where the pump
	// is on the maximum speed, but is still pulsing from the drip goroutine.
	if d.GetState() == DRIP {
		d.Off()
	}

	err := d.setSpeed(d.Settings.RunSpeed)
	if err != nil {
		return err
	}

	d.setState(RUN)

	err = d.on()
	if err != nil {
		return err
	}

	return nil
}

// Off ensures the dripper is completely stopped.
func (d *Dripper) Off() error {
	if d.GetState() == DRIP {
		d.stopDripper <- true
		d.dripperWG.Wait()
	}

	d.setState(OFF)

	return d.stop()
}

// SetDripsPerMinute will update the dripper with the desired drip rate.
func (d *Dripper) SetDripsPerMinute(dripsPerMin float64) {
	d.dripsPerMinMutex.Lock()
	d.dripsPerMin = dripsPerMin
	d.dripsPerMinMutex.Unlock()
}

// GetState returns the internal state of the dripper. This is useful for
// clients to determine the state of the dripper when reconnecting.
func (d *Dripper) GetState() string {
	d.stateMutex.Lock()
	state := d.state
	d.stateMutex.Unlock()

	return state
}

// GetDripsPerMinute will return the current drip rate from the dripper.
func (d *Dripper) GetDripsPerMinute() float64 {
	d.dripsPerMinMutex.Lock()
	dpm := d.dripsPerMin
	d.dripsPerMinMutex.Unlock()

	return dpm
}

// runDrip runs a goroutine to pulse the dripper at the desired drip rate. This
// is mainly to account for the fact that the cheap peristaltic pump used in
// this design as the motor cannot rotate much lower then the drip speed
// constant. In order to get just a small amount of water out of the pump, it is
// simply turned on at the lowest speed for some amount of time before being
// stopped to simulate a single drip.
func (d *Dripper) runDrip() {
	defer d.dripperWG.Done()

	for {
		select {
		case <-d.stopDripper:
			return
		default:
			go d.drip()
			dpm := d.GetDripsPerMinute()
			dripDuration := d.Settings.DripDuration
			stopDuration := calcStopDuration(dpm, dripDuration)
			time.Sleep(time.Duration((stopDuration * 1000)) * time.Millisecond)
		}
	}
}

// drip is a low level method to produce one drip from the dripper.
func (d *Dripper) drip() {
	err := d.on()
	if err != nil {
		log.Println(err)
	}

	time.Sleep(time.Duration(d.Settings.DripDuration) * time.Millisecond)

	err = d.stop()
	if err != nil {
		log.Println(err)
	}
}

// on is a low level mehtod to start the rotation of the motor.
func (d *Dripper) on() error {
	return d.pump.RunDCMotor(d.motorNum, i2c.AdafruitForward)
}

// setSpeed is a low level method to set the motor speed.
func (d *Dripper) setSpeed(speed int32) error {
	return d.pump.SetDCMotorSpeed(d.motorNum, speed)
}

// stop is a low level method to stop the rotation of the motor.
func (d *Dripper) stop() error {
	return d.pump.RunDCMotor(d.motorNum, i2c.AdafruitRelease)
}

// calcStopDuration calculates the amount of time between drips is necessary
// to achieve the desired drip rate.
func calcStopDuration(dripsPerMin float64, dripDuration int64) float64 {
	if dripsPerMin > 240 {
		dripsPerMin = 240
	}

	secondsPerDrip := secondsPerMin / dripsPerMin
	stop := secondsPerDrip - (float64(dripDuration) / millisecondsPerSec)
	return stop
}
