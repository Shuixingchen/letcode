package bitcoin

import (
	"github.com/gin-gonic/gin"
	"letcode/bitcoin/web"
)

var BC *BlockChain

type TxParam struct {
	Pub string `form:"pub"`
	Pri string `form:"prv"`
	Amount int `form:"amount"`
	PayeeAddr string `form:"PayeeAddr"`
}

func Serve() {
	BC = CreateBlockChain([]byte("aaa"))
	txChan := make(chan *TxParam, 10)
	go Run(txChan)
	for {
		<-txChan
	}

}

//使用gin写一个web服务，接收用户转账，查看区块链信息
func Run(txChan chan *TxParam) {
	r := SetRouter(txChan)
	r.Run(":8080")
}

func SetRouter(txChan chan *TxParam) *gin.Engine{
	r := gin.Default()
	r.LoadHTMLGlob("templates/bitcoin/*")

	// 获取区块链数据路由
	r.GET("/blockchain", web.BlockList)
	return r
}