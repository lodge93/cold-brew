// Copyright © 2018 Mark Spicer
// Made available under the MIT license.

package server

import (
	"errors"

	"github.com/spf13/viper"
)

const (
	// EnvDevelopment is a constant used to determine if the application is in a
	// development environment.
	EnvDevelopment = "development"

	// EnvProduction is a constant used to determine if the application is in a
	// production environment.
	EnvProduction = "production"

	// EnvTesting is a constant used to determine if the application is in a
	// testing environment.
	EnvTesting = "testing"
)

// Config is a configuration struct used by the server package to configure
// downstream dependencies.
type Config struct {
	// Environment is used to determine which environment the cold brew dirpper
	// application is being run in.
	Environment string

	// DBFile is the aboslute path of the SQLite database.
	DBFile string
}

// NewConfig returns a new configuration struct populated from a config file.
func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("testdata")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	if !viper.IsSet("databaseFile") {
		return nil, errors.New("databaseFile key is not set in the configuration file")
	}
	dbFile := viper.GetString("databaseFile")

	if !viper.IsSet("environment") {
		return nil, errors.New("environment key is not set in the configuration file")
	}
	env := viper.GetString("environment")
	if !isValidEnvironment(env) {
		return nil, errors.New("environment is not valid")
	}

	return &Config{
		Environment: env,
		DBFile:      dbFile,
	}, nil
}

// isValidEnvironment validates that an environment is one that the application
// respects.
func isValidEnvironment(env string) bool {
	if env == EnvDevelopment {
		return true
	}

	if env == EnvProduction {
		return true
	}

	if env == EnvTesting {
		return true
	}

	return false
}
