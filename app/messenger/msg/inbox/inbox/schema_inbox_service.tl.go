/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package inbox

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

// TLInboxEditUserMessageToInbox <--
type TLInboxEditUserMessageToInbox struct {
	ClazzID    uint32      `json:"_id"`
	FromId     int64       `json:"from_id"`
	PeerUserId int64       `json:"peer_user_id"`
	Message    *tg.Message `json:"message"`
}

func (m *TLInboxEditUserMessageToInbox) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxEditUserMessageToInbox) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x5cfb37a8: func() error {
			x.PutClazzID(0x5cfb37a8)

			x.PutInt64(m.FromId)
			x.PutInt64(m.PeerUserId)
			_ = m.Message.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inbox_editUserMessageToInbox, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inbox_editUserMessageToInbox, layer)
	}
}

// Decode <--
func (m *TLInboxEditUserMessageToInbox) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x5cfb37a8: func() (err error) {
			m.FromId, err = d.Int64()
			m.PeerUserId, err = d.Int64()

			m3 := &tg.Message{}
			_ = m3.Decode(d)
			m.Message = m3

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

// TLInboxEditChatMessageToInbox <--
type TLInboxEditChatMessageToInbox struct {
	ClazzID    uint32      `json:"_id"`
	FromId     int64       `json:"from_id"`
	PeerChatId int64       `json:"peer_chat_id"`
	Message    *tg.Message `json:"message"`
}

func (m *TLInboxEditChatMessageToInbox) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxEditChatMessageToInbox) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x79107a0f: func() error {
			x.PutClazzID(0x79107a0f)

			x.PutInt64(m.FromId)
			x.PutInt64(m.PeerChatId)
			_ = m.Message.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inbox_editChatMessageToInbox, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inbox_editChatMessageToInbox, layer)
	}
}

// Decode <--
func (m *TLInboxEditChatMessageToInbox) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x79107a0f: func() (err error) {
			m.FromId, err = d.Int64()
			m.PeerChatId, err = d.Int64()

			m3 := &tg.Message{}
			_ = m3.Decode(d)
			m.Message = m3

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

// TLInboxDeleteMessagesToInbox <--
type TLInboxDeleteMessagesToInbox struct {
	ClazzID  uint32  `json:"_id"`
	FromId   int64   `json:"from_id"`
	PeerType int32   `json:"peer_type"`
	PeerId   int64   `json:"peer_id"`
	Id       []int64 `json:"id"`
}

func (m *TLInboxDeleteMessagesToInbox) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxDeleteMessagesToInbox) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x851c6e34: func() error {
			x.PutClazzID(0x851c6e34)

			x.PutInt64(m.FromId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			iface.EncodeInt64List(x, m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inbox_deleteMessagesToInbox, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inbox_deleteMessagesToInbox, layer)
	}
}

// Decode <--
func (m *TLInboxDeleteMessagesToInbox) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x851c6e34: func() (err error) {
			m.FromId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()

			m.Id, err = iface.DecodeInt64List(d)

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

// TLInboxDeleteUserHistoryToInbox <--
type TLInboxDeleteUserHistoryToInbox struct {
	ClazzID    uint32 `json:"_id"`
	FromId     int64  `json:"from_id"`
	PeerUserId int64  `json:"peer_user_id"`
	JustClear  bool   `json:"just_clear"`
	MaxId      int32  `json:"max_id"`
}

func (m *TLInboxDeleteUserHistoryToInbox) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxDeleteUserHistoryToInbox) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x140a8158: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inbox_deleteUserHistoryToInbox, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inbox_deleteUserHistoryToInbox, layer)
	}
}

// Decode <--
func (m *TLInboxDeleteUserHistoryToInbox) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x140a8158: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.FromId, err = d.Int64()
			m.PeerUserId, err = d.Int64()
			if (flags & (1 << 1)) != 0 {
				m.JustClear = true
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

// TLInboxDeleteChatHistoryToInbox <--
type TLInboxDeleteChatHistoryToInbox struct {
	ClazzID    uint32 `json:"_id"`
	FromId     int64  `json:"from_id"`
	PeerChatId int64  `json:"peer_chat_id"`
	MaxId      int32  `json:"max_id"`
}

func (m *TLInboxDeleteChatHistoryToInbox) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxDeleteChatHistoryToInbox) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xd8aaa602: func() error {
			x.PutClazzID(0xd8aaa602)

			x.PutInt64(m.FromId)
			x.PutInt64(m.PeerChatId)
			x.PutInt32(m.MaxId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inbox_deleteChatHistoryToInbox, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inbox_deleteChatHistoryToInbox, layer)
	}
}

// Decode <--
func (m *TLInboxDeleteChatHistoryToInbox) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xd8aaa602: func() (err error) {
			m.FromId, err = d.Int64()
			m.PeerChatId, err = d.Int64()
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

// TLInboxReadUserMediaUnreadToInbox <--
type TLInboxReadUserMediaUnreadToInbox struct {
	ClazzID    uint32            `json:"_id"`
	FromId     int64             `json:"from_id"`
	PeerUserId int64             `json:"peer_user_id"`
	Id         []*InboxMessageId `json:"id"`
}

func (m *TLInboxReadUserMediaUnreadToInbox) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxReadUserMediaUnreadToInbox) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x15c1034b: func() error {
			x.PutClazzID(0x15c1034b)

			x.PutInt64(m.FromId)
			x.PutInt64(m.PeerUserId)

			_ = iface.EncodeObjectList(x, m.Id, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inbox_readUserMediaUnreadToInbox, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inbox_readUserMediaUnreadToInbox, layer)
	}
}

// Decode <--
func (m *TLInboxReadUserMediaUnreadToInbox) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x15c1034b: func() (err error) {
			m.FromId, err = d.Int64()
			m.PeerUserId, err = d.Int64()
			c3, err2 := d.ClazzID()
			if c3 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 3, c3)
				return err2
			}
			l3, err3 := d.Int()
			v3 := make([]*InboxMessageId, l3)
			for i := 0; i < l3; i++ {
				vv := new(InboxMessageId)
				err3 = vv.Decode(d)
				_ = err3
				v3[i] = vv
			}
			m.Id = v3

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

// TLInboxReadChatMediaUnreadToInbox <--
type TLInboxReadChatMediaUnreadToInbox struct {
	ClazzID    uint32            `json:"_id"`
	FromId     int64             `json:"from_id"`
	PeerChatId int64             `json:"peer_chat_id"`
	Id         []*InboxMessageId `json:"id"`
}

func (m *TLInboxReadChatMediaUnreadToInbox) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxReadChatMediaUnreadToInbox) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x55415dd4: func() error {
			x.PutClazzID(0x55415dd4)

			x.PutInt64(m.FromId)
			x.PutInt64(m.PeerChatId)

			_ = iface.EncodeObjectList(x, m.Id, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inbox_readChatMediaUnreadToInbox, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inbox_readChatMediaUnreadToInbox, layer)
	}
}

// Decode <--
func (m *TLInboxReadChatMediaUnreadToInbox) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x55415dd4: func() (err error) {
			m.FromId, err = d.Int64()
			m.PeerChatId, err = d.Int64()
			c3, err2 := d.ClazzID()
			if c3 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 3, c3)
				return err2
			}
			l3, err3 := d.Int()
			v3 := make([]*InboxMessageId, l3)
			for i := 0; i < l3; i++ {
				vv := new(InboxMessageId)
				err3 = vv.Decode(d)
				_ = err3
				v3[i] = vv
			}
			m.Id = v3

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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxUpdateHistoryReaded) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc3c84ce0: func() error {
			x.PutClazzID(0xc3c84ce0)

			x.PutInt64(m.FromId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			x.PutInt32(m.MaxId)
			x.PutInt64(m.Sender)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inbox_updateHistoryReaded, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inbox_updateHistoryReaded, layer)
	}
}

// Decode <--
func (m *TLInboxUpdateHistoryReaded) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xc3c84ce0: func() (err error) {
			m.FromId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.MaxId, err = d.Int32()
			m.Sender, err = d.Int64()

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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxUpdatePinnedMessage) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa96c2af4: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inbox_updatePinnedMessage, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inbox_updatePinnedMessage, layer)
	}
}

// Decode <--
func (m *TLInboxUpdatePinnedMessage) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa96c2af4: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.UserId, err = d.Int64()
			if (flags & (1 << 1)) != 0 {
				m.Unpin = true
			}
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.Id, err = d.Int32()
			m.DialogMessageId, err = d.Int64()

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

// TLInboxUnpinAllMessages <--
type TLInboxUnpinAllMessages struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	PeerType  int32  `json:"peer_type"`
	PeerId    int64  `json:"peer_id"`
}

func (m *TLInboxUnpinAllMessages) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxUnpinAllMessages) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x231ca261: func() error {
			x.PutClazzID(0x231ca261)

			x.PutInt64(m.UserId)
			x.PutInt64(m.AuthKeyId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inbox_unpinAllMessages, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inbox_unpinAllMessages, layer)
	}
}

// Decode <--
func (m *TLInboxUnpinAllMessages) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x231ca261: func() (err error) {
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

// TLInboxSendUserMessageToInboxV2 <--
type TLInboxSendUserMessageToInboxV2 struct {
	ClazzID        uint32           `json:"_id"`
	UserId         int64            `json:"user_id"`
	Out            bool             `json:"out"`
	FromId         int64            `json:"from_id"`
	FromAuthKeyId  int64            `json:"from_auth_keyId"`
	PeerType       int32            `json:"peer_type"`
	PeerId         int64            `json:"peer_id"`
	BoxList        []*tg.MessageBox `json:"box_list"`
	Users          []*tg.User       `json:"users"`
	Chats          []*tg.Chat       `json:"chats"`
	Layer          *int32           `json:"layer"`
	ServerId       *string          `json:"server_id"`
	SessionId      *int64           `json:"session_id"`
	ClientReqMsgId *int64           `json:"client_req_msg_id"`
	AuthKeyId      *int64           `json:"auth_key_id"`
}

func (m *TLInboxSendUserMessageToInboxV2) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxSendUserMessageToInboxV2) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x5bd7522: func() error {
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

			_ = iface.EncodeObjectList(x, m.BoxList, layer)

			if m.Users != nil {
				_ = iface.EncodeObjectList(x, m.Users, layer)
			}
			if m.Chats != nil {
				_ = iface.EncodeObjectList(x, m.Chats, layer)
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inbox_sendUserMessageToInboxV2, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inbox_sendUserMessageToInboxV2, layer)
	}
}

// Decode <--
func (m *TLInboxSendUserMessageToInboxV2) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x5bd7522: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.UserId, err = d.Int64()
			if (flags & (1 << 0)) != 0 {
				m.Out = true
			}
			m.FromId, err = d.Int64()
			m.FromAuthKeyId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			c8, err2 := d.ClazzID()
			if c8 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 8, c8)
				return err2
			}
			l8, err3 := d.Int()
			v8 := make([]*tg.MessageBox, l8)
			for i := 0; i < l8; i++ {
				vv := new(tg.MessageBox)
				err3 = vv.Decode(d)
				_ = err3
				v8[i] = vv
			}
			m.BoxList = v8

			if (flags & (1 << 1)) != 0 {
				c9, err2 := d.ClazzID()
				if c9 != iface.ClazzID_vector {
					// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 9, c9)
					return err2
				}
				l9, err3 := d.Int()
				v9 := make([]*tg.User, l9)
				for i := 0; i < l9; i++ {
					vv := new(tg.User)
					err3 = vv.Decode(d)
					_ = err3
					v9[i] = vv
				}
				m.Users = v9
			}
			if (flags & (1 << 2)) != 0 {
				c10, err2 := d.ClazzID()
				if c10 != iface.ClazzID_vector {
					// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 10, c10)
					return err2
				}
				l10, err3 := d.Int()
				v10 := make([]*tg.Chat, l10)
				for i := 0; i < l10; i++ {
					vv := new(tg.Chat)
					err3 = vv.Decode(d)
					_ = err3
					v10[i] = vv
				}
				m.Chats = v10
			}
			if (flags & (1 << 3)) != 0 {
				m.Layer = new(int32)
				*m.Layer, err = d.Int32()
			}
			if (flags & (1 << 4)) != 0 {
				m.ServerId = new(string)
				*m.ServerId, err = d.String()
			}

			if (flags & (1 << 5)) != 0 {
				m.SessionId = new(int64)
				*m.SessionId, err = d.Int64()
			}

			if (flags & (1 << 6)) != 0 {
				m.ClientReqMsgId = new(int64)
				*m.ClientReqMsgId, err = d.Int64()
			}

			if (flags & (1 << 7)) != 0 {
				m.AuthKeyId = new(int64)
				*m.AuthKeyId, err = d.Int64()
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

// TLInboxEditMessageToInboxV2 <--
type TLInboxEditMessageToInboxV2 struct {
	ClazzID       uint32         `json:"_id"`
	UserId        int64          `json:"user_id"`
	Out           bool           `json:"out"`
	FromId        int64          `json:"from_id"`
	FromAuthKeyId int64          `json:"from_auth_keyId"`
	PeerType      int32          `json:"peer_type"`
	PeerId        int64          `json:"peer_id"`
	NewMessage    *tg.MessageBox `json:"new_message"`
	DstMessage    *tg.MessageBox `json:"dst_message"`
	Users         []*tg.User     `json:"users"`
	Chats         []*tg.Chat     `json:"chats"`
}

func (m *TLInboxEditMessageToInboxV2) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxEditMessageToInboxV2) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xdabb9e69: func() error {
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
			_ = m.NewMessage.Encode(x, layer)
			if m.DstMessage != nil {
				_ = m.DstMessage.Encode(x, layer)
			}

			if m.Users != nil {
				_ = iface.EncodeObjectList(x, m.Users, layer)
			}
			if m.Chats != nil {
				_ = iface.EncodeObjectList(x, m.Chats, layer)
			}

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inbox_editMessageToInboxV2, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inbox_editMessageToInboxV2, layer)
	}
}

// Decode <--
func (m *TLInboxEditMessageToInboxV2) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xdabb9e69: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.UserId, err = d.Int64()
			if (flags & (1 << 0)) != 0 {
				m.Out = true
			}
			m.FromId, err = d.Int64()
			m.FromAuthKeyId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()

			m8 := &tg.MessageBox{}
			_ = m8.Decode(d)
			m.NewMessage = m8

			if (flags & (1 << 1)) != 0 {
				m9 := &tg.MessageBox{}
				_ = m9.Decode(d)
				m.DstMessage = m9
			}
			if (flags & (1 << 2)) != 0 {
				c10, err2 := d.ClazzID()
				if c10 != iface.ClazzID_vector {
					// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 10, c10)
					return err2
				}
				l10, err3 := d.Int()
				v10 := make([]*tg.User, l10)
				for i := 0; i < l10; i++ {
					vv := new(tg.User)
					err3 = vv.Decode(d)
					_ = err3
					v10[i] = vv
				}
				m.Users = v10
			}
			if (flags & (1 << 3)) != 0 {
				c11, err2 := d.ClazzID()
				if c11 != iface.ClazzID_vector {
					// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 11, c11)
					return err2
				}
				l11, err3 := d.Int()
				v11 := make([]*tg.Chat, l11)
				for i := 0; i < l11; i++ {
					vv := new(tg.Chat)
					err3 = vv.Decode(d)
					_ = err3
					v11[i] = vv
				}
				m.Chats = v11
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxReadInboxHistory) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x1f73675: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inbox_readInboxHistory, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inbox_readInboxHistory, layer)
	}
}

// Decode <--
func (m *TLInboxReadInboxHistory) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x1f73675: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.UserId, err = d.Int64()
			m.AuthKeyId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.Pts, err = d.Int32()
			m.PtsCount, err = d.Int32()
			m.UnreadCount, err = d.Int32()
			m.ReadInboxMaxId, err = d.Int32()
			m.MaxId, err = d.Int32()
			if (flags & (1 << 3)) != 0 {
				m.Layer = new(int32)
				*m.Layer, err = d.Int32()
			}
			if (flags & (1 << 4)) != 0 {
				m.ServerId = new(string)
				*m.ServerId, err = d.String()
			}

			if (flags & (1 << 5)) != 0 {
				m.SessionId = new(int64)
				*m.SessionId, err = d.Int64()
			}

			if (flags & (1 << 6)) != 0 {
				m.ClientReqMsgId = new(int64)
				*m.ClientReqMsgId, err = d.Int64()
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

// TLInboxReadOutboxHistory <--
type TLInboxReadOutboxHistory struct {
	ClazzID            uint32 `json:"_id"`
	UserId             int64  `json:"user_id"`
	PeerType           int32  `json:"peer_type"`
	PeerId             int64  `json:"peer_id"`
	MaxDialogMessageId int64  `json:"max_dialog_message_id"`
}

func (m *TLInboxReadOutboxHistory) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxReadOutboxHistory) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x1c7036ca: func() error {
			x.PutClazzID(0x1c7036ca)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			x.PutInt64(m.MaxDialogMessageId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inbox_readOutboxHistory, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inbox_readOutboxHistory, layer)
	}
}

// Decode <--
func (m *TLInboxReadOutboxHistory) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x1c7036ca: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.MaxDialogMessageId, err = d.Int64()

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

// TLInboxReadMediaUnreadToInboxV2 <--
type TLInboxReadMediaUnreadToInboxV2 struct {
	ClazzID         uint32 `json:"_id"`
	UserId          int64  `json:"user_id"`
	PeerType        int32  `json:"peer_type"`
	PeerId          int64  `json:"peer_id"`
	DialogMessageId int64  `json:"dialog_message_id"`
}

func (m *TLInboxReadMediaUnreadToInboxV2) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxReadMediaUnreadToInboxV2) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xeac54342: func() error {
			x.PutClazzID(0xeac54342)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			x.PutInt64(m.DialogMessageId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inbox_readMediaUnreadToInboxV2, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inbox_readMediaUnreadToInboxV2, layer)
	}
}

// Decode <--
func (m *TLInboxReadMediaUnreadToInboxV2) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xeac54342: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.DialogMessageId, err = d.Int64()

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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLInboxUpdatePinnedMessageV2) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x56b79e7c: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inbox_updatePinnedMessageV2, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inbox_updatePinnedMessageV2, layer)
	}
}

// Decode <--
func (m *TLInboxUpdatePinnedMessageV2) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x56b79e7c: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.UserId, err = d.Int64()
			if (flags & (1 << 1)) != 0 {
				m.Unpin = true
			}
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.Id, err = d.Int32()
			m.DialogMessageId, err = d.Int64()
			if (flags & (1 << 3)) != 0 {
				m.Layer = new(int32)
				*m.Layer, err = d.Int32()
			}
			if (flags & (1 << 4)) != 0 {
				m.ServerId = new(string)
				*m.ServerId, err = d.String()
			}

			if (flags & (1 << 5)) != 0 {
				m.SessionId = new(int64)
				*m.SessionId, err = d.Int64()
			}

			if (flags & (1 << 6)) != 0 {
				m.ClientReqMsgId = new(int64)
				*m.ClientReqMsgId, err = d.Int64()
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
