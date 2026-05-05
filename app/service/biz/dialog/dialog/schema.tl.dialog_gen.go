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

// DialogCursorClazz <--
//   - TL_DialogCursor
type DialogCursorClazz = *TLDialogCursor

func DecodeDialogCursorClazz(d *bin.Decoder) (DialogCursorClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode DialogCursor: constructor: %w", err)
	}

	switch id {
	case 0x5b146008:
		x := &TLDialogCursor{ClazzID: id, ClazzName2: ClazzName_dialogCursor}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode DialogCursor: invalid constructor %x", id)
	}

}

// TLDialogCursor <--
type TLDialogCursor struct {
	ClazzID               uint32 `json:"_id"`
	ClazzName2            string `json:"_name"`
	FolderId              int32  `json:"folder_id"`
	Section               string `json:"section"`
	PinnedSnapshotVersion int64  `json:"pinned_snapshot_version"`
	PinOrder              int64  `json:"pin_order"`
	TopMessageDate        int64  `json:"top_message_date"`
	TopPeerSeq            int64  `json:"top_peer_seq"`
	PeerType              int32  `json:"peer_type"`
	PeerId                int64  `json:"peer_id"`
}

func MakeTLDialogCursor(m *TLDialogCursor) *TLDialogCursor {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_dialogCursor

	return m
}

func (m *TLDialogCursor) String() string {
	return iface.DebugStringWithName("dialogCursor", m)
}

func (m *TLDialogCursor) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("dialogCursor", m)
}

// DialogCursorClazzName <--
func (m *TLDialogCursor) DialogCursorClazzName() string {
	return ClazzName_dialogCursor
}

// ClazzName <--
func (m *TLDialogCursor) ClazzName() string {
	return m.ClazzName2
}

// ToDialogCursor <--
func (m *TLDialogCursor) ToDialogCursor() *DialogCursor {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLDialogCursor) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogCursor, int(layer)); clazzId {
	case 0x5b146008:
		size := 4
		size += 4
		size += 4
		size += iface.CalcStringSize(m.Section)
		size += 8
		size += 8
		size += 8
		size += 8
		size += 4
		size += 8

		return size
	default:
		return 0
	}
}

func (m *TLDialogCursor) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogCursor, int(layer)); clazzId {
	case 0x5b146008:
		if err := iface.ValidateRequiredString("section", m.Section); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode dialogCursor: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLDialogCursor) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogCursor, int(layer)); clazzId {
	case 0x5b146008:
		x.PutClazzID(0x5b146008)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt32(m.FolderId)
		x.PutString(m.Section)
		x.PutInt64(m.PinnedSnapshotVersion)
		x.PutInt64(m.PinOrder)
		x.PutInt64(m.TopMessageDate)
		x.PutInt64(m.TopPeerSeq)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		return fmt.Errorf("unable to encode dialogCursor: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDialogCursor) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x5b146008:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode dialogCursor#0x5b146008: field flags: %w", err)
		}
		_ = flags
		m.FolderId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dialogCursor#0x5b146008: field folder_id: %w", err)
		}
		m.Section, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode dialogCursor#0x5b146008: field section: %w", err)
		}
		m.PinnedSnapshotVersion, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogCursor#0x5b146008: field pinned_snapshot_version: %w", err)
		}
		m.PinOrder, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogCursor#0x5b146008: field pin_order: %w", err)
		}
		m.TopMessageDate, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogCursor#0x5b146008: field top_message_date: %w", err)
		}
		m.TopPeerSeq, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogCursor#0x5b146008: field top_peer_seq: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dialogCursor#0x5b146008: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogCursor#0x5b146008: field peer_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dialogCursor: invalid constructor %x", m.ClazzID)
	}
}

// DialogCursor <--
type DialogCursor = TLDialogCursor

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

// DialogExtV2Clazz <--
//   - TL_DialogExtV2
type DialogExtV2Clazz = *TLDialogExtV2

func DecodeDialogExtV2Clazz(d *bin.Decoder) (DialogExtV2Clazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode DialogExtV2: constructor: %w", err)
	}

	switch id {
	case 0x7c9d7c44:
		x := &TLDialogExtV2{ClazzID: id, ClazzName2: ClazzName_dialogExtV2}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode DialogExtV2: invalid constructor %x", id)
	}

}

// TLDialogExtV2 <--
type TLDialogExtV2 struct {
	ClazzID                  uint32            `json:"_id"`
	ClazzName2               string            `json:"_name"`
	PeerType                 int32             `json:"peer_type"`
	PeerId                   int64             `json:"peer_id"`
	TopPeerSeq               int64             `json:"top_peer_seq"`
	TopCanonicalMessageId    int64             `json:"top_canonical_message_id"`
	TopMessageDate           int64             `json:"top_message_date"`
	UnreadCount              int32             `json:"unread_count"`
	UnreadMentionsCount      int32             `json:"unread_mentions_count"`
	UnreadReactionsCount     int32             `json:"unread_reactions_count"`
	UnreadMark               bool              `json:"unread_mark"`
	PinnedPeerSeq            int64             `json:"pinned_peer_seq"`
	PinnedCanonicalMessageId int64             `json:"pinned_canonical_message_id"`
	HasScheduled             bool              `json:"has_scheduled"`
	AvailableMinPeerSeq      int64             `json:"available_min_peer_seq"`
	FolderId                 int32             `json:"folder_id"`
	MainPinnedOrder          int64             `json:"main_pinned_order"`
	FolderPinnedOrder        int64             `json:"folder_pinned_order"`
	Extras                   DialogExtrasClazz `json:"extras"`
}

func MakeTLDialogExtV2(m *TLDialogExtV2) *TLDialogExtV2 {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_dialogExtV2

	return m
}

func (m *TLDialogExtV2) String() string {
	return iface.DebugStringWithName("dialogExtV2", m)
}

func (m *TLDialogExtV2) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("dialogExtV2", m)
}

// DialogExtV2ClazzName <--
func (m *TLDialogExtV2) DialogExtV2ClazzName() string {
	return ClazzName_dialogExtV2
}

// ClazzName <--
func (m *TLDialogExtV2) ClazzName() string {
	return m.ClazzName2
}

// ToDialogExtV2 <--
func (m *TLDialogExtV2) ToDialogExtV2() *DialogExtV2 {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLDialogExtV2) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogExtV2, int(layer)); clazzId {
	case 0x7c9d7c44:
		size := 4
		size += 4
		size += 4
		size += 8
		size += 8
		size += 8
		size += 8
		size += 4
		size += 4
		size += 4
		size += 8
		size += 8
		size += 8
		size += 4
		size += 8
		size += 8
		size += iface.CalcObjectSize(m.Extras, layer)

		return size
	default:
		return 0
	}
}

func (m *TLDialogExtV2) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogExtV2, int(layer)); clazzId {
	case 0x7c9d7c44:
		if err := iface.ValidateRequiredObject("extras", m.Extras); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode dialogExtV2: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLDialogExtV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogExtV2, int(layer)); clazzId {
	case 0x7c9d7c44:
		x.PutClazzID(0x7c9d7c44)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.UnreadMark == true {
				flags |= 1 << 0
			}

			if m.HasScheduled == true {
				flags |= 1 << 1
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt64(m.TopPeerSeq)
		x.PutInt64(m.TopCanonicalMessageId)
		x.PutInt64(m.TopMessageDate)
		x.PutInt32(m.UnreadCount)
		x.PutInt32(m.UnreadMentionsCount)
		x.PutInt32(m.UnreadReactionsCount)
		x.PutInt64(m.PinnedPeerSeq)
		x.PutInt64(m.PinnedCanonicalMessageId)
		x.PutInt64(m.AvailableMinPeerSeq)
		x.PutInt32(m.FolderId)
		x.PutInt64(m.MainPinnedOrder)
		x.PutInt64(m.FolderPinnedOrder)
		if m.Extras == nil {
			return fmt.Errorf("unable to encode dialogExtV2#0x7c9d7c44: field extras is nil")
		}
		if err := m.Extras.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode dialogExtV2#0x7c9d7c44: field extras: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode dialogExtV2: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDialogExtV2) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x7c9d7c44:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtV2#0x7c9d7c44: field flags: %w", err)
		}
		_ = flags
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtV2#0x7c9d7c44: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtV2#0x7c9d7c44: field peer_id: %w", err)
		}
		m.TopPeerSeq, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtV2#0x7c9d7c44: field top_peer_seq: %w", err)
		}
		m.TopCanonicalMessageId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtV2#0x7c9d7c44: field top_canonical_message_id: %w", err)
		}
		m.TopMessageDate, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtV2#0x7c9d7c44: field top_message_date: %w", err)
		}
		m.UnreadCount, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtV2#0x7c9d7c44: field unread_count: %w", err)
		}
		m.UnreadMentionsCount, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtV2#0x7c9d7c44: field unread_mentions_count: %w", err)
		}
		m.UnreadReactionsCount, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtV2#0x7c9d7c44: field unread_reactions_count: %w", err)
		}
		if (flags & (1 << 0)) != 0 {
			m.UnreadMark = true
		}
		m.PinnedPeerSeq, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtV2#0x7c9d7c44: field pinned_peer_seq: %w", err)
		}
		m.PinnedCanonicalMessageId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtV2#0x7c9d7c44: field pinned_canonical_message_id: %w", err)
		}
		if (flags & (1 << 1)) != 0 {
			m.HasScheduled = true
		}
		m.AvailableMinPeerSeq, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtV2#0x7c9d7c44: field available_min_peer_seq: %w", err)
		}
		m.FolderId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtV2#0x7c9d7c44: field folder_id: %w", err)
		}
		m.MainPinnedOrder, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtV2#0x7c9d7c44: field main_pinned_order: %w", err)
		}
		m.FolderPinnedOrder, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtV2#0x7c9d7c44: field folder_pinned_order: %w", err)
		}

		m.Extras, err = DecodeDialogExtrasClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtV2#0x7c9d7c44: field extras: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dialogExtV2: invalid constructor %x", m.ClazzID)
	}
}

// DialogExtV2 <--
type DialogExtV2 = TLDialogExtV2

// DialogExtrasClazz <--
//   - TL_DialogExtras
type DialogExtrasClazz = *TLDialogExtras

func DecodeDialogExtrasClazz(d *bin.Decoder) (DialogExtrasClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode DialogExtras: constructor: %w", err)
	}

	switch id {
	case 0x7a69125:
		x := &TLDialogExtras{ClazzID: id, ClazzName2: ClazzName_dialogExtras}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode DialogExtras: invalid constructor %x", id)
	}

}

// TLDialogExtras <--
type TLDialogExtras struct {
	ClazzID              uint32  `json:"_id"`
	ClazzName2           string  `json:"_name"`
	PeerType             int32   `json:"peer_type"`
	PeerId               int64   `json:"peer_id"`
	FolderId             int32   `json:"folder_id"`
	MainPinnedOrder      int64   `json:"main_pinned_order"`
	FolderPinnedOrder    int64   `json:"folder_pinned_order"`
	DraftPayload         []byte  `json:"draft_payload"`
	PrivateTtlPeriod     *int32  `json:"private_ttl_period"`
	PrivateThemeEmoticon *string `json:"private_theme_emoticon"`
	WallpaperId          *int64  `json:"wallpaper_id"`
	WallpaperOverridden  bool    `json:"wallpaper_overridden"`
}

func MakeTLDialogExtras(m *TLDialogExtras) *TLDialogExtras {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_dialogExtras

	return m
}

func (m *TLDialogExtras) String() string {
	return iface.DebugStringWithName("dialogExtras", m)
}

func (m *TLDialogExtras) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("dialogExtras", m)
}

// DialogExtrasClazzName <--
func (m *TLDialogExtras) DialogExtrasClazzName() string {
	return ClazzName_dialogExtras
}

// ClazzName <--
func (m *TLDialogExtras) ClazzName() string {
	return m.ClazzName2
}

// ToDialogExtras <--
func (m *TLDialogExtras) ToDialogExtras() *DialogExtras {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLDialogExtras) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogExtras, int(layer)); clazzId {
	case 0x7a69125:
		size := 4
		size += 4
		size += 4
		size += 8
		size += 4
		size += 8
		size += 8
		if m.DraftPayload != nil {
			size += iface.CalcBytesSize(m.DraftPayload)
		}

		if m.PrivateTtlPeriod != nil {
			size += 4
		}

		if m.PrivateThemeEmoticon != nil {
			size += iface.CalcStringSize(*m.PrivateThemeEmoticon)
		}

		if m.WallpaperId != nil {
			size += 8
		}

		return size
	default:
		return 0
	}
}

func (m *TLDialogExtras) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogExtras, int(layer)); clazzId {
	case 0x7a69125:

		return nil
	default:
		return fmt.Errorf("unable to encode dialogExtras: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLDialogExtras) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogExtras, int(layer)); clazzId {
	case 0x7a69125:
		x.PutClazzID(0x7a69125)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.DraftPayload != nil {
				flags |= 1 << 0
			}
			if m.PrivateTtlPeriod != nil {
				flags |= 1 << 1
			}
			if m.PrivateThemeEmoticon != nil {
				flags |= 1 << 2
			}
			if m.WallpaperId != nil {
				flags |= 1 << 3
			}
			if m.WallpaperOverridden == true {
				flags |= 1 << 4
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt32(m.FolderId)
		x.PutInt64(m.MainPinnedOrder)
		x.PutInt64(m.FolderPinnedOrder)
		if m.DraftPayload != nil {
			x.PutBytes(m.DraftPayload)
		}

		if m.PrivateTtlPeriod != nil {
			x.PutInt32(*m.PrivateTtlPeriod)
		}

		if m.PrivateThemeEmoticon != nil {
			x.PutString(*m.PrivateThemeEmoticon)
		}

		if m.WallpaperId != nil {
			x.PutInt64(*m.WallpaperId)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode dialogExtras: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDialogExtras) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x7a69125:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtras#0x7a69125: field flags: %w", err)
		}
		_ = flags
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtras#0x7a69125: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtras#0x7a69125: field peer_id: %w", err)
		}
		m.FolderId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtras#0x7a69125: field folder_id: %w", err)
		}
		m.MainPinnedOrder, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtras#0x7a69125: field main_pinned_order: %w", err)
		}
		m.FolderPinnedOrder, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogExtras#0x7a69125: field folder_pinned_order: %w", err)
		}
		if (flags & (1 << 0)) != 0 {
			m.DraftPayload, err = d.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode dialogExtras#0x7a69125: field draft_payload: %w", err)
			}
		}

		if (flags & (1 << 1)) != 0 {
			m.PrivateTtlPeriod = new(int32)
			*m.PrivateTtlPeriod, err = d.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode dialogExtras#0x7a69125: field private_ttl_period: %w", err)
			}
		}
		if (flags & (1 << 2)) != 0 {
			m.PrivateThemeEmoticon = new(string)
			*m.PrivateThemeEmoticon, err = d.String()
			if err != nil {
				return fmt.Errorf("unable to decode dialogExtras#0x7a69125: field private_theme_emoticon: %w", err)
			}
		}

		if (flags & (1 << 3)) != 0 {
			m.WallpaperId = new(int64)
			*m.WallpaperId, err = d.Int64()
			if err != nil {
				return fmt.Errorf("unable to decode dialogExtras#0x7a69125: field wallpaper_id: %w", err)
			}
		}

		if (flags & (1 << 4)) != 0 {
			m.WallpaperOverridden = true
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dialogExtras: invalid constructor %x", m.ClazzID)
	}
}

// DialogExtras <--
type DialogExtras = TLDialogExtras

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

// DialogPageClazz <--
//   - TL_DialogPage
type DialogPageClazz = *TLDialogPage

func DecodeDialogPageClazz(d *bin.Decoder) (DialogPageClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode DialogPage: constructor: %w", err)
	}

	switch id {
	case 0x92dbd5aa:
		x := &TLDialogPage{ClazzID: id, ClazzName2: ClazzName_dialogPage}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode DialogPage: invalid constructor %x", id)
	}

}

// TLDialogPage <--
type TLDialogPage struct {
	ClazzID    uint32             `json:"_id"`
	ClazzName2 string             `json:"_name"`
	Dialogs    []DialogExtV2Clazz `json:"dialogs"`
	NextCursor DialogCursorClazz  `json:"next_cursor"`
	Exhausted  tg.BoolClazz       `json:"exhausted"`
}

func MakeTLDialogPage(m *TLDialogPage) *TLDialogPage {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_dialogPage

	return m
}

func (m *TLDialogPage) String() string {
	return iface.DebugStringWithName("dialogPage", m)
}

func (m *TLDialogPage) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("dialogPage", m)
}

// DialogPageClazzName <--
func (m *TLDialogPage) DialogPageClazzName() string {
	return ClazzName_dialogPage
}

// ClazzName <--
func (m *TLDialogPage) ClazzName() string {
	return m.ClazzName2
}

// ToDialogPage <--
func (m *TLDialogPage) ToDialogPage() *DialogPage {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLDialogPage) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogPage, int(layer)); clazzId {
	case 0x92dbd5aa:
		size := 4
		size += iface.CalcObjectListSize(m.Dialogs, layer)
		size += iface.CalcObjectSize(m.NextCursor, layer)
		size += iface.CalcObjectSize(m.Exhausted, layer)

		return size
	default:
		return 0
	}
}

func (m *TLDialogPage) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogPage, int(layer)); clazzId {
	case 0x92dbd5aa:
		if err := iface.ValidateRequiredSlice("dialogs", m.Dialogs); err != nil {
			return err
		}

		if err := iface.ValidateRequiredObject("next_cursor", m.NextCursor); err != nil {
			return err
		}

		if err := iface.ValidateRequiredObject("exhausted", m.Exhausted); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode dialogPage: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLDialogPage) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogPage, int(layer)); clazzId {
	case 0x92dbd5aa:
		x.PutClazzID(0x92dbd5aa)

		if err := iface.EncodeObjectList(x, m.Dialogs, layer); err != nil {
			return fmt.Errorf("unable to encode dialogPage#0x92dbd5aa: field dialogs: %w", err)
		}

		if m.NextCursor == nil {
			return fmt.Errorf("unable to encode dialogPage#0x92dbd5aa: field next_cursor is nil")
		}
		if err := m.NextCursor.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode dialogPage#0x92dbd5aa: field next_cursor: %w", err)
		}
		if m.Exhausted == nil {
			return fmt.Errorf("unable to encode dialogPage#0x92dbd5aa: field exhausted is nil")
		}
		if err := m.Exhausted.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode dialogPage#0x92dbd5aa: field exhausted: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode dialogPage: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDialogPage) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x92dbd5aa:
		l0, err3 := d.VectorHeader()
		if err3 != nil {
			return fmt.Errorf("unable to decode dialogPage#0x92dbd5aa: field dialogs: %w", err3)
		}
		if l0 > bin.MaxVectorLen {
			return fmt.Errorf("unable to decode dialogPage#0x92dbd5aa: field dialogs: %w", &bin.InvalidLengthError{Type: "vector", Length: int(l0)})
		}
		prealloc0 := int(l0)
		if prealloc0 > bin.PreallocateLimit {
			prealloc0 = bin.PreallocateLimit
		}
		v0 := make([]DialogExtV2Clazz, 0, prealloc0)
		for i := int32(0); i < l0; i++ {
			vv0, err3 := DecodeDialogExtV2Clazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode dialogPage#0x92dbd5aa: field dialogs: %w", err3)
			}
			v0 = append(v0, vv0)
		}
		m.Dialogs = v0

		m.NextCursor, err = DecodeDialogCursorClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode dialogPage#0x92dbd5aa: field next_cursor: %w", err)
		}

		m.Exhausted, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode dialogPage#0x92dbd5aa: field exhausted: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dialogPage: invalid constructor %x", m.ClazzID)
	}
}

// DialogPage <--
type DialogPage = TLDialogPage

// DialogPeerClazz <--
//   - TL_DialogPeer
type DialogPeerClazz = *TLDialogPeer

func DecodeDialogPeerClazz(d *bin.Decoder) (DialogPeerClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode DialogPeer: constructor: %w", err)
	}

	switch id {
	case 0xb7789d79:
		x := &TLDialogPeer{ClazzID: id, ClazzName2: ClazzName_dialogPeer}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode DialogPeer: invalid constructor %x", id)
	}

}

// TLDialogPeer <--
type TLDialogPeer struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	PeerType   int32  `json:"peer_type"`
	PeerId     int64  `json:"peer_id"`
}

func MakeTLDialogPeer(m *TLDialogPeer) *TLDialogPeer {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_dialogPeer

	return m
}

func (m *TLDialogPeer) String() string {
	return iface.DebugStringWithName("dialogPeer", m)
}

func (m *TLDialogPeer) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("dialogPeer", m)
}

// DialogPeerClazzName <--
func (m *TLDialogPeer) DialogPeerClazzName() string {
	return ClazzName_dialogPeer
}

// ClazzName <--
func (m *TLDialogPeer) ClazzName() string {
	return m.ClazzName2
}

// ToDialogPeer <--
func (m *TLDialogPeer) ToDialogPeer() *DialogPeer {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLDialogPeer) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogPeer, int(layer)); clazzId {
	case 0xb7789d79:
		size := 4
		size += 4
		size += 8

		return size
	default:
		return 0
	}
}

func (m *TLDialogPeer) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogPeer, int(layer)); clazzId {
	case 0xb7789d79:

		return nil
	default:
		return fmt.Errorf("unable to encode dialogPeer: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLDialogPeer) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialogPeer, int(layer)); clazzId {
	case 0xb7789d79:
		x.PutClazzID(0xb7789d79)

		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		return fmt.Errorf("unable to encode dialogPeer: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDialogPeer) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xb7789d79:
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dialogPeer#0xb7789d79: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dialogPeer#0xb7789d79: field peer_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dialogPeer: invalid constructor %x", m.ClazzID)
	}
}

// DialogPeer <--
type DialogPeer = TLDialogPeer

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
