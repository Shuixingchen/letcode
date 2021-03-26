package bitcoin

import (
	"github.com/gin-gonic/gin"
)

type TxParam struct {
	Pub string `form:"pub"`
	Pri string `form:"prv"`
	Amount int `form:"amount"`
	PayeeAddr string `form:"PayeeAddr"`
}

func Serve() {
	bc := CreateBlockChain([]byte("aaa"))
	txChan := make(chan *TxParam, 10)
	go Run(bc,txChan)
	for {
		<-txChan
	}

}

//使用gin写一个web服务，接收用户转账，查看区块链信息
func Run(bc *BlockChain, txChan chan *TxParam) {
	r := SetRouter(bc, txChan)
	r.Run(":8080")
}

func SetRouter(bc *BlockChain, txChan chan *TxParam) *gin.Engine{
	r := gin.Default()
	// 获取用户数据路由
	r.GET("/blockchain", func(c *gin.Context) {
		bc.Print()
		c.String(200,res)
	})
	return r
}