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

// TLAuthsessionGetAuthorizations <--
type TLAuthsessionGetAuthorizations struct {
	ClazzID          uint32 `json:"_id"`
	UserId           int64  `json:"user_id"`
	ExcludeAuthKeyId int64  `json:"exclude_auth_keyId"`
}

func (m *TLAuthsessionGetAuthorizations) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_getAuthorizations, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionGetAuthorizations) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_getAuthorizations, int(layer)); clazzId {
	case 0x30e21244:
		x.PutClazzID(0x30e21244)

		x.PutInt64(m.UserId)
		x.PutInt64(m.ExcludeAuthKeyId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getAuthorizations, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetAuthorizations) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x30e21244:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.ExcludeAuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
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

func (m *TLAuthsessionResetAuthorization) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_resetAuthorization, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionResetAuthorization) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_resetAuthorization, int(layer)); clazzId {
	case 0x8d5f6ca6:
		x.PutClazzID(0x8d5f6ca6)

		x.PutInt64(m.UserId)
		x.PutInt64(m.AuthKeyId)
		x.PutInt64(m.Hash)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_resetAuthorization, layer)
	}
}

// Decode <--
func (m *TLAuthsessionResetAuthorization) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x8d5f6ca6:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.Hash, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionGetLayer <--
type TLAuthsessionGetLayer struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

func (m *TLAuthsessionGetLayer) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_getLayer, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionGetLayer) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_getLayer, int(layer)); clazzId {
	case 0xa82f16a9:
		x.PutClazzID(0xa82f16a9)

		x.PutInt64(m.AuthKeyId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getLayer, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetLayer) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xa82f16a9:
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionGetLangPack <--
type TLAuthsessionGetLangPack struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

func (m *TLAuthsessionGetLangPack) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_getLangPack, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionGetLangPack) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_getLangPack, int(layer)); clazzId {
	case 0x29bbc166:
		x.PutClazzID(0x29bbc166)

		x.PutInt64(m.AuthKeyId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getLangPack, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetLangPack) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x29bbc166:
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionGetClient <--
type TLAuthsessionGetClient struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

func (m *TLAuthsessionGetClient) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_getClient, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionGetClient) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_getClient, int(layer)); clazzId {
	case 0x605855be:
		x.PutClazzID(0x605855be)

		x.PutInt64(m.AuthKeyId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getClient, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetClient) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x605855be:
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionGetLangCode <--
type TLAuthsessionGetLangCode struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

func (m *TLAuthsessionGetLangCode) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_getLangCode, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionGetLangCode) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_getLangCode, int(layer)); clazzId {
	case 0x5899b559:
		x.PutClazzID(0x5899b559)

		x.PutInt64(m.AuthKeyId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getLangCode, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetLangCode) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x5899b559:
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionGetUserId <--
type TLAuthsessionGetUserId struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

func (m *TLAuthsessionGetUserId) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_getUserId, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionGetUserId) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_getUserId, int(layer)); clazzId {
	case 0x57491cac:
		x.PutClazzID(0x57491cac)

		x.PutInt64(m.AuthKeyId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getUserId, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetUserId) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x57491cac:
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
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

func (m *TLAuthsessionGetPushSessionId) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_getPushSessionId, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionGetPushSessionId) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_getPushSessionId, int(layer)); clazzId {
	case 0xb3c23141:
		x.PutClazzID(0xb3c23141)

		x.PutInt64(m.UserId)
		x.PutInt64(m.AuthKeyId)
		x.PutInt32(m.TokenType)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getPushSessionId, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetPushSessionId) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xb3c23141:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.TokenType, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionGetFutureSalts <--
type TLAuthsessionGetFutureSalts struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	Num       int32  `json:"num"`
}

func (m *TLAuthsessionGetFutureSalts) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_getFutureSalts, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionGetFutureSalts) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_getFutureSalts, int(layer)); clazzId {
	case 0xb8cf5815:
		x.PutClazzID(0xb8cf5815)

		x.PutInt64(m.AuthKeyId)
		x.PutInt32(m.Num)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getFutureSalts, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetFutureSalts) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xb8cf5815:
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.Num, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionQueryAuthKey <--
type TLAuthsessionQueryAuthKey struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

func (m *TLAuthsessionQueryAuthKey) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_queryAuthKey, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionQueryAuthKey) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_queryAuthKey, int(layer)); clazzId {
	case 0x54b73828:
		x.PutClazzID(0x54b73828)

		x.PutInt64(m.AuthKeyId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_queryAuthKey, layer)
	}
}

// Decode <--
func (m *TLAuthsessionQueryAuthKey) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x54b73828:
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionSetAuthKey <--
type TLAuthsessionSetAuthKey struct {
	ClazzID    uint32              `json:"_id"`
	AuthKey    tg.AuthKeyInfoClazz `json:"auth_key"`
	FutureSalt tg.FutureSaltClazz  `json:"future_salt"`
	ExpiresIn  int32               `json:"expires_in"`
}

func (m *TLAuthsessionSetAuthKey) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_setAuthKey, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionSetAuthKey) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_setAuthKey, int(layer)); clazzId {
	case 0x3e940c91:
		x.PutClazzID(0x3e940c91)

		_ = m.AuthKey.Encode(x, layer)
		_ = m.FutureSalt.Encode(x, layer)
		x.PutInt32(m.ExpiresIn)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_setAuthKey, layer)
	}
}

// Decode <--
func (m *TLAuthsessionSetAuthKey) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x3e940c91:

		m.AuthKey, err = tg.DecodeAuthKeyInfoClazz(d)
		if err != nil {
			return err
		}

		m.FutureSalt, err = tg.DecodeFutureSaltClazz(d)
		if err != nil {
			return err
		}

		m.ExpiresIn, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionBindAuthKeyUser <--
type TLAuthsessionBindAuthKeyUser struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	UserId    int64  `json:"user_id"`
}

func (m *TLAuthsessionBindAuthKeyUser) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_bindAuthKeyUser, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionBindAuthKeyUser) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_bindAuthKeyUser, int(layer)); clazzId {
	case 0xbce0423:
		x.PutClazzID(0xbce0423)

		x.PutInt64(m.AuthKeyId)
		x.PutInt64(m.UserId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_bindAuthKeyUser, layer)
	}
}

// Decode <--
func (m *TLAuthsessionBindAuthKeyUser) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xbce0423:
		m.AuthKeyId, err = d.Int64()
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

// TLAuthsessionUnbindAuthKeyUser <--
type TLAuthsessionUnbindAuthKeyUser struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	UserId    int64  `json:"user_id"`
}

func (m *TLAuthsessionUnbindAuthKeyUser) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_unbindAuthKeyUser, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionUnbindAuthKeyUser) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_unbindAuthKeyUser, int(layer)); clazzId {
	case 0x758c648:
		x.PutClazzID(0x758c648)

		x.PutInt64(m.AuthKeyId)
		x.PutInt64(m.UserId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_unbindAuthKeyUser, layer)
	}
}

// Decode <--
func (m *TLAuthsessionUnbindAuthKeyUser) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x758c648:
		m.AuthKeyId, err = d.Int64()
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

// TLAuthsessionGetPermAuthKeyId <--
type TLAuthsessionGetPermAuthKeyId struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

func (m *TLAuthsessionGetPermAuthKeyId) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_getPermAuthKeyId, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionGetPermAuthKeyId) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_getPermAuthKeyId, int(layer)); clazzId {
	case 0x907464d6:
		x.PutClazzID(0x907464d6)

		x.PutInt64(m.AuthKeyId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getPermAuthKeyId, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetPermAuthKeyId) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x907464d6:
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
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

func (m *TLAuthsessionBindTempAuthKey) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_bindTempAuthKey, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionBindTempAuthKey) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_bindTempAuthKey, int(layer)); clazzId {
	case 0x608f4f86:
		x.PutClazzID(0x608f4f86)

		x.PutInt64(m.PermAuthKeyId)
		x.PutInt64(m.Nonce)
		x.PutInt32(m.ExpiresAt)
		x.PutBytes(m.EncryptedMessage)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_bindTempAuthKey, layer)
	}
}

// Decode <--
func (m *TLAuthsessionBindTempAuthKey) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x608f4f86:
		m.PermAuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.Nonce, err = d.Int64()
		if err != nil {
			return err
		}
		m.ExpiresAt, err = d.Int32()
		if err != nil {
			return err
		}
		m.EncryptedMessage, err = d.Bytes()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionSetClientSessionInfo <--
type TLAuthsessionSetClientSessionInfo struct {
	ClazzID uint32             `json:"_id"`
	Data    ClientSessionClazz `json:"data"`
}

func (m *TLAuthsessionSetClientSessionInfo) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_setClientSessionInfo, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionSetClientSessionInfo) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_setClientSessionInfo, int(layer)); clazzId {
	case 0x2d9ff94:
		x.PutClazzID(0x2d9ff94)

		_ = m.Data.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_setClientSessionInfo, layer)
	}
}

// Decode <--
func (m *TLAuthsessionSetClientSessionInfo) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x2d9ff94:

		m.Data, err = DecodeClientSessionClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionGetAuthorization <--
type TLAuthsessionGetAuthorization struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

func (m *TLAuthsessionGetAuthorization) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_getAuthorization, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionGetAuthorization) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_getAuthorization, int(layer)); clazzId {
	case 0x6e5e1923:
		x.PutClazzID(0x6e5e1923)

		x.PutInt64(m.AuthKeyId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getAuthorization, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetAuthorization) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x6e5e1923:
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLAuthsessionGetAuthStateData <--
type TLAuthsessionGetAuthStateData struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

func (m *TLAuthsessionGetAuthStateData) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_getAuthStateData, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionGetAuthStateData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_getAuthStateData, int(layer)); clazzId {
	case 0x4f5e3131:
		x.PutClazzID(0x4f5e3131)

		x.PutInt64(m.AuthKeyId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_getAuthStateData, layer)
	}
}

// Decode <--
func (m *TLAuthsessionGetAuthStateData) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x4f5e3131:
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
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

func (m *TLAuthsessionSetLayer) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_setLayer, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionSetLayer) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_setLayer, int(layer)); clazzId {
	case 0x44651485:
		x.PutClazzID(0x44651485)

		x.PutInt64(m.AuthKeyId)
		x.PutString(m.Ip)
		x.PutInt32(m.Layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_setLayer, layer)
	}
}

// Decode <--
func (m *TLAuthsessionSetLayer) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x44651485:
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

		return nil
	default:
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

func (m *TLAuthsessionSetInitConnection) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_setInitConnection, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionSetInitConnection) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_setInitConnection, int(layer)); clazzId {
	case 0x7cdf8a8c:
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
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_setInitConnection, layer)
	}
}

// Decode <--
func (m *TLAuthsessionSetInitConnection) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x7cdf8a8c:
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.Ip, err = d.String()
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

// TLAuthsessionSetAndroidPushSessionId <--
type TLAuthsessionSetAndroidPushSessionId struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	SessionId int64  `json:"session_id"`
}

func (m *TLAuthsessionSetAndroidPushSessionId) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_authsession_setAndroidPushSessionId, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLAuthsessionSetAndroidPushSessionId) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_authsession_setAndroidPushSessionId, int(layer)); clazzId {
	case 0x92a8233c:
		x.PutClazzID(0x92a8233c)

		x.PutInt64(m.UserId)
		x.PutInt64(m.AuthKeyId)
		x.PutInt64(m.SessionId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_authsession_setAndroidPushSessionId, layer)
	}
}

// Decode <--
func (m *TLAuthsessionSetAndroidPushSessionId) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x92a8233c:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.SessionId, err = d.Int64()
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

type RPCAuthsession interface {
	AuthsessionGetAuthorizations(ctx context.Context, in *TLAuthsessionGetAuthorizations) (*tg.AccountAuthorizations, error)
	AuthsessionResetAuthorization(ctx context.Context, in *TLAuthsessionResetAuthorization) (*VectorLong, error)
	AuthsessionGetLayer(ctx context.Context, in *TLAuthsessionGetLayer) (*tg.Int32, error)
	AuthsessionGetLangPack(ctx context.Context, in *TLAuthsessionGetLangPack) (*tg.String, error)
	AuthsessionGetClient(ctx context.Context, in *TLAuthsessionGetClient) (*tg.String, error)
	AuthsessionGetLangCode(ctx context.Context, in *TLAuthsessionGetLangCode) (*tg.String, error)
	AuthsessionGetUserId(ctx context.Context, in *TLAuthsessionGetUserId) (*tg.Int64, error)
	AuthsessionGetPushSessionId(ctx context.Context, in *TLAuthsessionGetPushSessionId) (*tg.Int64, error)
	AuthsessionGetFutureSalts(ctx context.Context, in *TLAuthsessionGetFutureSalts) (*tg.FutureSalts, error)
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
