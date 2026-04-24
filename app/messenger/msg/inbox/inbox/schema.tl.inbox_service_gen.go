/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package inbox

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

// TLInboxEditUserMessageToInbox <--
type TLInboxEditUserMessageToInbox struct {
	ClazzID    uint32          `json:"_id"`
	FromId     int64           `json:"from_id"`
	PeerUserId int64           `json:"peer_user_id"`
	Message    tg.MessageClazz `json:"message"`
}

func (m *TLInboxEditUserMessageToInbox) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_inbox_editUserMessageToInbox, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxEditUserMessageToInbox) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inbox_editUserMessageToInbox, int(layer)); clazzId {
	case 0x5cfb37a8:
		x.PutClazzID(0x5cfb37a8)

		x.PutInt64(m.FromId)
		x.PutInt64(m.PeerUserId)
		if m.Message == nil {
			return fmt.Errorf("unable to encode inbox_editUserMessageToInbox#0x5cfb37a8: field message is nil")
		}
		if err := m.Message.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode inbox_editUserMessageToInbox#0x5cfb37a8: field message: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode inbox_editUserMessageToInbox: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLInboxEditUserMessageToInbox) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_editUserMessageToInbox: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x5cfb37a8:
		m.FromId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_editUserMessageToInbox#0x5cfb37a8: field from_id: %w", err)
		}
		m.PeerUserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_editUserMessageToInbox#0x5cfb37a8: field peer_user_id: %w", err)
		}

		m.Message, err = tg.DecodeMessageClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode inbox_editUserMessageToInbox#0x5cfb37a8: field message: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode inbox_editUserMessageToInbox: invalid constructor %x", m.ClazzID)
	}
}

// TLInboxEditChatMessageToInbox <--
type TLInboxEditChatMessageToInbox struct {
	ClazzID    uint32          `json:"_id"`
	FromId     int64           `json:"from_id"`
	PeerChatId int64           `json:"peer_chat_id"`
	Message    tg.MessageClazz `json:"message"`
}

func (m *TLInboxEditChatMessageToInbox) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_inbox_editChatMessageToInbox, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxEditChatMessageToInbox) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inbox_editChatMessageToInbox, int(layer)); clazzId {
	case 0x79107a0f:
		x.PutClazzID(0x79107a0f)

		x.PutInt64(m.FromId)
		x.PutInt64(m.PeerChatId)
		if m.Message == nil {
			return fmt.Errorf("unable to encode inbox_editChatMessageToInbox#0x79107a0f: field message is nil")
		}
		if err := m.Message.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode inbox_editChatMessageToInbox#0x79107a0f: field message: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode inbox_editChatMessageToInbox: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLInboxEditChatMessageToInbox) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_editChatMessageToInbox: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x79107a0f:
		m.FromId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_editChatMessageToInbox#0x79107a0f: field from_id: %w", err)
		}
		m.PeerChatId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_editChatMessageToInbox#0x79107a0f: field peer_chat_id: %w", err)
		}

		m.Message, err = tg.DecodeMessageClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode inbox_editChatMessageToInbox#0x79107a0f: field message: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode inbox_editChatMessageToInbox: invalid constructor %x", m.ClazzID)
	}
}

// TLInboxDeleteMessagesToInbox <--
type TLInboxDeleteMessagesToInbox struct {
	ClazzID  uint32  `json:"_id"`
	FromId   int64   `json:"from_id"`
	PeerType int32   `json:"peer_type"`
	PeerId   int64   `json:"peer_id"`
	Id       []int64 `json:"id"`
}

func (m *TLInboxDeleteMessagesToInbox) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_inbox_deleteMessagesToInbox, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxDeleteMessagesToInbox) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inbox_deleteMessagesToInbox, int(layer)); clazzId {
	case 0x851c6e34:
		x.PutClazzID(0x851c6e34)

		x.PutInt64(m.FromId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		iface.EncodeInt64List(x, m.Id)

		return nil
	default:
		return fmt.Errorf("unable to encode inbox_deleteMessagesToInbox: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLInboxDeleteMessagesToInbox) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_deleteMessagesToInbox: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x851c6e34:
		m.FromId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_deleteMessagesToInbox#0x851c6e34: field from_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_deleteMessagesToInbox#0x851c6e34: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_deleteMessagesToInbox#0x851c6e34: field peer_id: %w", err)
		}

		m.Id, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode inbox_deleteMessagesToInbox#0x851c6e34: field id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode inbox_deleteMessagesToInbox: invalid constructor %x", m.ClazzID)
	}
}

// TLInboxDeleteUserHistoryToInbox <--
type TLInboxDeleteUserHistoryToInbox struct {
	ClazzID    uint32 `json:"_id"`
	FromId     int64  `json:"from_id"`
	PeerUserId int64  `json:"peer_user_id"`
	JustClear  bool   `json:"just_clear"`
	MaxId      int32  `json:"max_id"`
}

func (m *TLInboxDeleteUserHistoryToInbox) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_inbox_deleteUserHistoryToInbox, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxDeleteUserHistoryToInbox) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inbox_deleteUserHistoryToInbox, int(layer)); clazzId {
	case 0x140a8158:
		x.PutClazzID(0x140a8158)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.JustClear == true {
				flags |= 1 << 1
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.FromId)
		x.PutInt64(m.PeerUserId)
		x.PutInt32(m.MaxId)

		return nil
	default:
		return fmt.Errorf("unable to encode inbox_deleteUserHistoryToInbox: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLInboxDeleteUserHistoryToInbox) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_deleteUserHistoryToInbox: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x140a8158:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_deleteUserHistoryToInbox: field flags: %w", err)
		}
		_ = flags
		m.FromId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_deleteUserHistoryToInbox#0x140a8158: field from_id: %w", err)
		}
		m.PeerUserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_deleteUserHistoryToInbox#0x140a8158: field peer_user_id: %w", err)
		}
		if (flags & (1 << 1)) != 0 {
			m.JustClear = true
		}
		m.MaxId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_deleteUserHistoryToInbox#0x140a8158: field max_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode inbox_deleteUserHistoryToInbox: invalid constructor %x", m.ClazzID)
	}
}

// TLInboxDeleteChatHistoryToInbox <--
type TLInboxDeleteChatHistoryToInbox struct {
	ClazzID    uint32 `json:"_id"`
	FromId     int64  `json:"from_id"`
	PeerChatId int64  `json:"peer_chat_id"`
	MaxId      int32  `json:"max_id"`
}

func (m *TLInboxDeleteChatHistoryToInbox) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_inbox_deleteChatHistoryToInbox, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxDeleteChatHistoryToInbox) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inbox_deleteChatHistoryToInbox, int(layer)); clazzId {
	case 0xd8aaa602:
		x.PutClazzID(0xd8aaa602)

		x.PutInt64(m.FromId)
		x.PutInt64(m.PeerChatId)
		x.PutInt32(m.MaxId)

		return nil
	default:
		return fmt.Errorf("unable to encode inbox_deleteChatHistoryToInbox: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLInboxDeleteChatHistoryToInbox) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_deleteChatHistoryToInbox: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xd8aaa602:
		m.FromId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_deleteChatHistoryToInbox#0xd8aaa602: field from_id: %w", err)
		}
		m.PeerChatId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_deleteChatHistoryToInbox#0xd8aaa602: field peer_chat_id: %w", err)
		}
		m.MaxId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_deleteChatHistoryToInbox#0xd8aaa602: field max_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode inbox_deleteChatHistoryToInbox: invalid constructor %x", m.ClazzID)
	}
}

// TLInboxReadUserMediaUnreadToInbox <--
type TLInboxReadUserMediaUnreadToInbox struct {
	ClazzID    uint32                `json:"_id"`
	FromId     int64                 `json:"from_id"`
	PeerUserId int64                 `json:"peer_user_id"`
	Id         []InboxMessageIdClazz `json:"id"`
}

func (m *TLInboxReadUserMediaUnreadToInbox) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_inbox_readUserMediaUnreadToInbox, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxReadUserMediaUnreadToInbox) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inbox_readUserMediaUnreadToInbox, int(layer)); clazzId {
	case 0x15c1034b:
		x.PutClazzID(0x15c1034b)

		x.PutInt64(m.FromId)
		x.PutInt64(m.PeerUserId)

		if err := iface.EncodeObjectList(x, m.Id, layer); err != nil {
			return fmt.Errorf("unable to encode inbox_readUserMediaUnreadToInbox#0x15c1034b: field id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode inbox_readUserMediaUnreadToInbox: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLInboxReadUserMediaUnreadToInbox) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readUserMediaUnreadToInbox: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x15c1034b:
		m.FromId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readUserMediaUnreadToInbox#0x15c1034b: field from_id: %w", err)
		}
		m.PeerUserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readUserMediaUnreadToInbox#0x15c1034b: field peer_user_id: %w", err)
		}
		c3, err2 := d.ClazzID()
		if err2 != nil {
			return fmt.Errorf("unable to decode inbox_readUserMediaUnreadToInbox#0x15c1034b: field id: %w", err2)
		}
		if c3 != iface.ClazzID_vector {
			return fmt.Errorf("unable to decode inbox_readUserMediaUnreadToInbox#0x15c1034b: field id: invalid vector constructor %x", c3)
		}
		l3, err3 := d.Int()
		if err3 != nil {
			return fmt.Errorf("unable to decode inbox_readUserMediaUnreadToInbox#0x15c1034b: field id: %w", err3)
		}
		v3 := make([]InboxMessageIdClazz, l3)
		for i := 0; i < l3; i++ {
			v3[i], err3 = DecodeInboxMessageIdClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode inbox_readUserMediaUnreadToInbox#0x15c1034b: field id: %w", err3)
			}
		}
		m.Id = v3

		return nil
	default:
		return fmt.Errorf("unable to decode inbox_readUserMediaUnreadToInbox: invalid constructor %x", m.ClazzID)
	}
}

// TLInboxReadChatMediaUnreadToInbox <--
type TLInboxReadChatMediaUnreadToInbox struct {
	ClazzID    uint32                `json:"_id"`
	FromId     int64                 `json:"from_id"`
	PeerChatId int64                 `json:"peer_chat_id"`
	Id         []InboxMessageIdClazz `json:"id"`
}

func (m *TLInboxReadChatMediaUnreadToInbox) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_inbox_readChatMediaUnreadToInbox, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxReadChatMediaUnreadToInbox) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inbox_readChatMediaUnreadToInbox, int(layer)); clazzId {
	case 0x55415dd4:
		x.PutClazzID(0x55415dd4)

		x.PutInt64(m.FromId)
		x.PutInt64(m.PeerChatId)

		if err := iface.EncodeObjectList(x, m.Id, layer); err != nil {
			return fmt.Errorf("unable to encode inbox_readChatMediaUnreadToInbox#0x55415dd4: field id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode inbox_readChatMediaUnreadToInbox: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLInboxReadChatMediaUnreadToInbox) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readChatMediaUnreadToInbox: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x55415dd4:
		m.FromId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readChatMediaUnreadToInbox#0x55415dd4: field from_id: %w", err)
		}
		m.PeerChatId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readChatMediaUnreadToInbox#0x55415dd4: field peer_chat_id: %w", err)
		}
		c3, err2 := d.ClazzID()
		if err2 != nil {
			return fmt.Errorf("unable to decode inbox_readChatMediaUnreadToInbox#0x55415dd4: field id: %w", err2)
		}
		if c3 != iface.ClazzID_vector {
			return fmt.Errorf("unable to decode inbox_readChatMediaUnreadToInbox#0x55415dd4: field id: invalid vector constructor %x", c3)
		}
		l3, err3 := d.Int()
		if err3 != nil {
			return fmt.Errorf("unable to decode inbox_readChatMediaUnreadToInbox#0x55415dd4: field id: %w", err3)
		}
		v3 := make([]InboxMessageIdClazz, l3)
		for i := 0; i < l3; i++ {
			v3[i], err3 = DecodeInboxMessageIdClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode inbox_readChatMediaUnreadToInbox#0x55415dd4: field id: %w", err3)
			}
		}
		m.Id = v3

		return nil
	default:
		return fmt.Errorf("unable to decode inbox_readChatMediaUnreadToInbox: invalid constructor %x", m.ClazzID)
	}
}

// TLInboxUpdateHistoryReaded <--
type TLInboxUpdateHistoryReaded struct {
	ClazzID  uint32 `json:"_id"`
	FromId   int64  `json:"from_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
	MaxId    int32  `json:"max_id"`
	Sender   int64  `json:"sender"`
}

func (m *TLInboxUpdateHistoryReaded) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_inbox_updateHistoryReaded, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxUpdateHistoryReaded) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inbox_updateHistoryReaded, int(layer)); clazzId {
	case 0xc3c84ce0:
		x.PutClazzID(0xc3c84ce0)

		x.PutInt64(m.FromId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt32(m.MaxId)
		x.PutInt64(m.Sender)

		return nil
	default:
		return fmt.Errorf("unable to encode inbox_updateHistoryReaded: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLInboxUpdateHistoryReaded) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updateHistoryReaded: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xc3c84ce0:
		m.FromId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updateHistoryReaded#0xc3c84ce0: field from_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updateHistoryReaded#0xc3c84ce0: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updateHistoryReaded#0xc3c84ce0: field peer_id: %w", err)
		}
		m.MaxId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updateHistoryReaded#0xc3c84ce0: field max_id: %w", err)
		}
		m.Sender, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updateHistoryReaded#0xc3c84ce0: field sender: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode inbox_updateHistoryReaded: invalid constructor %x", m.ClazzID)
	}
}

// TLInboxUpdatePinnedMessage <--
type TLInboxUpdatePinnedMessage struct {
	ClazzID         uint32 `json:"_id"`
	UserId          int64  `json:"user_id"`
	Unpin           bool   `json:"unpin"`
	PeerType        int32  `json:"peer_type"`
	PeerId          int64  `json:"peer_id"`
	Id              int32  `json:"id"`
	DialogMessageId int64  `json:"dialog_message_id"`
}

func (m *TLInboxUpdatePinnedMessage) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_inbox_updatePinnedMessage, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxUpdatePinnedMessage) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inbox_updatePinnedMessage, int(layer)); clazzId {
	case 0xa96c2af4:
		x.PutClazzID(0xa96c2af4)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.Unpin == true {
				flags |= 1 << 1
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt32(m.Id)
		x.PutInt64(m.DialogMessageId)

		return nil
	default:
		return fmt.Errorf("unable to encode inbox_updatePinnedMessage: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLInboxUpdatePinnedMessage) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updatePinnedMessage: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xa96c2af4:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updatePinnedMessage: field flags: %w", err)
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updatePinnedMessage#0xa96c2af4: field user_id: %w", err)
		}
		if (flags & (1 << 1)) != 0 {
			m.Unpin = true
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updatePinnedMessage#0xa96c2af4: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updatePinnedMessage#0xa96c2af4: field peer_id: %w", err)
		}
		m.Id, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updatePinnedMessage#0xa96c2af4: field id: %w", err)
		}
		m.DialogMessageId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updatePinnedMessage#0xa96c2af4: field dialog_message_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode inbox_updatePinnedMessage: invalid constructor %x", m.ClazzID)
	}
}

// TLInboxUnpinAllMessages <--
type TLInboxUnpinAllMessages struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	PeerType  int32  `json:"peer_type"`
	PeerId    int64  `json:"peer_id"`
}

func (m *TLInboxUnpinAllMessages) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_inbox_unpinAllMessages, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxUnpinAllMessages) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inbox_unpinAllMessages, int(layer)); clazzId {
	case 0x231ca261:
		x.PutClazzID(0x231ca261)

		x.PutInt64(m.UserId)
		x.PutInt64(m.AuthKeyId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		return fmt.Errorf("unable to encode inbox_unpinAllMessages: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLInboxUnpinAllMessages) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_unpinAllMessages: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x231ca261:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_unpinAllMessages#0x231ca261: field user_id: %w", err)
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_unpinAllMessages#0x231ca261: field auth_key_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_unpinAllMessages#0x231ca261: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_unpinAllMessages#0x231ca261: field peer_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode inbox_unpinAllMessages: invalid constructor %x", m.ClazzID)
	}
}

// TLInboxSendUserMessageToInboxV2 <--
type TLInboxSendUserMessageToInboxV2 struct {
	ClazzID        uint32               `json:"_id"`
	UserId         int64                `json:"user_id"`
	Out            bool                 `json:"out"`
	FromId         int64                `json:"from_id"`
	FromAuthKeyId  int64                `json:"from_auth_keyId"`
	PeerType       int32                `json:"peer_type"`
	PeerId         int64                `json:"peer_id"`
	BoxList        []tg.MessageBoxClazz `json:"box_list"`
	Users          []tg.UserClazz       `json:"users"`
	Chats          []tg.ChatClazz       `json:"chats"`
	Layer          *int32               `json:"layer"`
	ServerId       *string              `json:"server_id"`
	SessionId      *int64               `json:"session_id"`
	ClientReqMsgId *int64               `json:"client_req_msg_id"`
	AuthKeyId      *int64               `json:"auth_key_id"`
}

func (m *TLInboxSendUserMessageToInboxV2) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_inbox_sendUserMessageToInboxV2, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxSendUserMessageToInboxV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inbox_sendUserMessageToInboxV2, int(layer)); clazzId {
	case 0x5bd7522:
		x.PutClazzID(0x5bd7522)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.Out == true {
				flags |= 1 << 0
			}

			if m.Users != nil {
				flags |= 1 << 1
			}
			if m.Chats != nil {
				flags |= 1 << 2
			}
			if m.Layer != nil {
				flags |= 1 << 3
			}
			if m.ServerId != nil {
				flags |= 1 << 4
			}
			if m.SessionId != nil {
				flags |= 1 << 5
			}
			if m.ClientReqMsgId != nil {
				flags |= 1 << 6
			}
			if m.AuthKeyId != nil {
				flags |= 1 << 7
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.UserId)
		x.PutInt64(m.FromId)
		x.PutInt64(m.FromAuthKeyId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		if err := iface.EncodeObjectList(x, m.BoxList, layer); err != nil {
			return fmt.Errorf("unable to encode inbox_sendUserMessageToInboxV2#0x5bd7522: field box_list: %w", err)
		}

		if m.Users != nil {
			if err := iface.EncodeObjectList(x, m.Users, layer); err != nil {
				return fmt.Errorf("unable to encode inbox_sendUserMessageToInboxV2#0x5bd7522: field users: %w", err)
			}
		}
		if m.Chats != nil {
			if err := iface.EncodeObjectList(x, m.Chats, layer); err != nil {
				return fmt.Errorf("unable to encode inbox_sendUserMessageToInboxV2#0x5bd7522: field chats: %w", err)
			}
		}
		if m.Layer != nil {
			x.PutInt32(*m.Layer)
		}

		if m.ServerId != nil {
			x.PutString(*m.ServerId)
		}

		if m.SessionId != nil {
			x.PutInt64(*m.SessionId)
		}

		if m.ClientReqMsgId != nil {
			x.PutInt64(*m.ClientReqMsgId)
		}

		if m.AuthKeyId != nil {
			x.PutInt64(*m.AuthKeyId)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode inbox_sendUserMessageToInboxV2: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLInboxSendUserMessageToInboxV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x5bd7522:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2: field flags: %w", err)
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field user_id: %w", err)
		}
		if (flags & (1 << 0)) != 0 {
			m.Out = true
		}
		m.FromId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field from_id: %w", err)
		}
		m.FromAuthKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field from_auth_keyId: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field peer_id: %w", err)
		}
		c8, err2 := d.ClazzID()
		if err2 != nil {
			return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field box_list: %w", err2)
		}
		if c8 != iface.ClazzID_vector {
			return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field box_list: invalid vector constructor %x", c8)
		}
		l8, err3 := d.Int()
		if err3 != nil {
			return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field box_list: %w", err3)
		}
		v8 := make([]tg.MessageBoxClazz, l8)
		for i := 0; i < l8; i++ {
			v8[i], err3 = tg.DecodeMessageBoxClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field box_list: %w", err3)
			}
		}
		m.BoxList = v8

		if (flags & (1 << 1)) != 0 {
			c9, err2 := d.ClazzID()
			if err2 != nil {
				return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field users: %w", err2)
			}
			if c9 != iface.ClazzID_vector {
				return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field users: invalid vector constructor %x", c9)
			}
			l9, err3 := d.Int()
			if err3 != nil {
				return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field users: %w", err3)
			}
			v9 := make([]tg.UserClazz, l9)
			for i := 0; i < l9; i++ {
				v9[i], err3 = tg.DecodeUserClazz(d)
				if err3 != nil {
					return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field users: %w", err3)
				}
			}
			m.Users = v9
		}
		if (flags & (1 << 2)) != 0 {
			c10, err2 := d.ClazzID()
			if err2 != nil {
				return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field chats: %w", err2)
			}
			if c10 != iface.ClazzID_vector {
				return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field chats: invalid vector constructor %x", c10)
			}
			l10, err3 := d.Int()
			if err3 != nil {
				return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field chats: %w", err3)
			}
			v10 := make([]tg.ChatClazz, l10)
			for i := 0; i < l10; i++ {
				v10[i], err3 = tg.DecodeChatClazz(d)
				if err3 != nil {
					return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field chats: %w", err3)
				}
			}
			m.Chats = v10
		}
		if (flags & (1 << 3)) != 0 {
			m.Layer = new(int32)
			*m.Layer, err = d.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field layer: %w", err)
			}
		}
		if (flags & (1 << 4)) != 0 {
			m.ServerId = new(string)
			*m.ServerId, err = d.String()
			if err != nil {
				return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field server_id: %w", err)
			}
		}

		if (flags & (1 << 5)) != 0 {
			m.SessionId = new(int64)
			*m.SessionId, err = d.Int64()
			if err != nil {
				return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field session_id: %w", err)
			}
		}

		if (flags & (1 << 6)) != 0 {
			m.ClientReqMsgId = new(int64)
			*m.ClientReqMsgId, err = d.Int64()
			if err != nil {
				return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field client_req_msg_id: %w", err)
			}
		}

		if (flags & (1 << 7)) != 0 {
			m.AuthKeyId = new(int64)
			*m.AuthKeyId, err = d.Int64()
			if err != nil {
				return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2#0x5bd7522: field auth_key_id: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to decode inbox_sendUserMessageToInboxV2: invalid constructor %x", m.ClazzID)
	}
}

// TLInboxEditMessageToInboxV2 <--
type TLInboxEditMessageToInboxV2 struct {
	ClazzID       uint32             `json:"_id"`
	UserId        int64              `json:"user_id"`
	Out           bool               `json:"out"`
	FromId        int64              `json:"from_id"`
	FromAuthKeyId int64              `json:"from_auth_keyId"`
	PeerType      int32              `json:"peer_type"`
	PeerId        int64              `json:"peer_id"`
	NewMessage    tg.MessageBoxClazz `json:"new_message"`
	DstMessage    tg.MessageBoxClazz `json:"dst_message"`
	Users         []tg.UserClazz     `json:"users"`
	Chats         []tg.ChatClazz     `json:"chats"`
}

func (m *TLInboxEditMessageToInboxV2) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_inbox_editMessageToInboxV2, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxEditMessageToInboxV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inbox_editMessageToInboxV2, int(layer)); clazzId {
	case 0xdabb9e69:
		x.PutClazzID(0xdabb9e69)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.Out == true {
				flags |= 1 << 0
			}

			if m.DstMessage != nil {
				flags |= 1 << 1
			}
			if m.Users != nil {
				flags |= 1 << 2
			}
			if m.Chats != nil {
				flags |= 1 << 3
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.UserId)
		x.PutInt64(m.FromId)
		x.PutInt64(m.FromAuthKeyId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		if m.NewMessage == nil {
			return fmt.Errorf("unable to encode inbox_editMessageToInboxV2#0xdabb9e69: field new_message is nil")
		}
		if err := m.NewMessage.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode inbox_editMessageToInboxV2#0xdabb9e69: field new_message: %w", err)
		}
		if m.DstMessage != nil {
			if err := m.DstMessage.Encode(x, layer); err != nil {
				return fmt.Errorf("unable to encode inbox_editMessageToInboxV2#0xdabb9e69: field dst_message: %w", err)
			}
		}

		if m.Users != nil {
			if err := iface.EncodeObjectList(x, m.Users, layer); err != nil {
				return fmt.Errorf("unable to encode inbox_editMessageToInboxV2#0xdabb9e69: field users: %w", err)
			}
		}
		if m.Chats != nil {
			if err := iface.EncodeObjectList(x, m.Chats, layer); err != nil {
				return fmt.Errorf("unable to encode inbox_editMessageToInboxV2#0xdabb9e69: field chats: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to encode inbox_editMessageToInboxV2: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLInboxEditMessageToInboxV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_editMessageToInboxV2: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xdabb9e69:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_editMessageToInboxV2: field flags: %w", err)
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_editMessageToInboxV2#0xdabb9e69: field user_id: %w", err)
		}
		if (flags & (1 << 0)) != 0 {
			m.Out = true
		}
		m.FromId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_editMessageToInboxV2#0xdabb9e69: field from_id: %w", err)
		}
		m.FromAuthKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_editMessageToInboxV2#0xdabb9e69: field from_auth_keyId: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_editMessageToInboxV2#0xdabb9e69: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_editMessageToInboxV2#0xdabb9e69: field peer_id: %w", err)
		}

		m.NewMessage, err = tg.DecodeMessageBoxClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode inbox_editMessageToInboxV2#0xdabb9e69: field new_message: %w", err)
		}

		if (flags & (1 << 1)) != 0 {
			m.DstMessage, err = tg.DecodeMessageBoxClazz(d)
			if err != nil {
				return fmt.Errorf("unable to decode inbox_editMessageToInboxV2#0xdabb9e69: field dst_message: %w", err)
			}
		}
		if (flags & (1 << 2)) != 0 {
			c10, err2 := d.ClazzID()
			if err2 != nil {
				return fmt.Errorf("unable to decode inbox_editMessageToInboxV2#0xdabb9e69: field users: %w", err2)
			}
			if c10 != iface.ClazzID_vector {
				return fmt.Errorf("unable to decode inbox_editMessageToInboxV2#0xdabb9e69: field users: invalid vector constructor %x", c10)
			}
			l10, err3 := d.Int()
			if err3 != nil {
				return fmt.Errorf("unable to decode inbox_editMessageToInboxV2#0xdabb9e69: field users: %w", err3)
			}
			v10 := make([]tg.UserClazz, l10)
			for i := 0; i < l10; i++ {
				v10[i], err3 = tg.DecodeUserClazz(d)
				if err3 != nil {
					return fmt.Errorf("unable to decode inbox_editMessageToInboxV2#0xdabb9e69: field users: %w", err3)
				}
			}
			m.Users = v10
		}
		if (flags & (1 << 3)) != 0 {
			c11, err2 := d.ClazzID()
			if err2 != nil {
				return fmt.Errorf("unable to decode inbox_editMessageToInboxV2#0xdabb9e69: field chats: %w", err2)
			}
			if c11 != iface.ClazzID_vector {
				return fmt.Errorf("unable to decode inbox_editMessageToInboxV2#0xdabb9e69: field chats: invalid vector constructor %x", c11)
			}
			l11, err3 := d.Int()
			if err3 != nil {
				return fmt.Errorf("unable to decode inbox_editMessageToInboxV2#0xdabb9e69: field chats: %w", err3)
			}
			v11 := make([]tg.ChatClazz, l11)
			for i := 0; i < l11; i++ {
				v11[i], err3 = tg.DecodeChatClazz(d)
				if err3 != nil {
					return fmt.Errorf("unable to decode inbox_editMessageToInboxV2#0xdabb9e69: field chats: %w", err3)
				}
			}
			m.Chats = v11
		}

		return nil
	default:
		return fmt.Errorf("unable to decode inbox_editMessageToInboxV2: invalid constructor %x", m.ClazzID)
	}
}

// TLInboxReadInboxHistory <--
type TLInboxReadInboxHistory struct {
	ClazzID        uint32  `json:"_id"`
	UserId         int64   `json:"user_id"`
	AuthKeyId      int64   `json:"auth_key_id"`
	PeerType       int32   `json:"peer_type"`
	PeerId         int64   `json:"peer_id"`
	Pts            int32   `json:"pts"`
	PtsCount       int32   `json:"pts_count"`
	UnreadCount    int32   `json:"unread_count"`
	ReadInboxMaxId int32   `json:"read_inbox_max_id"`
	MaxId          int32   `json:"max_id"`
	Layer          *int32  `json:"layer"`
	ServerId       *string `json:"server_id"`
	SessionId      *int64  `json:"session_id"`
	ClientReqMsgId *int64  `json:"client_req_msg_id"`
}

func (m *TLInboxReadInboxHistory) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_inbox_readInboxHistory, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxReadInboxHistory) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inbox_readInboxHistory, int(layer)); clazzId {
	case 0x1f73675:
		x.PutClazzID(0x1f73675)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.Layer != nil {
				flags |= 1 << 3
			}
			if m.ServerId != nil {
				flags |= 1 << 4
			}
			if m.SessionId != nil {
				flags |= 1 << 5
			}
			if m.ClientReqMsgId != nil {
				flags |= 1 << 6
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
		x.PutInt32(m.Pts)
		x.PutInt32(m.PtsCount)
		x.PutInt32(m.UnreadCount)
		x.PutInt32(m.ReadInboxMaxId)
		x.PutInt32(m.MaxId)
		if m.Layer != nil {
			x.PutInt32(*m.Layer)
		}

		if m.ServerId != nil {
			x.PutString(*m.ServerId)
		}

		if m.SessionId != nil {
			x.PutInt64(*m.SessionId)
		}

		if m.ClientReqMsgId != nil {
			x.PutInt64(*m.ClientReqMsgId)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode inbox_readInboxHistory: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLInboxReadInboxHistory) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readInboxHistory: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x1f73675:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readInboxHistory: field flags: %w", err)
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readInboxHistory#0x1f73675: field user_id: %w", err)
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readInboxHistory#0x1f73675: field auth_key_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readInboxHistory#0x1f73675: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readInboxHistory#0x1f73675: field peer_id: %w", err)
		}
		m.Pts, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readInboxHistory#0x1f73675: field pts: %w", err)
		}
		m.PtsCount, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readInboxHistory#0x1f73675: field pts_count: %w", err)
		}
		m.UnreadCount, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readInboxHistory#0x1f73675: field unread_count: %w", err)
		}
		m.ReadInboxMaxId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readInboxHistory#0x1f73675: field read_inbox_max_id: %w", err)
		}
		m.MaxId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readInboxHistory#0x1f73675: field max_id: %w", err)
		}
		if (flags & (1 << 3)) != 0 {
			m.Layer = new(int32)
			*m.Layer, err = d.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode inbox_readInboxHistory#0x1f73675: field layer: %w", err)
			}
		}
		if (flags & (1 << 4)) != 0 {
			m.ServerId = new(string)
			*m.ServerId, err = d.String()
			if err != nil {
				return fmt.Errorf("unable to decode inbox_readInboxHistory#0x1f73675: field server_id: %w", err)
			}
		}

		if (flags & (1 << 5)) != 0 {
			m.SessionId = new(int64)
			*m.SessionId, err = d.Int64()
			if err != nil {
				return fmt.Errorf("unable to decode inbox_readInboxHistory#0x1f73675: field session_id: %w", err)
			}
		}

		if (flags & (1 << 6)) != 0 {
			m.ClientReqMsgId = new(int64)
			*m.ClientReqMsgId, err = d.Int64()
			if err != nil {
				return fmt.Errorf("unable to decode inbox_readInboxHistory#0x1f73675: field client_req_msg_id: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to decode inbox_readInboxHistory: invalid constructor %x", m.ClazzID)
	}
}

// TLInboxReadOutboxHistory <--
type TLInboxReadOutboxHistory struct {
	ClazzID            uint32 `json:"_id"`
	UserId             int64  `json:"user_id"`
	PeerType           int32  `json:"peer_type"`
	PeerId             int64  `json:"peer_id"`
	MaxDialogMessageId int64  `json:"max_dialog_message_id"`
}

func (m *TLInboxReadOutboxHistory) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_inbox_readOutboxHistory, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxReadOutboxHistory) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inbox_readOutboxHistory, int(layer)); clazzId {
	case 0x1c7036ca:
		x.PutClazzID(0x1c7036ca)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt64(m.MaxDialogMessageId)

		return nil
	default:
		return fmt.Errorf("unable to encode inbox_readOutboxHistory: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLInboxReadOutboxHistory) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readOutboxHistory: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x1c7036ca:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readOutboxHistory#0x1c7036ca: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readOutboxHistory#0x1c7036ca: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readOutboxHistory#0x1c7036ca: field peer_id: %w", err)
		}
		m.MaxDialogMessageId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readOutboxHistory#0x1c7036ca: field max_dialog_message_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode inbox_readOutboxHistory: invalid constructor %x", m.ClazzID)
	}
}

// TLInboxReadMediaUnreadToInboxV2 <--
type TLInboxReadMediaUnreadToInboxV2 struct {
	ClazzID         uint32 `json:"_id"`
	UserId          int64  `json:"user_id"`
	PeerType        int32  `json:"peer_type"`
	PeerId          int64  `json:"peer_id"`
	DialogMessageId int64  `json:"dialog_message_id"`
}

func (m *TLInboxReadMediaUnreadToInboxV2) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_inbox_readMediaUnreadToInboxV2, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxReadMediaUnreadToInboxV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inbox_readMediaUnreadToInboxV2, int(layer)); clazzId {
	case 0xeac54342:
		x.PutClazzID(0xeac54342)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt64(m.DialogMessageId)

		return nil
	default:
		return fmt.Errorf("unable to encode inbox_readMediaUnreadToInboxV2: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLInboxReadMediaUnreadToInboxV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readMediaUnreadToInboxV2: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xeac54342:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readMediaUnreadToInboxV2#0xeac54342: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readMediaUnreadToInboxV2#0xeac54342: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readMediaUnreadToInboxV2#0xeac54342: field peer_id: %w", err)
		}
		m.DialogMessageId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_readMediaUnreadToInboxV2#0xeac54342: field dialog_message_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode inbox_readMediaUnreadToInboxV2: invalid constructor %x", m.ClazzID)
	}
}

// TLInboxUpdatePinnedMessageV2 <--
type TLInboxUpdatePinnedMessageV2 struct {
	ClazzID         uint32  `json:"_id"`
	UserId          int64   `json:"user_id"`
	Unpin           bool    `json:"unpin"`
	PeerType        int32   `json:"peer_type"`
	PeerId          int64   `json:"peer_id"`
	Id              int32   `json:"id"`
	DialogMessageId int64   `json:"dialog_message_id"`
	Layer           *int32  `json:"layer"`
	ServerId        *string `json:"server_id"`
	SessionId       *int64  `json:"session_id"`
	ClientReqMsgId  *int64  `json:"client_req_msg_id"`
}

func (m *TLInboxUpdatePinnedMessageV2) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_inbox_updatePinnedMessageV2, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxUpdatePinnedMessageV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inbox_updatePinnedMessageV2, int(layer)); clazzId {
	case 0x56b79e7c:
		x.PutClazzID(0x56b79e7c)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.Unpin == true {
				flags |= 1 << 1
			}

			if m.Layer != nil {
				flags |= 1 << 3
			}
			if m.ServerId != nil {
				flags |= 1 << 4
			}
			if m.SessionId != nil {
				flags |= 1 << 5
			}
			if m.ClientReqMsgId != nil {
				flags |= 1 << 6
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt32(m.Id)
		x.PutInt64(m.DialogMessageId)
		if m.Layer != nil {
			x.PutInt32(*m.Layer)
		}

		if m.ServerId != nil {
			x.PutString(*m.ServerId)
		}

		if m.SessionId != nil {
			x.PutInt64(*m.SessionId)
		}

		if m.ClientReqMsgId != nil {
			x.PutInt64(*m.ClientReqMsgId)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode inbox_updatePinnedMessageV2: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLInboxUpdatePinnedMessageV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updatePinnedMessageV2: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x56b79e7c:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updatePinnedMessageV2: field flags: %w", err)
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updatePinnedMessageV2#0x56b79e7c: field user_id: %w", err)
		}
		if (flags & (1 << 1)) != 0 {
			m.Unpin = true
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updatePinnedMessageV2#0x56b79e7c: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updatePinnedMessageV2#0x56b79e7c: field peer_id: %w", err)
		}
		m.Id, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updatePinnedMessageV2#0x56b79e7c: field id: %w", err)
		}
		m.DialogMessageId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode inbox_updatePinnedMessageV2#0x56b79e7c: field dialog_message_id: %w", err)
		}
		if (flags & (1 << 3)) != 0 {
			m.Layer = new(int32)
			*m.Layer, err = d.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode inbox_updatePinnedMessageV2#0x56b79e7c: field layer: %w", err)
			}
		}
		if (flags & (1 << 4)) != 0 {
			m.ServerId = new(string)
			*m.ServerId, err = d.String()
			if err != nil {
				return fmt.Errorf("unable to decode inbox_updatePinnedMessageV2#0x56b79e7c: field server_id: %w", err)
			}
		}

		if (flags & (1 << 5)) != 0 {
			m.SessionId = new(int64)
			*m.SessionId, err = d.Int64()
			if err != nil {
				return fmt.Errorf("unable to decode inbox_updatePinnedMessageV2#0x56b79e7c: field session_id: %w", err)
			}
		}

		if (flags & (1 << 6)) != 0 {
			m.ClientReqMsgId = new(int64)
			*m.ClientReqMsgId, err = d.Int64()
			if err != nil {
				return fmt.Errorf("unable to decode inbox_updatePinnedMessageV2#0x56b79e7c: field client_req_msg_id: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to decode inbox_updatePinnedMessageV2: invalid constructor %x", m.ClazzID)
	}
}

// Vector api result type
// ----------------------------------------------------------------------------
// VectorResList <--

// ----------------------------------------------------------------------------
// rpc

type RPCInbox interface {
	InboxEditUserMessageToInbox(ctx context.Context, in *TLInboxEditUserMessageToInbox) (*tg.Void, error)
	InboxEditChatMessageToInbox(ctx context.Context, in *TLInboxEditChatMessageToInbox) (*tg.Void, error)
	InboxDeleteMessagesToInbox(ctx context.Context, in *TLInboxDeleteMessagesToInbox) (*tg.Void, error)
	InboxDeleteUserHistoryToInbox(ctx context.Context, in *TLInboxDeleteUserHistoryToInbox) (*tg.Void, error)
	InboxDeleteChatHistoryToInbox(ctx context.Context, in *TLInboxDeleteChatHistoryToInbox) (*tg.Void, error)
	InboxReadUserMediaUnreadToInbox(ctx context.Context, in *TLInboxReadUserMediaUnreadToInbox) (*tg.Void, error)
	InboxReadChatMediaUnreadToInbox(ctx context.Context, in *TLInboxReadChatMediaUnreadToInbox) (*tg.Void, error)
	InboxUpdateHistoryReaded(ctx context.Context, in *TLInboxUpdateHistoryReaded) (*tg.Void, error)
	InboxUpdatePinnedMessage(ctx context.Context, in *TLInboxUpdatePinnedMessage) (*tg.Void, error)
	InboxUnpinAllMessages(ctx context.Context, in *TLInboxUnpinAllMessages) (*tg.Void, error)
	InboxSendUserMessageToInboxV2(ctx context.Context, in *TLInboxSendUserMessageToInboxV2) (*tg.Void, error)
	InboxEditMessageToInboxV2(ctx context.Context, in *TLInboxEditMessageToInboxV2) (*tg.Void, error)
	InboxReadInboxHistory(ctx context.Context, in *TLInboxReadInboxHistory) (*tg.Void, error)
	InboxReadOutboxHistory(ctx context.Context, in *TLInboxReadOutboxHistory) (*tg.Void, error)
	InboxReadMediaUnreadToInboxV2(ctx context.Context, in *TLInboxReadMediaUnreadToInboxV2) (*tg.Void, error)
	InboxUpdatePinnedMessageV2(ctx context.Context, in *TLInboxUpdatePinnedMessageV2) (*tg.Void, error)
}
