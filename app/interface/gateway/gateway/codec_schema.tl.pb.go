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

package gateway

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
	645953552: func() mtproto.TLObject { // 0x26807810
		return &TLGatewaySendDataToGateway{
			Constructor: 645953552,
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
// TLGatewaySendDataToGateway
///////////////////////////////////////////////////////////////////////////////
func (m *TLGatewaySendDataToGateway) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_gateway_sendDataToGateway))

	switch uint32(m.Constructor) {
	case 0x26807810:
		// gateway.sendDataToGateway auth_key_id:long session_id:long payload:bytes = Bool;
		x.UInt(0x26807810)

		// no flags

		x.Long(m.GetAuthKeyId())
		x.Long(m.GetSessionId())
		x.StringBytes(m.GetPayload())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLGatewaySendDataToGateway) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLGatewaySendDataToGateway) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x26807810:
		// gateway.sendDataToGateway auth_key_id:long session_id:long payload:bytes = Bool;

		// not has flags

		m.AuthKeyId = dBuf.Long()
		m.SessionId = dBuf.Long()
		m.Payload = dBuf.StringBytes()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLGatewaySendDataToGateway) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
