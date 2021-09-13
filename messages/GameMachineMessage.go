// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package messages

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type GameMachineMessage struct {
	_tab flatbuffers.Table
}

func GetRootAsGameMachineMessage(buf []byte, offset flatbuffers.UOffsetT) *GameMachineMessage {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &GameMachineMessage{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsGameMachineMessage(buf []byte, offset flatbuffers.UOffsetT) *GameMachineMessage {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &GameMachineMessage{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *GameMachineMessage) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *GameMachineMessage) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *GameMachineMessage) MType() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *GameMachineMessage) MutateMType(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *GameMachineMessage) Contents(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *GameMachineMessage) ContentsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *GameMachineMessage) ContentsBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *GameMachineMessage) MutateContents(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func GameMachineMessageStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func GameMachineMessageAddMType(builder *flatbuffers.Builder, mType int32) {
	builder.PrependInt32Slot(0, mType, 0)
}
func GameMachineMessageAddContents(builder *flatbuffers.Builder, contents flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(contents), 0)
}
func GameMachineMessageStartContentsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func GameMachineMessageEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
