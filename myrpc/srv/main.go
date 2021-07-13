package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/util/log"
	"letcode/myrpc/common"
	"net"
	"net/rpc"
)

type HelloService struct {}

func (p *HelloService) Hello(request common.Request, response *common.Response) error {
	*response = common.Response{
		UserId: request.UserId,
		UserName: "lalalal",
	}
	fmt.Println("request:", request)
	return nil
}

func startRpc(){
	//注册服务
	rpc.RegisterName("HelloService", new(HelloService))
	//监听访问，
	listener, err := net.Listen("ws", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}

func startRpcHttp(){
	rpc.RegisterName("HelloService", new(HelloService))
	//listen,err := net.Listen("")
}

func main() {
	startRpc()
}