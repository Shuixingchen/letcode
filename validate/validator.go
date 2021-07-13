package main

import (
	"fmt"
	"github.com/go-playground/validator"
)

var validate *validator.Validate //定义

type User struct {
	Name  string `validate:"required"` //非空
	Age   int  `validate:"gte=0,lte=130"` //  0<=Age<=130
	Email string `validate:"required,email"` //非空，email格式
}

func CheckUser(sl validator.StructLevel){
	param := sl.Current().Interface().(User) //获取需要验证的对象
	if param.Name != "aa" {
		sl.ReportError(param.Name, "Name", "name", "aa", "")
	}
	if len(param.Email) == 0 {
		sl.ReportError(param.Email, "Email", "email", "required", "")
	}
}

func main() {
	user := &User{
		Name:    "ddd",
		Age:     1330,
		Email:   "1232@qq.com",
	}
	validate = validator.New()
	validate.RegisterStructValidation(CheckUser, User{}) //注册struct验证函数
	err := validate.Struct(user) //因为注册了验证函数，除了tag验证，自定义函数也会走
	if err != nil {
		for _,e:=range err.(validator.ValidationErrors){
			fmt.Println("错误字段：",e.Field())
			fmt.Println("错误值：",e.Value())
			fmt.Println("错误tag：",e.Tag())
		}
		return
	}
	fmt.Println("success")
}
