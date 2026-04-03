/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package session

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

// HttpSessionDataClazz <--
//   - TL_HttpSessionData
type HttpSessionDataClazz = *TLHttpSessionData

func DecodeHttpSessionDataClazz(d *bin.Decoder) (HttpSessionDataClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0xdbd8534f:
		x := &TLHttpSessionData{ClazzID: id, ClazzName2: ClazzName_httpSessionData}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeHttpSessionData - unexpected clazzId: %d", id)
	}

}

// TLHttpSessionData <--
type TLHttpSessionData struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Payload    []byte `json:"payload"`
}

func MakeTLHttpSessionData(m *TLHttpSessionData) *TLHttpSessionData {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_httpSessionData

	return m
}

func (m *TLHttpSessionData) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLHttpSessionData) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("httpSessionData", m)
}

// HttpSessionDataClazzName <--
func (m *TLHttpSessionData) HttpSessionDataClazzName() string {
	return ClazzName_httpSessionData
}

// ClazzName <--
func (m *TLHttpSessionData) ClazzName() string {
	return m.ClazzName2
}

// ToHttpSessionData <--
func (m *TLHttpSessionData) ToHttpSessionData() *HttpSessionData {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLHttpSessionData) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_httpSessionData, int(layer)); clazzId {
	case 0xdbd8534f:
		size := 4
		size += iface.CalcBytesSize(m.Payload)

		return size
	default:
		return 0
	}
}

func (m *TLHttpSessionData) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_httpSessionData, int(layer)); clazzId {
	case 0xdbd8534f:
		if err := iface.ValidateRequiredBytes("payload", m.Payload); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_httpSessionData, layer)
	}
}

// Encode <--
func (m *TLHttpSessionData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_httpSessionData, int(layer)); clazzId {
	case 0xdbd8534f:
		x.PutClazzID(0xdbd8534f)

		x.PutBytes(m.Payload)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_httpSessionData, layer)
	}
}

// Decode <--
func (m *TLHttpSessionData) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xdbd8534f:
		m.Payload, err = d.Bytes()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// HttpSessionData <--
type HttpSessionData = TLHttpSessionData

// SessionClientDataClazz <--
//   - TL_SessionClientData
type SessionClientDataClazz = *TLSessionClientData

func DecodeSessionClientDataClazz(d *bin.Decoder) (SessionClientDataClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x41a20c4e:
		x := &TLSessionClientData{ClazzID: id, ClazzName2: ClazzName_sessionClientData}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeSessionClientData - unexpected clazzId: %d", id)
	}

}

// TLSessionClientData <--
type TLSessionClientData struct {
	ClazzID       uint32 `json:"_id"`
	ClazzName2    string `json:"_name"`
	ServerId      string `json:"server_id"`
	ConnType      int32  `json:"conn_type"`
	AuthKeyId     int64  `json:"auth_key_id"`
	KeyType       int32  `json:"key_type"`
	PermAuthKeyId int64  `json:"perm_auth_key_id"`
	SessionId     int64  `json:"session_id"`
	ClientIp      string `json:"client_ip"`
	QuickAck      int32  `json:"quick_ack"`
	Salt          int64  `json:"salt"`
	Payload       []byte `json:"payload"`
}

func MakeTLSessionClientData(m *TLSessionClientData) *TLSessionClientData {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_sessionClientData

	return m
}

func (m *TLSessionClientData) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLSessionClientData) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("sessionClientData", m)
}

// SessionClientDataClazzName <--
func (m *TLSessionClientData) SessionClientDataClazzName() string {
	return ClazzName_sessionClientData
}

// ClazzName <--
func (m *TLSessionClientData) ClazzName() string {
	return m.ClazzName2
}

// ToSessionClientData <--
func (m *TLSessionClientData) ToSessionClientData() *SessionClientData {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLSessionClientData) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_sessionClientData, int(layer)); clazzId {
	case 0x41a20c4e:
		size := 4
		size += iface.CalcStringSize(m.ServerId)
		size += 4
		size += 8
		size += 4
		size += 8
		size += 8
		size += iface.CalcStringSize(m.ClientIp)
		size += 4
		size += 8
		size += iface.CalcBytesSize(m.Payload)

		return size
	default:
		return 0
	}
}

func (m *TLSessionClientData) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_sessionClientData, int(layer)); clazzId {
	case 0x41a20c4e:
		if err := iface.ValidateRequiredString("server_id", m.ServerId); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("client_ip", m.ClientIp); err != nil {
			return err
		}

		if err := iface.ValidateRequiredBytes("payload", m.Payload); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_sessionClientData, layer)
	}
}

// Encode <--
func (m *TLSessionClientData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_sessionClientData, int(layer)); clazzId {
	case 0x41a20c4e:
		x.PutClazzID(0x41a20c4e)

		x.PutString(m.ServerId)
		x.PutInt32(m.ConnType)
		x.PutInt64(m.AuthKeyId)
		x.PutInt32(m.KeyType)
		x.PutInt64(m.PermAuthKeyId)
		x.PutInt64(m.SessionId)
		x.PutString(m.ClientIp)
		x.PutInt32(m.QuickAck)
		x.PutInt64(m.Salt)
		x.PutBytes(m.Payload)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_sessionClientData, layer)
	}
}

// Decode <--
func (m *TLSessionClientData) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x41a20c4e:
		m.ServerId, err = d.String()
		if err != nil {
			return err
		}
		m.ConnType, err = d.Int32()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.KeyType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PermAuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.SessionId, err = d.Int64()
		if err != nil {
			return err
		}
		m.ClientIp, err = d.String()
		if err != nil {
			return err
		}
		m.QuickAck, err = d.Int32()
		if err != nil {
			return err
		}
		m.Salt, err = d.Int64()
		if err != nil {
			return err
		}
		m.Payload, err = d.Bytes()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// SessionClientData <--
type SessionClientData = TLSessionClientData

// SessionClientEventClazz <--
//   - TL_SessionClientEvent
type SessionClientEventClazz = *TLSessionClientEvent

func DecodeSessionClientEventClazz(d *bin.Decoder) (SessionClientEventClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0xf17f375f:
		x := &TLSessionClientEvent{ClazzID: id, ClazzName2: ClazzName_sessionClientEvent}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeSessionClientEvent - unexpected clazzId: %d", id)
	}

}

// TLSessionClientEvent <--
type TLSessionClientEvent struct {
	ClazzID       uint32 `json:"_id"`
	ClazzName2    string `json:"_name"`
	ServerId      string `json:"server_id"`
	ConnType      int32  `json:"conn_type"`
	AuthKeyId     int64  `json:"auth_key_id"`
	KeyType       int32  `json:"key_type"`
	PermAuthKeyId int64  `json:"perm_auth_key_id"`
	SessionId     int64  `json:"session_id"`
	ClientIp      string `json:"client_ip"`
}

func MakeTLSessionClientEvent(m *TLSessionClientEvent) *TLSessionClientEvent {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_sessionClientEvent

	return m
}

func (m *TLSessionClientEvent) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLSessionClientEvent) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("sessionClientEvent", m)
}

// SessionClientEventClazzName <--
func (m *TLSessionClientEvent) SessionClientEventClazzName() string {
	return ClazzName_sessionClientEvent
}

// ClazzName <--
func (m *TLSessionClientEvent) ClazzName() string {
	return m.ClazzName2
}

// ToSessionClientEvent <--
func (m *TLSessionClientEvent) ToSessionClientEvent() *SessionClientEvent {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLSessionClientEvent) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_sessionClientEvent, int(layer)); clazzId {
	case 0xf17f375f:
		size := 4
		size += iface.CalcStringSize(m.ServerId)
		size += 4
		size += 8
		size += 4
		size += 8
		size += 8
		size += iface.CalcStringSize(m.ClientIp)

		return size
	default:
		return 0
	}
}

func (m *TLSessionClientEvent) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_sessionClientEvent, int(layer)); clazzId {
	case 0xf17f375f:
		if err := iface.ValidateRequiredString("server_id", m.ServerId); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("client_ip", m.ClientIp); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_sessionClientEvent, layer)
	}
}

// Encode <--
func (m *TLSessionClientEvent) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_sessionClientEvent, int(layer)); clazzId {
	case 0xf17f375f:
		x.PutClazzID(0xf17f375f)

		x.PutString(m.ServerId)
		x.PutInt32(m.ConnType)
		x.PutInt64(m.AuthKeyId)
		x.PutInt32(m.KeyType)
		x.PutInt64(m.PermAuthKeyId)
		x.PutInt64(m.SessionId)
		x.PutString(m.ClientIp)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_sessionClientEvent, layer)
	}
}

// Decode <--
func (m *TLSessionClientEvent) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xf17f375f:
		m.ServerId, err = d.String()
		if err != nil {
			return err
		}
		m.ConnType, err = d.Int32()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.KeyType, err = d.Int32()
		if err != nil {
			return err
		}
		m.PermAuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.SessionId, err = d.Int64()
		if err != nil {
			return err
		}
		m.ClientIp, err = d.String()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// SessionClientEvent <--
type SessionClientEvent = TLSessionClientEvent
