// Copyright © 2018 Mark Spicer
// Made available under the MIT license.

package dripper

import (
	"testing"

	"github.com/betterengineering/cold-brew/pkg/dripper/mock_dripper"
	"github.com/golang/mock/gomock"
)

type testDripper struct {
	dripper             *Dripper
	mockCtrl            *gomock.Controller
	mockMotorController *mock_dripper.MockMotorController
	t                   *testing.T
}

func setup(t *testing.T) testDripper {
	mockCtrl := gomock.NewController(t)
	mockMotorController := mock_dripper.NewMockMotorController(mockCtrl)

	config := DefaultSettings()
	d := Dripper{
		motorNum:    2,
		pump:        mockMotorController,
		state:       OFF,
		stopDripper: make(chan bool, 1),
		Settings:    config,
	}

	return testDripper{
		dripper:             &d,
		mockCtrl:            mockCtrl,
		mockMotorController: mockMotorController,
		t:                   t,
	}

}

func TestDripWhenStateIsNotDrip(t *testing.T) {
	d := setup(t)
	defer d.teardown()

	dripsPerMinute := 60.0

	d.givenMotorIsSetToDrip()
	d.whenDrip(dripsPerMinute)
	d.ensureDripsPerMinuteIsSet(dripsPerMinute)
	d.ensureStateIsDrip()
	// TODO: ensure drip go routine is started.
}

func TestDripWhenStateIsDrip(t *testing.T) {
	d := setup(t)
	defer d.teardown()

	dripsPerMinute := 60.0

	d.givenMotorIsSetToDrip()
	d.givenStateIsDrip()
	d.whenDrip(dripsPerMinute)
	d.ensureDripsPerMinuteIsSet(dripsPerMinute)
	d.ensureStateIsDrip()
	// TODO: ensure drip go routine was not started.
}

func (d *testDripper) teardown() {
	d.dripper.stopDripper <- true
	d.dripper.dripperWG.Wait()
	d.mockCtrl.Finish()
}

func (d *testDripper) givenStateIsDrip() {
	d.dripper.state = DRIP
}

func (d *testDripper) givenMotorIsSetToDrip() {
	d.mockMotorController.EXPECT().SetDCMotorSpeed(d.dripper.motorNum, d.dripper.Settings.DripSpeed)
}

func (d *testDripper) whenDrip(dripsPerMinute float64) {
	d.dripper.Drip(dripsPerMinute)
}

func (d *testDripper) ensureDripsPerMinuteIsSet(dripsPerMinute float64) {
	if d.dripper.dripsPerMin != dripsPerMinute {
		d.t.Error("drips per minute was not properly set")
	}
}

func (d *testDripper) ensureStateIsDrip() {
	if d.dripper.state != DRIP {
		d.t.Error("dripper state was not set to drip")
	}
}
