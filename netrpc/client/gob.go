package main

import (
	"fmt"
	"letcode/netrpc/common"
	"net/rpc"
)

//默认使用gob编解码

func GobRpcClient() {
	rpcClient, err := rpc.Dial("tcp", "localhost:1234")
	// rpcClient, err := rpc.DialHTTP("tcp", "localhost:8888")
	if err != nil {
		fmt.Println(err)
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

func main() {
	HttpRpcClient()
}
