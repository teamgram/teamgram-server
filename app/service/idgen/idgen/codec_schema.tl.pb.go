/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teagramio (teagram.io@gmail.com)
 */

// ConstructorList
// RequestList

package idgen

import (
	"fmt"

	"github.com/teamgram/proto/mtproto"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/types"
)

//////////////////////////////////////////////////////////////////////////////////////////
var _ *types.Int32Value
var _ *mtproto.Bool
var _ fmt.GoStringer

var clazzIdRegisters2 = map[int32]func() mtproto.TLObject{
	// Constructor

	// Method
	-1099886560: func() mtproto.TLObject { // 0xbe711020
		return &TLIdgenNextId{
			Constructor: -1099886560,
		}
	},
	1204121518: func() mtproto.TLObject { // 0x47c56fae
		return &TLIdgenNextIds{
			Constructor: 1204121518,
		}
	},
	-1654936704: func() mtproto.TLObject { // 0x9d5bab80
		return &TLIdgenGetCurrentSeqId{
			Constructor: -1654936704,
		}
	},
	-852747923: func() mtproto.TLObject { // 0xcd2c196d
		return &TLIdgenSetCurrentSeqId{
			Constructor: -852747923,
		}
	},
	-160339608: func() mtproto.TLObject { // 0xf6716968
		return &TLIdgenGetNextSeqId{
			Constructor: -160339608,
		}
	},
	-1479226258: func() mtproto.TLObject { // 0xa7d4cc6e
		return &TLIdgenGetNextNSeqId{
			Constructor: -1479226258,
		}
	},
}

func NewTLObjectByClassID(classId int32) mtproto.TLObject {
	f, ok := clazzIdRegisters2[classId]
	if !ok {
		return nil
	}
	return f()
}

func CheckClassID(classId int32) (ok bool) {
	_, ok = clazzIdRegisters2[classId]
	return
}

//----------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------
// TLIdgenNextId
///////////////////////////////////////////////////////////////////////////////
func (m *TLIdgenNextId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_idgen_nextId))

	switch uint32(m.Constructor) {
	case 0xbe711020:
		// idgen.nextId = Int64;
		x.UInt(0xbe711020)

		// no flags

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLIdgenNextId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLIdgenNextId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xbe711020:
		// idgen.nextId = Int64;

		// not has flags

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLIdgenNextId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLIdgenNextIds
///////////////////////////////////////////////////////////////////////////////
func (m *TLIdgenNextIds) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_idgen_nextIds))

	switch uint32(m.Constructor) {
	case 0x47c56fae:
		// idgen.nextIds num:int = Vector<long>;
		x.UInt(0x47c56fae)

		// no flags

		x.Int(m.GetNum())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLIdgenNextIds) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLIdgenNextIds) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x47c56fae:
		// idgen.nextIds num:int = Vector<long>;

		// not has flags

		m.Num = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLIdgenNextIds) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLIdgenGetCurrentSeqId
///////////////////////////////////////////////////////////////////////////////
func (m *TLIdgenGetCurrentSeqId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_idgen_getCurrentSeqId))

	switch uint32(m.Constructor) {
	case 0x9d5bab80:
		// idgen.getCurrentSeqId key:string = Int64;
		x.UInt(0x9d5bab80)

		// no flags

		x.String(m.GetKey())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLIdgenGetCurrentSeqId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLIdgenGetCurrentSeqId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x9d5bab80:
		// idgen.getCurrentSeqId key:string = Int64;

		// not has flags

		m.Key = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLIdgenGetCurrentSeqId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLIdgenSetCurrentSeqId
///////////////////////////////////////////////////////////////////////////////
func (m *TLIdgenSetCurrentSeqId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_idgen_setCurrentSeqId))

	switch uint32(m.Constructor) {
	case 0xcd2c196d:
		// idgen.setCurrentSeqId key:string id:long = Bool;
		x.UInt(0xcd2c196d)

		// no flags

		x.String(m.GetKey())
		x.Long(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLIdgenSetCurrentSeqId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLIdgenSetCurrentSeqId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xcd2c196d:
		// idgen.setCurrentSeqId key:string id:long = Bool;

		// not has flags

		m.Key = dBuf.String()
		m.Id = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLIdgenSetCurrentSeqId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLIdgenGetNextSeqId
///////////////////////////////////////////////////////////////////////////////
func (m *TLIdgenGetNextSeqId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_idgen_getNextSeqId))

	switch uint32(m.Constructor) {
	case 0xf6716968:
		// idgen.getNextSeqId key:string = Int64;
		x.UInt(0xf6716968)

		// no flags

		x.String(m.GetKey())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLIdgenGetNextSeqId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLIdgenGetNextSeqId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf6716968:
		// idgen.getNextSeqId key:string = Int64;

		// not has flags

		m.Key = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLIdgenGetNextSeqId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLIdgenGetNextNSeqId
///////////////////////////////////////////////////////////////////////////////
func (m *TLIdgenGetNextNSeqId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_idgen_getNextNSeqId))

	switch uint32(m.Constructor) {
	case 0xa7d4cc6e:
		// idgen.getNextNSeqId key:string n:int = Int64;
		x.UInt(0xa7d4cc6e)

		// no flags

		x.String(m.GetKey())
		x.Int(m.GetN())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLIdgenGetNextNSeqId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLIdgenGetNextNSeqId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xa7d4cc6e:
		// idgen.getNextNSeqId key:string n:int = Int64;

		// not has flags

		m.Key = dBuf.String()
		m.N = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLIdgenGetNextNSeqId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// Vector_Long
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_Long) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.VectorLong(m.Datas)

	return x.GetBuf()
}

func (m *Vector_Long) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Datas = dBuf.VectorLong()

	return dBuf.GetError()
}

func (m *Vector_Long) CalcByteSize(layer int32) int {
	return 0
}

func (m *Vector_Long) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}
