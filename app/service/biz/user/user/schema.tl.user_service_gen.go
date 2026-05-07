/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package user

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

// TLUserGetLastSeens <--
type TLUserGetLastSeens struct {
	ClazzID uint32  `json:"_id"`
	Id      []int64 `json:"id"`
}

func (m *TLUserGetLastSeens) String() string {
	return iface.DebugStringWithName(ClazzName_user_getLastSeens, m)
}

// Encode <--
func (m *TLUserGetLastSeens) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getLastSeens, int(layer)); clazzId {
	case 0x7ca17e01:
		x.PutClazzID(0x7ca17e01)

		iface.EncodeInt64List(x, m.Id)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getLastSeens: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetLastSeens) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getLastSeens: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x7ca17e01:

		m.Id, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_getLastSeens#0x7ca17e01: field id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getLastSeens: invalid constructor %x", m.ClazzID)
	}
}

// TLUserUpdateLastSeen <--
type TLUserUpdateLastSeen struct {
	ClazzID    uint32 `json:"_id"`
	Id         int64  `json:"id"`
	LastSeenAt int64  `json:"last_seen_at"`
	Expires    int32  `json:"expires"`
}

func (m *TLUserUpdateLastSeen) String() string {
	return iface.DebugStringWithName(ClazzName_user_updateLastSeen, m)
}

// Encode <--
func (m *TLUserUpdateLastSeen) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_updateLastSeen, int(layer)); clazzId {
	case 0xfd405a2d:
		x.PutClazzID(0xfd405a2d)

		x.PutInt64(m.Id)
		x.PutInt64(m.LastSeenAt)
		x.PutInt32(m.Expires)

		return nil
	default:
		return fmt.Errorf("unable to encode user_updateLastSeen: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserUpdateLastSeen) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateLastSeen: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xfd405a2d:
		m.Id, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateLastSeen#0xfd405a2d: field id: %w", err)
		}
		m.LastSeenAt, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateLastSeen#0xfd405a2d: field last_seen_at: %w", err)
		}
		m.Expires, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateLastSeen#0xfd405a2d: field expires: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_updateLastSeen: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetLastSeen <--
type TLUserGetLastSeen struct {
	ClazzID uint32 `json:"_id"`
	Id      int64  `json:"id"`
}

func (m *TLUserGetLastSeen) String() string {
	return iface.DebugStringWithName(ClazzName_user_getLastSeen, m)
}

// Encode <--
func (m *TLUserGetLastSeen) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getLastSeen, int(layer)); clazzId {
	case 0x9119c8de:
		x.PutClazzID(0x9119c8de)

		x.PutInt64(m.Id)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getLastSeen: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetLastSeen) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getLastSeen: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x9119c8de:
		m.Id, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getLastSeen#0x9119c8de: field id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getLastSeen: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetImmutableUser <--
type TLUserGetImmutableUser struct {
	ClazzID  uint32  `json:"_id"`
	Id       int64   `json:"id"`
	Privacy  bool    `json:"privacy"`
	Contacts []int64 `json:"contacts"`
}

func (m *TLUserGetImmutableUser) String() string {
	return iface.DebugStringWithName(ClazzName_user_getImmutableUser, m)
}

// Encode <--
func (m *TLUserGetImmutableUser) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getImmutableUser, int(layer)); clazzId {
	case 0x376a6744:
		x.PutClazzID(0x376a6744)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.Privacy == true {
				flags |= 1 << 1
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.Id)

		iface.EncodeInt64List(x, m.Contacts)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getImmutableUser: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetImmutableUser) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getImmutableUser: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x376a6744:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode user_getImmutableUser: field flags: %w", err)
		}
		_ = flags
		m.Id, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getImmutableUser#0x376a6744: field id: %w", err)
		}
		if (flags & (1 << 1)) != 0 {
			m.Privacy = true
		}

		m.Contacts, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_getImmutableUser#0x376a6744: field contacts: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getImmutableUser: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetMutableUsers <--
type TLUserGetMutableUsers struct {
	ClazzID uint32  `json:"_id"`
	Id      []int64 `json:"id"`
	To      []int64 `json:"to"`
}

func (m *TLUserGetMutableUsers) String() string {
	return iface.DebugStringWithName(ClazzName_user_getMutableUsers, m)
}

// Encode <--
func (m *TLUserGetMutableUsers) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getMutableUsers, int(layer)); clazzId {
	case 0x9d3b23d7:
		x.PutClazzID(0x9d3b23d7)

		iface.EncodeInt64List(x, m.Id)

		iface.EncodeInt64List(x, m.To)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getMutableUsers: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetMutableUsers) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getMutableUsers: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x9d3b23d7:

		m.Id, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_getMutableUsers#0x9d3b23d7: field id: %w", err)
		}

		m.To, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_getMutableUsers#0x9d3b23d7: field to: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getMutableUsers: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetImmutableUserByPhone <--
type TLUserGetImmutableUserByPhone struct {
	ClazzID uint32 `json:"_id"`
	Phone   string `json:"phone"`
}

func (m *TLUserGetImmutableUserByPhone) String() string {
	return iface.DebugStringWithName(ClazzName_user_getImmutableUserByPhone, m)
}

// Encode <--
func (m *TLUserGetImmutableUserByPhone) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getImmutableUserByPhone, int(layer)); clazzId {
	case 0xe9c36fe4:
		x.PutClazzID(0xe9c36fe4)

		x.PutString(m.Phone)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getImmutableUserByPhone: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetImmutableUserByPhone) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getImmutableUserByPhone: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xe9c36fe4:
		m.Phone, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_getImmutableUserByPhone#0xe9c36fe4: field phone: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getImmutableUserByPhone: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetImmutableUserByToken <--
type TLUserGetImmutableUserByToken struct {
	ClazzID uint32 `json:"_id"`
	Token   string `json:"token"`
}

func (m *TLUserGetImmutableUserByToken) String() string {
	return iface.DebugStringWithName(ClazzName_user_getImmutableUserByToken, m)
}

// Encode <--
func (m *TLUserGetImmutableUserByToken) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getImmutableUserByToken, int(layer)); clazzId {
	case 0xff3e1373:
		x.PutClazzID(0xff3e1373)

		x.PutString(m.Token)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getImmutableUserByToken: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetImmutableUserByToken) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getImmutableUserByToken: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xff3e1373:
		m.Token, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_getImmutableUserByToken#0xff3e1373: field token: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getImmutableUserByToken: invalid constructor %x", m.ClazzID)
	}
}

// TLUserSetAccountDaysTTL <--
type TLUserSetAccountDaysTTL struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Ttl     int32  `json:"ttl"`
}

func (m *TLUserSetAccountDaysTTL) String() string {
	return iface.DebugStringWithName(ClazzName_user_setAccountDaysTTL, m)
}

// Encode <--
func (m *TLUserSetAccountDaysTTL) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_setAccountDaysTTL, int(layer)); clazzId {
	case 0xd2550b4c:
		x.PutClazzID(0xd2550b4c)

		x.PutInt64(m.UserId)
		x.PutInt32(m.Ttl)

		return nil
	default:
		return fmt.Errorf("unable to encode user_setAccountDaysTTL: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserSetAccountDaysTTL) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_setAccountDaysTTL: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xd2550b4c:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_setAccountDaysTTL#0xd2550b4c: field user_id: %w", err)
		}
		m.Ttl, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_setAccountDaysTTL#0xd2550b4c: field ttl: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_setAccountDaysTTL: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetAccountDaysTTL <--
type TLUserGetAccountDaysTTL struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLUserGetAccountDaysTTL) String() string {
	return iface.DebugStringWithName(ClazzName_user_getAccountDaysTTL, m)
}

// Encode <--
func (m *TLUserGetAccountDaysTTL) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getAccountDaysTTL, int(layer)); clazzId {
	case 0xb2843ee0:
		x.PutClazzID(0xb2843ee0)

		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getAccountDaysTTL: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetAccountDaysTTL) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getAccountDaysTTL: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xb2843ee0:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getAccountDaysTTL#0xb2843ee0: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getAccountDaysTTL: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetNotifySettings <--
type TLUserGetNotifySettings struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

func (m *TLUserGetNotifySettings) String() string {
	return iface.DebugStringWithName(ClazzName_user_getNotifySettings, m)
}

// Encode <--
func (m *TLUserGetNotifySettings) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getNotifySettings, int(layer)); clazzId {
	case 0x40ac3766:
		x.PutClazzID(0x40ac3766)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getNotifySettings: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetNotifySettings) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getNotifySettings: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x40ac3766:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getNotifySettings#0x40ac3766: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_getNotifySettings#0x40ac3766: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getNotifySettings#0x40ac3766: field peer_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getNotifySettings: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetNotifySettingsList <--
type TLUserGetNotifySettingsList struct {
	ClazzID uint32             `json:"_id"`
	UserId  int64              `json:"user_id"`
	Peers   []tg.PeerUtilClazz `json:"peers"`
}

func (m *TLUserGetNotifySettingsList) String() string {
	return iface.DebugStringWithName(ClazzName_user_getNotifySettingsList, m)
}

// Encode <--
func (m *TLUserGetNotifySettingsList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getNotifySettingsList, int(layer)); clazzId {
	case 0xe465159c:
		x.PutClazzID(0xe465159c)

		x.PutInt64(m.UserId)

		if err := iface.EncodeObjectList(x, m.Peers, layer); err != nil {
			return fmt.Errorf("unable to encode user_getNotifySettingsList#0xe465159c: field peers: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode user_getNotifySettingsList: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetNotifySettingsList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getNotifySettingsList: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xe465159c:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getNotifySettingsList#0xe465159c: field user_id: %w", err)
		}
		l2, err3 := d.VectorHeader()
		if err3 != nil {
			return fmt.Errorf("unable to decode user_getNotifySettingsList#0xe465159c: field peers: %w", err3)
		}
		if l2 > bin.MaxVectorLen {
			return fmt.Errorf("unable to decode user_getNotifySettingsList#0xe465159c: field peers: %w", &bin.InvalidLengthError{Type: "vector", Length: int(l2)})
		}
		prealloc2 := int(l2)
		if prealloc2 > bin.PreallocateLimit {
			prealloc2 = bin.PreallocateLimit
		}
		v2 := make([]tg.PeerUtilClazz, 0, prealloc2)
		for i := int32(0); i < l2; i++ {
			vv2, err3 := tg.DecodePeerUtilClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode user_getNotifySettingsList#0xe465159c: field peers: %w", err3)
			}
			v2 = append(v2, vv2)
		}
		m.Peers = v2

		return nil
	default:
		return fmt.Errorf("unable to decode user_getNotifySettingsList: invalid constructor %x", m.ClazzID)
	}
}

// TLUserSetNotifySettings <--
type TLUserSetNotifySettings struct {
	ClazzID  uint32                     `json:"_id"`
	UserId   int64                      `json:"user_id"`
	PeerType int32                      `json:"peer_type"`
	PeerId   int64                      `json:"peer_id"`
	Settings tg.PeerNotifySettingsClazz `json:"settings"`
}

func (m *TLUserSetNotifySettings) String() string {
	return iface.DebugStringWithName(ClazzName_user_setNotifySettings, m)
}

// Encode <--
func (m *TLUserSetNotifySettings) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_setNotifySettings, int(layer)); clazzId {
	case 0xc9ed65e5:
		x.PutClazzID(0xc9ed65e5)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		if m.Settings == nil {
			return fmt.Errorf("unable to encode user_setNotifySettings#0xc9ed65e5: field settings is nil")
		}
		if err := m.Settings.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode user_setNotifySettings#0xc9ed65e5: field settings: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode user_setNotifySettings: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserSetNotifySettings) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_setNotifySettings: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xc9ed65e5:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_setNotifySettings#0xc9ed65e5: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_setNotifySettings#0xc9ed65e5: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_setNotifySettings#0xc9ed65e5: field peer_id: %w", err)
		}

		m.Settings, err = tg.DecodePeerNotifySettingsClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_setNotifySettings#0xc9ed65e5: field settings: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_setNotifySettings: invalid constructor %x", m.ClazzID)
	}
}

// TLUserResetNotifySettings <--
type TLUserResetNotifySettings struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLUserResetNotifySettings) String() string {
	return iface.DebugStringWithName(ClazzName_user_resetNotifySettings, m)
}

// Encode <--
func (m *TLUserResetNotifySettings) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_resetNotifySettings, int(layer)); clazzId {
	case 0xe079d74:
		x.PutClazzID(0xe079d74)

		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_resetNotifySettings: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserResetNotifySettings) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_resetNotifySettings: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xe079d74:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_resetNotifySettings#0xe079d74: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_resetNotifySettings: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetAllNotifySettings <--
type TLUserGetAllNotifySettings struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLUserGetAllNotifySettings) String() string {
	return iface.DebugStringWithName(ClazzName_user_getAllNotifySettings, m)
}

// Encode <--
func (m *TLUserGetAllNotifySettings) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getAllNotifySettings, int(layer)); clazzId {
	case 0x55926875:
		x.PutClazzID(0x55926875)

		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getAllNotifySettings: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetAllNotifySettings) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getAllNotifySettings: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x55926875:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getAllNotifySettings#0x55926875: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getAllNotifySettings: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetGlobalPrivacySettings <--
type TLUserGetGlobalPrivacySettings struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLUserGetGlobalPrivacySettings) String() string {
	return iface.DebugStringWithName(ClazzName_user_getGlobalPrivacySettings, m)
}

// Encode <--
func (m *TLUserGetGlobalPrivacySettings) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getGlobalPrivacySettings, int(layer)); clazzId {
	case 0x77f6f112:
		x.PutClazzID(0x77f6f112)

		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getGlobalPrivacySettings: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetGlobalPrivacySettings) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getGlobalPrivacySettings: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x77f6f112:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getGlobalPrivacySettings#0x77f6f112: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getGlobalPrivacySettings: invalid constructor %x", m.ClazzID)
	}
}

// TLUserSetGlobalPrivacySettings <--
type TLUserSetGlobalPrivacySettings struct {
	ClazzID  uint32                        `json:"_id"`
	UserId   int64                         `json:"user_id"`
	Settings tg.GlobalPrivacySettingsClazz `json:"settings"`
}

func (m *TLUserSetGlobalPrivacySettings) String() string {
	return iface.DebugStringWithName(ClazzName_user_setGlobalPrivacySettings, m)
}

// Encode <--
func (m *TLUserSetGlobalPrivacySettings) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_setGlobalPrivacySettings, int(layer)); clazzId {
	case 0x8cb592ae:
		x.PutClazzID(0x8cb592ae)

		x.PutInt64(m.UserId)
		if m.Settings == nil {
			return fmt.Errorf("unable to encode user_setGlobalPrivacySettings#0x8cb592ae: field settings is nil")
		}
		if err := m.Settings.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode user_setGlobalPrivacySettings#0x8cb592ae: field settings: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode user_setGlobalPrivacySettings: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserSetGlobalPrivacySettings) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_setGlobalPrivacySettings: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x8cb592ae:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_setGlobalPrivacySettings#0x8cb592ae: field user_id: %w", err)
		}

		m.Settings, err = tg.DecodeGlobalPrivacySettingsClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_setGlobalPrivacySettings#0x8cb592ae: field settings: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_setGlobalPrivacySettings: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetPrivacy <--
type TLUserGetPrivacy struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	KeyType int32  `json:"key_type"`
}

func (m *TLUserGetPrivacy) String() string {
	return iface.DebugStringWithName(ClazzName_user_getPrivacy, m)
}

// Encode <--
func (m *TLUserGetPrivacy) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getPrivacy, int(layer)); clazzId {
	case 0x9d40a3b4:
		x.PutClazzID(0x9d40a3b4)

		x.PutInt64(m.UserId)
		x.PutInt32(m.KeyType)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getPrivacy: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetPrivacy) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getPrivacy: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x9d40a3b4:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getPrivacy#0x9d40a3b4: field user_id: %w", err)
		}
		m.KeyType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_getPrivacy#0x9d40a3b4: field key_type: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getPrivacy: invalid constructor %x", m.ClazzID)
	}
}

// TLUserSetPrivacy <--
type TLUserSetPrivacy struct {
	ClazzID uint32                `json:"_id"`
	UserId  int64                 `json:"user_id"`
	KeyType int32                 `json:"key_type"`
	Rules   []tg.PrivacyRuleClazz `json:"rules"`
}

func (m *TLUserSetPrivacy) String() string {
	return iface.DebugStringWithName(ClazzName_user_setPrivacy, m)
}

// Encode <--
func (m *TLUserSetPrivacy) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_setPrivacy, int(layer)); clazzId {
	case 0x8855ad8f:
		x.PutClazzID(0x8855ad8f)

		x.PutInt64(m.UserId)
		x.PutInt32(m.KeyType)

		if err := iface.EncodeObjectList(x, m.Rules, layer); err != nil {
			return fmt.Errorf("unable to encode user_setPrivacy#0x8855ad8f: field rules: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode user_setPrivacy: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserSetPrivacy) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_setPrivacy: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x8855ad8f:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_setPrivacy#0x8855ad8f: field user_id: %w", err)
		}
		m.KeyType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_setPrivacy#0x8855ad8f: field key_type: %w", err)
		}
		l3, err3 := d.VectorHeader()
		if err3 != nil {
			return fmt.Errorf("unable to decode user_setPrivacy#0x8855ad8f: field rules: %w", err3)
		}
		if l3 > bin.MaxVectorLen {
			return fmt.Errorf("unable to decode user_setPrivacy#0x8855ad8f: field rules: %w", &bin.InvalidLengthError{Type: "vector", Length: int(l3)})
		}
		prealloc3 := int(l3)
		if prealloc3 > bin.PreallocateLimit {
			prealloc3 = bin.PreallocateLimit
		}
		v3 := make([]tg.PrivacyRuleClazz, 0, prealloc3)
		for i := int32(0); i < l3; i++ {
			vv3, err3 := tg.DecodePrivacyRuleClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode user_setPrivacy#0x8855ad8f: field rules: %w", err3)
			}
			v3 = append(v3, vv3)
		}
		m.Rules = v3

		return nil
	default:
		return fmt.Errorf("unable to decode user_setPrivacy: invalid constructor %x", m.ClazzID)
	}
}

// TLUserCheckPrivacy <--
type TLUserCheckPrivacy struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	KeyType int32  `json:"key_type"`
	PeerId  int64  `json:"peer_id"`
}

func (m *TLUserCheckPrivacy) String() string {
	return iface.DebugStringWithName(ClazzName_user_checkPrivacy, m)
}

// Encode <--
func (m *TLUserCheckPrivacy) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_checkPrivacy, int(layer)); clazzId {
	case 0xc56e1eaa:
		x.PutClazzID(0xc56e1eaa)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.UserId)
		x.PutInt32(m.KeyType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_checkPrivacy: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserCheckPrivacy) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_checkPrivacy: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xc56e1eaa:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode user_checkPrivacy: field flags: %w", err)
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_checkPrivacy#0xc56e1eaa: field user_id: %w", err)
		}
		m.KeyType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_checkPrivacy#0xc56e1eaa: field key_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_checkPrivacy#0xc56e1eaa: field peer_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_checkPrivacy: invalid constructor %x", m.ClazzID)
	}
}

// TLUserAddPeerSettings <--
type TLUserAddPeerSettings struct {
	ClazzID  uint32               `json:"_id"`
	UserId   int64                `json:"user_id"`
	PeerType int32                `json:"peer_type"`
	PeerId   int64                `json:"peer_id"`
	Settings tg.PeerSettingsClazz `json:"settings"`
}

func (m *TLUserAddPeerSettings) String() string {
	return iface.DebugStringWithName(ClazzName_user_addPeerSettings, m)
}

// Encode <--
func (m *TLUserAddPeerSettings) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_addPeerSettings, int(layer)); clazzId {
	case 0xcae22763:
		x.PutClazzID(0xcae22763)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		if m.Settings == nil {
			return fmt.Errorf("unable to encode user_addPeerSettings#0xcae22763: field settings is nil")
		}
		if err := m.Settings.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode user_addPeerSettings#0xcae22763: field settings: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode user_addPeerSettings: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserAddPeerSettings) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_addPeerSettings: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xcae22763:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_addPeerSettings#0xcae22763: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_addPeerSettings#0xcae22763: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_addPeerSettings#0xcae22763: field peer_id: %w", err)
		}

		m.Settings, err = tg.DecodePeerSettingsClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_addPeerSettings#0xcae22763: field settings: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_addPeerSettings: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetPeerSettings <--
type TLUserGetPeerSettings struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

func (m *TLUserGetPeerSettings) String() string {
	return iface.DebugStringWithName(ClazzName_user_getPeerSettings, m)
}

// Encode <--
func (m *TLUserGetPeerSettings) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getPeerSettings, int(layer)); clazzId {
	case 0xd02ef67:
		x.PutClazzID(0xd02ef67)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getPeerSettings: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetPeerSettings) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getPeerSettings: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xd02ef67:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getPeerSettings#0xd02ef67: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_getPeerSettings#0xd02ef67: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getPeerSettings#0xd02ef67: field peer_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getPeerSettings: invalid constructor %x", m.ClazzID)
	}
}

// TLUserDeletePeerSettings <--
type TLUserDeletePeerSettings struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

func (m *TLUserDeletePeerSettings) String() string {
	return iface.DebugStringWithName(ClazzName_user_deletePeerSettings, m)
}

// Encode <--
func (m *TLUserDeletePeerSettings) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_deletePeerSettings, int(layer)); clazzId {
	case 0x5e891967:
		x.PutClazzID(0x5e891967)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_deletePeerSettings: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserDeletePeerSettings) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_deletePeerSettings: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x5e891967:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_deletePeerSettings#0x5e891967: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_deletePeerSettings#0x5e891967: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_deletePeerSettings#0x5e891967: field peer_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_deletePeerSettings: invalid constructor %x", m.ClazzID)
	}
}

// TLUserChangePhone <--
type TLUserChangePhone struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Phone   string `json:"phone"`
}

func (m *TLUserChangePhone) String() string {
	return iface.DebugStringWithName(ClazzName_user_changePhone, m)
}

// Encode <--
func (m *TLUserChangePhone) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_changePhone, int(layer)); clazzId {
	case 0xff7a575b:
		x.PutClazzID(0xff7a575b)

		x.PutInt64(m.UserId)
		x.PutString(m.Phone)

		return nil
	default:
		return fmt.Errorf("unable to encode user_changePhone: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserChangePhone) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_changePhone: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xff7a575b:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_changePhone#0xff7a575b: field user_id: %w", err)
		}
		m.Phone, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_changePhone#0xff7a575b: field phone: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_changePhone: invalid constructor %x", m.ClazzID)
	}
}

// TLUserCreateNewUser <--
type TLUserCreateNewUser struct {
	ClazzID     uint32 `json:"_id"`
	SecretKeyId int64  `json:"secret_key_id"`
	Phone       string `json:"phone"`
	CountryCode string `json:"country_code"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

func (m *TLUserCreateNewUser) String() string {
	return iface.DebugStringWithName(ClazzName_user_createNewUser, m)
}

// Encode <--
func (m *TLUserCreateNewUser) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_createNewUser, int(layer)); clazzId {
	case 0x79e01881:
		x.PutClazzID(0x79e01881)

		x.PutInt64(m.SecretKeyId)
		x.PutString(m.Phone)
		x.PutString(m.CountryCode)
		x.PutString(m.FirstName)
		x.PutString(m.LastName)

		return nil
	default:
		return fmt.Errorf("unable to encode user_createNewUser: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserCreateNewUser) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_createNewUser: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x79e01881:
		m.SecretKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_createNewUser#0x79e01881: field secret_key_id: %w", err)
		}
		m.Phone, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_createNewUser#0x79e01881: field phone: %w", err)
		}
		m.CountryCode, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_createNewUser#0x79e01881: field country_code: %w", err)
		}
		m.FirstName, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_createNewUser#0x79e01881: field first_name: %w", err)
		}
		m.LastName, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_createNewUser#0x79e01881: field last_name: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_createNewUser: invalid constructor %x", m.ClazzID)
	}
}

// TLUserDeleteUser <--
type TLUserDeleteUser struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Reason  string `json:"reason"`
	Phone   string `json:"phone"`
}

func (m *TLUserDeleteUser) String() string {
	return iface.DebugStringWithName(ClazzName_user_deleteUser, m)
}

// Encode <--
func (m *TLUserDeleteUser) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_deleteUser, int(layer)); clazzId {
	case 0x626dbd10:
		x.PutClazzID(0x626dbd10)

		x.PutInt64(m.UserId)
		x.PutString(m.Reason)
		x.PutString(m.Phone)

		return nil
	default:
		return fmt.Errorf("unable to encode user_deleteUser: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserDeleteUser) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_deleteUser: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x626dbd10:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_deleteUser#0x626dbd10: field user_id: %w", err)
		}
		m.Reason, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_deleteUser#0x626dbd10: field reason: %w", err)
		}
		m.Phone, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_deleteUser#0x626dbd10: field phone: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_deleteUser: invalid constructor %x", m.ClazzID)
	}
}

// TLUserBlockPeer <--
type TLUserBlockPeer struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

func (m *TLUserBlockPeer) String() string {
	return iface.DebugStringWithName(ClazzName_user_blockPeer, m)
}

// Encode <--
func (m *TLUserBlockPeer) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_blockPeer, int(layer)); clazzId {
	case 0x81062eb0:
		x.PutClazzID(0x81062eb0)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_blockPeer: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserBlockPeer) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_blockPeer: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x81062eb0:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_blockPeer#0x81062eb0: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_blockPeer#0x81062eb0: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_blockPeer#0x81062eb0: field peer_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_blockPeer: invalid constructor %x", m.ClazzID)
	}
}

// TLUserUnBlockPeer <--
type TLUserUnBlockPeer struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

func (m *TLUserUnBlockPeer) String() string {
	return iface.DebugStringWithName(ClazzName_user_unBlockPeer, m)
}

// Encode <--
func (m *TLUserUnBlockPeer) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_unBlockPeer, int(layer)); clazzId {
	case 0xdee7160d:
		x.PutClazzID(0xdee7160d)

		x.PutInt64(m.UserId)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_unBlockPeer: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserUnBlockPeer) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_unBlockPeer: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xdee7160d:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_unBlockPeer#0xdee7160d: field user_id: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_unBlockPeer#0xdee7160d: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_unBlockPeer#0xdee7160d: field peer_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_unBlockPeer: invalid constructor %x", m.ClazzID)
	}
}

// TLUserBlockedByUser <--
type TLUserBlockedByUser struct {
	ClazzID    uint32 `json:"_id"`
	UserId     int64  `json:"user_id"`
	PeerUserId int64  `json:"peer_user_id"`
}

func (m *TLUserBlockedByUser) String() string {
	return iface.DebugStringWithName(ClazzName_user_blockedByUser, m)
}

// Encode <--
func (m *TLUserBlockedByUser) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_blockedByUser, int(layer)); clazzId {
	case 0xbba0058e:
		x.PutClazzID(0xbba0058e)

		x.PutInt64(m.UserId)
		x.PutInt64(m.PeerUserId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_blockedByUser: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserBlockedByUser) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_blockedByUser: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xbba0058e:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_blockedByUser#0xbba0058e: field user_id: %w", err)
		}
		m.PeerUserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_blockedByUser#0xbba0058e: field peer_user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_blockedByUser: invalid constructor %x", m.ClazzID)
	}
}

// TLUserIsBlockedByUser <--
type TLUserIsBlockedByUser struct {
	ClazzID    uint32 `json:"_id"`
	UserId     int64  `json:"user_id"`
	PeerUserId int64  `json:"peer_user_id"`
}

func (m *TLUserIsBlockedByUser) String() string {
	return iface.DebugStringWithName(ClazzName_user_isBlockedByUser, m)
}

// Encode <--
func (m *TLUserIsBlockedByUser) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_isBlockedByUser, int(layer)); clazzId {
	case 0x8caeb1df:
		x.PutClazzID(0x8caeb1df)

		x.PutInt64(m.UserId)
		x.PutInt64(m.PeerUserId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_isBlockedByUser: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserIsBlockedByUser) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_isBlockedByUser: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x8caeb1df:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_isBlockedByUser#0x8caeb1df: field user_id: %w", err)
		}
		m.PeerUserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_isBlockedByUser#0x8caeb1df: field peer_user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_isBlockedByUser: invalid constructor %x", m.ClazzID)
	}
}

// TLUserCheckBlockUserList <--
type TLUserCheckBlockUserList struct {
	ClazzID uint32  `json:"_id"`
	UserId  int64   `json:"user_id"`
	Id      []int64 `json:"id"`
}

func (m *TLUserCheckBlockUserList) String() string {
	return iface.DebugStringWithName(ClazzName_user_checkBlockUserList, m)
}

// Encode <--
func (m *TLUserCheckBlockUserList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_checkBlockUserList, int(layer)); clazzId {
	case 0xc3fd70f0:
		x.PutClazzID(0xc3fd70f0)

		x.PutInt64(m.UserId)

		iface.EncodeInt64List(x, m.Id)

		return nil
	default:
		return fmt.Errorf("unable to encode user_checkBlockUserList: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserCheckBlockUserList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_checkBlockUserList: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xc3fd70f0:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_checkBlockUserList#0xc3fd70f0: field user_id: %w", err)
		}

		m.Id, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_checkBlockUserList#0xc3fd70f0: field id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_checkBlockUserList: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetBlockedList <--
type TLUserGetBlockedList struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Offset  int32  `json:"offset"`
	Limit   int32  `json:"limit"`
}

func (m *TLUserGetBlockedList) String() string {
	return iface.DebugStringWithName(ClazzName_user_getBlockedList, m)
}

// Encode <--
func (m *TLUserGetBlockedList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getBlockedList, int(layer)); clazzId {
	case 0x23ffc348:
		x.PutClazzID(0x23ffc348)

		x.PutInt64(m.UserId)
		x.PutInt32(m.Offset)
		x.PutInt32(m.Limit)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getBlockedList: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetBlockedList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getBlockedList: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x23ffc348:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getBlockedList#0x23ffc348: field user_id: %w", err)
		}
		m.Offset, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_getBlockedList#0x23ffc348: field offset: %w", err)
		}
		m.Limit, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_getBlockedList#0x23ffc348: field limit: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getBlockedList: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetContactSignUpNotification <--
type TLUserGetContactSignUpNotification struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLUserGetContactSignUpNotification) String() string {
	return iface.DebugStringWithName(ClazzName_user_getContactSignUpNotification, m)
}

// Encode <--
func (m *TLUserGetContactSignUpNotification) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getContactSignUpNotification, int(layer)); clazzId {
	case 0xe4d1d3d6:
		x.PutClazzID(0xe4d1d3d6)

		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getContactSignUpNotification: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetContactSignUpNotification) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getContactSignUpNotification: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xe4d1d3d6:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getContactSignUpNotification#0xe4d1d3d6: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getContactSignUpNotification: invalid constructor %x", m.ClazzID)
	}
}

// TLUserSetContactSignUpNotification <--
type TLUserSetContactSignUpNotification struct {
	ClazzID uint32       `json:"_id"`
	UserId  int64        `json:"user_id"`
	Silent  tg.BoolClazz `json:"silent"`
}

func (m *TLUserSetContactSignUpNotification) String() string {
	return iface.DebugStringWithName(ClazzName_user_setContactSignUpNotification, m)
}

// Encode <--
func (m *TLUserSetContactSignUpNotification) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_setContactSignUpNotification, int(layer)); clazzId {
	case 0x85a17361:
		x.PutClazzID(0x85a17361)

		x.PutInt64(m.UserId)
		if m.Silent == nil {
			return fmt.Errorf("unable to encode user_setContactSignUpNotification#0x85a17361: field silent is nil")
		}
		if err := m.Silent.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode user_setContactSignUpNotification#0x85a17361: field silent: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode user_setContactSignUpNotification: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserSetContactSignUpNotification) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_setContactSignUpNotification: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x85a17361:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_setContactSignUpNotification#0x85a17361: field user_id: %w", err)
		}

		m.Silent, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_setContactSignUpNotification#0x85a17361: field silent: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_setContactSignUpNotification: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetContentSettings <--
type TLUserGetContentSettings struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLUserGetContentSettings) String() string {
	return iface.DebugStringWithName(ClazzName_user_getContentSettings, m)
}

// Encode <--
func (m *TLUserGetContentSettings) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getContentSettings, int(layer)); clazzId {
	case 0x94c3ad9f:
		x.PutClazzID(0x94c3ad9f)

		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getContentSettings: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetContentSettings) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getContentSettings: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x94c3ad9f:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getContentSettings#0x94c3ad9f: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getContentSettings: invalid constructor %x", m.ClazzID)
	}
}

// TLUserSetContentSettings <--
type TLUserSetContentSettings struct {
	ClazzID          uint32 `json:"_id"`
	UserId           int64  `json:"user_id"`
	SensitiveEnabled bool   `json:"sensitive_enabled"`
}

func (m *TLUserSetContentSettings) String() string {
	return iface.DebugStringWithName(ClazzName_user_setContentSettings, m)
}

// Encode <--
func (m *TLUserSetContentSettings) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_setContentSettings, int(layer)); clazzId {
	case 0x9d63fe6b:
		x.PutClazzID(0x9d63fe6b)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.SensitiveEnabled == true {
				flags |= 1 << 0
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_setContentSettings: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserSetContentSettings) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_setContentSettings: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x9d63fe6b:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode user_setContentSettings: field flags: %w", err)
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_setContentSettings#0x9d63fe6b: field user_id: %w", err)
		}
		if (flags & (1 << 0)) != 0 {
			m.SensitiveEnabled = true
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_setContentSettings: invalid constructor %x", m.ClazzID)
	}
}

// TLUserDeleteContact <--
type TLUserDeleteContact struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Id      int64  `json:"id"`
}

func (m *TLUserDeleteContact) String() string {
	return iface.DebugStringWithName(ClazzName_user_deleteContact, m)
}

// Encode <--
func (m *TLUserDeleteContact) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_deleteContact, int(layer)); clazzId {
	case 0xc6018219:
		x.PutClazzID(0xc6018219)

		x.PutInt64(m.UserId)
		x.PutInt64(m.Id)

		return nil
	default:
		return fmt.Errorf("unable to encode user_deleteContact: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserDeleteContact) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_deleteContact: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xc6018219:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_deleteContact#0xc6018219: field user_id: %w", err)
		}
		m.Id, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_deleteContact#0xc6018219: field id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_deleteContact: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetContactList <--
type TLUserGetContactList struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLUserGetContactList) String() string {
	return iface.DebugStringWithName(ClazzName_user_getContactList, m)
}

// Encode <--
func (m *TLUserGetContactList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getContactList, int(layer)); clazzId {
	case 0xc74bd161:
		x.PutClazzID(0xc74bd161)

		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getContactList: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetContactList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getContactList: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xc74bd161:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getContactList#0xc74bd161: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getContactList: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetContactIdList <--
type TLUserGetContactIdList struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLUserGetContactIdList) String() string {
	return iface.DebugStringWithName(ClazzName_user_getContactIdList, m)
}

// Encode <--
func (m *TLUserGetContactIdList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getContactIdList, int(layer)); clazzId {
	case 0xf1dd983e:
		x.PutClazzID(0xf1dd983e)

		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getContactIdList: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetContactIdList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getContactIdList: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xf1dd983e:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getContactIdList#0xf1dd983e: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getContactIdList: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetContact <--
type TLUserGetContact struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Id      int64  `json:"id"`
}

func (m *TLUserGetContact) String() string {
	return iface.DebugStringWithName(ClazzName_user_getContact, m)
}

// Encode <--
func (m *TLUserGetContact) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getContact, int(layer)); clazzId {
	case 0xdb728be3:
		x.PutClazzID(0xdb728be3)

		x.PutInt64(m.UserId)
		x.PutInt64(m.Id)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getContact: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetContact) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getContact: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xdb728be3:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getContact#0xdb728be3: field user_id: %w", err)
		}
		m.Id, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getContact#0xdb728be3: field id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getContact: invalid constructor %x", m.ClazzID)
	}
}

// TLUserAddContact <--
type TLUserAddContact struct {
	ClazzID                  uint32       `json:"_id"`
	UserId                   int64        `json:"user_id"`
	AddPhonePrivacyException tg.BoolClazz `json:"add_phone_privacy_exception"`
	Id                       int64        `json:"id"`
	FirstName                string       `json:"first_name"`
	LastName                 string       `json:"last_name"`
	Phone                    string       `json:"phone"`
}

func (m *TLUserAddContact) String() string {
	return iface.DebugStringWithName(ClazzName_user_addContact, m)
}

// Encode <--
func (m *TLUserAddContact) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_addContact, int(layer)); clazzId {
	case 0x79c4bd0e:
		x.PutClazzID(0x79c4bd0e)

		x.PutInt64(m.UserId)
		if m.AddPhonePrivacyException == nil {
			return fmt.Errorf("unable to encode user_addContact#0x79c4bd0e: field add_phone_privacy_exception is nil")
		}
		if err := m.AddPhonePrivacyException.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode user_addContact#0x79c4bd0e: field add_phone_privacy_exception: %w", err)
		}
		x.PutInt64(m.Id)
		x.PutString(m.FirstName)
		x.PutString(m.LastName)
		x.PutString(m.Phone)

		return nil
	default:
		return fmt.Errorf("unable to encode user_addContact: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserAddContact) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_addContact: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x79c4bd0e:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_addContact#0x79c4bd0e: field user_id: %w", err)
		}

		m.AddPhonePrivacyException, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_addContact#0x79c4bd0e: field add_phone_privacy_exception: %w", err)
		}

		m.Id, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_addContact#0x79c4bd0e: field id: %w", err)
		}
		m.FirstName, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_addContact#0x79c4bd0e: field first_name: %w", err)
		}
		m.LastName, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_addContact#0x79c4bd0e: field last_name: %w", err)
		}
		m.Phone, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_addContact#0x79c4bd0e: field phone: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_addContact: invalid constructor %x", m.ClazzID)
	}
}

// TLUserCheckContact <--
type TLUserCheckContact struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Id      int64  `json:"id"`
}

func (m *TLUserCheckContact) String() string {
	return iface.DebugStringWithName(ClazzName_user_checkContact, m)
}

// Encode <--
func (m *TLUserCheckContact) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_checkContact, int(layer)); clazzId {
	case 0x82a758a4:
		x.PutClazzID(0x82a758a4)

		x.PutInt64(m.UserId)
		x.PutInt64(m.Id)

		return nil
	default:
		return fmt.Errorf("unable to encode user_checkContact: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserCheckContact) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_checkContact: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x82a758a4:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_checkContact#0x82a758a4: field user_id: %w", err)
		}
		m.Id, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_checkContact#0x82a758a4: field id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_checkContact: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetImportersByPhone <--
type TLUserGetImportersByPhone struct {
	ClazzID uint32 `json:"_id"`
	Phone   string `json:"phone"`
}

func (m *TLUserGetImportersByPhone) String() string {
	return iface.DebugStringWithName(ClazzName_user_getImportersByPhone, m)
}

// Encode <--
func (m *TLUserGetImportersByPhone) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getImportersByPhone, int(layer)); clazzId {
	case 0x47aa8212:
		x.PutClazzID(0x47aa8212)

		x.PutString(m.Phone)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getImportersByPhone: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetImportersByPhone) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getImportersByPhone: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x47aa8212:
		m.Phone, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_getImportersByPhone#0x47aa8212: field phone: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getImportersByPhone: invalid constructor %x", m.ClazzID)
	}
}

// TLUserDeleteImportersByPhone <--
type TLUserDeleteImportersByPhone struct {
	ClazzID uint32 `json:"_id"`
	Phone   string `json:"phone"`
}

func (m *TLUserDeleteImportersByPhone) String() string {
	return iface.DebugStringWithName(ClazzName_user_deleteImportersByPhone, m)
}

// Encode <--
func (m *TLUserDeleteImportersByPhone) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_deleteImportersByPhone, int(layer)); clazzId {
	case 0x174ddc54:
		x.PutClazzID(0x174ddc54)

		x.PutString(m.Phone)

		return nil
	default:
		return fmt.Errorf("unable to encode user_deleteImportersByPhone: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserDeleteImportersByPhone) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_deleteImportersByPhone: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x174ddc54:
		m.Phone, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_deleteImportersByPhone#0x174ddc54: field phone: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_deleteImportersByPhone: invalid constructor %x", m.ClazzID)
	}
}

// TLUserImportContacts <--
type TLUserImportContacts struct {
	ClazzID  uint32                 `json:"_id"`
	UserId   int64                  `json:"user_id"`
	Contacts []tg.InputContactClazz `json:"contacts"`
}

func (m *TLUserImportContacts) String() string {
	return iface.DebugStringWithName(ClazzName_user_importContacts, m)
}

// Encode <--
func (m *TLUserImportContacts) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_importContacts, int(layer)); clazzId {
	case 0x9a00f792:
		x.PutClazzID(0x9a00f792)

		x.PutInt64(m.UserId)

		if err := iface.EncodeObjectList(x, m.Contacts, layer); err != nil {
			return fmt.Errorf("unable to encode user_importContacts#0x9a00f792: field contacts: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode user_importContacts: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserImportContacts) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_importContacts: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x9a00f792:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_importContacts#0x9a00f792: field user_id: %w", err)
		}
		l2, err3 := d.VectorHeader()
		if err3 != nil {
			return fmt.Errorf("unable to decode user_importContacts#0x9a00f792: field contacts: %w", err3)
		}
		if l2 > bin.MaxVectorLen {
			return fmt.Errorf("unable to decode user_importContacts#0x9a00f792: field contacts: %w", &bin.InvalidLengthError{Type: "vector", Length: int(l2)})
		}
		prealloc2 := int(l2)
		if prealloc2 > bin.PreallocateLimit {
			prealloc2 = bin.PreallocateLimit
		}
		v2 := make([]tg.InputContactClazz, 0, prealloc2)
		for i := int32(0); i < l2; i++ {
			vv2, err3 := tg.DecodeInputContactClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode user_importContacts#0x9a00f792: field contacts: %w", err3)
			}
			v2 = append(v2, vv2)
		}
		m.Contacts = v2

		return nil
	default:
		return fmt.Errorf("unable to decode user_importContacts: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetCountryCode <--
type TLUserGetCountryCode struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLUserGetCountryCode) String() string {
	return iface.DebugStringWithName(ClazzName_user_getCountryCode, m)
}

// Encode <--
func (m *TLUserGetCountryCode) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getCountryCode, int(layer)); clazzId {
	case 0x12006832:
		x.PutClazzID(0x12006832)

		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getCountryCode: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetCountryCode) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getCountryCode: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x12006832:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getCountryCode#0x12006832: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getCountryCode: invalid constructor %x", m.ClazzID)
	}
}

// TLUserUpdateAbout <--
type TLUserUpdateAbout struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	About   string `json:"about"`
}

func (m *TLUserUpdateAbout) String() string {
	return iface.DebugStringWithName(ClazzName_user_updateAbout, m)
}

// Encode <--
func (m *TLUserUpdateAbout) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_updateAbout, int(layer)); clazzId {
	case 0xdb00f187:
		x.PutClazzID(0xdb00f187)

		x.PutInt64(m.UserId)
		x.PutString(m.About)

		return nil
	default:
		return fmt.Errorf("unable to encode user_updateAbout: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserUpdateAbout) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateAbout: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xdb00f187:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateAbout#0xdb00f187: field user_id: %w", err)
		}
		m.About, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateAbout#0xdb00f187: field about: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_updateAbout: invalid constructor %x", m.ClazzID)
	}
}

// TLUserUpdateFirstAndLastName <--
type TLUserUpdateFirstAndLastName struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (m *TLUserUpdateFirstAndLastName) String() string {
	return iface.DebugStringWithName(ClazzName_user_updateFirstAndLastName, m)
}

// Encode <--
func (m *TLUserUpdateFirstAndLastName) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_updateFirstAndLastName, int(layer)); clazzId {
	case 0xcb6685ec:
		x.PutClazzID(0xcb6685ec)

		x.PutInt64(m.UserId)
		x.PutString(m.FirstName)
		x.PutString(m.LastName)

		return nil
	default:
		return fmt.Errorf("unable to encode user_updateFirstAndLastName: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserUpdateFirstAndLastName) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateFirstAndLastName: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xcb6685ec:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateFirstAndLastName#0xcb6685ec: field user_id: %w", err)
		}
		m.FirstName, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateFirstAndLastName#0xcb6685ec: field first_name: %w", err)
		}
		m.LastName, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateFirstAndLastName#0xcb6685ec: field last_name: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_updateFirstAndLastName: invalid constructor %x", m.ClazzID)
	}
}

// TLUserUpdateVerified <--
type TLUserUpdateVerified struct {
	ClazzID  uint32       `json:"_id"`
	UserId   int64        `json:"user_id"`
	Verified tg.BoolClazz `json:"verified"`
}

func (m *TLUserUpdateVerified) String() string {
	return iface.DebugStringWithName(ClazzName_user_updateVerified, m)
}

// Encode <--
func (m *TLUserUpdateVerified) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_updateVerified, int(layer)); clazzId {
	case 0x24c92963:
		x.PutClazzID(0x24c92963)

		x.PutInt64(m.UserId)
		if m.Verified == nil {
			return fmt.Errorf("unable to encode user_updateVerified#0x24c92963: field verified is nil")
		}
		if err := m.Verified.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode user_updateVerified#0x24c92963: field verified: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode user_updateVerified: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserUpdateVerified) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateVerified: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x24c92963:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateVerified#0x24c92963: field user_id: %w", err)
		}

		m.Verified, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_updateVerified#0x24c92963: field verified: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_updateVerified: invalid constructor %x", m.ClazzID)
	}
}

// TLUserUpdateUsername <--
type TLUserUpdateUsername struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
}

func (m *TLUserUpdateUsername) String() string {
	return iface.DebugStringWithName(ClazzName_user_updateUsername, m)
}

// Encode <--
func (m *TLUserUpdateUsername) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_updateUsername, int(layer)); clazzId {
	case 0xf54d1e71:
		x.PutClazzID(0xf54d1e71)

		x.PutInt64(m.UserId)
		x.PutString(m.Username)

		return nil
	default:
		return fmt.Errorf("unable to encode user_updateUsername: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserUpdateUsername) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateUsername: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xf54d1e71:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateUsername#0xf54d1e71: field user_id: %w", err)
		}
		m.Username, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateUsername#0xf54d1e71: field username: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_updateUsername: invalid constructor %x", m.ClazzID)
	}
}

// TLUserUpdateProfilePhoto <--
type TLUserUpdateProfilePhoto struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Id      int64  `json:"id"`
}

func (m *TLUserUpdateProfilePhoto) String() string {
	return iface.DebugStringWithName(ClazzName_user_updateProfilePhoto, m)
}

// Encode <--
func (m *TLUserUpdateProfilePhoto) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_updateProfilePhoto, int(layer)); clazzId {
	case 0x3b740f87:
		x.PutClazzID(0x3b740f87)

		x.PutInt64(m.UserId)
		x.PutInt64(m.Id)

		return nil
	default:
		return fmt.Errorf("unable to encode user_updateProfilePhoto: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserUpdateProfilePhoto) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateProfilePhoto: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x3b740f87:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateProfilePhoto#0x3b740f87: field user_id: %w", err)
		}
		m.Id, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateProfilePhoto#0x3b740f87: field id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_updateProfilePhoto: invalid constructor %x", m.ClazzID)
	}
}

// TLUserDeleteProfilePhotos <--
type TLUserDeleteProfilePhotos struct {
	ClazzID uint32  `json:"_id"`
	UserId  int64   `json:"user_id"`
	Id      []int64 `json:"id"`
}

func (m *TLUserDeleteProfilePhotos) String() string {
	return iface.DebugStringWithName(ClazzName_user_deleteProfilePhotos, m)
}

// Encode <--
func (m *TLUserDeleteProfilePhotos) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_deleteProfilePhotos, int(layer)); clazzId {
	case 0x2be3620e:
		x.PutClazzID(0x2be3620e)

		x.PutInt64(m.UserId)

		iface.EncodeInt64List(x, m.Id)

		return nil
	default:
		return fmt.Errorf("unable to encode user_deleteProfilePhotos: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserDeleteProfilePhotos) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_deleteProfilePhotos: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x2be3620e:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_deleteProfilePhotos#0x2be3620e: field user_id: %w", err)
		}

		m.Id, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_deleteProfilePhotos#0x2be3620e: field id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_deleteProfilePhotos: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetProfilePhotos <--
type TLUserGetProfilePhotos struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLUserGetProfilePhotos) String() string {
	return iface.DebugStringWithName(ClazzName_user_getProfilePhotos, m)
}

// Encode <--
func (m *TLUserGetProfilePhotos) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getProfilePhotos, int(layer)); clazzId {
	case 0xdc66c146:
		x.PutClazzID(0xdc66c146)

		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getProfilePhotos: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetProfilePhotos) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getProfilePhotos: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xdc66c146:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getProfilePhotos#0xdc66c146: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getProfilePhotos: invalid constructor %x", m.ClazzID)
	}
}

// TLUserSetBotCommands <--
type TLUserSetBotCommands struct {
	ClazzID  uint32               `json:"_id"`
	UserId   int64                `json:"user_id"`
	BotId    int64                `json:"bot_id"`
	Commands []tg.BotCommandClazz `json:"commands"`
}

func (m *TLUserSetBotCommands) String() string {
	return iface.DebugStringWithName(ClazzName_user_setBotCommands, m)
}

// Encode <--
func (m *TLUserSetBotCommands) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_setBotCommands, int(layer)); clazzId {
	case 0x753ba916:
		x.PutClazzID(0x753ba916)

		x.PutInt64(m.UserId)
		x.PutInt64(m.BotId)

		if err := iface.EncodeObjectList(x, m.Commands, layer); err != nil {
			return fmt.Errorf("unable to encode user_setBotCommands#0x753ba916: field commands: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode user_setBotCommands: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserSetBotCommands) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_setBotCommands: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x753ba916:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_setBotCommands#0x753ba916: field user_id: %w", err)
		}
		m.BotId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_setBotCommands#0x753ba916: field bot_id: %w", err)
		}
		l3, err3 := d.VectorHeader()
		if err3 != nil {
			return fmt.Errorf("unable to decode user_setBotCommands#0x753ba916: field commands: %w", err3)
		}
		if l3 > bin.MaxVectorLen {
			return fmt.Errorf("unable to decode user_setBotCommands#0x753ba916: field commands: %w", &bin.InvalidLengthError{Type: "vector", Length: int(l3)})
		}
		prealloc3 := int(l3)
		if prealloc3 > bin.PreallocateLimit {
			prealloc3 = bin.PreallocateLimit
		}
		v3 := make([]tg.BotCommandClazz, 0, prealloc3)
		for i := int32(0); i < l3; i++ {
			vv3, err3 := tg.DecodeBotCommandClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode user_setBotCommands#0x753ba916: field commands: %w", err3)
			}
			v3 = append(v3, vv3)
		}
		m.Commands = v3

		return nil
	default:
		return fmt.Errorf("unable to decode user_setBotCommands: invalid constructor %x", m.ClazzID)
	}
}

// TLUserIsBot <--
type TLUserIsBot struct {
	ClazzID uint32 `json:"_id"`
	Id      int64  `json:"id"`
}

func (m *TLUserIsBot) String() string {
	return iface.DebugStringWithName(ClazzName_user_isBot, m)
}

// Encode <--
func (m *TLUserIsBot) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_isBot, int(layer)); clazzId {
	case 0xc772c7ee:
		x.PutClazzID(0xc772c7ee)

		x.PutInt64(m.Id)

		return nil
	default:
		return fmt.Errorf("unable to encode user_isBot: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserIsBot) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_isBot: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xc772c7ee:
		m.Id, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_isBot#0xc772c7ee: field id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_isBot: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetBotInfo <--
type TLUserGetBotInfo struct {
	ClazzID uint32 `json:"_id"`
	BotId   int64  `json:"bot_id"`
}

func (m *TLUserGetBotInfo) String() string {
	return iface.DebugStringWithName(ClazzName_user_getBotInfo, m)
}

// Encode <--
func (m *TLUserGetBotInfo) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getBotInfo, int(layer)); clazzId {
	case 0x34663710:
		x.PutClazzID(0x34663710)

		x.PutInt64(m.BotId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getBotInfo: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetBotInfo) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getBotInfo: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x34663710:
		m.BotId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getBotInfo#0x34663710: field bot_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getBotInfo: invalid constructor %x", m.ClazzID)
	}
}

// TLUserCheckBots <--
type TLUserCheckBots struct {
	ClazzID uint32  `json:"_id"`
	Id      []int64 `json:"id"`
}

func (m *TLUserCheckBots) String() string {
	return iface.DebugStringWithName(ClazzName_user_checkBots, m)
}

// Encode <--
func (m *TLUserCheckBots) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_checkBots, int(layer)); clazzId {
	case 0x736500c1:
		x.PutClazzID(0x736500c1)

		iface.EncodeInt64List(x, m.Id)

		return nil
	default:
		return fmt.Errorf("unable to encode user_checkBots: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserCheckBots) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_checkBots: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x736500c1:

		m.Id, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_checkBots#0x736500c1: field id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_checkBots: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetFullUser <--
type TLUserGetFullUser struct {
	ClazzID    uint32 `json:"_id"`
	SelfUserId int64  `json:"self_user_id"`
	Id         int64  `json:"id"`
}

func (m *TLUserGetFullUser) String() string {
	return iface.DebugStringWithName(ClazzName_user_getFullUser, m)
}

// Encode <--
func (m *TLUserGetFullUser) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getFullUser, int(layer)); clazzId {
	case 0xfd10e13a:
		x.PutClazzID(0xfd10e13a)

		x.PutInt64(m.SelfUserId)
		x.PutInt64(m.Id)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getFullUser: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetFullUser) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getFullUser: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xfd10e13a:
		m.SelfUserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getFullUser#0xfd10e13a: field self_user_id: %w", err)
		}
		m.Id, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getFullUser#0xfd10e13a: field id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getFullUser: invalid constructor %x", m.ClazzID)
	}
}

// TLUserUpdateEmojiStatus <--
type TLUserUpdateEmojiStatus struct {
	ClazzID               uint32 `json:"_id"`
	UserId                int64  `json:"user_id"`
	EmojiStatusDocumentId int64  `json:"emoji_status_document_id"`
	EmojiStatusUntil      int32  `json:"emoji_status_until"`
}

func (m *TLUserUpdateEmojiStatus) String() string {
	return iface.DebugStringWithName(ClazzName_user_updateEmojiStatus, m)
}

// Encode <--
func (m *TLUserUpdateEmojiStatus) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_updateEmojiStatus, int(layer)); clazzId {
	case 0xf8c8bad8:
		x.PutClazzID(0xf8c8bad8)

		x.PutInt64(m.UserId)
		x.PutInt64(m.EmojiStatusDocumentId)
		x.PutInt32(m.EmojiStatusUntil)

		return nil
	default:
		return fmt.Errorf("unable to encode user_updateEmojiStatus: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserUpdateEmojiStatus) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateEmojiStatus: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xf8c8bad8:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateEmojiStatus#0xf8c8bad8: field user_id: %w", err)
		}
		m.EmojiStatusDocumentId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateEmojiStatus#0xf8c8bad8: field emoji_status_document_id: %w", err)
		}
		m.EmojiStatusUntil, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateEmojiStatus#0xf8c8bad8: field emoji_status_until: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_updateEmojiStatus: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetUserDataById <--
type TLUserGetUserDataById struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLUserGetUserDataById) String() string {
	return iface.DebugStringWithName(ClazzName_user_getUserDataById, m)
}

// Encode <--
func (m *TLUserGetUserDataById) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getUserDataById, int(layer)); clazzId {
	case 0x3bb7103:
		x.PutClazzID(0x3bb7103)

		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getUserDataById: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetUserDataById) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getUserDataById: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x3bb7103:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getUserDataById#0x3bb7103: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getUserDataById: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetUserDataListByIdList <--
type TLUserGetUserDataListByIdList struct {
	ClazzID    uint32  `json:"_id"`
	UserIdList []int64 `json:"user_id_list"`
}

func (m *TLUserGetUserDataListByIdList) String() string {
	return iface.DebugStringWithName(ClazzName_user_getUserDataListByIdList, m)
}

// Encode <--
func (m *TLUserGetUserDataListByIdList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getUserDataListByIdList, int(layer)); clazzId {
	case 0x8191eff9:
		x.PutClazzID(0x8191eff9)

		iface.EncodeInt64List(x, m.UserIdList)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getUserDataListByIdList: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetUserDataListByIdList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getUserDataListByIdList: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x8191eff9:

		m.UserIdList, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_getUserDataListByIdList#0x8191eff9: field user_id_list: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getUserDataListByIdList: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetUserDataByToken <--
type TLUserGetUserDataByToken struct {
	ClazzID uint32 `json:"_id"`
	Token   string `json:"token"`
}

func (m *TLUserGetUserDataByToken) String() string {
	return iface.DebugStringWithName(ClazzName_user_getUserDataByToken, m)
}

// Encode <--
func (m *TLUserGetUserDataByToken) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getUserDataByToken, int(layer)); clazzId {
	case 0x3f09659e:
		x.PutClazzID(0x3f09659e)

		x.PutString(m.Token)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getUserDataByToken: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetUserDataByToken) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getUserDataByToken: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x3f09659e:
		m.Token, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_getUserDataByToken#0x3f09659e: field token: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getUserDataByToken: invalid constructor %x", m.ClazzID)
	}
}

// TLUserSearch <--
type TLUserSearch struct {
	ClazzID          uint32  `json:"_id"`
	Q                string  `json:"q"`
	ExcludedContacts []int64 `json:"excluded_contacts"`
	Offset           int64   `json:"offset"`
	Limit            int32   `json:"limit"`
}

func (m *TLUserSearch) String() string {
	return iface.DebugStringWithName(ClazzName_user_search, m)
}

// Encode <--
func (m *TLUserSearch) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_search, int(layer)); clazzId {
	case 0x7035b6cd:
		x.PutClazzID(0x7035b6cd)

		x.PutString(m.Q)

		iface.EncodeInt64List(x, m.ExcludedContacts)

		x.PutInt64(m.Offset)
		x.PutInt32(m.Limit)

		return nil
	default:
		return fmt.Errorf("unable to encode user_search: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserSearch) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_search: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x7035b6cd:
		m.Q, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_search#0x7035b6cd: field q: %w", err)
		}

		m.ExcludedContacts, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_search#0x7035b6cd: field excluded_contacts: %w", err)
		}

		m.Offset, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_search#0x7035b6cd: field offset: %w", err)
		}
		m.Limit, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_search#0x7035b6cd: field limit: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_search: invalid constructor %x", m.ClazzID)
	}
}

// TLUserUpdateBotData <--
type TLUserUpdateBotData struct {
	ClazzID              uint32       `json:"_id"`
	BotId                int64        `json:"bot_id"`
	BotChatHistory       tg.BoolClazz `json:"bot_chat_history"`
	BotNochats           tg.BoolClazz `json:"bot_nochats"`
	BotInlineGeo         tg.BoolClazz `json:"bot_inline_geo"`
	BotAttachMenu        tg.BoolClazz `json:"bot_attach_menu"`
	BotInlinePlaceholder *string      `json:"bot_inline_placeholder"`
	BotHasMainApp        tg.BoolClazz `json:"bot_has_main_app"`
}

func (m *TLUserUpdateBotData) String() string {
	return iface.DebugStringWithName(ClazzName_user_updateBotData, m)
}

// Encode <--
func (m *TLUserUpdateBotData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_updateBotData, int(layer)); clazzId {
	case 0x60f35d28:
		x.PutClazzID(0x60f35d28)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.BotChatHistory != nil {
				flags |= 1 << 15
			}
			if m.BotNochats != nil {
				flags |= 1 << 16
			}
			if m.BotInlineGeo != nil {
				flags |= 1 << 21
			}
			if m.BotAttachMenu != nil {
				flags |= 1 << 27
			}
			if m.BotInlinePlaceholder != nil {
				flags |= 1 << 19
			}
			if m.BotHasMainApp != nil {
				flags |= 1 << 13
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.BotId)
		if m.BotChatHistory != nil {
			if err := m.BotChatHistory.Encode(x, layer); err != nil {
				return fmt.Errorf("unable to encode user_updateBotData#0x60f35d28: field bot_chat_history: %w", err)
			}
		}

		if m.BotNochats != nil {
			if err := m.BotNochats.Encode(x, layer); err != nil {
				return fmt.Errorf("unable to encode user_updateBotData#0x60f35d28: field bot_nochats: %w", err)
			}
		}

		if m.BotInlineGeo != nil {
			if err := m.BotInlineGeo.Encode(x, layer); err != nil {
				return fmt.Errorf("unable to encode user_updateBotData#0x60f35d28: field bot_inline_geo: %w", err)
			}
		}

		if m.BotAttachMenu != nil {
			if err := m.BotAttachMenu.Encode(x, layer); err != nil {
				return fmt.Errorf("unable to encode user_updateBotData#0x60f35d28: field bot_attach_menu: %w", err)
			}
		}

		if m.BotInlinePlaceholder != nil {
			x.PutString(*m.BotInlinePlaceholder)
		}

		if m.BotHasMainApp != nil {
			if err := m.BotHasMainApp.Encode(x, layer); err != nil {
				return fmt.Errorf("unable to encode user_updateBotData#0x60f35d28: field bot_has_main_app: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to encode user_updateBotData: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserUpdateBotData) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateBotData: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x60f35d28:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateBotData: field flags: %w", err)
		}
		_ = flags
		m.BotId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateBotData#0x60f35d28: field bot_id: %w", err)
		}
		if (flags & (1 << 15)) != 0 {
			m.BotChatHistory, err = tg.DecodeBoolClazz(d)
			if err != nil {
				return fmt.Errorf("unable to decode user_updateBotData#0x60f35d28: field bot_chat_history: %w", err)
			}
		}
		if (flags & (1 << 16)) != 0 {
			m.BotNochats, err = tg.DecodeBoolClazz(d)
			if err != nil {
				return fmt.Errorf("unable to decode user_updateBotData#0x60f35d28: field bot_nochats: %w", err)
			}
		}
		if (flags & (1 << 21)) != 0 {
			m.BotInlineGeo, err = tg.DecodeBoolClazz(d)
			if err != nil {
				return fmt.Errorf("unable to decode user_updateBotData#0x60f35d28: field bot_inline_geo: %w", err)
			}
		}
		if (flags & (1 << 27)) != 0 {
			m.BotAttachMenu, err = tg.DecodeBoolClazz(d)
			if err != nil {
				return fmt.Errorf("unable to decode user_updateBotData#0x60f35d28: field bot_attach_menu: %w", err)
			}
		}
		if (flags & (1 << 19)) != 0 {
			m.BotInlinePlaceholder = new(string)
			*m.BotInlinePlaceholder, err = d.String()
			if err != nil {
				return fmt.Errorf("unable to decode user_updateBotData#0x60f35d28: field bot_inline_placeholder: %w", err)
			}
		}

		if (flags & (1 << 13)) != 0 {
			m.BotHasMainApp, err = tg.DecodeBoolClazz(d)
			if err != nil {
				return fmt.Errorf("unable to decode user_updateBotData#0x60f35d28: field bot_has_main_app: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_updateBotData: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetImmutableUserV2 <--
type TLUserGetImmutableUserV2 struct {
	ClazzID uint32  `json:"_id"`
	Id      int64   `json:"id"`
	Privacy bool    `json:"privacy"`
	HasTo   bool    `json:"has_to"`
	To      []int64 `json:"to"`
}

func (m *TLUserGetImmutableUserV2) String() string {
	return iface.DebugStringWithName(ClazzName_user_getImmutableUserV2, m)
}

// Encode <--
func (m *TLUserGetImmutableUserV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getImmutableUserV2, int(layer)); clazzId {
	case 0x300aba4c:
		x.PutClazzID(0x300aba4c)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.Privacy == true {
				flags |= 1 << 0
			}
			if m.HasTo == true {
				flags |= 1 << 2
			}
			if m.To != nil {
				flags |= 1 << 2
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.Id)
		if m.To != nil {
			iface.EncodeInt64List(x, m.To)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode user_getImmutableUserV2: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetImmutableUserV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getImmutableUserV2: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x300aba4c:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode user_getImmutableUserV2: field flags: %w", err)
		}
		_ = flags
		m.Id, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getImmutableUserV2#0x300aba4c: field id: %w", err)
		}
		if (flags & (1 << 0)) != 0 {
			m.Privacy = true
		}
		if (flags & (1 << 2)) != 0 {
			m.HasTo = true
		}
		if (flags & (1 << 2)) != 0 {
			m.To, err = iface.DecodeInt64List(d)
			if err != nil {
				return fmt.Errorf("unable to decode user_getImmutableUserV2#0x300aba4c: field to: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getImmutableUserV2: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetMutableUsersV2 <--
type TLUserGetMutableUsersV2 struct {
	ClazzID uint32  `json:"_id"`
	Id      []int64 `json:"id"`
	Privacy bool    `json:"privacy"`
	HasTo   bool    `json:"has_to"`
	To      []int64 `json:"to"`
}

func (m *TLUserGetMutableUsersV2) String() string {
	return iface.DebugStringWithName(ClazzName_user_getMutableUsersV2, m)
}

// Encode <--
func (m *TLUserGetMutableUsersV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getMutableUsersV2, int(layer)); clazzId {
	case 0x94f98b28:
		x.PutClazzID(0x94f98b28)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.Privacy == true {
				flags |= 1 << 0
			}
			if m.HasTo == true {
				flags |= 1 << 2
			}
			if m.To != nil {
				flags |= 1 << 2
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)

		iface.EncodeInt64List(x, m.Id)

		if m.To != nil {
			iface.EncodeInt64List(x, m.To)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode user_getMutableUsersV2: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetMutableUsersV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getMutableUsersV2: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x94f98b28:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode user_getMutableUsersV2: field flags: %w", err)
		}
		_ = flags

		m.Id, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_getMutableUsersV2#0x94f98b28: field id: %w", err)
		}

		if (flags & (1 << 0)) != 0 {
			m.Privacy = true
		}
		if (flags & (1 << 2)) != 0 {
			m.HasTo = true
		}
		if (flags & (1 << 2)) != 0 {
			m.To, err = iface.DecodeInt64List(d)
			if err != nil {
				return fmt.Errorf("unable to decode user_getMutableUsersV2#0x94f98b28: field to: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getMutableUsersV2: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetUserProjectionBundle <--
type TLUserGetUserProjectionBundle struct {
	ClazzID       uint32  `json:"_id"`
	WithFacts     bool    `json:"with_facts"`
	ViewerUserIds []int64 `json:"viewer_user_ids"`
	TargetUserIds []int64 `json:"target_user_ids"`
}

func (m *TLUserGetUserProjectionBundle) String() string {
	return iface.DebugStringWithName(ClazzName_user_getUserProjectionBundle, m)
}

// Encode <--
func (m *TLUserGetUserProjectionBundle) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getUserProjectionBundle, int(layer)); clazzId {
	case 0x3fc25f21:
		x.PutClazzID(0x3fc25f21)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.WithFacts == true {
				flags |= 1 << 0
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)

		iface.EncodeInt64List(x, m.ViewerUserIds)

		iface.EncodeInt64List(x, m.TargetUserIds)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getUserProjectionBundle: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetUserProjectionBundle) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getUserProjectionBundle: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x3fc25f21:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode user_getUserProjectionBundle: field flags: %w", err)
		}
		_ = flags
		if (flags & (1 << 0)) != 0 {
			m.WithFacts = true
		}

		m.ViewerUserIds, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_getUserProjectionBundle#0x3fc25f21: field viewer_user_ids: %w", err)
		}

		m.TargetUserIds, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_getUserProjectionBundle#0x3fc25f21: field target_user_ids: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getUserProjectionBundle: invalid constructor %x", m.ClazzID)
	}
}

// TLUserCreateNewTestUser <--
type TLUserCreateNewTestUser struct {
	ClazzID     uint32 `json:"_id"`
	SecretKeyId int64  `json:"secret_key_id"`
	MinId       int64  `json:"min_id"`
	MaxId       int64  `json:"max_id"`
}

func (m *TLUserCreateNewTestUser) String() string {
	return iface.DebugStringWithName(ClazzName_user_createNewTestUser, m)
}

// Encode <--
func (m *TLUserCreateNewTestUser) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_createNewTestUser, int(layer)); clazzId {
	case 0x4c6eccab:
		x.PutClazzID(0x4c6eccab)

		x.PutInt64(m.SecretKeyId)
		x.PutInt64(m.MinId)
		x.PutInt64(m.MaxId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_createNewTestUser: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserCreateNewTestUser) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_createNewTestUser: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x4c6eccab:
		m.SecretKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_createNewTestUser#0x4c6eccab: field secret_key_id: %w", err)
		}
		m.MinId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_createNewTestUser#0x4c6eccab: field min_id: %w", err)
		}
		m.MaxId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_createNewTestUser#0x4c6eccab: field max_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_createNewTestUser: invalid constructor %x", m.ClazzID)
	}
}

// TLUserEditCloseFriends <--
type TLUserEditCloseFriends struct {
	ClazzID uint32  `json:"_id"`
	UserId  int64   `json:"user_id"`
	Id      []int64 `json:"id"`
}

func (m *TLUserEditCloseFriends) String() string {
	return iface.DebugStringWithName(ClazzName_user_editCloseFriends, m)
}

// Encode <--
func (m *TLUserEditCloseFriends) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_editCloseFriends, int(layer)); clazzId {
	case 0x86247b05:
		x.PutClazzID(0x86247b05)

		x.PutInt64(m.UserId)

		iface.EncodeInt64List(x, m.Id)

		return nil
	default:
		return fmt.Errorf("unable to encode user_editCloseFriends: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserEditCloseFriends) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_editCloseFriends: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x86247b05:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_editCloseFriends#0x86247b05: field user_id: %w", err)
		}

		m.Id, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_editCloseFriends#0x86247b05: field id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_editCloseFriends: invalid constructor %x", m.ClazzID)
	}
}

// TLUserSetStoriesMaxId <--
type TLUserSetStoriesMaxId struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Id      int32  `json:"id"`
}

func (m *TLUserSetStoriesMaxId) String() string {
	return iface.DebugStringWithName(ClazzName_user_setStoriesMaxId, m)
}

// Encode <--
func (m *TLUserSetStoriesMaxId) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_setStoriesMaxId, int(layer)); clazzId {
	case 0x52f5b670:
		x.PutClazzID(0x52f5b670)

		x.PutInt64(m.UserId)
		x.PutInt32(m.Id)

		return nil
	default:
		return fmt.Errorf("unable to encode user_setStoriesMaxId: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserSetStoriesMaxId) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_setStoriesMaxId: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x52f5b670:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_setStoriesMaxId#0x52f5b670: field user_id: %w", err)
		}
		m.Id, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_setStoriesMaxId#0x52f5b670: field id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_setStoriesMaxId: invalid constructor %x", m.ClazzID)
	}
}

// TLUserSetColor <--
type TLUserSetColor struct {
	ClazzID           uint32 `json:"_id"`
	UserId            int64  `json:"user_id"`
	ForProfile        bool   `json:"for_profile"`
	Color             int32  `json:"color"`
	BackgroundEmojiId int64  `json:"background_emoji_id"`
}

func (m *TLUserSetColor) String() string {
	return iface.DebugStringWithName(ClazzName_user_setColor, m)
}

// Encode <--
func (m *TLUserSetColor) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_setColor, int(layer)); clazzId {
	case 0x22fa0d77:
		x.PutClazzID(0x22fa0d77)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.ForProfile == true {
				flags |= 1 << 1
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.UserId)
		x.PutInt32(m.Color)
		x.PutInt64(m.BackgroundEmojiId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_setColor: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserSetColor) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_setColor: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x22fa0d77:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode user_setColor: field flags: %w", err)
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_setColor#0x22fa0d77: field user_id: %w", err)
		}
		if (flags & (1 << 1)) != 0 {
			m.ForProfile = true
		}
		m.Color, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_setColor#0x22fa0d77: field color: %w", err)
		}
		m.BackgroundEmojiId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_setColor#0x22fa0d77: field background_emoji_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_setColor: invalid constructor %x", m.ClazzID)
	}
}

// TLUserUpdateBirthday <--
type TLUserUpdateBirthday struct {
	ClazzID  uint32           `json:"_id"`
	UserId   int64            `json:"user_id"`
	Birthday tg.BirthdayClazz `json:"birthday"`
}

func (m *TLUserUpdateBirthday) String() string {
	return iface.DebugStringWithName(ClazzName_user_updateBirthday, m)
}

// Encode <--
func (m *TLUserUpdateBirthday) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_updateBirthday, int(layer)); clazzId {
	case 0x587aab92:
		x.PutClazzID(0x587aab92)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.Birthday != nil {
				flags |= 1 << 1
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.UserId)
		if m.Birthday != nil {
			if err := m.Birthday.Encode(x, layer); err != nil {
				return fmt.Errorf("unable to encode user_updateBirthday#0x587aab92: field birthday: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to encode user_updateBirthday: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserUpdateBirthday) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateBirthday: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x587aab92:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateBirthday: field flags: %w", err)
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateBirthday#0x587aab92: field user_id: %w", err)
		}
		if (flags & (1 << 1)) != 0 {
			m.Birthday, err = tg.DecodeBirthdayClazz(d)
			if err != nil {
				return fmt.Errorf("unable to decode user_updateBirthday#0x587aab92: field birthday: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_updateBirthday: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetBirthdays <--
type TLUserGetBirthdays struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLUserGetBirthdays) String() string {
	return iface.DebugStringWithName(ClazzName_user_getBirthdays, m)
}

// Encode <--
func (m *TLUserGetBirthdays) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getBirthdays, int(layer)); clazzId {
	case 0xfe8ebfa6:
		x.PutClazzID(0xfe8ebfa6)

		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getBirthdays: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetBirthdays) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getBirthdays: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xfe8ebfa6:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getBirthdays#0xfe8ebfa6: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getBirthdays: invalid constructor %x", m.ClazzID)
	}
}

// TLUserSetStoriesHidden <--
type TLUserSetStoriesHidden struct {
	ClazzID uint32       `json:"_id"`
	UserId  int64        `json:"user_id"`
	Id      int64        `json:"id"`
	Hidden  tg.BoolClazz `json:"hidden"`
}

func (m *TLUserSetStoriesHidden) String() string {
	return iface.DebugStringWithName(ClazzName_user_setStoriesHidden, m)
}

// Encode <--
func (m *TLUserSetStoriesHidden) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_setStoriesHidden, int(layer)); clazzId {
	case 0xf7c61858:
		x.PutClazzID(0xf7c61858)

		x.PutInt64(m.UserId)
		x.PutInt64(m.Id)
		if m.Hidden == nil {
			return fmt.Errorf("unable to encode user_setStoriesHidden#0xf7c61858: field hidden is nil")
		}
		if err := m.Hidden.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode user_setStoriesHidden#0xf7c61858: field hidden: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode user_setStoriesHidden: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserSetStoriesHidden) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_setStoriesHidden: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xf7c61858:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_setStoriesHidden#0xf7c61858: field user_id: %w", err)
		}
		m.Id, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_setStoriesHidden#0xf7c61858: field id: %w", err)
		}

		m.Hidden, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_setStoriesHidden#0xf7c61858: field hidden: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_setStoriesHidden: invalid constructor %x", m.ClazzID)
	}
}

// TLUserUpdatePersonalChannel <--
type TLUserUpdatePersonalChannel struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	ChannelId int64  `json:"channel_id"`
}

func (m *TLUserUpdatePersonalChannel) String() string {
	return iface.DebugStringWithName(ClazzName_user_updatePersonalChannel, m)
}

// Encode <--
func (m *TLUserUpdatePersonalChannel) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_updatePersonalChannel, int(layer)); clazzId {
	case 0xc7f7bed0:
		x.PutClazzID(0xc7f7bed0)

		x.PutInt64(m.UserId)
		x.PutInt64(m.ChannelId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_updatePersonalChannel: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserUpdatePersonalChannel) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_updatePersonalChannel: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xc7f7bed0:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_updatePersonalChannel#0xc7f7bed0: field user_id: %w", err)
		}
		m.ChannelId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_updatePersonalChannel#0xc7f7bed0: field channel_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_updatePersonalChannel: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetUserIdByPhone <--
type TLUserGetUserIdByPhone struct {
	ClazzID uint32 `json:"_id"`
	Phone   string `json:"phone"`
}

func (m *TLUserGetUserIdByPhone) String() string {
	return iface.DebugStringWithName(ClazzName_user_getUserIdByPhone, m)
}

// Encode <--
func (m *TLUserGetUserIdByPhone) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getUserIdByPhone, int(layer)); clazzId {
	case 0xfbab83c2:
		x.PutClazzID(0xfbab83c2)

		x.PutString(m.Phone)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getUserIdByPhone: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetUserIdByPhone) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getUserIdByPhone: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xfbab83c2:
		m.Phone, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_getUserIdByPhone#0xfbab83c2: field phone: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getUserIdByPhone: invalid constructor %x", m.ClazzID)
	}
}

// TLUserSetAuthorizationTTL <--
type TLUserSetAuthorizationTTL struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Ttl     int32  `json:"ttl"`
}

func (m *TLUserSetAuthorizationTTL) String() string {
	return iface.DebugStringWithName(ClazzName_user_setAuthorizationTTL, m)
}

// Encode <--
func (m *TLUserSetAuthorizationTTL) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_setAuthorizationTTL, int(layer)); clazzId {
	case 0xd621f3f0:
		x.PutClazzID(0xd621f3f0)

		x.PutInt64(m.UserId)
		x.PutInt32(m.Ttl)

		return nil
	default:
		return fmt.Errorf("unable to encode user_setAuthorizationTTL: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserSetAuthorizationTTL) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_setAuthorizationTTL: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xd621f3f0:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_setAuthorizationTTL#0xd621f3f0: field user_id: %w", err)
		}
		m.Ttl, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_setAuthorizationTTL#0xd621f3f0: field ttl: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_setAuthorizationTTL: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetAuthorizationTTL <--
type TLUserGetAuthorizationTTL struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLUserGetAuthorizationTTL) String() string {
	return iface.DebugStringWithName(ClazzName_user_getAuthorizationTTL, m)
}

// Encode <--
func (m *TLUserGetAuthorizationTTL) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getAuthorizationTTL, int(layer)); clazzId {
	case 0xde6e493c:
		x.PutClazzID(0xde6e493c)

		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getAuthorizationTTL: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetAuthorizationTTL) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getAuthorizationTTL: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xde6e493c:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getAuthorizationTTL#0xde6e493c: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getAuthorizationTTL: invalid constructor %x", m.ClazzID)
	}
}

// TLUserUpdatePremium <--
type TLUserUpdatePremium struct {
	ClazzID uint32       `json:"_id"`
	UserId  int64        `json:"user_id"`
	Premium tg.BoolClazz `json:"premium"`
	Months  *int32       `json:"months"`
}

func (m *TLUserUpdatePremium) String() string {
	return iface.DebugStringWithName(ClazzName_user_updatePremium, m)
}

// Encode <--
func (m *TLUserUpdatePremium) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_updatePremium, int(layer)); clazzId {
	case 0xba08dc99:
		x.PutClazzID(0xba08dc99)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.Months != nil {
				flags |= 1 << 1
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.UserId)
		if m.Premium == nil {
			return fmt.Errorf("unable to encode user_updatePremium#0xba08dc99: field premium is nil")
		}
		if err := m.Premium.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode user_updatePremium#0xba08dc99: field premium: %w", err)
		}
		if m.Months != nil {
			x.PutInt32(*m.Months)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode user_updatePremium: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserUpdatePremium) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_updatePremium: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xba08dc99:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode user_updatePremium: field flags: %w", err)
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_updatePremium#0xba08dc99: field user_id: %w", err)
		}

		m.Premium, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_updatePremium#0xba08dc99: field premium: %w", err)
		}

		if (flags & (1 << 1)) != 0 {
			m.Months = new(int32)
			*m.Months, err = d.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode user_updatePremium#0xba08dc99: field months: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_updatePremium: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetBotInfoV2 <--
type TLUserGetBotInfoV2 struct {
	ClazzID uint32 `json:"_id"`
	BotId   int64  `json:"bot_id"`
}

func (m *TLUserGetBotInfoV2) String() string {
	return iface.DebugStringWithName(ClazzName_user_getBotInfoV2, m)
}

// Encode <--
func (m *TLUserGetBotInfoV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getBotInfoV2, int(layer)); clazzId {
	case 0xd3fc9ca5:
		x.PutClazzID(0xd3fc9ca5)

		x.PutInt64(m.BotId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getBotInfoV2: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetBotInfoV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getBotInfoV2: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xd3fc9ca5:
		m.BotId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getBotInfoV2#0xd3fc9ca5: field bot_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getBotInfoV2: invalid constructor %x", m.ClazzID)
	}
}

// TLUserSaveMusic <--
type TLUserSaveMusic struct {
	ClazzID uint32 `json:"_id"`
	Unsave  bool   `json:"unsave"`
	UserId  int64  `json:"user_id"`
	Id      int64  `json:"id"`
	AfterId *int64 `json:"after_id"`
}

func (m *TLUserSaveMusic) String() string {
	return iface.DebugStringWithName(ClazzName_user_saveMusic, m)
}

// Encode <--
func (m *TLUserSaveMusic) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_saveMusic, int(layer)); clazzId {
	case 0xda28349:
		x.PutClazzID(0xda28349)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.Unsave == true {
				flags |= 1 << 0
			}

			if m.AfterId != nil {
				flags |= 1 << 15
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.UserId)
		x.PutInt64(m.Id)
		if m.AfterId != nil {
			x.PutInt64(*m.AfterId)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode user_saveMusic: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserSaveMusic) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_saveMusic: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xda28349:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode user_saveMusic: field flags: %w", err)
		}
		_ = flags
		if (flags & (1 << 0)) != 0 {
			m.Unsave = true
		}
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_saveMusic#0xda28349: field user_id: %w", err)
		}
		m.Id, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_saveMusic#0xda28349: field id: %w", err)
		}
		if (flags & (1 << 15)) != 0 {
			m.AfterId = new(int64)
			*m.AfterId, err = d.Int64()
			if err != nil {
				return fmt.Errorf("unable to decode user_saveMusic#0xda28349: field after_id: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_saveMusic: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetSavedMusicIdList <--
type TLUserGetSavedMusicIdList struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLUserGetSavedMusicIdList) String() string {
	return iface.DebugStringWithName(ClazzName_user_getSavedMusicIdList, m)
}

// Encode <--
func (m *TLUserGetSavedMusicIdList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getSavedMusicIdList, int(layer)); clazzId {
	case 0x5b4ac25f:
		x.PutClazzID(0x5b4ac25f)

		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getSavedMusicIdList: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetSavedMusicIdList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getSavedMusicIdList: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x5b4ac25f:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getSavedMusicIdList#0x5b4ac25f: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getSavedMusicIdList: invalid constructor %x", m.ClazzID)
	}
}

// TLUserSetMainProfileTab <--
type TLUserSetMainProfileTab struct {
	ClazzID uint32             `json:"_id"`
	UserId  int64              `json:"user_id"`
	Tab     tg.ProfileTabClazz `json:"tab"`
}

func (m *TLUserSetMainProfileTab) String() string {
	return iface.DebugStringWithName(ClazzName_user_setMainProfileTab, m)
}

// Encode <--
func (m *TLUserSetMainProfileTab) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_setMainProfileTab, int(layer)); clazzId {
	case 0x9d48a89c:
		x.PutClazzID(0x9d48a89c)

		x.PutInt64(m.UserId)
		if m.Tab == nil {
			return fmt.Errorf("unable to encode user_setMainProfileTab#0x9d48a89c: field tab is nil")
		}
		if err := m.Tab.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode user_setMainProfileTab#0x9d48a89c: field tab: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode user_setMainProfileTab: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserSetMainProfileTab) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_setMainProfileTab: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x9d48a89c:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_setMainProfileTab#0x9d48a89c: field user_id: %w", err)
		}

		m.Tab, err = tg.DecodeProfileTabClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_setMainProfileTab#0x9d48a89c: field tab: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_setMainProfileTab: invalid constructor %x", m.ClazzID)
	}
}

// TLUserSetDefaultHistoryTTL <--
type TLUserSetDefaultHistoryTTL struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Ttl     int32  `json:"ttl"`
}

func (m *TLUserSetDefaultHistoryTTL) String() string {
	return iface.DebugStringWithName(ClazzName_user_setDefaultHistoryTTL, m)
}

// Encode <--
func (m *TLUserSetDefaultHistoryTTL) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_setDefaultHistoryTTL, int(layer)); clazzId {
	case 0x8f09517f:
		x.PutClazzID(0x8f09517f)

		x.PutInt64(m.UserId)
		x.PutInt32(m.Ttl)

		return nil
	default:
		return fmt.Errorf("unable to encode user_setDefaultHistoryTTL: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserSetDefaultHistoryTTL) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_setDefaultHistoryTTL: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x8f09517f:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_setDefaultHistoryTTL#0x8f09517f: field user_id: %w", err)
		}
		m.Ttl, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_setDefaultHistoryTTL#0x8f09517f: field ttl: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_setDefaultHistoryTTL: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetDefaultHistoryTTL <--
type TLUserGetDefaultHistoryTTL struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLUserGetDefaultHistoryTTL) String() string {
	return iface.DebugStringWithName(ClazzName_user_getDefaultHistoryTTL, m)
}

// Encode <--
func (m *TLUserGetDefaultHistoryTTL) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getDefaultHistoryTTL, int(layer)); clazzId {
	case 0x4d4c2fe0:
		x.PutClazzID(0x4d4c2fe0)

		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getDefaultHistoryTTL: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetDefaultHistoryTTL) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getDefaultHistoryTTL: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x4d4c2fe0:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getDefaultHistoryTTL#0x4d4c2fe0: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getDefaultHistoryTTL: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetAccountUsername <--
type TLUserGetAccountUsername struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLUserGetAccountUsername) String() string {
	return iface.DebugStringWithName(ClazzName_user_getAccountUsername, m)
}

// Encode <--
func (m *TLUserGetAccountUsername) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getAccountUsername, int(layer)); clazzId {
	case 0xff8b61cf:
		x.PutClazzID(0xff8b61cf)

		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getAccountUsername: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetAccountUsername) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getAccountUsername: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xff8b61cf:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getAccountUsername#0xff8b61cf: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getAccountUsername: invalid constructor %x", m.ClazzID)
	}
}

// TLUserCheckAccountUsername <--
type TLUserCheckAccountUsername struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
}

func (m *TLUserCheckAccountUsername) String() string {
	return iface.DebugStringWithName(ClazzName_user_checkAccountUsername, m)
}

// Encode <--
func (m *TLUserCheckAccountUsername) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_checkAccountUsername, int(layer)); clazzId {
	case 0xefe05198:
		x.PutClazzID(0xefe05198)

		x.PutInt64(m.UserId)
		x.PutString(m.Username)

		return nil
	default:
		return fmt.Errorf("unable to encode user_checkAccountUsername: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserCheckAccountUsername) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_checkAccountUsername: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xefe05198:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_checkAccountUsername#0xefe05198: field user_id: %w", err)
		}
		m.Username, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_checkAccountUsername#0xefe05198: field username: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_checkAccountUsername: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetChannelUsername <--
type TLUserGetChannelUsername struct {
	ClazzID   uint32 `json:"_id"`
	ChannelId int64  `json:"channel_id"`
}

func (m *TLUserGetChannelUsername) String() string {
	return iface.DebugStringWithName(ClazzName_user_getChannelUsername, m)
}

// Encode <--
func (m *TLUserGetChannelUsername) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getChannelUsername, int(layer)); clazzId {
	case 0x910ac7b1:
		x.PutClazzID(0x910ac7b1)

		x.PutInt64(m.ChannelId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getChannelUsername: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetChannelUsername) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getChannelUsername: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x910ac7b1:
		m.ChannelId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_getChannelUsername#0x910ac7b1: field channel_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getChannelUsername: invalid constructor %x", m.ClazzID)
	}
}

// TLUserCheckChannelUsername <--
type TLUserCheckChannelUsername struct {
	ClazzID   uint32 `json:"_id"`
	ChannelId int64  `json:"channel_id"`
	Username  string `json:"username"`
}

func (m *TLUserCheckChannelUsername) String() string {
	return iface.DebugStringWithName(ClazzName_user_checkChannelUsername, m)
}

// Encode <--
func (m *TLUserCheckChannelUsername) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_checkChannelUsername, int(layer)); clazzId {
	case 0xfecbacb6:
		x.PutClazzID(0xfecbacb6)

		x.PutInt64(m.ChannelId)
		x.PutString(m.Username)

		return nil
	default:
		return fmt.Errorf("unable to encode user_checkChannelUsername: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserCheckChannelUsername) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_checkChannelUsername: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xfecbacb6:
		m.ChannelId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_checkChannelUsername#0xfecbacb6: field channel_id: %w", err)
		}
		m.Username, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_checkChannelUsername#0xfecbacb6: field username: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_checkChannelUsername: invalid constructor %x", m.ClazzID)
	}
}

// TLUserUpdateUsernameByPeer <--
type TLUserUpdateUsernameByPeer struct {
	ClazzID  uint32 `json:"_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
	Username string `json:"username"`
}

func (m *TLUserUpdateUsernameByPeer) String() string {
	return iface.DebugStringWithName(ClazzName_user_updateUsernameByPeer, m)
}

// Encode <--
func (m *TLUserUpdateUsernameByPeer) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_updateUsernameByPeer, int(layer)); clazzId {
	case 0xe3a0e9e2:
		x.PutClazzID(0xe3a0e9e2)

		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutString(m.Username)

		return nil
	default:
		return fmt.Errorf("unable to encode user_updateUsernameByPeer: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserUpdateUsernameByPeer) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateUsernameByPeer: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xe3a0e9e2:
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateUsernameByPeer#0xe3a0e9e2: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateUsernameByPeer#0xe3a0e9e2: field peer_id: %w", err)
		}
		m.Username, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateUsernameByPeer#0xe3a0e9e2: field username: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_updateUsernameByPeer: invalid constructor %x", m.ClazzID)
	}
}

// TLUserCheckUsername <--
type TLUserCheckUsername struct {
	ClazzID  uint32 `json:"_id"`
	Username string `json:"username"`
}

func (m *TLUserCheckUsername) String() string {
	return iface.DebugStringWithName(ClazzName_user_checkUsername, m)
}

// Encode <--
func (m *TLUserCheckUsername) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_checkUsername, int(layer)); clazzId {
	case 0x3475e700:
		x.PutClazzID(0x3475e700)

		x.PutString(m.Username)

		return nil
	default:
		return fmt.Errorf("unable to encode user_checkUsername: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserCheckUsername) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_checkUsername: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x3475e700:
		m.Username, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_checkUsername#0x3475e700: field username: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_checkUsername: invalid constructor %x", m.ClazzID)
	}
}

// TLUserUpdateUsernameByUsername <--
type TLUserUpdateUsernameByUsername struct {
	ClazzID  uint32 `json:"_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
	Username string `json:"username"`
}

func (m *TLUserUpdateUsernameByUsername) String() string {
	return iface.DebugStringWithName(ClazzName_user_updateUsernameByUsername, m)
}

// Encode <--
func (m *TLUserUpdateUsernameByUsername) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_updateUsernameByUsername, int(layer)); clazzId {
	case 0x13841a86:
		x.PutClazzID(0x13841a86)

		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutString(m.Username)

		return nil
	default:
		return fmt.Errorf("unable to encode user_updateUsernameByUsername: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserUpdateUsernameByUsername) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateUsernameByUsername: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x13841a86:
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateUsernameByUsername#0x13841a86: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateUsernameByUsername#0x13841a86: field peer_id: %w", err)
		}
		m.Username, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_updateUsernameByUsername#0x13841a86: field username: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_updateUsernameByUsername: invalid constructor %x", m.ClazzID)
	}
}

// TLUserDeleteUsername <--
type TLUserDeleteUsername struct {
	ClazzID  uint32 `json:"_id"`
	Username string `json:"username"`
}

func (m *TLUserDeleteUsername) String() string {
	return iface.DebugStringWithName(ClazzName_user_deleteUsername, m)
}

// Encode <--
func (m *TLUserDeleteUsername) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_deleteUsername, int(layer)); clazzId {
	case 0x85e677a4:
		x.PutClazzID(0x85e677a4)

		x.PutString(m.Username)

		return nil
	default:
		return fmt.Errorf("unable to encode user_deleteUsername: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserDeleteUsername) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_deleteUsername: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x85e677a4:
		m.Username, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_deleteUsername#0x85e677a4: field username: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_deleteUsername: invalid constructor %x", m.ClazzID)
	}
}

// TLUserResolveUsername <--
type TLUserResolveUsername struct {
	ClazzID  uint32 `json:"_id"`
	Username string `json:"username"`
}

func (m *TLUserResolveUsername) String() string {
	return iface.DebugStringWithName(ClazzName_user_resolveUsername, m)
}

// Encode <--
func (m *TLUserResolveUsername) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_resolveUsername, int(layer)); clazzId {
	case 0x4527d121:
		x.PutClazzID(0x4527d121)

		x.PutString(m.Username)

		return nil
	default:
		return fmt.Errorf("unable to encode user_resolveUsername: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserResolveUsername) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_resolveUsername: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x4527d121:
		m.Username, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_resolveUsername#0x4527d121: field username: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_resolveUsername: invalid constructor %x", m.ClazzID)
	}
}

// TLUserGetListByUsernameList <--
type TLUserGetListByUsernameList struct {
	ClazzID uint32   `json:"_id"`
	Names   []string `json:"names"`
}

func (m *TLUserGetListByUsernameList) String() string {
	return iface.DebugStringWithName(ClazzName_user_getListByUsernameList, m)
}

// Encode <--
func (m *TLUserGetListByUsernameList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_getListByUsernameList, int(layer)); clazzId {
	case 0x6e606b62:
		x.PutClazzID(0x6e606b62)

		iface.EncodeStringList(x, m.Names)

		return nil
	default:
		return fmt.Errorf("unable to encode user_getListByUsernameList: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserGetListByUsernameList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_getListByUsernameList: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x6e606b62:

		m.Names, err = iface.DecodeStringList(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_getListByUsernameList#0x6e606b62: field names: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_getListByUsernameList: invalid constructor %x", m.ClazzID)
	}
}

// TLUserDeleteUsernameByPeer <--
type TLUserDeleteUsernameByPeer struct {
	ClazzID  uint32 `json:"_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

func (m *TLUserDeleteUsernameByPeer) String() string {
	return iface.DebugStringWithName(ClazzName_user_deleteUsernameByPeer, m)
}

// Encode <--
func (m *TLUserDeleteUsernameByPeer) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_deleteUsernameByPeer, int(layer)); clazzId {
	case 0x7cafbc1:
		x.PutClazzID(0x7cafbc1)

		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_deleteUsernameByPeer: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserDeleteUsernameByPeer) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_deleteUsernameByPeer: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x7cafbc1:
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_deleteUsernameByPeer#0x7cafbc1: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_deleteUsernameByPeer#0x7cafbc1: field peer_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_deleteUsernameByPeer: invalid constructor %x", m.ClazzID)
	}
}

// TLUserSearchUsername <--
type TLUserSearchUsername struct {
	ClazzID          uint32  `json:"_id"`
	Q                string  `json:"q"`
	ExcludedContacts []int64 `json:"excluded_contacts"`
	Limit            int32   `json:"limit"`
}

func (m *TLUserSearchUsername) String() string {
	return iface.DebugStringWithName(ClazzName_user_searchUsername, m)
}

// Encode <--
func (m *TLUserSearchUsername) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_searchUsername, int(layer)); clazzId {
	case 0x266eb885:
		x.PutClazzID(0x266eb885)

		x.PutString(m.Q)

		iface.EncodeInt64List(x, m.ExcludedContacts)

		x.PutInt32(m.Limit)

		return nil
	default:
		return fmt.Errorf("unable to encode user_searchUsername: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserSearchUsername) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_searchUsername: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x266eb885:
		m.Q, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_searchUsername#0x266eb885: field q: %w", err)
		}

		m.ExcludedContacts, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_searchUsername#0x266eb885: field excluded_contacts: %w", err)
		}

		m.Limit, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_searchUsername#0x266eb885: field limit: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_searchUsername: invalid constructor %x", m.ClazzID)
	}
}

// TLUserToggleUsername <--
type TLUserToggleUsername struct {
	ClazzID  uint32       `json:"_id"`
	PeerType int32        `json:"peer_type"`
	PeerId   int64        `json:"peer_id"`
	Username string       `json:"username"`
	Active   tg.BoolClazz `json:"active"`
}

func (m *TLUserToggleUsername) String() string {
	return iface.DebugStringWithName(ClazzName_user_toggleUsername, m)
}

// Encode <--
func (m *TLUserToggleUsername) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_toggleUsername, int(layer)); clazzId {
	case 0xdd3b5a14:
		x.PutClazzID(0xdd3b5a14)

		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutString(m.Username)
		if m.Active == nil {
			return fmt.Errorf("unable to encode user_toggleUsername#0xdd3b5a14: field active is nil")
		}
		if err := m.Active.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode user_toggleUsername#0xdd3b5a14: field active: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode user_toggleUsername: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserToggleUsername) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_toggleUsername: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xdd3b5a14:
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_toggleUsername#0xdd3b5a14: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_toggleUsername#0xdd3b5a14: field peer_id: %w", err)
		}
		m.Username, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode user_toggleUsername#0xdd3b5a14: field username: %w", err)
		}

		m.Active, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_toggleUsername#0xdd3b5a14: field active: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_toggleUsername: invalid constructor %x", m.ClazzID)
	}
}

// TLUserReorderUsernames <--
type TLUserReorderUsernames struct {
	ClazzID      uint32   `json:"_id"`
	PeerType     int32    `json:"peer_type"`
	PeerId       int64    `json:"peer_id"`
	UsernameList []string `json:"username_list"`
}

func (m *TLUserReorderUsernames) String() string {
	return iface.DebugStringWithName(ClazzName_user_reorderUsernames, m)
}

// Encode <--
func (m *TLUserReorderUsernames) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_reorderUsernames, int(layer)); clazzId {
	case 0x3a61bdc0:
		x.PutClazzID(0x3a61bdc0)

		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)

		iface.EncodeStringList(x, m.UsernameList)

		return nil
	default:
		return fmt.Errorf("unable to encode user_reorderUsernames: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserReorderUsernames) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_reorderUsernames: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x3a61bdc0:
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode user_reorderUsernames#0x3a61bdc0: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_reorderUsernames#0x3a61bdc0: field peer_id: %w", err)
		}

		m.UsernameList, err = iface.DecodeStringList(d)
		if err != nil {
			return fmt.Errorf("unable to decode user_reorderUsernames#0x3a61bdc0: field username_list: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_reorderUsernames: invalid constructor %x", m.ClazzID)
	}
}

// TLUserDeactivateAllChannelUsernames <--
type TLUserDeactivateAllChannelUsernames struct {
	ClazzID   uint32 `json:"_id"`
	ChannelId int64  `json:"channel_id"`
}

func (m *TLUserDeactivateAllChannelUsernames) String() string {
	return iface.DebugStringWithName(ClazzName_user_deactivateAllChannelUsernames, m)
}

// Encode <--
func (m *TLUserDeactivateAllChannelUsernames) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_user_deactivateAllChannelUsernames, int(layer)); clazzId {
	case 0x9a5fe53c:
		x.PutClazzID(0x9a5fe53c)

		x.PutInt64(m.ChannelId)

		return nil
	default:
		return fmt.Errorf("unable to encode user_deactivateAllChannelUsernames: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserDeactivateAllChannelUsernames) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode user_deactivateAllChannelUsernames: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x9a5fe53c:
		m.ChannelId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode user_deactivateAllChannelUsernames#0x9a5fe53c: field channel_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode user_deactivateAllChannelUsernames: invalid constructor %x", m.ClazzID)
	}
}

// Vector api result type
// ----------------------------------------------------------------------------
// VectorResList <--

// VectorLastSeenData <--
type VectorLastSeenData struct {
	Datas []LastSeenDataClazz `json:"_datas"`
}

func (m *VectorLastSeenData) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorLastSeenData) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorLastSeenData) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[LastSeenDataClazz](d)

	return err
}

// VectorImmutableUser <--
type VectorImmutableUser struct {
	Datas []tg.ImmutableUserClazz `json:"_datas"`
}

func (m *VectorImmutableUser) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorImmutableUser) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorImmutableUser) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[tg.ImmutableUserClazz](d)

	return err
}

// VectorPeerPeerNotifySettings <--
type VectorPeerPeerNotifySettings struct {
	Datas []PeerPeerNotifySettingsClazz `json:"_datas"`
}

func (m *VectorPeerPeerNotifySettings) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorPeerPeerNotifySettings) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorPeerPeerNotifySettings) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[PeerPeerNotifySettingsClazz](d)

	return err
}

// VectorPrivacyRule <--
type VectorPrivacyRule struct {
	Datas []tg.PrivacyRuleClazz `json:"_datas"`
}

func (m *VectorPrivacyRule) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorPrivacyRule) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorPrivacyRule) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[tg.PrivacyRuleClazz](d)

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

// VectorPeerBlocked <--
type VectorPeerBlocked struct {
	Datas []tg.PeerBlockedClazz `json:"_datas"`
}

func (m *VectorPeerBlocked) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorPeerBlocked) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorPeerBlocked) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[tg.PeerBlockedClazz](d)

	return err
}

// VectorContactData <--
type VectorContactData struct {
	Datas []tg.ContactDataClazz `json:"_datas"`
}

func (m *VectorContactData) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorContactData) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorContactData) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[tg.ContactDataClazz](d)

	return err
}

// VectorInputContact <--
type VectorInputContact struct {
	Datas []tg.InputContactClazz `json:"_datas"`
}

func (m *VectorInputContact) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorInputContact) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorInputContact) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[tg.InputContactClazz](d)

	return err
}

// VectorUserData <--
type VectorUserData struct {
	Datas []tg.UserDataClazz `json:"_datas"`
}

func (m *VectorUserData) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorUserData) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorUserData) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[tg.UserDataClazz](d)

	return err
}

// VectorContactBirthday <--
type VectorContactBirthday struct {
	Datas []tg.ContactBirthdayClazz `json:"_datas"`
}

func (m *VectorContactBirthday) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorContactBirthday) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorContactBirthday) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[tg.ContactBirthdayClazz](d)

	return err
}

// VectorUsernameData <--
type VectorUsernameData struct {
	Datas []UsernameDataClazz `json:"_datas"`
}

func (m *VectorUsernameData) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorUsernameData) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorUsernameData) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[UsernameDataClazz](d)

	return err
}

// ----------------------------------------------------------------------------
// rpc

type RPCUser interface {
	UserGetLastSeens(ctx context.Context, in *TLUserGetLastSeens) (*VectorLastSeenData, error)
	UserUpdateLastSeen(ctx context.Context, in *TLUserUpdateLastSeen) (*tg.Bool, error)
	UserGetLastSeen(ctx context.Context, in *TLUserGetLastSeen) (*LastSeenData, error)
	UserGetImmutableUser(ctx context.Context, in *TLUserGetImmutableUser) (*tg.ImmutableUser, error)
	UserGetMutableUsers(ctx context.Context, in *TLUserGetMutableUsers) (*VectorImmutableUser, error)
	UserGetImmutableUserByPhone(ctx context.Context, in *TLUserGetImmutableUserByPhone) (*tg.ImmutableUser, error)
	UserGetImmutableUserByToken(ctx context.Context, in *TLUserGetImmutableUserByToken) (*tg.ImmutableUser, error)
	UserSetAccountDaysTTL(ctx context.Context, in *TLUserSetAccountDaysTTL) (*tg.Bool, error)
	UserGetAccountDaysTTL(ctx context.Context, in *TLUserGetAccountDaysTTL) (*tg.AccountDaysTTL, error)
	UserGetNotifySettings(ctx context.Context, in *TLUserGetNotifySettings) (*tg.PeerNotifySettings, error)
	UserGetNotifySettingsList(ctx context.Context, in *TLUserGetNotifySettingsList) (*VectorPeerPeerNotifySettings, error)
	UserSetNotifySettings(ctx context.Context, in *TLUserSetNotifySettings) (*tg.Bool, error)
	UserResetNotifySettings(ctx context.Context, in *TLUserResetNotifySettings) (*tg.Bool, error)
	UserGetAllNotifySettings(ctx context.Context, in *TLUserGetAllNotifySettings) (*VectorPeerPeerNotifySettings, error)
	UserGetGlobalPrivacySettings(ctx context.Context, in *TLUserGetGlobalPrivacySettings) (*tg.GlobalPrivacySettings, error)
	UserSetGlobalPrivacySettings(ctx context.Context, in *TLUserSetGlobalPrivacySettings) (*tg.Bool, error)
	UserGetPrivacy(ctx context.Context, in *TLUserGetPrivacy) (*VectorPrivacyRule, error)
	UserSetPrivacy(ctx context.Context, in *TLUserSetPrivacy) (*tg.Bool, error)
	UserCheckPrivacy(ctx context.Context, in *TLUserCheckPrivacy) (*tg.Bool, error)
	UserAddPeerSettings(ctx context.Context, in *TLUserAddPeerSettings) (*tg.Bool, error)
	UserGetPeerSettings(ctx context.Context, in *TLUserGetPeerSettings) (*tg.PeerSettings, error)
	UserDeletePeerSettings(ctx context.Context, in *TLUserDeletePeerSettings) (*tg.Bool, error)
	UserChangePhone(ctx context.Context, in *TLUserChangePhone) (*tg.Bool, error)
	UserCreateNewUser(ctx context.Context, in *TLUserCreateNewUser) (*tg.ImmutableUser, error)
	UserDeleteUser(ctx context.Context, in *TLUserDeleteUser) (*tg.Bool, error)
	UserBlockPeer(ctx context.Context, in *TLUserBlockPeer) (*tg.Bool, error)
	UserUnBlockPeer(ctx context.Context, in *TLUserUnBlockPeer) (*tg.Bool, error)
	UserBlockedByUser(ctx context.Context, in *TLUserBlockedByUser) (*tg.Bool, error)
	UserIsBlockedByUser(ctx context.Context, in *TLUserIsBlockedByUser) (*tg.Bool, error)
	UserCheckBlockUserList(ctx context.Context, in *TLUserCheckBlockUserList) (*VectorLong, error)
	UserGetBlockedList(ctx context.Context, in *TLUserGetBlockedList) (*VectorPeerBlocked, error)
	UserGetContactSignUpNotification(ctx context.Context, in *TLUserGetContactSignUpNotification) (*tg.Bool, error)
	UserSetContactSignUpNotification(ctx context.Context, in *TLUserSetContactSignUpNotification) (*tg.Bool, error)
	UserGetContentSettings(ctx context.Context, in *TLUserGetContentSettings) (*tg.AccountContentSettings, error)
	UserSetContentSettings(ctx context.Context, in *TLUserSetContentSettings) (*tg.Bool, error)
	UserDeleteContact(ctx context.Context, in *TLUserDeleteContact) (*tg.Bool, error)
	UserGetContactList(ctx context.Context, in *TLUserGetContactList) (*VectorContactData, error)
	UserGetContactIdList(ctx context.Context, in *TLUserGetContactIdList) (*VectorLong, error)
	UserGetContact(ctx context.Context, in *TLUserGetContact) (*tg.ContactData, error)
	UserAddContact(ctx context.Context, in *TLUserAddContact) (*tg.Bool, error)
	UserCheckContact(ctx context.Context, in *TLUserCheckContact) (*tg.Bool, error)
	UserGetImportersByPhone(ctx context.Context, in *TLUserGetImportersByPhone) (*VectorInputContact, error)
	UserDeleteImportersByPhone(ctx context.Context, in *TLUserDeleteImportersByPhone) (*tg.Bool, error)
	UserImportContacts(ctx context.Context, in *TLUserImportContacts) (*UserImportedContacts, error)
	UserGetCountryCode(ctx context.Context, in *TLUserGetCountryCode) (*tg.String, error)
	UserUpdateAbout(ctx context.Context, in *TLUserUpdateAbout) (*tg.Bool, error)
	UserUpdateFirstAndLastName(ctx context.Context, in *TLUserUpdateFirstAndLastName) (*tg.Bool, error)
	UserUpdateVerified(ctx context.Context, in *TLUserUpdateVerified) (*tg.Bool, error)
	UserUpdateUsername(ctx context.Context, in *TLUserUpdateUsername) (*tg.Bool, error)
	UserUpdateProfilePhoto(ctx context.Context, in *TLUserUpdateProfilePhoto) (*tg.Int64, error)
	UserDeleteProfilePhotos(ctx context.Context, in *TLUserDeleteProfilePhotos) (*tg.Int64, error)
	UserGetProfilePhotos(ctx context.Context, in *TLUserGetProfilePhotos) (*VectorLong, error)
	UserSetBotCommands(ctx context.Context, in *TLUserSetBotCommands) (*tg.Bool, error)
	UserIsBot(ctx context.Context, in *TLUserIsBot) (*tg.Bool, error)
	UserGetBotInfo(ctx context.Context, in *TLUserGetBotInfo) (*tg.BotInfo, error)
	UserCheckBots(ctx context.Context, in *TLUserCheckBots) (*VectorLong, error)
	UserGetFullUser(ctx context.Context, in *TLUserGetFullUser) (*tg.UsersUserFull, error)
	UserUpdateEmojiStatus(ctx context.Context, in *TLUserUpdateEmojiStatus) (*tg.Bool, error)
	UserGetUserDataById(ctx context.Context, in *TLUserGetUserDataById) (*tg.UserData, error)
	UserGetUserDataListByIdList(ctx context.Context, in *TLUserGetUserDataListByIdList) (*VectorUserData, error)
	UserGetUserDataByToken(ctx context.Context, in *TLUserGetUserDataByToken) (*tg.UserData, error)
	UserSearch(ctx context.Context, in *TLUserSearch) (*UsersFound, error)
	UserUpdateBotData(ctx context.Context, in *TLUserUpdateBotData) (*tg.Bool, error)
	UserGetImmutableUserV2(ctx context.Context, in *TLUserGetImmutableUserV2) (*tg.ImmutableUser, error)
	UserGetMutableUsersV2(ctx context.Context, in *TLUserGetMutableUsersV2) (*tg.MutableUsers, error)
	UserGetUserProjectionBundle(ctx context.Context, in *TLUserGetUserProjectionBundle) (*UserProjectionBundle, error)
	UserCreateNewTestUser(ctx context.Context, in *TLUserCreateNewTestUser) (*tg.ImmutableUser, error)
	UserEditCloseFriends(ctx context.Context, in *TLUserEditCloseFriends) (*tg.Bool, error)
	UserSetStoriesMaxId(ctx context.Context, in *TLUserSetStoriesMaxId) (*tg.Bool, error)
	UserSetColor(ctx context.Context, in *TLUserSetColor) (*tg.Bool, error)
	UserUpdateBirthday(ctx context.Context, in *TLUserUpdateBirthday) (*tg.Bool, error)
	UserGetBirthdays(ctx context.Context, in *TLUserGetBirthdays) (*VectorContactBirthday, error)
	UserSetStoriesHidden(ctx context.Context, in *TLUserSetStoriesHidden) (*tg.Bool, error)
	UserUpdatePersonalChannel(ctx context.Context, in *TLUserUpdatePersonalChannel) (*tg.Bool, error)
	UserGetUserIdByPhone(ctx context.Context, in *TLUserGetUserIdByPhone) (*tg.Int64, error)
	UserSetAuthorizationTTL(ctx context.Context, in *TLUserSetAuthorizationTTL) (*tg.Bool, error)
	UserGetAuthorizationTTL(ctx context.Context, in *TLUserGetAuthorizationTTL) (*tg.AccountDaysTTL, error)
	UserUpdatePremium(ctx context.Context, in *TLUserUpdatePremium) (*tg.Bool, error)
	UserGetBotInfoV2(ctx context.Context, in *TLUserGetBotInfoV2) (*BotInfoData, error)
	UserSaveMusic(ctx context.Context, in *TLUserSaveMusic) (*tg.Bool, error)
	UserGetSavedMusicIdList(ctx context.Context, in *TLUserGetSavedMusicIdList) (*VectorLong, error)
	UserSetMainProfileTab(ctx context.Context, in *TLUserSetMainProfileTab) (*tg.Bool, error)
	UserSetDefaultHistoryTTL(ctx context.Context, in *TLUserSetDefaultHistoryTTL) (*tg.Bool, error)
	UserGetDefaultHistoryTTL(ctx context.Context, in *TLUserGetDefaultHistoryTTL) (*tg.DefaultHistoryTTL, error)
	UserGetAccountUsername(ctx context.Context, in *TLUserGetAccountUsername) (*UsernameData, error)
	UserCheckAccountUsername(ctx context.Context, in *TLUserCheckAccountUsername) (*UsernameExist, error)
	UserGetChannelUsername(ctx context.Context, in *TLUserGetChannelUsername) (*UsernameData, error)
	UserCheckChannelUsername(ctx context.Context, in *TLUserCheckChannelUsername) (*UsernameExist, error)
	UserUpdateUsernameByPeer(ctx context.Context, in *TLUserUpdateUsernameByPeer) (*tg.Bool, error)
	UserCheckUsername(ctx context.Context, in *TLUserCheckUsername) (*UsernameExist, error)
	UserUpdateUsernameByUsername(ctx context.Context, in *TLUserUpdateUsernameByUsername) (*tg.Bool, error)
	UserDeleteUsername(ctx context.Context, in *TLUserDeleteUsername) (*tg.Bool, error)
	UserResolveUsername(ctx context.Context, in *TLUserResolveUsername) (*tg.Peer, error)
	UserGetListByUsernameList(ctx context.Context, in *TLUserGetListByUsernameList) (*VectorUsernameData, error)
	UserDeleteUsernameByPeer(ctx context.Context, in *TLUserDeleteUsernameByPeer) (*tg.Bool, error)
	UserSearchUsername(ctx context.Context, in *TLUserSearchUsername) (*VectorUsernameData, error)
	UserToggleUsername(ctx context.Context, in *TLUserToggleUsername) (*tg.Bool, error)
	UserReorderUsernames(ctx context.Context, in *TLUserReorderUsernames) (*tg.Bool, error)
	UserDeactivateAllChannelUsernames(ctx context.Context, in *TLUserDeactivateAllChannelUsernames) (*tg.Bool, error)
}
