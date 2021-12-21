package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

//使用gob实现编码器接口

type GobCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer
	dec  *gob.Decoder
	enc  *gob.Encoder
}

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	return &GobCodec{
		conn: conn,
		buf:  bufio.NewWriter(conn),
		dec:  gob.NewDecoder(conn), //得到解码器
		enc:  gob.NewEncoder(conn), //得到编码器
	}
}

func (g *GobCodec) ReadHeader(h *Header) error {
	return g.dec.Decode(h)
}

func (g *GobCodec) ReadBody(body interface{}) error {
	return g.dec.Decode(body)
}

func (g *GobCodec) Write(h *Header, body interface{}) (err error) {
	defer func() {
		g.buf.Flush()
		if err != nil {
			g.Close()
		}
	}()
	if err := g.enc.Encode(h); err != nil {
		log.Println(err)
		return err
	}
	if err := g.enc.Encode(body); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (g *GobCodec) Close() error {
	return g.conn.Close()
}
