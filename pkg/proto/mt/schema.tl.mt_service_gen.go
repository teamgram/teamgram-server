/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mt

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

var (
	_ iface.TLObject
)

// TLReqPq <--
type TLReqPq struct {
	ClazzID uint32     `json:"_id"`
	Nonce   bin.Int128 `json:"Nonce"`
}

func (m *TLReqPq) String() string {
	return iface.DebugStringWithName(ClazzName_req_pq, m)
}

func (m *TLReqPq) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_req_pq, int(layer)); clazzId {
	case 0x60469778:
		x.PutClazzID(0x60469778)

		x.PutInt128(m.Nonce)

		return nil
	default:
		return fmt.Errorf("unable to encode req_pq: unsupported layer %d", layer)
	}
}

func (m *TLReqPq) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode req_pq: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x60469778:
		err = m.Nonce.Decode(d)
		if err != nil {
			return fmt.Errorf("unable to decode req_pq#0x60469778: field nonce: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode req_pq: invalid constructor %x", m.ClazzID)
	}
}

// TLReqPqMulti <--
type TLReqPqMulti struct {
	ClazzID uint32     `json:"_id"`
	Nonce   bin.Int128 `json:"Nonce"`
}

func (m *TLReqPqMulti) String() string {
	return iface.DebugStringWithName(ClazzName_req_pq_multi, m)
}

func (m *TLReqPqMulti) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_req_pq_multi, int(layer)); clazzId {
	case 0xbe7e8ef1:
		x.PutClazzID(0xbe7e8ef1)

		x.PutInt128(m.Nonce)

		return nil
	default:
		return fmt.Errorf("unable to encode req_pq_multi: unsupported layer %d", layer)
	}
}

func (m *TLReqPqMulti) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode req_pq_multi: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xbe7e8ef1:
		err = m.Nonce.Decode(d)
		if err != nil {
			return fmt.Errorf("unable to decode req_pq_multi#0xbe7e8ef1: field nonce: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode req_pq_multi: invalid constructor %x", m.ClazzID)
	}
}

// TLReqDHParams <--
type TLReqDHParams struct {
	ClazzID              uint32     `json:"_id"`
	Nonce                bin.Int128 `json:"Nonce"`
	ServerNonce          bin.Int128 `json:"ServerNonce"`
	P                    string     `json:"P"`
	Q                    string     `json:"Q"`
	PublicKeyFingerprint int64      `json:"PublicKeyFingerprint"`
	EncryptedData        string     `json:"EncryptedData"`
}

func (m *TLReqDHParams) String() string {
	return iface.DebugStringWithName(ClazzName_req_DH_params, m)
}

func (m *TLReqDHParams) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_req_DH_params, int(layer)); clazzId {
	case 0xd712e4be:
		x.PutClazzID(0xd712e4be)

		x.PutInt128(m.Nonce)
		x.PutInt128(m.ServerNonce)
		x.PutString(m.P)
		x.PutString(m.Q)
		x.PutInt64(m.PublicKeyFingerprint)
		x.PutString(m.EncryptedData)

		return nil
	default:
		return fmt.Errorf("unable to encode req_DH_params: unsupported layer %d", layer)
	}
}

func (m *TLReqDHParams) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode req_DH_params: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xd712e4be:
		err = m.Nonce.Decode(d)
		if err != nil {
			return fmt.Errorf("unable to decode req_DH_params#0xd712e4be: field nonce: %w", err)
		}
		err = m.ServerNonce.Decode(d)
		if err != nil {
			return fmt.Errorf("unable to decode req_DH_params#0xd712e4be: field server_nonce: %w", err)
		}
		m.P, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode req_DH_params#0xd712e4be: field p: %w", err)
		}
		m.Q, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode req_DH_params#0xd712e4be: field q: %w", err)
		}
		m.PublicKeyFingerprint, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode req_DH_params#0xd712e4be: field public_key_fingerprint: %w", err)
		}
		m.EncryptedData, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode req_DH_params#0xd712e4be: field encrypted_data: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode req_DH_params: invalid constructor %x", m.ClazzID)
	}
}

// TLSetClientDHParams <--
type TLSetClientDHParams struct {
	ClazzID       uint32     `json:"_id"`
	Nonce         bin.Int128 `json:"Nonce"`
	ServerNonce   bin.Int128 `json:"ServerNonce"`
	EncryptedData string     `json:"EncryptedData"`
}

func (m *TLSetClientDHParams) String() string {
	return iface.DebugStringWithName(ClazzName_set_client_DH_params, m)
}

func (m *TLSetClientDHParams) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_set_client_DH_params, int(layer)); clazzId {
	case 0xf5045f1f:
		x.PutClazzID(0xf5045f1f)

		x.PutInt128(m.Nonce)
		x.PutInt128(m.ServerNonce)
		x.PutString(m.EncryptedData)

		return nil
	default:
		return fmt.Errorf("unable to encode set_client_DH_params: unsupported layer %d", layer)
	}
}

func (m *TLSetClientDHParams) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode set_client_DH_params: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xf5045f1f:
		err = m.Nonce.Decode(d)
		if err != nil {
			return fmt.Errorf("unable to decode set_client_DH_params#0xf5045f1f: field nonce: %w", err)
		}
		err = m.ServerNonce.Decode(d)
		if err != nil {
			return fmt.Errorf("unable to decode set_client_DH_params#0xf5045f1f: field server_nonce: %w", err)
		}
		m.EncryptedData, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode set_client_DH_params#0xf5045f1f: field encrypted_data: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode set_client_DH_params: invalid constructor %x", m.ClazzID)
	}
}

// TLDestroyAuthKey <--
type TLDestroyAuthKey struct {
	ClazzID uint32 `json:"_id"`
}

func (m *TLDestroyAuthKey) String() string {
	return iface.DebugStringWithName(ClazzName_destroy_auth_key, m)
}

func (m *TLDestroyAuthKey) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_destroy_auth_key, int(layer)); clazzId {
	case 0xd1435160:
		x.PutClazzID(0xd1435160)

		return nil
	default:
		return fmt.Errorf("unable to encode destroy_auth_key: unsupported layer %d", layer)
	}
}

func (m *TLDestroyAuthKey) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode destroy_auth_key: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xd1435160:

		return nil
	default:
		return fmt.Errorf("unable to decode destroy_auth_key: invalid constructor %x", m.ClazzID)
	}
}

// TLHelpTest <--
type TLHelpTest struct {
	ClazzID uint32 `json:"_id"`
}

func (m *TLHelpTest) String() string {
	return iface.DebugStringWithName(ClazzName_help_test, m)
}

func (m *TLHelpTest) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_help_test, int(layer)); clazzId {
	case 0xc0e202f7:
		x.PutClazzID(0xc0e202f7)

		return nil
	default:
		return fmt.Errorf("unable to encode help_test: unsupported layer %d", layer)
	}
}

func (m *TLHelpTest) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode help_test: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xc0e202f7:

		return nil
	default:
		return fmt.Errorf("unable to decode help_test: invalid constructor %x", m.ClazzID)
	}
}

// TLTestUseError <--
type TLTestUseError struct {
	ClazzID uint32 `json:"_id"`
}

func (m *TLTestUseError) String() string {
	return iface.DebugStringWithName(ClazzName_test_useError, m)
}

func (m *TLTestUseError) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_test_useError, int(layer)); clazzId {
	case 0xee75af01:
		x.PutClazzID(0xee75af01)

		return nil
	default:
		return fmt.Errorf("unable to encode test_useError: unsupported layer %d", layer)
	}
}

func (m *TLTestUseError) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode test_useError: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xee75af01:

		return nil
	default:
		return fmt.Errorf("unable to decode test_useError: invalid constructor %x", m.ClazzID)
	}
}

// TLTestUseConfigSimple <--
type TLTestUseConfigSimple struct {
	ClazzID uint32 `json:"_id"`
}

func (m *TLTestUseConfigSimple) String() string {
	return iface.DebugStringWithName(ClazzName_test_useConfigSimple, m)
}

func (m *TLTestUseConfigSimple) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_test_useConfigSimple, int(layer)); clazzId {
	case 0xf9b7b23d:
		x.PutClazzID(0xf9b7b23d)

		return nil
	default:
		return fmt.Errorf("unable to encode test_useConfigSimple: unsupported layer %d", layer)
	}
}

func (m *TLTestUseConfigSimple) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode test_useConfigSimple: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xf9b7b23d:

		return nil
	default:
		return fmt.Errorf("unable to decode test_useConfigSimple: invalid constructor %x", m.ClazzID)
	}
}

// TLRpcDropAnswer <--
type TLRpcDropAnswer struct {
	ClazzID  uint32 `json:"_id"`
	ReqMsgId int64  `json:"ReqMsgId"`
}

func (m *TLRpcDropAnswer) String() string {
	return iface.DebugStringWithName(ClazzName_rpc_drop_answer, m)
}

func (m *TLRpcDropAnswer) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_rpc_drop_answer, int(layer)); clazzId {
	case 0x58e4a740:
		x.PutClazzID(0x58e4a740)

		x.PutInt64(m.ReqMsgId)

		return nil
	default:
		return fmt.Errorf("unable to encode rpc_drop_answer: unsupported layer %d", layer)
	}
}

func (m *TLRpcDropAnswer) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode rpc_drop_answer: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x58e4a740:
		m.ReqMsgId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode rpc_drop_answer#0x58e4a740: field req_msg_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode rpc_drop_answer: invalid constructor %x", m.ClazzID)
	}
}

// TLGetFutureSalts <--
type TLGetFutureSalts struct {
	ClazzID uint32 `json:"_id"`
	Num     int32  `json:"Num"`
}

func (m *TLGetFutureSalts) String() string {
	return iface.DebugStringWithName(ClazzName_get_future_salts, m)
}

func (m *TLGetFutureSalts) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_get_future_salts, int(layer)); clazzId {
	case 0xb921bd04:
		x.PutClazzID(0xb921bd04)

		x.PutInt32(m.Num)

		return nil
	default:
		return fmt.Errorf("unable to encode get_future_salts: unsupported layer %d", layer)
	}
}

func (m *TLGetFutureSalts) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode get_future_salts: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xb921bd04:
		m.Num, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode get_future_salts#0xb921bd04: field num: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode get_future_salts: invalid constructor %x", m.ClazzID)
	}
}

// TLPing <--
type TLPing struct {
	ClazzID uint32 `json:"_id"`
	PingId  int64  `json:"PingId"`
}

func (m *TLPing) String() string {
	return iface.DebugStringWithName(ClazzName_ping, m)
}

func (m *TLPing) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_ping, int(layer)); clazzId {
	case 0x7abe77ec:
		x.PutClazzID(0x7abe77ec)

		x.PutInt64(m.PingId)

		return nil
	default:
		return fmt.Errorf("unable to encode ping: unsupported layer %d", layer)
	}
}

func (m *TLPing) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode ping: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x7abe77ec:
		m.PingId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode ping#0x7abe77ec: field ping_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode ping: invalid constructor %x", m.ClazzID)
	}
}

// TLPingDelayDisconnect <--
type TLPingDelayDisconnect struct {
	ClazzID         uint32 `json:"_id"`
	PingId          int64  `json:"PingId"`
	DisconnectDelay int32  `json:"DisconnectDelay"`
}

func (m *TLPingDelayDisconnect) String() string {
	return iface.DebugStringWithName(ClazzName_ping_delay_disconnect, m)
}

func (m *TLPingDelayDisconnect) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_ping_delay_disconnect, int(layer)); clazzId {
	case 0xf3427b8c:
		x.PutClazzID(0xf3427b8c)

		x.PutInt64(m.PingId)
		x.PutInt32(m.DisconnectDelay)

		return nil
	default:
		return fmt.Errorf("unable to encode ping_delay_disconnect: unsupported layer %d", layer)
	}
}

func (m *TLPingDelayDisconnect) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode ping_delay_disconnect: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xf3427b8c:
		m.PingId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode ping_delay_disconnect#0xf3427b8c: field ping_id: %w", err)
		}
		m.DisconnectDelay, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode ping_delay_disconnect#0xf3427b8c: field disconnect_delay: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode ping_delay_disconnect: invalid constructor %x", m.ClazzID)
	}
}

// TLDestroySession <--
type TLDestroySession struct {
	ClazzID   uint32 `json:"_id"`
	SessionId int64  `json:"SessionId"`
}

func (m *TLDestroySession) String() string {
	return iface.DebugStringWithName(ClazzName_destroy_session, m)
}

func (m *TLDestroySession) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_destroy_session, int(layer)); clazzId {
	case 0xe7512126:
		x.PutClazzID(0xe7512126)

		x.PutInt64(m.SessionId)

		return nil
	default:
		return fmt.Errorf("unable to encode destroy_session: unsupported layer %d", layer)
	}
}

func (m *TLDestroySession) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode destroy_session: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xe7512126:
		m.SessionId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode destroy_session#0xe7512126: field session_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode destroy_session: invalid constructor %x", m.ClazzID)
	}
}
