// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbe

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TxOutputLite struct {
	_tab flatbuffers.Table
}

func GetRootAsTxOutputLite(buf []byte, offset flatbuffers.UOffsetT) *TxOutputLite {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TxOutputLite{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTxOutputLite(buf []byte, offset flatbuffers.UOffsetT) *TxOutputLite {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &TxOutputLite{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *TxOutputLite) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TxOutputLite) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TxOutputLite) Value() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TxOutputLite) MutateValue(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *TxOutputLite) Addresses(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *TxOutputLite) AddressesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *TxOutputLite) ScriptBin() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func TxOutputLiteStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func TxOutputLiteAddValue(builder *flatbuffers.Builder, value int64) {
	builder.PrependInt64Slot(0, value, 0)
}
func TxOutputLiteAddAddresses(builder *flatbuffers.Builder, addresses flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(addresses), 0)
}
func TxOutputLiteStartAddressesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TxOutputLiteAddScriptBin(builder *flatbuffers.Builder, scriptBin flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(scriptBin), 0)
}
func TxOutputLiteEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
