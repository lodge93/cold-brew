// Copyright © 2018 Mark Spicer
// Made available under the MIT license.

// Cold Brew is a project to control a kyoto cold brew tower via software.
package main

import (
	"github.com/betterengineering/cold-brew/internal/server"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	s := server.New()
	defer s.Dripper.Off()

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./assets/dist", true)))
	r.GET("/api/cold-brew/v1/dripper", s.GetDripper)
	r.GET("/api/cold-brew/v1/dripper/settings", s.GetDripperSettings)
	r.POST("/api/cold-brew/v1/dripper/settings", s.SetDripperSettings)
	r.POST("/api/cold-brew/v1/dripper/run", s.SetDripperRun)
	r.POST("/api/cold-brew/v1/dripper/off", s.SetDripperOff)
	r.POST("/api/cold-brew/v1/dripper/drip", s.SetDripperDrip)
	r.Run()
}
