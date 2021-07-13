package main

import (
	"fmt"
	"letcode/myrpc/common"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println(err)
	}

	var reply common.Response
	var req common.Request
	req.UserId = 11
	err = client.Call("HelloService.Hello", req, &reply)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reply)
}