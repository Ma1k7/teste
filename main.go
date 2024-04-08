package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.Render(200, render.JSON{Data: any(gin.H{"message": "aaaaa"})})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
