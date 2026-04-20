/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package msg

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

// TLMsgPushUserMessage <--
type TLMsgPushUserMessage struct {
	ClazzID   uint32             `json:"_id"`
	UserId    int64              `json:"user_id"`
	AuthKeyId int64              `json:"auth_key_id"`
	PeerType  int32              `json:"peer_type"`
	PeerId    int64              `json:"peer_id"`
	PushType  int32              `json:"push_type"`
	Message   OutboxMessageClazz `json:"message"`
}

func (m *TLMsgPushUserMessage) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_msg_pushUserMessage, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgPushUserMessage) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_msg_pushUserMessage, int(layer)); clazzId {
	case 0x35d0fa1a:
		x.PutClazzID(0x35d0fa1a)

		x.PutInt64(m.UserId)
		x.PutInt64(m.AuthKeyId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt32(m.PushType)
		_ = m.Message.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_pushUserMessage, layer)
	}
}

// Decode <--
func (m *TLMsgPushUserMessage) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x35d0fa1a:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PushType, err = d.Int32()
		if err != nil {
			return err
		}

		m.Message, err = DecodeOutboxMessageClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMsgReadMessageContents <--
type TLMsgReadMessageContents struct {
	ClazzID   uint32                `json:"_id"`
	UserId    int64                 `json:"user_id"`
	AuthKeyId int64                 `json:"auth_key_id"`
	PeerType  int32                 `json:"peer_type"`
	PeerId    int64                 `json:"peer_id"`
	Id        []ContentMessageClazz `json:"id"`
}

func (m *TLMsgReadMessageContents) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_msg_readMessageContents, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgReadMessageContents) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_msg_readMessageContents, int(layer)); clazzId {
	case 0x282484d4:
		x.PutClazzID(0x282484d4)

		x.PutInt64(m.UserId)
		x.PutInt64(m.AuthKeyId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		if err := iface.EncodeObjectList(x, m.Id, layer); err != nil {
			return err
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_readMessageContents, layer)
	}
}

// Decode <--
func (m *TLMsgReadMessageContents) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x282484d4:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}
		c5, err2 := d.ClazzID()
		if err2 != nil {
			return err2
		}
		if c5 != iface.ClazzID_vector {
			return fmt.Errorf("invalid ClazzID_vector, c%d: %d", 5, c5)
		}
		l5, err3 := d.Int()
		if err3 != nil {
			return err3
		}
		v5 := make([]ContentMessageClazz, l5)
		for i := 0; i < l5; i++ {
			v5[i], err3 = DecodeContentMessageClazz(d)
			if err3 != nil {
				return err3
			}
		}
		m.Id = v5

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMsgSendMessageV2 <--
type TLMsgSendMessageV2 struct {
	ClazzID   uint32               `json:"_id"`
	UserId    int64                `json:"user_id"`
	AuthKeyId int64                `json:"auth_key_id"`
	PeerType  int32                `json:"peer_type"`
	PeerId    int64                `json:"peer_id"`
	Message   []OutboxMessageClazz `json:"message"`
}

func (m *TLMsgSendMessageV2) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_msg_sendMessageV2, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgSendMessageV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_msg_sendMessageV2, int(layer)); clazzId {
	case 0xf4ca7cc4:
		x.PutClazzID(0xf4ca7cc4)

		x.PutInt64(m.UserId)
		x.PutInt64(m.AuthKeyId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		if err := iface.EncodeObjectList(x, m.Message, layer); err != nil {
			return err
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_sendMessageV2, layer)
	}
}

// Decode <--
func (m *TLMsgSendMessageV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xf4ca7cc4:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}
		c5, err2 := d.ClazzID()
		if err2 != nil {
			return err2
		}
		if c5 != iface.ClazzID_vector {
			return fmt.Errorf("invalid ClazzID_vector, c%d: %d", 5, c5)
		}
		l5, err3 := d.Int()
		if err3 != nil {
			return err3
		}
		v5 := make([]OutboxMessageClazz, l5)
		for i := 0; i < l5; i++ {
			v5[i], err3 = DecodeOutboxMessageClazz(d)
			if err3 != nil {
				return err3
			}
		}
		m.Message = v5

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMsgEditMessageV2 <--
type TLMsgEditMessageV2 struct {
	ClazzID    uint32             `json:"_id"`
	UserId     int64              `json:"user_id"`
	AuthKeyId  int64              `json:"auth_key_id"`
	PeerType   int32              `json:"peer_type"`
	PeerId     int64              `json:"peer_id"`
	EditType   int32              `json:"edit_type"`
	NewMessage OutboxMessageClazz `json:"new_message"`
	DstMessage tg.MessageBoxClazz `json:"dst_message"`
}

func (m *TLMsgEditMessageV2) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_msg_editMessageV2, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgEditMessageV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_msg_editMessageV2, int(layer)); clazzId {
	case 0x69fe5fe1:
		x.PutClazzID(0x69fe5fe1)

		x.PutInt64(m.UserId)
		x.PutInt64(m.AuthKeyId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt32(m.EditType)
		_ = m.NewMessage.Encode(x, layer)
		_ = m.DstMessage.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_editMessageV2, layer)
	}
}

// Decode <--
func (m *TLMsgEditMessageV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x69fe5fe1:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}
		m.EditType, err = d.Int32()
		if err != nil {
			return err
		}

		m.NewMessage, err = DecodeOutboxMessageClazz(d)
		if err != nil {
			return err
		}

		m.DstMessage, err = tg.DecodeMessageBoxClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_msg_deleteMessages, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgDeleteMessages) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_msg_deleteMessages, int(layer)); clazzId {
	case 0x21e80a1d:
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
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_deleteMessages, layer)
	}
}

// Decode <--
func (m *TLMsgDeleteMessages) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x21e80a1d:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}
		if (flags & (1 << 1)) != 0 {
			m.Revoke = true
		}

		m.Id, err = iface.DecodeInt32List(d)
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_msg_deleteHistory, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgDeleteHistory) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_msg_deleteHistory, int(layer)); clazzId {
	case 0x75c0e8ca:
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
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_deleteHistory, layer)
	}
}

// Decode <--
func (m *TLMsgDeleteHistory) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x75c0e8ca:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}
		if (flags & (1 << 0)) != 0 {
			m.JustClear = true
		}
		if (flags & (1 << 1)) != 0 {
			m.Revoke = true
		}
		m.MaxId, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_msg_deletePhoneCallHistory, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgDeletePhoneCallHistory) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_msg_deletePhoneCallHistory, int(layer)); clazzId {
	case 0x26b7a13e:
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
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_deletePhoneCallHistory, layer)
	}
}

// Decode <--
func (m *TLMsgDeletePhoneCallHistory) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x26b7a13e:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		if (flags & (1 << 1)) != 0 {
			m.Revoke = true
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_msg_deleteChatHistory, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgDeleteChatHistory) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_msg_deleteChatHistory, int(layer)); clazzId {
	case 0xef1f62db:
		x.PutClazzID(0xef1f62db)

		x.PutInt64(m.ChatId)
		x.PutInt64(m.DeleteUserId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_deleteChatHistory, layer)
	}
}

// Decode <--
func (m *TLMsgDeleteChatHistory) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xef1f62db:
		m.ChatId, err = d.Int64()
		if err != nil {
			return err
		}
		m.DeleteUserId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_msg_readHistory, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgReadHistory) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_msg_readHistory, int(layer)); clazzId {
	case 0x5a0f6e12:
		x.PutClazzID(0x5a0f6e12)

		x.PutInt64(m.UserId)
		x.PutInt64(m.AuthKeyId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt32(m.MaxId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_readHistory, layer)
	}
}

// Decode <--
func (m *TLMsgReadHistory) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x5a0f6e12:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}
		m.MaxId, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_msg_readHistoryV2, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgReadHistoryV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_msg_readHistoryV2, int(layer)); clazzId {
	case 0xfb9b206:
		x.PutClazzID(0xfb9b206)

		x.PutInt64(m.UserId)
		x.PutInt64(m.AuthKeyId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt32(m.MaxId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_readHistoryV2, layer)
	}
}

// Decode <--
func (m *TLMsgReadHistoryV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xfb9b206:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}
		m.MaxId, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_msg_updatePinnedMessage, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgUpdatePinnedMessage) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_msg_updatePinnedMessage, int(layer)); clazzId {
	case 0xe5ae51a9:
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
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_updatePinnedMessage, layer)
	}
}

// Decode <--
func (m *TLMsgUpdatePinnedMessage) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xe5ae51a9:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
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
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}
		m.Id, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_msg_unpinAllMessages, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMsgUnpinAllMessages) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_msg_unpinAllMessages, int(layer)); clazzId {
	case 0xb8865f25:
		x.PutClazzID(0xb8865f25)

		x.PutInt64(m.UserId)
		x.PutInt64(m.AuthKeyId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_unpinAllMessages, layer)
	}
}

// Decode <--
func (m *TLMsgUnpinAllMessages) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xb8865f25:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
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
