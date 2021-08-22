package main

import (
	"fmt"
	"letcode/myrpc/common"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

type SayService struct{}

func (p *HelloService) Hello(request common.Request, response *common.Response) error {
	*response = common.Response{
		UserId:   request.UserId,
		UserName: "aaaaaa",
	}
	fmt.Println("request:", request)
	return nil
}
func (p *SayService) Say(request common.Request, response *common.Response) error {
	*response = common.Response{
		UserId:   request.UserId,
		UserName: "bbbbbb",
	}
	fmt.Println("request:", request)
	return nil
}

func startRpc() {
	//注册一个name，和实例，客户端调用rpcClient.Call("HelloService.Hello", req, &reply)
	//重复的name会覆盖前面相同的实例
	serve := rpc.NewServer()
	serve.RegisterName("eth", new(SayService))
	serve.RegisterName("eth", new(HelloService))

	//监听访问，
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	//获取conn
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		rpc.ServeConn(conn)
	}

}

func main() {
	startRpc()
}
