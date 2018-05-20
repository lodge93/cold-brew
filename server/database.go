// Copyright © 2018 Mark Spicer
// Made available under the MIT license.

package server

import (
	scribble "github.com/nanobox-io/golang-scribble"
)

// NewDatabase creates a new filesystem based database given the base directory.
func NewDatabase(dir string) (*scribble.Driver, error) {
	db, err := scribble.New(dir, nil)
	if err != nil {
		return nil, err
	}

	return db, nil
}
