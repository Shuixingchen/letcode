package main

import (
	"log"
	"net/http"
)

//sayHelloWorld需要实现http.Handler接口
func sayHelloWorld(w http.ResponseWriter, r *http.Request)  {
	values := r.URL.Query()
	id :=values.Get("id")
	w.Write([]byte("id"+id))
}

func main()  {
	http.HandleFunc("/", sayHelloWorld)      //路由处理
	err := http.ListenAndServe(":9091", nil) //监听9091端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
