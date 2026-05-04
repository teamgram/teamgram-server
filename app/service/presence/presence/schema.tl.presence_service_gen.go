/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package presence

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

// TLPresenceSetSessionOnline <--
type TLPresenceSetSessionOnline struct {
	ClazzID uint32             `json:"_id"`
	Session OnlineSessionClazz `json:"session"`
}

func (m *TLPresenceSetSessionOnline) String() string {
	return iface.DebugStringWithName(ClazzName_presence_setSessionOnline, m)
}

// Encode <--
func (m *TLPresenceSetSessionOnline) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_presence_setSessionOnline, int(layer)); clazzId {
	case 0x75df4bac:
		x.PutClazzID(0x75df4bac)

		if m.Session == nil {
			return fmt.Errorf("unable to encode presence_setSessionOnline#0x75df4bac: field session is nil")
		}
		if err := m.Session.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode presence_setSessionOnline#0x75df4bac: field session: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode presence_setSessionOnline: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLPresenceSetSessionOnline) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode presence_setSessionOnline: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x75df4bac:

		m.Session, err = DecodeOnlineSessionClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode presence_setSessionOnline#0x75df4bac: field session: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode presence_setSessionOnline: invalid constructor %x", m.ClazzID)
	}
}

// TLPresenceSetSessionOffline <--
type TLPresenceSetSessionOffline struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	SessionId int64  `json:"session_id"`
}

func (m *TLPresenceSetSessionOffline) String() string {
	return iface.DebugStringWithName(ClazzName_presence_setSessionOffline, m)
}

// Encode <--
func (m *TLPresenceSetSessionOffline) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_presence_setSessionOffline, int(layer)); clazzId {
	case 0x71ed7afb:
		x.PutClazzID(0x71ed7afb)

		x.PutInt64(m.UserId)
		x.PutInt64(m.AuthKeyId)
		x.PutInt64(m.SessionId)

		return nil
	default:
		return fmt.Errorf("unable to encode presence_setSessionOffline: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLPresenceSetSessionOffline) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode presence_setSessionOffline: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x71ed7afb:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode presence_setSessionOffline#0x71ed7afb: field user_id: %w", err)
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode presence_setSessionOffline#0x71ed7afb: field auth_key_id: %w", err)
		}
		m.SessionId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode presence_setSessionOffline#0x71ed7afb: field session_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode presence_setSessionOffline: invalid constructor %x", m.ClazzID)
	}
}

// TLPresenceGetUserOnlineSessions <--
type TLPresenceGetUserOnlineSessions struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLPresenceGetUserOnlineSessions) String() string {
	return iface.DebugStringWithName(ClazzName_presence_getUserOnlineSessions, m)
}

// Encode <--
func (m *TLPresenceGetUserOnlineSessions) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_presence_getUserOnlineSessions, int(layer)); clazzId {
	case 0x5aede88d:
		x.PutClazzID(0x5aede88d)

		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode presence_getUserOnlineSessions: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLPresenceGetUserOnlineSessions) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode presence_getUserOnlineSessions: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x5aede88d:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode presence_getUserOnlineSessions#0x5aede88d: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode presence_getUserOnlineSessions: invalid constructor %x", m.ClazzID)
	}
}

// TLPresenceGetUsersOnlineSessions <--
type TLPresenceGetUsersOnlineSessions struct {
	ClazzID uint32  `json:"_id"`
	Users   []int64 `json:"users"`
}

func (m *TLPresenceGetUsersOnlineSessions) String() string {
	return iface.DebugStringWithName(ClazzName_presence_getUsersOnlineSessions, m)
}

// Encode <--
func (m *TLPresenceGetUsersOnlineSessions) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_presence_getUsersOnlineSessions, int(layer)); clazzId {
	case 0x1e5a09c5:
		x.PutClazzID(0x1e5a09c5)

		iface.EncodeInt64List(x, m.Users)

		return nil
	default:
		return fmt.Errorf("unable to encode presence_getUsersOnlineSessions: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLPresenceGetUsersOnlineSessions) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode presence_getUsersOnlineSessions: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x1e5a09c5:

		m.Users, err = iface.DecodeInt64List(d)
		if err != nil {
			return fmt.Errorf("unable to decode presence_getUsersOnlineSessions#0x1e5a09c5: field users: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode presence_getUsersOnlineSessions: invalid constructor %x", m.ClazzID)
	}
}

// TLPresenceGetGatewaySessions <--
type TLPresenceGetGatewaySessions struct {
	ClazzID   uint32 `json:"_id"`
	GatewayId string `json:"gateway_id"`
}

func (m *TLPresenceGetGatewaySessions) String() string {
	return iface.DebugStringWithName(ClazzName_presence_getGatewaySessions, m)
}

// Encode <--
func (m *TLPresenceGetGatewaySessions) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_presence_getGatewaySessions, int(layer)); clazzId {
	case 0xe5c814c:
		x.PutClazzID(0xe5c814c)

		x.PutString(m.GatewayId)

		return nil
	default:
		return fmt.Errorf("unable to encode presence_getGatewaySessions: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLPresenceGetGatewaySessions) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode presence_getGatewaySessions: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xe5c814c:
		m.GatewayId, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode presence_getGatewaySessions#0xe5c814c: field gateway_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode presence_getGatewaySessions: invalid constructor %x", m.ClazzID)
	}
}

// Vector api result type
// ----------------------------------------------------------------------------
// VectorResList <--

// VectorUserOnlineSessions <--
type VectorUserOnlineSessions struct {
	Datas []UserOnlineSessionsClazz `json:"_datas"`
}

func (m *VectorUserOnlineSessions) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorUserOnlineSessions) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorUserOnlineSessions) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[UserOnlineSessionsClazz](d)

	return err
}

// VectorOnlineSession <--
type VectorOnlineSession struct {
	Datas []OnlineSessionClazz `json:"_datas"`
}

func (m *VectorOnlineSession) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorOnlineSession) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorOnlineSession) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[OnlineSessionClazz](d)

	return err
}

// ----------------------------------------------------------------------------
// rpc

type RPCPresence interface {
	PresenceSetSessionOnline(ctx context.Context, in *TLPresenceSetSessionOnline) (*tg.Bool, error)
	PresenceSetSessionOffline(ctx context.Context, in *TLPresenceSetSessionOffline) (*tg.Bool, error)
	PresenceGetUserOnlineSessions(ctx context.Context, in *TLPresenceGetUserOnlineSessions) (*UserOnlineSessions, error)
	PresenceGetUsersOnlineSessions(ctx context.Context, in *TLPresenceGetUsersOnlineSessions) (*VectorUserOnlineSessions, error)
	PresenceGetGatewaySessions(ctx context.Context, in *TLPresenceGetGatewaySessions) (*VectorOnlineSession, error)
}
