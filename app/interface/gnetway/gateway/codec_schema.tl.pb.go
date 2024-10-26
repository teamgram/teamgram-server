/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

// ConstructorList
// RequestList

package gateway

import (
	"fmt"

	"github.com/teamgram/proto/mtproto"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

//////////////////////////////////////////////////////////////////////////////////////////

var _ *wrapperspb.Int32Value
var _ *mtproto.Bool
var _ fmt.Stringer

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

func (m *TLGatewaySendDataToGateway) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x26807810:
		x.UInt(0x26807810)

		// no flags

		x.Long(m.GetAuthKeyId())
		x.Long(m.GetSessionId())
		x.StringBytes(m.GetPayload())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLGatewaySendDataToGateway) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLGatewaySendDataToGateway) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x26807810:

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
