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

// TLChatGetMutableChat <--
type TLChatGetMutableChat struct {
	ClazzID uint32 `json:"_id"`
	ChatId  int64  `json:"chat_id"`
}

func (m *TLChatGetMutableChat) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatGetMutableChat) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x2c2c25d2: func() error {
			x.PutClazzID(0x2c2c25d2)

			x.PutInt64(m.ChatId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_getMutableChat, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_getMutableChat, layer)
	}
}

// Decode <--
func (m *TLChatGetMutableChat) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x2c2c25d2: func() (err error) {
			m.ChatId, err = d.Int64()

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

// TLChatGetChatListByIdList <--
type TLChatGetChatListByIdList struct {
	ClazzID uint32  `json:"_id"`
	SelfId  int64   `json:"self_id"`
	IdList  []int64 `json:"id_list"`
}

func (m *TLChatGetChatListByIdList) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatGetChatListByIdList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xe740f539: func() error {
			x.PutClazzID(0xe740f539)

			x.PutInt64(m.SelfId)

			iface.EncodeInt64List(x, m.IdList)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_getChatListByIdList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_getChatListByIdList, layer)
	}
}

// Decode <--
func (m *TLChatGetChatListByIdList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xe740f539: func() (err error) {
			m.SelfId, err = d.Int64()

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

// TLChatGetChatBySelfId <--
type TLChatGetChatBySelfId struct {
	ClazzID uint32 `json:"_id"`
	SelfId  int64  `json:"self_id"`
	ChatId  int64  `json:"chat_id"`
}

func (m *TLChatGetChatBySelfId) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatGetChatBySelfId) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x49b71a48: func() error {
			x.PutClazzID(0x49b71a48)

			x.PutInt64(m.SelfId)
			x.PutInt64(m.ChatId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_getChatBySelfId, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_getChatBySelfId, layer)
	}
}

// Decode <--
func (m *TLChatGetChatBySelfId) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x49b71a48: func() (err error) {
			m.SelfId, err = d.Int64()
			m.ChatId, err = d.Int64()

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

// TLChatCreateChat2 <--
type TLChatCreateChat2 struct {
	ClazzID    uint32  `json:"_id"`
	CreatorId  int64   `json:"creator_id"`
	UserIdList []int64 `json:"user_id_list"`
	Title      string  `json:"title"`
	Bots       []int64 `json:"bots"`
}

func (m *TLChatCreateChat2) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatCreateChat2) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xf77448d2: func() error {
			x.PutClazzID(0xf77448d2)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.Bots != nil {
					flags |= 1 << 0
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.CreatorId)

			iface.EncodeInt64List(x, m.UserIdList)

			x.PutString(m.Title)
			if m.Bots != nil {
				iface.EncodeInt64List(x, m.Bots)
			}

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_createChat2, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_createChat2, layer)
	}
}

// Decode <--
func (m *TLChatCreateChat2) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xf77448d2: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.CreatorId, err = d.Int64()

			m.UserIdList, err = iface.DecodeInt64List(d)

			m.Title, err = d.String()
			if (flags & (1 << 0)) != 0 {
				m.Bots, err = iface.DecodeInt64List(d)
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

// TLChatDeleteChat <--
type TLChatDeleteChat struct {
	ClazzID    uint32 `json:"_id"`
	ChatId     int64  `json:"chat_id"`
	OperatorId int64  `json:"operator_id"`
}

func (m *TLChatDeleteChat) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatDeleteChat) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x6d11ec1e: func() error {
			x.PutClazzID(0x6d11ec1e)

			x.PutInt64(m.ChatId)
			x.PutInt64(m.OperatorId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_deleteChat, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_deleteChat, layer)
	}
}

// Decode <--
func (m *TLChatDeleteChat) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x6d11ec1e: func() (err error) {
			m.ChatId, err = d.Int64()
			m.OperatorId, err = d.Int64()

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

// TLChatDeleteChatUser <--
type TLChatDeleteChatUser struct {
	ClazzID      uint32 `json:"_id"`
	ChatId       int64  `json:"chat_id"`
	OperatorId   int64  `json:"operator_id"`
	DeleteUserId int64  `json:"delete_user_id"`
}

func (m *TLChatDeleteChatUser) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatDeleteChatUser) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xb270fd5: func() error {
			x.PutClazzID(0xb270fd5)

			x.PutInt64(m.ChatId)
			x.PutInt64(m.OperatorId)
			x.PutInt64(m.DeleteUserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_deleteChatUser, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_deleteChatUser, layer)
	}
}

// Decode <--
func (m *TLChatDeleteChatUser) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xb270fd5: func() (err error) {
			m.ChatId, err = d.Int64()
			m.OperatorId, err = d.Int64()
			m.DeleteUserId, err = d.Int64()

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

// TLChatEditChatTitle <--
type TLChatEditChatTitle struct {
	ClazzID    uint32 `json:"_id"`
	ChatId     int64  `json:"chat_id"`
	EditUserId int64  `json:"edit_user_id"`
	Title      string `json:"title"`
}

func (m *TLChatEditChatTitle) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatEditChatTitle) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x95c59ea7: func() error {
			x.PutClazzID(0x95c59ea7)

			x.PutInt64(m.ChatId)
			x.PutInt64(m.EditUserId)
			x.PutString(m.Title)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_editChatTitle, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_editChatTitle, layer)
	}
}

// Decode <--
func (m *TLChatEditChatTitle) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x95c59ea7: func() (err error) {
			m.ChatId, err = d.Int64()
			m.EditUserId, err = d.Int64()
			m.Title, err = d.String()

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

// TLChatEditChatAbout <--
type TLChatEditChatAbout struct {
	ClazzID    uint32 `json:"_id"`
	ChatId     int64  `json:"chat_id"`
	EditUserId int64  `json:"edit_user_id"`
	About      string `json:"about"`
}

func (m *TLChatEditChatAbout) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatEditChatAbout) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x5c737c78: func() error {
			x.PutClazzID(0x5c737c78)

			x.PutInt64(m.ChatId)
			x.PutInt64(m.EditUserId)
			x.PutString(m.About)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_editChatAbout, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_editChatAbout, layer)
	}
}

// Decode <--
func (m *TLChatEditChatAbout) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x5c737c78: func() (err error) {
			m.ChatId, err = d.Int64()
			m.EditUserId, err = d.Int64()
			m.About, err = d.String()

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

// TLChatEditChatPhoto <--
type TLChatEditChatPhoto struct {
	ClazzID    uint32    `json:"_id"`
	ChatId     int64     `json:"chat_id"`
	EditUserId int64     `json:"edit_user_id"`
	ChatPhoto  *tg.Photo `json:"chat_photo"`
}

func (m *TLChatEditChatPhoto) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatEditChatPhoto) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x45c2a668: func() error {
			x.PutClazzID(0x45c2a668)

			x.PutInt64(m.ChatId)
			x.PutInt64(m.EditUserId)
			_ = m.ChatPhoto.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_editChatPhoto, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_editChatPhoto, layer)
	}
}

// Decode <--
func (m *TLChatEditChatPhoto) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x45c2a668: func() (err error) {
			m.ChatId, err = d.Int64()
			m.EditUserId, err = d.Int64()

			m3 := &tg.Photo{}
			_ = m3.Decode(d)
			m.ChatPhoto = m3

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

// TLChatEditChatAdmin <--
type TLChatEditChatAdmin struct {
	ClazzID         uint32   `json:"_id"`
	ChatId          int64    `json:"chat_id"`
	OperatorId      int64    `json:"operator_id"`
	EditChatAdminId int64    `json:"edit_chat_admin_id"`
	IsAdmin         *tg.Bool `json:"is_admin"`
}

func (m *TLChatEditChatAdmin) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatEditChatAdmin) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x1905e5ec: func() error {
			x.PutClazzID(0x1905e5ec)

			x.PutInt64(m.ChatId)
			x.PutInt64(m.OperatorId)
			x.PutInt64(m.EditChatAdminId)
			_ = m.IsAdmin.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_editChatAdmin, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_editChatAdmin, layer)
	}
}

// Decode <--
func (m *TLChatEditChatAdmin) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x1905e5ec: func() (err error) {
			m.ChatId, err = d.Int64()
			m.OperatorId, err = d.Int64()
			m.EditChatAdminId, err = d.Int64()

			m4 := &tg.Bool{}
			_ = m4.Decode(d)
			m.IsAdmin = m4

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

// TLChatEditChatDefaultBannedRights <--
type TLChatEditChatDefaultBannedRights struct {
	ClazzID      uint32               `json:"_id"`
	ChatId       int64                `json:"chat_id"`
	OperatorId   int64                `json:"operator_id"`
	BannedRights *tg.ChatBannedRights `json:"banned_rights"`
}

func (m *TLChatEditChatDefaultBannedRights) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatEditChatDefaultBannedRights) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x5a34a687: func() error {
			x.PutClazzID(0x5a34a687)

			x.PutInt64(m.ChatId)
			x.PutInt64(m.OperatorId)
			_ = m.BannedRights.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_editChatDefaultBannedRights, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_editChatDefaultBannedRights, layer)
	}
}

// Decode <--
func (m *TLChatEditChatDefaultBannedRights) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x5a34a687: func() (err error) {
			m.ChatId, err = d.Int64()
			m.OperatorId, err = d.Int64()

			m3 := &tg.ChatBannedRights{}
			_ = m3.Decode(d)
			m.BannedRights = m3

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

// TLChatAddChatUser <--
type TLChatAddChatUser struct {
	ClazzID   uint32 `json:"_id"`
	ChatId    int64  `json:"chat_id"`
	InviterId int64  `json:"inviter_id"`
	UserId    int64  `json:"user_id"`
	IsBot     bool   `json:"is_bot"`
}

func (m *TLChatAddChatUser) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatAddChatUser) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xe5554168: func() error {
			x.PutClazzID(0xe5554168)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.IsBot == true {
					flags |= 1 << 0
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.ChatId)
			x.PutInt64(m.InviterId)
			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_addChatUser, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_addChatUser, layer)
	}
}

// Decode <--
func (m *TLChatAddChatUser) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xe5554168: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.ChatId, err = d.Int64()
			m.InviterId, err = d.Int64()
			m.UserId, err = d.Int64()
			if (flags & (1 << 0)) != 0 {
				m.IsBot = true
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

// TLChatGetMutableChatByLink <--
type TLChatGetMutableChatByLink struct {
	ClazzID uint32 `json:"_id"`
	Link    string `json:"link"`
}

func (m *TLChatGetMutableChatByLink) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatGetMutableChatByLink) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa266278b: func() error {
			x.PutClazzID(0xa266278b)

			x.PutString(m.Link)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_getMutableChatByLink, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_getMutableChatByLink, layer)
	}
}

// Decode <--
func (m *TLChatGetMutableChatByLink) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa266278b: func() (err error) {
			m.Link, err = d.String()

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

// TLChatToggleNoForwards <--
type TLChatToggleNoForwards struct {
	ClazzID    uint32   `json:"_id"`
	ChatId     int64    `json:"chat_id"`
	OperatorId int64    `json:"operator_id"`
	Enabled    *tg.Bool `json:"enabled"`
}

func (m *TLChatToggleNoForwards) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatToggleNoForwards) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xd5952af9: func() error {
			x.PutClazzID(0xd5952af9)

			x.PutInt64(m.ChatId)
			x.PutInt64(m.OperatorId)
			_ = m.Enabled.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_toggleNoForwards, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_toggleNoForwards, layer)
	}
}

// Decode <--
func (m *TLChatToggleNoForwards) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xd5952af9: func() (err error) {
			m.ChatId, err = d.Int64()
			m.OperatorId, err = d.Int64()

			m3 := &tg.Bool{}
			_ = m3.Decode(d)
			m.Enabled = m3

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

// TLChatMigratedToChannel <--
type TLChatMigratedToChannel struct {
	ClazzID    uint32          `json:"_id"`
	Chat       *tg.MutableChat `json:"chat"`
	Id         int64           `json:"id"`
	AccessHash int64           `json:"access_hash"`
}

func (m *TLChatMigratedToChannel) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatMigratedToChannel) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x83faadf: func() error {
			x.PutClazzID(0x83faadf)

			_ = m.Chat.Encode(x, layer)
			x.PutInt64(m.Id)
			x.PutInt64(m.AccessHash)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_migratedToChannel, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_migratedToChannel, layer)
	}
}

// Decode <--
func (m *TLChatMigratedToChannel) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x83faadf: func() (err error) {

			m1 := &tg.MutableChat{}
			_ = m1.Decode(d)
			m.Chat = m1

			m.Id, err = d.Int64()
			m.AccessHash, err = d.Int64()

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

// TLChatGetChatParticipantIdList <--
type TLChatGetChatParticipantIdList struct {
	ClazzID uint32 `json:"_id"`
	ChatId  int64  `json:"chat_id"`
}

func (m *TLChatGetChatParticipantIdList) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatGetChatParticipantIdList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x329622a9: func() error {
			x.PutClazzID(0x329622a9)

			x.PutInt64(m.ChatId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_getChatParticipantIdList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_getChatParticipantIdList, layer)
	}
}

// Decode <--
func (m *TLChatGetChatParticipantIdList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x329622a9: func() (err error) {
			m.ChatId, err = d.Int64()

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

// TLChatGetUsersChatIdList <--
type TLChatGetUsersChatIdList struct {
	ClazzID uint32  `json:"_id"`
	Id      []int64 `json:"id"`
}

func (m *TLChatGetUsersChatIdList) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatGetUsersChatIdList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x2f36ab4c: func() error {
			x.PutClazzID(0x2f36ab4c)

			iface.EncodeInt64List(x, m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_getUsersChatIdList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_getUsersChatIdList, layer)
	}
}

// Decode <--
func (m *TLChatGetUsersChatIdList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x2f36ab4c: func() (err error) {

			m.Id, err = iface.DecodeInt64List(d)

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

// TLChatGetMyChatList <--
type TLChatGetMyChatList struct {
	ClazzID   uint32   `json:"_id"`
	UserId    int64    `json:"user_id"`
	IsCreator *tg.Bool `json:"is_creator"`
}

func (m *TLChatGetMyChatList) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatGetMyChatList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xf3756c88: func() error {
			x.PutClazzID(0xf3756c88)

			x.PutInt64(m.UserId)
			_ = m.IsCreator.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_getMyChatList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_getMyChatList, layer)
	}
}

// Decode <--
func (m *TLChatGetMyChatList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xf3756c88: func() (err error) {
			m.UserId, err = d.Int64()

			m2 := &tg.Bool{}
			_ = m2.Decode(d)
			m.IsCreator = m2

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

// TLChatExportChatInvite <--
type TLChatExportChatInvite struct {
	ClazzID               uint32  `json:"_id"`
	ChatId                int64   `json:"chat_id"`
	AdminId               int64   `json:"admin_id"`
	LegacyRevokePermanent bool    `json:"legacy_revoke_permanent"`
	RequestNeeded         bool    `json:"request_needed"`
	ExpireDate            *int32  `json:"expire_date"`
	UsageLimit            *int32  `json:"usage_limit"`
	Title                 *string `json:"title"`
}

func (m *TLChatExportChatInvite) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatExportChatInvite) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc5cf804b: func() error {
			x.PutClazzID(0xc5cf804b)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.LegacyRevokePermanent == true {
					flags |= 1 << 2
				}
				if m.RequestNeeded == true {
					flags |= 1 << 3
				}
				if m.ExpireDate != nil {
					flags |= 1 << 0
				}
				if m.UsageLimit != nil {
					flags |= 1 << 1
				}
				if m.Title != nil {
					flags |= 1 << 4
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.ChatId)
			x.PutInt64(m.AdminId)
			if m.ExpireDate != nil {
				x.PutInt32(*m.ExpireDate)
			}

			if m.UsageLimit != nil {
				x.PutInt32(*m.UsageLimit)
			}

			if m.Title != nil {
				x.PutString(*m.Title)
			}

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_exportChatInvite, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_exportChatInvite, layer)
	}
}

// Decode <--
func (m *TLChatExportChatInvite) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xc5cf804b: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.ChatId, err = d.Int64()
			m.AdminId, err = d.Int64()
			if (flags & (1 << 2)) != 0 {
				m.LegacyRevokePermanent = true
			}
			if (flags & (1 << 3)) != 0 {
				m.RequestNeeded = true
			}
			if (flags & (1 << 0)) != 0 {
				m.ExpireDate = new(int32)
				*m.ExpireDate, err = d.Int32()
			}
			if (flags & (1 << 1)) != 0 {
				m.UsageLimit = new(int32)
				*m.UsageLimit, err = d.Int32()
			}
			if (flags & (1 << 4)) != 0 {
				m.Title = new(string)
				*m.Title, err = d.String()
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

// TLChatGetAdminsWithInvites <--
type TLChatGetAdminsWithInvites struct {
	ClazzID uint32 `json:"_id"`
	SelfId  int64  `json:"self_id"`
	ChatId  int64  `json:"chat_id"`
}

func (m *TLChatGetAdminsWithInvites) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatGetAdminsWithInvites) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xd2ea41d2: func() error {
			x.PutClazzID(0xd2ea41d2)

			x.PutInt64(m.SelfId)
			x.PutInt64(m.ChatId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_getAdminsWithInvites, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_getAdminsWithInvites, layer)
	}
}

// Decode <--
func (m *TLChatGetAdminsWithInvites) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xd2ea41d2: func() (err error) {
			m.SelfId, err = d.Int64()
			m.ChatId, err = d.Int64()

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

// TLChatGetExportedChatInvite <--
type TLChatGetExportedChatInvite struct {
	ClazzID uint32 `json:"_id"`
	ChatId  int64  `json:"chat_id"`
	Link    string `json:"link"`
}

func (m *TLChatGetExportedChatInvite) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatGetExportedChatInvite) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xddea3250: func() error {
			x.PutClazzID(0xddea3250)

			x.PutInt64(m.ChatId)
			x.PutString(m.Link)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_getExportedChatInvite, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_getExportedChatInvite, layer)
	}
}

// Decode <--
func (m *TLChatGetExportedChatInvite) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xddea3250: func() (err error) {
			m.ChatId, err = d.Int64()
			m.Link, err = d.String()

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

// TLChatGetExportedChatInvites <--
type TLChatGetExportedChatInvites struct {
	ClazzID    uint32  `json:"_id"`
	ChatId     int64   `json:"chat_id"`
	AdminId    int64   `json:"admin_id"`
	Revoked    bool    `json:"revoked"`
	OffsetDate *int32  `json:"offset_date"`
	OffsetLink *string `json:"offset_link"`
	Limit      int32   `json:"limit"`
}

func (m *TLChatGetExportedChatInvites) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatGetExportedChatInvites) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xb48f18f6: func() error {
			x.PutClazzID(0xb48f18f6)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.Revoked == true {
					flags |= 1 << 3
				}
				if m.OffsetDate != nil {
					flags |= 1 << 2
				}
				if m.OffsetLink != nil {
					flags |= 1 << 2
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.ChatId)
			x.PutInt64(m.AdminId)
			if m.OffsetDate != nil {
				x.PutInt32(*m.OffsetDate)
			}

			if m.OffsetLink != nil {
				x.PutString(*m.OffsetLink)
			}

			x.PutInt32(m.Limit)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_getExportedChatInvites, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_getExportedChatInvites, layer)
	}
}

// Decode <--
func (m *TLChatGetExportedChatInvites) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xb48f18f6: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.ChatId, err = d.Int64()
			m.AdminId, err = d.Int64()
			if (flags & (1 << 3)) != 0 {
				m.Revoked = true
			}
			if (flags & (1 << 2)) != 0 {
				m.OffsetDate = new(int32)
				*m.OffsetDate, err = d.Int32()
			}
			if (flags & (1 << 2)) != 0 {
				m.OffsetLink = new(string)
				*m.OffsetLink, err = d.String()
			}

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

// TLChatCheckChatInvite <--
type TLChatCheckChatInvite struct {
	ClazzID uint32 `json:"_id"`
	SelfId  int64  `json:"self_id"`
	Hash    string `json:"hash"`
}

func (m *TLChatCheckChatInvite) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatCheckChatInvite) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x7387f28c: func() error {
			x.PutClazzID(0x7387f28c)

			x.PutInt64(m.SelfId)
			x.PutString(m.Hash)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_checkChatInvite, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_checkChatInvite, layer)
	}
}

// Decode <--
func (m *TLChatCheckChatInvite) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x7387f28c: func() (err error) {
			m.SelfId, err = d.Int64()
			m.Hash, err = d.String()

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

// TLChatImportChatInvite <--
type TLChatImportChatInvite struct {
	ClazzID uint32 `json:"_id"`
	SelfId  int64  `json:"self_id"`
	Hash    string `json:"hash"`
}

func (m *TLChatImportChatInvite) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatImportChatInvite) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x58e660d4: func() error {
			x.PutClazzID(0x58e660d4)

			x.PutInt64(m.SelfId)
			x.PutString(m.Hash)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_importChatInvite, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_importChatInvite, layer)
	}
}

// Decode <--
func (m *TLChatImportChatInvite) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x58e660d4: func() (err error) {
			m.SelfId, err = d.Int64()
			m.Hash, err = d.String()

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

// TLChatGetChatInviteImporters <--
type TLChatGetChatInviteImporters struct {
	ClazzID    uint32  `json:"_id"`
	SelfId     int64   `json:"self_id"`
	ChatId     int64   `json:"chat_id"`
	Requested  bool    `json:"requested"`
	Link       *string `json:"link"`
	Q          *string `json:"q"`
	OffsetDate int32   `json:"offset_date"`
	OffsetUser int64   `json:"offset_user"`
	Limit      int32   `json:"limit"`
}

func (m *TLChatGetChatInviteImporters) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatGetChatInviteImporters) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x9846557f: func() error {
			x.PutClazzID(0x9846557f)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.Requested == true {
					flags |= 1 << 0
				}
				if m.Link != nil {
					flags |= 1 << 1
				}
				if m.Q != nil {
					flags |= 1 << 2
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.SelfId)
			x.PutInt64(m.ChatId)
			if m.Link != nil {
				x.PutString(*m.Link)
			}

			if m.Q != nil {
				x.PutString(*m.Q)
			}

			x.PutInt32(m.OffsetDate)
			x.PutInt64(m.OffsetUser)
			x.PutInt32(m.Limit)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_getChatInviteImporters, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_getChatInviteImporters, layer)
	}
}

// Decode <--
func (m *TLChatGetChatInviteImporters) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x9846557f: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.SelfId, err = d.Int64()
			m.ChatId, err = d.Int64()
			if (flags & (1 << 0)) != 0 {
				m.Requested = true
			}
			if (flags & (1 << 1)) != 0 {
				m.Link = new(string)
				*m.Link, err = d.String()
			}

			if (flags & (1 << 2)) != 0 {
				m.Q = new(string)
				*m.Q, err = d.String()
			}

			m.OffsetDate, err = d.Int32()
			m.OffsetUser, err = d.Int64()
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

// TLChatDeleteExportedChatInvite <--
type TLChatDeleteExportedChatInvite struct {
	ClazzID uint32 `json:"_id"`
	SelfId  int64  `json:"self_id"`
	ChatId  int64  `json:"chat_id"`
	Link    string `json:"link"`
}

func (m *TLChatDeleteExportedChatInvite) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatDeleteExportedChatInvite) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x562288b8: func() error {
			x.PutClazzID(0x562288b8)

			x.PutInt64(m.SelfId)
			x.PutInt64(m.ChatId)
			x.PutString(m.Link)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_deleteExportedChatInvite, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_deleteExportedChatInvite, layer)
	}
}

// Decode <--
func (m *TLChatDeleteExportedChatInvite) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x562288b8: func() (err error) {
			m.SelfId, err = d.Int64()
			m.ChatId, err = d.Int64()
			m.Link, err = d.String()

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

// TLChatDeleteRevokedExportedChatInvites <--
type TLChatDeleteRevokedExportedChatInvites struct {
	ClazzID uint32 `json:"_id"`
	SelfId  int64  `json:"self_id"`
	ChatId  int64  `json:"chat_id"`
	AdminId int64  `json:"admin_id"`
}

func (m *TLChatDeleteRevokedExportedChatInvites) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatDeleteRevokedExportedChatInvites) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xd0126269: func() error {
			x.PutClazzID(0xd0126269)

			x.PutInt64(m.SelfId)
			x.PutInt64(m.ChatId)
			x.PutInt64(m.AdminId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_deleteRevokedExportedChatInvites, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_deleteRevokedExportedChatInvites, layer)
	}
}

// Decode <--
func (m *TLChatDeleteRevokedExportedChatInvites) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xd0126269: func() (err error) {
			m.SelfId, err = d.Int64()
			m.ChatId, err = d.Int64()
			m.AdminId, err = d.Int64()

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

// TLChatEditExportedChatInvite <--
type TLChatEditExportedChatInvite struct {
	ClazzID       uint32   `json:"_id"`
	SelfId        int64    `json:"self_id"`
	ChatId        int64    `json:"chat_id"`
	Revoked       bool     `json:"revoked"`
	Link          string   `json:"link"`
	ExpireDate    *int32   `json:"expire_date"`
	UsageLimit    *int32   `json:"usage_limit"`
	RequestNeeded *tg.Bool `json:"request_needed"`
	Title         *string  `json:"title"`
}

func (m *TLChatEditExportedChatInvite) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatEditExportedChatInvite) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xaf994c76: func() error {
			x.PutClazzID(0xaf994c76)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.Revoked == true {
					flags |= 1 << 2
				}

				if m.ExpireDate != nil {
					flags |= 1 << 0
				}
				if m.UsageLimit != nil {
					flags |= 1 << 1
				}
				if m.RequestNeeded != nil {
					flags |= 1 << 3
				}
				if m.Title != nil {
					flags |= 1 << 4
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.SelfId)
			x.PutInt64(m.ChatId)
			x.PutString(m.Link)
			if m.ExpireDate != nil {
				x.PutInt32(*m.ExpireDate)
			}

			if m.UsageLimit != nil {
				x.PutInt32(*m.UsageLimit)
			}

			if m.RequestNeeded != nil {
				_ = m.RequestNeeded.Encode(x, layer)
			}

			if m.Title != nil {
				x.PutString(*m.Title)
			}

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_editExportedChatInvite, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_editExportedChatInvite, layer)
	}
}

// Decode <--
func (m *TLChatEditExportedChatInvite) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xaf994c76: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.SelfId, err = d.Int64()
			m.ChatId, err = d.Int64()
			if (flags & (1 << 2)) != 0 {
				m.Revoked = true
			}
			m.Link, err = d.String()
			if (flags & (1 << 0)) != 0 {
				m.ExpireDate = new(int32)
				*m.ExpireDate, err = d.Int32()
			}
			if (flags & (1 << 1)) != 0 {
				m.UsageLimit = new(int32)
				*m.UsageLimit, err = d.Int32()
			}
			if (flags & (1 << 3)) != 0 {
				m8 := &tg.Bool{}
				_ = m8.Decode(d)
				m.RequestNeeded = m8
			}
			if (flags & (1 << 4)) != 0 {
				m.Title = new(string)
				*m.Title, err = d.String()
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

// TLChatSetChatAvailableReactions <--
type TLChatSetChatAvailableReactions struct {
	ClazzID                uint32   `json:"_id"`
	SelfId                 int64    `json:"self_id"`
	ChatId                 int64    `json:"chat_id"`
	AvailableReactionsType int32    `json:"available_reactions_type"`
	AvailableReactions     []string `json:"available_reactions"`
}

func (m *TLChatSetChatAvailableReactions) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatSetChatAvailableReactions) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc4d08972: func() error {
			x.PutClazzID(0xc4d08972)

			x.PutInt64(m.SelfId)
			x.PutInt64(m.ChatId)
			x.PutInt32(m.AvailableReactionsType)

			iface.EncodeStringList(x, m.AvailableReactions)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_setChatAvailableReactions, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_setChatAvailableReactions, layer)
	}
}

// Decode <--
func (m *TLChatSetChatAvailableReactions) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xc4d08972: func() (err error) {
			m.SelfId, err = d.Int64()
			m.ChatId, err = d.Int64()
			m.AvailableReactionsType, err = d.Int32()

			m.AvailableReactions, err = iface.DecodeStringList(d)

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

// TLChatSetHistoryTTL <--
type TLChatSetHistoryTTL struct {
	ClazzID   uint32 `json:"_id"`
	SelfId    int64  `json:"self_id"`
	ChatId    int64  `json:"chat_id"`
	TtlPeriod int32  `json:"ttl_period"`
}

func (m *TLChatSetHistoryTTL) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatSetHistoryTTL) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x3cfb6384: func() error {
			x.PutClazzID(0x3cfb6384)

			x.PutInt64(m.SelfId)
			x.PutInt64(m.ChatId)
			x.PutInt32(m.TtlPeriod)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_setHistoryTTL, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_setHistoryTTL, layer)
	}
}

// Decode <--
func (m *TLChatSetHistoryTTL) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x3cfb6384: func() (err error) {
			m.SelfId, err = d.Int64()
			m.ChatId, err = d.Int64()
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

// TLChatSearch <--
type TLChatSearch struct {
	ClazzID uint32 `json:"_id"`
	SelfId  int64  `json:"self_id"`
	Q       string `json:"q"`
	Offset  int64  `json:"offset"`
	Limit   int32  `json:"limit"`
}

func (m *TLChatSearch) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatSearch) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x21e014fb: func() error {
			x.PutClazzID(0x21e014fb)

			x.PutInt64(m.SelfId)
			x.PutString(m.Q)
			x.PutInt64(m.Offset)
			x.PutInt32(m.Limit)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_search, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_search, layer)
	}
}

// Decode <--
func (m *TLChatSearch) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x21e014fb: func() (err error) {
			m.SelfId, err = d.Int64()
			m.Q, err = d.String()
			m.Offset, err = d.Int64()
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

// TLChatGetRecentChatInviteRequesters <--
type TLChatGetRecentChatInviteRequesters struct {
	ClazzID uint32 `json:"_id"`
	SelfId  int64  `json:"self_id"`
	ChatId  int64  `json:"chat_id"`
}

func (m *TLChatGetRecentChatInviteRequesters) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatGetRecentChatInviteRequesters) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xfedc1098: func() error {
			x.PutClazzID(0xfedc1098)

			x.PutInt64(m.SelfId)
			x.PutInt64(m.ChatId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_getRecentChatInviteRequesters, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_getRecentChatInviteRequesters, layer)
	}
}

// Decode <--
func (m *TLChatGetRecentChatInviteRequesters) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xfedc1098: func() (err error) {
			m.SelfId, err = d.Int64()
			m.ChatId, err = d.Int64()

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

// TLChatHideChatJoinRequests <--
type TLChatHideChatJoinRequests struct {
	ClazzID  uint32  `json:"_id"`
	SelfId   int64   `json:"self_id"`
	ChatId   int64   `json:"chat_id"`
	Approved bool    `json:"approved"`
	Link     *string `json:"link"`
	UserId   *int64  `json:"user_id"`
}

func (m *TLChatHideChatJoinRequests) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatHideChatJoinRequests) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x3ea52cd1: func() error {
			x.PutClazzID(0x3ea52cd1)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.Approved == true {
					flags |= 1 << 0
				}
				if m.Link != nil {
					flags |= 1 << 1
				}
				if m.UserId != nil {
					flags |= 1 << 2
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.SelfId)
			x.PutInt64(m.ChatId)
			if m.Link != nil {
				x.PutString(*m.Link)
			}

			if m.UserId != nil {
				x.PutInt64(*m.UserId)
			}

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_hideChatJoinRequests, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_hideChatJoinRequests, layer)
	}
}

// Decode <--
func (m *TLChatHideChatJoinRequests) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x3ea52cd1: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.SelfId, err = d.Int64()
			m.ChatId, err = d.Int64()
			if (flags & (1 << 0)) != 0 {
				m.Approved = true
			}
			if (flags & (1 << 1)) != 0 {
				m.Link = new(string)
				*m.Link, err = d.String()
			}

			if (flags & (1 << 2)) != 0 {
				m.UserId = new(int64)
				*m.UserId, err = d.Int64()
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

// TLChatImportChatInvite2 <--
type TLChatImportChatInvite2 struct {
	ClazzID uint32 `json:"_id"`
	SelfId  int64  `json:"self_id"`
	Hash    string `json:"hash"`
}

func (m *TLChatImportChatInvite2) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLChatImportChatInvite2) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xdcd93dbf: func() error {
			x.PutClazzID(0xdcd93dbf)

			x.PutInt64(m.SelfId)
			x.PutString(m.Hash)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_chat_importChatInvite2, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_chat_importChatInvite2, layer)
	}
}

// Decode <--
func (m *TLChatImportChatInvite2) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xdcd93dbf: func() (err error) {
			m.SelfId, err = d.Int64()
			m.Hash, err = d.String()

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

// VectorMutableChat <--
type VectorMutableChat struct {
	Datas []*tg.MutableChat `json:"_datas"`
}

func (m *VectorMutableChat) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorMutableChat) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorMutableChat) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*tg.MutableChat](d)

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

// VectorUserChatIdList <--
type VectorUserChatIdList struct {
	Datas []*UserChatIdList `json:"_datas"`
}

func (m *VectorUserChatIdList) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorUserChatIdList) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorUserChatIdList) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*UserChatIdList](d)

	return err
}

// VectorChatAdminWithInvites <--
type VectorChatAdminWithInvites struct {
	Datas []*tg.ChatAdminWithInvites `json:"_datas"`
}

func (m *VectorChatAdminWithInvites) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorChatAdminWithInvites) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorChatAdminWithInvites) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*tg.ChatAdminWithInvites](d)

	return err
}

// VectorExportedChatInvite <--
type VectorExportedChatInvite struct {
	Datas []*tg.ExportedChatInvite `json:"_datas"`
}

func (m *VectorExportedChatInvite) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorExportedChatInvite) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorExportedChatInvite) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*tg.ExportedChatInvite](d)

	return err
}

// VectorChatInviteImporter <--
type VectorChatInviteImporter struct {
	Datas []*tg.ChatInviteImporter `json:"_datas"`
}

func (m *VectorChatInviteImporter) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorChatInviteImporter) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorChatInviteImporter) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*tg.ChatInviteImporter](d)

	return err
}

// ----------------------------------------------------------------------------
// rpc

type RPCChat interface {
	ChatGetMutableChat(ctx context.Context, in *TLChatGetMutableChat) (*tg.MutableChat, error)
	ChatGetChatListByIdList(ctx context.Context, in *TLChatGetChatListByIdList) (*VectorMutableChat, error)
	ChatGetChatBySelfId(ctx context.Context, in *TLChatGetChatBySelfId) (*tg.MutableChat, error)
	ChatCreateChat2(ctx context.Context, in *TLChatCreateChat2) (*tg.MutableChat, error)
	ChatDeleteChat(ctx context.Context, in *TLChatDeleteChat) (*tg.MutableChat, error)
	ChatDeleteChatUser(ctx context.Context, in *TLChatDeleteChatUser) (*tg.MutableChat, error)
	ChatEditChatTitle(ctx context.Context, in *TLChatEditChatTitle) (*tg.MutableChat, error)
	ChatEditChatAbout(ctx context.Context, in *TLChatEditChatAbout) (*tg.MutableChat, error)
	ChatEditChatPhoto(ctx context.Context, in *TLChatEditChatPhoto) (*tg.MutableChat, error)
	ChatEditChatAdmin(ctx context.Context, in *TLChatEditChatAdmin) (*tg.MutableChat, error)
	ChatEditChatDefaultBannedRights(ctx context.Context, in *TLChatEditChatDefaultBannedRights) (*tg.MutableChat, error)
	ChatAddChatUser(ctx context.Context, in *TLChatAddChatUser) (*tg.MutableChat, error)
	ChatGetMutableChatByLink(ctx context.Context, in *TLChatGetMutableChatByLink) (*tg.MutableChat, error)
	ChatToggleNoForwards(ctx context.Context, in *TLChatToggleNoForwards) (*tg.MutableChat, error)
	ChatMigratedToChannel(ctx context.Context, in *TLChatMigratedToChannel) (*tg.Bool, error)
	ChatGetChatParticipantIdList(ctx context.Context, in *TLChatGetChatParticipantIdList) (*VectorLong, error)
	ChatGetUsersChatIdList(ctx context.Context, in *TLChatGetUsersChatIdList) (*VectorUserChatIdList, error)
	ChatGetMyChatList(ctx context.Context, in *TLChatGetMyChatList) (*VectorMutableChat, error)
	ChatExportChatInvite(ctx context.Context, in *TLChatExportChatInvite) (*tg.ExportedChatInvite, error)
	ChatGetAdminsWithInvites(ctx context.Context, in *TLChatGetAdminsWithInvites) (*VectorChatAdminWithInvites, error)
	ChatGetExportedChatInvite(ctx context.Context, in *TLChatGetExportedChatInvite) (*tg.ExportedChatInvite, error)
	ChatGetExportedChatInvites(ctx context.Context, in *TLChatGetExportedChatInvites) (*VectorExportedChatInvite, error)
	ChatCheckChatInvite(ctx context.Context, in *TLChatCheckChatInvite) (*ChatInviteExt, error)
	ChatImportChatInvite(ctx context.Context, in *TLChatImportChatInvite) (*tg.MutableChat, error)
	ChatGetChatInviteImporters(ctx context.Context, in *TLChatGetChatInviteImporters) (*VectorChatInviteImporter, error)
	ChatDeleteExportedChatInvite(ctx context.Context, in *TLChatDeleteExportedChatInvite) (*tg.Bool, error)
	ChatDeleteRevokedExportedChatInvites(ctx context.Context, in *TLChatDeleteRevokedExportedChatInvites) (*tg.Bool, error)
	ChatEditExportedChatInvite(ctx context.Context, in *TLChatEditExportedChatInvite) (*VectorExportedChatInvite, error)
	ChatSetChatAvailableReactions(ctx context.Context, in *TLChatSetChatAvailableReactions) (*tg.MutableChat, error)
	ChatSetHistoryTTL(ctx context.Context, in *TLChatSetHistoryTTL) (*tg.MutableChat, error)
	ChatSearch(ctx context.Context, in *TLChatSearch) (*VectorMutableChat, error)
	ChatGetRecentChatInviteRequesters(ctx context.Context, in *TLChatGetRecentChatInviteRequesters) (*RecentChatInviteRequesters, error)
	ChatHideChatJoinRequests(ctx context.Context, in *TLChatHideChatJoinRequests) (*RecentChatInviteRequesters, error)
	ChatImportChatInvite2(ctx context.Context, in *TLChatImportChatInvite2) (*ChatInviteImported, error)
}
