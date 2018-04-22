// Copyright © 2018 Mark Spicer
// Made available under the MIT license.
package dripper

import (
	"testing"
)

func TestDefaultConfigReturnsDefaultValues(t *testing.T) {
	config := DefaultConfig()

	if config.DripSpeed != DefaultDripSpeed {
		t.Fail()
	}

	if config.RunSpeed != DefaultRunSpeed {
		t.Fail()
	}

	if config.DripDuration != DefaultDripDuration {
		t.Fail()
	}
}
