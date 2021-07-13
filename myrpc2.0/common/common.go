package common

import "net/rpc"

const UserServiceName = "rpc/UserService"

type UserServiceInterface interface {
	Hello(request Request, reply *Response) error
}

type Request struct {
	UserId int
}

type Response struct {
	UserId int
	UserName string
}

func RegisterUserService(srv UserServiceInterface) error{
	return rpc.RegisterName(UserServiceName, srv)
}