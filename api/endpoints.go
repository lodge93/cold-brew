// Copyright © 2018 Mark Spicer
// Made available under the MIT license.

// Package api provides a shared data model for the application and potential
// clients.
package api

// DripperEndpoint is a data model for the dripper endpoints.
type DripperEndpoint struct {
	DripsPerMinute float64 `json:"dripsPerMinute" binding:"required"`
	State          string  `json:"state"`
}
