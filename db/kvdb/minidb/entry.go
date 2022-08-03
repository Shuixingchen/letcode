package minidb

import (
	"encoding/binary"
	"errors"
)

const entryHeaderSize = 10

type Entry struct {
	Key       []byte
	Value     []byte
	KeySize   uint32
	ValueSize uint32
	Mark      uint16
}

const (
	SET    = 1
	DELETE = 2
)

func NewEntry(key, value []byte, mark uint16) *Entry {
	return &Entry{
		Key:       key,
		Value:     value,
		KeySize:   uint32(len(key)),
		ValueSize: uint32(len(value)),
		Mark:      mark,
	}
}

func (e *Entry) GetSize() int64 {
	return int64(entryHeaderSize + e.KeySize + e.ValueSize)
}

// 编码为二进制
func (e *Entry) Encode() []byte {
	buf := make([]byte, e.GetSize())
	binary.BigEndian.PutUint32(buf[0:4], e.KeySize)
	binary.BigEndian.PutUint32(buf[4:8], e.ValueSize)
	binary.BigEndian.PutUint16(buf[8:10], e.Mark)
	copy(buf[entryHeaderSize:entryHeaderSize+e.KeySize], e.Key)
	copy(buf[entryHeaderSize+e.KeySize:], e.Value)
	return buf
}

func DecodeHeader(buf []byte) (*Entry, error) {
	var e Entry
	if len(buf) < entryHeaderSize {
		return nil, errors.New("buf error")
	}
	e.KeySize = binary.BigEndian.Uint32(buf[0:4])
	e.ValueSize = binary.BigEndian.Uint32(buf[4:8])
	e.Mark = binary.BigEndian.Uint16(buf[8:10])
	return &e, nil
}
