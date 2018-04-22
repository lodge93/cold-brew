// Copyright © 2018 Mark Spicer
// Made available under the MIT license.

// Package server provides a RESTful HTTP interface for the cold brew coffee
// dripper.
package server

import (
	"log"

	"github.com/lodge93/cold-brew/dripper"
)

// Server is a base object which provide HTTP requests access to the dripper.
type Server struct {
	Dripper *dripper.Dripper
}

// New creates a new server instance.
func New() *Server {
	d, err := dripper.New(dripper.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}

	s := Server{
		Dripper: d,
	}

	return &s
}
