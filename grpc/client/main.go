package main

import (
	"context"
	"fmt"
	"letcode/grpc/protos"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := protos.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &protos.Request{Value: "aaa"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetData())
}
