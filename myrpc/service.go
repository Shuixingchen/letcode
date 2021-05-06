package myrpc

import (
	"letcode/myrpc/common"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type UserService struct {

}

//查询info的信息
func (us *UserService)Info(arg *common.Args, reply *common.Reply) error{
	userId := arg.UserId
	*reply = common.Reply{UserId:userId,UserName:"ssss"}
	return nil
}

func ServeRun () {
	Us := new(UserService)
	rpc.Register(Us)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("启动服务监听失败:", err)
	}
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("启动 HTTP 服务失败:", err)
	}
}