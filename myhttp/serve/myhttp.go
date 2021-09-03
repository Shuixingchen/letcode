package serve

import (
	"fmt"
	"log"
	"net"
)

//自定义httpserver

func HttpStart() {
	fmt.Println("start listen 8080")
	serve, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen 8080 faild", err)
	}
	for {
		conn, err := serve.Accept()
		if err != nil {
			log.Fatal("accept faild", err)
		}
		go HandConn(conn)
	}
}

func HandConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1000)
	for {
		n, err := conn.Read(buf)
		if err != nil || n == 0 {
			log.Fatal("n,err", n, err)
			return
		}

	}
}
