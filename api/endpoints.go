// Copyright © 2018 Mark Spicer
// Made available under the MIT license.

package api

// DripperEndpoint is a data model for the dripper endpoints.
type DripperEndpoint struct {
	DripsPerMinute float64 `json:"dripsPerMinute" binding:"required"`
	State          string  `json:"state"`
}
