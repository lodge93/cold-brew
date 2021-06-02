// Copyright © 2018 Mark Spicer
// Made available under the MIT license.

// Package server provides a RESTful HTTP interface for the cold brew coffee
// dripper.
package server

import (
	"log"

	"github.com/betterengineering/cold-brew/pkg/dripper"
	scribble "github.com/nanobox-io/golang-scribble"
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

	db, err := NewDatabase(config.DatabaseDir)
	if err != nil {
		log.Fatal(err)
	}

	s := Server{
		Config: config,
		DB:     db,
	}

	settings := s.readSettingsOrDefault()

	d, err := dripper.New(settings)
	if err != nil {
		if config.Environment == EnvDevelopment {
			log.Println(err)
		} else {
			log.Fatal(err)
		}
	}

	s.Dripper = d

	return &s
}
