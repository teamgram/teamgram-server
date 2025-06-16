/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package user

import (
	"context"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
)

var _ iface.TLObject
var _ fmt.Stringer
var _ *tg.Bool
var _ bin.Fields

// TLUserGetLastSeens <--
type TLUserGetLastSeens struct {
	ClazzID uint32  `json:"_id"`
	Id      []int64 `json:"id"`
}

// Encode <--
func (m *TLUserGetLastSeens) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x7ca17e01: func() error {
			x.PutClazzID(0x7ca17e01)

			iface.EncodeInt64List(x, m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getLastSeens, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getLastSeens, layer)
	}
}

// Decode <--
func (m *TLUserGetLastSeens) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x7ca17e01: func() (err error) {

			m.Id, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserUpdateLastSeen <--
type TLUserUpdateLastSeen struct {
	ClazzID    uint32 `json:"_id"`
	Id         int64  `json:"id"`
	LastSeenAt int64  `json:"last_seen_at"`
	Expires    int32  `json:"expires"`
}

// Encode <--
func (m *TLUserUpdateLastSeen) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xfd405a2d: func() error {
			x.PutClazzID(0xfd405a2d)

			x.PutInt64(m.Id)
			x.PutInt64(m.LastSeenAt)
			x.PutInt32(m.Expires)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_updateLastSeen, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_updateLastSeen, layer)
	}
}

// Decode <--
func (m *TLUserUpdateLastSeen) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xfd405a2d: func() (err error) {
			m.Id, err = d.Int64()
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

// TLUserGetLastSeen <--
type TLUserGetLastSeen struct {
	ClazzID uint32 `json:"_id"`
	Id      int64  `json:"id"`
}

// Encode <--
func (m *TLUserGetLastSeen) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x9119c8de: func() error {
			x.PutClazzID(0x9119c8de)

			x.PutInt64(m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getLastSeen, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getLastSeen, layer)
	}
}

// Decode <--
func (m *TLUserGetLastSeen) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x9119c8de: func() (err error) {
			m.Id, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetImmutableUser <--
type TLUserGetImmutableUser struct {
	ClazzID  uint32  `json:"_id"`
	Id       int64   `json:"id"`
	Privacy  bool    `json:"privacy"`
	Contacts []int64 `json:"contacts"`
}

// Encode <--
func (m *TLUserGetImmutableUser) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x376a6744: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getImmutableUser, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getImmutableUser, layer)
	}
}

// Decode <--
func (m *TLUserGetImmutableUser) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x376a6744: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.Id, err = d.Int64()
			if (flags & (1 << 1)) != 0 {
				m.Privacy = true
			}

			m.Contacts, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetMutableUsers <--
type TLUserGetMutableUsers struct {
	ClazzID uint32  `json:"_id"`
	Id      []int64 `json:"id"`
	To      []int64 `json:"to"`
}

// Encode <--
func (m *TLUserGetMutableUsers) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x9d3b23d7: func() error {
			x.PutClazzID(0x9d3b23d7)

			iface.EncodeInt64List(x, m.Id)

			iface.EncodeInt64List(x, m.To)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getMutableUsers, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getMutableUsers, layer)
	}
}

// Decode <--
func (m *TLUserGetMutableUsers) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x9d3b23d7: func() (err error) {

			m.Id, err = iface.DecodeInt64List(d)

			m.To, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetImmutableUserByPhone <--
type TLUserGetImmutableUserByPhone struct {
	ClazzID uint32 `json:"_id"`
	Phone   string `json:"phone"`
}

// Encode <--
func (m *TLUserGetImmutableUserByPhone) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xe9c36fe4: func() error {
			x.PutClazzID(0xe9c36fe4)

			x.PutString(m.Phone)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getImmutableUserByPhone, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getImmutableUserByPhone, layer)
	}
}

// Decode <--
func (m *TLUserGetImmutableUserByPhone) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xe9c36fe4: func() (err error) {
			m.Phone, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetImmutableUserByToken <--
type TLUserGetImmutableUserByToken struct {
	ClazzID uint32 `json:"_id"`
	Token   string `json:"token"`
}

// Encode <--
func (m *TLUserGetImmutableUserByToken) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xff3e1373: func() error {
			x.PutClazzID(0xff3e1373)

			x.PutString(m.Token)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getImmutableUserByToken, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getImmutableUserByToken, layer)
	}
}

// Decode <--
func (m *TLUserGetImmutableUserByToken) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xff3e1373: func() (err error) {
			m.Token, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserSetAccountDaysTTL <--
type TLUserSetAccountDaysTTL struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Ttl     int32  `json:"ttl"`
}

// Encode <--
func (m *TLUserSetAccountDaysTTL) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xd2550b4c: func() error {
			x.PutClazzID(0xd2550b4c)

			x.PutInt64(m.UserId)
			x.PutInt32(m.Ttl)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_setAccountDaysTTL, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_setAccountDaysTTL, layer)
	}
}

// Decode <--
func (m *TLUserSetAccountDaysTTL) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xd2550b4c: func() (err error) {
			m.UserId, err = d.Int64()
			m.Ttl, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetAccountDaysTTL <--
type TLUserGetAccountDaysTTL struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

// Encode <--
func (m *TLUserGetAccountDaysTTL) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xb2843ee0: func() error {
			x.PutClazzID(0xb2843ee0)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getAccountDaysTTL, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getAccountDaysTTL, layer)
	}
}

// Decode <--
func (m *TLUserGetAccountDaysTTL) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xb2843ee0: func() (err error) {
			m.UserId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetNotifySettings <--
type TLUserGetNotifySettings struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

// Encode <--
func (m *TLUserGetNotifySettings) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x40ac3766: func() error {
			x.PutClazzID(0x40ac3766)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getNotifySettings, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getNotifySettings, layer)
	}
}

// Decode <--
func (m *TLUserGetNotifySettings) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x40ac3766: func() (err error) {
			m.UserId, err = d.Int64()
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

// TLUserGetNotifySettingsList <--
type TLUserGetNotifySettingsList struct {
	ClazzID uint32         `json:"_id"`
	UserId  int64          `json:"user_id"`
	Peers   []*tg.PeerUtil `json:"peers"`
}

// Encode <--
func (m *TLUserGetNotifySettingsList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xe465159c: func() error {
			x.PutClazzID(0xe465159c)

			x.PutInt64(m.UserId)

			_ = iface.EncodeObjectList(x, m.Peers, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getNotifySettingsList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getNotifySettingsList, layer)
	}
}

// Decode <--
func (m *TLUserGetNotifySettingsList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xe465159c: func() (err error) {
			m.UserId, err = d.Int64()
			c2, err2 := d.ClazzID()
			if c2 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 2, c2)
				return err2
			}
			l2, err3 := d.Int()
			v2 := make([]*tg.PeerUtil, l2)
			for i := 0; i < l2; i++ {
				vv := new(tg.PeerUtil)
				err3 = vv.Decode(d)
				_ = err3
				v2[i] = vv
			}
			m.Peers = v2

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserSetNotifySettings <--
type TLUserSetNotifySettings struct {
	ClazzID  uint32                 `json:"_id"`
	UserId   int64                  `json:"user_id"`
	PeerType int32                  `json:"peer_type"`
	PeerId   int64                  `json:"peer_id"`
	Settings *tg.PeerNotifySettings `json:"settings"`
}

// Encode <--
func (m *TLUserSetNotifySettings) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc9ed65e5: func() error {
			x.PutClazzID(0xc9ed65e5)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			_ = m.Settings.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_setNotifySettings, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_setNotifySettings, layer)
	}
}

// Decode <--
func (m *TLUserSetNotifySettings) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xc9ed65e5: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()

			m4 := &tg.PeerNotifySettings{}
			_ = m4.Decode(d)
			m.Settings = m4

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserResetNotifySettings <--
type TLUserResetNotifySettings struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

// Encode <--
func (m *TLUserResetNotifySettings) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xe079d74: func() error {
			x.PutClazzID(0xe079d74)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_resetNotifySettings, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_resetNotifySettings, layer)
	}
}

// Decode <--
func (m *TLUserResetNotifySettings) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xe079d74: func() (err error) {
			m.UserId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetAllNotifySettings <--
type TLUserGetAllNotifySettings struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

// Encode <--
func (m *TLUserGetAllNotifySettings) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x55926875: func() error {
			x.PutClazzID(0x55926875)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getAllNotifySettings, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getAllNotifySettings, layer)
	}
}

// Decode <--
func (m *TLUserGetAllNotifySettings) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x55926875: func() (err error) {
			m.UserId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetGlobalPrivacySettings <--
type TLUserGetGlobalPrivacySettings struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

// Encode <--
func (m *TLUserGetGlobalPrivacySettings) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x77f6f112: func() error {
			x.PutClazzID(0x77f6f112)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getGlobalPrivacySettings, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getGlobalPrivacySettings, layer)
	}
}

// Decode <--
func (m *TLUserGetGlobalPrivacySettings) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x77f6f112: func() (err error) {
			m.UserId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserSetGlobalPrivacySettings <--
type TLUserSetGlobalPrivacySettings struct {
	ClazzID  uint32                    `json:"_id"`
	UserId   int64                     `json:"user_id"`
	Settings *tg.GlobalPrivacySettings `json:"settings"`
}

// Encode <--
func (m *TLUserSetGlobalPrivacySettings) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x8cb592ae: func() error {
			x.PutClazzID(0x8cb592ae)

			x.PutInt64(m.UserId)
			_ = m.Settings.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_setGlobalPrivacySettings, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_setGlobalPrivacySettings, layer)
	}
}

// Decode <--
func (m *TLUserSetGlobalPrivacySettings) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x8cb592ae: func() (err error) {
			m.UserId, err = d.Int64()

			m2 := &tg.GlobalPrivacySettings{}
			_ = m2.Decode(d)
			m.Settings = m2

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetPrivacy <--
type TLUserGetPrivacy struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	KeyType int32  `json:"key_type"`
}

// Encode <--
func (m *TLUserGetPrivacy) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x9d40a3b4: func() error {
			x.PutClazzID(0x9d40a3b4)

			x.PutInt64(m.UserId)
			x.PutInt32(m.KeyType)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getPrivacy, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getPrivacy, layer)
	}
}

// Decode <--
func (m *TLUserGetPrivacy) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x9d40a3b4: func() (err error) {
			m.UserId, err = d.Int64()
			m.KeyType, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserSetPrivacy <--
type TLUserSetPrivacy struct {
	ClazzID uint32            `json:"_id"`
	UserId  int64             `json:"user_id"`
	KeyType int32             `json:"key_type"`
	Rules   []*tg.PrivacyRule `json:"rules"`
}

// Encode <--
func (m *TLUserSetPrivacy) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x8855ad8f: func() error {
			x.PutClazzID(0x8855ad8f)

			x.PutInt64(m.UserId)
			x.PutInt32(m.KeyType)

			_ = iface.EncodeObjectList(x, m.Rules, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_setPrivacy, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_setPrivacy, layer)
	}
}

// Decode <--
func (m *TLUserSetPrivacy) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x8855ad8f: func() (err error) {
			m.UserId, err = d.Int64()
			m.KeyType, err = d.Int32()
			c3, err2 := d.ClazzID()
			if c3 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 3, c3)
				return err2
			}
			l3, err3 := d.Int()
			v3 := make([]*tg.PrivacyRule, l3)
			for i := 0; i < l3; i++ {
				vv := new(tg.PrivacyRule)
				err3 = vv.Decode(d)
				_ = err3
				v3[i] = vv
			}
			m.Rules = v3

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserCheckPrivacy <--
type TLUserCheckPrivacy struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	KeyType int32  `json:"key_type"`
	PeerId  int64  `json:"peer_id"`
}

// Encode <--
func (m *TLUserCheckPrivacy) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc56e1eaa: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_checkPrivacy, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_checkPrivacy, layer)
	}
}

// Decode <--
func (m *TLUserCheckPrivacy) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xc56e1eaa: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.UserId, err = d.Int64()
			m.KeyType, err = d.Int32()
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

// TLUserAddPeerSettings <--
type TLUserAddPeerSettings struct {
	ClazzID  uint32           `json:"_id"`
	UserId   int64            `json:"user_id"`
	PeerType int32            `json:"peer_type"`
	PeerId   int64            `json:"peer_id"`
	Settings *tg.PeerSettings `json:"settings"`
}

// Encode <--
func (m *TLUserAddPeerSettings) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xcae22763: func() error {
			x.PutClazzID(0xcae22763)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			_ = m.Settings.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_addPeerSettings, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_addPeerSettings, layer)
	}
}

// Decode <--
func (m *TLUserAddPeerSettings) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xcae22763: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()

			m4 := &tg.PeerSettings{}
			_ = m4.Decode(d)
			m.Settings = m4

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetPeerSettings <--
type TLUserGetPeerSettings struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

// Encode <--
func (m *TLUserGetPeerSettings) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xd02ef67: func() error {
			x.PutClazzID(0xd02ef67)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getPeerSettings, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getPeerSettings, layer)
	}
}

// Decode <--
func (m *TLUserGetPeerSettings) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xd02ef67: func() (err error) {
			m.UserId, err = d.Int64()
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

// TLUserDeletePeerSettings <--
type TLUserDeletePeerSettings struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

// Encode <--
func (m *TLUserDeletePeerSettings) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x5e891967: func() error {
			x.PutClazzID(0x5e891967)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_deletePeerSettings, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_deletePeerSettings, layer)
	}
}

// Decode <--
func (m *TLUserDeletePeerSettings) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x5e891967: func() (err error) {
			m.UserId, err = d.Int64()
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

// TLUserChangePhone <--
type TLUserChangePhone struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Phone   string `json:"phone"`
}

// Encode <--
func (m *TLUserChangePhone) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xff7a575b: func() error {
			x.PutClazzID(0xff7a575b)

			x.PutInt64(m.UserId)
			x.PutString(m.Phone)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_changePhone, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_changePhone, layer)
	}
}

// Decode <--
func (m *TLUserChangePhone) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xff7a575b: func() (err error) {
			m.UserId, err = d.Int64()
			m.Phone, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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

// Encode <--
func (m *TLUserCreateNewUser) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x79e01881: func() error {
			x.PutClazzID(0x79e01881)

			x.PutInt64(m.SecretKeyId)
			x.PutString(m.Phone)
			x.PutString(m.CountryCode)
			x.PutString(m.FirstName)
			x.PutString(m.LastName)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_createNewUser, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_createNewUser, layer)
	}
}

// Decode <--
func (m *TLUserCreateNewUser) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x79e01881: func() (err error) {
			m.SecretKeyId, err = d.Int64()
			m.Phone, err = d.String()
			m.CountryCode, err = d.String()
			m.FirstName, err = d.String()
			m.LastName, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserDeleteUser <--
type TLUserDeleteUser struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Reason  string `json:"reason"`
	Phone   string `json:"phone"`
}

// Encode <--
func (m *TLUserDeleteUser) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x626dbd10: func() error {
			x.PutClazzID(0x626dbd10)

			x.PutInt64(m.UserId)
			x.PutString(m.Reason)
			x.PutString(m.Phone)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_deleteUser, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_deleteUser, layer)
	}
}

// Decode <--
func (m *TLUserDeleteUser) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x626dbd10: func() (err error) {
			m.UserId, err = d.Int64()
			m.Reason, err = d.String()
			m.Phone, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserBlockPeer <--
type TLUserBlockPeer struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

// Encode <--
func (m *TLUserBlockPeer) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x81062eb0: func() error {
			x.PutClazzID(0x81062eb0)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_blockPeer, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_blockPeer, layer)
	}
}

// Decode <--
func (m *TLUserBlockPeer) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x81062eb0: func() (err error) {
			m.UserId, err = d.Int64()
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

// TLUserUnBlockPeer <--
type TLUserUnBlockPeer struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

// Encode <--
func (m *TLUserUnBlockPeer) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xdee7160d: func() error {
			x.PutClazzID(0xdee7160d)

			x.PutInt64(m.UserId)
			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_unBlockPeer, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_unBlockPeer, layer)
	}
}

// Decode <--
func (m *TLUserUnBlockPeer) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xdee7160d: func() (err error) {
			m.UserId, err = d.Int64()
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

// TLUserBlockedByUser <--
type TLUserBlockedByUser struct {
	ClazzID    uint32 `json:"_id"`
	UserId     int64  `json:"user_id"`
	PeerUserId int64  `json:"peer_user_id"`
}

// Encode <--
func (m *TLUserBlockedByUser) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xbba0058e: func() error {
			x.PutClazzID(0xbba0058e)

			x.PutInt64(m.UserId)
			x.PutInt64(m.PeerUserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_blockedByUser, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_blockedByUser, layer)
	}
}

// Decode <--
func (m *TLUserBlockedByUser) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xbba0058e: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerUserId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserIsBlockedByUser <--
type TLUserIsBlockedByUser struct {
	ClazzID    uint32 `json:"_id"`
	UserId     int64  `json:"user_id"`
	PeerUserId int64  `json:"peer_user_id"`
}

// Encode <--
func (m *TLUserIsBlockedByUser) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x8caeb1df: func() error {
			x.PutClazzID(0x8caeb1df)

			x.PutInt64(m.UserId)
			x.PutInt64(m.PeerUserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_isBlockedByUser, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_isBlockedByUser, layer)
	}
}

// Decode <--
func (m *TLUserIsBlockedByUser) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x8caeb1df: func() (err error) {
			m.UserId, err = d.Int64()
			m.PeerUserId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserCheckBlockUserList <--
type TLUserCheckBlockUserList struct {
	ClazzID uint32  `json:"_id"`
	UserId  int64   `json:"user_id"`
	Id      []int64 `json:"id"`
}

// Encode <--
func (m *TLUserCheckBlockUserList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc3fd70f0: func() error {
			x.PutClazzID(0xc3fd70f0)

			x.PutInt64(m.UserId)

			iface.EncodeInt64List(x, m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_checkBlockUserList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_checkBlockUserList, layer)
	}
}

// Decode <--
func (m *TLUserCheckBlockUserList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xc3fd70f0: func() (err error) {
			m.UserId, err = d.Int64()

			m.Id, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetBlockedList <--
type TLUserGetBlockedList struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Offset  int32  `json:"offset"`
	Limit   int32  `json:"limit"`
}

// Encode <--
func (m *TLUserGetBlockedList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x23ffc348: func() error {
			x.PutClazzID(0x23ffc348)

			x.PutInt64(m.UserId)
			x.PutInt32(m.Offset)
			x.PutInt32(m.Limit)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getBlockedList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getBlockedList, layer)
	}
}

// Decode <--
func (m *TLUserGetBlockedList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x23ffc348: func() (err error) {
			m.UserId, err = d.Int64()
			m.Offset, err = d.Int32()
			m.Limit, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetContactSignUpNotification <--
type TLUserGetContactSignUpNotification struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

// Encode <--
func (m *TLUserGetContactSignUpNotification) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xe4d1d3d6: func() error {
			x.PutClazzID(0xe4d1d3d6)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getContactSignUpNotification, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getContactSignUpNotification, layer)
	}
}

// Decode <--
func (m *TLUserGetContactSignUpNotification) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xe4d1d3d6: func() (err error) {
			m.UserId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserSetContactSignUpNotification <--
type TLUserSetContactSignUpNotification struct {
	ClazzID uint32   `json:"_id"`
	UserId  int64    `json:"user_id"`
	Silent  *tg.Bool `json:"silent"`
}

// Encode <--
func (m *TLUserSetContactSignUpNotification) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x85a17361: func() error {
			x.PutClazzID(0x85a17361)

			x.PutInt64(m.UserId)
			_ = m.Silent.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_setContactSignUpNotification, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_setContactSignUpNotification, layer)
	}
}

// Decode <--
func (m *TLUserSetContactSignUpNotification) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x85a17361: func() (err error) {
			m.UserId, err = d.Int64()

			m2 := &tg.Bool{}
			_ = m2.Decode(d)
			m.Silent = m2

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetContentSettings <--
type TLUserGetContentSettings struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

// Encode <--
func (m *TLUserGetContentSettings) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x94c3ad9f: func() error {
			x.PutClazzID(0x94c3ad9f)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getContentSettings, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getContentSettings, layer)
	}
}

// Decode <--
func (m *TLUserGetContentSettings) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x94c3ad9f: func() (err error) {
			m.UserId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserSetContentSettings <--
type TLUserSetContentSettings struct {
	ClazzID          uint32 `json:"_id"`
	UserId           int64  `json:"user_id"`
	SensitiveEnabled bool   `json:"sensitive_enabled"`
}

// Encode <--
func (m *TLUserSetContentSettings) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x9d63fe6b: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_setContentSettings, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_setContentSettings, layer)
	}
}

// Decode <--
func (m *TLUserSetContentSettings) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x9d63fe6b: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.UserId, err = d.Int64()
			if (flags & (1 << 0)) != 0 {
				m.SensitiveEnabled = true
			}

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserDeleteContact <--
type TLUserDeleteContact struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Id      int64  `json:"id"`
}

// Encode <--
func (m *TLUserDeleteContact) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc6018219: func() error {
			x.PutClazzID(0xc6018219)

			x.PutInt64(m.UserId)
			x.PutInt64(m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_deleteContact, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_deleteContact, layer)
	}
}

// Decode <--
func (m *TLUserDeleteContact) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xc6018219: func() (err error) {
			m.UserId, err = d.Int64()
			m.Id, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetContactList <--
type TLUserGetContactList struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

// Encode <--
func (m *TLUserGetContactList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc74bd161: func() error {
			x.PutClazzID(0xc74bd161)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getContactList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getContactList, layer)
	}
}

// Decode <--
func (m *TLUserGetContactList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xc74bd161: func() (err error) {
			m.UserId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetContactIdList <--
type TLUserGetContactIdList struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

// Encode <--
func (m *TLUserGetContactIdList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xf1dd983e: func() error {
			x.PutClazzID(0xf1dd983e)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getContactIdList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getContactIdList, layer)
	}
}

// Decode <--
func (m *TLUserGetContactIdList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xf1dd983e: func() (err error) {
			m.UserId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetContact <--
type TLUserGetContact struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Id      int64  `json:"id"`
}

// Encode <--
func (m *TLUserGetContact) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xdb728be3: func() error {
			x.PutClazzID(0xdb728be3)

			x.PutInt64(m.UserId)
			x.PutInt64(m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getContact, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getContact, layer)
	}
}

// Decode <--
func (m *TLUserGetContact) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xdb728be3: func() (err error) {
			m.UserId, err = d.Int64()
			m.Id, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserAddContact <--
type TLUserAddContact struct {
	ClazzID                  uint32   `json:"_id"`
	UserId                   int64    `json:"user_id"`
	AddPhonePrivacyException *tg.Bool `json:"add_phone_privacy_exception"`
	Id                       int64    `json:"id"`
	FirstName                string   `json:"first_name"`
	LastName                 string   `json:"last_name"`
	Phone                    string   `json:"phone"`
}

// Encode <--
func (m *TLUserAddContact) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x79c4bd0e: func() error {
			x.PutClazzID(0x79c4bd0e)

			x.PutInt64(m.UserId)
			_ = m.AddPhonePrivacyException.Encode(x, layer)
			x.PutInt64(m.Id)
			x.PutString(m.FirstName)
			x.PutString(m.LastName)
			x.PutString(m.Phone)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_addContact, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_addContact, layer)
	}
}

// Decode <--
func (m *TLUserAddContact) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x79c4bd0e: func() (err error) {
			m.UserId, err = d.Int64()

			m2 := &tg.Bool{}
			_ = m2.Decode(d)
			m.AddPhonePrivacyException = m2

			m.Id, err = d.Int64()
			m.FirstName, err = d.String()
			m.LastName, err = d.String()
			m.Phone, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserCheckContact <--
type TLUserCheckContact struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Id      int64  `json:"id"`
}

// Encode <--
func (m *TLUserCheckContact) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x82a758a4: func() error {
			x.PutClazzID(0x82a758a4)

			x.PutInt64(m.UserId)
			x.PutInt64(m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_checkContact, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_checkContact, layer)
	}
}

// Decode <--
func (m *TLUserCheckContact) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x82a758a4: func() (err error) {
			m.UserId, err = d.Int64()
			m.Id, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetImportersByPhone <--
type TLUserGetImportersByPhone struct {
	ClazzID uint32 `json:"_id"`
	Phone   string `json:"phone"`
}

// Encode <--
func (m *TLUserGetImportersByPhone) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x47aa8212: func() error {
			x.PutClazzID(0x47aa8212)

			x.PutString(m.Phone)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getImportersByPhone, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getImportersByPhone, layer)
	}
}

// Decode <--
func (m *TLUserGetImportersByPhone) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x47aa8212: func() (err error) {
			m.Phone, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserDeleteImportersByPhone <--
type TLUserDeleteImportersByPhone struct {
	ClazzID uint32 `json:"_id"`
	Phone   string `json:"phone"`
}

// Encode <--
func (m *TLUserDeleteImportersByPhone) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x174ddc54: func() error {
			x.PutClazzID(0x174ddc54)

			x.PutString(m.Phone)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_deleteImportersByPhone, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_deleteImportersByPhone, layer)
	}
}

// Decode <--
func (m *TLUserDeleteImportersByPhone) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x174ddc54: func() (err error) {
			m.Phone, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserImportContacts <--
type TLUserImportContacts struct {
	ClazzID  uint32             `json:"_id"`
	UserId   int64              `json:"user_id"`
	Contacts []*tg.InputContact `json:"contacts"`
}

// Encode <--
func (m *TLUserImportContacts) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x9a00f792: func() error {
			x.PutClazzID(0x9a00f792)

			x.PutInt64(m.UserId)

			_ = iface.EncodeObjectList(x, m.Contacts, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_importContacts, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_importContacts, layer)
	}
}

// Decode <--
func (m *TLUserImportContacts) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x9a00f792: func() (err error) {
			m.UserId, err = d.Int64()
			c2, err2 := d.ClazzID()
			if c2 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 2, c2)
				return err2
			}
			l2, err3 := d.Int()
			v2 := make([]*tg.InputContact, l2)
			for i := 0; i < l2; i++ {
				vv := new(tg.InputContact)
				err3 = vv.Decode(d)
				_ = err3
				v2[i] = vv
			}
			m.Contacts = v2

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetCountryCode <--
type TLUserGetCountryCode struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

// Encode <--
func (m *TLUserGetCountryCode) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x12006832: func() error {
			x.PutClazzID(0x12006832)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getCountryCode, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getCountryCode, layer)
	}
}

// Decode <--
func (m *TLUserGetCountryCode) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x12006832: func() (err error) {
			m.UserId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserUpdateAbout <--
type TLUserUpdateAbout struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	About   string `json:"about"`
}

// Encode <--
func (m *TLUserUpdateAbout) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xdb00f187: func() error {
			x.PutClazzID(0xdb00f187)

			x.PutInt64(m.UserId)
			x.PutString(m.About)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_updateAbout, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_updateAbout, layer)
	}
}

// Decode <--
func (m *TLUserUpdateAbout) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xdb00f187: func() (err error) {
			m.UserId, err = d.Int64()
			m.About, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserUpdateFirstAndLastName <--
type TLUserUpdateFirstAndLastName struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Encode <--
func (m *TLUserUpdateFirstAndLastName) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xcb6685ec: func() error {
			x.PutClazzID(0xcb6685ec)

			x.PutInt64(m.UserId)
			x.PutString(m.FirstName)
			x.PutString(m.LastName)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_updateFirstAndLastName, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_updateFirstAndLastName, layer)
	}
}

// Decode <--
func (m *TLUserUpdateFirstAndLastName) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xcb6685ec: func() (err error) {
			m.UserId, err = d.Int64()
			m.FirstName, err = d.String()
			m.LastName, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserUpdateVerified <--
type TLUserUpdateVerified struct {
	ClazzID  uint32   `json:"_id"`
	UserId   int64    `json:"user_id"`
	Verified *tg.Bool `json:"verified"`
}

// Encode <--
func (m *TLUserUpdateVerified) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x24c92963: func() error {
			x.PutClazzID(0x24c92963)

			x.PutInt64(m.UserId)
			_ = m.Verified.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_updateVerified, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_updateVerified, layer)
	}
}

// Decode <--
func (m *TLUserUpdateVerified) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x24c92963: func() (err error) {
			m.UserId, err = d.Int64()

			m2 := &tg.Bool{}
			_ = m2.Decode(d)
			m.Verified = m2

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserUpdateUsername <--
type TLUserUpdateUsername struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
}

// Encode <--
func (m *TLUserUpdateUsername) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xf54d1e71: func() error {
			x.PutClazzID(0xf54d1e71)

			x.PutInt64(m.UserId)
			x.PutString(m.Username)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_updateUsername, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_updateUsername, layer)
	}
}

// Decode <--
func (m *TLUserUpdateUsername) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xf54d1e71: func() (err error) {
			m.UserId, err = d.Int64()
			m.Username, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserUpdateProfilePhoto <--
type TLUserUpdateProfilePhoto struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Id      int64  `json:"id"`
}

// Encode <--
func (m *TLUserUpdateProfilePhoto) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x3b740f87: func() error {
			x.PutClazzID(0x3b740f87)

			x.PutInt64(m.UserId)
			x.PutInt64(m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_updateProfilePhoto, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_updateProfilePhoto, layer)
	}
}

// Decode <--
func (m *TLUserUpdateProfilePhoto) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x3b740f87: func() (err error) {
			m.UserId, err = d.Int64()
			m.Id, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserDeleteProfilePhotos <--
type TLUserDeleteProfilePhotos struct {
	ClazzID uint32  `json:"_id"`
	UserId  int64   `json:"user_id"`
	Id      []int64 `json:"id"`
}

// Encode <--
func (m *TLUserDeleteProfilePhotos) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x2be3620e: func() error {
			x.PutClazzID(0x2be3620e)

			x.PutInt64(m.UserId)

			iface.EncodeInt64List(x, m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_deleteProfilePhotos, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_deleteProfilePhotos, layer)
	}
}

// Decode <--
func (m *TLUserDeleteProfilePhotos) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x2be3620e: func() (err error) {
			m.UserId, err = d.Int64()

			m.Id, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetProfilePhotos <--
type TLUserGetProfilePhotos struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

// Encode <--
func (m *TLUserGetProfilePhotos) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xdc66c146: func() error {
			x.PutClazzID(0xdc66c146)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getProfilePhotos, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getProfilePhotos, layer)
	}
}

// Decode <--
func (m *TLUserGetProfilePhotos) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xdc66c146: func() (err error) {
			m.UserId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserSetBotCommands <--
type TLUserSetBotCommands struct {
	ClazzID  uint32           `json:"_id"`
	UserId   int64            `json:"user_id"`
	BotId    int64            `json:"bot_id"`
	Commands []*tg.BotCommand `json:"commands"`
}

// Encode <--
func (m *TLUserSetBotCommands) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x753ba916: func() error {
			x.PutClazzID(0x753ba916)

			x.PutInt64(m.UserId)
			x.PutInt64(m.BotId)

			_ = iface.EncodeObjectList(x, m.Commands, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_setBotCommands, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_setBotCommands, layer)
	}
}

// Decode <--
func (m *TLUserSetBotCommands) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x753ba916: func() (err error) {
			m.UserId, err = d.Int64()
			m.BotId, err = d.Int64()
			c3, err2 := d.ClazzID()
			if c3 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 3, c3)
				return err2
			}
			l3, err3 := d.Int()
			v3 := make([]*tg.BotCommand, l3)
			for i := 0; i < l3; i++ {
				vv := new(tg.BotCommand)
				err3 = vv.Decode(d)
				_ = err3
				v3[i] = vv
			}
			m.Commands = v3

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserIsBot <--
type TLUserIsBot struct {
	ClazzID uint32 `json:"_id"`
	Id      int64  `json:"id"`
}

// Encode <--
func (m *TLUserIsBot) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc772c7ee: func() error {
			x.PutClazzID(0xc772c7ee)

			x.PutInt64(m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_isBot, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_isBot, layer)
	}
}

// Decode <--
func (m *TLUserIsBot) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xc772c7ee: func() (err error) {
			m.Id, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetBotInfo <--
type TLUserGetBotInfo struct {
	ClazzID uint32 `json:"_id"`
	BotId   int64  `json:"bot_id"`
}

// Encode <--
func (m *TLUserGetBotInfo) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x34663710: func() error {
			x.PutClazzID(0x34663710)

			x.PutInt64(m.BotId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getBotInfo, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getBotInfo, layer)
	}
}

// Decode <--
func (m *TLUserGetBotInfo) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x34663710: func() (err error) {
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

// TLUserCheckBots <--
type TLUserCheckBots struct {
	ClazzID uint32  `json:"_id"`
	Id      []int64 `json:"id"`
}

// Encode <--
func (m *TLUserCheckBots) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x736500c1: func() error {
			x.PutClazzID(0x736500c1)

			iface.EncodeInt64List(x, m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_checkBots, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_checkBots, layer)
	}
}

// Decode <--
func (m *TLUserCheckBots) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x736500c1: func() (err error) {

			m.Id, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetFullUser <--
type TLUserGetFullUser struct {
	ClazzID    uint32 `json:"_id"`
	SelfUserId int64  `json:"self_user_id"`
	Id         int64  `json:"id"`
}

// Encode <--
func (m *TLUserGetFullUser) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xfd10e13a: func() error {
			x.PutClazzID(0xfd10e13a)

			x.PutInt64(m.SelfUserId)
			x.PutInt64(m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getFullUser, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getFullUser, layer)
	}
}

// Decode <--
func (m *TLUserGetFullUser) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xfd10e13a: func() (err error) {
			m.SelfUserId, err = d.Int64()
			m.Id, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserUpdateEmojiStatus <--
type TLUserUpdateEmojiStatus struct {
	ClazzID               uint32 `json:"_id"`
	UserId                int64  `json:"user_id"`
	EmojiStatusDocumentId int64  `json:"emoji_status_document_id"`
	EmojiStatusUntil      int32  `json:"emoji_status_until"`
}

// Encode <--
func (m *TLUserUpdateEmojiStatus) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xf8c8bad8: func() error {
			x.PutClazzID(0xf8c8bad8)

			x.PutInt64(m.UserId)
			x.PutInt64(m.EmojiStatusDocumentId)
			x.PutInt32(m.EmojiStatusUntil)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_updateEmojiStatus, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_updateEmojiStatus, layer)
	}
}

// Decode <--
func (m *TLUserUpdateEmojiStatus) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xf8c8bad8: func() (err error) {
			m.UserId, err = d.Int64()
			m.EmojiStatusDocumentId, err = d.Int64()
			m.EmojiStatusUntil, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetUserDataById <--
type TLUserGetUserDataById struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

// Encode <--
func (m *TLUserGetUserDataById) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x3bb7103: func() error {
			x.PutClazzID(0x3bb7103)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getUserDataById, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getUserDataById, layer)
	}
}

// Decode <--
func (m *TLUserGetUserDataById) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x3bb7103: func() (err error) {
			m.UserId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetUserDataListByIdList <--
type TLUserGetUserDataListByIdList struct {
	ClazzID    uint32  `json:"_id"`
	UserIdList []int64 `json:"user_id_list"`
}

// Encode <--
func (m *TLUserGetUserDataListByIdList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x8191eff9: func() error {
			x.PutClazzID(0x8191eff9)

			iface.EncodeInt64List(x, m.UserIdList)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getUserDataListByIdList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getUserDataListByIdList, layer)
	}
}

// Decode <--
func (m *TLUserGetUserDataListByIdList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x8191eff9: func() (err error) {

			m.UserIdList, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetUserDataByToken <--
type TLUserGetUserDataByToken struct {
	ClazzID uint32 `json:"_id"`
	Token   string `json:"token"`
}

// Encode <--
func (m *TLUserGetUserDataByToken) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x3f09659e: func() error {
			x.PutClazzID(0x3f09659e)

			x.PutString(m.Token)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getUserDataByToken, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getUserDataByToken, layer)
	}
}

// Decode <--
func (m *TLUserGetUserDataByToken) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x3f09659e: func() (err error) {
			m.Token, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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

// Encode <--
func (m *TLUserSearch) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x7035b6cd: func() error {
			x.PutClazzID(0x7035b6cd)

			x.PutString(m.Q)

			iface.EncodeInt64List(x, m.ExcludedContacts)

			x.PutInt64(m.Offset)
			x.PutInt32(m.Limit)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_search, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_search, layer)
	}
}

// Decode <--
func (m *TLUserSearch) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x7035b6cd: func() (err error) {
			m.Q, err = d.String()

			m.ExcludedContacts, err = iface.DecodeInt64List(d)

			m.Offset, err = d.Int64()
			m.Limit, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserUpdateBotData <--
type TLUserUpdateBotData struct {
	ClazzID              uint32   `json:"_id"`
	BotId                int64    `json:"bot_id"`
	BotChatHistory       *tg.Bool `json:"bot_chat_history"`
	BotNochats           *tg.Bool `json:"bot_nochats"`
	BotInlineGeo         *tg.Bool `json:"bot_inline_geo"`
	BotAttachMenu        *tg.Bool `json:"bot_attach_menu"`
	BotInlinePlaceholder *tg.Bool `json:"bot_inline_placeholder"`
}

// Encode <--
func (m *TLUserUpdateBotData) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xb9fd39ee: func() error {
			x.PutClazzID(0xb9fd39ee)

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

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.BotId)
			if m.BotChatHistory != nil {
				_ = m.BotChatHistory.Encode(x, layer)
			}

			if m.BotNochats != nil {
				_ = m.BotNochats.Encode(x, layer)
			}

			if m.BotInlineGeo != nil {
				_ = m.BotInlineGeo.Encode(x, layer)
			}

			if m.BotAttachMenu != nil {
				_ = m.BotAttachMenu.Encode(x, layer)
			}

			if m.BotInlinePlaceholder != nil {
				_ = m.BotInlinePlaceholder.Encode(x, layer)
			}

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_updateBotData, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_updateBotData, layer)
	}
}

// Decode <--
func (m *TLUserUpdateBotData) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xb9fd39ee: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.BotId, err = d.Int64()
			if (flags & (1 << 15)) != 0 {
				m3 := &tg.Bool{}
				_ = m3.Decode(d)
				m.BotChatHistory = m3
			}
			if (flags & (1 << 16)) != 0 {
				m4 := &tg.Bool{}
				_ = m4.Decode(d)
				m.BotNochats = m4
			}
			if (flags & (1 << 21)) != 0 {
				m5 := &tg.Bool{}
				_ = m5.Decode(d)
				m.BotInlineGeo = m5
			}
			if (flags & (1 << 27)) != 0 {
				m6 := &tg.Bool{}
				_ = m6.Decode(d)
				m.BotAttachMenu = m6
			}
			if (flags & (1 << 19)) != 0 {
				m7 := &tg.Bool{}
				_ = m7.Decode(d)
				m.BotInlinePlaceholder = m7
			}

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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

// Encode <--
func (m *TLUserGetImmutableUserV2) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x300aba4c: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getImmutableUserV2, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getImmutableUserV2, layer)
	}
}

// Decode <--
func (m *TLUserGetImmutableUserV2) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x300aba4c: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.Id, err = d.Int64()
			if (flags & (1 << 0)) != 0 {
				m.Privacy = true
			}
			if (flags & (1 << 2)) != 0 {
				m.HasTo = true
			}
			if (flags & (1 << 2)) != 0 {
				m.To, err = iface.DecodeInt64List(d)
			}

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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

// Encode <--
func (m *TLUserGetMutableUsersV2) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x94f98b28: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getMutableUsersV2, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getMutableUsersV2, layer)
	}
}

// Decode <--
func (m *TLUserGetMutableUsersV2) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x94f98b28: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags

			m.Id, err = iface.DecodeInt64List(d)

			if (flags & (1 << 0)) != 0 {
				m.Privacy = true
			}
			if (flags & (1 << 2)) != 0 {
				m.HasTo = true
			}
			if (flags & (1 << 2)) != 0 {
				m.To, err = iface.DecodeInt64List(d)
			}

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserCreateNewTestUser <--
type TLUserCreateNewTestUser struct {
	ClazzID     uint32 `json:"_id"`
	SecretKeyId int64  `json:"secret_key_id"`
	MinId       int64  `json:"min_id"`
	MaxId       int64  `json:"max_id"`
}

// Encode <--
func (m *TLUserCreateNewTestUser) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x4c6eccab: func() error {
			x.PutClazzID(0x4c6eccab)

			x.PutInt64(m.SecretKeyId)
			x.PutInt64(m.MinId)
			x.PutInt64(m.MaxId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_createNewTestUser, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_createNewTestUser, layer)
	}
}

// Decode <--
func (m *TLUserCreateNewTestUser) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x4c6eccab: func() (err error) {
			m.SecretKeyId, err = d.Int64()
			m.MinId, err = d.Int64()
			m.MaxId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserEditCloseFriends <--
type TLUserEditCloseFriends struct {
	ClazzID uint32  `json:"_id"`
	UserId  int64   `json:"user_id"`
	Id      []int64 `json:"id"`
}

// Encode <--
func (m *TLUserEditCloseFriends) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x86247b05: func() error {
			x.PutClazzID(0x86247b05)

			x.PutInt64(m.UserId)

			iface.EncodeInt64List(x, m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_editCloseFriends, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_editCloseFriends, layer)
	}
}

// Decode <--
func (m *TLUserEditCloseFriends) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x86247b05: func() (err error) {
			m.UserId, err = d.Int64()

			m.Id, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserSetStoriesMaxId <--
type TLUserSetStoriesMaxId struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Id      int32  `json:"id"`
}

// Encode <--
func (m *TLUserSetStoriesMaxId) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x52f5b670: func() error {
			x.PutClazzID(0x52f5b670)

			x.PutInt64(m.UserId)
			x.PutInt32(m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_setStoriesMaxId, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_setStoriesMaxId, layer)
	}
}

// Decode <--
func (m *TLUserSetStoriesMaxId) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x52f5b670: func() (err error) {
			m.UserId, err = d.Int64()
			m.Id, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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

// Encode <--
func (m *TLUserSetColor) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x22fa0d77: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_setColor, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_setColor, layer)
	}
}

// Decode <--
func (m *TLUserSetColor) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x22fa0d77: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.UserId, err = d.Int64()
			if (flags & (1 << 1)) != 0 {
				m.ForProfile = true
			}
			m.Color, err = d.Int32()
			m.BackgroundEmojiId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserUpdateBirthday <--
type TLUserUpdateBirthday struct {
	ClazzID  uint32       `json:"_id"`
	UserId   int64        `json:"user_id"`
	Birthday *tg.Birthday `json:"birthday"`
}

// Encode <--
func (m *TLUserUpdateBirthday) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x587aab92: func() error {
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
				_ = m.Birthday.Encode(x, layer)
			}

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_updateBirthday, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_updateBirthday, layer)
	}
}

// Decode <--
func (m *TLUserUpdateBirthday) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x587aab92: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.UserId, err = d.Int64()
			if (flags & (1 << 1)) != 0 {
				m3 := &tg.Birthday{}
				_ = m3.Decode(d)
				m.Birthday = m3
			}

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetBirthdays <--
type TLUserGetBirthdays struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

// Encode <--
func (m *TLUserGetBirthdays) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xfe8ebfa6: func() error {
			x.PutClazzID(0xfe8ebfa6)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getBirthdays, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getBirthdays, layer)
	}
}

// Decode <--
func (m *TLUserGetBirthdays) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xfe8ebfa6: func() (err error) {
			m.UserId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserSetStoriesHidden <--
type TLUserSetStoriesHidden struct {
	ClazzID uint32   `json:"_id"`
	UserId  int64    `json:"user_id"`
	Id      int64    `json:"id"`
	Hidden  *tg.Bool `json:"hidden"`
}

// Encode <--
func (m *TLUserSetStoriesHidden) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xf7c61858: func() error {
			x.PutClazzID(0xf7c61858)

			x.PutInt64(m.UserId)
			x.PutInt64(m.Id)
			_ = m.Hidden.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_setStoriesHidden, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_setStoriesHidden, layer)
	}
}

// Decode <--
func (m *TLUserSetStoriesHidden) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xf7c61858: func() (err error) {
			m.UserId, err = d.Int64()
			m.Id, err = d.Int64()

			m3 := &tg.Bool{}
			_ = m3.Decode(d)
			m.Hidden = m3

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserUpdatePersonalChannel <--
type TLUserUpdatePersonalChannel struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	ChannelId int64  `json:"channel_id"`
}

// Encode <--
func (m *TLUserUpdatePersonalChannel) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc7f7bed0: func() error {
			x.PutClazzID(0xc7f7bed0)

			x.PutInt64(m.UserId)
			x.PutInt64(m.ChannelId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_updatePersonalChannel, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_updatePersonalChannel, layer)
	}
}

// Decode <--
func (m *TLUserUpdatePersonalChannel) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xc7f7bed0: func() (err error) {
			m.UserId, err = d.Int64()
			m.ChannelId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetUserIdByPhone <--
type TLUserGetUserIdByPhone struct {
	ClazzID uint32 `json:"_id"`
	Phone   string `json:"phone"`
}

// Encode <--
func (m *TLUserGetUserIdByPhone) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xfbab83c2: func() error {
			x.PutClazzID(0xfbab83c2)

			x.PutString(m.Phone)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getUserIdByPhone, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getUserIdByPhone, layer)
	}
}

// Decode <--
func (m *TLUserGetUserIdByPhone) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xfbab83c2: func() (err error) {
			m.Phone, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserSetAuthorizationTTL <--
type TLUserSetAuthorizationTTL struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
	Ttl     int32  `json:"ttl"`
}

// Encode <--
func (m *TLUserSetAuthorizationTTL) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xd621f3f0: func() error {
			x.PutClazzID(0xd621f3f0)

			x.PutInt64(m.UserId)
			x.PutInt32(m.Ttl)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_setAuthorizationTTL, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_setAuthorizationTTL, layer)
	}
}

// Decode <--
func (m *TLUserSetAuthorizationTTL) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xd621f3f0: func() (err error) {
			m.UserId, err = d.Int64()
			m.Ttl, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUserGetAuthorizationTTL <--
type TLUserGetAuthorizationTTL struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

// Encode <--
func (m *TLUserGetAuthorizationTTL) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xde6e493c: func() error {
			x.PutClazzID(0xde6e493c)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_user_getAuthorizationTTL, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_user_getAuthorizationTTL, layer)
	}
}

// Decode <--
func (m *TLUserGetAuthorizationTTL) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xde6e493c: func() (err error) {
			m.UserId, err = d.Int64()

			return nil
		},
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

// VectorLastSeenData <--
type VectorLastSeenData struct {
	Datas []*LastSeenData `json:"datas"`
}

// Encode <--
func (m *VectorLastSeenData) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorLastSeenData) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*LastSeenData](d)

	return err
}

// VectorImmutableUser <--
type VectorImmutableUser struct {
	Datas []*tg.ImmutableUser `json:"datas"`
}

// Encode <--
func (m *VectorImmutableUser) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorImmutableUser) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*tg.ImmutableUser](d)

	return err
}

// VectorPeerPeerNotifySettings <--
type VectorPeerPeerNotifySettings struct {
	Datas []*PeerPeerNotifySettings `json:"datas"`
}

// Encode <--
func (m *VectorPeerPeerNotifySettings) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorPeerPeerNotifySettings) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*PeerPeerNotifySettings](d)

	return err
}

// VectorPrivacyRule <--
type VectorPrivacyRule struct {
	Datas []*tg.PrivacyRule `json:"datas"`
}

// Encode <--
func (m *VectorPrivacyRule) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorPrivacyRule) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*tg.PrivacyRule](d)

	return err
}

// VectorLong <--
type VectorLong struct {
	Datas []int64 `json:"datas"`
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
	Datas []*tg.PeerBlocked `json:"datas"`
}

// Encode <--
func (m *VectorPeerBlocked) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorPeerBlocked) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*tg.PeerBlocked](d)

	return err
}

// VectorContactData <--
type VectorContactData struct {
	Datas []*tg.ContactData `json:"datas"`
}

// Encode <--
func (m *VectorContactData) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorContactData) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*tg.ContactData](d)

	return err
}

// VectorInputContact <--
type VectorInputContact struct {
	Datas []*tg.InputContact `json:"datas"`
}

// Encode <--
func (m *VectorInputContact) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorInputContact) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*tg.InputContact](d)

	return err
}

// VectorUserData <--
type VectorUserData struct {
	Datas []*tg.UserData `json:"datas"`
}

// Encode <--
func (m *VectorUserData) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorUserData) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*tg.UserData](d)

	return err
}

// VectorContactBirthday <--
type VectorContactBirthday struct {
	Datas []*tg.ContactBirthday `json:"datas"`
}

// Encode <--
func (m *VectorContactBirthday) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorContactBirthday) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*tg.ContactBirthday](d)

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
}
