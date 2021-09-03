package serve

import (
	"bufio"
	"fmt"
	"net"
)

type Conn struct {
	Conn *net.Conn
	bufr *bufio.Reader
	bufw *bufio.Writer
}

type Request struct {
}
type Response struct {
}

func NewConn(c *net.Conn) *Conn {
	return &Conn{
		Conn: c,
	}
}

//读取封装数据
func (c *Conn) Serve() {
	c.bufr = bufio.NewReader(*c.Conn)
	c.bufw = bufio.NewWriterSize(*c.Conn, 4<<10)
	c.ReadRequest()
}
func (c *Conn) ReadRequest() (line []byte, more bool, err error) {
	buf := make([]byte, 1<<10)
	c.bufr.Read(buf)
	fmt.Println(string(buf))
	return
}
