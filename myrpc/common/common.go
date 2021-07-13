package common

//rps客户端的请求参数

type Request struct {
	UserId int
}

type Response struct {
	UserId int
	UserName string
}