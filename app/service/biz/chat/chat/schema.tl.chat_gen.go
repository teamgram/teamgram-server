/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package chat

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

// ChatAccessCheckResultClazz <--
//   - TL_ChatAccessCheckResult
type ChatAccessCheckResultClazz = *TLChatAccessCheckResult

func DecodeChatAccessCheckResultClazz(d *bin.Decoder) (ChatAccessCheckResultClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode ChatAccessCheckResult: constructor: %w", err)
	}

	switch id {
	case 0xc9b5daa6:
		x := &TLChatAccessCheckResult{ClazzID: id, ClazzName2: ClazzName_chatAccessCheckResult}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatAccessCheckResult: invalid constructor %x", id)
	}

}

// TLChatAccessCheckResult <--
type TLChatAccessCheckResult struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	SelfId     int64  `json:"self_id"`
	ChatId     int64  `json:"chat_id"`
	AccessKind string `json:"access_kind"`
}

func MakeTLChatAccessCheckResult(m *TLChatAccessCheckResult) *TLChatAccessCheckResult {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_chatAccessCheckResult

	return m
}

func (m *TLChatAccessCheckResult) String() string {
	return iface.DebugStringWithName("chatAccessCheckResult", m)
}

func (m *TLChatAccessCheckResult) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("chatAccessCheckResult", m)
}

// ChatAccessCheckResultClazzName <--
func (m *TLChatAccessCheckResult) ChatAccessCheckResultClazzName() string {
	return ClazzName_chatAccessCheckResult
}

// ClazzName <--
func (m *TLChatAccessCheckResult) ClazzName() string {
	return m.ClazzName2
}

// ToChatAccessCheckResult <--
func (m *TLChatAccessCheckResult) ToChatAccessCheckResult() *ChatAccessCheckResult {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLChatAccessCheckResult) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatAccessCheckResult, int(layer)); clazzId {
	case 0xc9b5daa6:
		size := 4
		size += 8
		size += 8
		size += iface.CalcStringSize(m.AccessKind)

		return size
	default:
		return 0
	}
}

func (m *TLChatAccessCheckResult) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatAccessCheckResult, int(layer)); clazzId {
	case 0xc9b5daa6:
		if err := iface.ValidateRequiredString("access_kind", m.AccessKind); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode chatAccessCheckResult: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLChatAccessCheckResult) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatAccessCheckResult, int(layer)); clazzId {
	case 0xc9b5daa6:
		x.PutClazzID(0xc9b5daa6)

		x.PutInt64(m.SelfId)
		x.PutInt64(m.ChatId)
		x.PutString(m.AccessKind)

		return nil
	default:
		return fmt.Errorf("unable to encode chatAccessCheckResult: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLChatAccessCheckResult) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xc9b5daa6:
		m.SelfId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode chatAccessCheckResult#0xc9b5daa6: field self_id: %w", err)
		}
		m.ChatId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode chatAccessCheckResult#0xc9b5daa6: field chat_id: %w", err)
		}
		m.AccessKind, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatAccessCheckResult#0xc9b5daa6: field access_kind: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode chatAccessCheckResult: invalid constructor %x", m.ClazzID)
	}
}

// ChatAccessCheckResult <--
type ChatAccessCheckResult = TLChatAccessCheckResult

// ChatInviteExtClazz <--
//   - TL_ChatInviteAlready
//   - TL_ChatInvite
//   - TL_ChatInvitePeek
type ChatInviteExtClazz interface {
	iface.TLObject
	ChatInviteExtClazzName() string
}

func DecodeChatInviteExtClazz(d *bin.Decoder) (ChatInviteExtClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode ChatInviteExt: constructor: %w", err)
	}

	switch id {
	case 0xa40e7d5e:
		x := &TLChatInviteAlready{ClazzID: id, ClazzName2: ClazzName_chatInviteAlready}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0xdb75d1a7:
		x := &TLChatInvite{ClazzID: id, ClazzName2: ClazzName_chatInvite}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0xace3e26e:
		x := &TLChatInvitePeek{ClazzID: id, ClazzName2: ClazzName_chatInvitePeek}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatInviteExt: invalid constructor %x", id)
	}

}

// TLChatInviteAlready <--
type TLChatInviteAlready struct {
	ClazzID    uint32              `json:"_id"`
	ClazzName2 string              `json:"_name"`
	Chat       tg.MutableChatClazz `json:"chat"`
}

func MakeTLChatInviteAlready(m *TLChatInviteAlready) *TLChatInviteAlready {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_chatInviteAlready

	return m
}

func (m *TLChatInviteAlready) String() string {
	return iface.DebugStringWithName("chatInviteAlready", m)
}

func (m *TLChatInviteAlready) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("chatInviteAlready", m)
}

// ChatInviteExtClazzName <--
func (m *TLChatInviteAlready) ChatInviteExtClazzName() string {
	return ClazzName_chatInviteAlready
}

// ClazzName <--
func (m *TLChatInviteAlready) ClazzName() string {
	return m.ClazzName2
}

// ToChatInviteExt <--
func (m *TLChatInviteAlready) ToChatInviteExt() *ChatInviteExt {
	if m == nil {
		return nil
	}

	return &ChatInviteExt{Clazz: m}

}

func (m *TLChatInviteAlready) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatInviteAlready, int(layer)); clazzId {
	case 0xa40e7d5e:
		size := 4
		size += iface.CalcObjectSize(m.Chat, layer)

		return size
	default:
		return 0
	}
}

func (m *TLChatInviteAlready) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatInviteAlready, int(layer)); clazzId {
	case 0xa40e7d5e:
		if err := iface.ValidateRequiredObject("chat", m.Chat); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode chatInviteAlready: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLChatInviteAlready) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatInviteAlready, int(layer)); clazzId {
	case 0xa40e7d5e:
		x.PutClazzID(0xa40e7d5e)

		if m.Chat == nil {
			return fmt.Errorf("unable to encode chatInviteAlready#0xa40e7d5e: field chat is nil")
		}
		if err := m.Chat.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode chatInviteAlready#0xa40e7d5e: field chat: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode chatInviteAlready: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLChatInviteAlready) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xa40e7d5e:

		m.Chat, err = tg.DecodeMutableChatClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteAlready#0xa40e7d5e: field chat: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode chatInviteAlready: invalid constructor %x", m.ClazzID)
	}
}

// TLChatInvite <--
type TLChatInvite struct {
	ClazzID           uint32        `json:"_id"`
	ClazzName2        string        `json:"_name"`
	RequestNeeded     bool          `json:"request_needed"`
	Title             string        `json:"title"`
	About             *string       `json:"about"`
	Photo             tg.PhotoClazz `json:"photo"`
	ParticipantsCount int32         `json:"participants_count"`
	Participants      []int64       `json:"participants"`
}

func MakeTLChatInvite(m *TLChatInvite) *TLChatInvite {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_chatInvite

	return m
}

func (m *TLChatInvite) String() string {
	return iface.DebugStringWithName("chatInvite", m)
}

func (m *TLChatInvite) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("chatInvite", m)
}

// ChatInviteExtClazzName <--
func (m *TLChatInvite) ChatInviteExtClazzName() string {
	return ClazzName_chatInvite
}

// ClazzName <--
func (m *TLChatInvite) ClazzName() string {
	return m.ClazzName2
}

// ToChatInviteExt <--
func (m *TLChatInvite) ToChatInviteExt() *ChatInviteExt {
	if m == nil {
		return nil
	}

	return &ChatInviteExt{Clazz: m}

}

func (m *TLChatInvite) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatInvite, int(layer)); clazzId {
	case 0xdb75d1a7:
		size := 4
		size += 4
		size += iface.CalcStringSize(m.Title)
		if m.About != nil {
			size += iface.CalcStringSize(*m.About)
		}

		size += iface.CalcObjectSize(m.Photo, layer)
		size += 4
		if m.Participants != nil {
			size += iface.CalcInt64ListSize(m.Participants)
		}

		return size
	default:
		return 0
	}
}

func (m *TLChatInvite) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatInvite, int(layer)); clazzId {
	case 0xdb75d1a7:
		if err := iface.ValidateRequiredString("title", m.Title); err != nil {
			return err
		}

		if err := iface.ValidateRequiredObject("photo", m.Photo); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode chatInvite: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLChatInvite) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatInvite, int(layer)); clazzId {
	case 0xdb75d1a7:
		x.PutClazzID(0xdb75d1a7)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.RequestNeeded == true {
				flags |= 1 << 6
			}

			if m.About != nil {
				flags |= 1 << 5
			}

			if m.Participants != nil {
				flags |= 1 << 4
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutString(m.Title)
		if m.About != nil {
			x.PutString(*m.About)
		}

		if m.Photo == nil {
			return fmt.Errorf("unable to encode chatInvite#0xdb75d1a7: field photo is nil")
		}
		if err := m.Photo.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode chatInvite#0xdb75d1a7: field photo: %w", err)
		}
		x.PutInt32(m.ParticipantsCount)
		if m.Participants != nil {
			iface.EncodeInt64List(x, m.Participants)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode chatInvite: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLChatInvite) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xdb75d1a7:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode chatInvite#0xdb75d1a7: field flags: %w", err)
		}
		_ = flags
		if (flags & (1 << 6)) != 0 {
			m.RequestNeeded = true
		}
		m.Title, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatInvite#0xdb75d1a7: field title: %w", err)
		}
		if (flags & (1 << 5)) != 0 {
			m.About = new(string)
			*m.About, err = d.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatInvite#0xdb75d1a7: field about: %w", err)
			}
		}

		m.Photo, err = tg.DecodePhotoClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode chatInvite#0xdb75d1a7: field photo: %w", err)
		}

		m.ParticipantsCount, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatInvite#0xdb75d1a7: field participants_count: %w", err)
		}
		if (flags & (1 << 4)) != 0 {
			m.Participants, err = iface.DecodeInt64List(d)
			if err != nil {
				return fmt.Errorf("unable to decode chatInvite#0xdb75d1a7: field participants: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to decode chatInvite: invalid constructor %x", m.ClazzID)
	}
}

// TLChatInvitePeek <--
type TLChatInvitePeek struct {
	ClazzID    uint32              `json:"_id"`
	ClazzName2 string              `json:"_name"`
	Chat       tg.MutableChatClazz `json:"chat"`
	Expires    int32               `json:"expires"`
}

func MakeTLChatInvitePeek(m *TLChatInvitePeek) *TLChatInvitePeek {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_chatInvitePeek

	return m
}

func (m *TLChatInvitePeek) String() string {
	return iface.DebugStringWithName("chatInvitePeek", m)
}

func (m *TLChatInvitePeek) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("chatInvitePeek", m)
}

// ChatInviteExtClazzName <--
func (m *TLChatInvitePeek) ChatInviteExtClazzName() string {
	return ClazzName_chatInvitePeek
}

// ClazzName <--
func (m *TLChatInvitePeek) ClazzName() string {
	return m.ClazzName2
}

// ToChatInviteExt <--
func (m *TLChatInvitePeek) ToChatInviteExt() *ChatInviteExt {
	if m == nil {
		return nil
	}

	return &ChatInviteExt{Clazz: m}

}

func (m *TLChatInvitePeek) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatInvitePeek, int(layer)); clazzId {
	case 0xace3e26e:
		size := 4
		size += iface.CalcObjectSize(m.Chat, layer)
		size += 4

		return size
	default:
		return 0
	}
}

func (m *TLChatInvitePeek) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatInvitePeek, int(layer)); clazzId {
	case 0xace3e26e:
		if err := iface.ValidateRequiredObject("chat", m.Chat); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode chatInvitePeek: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLChatInvitePeek) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatInvitePeek, int(layer)); clazzId {
	case 0xace3e26e:
		x.PutClazzID(0xace3e26e)

		if m.Chat == nil {
			return fmt.Errorf("unable to encode chatInvitePeek#0xace3e26e: field chat is nil")
		}
		if err := m.Chat.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode chatInvitePeek#0xace3e26e: field chat: %w", err)
		}
		x.PutInt32(m.Expires)

		return nil
	default:
		return fmt.Errorf("unable to encode chatInvitePeek: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLChatInvitePeek) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xace3e26e:

		m.Chat, err = tg.DecodeMutableChatClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode chatInvitePeek#0xace3e26e: field chat: %w", err)
		}

		m.Expires, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatInvitePeek#0xace3e26e: field expires: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode chatInvitePeek: invalid constructor %x", m.ClazzID)
	}
}

// ChatInviteExt <--
type ChatInviteExt struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz ChatInviteExtClazz `json:"_clazz"`
}

func (m *ChatInviteExt) String() string {
	return iface.DebugStringWithName(m.ClazzName(), m)
}

func (m *ChatInviteExt) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName(m.ClazzName(), m)
}

func (m *ChatInviteExt) CalcSize(layer int32) int {
	if m == nil || m.Clazz == nil {
		return 0
	}
	return iface.CalcObjectSize(m.Clazz, layer)
}

func (m *ChatInviteExt) Validate(layer int32) error {
	if m == nil {
		return fmt.Errorf("ChatInviteExt is required")
	}
	if m.Clazz == nil {
		return fmt.Errorf("ChatInviteExt.Clazz is required")
	}
	if v, ok := m.Clazz.(iface.TLObjectValidator); ok {
		if err := v.Validate(layer); err != nil {
			return fmt.Errorf("unable to validate ChatInviteExt.Clazz: %w", err)
		}
	}
	return nil
}

func (m *ChatInviteExt) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.ChatInviteExtClazzName()
	}
}

// Encode <--
func (m *ChatInviteExt) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		if err := m.Clazz.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode ChatInviteExt.Clazz: %w", err)
		}
		return nil
	}

	return fmt.Errorf("ChatInviteExt - invalid Clazz")
}

// Decode <--
func (m *ChatInviteExt) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeChatInviteExtClazz(d)
	return
}

// ToChatInviteAlready <--
func (m *ChatInviteExt) ToChatInviteAlready() (*TLChatInviteAlready, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLChatInviteAlready); ok {
		return x, true
	}

	return nil, false
}

// ToChatInvite <--
func (m *ChatInviteExt) ToChatInvite() (*TLChatInvite, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLChatInvite); ok {
		return x, true
	}

	return nil, false
}

// ToChatInvitePeek <--
func (m *ChatInviteExt) ToChatInvitePeek() (*TLChatInvitePeek, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLChatInvitePeek); ok {
		return x, true
	}

	return nil, false
}

// ChatInviteImportedClazz <--
//   - TL_ChatInviteImported
type ChatInviteImportedClazz = *TLChatInviteImported

func DecodeChatInviteImportedClazz(d *bin.Decoder) (ChatInviteImportedClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode ChatInviteImported: constructor: %w", err)
	}

	switch id {
	case 0x721051f6:
		x := &TLChatInviteImported{ClazzID: id, ClazzName2: ClazzName_chatInviteImported}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatInviteImported: invalid constructor %x", id)
	}

}

// TLChatInviteImported <--
type TLChatInviteImported struct {
	ClazzID    uint32                          `json:"_id"`
	ClazzName2 string                          `json:"_name"`
	Chat       tg.MutableChatClazz             `json:"chat"`
	Requesters RecentChatInviteRequestersClazz `json:"requesters"`
}

func MakeTLChatInviteImported(m *TLChatInviteImported) *TLChatInviteImported {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_chatInviteImported

	return m
}

func (m *TLChatInviteImported) String() string {
	return iface.DebugStringWithName("chatInviteImported", m)
}

func (m *TLChatInviteImported) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("chatInviteImported", m)
}

// ChatInviteImportedClazzName <--
func (m *TLChatInviteImported) ChatInviteImportedClazzName() string {
	return ClazzName_chatInviteImported
}

// ClazzName <--
func (m *TLChatInviteImported) ClazzName() string {
	return m.ClazzName2
}

// ToChatInviteImported <--
func (m *TLChatInviteImported) ToChatInviteImported() *ChatInviteImported {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLChatInviteImported) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatInviteImported, int(layer)); clazzId {
	case 0x721051f6:
		size := 4
		size += 4
		size += iface.CalcObjectSize(m.Chat, layer)
		if m.Requesters != nil {
			size += iface.CalcObjectSize(m.Requesters, layer)
		}

		return size
	default:
		return 0
	}
}

func (m *TLChatInviteImported) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatInviteImported, int(layer)); clazzId {
	case 0x721051f6:
		if err := iface.ValidateRequiredObject("chat", m.Chat); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode chatInviteImported: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLChatInviteImported) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatInviteImported, int(layer)); clazzId {
	case 0x721051f6:
		x.PutClazzID(0x721051f6)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.Requesters != nil {
				flags |= 1 << 0
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		if m.Chat == nil {
			return fmt.Errorf("unable to encode chatInviteImported#0x721051f6: field chat is nil")
		}
		if err := m.Chat.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode chatInviteImported#0x721051f6: field chat: %w", err)
		}
		if m.Requesters != nil {
			if err := m.Requesters.Encode(x, layer); err != nil {
				return fmt.Errorf("unable to encode chatInviteImported#0x721051f6: field requesters: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to encode chatInviteImported: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLChatInviteImported) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x721051f6:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteImported#0x721051f6: field flags: %w", err)
		}
		_ = flags

		m.Chat, err = tg.DecodeMutableChatClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteImported#0x721051f6: field chat: %w", err)
		}

		if (flags & (1 << 0)) != 0 {
			m.Requesters, err = DecodeRecentChatInviteRequestersClazz(d)
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteImported#0x721051f6: field requesters: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to decode chatInviteImported: invalid constructor %x", m.ClazzID)
	}
}

// ChatInviteImported <--
type ChatInviteImported = TLChatInviteImported

// ChatProjectionBundleClazz <--
//   - TL_ChatProjectionBundle
type ChatProjectionBundleClazz = *TLChatProjectionBundle

func DecodeChatProjectionBundleClazz(d *bin.Decoder) (ChatProjectionBundleClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode ChatProjectionBundle: constructor: %w", err)
	}

	switch id {
	case 0xff3f1aa4:
		x := &TLChatProjectionBundle{ClazzID: id, ClazzName2: ClazzName_chatProjectionBundle}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatProjectionBundle: invalid constructor %x", id)
	}

}

// TLChatProjectionBundle <--
type TLChatProjectionBundle struct {
	ClazzID        uint32             `json:"_id"`
	ClazzName2     string             `json:"_name"`
	ViewerChats    []ViewerChatsClazz `json:"viewer_chats"`
	MissingChatIds []int64            `json:"missing_chat_ids"`
}

func MakeTLChatProjectionBundle(m *TLChatProjectionBundle) *TLChatProjectionBundle {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_chatProjectionBundle

	return m
}

func (m *TLChatProjectionBundle) String() string {
	return iface.DebugStringWithName("chatProjectionBundle", m)
}

func (m *TLChatProjectionBundle) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("chatProjectionBundle", m)
}

// ChatProjectionBundleClazzName <--
func (m *TLChatProjectionBundle) ChatProjectionBundleClazzName() string {
	return ClazzName_chatProjectionBundle
}

// ClazzName <--
func (m *TLChatProjectionBundle) ClazzName() string {
	return m.ClazzName2
}

// ToChatProjectionBundle <--
func (m *TLChatProjectionBundle) ToChatProjectionBundle() *ChatProjectionBundle {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLChatProjectionBundle) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatProjectionBundle, int(layer)); clazzId {
	case 0xff3f1aa4:
		size := 4
		size += iface.CalcObjectListSize(m.ViewerChats, layer)
		size += iface.CalcInt64ListSize(m.MissingChatIds)

		return size
	default:
		return 0
	}
}

func (m *TLChatProjectionBundle) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatProjectionBundle, int(layer)); clazzId {
	case 0xff3f1aa4:
		if err := iface.ValidateRequiredSlice("viewer_chats", m.ViewerChats); err != nil {
			return err
		}

		if err := iface.ValidateRequiredSlice("missing_chat_ids", m.MissingChatIds); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode chatProjectionBundle: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLChatProjectionBundle) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatProjectionBundle, int(layer)); clazzId {
	case 0xff3f1aa4:
		x.PutClazzID(0xff3f1aa4)

		if err := iface.EncodeObjectList(x, m.ViewerChats, layer); err != nil {
			return fmt.Errorf("unable to encode chatProjectionBundle#0xff3f1aa4: field viewer_chats: %w", err)
		}

		iface.EncodeInt64List(x, m.MissingChatIds)

		return nil
	default:
		return fmt.Errorf("unable to encode chatProjectionBundle: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLChatProjectionBundle) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xff3f1aa4:
		l0, err3 := d.VectorHeader()
		if err3 != nil {
			return fmt.Errorf("unable to decode chatProjectionBundle#0xff3f1aa4: field viewer_chats: %w", err3)
		}
		if l0 > bin.MaxVectorLen {
			return fmt.Errorf("unable to decode chatProjectionBundle#0xff3f1aa4: field viewer_chats: %w", &bin.InvalidLengthError{Type: "vector", Length: int(l0)})
		}
		prealloc0 := int(l0)
		if prealloc0 > bin.PreallocateLimit {
			prealloc0 = bin.PreallocateLimit
		}
		v0 := make([]ViewerChatsClazz, 0, prealloc0)
		for i := int32(0); i < l0; i++ {
			vv0, err3 := DecodeViewerChatsClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode chatProjectionBundle#0xff3f1aa4: field viewer_chats: %w", err3)
			}
			v0 = append(v0, vv0)
		}
		m.ViewerChats = v0

		m.MissingChatIds, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode chatProjectionBundle#0xff3f1aa4: field missing_chat_ids: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode chatProjectionBundle: invalid constructor %x", m.ClazzID)
	}
}

// ChatProjectionBundle <--
type ChatProjectionBundle = TLChatProjectionBundle

// MessageActionCheckResultClazz <--
//   - TL_MessageActionCheckResult
type MessageActionCheckResultClazz = *TLMessageActionCheckResult

func DecodeMessageActionCheckResultClazz(d *bin.Decoder) (MessageActionCheckResultClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode MessageActionCheckResult: constructor: %w", err)
	}

	switch id {
	case 0x667011da:
		x := &TLMessageActionCheckResult{ClazzID: id, ClazzName2: ClazzName_messageActionCheckResult}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode MessageActionCheckResult: invalid constructor %x", id)
	}

}

// TLMessageActionCheckResult <--
type TLMessageActionCheckResult struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	SelfId     int64  `json:"self_id"`
	ChatId     int64  `json:"chat_id"`
	Action     string `json:"action"`
	MediaKind  string `json:"media_kind"`
}

func MakeTLMessageActionCheckResult(m *TLMessageActionCheckResult) *TLMessageActionCheckResult {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_messageActionCheckResult

	return m
}

func (m *TLMessageActionCheckResult) String() string {
	return iface.DebugStringWithName("messageActionCheckResult", m)
}

func (m *TLMessageActionCheckResult) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("messageActionCheckResult", m)
}

// MessageActionCheckResultClazzName <--
func (m *TLMessageActionCheckResult) MessageActionCheckResultClazzName() string {
	return ClazzName_messageActionCheckResult
}

// ClazzName <--
func (m *TLMessageActionCheckResult) ClazzName() string {
	return m.ClazzName2
}

// ToMessageActionCheckResult <--
func (m *TLMessageActionCheckResult) ToMessageActionCheckResult() *MessageActionCheckResult {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLMessageActionCheckResult) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_messageActionCheckResult, int(layer)); clazzId {
	case 0x667011da:
		size := 4
		size += 8
		size += 8
		size += iface.CalcStringSize(m.Action)
		size += iface.CalcStringSize(m.MediaKind)

		return size
	default:
		return 0
	}
}

func (m *TLMessageActionCheckResult) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_messageActionCheckResult, int(layer)); clazzId {
	case 0x667011da:
		if err := iface.ValidateRequiredString("action", m.Action); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("media_kind", m.MediaKind); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode messageActionCheckResult: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLMessageActionCheckResult) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_messageActionCheckResult, int(layer)); clazzId {
	case 0x667011da:
		x.PutClazzID(0x667011da)

		x.PutInt64(m.SelfId)
		x.PutInt64(m.ChatId)
		x.PutString(m.Action)
		x.PutString(m.MediaKind)

		return nil
	default:
		return fmt.Errorf("unable to encode messageActionCheckResult: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMessageActionCheckResult) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x667011da:
		m.SelfId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode messageActionCheckResult#0x667011da: field self_id: %w", err)
		}
		m.ChatId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode messageActionCheckResult#0x667011da: field chat_id: %w", err)
		}
		m.Action, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode messageActionCheckResult#0x667011da: field action: %w", err)
		}
		m.MediaKind, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode messageActionCheckResult#0x667011da: field media_kind: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode messageActionCheckResult: invalid constructor %x", m.ClazzID)
	}
}

// MessageActionCheckResult <--
type MessageActionCheckResult = TLMessageActionCheckResult

// RecentChatInviteRequestersClazz <--
//   - TL_RecentChatInviteRequesters
type RecentChatInviteRequestersClazz = *TLRecentChatInviteRequesters

func DecodeRecentChatInviteRequestersClazz(d *bin.Decoder) (RecentChatInviteRequestersClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode RecentChatInviteRequesters: constructor: %w", err)
	}

	switch id {
	case 0x1c6e3c54:
		x := &TLRecentChatInviteRequesters{ClazzID: id, ClazzName2: ClazzName_recentChatInviteRequesters}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode RecentChatInviteRequesters: invalid constructor %x", id)
	}

}

// TLRecentChatInviteRequesters <--
type TLRecentChatInviteRequesters struct {
	ClazzID          uint32  `json:"_id"`
	ClazzName2       string  `json:"_name"`
	RequestsPending  int32   `json:"requests_pending"`
	RecentRequesters []int64 `json:"recent_requesters"`
}

func MakeTLRecentChatInviteRequesters(m *TLRecentChatInviteRequesters) *TLRecentChatInviteRequesters {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_recentChatInviteRequesters

	return m
}

func (m *TLRecentChatInviteRequesters) String() string {
	return iface.DebugStringWithName("recentChatInviteRequesters", m)
}

func (m *TLRecentChatInviteRequesters) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("recentChatInviteRequesters", m)
}

// RecentChatInviteRequestersClazzName <--
func (m *TLRecentChatInviteRequesters) RecentChatInviteRequestersClazzName() string {
	return ClazzName_recentChatInviteRequesters
}

// ClazzName <--
func (m *TLRecentChatInviteRequesters) ClazzName() string {
	return m.ClazzName2
}

// ToRecentChatInviteRequesters <--
func (m *TLRecentChatInviteRequesters) ToRecentChatInviteRequesters() *RecentChatInviteRequesters {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLRecentChatInviteRequesters) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_recentChatInviteRequesters, int(layer)); clazzId {
	case 0x1c6e3c54:
		size := 4
		size += 4
		size += iface.CalcInt64ListSize(m.RecentRequesters)

		return size
	default:
		return 0
	}
}

func (m *TLRecentChatInviteRequesters) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_recentChatInviteRequesters, int(layer)); clazzId {
	case 0x1c6e3c54:
		if err := iface.ValidateRequiredSlice("recent_requesters", m.RecentRequesters); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode recentChatInviteRequesters: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLRecentChatInviteRequesters) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_recentChatInviteRequesters, int(layer)); clazzId {
	case 0x1c6e3c54:
		x.PutClazzID(0x1c6e3c54)

		x.PutInt32(m.RequestsPending)

		iface.EncodeInt64List(x, m.RecentRequesters)

		return nil
	default:
		return fmt.Errorf("unable to encode recentChatInviteRequesters: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLRecentChatInviteRequesters) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x1c6e3c54:
		m.RequestsPending, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode recentChatInviteRequesters#0x1c6e3c54: field requests_pending: %w", err)
		}

		m.RecentRequesters, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode recentChatInviteRequesters#0x1c6e3c54: field recent_requesters: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode recentChatInviteRequesters: invalid constructor %x", m.ClazzID)
	}
}

// RecentChatInviteRequesters <--
type RecentChatInviteRequesters = TLRecentChatInviteRequesters

// UserChatIdListClazz <--
//   - TL_UserChatIdList
type UserChatIdListClazz = *TLUserChatIdList

func DecodeUserChatIdListClazz(d *bin.Decoder) (UserChatIdListClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode UserChatIdList: constructor: %w", err)
	}

	switch id {
	case 0x50067224:
		x := &TLUserChatIdList{ClazzID: id, ClazzName2: ClazzName_userChatIdList}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode UserChatIdList: invalid constructor %x", id)
	}

}

// TLUserChatIdList <--
type TLUserChatIdList struct {
	ClazzID    uint32  `json:"_id"`
	ClazzName2 string  `json:"_name"`
	UserId     int64   `json:"user_id"`
	ChatIdList []int64 `json:"chat_id_list"`
}

func MakeTLUserChatIdList(m *TLUserChatIdList) *TLUserChatIdList {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_userChatIdList

	return m
}

func (m *TLUserChatIdList) String() string {
	return iface.DebugStringWithName("userChatIdList", m)
}

func (m *TLUserChatIdList) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("userChatIdList", m)
}

// UserChatIdListClazzName <--
func (m *TLUserChatIdList) UserChatIdListClazzName() string {
	return ClazzName_userChatIdList
}

// ClazzName <--
func (m *TLUserChatIdList) ClazzName() string {
	return m.ClazzName2
}

// ToUserChatIdList <--
func (m *TLUserChatIdList) ToUserChatIdList() *UserChatIdList {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLUserChatIdList) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userChatIdList, int(layer)); clazzId {
	case 0x50067224:
		size := 4
		size += 8
		size += iface.CalcInt64ListSize(m.ChatIdList)

		return size
	default:
		return 0
	}
}

func (m *TLUserChatIdList) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userChatIdList, int(layer)); clazzId {
	case 0x50067224:
		if err := iface.ValidateRequiredSlice("chat_id_list", m.ChatIdList); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode userChatIdList: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLUserChatIdList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userChatIdList, int(layer)); clazzId {
	case 0x50067224:
		x.PutClazzID(0x50067224)

		x.PutInt64(m.UserId)

		iface.EncodeInt64List(x, m.ChatIdList)

		return nil
	default:
		return fmt.Errorf("unable to encode userChatIdList: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserChatIdList) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x50067224:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userChatIdList#0x50067224: field user_id: %w", err)
		}

		m.ChatIdList, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode userChatIdList#0x50067224: field chat_id_list: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode userChatIdList: invalid constructor %x", m.ClazzID)
	}
}

// UserChatIdList <--
type UserChatIdList = TLUserChatIdList

// ViewerChatsClazz <--
//   - TL_ViewerChats
type ViewerChatsClazz = *TLViewerChats

func DecodeViewerChatsClazz(d *bin.Decoder) (ViewerChatsClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode ViewerChats: constructor: %w", err)
	}

	switch id {
	case 0x689a4cf:
		x := &TLViewerChats{ClazzID: id, ClazzName2: ClazzName_viewerChats}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode ViewerChats: invalid constructor %x", id)
	}

}

// TLViewerChats <--
type TLViewerChats struct {
	ClazzID      uint32         `json:"_id"`
	ClazzName2   string         `json:"_name"`
	ViewerUserId int64          `json:"viewer_user_id"`
	Chats        []tg.ChatClazz `json:"chats"`
}

func MakeTLViewerChats(m *TLViewerChats) *TLViewerChats {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_viewerChats

	return m
}

func (m *TLViewerChats) String() string {
	return iface.DebugStringWithName("viewerChats", m)
}

func (m *TLViewerChats) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("viewerChats", m)
}

// ViewerChatsClazzName <--
func (m *TLViewerChats) ViewerChatsClazzName() string {
	return ClazzName_viewerChats
}

// ClazzName <--
func (m *TLViewerChats) ClazzName() string {
	return m.ClazzName2
}

// ToViewerChats <--
func (m *TLViewerChats) ToViewerChats() *ViewerChats {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLViewerChats) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_viewerChats, int(layer)); clazzId {
	case 0x689a4cf:
		size := 4
		size += 8
		size += iface.CalcObjectListSize(m.Chats, layer)

		return size
	default:
		return 0
	}
}

func (m *TLViewerChats) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_viewerChats, int(layer)); clazzId {
	case 0x689a4cf:
		if err := iface.ValidateRequiredSlice("chats", m.Chats); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode viewerChats: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLViewerChats) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_viewerChats, int(layer)); clazzId {
	case 0x689a4cf:
		x.PutClazzID(0x689a4cf)

		x.PutInt64(m.ViewerUserId)

		if err := iface.EncodeObjectList(x, m.Chats, layer); err != nil {
			return fmt.Errorf("unable to encode viewerChats#0x689a4cf: field chats: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode viewerChats: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLViewerChats) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x689a4cf:
		m.ViewerUserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode viewerChats#0x689a4cf: field viewer_user_id: %w", err)
		}
		l1, err3 := d.VectorHeader()
		if err3 != nil {
			return fmt.Errorf("unable to decode viewerChats#0x689a4cf: field chats: %w", err3)
		}
		if l1 > bin.MaxVectorLen {
			return fmt.Errorf("unable to decode viewerChats#0x689a4cf: field chats: %w", &bin.InvalidLengthError{Type: "vector", Length: int(l1)})
		}
		prealloc1 := int(l1)
		if prealloc1 > bin.PreallocateLimit {
			prealloc1 = bin.PreallocateLimit
		}
		v1 := make([]tg.ChatClazz, 0, prealloc1)
		for i := int32(0); i < l1; i++ {
			vv1, err3 := tg.DecodeChatClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode viewerChats#0x689a4cf: field chats: %w", err3)
			}
			v1 = append(v1, vv1)
		}
		m.Chats = v1

		return nil
	default:
		return fmt.Errorf("unable to decode viewerChats: invalid constructor %x", m.ClazzID)
	}
}

// ViewerChats <--
type ViewerChats = TLViewerChats
