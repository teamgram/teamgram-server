/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
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
		return nil, err
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
		return nil, fmt.Errorf("DecodeChatInviteExt - unexpected clazzId: %d", id)
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
	data, _ := json.Marshal(m)
	return string(data)
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
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chatInviteAlready, layer)
	}
}

// Encode <--
func (m *TLChatInviteAlready) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatInviteAlready, int(layer)); clazzId {
	case 0xa40e7d5e:
		x.PutClazzID(0xa40e7d5e)

		_ = m.Chat.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chatInviteAlready, layer)
	}
}

// Decode <--
func (m *TLChatInviteAlready) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xa40e7d5e:

		m.Chat, err = tg.DecodeMutableChatClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	data, _ := json.Marshal(m)
	return string(data)
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
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chatInvite, layer)
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

		_ = m.Photo.Encode(x, layer)
		x.PutInt32(m.ParticipantsCount)
		if m.Participants != nil {
			iface.EncodeInt64List(x, m.Participants)
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chatInvite, layer)
	}
}

// Decode <--
func (m *TLChatInvite) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xdb75d1a7:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		if (flags & (1 << 6)) != 0 {
			m.RequestNeeded = true
		}
		m.Title, err = d.String()
		if err != nil {
			return err
		}
		if (flags & (1 << 5)) != 0 {
			m.About = new(string)
			*m.About, err = d.String()
			if err != nil {
				return err
			}
		}

		m.Photo, err = tg.DecodePhotoClazz(d)
		if err != nil {
			return err
		}

		m.ParticipantsCount, err = d.Int32()
		if err != nil {
			return err
		}
		if (flags & (1 << 4)) != 0 {
			m.Participants, err = iface.DecodeInt64List(d)
			if err != nil {
				return err
			}
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	data, _ := json.Marshal(m)
	return string(data)
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
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chatInvitePeek, layer)
	}
}

// Encode <--
func (m *TLChatInvitePeek) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_chatInvitePeek, int(layer)); clazzId {
	case 0xace3e26e:
		x.PutClazzID(0xace3e26e)

		_ = m.Chat.Encode(x, layer)
		x.PutInt32(m.Expires)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chatInvitePeek, layer)
	}
}

// Decode <--
func (m *TLChatInvitePeek) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xace3e26e:

		m.Chat, err = tg.DecodeMutableChatClazz(d)
		if err != nil {
			return err
		}

		m.Expires, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// ChatInviteExt <--
type ChatInviteExt struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz ChatInviteExtClazz `json:"_clazz"`
}

func (m *ChatInviteExt) String() string {
	data, _ := json.Marshal(m)
	return string(data)
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
		return v.Validate(layer)
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
		return m.Clazz.Encode(x, layer)
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
		return nil, err
	}

	switch id {
	case 0x721051f6:
		x := &TLChatInviteImported{ClazzID: id, ClazzName2: ClazzName_chatInviteImported}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeChatInviteImported - unexpected clazzId: %d", id)
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
	data, _ := json.Marshal(m)
	return string(data)
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
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chatInviteImported, layer)
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
		_ = m.Chat.Encode(x, layer)
		if m.Requesters != nil {
			_ = m.Requesters.Encode(x, layer)
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chatInviteImported, layer)
	}
}

// Decode <--
func (m *TLChatInviteImported) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x721051f6:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags

		m.Chat, err = tg.DecodeMutableChatClazz(d)
		if err != nil {
			return err
		}

		if (flags & (1 << 0)) != 0 {
			m.Requesters, err = DecodeRecentChatInviteRequestersClazz(d)
			if err != nil {
				return err
			}
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// ChatInviteImported <--
type ChatInviteImported = TLChatInviteImported

// RecentChatInviteRequestersClazz <--
//   - TL_RecentChatInviteRequesters
type RecentChatInviteRequestersClazz = *TLRecentChatInviteRequesters

func DecodeRecentChatInviteRequestersClazz(d *bin.Decoder) (RecentChatInviteRequestersClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x1c6e3c54:
		x := &TLRecentChatInviteRequesters{ClazzID: id, ClazzName2: ClazzName_recentChatInviteRequesters}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeRecentChatInviteRequesters - unexpected clazzId: %d", id)
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
	data, _ := json.Marshal(m)
	return string(data)
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
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_recentChatInviteRequesters, layer)
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
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_recentChatInviteRequesters, layer)
	}
}

// Decode <--
func (m *TLRecentChatInviteRequesters) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x1c6e3c54:
		m.RequestsPending, err = d.Int32()
		if err != nil {
			return err
		}

		m.RecentRequesters, err = iface.DecodeInt64List(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
		return nil, err
	}

	switch id {
	case 0x50067224:
		x := &TLUserChatIdList{ClazzID: id, ClazzName2: ClazzName_userChatIdList}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeUserChatIdList - unexpected clazzId: %d", id)
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
	data, _ := json.Marshal(m)
	return string(data)
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
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_userChatIdList, layer)
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
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_userChatIdList, layer)
	}
}

// Decode <--
func (m *TLUserChatIdList) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x50067224:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		m.ChatIdList, err = iface.DecodeInt64List(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// UserChatIdList <--
type UserChatIdList = TLUserChatIdList
