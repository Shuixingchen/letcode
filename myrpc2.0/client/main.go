package main

import (
	"fmt"
	"letcode/myrpc2.0/common"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	request := common.Request{
		UserId: 1111,
	}
	var response common.Response
	err = client.Call(common.UserServiceName+".Hello", request, &response)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
