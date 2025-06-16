/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package dialog

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

// DialogExtClazz <--
//   - TL_DialogExt
type DialogExtClazz interface {
	iface.TLObject
	DialogExtClazzName() string
}

func DecodeDialogExtClazz(d *bin.Decoder) (DialogExtClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_dialogExt:
		x := &TLDialogExt{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeDialogExt - unexpected clazzId: %d", id)
	}
}

// TLDialogExt <--
type TLDialogExt struct {
	ClazzID        uint32     `json:"_id"`
	Order          int64      `json:"order"`
	Dialog         *tg.Dialog `json:"dialog"`
	AvailableMinId int32      `json:"available_min_id"`
	Date           int64      `json:"date"`
	ThemeEmoticon  string     `json:"theme_emoticon"`
	TtlPeriod      int32      `json:"ttl_period"`
}

// DialogExtClazzName <--
func (m *TLDialogExt) DialogExtClazzName() string {
	return ClazzName_dialogExt
}

// ClazzName <--
func (m *TLDialogExt) ClazzName() string {
	return ClazzName_dialogExt
}

// ToDialogExt <--
func (m *TLDialogExt) ToDialogExt() *DialogExt {
	return MakeDialogExt(m)
}

// Encode <--
func (m *TLDialogExt) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xbdd9a860: func() error {
			x.PutClazzID(0xbdd9a860)

			x.PutInt64(m.Order)
			_ = m.Dialog.Encode(x, layer)
			x.PutInt32(m.AvailableMinId)
			x.PutInt64(m.Date)
			x.PutString(m.ThemeEmoticon)
			x.PutInt32(m.TtlPeriod)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialogExt, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialogExt, layer)
	}
}

// Decode <--
func (m *TLDialogExt) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xbdd9a860: func() (err error) {
			m.Order, err = d.Int64()

			m1 := &tg.Dialog{}
			_ = m1.Decode(d)
			m.Dialog = m1

			m.AvailableMinId, err = d.Int32()
			m.Date, err = d.Int64()
			m.ThemeEmoticon, err = d.String()
			m.TtlPeriod, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// DialogExt <--
type DialogExt struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	DialogExtClazz
}

// MakeDialogExt <--
func MakeDialogExt(c DialogExtClazz) *DialogExt {
	return &DialogExt{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		DialogExtClazz: c,
	}
}

// Encode <--
func (m *DialogExt) Encode(x *bin.Encoder, layer int32) error {
	if m.DialogExtClazz != nil {
		return m.DialogExtClazz.Encode(x, layer)
	}

	return fmt.Errorf("DialogExt - invalid Clazz")
}

// Decode <--
func (m *DialogExt) Decode(d *bin.Decoder) (err error) {
	m.DialogExtClazz, err = DecodeDialogExtClazz(d)
	return
}

// Match <--
func (m *DialogExt) Match(f ...interface{}) {
	switch c := m.DialogExtClazz.(type) {
	case *TLDialogExt:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDialogExt) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToDialogExt <--
func (m *DialogExt) ToDialogExt() (*TLDialogExt, bool) {
	if m.DialogExtClazz == nil {
		return nil, false
	}

	if x, ok := m.DialogExtClazz.(*TLDialogExt); ok {
		return x, true
	}

	return nil, false
}

// DialogFilterExtClazz <--
//   - TL_DialogFilterExt
type DialogFilterExtClazz interface {
	iface.TLObject
	DialogFilterExtClazzName() string
}

func DecodeDialogFilterExtClazz(d *bin.Decoder) (DialogFilterExtClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_dialogFilterExt:
		x := &TLDialogFilterExt{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeDialogFilterExt - unexpected clazzId: %d", id)
	}
}

// TLDialogFilterExt <--
type TLDialogFilterExt struct {
	ClazzID      uint32           `json:"_id"`
	Id           int32            `json:"id"`
	JoinedBySlug bool             `json:"joined_by_slug"`
	Slug         string           `json:"slug"`
	DialogFilter *tg.DialogFilter `json:"dialog_filter"`
	Order        int64            `json:"order"`
}

// DialogFilterExtClazzName <--
func (m *TLDialogFilterExt) DialogFilterExtClazzName() string {
	return ClazzName_dialogFilterExt
}

// ClazzName <--
func (m *TLDialogFilterExt) ClazzName() string {
	return ClazzName_dialogFilterExt
}

// ToDialogFilterExt <--
func (m *TLDialogFilterExt) ToDialogFilterExt() *DialogFilterExt {
	return MakeDialogFilterExt(m)
}

// Encode <--
func (m *TLDialogFilterExt) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa6d498fe: func() error {
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
			_ = m.DialogFilter.Encode(x, layer)
			x.PutInt64(m.Order)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialogFilterExt, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialogFilterExt, layer)
	}
}

// Decode <--
func (m *TLDialogFilterExt) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa6d498fe: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.Id, err = d.Int32()
			if (flags & (1 << 0)) != 0 {
				m.JoinedBySlug = true
			}
			m.Slug, err = d.String()

			m4 := &tg.DialogFilter{}
			_ = m4.Decode(d)
			m.DialogFilter = m4

			m.Order, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// DialogFilterExt <--
type DialogFilterExt struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	DialogFilterExtClazz
}

// MakeDialogFilterExt <--
func MakeDialogFilterExt(c DialogFilterExtClazz) *DialogFilterExt {
	return &DialogFilterExt{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		DialogFilterExtClazz: c,
	}
}

// Encode <--
func (m *DialogFilterExt) Encode(x *bin.Encoder, layer int32) error {
	if m.DialogFilterExtClazz != nil {
		return m.DialogFilterExtClazz.Encode(x, layer)
	}

	return fmt.Errorf("DialogFilterExt - invalid Clazz")
}

// Decode <--
func (m *DialogFilterExt) Decode(d *bin.Decoder) (err error) {
	m.DialogFilterExtClazz, err = DecodeDialogFilterExtClazz(d)
	return
}

// Match <--
func (m *DialogFilterExt) Match(f ...interface{}) {
	switch c := m.DialogFilterExtClazz.(type) {
	case *TLDialogFilterExt:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDialogFilterExt) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToDialogFilterExt <--
func (m *DialogFilterExt) ToDialogFilterExt() (*TLDialogFilterExt, bool) {
	if m.DialogFilterExtClazz == nil {
		return nil, false
	}

	if x, ok := m.DialogFilterExtClazz.(*TLDialogFilterExt); ok {
		return x, true
	}

	return nil, false
}

// DialogPinnedExtClazz <--
//   - TL_DialogPinnedExt
type DialogPinnedExtClazz interface {
	iface.TLObject
	DialogPinnedExtClazzName() string
}

func DecodeDialogPinnedExtClazz(d *bin.Decoder) (DialogPinnedExtClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_dialogPinnedExt:
		x := &TLDialogPinnedExt{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeDialogPinnedExt - unexpected clazzId: %d", id)
	}
}

// TLDialogPinnedExt <--
type TLDialogPinnedExt struct {
	ClazzID  uint32 `json:"_id"`
	Order    int64  `json:"order"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

// DialogPinnedExtClazzName <--
func (m *TLDialogPinnedExt) DialogPinnedExtClazzName() string {
	return ClazzName_dialogPinnedExt
}

// ClazzName <--
func (m *TLDialogPinnedExt) ClazzName() string {
	return ClazzName_dialogPinnedExt
}

// ToDialogPinnedExt <--
func (m *TLDialogPinnedExt) ToDialogPinnedExt() *DialogPinnedExt {
	return MakeDialogPinnedExt(m)
}

// Encode <--
func (m *TLDialogPinnedExt) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xea7222c: func() error {
			x.PutClazzID(0xea7222c)

			x.PutInt64(m.Order)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dialogPinnedExt, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dialogPinnedExt, layer)
	}
}

// Decode <--
func (m *TLDialogPinnedExt) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xea7222c: func() (err error) {
			m.Order, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// DialogPinnedExt <--
type DialogPinnedExt struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	DialogPinnedExtClazz
}

// MakeDialogPinnedExt <--
func MakeDialogPinnedExt(c DialogPinnedExtClazz) *DialogPinnedExt {
	return &DialogPinnedExt{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		DialogPinnedExtClazz: c,
	}
}

// Encode <--
func (m *DialogPinnedExt) Encode(x *bin.Encoder, layer int32) error {
	if m.DialogPinnedExtClazz != nil {
		return m.DialogPinnedExtClazz.Encode(x, layer)
	}

	return fmt.Errorf("DialogPinnedExt - invalid Clazz")
}

// Decode <--
func (m *DialogPinnedExt) Decode(d *bin.Decoder) (err error) {
	m.DialogPinnedExtClazz, err = DecodeDialogPinnedExtClazz(d)
	return
}

// Match <--
func (m *DialogPinnedExt) Match(f ...interface{}) {
	switch c := m.DialogPinnedExtClazz.(type) {
	case *TLDialogPinnedExt:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDialogPinnedExt) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToDialogPinnedExt <--
func (m *DialogPinnedExt) ToDialogPinnedExt() (*TLDialogPinnedExt, bool) {
	if m.DialogPinnedExtClazz == nil {
		return nil, false
	}

	if x, ok := m.DialogPinnedExtClazz.(*TLDialogPinnedExt); ok {
		return x, true
	}

	return nil, false
}

// DialogsDataClazz <--
//   - TL_SimpleDialogsData
type DialogsDataClazz interface {
	iface.TLObject
	DialogsDataClazzName() string
}

func DecodeDialogsDataClazz(d *bin.Decoder) (DialogsDataClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_simpleDialogsData:
		x := &TLSimpleDialogsData{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeDialogsData - unexpected clazzId: %d", id)
	}
}

// TLSimpleDialogsData <--
type TLSimpleDialogsData struct {
	ClazzID  uint32  `json:"_id"`
	Users    []int64 `json:"users"`
	Chats    []int64 `json:"chats"`
	Channels []int64 `json:"channels"`
}

// DialogsDataClazzName <--
func (m *TLSimpleDialogsData) DialogsDataClazzName() string {
	return ClazzName_simpleDialogsData
}

// ClazzName <--
func (m *TLSimpleDialogsData) ClazzName() string {
	return ClazzName_simpleDialogsData
}

// ToDialogsData <--
func (m *TLSimpleDialogsData) ToDialogsData() *DialogsData {
	return MakeDialogsData(m)
}

// Encode <--
func (m *TLSimpleDialogsData) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x1d59b45d: func() error {
			x.PutClazzID(0x1d59b45d)

			iface.EncodeInt64List(x, m.Users)

			iface.EncodeInt64List(x, m.Chats)

			iface.EncodeInt64List(x, m.Channels)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_simpleDialogsData, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_simpleDialogsData, layer)
	}
}

// Decode <--
func (m *TLSimpleDialogsData) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x1d59b45d: func() (err error) {

			m.Users, err = iface.DecodeInt64List(d)

			m.Chats, err = iface.DecodeInt64List(d)

			m.Channels, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// DialogsData <--
type DialogsData struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	DialogsDataClazz
}

// MakeDialogsData <--
func MakeDialogsData(c DialogsDataClazz) *DialogsData {
	return &DialogsData{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		DialogsDataClazz: c,
	}
}

// Encode <--
func (m *DialogsData) Encode(x *bin.Encoder, layer int32) error {
	if m.DialogsDataClazz != nil {
		return m.DialogsDataClazz.Encode(x, layer)
	}

	return fmt.Errorf("DialogsData - invalid Clazz")
}

// Decode <--
func (m *DialogsData) Decode(d *bin.Decoder) (err error) {
	m.DialogsDataClazz, err = DecodeDialogsDataClazz(d)
	return
}

// Match <--
func (m *DialogsData) Match(f ...interface{}) {
	switch c := m.DialogsDataClazz.(type) {
	case *TLSimpleDialogsData:
		for _, v := range f {
			if f1, ok := v.(func(c *TLSimpleDialogsData) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToSimpleDialogsData <--
func (m *DialogsData) ToSimpleDialogsData() (*TLSimpleDialogsData, bool) {
	if m.DialogsDataClazz == nil {
		return nil, false
	}

	if x, ok := m.DialogsDataClazz.(*TLSimpleDialogsData); ok {
		return x, true
	}

	return nil, false
}

// PeerWithDraftMessageClazz <--
//   - TL_UpdateDraftMessage
type PeerWithDraftMessageClazz interface {
	iface.TLObject
	PeerWithDraftMessageClazzName() string
}

func DecodePeerWithDraftMessageClazz(d *bin.Decoder) (PeerWithDraftMessageClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_updateDraftMessage:
		x := &TLUpdateDraftMessage{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodePeerWithDraftMessage - unexpected clazzId: %d", id)
	}
}

// TLUpdateDraftMessage <--
type TLUpdateDraftMessage struct {
	ClazzID uint32           `json:"_id"`
	Peer    *tg.Peer         `json:"peer"`
	Draft   *tg.DraftMessage `json:"draft"`
}

// PeerWithDraftMessageClazzName <--
func (m *TLUpdateDraftMessage) PeerWithDraftMessageClazzName() string {
	return ClazzName_updateDraftMessage
}

// ClazzName <--
func (m *TLUpdateDraftMessage) ClazzName() string {
	return ClazzName_updateDraftMessage
}

// ToPeerWithDraftMessage <--
func (m *TLUpdateDraftMessage) ToPeerWithDraftMessage() *PeerWithDraftMessage {
	return MakePeerWithDraftMessage(m)
}

// Encode <--
func (m *TLUpdateDraftMessage) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xf6bdc4b2: func() error {
			x.PutClazzID(0xf6bdc4b2)

			_ = m.Peer.Encode(x, layer)
			_ = m.Draft.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_updateDraftMessage, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_updateDraftMessage, layer)
	}
}

// Decode <--
func (m *TLUpdateDraftMessage) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xf6bdc4b2: func() (err error) {

			m0 := &tg.Peer{}
			_ = m0.Decode(d)
			m.Peer = m0

			m1 := &tg.DraftMessage{}
			_ = m1.Decode(d)
			m.Draft = m1

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// PeerWithDraftMessage <--
type PeerWithDraftMessage struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	PeerWithDraftMessageClazz
}

// MakePeerWithDraftMessage <--
func MakePeerWithDraftMessage(c PeerWithDraftMessageClazz) *PeerWithDraftMessage {
	return &PeerWithDraftMessage{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		PeerWithDraftMessageClazz: c,
	}
}

// Encode <--
func (m *PeerWithDraftMessage) Encode(x *bin.Encoder, layer int32) error {
	if m.PeerWithDraftMessageClazz != nil {
		return m.PeerWithDraftMessageClazz.Encode(x, layer)
	}

	return fmt.Errorf("PeerWithDraftMessage - invalid Clazz")
}

// Decode <--
func (m *PeerWithDraftMessage) Decode(d *bin.Decoder) (err error) {
	m.PeerWithDraftMessageClazz, err = DecodePeerWithDraftMessageClazz(d)
	return
}

// Match <--
func (m *PeerWithDraftMessage) Match(f ...interface{}) {
	switch c := m.PeerWithDraftMessageClazz.(type) {
	case *TLUpdateDraftMessage:
		for _, v := range f {
			if f1, ok := v.(func(c *TLUpdateDraftMessage) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToUpdateDraftMessage <--
func (m *PeerWithDraftMessage) ToUpdateDraftMessage() (*TLUpdateDraftMessage, bool) {
	if m.PeerWithDraftMessageClazz == nil {
		return nil, false
	}

	if x, ok := m.PeerWithDraftMessageClazz.(*TLUpdateDraftMessage); ok {
		return x, true
	}

	return nil, false
}

// SavedDialogListClazz <--
//   - TL_SavedDialogList
type SavedDialogListClazz interface {
	iface.TLObject
	SavedDialogListClazzName() string
}

func DecodeSavedDialogListClazz(d *bin.Decoder) (SavedDialogListClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_savedDialogList:
		x := &TLSavedDialogList{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeSavedDialogList - unexpected clazzId: %d", id)
	}
}

// TLSavedDialogList <--
type TLSavedDialogList struct {
	ClazzID uint32            `json:"_id"`
	Count   int32             `json:"count"`
	Dialogs []*tg.SavedDialog `json:"dialogs"`
}

// SavedDialogListClazzName <--
func (m *TLSavedDialogList) SavedDialogListClazzName() string {
	return ClazzName_savedDialogList
}

// ClazzName <--
func (m *TLSavedDialogList) ClazzName() string {
	return ClazzName_savedDialogList
}

// ToSavedDialogList <--
func (m *TLSavedDialogList) ToSavedDialogList() *SavedDialogList {
	return MakeSavedDialogList(m)
}

// Encode <--
func (m *TLSavedDialogList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x778fe85a: func() error {
			x.PutClazzID(0x778fe85a)

			x.PutInt32(m.Count)

			_ = iface.EncodeObjectList(x, m.Dialogs, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_savedDialogList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_savedDialogList, layer)
	}
}

// Decode <--
func (m *TLSavedDialogList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x778fe85a: func() (err error) {
			m.Count, err = d.Int32()
			c1, err2 := d.ClazzID()
			if c1 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 1, c1)
				return err2
			}
			l1, err3 := d.Int()
			v1 := make([]*tg.SavedDialog, l1)
			for i := 0; i < l1; i++ {
				vv := new(tg.SavedDialog)
				err3 = vv.Decode(d)
				_ = err3
				v1[i] = vv
			}
			m.Dialogs = v1

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// SavedDialogList <--
type SavedDialogList struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	SavedDialogListClazz
}

// MakeSavedDialogList <--
func MakeSavedDialogList(c SavedDialogListClazz) *SavedDialogList {
	return &SavedDialogList{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		SavedDialogListClazz: c,
	}
}

// Encode <--
func (m *SavedDialogList) Encode(x *bin.Encoder, layer int32) error {
	if m.SavedDialogListClazz != nil {
		return m.SavedDialogListClazz.Encode(x, layer)
	}

	return fmt.Errorf("SavedDialogList - invalid Clazz")
}

// Decode <--
func (m *SavedDialogList) Decode(d *bin.Decoder) (err error) {
	m.SavedDialogListClazz, err = DecodeSavedDialogListClazz(d)
	return
}

// Match <--
func (m *SavedDialogList) Match(f ...interface{}) {
	switch c := m.SavedDialogListClazz.(type) {
	case *TLSavedDialogList:
		for _, v := range f {
			if f1, ok := v.(func(c *TLSavedDialogList) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToSavedDialogList <--
func (m *SavedDialogList) ToSavedDialogList() (*TLSavedDialogList, bool) {
	if m.SavedDialogListClazz == nil {
		return nil, false
	}

	if x, ok := m.SavedDialogListClazz.(*TLSavedDialogList); ok {
		return x, true
	}

	return nil, false
}
