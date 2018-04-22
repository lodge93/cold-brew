// Copyright © 2018 Mark Spicer
// Made available under the MIT license.

package dripper

const (
	// DefaultDripDuration is a sane default for the DripDuration configuration
	// settting.
	DefaultDripDuration = 250

	// DefaultDripSpeed is a sane default for the DripSpeed configuration
	// setting.
	DefaultDripSpeed = 100

	// DefaultRunSpeed is a sane default for the RunSpeed configuration setting.
	DefaultRunSpeed = 255
)

// Config is a configuration object used to configure dripper settings.
type Config struct {
	// DripDuration is the time in milliseconds the motor is turned on for in
	// order to produce one drip at the drip speed.
	DripDuration int64

	// DripSpeed is the slowest speed at which the motor still rotates.
	DripSpeed int32

	// RunSpeed is the fastest speed the motor will rotate.
	RunSpeed int32
}

// DefaultConfig config returns a configuration object with sane defaults.
func DefaultConfig() Config {
	return Config{
		DripDuration: DefaultDripDuration,
		DripSpeed:    DefaultDripSpeed,
		RunSpeed:     DefaultRunSpeed,
	}
}
