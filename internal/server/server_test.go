// Copyright © 2018 Mark Spicer
// Made available under the MIT license.

package server

func withTestServerStruct() (*Server, string, error) {
	db, dir, err := withNewTempDatabase()
	if err != nil {
		return nil, "", err
	}

	config, err := NewConfig()
	if err != nil {
		return nil, "", err
	}

	s := Server{
		Config: config,
		DB:     db,
	}

	return &s, dir, nil
}
