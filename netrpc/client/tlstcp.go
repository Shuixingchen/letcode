package main

import (
	"crypto/tls"
	"fmt"
	"letcode/netrpc/common"
	"net/rpc"
)

func TlsTcpClient() {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, //如果客户端不需要对服务端鉴权，可以设置true
	}
	tlsconn, _ := tls.Dial("tcp", ":1234", tlsConfig)
	rpcClient := rpc.NewClient(tlsconn)

	var reply common.Response
	var req common.Request
	req.UserId = 11
	err := rpcClient.Call("HelloService.Hello", req, &reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(reply)
}
