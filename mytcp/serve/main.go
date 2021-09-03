package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	serve, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("开启socket失败")
	}
	fmt.Print("正在监听端口8080")
	for {
		conn, err := serve.Accept()
		if err != nil {
			fmt.Print("连接出错")
		}
		go handlerConn(conn)
	}
}

func handlerConn(conn net.Conn) {
	buf := make([]byte, 16)
	for {
		n, err := conn.Read(buf)
		if n == 0 || err != nil {
			conn.Close()
			break
		}
		fmt.Print("buf:", buf)
		res := binary.BigEndian.Uint16(buf)
		fmt.Println("buftoint:", res)
	}
	fmt.Printf("来自 %v 的连接关闭\n", conn.RemoteAddr())
}
