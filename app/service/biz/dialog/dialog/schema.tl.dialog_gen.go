/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dialog

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

// DialogExtClazz <--
//   - TL_DialogExt
type DialogExtClazz = *TLDialogExt

func DecodeDialogExtClazz(d *bin.Decoder) (DialogExtClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode DialogExt: constructor: %w", err)
	}

	switch id {
	case 0x730ba93f:
		x := &TLDialogExt{ClazzID: id, ClazzName2: ClazzName_dialogExt}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode DialogExt: invalid constructor %x", id)
	}

}

// TLDialogExt <--
type TLDialogExt struct {
	ClazzID             uint32         `json:"_id"`
	ClazzName2          string         `json:"_name"`
	Order               int64          `json:"order"`
	Dialog              tg.DialogClazz `json:"dialog"`
	AvailableMinId      int32          `json:"available_min_id"`
	Date                int64          `json:"date"`
	ThemeEmoticon       string         `json:"theme_emoticon"`
	TtlPeriod           int32          `json:"ttl_period"`
	WallpaperId         int64          `json:"wallpaper_id"`
	WallpaperOverridden bool           `json:"wallpaper_overridden"`
}

func MakeTLDialogExt(m *TLDialogExt) *TLDialogExt {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_dialogExt

	return m
}

func (m *TLDialogExt) String() string {
	return iface.DebugStringWithName("dialogExt", m)
}

func (m *TLDialogExt) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("dialogExt", m)
}

// DialogExtClazzName <--
func (m *TLDialogExt) DialogExtClazzName() string {
	return ClazzName_dialogExt
}

// ClazzName <--
func (m *TLDialogExt) ClazzName() string {
	return m.ClazzName2
}

// ToDialogExt <--
func (m *TLDialogExt) ToDialogExt() *DialogExt {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLDialogExt) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogExt, int(layer)); clazzId {
	case 0x730ba93f:
		size := 4
		size += 4
		size += 8
		size += iface.CalcObjectSize(m.Dialog, layer)
		size += 4
		size += 8
		size += iface.CalcStringSize(m.ThemeEmoticon)
		size += 4
		size += 8

		return size
	default:
		return 0
	}
}

func (m *TLDialogExt) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogExt, int(layer)); clazzId {
	case 0x730ba93f:
		if err := iface.ValidateRequiredObject("dialog", m.Dialog); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("theme_emoticon", m.ThemeEmoticon); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode dialogExt: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLDialogExt) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogExt, int(layer)); clazzId {
	case 0x730ba93f:
		x.PutClazzID(0x730ba93f)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.WallpaperOverridden == true {
				flags |= 1 << 0
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.Order)
		if m.Dialog == nil {
			return fmt.Errorf("unable to encode dialogExt#0x730ba93f: field dialog is nil")
		}
		if err := m.Dialog.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode dialogExt#0x730ba93f: field dialog: %w", err)
		}
		x.PutInt32(m.AvailableMinId)
		x.PutInt64(m.Date)
		x.PutString(m.ThemeEmoticon)
		x.PutInt32(m.TtlPeriod)
		x.PutInt64(m.WallpaperId)

		return nil
	default:
		return fmt.Errorf("unable to encode dialogExt: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDialogExt) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x730ba93f:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExt#0x730ba93f: field flags: %w", err)
		}
		_ = flags
		m.Order, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExt#0x730ba93f: field order: %w", err)
		}

		m.Dialog, err = tg.DecodeDialogClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode dialogExt#0x730ba93f: field dialog: %w", err)
		}

		m.AvailableMinId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExt#0x730ba93f: field available_min_id: %w", err)
		}
		m.Date, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExt#0x730ba93f: field date: %w", err)
		}
		m.ThemeEmoticon, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExt#0x730ba93f: field theme_emoticon: %w", err)
		}
		m.TtlPeriod, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExt#0x730ba93f: field ttl_period: %w", err)
		}
		m.WallpaperId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExt#0x730ba93f: field wallpaper_id: %w", err)
		}
		if (flags & (1 << 0)) != 0 {
			m.WallpaperOverridden = true
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dialogExt: invalid constructor %x", m.ClazzID)
	}
}

// DialogExt <--
type DialogExt = TLDialogExt

// DialogFilterExtClazz <--
//   - TL_DialogFilterExt
type DialogFilterExtClazz = *TLDialogFilterExt

func DecodeDialogFilterExtClazz(d *bin.Decoder) (DialogFilterExtClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode DialogFilterExt: constructor: %w", err)
	}

	switch id {
	case 0xa6d498fe:
		x := &TLDialogFilterExt{ClazzID: id, ClazzName2: ClazzName_dialogFilterExt}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode DialogFilterExt: invalid constructor %x", id)
	}

}

// TLDialogFilterExt <--
type TLDialogFilterExt struct {
	ClazzID      uint32               `json:"_id"`
	ClazzName2   string               `json:"_name"`
	Id           int32                `json:"id"`
	JoinedBySlug bool                 `json:"joined_by_slug"`
	Slug         string               `json:"slug"`
	DialogFilter tg.DialogFilterClazz `json:"dialog_filter"`
	Order        int64                `json:"order"`
}

func MakeTLDialogFilterExt(m *TLDialogFilterExt) *TLDialogFilterExt {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_dialogFilterExt

	return m
}

func (m *TLDialogFilterExt) String() string {
	return iface.DebugStringWithName("dialogFilterExt", m)
}

func (m *TLDialogFilterExt) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("dialogFilterExt", m)
}

// DialogFilterExtClazzName <--
func (m *TLDialogFilterExt) DialogFilterExtClazzName() string {
	return ClazzName_dialogFilterExt
}

// ClazzName <--
func (m *TLDialogFilterExt) ClazzName() string {
	return m.ClazzName2
}

// ToDialogFilterExt <--
func (m *TLDialogFilterExt) ToDialogFilterExt() *DialogFilterExt {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLDialogFilterExt) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogFilterExt, int(layer)); clazzId {
	case 0xa6d498fe:
		size := 4
		size += 4
		size += 4
		size += iface.CalcStringSize(m.Slug)
		size += iface.CalcObjectSize(m.DialogFilter, layer)
		size += 8

		return size
	default:
		return 0
	}
}

func (m *TLDialogFilterExt) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogFilterExt, int(layer)); clazzId {
	case 0xa6d498fe:
		if err := iface.ValidateRequiredString("slug", m.Slug); err != nil {
			return err
		}

		if err := iface.ValidateRequiredObject("dialog_filter", m.DialogFilter); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode dialogFilterExt: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLDialogFilterExt) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogFilterExt, int(layer)); clazzId {
	case 0xa6d498fe:
		x.PutClazzID(0xa6d498fe)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.JoinedBySlug == true {
				flags |= 1 << 0
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt32(m.Id)
		x.PutString(m.Slug)
		if m.DialogFilter == nil {
			return fmt.Errorf("unable to encode dialogFilterExt#0xa6d498fe: field dialog_filter is nil")
		}
		if err := m.DialogFilter.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode dialogFilterExt#0xa6d498fe: field dialog_filter: %w", err)
		}
		x.PutInt64(m.Order)

		return nil
	default:
		return fmt.Errorf("unable to encode dialogFilterExt: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDialogFilterExt) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xa6d498fe:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilterExt#0xa6d498fe: field flags: %w", err)
		}
		_ = flags
		m.Id, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilterExt#0xa6d498fe: field id: %w", err)
		}
		if (flags & (1 << 0)) != 0 {
			m.JoinedBySlug = true
		}
		m.Slug, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilterExt#0xa6d498fe: field slug: %w", err)
		}

		m.DialogFilter, err = tg.DecodeDialogFilterClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilterExt#0xa6d498fe: field dialog_filter: %w", err)
		}

		m.Order, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilterExt#0xa6d498fe: field order: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dialogFilterExt: invalid constructor %x", m.ClazzID)
	}
}

// DialogFilterExt <--
type DialogFilterExt = TLDialogFilterExt

// DialogPinnedExtClazz <--
//   - TL_DialogPinnedExt
type DialogPinnedExtClazz = *TLDialogPinnedExt

func DecodeDialogPinnedExtClazz(d *bin.Decoder) (DialogPinnedExtClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode DialogPinnedExt: constructor: %w", err)
	}

	switch id {
	case 0xea7222c:
		x := &TLDialogPinnedExt{ClazzID: id, ClazzName2: ClazzName_dialogPinnedExt}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode DialogPinnedExt: invalid constructor %x", id)
	}

}

// TLDialogPinnedExt <--
type TLDialogPinnedExt struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Order      int64  `json:"order"`
	PeerType   int32  `json:"peer_type"`
	PeerId     int64  `json:"peer_id"`
}

func MakeTLDialogPinnedExt(m *TLDialogPinnedExt) *TLDialogPinnedExt {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_dialogPinnedExt

	return m
}

func (m *TLDialogPinnedExt) String() string {
	return iface.DebugStringWithName("dialogPinnedExt", m)
}

func (m *TLDialogPinnedExt) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("dialogPinnedExt", m)
}

// DialogPinnedExtClazzName <--
func (m *TLDialogPinnedExt) DialogPinnedExtClazzName() string {
	return ClazzName_dialogPinnedExt
}

// ClazzName <--
func (m *TLDialogPinnedExt) ClazzName() string {
	return m.ClazzName2
}

// ToDialogPinnedExt <--
func (m *TLDialogPinnedExt) ToDialogPinnedExt() *DialogPinnedExt {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLDialogPinnedExt) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogPinnedExt, int(layer)); clazzId {
	case 0xea7222c:
		size := 4
		size += 8
		size += 4
		size += 8

		return size
	default:
		return 0
	}
}

func (m *TLDialogPinnedExt) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogPinnedExt, int(layer)); clazzId {
	case 0xea7222c:

		return nil
	default:
		return fmt.Errorf("unable to encode dialogPinnedExt: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLDialogPinnedExt) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogPinnedExt, int(layer)); clazzId {
	case 0xea7222c:
		x.PutClazzID(0xea7222c)

		x.PutInt64(m.Order)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		return fmt.Errorf("unable to encode dialogPinnedExt: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDialogPinnedExt) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xea7222c:
		m.Order, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogPinnedExt#0xea7222c: field order: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dialogPinnedExt#0xea7222c: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogPinnedExt#0xea7222c: field peer_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dialogPinnedExt: invalid constructor %x", m.ClazzID)
	}
}

// DialogPinnedExt <--
type DialogPinnedExt = TLDialogPinnedExt

// DialogsDataClazz <--
//   - TL_SimpleDialogsData
type DialogsDataClazz = *TLSimpleDialogsData

func DecodeDialogsDataClazz(d *bin.Decoder) (DialogsDataClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode DialogsData: constructor: %w", err)
	}

	switch id {
	case 0x1d59b45d:
		x := &TLSimpleDialogsData{ClazzID: id, ClazzName2: ClazzName_simpleDialogsData}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode DialogsData: invalid constructor %x", id)
	}

}

// TLSimpleDialogsData <--
type TLSimpleDialogsData struct {
	ClazzID    uint32  `json:"_id"`
	ClazzName2 string  `json:"_name"`
	Users      []int64 `json:"users"`
	Chats      []int64 `json:"chats"`
	Channels   []int64 `json:"channels"`
}

func MakeTLSimpleDialogsData(m *TLSimpleDialogsData) *TLSimpleDialogsData {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_simpleDialogsData

	return m
}

func (m *TLSimpleDialogsData) String() string {
	return iface.DebugStringWithName("simpleDialogsData", m)
}

func (m *TLSimpleDialogsData) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("simpleDialogsData", m)
}

// DialogsDataClazzName <--
func (m *TLSimpleDialogsData) DialogsDataClazzName() string {
	return ClazzName_simpleDialogsData
}

// ClazzName <--
func (m *TLSimpleDialogsData) ClazzName() string {
	return m.ClazzName2
}

// ToDialogsData <--
func (m *TLSimpleDialogsData) ToDialogsData() *DialogsData {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLSimpleDialogsData) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_simpleDialogsData, int(layer)); clazzId {
	case 0x1d59b45d:
		size := 4
		size += iface.CalcInt64ListSize(m.Users)
		size += iface.CalcInt64ListSize(m.Chats)
		size += iface.CalcInt64ListSize(m.Channels)

		return size
	default:
		return 0
	}
}

func (m *TLSimpleDialogsData) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_simpleDialogsData, int(layer)); clazzId {
	case 0x1d59b45d:
		if err := iface.ValidateRequiredSlice("users", m.Users); err != nil {
			return err
		}

		if err := iface.ValidateRequiredSlice("chats", m.Chats); err != nil {
			return err
		}

		if err := iface.ValidateRequiredSlice("channels", m.Channels); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode simpleDialogsData: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLSimpleDialogsData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_simpleDialogsData, int(layer)); clazzId {
	case 0x1d59b45d:
		x.PutClazzID(0x1d59b45d)

		iface.EncodeInt64List(x, m.Users)

		iface.EncodeInt64List(x, m.Chats)

		iface.EncodeInt64List(x, m.Channels)

		return nil
	default:
		return fmt.Errorf("unable to encode simpleDialogsData: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLSimpleDialogsData) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x1d59b45d:

		m.Users, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode simpleDialogsData#0x1d59b45d: field users: %w", err)
		}

		m.Chats, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode simpleDialogsData#0x1d59b45d: field chats: %w", err)
		}

		m.Channels, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode simpleDialogsData#0x1d59b45d: field channels: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode simpleDialogsData: invalid constructor %x", m.ClazzID)
	}
}

// DialogsData <--
type DialogsData = TLSimpleDialogsData

// PeerWithDraftMessageClazz <--
//   - TL_UpdateDraftMessage
type PeerWithDraftMessageClazz = *TLUpdateDraftMessage

func DecodePeerWithDraftMessageClazz(d *bin.Decoder) (PeerWithDraftMessageClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode PeerWithDraftMessage: constructor: %w", err)
	}

	switch id {
	case 0xf6bdc4b2:
		x := &TLUpdateDraftMessage{ClazzID: id, ClazzName2: ClazzName_updateDraftMessage}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode PeerWithDraftMessage: invalid constructor %x", id)
	}

}

// TLUpdateDraftMessage <--
type TLUpdateDraftMessage struct {
	ClazzID    uint32               `json:"_id"`
	ClazzName2 string               `json:"_name"`
	Peer       tg.PeerClazz         `json:"peer"`
	Draft      tg.DraftMessageClazz `json:"draft"`
}

func MakeTLUpdateDraftMessage(m *TLUpdateDraftMessage) *TLUpdateDraftMessage {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_updateDraftMessage

	return m
}

func (m *TLUpdateDraftMessage) String() string {
	return iface.DebugStringWithName("updateDraftMessage", m)
}

func (m *TLUpdateDraftMessage) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("updateDraftMessage", m)
}

// PeerWithDraftMessageClazzName <--
func (m *TLUpdateDraftMessage) PeerWithDraftMessageClazzName() string {
	return ClazzName_updateDraftMessage
}

// ClazzName <--
func (m *TLUpdateDraftMessage) ClazzName() string {
	return m.ClazzName2
}

// ToPeerWithDraftMessage <--
func (m *TLUpdateDraftMessage) ToPeerWithDraftMessage() *PeerWithDraftMessage {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLUpdateDraftMessage) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_updateDraftMessage, int(layer)); clazzId {
	case 0xf6bdc4b2:
		size := 4
		size += iface.CalcObjectSize(m.Peer, layer)
		size += iface.CalcObjectSize(m.Draft, layer)

		return size
	default:
		return 0
	}
}

func (m *TLUpdateDraftMessage) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_updateDraftMessage, int(layer)); clazzId {
	case 0xf6bdc4b2:
		if err := iface.ValidateRequiredObject("peer", m.Peer); err != nil {
			return err
		}

		if err := iface.ValidateRequiredObject("draft", m.Draft); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode updateDraftMessage: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLUpdateDraftMessage) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_updateDraftMessage, int(layer)); clazzId {
	case 0xf6bdc4b2:
		x.PutClazzID(0xf6bdc4b2)

		if m.Peer == nil {
			return fmt.Errorf("unable to encode updateDraftMessage#0xf6bdc4b2: field peer is nil")
		}
		if err := m.Peer.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode updateDraftMessage#0xf6bdc4b2: field peer: %w", err)
		}
		if m.Draft == nil {
			return fmt.Errorf("unable to encode updateDraftMessage#0xf6bdc4b2: field draft is nil")
		}
		if err := m.Draft.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode updateDraftMessage#0xf6bdc4b2: field draft: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode updateDraftMessage: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUpdateDraftMessage) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xf6bdc4b2:

		m.Peer, err = tg.DecodePeerClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode updateDraftMessage#0xf6bdc4b2: field peer: %w", err)
		}

		m.Draft, err = tg.DecodeDraftMessageClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode updateDraftMessage#0xf6bdc4b2: field draft: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode updateDraftMessage: invalid constructor %x", m.ClazzID)
	}
}

// PeerWithDraftMessage <--
type PeerWithDraftMessage = TLUpdateDraftMessage

// SavedDialogListClazz <--
//   - TL_SavedDialogList
type SavedDialogListClazz = *TLSavedDialogList

func DecodeSavedDialogListClazz(d *bin.Decoder) (SavedDialogListClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode SavedDialogList: constructor: %w", err)
	}

	switch id {
	case 0x778fe85a:
		x := &TLSavedDialogList{ClazzID: id, ClazzName2: ClazzName_savedDialogList}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode SavedDialogList: invalid constructor %x", id)
	}

}

// TLSavedDialogList <--
type TLSavedDialogList struct {
	ClazzID    uint32                `json:"_id"`
	ClazzName2 string                `json:"_name"`
	Count      int32                 `json:"count"`
	Dialogs    []tg.SavedDialogClazz `json:"dialogs"`
}

func MakeTLSavedDialogList(m *TLSavedDialogList) *TLSavedDialogList {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_savedDialogList

	return m
}

func (m *TLSavedDialogList) String() string {
	return iface.DebugStringWithName("savedDialogList", m)
}

func (m *TLSavedDialogList) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("savedDialogList", m)
}

// SavedDialogListClazzName <--
func (m *TLSavedDialogList) SavedDialogListClazzName() string {
	return ClazzName_savedDialogList
}

// ClazzName <--
func (m *TLSavedDialogList) ClazzName() string {
	return m.ClazzName2
}

// ToSavedDialogList <--
func (m *TLSavedDialogList) ToSavedDialogList() *SavedDialogList {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLSavedDialogList) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_savedDialogList, int(layer)); clazzId {
	case 0x778fe85a:
		size := 4
		size += 4
		size += iface.CalcObjectListSize(m.Dialogs, layer)

		return size
	default:
		return 0
	}
}

func (m *TLSavedDialogList) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_savedDialogList, int(layer)); clazzId {
	case 0x778fe85a:
		if err := iface.ValidateRequiredSlice("dialogs", m.Dialogs); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode savedDialogList: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLSavedDialogList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_savedDialogList, int(layer)); clazzId {
	case 0x778fe85a:
		x.PutClazzID(0x778fe85a)

		x.PutInt32(m.Count)

		if err := iface.EncodeObjectList(x, m.Dialogs, layer); err != nil {
			return fmt.Errorf("unable to encode savedDialogList#0x778fe85a: field dialogs: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode savedDialogList: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLSavedDialogList) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x778fe85a:
		m.Count, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode savedDialogList#0x778fe85a: field count: %w", err)
		}
		l1, err3 := d.VectorHeader()
		if err3 != nil {
			return fmt.Errorf("unable to decode savedDialogList#0x778fe85a: field dialogs: %w", err3)
		}
		if l1 > bin.MaxVectorLen {
			return fmt.Errorf("unable to decode savedDialogList#0x778fe85a: field dialogs: %w", &bin.InvalidLengthError{Type: "vector", Length: int(l1)})
		}
		prealloc1 := int(l1)
		if prealloc1 > bin.PreallocateLimit {
			prealloc1 = bin.PreallocateLimit
		}
		v1 := make([]tg.SavedDialogClazz, 0, prealloc1)
		for i := int32(0); i < l1; i++ {
			vv1, err3 := tg.DecodeSavedDialogClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode savedDialogList#0x778fe85a: field dialogs: %w", err3)
			}
			v1 = append(v1, vv1)
		}
		m.Dialogs = v1

		return nil
	default:
		return fmt.Errorf("unable to decode savedDialogList: invalid constructor %x", m.ClazzID)
	}
}

// SavedDialogList <--
type SavedDialogList = TLSavedDialogList
