package main

import (
	"log"
	"net/http"
)

func main() {
	// serve.NewServe(":8080")
	defaultServe()
}

//sayHelloWorld需要实现http.Handler接口
func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values.Get("id")
	w.Write([]byte("id" + id))
}

func defaultServe() {
	http.HandleFunc("/", sayHelloWorld)      //路由处理
	err := http.ListenAndServe(":8080", nil) //监听9091端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
