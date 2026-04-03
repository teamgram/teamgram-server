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

// ContentMessageClazz <--
//   - TL_ContentMessage
type ContentMessageClazz = *TLContentMessage

func DecodeContentMessageClazz(d *bin.Decoder) (ContentMessageClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x8d64b133:
		x := &TLContentMessage{ClazzID: id, ClazzName2: ClazzName_contentMessage}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeContentMessage - unexpected clazzId: %d", id)
	}

}

// TLContentMessage <--
type TLContentMessage struct {
	ClazzID         uint32 `json:"_id"`
	ClazzName2      string `json:"_name"`
	Id              int32  `json:"id"`
	DialogMessageId int64  `json:"dialog_message_id"`
	Mentioned       bool   `json:"mentioned"`
	MediaUnread     bool   `json:"media_unread"`
	Reaction        bool   `json:"reaction"`
	SendUserId      int64  `json:"send_user_id"`
}

func MakeTLContentMessage(m *TLContentMessage) *TLContentMessage {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_contentMessage

	return m
}

func (m *TLContentMessage) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLContentMessage) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("contentMessage", m)
}

// ContentMessageClazzName <--
func (m *TLContentMessage) ContentMessageClazzName() string {
	return ClazzName_contentMessage
}

// ClazzName <--
func (m *TLContentMessage) ClazzName() string {
	return m.ClazzName2
}

// ToContentMessage <--
func (m *TLContentMessage) ToContentMessage() *ContentMessage {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLContentMessage) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_contentMessage, int(layer)); clazzId {
	case 0x8d64b133:
		size := 4
		size += 4
		size += 4
		size += 8
		size += 8

		return size
	default:
		return 0
	}
}

func (m *TLContentMessage) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_contentMessage, int(layer)); clazzId {
	case 0x8d64b133:

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_contentMessage, layer)
	}
}

// Encode <--
func (m *TLContentMessage) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_contentMessage, int(layer)); clazzId {
	case 0x8d64b133:
		x.PutClazzID(0x8d64b133)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.Mentioned == true {
				flags |= 1 << 0
			}
			if m.MediaUnread == true {
				flags |= 1 << 1
			}
			if m.Reaction == true {
				flags |= 1 << 2
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt32(m.Id)
		x.PutInt64(m.DialogMessageId)
		x.PutInt64(m.SendUserId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_contentMessage, layer)
	}
}

// Decode <--
func (m *TLContentMessage) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x8d64b133:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.Id, err = d.Int32()
		if err != nil {
			return err
		}
		m.DialogMessageId, err = d.Int64()
		if err != nil {
			return err
		}
		if (flags & (1 << 0)) != 0 {
			m.Mentioned = true
		}
		if (flags & (1 << 1)) != 0 {
			m.MediaUnread = true
		}
		if (flags & (1 << 2)) != 0 {
			m.Reaction = true
		}
		m.SendUserId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// ContentMessage <--
type ContentMessage = TLContentMessage

// OutboxMessageClazz <--
//   - TL_OutboxMessage
type OutboxMessageClazz = *TLOutboxMessage

func DecodeOutboxMessageClazz(d *bin.Decoder) (OutboxMessageClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x539524b1:
		x := &TLOutboxMessage{ClazzID: id, ClazzName2: ClazzName_outboxMessage}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeOutboxMessage - unexpected clazzId: %d", id)
	}

}

// TLOutboxMessage <--
type TLOutboxMessage struct {
	ClazzID      uint32          `json:"_id"`
	ClazzName2   string          `json:"_name"`
	NoWebpage    bool            `json:"no_webpage"`
	Background   bool            `json:"background"`
	RandomId     int64           `json:"random_id"`
	Message      tg.MessageClazz `json:"message"`
	ScheduleDate *int32          `json:"schedule_date"`
}

func MakeTLOutboxMessage(m *TLOutboxMessage) *TLOutboxMessage {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_outboxMessage

	return m
}

func (m *TLOutboxMessage) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLOutboxMessage) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("outboxMessage", m)
}

// OutboxMessageClazzName <--
func (m *TLOutboxMessage) OutboxMessageClazzName() string {
	return ClazzName_outboxMessage
}

// ClazzName <--
func (m *TLOutboxMessage) ClazzName() string {
	return m.ClazzName2
}

// ToOutboxMessage <--
func (m *TLOutboxMessage) ToOutboxMessage() *OutboxMessage {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLOutboxMessage) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_outboxMessage, int(layer)); clazzId {
	case 0x539524b1:
		size := 4
		size += 4
		size += 8
		size += iface.CalcObjectSize(m.Message, layer)
		if m.ScheduleDate != nil {
			size += 4
		}

		return size
	default:
		return 0
	}
}

func (m *TLOutboxMessage) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_outboxMessage, int(layer)); clazzId {
	case 0x539524b1:
		if err := iface.ValidateRequiredObject("message", m.Message); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_outboxMessage, layer)
	}
}

// Encode <--
func (m *TLOutboxMessage) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_outboxMessage, int(layer)); clazzId {
	case 0x539524b1:
		x.PutClazzID(0x539524b1)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.NoWebpage == true {
				flags |= 1 << 0
			}
			if m.Background == true {
				flags |= 1 << 1
			}

			if m.ScheduleDate != nil {
				flags |= 1 << 2
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.RandomId)
		_ = m.Message.Encode(x, layer)
		if m.ScheduleDate != nil {
			x.PutInt32(*m.ScheduleDate)
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_outboxMessage, layer)
	}
}

// Decode <--
func (m *TLOutboxMessage) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x539524b1:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		if (flags & (1 << 0)) != 0 {
			m.NoWebpage = true
		}
		if (flags & (1 << 1)) != 0 {
			m.Background = true
		}
		m.RandomId, err = d.Int64()
		if err != nil {
			return err
		}

		m.Message, err = tg.DecodeMessageClazz(d)
		if err != nil {
			return err
		}

		if (flags & (1 << 2)) != 0 {
			m.ScheduleDate = new(int32)
			*m.ScheduleDate, err = d.Int32()
			if err != nil {
				return err
			}
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// OutboxMessage <--
type OutboxMessage = TLOutboxMessage

// SenderClazz <--
//   - TL_Sender
type SenderClazz = *TLSender

func DecodeSenderClazz(d *bin.Decoder) (SenderClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x5a3864ba:
		x := &TLSender{ClazzID: id, ClazzName2: ClazzName_sender}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeSender - unexpected clazzId: %d", id)
	}

}

// TLSender <--
type TLSender struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	UserId     int64  `json:"user_id"`
	Type       int32  `json:"type"`
	AuthKeyId  int64  `json:"auth_key_id"`
}

func MakeTLSender(m *TLSender) *TLSender {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_sender

	return m
}

func (m *TLSender) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLSender) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("sender", m)
}

// SenderClazzName <--
func (m *TLSender) SenderClazzName() string {
	return ClazzName_sender
}

// ClazzName <--
func (m *TLSender) ClazzName() string {
	return m.ClazzName2
}

// ToSender <--
func (m *TLSender) ToSender() *Sender {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLSender) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_sender, int(layer)); clazzId {
	case 0x5a3864ba:
		size := 4
		size += 8
		size += 4
		size += 8

		return size
	default:
		return 0
	}
}

func (m *TLSender) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_sender, int(layer)); clazzId {
	case 0x5a3864ba:

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_sender, layer)
	}
}

// Encode <--
func (m *TLSender) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_sender, int(layer)); clazzId {
	case 0x5a3864ba:
		x.PutClazzID(0x5a3864ba)

		x.PutInt64(m.UserId)
		x.PutInt32(m.Type)
		x.PutInt64(m.AuthKeyId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_sender, layer)
	}
}

// Decode <--
func (m *TLSender) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x5a3864ba:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.Type, err = d.Int32()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// Sender <--
type Sender = TLSender
