package common

//rps客户端的请求参数

type Args struct {
	UserId int
}

type Reply struct {
	UserId int
	UserName string
}