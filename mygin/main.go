package main

import (
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
	r.GET("/", Hello1)
	versionRoute := r.Group("/v1")
	versionRoute.GET("/token/:addr", Hello)
	versionRoute.GET("/account/txns_all/data", Hello1)
	versionRoute.GET("/token/:token_addr/acct/:acct_addr", Hello3)
	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}
