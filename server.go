package main

import (
	"io"
	"os"

	"github.com/Omar-Belghaouti/gin-framework-course/controller"
	"github.com/Omar-Belghaouti/gin-framework-course/middlewares"
	"github.com/Omar-Belghaouti/gin-framework-course/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()

	server.Use(
		gin.Recovery(),
		middlewares.Logger(),
		middlewares.BasicAuth(),
		gindump.Dump(),
	)

	server.GET("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.Save(c))
	})

	server.Run(":8080")
}
