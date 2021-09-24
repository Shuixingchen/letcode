package serve

import (
	"context"
	"fmt"
	"log"
	"net"
)

type Serve struct {
	Addr string
	Mux  *ServeMux
}

//管理路由
type ServerHandler struct {
	srv *Serve
}

func (sh ServerHandler) ServeHTTP(rw Response, req *Request) {
	route := sh.srv.Mux
	doF, _ := route.Handler(req)
	doF(rw, req)
}

func NewServe(addr string) *Serve {
	return &Serve{Addr: addr, Mux: NewServeMux()}
}

func (s *Serve) AddFunc(pattern string, handler DoFunc) {
	s.Mux.Register(pattern, handler)
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
		mConn.Ser = s
		go mConn.Serve(context.Background())
	}
}
