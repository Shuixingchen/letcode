package main

import (
	"fmt"
	"letcode/myjsonrpc/common"
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}
	var reply common.Response
	var req common.Request
	req.UserId = 11
	err = client.Call("Hello.Say", req, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
