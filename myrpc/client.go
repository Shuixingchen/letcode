package myrpc

import (
"fmt"
"letcode/myrpc/common"
"log"
"net/rpc"
)

func ClientRun () {
	var serverAddress = "localhost"
	client, err := rpc.DialHTTP("tcp", serverAddress + ":8080")
	if err != nil {
		log.Fatal("建立与服务端连接失败:", err)
	}

	args := &common.Args{10}
	var reply common.Reply
	err = client.Call("UserService.Info", args, &reply)
	if err != nil {
		log.Fatal("调用远程方法 MathService.Multiply 失败:", err)
	}
	fmt.Println(args,reply)
}
