/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package chat

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

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_chatInviteAlready:
		x := &TLChatInviteAlready{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_chatInvite:
		x := &TLChatInvite{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_chatInvitePeek:
		x := &TLChatInvitePeek{ClazzID: id}
		_ = x.Decode(d)
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
	wrapper := iface.WithNameWrapper{"chatInviteAlready", m}
	return wrapper.String()
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

// Encode <--
func (m *TLChatInviteAlready) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa40e7d5e: func() error {
			x.PutClazzID(0xa40e7d5e)

			_ = m.Chat.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chatInviteAlready, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chatInviteAlready, layer)
	}
}

// Decode <--
func (m *TLChatInviteAlready) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa40e7d5e: func() (err error) {

			// m0 := &tg.MutableChat{}
			// _ = m0.Decode(d)
			// m.Chat = m0
			m.Chat, _ = tg.DecodeMutableChatClazz(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
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
	wrapper := iface.WithNameWrapper{"chatInvite", m}
	return wrapper.String()
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

// Encode <--
func (m *TLChatInvite) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xdb75d1a7: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chatInvite, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chatInvite, layer)
	}
}

// Decode <--
func (m *TLChatInvite) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xdb75d1a7: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			if (flags & (1 << 6)) != 0 {
				m.RequestNeeded = true
			}
			m.Title, err = d.String()
			if (flags & (1 << 5)) != 0 {
				m.About = new(string)
				*m.About, err = d.String()
			}

			// m5 := &tg.Photo{}
			// _ = m5.Decode(d)
			// m.Photo = m5
			m.Photo, _ = tg.DecodePhotoClazz(d)

			m.ParticipantsCount, err = d.Int32()
			if (flags & (1 << 4)) != 0 {
				m.Participants, err = iface.DecodeInt64List(d)
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
	wrapper := iface.WithNameWrapper{"chatInvitePeek", m}
	return wrapper.String()
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

// Encode <--
func (m *TLChatInvitePeek) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xace3e26e: func() error {
			x.PutClazzID(0xace3e26e)

			_ = m.Chat.Encode(x, layer)
			x.PutInt32(m.Expires)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chatInvitePeek, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chatInvitePeek, layer)
	}
}

// Decode <--
func (m *TLChatInvitePeek) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xace3e26e: func() (err error) {

			// m0 := &tg.MutableChat{}
			// _ = m0.Decode(d)
			// m.Chat = m0
			m.Chat, _ = tg.DecodeMutableChatClazz(d)

			m.Expires, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
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
	wrapper := iface.WithNameWrapper{m.ClazzName(), m}
	return wrapper.String()
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

// Match <--
func (m *ChatInviteExt) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLChatInviteAlready:
		for _, v := range f {
			if f1, ok := v.(func(c *TLChatInviteAlready) interface{}); ok {
				f1(c)
			}
		}
	case *TLChatInvite:
		for _, v := range f {
			if f1, ok := v.(func(c *TLChatInvite) interface{}); ok {
				f1(c)
			}
		}
	case *TLChatInvitePeek:
		for _, v := range f {
			if f1, ok := v.(func(c *TLChatInvitePeek) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
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
type ChatInviteImportedClazz interface {
	iface.TLObject
	ChatInviteImportedClazzName() string
}

func DecodeChatInviteImportedClazz(d *bin.Decoder) (ChatInviteImportedClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_chatInviteImported:
		x := &TLChatInviteImported{ClazzID: id}
		_ = x.Decode(d)
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
	wrapper := iface.WithNameWrapper{"chatInviteImported", m}
	return wrapper.String()
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

	return &ChatInviteImported{Clazz: m}
}

// Encode <--
func (m *TLChatInviteImported) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x721051f6: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chatInviteImported, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chatInviteImported, layer)
	}
}

// Decode <--
func (m *TLChatInviteImported) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x721051f6: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags

			// m1 := &tg.MutableChat{}
			// _ = m1.Decode(d)
			// m.Chat = m1
			m.Chat, _ = tg.DecodeMutableChatClazz(d)

			if (flags & (1 << 0)) != 0 {
				// m2 := &RecentChatInviteRequesters{}
				// _ = m2.Decode(d)
				// m.Requesters = m2
				m.Requesters, _ = DecodeRecentChatInviteRequestersClazz(d)
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

// ChatInviteImported <--
type ChatInviteImported struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz ChatInviteImportedClazz `json:"_clazz"`
}

func (m *ChatInviteImported) String() string {
	wrapper := iface.WithNameWrapper{m.ClazzName(), m}
	return wrapper.String()
}

func (m *ChatInviteImported) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.ChatInviteImportedClazzName()
	}
}

// Encode <--
func (m *ChatInviteImported) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("ChatInviteImported - invalid Clazz")
}

// Decode <--
func (m *ChatInviteImported) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeChatInviteImportedClazz(d)
	return
}

// Match <--
func (m *ChatInviteImported) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLChatInviteImported:
		for _, v := range f {
			if f1, ok := v.(func(c *TLChatInviteImported) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToChatInviteImported <--
func (m *ChatInviteImported) ToChatInviteImported() (*TLChatInviteImported, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLChatInviteImported); ok {
		return x, true
	}

	return nil, false
}

// RecentChatInviteRequestersClazz <--
//   - TL_RecentChatInviteRequesters
type RecentChatInviteRequestersClazz interface {
	iface.TLObject
	RecentChatInviteRequestersClazzName() string
}

func DecodeRecentChatInviteRequestersClazz(d *bin.Decoder) (RecentChatInviteRequestersClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_recentChatInviteRequesters:
		x := &TLRecentChatInviteRequesters{ClazzID: id}
		_ = x.Decode(d)
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
	wrapper := iface.WithNameWrapper{"recentChatInviteRequesters", m}
	return wrapper.String()
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

	return &RecentChatInviteRequesters{Clazz: m}
}

// Encode <--
func (m *TLRecentChatInviteRequesters) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x1c6e3c54: func() error {
			x.PutClazzID(0x1c6e3c54)

			x.PutInt32(m.RequestsPending)

			iface.EncodeInt64List(x, m.RecentRequesters)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_recentChatInviteRequesters, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_recentChatInviteRequesters, layer)
	}
}

// Decode <--
func (m *TLRecentChatInviteRequesters) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x1c6e3c54: func() (err error) {
			m.RequestsPending, err = d.Int32()

			m.RecentRequesters, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// RecentChatInviteRequesters <--
type RecentChatInviteRequesters struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz RecentChatInviteRequestersClazz `json:"_clazz"`
}

func (m *RecentChatInviteRequesters) String() string {
	wrapper := iface.WithNameWrapper{m.ClazzName(), m}
	return wrapper.String()
}

func (m *RecentChatInviteRequesters) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.RecentChatInviteRequestersClazzName()
	}
}

// Encode <--
func (m *RecentChatInviteRequesters) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("RecentChatInviteRequesters - invalid Clazz")
}

// Decode <--
func (m *RecentChatInviteRequesters) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeRecentChatInviteRequestersClazz(d)
	return
}

// Match <--
func (m *RecentChatInviteRequesters) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLRecentChatInviteRequesters:
		for _, v := range f {
			if f1, ok := v.(func(c *TLRecentChatInviteRequesters) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToRecentChatInviteRequesters <--
func (m *RecentChatInviteRequesters) ToRecentChatInviteRequesters() (*TLRecentChatInviteRequesters, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLRecentChatInviteRequesters); ok {
		return x, true
	}

	return nil, false
}

// UserChatIdListClazz <--
//   - TL_UserChatIdList
type UserChatIdListClazz interface {
	iface.TLObject
	UserChatIdListClazzName() string
}

func DecodeUserChatIdListClazz(d *bin.Decoder) (UserChatIdListClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_userChatIdList:
		x := &TLUserChatIdList{ClazzID: id}
		_ = x.Decode(d)
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
	wrapper := iface.WithNameWrapper{"userChatIdList", m}
	return wrapper.String()
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

	return &UserChatIdList{Clazz: m}
}

// Encode <--
func (m *TLUserChatIdList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x50067224: func() error {
			x.PutClazzID(0x50067224)

			x.PutInt64(m.UserId)

			iface.EncodeInt64List(x, m.ChatIdList)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_userChatIdList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_userChatIdList, layer)
	}
}

// Decode <--
func (m *TLUserChatIdList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x50067224: func() (err error) {
			m.UserId, err = d.Int64()

			m.ChatIdList, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// UserChatIdList <--
type UserChatIdList struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz UserChatIdListClazz `json:"_clazz"`
}

func (m *UserChatIdList) String() string {
	wrapper := iface.WithNameWrapper{m.ClazzName(), m}
	return wrapper.String()
}

func (m *UserChatIdList) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.UserChatIdListClazzName()
	}
}

// Encode <--
func (m *UserChatIdList) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("UserChatIdList - invalid Clazz")
}

// Decode <--
func (m *UserChatIdList) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeUserChatIdListClazz(d)
	return
}

// Match <--
func (m *UserChatIdList) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLUserChatIdList:
		for _, v := range f {
			if f1, ok := v.(func(c *TLUserChatIdList) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToUserChatIdList <--
func (m *UserChatIdList) ToUserChatIdList() (*TLUserChatIdList, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLUserChatIdList); ok {
		return x, true
	}

	return nil, false
}
