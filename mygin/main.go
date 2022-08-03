package main

import (
	"letcode/mygin/middleware"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Hello(c *gin.Context) {
	addr := c.Param("addr")
	log.Info("addr:", addr)
}
func Hello1(c *gin.Context) {
	log.Info("addr1:")
}
func Hello3(c *gin.Context) {
	addr := c.Param("addr")
	log.Info("addr:", addr)
}

func main() {
	r := gin.Default()
	r.Use(middleware.UseOpenTracing())
	r.GET("/", Hello1)
	versionRoute := r.Group("/v1")
	versionRoute.GET("/token", Hello)

	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}
