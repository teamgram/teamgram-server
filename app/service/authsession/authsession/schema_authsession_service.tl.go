/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package authsession

import (
	"context"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/mt"
	"github.com/teamgram/proto/v2/tg"
)

var _ iface.TLObject
var _ fmt.Stringer
var _ *tg.Bool
var _ bin.Fields

// TLAuthsessionGetAuthorizations <--
type TLAuthsessionGetAuthorizations struct {
	ClazzID          uint32 `json:"_id"`
	UserId           int64  `json:"user_id"`
	ExcludeAuthKeyId int64  `json:"exclude_auth_keyId"`
}

// Encode <--
func (m *TLAuthsessionGetAuthorizations) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x30e21244: func() error {
			x.PutClazzID(0x30e21244)

			x.PutInt64(m.UserId)
			x.PutInt64(m.ExcludeAuthKeyId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_getAuthorizations, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getAuthorizations, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetAuthorizations) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x30e21244: func() (err error) {
			m.UserId, err = d.Int64()
			m.ExcludeAuthKeyId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionResetAuthorization <--
type TLAuthsessionResetAuthorization struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	Hash      int64  `json:"hash"`
}

// Encode <--
func (m *TLAuthsessionResetAuthorization) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x8d5f6ca6: func() error {
			x.PutClazzID(0x8d5f6ca6)

			x.PutInt64(m.UserId)
			x.PutInt64(m.AuthKeyId)
			x.PutInt64(m.Hash)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_resetAuthorization, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_resetAuthorization, layer)
	}
}

// Decode <--
func (m *TLAuthsessionResetAuthorization) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x8d5f6ca6: func() (err error) {
			m.UserId, err = d.Int64()
			m.AuthKeyId, err = d.Int64()
			m.Hash, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionGetLayer <--
type TLAuthsessionGetLayer struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

// Encode <--
func (m *TLAuthsessionGetLayer) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa82f16a9: func() error {
			x.PutClazzID(0xa82f16a9)

			x.PutInt64(m.AuthKeyId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_getLayer, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getLayer, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetLayer) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa82f16a9: func() (err error) {
			m.AuthKeyId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionGetLangPack <--
type TLAuthsessionGetLangPack struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

// Encode <--
func (m *TLAuthsessionGetLangPack) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x29bbc166: func() error {
			x.PutClazzID(0x29bbc166)

			x.PutInt64(m.AuthKeyId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_getLangPack, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getLangPack, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetLangPack) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x29bbc166: func() (err error) {
			m.AuthKeyId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionGetClient <--
type TLAuthsessionGetClient struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

// Encode <--
func (m *TLAuthsessionGetClient) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x605855be: func() error {
			x.PutClazzID(0x605855be)

			x.PutInt64(m.AuthKeyId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_getClient, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getClient, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetClient) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x605855be: func() (err error) {
			m.AuthKeyId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionGetLangCode <--
type TLAuthsessionGetLangCode struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

// Encode <--
func (m *TLAuthsessionGetLangCode) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x5899b559: func() error {
			x.PutClazzID(0x5899b559)

			x.PutInt64(m.AuthKeyId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_getLangCode, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getLangCode, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetLangCode) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x5899b559: func() (err error) {
			m.AuthKeyId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionGetUserId <--
type TLAuthsessionGetUserId struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

// Encode <--
func (m *TLAuthsessionGetUserId) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x57491cac: func() error {
			x.PutClazzID(0x57491cac)

			x.PutInt64(m.AuthKeyId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_getUserId, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getUserId, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetUserId) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x57491cac: func() (err error) {
			m.AuthKeyId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionGetPushSessionId <--
type TLAuthsessionGetPushSessionId struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	TokenType int32  `json:"token_type"`
}

// Encode <--
func (m *TLAuthsessionGetPushSessionId) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xb3c23141: func() error {
			x.PutClazzID(0xb3c23141)

			x.PutInt64(m.UserId)
			x.PutInt64(m.AuthKeyId)
			x.PutInt32(m.TokenType)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_getPushSessionId, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getPushSessionId, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetPushSessionId) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xb3c23141: func() (err error) {
			m.UserId, err = d.Int64()
			m.AuthKeyId, err = d.Int64()
			m.TokenType, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionGetFutureSalts <--
type TLAuthsessionGetFutureSalts struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	Num       int32  `json:"num"`
}

// Encode <--
func (m *TLAuthsessionGetFutureSalts) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xb8cf5815: func() error {
			x.PutClazzID(0xb8cf5815)

			x.PutInt64(m.AuthKeyId)
			x.PutInt32(m.Num)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_getFutureSalts, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getFutureSalts, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetFutureSalts) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xb8cf5815: func() (err error) {
			m.AuthKeyId, err = d.Int64()
			m.Num, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionQueryAuthKey <--
type TLAuthsessionQueryAuthKey struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

// Encode <--
func (m *TLAuthsessionQueryAuthKey) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x54b73828: func() error {
			x.PutClazzID(0x54b73828)

			x.PutInt64(m.AuthKeyId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_queryAuthKey, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_queryAuthKey, layer)
	}
}

// Decode <--
func (m *TLAuthsessionQueryAuthKey) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x54b73828: func() (err error) {
			m.AuthKeyId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionSetAuthKey <--
type TLAuthsessionSetAuthKey struct {
	ClazzID    uint32          `json:"_id"`
	AuthKey    *tg.AuthKeyInfo `json:"auth_key"`
	FutureSalt *mt.FutureSalt  `json:"future_salt"`
	ExpiresIn  int32           `json:"expires_in"`
}

// Encode <--
func (m *TLAuthsessionSetAuthKey) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x3e940c91: func() error {
			x.PutClazzID(0x3e940c91)

			_ = m.AuthKey.Encode(x, layer)
			_ = m.FutureSalt.Encode(x, layer)
			x.PutInt32(m.ExpiresIn)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_setAuthKey, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_setAuthKey, layer)
	}
}

// Decode <--
func (m *TLAuthsessionSetAuthKey) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x3e940c91: func() (err error) {

			m1 := &tg.AuthKeyInfo{}
			_ = m1.Decode(d)
			m.AuthKey = m1

			m2 := &mt.FutureSalt{}
			_ = m2.Decode(d)
			m.FutureSalt = m2

			m.ExpiresIn, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionBindAuthKeyUser <--
type TLAuthsessionBindAuthKeyUser struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	UserId    int64  `json:"user_id"`
}

// Encode <--
func (m *TLAuthsessionBindAuthKeyUser) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xbce0423: func() error {
			x.PutClazzID(0xbce0423)

			x.PutInt64(m.AuthKeyId)
			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_bindAuthKeyUser, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_bindAuthKeyUser, layer)
	}
}

// Decode <--
func (m *TLAuthsessionBindAuthKeyUser) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xbce0423: func() (err error) {
			m.AuthKeyId, err = d.Int64()
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

// TLAuthsessionUnbindAuthKeyUser <--
type TLAuthsessionUnbindAuthKeyUser struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	UserId    int64  `json:"user_id"`
}

// Encode <--
func (m *TLAuthsessionUnbindAuthKeyUser) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x758c648: func() error {
			x.PutClazzID(0x758c648)

			x.PutInt64(m.AuthKeyId)
			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_unbindAuthKeyUser, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_unbindAuthKeyUser, layer)
	}
}

// Decode <--
func (m *TLAuthsessionUnbindAuthKeyUser) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x758c648: func() (err error) {
			m.AuthKeyId, err = d.Int64()
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

// TLAuthsessionGetPermAuthKeyId <--
type TLAuthsessionGetPermAuthKeyId struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

// Encode <--
func (m *TLAuthsessionGetPermAuthKeyId) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x907464d6: func() error {
			x.PutClazzID(0x907464d6)

			x.PutInt64(m.AuthKeyId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_getPermAuthKeyId, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getPermAuthKeyId, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetPermAuthKeyId) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x907464d6: func() (err error) {
			m.AuthKeyId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionBindTempAuthKey <--
type TLAuthsessionBindTempAuthKey struct {
	ClazzID          uint32 `json:"_id"`
	PermAuthKeyId    int64  `json:"perm_auth_key_id"`
	Nonce            int64  `json:"nonce"`
	ExpiresAt        int32  `json:"expires_at"`
	EncryptedMessage []byte `json:"encrypted_message"`
}

// Encode <--
func (m *TLAuthsessionBindTempAuthKey) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x608f4f86: func() error {
			x.PutClazzID(0x608f4f86)

			x.PutInt64(m.PermAuthKeyId)
			x.PutInt64(m.Nonce)
			x.PutInt32(m.ExpiresAt)
			x.PutBytes(m.EncryptedMessage)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_bindTempAuthKey, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_bindTempAuthKey, layer)
	}
}

// Decode <--
func (m *TLAuthsessionBindTempAuthKey) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x608f4f86: func() (err error) {
			m.PermAuthKeyId, err = d.Int64()
			m.Nonce, err = d.Int64()
			m.ExpiresAt, err = d.Int32()
			m.EncryptedMessage, err = d.Bytes()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionSetClientSessionInfo <--
type TLAuthsessionSetClientSessionInfo struct {
	ClazzID uint32         `json:"_id"`
	Data    *ClientSession `json:"data"`
}

// Encode <--
func (m *TLAuthsessionSetClientSessionInfo) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x2d9ff94: func() error {
			x.PutClazzID(0x2d9ff94)

			_ = m.Data.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_setClientSessionInfo, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_setClientSessionInfo, layer)
	}
}

// Decode <--
func (m *TLAuthsessionSetClientSessionInfo) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x2d9ff94: func() (err error) {

			m1 := &ClientSession{}
			_ = m1.Decode(d)
			m.Data = m1

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionGetAuthorization <--
type TLAuthsessionGetAuthorization struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

// Encode <--
func (m *TLAuthsessionGetAuthorization) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x6e5e1923: func() error {
			x.PutClazzID(0x6e5e1923)

			x.PutInt64(m.AuthKeyId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_getAuthorization, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getAuthorization, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetAuthorization) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x6e5e1923: func() (err error) {
			m.AuthKeyId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionGetAuthStateData <--
type TLAuthsessionGetAuthStateData struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

// Encode <--
func (m *TLAuthsessionGetAuthStateData) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x4f5e3131: func() error {
			x.PutClazzID(0x4f5e3131)

			x.PutInt64(m.AuthKeyId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_getAuthStateData, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getAuthStateData, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetAuthStateData) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x4f5e3131: func() (err error) {
			m.AuthKeyId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionSetLayer <--
type TLAuthsessionSetLayer struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	Ip        string `json:"ip"`
	Layer     int32  `json:"layer"`
}

// Encode <--
func (m *TLAuthsessionSetLayer) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x44651485: func() error {
			x.PutClazzID(0x44651485)

			x.PutInt64(m.AuthKeyId)
			x.PutString(m.Ip)
			x.PutInt32(m.Layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_setLayer, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_setLayer, layer)
	}
}

// Decode <--
func (m *TLAuthsessionSetLayer) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x44651485: func() (err error) {
			m.AuthKeyId, err = d.Int64()
			m.Ip, err = d.String()
			m.Layer, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionSetInitConnection <--
type TLAuthsessionSetInitConnection struct {
	ClazzID        uint32 `json:"_id"`
	AuthKeyId      int64  `json:"auth_key_id"`
	Ip             string `json:"ip"`
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

// Encode <--
func (m *TLAuthsessionSetInitConnection) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x7cdf8a8c: func() error {
			x.PutClazzID(0x7cdf8a8c)

			x.PutInt64(m.AuthKeyId)
			x.PutString(m.Ip)
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

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_setInitConnection, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_setInitConnection, layer)
	}
}

// Decode <--
func (m *TLAuthsessionSetInitConnection) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x7cdf8a8c: func() (err error) {
			m.AuthKeyId, err = d.Int64()
			m.Ip, err = d.String()
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

// TLAuthsessionSetAndroidPushSessionId <--
type TLAuthsessionSetAndroidPushSessionId struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	SessionId int64  `json:"session_id"`
}

// Encode <--
func (m *TLAuthsessionSetAndroidPushSessionId) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x92a8233c: func() error {
			x.PutClazzID(0x92a8233c)

			x.PutInt64(m.UserId)
			x.PutInt64(m.AuthKeyId)
			x.PutInt64(m.SessionId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_authsession_setAndroidPushSessionId, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_setAndroidPushSessionId, layer)
	}
}

// Decode <--
func (m *TLAuthsessionSetAndroidPushSessionId) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x92a8233c: func() (err error) {
			m.UserId, err = d.Int64()
			m.AuthKeyId, err = d.Int64()
			m.SessionId, err = d.Int64()

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

// ----------------------------------------------------------------------------
// rpc

type RPCAuthsession interface {
	AuthsessionGetAuthorizations(ctx context.Context, in *TLAuthsessionGetAuthorizations) (*tg.AccountAuthorizations, error)
	AuthsessionResetAuthorization(ctx context.Context, in *TLAuthsessionResetAuthorization) (*VectorLong, error)
	AuthsessionGetLayer(ctx context.Context, in *TLAuthsessionGetLayer) (*tg.Int32, error)
	AuthsessionGetLangPack(ctx context.Context, in *TLAuthsessionGetLangPack) (*tg.String, error)
	AuthsessionGetClient(ctx context.Context, in *TLAuthsessionGetClient) (*tg.String, error)
	AuthsessionGetLangCode(ctx context.Context, in *TLAuthsessionGetLangCode) (*tg.String, error)
	AuthsessionGetUserId(ctx context.Context, in *TLAuthsessionGetUserId) (*tg.Int64, error)
	AuthsessionGetPushSessionId(ctx context.Context, in *TLAuthsessionGetPushSessionId) (*tg.Int64, error)
	AuthsessionGetFutureSalts(ctx context.Context, in *TLAuthsessionGetFutureSalts) (*mt.FutureSalts, error)
	AuthsessionQueryAuthKey(ctx context.Context, in *TLAuthsessionQueryAuthKey) (*tg.AuthKeyInfo, error)
	AuthsessionSetAuthKey(ctx context.Context, in *TLAuthsessionSetAuthKey) (*tg.Bool, error)
	AuthsessionBindAuthKeyUser(ctx context.Context, in *TLAuthsessionBindAuthKeyUser) (*tg.Int64, error)
	AuthsessionUnbindAuthKeyUser(ctx context.Context, in *TLAuthsessionUnbindAuthKeyUser) (*tg.Bool, error)
	AuthsessionGetPermAuthKeyId(ctx context.Context, in *TLAuthsessionGetPermAuthKeyId) (*tg.Int64, error)
	AuthsessionBindTempAuthKey(ctx context.Context, in *TLAuthsessionBindTempAuthKey) (*tg.Bool, error)
	AuthsessionSetClientSessionInfo(ctx context.Context, in *TLAuthsessionSetClientSessionInfo) (*tg.Bool, error)
	AuthsessionGetAuthorization(ctx context.Context, in *TLAuthsessionGetAuthorization) (*tg.Authorization, error)
	AuthsessionGetAuthStateData(ctx context.Context, in *TLAuthsessionGetAuthStateData) (*AuthKeyStateData, error)
	AuthsessionSetLayer(ctx context.Context, in *TLAuthsessionSetLayer) (*tg.Bool, error)
	AuthsessionSetInitConnection(ctx context.Context, in *TLAuthsessionSetInitConnection) (*tg.Bool, error)
	AuthsessionSetAndroidPushSessionId(ctx context.Context, in *TLAuthsessionSetAndroidPushSessionId) (*tg.Bool, error)
}
