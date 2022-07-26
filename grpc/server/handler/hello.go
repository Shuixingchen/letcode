package handler

import (
	"context"
	"letcode/grpc/protos"
)

type HelloService struct {
	protos.UnimplementedHelloServiceServer
}

func NewHellowService() *HelloService {
	return new(HelloService)
}

func (h *HelloService) Hello(ctx context.Context, req *protos.Request) (*protos.Response, error) {
	res := protos.Response{Data: "hello " + req.Value}
	return &res, nil
}
