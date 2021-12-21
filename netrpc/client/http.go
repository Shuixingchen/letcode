package main

import (
	"fmt"
	"letcode/netrpc/common"
	"net/rpc"
)

func HttpRpcClient() {
	rpcClient, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var reply common.Response
	var req common.Request
	req.UserId = 11
	err = rpcClient.Call("HelloService.Hello", req, &reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(reply)
}

func HttpRpcAsync() {
	rpcClient, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var reply common.Response
	var req common.Request
	req.UserId = 11
	async := rpcClient.Go("HelloService.Hello", req, &reply, nil)
	<-async.Done
	fmt.Print(reply)
}
