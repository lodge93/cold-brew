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

package config_test

import (
  "github.com/lodge93/cold-brew/internal/config"
  "github.com/sirupsen/logrus"
  "os"
  "testing"
)

const (
  TestAppName = "foo-service"
)

func TestNewConfigFromFile(t *testing.T) {
  expected := config.Config{
    Env: config.Development,
    Logging: &config.LoggingConfig{
      LogLevel: logrus.WarnLevel,
    },
  }

  cfg, err := config.NewConfig(TestAppName)
  if err != nil {
    t.Fatalf("recieved and error when expected to be nil: %s", err)
  }

  if cfg.Env != expected.Env {
    t.Errorf("actual: '%s' does not match expected: '%s'", cfg.Env, expected.Env)
  }

  if cfg.Logging.LogLevel != expected.Logging.LogLevel {
    t.Errorf("actual: '%s' does not match expected: '%s'", cfg.Logging.LogLevel, expected.Logging.LogLevel)
  }
}

func TestEnvironmentOverride(t *testing.T) {
  err := os.Setenv("COLD_BREW_ENVIRONMENT", config.Production)
  if err != nil {
    t.Fatalf("could not instantiate test: %s", err)
  }
  defer os.Clearenv()

  expected := config.Config{
    Env: config.Production,
  }

  cfg, err := config.NewConfig(TestAppName)
  if err != nil {
    t.Fatalf("recieved and error when expected to be nil: %s", err)
  }

  if cfg.Env != expected.Env {
    t.Errorf("actual: '%s' does not match expected: '%s'", cfg.Env, expected.Env)
  }
}

