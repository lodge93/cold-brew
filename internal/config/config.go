// Copyright 2019 Mark Spicer
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
// documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
// Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
// WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Package config provides configuration for the rest of the application.
package config

import (
  "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"
  "strings"
)

// Config is a configuration object for an application.
type Config struct {
  // Env is the environment in which the application is running in.
  Env Environment

  // Logging is a configuration object to hold log configuration information.
  Logging *LoggingConfig
}

// NewConfig provides an instantiated configuration object.
func NewConfig(appName string) (*Config, error) {
  err := initConfig(appName)
  if err != nil {
    return nil, err
  }

  env, err := getEnvironment()
  if err != nil {
    return nil, err
  }

  logCfg, err := getLoggingConfig()
  if err != nil {
    return nil, err
  }

  return &Config{
    Env: env,
    Logging: logCfg,
  }, nil
}

func initConfig(appName string) error {
  userHomeDir, err := homedir.Dir()
  if err != nil {
    return err
  }

  viper.SetEnvPrefix("COLD_BREW")
  viper.AutomaticEnv()
  viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

  viper.SetConfigName("config")
  viper.AddConfigPath("/etc/" + appName + "/")
  viper.AddConfigPath("testdata/")
  viper.AddConfigPath(userHomeDir + "/." + appName + "/")

  return viper.ReadInConfig()
}
