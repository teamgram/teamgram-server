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

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

//////////////////////////////////////////////////////////////////////////////////////////

var _ *wrapperspb.Int32Value
var _ *mtproto.Bool
var _ fmt.Stringer

var clazzIdRegisters2 = map[int32]func() mtproto.TLObject{
	// Constructor

	// Method
	1614568688: func() mtproto.TLObject { // 0x603c5cf0
		return &TLSyncUpdatesMe{
			Constructor: 1614568688,
		}
	},
	16458447: func() mtproto.TLObject { // 0xfb22cf
		return &TLSyncUpdatesNotMe{
			Constructor: 16458447,
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
	-1874085983: func() mtproto.TLObject { // 0x904bb7a1
		return &TLSyncPushRpcResult{
			Constructor: -1874085983,
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
	case 0x603c5cf0:
		x.UInt(0x603c5cf0)

		// set flags
		var flags uint32 = 0

		if m.GetSessionId() != nil {
			flags |= 1 << 0
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.String(m.GetServerId())
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
	case 0x603c5cf0:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		m.ServerId = dBuf.String()
		if (flags & (1 << 0)) != 0 {
			m.SessionId = &wrapperspb.Int64Value{Value: dBuf.Long()}
		}

		m6 := &mtproto.Updates{}
		m6.Decode(dBuf)
		m.Updates = m6

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLSyncUpdatesMe) DebugString() string {
	v, _ := protojson.Marshal(m)
	return string(v)
}

// TLSyncUpdatesNotMe
///////////////////////////////////////////////////////////////////////////////

func (m *TLSyncUpdatesNotMe) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xfb22cf:
		x.UInt(0xfb22cf)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
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
	case 0xfb22cf:

		// not has flags

		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()

		m3 := &mtproto.Updates{}
		m3.Decode(dBuf)
		m.Updates = m3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLSyncUpdatesNotMe) DebugString() string {
	v, _ := protojson.Marshal(m)
	return string(v)
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

func (m *TLSyncPushUpdates) DebugString() string {
	v, _ := protojson.Marshal(m)
	return string(v)
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

func (m *TLSyncPushUpdatesIfNot) DebugString() string {
	v, _ := protojson.Marshal(m)
	return string(v)
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

func (m *TLSyncPushBotUpdates) DebugString() string {
	v, _ := protojson.Marshal(m)
	return string(v)
}

// TLSyncPushRpcResult
///////////////////////////////////////////////////////////////////////////////

func (m *TLSyncPushRpcResult) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x904bb7a1:
		x.UInt(0x904bb7a1)

		// no flags

		x.Long(m.GetAuthKeyId())
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
	case 0x904bb7a1:

		// not has flags

		m.AuthKeyId = dBuf.Long()
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

func (m *TLSyncPushRpcResult) DebugString() string {
	v, _ := protojson.Marshal(m)
	return string(v)
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

func (m *TLSyncBroadcastUpdates) DebugString() string {
	v, _ := protojson.Marshal(m)
	return string(v)
}

//----------------------------------------------------------------------------------------------------------------
