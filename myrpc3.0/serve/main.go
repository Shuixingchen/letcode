package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"letcode/proto"
)

type GreeterHandler struct {
}
func(handler *GreeterHandler)Hello(ctx context.Context, req *proto.HelloRequest, resp *proto.HelloResponse)error{
	resp.Greeting = "ok"
	user1 := &proto.User{
		Id:                   1,
		Name:                 "aa",
	}
	user2 := &proto.User{
		Id:                   2,
		Name:                 "bb",
	}

	resp.Users = []*proto.User{user1,user2}
	return nil
}

func main(){
	//create service
	serve := micro.NewService(micro.Name("Greeter"))

	serve.Init()

	//register handler
	proto.RegisterGreeterHandler(serve.Server(), new(GreeterHandler))

	//runserve
	if err := serve.Run();err!=nil{
		fmt.Println(err)
	}

}