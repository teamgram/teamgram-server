/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

// ConstructorList
// RequestList

package message

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
	2060235208: func() mtproto.TLObject { // 0x7accb1c8
		return &TLMessageGetUserMessage{
			Constructor: 2060235208,
		}
	},
	-749200346: func() mtproto.TLObject { // 0xd3581c26
		return &TLMessageGetUserMessageList{
			Constructor: -749200346,
		}
	},
	290824571: func() mtproto.TLObject { // 0x1155a17b
		return &TLMessageGetUserMessageListByDataIdList{
			Constructor: 290824571,
		}
	},
	749890097: func() mtproto.TLObject { // 0x2cb26a31
		return &TLMessageGetUserMessageListByDataIdUserIdList{
			Constructor: 749890097,
		}
	},
	50897728: func() mtproto.TLObject { // 0x308a340
		return &TLMessageGetHistoryMessages{
			Constructor: 50897728,
		}
	},
	256933395: func() mtproto.TLObject { // 0xf507e13
		return &TLMessageGetHistoryMessagesCount{
			Constructor: 256933395,
		}
	},
	1940829983: func() mtproto.TLObject { // 0x73aeb71f
		return &TLMessageGetPeerUserMessageId{
			Constructor: 1940829983,
		}
	},
	1662161426: func() mtproto.TLObject { // 0x63129212
		return &TLMessageGetPeerUserMessage{
			Constructor: 1662161426,
		}
	},
	287058243: func() mtproto.TLObject { // 0x111c2943
		return &TLMessageSearchByMediaType{
			Constructor: 287058243,
		}
	},
	1748348963: func() mtproto.TLObject { // 0x6835b023
		return &TLMessageSearch{
			Constructor: 1748348963,
		}
	},
	-1281860155: func() mtproto.TLObject { // 0xb3985dc5
		return &TLMessageSearchGlobal{
			Constructor: -1281860155,
		}
	},
	1853053781: func() mtproto.TLObject { // 0x6e735b55
		return &TLMessageSearchByPinned{
			Constructor: 1853053781,
		}
	},
	-489963706: func() mtproto.TLObject { // 0xe2cbbf46
		return &TLMessageGetSearchCounter{
			Constructor: -489963706,
		}
	},
	-1580848351: func() mtproto.TLObject { // 0xa1c62b21
		return &TLMessageSearchV2{
			Constructor: -1580848351,
		}
	},
	-1348859861: func() mtproto.TLObject { // 0xaf9a082b
		return &TLMessageGetLastTwoPinnedMessageId{
			Constructor: -1348859861,
		}
	},
	-182391344: func() mtproto.TLObject { // 0xf520edd0
		return &TLMessageUpdatePinnedMessageId{
			Constructor: -182391344,
		}
	},
	-637415203: func() mtproto.TLObject { // 0xda01d0dd
		return &TLMessageGetPinnedMessageIdList{
			Constructor: -637415203,
		}
	},
	-368432525: func() mtproto.TLObject { // 0xea0a2a73
		return &TLMessageUnPinAllMessages{
			Constructor: -368432525,
		}
	},
	1877050548: func() mtproto.TLObject { // 0x6fe184b4
		return &TLMessageGetUnreadMentions{
			Constructor: 1877050548,
		}
	},
	-1254023095: func() mtproto.TLObject { // 0xb5412049
		return &TLMessageGetUnreadMentionsCount{
			Constructor: -1254023095,
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
// TLMessageGetUserMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageGetUserMessage) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_getUserMessage))

	switch uint32(m.Constructor) {
	case 0x7accb1c8:
		x.UInt(0x7accb1c8)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageGetUserMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageGetUserMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x7accb1c8:

		// not has flags

		m.UserId = dBuf.Long()
		m.Id = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageGetUserMessage) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMessageGetUserMessageList
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageGetUserMessageList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_getUserMessageList))

	switch uint32(m.Constructor) {
	case 0xd3581c26:
		x.UInt(0xd3581c26)

		// no flags

		x.Long(m.GetUserId())

		x.VectorInt(m.GetIdList())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageGetUserMessageList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageGetUserMessageList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xd3581c26:

		// not has flags

		m.UserId = dBuf.Long()

		m.IdList = dBuf.VectorInt()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageGetUserMessageList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMessageGetUserMessageListByDataIdList
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageGetUserMessageListByDataIdList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_getUserMessageListByDataIdList))

	switch uint32(m.Constructor) {
	case 0x1155a17b:
		x.UInt(0x1155a17b)

		// no flags

		x.Long(m.GetUserId())

		x.VectorLong(m.GetIdList())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageGetUserMessageListByDataIdList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageGetUserMessageListByDataIdList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x1155a17b:

		// not has flags

		m.UserId = dBuf.Long()

		m.IdList = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageGetUserMessageListByDataIdList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMessageGetUserMessageListByDataIdUserIdList
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageGetUserMessageListByDataIdUserIdList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_getUserMessageListByDataIdUserIdList))

	switch uint32(m.Constructor) {
	case 0x2cb26a31:
		x.UInt(0x2cb26a31)

		// no flags

		x.Long(m.GetId())

		x.VectorLong(m.GetUserIdList())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageGetUserMessageListByDataIdUserIdList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageGetUserMessageListByDataIdUserIdList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x2cb26a31:

		// not has flags

		m.Id = dBuf.Long()

		m.UserIdList = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageGetUserMessageListByDataIdUserIdList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMessageGetHistoryMessages
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageGetHistoryMessages) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_getHistoryMessages))

	switch uint32(m.Constructor) {
	case 0x308a340:
		x.UInt(0x308a340)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetOffsetId())
		x.Int(m.GetOffsetDate())
		x.Int(m.GetAddOffset())
		x.Int(m.GetLimit())
		x.Int(m.GetMaxId())
		x.Int(m.GetMinId())
		x.Long(m.GetHash())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageGetHistoryMessages) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageGetHistoryMessages) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x308a340:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.OffsetId = dBuf.Int()
		m.OffsetDate = dBuf.Int()
		m.AddOffset = dBuf.Int()
		m.Limit = dBuf.Int()
		m.MaxId = dBuf.Int()
		m.MinId = dBuf.Int()
		m.Hash = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageGetHistoryMessages) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMessageGetHistoryMessagesCount
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageGetHistoryMessagesCount) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_getHistoryMessagesCount))

	switch uint32(m.Constructor) {
	case 0xf507e13:
		x.UInt(0xf507e13)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageGetHistoryMessagesCount) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageGetHistoryMessagesCount) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf507e13:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageGetHistoryMessagesCount) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMessageGetPeerUserMessageId
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageGetPeerUserMessageId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_getPeerUserMessageId))

	switch uint32(m.Constructor) {
	case 0x73aeb71f:
		x.UInt(0x73aeb71f)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetPeerUserId())
		x.Int(m.GetMsgId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageGetPeerUserMessageId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageGetPeerUserMessageId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x73aeb71f:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerUserId = dBuf.Long()
		m.MsgId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageGetPeerUserMessageId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMessageGetPeerUserMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageGetPeerUserMessage) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_getPeerUserMessage))

	switch uint32(m.Constructor) {
	case 0x63129212:
		x.UInt(0x63129212)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetPeerUserId())
		x.Int(m.GetMsgId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageGetPeerUserMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageGetPeerUserMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x63129212:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerUserId = dBuf.Long()
		m.MsgId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageGetPeerUserMessage) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMessageSearchByMediaType
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageSearchByMediaType) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_searchByMediaType))

	switch uint32(m.Constructor) {
	case 0x111c2943:
		x.UInt(0x111c2943)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetMediaType())
		x.Int(m.GetOffset())
		x.Int(m.GetLimit())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageSearchByMediaType) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageSearchByMediaType) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x111c2943:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.MediaType = dBuf.Int()
		m.Offset = dBuf.Int()
		m.Limit = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageSearchByMediaType) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMessageSearch
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageSearch) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_search))

	switch uint32(m.Constructor) {
	case 0x6835b023:
		x.UInt(0x6835b023)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.String(m.GetQ())
		x.Int(m.GetOffset())
		x.Int(m.GetLimit())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageSearch) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageSearch) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x6835b023:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.Q = dBuf.String()
		m.Offset = dBuf.Int()
		m.Limit = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageSearch) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMessageSearchGlobal
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageSearchGlobal) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_searchGlobal))

	switch uint32(m.Constructor) {
	case 0xb3985dc5:
		x.UInt(0xb3985dc5)

		// no flags

		x.Long(m.GetUserId())
		x.String(m.GetQ())
		x.Int(m.GetOffset())
		x.Int(m.GetLimit())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageSearchGlobal) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageSearchGlobal) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb3985dc5:

		// not has flags

		m.UserId = dBuf.Long()
		m.Q = dBuf.String()
		m.Offset = dBuf.Int()
		m.Limit = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageSearchGlobal) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMessageSearchByPinned
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageSearchByPinned) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_searchByPinned))

	switch uint32(m.Constructor) {
	case 0x6e735b55:
		x.UInt(0x6e735b55)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageSearchByPinned) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageSearchByPinned) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x6e735b55:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageSearchByPinned) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMessageGetSearchCounter
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageGetSearchCounter) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_getSearchCounter))

	switch uint32(m.Constructor) {
	case 0xe2cbbf46:
		x.UInt(0xe2cbbf46)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetMediaType())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageGetSearchCounter) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageGetSearchCounter) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe2cbbf46:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.MediaType = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageGetSearchCounter) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMessageSearchV2
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageSearchV2) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_searchV2))

	switch uint32(m.Constructor) {
	case 0xa1c62b21:
		x.UInt(0xa1c62b21)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.String(m.GetQ())
		x.Long(m.GetFromId())
		x.Int(m.GetMinDate())
		x.Int(m.GetMaxDate())
		x.Int(m.GetOffsetId())
		x.Int(m.GetAddOffset())
		x.Int(m.GetLimit())
		x.Int(m.GetMaxId())
		x.Int(m.GetMinId())
		x.Long(m.GetHash())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageSearchV2) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageSearchV2) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xa1c62b21:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.Q = dBuf.String()
		m.FromId = dBuf.Long()
		m.MinDate = dBuf.Int()
		m.MaxDate = dBuf.Int()
		m.OffsetId = dBuf.Int()
		m.AddOffset = dBuf.Int()
		m.Limit = dBuf.Int()
		m.MaxId = dBuf.Int()
		m.MinId = dBuf.Int()
		m.Hash = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageSearchV2) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMessageGetLastTwoPinnedMessageId
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageGetLastTwoPinnedMessageId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_getLastTwoPinnedMessageId))

	switch uint32(m.Constructor) {
	case 0xaf9a082b:
		x.UInt(0xaf9a082b)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageGetLastTwoPinnedMessageId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageGetLastTwoPinnedMessageId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xaf9a082b:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageGetLastTwoPinnedMessageId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMessageUpdatePinnedMessageId
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageUpdatePinnedMessageId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_updatePinnedMessageId))

	switch uint32(m.Constructor) {
	case 0xf520edd0:
		x.UInt(0xf520edd0)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetId())
		x.Bytes(m.GetPinned().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageUpdatePinnedMessageId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageUpdatePinnedMessageId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf520edd0:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.Id = dBuf.Int()

		m5 := &mtproto.Bool{}
		m5.Decode(dBuf)
		m.Pinned = m5

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageUpdatePinnedMessageId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMessageGetPinnedMessageIdList
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageGetPinnedMessageIdList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_getPinnedMessageIdList))

	switch uint32(m.Constructor) {
	case 0xda01d0dd:
		x.UInt(0xda01d0dd)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageGetPinnedMessageIdList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageGetPinnedMessageIdList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xda01d0dd:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageGetPinnedMessageIdList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMessageUnPinAllMessages
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageUnPinAllMessages) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_unPinAllMessages))

	switch uint32(m.Constructor) {
	case 0xea0a2a73:
		x.UInt(0xea0a2a73)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageUnPinAllMessages) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageUnPinAllMessages) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xea0a2a73:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageUnPinAllMessages) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMessageGetUnreadMentions
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageGetUnreadMentions) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_getUnreadMentions))

	switch uint32(m.Constructor) {
	case 0x6fe184b4:
		x.UInt(0x6fe184b4)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetOffsetId())
		x.Int(m.GetAddOffset())
		x.Int(m.GetLimit())
		x.Int(m.GetMinId())
		x.Int(m.GetMaxInt())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageGetUnreadMentions) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageGetUnreadMentions) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x6fe184b4:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.OffsetId = dBuf.Int()
		m.AddOffset = dBuf.Int()
		m.Limit = dBuf.Int()
		m.MinId = dBuf.Int()
		m.MaxInt = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageGetUnreadMentions) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMessageGetUnreadMentionsCount
///////////////////////////////////////////////////////////////////////////////

func (m *TLMessageGetUnreadMentionsCount) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_message_getUnreadMentionsCount))

	switch uint32(m.Constructor) {
	case 0xb5412049:
		x.UInt(0xb5412049)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMessageGetUnreadMentionsCount) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMessageGetUnreadMentionsCount) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb5412049:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMessageGetUnreadMentionsCount) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// Vector_MessageBox
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_MessageBox) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
}

func (m *Vector_MessageBox) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*mtproto.MessageBox, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(mtproto.MessageBox)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_MessageBox) CalcByteSize(layer int32) int {
	return 0
}

func (m *Vector_MessageBox) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// Vector_Int
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_Int) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.VectorInt(m.Datas)

	return x.GetBuf()
}

func (m *Vector_Int) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Datas = dBuf.VectorInt()

	return dBuf.GetError()
}

func (m *Vector_Int) CalcByteSize(layer int32) int {
	return 0
}

func (m *Vector_Int) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}
