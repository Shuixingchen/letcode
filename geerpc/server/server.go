package server

import (
	"encoding/json"
	"fmt"
	"io"
	"letcode/geerpc/codec"
	"log"
	"net"
	"reflect"
	"sync"
)

type Serve struct {
	Option
	sending sync.Mutex //回复客户端过程需要一个一个发
}
type Option struct {
	Type codec.Type //使用哪种编解码"gob/json""
}

var DefaultOption = &Option{
	Type: codec.GobType,
}

type request struct {
	h            *codec.Header // header of request
	argv, replyv reflect.Value // argv and replyv of request
}

var DefaultServer = NewServer(Option{Type: "application/gob"})

func NewServer(option Option) *Serve {
	return &Serve{
		Option: option,
	}
}

func Accept(l net.Listener) {
	DefaultServer.Accept(l)
}

func (s *Serve) Accept(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		s.ServeConn(conn)
	}
}
func (s *Serve) ServeConn(conn io.ReadWriteCloser) {
	defer func() { _ = conn.Close() }()
	var opt Option
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		log.Println("rpc server: options error: ", err)
		return
	}

	f := codec.NewCodecFuncMap[opt.Type]
	if f == nil {
		log.Printf("rpc server: invalid codec type %s", opt.Type)
		return
	}
	s.serveCodec(f(conn))
}

var invalidRequest = struct{}{}

func (s *Serve) serveCodec(c codec.Codec) {
	var sending sync.Mutex
	var wg sync.WaitGroup
	for {
		req, err := s.readRequest(c) //循环读取客户端发送的请求
		if err != nil {
			if req == nil { //表示这个连接已经关闭了，我们也要关闭它
				break
			}
			s.sendResponse(c, req.h, invalidRequest, &sending)
			continue
		}
		wg.Add(1)
		go s.handleRequest(req, c, &sending, &wg) //每个请求都新建一个协程处理
	}
	wg.Wait()
	c.Close()
}

func (s *Serve) readRequest(cc codec.Codec) (*request, error) {
	h, err := s.readRequestHeader(cc)
	if err != nil {
		return &request{}, err
	}
	req := &request{h: h}
	// TODO: now we don't know the type of request argv
	// day 1, just suppose it's string
	req.argv = reflect.New(reflect.TypeOf(""))
	if err = cc.ReadBody(req.argv.Interface()); err != nil {
		log.Println("rpc server: read argv err:", err)
	}
	return req, nil
}

func (s *Serve) readRequestHeader(cc codec.Codec) (*codec.Header, error) {
	var header codec.Header
	if err := cc.ReadHeader(&header); err != nil {
		if err != io.EOF || err != io.ErrUnexpectedEOF {
			log.Println("rpc server: read header error:", err)
		}
		return nil, err
	}
	return &header, nil
}

func (s *Serve) handleRequest(req *request, cc codec.Codec, sending *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println(req.h, req.argv.Elem())
	req.replyv = reflect.ValueOf(fmt.Sprintf("geerpc resp %d", req.h.Seq))
	s.sendResponse(cc, req.h, req.replyv.Interface(), sending)
}

func (s *Serve) sendResponse(cc codec.Codec, h *codec.Header, body interface{}, sending *sync.Mutex) {
	s.sending.Lock()
	defer s.sending.Unlock()
	if err := cc.Write(h, body); err != nil {
		log.Println("rpc server: read argv err:", err)
	}
}
