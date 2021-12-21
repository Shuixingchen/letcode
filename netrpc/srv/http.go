package main

import (
	"log"
	"net/http"
	"net/rpc"
)

func HttpRpcServer() {
	rpc.RegisterName("HelloService", new(HelloService)) //把服务注册到rpc.Server.serviceMap
	rpc.HandleHTTP()                                    //注册rpc默认路由
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("Error serving: ", err)
	}
}
