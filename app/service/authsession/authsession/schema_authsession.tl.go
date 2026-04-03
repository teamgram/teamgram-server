/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package authsession

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

// AuthKeyStateDataClazz <--
//   - TL_AuthKeyStateData
type AuthKeyStateDataClazz = *TLAuthKeyStateData

func DecodeAuthKeyStateDataClazz(d *bin.Decoder) (AuthKeyStateDataClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0xe0408f17:
		x := &TLAuthKeyStateData{ClazzID: id, ClazzName2: ClazzName_authKeyStateData}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeAuthKeyStateData - unexpected clazzId: %d", id)
	}

}

// TLAuthKeyStateData <--
type TLAuthKeyStateData struct {
	ClazzID              uint32             `json:"_id"`
	ClazzName2           string             `json:"_name"`
	AuthKeyId            int64              `json:"auth_key_id"`
	KeyState             int32              `json:"key_state"`
	UserId               int64              `json:"user_id"`
	AccessHash           int64              `json:"access_hash"`
	Client               ClientSessionClazz `json:"client"`
	AndroidPushSessionId *int64             `json:"android_push_session_id"`
}

func MakeTLAuthKeyStateData(m *TLAuthKeyStateData) *TLAuthKeyStateData {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_authKeyStateData

	return m
}

func (m *TLAuthKeyStateData) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLAuthKeyStateData) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("authKeyStateData", m)
}

// AuthKeyStateDataClazzName <--
func (m *TLAuthKeyStateData) AuthKeyStateDataClazzName() string {
	return ClazzName_authKeyStateData
}

// ClazzName <--
func (m *TLAuthKeyStateData) ClazzName() string {
	return m.ClazzName2
}

// ToAuthKeyStateData <--
func (m *TLAuthKeyStateData) ToAuthKeyStateData() *AuthKeyStateData {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLAuthKeyStateData) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authKeyStateData, int(layer)); clazzId {
	case 0xe0408f17:
		size := 4
		size += 4
		size += 8
		size += 4
		size += 8
		size += 8
		if m.Client != nil {
			size += iface.CalcObjectSize(m.Client, layer)
		}

		if m.AndroidPushSessionId != nil {
			size += 8
		}

		return size
	default:
		return 0
	}
}

func (m *TLAuthKeyStateData) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authKeyStateData, int(layer)); clazzId {
	case 0xe0408f17:

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authKeyStateData, layer)
	}
}

// Encode <--
func (m *TLAuthKeyStateData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authKeyStateData, int(layer)); clazzId {
	case 0xe0408f17:
		x.PutClazzID(0xe0408f17)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.Client != nil {
				flags |= 1 << 0
			}
			if m.AndroidPushSessionId != nil {
				flags |= 1 << 1
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.AuthKeyId)
		x.PutInt32(m.KeyState)
		x.PutInt64(m.UserId)
		x.PutInt64(m.AccessHash)
		if m.Client != nil {
			_ = m.Client.Encode(x, layer)
		}

		if m.AndroidPushSessionId != nil {
			x.PutInt64(*m.AndroidPushSessionId)
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authKeyStateData, layer)
	}
}

// Decode <--
func (m *TLAuthKeyStateData) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xe0408f17:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.KeyState, err = d.Int32()
		if err != nil {
			return err
		}
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AccessHash, err = d.Int64()
		if err != nil {
			return err
		}
		if (flags & (1 << 0)) != 0 {
			m.Client, err = DecodeClientSessionClazz(d)
			if err != nil {
				return err
			}
		}
		if (flags & (1 << 1)) != 0 {
			m.AndroidPushSessionId = new(int64)
			*m.AndroidPushSessionId, err = d.Int64()
			if err != nil {
				return err
			}
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// AuthKeyStateData <--
type AuthKeyStateData = TLAuthKeyStateData

// ClientSessionClazz <--
//   - TL_ClientSession
type ClientSessionClazz = *TLClientSession

func DecodeClientSessionClazz(d *bin.Decoder) (ClientSessionClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x9a8e71b0:
		x := &TLClientSession{ClazzID: id, ClazzName2: ClazzName_clientSession}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeClientSession - unexpected clazzId: %d", id)
	}

}

// TLClientSession <--
type TLClientSession struct {
	ClazzID        uint32 `json:"_id"`
	ClazzName2     string `json:"_name"`
	AuthKeyId      int64  `json:"auth_key_id"`
	Ip             string `json:"ip"`
	Layer          int32  `json:"layer"`
	ApiId          int32  `json:"api_id"`
	DeviceModel    string `json:"device_model"`
	SystemVersion  string `json:"system_version"`
	AppVersion     string `json:"app_version"`
	SystemLangCode string `json:"system_lang_code"`
	LangPack       string `json:"lang_pack"`
	LangCode       string `json:"lang_code"`
	Proxy          string `json:"proxy"`
	Params         string `json:"params"`
}

func MakeTLClientSession(m *TLClientSession) *TLClientSession {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_clientSession

	return m
}

func (m *TLClientSession) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLClientSession) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("clientSession", m)
}

// ClientSessionClazzName <--
func (m *TLClientSession) ClientSessionClazzName() string {
	return ClazzName_clientSession
}

// ClazzName <--
func (m *TLClientSession) ClazzName() string {
	return m.ClazzName2
}

// ToClientSession <--
func (m *TLClientSession) ToClientSession() *ClientSession {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLClientSession) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_clientSession, int(layer)); clazzId {
	case 0x9a8e71b0:
		size := 4
		size += 8
		size += iface.CalcStringSize(m.Ip)
		size += 4
		size += 4
		size += iface.CalcStringSize(m.DeviceModel)
		size += iface.CalcStringSize(m.SystemVersion)
		size += iface.CalcStringSize(m.AppVersion)
		size += iface.CalcStringSize(m.SystemLangCode)
		size += iface.CalcStringSize(m.LangPack)
		size += iface.CalcStringSize(m.LangCode)
		size += iface.CalcStringSize(m.Proxy)
		size += iface.CalcStringSize(m.Params)

		return size
	default:
		return 0
	}
}

func (m *TLClientSession) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_clientSession, int(layer)); clazzId {
	case 0x9a8e71b0:
		if err := iface.ValidateRequiredString("ip", m.Ip); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("device_model", m.DeviceModel); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("system_version", m.SystemVersion); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("app_version", m.AppVersion); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("system_lang_code", m.SystemLangCode); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("lang_pack", m.LangPack); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("lang_code", m.LangCode); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("proxy", m.Proxy); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("params", m.Params); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_clientSession, layer)
	}
}

// Encode <--
func (m *TLClientSession) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_clientSession, int(layer)); clazzId {
	case 0x9a8e71b0:
		x.PutClazzID(0x9a8e71b0)

		x.PutInt64(m.AuthKeyId)
		x.PutString(m.Ip)
		x.PutInt32(m.Layer)
		x.PutInt32(m.ApiId)
		x.PutString(m.DeviceModel)
		x.PutString(m.SystemVersion)
		x.PutString(m.AppVersion)
		x.PutString(m.SystemLangCode)
		x.PutString(m.LangPack)
		x.PutString(m.LangCode)
		x.PutString(m.Proxy)
		x.PutString(m.Params)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_clientSession, layer)
	}
}

// Decode <--
func (m *TLClientSession) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x9a8e71b0:
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.Ip, err = d.String()
		if err != nil {
			return err
		}
		m.Layer, err = d.Int32()
		if err != nil {
			return err
		}
		m.ApiId, err = d.Int32()
		if err != nil {
			return err
		}
		m.DeviceModel, err = d.String()
		if err != nil {
			return err
		}
		m.SystemVersion, err = d.String()
		if err != nil {
			return err
		}
		m.AppVersion, err = d.String()
		if err != nil {
			return err
		}
		m.SystemLangCode, err = d.String()
		if err != nil {
			return err
		}
		m.LangPack, err = d.String()
		if err != nil {
			return err
		}
		m.LangCode, err = d.String()
		if err != nil {
			return err
		}
		m.Proxy, err = d.String()
		if err != nil {
			return err
		}
		m.Params, err = d.String()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// ClientSession <--
type ClientSession = TLClientSession
