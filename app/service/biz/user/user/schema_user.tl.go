/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package user

import (
	"encoding/json"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
)

var _ iface.TLObject
var _ fmt.Stringer
var _ *tg.Bool
var _ bin.Fields

// BotInfoDataClazz <--
//   - TL_BotInfoData
type BotInfoDataClazz interface {
	iface.TLObject
	BotInfoDataClazzName() string
}

func DecodeBotInfoDataClazz(d *bin.Decoder) (BotInfoDataClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_botInfoData:
		x := &TLBotInfoData{ClazzID: id, ClazzName2: ClazzName_botInfoData}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeBotInfoData - unexpected clazzId: %d", id)
	}
}

// TLBotInfoData <--
type TLBotInfoData struct {
	ClazzID    uint32          `json:"_id"`
	ClazzName2 string          `json:"_name"`
	BotInfo    tg.BotInfoClazz `json:"bot_info"`
	MainAppUrl *string         `json:"main_app_url"`
	BotInline  bool            `json:"bot_inline"`
	Token      string          `json:"token"`
	BotId      int64           `json:"bot_id"`
}

func MakeTLBotInfoData(m *TLBotInfoData) *TLBotInfoData {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_botInfoData

	return m
}

func (m *TLBotInfoData) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// BotInfoDataClazzName <--
func (m *TLBotInfoData) BotInfoDataClazzName() string {
	return ClazzName_botInfoData
}

// ClazzName <--
func (m *TLBotInfoData) ClazzName() string {
	return m.ClazzName2
}

// ToBotInfoData <--
func (m *TLBotInfoData) ToBotInfoData() *BotInfoData {
	if m == nil {
		return nil
	}

	return &BotInfoData{Clazz: m}
}

// Encode <--
func (m *TLBotInfoData) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x1835d1c: func() error {
			x.PutClazzID(0x1835d1c)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.MainAppUrl != nil {
					flags |= 1 << 0
				}
				if m.BotInline == true {
					flags |= 1 << 1
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			_ = m.BotInfo.Encode(x, layer)
			if m.MainAppUrl != nil {
				x.PutString(*m.MainAppUrl)
			}

			x.PutString(m.Token)
			x.PutInt64(m.BotId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_botInfoData, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_botInfoData, layer)
	}
}

// Decode <--
func (m *TLBotInfoData) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x1835d1c: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags

			// m1 := &tg.BotInfo{}
			// _ = m1.Decode(d)
			// m.BotInfo = m1
			m.BotInfo, _ = tg.DecodeBotInfoClazz(d)

			if (flags & (1 << 0)) != 0 {
				m.MainAppUrl = new(string)
				*m.MainAppUrl, err = d.String()
			}

			if (flags & (1 << 1)) != 0 {
				m.BotInline = true
			}
			m.Token, err = d.String()
			m.BotId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// BotInfoData <--
type BotInfoData struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz BotInfoDataClazz `json:"_clazz"`
}

func (m *BotInfoData) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *BotInfoData) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.BotInfoDataClazzName()
	}
}

// Encode <--
func (m *BotInfoData) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("BotInfoData - invalid Clazz")
}

// Decode <--
func (m *BotInfoData) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeBotInfoDataClazz(d)
	return
}

// Match <--
func (m *BotInfoData) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLBotInfoData:
		for _, v := range f {
			if f1, ok := v.(func(c *TLBotInfoData) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToBotInfoData <--
func (m *BotInfoData) ToBotInfoData() (*TLBotInfoData, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLBotInfoData); ok {
		return x, true
	}

	return nil, false
}

// LastSeenDataClazz <--
//   - TL_LastSeenData
type LastSeenDataClazz interface {
	iface.TLObject
	LastSeenDataClazzName() string
}

func DecodeLastSeenDataClazz(d *bin.Decoder) (LastSeenDataClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_lastSeenData:
		x := &TLLastSeenData{ClazzID: id, ClazzName2: ClazzName_lastSeenData}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeLastSeenData - unexpected clazzId: %d", id)
	}
}

// TLLastSeenData <--
type TLLastSeenData struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	UserId     int64  `json:"user_id"`
	LastSeenAt int64  `json:"last_seen_at"`
	Expires    int32  `json:"expires"`
}

func MakeTLLastSeenData(m *TLLastSeenData) *TLLastSeenData {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_lastSeenData

	return m
}

func (m *TLLastSeenData) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// LastSeenDataClazzName <--
func (m *TLLastSeenData) LastSeenDataClazzName() string {
	return ClazzName_lastSeenData
}

// ClazzName <--
func (m *TLLastSeenData) ClazzName() string {
	return m.ClazzName2
}

// ToLastSeenData <--
func (m *TLLastSeenData) ToLastSeenData() *LastSeenData {
	if m == nil {
		return nil
	}

	return &LastSeenData{Clazz: m}
}

// Encode <--
func (m *TLLastSeenData) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xb3b1a1df: func() error {
			x.PutClazzID(0xb3b1a1df)

			x.PutInt64(m.UserId)
			x.PutInt64(m.LastSeenAt)
			x.PutInt32(m.Expires)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_lastSeenData, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_lastSeenData, layer)
	}
}

// Decode <--
func (m *TLLastSeenData) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xb3b1a1df: func() (err error) {
			m.UserId, err = d.Int64()
			m.LastSeenAt, err = d.Int64()
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

// LastSeenData <--
type LastSeenData struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz LastSeenDataClazz `json:"_clazz"`
}

func (m *LastSeenData) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *LastSeenData) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.LastSeenDataClazzName()
	}
}

// Encode <--
func (m *LastSeenData) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("LastSeenData - invalid Clazz")
}

// Decode <--
func (m *LastSeenData) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeLastSeenDataClazz(d)
	return
}

// Match <--
func (m *LastSeenData) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLLastSeenData:
		for _, v := range f {
			if f1, ok := v.(func(c *TLLastSeenData) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToLastSeenData <--
func (m *LastSeenData) ToLastSeenData() (*TLLastSeenData, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLLastSeenData); ok {
		return x, true
	}

	return nil, false
}

// PeerPeerNotifySettingsClazz <--
//   - TL_PeerPeerNotifySettings
type PeerPeerNotifySettingsClazz interface {
	iface.TLObject
	PeerPeerNotifySettingsClazzName() string
}

func DecodePeerPeerNotifySettingsClazz(d *bin.Decoder) (PeerPeerNotifySettingsClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_peerPeerNotifySettings:
		x := &TLPeerPeerNotifySettings{ClazzID: id, ClazzName2: ClazzName_peerPeerNotifySettings}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodePeerPeerNotifySettings - unexpected clazzId: %d", id)
	}
}

// TLPeerPeerNotifySettings <--
type TLPeerPeerNotifySettings struct {
	ClazzID    uint32                     `json:"_id"`
	ClazzName2 string                     `json:"_name"`
	PeerType   int32                      `json:"peer_type"`
	PeerId     int64                      `json:"peer_id"`
	Settings   tg.PeerNotifySettingsClazz `json:"settings"`
}

func MakeTLPeerPeerNotifySettings(m *TLPeerPeerNotifySettings) *TLPeerPeerNotifySettings {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_peerPeerNotifySettings

	return m
}

func (m *TLPeerPeerNotifySettings) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// PeerPeerNotifySettingsClazzName <--
func (m *TLPeerPeerNotifySettings) PeerPeerNotifySettingsClazzName() string {
	return ClazzName_peerPeerNotifySettings
}

// ClazzName <--
func (m *TLPeerPeerNotifySettings) ClazzName() string {
	return m.ClazzName2
}

// ToPeerPeerNotifySettings <--
func (m *TLPeerPeerNotifySettings) ToPeerPeerNotifySettings() *PeerPeerNotifySettings {
	if m == nil {
		return nil
	}

	return &PeerPeerNotifySettings{Clazz: m}
}

// Encode <--
func (m *TLPeerPeerNotifySettings) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x70ea3fa9: func() error {
			x.PutClazzID(0x70ea3fa9)

			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			_ = m.Settings.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_peerPeerNotifySettings, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_peerPeerNotifySettings, layer)
	}
}

// Decode <--
func (m *TLPeerPeerNotifySettings) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x70ea3fa9: func() (err error) {
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()

			// m2 := &tg.PeerNotifySettings{}
			// _ = m2.Decode(d)
			// m.Settings = m2
			m.Settings, _ = tg.DecodePeerNotifySettingsClazz(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// PeerPeerNotifySettings <--
type PeerPeerNotifySettings struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz PeerPeerNotifySettingsClazz `json:"_clazz"`
}

func (m *PeerPeerNotifySettings) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *PeerPeerNotifySettings) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.PeerPeerNotifySettingsClazzName()
	}
}

// Encode <--
func (m *PeerPeerNotifySettings) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("PeerPeerNotifySettings - invalid Clazz")
}

// Decode <--
func (m *PeerPeerNotifySettings) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodePeerPeerNotifySettingsClazz(d)
	return
}

// Match <--
func (m *PeerPeerNotifySettings) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLPeerPeerNotifySettings:
		for _, v := range f {
			if f1, ok := v.(func(c *TLPeerPeerNotifySettings) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToPeerPeerNotifySettings <--
func (m *PeerPeerNotifySettings) ToPeerPeerNotifySettings() (*TLPeerPeerNotifySettings, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLPeerPeerNotifySettings); ok {
		return x, true
	}

	return nil, false
}

// UserImportedContactsClazz <--
//   - TL_UserImportedContacts
type UserImportedContactsClazz interface {
	iface.TLObject
	UserImportedContactsClazzName() string
}

func DecodeUserImportedContactsClazz(d *bin.Decoder) (UserImportedContactsClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_userImportedContacts:
		x := &TLUserImportedContacts{ClazzID: id, ClazzName2: ClazzName_userImportedContacts}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeUserImportedContacts - unexpected clazzId: %d", id)
	}
}

// TLUserImportedContacts <--
type TLUserImportedContacts struct {
	ClazzID        uint32                    `json:"_id"`
	ClazzName2     string                    `json:"_name"`
	Imported       []tg.ImportedContactClazz `json:"imported"`
	PopularInvites []tg.PopularContactClazz  `json:"popular_invites"`
	RetryContacts  []int64                   `json:"retry_contacts"`
	Users          []tg.UserClazz            `json:"users"`
	UpdateIdList   []int64                   `json:"update_id_list"`
}

func MakeTLUserImportedContacts(m *TLUserImportedContacts) *TLUserImportedContacts {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_userImportedContacts

	return m
}

func (m *TLUserImportedContacts) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// UserImportedContactsClazzName <--
func (m *TLUserImportedContacts) UserImportedContactsClazzName() string {
	return ClazzName_userImportedContacts
}

// ClazzName <--
func (m *TLUserImportedContacts) ClazzName() string {
	return m.ClazzName2
}

// ToUserImportedContacts <--
func (m *TLUserImportedContacts) ToUserImportedContacts() *UserImportedContacts {
	if m == nil {
		return nil
	}

	return &UserImportedContacts{Clazz: m}
}

// Encode <--
func (m *TLUserImportedContacts) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x4adf7bc0: func() error {
			x.PutClazzID(0x4adf7bc0)

			_ = iface.EncodeObjectList(x, m.Imported, layer)

			_ = iface.EncodeObjectList(x, m.PopularInvites, layer)

			iface.EncodeInt64List(x, m.RetryContacts)

			_ = iface.EncodeObjectList(x, m.Users, layer)

			iface.EncodeInt64List(x, m.UpdateIdList)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_userImportedContacts, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_userImportedContacts, layer)
	}
}

// Decode <--
func (m *TLUserImportedContacts) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x4adf7bc0: func() (err error) {
			c0, err2 := d.ClazzID()
			if c0 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 0, c0)
				return err2
			}
			l0, err3 := d.Int()
			v0 := make([]tg.ImportedContactClazz, l0)
			for i := 0; i < l0; i++ {
				// vv := new(ImportedContact)
				// err3 = vv.Decode(d)
				// _ = err3
				// v0[i] = vv
				v0[i], err3 = tg.DecodeImportedContactClazz(d)
				_ = err3
			}
			m.Imported = v0

			c1, err2 := d.ClazzID()
			if c1 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 1, c1)
				return err2
			}
			l1, err3 := d.Int()
			v1 := make([]tg.PopularContactClazz, l1)
			for i := 0; i < l1; i++ {
				// vv := new(PopularContact)
				// err3 = vv.Decode(d)
				// _ = err3
				// v1[i] = vv
				v1[i], err3 = tg.DecodePopularContactClazz(d)
				_ = err3
			}
			m.PopularInvites = v1

			m.RetryContacts, err = iface.DecodeInt64List(d)

			c3, err2 := d.ClazzID()
			if c3 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 3, c3)
				return err2
			}
			l3, err3 := d.Int()
			v3 := make([]tg.UserClazz, l3)
			for i := 0; i < l3; i++ {
				// vv := new(User)
				// err3 = vv.Decode(d)
				// _ = err3
				// v3[i] = vv
				v3[i], err3 = tg.DecodeUserClazz(d)
				_ = err3
			}
			m.Users = v3

			m.UpdateIdList, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// UserImportedContacts <--
type UserImportedContacts struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz UserImportedContactsClazz `json:"_clazz"`
}

func (m *UserImportedContacts) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *UserImportedContacts) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.UserImportedContactsClazzName()
	}
}

// Encode <--
func (m *UserImportedContacts) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("UserImportedContacts - invalid Clazz")
}

// Decode <--
func (m *UserImportedContacts) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeUserImportedContactsClazz(d)
	return
}

// Match <--
func (m *UserImportedContacts) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLUserImportedContacts:
		for _, v := range f {
			if f1, ok := v.(func(c *TLUserImportedContacts) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToUserImportedContacts <--
func (m *UserImportedContacts) ToUserImportedContacts() (*TLUserImportedContacts, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLUserImportedContacts); ok {
		return x, true
	}

	return nil, false
}

// UsersFoundClazz <--
//   - TL_UsersDataFound
//   - TL_UsersIdFound
type UsersFoundClazz interface {
	iface.TLObject
	UsersFoundClazzName() string
}

func DecodeUsersFoundClazz(d *bin.Decoder) (UsersFoundClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_usersDataFound:
		x := &TLUsersDataFound{ClazzID: id, ClazzName2: ClazzName_usersDataFound}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_usersIdFound:
		x := &TLUsersIdFound{ClazzID: id, ClazzName2: ClazzName_usersIdFound}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeUsersFound - unexpected clazzId: %d", id)
	}
}

// TLUsersDataFound <--
type TLUsersDataFound struct {
	ClazzID    uint32             `json:"_id"`
	ClazzName2 string             `json:"_name"`
	Count      int32              `json:"count"`
	Users      []tg.UserDataClazz `json:"users"`
	NextOffset string             `json:"next_offset"`
}

func MakeTLUsersDataFound(m *TLUsersDataFound) *TLUsersDataFound {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_usersDataFound

	return m
}

func (m *TLUsersDataFound) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// UsersFoundClazzName <--
func (m *TLUsersDataFound) UsersFoundClazzName() string {
	return ClazzName_usersDataFound
}

// ClazzName <--
func (m *TLUsersDataFound) ClazzName() string {
	return m.ClazzName2
}

// ToUsersFound <--
func (m *TLUsersDataFound) ToUsersFound() *UsersFound {
	if m == nil {
		return nil
	}

	return &UsersFound{Clazz: m}
}

// Encode <--
func (m *TLUsersDataFound) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x3fa3dbc7: func() error {
			x.PutClazzID(0x3fa3dbc7)

			x.PutInt32(m.Count)

			_ = iface.EncodeObjectList(x, m.Users, layer)

			x.PutString(m.NextOffset)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_usersDataFound, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usersDataFound, layer)
	}
}

// Decode <--
func (m *TLUsersDataFound) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x3fa3dbc7: func() (err error) {
			m.Count, err = d.Int32()
			c1, err2 := d.ClazzID()
			if c1 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 1, c1)
				return err2
			}
			l1, err3 := d.Int()
			v1 := make([]tg.UserDataClazz, l1)
			for i := 0; i < l1; i++ {
				// vv := new(UserData)
				// err3 = vv.Decode(d)
				// _ = err3
				// v1[i] = vv
				v1[i], err3 = tg.DecodeUserDataClazz(d)
				_ = err3
			}
			m.Users = v1

			m.NextOffset, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUsersIdFound <--
type TLUsersIdFound struct {
	ClazzID    uint32  `json:"_id"`
	ClazzName2 string  `json:"_name"`
	IdList     []int64 `json:"id_list"`
}

func MakeTLUsersIdFound(m *TLUsersIdFound) *TLUsersIdFound {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_usersIdFound

	return m
}

func (m *TLUsersIdFound) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// UsersFoundClazzName <--
func (m *TLUsersIdFound) UsersFoundClazzName() string {
	return ClazzName_usersIdFound
}

// ClazzName <--
func (m *TLUsersIdFound) ClazzName() string {
	return m.ClazzName2
}

// ToUsersFound <--
func (m *TLUsersIdFound) ToUsersFound() *UsersFound {
	if m == nil {
		return nil
	}

	return &UsersFound{Clazz: m}
}

// Encode <--
func (m *TLUsersIdFound) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x80c4adfa: func() error {
			x.PutClazzID(0x80c4adfa)

			iface.EncodeInt64List(x, m.IdList)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_usersIdFound, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usersIdFound, layer)
	}
}

// Decode <--
func (m *TLUsersIdFound) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x80c4adfa: func() (err error) {

			m.IdList, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// UsersFound <--
type UsersFound struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz UsersFoundClazz `json:"_clazz"`
}

func (m *UsersFound) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *UsersFound) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.UsersFoundClazzName()
	}
}

// Encode <--
func (m *UsersFound) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("UsersFound - invalid Clazz")
}

// Decode <--
func (m *UsersFound) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeUsersFoundClazz(d)
	return
}

// Match <--
func (m *UsersFound) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLUsersDataFound:
		for _, v := range f {
			if f1, ok := v.(func(c *TLUsersDataFound) interface{}); ok {
				f1(c)
			}
		}
	case *TLUsersIdFound:
		for _, v := range f {
			if f1, ok := v.(func(c *TLUsersIdFound) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToUsersDataFound <--
func (m *UsersFound) ToUsersDataFound() (*TLUsersDataFound, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLUsersDataFound); ok {
		return x, true
	}

	return nil, false
}

// ToUsersIdFound <--
func (m *UsersFound) ToUsersIdFound() (*TLUsersIdFound, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLUsersIdFound); ok {
		return x, true
	}

	return nil, false
}
