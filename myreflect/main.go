package main

import (
	"fmt"
	"letcode/myreflect/handle"
	"letcode/myreflect/services"
)

func Register() {
	var user services.User
	// 注册user服务
	handle.RegisterService(&user)
	// 调用user服务
	handle.CallServerMethod("Say", []string{"hello"})
}

func main() {
	u1 := new(services.User)
	u2 := new(services.User)
	u1.Id = 11
	u1.SetUse("aa")
	fmt.Println(u1.GetUse(), u1.Id)
	handle.CopyData(u2, u1)
	fmt.Println(u2.GetUse(), u2.Id)
}
