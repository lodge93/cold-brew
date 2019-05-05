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

package config

import (
  "github.com/spf13/viper"
  "testing"
)

func TestEnvironmentProduction(t *testing.T) {
  viper.Set(environmentViperKey, Production)
  defer viper.Reset()

  env, err := getEnvironment()
  if err != nil {
    t.Fatalf("recieved and error when expected to be nil: %s", err)
  }

  if env != Production {
    t.Errorf("actual: '%s' does not match expected: '%s'", env, Production)
  }
}

func TestEnvironmentDevelopment(t *testing.T) {
  viper.Set(environmentViperKey, Development)
  defer viper.Reset()

  env, err := getEnvironment()
  if err != nil {
    t.Fatalf("recieved and error when expected to be nil: %s", err)
  }

  if env != Development {
    t.Errorf("actual: '%s' does not match expected: '%s'", env, Development)
  }
}

func TestEnvironmentUnknown(t *testing.T) {
  viper.Set(environmentViperKey, "foobar")
  defer viper.Reset()

  _, err := getEnvironment()
  if err != ErrInvalidEnvironment {
    t.Errorf("actual: '%s' does not match expected: '%s'", err, ErrInvalidEnvironment)
  }
}

func TestEnvironmentDefault(t *testing.T) {
  env, err := getEnvironment()
  if err != nil {
    t.Fatalf("recieved and error when expected to be nil: %s", err)
  }

  if env != Production {
    t.Errorf("actual: '%s' does not match expected: '%s'", env, Production)
  }
}
