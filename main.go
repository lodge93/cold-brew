// Copyright © 2018 Mark Spicer
// Made available under the MIT license.

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lodge93/cold-brew/server"
)

func main() {
	s := server.New()
	defer s.Dripper.Off()

	r := gin.Default()
	r.GET("/api/cold-brew/v1/dripper", s.GetDripper)
	r.POST("/api/cold-brew/v1/dripper/run", s.SetDripperRun)
	r.POST("/api/cold-brew/v1/dripper/off", s.SetDripperOff)
	r.POST("/api/cold-brew/v1/dripper/drip", s.SetDripperDrip)
	r.Run()
}
