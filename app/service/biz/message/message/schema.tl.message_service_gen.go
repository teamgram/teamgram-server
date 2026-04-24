/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package message

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var (
	_ iface.TLObject
	_ fmt.Stringer
	_ *tg.Bool
	_ bin.Fields
	_ json.Marshaler
)

// TLMessageGetUserMessage <--
type TLMessageGetUserMessage struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Id      int32  `json:"id"`
}

func (m *TLMessageGetUserMessage) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_getUserMessage, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetUserMessage) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_getUserMessage, int(layer)); clazzId {
	case 0x7accb1c8:
		x.PutClazzID(0x7accb1c8)

		x.PutInt64(m.UserId)
		x.PutInt32(m.Id)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_getUserMessage: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageGetUserMessage) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUserMessage: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x7accb1c8:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUserMessage#0x7accb1c8: field user_id: %w", err)
		}
		m.Id, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUserMessage#0x7accb1c8: field id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_getUserMessage: invalid constructor %x", m.ClazzID)
	}
}

// TLMessageGetUserMessageList <--
type TLMessageGetUserMessageList struct {
	ClazzID uint32  `json:"_id"`
	UserId  int64   `json:"user_id"`
	IdList  []int32 `json:"id_list"`
}

func (m *TLMessageGetUserMessageList) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_getUserMessageList, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetUserMessageList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_getUserMessageList, int(layer)); clazzId {
	case 0xd3581c26:
		x.PutClazzID(0xd3581c26)

		x.PutInt64(m.UserId)

		iface.EncodeInt32List(x, m.IdList)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_getUserMessageList: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageGetUserMessageList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUserMessageList: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xd3581c26:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUserMessageList#0xd3581c26: field user_id: %w", err)
		}

		m.IdList, err = iface.DecodeInt32List(d)
		if err != nil {
			return fmt.Errorf("unable to decode message_getUserMessageList#0xd3581c26: field id_list: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_getUserMessageList: invalid constructor %x", m.ClazzID)
	}
}

// TLMessageGetUserMessageListByDataIdList <--
type TLMessageGetUserMessageListByDataIdList struct {
	ClazzID uint32  `json:"_id"`
	UserId  int64   `json:"user_id"`
	IdList  []int64 `json:"id_list"`
}

func (m *TLMessageGetUserMessageListByDataIdList) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_getUserMessageListByDataIdList, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetUserMessageListByDataIdList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_getUserMessageListByDataIdList, int(layer)); clazzId {
	case 0x1155a17b:
		x.PutClazzID(0x1155a17b)

		x.PutInt64(m.UserId)

		iface.EncodeInt64List(x, m.IdList)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_getUserMessageListByDataIdList: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageGetUserMessageListByDataIdList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUserMessageListByDataIdList: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x1155a17b:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUserMessageListByDataIdList#0x1155a17b: field user_id: %w", err)
		}

		m.IdList, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode message_getUserMessageListByDataIdList#0x1155a17b: field id_list: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_getUserMessageListByDataIdList: invalid constructor %x", m.ClazzID)
	}
}

// TLMessageGetUserMessageListByDataIdUserIdList <--
type TLMessageGetUserMessageListByDataIdUserIdList struct {
	ClazzID    uint32  `json:"_id"`
	Id         int64   `json:"id"`
	UserIdList []int64 `json:"user_id_list"`
}

func (m *TLMessageGetUserMessageListByDataIdUserIdList) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_getUserMessageListByDataIdUserIdList, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetUserMessageListByDataIdUserIdList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_getUserMessageListByDataIdUserIdList, int(layer)); clazzId {
	case 0x2cb26a31:
		x.PutClazzID(0x2cb26a31)

		x.PutInt64(m.Id)

		iface.EncodeInt64List(x, m.UserIdList)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_getUserMessageListByDataIdUserIdList: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageGetUserMessageListByDataIdUserIdList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUserMessageListByDataIdUserIdList: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x2cb26a31:
		m.Id, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUserMessageListByDataIdUserIdList#0x2cb26a31: field id: %w", err)
		}

		m.UserIdList, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode message_getUserMessageListByDataIdUserIdList#0x2cb26a31: field user_id_list: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_getUserMessageListByDataIdUserIdList: invalid constructor %x", m.ClazzID)
	}
}

// TLMessageGetHistoryMessages <--
type TLMessageGetHistoryMessages struct {
	ClazzID    uint32 `json:"_id"`
	UserId     int64  `json:"user_id"`
	PeerType   int32  `json:"peer_type"`
	PeerId     int64  `json:"peer_id"`
	OffsetId   int32  `json:"offset_id"`
	OffsetDate int32  `json:"offset_date"`
	AddOffset  int32  `json:"add_offset"`
	Limit      int32  `json:"limit"`
	MaxId      int32  `json:"max_id"`
	MinId      int32  `json:"min_id"`
	Hash       int64  `json:"hash"`
}

func (m *TLMessageGetHistoryMessages) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_getHistoryMessages, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetHistoryMessages) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_getHistoryMessages, int(layer)); clazzId {
	case 0x308a340:
		x.PutClazzID(0x308a340)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt32(m.OffsetId)
		x.PutInt32(m.OffsetDate)
		x.PutInt32(m.AddOffset)
		x.PutInt32(m.Limit)
		x.PutInt32(m.MaxId)
		x.PutInt32(m.MinId)
		x.PutInt64(m.Hash)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_getHistoryMessages: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageGetHistoryMessages) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_getHistoryMessages: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x308a340:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getHistoryMessages#0x308a340: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getHistoryMessages#0x308a340: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getHistoryMessages#0x308a340: field peer_id: %w", err)
		}
		m.OffsetId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getHistoryMessages#0x308a340: field offset_id: %w", err)
		}
		m.OffsetDate, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getHistoryMessages#0x308a340: field offset_date: %w", err)
		}
		m.AddOffset, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getHistoryMessages#0x308a340: field add_offset: %w", err)
		}
		m.Limit, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getHistoryMessages#0x308a340: field limit: %w", err)
		}
		m.MaxId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getHistoryMessages#0x308a340: field max_id: %w", err)
		}
		m.MinId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getHistoryMessages#0x308a340: field min_id: %w", err)
		}
		m.Hash, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getHistoryMessages#0x308a340: field hash: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_getHistoryMessages: invalid constructor %x", m.ClazzID)
	}
}

// TLMessageGetHistoryMessagesCount <--
type TLMessageGetHistoryMessagesCount struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

func (m *TLMessageGetHistoryMessagesCount) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_getHistoryMessagesCount, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetHistoryMessagesCount) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_getHistoryMessagesCount, int(layer)); clazzId {
	case 0xf507e13:
		x.PutClazzID(0xf507e13)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_getHistoryMessagesCount: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageGetHistoryMessagesCount) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_getHistoryMessagesCount: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xf507e13:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getHistoryMessagesCount#0xf507e13: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getHistoryMessagesCount#0xf507e13: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getHistoryMessagesCount#0xf507e13: field peer_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_getHistoryMessagesCount: invalid constructor %x", m.ClazzID)
	}
}

// TLMessageGetPeerUserMessageId <--
type TLMessageGetPeerUserMessageId struct {
	ClazzID    uint32 `json:"_id"`
	UserId     int64  `json:"user_id"`
	PeerUserId int64  `json:"peer_user_id"`
	MsgId      int32  `json:"msg_id"`
}

func (m *TLMessageGetPeerUserMessageId) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_getPeerUserMessageId, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetPeerUserMessageId) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_getPeerUserMessageId, int(layer)); clazzId {
	case 0x73aeb71f:
		x.PutClazzID(0x73aeb71f)

		x.PutInt64(m.UserId)
		x.PutInt64(m.PeerUserId)
		x.PutInt32(m.MsgId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_getPeerUserMessageId: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageGetPeerUserMessageId) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_getPeerUserMessageId: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x73aeb71f:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getPeerUserMessageId#0x73aeb71f: field user_id: %w", err)
		}
		m.PeerUserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getPeerUserMessageId#0x73aeb71f: field peer_user_id: %w", err)
		}
		m.MsgId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getPeerUserMessageId#0x73aeb71f: field msg_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_getPeerUserMessageId: invalid constructor %x", m.ClazzID)
	}
}

// TLMessageGetPeerUserMessage <--
type TLMessageGetPeerUserMessage struct {
	ClazzID    uint32 `json:"_id"`
	UserId     int64  `json:"user_id"`
	PeerUserId int64  `json:"peer_user_id"`
	MsgId      int32  `json:"msg_id"`
}

func (m *TLMessageGetPeerUserMessage) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_getPeerUserMessage, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetPeerUserMessage) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_getPeerUserMessage, int(layer)); clazzId {
	case 0x63129212:
		x.PutClazzID(0x63129212)

		x.PutInt64(m.UserId)
		x.PutInt64(m.PeerUserId)
		x.PutInt32(m.MsgId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_getPeerUserMessage: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageGetPeerUserMessage) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_getPeerUserMessage: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x63129212:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getPeerUserMessage#0x63129212: field user_id: %w", err)
		}
		m.PeerUserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getPeerUserMessage#0x63129212: field peer_user_id: %w", err)
		}
		m.MsgId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getPeerUserMessage#0x63129212: field msg_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_getPeerUserMessage: invalid constructor %x", m.ClazzID)
	}
}

// TLMessageSearchByMediaType <--
type TLMessageSearchByMediaType struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	PeerType  int32  `json:"peer_type"`
	PeerId    int64  `json:"peer_id"`
	MediaType int32  `json:"media_type"`
	Offset    int32  `json:"offset"`
	Limit     int32  `json:"limit"`
}

func (m *TLMessageSearchByMediaType) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_searchByMediaType, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageSearchByMediaType) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_searchByMediaType, int(layer)); clazzId {
	case 0x111c2943:
		x.PutClazzID(0x111c2943)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt32(m.MediaType)
		x.PutInt32(m.Offset)
		x.PutInt32(m.Limit)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_searchByMediaType: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageSearchByMediaType) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchByMediaType: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x111c2943:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchByMediaType#0x111c2943: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchByMediaType#0x111c2943: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchByMediaType#0x111c2943: field peer_id: %w", err)
		}
		m.MediaType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchByMediaType#0x111c2943: field media_type: %w", err)
		}
		m.Offset, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchByMediaType#0x111c2943: field offset: %w", err)
		}
		m.Limit, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchByMediaType#0x111c2943: field limit: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_searchByMediaType: invalid constructor %x", m.ClazzID)
	}
}

// TLMessageSearch <--
type TLMessageSearch struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
	Q        string `json:"q"`
	Offset   int32  `json:"offset"`
	Limit    int32  `json:"limit"`
}

func (m *TLMessageSearch) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_search, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageSearch) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_search, int(layer)); clazzId {
	case 0x6835b023:
		x.PutClazzID(0x6835b023)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutString(m.Q)
		x.PutInt32(m.Offset)
		x.PutInt32(m.Limit)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_search: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageSearch) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_search: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x6835b023:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_search#0x6835b023: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_search#0x6835b023: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_search#0x6835b023: field peer_id: %w", err)
		}
		m.Q, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode message_search#0x6835b023: field q: %w", err)
		}
		m.Offset, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_search#0x6835b023: field offset: %w", err)
		}
		m.Limit, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_search#0x6835b023: field limit: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_search: invalid constructor %x", m.ClazzID)
	}
}

// TLMessageSearchGlobal <--
type TLMessageSearchGlobal struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Q       string `json:"q"`
	Offset  int32  `json:"offset"`
	Limit   int32  `json:"limit"`
}

func (m *TLMessageSearchGlobal) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_searchGlobal, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageSearchGlobal) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_searchGlobal, int(layer)); clazzId {
	case 0xb3985dc5:
		x.PutClazzID(0xb3985dc5)

		x.PutInt64(m.UserId)
		x.PutString(m.Q)
		x.PutInt32(m.Offset)
		x.PutInt32(m.Limit)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_searchGlobal: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageSearchGlobal) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchGlobal: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xb3985dc5:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchGlobal#0xb3985dc5: field user_id: %w", err)
		}
		m.Q, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchGlobal#0xb3985dc5: field q: %w", err)
		}
		m.Offset, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchGlobal#0xb3985dc5: field offset: %w", err)
		}
		m.Limit, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchGlobal#0xb3985dc5: field limit: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_searchGlobal: invalid constructor %x", m.ClazzID)
	}
}

// TLMessageSearchByPinned <--
type TLMessageSearchByPinned struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

func (m *TLMessageSearchByPinned) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_searchByPinned, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageSearchByPinned) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_searchByPinned, int(layer)); clazzId {
	case 0x6e735b55:
		x.PutClazzID(0x6e735b55)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_searchByPinned: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageSearchByPinned) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchByPinned: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x6e735b55:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchByPinned#0x6e735b55: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchByPinned#0x6e735b55: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchByPinned#0x6e735b55: field peer_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_searchByPinned: invalid constructor %x", m.ClazzID)
	}
}

// TLMessageGetSearchCounter <--
type TLMessageGetSearchCounter struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	PeerType  int32  `json:"peer_type"`
	PeerId    int64  `json:"peer_id"`
	MediaType int32  `json:"media_type"`
}

func (m *TLMessageGetSearchCounter) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_getSearchCounter, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetSearchCounter) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_getSearchCounter, int(layer)); clazzId {
	case 0xe2cbbf46:
		x.PutClazzID(0xe2cbbf46)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt32(m.MediaType)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_getSearchCounter: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageGetSearchCounter) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_getSearchCounter: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xe2cbbf46:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getSearchCounter#0xe2cbbf46: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getSearchCounter#0xe2cbbf46: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getSearchCounter#0xe2cbbf46: field peer_id: %w", err)
		}
		m.MediaType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getSearchCounter#0xe2cbbf46: field media_type: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_getSearchCounter: invalid constructor %x", m.ClazzID)
	}
}

// TLMessageSearchV2 <--
type TLMessageSearchV2 struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	PeerType  int32  `json:"peer_type"`
	PeerId    int64  `json:"peer_id"`
	Q         string `json:"q"`
	FromId    int64  `json:"from_id"`
	MinDate   int32  `json:"min_date"`
	MaxDate   int32  `json:"max_date"`
	OffsetId  int32  `json:"offset_id"`
	AddOffset int32  `json:"add_offset"`
	Limit     int32  `json:"limit"`
	MaxId     int32  `json:"max_id"`
	MinId     int32  `json:"min_id"`
	Hash      int64  `json:"hash"`
}

func (m *TLMessageSearchV2) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_searchV2, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageSearchV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_searchV2, int(layer)); clazzId {
	case 0xa1c62b21:
		x.PutClazzID(0xa1c62b21)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutString(m.Q)
		x.PutInt64(m.FromId)
		x.PutInt32(m.MinDate)
		x.PutInt32(m.MaxDate)
		x.PutInt32(m.OffsetId)
		x.PutInt32(m.AddOffset)
		x.PutInt32(m.Limit)
		x.PutInt32(m.MaxId)
		x.PutInt32(m.MinId)
		x.PutInt64(m.Hash)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_searchV2: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageSearchV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchV2: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xa1c62b21:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchV2#0xa1c62b21: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchV2#0xa1c62b21: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchV2#0xa1c62b21: field peer_id: %w", err)
		}
		m.Q, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchV2#0xa1c62b21: field q: %w", err)
		}
		m.FromId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchV2#0xa1c62b21: field from_id: %w", err)
		}
		m.MinDate, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchV2#0xa1c62b21: field min_date: %w", err)
		}
		m.MaxDate, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchV2#0xa1c62b21: field max_date: %w", err)
		}
		m.OffsetId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchV2#0xa1c62b21: field offset_id: %w", err)
		}
		m.AddOffset, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchV2#0xa1c62b21: field add_offset: %w", err)
		}
		m.Limit, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchV2#0xa1c62b21: field limit: %w", err)
		}
		m.MaxId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchV2#0xa1c62b21: field max_id: %w", err)
		}
		m.MinId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchV2#0xa1c62b21: field min_id: %w", err)
		}
		m.Hash, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_searchV2#0xa1c62b21: field hash: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_searchV2: invalid constructor %x", m.ClazzID)
	}
}

// TLMessageGetLastTwoPinnedMessageId <--
type TLMessageGetLastTwoPinnedMessageId struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

func (m *TLMessageGetLastTwoPinnedMessageId) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_getLastTwoPinnedMessageId, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetLastTwoPinnedMessageId) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_getLastTwoPinnedMessageId, int(layer)); clazzId {
	case 0xaf9a082b:
		x.PutClazzID(0xaf9a082b)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_getLastTwoPinnedMessageId: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageGetLastTwoPinnedMessageId) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_getLastTwoPinnedMessageId: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xaf9a082b:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getLastTwoPinnedMessageId#0xaf9a082b: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getLastTwoPinnedMessageId#0xaf9a082b: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getLastTwoPinnedMessageId#0xaf9a082b: field peer_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_getLastTwoPinnedMessageId: invalid constructor %x", m.ClazzID)
	}
}

// TLMessageUpdatePinnedMessageId <--
type TLMessageUpdatePinnedMessageId struct {
	ClazzID  uint32       `json:"_id"`
	UserId   int64        `json:"user_id"`
	PeerType int32        `json:"peer_type"`
	PeerId   int64        `json:"peer_id"`
	Id       int32        `json:"id"`
	Pinned   tg.BoolClazz `json:"pinned"`
}

func (m *TLMessageUpdatePinnedMessageId) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_updatePinnedMessageId, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageUpdatePinnedMessageId) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_updatePinnedMessageId, int(layer)); clazzId {
	case 0xf520edd0:
		x.PutClazzID(0xf520edd0)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt32(m.Id)
		if m.Pinned == nil {
			return fmt.Errorf("unable to encode message_updatePinnedMessageId#0xf520edd0: field pinned is nil")
		}
		if err := m.Pinned.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode message_updatePinnedMessageId#0xf520edd0: field pinned: %w", err)
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_updatePinnedMessageId: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageUpdatePinnedMessageId) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_updatePinnedMessageId: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xf520edd0:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_updatePinnedMessageId#0xf520edd0: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_updatePinnedMessageId#0xf520edd0: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_updatePinnedMessageId#0xf520edd0: field peer_id: %w", err)
		}
		m.Id, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_updatePinnedMessageId#0xf520edd0: field id: %w", err)
		}

		m.Pinned, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode message_updatePinnedMessageId#0xf520edd0: field pinned: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_updatePinnedMessageId: invalid constructor %x", m.ClazzID)
	}
}

// TLMessageGetPinnedMessageIdList <--
type TLMessageGetPinnedMessageIdList struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

func (m *TLMessageGetPinnedMessageIdList) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_getPinnedMessageIdList, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetPinnedMessageIdList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_getPinnedMessageIdList, int(layer)); clazzId {
	case 0xda01d0dd:
		x.PutClazzID(0xda01d0dd)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_getPinnedMessageIdList: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageGetPinnedMessageIdList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_getPinnedMessageIdList: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xda01d0dd:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getPinnedMessageIdList#0xda01d0dd: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getPinnedMessageIdList#0xda01d0dd: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getPinnedMessageIdList#0xda01d0dd: field peer_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_getPinnedMessageIdList: invalid constructor %x", m.ClazzID)
	}
}

// TLMessageUnPinAllMessages <--
type TLMessageUnPinAllMessages struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

func (m *TLMessageUnPinAllMessages) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_unPinAllMessages, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageUnPinAllMessages) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_unPinAllMessages, int(layer)); clazzId {
	case 0xea0a2a73:
		x.PutClazzID(0xea0a2a73)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_unPinAllMessages: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageUnPinAllMessages) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_unPinAllMessages: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xea0a2a73:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_unPinAllMessages#0xea0a2a73: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_unPinAllMessages#0xea0a2a73: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_unPinAllMessages#0xea0a2a73: field peer_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_unPinAllMessages: invalid constructor %x", m.ClazzID)
	}
}

// TLMessageGetUnreadMentions <--
type TLMessageGetUnreadMentions struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	PeerType  int32  `json:"peer_type"`
	PeerId    int64  `json:"peer_id"`
	OffsetId  int32  `json:"offset_id"`
	AddOffset int32  `json:"add_offset"`
	Limit     int32  `json:"limit"`
	MinId     int32  `json:"min_id"`
	MaxInt    int32  `json:"max_int"`
}

func (m *TLMessageGetUnreadMentions) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_getUnreadMentions, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetUnreadMentions) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_getUnreadMentions, int(layer)); clazzId {
	case 0x6fe184b4:
		x.PutClazzID(0x6fe184b4)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt32(m.OffsetId)
		x.PutInt32(m.AddOffset)
		x.PutInt32(m.Limit)
		x.PutInt32(m.MinId)
		x.PutInt32(m.MaxInt)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_getUnreadMentions: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageGetUnreadMentions) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUnreadMentions: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x6fe184b4:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUnreadMentions#0x6fe184b4: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUnreadMentions#0x6fe184b4: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUnreadMentions#0x6fe184b4: field peer_id: %w", err)
		}
		m.OffsetId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUnreadMentions#0x6fe184b4: field offset_id: %w", err)
		}
		m.AddOffset, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUnreadMentions#0x6fe184b4: field add_offset: %w", err)
		}
		m.Limit, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUnreadMentions#0x6fe184b4: field limit: %w", err)
		}
		m.MinId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUnreadMentions#0x6fe184b4: field min_id: %w", err)
		}
		m.MaxInt, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUnreadMentions#0x6fe184b4: field max_int: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_getUnreadMentions: invalid constructor %x", m.ClazzID)
	}
}

// TLMessageGetUnreadMentionsCount <--
type TLMessageGetUnreadMentionsCount struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

func (m *TLMessageGetUnreadMentionsCount) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_message_getUnreadMentionsCount, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetUnreadMentionsCount) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_message_getUnreadMentionsCount, int(layer)); clazzId {
	case 0xb5412049:
		x.PutClazzID(0xb5412049)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("unable to encode message_getUnreadMentionsCount: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageGetUnreadMentionsCount) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUnreadMentionsCount: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xb5412049:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUnreadMentionsCount#0xb5412049: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUnreadMentionsCount#0xb5412049: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode message_getUnreadMentionsCount#0xb5412049: field peer_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode message_getUnreadMentionsCount: invalid constructor %x", m.ClazzID)
	}
}

// Vector api result type
// ----------------------------------------------------------------------------
// VectorResList <--

// VectorMessageBox <--
type VectorMessageBox struct {
	Datas []tg.MessageBoxClazz `json:"_datas"`
}

func (m *VectorMessageBox) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorMessageBox) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorMessageBox) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[tg.MessageBoxClazz](d)

	return err
}

// VectorInt <--
type VectorInt struct {
	Datas []int32 `json:"_datas"`
}

func (m *VectorInt) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorInt) Encode(x *bin.Encoder, layer int32) error {
	iface.EncodeInt32List(x, m.Datas)

	return nil
}

// Decode <--
func (m *VectorInt) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeInt32List(d)

	return err
}

// ----------------------------------------------------------------------------
// rpc

type RPCMessage interface {
	MessageGetUserMessage(ctx context.Context, in *TLMessageGetUserMessage) (*tg.MessageBox, error)
	MessageGetUserMessageList(ctx context.Context, in *TLMessageGetUserMessageList) (*VectorMessageBox, error)
	MessageGetUserMessageListByDataIdList(ctx context.Context, in *TLMessageGetUserMessageListByDataIdList) (*VectorMessageBox, error)
	MessageGetUserMessageListByDataIdUserIdList(ctx context.Context, in *TLMessageGetUserMessageListByDataIdUserIdList) (*VectorMessageBox, error)
	MessageGetHistoryMessages(ctx context.Context, in *TLMessageGetHistoryMessages) (*VectorMessageBox, error)
	MessageGetHistoryMessagesCount(ctx context.Context, in *TLMessageGetHistoryMessagesCount) (*tg.Int32, error)
	MessageGetPeerUserMessageId(ctx context.Context, in *TLMessageGetPeerUserMessageId) (*tg.Int32, error)
	MessageGetPeerUserMessage(ctx context.Context, in *TLMessageGetPeerUserMessage) (*tg.MessageBox, error)
	MessageSearchByMediaType(ctx context.Context, in *TLMessageSearchByMediaType) (*VectorMessageBox, error)
	MessageSearch(ctx context.Context, in *TLMessageSearch) (*VectorMessageBox, error)
	MessageSearchGlobal(ctx context.Context, in *TLMessageSearchGlobal) (*VectorMessageBox, error)
	MessageSearchByPinned(ctx context.Context, in *TLMessageSearchByPinned) (*VectorMessageBox, error)
	MessageGetSearchCounter(ctx context.Context, in *TLMessageGetSearchCounter) (*tg.Int32, error)
	MessageSearchV2(ctx context.Context, in *TLMessageSearchV2) (*VectorMessageBox, error)
	MessageGetLastTwoPinnedMessageId(ctx context.Context, in *TLMessageGetLastTwoPinnedMessageId) (*VectorInt, error)
	MessageUpdatePinnedMessageId(ctx context.Context, in *TLMessageUpdatePinnedMessageId) (*tg.Bool, error)
	MessageGetPinnedMessageIdList(ctx context.Context, in *TLMessageGetPinnedMessageIdList) (*VectorInt, error)
	MessageUnPinAllMessages(ctx context.Context, in *TLMessageUnPinAllMessages) (*VectorInt, error)
	MessageGetUnreadMentions(ctx context.Context, in *TLMessageGetUnreadMentions) (*VectorMessageBox, error)
	MessageGetUnreadMentionsCount(ctx context.Context, in *TLMessageGetUnreadMentionsCount) (*tg.Int32, error)
}
