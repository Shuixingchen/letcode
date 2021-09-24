package main

import (
	"fmt"
	"letcode/myhttp/serve"
	"log"
	"net/http"
	"time"
)

func myHandler(w serve.Response, r *serve.Request) {
	// fmt.Print(r)
	//向response写入
	_, err := w.Write([]byte("id"))
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Print("finish response")
}

func main() {
	// s := serve.NewServe(":8080")
	// s.AddFunc("/", myHandler)
	// s.Listen()

	defaultServe()
}

//sayHelloWorld需要实现http.Handler接口
func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	//直接从url获取参数
	values := r.URL.Query()
	id := values.Get("id")

	//从form表单获取数据
	r.ParseForm()
	_ = r.PostForm["user"] //post提交的数据

	//从json字符串获取
	body := make([]byte, r.ContentLength) //读取request的body
	r.Body.Read(body)

	//向response写入
	w.Write([]byte("id" + id))
}

func defaultServe() {
	http.HandleFunc("/", sayHelloWorld) //路由处理

	server := http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
