package main

import (
	_ "embed"
	"fmt"
	"letcode/goweb/db"
	"log"
	"net"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/skratchdot/open-golang/open"
)

// 这里不能有空格
//go:embed templates/transaction.html
var content []byte
var DB db.DBInterface

func main() {
	l, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	r := SetRouter()

	open.RunWith("http://localhost:3000/transaction", "chrome")
	// Start the blocking server loop
	http.Serve(l, r)
}

func SetRouter() *gin.Engine {
	r := gin.Default()
	// 显示发起交易页面
	r.GET("/transaction", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", content)
	})
	r.POST("/transaction", func(c *gin.Context) {
		DB.Set([]byte("payeeAddr"), []byte(c.PostForm("payeeAddr")))
		DB.Set([]byte("amount"), []byte(c.PostForm("amount")))
		DB.Set([]byte("pubKey"), []byte(c.PostForm("pubKey")))
		DB.Set([]byte("prvKey"), []byte(c.PostForm("prvKey")))

		c.JSON(http.StatusOK, "ooook")
	})
	return r
}

func Init() {
	DB = db.NewBadgerDB()
	DB.Open("./db", true)
}

func Exec(name string, arg []string) {
	// 判断命令是否存在
	_, err := exec.LookPath(name)
	if err != nil {
		fmt.Printf("command $s not found", name)
		return
	}
	cmd := exec.Command(name, arg...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("command exec error %v", err)
		return
	}
	fmt.Printf("%v", output)
}
