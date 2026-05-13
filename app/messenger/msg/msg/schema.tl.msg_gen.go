/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
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
		return nil, fmt.Errorf("unable to decode ContentMessage: constructor: %w", err)
	}

	switch id {
	case 0x8d64b133:
		x := &TLContentMessage{ClazzID: id, ClazzName2: ClazzName_contentMessage}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode ContentMessage: invalid constructor %x", id)
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
	return iface.DebugStringWithName("contentMessage", m)
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
		return fmt.Errorf("unable to encode contentMessage: unsupported layer %d", layer)
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
		return fmt.Errorf("unable to encode contentMessage: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLContentMessage) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x8d64b133:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode contentMessage#0x8d64b133: field flags: %w", err)
		}
		_ = flags
		m.Id, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode contentMessage#0x8d64b133: field id: %w", err)
		}
		m.DialogMessageId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode contentMessage#0x8d64b133: field dialog_message_id: %w", err)
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
			return fmt.Errorf("unable to decode contentMessage#0x8d64b133: field send_user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode contentMessage: invalid constructor %x", m.ClazzID)
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
		return nil, fmt.Errorf("unable to decode OutboxMessage: constructor: %w", err)
	}

	switch id {
	case 0x625d8b25:
		x := &TLOutboxMessage{ClazzID: id, ClazzName2: ClazzName_outboxMessage}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode OutboxMessage: invalid constructor %x", id)
	}

}

// TLOutboxMessage <--
type TLOutboxMessage struct {
	ClazzID         uint32          `json:"_id"`
	ClazzName2      string          `json:"_name"`
	NoWebpage       bool            `json:"no_webpage"`
	Background      bool            `json:"background"`
	RandomId        int64           `json:"random_id"`
	Message         tg.MessageClazz `json:"message"`
	ScheduleDate    *int32          `json:"schedule_date"`
	ForwardSourceId *int32          `json:"forward_source_id"`
}

func MakeTLOutboxMessage(m *TLOutboxMessage) *TLOutboxMessage {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_outboxMessage

	return m
}

func (m *TLOutboxMessage) String() string {
	return iface.DebugStringWithName("outboxMessage", m)
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
	case 0x625d8b25:
		size := 4
		size += 4
		size += 8
		size += iface.CalcObjectSize(m.Message, layer)
		if m.ScheduleDate != nil {
			size += 4
		}

		if m.ForwardSourceId != nil {
			size += 4
		}

		return size
	default:
		return 0
	}
}

func (m *TLOutboxMessage) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_outboxMessage, int(layer)); clazzId {
	case 0x625d8b25:
		if err := iface.ValidateRequiredObject("message", m.Message); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode outboxMessage: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLOutboxMessage) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_outboxMessage, int(layer)); clazzId {
	case 0x625d8b25:
		x.PutClazzID(0x625d8b25)

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
			if m.ForwardSourceId != nil {
				flags |= 1 << 3
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.RandomId)
		if m.Message == nil {
			return fmt.Errorf("unable to encode outboxMessage#0x625d8b25: field message is nil")
		}
		if err := m.Message.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode outboxMessage#0x625d8b25: field message: %w", err)
		}
		if m.ScheduleDate != nil {
			x.PutInt32(*m.ScheduleDate)
		}

		if m.ForwardSourceId != nil {
			x.PutInt32(*m.ForwardSourceId)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode outboxMessage: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLOutboxMessage) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x625d8b25:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode outboxMessage#0x625d8b25: field flags: %w", err)
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
			return fmt.Errorf("unable to decode outboxMessage#0x625d8b25: field random_id: %w", err)
		}

		m.Message, err = tg.DecodeMessageClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode outboxMessage#0x625d8b25: field message: %w", err)
		}

		if (flags & (1 << 2)) != 0 {
			m.ScheduleDate = new(int32)
			*m.ScheduleDate, err = d.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode outboxMessage#0x625d8b25: field schedule_date: %w", err)
			}
		}
		if (flags & (1 << 3)) != 0 {
			m.ForwardSourceId = new(int32)
			*m.ForwardSourceId, err = d.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode outboxMessage#0x625d8b25: field forward_source_id: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to decode outboxMessage: invalid constructor %x", m.ClazzID)
	}
}

// OutboxMessage <--
type OutboxMessage = TLOutboxMessage

// ResolvedDialogCursorClazz <--
//   - TL_ResolvedDialogCursor
type ResolvedDialogCursorClazz = *TLResolvedDialogCursor

func DecodeResolvedDialogCursorClazz(d *bin.Decoder) (ResolvedDialogCursorClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode ResolvedDialogCursor: constructor: %w", err)
	}

	switch id {
	case 0x7debda91:
		x := &TLResolvedDialogCursor{ClazzID: id, ClazzName2: ClazzName_resolvedDialogCursor}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode ResolvedDialogCursor: invalid constructor %x", id)
	}

}

// TLResolvedDialogCursor <--
type TLResolvedDialogCursor struct {
	ClazzID     uint32       `json:"_id"`
	ClazzName2  string       `json:"_name"`
	Found       tg.BoolClazz `json:"found"`
	PeerType    int32        `json:"peer_type"`
	PeerId      int64        `json:"peer_id"`
	PeerSeq     int64        `json:"peer_seq"`
	MessageDate int64        `json:"message_date"`
}

func MakeTLResolvedDialogCursor(m *TLResolvedDialogCursor) *TLResolvedDialogCursor {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_resolvedDialogCursor

	return m
}

func (m *TLResolvedDialogCursor) String() string {
	return iface.DebugStringWithName("resolvedDialogCursor", m)
}

func (m *TLResolvedDialogCursor) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("resolvedDialogCursor", m)
}

// ResolvedDialogCursorClazzName <--
func (m *TLResolvedDialogCursor) ResolvedDialogCursorClazzName() string {
	return ClazzName_resolvedDialogCursor
}

// ClazzName <--
func (m *TLResolvedDialogCursor) ClazzName() string {
	return m.ClazzName2
}

// ToResolvedDialogCursor <--
func (m *TLResolvedDialogCursor) ToResolvedDialogCursor() *ResolvedDialogCursor {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLResolvedDialogCursor) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_resolvedDialogCursor, int(layer)); clazzId {
	case 0x7debda91:
		size := 4
		size += iface.CalcObjectSize(m.Found, layer)
		size += 4
		size += 8
		size += 8
		size += 8

		return size
	default:
		return 0
	}
}

func (m *TLResolvedDialogCursor) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_resolvedDialogCursor, int(layer)); clazzId {
	case 0x7debda91:
		if err := iface.ValidateRequiredObject("found", m.Found); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode resolvedDialogCursor: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLResolvedDialogCursor) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_resolvedDialogCursor, int(layer)); clazzId {
	case 0x7debda91:
		x.PutClazzID(0x7debda91)

		if m.Found == nil {
			return fmt.Errorf("unable to encode resolvedDialogCursor#0x7debda91: field found is nil")
		}
		if err := m.Found.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode resolvedDialogCursor#0x7debda91: field found: %w", err)
		}
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt64(m.PeerSeq)
		x.PutInt64(m.MessageDate)

		return nil
	default:
		return fmt.Errorf("unable to encode resolvedDialogCursor: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLResolvedDialogCursor) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x7debda91:

		m.Found, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode resolvedDialogCursor#0x7debda91: field found: %w", err)
		}

		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode resolvedDialogCursor#0x7debda91: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode resolvedDialogCursor#0x7debda91: field peer_id: %w", err)
		}
		m.PeerSeq, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode resolvedDialogCursor#0x7debda91: field peer_seq: %w", err)
		}
		m.MessageDate, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode resolvedDialogCursor#0x7debda91: field message_date: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode resolvedDialogCursor: invalid constructor %x", m.ClazzID)
	}
}

// ResolvedDialogCursor <--
type ResolvedDialogCursor = TLResolvedDialogCursor

// SenderClazz <--
//   - TL_Sender
type SenderClazz = *TLSender

func DecodeSenderClazz(d *bin.Decoder) (SenderClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode Sender: constructor: %w", err)
	}

	switch id {
	case 0x5a3864ba:
		x := &TLSender{ClazzID: id, ClazzName2: ClazzName_sender}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode Sender: invalid constructor %x", id)
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
	return iface.DebugStringWithName("sender", m)
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
		return fmt.Errorf("unable to encode sender: unsupported layer %d", layer)
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
		return fmt.Errorf("unable to encode sender: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLSender) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x5a3864ba:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode sender#0x5a3864ba: field user_id: %w", err)
		}
		m.Type, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode sender#0x5a3864ba: field type: %w", err)
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode sender#0x5a3864ba: field auth_key_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode sender: invalid constructor %x", m.ClazzID)
	}
}

// Sender <--
type Sender = TLSender

// UpdateFactClazz <--
//   - TL_UpdateFact
type UpdateFactClazz = *TLUpdateFact

func DecodeUpdateFactClazz(d *bin.Decoder) (UpdateFactClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode UpdateFact: constructor: %w", err)
	}

	switch id {
	case 0x4561a083:
		x := &TLUpdateFact{ClazzID: id, ClazzName2: ClazzName_updateFact}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode UpdateFact: invalid constructor %x", id)
	}

}

// TLUpdateFact <--
type TLUpdateFact struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Kind       string `json:"kind"`
	Payload    []byte `json:"payload"`
}

func MakeTLUpdateFact(m *TLUpdateFact) *TLUpdateFact {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_updateFact

	return m
}

func (m *TLUpdateFact) String() string {
	return iface.DebugStringWithName("updateFact", m)
}

func (m *TLUpdateFact) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("updateFact", m)
}

// UpdateFactClazzName <--
func (m *TLUpdateFact) UpdateFactClazzName() string {
	return ClazzName_updateFact
}

// ClazzName <--
func (m *TLUpdateFact) ClazzName() string {
	return m.ClazzName2
}

// ToUpdateFact <--
func (m *TLUpdateFact) ToUpdateFact() *UpdateFact {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLUpdateFact) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_updateFact, int(layer)); clazzId {
	case 0x4561a083:
		size := 4
		size += iface.CalcStringSize(m.Kind)
		size += iface.CalcBytesSize(m.Payload)

		return size
	default:
		return 0
	}
}

func (m *TLUpdateFact) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_updateFact, int(layer)); clazzId {
	case 0x4561a083:
		if err := iface.ValidateRequiredString("kind", m.Kind); err != nil {
			return err
		}

		if err := iface.ValidateRequiredBytes("payload", m.Payload); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode updateFact: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLUpdateFact) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_updateFact, int(layer)); clazzId {
	case 0x4561a083:
		x.PutClazzID(0x4561a083)

		x.PutString(m.Kind)
		x.PutBytes(m.Payload)

		return nil
	default:
		return fmt.Errorf("unable to encode updateFact: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUpdateFact) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x4561a083:
		m.Kind, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode updateFact#0x4561a083: field kind: %w", err)
		}
		m.Payload, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode updateFact#0x4561a083: field payload: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode updateFact: invalid constructor %x", m.ClazzID)
	}
}

// UpdateFact <--
type UpdateFact = TLUpdateFact
