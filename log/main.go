package main

import "log"

//官方自带log,log是多 goroutine 安全的

func init() {
	log.SetPrefix("TRACE: ") //日志前缀
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile) //显示的信息
}

func main() {
	log.Println("message")
	log.Fatalln("fatal message")
	log.Panicln("panic message") //会调用panic
}
