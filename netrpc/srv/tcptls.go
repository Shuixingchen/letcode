package main

import (
	"crypto/tls"
	"log"
	"net/rpc"
)

func TcpTlsServer() {
	rpc.RegisterName("HelloService", new(HelloService))
	cert, err := tls.LoadX509KeyPair("cert/server.crt", "cert/server.key")
	if err != nil {
		panic(err)
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	listener, _ := tls.Listen("tcp", ":1234", tlsConfig)
	for {
		conn, err := listener.Accept()
		defer conn.Close()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
