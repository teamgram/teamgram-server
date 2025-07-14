/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package message

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetUserMessage) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x7accb1c8: func() error {
			x.PutClazzID(0x7accb1c8)

			x.PutInt64(m.UserId)
			x.PutInt32(m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_getUserMessage, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_getUserMessage, layer)
	}
}

// Decode <--
func (m *TLMessageGetUserMessage) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x7accb1c8: func() (err error) {
			m.UserId, err = d.Int64()
			m.Id, err = d.Int32()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMessageGetUserMessageList <--
type TLMessageGetUserMessageList struct {
	ClazzID uint32  `json:"_id"`
	UserId  int64   `json:"user_id"`
	IdList  []int32 `json:"id_list"`
}

func (m *TLMessageGetUserMessageList) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetUserMessageList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xd3581c26: func() error {
			x.PutClazzID(0xd3581c26)

			x.PutInt64(m.UserId)

			iface.EncodeInt32List(x, m.IdList)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_getUserMessageList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_getUserMessageList, layer)
	}
}

// Decode <--
func (m *TLMessageGetUserMessageList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xd3581c26: func() (err error) {
			m.UserId, err = d.Int64()

			m.IdList, err = iface.DecodeInt32List(d)

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMessageGetUserMessageListByDataIdList <--
type TLMessageGetUserMessageListByDataIdList struct {
	ClazzID uint32  `json:"_id"`
	UserId  int64   `json:"user_id"`
	IdList  []int64 `json:"id_list"`
}

func (m *TLMessageGetUserMessageListByDataIdList) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetUserMessageListByDataIdList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x1155a17b: func() error {
			x.PutClazzID(0x1155a17b)

			x.PutInt64(m.UserId)

			iface.EncodeInt64List(x, m.IdList)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_getUserMessageListByDataIdList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_getUserMessageListByDataIdList, layer)
	}
}

// Decode <--
func (m *TLMessageGetUserMessageListByDataIdList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x1155a17b: func() (err error) {
			m.UserId, err = d.Int64()

			m.IdList, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMessageGetUserMessageListByDataIdUserIdList <--
type TLMessageGetUserMessageListByDataIdUserIdList struct {
	ClazzID    uint32  `json:"_id"`
	Id         int64   `json:"id"`
	UserIdList []int64 `json:"user_id_list"`
}

func (m *TLMessageGetUserMessageListByDataIdUserIdList) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetUserMessageListByDataIdUserIdList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x2cb26a31: func() error {
			x.PutClazzID(0x2cb26a31)

			x.PutInt64(m.Id)

			iface.EncodeInt64List(x, m.UserIdList)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_getUserMessageListByDataIdUserIdList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_getUserMessageListByDataIdUserIdList, layer)
	}
}

// Decode <--
func (m *TLMessageGetUserMessageListByDataIdUserIdList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x2cb26a31: func() (err error) {
			m.Id, err = d.Int64()

			m.UserIdList, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetHistoryMessages) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x308a340: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_getHistoryMessages, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_getHistoryMessages, layer)
	}
}

// Decode <--
func (m *TLMessageGetHistoryMessages) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x308a340: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.OffsetId, err = d.Int32()
			m.OffsetDate, err = d.Int32()
			m.AddOffset, err = d.Int32()
			m.Limit, err = d.Int32()
			m.MaxId, err = d.Int32()
			m.MinId, err = d.Int32()
			m.Hash, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetHistoryMessagesCount) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xf507e13: func() error {
			x.PutClazzID(0xf507e13)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_getHistoryMessagesCount, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_getHistoryMessagesCount, layer)
	}
}

// Decode <--
func (m *TLMessageGetHistoryMessagesCount) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xf507e13: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetPeerUserMessageId) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x73aeb71f: func() error {
			x.PutClazzID(0x73aeb71f)

			x.PutInt64(m.UserId)
			x.PutInt64(m.PeerUserId)
			x.PutInt32(m.MsgId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_getPeerUserMessageId, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_getPeerUserMessageId, layer)
	}
}

// Decode <--
func (m *TLMessageGetPeerUserMessageId) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x73aeb71f: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerUserId, err = d.Int64()
			m.MsgId, err = d.Int32()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetPeerUserMessage) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x63129212: func() error {
			x.PutClazzID(0x63129212)

			x.PutInt64(m.UserId)
			x.PutInt64(m.PeerUserId)
			x.PutInt32(m.MsgId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_getPeerUserMessage, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_getPeerUserMessage, layer)
	}
}

// Decode <--
func (m *TLMessageGetPeerUserMessage) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x63129212: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerUserId, err = d.Int64()
			m.MsgId, err = d.Int32()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageSearchByMediaType) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x111c2943: func() error {
			x.PutClazzID(0x111c2943)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			x.PutInt32(m.MediaType)
			x.PutInt32(m.Offset)
			x.PutInt32(m.Limit)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_searchByMediaType, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_searchByMediaType, layer)
	}
}

// Decode <--
func (m *TLMessageSearchByMediaType) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x111c2943: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.MediaType, err = d.Int32()
			m.Offset, err = d.Int32()
			m.Limit, err = d.Int32()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageSearch) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x6835b023: func() error {
			x.PutClazzID(0x6835b023)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			x.PutString(m.Q)
			x.PutInt32(m.Offset)
			x.PutInt32(m.Limit)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_search, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_search, layer)
	}
}

// Decode <--
func (m *TLMessageSearch) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x6835b023: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.Q, err = d.String()
			m.Offset, err = d.Int32()
			m.Limit, err = d.Int32()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageSearchGlobal) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xb3985dc5: func() error {
			x.PutClazzID(0xb3985dc5)

			x.PutInt64(m.UserId)
			x.PutString(m.Q)
			x.PutInt32(m.Offset)
			x.PutInt32(m.Limit)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_searchGlobal, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_searchGlobal, layer)
	}
}

// Decode <--
func (m *TLMessageSearchGlobal) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xb3985dc5: func() (err error) {
			m.UserId, err = d.Int64()
			m.Q, err = d.String()
			m.Offset, err = d.Int32()
			m.Limit, err = d.Int32()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageSearchByPinned) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x6e735b55: func() error {
			x.PutClazzID(0x6e735b55)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_searchByPinned, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_searchByPinned, layer)
	}
}

// Decode <--
func (m *TLMessageSearchByPinned) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x6e735b55: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetSearchCounter) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xe2cbbf46: func() error {
			x.PutClazzID(0xe2cbbf46)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			x.PutInt32(m.MediaType)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_getSearchCounter, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_getSearchCounter, layer)
	}
}

// Decode <--
func (m *TLMessageGetSearchCounter) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xe2cbbf46: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.MediaType, err = d.Int32()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageSearchV2) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa1c62b21: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_searchV2, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_searchV2, layer)
	}
}

// Decode <--
func (m *TLMessageSearchV2) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa1c62b21: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.Q, err = d.String()
			m.FromId, err = d.Int64()
			m.MinDate, err = d.Int32()
			m.MaxDate, err = d.Int32()
			m.OffsetId, err = d.Int32()
			m.AddOffset, err = d.Int32()
			m.Limit, err = d.Int32()
			m.MaxId, err = d.Int32()
			m.MinId, err = d.Int32()
			m.Hash, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetLastTwoPinnedMessageId) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xaf9a082b: func() error {
			x.PutClazzID(0xaf9a082b)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_getLastTwoPinnedMessageId, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_getLastTwoPinnedMessageId, layer)
	}
}

// Decode <--
func (m *TLMessageGetLastTwoPinnedMessageId) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xaf9a082b: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageUpdatePinnedMessageId) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xf520edd0: func() error {
			x.PutClazzID(0xf520edd0)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			x.PutInt32(m.Id)
			_ = m.Pinned.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_updatePinnedMessageId, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_updatePinnedMessageId, layer)
	}
}

// Decode <--
func (m *TLMessageUpdatePinnedMessageId) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xf520edd0: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.Id, err = d.Int32()

			// m5 := &tg.Bool{}
			// _ = m5.Decode(d)
			// m.Pinned = m5
			m.Pinned, _ = tg.DecodeBoolClazz(d)

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetPinnedMessageIdList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xda01d0dd: func() error {
			x.PutClazzID(0xda01d0dd)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_getPinnedMessageIdList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_getPinnedMessageIdList, layer)
	}
}

// Decode <--
func (m *TLMessageGetPinnedMessageIdList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xda01d0dd: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageUnPinAllMessages) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xea0a2a73: func() error {
			x.PutClazzID(0xea0a2a73)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_unPinAllMessages, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_unPinAllMessages, layer)
	}
}

// Decode <--
func (m *TLMessageUnPinAllMessages) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xea0a2a73: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetUnreadMentions) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x6fe184b4: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_getUnreadMentions, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_getUnreadMentions, layer)
	}
}

// Decode <--
func (m *TLMessageGetUnreadMentions) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x6fe184b4: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.OffsetId, err = d.Int32()
			m.AddOffset, err = d.Int32()
			m.Limit, err = d.Int32()
			m.MinId, err = d.Int32()
			m.MaxInt, err = d.Int32()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMessageGetUnreadMentionsCount) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xb5412049: func() error {
			x.PutClazzID(0xb5412049)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_message_getUnreadMentionsCount, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_message_getUnreadMentionsCount, layer)
	}
}

// Decode <--
func (m *TLMessageGetUnreadMentionsCount) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xb5412049: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
