package serve

import (
	"fmt"
	"log"
	"net"
)

//管理连接
type Serve struct {
	Addr string
}

func NewServe(addr string) {
	s := Serve{Addr: addr}
	s.Listen()
}

func (s *Serve) Listen() {
	fmt.Println("start listen", s.Addr)
	l, err := net.Listen("tcp", s.Addr)
	if err != nil {
		log.Fatal("listen faild ", err)
	}
	for {
		rw, err := l.Accept()
		if err != nil {
			log.Fatal("accept faild ", err)
		}
		mConn := NewConn(&rw)
		go mConn.Serve()
	}
}
