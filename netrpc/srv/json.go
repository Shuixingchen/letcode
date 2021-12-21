package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func jsonServer() {
	rpc.RegisterName("HelloService", new(HelloService)) //把服务注册到rpc.Server.serviceMap
	listener, _ := net.Listen("tcp", ":1234")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
