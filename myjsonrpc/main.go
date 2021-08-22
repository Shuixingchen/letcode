package main

import (
	"fmt"
	"letcode/myjsonrpc/common"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"log"
)

type HelloService struct {
	Name string
	Age  int
}

func (h *HelloService) Say(req common.Request, res *common.Response) error {
	*res = common.Response{
		UserId:   11,
		UserName: "aa",
	}
	fmt.Println(req)
	return nil
}

func main() {
	rpc.RegisterName("Hello", new(HelloService))
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go jsonrpc.ServeConn(conn)
	}

}
