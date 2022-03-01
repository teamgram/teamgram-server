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

package auth

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
	-1210022402: func() mtproto.TLObject { // 0xb7e085fe
		return &TLAuthExportLoginToken{
			Constructor: -1210022402,
		}
	},
	-1783866140: func() mtproto.TLObject { // 0x95ac5ce4
		return &TLAuthImportLoginToken{
			Constructor: -1783866140,
		}
	},
	-392909491: func() mtproto.TLObject { // 0xe894ad4d
		return &TLAuthAcceptLoginToken{
			Constructor: -392909491,
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
// TLAuthExportLoginToken
///////////////////////////////////////////////////////////////////////////////
func (m *TLAuthExportLoginToken) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_auth_exportLoginToken))

	switch uint32(m.Constructor) {
	case 0xb7e085fe:
		// auth.exportLoginToken api_id:int api_hash:string except_ids:Vector<long> = auth.LoginToken;
		x.UInt(0xb7e085fe)

		// no flags

		x.Int(m.GetApiId())
		x.String(m.GetApiHash())

		x.VectorLong(m.GetExceptIds())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthExportLoginToken) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthExportLoginToken) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb7e085fe:
		// auth.exportLoginToken api_id:int api_hash:string except_ids:Vector<long> = auth.LoginToken;

		// not has flags

		m.ApiId = dBuf.Int()
		m.ApiHash = dBuf.String()

		m.ExceptIds = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthExportLoginToken) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLAuthImportLoginToken
///////////////////////////////////////////////////////////////////////////////
func (m *TLAuthImportLoginToken) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_auth_importLoginToken))

	switch uint32(m.Constructor) {
	case 0x95ac5ce4:
		// auth.importLoginToken token:bytes = auth.LoginToken;
		x.UInt(0x95ac5ce4)

		// no flags

		x.StringBytes(m.GetToken())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthImportLoginToken) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthImportLoginToken) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x95ac5ce4:
		// auth.importLoginToken token:bytes = auth.LoginToken;

		// not has flags

		m.Token = dBuf.StringBytes()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthImportLoginToken) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLAuthAcceptLoginToken
///////////////////////////////////////////////////////////////////////////////
func (m *TLAuthAcceptLoginToken) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_auth_acceptLoginToken))

	switch uint32(m.Constructor) {
	case 0xe894ad4d:
		// auth.acceptLoginToken token:bytes = Authorization;
		x.UInt(0xe894ad4d)

		// no flags

		x.StringBytes(m.GetToken())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthAcceptLoginToken) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthAcceptLoginToken) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe894ad4d:
		// auth.acceptLoginToken token:bytes = Authorization;

		// not has flags

		m.Token = dBuf.StringBytes()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthAcceptLoginToken) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
