// Copyright © 2018 Mark Spicer
// Made available under the MIT license.

package main

import (
	"log"
	"time"

	"github.com/lodge93/cold-brew/dripper"
)

func main() {
	d, err := dripper.New()
	if err != nil {
		log.Fatal(err)
	}

	err = d.Bloom(5 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	err = d.Drip(60.0)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(20 * time.Second)

	d.SetDripsPerMinute(120.0)
	time.Sleep(20 * time.Second)

	err = d.Stop()
	if err != nil {
		log.Fatal(err)
	}
}
