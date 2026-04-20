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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_status_setSessionOnline, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLStatusSetSessionOnline) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_status_setSessionOnline, int(layer)); clazzId {
	case 0x52518bcf:
		x.PutClazzID(0x52518bcf)

		x.PutInt64(m.UserId)
		_ = m.Session.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_setSessionOnline, layer)
	}
}

// Decode <--
func (m *TLStatusSetSessionOnline) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x52518bcf:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		m.Session, err = DecodeSessionEntryClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_status_setSessionOffline, TLObject: m}
	return wrapper.String()
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
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_setSessionOffline, layer)
	}
}

// Decode <--
func (m *TLStatusSetSessionOffline) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x25a66a5c:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLStatusGetUserOnlineSessions <--
type TLStatusGetUserOnlineSessions struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLStatusGetUserOnlineSessions) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_status_getUserOnlineSessions, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLStatusGetUserOnlineSessions) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_status_getUserOnlineSessions, int(layer)); clazzId {
	case 0xe7c0e5cd:
		x.PutClazzID(0xe7c0e5cd)

		x.PutInt64(m.UserId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_getUserOnlineSessions, layer)
	}
}

// Decode <--
func (m *TLStatusGetUserOnlineSessions) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xe7c0e5cd:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLStatusGetUsersOnlineSessionsList <--
type TLStatusGetUsersOnlineSessionsList struct {
	ClazzID uint32  `json:"_id"`
	Users   []int64 `json:"users"`
}

func (m *TLStatusGetUsersOnlineSessionsList) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_status_getUsersOnlineSessionsList, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLStatusGetUsersOnlineSessionsList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_status_getUsersOnlineSessionsList, int(layer)); clazzId {
	case 0x883b35c4:
		x.PutClazzID(0x883b35c4)

		iface.EncodeInt64List(x, m.Users)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_getUsersOnlineSessionsList, layer)
	}
}

// Decode <--
func (m *TLStatusGetUsersOnlineSessionsList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x883b35c4:

		m.Users, err = iface.DecodeInt64List(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLStatusGetChannelOnlineUsers <--
type TLStatusGetChannelOnlineUsers struct {
	ClazzID   uint32 `json:"_id"`
	ChannelId int64  `json:"channel_id"`
}

func (m *TLStatusGetChannelOnlineUsers) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_status_getChannelOnlineUsers, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLStatusGetChannelOnlineUsers) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_status_getChannelOnlineUsers, int(layer)); clazzId {
	case 0x4583ac55:
		x.PutClazzID(0x4583ac55)

		x.PutInt64(m.ChannelId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_getChannelOnlineUsers, layer)
	}
}

// Decode <--
func (m *TLStatusGetChannelOnlineUsers) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x4583ac55:
		m.ChannelId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_status_setUserChannelsOnline, TLObject: m}
	return wrapper.String()
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
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_setUserChannelsOnline, layer)
	}
}

// Decode <--
func (m *TLStatusSetUserChannelsOnline) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xcd39044d:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		m.Channels, err = iface.DecodeInt64List(d)
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_status_setUserChannelsOffline, TLObject: m}
	return wrapper.String()
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
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_setUserChannelsOffline, layer)
	}
}

// Decode <--
func (m *TLStatusSetUserChannelsOffline) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x6ca361aa:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		m.Channels, err = iface.DecodeInt64List(d)
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_status_setChannelUserOffline, TLObject: m}
	return wrapper.String()
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
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_setChannelUserOffline, layer)
	}
}

// Decode <--
func (m *TLStatusSetChannelUserOffline) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xc48bcb7c:
		m.ChannelId, err = d.Int64()
		if err != nil {
			return err
		}
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_status_setChannelUsersOnline, TLObject: m}
	return wrapper.String()
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
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_setChannelUsersOnline, layer)
	}
}

// Decode <--
func (m *TLStatusSetChannelUsersOnline) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xa69bdcf7:
		m.ChannelId, err = d.Int64()
		if err != nil {
			return err
		}

		m.Id, err = iface.DecodeInt64List(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLStatusSetChannelOffline <--
type TLStatusSetChannelOffline struct {
	ClazzID   uint32 `json:"_id"`
	ChannelId int64  `json:"channel_id"`
}

func (m *TLStatusSetChannelOffline) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_status_setChannelOffline, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLStatusSetChannelOffline) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_status_setChannelOffline, int(layer)); clazzId {
	case 0x4b7756f5:
		x.PutClazzID(0x4b7756f5)

		x.PutInt64(m.ChannelId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_status_setChannelOffline, layer)
	}
}

// Decode <--
func (m *TLStatusSetChannelOffline) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x4b7756f5:
		m.ChannelId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
