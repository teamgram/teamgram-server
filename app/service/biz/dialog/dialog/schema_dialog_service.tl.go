/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package dialog

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

// TLDialogSaveDraftMessage <--
type TLDialogSaveDraftMessage struct {
	ClazzID  uint32           `json:"_id"`
	UserId   int64            `json:"user_id"`
	PeerType int32            `json:"peer_type"`
	PeerId   int64            `json:"peer_id"`
	Message  *tg.DraftMessage `json:"message"`
}

func (m *TLDialogSaveDraftMessage) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogSaveDraftMessage) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x4ecad99a: func() error {
			x.PutClazzID(0x4ecad99a)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			_ = m.Message.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_saveDraftMessage, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_saveDraftMessage, layer)
	}
}

// Decode <--
func (m *TLDialogSaveDraftMessage) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x4ecad99a: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()

			m4 := &tg.DraftMessage{}
			_ = m4.Decode(d)
			m.Message = m4

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

// TLDialogClearDraftMessage <--
type TLDialogClearDraftMessage struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

func (m *TLDialogClearDraftMessage) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogClearDraftMessage) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xfb70b29a: func() error {
			x.PutClazzID(0xfb70b29a)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_clearDraftMessage, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_clearDraftMessage, layer)
	}
}

// Decode <--
func (m *TLDialogClearDraftMessage) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xfb70b29a: func() (err error) {
			m.UserId, err = d.Int64()
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

// TLDialogGetAllDrafts <--
type TLDialogGetAllDrafts struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLDialogGetAllDrafts) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetAllDrafts) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xacde4fe6: func() error {
			x.PutClazzID(0xacde4fe6)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_getAllDrafts, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getAllDrafts, layer)
	}
}

// Decode <--
func (m *TLDialogGetAllDrafts) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xacde4fe6: func() (err error) {
			m.UserId, err = d.Int64()

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

// TLDialogClearAllDrafts <--
type TLDialogClearAllDrafts struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLDialogClearAllDrafts) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogClearAllDrafts) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x41b890fc: func() error {
			x.PutClazzID(0x41b890fc)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_clearAllDrafts, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_clearAllDrafts, layer)
	}
}

// Decode <--
func (m *TLDialogClearAllDrafts) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x41b890fc: func() (err error) {
			m.UserId, err = d.Int64()

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

// TLDialogMarkDialogUnread <--
type TLDialogMarkDialogUnread struct {
	ClazzID    uint32   `json:"_id"`
	UserId     int64    `json:"user_id"`
	PeerType   int32    `json:"peer_type"`
	PeerId     int64    `json:"peer_id"`
	UnreadMark *tg.Bool `json:"unread_mark"`
}

func (m *TLDialogMarkDialogUnread) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogMarkDialogUnread) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x4532910e: func() error {
			x.PutClazzID(0x4532910e)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			_ = m.UnreadMark.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_markDialogUnread, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_markDialogUnread, layer)
	}
}

// Decode <--
func (m *TLDialogMarkDialogUnread) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x4532910e: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()

			m4 := &tg.Bool{}
			_ = m4.Decode(d)
			m.UnreadMark = m4

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

// TLDialogToggleDialogPin <--
type TLDialogToggleDialogPin struct {
	ClazzID  uint32   `json:"_id"`
	UserId   int64    `json:"user_id"`
	PeerType int32    `json:"peer_type"`
	PeerId   int64    `json:"peer_id"`
	Pinned   *tg.Bool `json:"pinned"`
}

func (m *TLDialogToggleDialogPin) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogToggleDialogPin) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x867ee52f: func() error {
			x.PutClazzID(0x867ee52f)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			_ = m.Pinned.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_toggleDialogPin, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_toggleDialogPin, layer)
	}
}

// Decode <--
func (m *TLDialogToggleDialogPin) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x867ee52f: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()

			m4 := &tg.Bool{}
			_ = m4.Decode(d)
			m.Pinned = m4

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

// TLDialogGetDialogUnreadMarkList <--
type TLDialogGetDialogUnreadMarkList struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLDialogGetDialogUnreadMarkList) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogUnreadMarkList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xcabc38f4: func() error {
			x.PutClazzID(0xcabc38f4)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogUnreadMarkList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogUnreadMarkList, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogUnreadMarkList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xcabc38f4: func() (err error) {
			m.UserId, err = d.Int64()

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

// TLDialogGetDialogsByOffsetDate <--
type TLDialogGetDialogsByOffsetDate struct {
	ClazzID       uint32   `json:"_id"`
	UserId        int64    `json:"user_id"`
	ExcludePinned *tg.Bool `json:"exclude_pinned"`
	OffsetDate    int32    `json:"offset_date"`
	Limit         int32    `json:"limit"`
}

func (m *TLDialogGetDialogsByOffsetDate) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogsByOffsetDate) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x9d7e8604: func() error {
			x.PutClazzID(0x9d7e8604)

			x.PutInt64(m.UserId)
			_ = m.ExcludePinned.Encode(x, layer)
			x.PutInt32(m.OffsetDate)
			x.PutInt32(m.Limit)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogsByOffsetDate, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogsByOffsetDate, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogsByOffsetDate) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x9d7e8604: func() (err error) {
			m.UserId, err = d.Int64()

			m2 := &tg.Bool{}
			_ = m2.Decode(d)
			m.ExcludePinned = m2

			m.OffsetDate, err = d.Int32()
			m.Limit, err = d.Int32()

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

// TLDialogGetDialogs <--
type TLDialogGetDialogs struct {
	ClazzID       uint32   `json:"_id"`
	UserId        int64    `json:"user_id"`
	ExcludePinned *tg.Bool `json:"exclude_pinned"`
	FolderId      int32    `json:"folder_id"`
}

func (m *TLDialogGetDialogs) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogs) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x860b1e16: func() error {
			x.PutClazzID(0x860b1e16)

			x.PutInt64(m.UserId)
			_ = m.ExcludePinned.Encode(x, layer)
			x.PutInt32(m.FolderId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogs, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogs, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogs) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x860b1e16: func() (err error) {
			m.UserId, err = d.Int64()

			m2 := &tg.Bool{}
			_ = m2.Decode(d)
			m.ExcludePinned = m2

			m.FolderId, err = d.Int32()

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

// TLDialogGetDialogsByIdList <--
type TLDialogGetDialogsByIdList struct {
	ClazzID uint32  `json:"_id"`
	UserId  int64   `json:"user_id"`
	IdList  []int64 `json:"id_list"`
}

func (m *TLDialogGetDialogsByIdList) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogsByIdList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xad258871: func() error {
			x.PutClazzID(0xad258871)

			x.PutInt64(m.UserId)

			iface.EncodeInt64List(x, m.IdList)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogsByIdList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogsByIdList, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogsByIdList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xad258871: func() (err error) {
			m.UserId, err = d.Int64()

			m.IdList, err = iface.DecodeInt64List(d)

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

// TLDialogGetDialogsCount <--
type TLDialogGetDialogsCount struct {
	ClazzID       uint32   `json:"_id"`
	UserId        int64    `json:"user_id"`
	ExcludePinned *tg.Bool `json:"exclude_pinned"`
	FolderId      int32    `json:"folder_id"`
}

func (m *TLDialogGetDialogsCount) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogsCount) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xe039b465: func() error {
			x.PutClazzID(0xe039b465)

			x.PutInt64(m.UserId)
			_ = m.ExcludePinned.Encode(x, layer)
			x.PutInt32(m.FolderId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogsCount, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogsCount, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogsCount) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xe039b465: func() (err error) {
			m.UserId, err = d.Int64()

			m2 := &tg.Bool{}
			_ = m2.Decode(d)
			m.ExcludePinned = m2

			m.FolderId, err = d.Int32()

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

// TLDialogGetPinnedDialogs <--
type TLDialogGetPinnedDialogs struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	FolderId int32  `json:"folder_id"`
}

func (m *TLDialogGetPinnedDialogs) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetPinnedDialogs) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa8c21bb5: func() error {
			x.PutClazzID(0xa8c21bb5)

			x.PutInt64(m.UserId)
			x.PutInt32(m.FolderId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_getPinnedDialogs, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getPinnedDialogs, layer)
	}
}

// Decode <--
func (m *TLDialogGetPinnedDialogs) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa8c21bb5: func() (err error) {
			m.UserId, err = d.Int64()
			m.FolderId, err = d.Int32()

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

// TLDialogReorderPinnedDialogs <--
type TLDialogReorderPinnedDialogs struct {
	ClazzID  uint32   `json:"_id"`
	UserId   int64    `json:"user_id"`
	Force    *tg.Bool `json:"force"`
	FolderId int32    `json:"folder_id"`
	IdList   []int64  `json:"id_list"`
}

func (m *TLDialogReorderPinnedDialogs) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogReorderPinnedDialogs) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xfee33567: func() error {
			x.PutClazzID(0xfee33567)

			x.PutInt64(m.UserId)
			_ = m.Force.Encode(x, layer)
			x.PutInt32(m.FolderId)

			iface.EncodeInt64List(x, m.IdList)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_reorderPinnedDialogs, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_reorderPinnedDialogs, layer)
	}
}

// Decode <--
func (m *TLDialogReorderPinnedDialogs) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xfee33567: func() (err error) {
			m.UserId, err = d.Int64()

			m2 := &tg.Bool{}
			_ = m2.Decode(d)
			m.Force = m2

			m.FolderId, err = d.Int32()

			m.IdList, err = iface.DecodeInt64List(d)

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

// TLDialogGetDialogById <--
type TLDialogGetDialogById struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

func (m *TLDialogGetDialogById) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogById) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa15f3bf5: func() error {
			x.PutClazzID(0xa15f3bf5)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogById, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogById, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogById) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa15f3bf5: func() (err error) {
			m.UserId, err = d.Int64()
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

// TLDialogGetTopMessage <--
type TLDialogGetTopMessage struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

func (m *TLDialogGetTopMessage) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetTopMessage) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xfa7db272: func() error {
			x.PutClazzID(0xfa7db272)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_getTopMessage, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getTopMessage, layer)
	}
}

// Decode <--
func (m *TLDialogGetTopMessage) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xfa7db272: func() (err error) {
			m.UserId, err = d.Int64()
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

// TLDialogInsertOrUpdateDialog <--
type TLDialogInsertOrUpdateDialog struct {
	ClazzID         uint32 `json:"_id"`
	UserId          int64  `json:"user_id"`
	PeerType        int32  `json:"peer_type"`
	PeerId          int64  `json:"peer_id"`
	TopMessage      *int32 `json:"top_message"`
	ReadOutboxMaxId *int32 `json:"read_outbox_max_id"`
	ReadInboxMaxId  *int32 `json:"read_inbox_max_id"`
	UnreadCount     *int32 `json:"unread_count"`
	UnreadMark      bool   `json:"unread_mark"`
	Date2           *int64 `json:"date2"`
	PinnedMsgId     *int32 `json:"pinned_msg_id"`
}

func (m *TLDialogInsertOrUpdateDialog) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogInsertOrUpdateDialog) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x5d2b8822: func() error {
			x.PutClazzID(0x5d2b8822)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.TopMessage != nil {
					flags |= 1 << 0
				}
				if m.ReadOutboxMaxId != nil {
					flags |= 1 << 1
				}
				if m.ReadInboxMaxId != nil {
					flags |= 1 << 2
				}
				if m.UnreadCount != nil {
					flags |= 1 << 3
				}
				if m.UnreadMark == true {
					flags |= 1 << 4
				}
				if m.Date2 != nil {
					flags |= 1 << 5
				}
				if m.PinnedMsgId != nil {
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
			if m.TopMessage != nil {
				x.PutInt32(*m.TopMessage)
			}

			if m.ReadOutboxMaxId != nil {
				x.PutInt32(*m.ReadOutboxMaxId)
			}

			if m.ReadInboxMaxId != nil {
				x.PutInt32(*m.ReadInboxMaxId)
			}

			if m.UnreadCount != nil {
				x.PutInt32(*m.UnreadCount)
			}

			if m.Date2 != nil {
				x.PutInt64(*m.Date2)
			}

			if m.PinnedMsgId != nil {
				x.PutInt32(*m.PinnedMsgId)
			}

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_insertOrUpdateDialog, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_insertOrUpdateDialog, layer)
	}
}

// Decode <--
func (m *TLDialogInsertOrUpdateDialog) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x5d2b8822: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			if (flags & (1 << 0)) != 0 {
				m.TopMessage = new(int32)
				*m.TopMessage, err = d.Int32()
			}
			if (flags & (1 << 1)) != 0 {
				m.ReadOutboxMaxId = new(int32)
				*m.ReadOutboxMaxId, err = d.Int32()
			}
			if (flags & (1 << 2)) != 0 {
				m.ReadInboxMaxId = new(int32)
				*m.ReadInboxMaxId, err = d.Int32()
			}
			if (flags & (1 << 3)) != 0 {
				m.UnreadCount = new(int32)
				*m.UnreadCount, err = d.Int32()
			}
			if (flags & (1 << 4)) != 0 {
				m.UnreadMark = true
			}
			if (flags & (1 << 5)) != 0 {
				m.Date2 = new(int64)
				*m.Date2, err = d.Int64()
			}

			if (flags & (1 << 6)) != 0 {
				m.PinnedMsgId = new(int32)
				*m.PinnedMsgId, err = d.Int32()
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

// TLDialogDeleteDialog <--
type TLDialogDeleteDialog struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

func (m *TLDialogDeleteDialog) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogDeleteDialog) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x1b31de3: func() error {
			x.PutClazzID(0x1b31de3)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_deleteDialog, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_deleteDialog, layer)
	}
}

// Decode <--
func (m *TLDialogDeleteDialog) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x1b31de3: func() (err error) {
			m.UserId, err = d.Int64()
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

// TLDialogGetUserPinnedMessage <--
type TLDialogGetUserPinnedMessage struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

func (m *TLDialogGetUserPinnedMessage) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetUserPinnedMessage) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x8f9bc2b1: func() error {
			x.PutClazzID(0x8f9bc2b1)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_getUserPinnedMessage, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getUserPinnedMessage, layer)
	}
}

// Decode <--
func (m *TLDialogGetUserPinnedMessage) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x8f9bc2b1: func() (err error) {
			m.UserId, err = d.Int64()
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

// TLDialogUpdateUserPinnedMessage <--
type TLDialogUpdateUserPinnedMessage struct {
	ClazzID     uint32 `json:"_id"`
	UserId      int64  `json:"user_id"`
	PeerType    int32  `json:"peer_type"`
	PeerId      int64  `json:"peer_id"`
	PinnedMsgId int32  `json:"pinned_msg_id"`
}

func (m *TLDialogUpdateUserPinnedMessage) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogUpdateUserPinnedMessage) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x1622f22a: func() error {
			x.PutClazzID(0x1622f22a)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			x.PutInt32(m.PinnedMsgId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_updateUserPinnedMessage, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_updateUserPinnedMessage, layer)
	}
}

// Decode <--
func (m *TLDialogUpdateUserPinnedMessage) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x1622f22a: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.PinnedMsgId, err = d.Int32()

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

// TLDialogInsertOrUpdateDialogFilter <--
type TLDialogInsertOrUpdateDialogFilter struct {
	ClazzID      uint32           `json:"_id"`
	UserId       int64            `json:"user_id"`
	Id           int32            `json:"id"`
	DialogFilter *tg.DialogFilter `json:"dialog_filter"`
}

func (m *TLDialogInsertOrUpdateDialogFilter) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogInsertOrUpdateDialogFilter) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xaa8a384: func() error {
			x.PutClazzID(0xaa8a384)

			x.PutInt64(m.UserId)
			x.PutInt32(m.Id)
			_ = m.DialogFilter.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_insertOrUpdateDialogFilter, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_insertOrUpdateDialogFilter, layer)
	}
}

// Decode <--
func (m *TLDialogInsertOrUpdateDialogFilter) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xaa8a384: func() (err error) {
			m.UserId, err = d.Int64()
			m.Id, err = d.Int32()

			m3 := &tg.DialogFilter{}
			_ = m3.Decode(d)
			m.DialogFilter = m3

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

// TLDialogDeleteDialogFilter <--
type TLDialogDeleteDialogFilter struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Id      int32  `json:"id"`
}

func (m *TLDialogDeleteDialogFilter) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogDeleteDialogFilter) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x1dd3e97: func() error {
			x.PutClazzID(0x1dd3e97)

			x.PutInt64(m.UserId)
			x.PutInt32(m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_deleteDialogFilter, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_deleteDialogFilter, layer)
	}
}

// Decode <--
func (m *TLDialogDeleteDialogFilter) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x1dd3e97: func() (err error) {
			m.UserId, err = d.Int64()
			m.Id, err = d.Int32()

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

// TLDialogUpdateDialogFiltersOrder <--
type TLDialogUpdateDialogFiltersOrder struct {
	ClazzID uint32  `json:"_id"`
	UserId  int64   `json:"user_id"`
	Order   []int32 `json:"order"`
}

func (m *TLDialogUpdateDialogFiltersOrder) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogUpdateDialogFiltersOrder) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xb13c0b3f: func() error {
			x.PutClazzID(0xb13c0b3f)

			x.PutInt64(m.UserId)

			iface.EncodeInt32List(x, m.Order)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_updateDialogFiltersOrder, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_updateDialogFiltersOrder, layer)
	}
}

// Decode <--
func (m *TLDialogUpdateDialogFiltersOrder) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xb13c0b3f: func() (err error) {
			m.UserId, err = d.Int64()

			m.Order, err = iface.DecodeInt32List(d)

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

// TLDialogGetDialogFilters <--
type TLDialogGetDialogFilters struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLDialogGetDialogFilters) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogFilters) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x6c676c3c: func() error {
			x.PutClazzID(0x6c676c3c)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogFilters, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogFilters, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogFilters) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x6c676c3c: func() (err error) {
			m.UserId, err = d.Int64()

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

// TLDialogGetDialogFolder <--
type TLDialogGetDialogFolder struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	FolderId int32  `json:"folder_id"`
}

func (m *TLDialogGetDialogFolder) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogFolder) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x411b8eb5: func() error {
			x.PutClazzID(0x411b8eb5)

			x.PutInt64(m.UserId)
			x.PutInt32(m.FolderId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogFolder, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogFolder, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogFolder) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x411b8eb5: func() (err error) {
			m.UserId, err = d.Int64()
			m.FolderId, err = d.Int32()

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

// TLDialogEditPeerFolders <--
type TLDialogEditPeerFolders struct {
	ClazzID        uint32  `json:"_id"`
	UserId         int64   `json:"user_id"`
	PeerDialogList []int64 `json:"peer_dialog_list"`
	FolderId       int32   `json:"folder_id"`
}

func (m *TLDialogEditPeerFolders) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogEditPeerFolders) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x2446869a: func() error {
			x.PutClazzID(0x2446869a)

			x.PutInt64(m.UserId)

			iface.EncodeInt64List(x, m.PeerDialogList)

			x.PutInt32(m.FolderId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_editPeerFolders, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_editPeerFolders, layer)
	}
}

// Decode <--
func (m *TLDialogEditPeerFolders) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x2446869a: func() (err error) {
			m.UserId, err = d.Int64()

			m.PeerDialogList, err = iface.DecodeInt64List(d)

			m.FolderId, err = d.Int32()

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

// TLDialogGetChannelMessageReadParticipants <--
type TLDialogGetChannelMessageReadParticipants struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	ChannelId int64  `json:"channel_id"`
	MsgId     int32  `json:"msg_id"`
}

func (m *TLDialogGetChannelMessageReadParticipants) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetChannelMessageReadParticipants) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x28bd4d3b: func() error {
			x.PutClazzID(0x28bd4d3b)

			x.PutInt64(m.UserId)
			x.PutInt64(m.ChannelId)
			x.PutInt32(m.MsgId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_getChannelMessageReadParticipants, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getChannelMessageReadParticipants, layer)
	}
}

// Decode <--
func (m *TLDialogGetChannelMessageReadParticipants) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x28bd4d3b: func() (err error) {
			m.UserId, err = d.Int64()
			m.ChannelId, err = d.Int64()
			m.MsgId, err = d.Int32()

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

// TLDialogSetChatTheme <--
type TLDialogSetChatTheme struct {
	ClazzID       uint32 `json:"_id"`
	UserId        int64  `json:"user_id"`
	PeerType      int32  `json:"peer_type"`
	PeerId        int64  `json:"peer_id"`
	ThemeEmoticon string `json:"theme_emoticon"`
}

func (m *TLDialogSetChatTheme) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogSetChatTheme) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xe9aea22a: func() error {
			x.PutClazzID(0xe9aea22a)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			x.PutString(m.ThemeEmoticon)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_setChatTheme, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_setChatTheme, layer)
	}
}

// Decode <--
func (m *TLDialogSetChatTheme) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xe9aea22a: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.ThemeEmoticon, err = d.String()

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

// TLDialogSetHistoryTTL <--
type TLDialogSetHistoryTTL struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	PeerType  int32  `json:"peer_type"`
	PeerId    int64  `json:"peer_id"`
	TtlPeriod int32  `json:"ttl_period"`
}

func (m *TLDialogSetHistoryTTL) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogSetHistoryTTL) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x9d9b8ac: func() error {
			x.PutClazzID(0x9d9b8ac)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			x.PutInt32(m.TtlPeriod)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_setHistoryTTL, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_setHistoryTTL, layer)
	}
}

// Decode <--
func (m *TLDialogSetHistoryTTL) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x9d9b8ac: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.TtlPeriod, err = d.Int32()

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

// TLDialogGetMyDialogsData <--
type TLDialogGetMyDialogsData struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	User    bool   `json:"user"`
	Chat    bool   `json:"chat"`
	Channel bool   `json:"channel"`
}

func (m *TLDialogGetMyDialogsData) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetMyDialogsData) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x7ee08f03: func() error {
			x.PutClazzID(0x7ee08f03)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.User == true {
					flags |= 1 << 0
				}
				if m.Chat == true {
					flags |= 1 << 1
				}
				if m.Channel == true {
					flags |= 1 << 2
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_getMyDialogsData, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getMyDialogsData, layer)
	}
}

// Decode <--
func (m *TLDialogGetMyDialogsData) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x7ee08f03: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.UserId, err = d.Int64()
			if (flags & (1 << 0)) != 0 {
				m.User = true
			}
			if (flags & (1 << 1)) != 0 {
				m.Chat = true
			}
			if (flags & (1 << 2)) != 0 {
				m.Channel = true
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

// TLDialogGetSavedDialogs <--
type TLDialogGetSavedDialogs struct {
	ClazzID       uint32       `json:"_id"`
	UserId        int64        `json:"user_id"`
	ExcludePinned *tg.Bool     `json:"exclude_pinned"`
	OffsetDate    int32        `json:"offset_date"`
	OffsetId      int32        `json:"offset_id"`
	OffsetPeer    *tg.PeerUtil `json:"offset_peer"`
	Limit         int32        `json:"limit"`
}

func (m *TLDialogGetSavedDialogs) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetSavedDialogs) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x38c1d668: func() error {
			x.PutClazzID(0x38c1d668)

			x.PutInt64(m.UserId)
			_ = m.ExcludePinned.Encode(x, layer)
			x.PutInt32(m.OffsetDate)
			x.PutInt32(m.OffsetId)
			_ = m.OffsetPeer.Encode(x, layer)
			x.PutInt32(m.Limit)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_getSavedDialogs, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getSavedDialogs, layer)
	}
}

// Decode <--
func (m *TLDialogGetSavedDialogs) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x38c1d668: func() (err error) {
			m.UserId, err = d.Int64()

			m2 := &tg.Bool{}
			_ = m2.Decode(d)
			m.ExcludePinned = m2

			m.OffsetDate, err = d.Int32()
			m.OffsetId, err = d.Int32()

			m5 := &tg.PeerUtil{}
			_ = m5.Decode(d)
			m.OffsetPeer = m5

			m.Limit, err = d.Int32()

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

// TLDialogGetPinnedSavedDialogs <--
type TLDialogGetPinnedSavedDialogs struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLDialogGetPinnedSavedDialogs) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetPinnedSavedDialogs) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x40a3b7e7: func() error {
			x.PutClazzID(0x40a3b7e7)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_getPinnedSavedDialogs, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getPinnedSavedDialogs, layer)
	}
}

// Decode <--
func (m *TLDialogGetPinnedSavedDialogs) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x40a3b7e7: func() (err error) {
			m.UserId, err = d.Int64()

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

// TLDialogToggleSavedDialogPin <--
type TLDialogToggleSavedDialogPin struct {
	ClazzID uint32       `json:"_id"`
	UserId  int64        `json:"user_id"`
	Peer    *tg.PeerUtil `json:"peer"`
	Pinned  *tg.Bool     `json:"pinned"`
}

func (m *TLDialogToggleSavedDialogPin) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogToggleSavedDialogPin) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x44f317d9: func() error {
			x.PutClazzID(0x44f317d9)

			x.PutInt64(m.UserId)
			_ = m.Peer.Encode(x, layer)
			_ = m.Pinned.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_toggleSavedDialogPin, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_toggleSavedDialogPin, layer)
	}
}

// Decode <--
func (m *TLDialogToggleSavedDialogPin) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x44f317d9: func() (err error) {
			m.UserId, err = d.Int64()

			m2 := &tg.PeerUtil{}
			_ = m2.Decode(d)
			m.Peer = m2

			m3 := &tg.Bool{}
			_ = m3.Decode(d)
			m.Pinned = m3

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

// TLDialogReorderPinnedSavedDialogs <--
type TLDialogReorderPinnedSavedDialogs struct {
	ClazzID uint32         `json:"_id"`
	UserId  int64          `json:"user_id"`
	Force   *tg.Bool       `json:"force"`
	Order   []*tg.PeerUtil `json:"order"`
}

func (m *TLDialogReorderPinnedSavedDialogs) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogReorderPinnedSavedDialogs) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xd85ccbd2: func() error {
			x.PutClazzID(0xd85ccbd2)

			x.PutInt64(m.UserId)
			_ = m.Force.Encode(x, layer)

			_ = iface.EncodeObjectList(x, m.Order, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_reorderPinnedSavedDialogs, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_reorderPinnedSavedDialogs, layer)
	}
}

// Decode <--
func (m *TLDialogReorderPinnedSavedDialogs) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xd85ccbd2: func() (err error) {
			m.UserId, err = d.Int64()

			m2 := &tg.Bool{}
			_ = m2.Decode(d)
			m.Force = m2

			c3, err2 := d.ClazzID()
			if c3 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 3, c3)
				return err2
			}
			l3, err3 := d.Int()
			v3 := make([]*tg.PeerUtil, l3)
			for i := 0; i < l3; i++ {
				vv := new(tg.PeerUtil)
				err3 = vv.Decode(d)
				_ = err3
				v3[i] = vv
			}
			m.Order = v3

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

// TLDialogGetDialogFilter <--
type TLDialogGetDialogFilter struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Id      int32  `json:"id"`
}

func (m *TLDialogGetDialogFilter) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogFilter) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xf388061c: func() error {
			x.PutClazzID(0xf388061c)

			x.PutInt64(m.UserId)
			x.PutInt32(m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogFilter, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogFilter, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogFilter) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xf388061c: func() (err error) {
			m.UserId, err = d.Int64()
			m.Id, err = d.Int32()

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

// TLDialogGetDialogFilterBySlug <--
type TLDialogGetDialogFilterBySlug struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Slug    string `json:"slug"`
}

func (m *TLDialogGetDialogFilterBySlug) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogFilterBySlug) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x4e457fef: func() error {
			x.PutClazzID(0x4e457fef)

			x.PutInt64(m.UserId)
			x.PutString(m.Slug)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogFilterBySlug, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogFilterBySlug, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogFilterBySlug) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x4e457fef: func() (err error) {
			m.UserId, err = d.Int64()
			m.Slug, err = d.String()

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

// TLDialogCreateDialogFilter <--
type TLDialogCreateDialogFilter struct {
	ClazzID      uint32           `json:"_id"`
	UserId       int64            `json:"user_id"`
	DialogFilter *DialogFilterExt `json:"dialog_filter"`
}

func (m *TLDialogCreateDialogFilter) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogCreateDialogFilter) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc6cb636f: func() error {
			x.PutClazzID(0xc6cb636f)

			x.PutInt64(m.UserId)
			_ = m.DialogFilter.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_createDialogFilter, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_createDialogFilter, layer)
	}
}

// Decode <--
func (m *TLDialogCreateDialogFilter) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xc6cb636f: func() (err error) {
			m.UserId, err = d.Int64()

			m2 := &DialogFilterExt{}
			_ = m2.Decode(d)
			m.DialogFilter = m2

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

// TLDialogUpdateUnreadCount <--
type TLDialogUpdateUnreadCount struct {
	ClazzID              uint32 `json:"_id"`
	UserId               int64  `json:"user_id"`
	PeerType             int32  `json:"peer_type"`
	PeerId               int64  `json:"peer_id"`
	UnreadCount          *int32 `json:"unread_count"`
	UnreadMentionsCount  *int32 `json:"unread_mentions_count"`
	UnreadReactionsCount *int32 `json:"unread_reactions_count"`
}

func (m *TLDialogUpdateUnreadCount) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogUpdateUnreadCount) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x2bac334d: func() error {
			x.PutClazzID(0x2bac334d)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.UnreadCount != nil {
					flags |= 1 << 0
				}
				if m.UnreadMentionsCount != nil {
					flags |= 1 << 1
				}
				if m.UnreadReactionsCount != nil {
					flags |= 1 << 2
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			if m.UnreadCount != nil {
				x.PutInt32(*m.UnreadCount)
			}

			if m.UnreadMentionsCount != nil {
				x.PutInt32(*m.UnreadMentionsCount)
			}

			if m.UnreadReactionsCount != nil {
				x.PutInt32(*m.UnreadReactionsCount)
			}

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_updateUnreadCount, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_updateUnreadCount, layer)
	}
}

// Decode <--
func (m *TLDialogUpdateUnreadCount) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x2bac334d: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			if (flags & (1 << 0)) != 0 {
				m.UnreadCount = new(int32)
				*m.UnreadCount, err = d.Int32()
			}
			if (flags & (1 << 1)) != 0 {
				m.UnreadMentionsCount = new(int32)
				*m.UnreadMentionsCount, err = d.Int32()
			}
			if (flags & (1 << 2)) != 0 {
				m.UnreadReactionsCount = new(int32)
				*m.UnreadReactionsCount, err = d.Int32()
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

// TLDialogToggleDialogFilterTags <--
type TLDialogToggleDialogFilterTags struct {
	ClazzID uint32   `json:"_id"`
	UserId  int64    `json:"user_id"`
	Enabled *tg.Bool `json:"enabled"`
}

func (m *TLDialogToggleDialogFilterTags) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogToggleDialogFilterTags) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa0cd6d89: func() error {
			x.PutClazzID(0xa0cd6d89)

			x.PutInt64(m.UserId)
			_ = m.Enabled.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_toggleDialogFilterTags, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_toggleDialogFilterTags, layer)
	}
}

// Decode <--
func (m *TLDialogToggleDialogFilterTags) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa0cd6d89: func() (err error) {
			m.UserId, err = d.Int64()

			m2 := &tg.Bool{}
			_ = m2.Decode(d)
			m.Enabled = m2

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

// TLDialogGetDialogFilterTags <--
type TLDialogGetDialogFilterTags struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLDialogGetDialogFilterTags) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogFilterTags) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xfaf0fa97: func() error {
			x.PutClazzID(0xfaf0fa97)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogFilterTags, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogFilterTags, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogFilterTags) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xfaf0fa97: func() (err error) {
			m.UserId, err = d.Int64()

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

// TLDialogSetChatWallpaper <--
type TLDialogSetChatWallpaper struct {
	ClazzID             uint32 `json:"_id"`
	UserId              int64  `json:"user_id"`
	PeerType            int32  `json:"peer_type"`
	PeerId              int64  `json:"peer_id"`
	WallpaperId         int64  `json:"wallpaper_id"`
	WallpaperOverridden bool   `json:"wallpaper_overridden"`
}

func (m *TLDialogSetChatWallpaper) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogSetChatWallpaper) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xb551db12: func() error {
			x.PutClazzID(0xb551db12)

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
			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			x.PutInt64(m.WallpaperId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialog_setChatWallpaper, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_setChatWallpaper, layer)
	}
}

// Decode <--
func (m *TLDialogSetChatWallpaper) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xb551db12: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.WallpaperId, err = d.Int64()
			if (flags & (1 << 0)) != 0 {
				m.WallpaperOverridden = true
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

// VectorPeerWithDraftMessage <--
type VectorPeerWithDraftMessage struct {
	Datas []*PeerWithDraftMessage `json:"_datas"`
}

func (m *VectorPeerWithDraftMessage) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorPeerWithDraftMessage) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorPeerWithDraftMessage) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*PeerWithDraftMessage](d)

	return err
}

// VectorDialogPeer <--
type VectorDialogPeer struct {
	Datas []*tg.DialogPeer `json:"_datas"`
}

func (m *VectorDialogPeer) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorDialogPeer) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorDialogPeer) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*tg.DialogPeer](d)

	return err
}

// VectorDialogExt <--
type VectorDialogExt struct {
	Datas []*DialogExt `json:"_datas"`
}

func (m *VectorDialogExt) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorDialogExt) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorDialogExt) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*DialogExt](d)

	return err
}

// VectorDialogFilterExt <--
type VectorDialogFilterExt struct {
	Datas []*DialogFilterExt `json:"_datas"`
}

func (m *VectorDialogFilterExt) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorDialogFilterExt) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorDialogFilterExt) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*DialogFilterExt](d)

	return err
}

// VectorDialogPinnedExt <--
type VectorDialogPinnedExt struct {
	Datas []*DialogPinnedExt `json:"_datas"`
}

func (m *VectorDialogPinnedExt) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorDialogPinnedExt) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorDialogPinnedExt) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*DialogPinnedExt](d)

	return err
}

// VectorLong <--
type VectorLong struct {
	Datas []int64 `json:"_datas"`
}

func (m *VectorLong) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorLong) Encode(x *bin.Encoder, layer int32) error {
	iface.EncodeInt64List(x, m.Datas)

	return nil
}

// Decode <--
func (m *VectorLong) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeInt64List(d)

	return err
}

// ----------------------------------------------------------------------------
// rpc

type RPCDialog interface {
	DialogSaveDraftMessage(ctx context.Context, in *TLDialogSaveDraftMessage) (*tg.Bool, error)
	DialogClearDraftMessage(ctx context.Context, in *TLDialogClearDraftMessage) (*tg.Bool, error)
	DialogGetAllDrafts(ctx context.Context, in *TLDialogGetAllDrafts) (*VectorPeerWithDraftMessage, error)
	DialogClearAllDrafts(ctx context.Context, in *TLDialogClearAllDrafts) (*VectorPeerWithDraftMessage, error)
	DialogMarkDialogUnread(ctx context.Context, in *TLDialogMarkDialogUnread) (*tg.Bool, error)
	DialogToggleDialogPin(ctx context.Context, in *TLDialogToggleDialogPin) (*tg.Int32, error)
	DialogGetDialogUnreadMarkList(ctx context.Context, in *TLDialogGetDialogUnreadMarkList) (*VectorDialogPeer, error)
	DialogGetDialogsByOffsetDate(ctx context.Context, in *TLDialogGetDialogsByOffsetDate) (*VectorDialogExt, error)
	DialogGetDialogs(ctx context.Context, in *TLDialogGetDialogs) (*VectorDialogExt, error)
	DialogGetDialogsByIdList(ctx context.Context, in *TLDialogGetDialogsByIdList) (*VectorDialogExt, error)
	DialogGetDialogsCount(ctx context.Context, in *TLDialogGetDialogsCount) (*tg.Int32, error)
	DialogGetPinnedDialogs(ctx context.Context, in *TLDialogGetPinnedDialogs) (*VectorDialogExt, error)
	DialogReorderPinnedDialogs(ctx context.Context, in *TLDialogReorderPinnedDialogs) (*tg.Bool, error)
	DialogGetDialogById(ctx context.Context, in *TLDialogGetDialogById) (*DialogExt, error)
	DialogGetTopMessage(ctx context.Context, in *TLDialogGetTopMessage) (*tg.Int32, error)
	DialogInsertOrUpdateDialog(ctx context.Context, in *TLDialogInsertOrUpdateDialog) (*tg.Bool, error)
	DialogDeleteDialog(ctx context.Context, in *TLDialogDeleteDialog) (*tg.Bool, error)
	DialogGetUserPinnedMessage(ctx context.Context, in *TLDialogGetUserPinnedMessage) (*tg.Int32, error)
	DialogUpdateUserPinnedMessage(ctx context.Context, in *TLDialogUpdateUserPinnedMessage) (*tg.Bool, error)
	DialogInsertOrUpdateDialogFilter(ctx context.Context, in *TLDialogInsertOrUpdateDialogFilter) (*tg.Bool, error)
	DialogDeleteDialogFilter(ctx context.Context, in *TLDialogDeleteDialogFilter) (*tg.Bool, error)
	DialogUpdateDialogFiltersOrder(ctx context.Context, in *TLDialogUpdateDialogFiltersOrder) (*tg.Bool, error)
	DialogGetDialogFilters(ctx context.Context, in *TLDialogGetDialogFilters) (*VectorDialogFilterExt, error)
	DialogGetDialogFolder(ctx context.Context, in *TLDialogGetDialogFolder) (*VectorDialogExt, error)
	DialogEditPeerFolders(ctx context.Context, in *TLDialogEditPeerFolders) (*VectorDialogPinnedExt, error)
	DialogGetChannelMessageReadParticipants(ctx context.Context, in *TLDialogGetChannelMessageReadParticipants) (*VectorLong, error)
	DialogSetChatTheme(ctx context.Context, in *TLDialogSetChatTheme) (*tg.Bool, error)
	DialogSetHistoryTTL(ctx context.Context, in *TLDialogSetHistoryTTL) (*tg.Bool, error)
	DialogGetMyDialogsData(ctx context.Context, in *TLDialogGetMyDialogsData) (*DialogsData, error)
	DialogGetSavedDialogs(ctx context.Context, in *TLDialogGetSavedDialogs) (*SavedDialogList, error)
	DialogGetPinnedSavedDialogs(ctx context.Context, in *TLDialogGetPinnedSavedDialogs) (*SavedDialogList, error)
	DialogToggleSavedDialogPin(ctx context.Context, in *TLDialogToggleSavedDialogPin) (*tg.Bool, error)
	DialogReorderPinnedSavedDialogs(ctx context.Context, in *TLDialogReorderPinnedSavedDialogs) (*tg.Bool, error)
	DialogGetDialogFilter(ctx context.Context, in *TLDialogGetDialogFilter) (*DialogFilterExt, error)
	DialogGetDialogFilterBySlug(ctx context.Context, in *TLDialogGetDialogFilterBySlug) (*DialogFilterExt, error)
	DialogCreateDialogFilter(ctx context.Context, in *TLDialogCreateDialogFilter) (*DialogFilterExt, error)
	DialogUpdateUnreadCount(ctx context.Context, in *TLDialogUpdateUnreadCount) (*tg.Bool, error)
	DialogToggleDialogFilterTags(ctx context.Context, in *TLDialogToggleDialogFilterTags) (*tg.Bool, error)
	DialogGetDialogFilterTags(ctx context.Context, in *TLDialogGetDialogFilterTags) (*tg.Bool, error)
	DialogSetChatWallpaper(ctx context.Context, in *TLDialogSetChatWallpaper) (*tg.Bool, error)
}
