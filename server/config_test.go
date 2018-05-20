// Copyright © 2018 Mark Spicer
// Made available under the MIT license.

package server

import (
	"testing"
)

func TestNewConfigReadsFromTestData(t *testing.T) {
	config, err := NewConfig()
	if err != nil {
		t.Error("could not load configuration file:", err)
	}

	if config.DatabaseDir != "/foo/bar" {
		t.Error("incorrect database file path was loaded from config file")
	}

	if config.Environment != "testing" {
		t.Error("incorrect environment loaded from config file")
	}
}

func TestIsValidEnvironmentWhenEnvironmentIsValid(t *testing.T) {
	if !isValidEnvironment(EnvDevelopment) {
		t.Error("valid environment did not return as valid")
	}

	if !isValidEnvironment(EnvProduction) {
		t.Error("valid environment did not return as valid")
	}

	if !isValidEnvironment(EnvTesting) {
		t.Error("valid environment did not return as valid")
	}
}

func TestIsValidEnvironmentWhenEnvironmenIsNotValid(t *testing.T) {
	if isValidEnvironment("foo") {
		t.Error("invalid environment returned as valid")
	}
}
