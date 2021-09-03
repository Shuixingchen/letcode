package main

import (
	"encoding/binary"
	"log"
	"net"
	"sync"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	cConnHandler(conn)
}

func cConnHandler(c net.Conn) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for {
			var buf = make([]byte, 32)
			n, err := c.Read(buf)
			if err != nil {
				log.Println("read error:", err)
				break
			} else {
				log.Printf("read % bytes, content is %s\n", n, string(buf[:n]))
			}
		}
		wg.Done()
	}()
	go func() {
		send := make([]byte, 2)
		binary.BigEndian.PutUint16(send, 300)
		c.Write(send)
		wg.Done()
	}()
	wg.Wait()
}
