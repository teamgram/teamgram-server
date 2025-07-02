/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package msg

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

// TLMsgPushUserMessage <--
type TLMsgPushUserMessage struct {
	ClazzID   uint32         `json:"_id"`
	UserId    int64          `json:"user_id"`
	AuthKeyId int64          `json:"auth_key_id"`
	PeerType  int32          `json:"peer_type"`
	PeerId    int64          `json:"peer_id"`
	PushType  int32          `json:"push_type"`
	Message   *OutboxMessage `json:"message"`
}

func (m *TLMsgPushUserMessage) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgPushUserMessage) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x35d0fa1a: func() error {
			x.PutClazzID(0x35d0fa1a)

			x.PutInt64(m.UserId)
			x.PutInt64(m.AuthKeyId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			x.PutInt32(m.PushType)
			_ = m.Message.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_msg_pushUserMessage, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_pushUserMessage, layer)
	}
}

// Decode <--
func (m *TLMsgPushUserMessage) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x35d0fa1a: func() (err error) {
			m.UserId, err = d.Int64()
			m.AuthKeyId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.PushType, err = d.Int32()

			m6 := &OutboxMessage{}
			_ = m6.Decode(d)
			m.Message = m6

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

// TLMsgReadMessageContents <--
type TLMsgReadMessageContents struct {
	ClazzID   uint32            `json:"_id"`
	UserId    int64             `json:"user_id"`
	AuthKeyId int64             `json:"auth_key_id"`
	PeerType  int32             `json:"peer_type"`
	PeerId    int64             `json:"peer_id"`
	Id        []*ContentMessage `json:"id"`
}

func (m *TLMsgReadMessageContents) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgReadMessageContents) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x282484d4: func() error {
			x.PutClazzID(0x282484d4)

			x.PutInt64(m.UserId)
			x.PutInt64(m.AuthKeyId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			_ = iface.EncodeObjectList(x, m.Id, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_msg_readMessageContents, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_readMessageContents, layer)
	}
}

// Decode <--
func (m *TLMsgReadMessageContents) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x282484d4: func() (err error) {
			m.UserId, err = d.Int64()
			m.AuthKeyId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			c5, err2 := d.ClazzID()
			if c5 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 5, c5)
				return err2
			}
			l5, err3 := d.Int()
			v5 := make([]*ContentMessage, l5)
			for i := 0; i < l5; i++ {
				vv := new(ContentMessage)
				err3 = vv.Decode(d)
				_ = err3
				v5[i] = vv
			}
			m.Id = v5

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

// TLMsgSendMessageV2 <--
type TLMsgSendMessageV2 struct {
	ClazzID   uint32           `json:"_id"`
	UserId    int64            `json:"user_id"`
	AuthKeyId int64            `json:"auth_key_id"`
	PeerType  int32            `json:"peer_type"`
	PeerId    int64            `json:"peer_id"`
	Message   []*OutboxMessage `json:"message"`
}

func (m *TLMsgSendMessageV2) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgSendMessageV2) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xf4ca7cc4: func() error {
			x.PutClazzID(0xf4ca7cc4)

			x.PutInt64(m.UserId)
			x.PutInt64(m.AuthKeyId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			_ = iface.EncodeObjectList(x, m.Message, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_msg_sendMessageV2, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_sendMessageV2, layer)
	}
}

// Decode <--
func (m *TLMsgSendMessageV2) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xf4ca7cc4: func() (err error) {
			m.UserId, err = d.Int64()
			m.AuthKeyId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			c5, err2 := d.ClazzID()
			if c5 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 5, c5)
				return err2
			}
			l5, err3 := d.Int()
			v5 := make([]*OutboxMessage, l5)
			for i := 0; i < l5; i++ {
				vv := new(OutboxMessage)
				err3 = vv.Decode(d)
				_ = err3
				v5[i] = vv
			}
			m.Message = v5

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

// TLMsgEditMessageV2 <--
type TLMsgEditMessageV2 struct {
	ClazzID    uint32         `json:"_id"`
	UserId     int64          `json:"user_id"`
	AuthKeyId  int64          `json:"auth_key_id"`
	PeerType   int32          `json:"peer_type"`
	PeerId     int64          `json:"peer_id"`
	EditType   int32          `json:"edit_type"`
	NewMessage *OutboxMessage `json:"new_message"`
	DstMessage *tg.MessageBox `json:"dst_message"`
}

func (m *TLMsgEditMessageV2) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgEditMessageV2) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x69fe5fe1: func() error {
			x.PutClazzID(0x69fe5fe1)

			x.PutInt64(m.UserId)
			x.PutInt64(m.AuthKeyId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			x.PutInt32(m.EditType)
			_ = m.NewMessage.Encode(x, layer)
			_ = m.DstMessage.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_msg_editMessageV2, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_editMessageV2, layer)
	}
}

// Decode <--
func (m *TLMsgEditMessageV2) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x69fe5fe1: func() (err error) {
			m.UserId, err = d.Int64()
			m.AuthKeyId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.EditType, err = d.Int32()

			m6 := &OutboxMessage{}
			_ = m6.Decode(d)
			m.NewMessage = m6

			m7 := &tg.MessageBox{}
			_ = m7.Decode(d)
			m.DstMessage = m7

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

// TLMsgDeleteMessages <--
type TLMsgDeleteMessages struct {
	ClazzID   uint32  `json:"_id"`
	UserId    int64   `json:"user_id"`
	AuthKeyId int64   `json:"auth_key_id"`
	PeerType  int32   `json:"peer_type"`
	PeerId    int64   `json:"peer_id"`
	Revoke    bool    `json:"revoke"`
	Id        []int32 `json:"id"`
}

func (m *TLMsgDeleteMessages) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgDeleteMessages) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x21e80a1d: func() error {
			x.PutClazzID(0x21e80a1d)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.Revoke == true {
					flags |= 1 << 1
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.UserId)
			x.PutInt64(m.AuthKeyId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			iface.EncodeInt32List(x, m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_msg_deleteMessages, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_deleteMessages, layer)
	}
}

// Decode <--
func (m *TLMsgDeleteMessages) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x21e80a1d: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.UserId, err = d.Int64()
			m.AuthKeyId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			if (flags & (1 << 1)) != 0 {
				m.Revoke = true
			}

			m.Id, err = iface.DecodeInt32List(d)

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

// TLMsgDeleteHistory <--
type TLMsgDeleteHistory struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	PeerType  int32  `json:"peer_type"`
	PeerId    int64  `json:"peer_id"`
	JustClear bool   `json:"just_clear"`
	Revoke    bool   `json:"revoke"`
	MaxId     int32  `json:"max_id"`
}

func (m *TLMsgDeleteHistory) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgDeleteHistory) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x75c0e8ca: func() error {
			x.PutClazzID(0x75c0e8ca)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.JustClear == true {
					flags |= 1 << 0
				}
				if m.Revoke == true {
					flags |= 1 << 1
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.UserId)
			x.PutInt64(m.AuthKeyId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			x.PutInt32(m.MaxId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_msg_deleteHistory, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_deleteHistory, layer)
	}
}

// Decode <--
func (m *TLMsgDeleteHistory) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x75c0e8ca: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.UserId, err = d.Int64()
			m.AuthKeyId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			if (flags & (1 << 0)) != 0 {
				m.JustClear = true
			}
			if (flags & (1 << 1)) != 0 {
				m.Revoke = true
			}
			m.MaxId, err = d.Int32()

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

// TLMsgDeletePhoneCallHistory <--
type TLMsgDeletePhoneCallHistory struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	Revoke    bool   `json:"revoke"`
}

func (m *TLMsgDeletePhoneCallHistory) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgDeletePhoneCallHistory) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x26b7a13e: func() error {
			x.PutClazzID(0x26b7a13e)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.Revoke == true {
					flags |= 1 << 1
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.UserId)
			x.PutInt64(m.AuthKeyId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_msg_deletePhoneCallHistory, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_deletePhoneCallHistory, layer)
	}
}

// Decode <--
func (m *TLMsgDeletePhoneCallHistory) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x26b7a13e: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.UserId, err = d.Int64()
			m.AuthKeyId, err = d.Int64()
			if (flags & (1 << 1)) != 0 {
				m.Revoke = true
			}

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

// TLMsgDeleteChatHistory <--
type TLMsgDeleteChatHistory struct {
	ClazzID      uint32 `json:"_id"`
	ChatId       int64  `json:"chat_id"`
	DeleteUserId int64  `json:"delete_user_id"`
}

func (m *TLMsgDeleteChatHistory) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgDeleteChatHistory) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xef1f62db: func() error {
			x.PutClazzID(0xef1f62db)

			x.PutInt64(m.ChatId)
			x.PutInt64(m.DeleteUserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_msg_deleteChatHistory, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_deleteChatHistory, layer)
	}
}

// Decode <--
func (m *TLMsgDeleteChatHistory) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xef1f62db: func() (err error) {
			m.ChatId, err = d.Int64()
			m.DeleteUserId, err = d.Int64()

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

// TLMsgReadHistory <--
type TLMsgReadHistory struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	PeerType  int32  `json:"peer_type"`
	PeerId    int64  `json:"peer_id"`
	MaxId     int32  `json:"max_id"`
}

func (m *TLMsgReadHistory) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgReadHistory) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x5a0f6e12: func() error {
			x.PutClazzID(0x5a0f6e12)

			x.PutInt64(m.UserId)
			x.PutInt64(m.AuthKeyId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			x.PutInt32(m.MaxId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_msg_readHistory, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_readHistory, layer)
	}
}

// Decode <--
func (m *TLMsgReadHistory) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x5a0f6e12: func() (err error) {
			m.UserId, err = d.Int64()
			m.AuthKeyId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.MaxId, err = d.Int32()

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

// TLMsgReadHistoryV2 <--
type TLMsgReadHistoryV2 struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	PeerType  int32  `json:"peer_type"`
	PeerId    int64  `json:"peer_id"`
	MaxId     int32  `json:"max_id"`
}

func (m *TLMsgReadHistoryV2) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgReadHistoryV2) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xfb9b206: func() error {
			x.PutClazzID(0xfb9b206)

			x.PutInt64(m.UserId)
			x.PutInt64(m.AuthKeyId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			x.PutInt32(m.MaxId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_msg_readHistoryV2, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_readHistoryV2, layer)
	}
}

// Decode <--
func (m *TLMsgReadHistoryV2) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xfb9b206: func() (err error) {
			m.UserId, err = d.Int64()
			m.AuthKeyId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.MaxId, err = d.Int32()

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

// TLMsgUpdatePinnedMessage <--
type TLMsgUpdatePinnedMessage struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	Silent    bool   `json:"silent"`
	Unpin     bool   `json:"unpin"`
	PmOneside bool   `json:"pm_oneside"`
	PeerType  int32  `json:"peer_type"`
	PeerId    int64  `json:"peer_id"`
	Id        int32  `json:"id"`
}

func (m *TLMsgUpdatePinnedMessage) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgUpdatePinnedMessage) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xe5ae51a9: func() error {
			x.PutClazzID(0xe5ae51a9)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.Silent == true {
					flags |= 1 << 0
				}
				if m.Unpin == true {
					flags |= 1 << 1
				}
				if m.PmOneside == true {
					flags |= 1 << 2
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.UserId)
			x.PutInt64(m.AuthKeyId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			x.PutInt32(m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_msg_updatePinnedMessage, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_updatePinnedMessage, layer)
	}
}

// Decode <--
func (m *TLMsgUpdatePinnedMessage) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xe5ae51a9: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.UserId, err = d.Int64()
			m.AuthKeyId, err = d.Int64()
			if (flags & (1 << 0)) != 0 {
				m.Silent = true
			}
			if (flags & (1 << 1)) != 0 {
				m.Unpin = true
			}
			if (flags & (1 << 2)) != 0 {
				m.PmOneside = true
			}
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
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

// TLMsgUnpinAllMessages <--
type TLMsgUnpinAllMessages struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	PeerType  int32  `json:"peer_type"`
	PeerId    int64  `json:"peer_id"`
}

func (m *TLMsgUnpinAllMessages) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgUnpinAllMessages) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xb8865f25: func() error {
			x.PutClazzID(0xb8865f25)

			x.PutInt64(m.UserId)
			x.PutInt64(m.AuthKeyId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_msg_unpinAllMessages, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_unpinAllMessages, layer)
	}
}

// Decode <--
func (m *TLMsgUnpinAllMessages) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xb8865f25: func() (err error) {
			m.UserId, err = d.Int64()
			m.AuthKeyId, err = d.Int64()
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

// ----------------------------------------------------------------------------
// rpc

type RPCMsg interface {
	MsgPushUserMessage(ctx context.Context, in *TLMsgPushUserMessage) (*tg.Bool, error)
	MsgReadMessageContents(ctx context.Context, in *TLMsgReadMessageContents) (*tg.MessagesAffectedMessages, error)
	MsgSendMessageV2(ctx context.Context, in *TLMsgSendMessageV2) (*tg.Updates, error)
	MsgEditMessageV2(ctx context.Context, in *TLMsgEditMessageV2) (*tg.Updates, error)
	MsgDeleteMessages(ctx context.Context, in *TLMsgDeleteMessages) (*tg.MessagesAffectedMessages, error)
	MsgDeleteHistory(ctx context.Context, in *TLMsgDeleteHistory) (*tg.MessagesAffectedHistory, error)
	MsgDeletePhoneCallHistory(ctx context.Context, in *TLMsgDeletePhoneCallHistory) (*tg.MessagesAffectedFoundMessages, error)
	MsgDeleteChatHistory(ctx context.Context, in *TLMsgDeleteChatHistory) (*tg.Bool, error)
	MsgReadHistory(ctx context.Context, in *TLMsgReadHistory) (*tg.MessagesAffectedMessages, error)
	MsgReadHistoryV2(ctx context.Context, in *TLMsgReadHistoryV2) (*tg.MessagesAffectedMessages, error)
	MsgUpdatePinnedMessage(ctx context.Context, in *TLMsgUpdatePinnedMessage) (*tg.Updates, error)
	MsgUnpinAllMessages(ctx context.Context, in *TLMsgUnpinAllMessages) (*tg.MessagesAffectedHistory, error)
}
