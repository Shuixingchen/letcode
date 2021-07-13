package serve

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"letcode/proto/spider"
	"net"
	"net/http"
)
const (
	// Address 监听地址
	Address string = "localhost:8080"
	// Method 通信方法
	Method string = "tcp"
)

type server struct {}
func (s *server) GetAddressResponse(ctx context.Context, a *spider.SendAddress) (*spider.GetResponse, error) {
	// 逻辑写在这里
	switch a.Method {
	case "get", "Get", "GET":
		// 演示微服务用,故只写get示例
		status, body, err := get(a.Address)
		if err != nil {
			return nil, err
		}
		res := spider.GetResponse{
			HttpCode: int32(status),
			Response: body,
		}
		return &res, nil
	}
	return nil, nil
}

func get(address string) (s int, r string, err error) {
	// get请求
	resp, err := http.Get(address)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	s = resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	r = string(body)
	return
}


func main(){
	// 监听本地端口
	listener, err := net.Listen(Method, Address)
	if err != nil {
		return
	}
	s := grpc.NewServer() // 创建GRPC
	spider.RegisterGoSpiderServer(s, &server{})
	reflection.Register(s) // 在GRPC服务器注册服务器反射服务
	err = s.Serve(listener)
	if err != nil {
		return
	}
}