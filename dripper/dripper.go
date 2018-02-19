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
}

// New creates a new dripper instance and initializes the motor HAT driver.
func New() (*Dripper, error) {
	var d Dripper

	// This is the motor number on the Adafruit Motor HAT indexed on zero. In
	// this case, 2 is M3.
	d.motorNum = 2

	r := raspi.NewAdaptor()
	d.pump = i2c.NewAdafruitMotorHatDriver(r)

	// TODO: The start method attempts to initialize both the servo and motor
	// drivers on the HAT. This is problematic because this package only uses
	// the motor interface. In order to account for this, edits have been made
	// to the vendored dependency, but this should really be a pull request:
	// https://github.com/hybridgroup/gobot/blob/master/drivers/i2c/adafruit_driver.go#L201
	err := d.pump.Start()
	if err != nil {
		return &d, err
	}

	d.stopDripper = make(chan bool)

	return &d, nil
}

// TODO: this will likely need refactored as there is no current way to stop
// a bloom. In addition, the bloom duration would logically mean the amount of
// time to wait after the water has been added where it is the amount of time
// the motor is on in the current state.
func (d *Dripper) Bloom(dur time.Duration) error {
	err := d.setSpeed(maxSpeed)
	if err != nil {
		return err
	}

	err = d.on()
	if err != nil {
		return err
	}

	time.Sleep(dur)

	err = d.off()
	if err != nil {
		return err
	}

	return nil
}

func (d *Dripper) Drip(dripsPerMin float64) error {
	err := d.setSpeed(dripSpeed)
	if err != nil {
		return err
	}

	d.dripsPerMin = dripsPerMin

	d.dripperWG.Add(1)
	go d.runDrip()

	return nil
}

func (d *Dripper) Stop() error {
	d.stopDripper <- true
	d.dripperWG.Wait()

	return d.off()
}

func (d *Dripper) SetDripsPerMinute(dripsPerMin float64) {
	d.mutex.Lock()
	d.dripsPerMin = dripsPerMin
	d.mutex.Unlock()
}

// TODO: this method relies heavily on time.Sleep which is problematic as it
// does not lead to a highly accurate drip rate. For example, 60 drips at 60
// drips per minute will execute in ~60.5 seconds. While good enough for the
// intial iteration, this should be refactored to reflect a more accurate drip
// rate as this will run for several hours and could drift quite a bit.
func (d *Dripper) runDrip() {
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

			err = d.off()
			if err != nil {
				log.Println(err)
			}
		}

		select {
		case <-d.stopDripper:
			return
		default:
			d.mutex.Lock()
			stopDuration := calcStopDuration(d.dripsPerMin)
			d.mutex.Unlock()
			time.Sleep(time.Duration((stopDuration * 1000)) * time.Millisecond)
		}
	}
}

func (d *Dripper) on() error {
	return d.pump.RunDCMotor(d.motorNum, i2c.AdafruitForward)
}

func (d *Dripper) setSpeed(speed int32) error {
	return d.pump.SetDCMotorSpeed(d.motorNum, speed)
}

func (d *Dripper) off() error {
	return d.pump.RunDCMotor(d.motorNum, i2c.AdafruitRelease)
}

func calcStopDuration(dripsPerMin float64) float64 {
	if dripsPerMin > 240 {
		dripsPerMin = 240
	}

	secondsPerDrip := secondsPerMin / dripsPerMin
	stop := secondsPerDrip - (dripDuration / millisecondsPerSec)
	return stop
}
