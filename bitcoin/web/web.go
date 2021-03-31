package web

import (
	"github.com/gin-gonic/gin"
)

type BlockShow struct {
	ID int
	Name string
}

func BlockList(ctx *gin.Context) {
	//ctx.HTML(http.StatusOK, "blockchain.html")
}