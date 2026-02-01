/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package mt

import (
	"fmt"

	"github.com/teamgooo/teamgooo-server/pkg/proto/bin"
	"github.com/teamgooo/teamgooo-server/pkg/proto/iface"
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

func (m *TLReqPq) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x60469778: func() error {
			x.PutClazzID(0x60469778)

			x.PutInt128(m.Nonce)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_req_pq, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_req_pq, layer)
	}
}

func (m *TLReqPq) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x60469778: func() error {
			err = m.Nonce.Decode(d)

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLReqPqMulti <--
type TLReqPqMulti struct {
	ClazzID uint32     `json:"_id"`
	Nonce   bin.Int128 `json:"Nonce"`
}

func (m *TLReqPqMulti) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

func (m *TLReqPqMulti) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xbe7e8ef1: func() error {
			x.PutClazzID(0xbe7e8ef1)

			x.PutInt128(m.Nonce)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_req_pq_multi, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_req_pq_multi, layer)
	}
}

func (m *TLReqPqMulti) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xbe7e8ef1: func() error {
			err = m.Nonce.Decode(d)

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

func (m *TLReqDHParams) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xd712e4be: func() error {
			x.PutClazzID(0xd712e4be)

			x.PutInt128(m.Nonce)
			x.PutInt128(m.ServerNonce)
			x.PutString(m.P)
			x.PutString(m.Q)
			x.PutInt64(m.PublicKeyFingerprint)
			x.PutString(m.EncryptedData)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_req_DH_params, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_req_DH_params, layer)
	}
}

func (m *TLReqDHParams) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xd712e4be: func() error {
			err = m.Nonce.Decode(d)
			err = m.ServerNonce.Decode(d)
			m.P, err = d.String()
			m.Q, err = d.String()
			m.PublicKeyFingerprint, err = d.Int64()
			m.EncryptedData, err = d.String()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
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
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

func (m *TLSetClientDHParams) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xf5045f1f: func() error {
			x.PutClazzID(0xf5045f1f)

			x.PutInt128(m.Nonce)
			x.PutInt128(m.ServerNonce)
			x.PutString(m.EncryptedData)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_set_client_DH_params, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_set_client_DH_params, layer)
	}
}

func (m *TLSetClientDHParams) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xf5045f1f: func() error {
			err = m.Nonce.Decode(d)
			err = m.ServerNonce.Decode(d)
			m.EncryptedData, err = d.String()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDestroyAuthKey <--
type TLDestroyAuthKey struct {
	ClazzID uint32 `json:"_id"`
}

func (m *TLDestroyAuthKey) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

func (m *TLDestroyAuthKey) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xd1435160: func() error {
			x.PutClazzID(0xd1435160)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_destroy_auth_key, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_destroy_auth_key, layer)
	}
}

func (m *TLDestroyAuthKey) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xd1435160: func() error {

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLHelpTest <--
type TLHelpTest struct {
	ClazzID uint32 `json:"_id"`
}

func (m *TLHelpTest) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

func (m *TLHelpTest) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc0e202f7: func() error {
			x.PutClazzID(0xc0e202f7)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_help_test, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_help_test, layer)
	}
}

func (m *TLHelpTest) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xc0e202f7: func() error {

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLTestUseError <--
type TLTestUseError struct {
	ClazzID uint32 `json:"_id"`
}

func (m *TLTestUseError) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

func (m *TLTestUseError) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xee75af01: func() error {
			x.PutClazzID(0xee75af01)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_test_useError, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_test_useError, layer)
	}
}

func (m *TLTestUseError) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xee75af01: func() error {

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLTestUseConfigSimple <--
type TLTestUseConfigSimple struct {
	ClazzID uint32 `json:"_id"`
}

func (m *TLTestUseConfigSimple) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

func (m *TLTestUseConfigSimple) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xf9b7b23d: func() error {
			x.PutClazzID(0xf9b7b23d)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_test_useConfigSimple, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_test_useConfigSimple, layer)
	}
}

func (m *TLTestUseConfigSimple) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xf9b7b23d: func() error {

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLRpcDropAnswer <--
type TLRpcDropAnswer struct {
	ClazzID  uint32 `json:"_id"`
	ReqMsgId int64  `json:"ReqMsgId"`
}

func (m *TLRpcDropAnswer) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

func (m *TLRpcDropAnswer) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x58e4a740: func() error {
			x.PutClazzID(0x58e4a740)

			x.PutInt64(m.ReqMsgId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_rpc_drop_answer, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_rpc_drop_answer, layer)
	}
}

func (m *TLRpcDropAnswer) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x58e4a740: func() error {
			m.ReqMsgId, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLGetFutureSalts <--
type TLGetFutureSalts struct {
	ClazzID uint32 `json:"_id"`
	Num     int32  `json:"Num"`
}

func (m *TLGetFutureSalts) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

func (m *TLGetFutureSalts) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xb921bd04: func() error {
			x.PutClazzID(0xb921bd04)

			x.PutInt32(m.Num)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_get_future_salts, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_get_future_salts, layer)
	}
}

func (m *TLGetFutureSalts) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xb921bd04: func() error {
			m.Num, err = d.Int32()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLPing <--
type TLPing struct {
	ClazzID uint32 `json:"_id"`
	PingId  int64  `json:"PingId"`
}

func (m *TLPing) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

func (m *TLPing) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x7abe77ec: func() error {
			x.PutClazzID(0x7abe77ec)

			x.PutInt64(m.PingId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_ping, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_ping, layer)
	}
}

func (m *TLPing) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x7abe77ec: func() error {
			m.PingId, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLPingDelayDisconnect <--
type TLPingDelayDisconnect struct {
	ClazzID         uint32 `json:"_id"`
	PingId          int64  `json:"PingId"`
	DisconnectDelay int32  `json:"DisconnectDelay"`
}

func (m *TLPingDelayDisconnect) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

func (m *TLPingDelayDisconnect) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xf3427b8c: func() error {
			x.PutClazzID(0xf3427b8c)

			x.PutInt64(m.PingId)
			x.PutInt32(m.DisconnectDelay)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_ping_delay_disconnect, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_ping_delay_disconnect, layer)
	}
}

func (m *TLPingDelayDisconnect) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xf3427b8c: func() error {
			m.PingId, err = d.Int64()
			m.DisconnectDelay, err = d.Int32()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDestroySession <--
type TLDestroySession struct {
	ClazzID   uint32 `json:"_id"`
	SessionId int64  `json:"SessionId"`
}

func (m *TLDestroySession) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

func (m *TLDestroySession) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xe7512126: func() error {
			x.PutClazzID(0xe7512126)

			x.PutInt64(m.SessionId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_destroy_session, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_destroy_session, layer)
	}
}

func (m *TLDestroySession) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xe7512126: func() error {
			m.SessionId, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}
