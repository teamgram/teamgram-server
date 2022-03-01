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

package sync

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

func (m *TLSyncUpdatesMe) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_sync_updatesMe))

	switch uint32(m.Constructor) {
	case 0x603c5cf0:
		// sync.updatesMe flags:# user_id:long auth_key_id:long server_id:string session_id:flags.0?long updates:Updates = Void;
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

		x.Bytes(m.GetUpdates().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLSyncUpdatesMe) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSyncUpdatesMe) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x603c5cf0:
		// sync.updatesMe flags:# user_id:long auth_key_id:long server_id:string session_id:flags.0?long updates:Updates = Void;

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		m.ServerId = dBuf.String()
		if (flags & (1 << 0)) != 0 {
			m.SessionId = &types.Int64Value{Value: dBuf.Long()}
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
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLSyncUpdatesNotMe
///////////////////////////////////////////////////////////////////////////////

func (m *TLSyncUpdatesNotMe) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_sync_updatesNotMe))

	switch uint32(m.Constructor) {
	case 0xfb22cf:
		// sync.updatesNotMe user_id:long auth_key_id:long updates:Updates = Void;
		x.UInt(0xfb22cf)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Bytes(m.GetUpdates().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLSyncUpdatesNotMe) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSyncUpdatesNotMe) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xfb22cf:
		// sync.updatesNotMe user_id:long auth_key_id:long updates:Updates = Void;

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
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLSyncPushUpdates
///////////////////////////////////////////////////////////////////////////////

func (m *TLSyncPushUpdates) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_sync_pushUpdates))

	switch uint32(m.Constructor) {
	case 0x8f0ad9be:
		// sync.pushUpdates user_id:long updates:Updates = Void;
		x.UInt(0x8f0ad9be)

		// no flags

		x.Long(m.GetUserId())
		x.Bytes(m.GetUpdates().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLSyncPushUpdates) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSyncPushUpdates) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x8f0ad9be:
		// sync.pushUpdates user_id:long updates:Updates = Void;

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
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLSyncPushUpdatesIfNot
///////////////////////////////////////////////////////////////////////////////

func (m *TLSyncPushUpdatesIfNot) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_sync_pushUpdatesIfNot))

	switch uint32(m.Constructor) {
	case 0x40053fe4:
		// sync.pushUpdatesIfNot user_id:long excludes:Vector<long> updates:Updates = Void;
		x.UInt(0x40053fe4)

		// no flags

		x.Long(m.GetUserId())

		x.VectorLong(m.GetExcludes())

		x.Bytes(m.GetUpdates().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLSyncPushUpdatesIfNot) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSyncPushUpdatesIfNot) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x40053fe4:
		// sync.pushUpdatesIfNot user_id:long excludes:Vector<long> updates:Updates = Void;

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
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLSyncPushBotUpdates
///////////////////////////////////////////////////////////////////////////////

func (m *TLSyncPushBotUpdates) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_sync_pushBotUpdates))

	switch uint32(m.Constructor) {
	case 0xadc3f000:
		// sync.pushBotUpdates user_id:long updates:Updates = Void;
		x.UInt(0xadc3f000)

		// no flags

		x.Long(m.GetUserId())
		x.Bytes(m.GetUpdates().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLSyncPushBotUpdates) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSyncPushBotUpdates) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xadc3f000:
		// sync.pushBotUpdates user_id:long updates:Updates = Void;

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
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLSyncPushRpcResult
///////////////////////////////////////////////////////////////////////////////

func (m *TLSyncPushRpcResult) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_sync_pushRpcResult))

	switch uint32(m.Constructor) {
	case 0x904bb7a1:
		// sync.pushRpcResult auth_key_id:long server_id:string session_id:long client_req_msg_id:long rpc_result:bytes = Void;
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

	return x.GetBuf()
}

func (m *TLSyncPushRpcResult) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSyncPushRpcResult) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x904bb7a1:
		// sync.pushRpcResult auth_key_id:long server_id:string session_id:long client_req_msg_id:long rpc_result:bytes = Void;

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
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLSyncBroadcastUpdates
///////////////////////////////////////////////////////////////////////////////

func (m *TLSyncBroadcastUpdates) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_sync_broadcastUpdates))

	switch uint32(m.Constructor) {
	case 0xf5e35cb6:
		// sync.broadcastUpdates broadcast_type:int chat_id:long exclude_id_list:Vector<long> updates:Updates = Void;
		x.UInt(0xf5e35cb6)

		// no flags

		x.Int(m.GetBroadcastType())
		x.Long(m.GetChatId())

		x.VectorLong(m.GetExcludeIdList())

		x.Bytes(m.GetUpdates().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLSyncBroadcastUpdates) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSyncBroadcastUpdates) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf5e35cb6:
		// sync.broadcastUpdates broadcast_type:int chat_id:long exclude_id_list:Vector<long> updates:Updates = Void;

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
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
