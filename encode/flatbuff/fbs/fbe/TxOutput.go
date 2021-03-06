// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbe

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TxOutput struct {
	_tab flatbuffers.Table
}

func GetRootAsTxOutput(buf []byte, offset flatbuffers.UOffsetT) *TxOutput {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TxOutput{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTxOutput(buf []byte, offset flatbuffers.UOffsetT) *TxOutput {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &TxOutput{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *TxOutput) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TxOutput) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TxOutput) Value() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TxOutput) MutateValue(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *TxOutput) Type() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TxOutput) MutateType(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *TxOutput) Addresses(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *TxOutput) AddressesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *TxOutput) ScriptAsm() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *TxOutput) ScriptHex() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func TxOutputStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func TxOutputAddValue(builder *flatbuffers.Builder, value int64) {
	builder.PrependInt64Slot(0, value, 0)
}
func TxOutputAddType(builder *flatbuffers.Builder, type_ int32) {
	builder.PrependInt32Slot(1, type_, 0)
}
func TxOutputAddAddresses(builder *flatbuffers.Builder, addresses flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(addresses), 0)
}
func TxOutputStartAddressesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TxOutputAddScriptAsm(builder *flatbuffers.Builder, scriptAsm flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(scriptAsm), 0)
}
func TxOutputAddScriptHex(builder *flatbuffers.Builder, scriptHex flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(scriptHex), 0)
}
func TxOutputEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
