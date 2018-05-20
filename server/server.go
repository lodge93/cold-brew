// Copyright © 2018 Mark Spicer
// Made available under the MIT license.

// Package server provides a RESTful HTTP interface for the cold brew coffee
// dripper.
package server

import (
	"log"

	"github.com/nanobox-io/golang-scribble"

	"github.com/lodge93/cold-brew/dripper"
)

// Server is a base object which provide HTTP requests access to the dripper.
type Server struct {
	Dripper *dripper.Dripper
	Config  *Config
	DB      *scribble.Driver
}

// New creates a new server instance.
func New() *Server {
	config, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	d, err := dripper.New(dripper.DefaultSettings())
	if err != nil {
		if config.Environment == EnvDevelopment {
			log.Println(err)
		} else {
			log.Fatal(err)
		}
	}

	db, err := NewDatabase(config.DatabaseDir)
	if err != nil {
		log.Fatal(err)
	}

	s := Server{
		Dripper: d,
		Config:  config,
		DB:      db,
	}

	return &s
}
