// Copyright © 2018 Mark Spicer
// Made available under the MIT license.

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lodge93/cold-brew/dripper"
)

const (
	settingsCollection = "settings"
	settingsResource   = "dripper"
)

// GetDripperSettings returns the current configuration of the dripper.
func (s *Server) GetDripperSettings(c *gin.Context) {
	c.JSON(http.StatusOK, s.Dripper.Settings)
}

// SetDripperSettings reinitializes the dripper with the supplied config.
func (s *Server) SetDripperSettings(c *gin.Context) {
	var settings dripper.Settings
	err := c.BindJSON(&settings)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the request submitted either was not JSON or did not contain the proper fields"})
		return
	}

	err = s.writeDripperSettingsToDB(settings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "the settings could not be written to the database"})
		return
	}

	s.Dripper.Off()

	d, err := dripper.New(settings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "the dripper could not be reinitialized"})
		return
	}

	s.Dripper = d
	c.JSON(http.StatusOK, s.Dripper.Settings)
}

// readDripperSettingsFromDB reads the dripper settings from the database.
func (s *Server) readDripperSettingsFromDB() (dripper.Settings, error) {
	settings := dripper.Settings{}
	err := s.DB.Read(settingsCollection, settingsResource, &settings)
	if err != nil {
		return settings, err
	}

	return settings, nil
}

// writeDripperSettingsToDB writes the supplied dripper settings to the
// database.
func (s *Server) writeDripperSettingsToDB(settings dripper.Settings) error {
	err := s.DB.Write(settingsCollection, settingsResource, settings)
	if err != nil {
		return err
	}

	return nil
}

// readSettingsOrDefault attempts to read the dripper settings from the database
// and returns the default dripper settings on error. This allows the server to
// start up even if the dripper settings have not yet been modified.
func (s *Server) readSettingsOrDefault() dripper.Settings {
	settings, err := s.readDripperSettingsFromDB()
	if err != nil {
		return dripper.DefaultSettings()
	}

	return settings
}
