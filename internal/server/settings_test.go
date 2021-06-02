// Copyright © 2018 Mark Spicer
// Made available under the MIT license.

package server

import (
	"errors"
	"reflect"
	"testing"

	"github.com/betterengineering/cold-brew/pkg/dripper"
)

func TestWriteDripperSettingsInDB(t *testing.T) {
	s, dir, err := withTestServerStruct()
	if err != nil {
		t.Fatal("could not create test server struct:", err)
	}
	defer cleanUpTempDatabase(dir)

	err = ensureSettingsCanBeWritten(s)
	if err != nil {
		t.Fatal("could not fetch settings from db:", err)
	}
}

func TestReadDripperSettingsOrDefaultWhenSettingsExist(t *testing.T) {
	s, dir, err := withTestServerStruct()
	if err != nil {
		t.Fatal("could not create test server struct:", err)
	}
	defer cleanUpTempDatabase(dir)

	settings := dripper.Settings{
		DripDuration: 200,
		DripSpeed:    200,
		RunSpeed:     200,
	}
	err = s.writeDripperSettingsToDB(settings)
	if err != nil {
		t.Fatal("could not write dripper settings to db")
	}

	readSettings := s.readSettingsOrDefault()
	if !reflect.DeepEqual(settings, readSettings) {
		t.Fatal("settings that existed were not returned")
	}
}

func TestReadDripperSettingsOrDefaultWhenSettingsDoNotExist(t *testing.T) {
	s, dir, err := withTestServerStruct()
	if err != nil {
		t.Fatal("could not create test server struct:", err)
	}
	defer cleanUpTempDatabase(dir)

	defaultSettings := dripper.DefaultSettings()

	readSettings := s.readSettingsOrDefault()
	if !reflect.DeepEqual(defaultSettings, readSettings) {
		t.Fatal("settings that existed were not returned")
	}
}

func ensureSettingsCanBeWritten(s *Server) error {
	settings := dripper.DefaultSettings()
	err := s.writeDripperSettingsToDB(settings)
	if err != nil {
		return err
	}

	readSettings, err := s.readDripperSettingsFromDB()
	if err != nil {
		return err
	}

	if !reflect.DeepEqual(settings, readSettings) {
		return errors.New("settings that were read do not match what was written")
	}

	return nil
}
