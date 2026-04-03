/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package user

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

// BotInfoDataClazz <--
//   - TL_BotInfoData
type BotInfoDataClazz = *TLBotInfoData

func DecodeBotInfoDataClazz(d *bin.Decoder) (BotInfoDataClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x1835d1c:
		x := &TLBotInfoData{ClazzID: id, ClazzName2: ClazzName_botInfoData}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
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

func (m *TLBotInfoData) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("botInfoData", m)
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

	return m

}

func (m *TLBotInfoData) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_botInfoData, int(layer)); clazzId {
	case 0x1835d1c:
		size := 4
		size += 4
		size += iface.CalcObjectSize(m.BotInfo, layer)
		if m.MainAppUrl != nil {
			size += iface.CalcStringSize(*m.MainAppUrl)
		}

		size += iface.CalcStringSize(m.Token)
		size += 8

		return size
	default:
		return 0
	}
}

func (m *TLBotInfoData) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_botInfoData, int(layer)); clazzId {
	case 0x1835d1c:
		if err := iface.ValidateRequiredObject("bot_info", m.BotInfo); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("token", m.Token); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_botInfoData, layer)
	}
}

// Encode <--
func (m *TLBotInfoData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_botInfoData, int(layer)); clazzId {
	case 0x1835d1c:
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
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_botInfoData, layer)
	}
}

// Decode <--
func (m *TLBotInfoData) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x1835d1c:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags

		m.BotInfo, err = tg.DecodeBotInfoClazz(d)
		if err != nil {
			return err
		}

		if (flags & (1 << 0)) != 0 {
			m.MainAppUrl = new(string)
			*m.MainAppUrl, err = d.String()
			if err != nil {
				return err
			}
		}

		if (flags & (1 << 1)) != 0 {
			m.BotInline = true
		}
		m.Token, err = d.String()
		if err != nil {
			return err
		}
		m.BotId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// BotInfoData <--
type BotInfoData = TLBotInfoData

// LastSeenDataClazz <--
//   - TL_LastSeenData
type LastSeenDataClazz = *TLLastSeenData

func DecodeLastSeenDataClazz(d *bin.Decoder) (LastSeenDataClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0xb3b1a1df:
		x := &TLLastSeenData{ClazzID: id, ClazzName2: ClazzName_lastSeenData}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
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

func (m *TLLastSeenData) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("lastSeenData", m)
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

	return m

}

func (m *TLLastSeenData) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_lastSeenData, int(layer)); clazzId {
	case 0xb3b1a1df:
		size := 4
		size += 8
		size += 8
		size += 4

		return size
	default:
		return 0
	}
}

func (m *TLLastSeenData) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_lastSeenData, int(layer)); clazzId {
	case 0xb3b1a1df:

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_lastSeenData, layer)
	}
}

// Encode <--
func (m *TLLastSeenData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_lastSeenData, int(layer)); clazzId {
	case 0xb3b1a1df:
		x.PutClazzID(0xb3b1a1df)

		x.PutInt64(m.UserId)
		x.PutInt64(m.LastSeenAt)
		x.PutInt32(m.Expires)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_lastSeenData, layer)
	}
}

// Decode <--
func (m *TLLastSeenData) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xb3b1a1df:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.LastSeenAt, err = d.Int64()
		if err != nil {
			return err
		}
		m.Expires, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// LastSeenData <--
type LastSeenData = TLLastSeenData

// PeerPeerNotifySettingsClazz <--
//   - TL_PeerPeerNotifySettings
type PeerPeerNotifySettingsClazz = *TLPeerPeerNotifySettings

func DecodePeerPeerNotifySettingsClazz(d *bin.Decoder) (PeerPeerNotifySettingsClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x70ea3fa9:
		x := &TLPeerPeerNotifySettings{ClazzID: id, ClazzName2: ClazzName_peerPeerNotifySettings}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
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

func (m *TLPeerPeerNotifySettings) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("peerPeerNotifySettings", m)
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

	return m

}

func (m *TLPeerPeerNotifySettings) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_peerPeerNotifySettings, int(layer)); clazzId {
	case 0x70ea3fa9:
		size := 4
		size += 4
		size += 8
		size += iface.CalcObjectSize(m.Settings, layer)

		return size
	default:
		return 0
	}
}

func (m *TLPeerPeerNotifySettings) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_peerPeerNotifySettings, int(layer)); clazzId {
	case 0x70ea3fa9:
		if err := iface.ValidateRequiredObject("settings", m.Settings); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_peerPeerNotifySettings, layer)
	}
}

// Encode <--
func (m *TLPeerPeerNotifySettings) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_peerPeerNotifySettings, int(layer)); clazzId {
	case 0x70ea3fa9:
		x.PutClazzID(0x70ea3fa9)

		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		_ = m.Settings.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_peerPeerNotifySettings, layer)
	}
}

// Decode <--
func (m *TLPeerPeerNotifySettings) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x70ea3fa9:
		m.PeerType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return err
		}

		m.Settings, err = tg.DecodePeerNotifySettingsClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// PeerPeerNotifySettings <--
type PeerPeerNotifySettings = TLPeerPeerNotifySettings

// UserImportedContactsClazz <--
//   - TL_UserImportedContacts
type UserImportedContactsClazz = *TLUserImportedContacts

func DecodeUserImportedContactsClazz(d *bin.Decoder) (UserImportedContactsClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x4adf7bc0:
		x := &TLUserImportedContacts{ClazzID: id, ClazzName2: ClazzName_userImportedContacts}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
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

func (m *TLUserImportedContacts) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("userImportedContacts", m)
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

	return m

}

func (m *TLUserImportedContacts) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userImportedContacts, int(layer)); clazzId {
	case 0x4adf7bc0:
		size := 4
		size += iface.CalcObjectListSize(m.Imported, layer)
		size += iface.CalcObjectListSize(m.PopularInvites, layer)
		size += iface.CalcInt64ListSize(m.RetryContacts)
		size += iface.CalcObjectListSize(m.Users, layer)
		size += iface.CalcInt64ListSize(m.UpdateIdList)

		return size
	default:
		return 0
	}
}

func (m *TLUserImportedContacts) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userImportedContacts, int(layer)); clazzId {
	case 0x4adf7bc0:
		if err := iface.ValidateRequiredSlice("imported", m.Imported); err != nil {
			return err
		}

		if err := iface.ValidateRequiredSlice("popular_invites", m.PopularInvites); err != nil {
			return err
		}

		if err := iface.ValidateRequiredSlice("retry_contacts", m.RetryContacts); err != nil {
			return err
		}

		if err := iface.ValidateRequiredSlice("users", m.Users); err != nil {
			return err
		}

		if err := iface.ValidateRequiredSlice("update_id_list", m.UpdateIdList); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_userImportedContacts, layer)
	}
}

// Encode <--
func (m *TLUserImportedContacts) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userImportedContacts, int(layer)); clazzId {
	case 0x4adf7bc0:
		x.PutClazzID(0x4adf7bc0)

		if err := iface.EncodeObjectList(x, m.Imported, layer); err != nil {
			return err
		}

		if err := iface.EncodeObjectList(x, m.PopularInvites, layer); err != nil {
			return err
		}

		iface.EncodeInt64List(x, m.RetryContacts)

		if err := iface.EncodeObjectList(x, m.Users, layer); err != nil {
			return err
		}

		iface.EncodeInt64List(x, m.UpdateIdList)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_userImportedContacts, layer)
	}
}

// Decode <--
func (m *TLUserImportedContacts) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x4adf7bc0:
		c0, err2 := d.ClazzID()
		if err2 != nil {
			return err2
		}
		if c0 != iface.ClazzID_vector {
			return fmt.Errorf("invalid ClazzID_vector, c%d: %d", 0, c0)
		}
		l0, err3 := d.Int()
		if err3 != nil {
			return err3
		}
		v0 := make([]tg.ImportedContactClazz, l0)
		for i := 0; i < l0; i++ {
			v0[i], err3 = tg.DecodeImportedContactClazz(d)
			if err3 != nil {
				return err3
			}
		}
		m.Imported = v0

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
		v1 := make([]tg.PopularContactClazz, l1)
		for i := 0; i < l1; i++ {
			v1[i], err3 = tg.DecodePopularContactClazz(d)
			if err3 != nil {
				return err3
			}
		}
		m.PopularInvites = v1

		m.RetryContacts, err = iface.DecodeInt64List(d)
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
		v3 := make([]tg.UserClazz, l3)
		for i := 0; i < l3; i++ {
			v3[i], err3 = tg.DecodeUserClazz(d)
			if err3 != nil {
				return err3
			}
		}
		m.Users = v3

		m.UpdateIdList, err = iface.DecodeInt64List(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// UserImportedContacts <--
type UserImportedContacts = TLUserImportedContacts

// UsernameDataClazz <--
//   - TL_UsernameData
type UsernameDataClazz = *TLUsernameData

func DecodeUsernameDataClazz(d *bin.Decoder) (UsernameDataClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0xaa4000bf:
		x := &TLUsernameData{ClazzID: id, ClazzName2: ClazzName_usernameData}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeUsernameData - unexpected clazzId: %d", id)
	}

}

// TLUsernameData <--
type TLUsernameData struct {
	ClazzID    uint32       `json:"_id"`
	ClazzName2 string       `json:"_name"`
	Username   string       `json:"username"`
	Peer       tg.PeerClazz `json:"peer"`
	Editable   bool         `json:"editable"`
	Active     bool         `json:"active"`
}

func MakeTLUsernameData(m *TLUsernameData) *TLUsernameData {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_usernameData

	return m
}

func (m *TLUsernameData) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLUsernameData) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("usernameData", m)
}

// UsernameDataClazzName <--
func (m *TLUsernameData) UsernameDataClazzName() string {
	return ClazzName_usernameData
}

// ClazzName <--
func (m *TLUsernameData) ClazzName() string {
	return m.ClazzName2
}

// ToUsernameData <--
func (m *TLUsernameData) ToUsernameData() *UsernameData {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLUsernameData) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usernameData, int(layer)); clazzId {
	case 0xaa4000bf:
		size := 4
		size += 4
		size += iface.CalcStringSize(m.Username)
		if m.Peer != nil {
			size += iface.CalcObjectSize(m.Peer, layer)
		}

		return size
	default:
		return 0
	}
}

func (m *TLUsernameData) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usernameData, int(layer)); clazzId {
	case 0xaa4000bf:
		if err := iface.ValidateRequiredString("username", m.Username); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usernameData, layer)
	}
}

// Encode <--
func (m *TLUsernameData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usernameData, int(layer)); clazzId {
	case 0xaa4000bf:
		x.PutClazzID(0xaa4000bf)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.Peer != nil {
				flags |= 1 << 0
			}
			if m.Editable == true {
				flags |= 1 << 1
			}
			if m.Active == true {
				flags |= 1 << 2
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutString(m.Username)
		if m.Peer != nil {
			_ = m.Peer.Encode(x, layer)
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usernameData, layer)
	}
}

// Decode <--
func (m *TLUsernameData) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xaa4000bf:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.Username, err = d.String()
		if err != nil {
			return err
		}
		if (flags & (1 << 0)) != 0 {
			m.Peer, err = tg.DecodePeerClazz(d)
			if err != nil {
				return err
			}
		}
		if (flags & (1 << 1)) != 0 {
			m.Editable = true
		}
		if (flags & (1 << 2)) != 0 {
			m.Active = true
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// UsernameData <--
type UsernameData = TLUsernameData

// UsernameExistClazz <--
//   - TL_UsernameNotExisted
//   - TL_UsernameExisted
//   - TL_UsernameExistedNotMe
//   - TL_UsernameExistedIsMe
type UsernameExistClazz interface {
	iface.TLObject
	UsernameExistClazzName() string
}

func DecodeUsernameExistClazz(d *bin.Decoder) (UsernameExistClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0xcb3cfb6d:
		x := &TLUsernameNotExisted{ClazzID: id, ClazzName2: ClazzName_usernameNotExisted}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0xace7f4cd:
		x := &TLUsernameExisted{ClazzID: id, ClazzName2: ClazzName_usernameExisted}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0xd01f47b1:
		x := &TLUsernameExistedNotMe{ClazzID: id, ClazzName2: ClazzName_usernameExistedNotMe}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x874e7771:
		x := &TLUsernameExistedIsMe{ClazzID: id, ClazzName2: ClazzName_usernameExistedIsMe}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeUsernameExist - unexpected clazzId: %d", id)
	}

}

// TLUsernameNotExisted <--
type TLUsernameNotExisted struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
}

func MakeTLUsernameNotExisted(m *TLUsernameNotExisted) *TLUsernameNotExisted {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_usernameNotExisted

	return m
}

func (m *TLUsernameNotExisted) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLUsernameNotExisted) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("usernameNotExisted", m)
}

// UsernameExistClazzName <--
func (m *TLUsernameNotExisted) UsernameExistClazzName() string {
	return ClazzName_usernameNotExisted
}

// ClazzName <--
func (m *TLUsernameNotExisted) ClazzName() string {
	return m.ClazzName2
}

// ToUsernameExist <--
func (m *TLUsernameNotExisted) ToUsernameExist() *UsernameExist {
	if m == nil {
		return nil
	}

	return &UsernameExist{Clazz: m}

}

func (m *TLUsernameNotExisted) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usernameNotExisted, int(layer)); clazzId {
	case 0xcb3cfb6d:
		size := 4

		return size
	default:
		return 0
	}
}

func (m *TLUsernameNotExisted) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usernameNotExisted, int(layer)); clazzId {
	case 0xcb3cfb6d:

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usernameNotExisted, layer)
	}
}

// Encode <--
func (m *TLUsernameNotExisted) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usernameNotExisted, int(layer)); clazzId {
	case 0xcb3cfb6d:
		x.PutClazzID(0xcb3cfb6d)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usernameNotExisted, layer)
	}
}

// Decode <--
func (m *TLUsernameNotExisted) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xcb3cfb6d:

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUsernameExisted <--
type TLUsernameExisted struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
}

func MakeTLUsernameExisted(m *TLUsernameExisted) *TLUsernameExisted {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_usernameExisted

	return m
}

func (m *TLUsernameExisted) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLUsernameExisted) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("usernameExisted", m)
}

// UsernameExistClazzName <--
func (m *TLUsernameExisted) UsernameExistClazzName() string {
	return ClazzName_usernameExisted
}

// ClazzName <--
func (m *TLUsernameExisted) ClazzName() string {
	return m.ClazzName2
}

// ToUsernameExist <--
func (m *TLUsernameExisted) ToUsernameExist() *UsernameExist {
	if m == nil {
		return nil
	}

	return &UsernameExist{Clazz: m}

}

func (m *TLUsernameExisted) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usernameExisted, int(layer)); clazzId {
	case 0xace7f4cd:
		size := 4

		return size
	default:
		return 0
	}
}

func (m *TLUsernameExisted) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usernameExisted, int(layer)); clazzId {
	case 0xace7f4cd:

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usernameExisted, layer)
	}
}

// Encode <--
func (m *TLUsernameExisted) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usernameExisted, int(layer)); clazzId {
	case 0xace7f4cd:
		x.PutClazzID(0xace7f4cd)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usernameExisted, layer)
	}
}

// Decode <--
func (m *TLUsernameExisted) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xace7f4cd:

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUsernameExistedNotMe <--
type TLUsernameExistedNotMe struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
}

func MakeTLUsernameExistedNotMe(m *TLUsernameExistedNotMe) *TLUsernameExistedNotMe {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_usernameExistedNotMe

	return m
}

func (m *TLUsernameExistedNotMe) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLUsernameExistedNotMe) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("usernameExistedNotMe", m)
}

// UsernameExistClazzName <--
func (m *TLUsernameExistedNotMe) UsernameExistClazzName() string {
	return ClazzName_usernameExistedNotMe
}

// ClazzName <--
func (m *TLUsernameExistedNotMe) ClazzName() string {
	return m.ClazzName2
}

// ToUsernameExist <--
func (m *TLUsernameExistedNotMe) ToUsernameExist() *UsernameExist {
	if m == nil {
		return nil
	}

	return &UsernameExist{Clazz: m}

}

func (m *TLUsernameExistedNotMe) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usernameExistedNotMe, int(layer)); clazzId {
	case 0xd01f47b1:
		size := 4

		return size
	default:
		return 0
	}
}

func (m *TLUsernameExistedNotMe) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usernameExistedNotMe, int(layer)); clazzId {
	case 0xd01f47b1:

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usernameExistedNotMe, layer)
	}
}

// Encode <--
func (m *TLUsernameExistedNotMe) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usernameExistedNotMe, int(layer)); clazzId {
	case 0xd01f47b1:
		x.PutClazzID(0xd01f47b1)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usernameExistedNotMe, layer)
	}
}

// Decode <--
func (m *TLUsernameExistedNotMe) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xd01f47b1:

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUsernameExistedIsMe <--
type TLUsernameExistedIsMe struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
}

func MakeTLUsernameExistedIsMe(m *TLUsernameExistedIsMe) *TLUsernameExistedIsMe {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_usernameExistedIsMe

	return m
}

func (m *TLUsernameExistedIsMe) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLUsernameExistedIsMe) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("usernameExistedIsMe", m)
}

// UsernameExistClazzName <--
func (m *TLUsernameExistedIsMe) UsernameExistClazzName() string {
	return ClazzName_usernameExistedIsMe
}

// ClazzName <--
func (m *TLUsernameExistedIsMe) ClazzName() string {
	return m.ClazzName2
}

// ToUsernameExist <--
func (m *TLUsernameExistedIsMe) ToUsernameExist() *UsernameExist {
	if m == nil {
		return nil
	}

	return &UsernameExist{Clazz: m}

}

func (m *TLUsernameExistedIsMe) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usernameExistedIsMe, int(layer)); clazzId {
	case 0x874e7771:
		size := 4

		return size
	default:
		return 0
	}
}

func (m *TLUsernameExistedIsMe) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usernameExistedIsMe, int(layer)); clazzId {
	case 0x874e7771:

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usernameExistedIsMe, layer)
	}
}

// Encode <--
func (m *TLUsernameExistedIsMe) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usernameExistedIsMe, int(layer)); clazzId {
	case 0x874e7771:
		x.PutClazzID(0x874e7771)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usernameExistedIsMe, layer)
	}
}

// Decode <--
func (m *TLUsernameExistedIsMe) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x874e7771:

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// UsernameExist <--
type UsernameExist struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz UsernameExistClazz `json:"_clazz"`
}

func (m *UsernameExist) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *UsernameExist) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName(m.ClazzName(), m)
}

func (m *UsernameExist) CalcSize(layer int32) int {
	if m == nil || m.Clazz == nil {
		return 0
	}
	return iface.CalcObjectSize(m.Clazz, layer)
}

func (m *UsernameExist) Validate(layer int32) error {
	if m == nil {
		return fmt.Errorf("UsernameExist is required")
	}
	if m.Clazz == nil {
		return fmt.Errorf("UsernameExist.Clazz is required")
	}
	if v, ok := m.Clazz.(iface.TLObjectValidator); ok {
		return v.Validate(layer)
	}
	return nil
}

func (m *UsernameExist) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.UsernameExistClazzName()
	}
}

// Encode <--
func (m *UsernameExist) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("UsernameExist - invalid Clazz")
}

// Decode <--
func (m *UsernameExist) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeUsernameExistClazz(d)
	return
}

// ToUsernameNotExisted <--
func (m *UsernameExist) ToUsernameNotExisted() (*TLUsernameNotExisted, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLUsernameNotExisted); ok {
		return x, true
	}

	return nil, false
}

// ToUsernameExisted <--
func (m *UsernameExist) ToUsernameExisted() (*TLUsernameExisted, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLUsernameExisted); ok {
		return x, true
	}

	return nil, false
}

// ToUsernameExistedNotMe <--
func (m *UsernameExist) ToUsernameExistedNotMe() (*TLUsernameExistedNotMe, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLUsernameExistedNotMe); ok {
		return x, true
	}

	return nil, false
}

// ToUsernameExistedIsMe <--
func (m *UsernameExist) ToUsernameExistedIsMe() (*TLUsernameExistedIsMe, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLUsernameExistedIsMe); ok {
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

	switch id {
	case 0x3fa3dbc7:
		x := &TLUsersDataFound{ClazzID: id, ClazzName2: ClazzName_usersDataFound}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x80c4adfa:
		x := &TLUsersIdFound{ClazzID: id, ClazzName2: ClazzName_usersIdFound}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
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

func (m *TLUsersDataFound) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("usersDataFound", m)
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

func (m *TLUsersDataFound) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usersDataFound, int(layer)); clazzId {
	case 0x3fa3dbc7:
		size := 4
		size += 4
		size += iface.CalcObjectListSize(m.Users, layer)
		size += iface.CalcStringSize(m.NextOffset)

		return size
	default:
		return 0
	}
}

func (m *TLUsersDataFound) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usersDataFound, int(layer)); clazzId {
	case 0x3fa3dbc7:
		if err := iface.ValidateRequiredSlice("users", m.Users); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("next_offset", m.NextOffset); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usersDataFound, layer)
	}
}

// Encode <--
func (m *TLUsersDataFound) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usersDataFound, int(layer)); clazzId {
	case 0x3fa3dbc7:
		x.PutClazzID(0x3fa3dbc7)

		x.PutInt32(m.Count)

		if err := iface.EncodeObjectList(x, m.Users, layer); err != nil {
			return err
		}

		x.PutString(m.NextOffset)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usersDataFound, layer)
	}
}

// Decode <--
func (m *TLUsersDataFound) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x3fa3dbc7:
		m.Count, err = d.Int32()
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
		v1 := make([]tg.UserDataClazz, l1)
		for i := 0; i < l1; i++ {
			v1[i], err3 = tg.DecodeUserDataClazz(d)
			if err3 != nil {
				return err3
			}
		}
		m.Users = v1

		m.NextOffset, err = d.String()
		if err != nil {
			return err
		}

		return nil
	default:
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

func (m *TLUsersIdFound) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("usersIdFound", m)
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

func (m *TLUsersIdFound) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usersIdFound, int(layer)); clazzId {
	case 0x80c4adfa:
		size := 4
		size += iface.CalcInt64ListSize(m.IdList)

		return size
	default:
		return 0
	}
}

func (m *TLUsersIdFound) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usersIdFound, int(layer)); clazzId {
	case 0x80c4adfa:
		if err := iface.ValidateRequiredSlice("id_list", m.IdList); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usersIdFound, layer)
	}
}

// Encode <--
func (m *TLUsersIdFound) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_usersIdFound, int(layer)); clazzId {
	case 0x80c4adfa:
		x.PutClazzID(0x80c4adfa)

		iface.EncodeInt64List(x, m.IdList)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usersIdFound, layer)
	}
}

// Decode <--
func (m *TLUsersIdFound) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x80c4adfa:

		m.IdList, err = iface.DecodeInt64List(d)
		if err != nil {
			return err
		}

		return nil
	default:
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

func (m *UsersFound) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName(m.ClazzName(), m)
}

func (m *UsersFound) CalcSize(layer int32) int {
	if m == nil || m.Clazz == nil {
		return 0
	}
	return iface.CalcObjectSize(m.Clazz, layer)
}

func (m *UsersFound) Validate(layer int32) error {
	if m == nil {
		return fmt.Errorf("UsersFound is required")
	}
	if m.Clazz == nil {
		return fmt.Errorf("UsersFound.Clazz is required")
	}
	if v, ok := m.Clazz.(iface.TLObjectValidator); ok {
		return v.Validate(layer)
	}
	return nil
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
