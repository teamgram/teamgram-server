/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package status

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

// TLStatusSetSessionOnline <--
type TLStatusSetSessionOnline struct {
	ClazzID uint32        `json:"_id"`
	UserId  int64         `json:"user_id"`
	Session *SessionEntry `json:"session"`
}

func (m *TLStatusSetSessionOnline) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLStatusSetSessionOnline) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x52518bcf: func() error {
			x.PutClazzID(0x52518bcf)

			x.PutInt64(m.UserId)
			_ = m.Session.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_status_setSessionOnline, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_setSessionOnline, layer)
	}
}

// Decode <--
func (m *TLStatusSetSessionOnline) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x52518bcf: func() (err error) {
			m.UserId, err = d.Int64()

			m2 := &SessionEntry{}
			_ = m2.Decode(d)
			m.Session = m2

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

// TLStatusSetSessionOffline <--
type TLStatusSetSessionOffline struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

func (m *TLStatusSetSessionOffline) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLStatusSetSessionOffline) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x25a66a5c: func() error {
			x.PutClazzID(0x25a66a5c)

			x.PutInt64(m.UserId)
			x.PutInt64(m.AuthKeyId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_status_setSessionOffline, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_setSessionOffline, layer)
	}
}

// Decode <--
func (m *TLStatusSetSessionOffline) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x25a66a5c: func() (err error) {
			m.UserId, err = d.Int64()
			m.AuthKeyId, err = d.Int64()

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

// TLStatusGetUserOnlineSessions <--
type TLStatusGetUserOnlineSessions struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLStatusGetUserOnlineSessions) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLStatusGetUserOnlineSessions) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xe7c0e5cd: func() error {
			x.PutClazzID(0xe7c0e5cd)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_status_getUserOnlineSessions, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_getUserOnlineSessions, layer)
	}
}

// Decode <--
func (m *TLStatusGetUserOnlineSessions) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xe7c0e5cd: func() (err error) {
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

// TLStatusGetUsersOnlineSessionsList <--
type TLStatusGetUsersOnlineSessionsList struct {
	ClazzID uint32  `json:"_id"`
	Users   []int64 `json:"users"`
}

func (m *TLStatusGetUsersOnlineSessionsList) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLStatusGetUsersOnlineSessionsList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x883b35c4: func() error {
			x.PutClazzID(0x883b35c4)

			iface.EncodeInt64List(x, m.Users)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_status_getUsersOnlineSessionsList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_getUsersOnlineSessionsList, layer)
	}
}

// Decode <--
func (m *TLStatusGetUsersOnlineSessionsList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x883b35c4: func() (err error) {

			m.Users, err = iface.DecodeInt64List(d)

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

// TLStatusGetChannelOnlineUsers <--
type TLStatusGetChannelOnlineUsers struct {
	ClazzID   uint32 `json:"_id"`
	ChannelId int64  `json:"channel_id"`
}

func (m *TLStatusGetChannelOnlineUsers) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLStatusGetChannelOnlineUsers) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x4583ac55: func() error {
			x.PutClazzID(0x4583ac55)

			x.PutInt64(m.ChannelId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_status_getChannelOnlineUsers, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_getChannelOnlineUsers, layer)
	}
}

// Decode <--
func (m *TLStatusGetChannelOnlineUsers) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x4583ac55: func() (err error) {
			m.ChannelId, err = d.Int64()

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

// TLStatusSetUserChannelsOnline <--
type TLStatusSetUserChannelsOnline struct {
	ClazzID  uint32  `json:"_id"`
	UserId   int64   `json:"user_id"`
	Channels []int64 `json:"channels"`
}

func (m *TLStatusSetUserChannelsOnline) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLStatusSetUserChannelsOnline) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xcd39044d: func() error {
			x.PutClazzID(0xcd39044d)

			x.PutInt64(m.UserId)

			iface.EncodeInt64List(x, m.Channels)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_status_setUserChannelsOnline, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_setUserChannelsOnline, layer)
	}
}

// Decode <--
func (m *TLStatusSetUserChannelsOnline) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xcd39044d: func() (err error) {
			m.UserId, err = d.Int64()

			m.Channels, err = iface.DecodeInt64List(d)

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

// TLStatusSetUserChannelsOffline <--
type TLStatusSetUserChannelsOffline struct {
	ClazzID  uint32  `json:"_id"`
	UserId   int64   `json:"user_id"`
	Channels []int64 `json:"channels"`
}

func (m *TLStatusSetUserChannelsOffline) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLStatusSetUserChannelsOffline) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x6ca361aa: func() error {
			x.PutClazzID(0x6ca361aa)

			x.PutInt64(m.UserId)

			iface.EncodeInt64List(x, m.Channels)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_status_setUserChannelsOffline, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_setUserChannelsOffline, layer)
	}
}

// Decode <--
func (m *TLStatusSetUserChannelsOffline) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x6ca361aa: func() (err error) {
			m.UserId, err = d.Int64()

			m.Channels, err = iface.DecodeInt64List(d)

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

// TLStatusSetChannelUserOffline <--
type TLStatusSetChannelUserOffline struct {
	ClazzID   uint32 `json:"_id"`
	ChannelId int64  `json:"channel_id"`
	UserId    int64  `json:"user_id"`
}

func (m *TLStatusSetChannelUserOffline) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLStatusSetChannelUserOffline) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc48bcb7c: func() error {
			x.PutClazzID(0xc48bcb7c)

			x.PutInt64(m.ChannelId)
			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_status_setChannelUserOffline, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_setChannelUserOffline, layer)
	}
}

// Decode <--
func (m *TLStatusSetChannelUserOffline) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xc48bcb7c: func() (err error) {
			m.ChannelId, err = d.Int64()
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

// TLStatusSetChannelUsersOnline <--
type TLStatusSetChannelUsersOnline struct {
	ClazzID   uint32  `json:"_id"`
	ChannelId int64   `json:"channel_id"`
	Id        []int64 `json:"id"`
}

func (m *TLStatusSetChannelUsersOnline) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLStatusSetChannelUsersOnline) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa69bdcf7: func() error {
			x.PutClazzID(0xa69bdcf7)

			x.PutInt64(m.ChannelId)

			iface.EncodeInt64List(x, m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_status_setChannelUsersOnline, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_setChannelUsersOnline, layer)
	}
}

// Decode <--
func (m *TLStatusSetChannelUsersOnline) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa69bdcf7: func() (err error) {
			m.ChannelId, err = d.Int64()

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

// TLStatusSetChannelOffline <--
type TLStatusSetChannelOffline struct {
	ClazzID   uint32 `json:"_id"`
	ChannelId int64  `json:"channel_id"`
}

func (m *TLStatusSetChannelOffline) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLStatusSetChannelOffline) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x4b7756f5: func() error {
			x.PutClazzID(0x4b7756f5)

			x.PutInt64(m.ChannelId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_status_setChannelOffline, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_setChannelOffline, layer)
	}
}

// Decode <--
func (m *TLStatusSetChannelOffline) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x4b7756f5: func() (err error) {
			m.ChannelId, err = d.Int64()

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

// VectorUserSessionEntryList <--
type VectorUserSessionEntryList struct {
	Datas []*UserSessionEntryList `json:"_datas"`
}

func (m *VectorUserSessionEntryList) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorUserSessionEntryList) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorUserSessionEntryList) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*UserSessionEntryList](d)

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

type RPCStatus interface {
	StatusSetSessionOnline(ctx context.Context, in *TLStatusSetSessionOnline) (*tg.Bool, error)
	StatusSetSessionOffline(ctx context.Context, in *TLStatusSetSessionOffline) (*tg.Bool, error)
	StatusGetUserOnlineSessions(ctx context.Context, in *TLStatusGetUserOnlineSessions) (*UserSessionEntryList, error)
	StatusGetUsersOnlineSessionsList(ctx context.Context, in *TLStatusGetUsersOnlineSessionsList) (*VectorUserSessionEntryList, error)
	StatusGetChannelOnlineUsers(ctx context.Context, in *TLStatusGetChannelOnlineUsers) (*VectorLong, error)
	StatusSetUserChannelsOnline(ctx context.Context, in *TLStatusSetUserChannelsOnline) (*tg.Bool, error)
	StatusSetUserChannelsOffline(ctx context.Context, in *TLStatusSetUserChannelsOffline) (*tg.Bool, error)
	StatusSetChannelUserOffline(ctx context.Context, in *TLStatusSetChannelUserOffline) (*tg.Bool, error)
	StatusSetChannelUsersOnline(ctx context.Context, in *TLStatusSetChannelUsersOnline) (*tg.Bool, error)
	StatusSetChannelOffline(ctx context.Context, in *TLStatusSetChannelOffline) (*tg.Bool, error)
}
