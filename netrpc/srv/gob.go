package main

import (
	"fmt"
	"letcode/netrpc/common"
	"log"
	"net"
	"net/rpc"
)

//默认使用gob编辑码

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

func GobRpcServer() {
	rpc.RegisterName("HelloService", new(HelloService))
	listener, _ := net.Listen("tcp", ":1234")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
func main() {
	HttpRpcServer()
}
