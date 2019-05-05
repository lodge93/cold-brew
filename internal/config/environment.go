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
  "github.com/pkg/errors"
  "github.com/spf13/viper"
)

const (
  // Production is an environment type representing a deployed Cold Brew application.
  Production = "production"

  // Development represents a Cold Brew application currently being developed upon.
  Development = "development"

  environmentViperKey = "environment"
)

var (
  // ErrInvalidEnvironment is returned when an environment was supplied that Cold Brew does not support.
  ErrInvalidEnvironment = errors.New("invalid environment supplied")
)

// Environment is a type to define the currently running environment.
type Environment string

func getEnvironment() (Environment, error) {
  if viper.IsSet(environmentViperKey) {
    return parseEnvironment(viper.GetString(environmentViperKey))
  }

  return Production, nil
}

func parseEnvironment(env string) (Environment, error) {
  switch env {
  case Production:
    return Production, nil
  case Development:
    return Development, nil
  default:
    return "", ErrInvalidEnvironment
  }
}
