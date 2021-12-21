package codec

import "io"

type Header struct {
	ServiceMethod string
	Seq           int64
	Err           string
}

//编码器接口，方便使用不同的编码器
type Codec interface {
	ReadHeader(*Header) error         //读取请求头
	ReadBody(interface{}) error       //读取请求体
	Write(*Header, interface{}) error //向对方发送请求
	Close() error
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "gob"
	JsonType Type = "json" // not implemented
)

//保存已经实现了的Codex
var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
