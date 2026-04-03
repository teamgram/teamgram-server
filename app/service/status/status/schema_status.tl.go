/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package status

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

// SessionEntryClazz <--
//   - TL_SessionEntry
type SessionEntryClazz = *TLSessionEntry

func DecodeSessionEntryClazz(d *bin.Decoder) (SessionEntryClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x1764ac31:
		x := &TLSessionEntry{ClazzID: id, ClazzName2: ClazzName_sessionEntry}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeSessionEntry - unexpected clazzId: %d", id)
	}

}

// TLSessionEntry <--
type TLSessionEntry struct {
	ClazzID       uint32 `json:"_id"`
	ClazzName2    string `json:"_name"`
	UserId        int64  `json:"user_id"`
	AuthKeyId     int64  `json:"auth_key_id"`
	Gateway       string `json:"gateway"`
	Expired       int64  `json:"expired"`
	Layer         int32  `json:"layer"`
	PermAuthKeyId int64  `json:"perm_auth_key_id"`
	Client        string `json:"client"`
}

func MakeTLSessionEntry(m *TLSessionEntry) *TLSessionEntry {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_sessionEntry

	return m
}

func (m *TLSessionEntry) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLSessionEntry) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("sessionEntry", m)
}

// SessionEntryClazzName <--
func (m *TLSessionEntry) SessionEntryClazzName() string {
	return ClazzName_sessionEntry
}

// ClazzName <--
func (m *TLSessionEntry) ClazzName() string {
	return m.ClazzName2
}

// ToSessionEntry <--
func (m *TLSessionEntry) ToSessionEntry() *SessionEntry {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLSessionEntry) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_sessionEntry, int(layer)); clazzId {
	case 0x1764ac31:
		size := 4
		size += 8
		size += 8
		size += iface.CalcStringSize(m.Gateway)
		size += 8
		size += 4
		size += 8
		size += iface.CalcStringSize(m.Client)

		return size
	default:
		return 0
	}
}

func (m *TLSessionEntry) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_sessionEntry, int(layer)); clazzId {
	case 0x1764ac31:
		if err := iface.ValidateRequiredString("gateway", m.Gateway); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("client", m.Client); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_sessionEntry, layer)
	}
}

// Encode <--
func (m *TLSessionEntry) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_sessionEntry, int(layer)); clazzId {
	case 0x1764ac31:
		x.PutClazzID(0x1764ac31)

		x.PutInt64(m.UserId)
		x.PutInt64(m.AuthKeyId)
		x.PutString(m.Gateway)
		x.PutInt64(m.Expired)
		x.PutInt32(m.Layer)
		x.PutInt64(m.PermAuthKeyId)
		x.PutString(m.Client)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_sessionEntry, layer)
	}
}

// Decode <--
func (m *TLSessionEntry) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x1764ac31:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.Gateway, err = d.String()
		if err != nil {
			return err
		}
		m.Expired, err = d.Int64()
		if err != nil {
			return err
		}
		m.Layer, err = d.Int32()
		if err != nil {
			return err
		}
		m.PermAuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.Client, err = d.String()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// SessionEntry <--
type SessionEntry = TLSessionEntry

// UserSessionEntryListClazz <--
//   - TL_UserSessionEntryList
type UserSessionEntryListClazz = *TLUserSessionEntryList

func DecodeUserSessionEntryListClazz(d *bin.Decoder) (UserSessionEntryListClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0xefecb398:
		x := &TLUserSessionEntryList{ClazzID: id, ClazzName2: ClazzName_userSessionEntryList}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeUserSessionEntryList - unexpected clazzId: %d", id)
	}

}

// TLUserSessionEntryList <--
type TLUserSessionEntryList struct {
	ClazzID      uint32              `json:"_id"`
	ClazzName2   string              `json:"_name"`
	UserId       int64               `json:"user_id"`
	UserSessions []SessionEntryClazz `json:"user_sessions"`
}

func MakeTLUserSessionEntryList(m *TLUserSessionEntryList) *TLUserSessionEntryList {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_userSessionEntryList

	return m
}

func (m *TLUserSessionEntryList) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLUserSessionEntryList) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("userSessionEntryList", m)
}

// UserSessionEntryListClazzName <--
func (m *TLUserSessionEntryList) UserSessionEntryListClazzName() string {
	return ClazzName_userSessionEntryList
}

// ClazzName <--
func (m *TLUserSessionEntryList) ClazzName() string {
	return m.ClazzName2
}

// ToUserSessionEntryList <--
func (m *TLUserSessionEntryList) ToUserSessionEntryList() *UserSessionEntryList {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLUserSessionEntryList) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userSessionEntryList, int(layer)); clazzId {
	case 0xefecb398:
		size := 4
		size += 8
		size += iface.CalcObjectListSize(m.UserSessions, layer)

		return size
	default:
		return 0
	}
}

func (m *TLUserSessionEntryList) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userSessionEntryList, int(layer)); clazzId {
	case 0xefecb398:
		if err := iface.ValidateRequiredSlice("user_sessions", m.UserSessions); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_userSessionEntryList, layer)
	}
}

// Encode <--
func (m *TLUserSessionEntryList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userSessionEntryList, int(layer)); clazzId {
	case 0xefecb398:
		x.PutClazzID(0xefecb398)

		x.PutInt64(m.UserId)

		if err := iface.EncodeObjectList(x, m.UserSessions, layer); err != nil {
			return err
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_userSessionEntryList, layer)
	}
}

// Decode <--
func (m *TLUserSessionEntryList) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xefecb398:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		c1, err2 := d.ClazzID()
		if err2 != nil {
			return err2
		}
		if c1 != iface.ClazzID_vector {
			return fmt.Errorf("invalid ClazzID_vector, c%d: %d", 1, c1)
		}
		l1, err3 := d.Int()
		if err3 != nil {
			return err3
		}
		v1 := make([]SessionEntryClazz, l1)
		for i := 0; i < l1; i++ {
			v1[i], err3 = DecodeSessionEntryClazz(d)
			if err3 != nil {
				return err3
			}
		}
		m.UserSessions = v1

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// UserSessionEntryList <--
type UserSessionEntryList = TLUserSessionEntryList
