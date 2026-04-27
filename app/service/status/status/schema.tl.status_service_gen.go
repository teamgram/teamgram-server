/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package status

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

// TLStatusSetSessionOnline <--
type TLStatusSetSessionOnline struct {
	ClazzID uint32            `json:"_id"`
	UserId  int64             `json:"user_id"`
	Session SessionEntryClazz `json:"session"`
}

func (m *TLStatusSetSessionOnline) String() string {
	return iface.DebugStringWithName(ClazzName_status_setSessionOnline, m)
}

// Encode <--
func (m *TLStatusSetSessionOnline) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_status_setSessionOnline, int(layer)); clazzId {
	case 0x52518bcf:
		x.PutClazzID(0x52518bcf)

		x.PutInt64(m.UserId)
		if m.Session == nil {
			return fmt.Errorf("unable to encode status_setSessionOnline#0x52518bcf: field session is nil")
		}
		if err := m.Session.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode status_setSessionOnline#0x52518bcf: field session: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode status_setSessionOnline: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLStatusSetSessionOnline) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode status_setSessionOnline: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x52518bcf:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode status_setSessionOnline#0x52518bcf: field user_id: %w", err)
		}

		m.Session, err = DecodeSessionEntryClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode status_setSessionOnline#0x52518bcf: field session: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode status_setSessionOnline: invalid constructor %x", m.ClazzID)
	}
}

// TLStatusSetSessionOffline <--
type TLStatusSetSessionOffline struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

func (m *TLStatusSetSessionOffline) String() string {
	return iface.DebugStringWithName(ClazzName_status_setSessionOffline, m)
}

// Encode <--
func (m *TLStatusSetSessionOffline) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_status_setSessionOffline, int(layer)); clazzId {
	case 0x25a66a5c:
		x.PutClazzID(0x25a66a5c)

		x.PutInt64(m.UserId)
		x.PutInt64(m.AuthKeyId)

		return nil
	default:
		return fmt.Errorf("unable to encode status_setSessionOffline: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLStatusSetSessionOffline) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode status_setSessionOffline: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x25a66a5c:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode status_setSessionOffline#0x25a66a5c: field user_id: %w", err)
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode status_setSessionOffline#0x25a66a5c: field auth_key_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode status_setSessionOffline: invalid constructor %x", m.ClazzID)
	}
}

// TLStatusGetUserOnlineSessions <--
type TLStatusGetUserOnlineSessions struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLStatusGetUserOnlineSessions) String() string {
	return iface.DebugStringWithName(ClazzName_status_getUserOnlineSessions, m)
}

// Encode <--
func (m *TLStatusGetUserOnlineSessions) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_status_getUserOnlineSessions, int(layer)); clazzId {
	case 0xe7c0e5cd:
		x.PutClazzID(0xe7c0e5cd)

		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode status_getUserOnlineSessions: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLStatusGetUserOnlineSessions) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode status_getUserOnlineSessions: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xe7c0e5cd:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode status_getUserOnlineSessions#0xe7c0e5cd: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode status_getUserOnlineSessions: invalid constructor %x", m.ClazzID)
	}
}

// TLStatusGetUsersOnlineSessionsList <--
type TLStatusGetUsersOnlineSessionsList struct {
	ClazzID uint32  `json:"_id"`
	Users   []int64 `json:"users"`
}

func (m *TLStatusGetUsersOnlineSessionsList) String() string {
	return iface.DebugStringWithName(ClazzName_status_getUsersOnlineSessionsList, m)
}

// Encode <--
func (m *TLStatusGetUsersOnlineSessionsList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_status_getUsersOnlineSessionsList, int(layer)); clazzId {
	case 0x883b35c4:
		x.PutClazzID(0x883b35c4)

		iface.EncodeInt64List(x, m.Users)

		return nil
	default:
		return fmt.Errorf("unable to encode status_getUsersOnlineSessionsList: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLStatusGetUsersOnlineSessionsList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode status_getUsersOnlineSessionsList: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x883b35c4:

		m.Users, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode status_getUsersOnlineSessionsList#0x883b35c4: field users: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode status_getUsersOnlineSessionsList: invalid constructor %x", m.ClazzID)
	}
}

// TLStatusGetChannelOnlineUsers <--
type TLStatusGetChannelOnlineUsers struct {
	ClazzID   uint32 `json:"_id"`
	ChannelId int64  `json:"channel_id"`
}

func (m *TLStatusGetChannelOnlineUsers) String() string {
	return iface.DebugStringWithName(ClazzName_status_getChannelOnlineUsers, m)
}

// Encode <--
func (m *TLStatusGetChannelOnlineUsers) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_status_getChannelOnlineUsers, int(layer)); clazzId {
	case 0x4583ac55:
		x.PutClazzID(0x4583ac55)

		x.PutInt64(m.ChannelId)

		return nil
	default:
		return fmt.Errorf("unable to encode status_getChannelOnlineUsers: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLStatusGetChannelOnlineUsers) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode status_getChannelOnlineUsers: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x4583ac55:
		m.ChannelId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode status_getChannelOnlineUsers#0x4583ac55: field channel_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode status_getChannelOnlineUsers: invalid constructor %x", m.ClazzID)
	}
}

// TLStatusSetUserChannelsOnline <--
type TLStatusSetUserChannelsOnline struct {
	ClazzID  uint32  `json:"_id"`
	UserId   int64   `json:"user_id"`
	Channels []int64 `json:"channels"`
}

func (m *TLStatusSetUserChannelsOnline) String() string {
	return iface.DebugStringWithName(ClazzName_status_setUserChannelsOnline, m)
}

// Encode <--
func (m *TLStatusSetUserChannelsOnline) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_status_setUserChannelsOnline, int(layer)); clazzId {
	case 0xcd39044d:
		x.PutClazzID(0xcd39044d)

		x.PutInt64(m.UserId)

		iface.EncodeInt64List(x, m.Channels)

		return nil
	default:
		return fmt.Errorf("unable to encode status_setUserChannelsOnline: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLStatusSetUserChannelsOnline) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode status_setUserChannelsOnline: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xcd39044d:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode status_setUserChannelsOnline#0xcd39044d: field user_id: %w", err)
		}

		m.Channels, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode status_setUserChannelsOnline#0xcd39044d: field channels: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode status_setUserChannelsOnline: invalid constructor %x", m.ClazzID)
	}
}

// TLStatusSetUserChannelsOffline <--
type TLStatusSetUserChannelsOffline struct {
	ClazzID  uint32  `json:"_id"`
	UserId   int64   `json:"user_id"`
	Channels []int64 `json:"channels"`
}

func (m *TLStatusSetUserChannelsOffline) String() string {
	return iface.DebugStringWithName(ClazzName_status_setUserChannelsOffline, m)
}

// Encode <--
func (m *TLStatusSetUserChannelsOffline) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_status_setUserChannelsOffline, int(layer)); clazzId {
	case 0x6ca361aa:
		x.PutClazzID(0x6ca361aa)

		x.PutInt64(m.UserId)

		iface.EncodeInt64List(x, m.Channels)

		return nil
	default:
		return fmt.Errorf("unable to encode status_setUserChannelsOffline: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLStatusSetUserChannelsOffline) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode status_setUserChannelsOffline: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x6ca361aa:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode status_setUserChannelsOffline#0x6ca361aa: field user_id: %w", err)
		}

		m.Channels, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode status_setUserChannelsOffline#0x6ca361aa: field channels: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode status_setUserChannelsOffline: invalid constructor %x", m.ClazzID)
	}
}

// TLStatusSetChannelUserOffline <--
type TLStatusSetChannelUserOffline struct {
	ClazzID   uint32 `json:"_id"`
	ChannelId int64  `json:"channel_id"`
	UserId    int64  `json:"user_id"`
}

func (m *TLStatusSetChannelUserOffline) String() string {
	return iface.DebugStringWithName(ClazzName_status_setChannelUserOffline, m)
}

// Encode <--
func (m *TLStatusSetChannelUserOffline) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_status_setChannelUserOffline, int(layer)); clazzId {
	case 0xc48bcb7c:
		x.PutClazzID(0xc48bcb7c)

		x.PutInt64(m.ChannelId)
		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode status_setChannelUserOffline: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLStatusSetChannelUserOffline) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode status_setChannelUserOffline: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xc48bcb7c:
		m.ChannelId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode status_setChannelUserOffline#0xc48bcb7c: field channel_id: %w", err)
		}
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode status_setChannelUserOffline#0xc48bcb7c: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode status_setChannelUserOffline: invalid constructor %x", m.ClazzID)
	}
}

// TLStatusSetChannelUsersOnline <--
type TLStatusSetChannelUsersOnline struct {
	ClazzID   uint32  `json:"_id"`
	ChannelId int64   `json:"channel_id"`
	Id        []int64 `json:"id"`
}

func (m *TLStatusSetChannelUsersOnline) String() string {
	return iface.DebugStringWithName(ClazzName_status_setChannelUsersOnline, m)
}

// Encode <--
func (m *TLStatusSetChannelUsersOnline) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_status_setChannelUsersOnline, int(layer)); clazzId {
	case 0xa69bdcf7:
		x.PutClazzID(0xa69bdcf7)

		x.PutInt64(m.ChannelId)

		iface.EncodeInt64List(x, m.Id)

		return nil
	default:
		return fmt.Errorf("unable to encode status_setChannelUsersOnline: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLStatusSetChannelUsersOnline) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode status_setChannelUsersOnline: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xa69bdcf7:
		m.ChannelId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode status_setChannelUsersOnline#0xa69bdcf7: field channel_id: %w", err)
		}

		m.Id, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode status_setChannelUsersOnline#0xa69bdcf7: field id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode status_setChannelUsersOnline: invalid constructor %x", m.ClazzID)
	}
}

// TLStatusSetChannelOffline <--
type TLStatusSetChannelOffline struct {
	ClazzID   uint32 `json:"_id"`
	ChannelId int64  `json:"channel_id"`
}

func (m *TLStatusSetChannelOffline) String() string {
	return iface.DebugStringWithName(ClazzName_status_setChannelOffline, m)
}

// Encode <--
func (m *TLStatusSetChannelOffline) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_status_setChannelOffline, int(layer)); clazzId {
	case 0x4b7756f5:
		x.PutClazzID(0x4b7756f5)

		x.PutInt64(m.ChannelId)

		return nil
	default:
		return fmt.Errorf("unable to encode status_setChannelOffline: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLStatusSetChannelOffline) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode status_setChannelOffline: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x4b7756f5:
		m.ChannelId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode status_setChannelOffline#0x4b7756f5: field channel_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode status_setChannelOffline: invalid constructor %x", m.ClazzID)
	}
}

// Vector api result type
// ----------------------------------------------------------------------------
// VectorResList <--

// VectorUserSessionEntryList <--
type VectorUserSessionEntryList struct {
	Datas []UserSessionEntryListClazz `json:"_datas"`
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
	m.Datas, err = iface.DecodeObjectList[UserSessionEntryListClazz](d)

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
