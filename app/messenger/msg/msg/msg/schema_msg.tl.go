/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package msg

import (
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
)

var _ iface.TLObject
var _ fmt.Stringer
var _ *tg.Bool
var _ bin.Fields

// ContentMessageClazz <--
//   - TL_ContentMessage
type ContentMessageClazz interface {
	iface.TLObject
	ContentMessageClazzName() string
}

func DecodeContentMessageClazz(d *bin.Decoder) (ContentMessageClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_contentMessage:
		x := &TLContentMessage{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeContentMessage - unexpected clazzId: %d", id)
	}
}

// TLContentMessage <--
type TLContentMessage struct {
	ClazzID         uint32 `json:"_id"`
	Id              int32  `json:"id"`
	DialogMessageId int64  `json:"dialog_message_id"`
	Mentioned       bool   `json:"mentioned"`
	MediaUnread     bool   `json:"media_unread"`
	Reaction        bool   `json:"reaction"`
	SendUserId      int64  `json:"send_user_id"`
}

// ContentMessageClazzName <--
func (m *TLContentMessage) ContentMessageClazzName() string {
	return ClazzName_contentMessage
}

// ClazzName <--
func (m *TLContentMessage) ClazzName() string {
	return ClazzName_contentMessage
}

// ToContentMessage <--
func (m *TLContentMessage) ToContentMessage() *ContentMessage {
	return MakeContentMessage(m)
}

// Encode <--
func (m *TLContentMessage) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x8d64b133: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_contentMessage, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_contentMessage, layer)
	}
}

// Decode <--
func (m *TLContentMessage) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x8d64b133: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.Id, err = d.Int32()
			m.DialogMessageId, err = d.Int64()
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

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// ContentMessage <--
type ContentMessage struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	ContentMessageClazz
}

// MakeContentMessage <--
func MakeContentMessage(c ContentMessageClazz) *ContentMessage {
	return &ContentMessage{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		ContentMessageClazz: c,
	}
}

// Encode <--
func (m *ContentMessage) Encode(x *bin.Encoder, layer int32) error {
	if m.ContentMessageClazz != nil {
		return m.ContentMessageClazz.Encode(x, layer)
	}

	return fmt.Errorf("ContentMessage - invalid Clazz")
}

// Decode <--
func (m *ContentMessage) Decode(d *bin.Decoder) (err error) {
	m.ContentMessageClazz, err = DecodeContentMessageClazz(d)
	return
}

// Match <--
func (m *ContentMessage) Match(f ...interface{}) {
	switch c := m.ContentMessageClazz.(type) {
	case *TLContentMessage:
		for _, v := range f {
			if f1, ok := v.(func(c *TLContentMessage) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToContentMessage <--
func (m *ContentMessage) ToContentMessage() (*TLContentMessage, bool) {
	if m.ContentMessageClazz == nil {
		return nil, false
	}

	if x, ok := m.ContentMessageClazz.(*TLContentMessage); ok {
		return x, true
	}

	return nil, false
}

// OutboxMessageClazz <--
//   - TL_OutboxMessage
type OutboxMessageClazz interface {
	iface.TLObject
	OutboxMessageClazzName() string
}

func DecodeOutboxMessageClazz(d *bin.Decoder) (OutboxMessageClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_outboxMessage:
		x := &TLOutboxMessage{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeOutboxMessage - unexpected clazzId: %d", id)
	}
}

// TLOutboxMessage <--
type TLOutboxMessage struct {
	ClazzID      uint32      `json:"_id"`
	NoWebpage    bool        `json:"no_webpage"`
	Background   bool        `json:"background"`
	RandomId     int64       `json:"random_id"`
	Message      *tg.Message `json:"message"`
	ScheduleDate *int32      `json:"schedule_date"`
}

// OutboxMessageClazzName <--
func (m *TLOutboxMessage) OutboxMessageClazzName() string {
	return ClazzName_outboxMessage
}

// ClazzName <--
func (m *TLOutboxMessage) ClazzName() string {
	return ClazzName_outboxMessage
}

// ToOutboxMessage <--
func (m *TLOutboxMessage) ToOutboxMessage() *OutboxMessage {
	return MakeOutboxMessage(m)
}

// Encode <--
func (m *TLOutboxMessage) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x539524b1: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_outboxMessage, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_outboxMessage, layer)
	}
}

// Decode <--
func (m *TLOutboxMessage) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x539524b1: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			if (flags & (1 << 0)) != 0 {
				m.NoWebpage = true
			}
			if (flags & (1 << 1)) != 0 {
				m.Background = true
			}
			m.RandomId, err = d.Int64()

			m4 := &tg.Message{}
			_ = m4.Decode(d)
			m.Message = m4

			if (flags & (1 << 2)) != 0 {
				m.ScheduleDate = new(int32)
				*m.ScheduleDate, err = d.Int32()
			}

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// OutboxMessage <--
type OutboxMessage struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	OutboxMessageClazz
}

// MakeOutboxMessage <--
func MakeOutboxMessage(c OutboxMessageClazz) *OutboxMessage {
	return &OutboxMessage{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		OutboxMessageClazz: c,
	}
}

// Encode <--
func (m *OutboxMessage) Encode(x *bin.Encoder, layer int32) error {
	if m.OutboxMessageClazz != nil {
		return m.OutboxMessageClazz.Encode(x, layer)
	}

	return fmt.Errorf("OutboxMessage - invalid Clazz")
}

// Decode <--
func (m *OutboxMessage) Decode(d *bin.Decoder) (err error) {
	m.OutboxMessageClazz, err = DecodeOutboxMessageClazz(d)
	return
}

// Match <--
func (m *OutboxMessage) Match(f ...interface{}) {
	switch c := m.OutboxMessageClazz.(type) {
	case *TLOutboxMessage:
		for _, v := range f {
			if f1, ok := v.(func(c *TLOutboxMessage) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToOutboxMessage <--
func (m *OutboxMessage) ToOutboxMessage() (*TLOutboxMessage, bool) {
	if m.OutboxMessageClazz == nil {
		return nil, false
	}

	if x, ok := m.OutboxMessageClazz.(*TLOutboxMessage); ok {
		return x, true
	}

	return nil, false
}

// SenderClazz <--
//   - TL_Sender
type SenderClazz interface {
	iface.TLObject
	SenderClazzName() string
}

func DecodeSenderClazz(d *bin.Decoder) (SenderClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_sender:
		x := &TLSender{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeSender - unexpected clazzId: %d", id)
	}
}

// TLSender <--
type TLSender struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	Type      int32  `json:"type"`
	AuthKeyId int64  `json:"auth_key_id"`
}

// SenderClazzName <--
func (m *TLSender) SenderClazzName() string {
	return ClazzName_sender
}

// ClazzName <--
func (m *TLSender) ClazzName() string {
	return ClazzName_sender
}

// ToSender <--
func (m *TLSender) ToSender() *Sender {
	return MakeSender(m)
}

// Encode <--
func (m *TLSender) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x5a3864ba: func() error {
			x.PutClazzID(0x5a3864ba)

			x.PutInt64(m.UserId)
			x.PutInt32(m.Type)
			x.PutInt64(m.AuthKeyId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_sender, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_sender, layer)
	}
}

// Decode <--
func (m *TLSender) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x5a3864ba: func() (err error) {
			m.UserId, err = d.Int64()
			m.Type, err = d.Int32()
			m.AuthKeyId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// Sender <--
type Sender struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	SenderClazz
}

// MakeSender <--
func MakeSender(c SenderClazz) *Sender {
	return &Sender{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		SenderClazz: c,
	}
}

// Encode <--
func (m *Sender) Encode(x *bin.Encoder, layer int32) error {
	if m.SenderClazz != nil {
		return m.SenderClazz.Encode(x, layer)
	}

	return fmt.Errorf("Sender - invalid Clazz")
}

// Decode <--
func (m *Sender) Decode(d *bin.Decoder) (err error) {
	m.SenderClazz, err = DecodeSenderClazz(d)
	return
}

// Match <--
func (m *Sender) Match(f ...interface{}) {
	switch c := m.SenderClazz.(type) {
	case *TLSender:
		for _, v := range f {
			if f1, ok := v.(func(c *TLSender) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToSender <--
func (m *Sender) ToSender() (*TLSender, bool) {
	if m.SenderClazz == nil {
		return nil, false
	}

	if x, ok := m.SenderClazz.(*TLSender); ok {
		return x, true
	}

	return nil, false
}
