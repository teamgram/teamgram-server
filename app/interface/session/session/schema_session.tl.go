/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package session

import (
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
)

var _ iface.TLObject
var _ fmt.Stringer
var _ *tg.Bool
var _ bin.Fields

// HttpSessionDataClazz <--
//   - TL_HttpSessionData
type HttpSessionDataClazz interface {
	iface.TLObject
	HttpSessionDataClazzName() string
}

func DecodeHttpSessionDataClazz(d *bin.Decoder) (HttpSessionDataClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_httpSessionData:
		x := &TLHttpSessionData{ClazzID: id}
		_ = x.Decode(d)
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
	wrapper := iface.WithNameWrapper{"httpSessionData", m}
	return wrapper.String()
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

	return &HttpSessionData{Clazz: m}
}

// Encode <--
func (m *TLHttpSessionData) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xdbd8534f: func() error {
			x.PutClazzID(0xdbd8534f)

			x.PutBytes(m.Payload)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_httpSessionData, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_httpSessionData, layer)
	}
}

// Decode <--
func (m *TLHttpSessionData) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xdbd8534f: func() (err error) {
			m.Payload, err = d.Bytes()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// HttpSessionData <--
type HttpSessionData struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz HttpSessionDataClazz `json:"_clazz"`
}

func (m *HttpSessionData) String() string {
	wrapper := iface.WithNameWrapper{m.ClazzName(), m}
	return wrapper.String()
}

func (m *HttpSessionData) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.HttpSessionDataClazzName()
	}
}

// Encode <--
func (m *HttpSessionData) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("HttpSessionData - invalid Clazz")
}

// Decode <--
func (m *HttpSessionData) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeHttpSessionDataClazz(d)
	return
}

// Match <--
func (m *HttpSessionData) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLHttpSessionData:
		for _, v := range f {
			if f1, ok := v.(func(c *TLHttpSessionData) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToHttpSessionData <--
func (m *HttpSessionData) ToHttpSessionData() (*TLHttpSessionData, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLHttpSessionData); ok {
		return x, true
	}

	return nil, false
}

// SessionClientDataClazz <--
//   - TL_SessionClientData
type SessionClientDataClazz interface {
	iface.TLObject
	SessionClientDataClazzName() string
}

func DecodeSessionClientDataClazz(d *bin.Decoder) (SessionClientDataClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_sessionClientData:
		x := &TLSessionClientData{ClazzID: id}
		_ = x.Decode(d)
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
	wrapper := iface.WithNameWrapper{"sessionClientData", m}
	return wrapper.String()
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

	return &SessionClientData{Clazz: m}
}

// Encode <--
func (m *TLSessionClientData) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x41a20c4e: func() error {
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
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_sessionClientData, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_sessionClientData, layer)
	}
}

// Decode <--
func (m *TLSessionClientData) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x41a20c4e: func() (err error) {
			m.ServerId, err = d.String()
			m.ConnType, err = d.Int32()
			m.AuthKeyId, err = d.Int64()
			m.KeyType, err = d.Int32()
			m.PermAuthKeyId, err = d.Int64()
			m.SessionId, err = d.Int64()
			m.ClientIp, err = d.String()
			m.QuickAck, err = d.Int32()
			m.Salt, err = d.Int64()
			m.Payload, err = d.Bytes()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// SessionClientData <--
type SessionClientData struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz SessionClientDataClazz `json:"_clazz"`
}

func (m *SessionClientData) String() string {
	wrapper := iface.WithNameWrapper{m.ClazzName(), m}
	return wrapper.String()
}

func (m *SessionClientData) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.SessionClientDataClazzName()
	}
}

// Encode <--
func (m *SessionClientData) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("SessionClientData - invalid Clazz")
}

// Decode <--
func (m *SessionClientData) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeSessionClientDataClazz(d)
	return
}

// Match <--
func (m *SessionClientData) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLSessionClientData:
		for _, v := range f {
			if f1, ok := v.(func(c *TLSessionClientData) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToSessionClientData <--
func (m *SessionClientData) ToSessionClientData() (*TLSessionClientData, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLSessionClientData); ok {
		return x, true
	}

	return nil, false
}

// SessionClientEventClazz <--
//   - TL_SessionClientEvent
type SessionClientEventClazz interface {
	iface.TLObject
	SessionClientEventClazzName() string
}

func DecodeSessionClientEventClazz(d *bin.Decoder) (SessionClientEventClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_sessionClientEvent:
		x := &TLSessionClientEvent{ClazzID: id}
		_ = x.Decode(d)
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
	wrapper := iface.WithNameWrapper{"sessionClientEvent", m}
	return wrapper.String()
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

	return &SessionClientEvent{Clazz: m}
}

// Encode <--
func (m *TLSessionClientEvent) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xf17f375f: func() error {
			x.PutClazzID(0xf17f375f)

			x.PutString(m.ServerId)
			x.PutInt32(m.ConnType)
			x.PutInt64(m.AuthKeyId)
			x.PutInt32(m.KeyType)
			x.PutInt64(m.PermAuthKeyId)
			x.PutInt64(m.SessionId)
			x.PutString(m.ClientIp)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_sessionClientEvent, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_sessionClientEvent, layer)
	}
}

// Decode <--
func (m *TLSessionClientEvent) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xf17f375f: func() (err error) {
			m.ServerId, err = d.String()
			m.ConnType, err = d.Int32()
			m.AuthKeyId, err = d.Int64()
			m.KeyType, err = d.Int32()
			m.PermAuthKeyId, err = d.Int64()
			m.SessionId, err = d.Int64()
			m.ClientIp, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// SessionClientEvent <--
type SessionClientEvent struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz SessionClientEventClazz `json:"_clazz"`
}

func (m *SessionClientEvent) String() string {
	wrapper := iface.WithNameWrapper{m.ClazzName(), m}
	return wrapper.String()
}

func (m *SessionClientEvent) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.SessionClientEventClazzName()
	}
}

// Encode <--
func (m *SessionClientEvent) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("SessionClientEvent - invalid Clazz")
}

// Decode <--
func (m *SessionClientEvent) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeSessionClientEventClazz(d)
	return
}

// Match <--
func (m *SessionClientEvent) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLSessionClientEvent:
		for _, v := range f {
			if f1, ok := v.(func(c *TLSessionClientEvent) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToSessionClientEvent <--
func (m *SessionClientEvent) ToSessionClientEvent() (*TLSessionClientEvent, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLSessionClientEvent); ok {
		return x, true
	}

	return nil, false
}
