package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		log.Fatal("create socket faild", err)
	}
	var address syscall.SockaddrInet4
	address.Port = 8080
	address.Addr = [4]byte{0, 0, 0, 0}
	err = syscall.Bind(fd, &address)
	if err != nil {
		log.Fatal("bind faild ", err)
	}
	syscall.Listen(fd, address.Port)

	for {
		nfd, addr, err := syscall.Accept(fd)
		if err != nil {
			log.Fatal("accept faild ", err)
		}
		fmt.Println("accept: ", nfd, addr, err)
		go HandConn(nfd)
	}
}
func HandConn(nfd int) {
	buf := make([]byte, 100)
	n, err := syscall.Read(nfd, buf)
	if err != nil {
		log.Fatal("read err:", err)
	}
	fmt.Println("read:", n, string(buf))

	n, err = syscall.Write(nfd, []byte("hello world"))
	if err != nil {
		log.Fatal("write faild: ", n, err)
	}
	fmt.Println("write success: ", n, err)
	return
}
