// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbe

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BlockLite struct {
	_tab flatbuffers.Table
}

func GetRootAsBlockLite(buf []byte, offset flatbuffers.UOffsetT) *BlockLite {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BlockLite{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsBlockLite(buf []byte, offset flatbuffers.UOffsetT) *BlockLite {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &BlockLite{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *BlockLite) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BlockLite) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BlockLite) Height() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BlockLite) MutateHeight(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *BlockLite) Timestamp() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BlockLite) MutateTimestamp(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *BlockLite) Hash() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *BlockLite) PrevHash() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func BlockLiteStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func BlockLiteAddHeight(builder *flatbuffers.Builder, height int32) {
	builder.PrependInt32Slot(0, height, 0)
}
func BlockLiteAddTimestamp(builder *flatbuffers.Builder, timestamp uint32) {
	builder.PrependUint32Slot(1, timestamp, 0)
}
func BlockLiteAddHash(builder *flatbuffers.Builder, hash flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(hash), 0)
}
func BlockLiteAddPrevHash(builder *flatbuffers.Builder, prevHash flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(prevHash), 0)
}
func BlockLiteEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
