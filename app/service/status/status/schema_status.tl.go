/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package status

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

// SessionEntryClazz <--
//   - TL_SessionEntry
type SessionEntryClazz interface {
	iface.TLObject
	SessionEntryClazzName() string
}

func DecodeSessionEntryClazz(d *bin.Decoder) (SessionEntryClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_sessionEntry:
		x := &TLSessionEntry{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeSessionEntry - unexpected clazzId: %d", id)
	}
}

// TLSessionEntry <--
type TLSessionEntry struct {
	ClazzID       uint32 `json:"_id"`
	UserId        int64  `json:"user_id"`
	AuthKeyId     int64  `json:"auth_key_id"`
	Gateway       string `json:"gateway"`
	Expired       int64  `json:"expired"`
	Layer         int32  `json:"layer"`
	PermAuthKeyId int64  `json:"perm_auth_key_id"`
	Client        string `json:"client"`
}

// SessionEntryClazzName <--
func (m *TLSessionEntry) SessionEntryClazzName() string {
	return ClazzName_sessionEntry
}

// ClazzName <--
func (m *TLSessionEntry) ClazzName() string {
	return ClazzName_sessionEntry
}

// ToSessionEntry <--
func (m *TLSessionEntry) ToSessionEntry() *SessionEntry {
	return MakeSessionEntry(m)
}

// Encode <--
func (m *TLSessionEntry) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x1764ac31: func() error {
			x.PutClazzID(0x1764ac31)

			x.PutInt64(m.UserId)
			x.PutInt64(m.AuthKeyId)
			x.PutString(m.Gateway)
			x.PutInt64(m.Expired)
			x.PutInt32(m.Layer)
			x.PutInt64(m.PermAuthKeyId)
			x.PutString(m.Client)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_sessionEntry, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_sessionEntry, layer)
	}
}

// Decode <--
func (m *TLSessionEntry) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x1764ac31: func() (err error) {
			m.UserId, err = d.Int64()
			m.AuthKeyId, err = d.Int64()
			m.Gateway, err = d.String()
			m.Expired, err = d.Int64()
			m.Layer, err = d.Int32()
			m.PermAuthKeyId, err = d.Int64()
			m.Client, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// SessionEntry <--
type SessionEntry struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	SessionEntryClazz
}

// MakeSessionEntry <--
func MakeSessionEntry(c SessionEntryClazz) *SessionEntry {
	return &SessionEntry{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		SessionEntryClazz: c,
	}
}

// Encode <--
func (m *SessionEntry) Encode(x *bin.Encoder, layer int32) error {
	if m.SessionEntryClazz != nil {
		return m.SessionEntryClazz.Encode(x, layer)
	}

	return fmt.Errorf("SessionEntry - invalid Clazz")
}

// Decode <--
func (m *SessionEntry) Decode(d *bin.Decoder) (err error) {
	m.SessionEntryClazz, err = DecodeSessionEntryClazz(d)
	return
}

// Match <--
func (m *SessionEntry) Match(f ...interface{}) {
	switch c := m.SessionEntryClazz.(type) {
	case *TLSessionEntry:
		for _, v := range f {
			if f1, ok := v.(func(c *TLSessionEntry) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToSessionEntry <--
func (m *SessionEntry) ToSessionEntry() (*TLSessionEntry, bool) {
	if m.SessionEntryClazz == nil {
		return nil, false
	}

	if x, ok := m.SessionEntryClazz.(*TLSessionEntry); ok {
		return x, true
	}

	return nil, false
}

// UserSessionEntryListClazz <--
//   - TL_UserSessionEntryList
type UserSessionEntryListClazz interface {
	iface.TLObject
	UserSessionEntryListClazzName() string
}

func DecodeUserSessionEntryListClazz(d *bin.Decoder) (UserSessionEntryListClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_userSessionEntryList:
		x := &TLUserSessionEntryList{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeUserSessionEntryList - unexpected clazzId: %d", id)
	}
}

// TLUserSessionEntryList <--
type TLUserSessionEntryList struct {
	ClazzID      uint32          `json:"_id"`
	UserId       int64           `json:"user_id"`
	UserSessions []*SessionEntry `json:"user_sessions"`
}

// UserSessionEntryListClazzName <--
func (m *TLUserSessionEntryList) UserSessionEntryListClazzName() string {
	return ClazzName_userSessionEntryList
}

// ClazzName <--
func (m *TLUserSessionEntryList) ClazzName() string {
	return ClazzName_userSessionEntryList
}

// ToUserSessionEntryList <--
func (m *TLUserSessionEntryList) ToUserSessionEntryList() *UserSessionEntryList {
	return MakeUserSessionEntryList(m)
}

// Encode <--
func (m *TLUserSessionEntryList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xefecb398: func() error {
			x.PutClazzID(0xefecb398)

			x.PutInt64(m.UserId)

			_ = iface.EncodeObjectList(x, m.UserSessions, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_userSessionEntryList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_userSessionEntryList, layer)
	}
}

// Decode <--
func (m *TLUserSessionEntryList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xefecb398: func() (err error) {
			m.UserId, err = d.Int64()
			c1, err2 := d.ClazzID()
			if c1 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 1, c1)
				return err2
			}
			l1, err3 := d.Int()
			v1 := make([]*SessionEntry, l1)
			for i := 0; i < l1; i++ {
				vv := new(SessionEntry)
				err3 = vv.Decode(d)
				_ = err3
				v1[i] = vv
			}
			m.UserSessions = v1

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// UserSessionEntryList <--
type UserSessionEntryList struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	UserSessionEntryListClazz
}

// MakeUserSessionEntryList <--
func MakeUserSessionEntryList(c UserSessionEntryListClazz) *UserSessionEntryList {
	return &UserSessionEntryList{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		UserSessionEntryListClazz: c,
	}
}

// Encode <--
func (m *UserSessionEntryList) Encode(x *bin.Encoder, layer int32) error {
	if m.UserSessionEntryListClazz != nil {
		return m.UserSessionEntryListClazz.Encode(x, layer)
	}

	return fmt.Errorf("UserSessionEntryList - invalid Clazz")
}

// Decode <--
func (m *UserSessionEntryList) Decode(d *bin.Decoder) (err error) {
	m.UserSessionEntryListClazz, err = DecodeUserSessionEntryListClazz(d)
	return
}

// Match <--
func (m *UserSessionEntryList) Match(f ...interface{}) {
	switch c := m.UserSessionEntryListClazz.(type) {
	case *TLUserSessionEntryList:
		for _, v := range f {
			if f1, ok := v.(func(c *TLUserSessionEntryList) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToUserSessionEntryList <--
func (m *UserSessionEntryList) ToUserSessionEntryList() (*TLUserSessionEntryList, bool) {
	if m.UserSessionEntryListClazz == nil {
		return nil, false
	}

	if x, ok := m.UserSessionEntryListClazz.(*TLUserSessionEntryList); ok {
		return x, true
	}

	return nil, false
}
