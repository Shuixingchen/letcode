package mytcp

import (
	"fmt"
	"net"
	"runtime"
)


func ServeRun() {
	serve,err := net.Listen("tcp", ":8080")
	if err !=nil {
		fmt.Println("开启socket失败")
	}
	fmt.Print("正在监听端口8080")
	for true {
		conn,err := serve.Accept()
		if err != nil {
			fmt.Print("连接出错")
		}
		go handlerConn(conn)
		num := runtime.NumGoroutine()
		fmt.Println(num)
	}
}

func handlerConn(conn net.Conn) {
	buf := make([]byte, 1024)
	for{
		n,err := conn.Read(buf)
		if n == 0 || err != nil {
			conn.Close()
			break
		}
		content := string(buf)
		switch content {
		case "ping":
			conn.Write([]byte("服务器端回复-> pong\n"))
		case "hello":
			conn.Write([]byte("服务器端回复-> world\n"))
		default:
			conn.Write([]byte("服务器端回复" + content + "\n"))
		}
	}
	fmt.Printf("来自 %v 的连接关闭\n", conn.RemoteAddr())
}
