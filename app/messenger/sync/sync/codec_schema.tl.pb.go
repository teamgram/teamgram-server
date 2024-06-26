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

package sync

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
	-444776161: func() mtproto.TLObject { // 0xe57d411f
		return &TLSyncUpdatesMe{
			Constructor: -444776161,
		}
	},
	-1750314959: func() mtproto.TLObject { // 0x97ac5031
		return &TLSyncUpdatesNotMe{
			Constructor: -1750314959,
		}
	},
	-1895114306: func() mtproto.TLObject { // 0x8f0ad9be
		return &TLSyncPushUpdates{
			Constructor: -1895114306,
		}
	},
	1074085860: func() mtproto.TLObject { // 0x40053fe4
		return &TLSyncPushUpdatesIfNot{
			Constructor: 1074085860,
		}
	},
	-1379667968: func() mtproto.TLObject { // 0xadc3f000
		return &TLSyncPushBotUpdates{
			Constructor: -1379667968,
		}
	},
	828180415: func() mtproto.TLObject { // 0x315d07bf
		return &TLSyncPushRpcResult{
			Constructor: 828180415,
		}
	},
	-169648970: func() mtproto.TLObject { // 0xf5e35cb6
		return &TLSyncBroadcastUpdates{
			Constructor: -169648970,
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
// TLSyncUpdatesMe
///////////////////////////////////////////////////////////////////////////////

func (m *TLSyncUpdatesMe) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xe57d411f:
		x.UInt(0xe57d411f)

		// set flags
		var flags uint32 = 0

		if m.GetServerId() != nil {
			flags |= 1 << 0
		}
		if m.GetAuthKeyId() != nil {
			flags |= 1 << 1
		}
		if m.GetSessionId() != nil {
			flags |= 1 << 1
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetUserId())
		x.Long(m.GetPermAuthKeyId())
		if m.GetServerId() != nil {
			x.String(m.GetServerId().Value)
		}

		if m.GetAuthKeyId() != nil {
			x.Long(m.GetAuthKeyId().Value)
		}

		if m.GetSessionId() != nil {
			x.Long(m.GetSessionId().Value)
		}

		m.GetUpdates().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLSyncUpdatesMe) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSyncUpdatesMe) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe57d411f:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.UserId = dBuf.Long()
		m.PermAuthKeyId = dBuf.Long()
		if (flags & (1 << 0)) != 0 {
			m.ServerId = &wrapperspb.StringValue{Value: dBuf.String()}
		}

		if (flags & (1 << 1)) != 0 {
			m.AuthKeyId = &wrapperspb.Int64Value{Value: dBuf.Long()}
		}

		if (flags & (1 << 1)) != 0 {
			m.SessionId = &wrapperspb.Int64Value{Value: dBuf.Long()}
		}

		m7 := &mtproto.Updates{}
		m7.Decode(dBuf)
		m.Updates = m7

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLSyncUpdatesNotMe
///////////////////////////////////////////////////////////////////////////////

func (m *TLSyncUpdatesNotMe) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x97ac5031:
		x.UInt(0x97ac5031)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetPermAuthKeyId())
		m.GetUpdates().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLSyncUpdatesNotMe) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSyncUpdatesNotMe) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x97ac5031:

		// not has flags

		m.UserId = dBuf.Long()
		m.PermAuthKeyId = dBuf.Long()

		m3 := &mtproto.Updates{}
		m3.Decode(dBuf)
		m.Updates = m3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLSyncPushUpdates
///////////////////////////////////////////////////////////////////////////////

func (m *TLSyncPushUpdates) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x8f0ad9be:
		x.UInt(0x8f0ad9be)

		// no flags

		x.Long(m.GetUserId())
		m.GetUpdates().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLSyncPushUpdates) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSyncPushUpdates) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x8f0ad9be:

		// not has flags

		m.UserId = dBuf.Long()

		m2 := &mtproto.Updates{}
		m2.Decode(dBuf)
		m.Updates = m2

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLSyncPushUpdatesIfNot
///////////////////////////////////////////////////////////////////////////////

func (m *TLSyncPushUpdatesIfNot) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x40053fe4:
		x.UInt(0x40053fe4)

		// no flags

		x.Long(m.GetUserId())

		x.VectorLong(m.GetExcludes())

		m.GetUpdates().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLSyncPushUpdatesIfNot) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSyncPushUpdatesIfNot) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x40053fe4:

		// not has flags

		m.UserId = dBuf.Long()

		m.Excludes = dBuf.VectorLong()

		m3 := &mtproto.Updates{}
		m3.Decode(dBuf)
		m.Updates = m3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLSyncPushBotUpdates
///////////////////////////////////////////////////////////////////////////////

func (m *TLSyncPushBotUpdates) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xadc3f000:
		x.UInt(0xadc3f000)

		// no flags

		x.Long(m.GetUserId())
		m.GetUpdates().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLSyncPushBotUpdates) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSyncPushBotUpdates) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xadc3f000:

		// not has flags

		m.UserId = dBuf.Long()

		m2 := &mtproto.Updates{}
		m2.Decode(dBuf)
		m.Updates = m2

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLSyncPushRpcResult
///////////////////////////////////////////////////////////////////////////////

func (m *TLSyncPushRpcResult) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x315d07bf:
		x.UInt(0x315d07bf)

		// no flags

		x.Long(m.GetAuthKeyId())
		x.Long(m.GetPermAuthKeyId())
		x.String(m.GetServerId())
		x.Long(m.GetSessionId())
		x.Long(m.GetClientReqMsgId())
		x.StringBytes(m.GetRpcResult())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLSyncPushRpcResult) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSyncPushRpcResult) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x315d07bf:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		m.PermAuthKeyId = dBuf.Long()
		m.ServerId = dBuf.String()
		m.SessionId = dBuf.Long()
		m.ClientReqMsgId = dBuf.Long()
		m.RpcResult = dBuf.StringBytes()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLSyncBroadcastUpdates
///////////////////////////////////////////////////////////////////////////////

func (m *TLSyncBroadcastUpdates) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xf5e35cb6:
		x.UInt(0xf5e35cb6)

		// no flags

		x.Int(m.GetBroadcastType())
		x.Long(m.GetChatId())

		x.VectorLong(m.GetExcludeIdList())

		m.GetUpdates().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLSyncBroadcastUpdates) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSyncBroadcastUpdates) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf5e35cb6:

		// not has flags

		m.BroadcastType = dBuf.Int()
		m.ChatId = dBuf.Long()

		m.ExcludeIdList = dBuf.VectorLong()

		m4 := &mtproto.Updates{}
		m4.Decode(dBuf)
		m.Updates = m4

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}
