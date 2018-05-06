// Copyright © 2018 Mark Spicer
// Made available under the MIT license.
package dripper

import (
	"testing"
)

func TestDefaultConfigReturnsDefaultValues(t *testing.T) {
	config := DefaultConfig()

	if config.DripSpeed != DefaultDripSpeed {
		t.Error("configured drip speed does not match default")
	}

	if config.RunSpeed != DefaultRunSpeed {
		t.Error("configured run speed does not match default")
	}

	if config.DripDuration != DefaultDripDuration {
		t.Error("configured duration does not match default")
	}
}
