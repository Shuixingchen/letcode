// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbe

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Tx struct {
	_tab flatbuffers.Table
}

func GetRootAsTx(buf []byte, offset flatbuffers.UOffsetT) *Tx {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Tx{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTx(buf []byte, offset flatbuffers.UOffsetT) *Tx {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Tx{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Tx) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Tx) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Tx) BlockHeight() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Tx) MutateBlockHeight(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *Tx) BlockTime() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Tx) MutateBlockTime(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *Tx) IsCoinbase() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Tx) MutateIsCoinbase(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

func (rcv *Tx) Version() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Tx) MutateVersion(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

func (rcv *Tx) LockTime() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Tx) MutateLockTime(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

func (rcv *Tx) Size() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Tx) MutateSize(n int32) bool {
	return rcv._tab.MutateInt32Slot(14, n)
}

func (rcv *Tx) Sigops() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Tx) MutateSigops(n int64) bool {
	return rcv._tab.MutateInt64Slot(16, n)
}

func (rcv *Tx) Fee() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Tx) MutateFee(n int64) bool {
	return rcv._tab.MutateInt64Slot(18, n)
}

func (rcv *Tx) Inputs(obj *TxInput, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Tx) InputsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Tx) InputsCount() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Tx) MutateInputsCount(n int32) bool {
	return rcv._tab.MutateInt32Slot(22, n)
}

func (rcv *Tx) InputsValue() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Tx) MutateInputsValue(n int64) bool {
	return rcv._tab.MutateInt64Slot(24, n)
}

func (rcv *Tx) Outputs(obj *TxOutput, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Tx) OutputsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Tx) OutputsCount() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Tx) MutateOutputsCount(n int32) bool {
	return rcv._tab.MutateInt32Slot(28, n)
}

func (rcv *Tx) OutputsValue() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Tx) MutateOutputsValue(n int64) bool {
	return rcv._tab.MutateInt64Slot(30, n)
}

func (rcv *Tx) CreatedAt() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Tx) MutateCreatedAt(n uint32) bool {
	return rcv._tab.MutateUint32Slot(32, n)
}

func (rcv *Tx) IsDoubleSpend() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Tx) MutateIsDoubleSpend(n bool) bool {
	return rcv._tab.MutateBoolSlot(34, n)
}

func (rcv *Tx) IsSwTx() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Tx) MutateIsSwTx(n bool) bool {
	return rcv._tab.MutateBoolSlot(36, n)
}

func (rcv *Tx) Weight() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Tx) MutateWeight(n int32) bool {
	return rcv._tab.MutateInt32Slot(38, n)
}

func (rcv *Tx) Vsize() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Tx) MutateVsize(n int32) bool {
	return rcv._tab.MutateInt32Slot(40, n)
}

func (rcv *Tx) WitnessHash() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func TxStart(builder *flatbuffers.Builder) {
	builder.StartObject(20)
}
func TxAddBlockHeight(builder *flatbuffers.Builder, blockHeight int32) {
	builder.PrependInt32Slot(0, blockHeight, 0)
}
func TxAddBlockTime(builder *flatbuffers.Builder, blockTime uint32) {
	builder.PrependUint32Slot(1, blockTime, 0)
}
func TxAddIsCoinbase(builder *flatbuffers.Builder, isCoinbase bool) {
	builder.PrependBoolSlot(2, isCoinbase, false)
}
func TxAddVersion(builder *flatbuffers.Builder, version int32) {
	builder.PrependInt32Slot(3, version, 0)
}
func TxAddLockTime(builder *flatbuffers.Builder, lockTime uint32) {
	builder.PrependUint32Slot(4, lockTime, 0)
}
func TxAddSize(builder *flatbuffers.Builder, size int32) {
	builder.PrependInt32Slot(5, size, 0)
}
func TxAddSigops(builder *flatbuffers.Builder, sigops int64) {
	builder.PrependInt64Slot(6, sigops, 0)
}
func TxAddFee(builder *flatbuffers.Builder, fee int64) {
	builder.PrependInt64Slot(7, fee, 0)
}
func TxAddInputs(builder *flatbuffers.Builder, inputs flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(inputs), 0)
}
func TxStartInputsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TxAddInputsCount(builder *flatbuffers.Builder, inputsCount int32) {
	builder.PrependInt32Slot(9, inputsCount, 0)
}
func TxAddInputsValue(builder *flatbuffers.Builder, inputsValue int64) {
	builder.PrependInt64Slot(10, inputsValue, 0)
}
func TxAddOutputs(builder *flatbuffers.Builder, outputs flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(outputs), 0)
}
func TxStartOutputsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TxAddOutputsCount(builder *flatbuffers.Builder, outputsCount int32) {
	builder.PrependInt32Slot(12, outputsCount, 0)
}
func TxAddOutputsValue(builder *flatbuffers.Builder, outputsValue int64) {
	builder.PrependInt64Slot(13, outputsValue, 0)
}
func TxAddCreatedAt(builder *flatbuffers.Builder, createdAt uint32) {
	builder.PrependUint32Slot(14, createdAt, 0)
}
func TxAddIsDoubleSpend(builder *flatbuffers.Builder, isDoubleSpend bool) {
	builder.PrependBoolSlot(15, isDoubleSpend, false)
}
func TxAddIsSwTx(builder *flatbuffers.Builder, isSwTx bool) {
	builder.PrependBoolSlot(16, isSwTx, false)
}
func TxAddWeight(builder *flatbuffers.Builder, weight int32) {
	builder.PrependInt32Slot(17, weight, 0)
}
func TxAddVsize(builder *flatbuffers.Builder, vsize int32) {
	builder.PrependInt32Slot(18, vsize, 0)
}
func TxAddWitnessHash(builder *flatbuffers.Builder, witnessHash flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(19, flatbuffers.UOffsetT(witnessHash), 0)
}
func TxEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
