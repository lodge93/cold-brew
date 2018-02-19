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

	// The slowest speed at which the motor still rotates.
	dripSpeed = 100

	// The maximum speed the motor can be set to.
	maxSpeed = 255

	// Time in milliseconds the motor is turned on for in order to produce one
	// drip at the drip speed.
	dripDuration = 250

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
	pump *i2c.AdafruitMotorHatDriver

	// The motor number on the Adafruit Motor HAT indexed on zero.
	motorNum int

	// A wait group used to ensure the dipper goroutine has been stopped
	// successfully.
	dripperWG sync.WaitGroup

	// A channel used to send a stop signal to the dripper goroutine.
	stopDripper chan bool

	// The drips per minute to set the dripper to.
	dripsPerMin float64

	// A mutex used to lock the dipper object when updating the drips per minute
	// outside the drip goroutine.
	mutex sync.Mutex

	// State is used internally to track the state of the dripper hardware.
	state string
}

// New creates a new dripper instance and initializes the motor HAT driver.
func New() (*Dripper, error) {
	r := raspi.NewAdaptor()

	d := Dripper{
		motorNum:    2,
		pump:        i2c.NewAdafruitMotorHatDriver(r),
		state:       OFF,
		stopDripper: make(chan bool),
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
	err := d.setSpeed(dripSpeed)
	if err != nil {
		return err
	}

	// This is a sanity check to ensure the dripper is not already in the drip
	// state.
	if d.state == DRIP {
		return nil
	}

	d.state = DRIP
	d.dripsPerMin = dripsPerMin

	d.dripperWG.Add(1)
	go d.runDrip()

	return nil
}

// Run turns on the dripper at the maximum pump speed. This is useful for
// blooming the batch of coffee, priming the pump with water, and draining the
// pump when finished brewing.
func (d *Dripper) Run() error {
	// This is a sanity check to ensure the drip goroutine is stopped before
	// trying to control the pump. This prevents the weird state where the pump
	// is on the maximum speed, but is still pulsing from the drip goroutine.
	if d.state == DRIP {
		d.Off()
	}

	err := d.setSpeed(maxSpeed)
	if err != nil {
		return err
	}

	d.state = RUN

	err = d.on()
	if err != nil {
		return err
	}

	return nil
}

// Off ensures the dripper is completely stopped.
func (d *Dripper) Off() error {
	if d.state == DRIP {
		d.stopDripper <- true
		d.dripperWG.Wait()
	}
	d.state = OFF

	return d.stop()
}

// SetDripsPerMinute will update the dripper with the desired drip rate.
func (d *Dripper) SetDripsPerMinute(dripsPerMin float64) {
	d.mutex.Lock()
	d.dripsPerMin = dripsPerMin
	d.mutex.Unlock()
}

// GetState returns the internal state of the dripper. This is useful for
// clients to determine the state of the dripper when reconnecting.
func (d *Dripper) GetState() string {
	// TODO: The mutexes in this package need updated so there are two separate
	// mutexes. Also, this will need a mutex wrapper once there are two separate
	// mutexes.
	return d.state
}

// GetDripsPerMinute will return the current drip rate from the dripper.
func (d *Dripper) GetDripsPerMinute() float64 {
	d.mutex.Lock()
	dpm := d.dripsPerMin
	d.mutex.Unlock()

	return dpm
}

// runDrip runs a goroutine to pulse the dripper at the desired drip rate. This
// is mainly to account for the fact that the cheap peristaltic pump used in
// this design as the motor cannot rotate much lower then the drip speed
// constant. In order to get just a small amount of water out of the pump, it is
// simply turned on at the lowest speed for some amount of time before being
// stopped to simulate a single drip.
func (d *Dripper) runDrip() {
	// TODO: this method relies heavily on time.Sleep which is problematic as it
	// does not lead to a highly accurate drip rate. For example, 60 drips at 60
	// drips per minute will execute in ~60.5 seconds. While good enough for the
	// initial iteration, this should be refactored to reflect a more accurate
	// drip rate as this will run for several hours and could drift quite a bit.
	defer d.dripperWG.Done()

	for {
		select {
		case <-d.stopDripper:
			return
		default:
			err := d.on()
			if err != nil {
				log.Println(err)
			}

			time.Sleep(dripDuration * time.Millisecond)

			err = d.stop()
			if err != nil {
				log.Println(err)
			}
		}

		select {
		case <-d.stopDripper:
			return
		default:
			d.mutex.Lock()
			dpm := d.dripsPerMin
			d.mutex.Unlock()
			stopDuration := calcStopDuration(dpm)
			time.Sleep(time.Duration((stopDuration * 1000)) * time.Millisecond)
		}
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
func calcStopDuration(dripsPerMin float64) float64 {
	if dripsPerMin > 240 {
		dripsPerMin = 240
	}

	secondsPerDrip := secondsPerMin / dripsPerMin
	stop := secondsPerDrip - (dripDuration / millisecondsPerSec)
	return stop
}
