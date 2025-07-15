/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package authsession

import (
	"encoding/json"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
)

var _ iface.TLObject
var _ fmt.Stringer
var _ *tg.Bool
var _ bin.Fields

// AuthKeyStateDataClazz <--
//   - TL_AuthKeyStateData
type AuthKeyStateDataClazz interface {
	iface.TLObject
	AuthKeyStateDataClazzName() string
}

func DecodeAuthKeyStateDataClazz(d *bin.Decoder) (AuthKeyStateDataClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_authKeyStateData:
		x := &TLAuthKeyStateData{ClazzID: id, ClazzName2: ClazzName_authKeyStateData}
		_ = x.Decode(d)
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

	return &AuthKeyStateData{Clazz: m}
}

// Encode <--
func (m *TLAuthKeyStateData) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xe0408f17: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authKeyStateData, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authKeyStateData, layer)
	}
}

// Decode <--
func (m *TLAuthKeyStateData) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xe0408f17: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.AuthKeyId, err = d.Int64()
			m.KeyState, err = d.Int32()
			m.UserId, err = d.Int64()
			m.AccessHash, err = d.Int64()
			if (flags & (1 << 0)) != 0 {
				// m5 := &ClientSession{}
				// _ = m5.Decode(d)
				// m.Client = m5
				m.Client, _ = DecodeClientSessionClazz(d)
			}
			if (flags & (1 << 1)) != 0 {
				m.AndroidPushSessionId = new(int64)
				*m.AndroidPushSessionId, err = d.Int64()
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

// AuthKeyStateData <--
type AuthKeyStateData struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz AuthKeyStateDataClazz `json:"_clazz"`
}

func (m *AuthKeyStateData) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *AuthKeyStateData) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.AuthKeyStateDataClazzName()
	}
}

// Encode <--
func (m *AuthKeyStateData) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("AuthKeyStateData - invalid Clazz")
}

// Decode <--
func (m *AuthKeyStateData) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeAuthKeyStateDataClazz(d)
	return
}

// Match <--
func (m *AuthKeyStateData) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLAuthKeyStateData:
		for _, v := range f {
			if f1, ok := v.(func(c *TLAuthKeyStateData) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToAuthKeyStateData <--
func (m *AuthKeyStateData) ToAuthKeyStateData() (*TLAuthKeyStateData, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLAuthKeyStateData); ok {
		return x, true
	}

	return nil, false
}

// ClientSessionClazz <--
//   - TL_ClientSession
type ClientSessionClazz interface {
	iface.TLObject
	ClientSessionClazzName() string
}

func DecodeClientSessionClazz(d *bin.Decoder) (ClientSessionClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_clientSession:
		x := &TLClientSession{ClazzID: id, ClazzName2: ClazzName_clientSession}
		_ = x.Decode(d)
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

	return &ClientSession{Clazz: m}
}

// Encode <--
func (m *TLClientSession) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x9a8e71b0: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_clientSession, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_clientSession, layer)
	}
}

// Decode <--
func (m *TLClientSession) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x9a8e71b0: func() (err error) {
			m.AuthKeyId, err = d.Int64()
			m.Ip, err = d.String()
			m.Layer, err = d.Int32()
			m.ApiId, err = d.Int32()
			m.DeviceModel, err = d.String()
			m.SystemVersion, err = d.String()
			m.AppVersion, err = d.String()
			m.SystemLangCode, err = d.String()
			m.LangPack, err = d.String()
			m.LangCode, err = d.String()
			m.Proxy, err = d.String()
			m.Params, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// ClientSession <--
type ClientSession struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz ClientSessionClazz `json:"_clazz"`
}

func (m *ClientSession) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *ClientSession) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.ClientSessionClazzName()
	}
}

// Encode <--
func (m *ClientSession) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("ClientSession - invalid Clazz")
}

// Decode <--
func (m *ClientSession) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeClientSessionClazz(d)
	return
}

// Match <--
func (m *ClientSession) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLClientSession:
		for _, v := range f {
			if f1, ok := v.(func(c *TLClientSession) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToClientSession <--
func (m *ClientSession) ToClientSession() (*TLClientSession, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLClientSession); ok {
		return x, true
	}

	return nil, false
}
