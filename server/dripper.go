// Copyright © 2018 Mark Spicer
// Made available under the MIT license.

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lodge93/cold-brew/api"
)

// GetDripper returns the current state of the cold brew dripper.
func (s *Server) GetDripper(c *gin.Context) {
	c.JSON(http.StatusOK, api.DripperEndpoint{
		State:          s.Dripper.GetState(),
		DripsPerMinute: s.Dripper.GetDripsPerMinute(),
	})
}

// SetDripperRun sets the dripper to the run state.
func (s *Server) SetDripperRun(c *gin.Context) {
	err := s.Dripper.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, api.DripperEndpoint{
		State:          s.Dripper.GetState(),
		DripsPerMinute: s.Dripper.GetDripsPerMinute(),
	})
}

// SetDripperOff sets the dripper to the off state.
func (s *Server) SetDripperOff(c *gin.Context) {
	err := s.Dripper.Off()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, api.DripperEndpoint{
		State:          s.Dripper.GetState(),
		DripsPerMinute: s.Dripper.GetDripsPerMinute(),
	})
}

// SetDripperDrip sets the dripper to the drip state.
func (s *Server) SetDripperDrip(c *gin.Context) {
	var json api.DripperEndpoint
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The request submitted either was not JSON or did not contain the proper fields."})
		return
	}

	if json.DripsPerMinute > 240.0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dripsPerMinute must not exceed 240."})
		return
	}

	err = s.Dripper.Drip(json.DripsPerMinute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, api.DripperEndpoint{
		State:          s.Dripper.GetState(),
		DripsPerMinute: s.Dripper.GetDripsPerMinute(),
	})
}
