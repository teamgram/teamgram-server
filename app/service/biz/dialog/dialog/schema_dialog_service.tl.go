/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dialog

import (
	"context"
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

// TLDialogSaveDraftMessage <--
type TLDialogSaveDraftMessage struct {
	ClazzID  uint32               `json:"_id"`
	UserId   int64                `json:"user_id"`
	PeerType int32                `json:"peer_type"`
	PeerId   int64                `json:"peer_id"`
	Message  tg.DraftMessageClazz `json:"message"`
}

func (m *TLDialogSaveDraftMessage) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_saveDraftMessage, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogSaveDraftMessage) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_saveDraftMessage, int(layer)); clazzId {
	case 0x4ecad99a:
		x.PutClazzID(0x4ecad99a)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		_ = m.Message.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_saveDraftMessage, layer)
	}
}

// Decode <--
func (m *TLDialogSaveDraftMessage) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x4ecad99a:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}

		m.Message, err = tg.DecodeDraftMessageClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_clearDraftMessage, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogClearDraftMessage) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_clearDraftMessage, int(layer)); clazzId {
	case 0xfb70b29a:
		x.PutClazzID(0xfb70b29a)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_clearDraftMessage, layer)
	}
}

// Decode <--
func (m *TLDialogClearDraftMessage) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xfb70b29a:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDialogGetAllDrafts <--
type TLDialogGetAllDrafts struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLDialogGetAllDrafts) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_getAllDrafts, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetAllDrafts) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_getAllDrafts, int(layer)); clazzId {
	case 0xacde4fe6:
		x.PutClazzID(0xacde4fe6)

		x.PutInt64(m.UserId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getAllDrafts, layer)
	}
}

// Decode <--
func (m *TLDialogGetAllDrafts) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xacde4fe6:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDialogClearAllDrafts <--
type TLDialogClearAllDrafts struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLDialogClearAllDrafts) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_clearAllDrafts, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogClearAllDrafts) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_clearAllDrafts, int(layer)); clazzId {
	case 0x41b890fc:
		x.PutClazzID(0x41b890fc)

		x.PutInt64(m.UserId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_clearAllDrafts, layer)
	}
}

// Decode <--
func (m *TLDialogClearAllDrafts) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x41b890fc:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDialogMarkDialogUnread <--
type TLDialogMarkDialogUnread struct {
	ClazzID    uint32       `json:"_id"`
	UserId     int64        `json:"user_id"`
	PeerType   int32        `json:"peer_type"`
	PeerId     int64        `json:"peer_id"`
	UnreadMark tg.BoolClazz `json:"unread_mark"`
}

func (m *TLDialogMarkDialogUnread) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_markDialogUnread, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogMarkDialogUnread) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_markDialogUnread, int(layer)); clazzId {
	case 0x4532910e:
		x.PutClazzID(0x4532910e)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		_ = m.UnreadMark.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_markDialogUnread, layer)
	}
}

// Decode <--
func (m *TLDialogMarkDialogUnread) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x4532910e:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}

		m.UnreadMark, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDialogToggleDialogPin <--
type TLDialogToggleDialogPin struct {
	ClazzID  uint32       `json:"_id"`
	UserId   int64        `json:"user_id"`
	PeerType int32        `json:"peer_type"`
	PeerId   int64        `json:"peer_id"`
	Pinned   tg.BoolClazz `json:"pinned"`
}

func (m *TLDialogToggleDialogPin) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_toggleDialogPin, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogToggleDialogPin) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_toggleDialogPin, int(layer)); clazzId {
	case 0x867ee52f:
		x.PutClazzID(0x867ee52f)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		_ = m.Pinned.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_toggleDialogPin, layer)
	}
}

// Decode <--
func (m *TLDialogToggleDialogPin) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x867ee52f:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}

		m.Pinned, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDialogGetDialogUnreadMarkList <--
type TLDialogGetDialogUnreadMarkList struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLDialogGetDialogUnreadMarkList) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_getDialogUnreadMarkList, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogUnreadMarkList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogUnreadMarkList, int(layer)); clazzId {
	case 0xcabc38f4:
		x.PutClazzID(0xcabc38f4)

		x.PutInt64(m.UserId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogUnreadMarkList, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogUnreadMarkList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xcabc38f4:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDialogGetDialogsByOffsetDate <--
type TLDialogGetDialogsByOffsetDate struct {
	ClazzID       uint32       `json:"_id"`
	UserId        int64        `json:"user_id"`
	ExcludePinned tg.BoolClazz `json:"exclude_pinned"`
	OffsetDate    int32        `json:"offset_date"`
	Limit         int32        `json:"limit"`
}

func (m *TLDialogGetDialogsByOffsetDate) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_getDialogsByOffsetDate, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogsByOffsetDate) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogsByOffsetDate, int(layer)); clazzId {
	case 0x9d7e8604:
		x.PutClazzID(0x9d7e8604)

		x.PutInt64(m.UserId)
		_ = m.ExcludePinned.Encode(x, layer)
		x.PutInt32(m.OffsetDate)
		x.PutInt32(m.Limit)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogsByOffsetDate, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogsByOffsetDate) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x9d7e8604:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		m.ExcludePinned, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return err
		}

		m.OffsetDate, err = d.Int32()
		if err != nil {
			return err
		}
		m.Limit, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDialogGetDialogs <--
type TLDialogGetDialogs struct {
	ClazzID       uint32       `json:"_id"`
	UserId        int64        `json:"user_id"`
	ExcludePinned tg.BoolClazz `json:"exclude_pinned"`
	FolderId      int32        `json:"folder_id"`
}

func (m *TLDialogGetDialogs) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_getDialogs, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogs) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogs, int(layer)); clazzId {
	case 0x860b1e16:
		x.PutClazzID(0x860b1e16)

		x.PutInt64(m.UserId)
		_ = m.ExcludePinned.Encode(x, layer)
		x.PutInt32(m.FolderId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogs, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogs) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x860b1e16:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		m.ExcludePinned, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return err
		}

		m.FolderId, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_getDialogsByIdList, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogsByIdList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogsByIdList, int(layer)); clazzId {
	case 0xad258871:
		x.PutClazzID(0xad258871)

		x.PutInt64(m.UserId)

		iface.EncodeInt64List(x, m.IdList)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogsByIdList, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogsByIdList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xad258871:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		m.IdList, err = iface.DecodeInt64List(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDialogGetDialogsCount <--
type TLDialogGetDialogsCount struct {
	ClazzID       uint32       `json:"_id"`
	UserId        int64        `json:"user_id"`
	ExcludePinned tg.BoolClazz `json:"exclude_pinned"`
	FolderId      int32        `json:"folder_id"`
}

func (m *TLDialogGetDialogsCount) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_getDialogsCount, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogsCount) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogsCount, int(layer)); clazzId {
	case 0xe039b465:
		x.PutClazzID(0xe039b465)

		x.PutInt64(m.UserId)
		_ = m.ExcludePinned.Encode(x, layer)
		x.PutInt32(m.FolderId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogsCount, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogsCount) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xe039b465:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		m.ExcludePinned, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return err
		}

		m.FolderId, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_getPinnedDialogs, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetPinnedDialogs) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_getPinnedDialogs, int(layer)); clazzId {
	case 0xa8c21bb5:
		x.PutClazzID(0xa8c21bb5)

		x.PutInt64(m.UserId)
		x.PutInt32(m.FolderId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getPinnedDialogs, layer)
	}
}

// Decode <--
func (m *TLDialogGetPinnedDialogs) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xa8c21bb5:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.FolderId, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDialogReorderPinnedDialogs <--
type TLDialogReorderPinnedDialogs struct {
	ClazzID  uint32       `json:"_id"`
	UserId   int64        `json:"user_id"`
	Force    tg.BoolClazz `json:"force"`
	FolderId int32        `json:"folder_id"`
	IdList   []int64      `json:"id_list"`
}

func (m *TLDialogReorderPinnedDialogs) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_reorderPinnedDialogs, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogReorderPinnedDialogs) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_reorderPinnedDialogs, int(layer)); clazzId {
	case 0xfee33567:
		x.PutClazzID(0xfee33567)

		x.PutInt64(m.UserId)
		_ = m.Force.Encode(x, layer)
		x.PutInt32(m.FolderId)

		iface.EncodeInt64List(x, m.IdList)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_reorderPinnedDialogs, layer)
	}
}

// Decode <--
func (m *TLDialogReorderPinnedDialogs) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xfee33567:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		m.Force, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return err
		}

		m.FolderId, err = d.Int32()
		if err != nil {
			return err
		}

		m.IdList, err = iface.DecodeInt64List(d)
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_getDialogById, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogById) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogById, int(layer)); clazzId {
	case 0xa15f3bf5:
		x.PutClazzID(0xa15f3bf5)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogById, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogById) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xa15f3bf5:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_getTopMessage, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetTopMessage) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_getTopMessage, int(layer)); clazzId {
	case 0xfa7db272:
		x.PutClazzID(0xfa7db272)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getTopMessage, layer)
	}
}

// Decode <--
func (m *TLDialogGetTopMessage) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xfa7db272:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_insertOrUpdateDialog, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogInsertOrUpdateDialog) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_insertOrUpdateDialog, int(layer)); clazzId {
	case 0x5d2b8822:
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
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_insertOrUpdateDialog, layer)
	}
}

// Decode <--
func (m *TLDialogInsertOrUpdateDialog) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x5d2b8822:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}
		if (flags & (1 << 0)) != 0 {
			m.TopMessage = new(int32)
			*m.TopMessage, err = d.Int32()
			if err != nil {
				return err
			}
		}
		if (flags & (1 << 1)) != 0 {
			m.ReadOutboxMaxId = new(int32)
			*m.ReadOutboxMaxId, err = d.Int32()
			if err != nil {
				return err
			}
		}
		if (flags & (1 << 2)) != 0 {
			m.ReadInboxMaxId = new(int32)
			*m.ReadInboxMaxId, err = d.Int32()
			if err != nil {
				return err
			}
		}
		if (flags & (1 << 3)) != 0 {
			m.UnreadCount = new(int32)
			*m.UnreadCount, err = d.Int32()
			if err != nil {
				return err
			}
		}
		if (flags & (1 << 4)) != 0 {
			m.UnreadMark = true
		}
		if (flags & (1 << 5)) != 0 {
			m.Date2 = new(int64)
			*m.Date2, err = d.Int64()
			if err != nil {
				return err
			}
		}

		if (flags & (1 << 6)) != 0 {
			m.PinnedMsgId = new(int32)
			*m.PinnedMsgId, err = d.Int32()
			if err != nil {
				return err
			}
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_deleteDialog, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogDeleteDialog) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_deleteDialog, int(layer)); clazzId {
	case 0x1b31de3:
		x.PutClazzID(0x1b31de3)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_deleteDialog, layer)
	}
}

// Decode <--
func (m *TLDialogDeleteDialog) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x1b31de3:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_getUserPinnedMessage, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetUserPinnedMessage) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_getUserPinnedMessage, int(layer)); clazzId {
	case 0x8f9bc2b1:
		x.PutClazzID(0x8f9bc2b1)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getUserPinnedMessage, layer)
	}
}

// Decode <--
func (m *TLDialogGetUserPinnedMessage) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x8f9bc2b1:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_updateUserPinnedMessage, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogUpdateUserPinnedMessage) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_updateUserPinnedMessage, int(layer)); clazzId {
	case 0x1622f22a:
		x.PutClazzID(0x1622f22a)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt32(m.PinnedMsgId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_updateUserPinnedMessage, layer)
	}
}

// Decode <--
func (m *TLDialogUpdateUserPinnedMessage) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x1622f22a:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PinnedMsgId, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDialogInsertOrUpdateDialogFilter <--
type TLDialogInsertOrUpdateDialogFilter struct {
	ClazzID      uint32               `json:"_id"`
	UserId       int64                `json:"user_id"`
	Id           int32                `json:"id"`
	DialogFilter tg.DialogFilterClazz `json:"dialog_filter"`
}

func (m *TLDialogInsertOrUpdateDialogFilter) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_insertOrUpdateDialogFilter, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogInsertOrUpdateDialogFilter) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_insertOrUpdateDialogFilter, int(layer)); clazzId {
	case 0xaa8a384:
		x.PutClazzID(0xaa8a384)

		x.PutInt64(m.UserId)
		x.PutInt32(m.Id)
		_ = m.DialogFilter.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_insertOrUpdateDialogFilter, layer)
	}
}

// Decode <--
func (m *TLDialogInsertOrUpdateDialogFilter) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xaa8a384:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.Id, err = d.Int32()
		if err != nil {
			return err
		}

		m.DialogFilter, err = tg.DecodeDialogFilterClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_deleteDialogFilter, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogDeleteDialogFilter) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_deleteDialogFilter, int(layer)); clazzId {
	case 0x1dd3e97:
		x.PutClazzID(0x1dd3e97)

		x.PutInt64(m.UserId)
		x.PutInt32(m.Id)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_deleteDialogFilter, layer)
	}
}

// Decode <--
func (m *TLDialogDeleteDialogFilter) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x1dd3e97:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.Id, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_updateDialogFiltersOrder, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogUpdateDialogFiltersOrder) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_updateDialogFiltersOrder, int(layer)); clazzId {
	case 0xb13c0b3f:
		x.PutClazzID(0xb13c0b3f)

		x.PutInt64(m.UserId)

		iface.EncodeInt32List(x, m.Order)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_updateDialogFiltersOrder, layer)
	}
}

// Decode <--
func (m *TLDialogUpdateDialogFiltersOrder) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xb13c0b3f:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		m.Order, err = iface.DecodeInt32List(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDialogGetDialogFilters <--
type TLDialogGetDialogFilters struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLDialogGetDialogFilters) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_getDialogFilters, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogFilters) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogFilters, int(layer)); clazzId {
	case 0x6c676c3c:
		x.PutClazzID(0x6c676c3c)

		x.PutInt64(m.UserId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogFilters, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogFilters) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x6c676c3c:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_getDialogFolder, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogFolder) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogFolder, int(layer)); clazzId {
	case 0x411b8eb5:
		x.PutClazzID(0x411b8eb5)

		x.PutInt64(m.UserId)
		x.PutInt32(m.FolderId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogFolder, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogFolder) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x411b8eb5:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.FolderId, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_editPeerFolders, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogEditPeerFolders) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_editPeerFolders, int(layer)); clazzId {
	case 0x2446869a:
		x.PutClazzID(0x2446869a)

		x.PutInt64(m.UserId)

		iface.EncodeInt64List(x, m.PeerDialogList)

		x.PutInt32(m.FolderId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_editPeerFolders, layer)
	}
}

// Decode <--
func (m *TLDialogEditPeerFolders) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x2446869a:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		m.PeerDialogList, err = iface.DecodeInt64List(d)
		if err != nil {
			return err
		}

		m.FolderId, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_getChannelMessageReadParticipants, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetChannelMessageReadParticipants) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_getChannelMessageReadParticipants, int(layer)); clazzId {
	case 0x28bd4d3b:
		x.PutClazzID(0x28bd4d3b)

		x.PutInt64(m.UserId)
		x.PutInt64(m.ChannelId)
		x.PutInt32(m.MsgId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getChannelMessageReadParticipants, layer)
	}
}

// Decode <--
func (m *TLDialogGetChannelMessageReadParticipants) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x28bd4d3b:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.ChannelId, err = d.Int64()
		if err != nil {
			return err
		}
		m.MsgId, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_setChatTheme, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogSetChatTheme) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_setChatTheme, int(layer)); clazzId {
	case 0xe9aea22a:
		x.PutClazzID(0xe9aea22a)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutString(m.ThemeEmoticon)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_setChatTheme, layer)
	}
}

// Decode <--
func (m *TLDialogSetChatTheme) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xe9aea22a:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}
		m.ThemeEmoticon, err = d.String()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_setHistoryTTL, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogSetHistoryTTL) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_setHistoryTTL, int(layer)); clazzId {
	case 0x9d9b8ac:
		x.PutClazzID(0x9d9b8ac)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt32(m.TtlPeriod)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_setHistoryTTL, layer)
	}
}

// Decode <--
func (m *TLDialogSetHistoryTTL) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x9d9b8ac:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}
		m.TtlPeriod, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_getMyDialogsData, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetMyDialogsData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_getMyDialogsData, int(layer)); clazzId {
	case 0x7ee08f03:
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
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getMyDialogsData, layer)
	}
}

// Decode <--
func (m *TLDialogGetMyDialogsData) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x7ee08f03:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
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
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDialogGetSavedDialogs <--
type TLDialogGetSavedDialogs struct {
	ClazzID       uint32           `json:"_id"`
	UserId        int64            `json:"user_id"`
	ExcludePinned tg.BoolClazz     `json:"exclude_pinned"`
	OffsetDate    int32            `json:"offset_date"`
	OffsetId      int32            `json:"offset_id"`
	OffsetPeer    tg.PeerUtilClazz `json:"offset_peer"`
	Limit         int32            `json:"limit"`
}

func (m *TLDialogGetSavedDialogs) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_getSavedDialogs, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetSavedDialogs) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_getSavedDialogs, int(layer)); clazzId {
	case 0x38c1d668:
		x.PutClazzID(0x38c1d668)

		x.PutInt64(m.UserId)
		_ = m.ExcludePinned.Encode(x, layer)
		x.PutInt32(m.OffsetDate)
		x.PutInt32(m.OffsetId)
		_ = m.OffsetPeer.Encode(x, layer)
		x.PutInt32(m.Limit)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getSavedDialogs, layer)
	}
}

// Decode <--
func (m *TLDialogGetSavedDialogs) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x38c1d668:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		m.ExcludePinned, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return err
		}

		m.OffsetDate, err = d.Int32()
		if err != nil {
			return err
		}
		m.OffsetId, err = d.Int32()
		if err != nil {
			return err
		}

		m.OffsetPeer, err = tg.DecodePeerUtilClazz(d)
		if err != nil {
			return err
		}

		m.Limit, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDialogGetPinnedSavedDialogs <--
type TLDialogGetPinnedSavedDialogs struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLDialogGetPinnedSavedDialogs) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_getPinnedSavedDialogs, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetPinnedSavedDialogs) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_getPinnedSavedDialogs, int(layer)); clazzId {
	case 0x40a3b7e7:
		x.PutClazzID(0x40a3b7e7)

		x.PutInt64(m.UserId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getPinnedSavedDialogs, layer)
	}
}

// Decode <--
func (m *TLDialogGetPinnedSavedDialogs) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x40a3b7e7:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDialogToggleSavedDialogPin <--
type TLDialogToggleSavedDialogPin struct {
	ClazzID uint32           `json:"_id"`
	UserId  int64            `json:"user_id"`
	Peer    tg.PeerUtilClazz `json:"peer"`
	Pinned  tg.BoolClazz     `json:"pinned"`
}

func (m *TLDialogToggleSavedDialogPin) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_toggleSavedDialogPin, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogToggleSavedDialogPin) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_toggleSavedDialogPin, int(layer)); clazzId {
	case 0x44f317d9:
		x.PutClazzID(0x44f317d9)

		x.PutInt64(m.UserId)
		_ = m.Peer.Encode(x, layer)
		_ = m.Pinned.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_toggleSavedDialogPin, layer)
	}
}

// Decode <--
func (m *TLDialogToggleSavedDialogPin) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x44f317d9:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		m.Peer, err = tg.DecodePeerUtilClazz(d)
		if err != nil {
			return err
		}

		m.Pinned, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDialogReorderPinnedSavedDialogs <--
type TLDialogReorderPinnedSavedDialogs struct {
	ClazzID uint32             `json:"_id"`
	UserId  int64              `json:"user_id"`
	Force   tg.BoolClazz       `json:"force"`
	Order   []tg.PeerUtilClazz `json:"order"`
}

func (m *TLDialogReorderPinnedSavedDialogs) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_reorderPinnedSavedDialogs, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogReorderPinnedSavedDialogs) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_reorderPinnedSavedDialogs, int(layer)); clazzId {
	case 0xd85ccbd2:
		x.PutClazzID(0xd85ccbd2)

		x.PutInt64(m.UserId)
		_ = m.Force.Encode(x, layer)

		if err := iface.EncodeObjectList(x, m.Order, layer); err != nil {
			return err
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_reorderPinnedSavedDialogs, layer)
	}
}

// Decode <--
func (m *TLDialogReorderPinnedSavedDialogs) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xd85ccbd2:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		m.Force, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return err
		}

		c3, err2 := d.ClazzID()
		if err2 != nil {
			return err2
		}
		if c3 != iface.ClazzID_vector {
			return fmt.Errorf("invalid ClazzID_vector, c%d: %d", 3, c3)
		}
		l3, err3 := d.Int()
		if err3 != nil {
			return err3
		}
		v3 := make([]tg.PeerUtilClazz, l3)
		for i := 0; i < l3; i++ {
			v3[i], err3 = tg.DecodePeerUtilClazz(d)
			if err3 != nil {
				return err3
			}
		}
		m.Order = v3

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_getDialogFilter, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogFilter) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogFilter, int(layer)); clazzId {
	case 0xf388061c:
		x.PutClazzID(0xf388061c)

		x.PutInt64(m.UserId)
		x.PutInt32(m.Id)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogFilter, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogFilter) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xf388061c:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.Id, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_getDialogFilterBySlug, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogFilterBySlug) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogFilterBySlug, int(layer)); clazzId {
	case 0x4e457fef:
		x.PutClazzID(0x4e457fef)

		x.PutInt64(m.UserId)
		x.PutString(m.Slug)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogFilterBySlug, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogFilterBySlug) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x4e457fef:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.Slug, err = d.String()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDialogCreateDialogFilter <--
type TLDialogCreateDialogFilter struct {
	ClazzID      uint32               `json:"_id"`
	UserId       int64                `json:"user_id"`
	DialogFilter DialogFilterExtClazz `json:"dialog_filter"`
}

func (m *TLDialogCreateDialogFilter) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_createDialogFilter, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogCreateDialogFilter) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_createDialogFilter, int(layer)); clazzId {
	case 0xc6cb636f:
		x.PutClazzID(0xc6cb636f)

		x.PutInt64(m.UserId)
		_ = m.DialogFilter.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_createDialogFilter, layer)
	}
}

// Decode <--
func (m *TLDialogCreateDialogFilter) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xc6cb636f:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		m.DialogFilter, err = DecodeDialogFilterExtClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_updateUnreadCount, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogUpdateUnreadCount) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_updateUnreadCount, int(layer)); clazzId {
	case 0x2bac334d:
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
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_updateUnreadCount, layer)
	}
}

// Decode <--
func (m *TLDialogUpdateUnreadCount) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x2bac334d:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}
		if (flags & (1 << 0)) != 0 {
			m.UnreadCount = new(int32)
			*m.UnreadCount, err = d.Int32()
			if err != nil {
				return err
			}
		}
		if (flags & (1 << 1)) != 0 {
			m.UnreadMentionsCount = new(int32)
			*m.UnreadMentionsCount, err = d.Int32()
			if err != nil {
				return err
			}
		}
		if (flags & (1 << 2)) != 0 {
			m.UnreadReactionsCount = new(int32)
			*m.UnreadReactionsCount, err = d.Int32()
			if err != nil {
				return err
			}
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDialogToggleDialogFilterTags <--
type TLDialogToggleDialogFilterTags struct {
	ClazzID uint32       `json:"_id"`
	UserId  int64        `json:"user_id"`
	Enabled tg.BoolClazz `json:"enabled"`
}

func (m *TLDialogToggleDialogFilterTags) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_toggleDialogFilterTags, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogToggleDialogFilterTags) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_toggleDialogFilterTags, int(layer)); clazzId {
	case 0xa0cd6d89:
		x.PutClazzID(0xa0cd6d89)

		x.PutInt64(m.UserId)
		_ = m.Enabled.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_toggleDialogFilterTags, layer)
	}
}

// Decode <--
func (m *TLDialogToggleDialogFilterTags) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xa0cd6d89:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		m.Enabled, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDialogGetDialogFilterTags <--
type TLDialogGetDialogFilterTags struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLDialogGetDialogFilterTags) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_getDialogFilterTags, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogGetDialogFilterTags) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_getDialogFilterTags, int(layer)); clazzId {
	case 0xfaf0fa97:
		x.PutClazzID(0xfaf0fa97)

		x.PutInt64(m.UserId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_getDialogFilterTags, layer)
	}
}

// Decode <--
func (m *TLDialogGetDialogFilterTags) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xfaf0fa97:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dialog_setChatWallpaper, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDialogSetChatWallpaper) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dialog_setChatWallpaper, int(layer)); clazzId {
	case 0xb551db12:
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
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialog_setChatWallpaper, layer)
	}
}

// Decode <--
func (m *TLDialogSetChatWallpaper) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xb551db12:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}
		m.WallpaperId, err = d.Int64()
		if err != nil {
			return err
		}
		if (flags & (1 << 0)) != 0 {
			m.WallpaperOverridden = true
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// Vector api result type
// ----------------------------------------------------------------------------
// VectorResList <--

// VectorPeerWithDraftMessage <--
type VectorPeerWithDraftMessage struct {
	Datas []PeerWithDraftMessageClazz `json:"_datas"`
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
	m.Datas, err = iface.DecodeObjectList[PeerWithDraftMessageClazz](d)

	return err
}

// VectorDialogPeer <--
type VectorDialogPeer struct {
	Datas []tg.DialogPeerClazz `json:"_datas"`
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
	m.Datas, err = iface.DecodeObjectList[tg.DialogPeerClazz](d)

	return err
}

// VectorDialogExt <--
type VectorDialogExt struct {
	Datas []DialogExtClazz `json:"_datas"`
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
	m.Datas, err = iface.DecodeObjectList[DialogExtClazz](d)

	return err
}

// VectorDialogFilterExt <--
type VectorDialogFilterExt struct {
	Datas []DialogFilterExtClazz `json:"_datas"`
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
	m.Datas, err = iface.DecodeObjectList[DialogFilterExtClazz](d)

	return err
}

// VectorDialogPinnedExt <--
type VectorDialogPinnedExt struct {
	Datas []DialogPinnedExtClazz `json:"_datas"`
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
	m.Datas, err = iface.DecodeObjectList[DialogPinnedExtClazz](d)

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
