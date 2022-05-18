// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbe

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BlockTxsHash struct {
	_tab flatbuffers.Table
}

func GetRootAsBlockTxsHash(buf []byte, offset flatbuffers.UOffsetT) *BlockTxsHash {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BlockTxsHash{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsBlockTxsHash(buf []byte, offset flatbuffers.UOffsetT) *BlockTxsHash {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &BlockTxsHash{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *BlockTxsHash) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BlockTxsHash) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BlockTxsHash) HashStr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func BlockTxsHashStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func BlockTxsHashAddHashStr(builder *flatbuffers.Builder, hashStr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(hashStr), 0)
}
func BlockTxsHashEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
