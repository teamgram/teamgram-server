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

package banned

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
	-515891261: func() mtproto.TLObject { // 0xe1401fc3
		return &TLBannedCheckPhoneNumberBanned{
			Constructor: -515891261,
		}
	},
	-453047268: func() mtproto.TLObject { // 0xe4ff0c1c
		return &TLBannedGetBannedByPhoneList{
			Constructor: -453047268,
		}
	},
	1037800024: func() mtproto.TLObject { // 0x3ddb9258
		return &TLBannedBan{
			Constructor: 1037800024,
		}
	},
	1912613348: func() mtproto.TLObject { // 0x720029e4
		return &TLBannedUnBan{
			Constructor: 1912613348,
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
// TLBannedCheckPhoneNumberBanned
///////////////////////////////////////////////////////////////////////////////

func (m *TLBannedCheckPhoneNumberBanned) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_banned_checkPhoneNumberBanned))

	switch uint32(m.Constructor) {
	case 0xe1401fc3:
		// banned.checkPhoneNumberBanned phone:string = Bool;
		x.UInt(0xe1401fc3)

		// no flags

		x.String(m.GetPhone())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLBannedCheckPhoneNumberBanned) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLBannedCheckPhoneNumberBanned) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe1401fc3:
		// banned.checkPhoneNumberBanned phone:string = Bool;

		// not has flags

		m.Phone = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLBannedCheckPhoneNumberBanned) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLBannedGetBannedByPhoneList
///////////////////////////////////////////////////////////////////////////////

func (m *TLBannedGetBannedByPhoneList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_banned_getBannedByPhoneList))

	switch uint32(m.Constructor) {
	case 0xe4ff0c1c:
		// banned.getBannedByPhoneList phones:Vector<string> = Vector<string>;
		x.UInt(0xe4ff0c1c)

		// no flags

		x.VectorString(m.GetPhones())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLBannedGetBannedByPhoneList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLBannedGetBannedByPhoneList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe4ff0c1c:
		// banned.getBannedByPhoneList phones:Vector<string> = Vector<string>;

		// not has flags

		m.Phones = dBuf.VectorString()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLBannedGetBannedByPhoneList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLBannedBan
///////////////////////////////////////////////////////////////////////////////

func (m *TLBannedBan) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_banned_ban))

	switch uint32(m.Constructor) {
	case 0x3ddb9258:
		// banned.ban phone:string expires:int reason:string = Bool;
		x.UInt(0x3ddb9258)

		// no flags

		x.String(m.GetPhone())
		x.Int(m.GetExpires())
		x.String(m.GetReason())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLBannedBan) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLBannedBan) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x3ddb9258:
		// banned.ban phone:string expires:int reason:string = Bool;

		// not has flags

		m.Phone = dBuf.String()
		m.Expires = dBuf.Int()
		m.Reason = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLBannedBan) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLBannedUnBan
///////////////////////////////////////////////////////////////////////////////

func (m *TLBannedUnBan) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_banned_unBan))

	switch uint32(m.Constructor) {
	case 0x720029e4:
		// banned.unBan phone:string = Bool;
		x.UInt(0x720029e4)

		// no flags

		x.String(m.GetPhone())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLBannedUnBan) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLBannedUnBan) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x720029e4:
		// banned.unBan phone:string = Bool;

		// not has flags

		m.Phone = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLBannedUnBan) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// Vector_String
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_String) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.String(v)
	}

	return x.GetBuf()
}

func (m *Vector_String) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]string, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = dBuf.String()
	}

	return dBuf.GetError()
}

func (m *Vector_String) CalcByteSize(layer int32) int {
	return 0
}

func (m *Vector_String) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}
