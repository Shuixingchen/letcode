package serve

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"strings"
)

type Conn struct {
	Conn *net.Conn
	bufr *bufio.Reader
	bufw *bufio.Writer
	Ser  *Serve
}

type Request struct {
	Host     string
	Method   string
	URI      string
	Protocol string
	Body     string
	Header   map[string]string
	Ctx      context.Context //response绑定了calcleFunc
}

func NewConn(c *net.Conn) *Conn {
	return &Conn{
		Conn: c,
	}
}

//读取requet,实例化response交给ServerHandler
func (c *Conn) Serve(ctx context.Context) {
	c.bufr = bufio.NewReader(*c.Conn)
	c.bufw = bufio.NewWriterSize(*c.Conn, 4<<10)
	ctx, cancelCtx := context.WithCancel(ctx)
	request := c.ReadRequest()
	response := Response{Conn: c, HandlerHeader: make(map[string]string, 0), ContentLength: 0, CancelCtx: cancelCtx}
	ServerHandler{srv: c.Ser}.ServeHTTP(response, request)
	response.CancelCtx() //response执行完毕
	fmt.Print("finish serve")
	return
}

//组装request
func (c *Conn) ReadRequest() *Request {
	//read first line
	s := ReadLine(c.bufr)
	firstLine := strings.Split(s, " ")
	request := new(Request)
	request.Method = firstLine[0]
	request.URI = firstLine[1]
	request.Protocol = firstLine[2]

	//readHeader
	request.Header = ReadHeader(c.bufr)

	//read body
	//request.Body = &body{src: io.LimitReader(r, realLength), closing: t.Close}
	return request
}

func ReadLine(r *bufio.Reader) string {
	var lines []byte
	for {
		line, isPrefix, err := r.ReadLine() //The text returned does not include the line end ("\r\n" or "\n")
		if err != nil {
			log.Fatal("read line err", err)
		}
		lines = append(lines, line...)
		if isPrefix == false {
			break
		}
	}
	return string(lines)
}

//读取到\r\n\r\n为止
func ReadHeader(r *bufio.Reader) map[string]string {
	r.Peek(1)
	header := make(map[string]string)
	for {
		line, _, err := r.ReadLine() //以\r\n读取
		if len(line) == 0 {          //最后一行是空行
			break
		}
		r.Peek(2)
		if err != nil {
			log.Fatal("read line err", err)
		}

		i := bytes.IndexByte(line, ':')
		if i < 0 {
			continue
		}
		key := string(line[:i])
		value := string(line[i+1:])
		header[key] = value
	}
	return header
}
