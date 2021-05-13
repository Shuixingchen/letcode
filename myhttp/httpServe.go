package myhttp

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "你好，学院君！")  // 发送响应到客户端
}

func main()  {
	http.HandleFunc("/", sayHelloWorld) //路由处理
	err := http.ListenAndServe(":9091", nil)//监听9091端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
