package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"golang.org/x/net/context"
	"letcode/proto"
)

type GreeterServiceHandler struct{}

func main(){
	//get serve
	service := micro.NewService(
		micro.Name("Greeter.client"),
		)

	service.Init()

	//get greeterService
	greeter := proto.NewGreeterService("Greeter", service.Client())
	request := proto.HelloRequest{
		Name:                 "11",
	}
	response,err := greeter.Hello(context.TODO(), &request)

	if err!= nil{
		fmt.Println(err)
	}
	fmt.Println(response)
}