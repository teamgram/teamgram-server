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

// OnlineSessionClazz <--
//   - TL_OnlineSession
type OnlineSessionClazz = *TLOnlineSession

func DecodeOnlineSessionClazz(d *bin.Decoder) (OnlineSessionClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode OnlineSession: constructor: %w", err)
	}

	switch id {
	case 0x8a390a9f:
		x := &TLOnlineSession{ClazzID: id, ClazzName2: ClazzName_onlineSession}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode OnlineSession: invalid constructor %x", id)
	}

}

// TLOnlineSession <--
type TLOnlineSession struct {
	ClazzID           uint32 `json:"_id"`
	ClazzName2        string `json:"_name"`
	UserId            int64  `json:"user_id"`
	PermAuthKeyId     int64  `json:"perm_auth_key_id"`
	AuthKeyId         int64  `json:"auth_key_id"`
	AuthKeyType       int32  `json:"auth_key_type"`
	SessionId         int64  `json:"session_id"`
	GatewayId         string `json:"gateway_id"`
	GatewayGeneration string `json:"gateway_generation"`
	GatewayRpcAddr    string `json:"gateway_rpc_addr"`
	Layer             int32  `json:"layer"`
	Client            string `json:"client"`
	UpdatedAt         int64  `json:"updated_at"`
	ExpiresAt         int64  `json:"expires_at"`
}

func MakeTLOnlineSession(m *TLOnlineSession) *TLOnlineSession {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_onlineSession

	return m
}

func (m *TLOnlineSession) String() string {
	return iface.DebugStringWithName("onlineSession", m)
}

func (m *TLOnlineSession) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("onlineSession", m)
}

// OnlineSessionClazzName <--
func (m *TLOnlineSession) OnlineSessionClazzName() string {
	return ClazzName_onlineSession
}

// ClazzName <--
func (m *TLOnlineSession) ClazzName() string {
	return m.ClazzName2
}

// ToOnlineSession <--
func (m *TLOnlineSession) ToOnlineSession() *OnlineSession {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLOnlineSession) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_onlineSession, int(layer)); clazzId {
	case 0x8a390a9f:
		size := 4
		size += 8
		size += 8
		size += 8
		size += 4
		size += 8
		size += iface.CalcStringSize(m.GatewayId)
		size += iface.CalcStringSize(m.GatewayGeneration)
		size += iface.CalcStringSize(m.GatewayRpcAddr)
		size += 4
		size += iface.CalcStringSize(m.Client)
		size += 8
		size += 8

		return size
	default:
		return 0
	}
}

func (m *TLOnlineSession) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_onlineSession, int(layer)); clazzId {
	case 0x8a390a9f:
		if err := iface.ValidateRequiredString("gateway_id", m.GatewayId); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("gateway_generation", m.GatewayGeneration); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("gateway_rpc_addr", m.GatewayRpcAddr); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("client", m.Client); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode onlineSession: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLOnlineSession) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_onlineSession, int(layer)); clazzId {
	case 0x8a390a9f:
		x.PutClazzID(0x8a390a9f)

		x.PutInt64(m.UserId)
		x.PutInt64(m.PermAuthKeyId)
		x.PutInt64(m.AuthKeyId)
		x.PutInt32(m.AuthKeyType)
		x.PutInt64(m.SessionId)
		x.PutString(m.GatewayId)
		x.PutString(m.GatewayGeneration)
		x.PutString(m.GatewayRpcAddr)
		x.PutInt32(m.Layer)
		x.PutString(m.Client)
		x.PutInt64(m.UpdatedAt)
		x.PutInt64(m.ExpiresAt)

		return nil
	default:
		return fmt.Errorf("unable to encode onlineSession: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLOnlineSession) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x8a390a9f:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode onlineSession#0x8a390a9f: field user_id: %w", err)
		}
		m.PermAuthKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode onlineSession#0x8a390a9f: field perm_auth_key_id: %w", err)
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode onlineSession#0x8a390a9f: field auth_key_id: %w", err)
		}
		m.AuthKeyType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode onlineSession#0x8a390a9f: field auth_key_type: %w", err)
		}
		m.SessionId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode onlineSession#0x8a390a9f: field session_id: %w", err)
		}
		m.GatewayId, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode onlineSession#0x8a390a9f: field gateway_id: %w", err)
		}
		m.GatewayGeneration, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode onlineSession#0x8a390a9f: field gateway_generation: %w", err)
		}
		m.GatewayRpcAddr, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode onlineSession#0x8a390a9f: field gateway_rpc_addr: %w", err)
		}
		m.Layer, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode onlineSession#0x8a390a9f: field layer: %w", err)
		}
		m.Client, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode onlineSession#0x8a390a9f: field client: %w", err)
		}
		m.UpdatedAt, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode onlineSession#0x8a390a9f: field updated_at: %w", err)
		}
		m.ExpiresAt, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode onlineSession#0x8a390a9f: field expires_at: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode onlineSession: invalid constructor %x", m.ClazzID)
	}
}

// OnlineSession <--
type OnlineSession = TLOnlineSession

// UserOnlineSessionsClazz <--
//   - TL_UserOnlineSessions
type UserOnlineSessionsClazz = *TLUserOnlineSessions

func DecodeUserOnlineSessionsClazz(d *bin.Decoder) (UserOnlineSessionsClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode UserOnlineSessions: constructor: %w", err)
	}

	switch id {
	case 0x11eacb03:
		x := &TLUserOnlineSessions{ClazzID: id, ClazzName2: ClazzName_userOnlineSessions}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode UserOnlineSessions: invalid constructor %x", id)
	}

}

// TLUserOnlineSessions <--
type TLUserOnlineSessions struct {
	ClazzID    uint32               `json:"_id"`
	ClazzName2 string               `json:"_name"`
	UserId     int64                `json:"user_id"`
	Sessions   []OnlineSessionClazz `json:"sessions"`
}

func MakeTLUserOnlineSessions(m *TLUserOnlineSessions) *TLUserOnlineSessions {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_userOnlineSessions

	return m
}

func (m *TLUserOnlineSessions) String() string {
	return iface.DebugStringWithName("userOnlineSessions", m)
}

func (m *TLUserOnlineSessions) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("userOnlineSessions", m)
}

// UserOnlineSessionsClazzName <--
func (m *TLUserOnlineSessions) UserOnlineSessionsClazzName() string {
	return ClazzName_userOnlineSessions
}

// ClazzName <--
func (m *TLUserOnlineSessions) ClazzName() string {
	return m.ClazzName2
}

// ToUserOnlineSessions <--
func (m *TLUserOnlineSessions) ToUserOnlineSessions() *UserOnlineSessions {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLUserOnlineSessions) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userOnlineSessions, int(layer)); clazzId {
	case 0x11eacb03:
		size := 4
		size += 8
		size += iface.CalcObjectListSize(m.Sessions, layer)

		return size
	default:
		return 0
	}
}

func (m *TLUserOnlineSessions) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userOnlineSessions, int(layer)); clazzId {
	case 0x11eacb03:
		if err := iface.ValidateRequiredSlice("sessions", m.Sessions); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode userOnlineSessions: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLUserOnlineSessions) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userOnlineSessions, int(layer)); clazzId {
	case 0x11eacb03:
		x.PutClazzID(0x11eacb03)

		x.PutInt64(m.UserId)

		if err := iface.EncodeObjectList(x, m.Sessions, layer); err != nil {
			return fmt.Errorf("unable to encode userOnlineSessions#0x11eacb03: field sessions: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode userOnlineSessions: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserOnlineSessions) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x11eacb03:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userOnlineSessions#0x11eacb03: field user_id: %w", err)
		}
		l1, err3 := d.VectorHeader()
		if err3 != nil {
			return fmt.Errorf("unable to decode userOnlineSessions#0x11eacb03: field sessions: %w", err3)
		}
		if l1 > bin.MaxVectorLen {
			return fmt.Errorf("unable to decode userOnlineSessions#0x11eacb03: field sessions: %w", &bin.InvalidLengthError{Type: "vector", Length: int(l1)})
		}
		prealloc1 := int(l1)
		if prealloc1 > bin.PreallocateLimit {
			prealloc1 = bin.PreallocateLimit
		}
		v1 := make([]OnlineSessionClazz, 0, prealloc1)
		for i := int32(0); i < l1; i++ {
			vv1, err3 := DecodeOnlineSessionClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode userOnlineSessions#0x11eacb03: field sessions: %w", err3)
			}
			v1 = append(v1, vv1)
		}
		m.Sessions = v1

		return nil
	default:
		return fmt.Errorf("unable to decode userOnlineSessions: invalid constructor %x", m.ClazzID)
	}
}

// UserOnlineSessions <--
type UserOnlineSessions = TLUserOnlineSessions
