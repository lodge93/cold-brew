// Copyright © 2018 Mark Spicer
// Made available under the MIT license.
package dripper

import (
	"testing"
)

func TestDefaultSettingsReturnsDefaultValues(t *testing.T) {
	settings := DefaultSettings()

	if settings.DripSpeed != DefaultDripSpeed {
		t.Error("configured drip speed does not match default")
	}

	if settings.RunSpeed != DefaultRunSpeed {
		t.Error("configured run speed does not match default")
	}

	if settings.DripDuration != DefaultDripDuration {
		t.Error("configured duration does not match default")
	}
}
