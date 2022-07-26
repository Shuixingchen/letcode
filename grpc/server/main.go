package main

import (
	"letcode/grpc/protos"
	"letcode/grpc/server/handler"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	grpcServer := grpc.NewServer()

	healthcheck := health.NewServer()
	hello := handler.NewHellowService()

	// 注册health服务，直接使用官方库生成的proto和handler
	healthpb.RegisterHealthServer(grpcServer, healthcheck)
	// 注册我们自己的hello服务，
	protos.RegisterHelloServiceServer(grpcServer, hello)
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
