package main

import (
	"letcode/grpc/protos"
	"letcode/grpc/server/handler"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	protos.RegisterHelloServiceServer(grpcServer, new(handler.HelloService))
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
