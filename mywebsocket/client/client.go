package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {

}

func MyClient() {
	conn, err := net.Dial("ws", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("connect fail", err)
	}
	conn.Write([]byte("GET /chat HTTP/1.1\r\n"))
	conn.Write([]byte("Host: server.example.com\r\n"))
	conn.Write([]byte("Upgrade: websocket\r\n"))
	conn.Write([]byte("Sec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==\r\n"))
	conn.Write([]byte("Origin: http://example.com\r\n"))
	conn.Write([]byte("Sec-WebSocket-Protocol: chat, superchat\r\n"))
	n, err := conn.Write([]byte("Sec-WebSocket-Version: 13\r\n"))
	fmt.Println("write:", n, err)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			buf := make([]byte, 1000)
			n, err := conn.Read(buf)
			fmt.Println("n,err", n, err)
			if n == 0 || err == nil {
				break
			}
			fmt.Println("receive from serve:", string(buf))
		}
	}()
	wg.Wait()
}
