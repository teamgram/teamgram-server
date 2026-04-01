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

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

var (
	_ iface.TLObject
)

// BindAuthKeyInnerClazz <--
//   - TL_BindAuthKeyInner
type BindAuthKeyInnerClazz = *TLBindAuthKeyInner

func DecodeBindAuthKeyInnerClazz(d *bin.Decoder) (BindAuthKeyInnerClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x75a3f765:
		x := &TLBindAuthKeyInner{ClazzID: id, ClazzName2: ClazzName_bind_auth_key_inner}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeBindAuthKeyInner - unexpected clazzId: %d", id)
	}

}

// TLBindAuthKeyInner <--
type TLBindAuthKeyInner struct {
	ClazzID       uint32 `json:"_id"`
	ClazzName2    string `json:"_name"`
	Nonce         int64  `json:"nonce"`
	TempAuthKeyId int64  `json:"temp_auth_key_id"`
	PermAuthKeyId int64  `json:"perm_auth_key_id"`
	TempSessionId int64  `json:"temp_session_id"`
	ExpiresAt     int32  `json:"expires_at"`
}

func MakeTLBindAuthKeyInner(m *TLBindAuthKeyInner) *TLBindAuthKeyInner {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_bind_auth_key_inner

	return m
}

func (m *TLBindAuthKeyInner) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "bind_auth_key_inner", TLObject: m}
	return wrapper.String()
}

// BindAuthKeyInnerClazzName <--
func (m *TLBindAuthKeyInner) BindAuthKeyInnerClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLBindAuthKeyInner) ClazzName() string {
	return ClazzName_bind_auth_key_inner
}

// ToBindAuthKeyInner <--
func (m *TLBindAuthKeyInner) ToBindAuthKeyInner() *BindAuthKeyInner {
	if m == nil {
		return nil
	}

	return m

}

// Encode <--
func (m *TLBindAuthKeyInner) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_bind_auth_key_inner, int(layer)); clazzId {
	case 0x75a3f765:
		x.PutClazzID(0x75a3f765)

		x.PutInt64(m.Nonce)
		x.PutInt64(m.TempAuthKeyId)
		x.PutInt64(m.PermAuthKeyId)
		x.PutInt64(m.TempSessionId)
		x.PutInt32(m.ExpiresAt)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_bind_auth_key_inner, layer)
	}
}

// Decode <--
func (m *TLBindAuthKeyInner) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x75a3f765:
		m.Nonce, err = d.Int64()
		m.TempAuthKeyId, err = d.Int64()
		m.PermAuthKeyId, err = d.Int64()
		m.TempSessionId, err = d.Int64()
		m.ExpiresAt, err = d.Int32()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// BindAuthKeyInner <--
type BindAuthKeyInner = TLBindAuthKeyInner

// ClientDHInnerDataClazz <--
//   - TL_ClientDHInnerData
type ClientDHInnerDataClazz = *TLClientDHInnerData

func DecodeClientDHInnerDataClazz(d *bin.Decoder) (ClientDHInnerDataClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x6643b654:
		x := &TLClientDHInnerData{ClazzID: id, ClazzName2: ClazzName_client_DH_inner_data}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeClientDHInnerData - unexpected clazzId: %d", id)
	}

}

// TLClientDHInnerData <--
type TLClientDHInnerData struct {
	ClazzID     uint32     `json:"_id"`
	ClazzName2  string     `json:"_name"`
	Nonce       bin.Int128 `json:"nonce"`
	ServerNonce bin.Int128 `json:"server_nonce"`
	RetryId     int64      `json:"retry_id"`
	GB          string     `json:"g_b"`
}

func MakeTLClientDHInnerData(m *TLClientDHInnerData) *TLClientDHInnerData {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_client_DH_inner_data

	return m
}

func (m *TLClientDHInnerData) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "client_DH_inner_data", TLObject: m}
	return wrapper.String()
}

// ClientDHInnerDataClazzName <--
func (m *TLClientDHInnerData) ClientDHInnerDataClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLClientDHInnerData) ClazzName() string {
	return ClazzName_client_DH_inner_data
}

// ToClientDHInnerData <--
func (m *TLClientDHInnerData) ToClientDHInnerData() *ClientDHInnerData {
	if m == nil {
		return nil
	}

	return m

}

// Encode <--
func (m *TLClientDHInnerData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_client_DH_inner_data, int(layer)); clazzId {
	case 0x6643b654:
		x.PutClazzID(0x6643b654)

		x.PutInt128(m.Nonce)
		x.PutInt128(m.ServerNonce)
		x.PutInt64(m.RetryId)
		x.PutString(m.GB)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_client_DH_inner_data, layer)
	}
}

// Decode <--
func (m *TLClientDHInnerData) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x6643b654:
		err = m.Nonce.Decode(d)
		err = m.ServerNonce.Decode(d)
		m.RetryId, err = d.Int64()
		m.GB, err = d.String()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// ClientDHInnerData <--
type ClientDHInnerData = TLClientDHInnerData

// DestroyAuthKeyResClazz <--
//   - TL_DestroyAuthKeyOk
//   - TL_DestroyAuthKeyNone
//   - TL_DestroyAuthKeyFail
type DestroyAuthKeyResClazz interface {
	iface.TLObject
	DestroyAuthKeyResClazzName() string
}

func DecodeDestroyAuthKeyResClazz(d *bin.Decoder) (DestroyAuthKeyResClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0xf660e1d4:
		x := &TLDestroyAuthKeyOk{ClazzID: id, ClazzName2: ClazzName_destroy_auth_key_ok}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0xa9f2259:
		x := &TLDestroyAuthKeyNone{ClazzID: id, ClazzName2: ClazzName_destroy_auth_key_none}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0xea109b13:
		x := &TLDestroyAuthKeyFail{ClazzID: id, ClazzName2: ClazzName_destroy_auth_key_fail}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeDestroyAuthKeyRes - unexpected clazzId: %d", id)
	}

}

// TLDestroyAuthKeyOk <--
type TLDestroyAuthKeyOk struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
}

func MakeTLDestroyAuthKeyOk(m *TLDestroyAuthKeyOk) *TLDestroyAuthKeyOk {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_destroy_auth_key_ok

	return m
}

func (m *TLDestroyAuthKeyOk) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "destroy_auth_key_ok", TLObject: m}
	return wrapper.String()
}

// DestroyAuthKeyResClazzName <--
func (m *TLDestroyAuthKeyOk) DestroyAuthKeyResClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLDestroyAuthKeyOk) ClazzName() string {
	return ClazzName_destroy_auth_key_ok
}

// ToDestroyAuthKeyRes <--
func (m *TLDestroyAuthKeyOk) ToDestroyAuthKeyRes() *DestroyAuthKeyRes {
	if m == nil {
		return nil
	}

	return &DestroyAuthKeyRes{Clazz: m}

}

// Encode <--
func (m *TLDestroyAuthKeyOk) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_destroy_auth_key_ok, int(layer)); clazzId {
	case 0xf660e1d4:
		x.PutClazzID(0xf660e1d4)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_destroy_auth_key_ok, layer)
	}
}

// Decode <--
func (m *TLDestroyAuthKeyOk) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xf660e1d4:

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDestroyAuthKeyNone <--
type TLDestroyAuthKeyNone struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
}

func MakeTLDestroyAuthKeyNone(m *TLDestroyAuthKeyNone) *TLDestroyAuthKeyNone {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_destroy_auth_key_none

	return m
}

func (m *TLDestroyAuthKeyNone) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "destroy_auth_key_none", TLObject: m}
	return wrapper.String()
}

// DestroyAuthKeyResClazzName <--
func (m *TLDestroyAuthKeyNone) DestroyAuthKeyResClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLDestroyAuthKeyNone) ClazzName() string {
	return ClazzName_destroy_auth_key_none
}

// ToDestroyAuthKeyRes <--
func (m *TLDestroyAuthKeyNone) ToDestroyAuthKeyRes() *DestroyAuthKeyRes {
	if m == nil {
		return nil
	}

	return &DestroyAuthKeyRes{Clazz: m}

}

// Encode <--
func (m *TLDestroyAuthKeyNone) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_destroy_auth_key_none, int(layer)); clazzId {
	case 0xa9f2259:
		x.PutClazzID(0xa9f2259)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_destroy_auth_key_none, layer)
	}
}

// Decode <--
func (m *TLDestroyAuthKeyNone) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xa9f2259:

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDestroyAuthKeyFail <--
type TLDestroyAuthKeyFail struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
}

func MakeTLDestroyAuthKeyFail(m *TLDestroyAuthKeyFail) *TLDestroyAuthKeyFail {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_destroy_auth_key_fail

	return m
}

func (m *TLDestroyAuthKeyFail) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "destroy_auth_key_fail", TLObject: m}
	return wrapper.String()
}

// DestroyAuthKeyResClazzName <--
func (m *TLDestroyAuthKeyFail) DestroyAuthKeyResClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLDestroyAuthKeyFail) ClazzName() string {
	return ClazzName_destroy_auth_key_fail
}

// ToDestroyAuthKeyRes <--
func (m *TLDestroyAuthKeyFail) ToDestroyAuthKeyRes() *DestroyAuthKeyRes {
	if m == nil {
		return nil
	}

	return &DestroyAuthKeyRes{Clazz: m}

}

// Encode <--
func (m *TLDestroyAuthKeyFail) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_destroy_auth_key_fail, int(layer)); clazzId {
	case 0xea109b13:
		x.PutClazzID(0xea109b13)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_destroy_auth_key_fail, layer)
	}
}

// Decode <--
func (m *TLDestroyAuthKeyFail) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xea109b13:

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// DestroyAuthKeyRes <--
type DestroyAuthKeyRes struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz DestroyAuthKeyResClazz `json:"_clazz"`
}

func (m *DestroyAuthKeyRes) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.DestroyAuthKeyResClazzName()
	}
}

func (m *DestroyAuthKeyRes) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: m.ClazzName(), TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *DestroyAuthKeyRes) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("DestroyAuthKeyRes - invalid Clazz")
}

// Decode <--
func (m *DestroyAuthKeyRes) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeDestroyAuthKeyResClazz(d)
	return
}

// Match <--
func (m *DestroyAuthKeyRes) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}

	switch c := m.Clazz.(type) {
	case *TLDestroyAuthKeyOk:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDestroyAuthKeyOk) interface{}); ok {
				f1(c)
			}
		}
	case *TLDestroyAuthKeyNone:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDestroyAuthKeyNone) interface{}); ok {
				f1(c)
			}
		}
	case *TLDestroyAuthKeyFail:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDestroyAuthKeyFail) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToDestroyAuthKeyOk <--
func (m *DestroyAuthKeyRes) ToDestroyAuthKeyOk() (*TLDestroyAuthKeyOk, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLDestroyAuthKeyOk); ok {
		return x, true
	}

	return nil, false
}

// ToDestroyAuthKeyNone <--
func (m *DestroyAuthKeyRes) ToDestroyAuthKeyNone() (*TLDestroyAuthKeyNone, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLDestroyAuthKeyNone); ok {
		return x, true
	}

	return nil, false
}

// ToDestroyAuthKeyFail <--
func (m *DestroyAuthKeyRes) ToDestroyAuthKeyFail() (*TLDestroyAuthKeyFail, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLDestroyAuthKeyFail); ok {
		return x, true
	}

	return nil, false
}

// PQInnerDataClazz <--
//   - TL_PQInnerData
//   - TL_PQInnerDataDc
//   - TL_PQInnerDataTemp
//   - TL_PQInnerDataTempDc
type PQInnerDataClazz interface {
	iface.TLObject
	PQInnerDataClazzName() string
}

func DecodePQInnerDataClazz(d *bin.Decoder) (PQInnerDataClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x83c95aec:
		x := &TLPQInnerData{ClazzID: id, ClazzName2: ClazzName_p_q_inner_data}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0xa9f55f95:
		x := &TLPQInnerDataDc{ClazzID: id, ClazzName2: ClazzName_p_q_inner_data_dc}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x3c6a84d4:
		x := &TLPQInnerDataTemp{ClazzID: id, ClazzName2: ClazzName_p_q_inner_data_temp}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x56fddf88:
		x := &TLPQInnerDataTempDc{ClazzID: id, ClazzName2: ClazzName_p_q_inner_data_temp_dc}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodePQInnerData - unexpected clazzId: %d", id)
	}

}

// TLPQInnerData <--
type TLPQInnerData struct {
	ClazzID     uint32     `json:"_id"`
	ClazzName2  string     `json:"_name"`
	Pq          string     `json:"pq"`
	P           string     `json:"p"`
	Q           string     `json:"q"`
	Nonce       bin.Int128 `json:"nonce"`
	ServerNonce bin.Int128 `json:"server_nonce"`
	NewNonce    bin.Int256 `json:"new_nonce"`
}

func MakeTLPQInnerData(m *TLPQInnerData) *TLPQInnerData {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_p_q_inner_data

	return m
}

func (m *TLPQInnerData) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "p_q_inner_data", TLObject: m}
	return wrapper.String()
}

// PQInnerDataClazzName <--
func (m *TLPQInnerData) PQInnerDataClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLPQInnerData) ClazzName() string {
	return ClazzName_p_q_inner_data
}

// ToPQInnerData <--
func (m *TLPQInnerData) ToPQInnerData() *PQInnerData {
	if m == nil {
		return nil
	}

	return &PQInnerData{Clazz: m}

}

// Encode <--
func (m *TLPQInnerData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_p_q_inner_data, int(layer)); clazzId {
	case 0x83c95aec:
		x.PutClazzID(0x83c95aec)

		x.PutString(m.Pq)
		x.PutString(m.P)
		x.PutString(m.Q)
		x.PutInt128(m.Nonce)
		x.PutInt128(m.ServerNonce)
		x.PutInt256(m.NewNonce)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_p_q_inner_data, layer)
	}
}

// Decode <--
func (m *TLPQInnerData) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x83c95aec:
		m.Pq, err = d.String()
		m.P, err = d.String()
		m.Q, err = d.String()
		err = m.Nonce.Decode(d)
		err = m.ServerNonce.Decode(d)
		err = m.NewNonce.Decode(d)

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLPQInnerDataDc <--
type TLPQInnerDataDc struct {
	ClazzID     uint32     `json:"_id"`
	ClazzName2  string     `json:"_name"`
	Pq          string     `json:"pq"`
	P           string     `json:"p"`
	Q           string     `json:"q"`
	Nonce       bin.Int128 `json:"nonce"`
	ServerNonce bin.Int128 `json:"server_nonce"`
	NewNonce    bin.Int256 `json:"new_nonce"`
	Dc          int32      `json:"dc"`
}

func MakeTLPQInnerDataDc(m *TLPQInnerDataDc) *TLPQInnerDataDc {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_p_q_inner_data_dc

	return m
}

func (m *TLPQInnerDataDc) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "p_q_inner_data_dc", TLObject: m}
	return wrapper.String()
}

// PQInnerDataClazzName <--
func (m *TLPQInnerDataDc) PQInnerDataClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLPQInnerDataDc) ClazzName() string {
	return ClazzName_p_q_inner_data_dc
}

// ToPQInnerData <--
func (m *TLPQInnerDataDc) ToPQInnerData() *PQInnerData {
	if m == nil {
		return nil
	}

	return &PQInnerData{Clazz: m}

}

// Encode <--
func (m *TLPQInnerDataDc) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_p_q_inner_data_dc, int(layer)); clazzId {
	case 0xa9f55f95:
		x.PutClazzID(0xa9f55f95)

		x.PutString(m.Pq)
		x.PutString(m.P)
		x.PutString(m.Q)
		x.PutInt128(m.Nonce)
		x.PutInt128(m.ServerNonce)
		x.PutInt256(m.NewNonce)
		x.PutInt32(m.Dc)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_p_q_inner_data_dc, layer)
	}
}

// Decode <--
func (m *TLPQInnerDataDc) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xa9f55f95:
		m.Pq, err = d.String()
		m.P, err = d.String()
		m.Q, err = d.String()
		err = m.Nonce.Decode(d)
		err = m.ServerNonce.Decode(d)
		err = m.NewNonce.Decode(d)
		m.Dc, err = d.Int32()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLPQInnerDataTemp <--
type TLPQInnerDataTemp struct {
	ClazzID     uint32     `json:"_id"`
	ClazzName2  string     `json:"_name"`
	Pq          string     `json:"pq"`
	P           string     `json:"p"`
	Q           string     `json:"q"`
	Nonce       bin.Int128 `json:"nonce"`
	ServerNonce bin.Int128 `json:"server_nonce"`
	NewNonce    bin.Int256 `json:"new_nonce"`
	ExpiresIn   int32      `json:"expires_in"`
}

func MakeTLPQInnerDataTemp(m *TLPQInnerDataTemp) *TLPQInnerDataTemp {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_p_q_inner_data_temp

	return m
}

func (m *TLPQInnerDataTemp) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "p_q_inner_data_temp", TLObject: m}
	return wrapper.String()
}

// PQInnerDataClazzName <--
func (m *TLPQInnerDataTemp) PQInnerDataClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLPQInnerDataTemp) ClazzName() string {
	return ClazzName_p_q_inner_data_temp
}

// ToPQInnerData <--
func (m *TLPQInnerDataTemp) ToPQInnerData() *PQInnerData {
	if m == nil {
		return nil
	}

	return &PQInnerData{Clazz: m}

}

// Encode <--
func (m *TLPQInnerDataTemp) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_p_q_inner_data_temp, int(layer)); clazzId {
	case 0x3c6a84d4:
		x.PutClazzID(0x3c6a84d4)

		x.PutString(m.Pq)
		x.PutString(m.P)
		x.PutString(m.Q)
		x.PutInt128(m.Nonce)
		x.PutInt128(m.ServerNonce)
		x.PutInt256(m.NewNonce)
		x.PutInt32(m.ExpiresIn)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_p_q_inner_data_temp, layer)
	}
}

// Decode <--
func (m *TLPQInnerDataTemp) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x3c6a84d4:
		m.Pq, err = d.String()
		m.P, err = d.String()
		m.Q, err = d.String()
		err = m.Nonce.Decode(d)
		err = m.ServerNonce.Decode(d)
		err = m.NewNonce.Decode(d)
		m.ExpiresIn, err = d.Int32()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLPQInnerDataTempDc <--
type TLPQInnerDataTempDc struct {
	ClazzID     uint32     `json:"_id"`
	ClazzName2  string     `json:"_name"`
	Pq          string     `json:"pq"`
	P           string     `json:"p"`
	Q           string     `json:"q"`
	Nonce       bin.Int128 `json:"nonce"`
	ServerNonce bin.Int128 `json:"server_nonce"`
	NewNonce    bin.Int256 `json:"new_nonce"`
	Dc          int32      `json:"dc"`
	ExpiresIn   int32      `json:"expires_in"`
}

func MakeTLPQInnerDataTempDc(m *TLPQInnerDataTempDc) *TLPQInnerDataTempDc {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_p_q_inner_data_temp_dc

	return m
}

func (m *TLPQInnerDataTempDc) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "p_q_inner_data_temp_dc", TLObject: m}
	return wrapper.String()
}

// PQInnerDataClazzName <--
func (m *TLPQInnerDataTempDc) PQInnerDataClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLPQInnerDataTempDc) ClazzName() string {
	return ClazzName_p_q_inner_data_temp_dc
}

// ToPQInnerData <--
func (m *TLPQInnerDataTempDc) ToPQInnerData() *PQInnerData {
	if m == nil {
		return nil
	}

	return &PQInnerData{Clazz: m}

}

// Encode <--
func (m *TLPQInnerDataTempDc) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_p_q_inner_data_temp_dc, int(layer)); clazzId {
	case 0x56fddf88:
		x.PutClazzID(0x56fddf88)

		x.PutString(m.Pq)
		x.PutString(m.P)
		x.PutString(m.Q)
		x.PutInt128(m.Nonce)
		x.PutInt128(m.ServerNonce)
		x.PutInt256(m.NewNonce)
		x.PutInt32(m.Dc)
		x.PutInt32(m.ExpiresIn)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_p_q_inner_data_temp_dc, layer)
	}
}

// Decode <--
func (m *TLPQInnerDataTempDc) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x56fddf88:
		m.Pq, err = d.String()
		m.P, err = d.String()
		m.Q, err = d.String()
		err = m.Nonce.Decode(d)
		err = m.ServerNonce.Decode(d)
		err = m.NewNonce.Decode(d)
		m.Dc, err = d.Int32()
		m.ExpiresIn, err = d.Int32()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// PQInnerData <--
type PQInnerData struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz PQInnerDataClazz `json:"_clazz"`
}

func (m *PQInnerData) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.PQInnerDataClazzName()
	}
}

func (m *PQInnerData) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: m.ClazzName(), TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *PQInnerData) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("PQInnerData - invalid Clazz")
}

// Decode <--
func (m *PQInnerData) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodePQInnerDataClazz(d)
	return
}

// Match <--
func (m *PQInnerData) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}

	switch c := m.Clazz.(type) {
	case *TLPQInnerData:
		for _, v := range f {
			if f1, ok := v.(func(c *TLPQInnerData) interface{}); ok {
				f1(c)
			}
		}
	case *TLPQInnerDataDc:
		for _, v := range f {
			if f1, ok := v.(func(c *TLPQInnerDataDc) interface{}); ok {
				f1(c)
			}
		}
	case *TLPQInnerDataTemp:
		for _, v := range f {
			if f1, ok := v.(func(c *TLPQInnerDataTemp) interface{}); ok {
				f1(c)
			}
		}
	case *TLPQInnerDataTempDc:
		for _, v := range f {
			if f1, ok := v.(func(c *TLPQInnerDataTempDc) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToPQInnerData <--
func (m *PQInnerData) ToPQInnerData() (*TLPQInnerData, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLPQInnerData); ok {
		return x, true
	}

	return nil, false
}

// ToPQInnerDataDc <--
func (m *PQInnerData) ToPQInnerDataDc() (*TLPQInnerDataDc, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLPQInnerDataDc); ok {
		return x, true
	}

	return nil, false
}

// ToPQInnerDataTemp <--
func (m *PQInnerData) ToPQInnerDataTemp() (*TLPQInnerDataTemp, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLPQInnerDataTemp); ok {
		return x, true
	}

	return nil, false
}

// ToPQInnerDataTempDc <--
func (m *PQInnerData) ToPQInnerDataTempDc() (*TLPQInnerDataTempDc, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLPQInnerDataTempDc); ok {
		return x, true
	}

	return nil, false
}

// ResPQClazz <--
//   - TL_ResPQ
type ResPQClazz = *TLResPQ

func DecodeResPQClazz(d *bin.Decoder) (ResPQClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x5162463:
		x := &TLResPQ{ClazzID: id, ClazzName2: ClazzName_resPQ}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeResPQ - unexpected clazzId: %d", id)
	}

}

// TLResPQ <--
type TLResPQ struct {
	ClazzID                     uint32     `json:"_id"`
	ClazzName2                  string     `json:"_name"`
	Nonce                       bin.Int128 `json:"nonce"`
	ServerNonce                 bin.Int128 `json:"server_nonce"`
	Pq                          string     `json:"pq"`
	ServerPublicKeyFingerprints []int64    `json:"server_public_key_fingerprints"`
}

func MakeTLResPQ(m *TLResPQ) *TLResPQ {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_resPQ

	return m
}

func (m *TLResPQ) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "resPQ", TLObject: m}
	return wrapper.String()
}

// ResPQClazzName <--
func (m *TLResPQ) ResPQClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLResPQ) ClazzName() string {
	return ClazzName_resPQ
}

// ToResPQ <--
func (m *TLResPQ) ToResPQ() *ResPQ {
	if m == nil {
		return nil
	}

	return m

}

// Encode <--
func (m *TLResPQ) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_resPQ, int(layer)); clazzId {
	case 0x5162463:
		x.PutClazzID(0x5162463)

		x.PutInt128(m.Nonce)
		x.PutInt128(m.ServerNonce)
		x.PutString(m.Pq)

		iface.EncodeInt64List(x, m.ServerPublicKeyFingerprints)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_resPQ, layer)
	}
}

// Decode <--
func (m *TLResPQ) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x5162463:
		err = m.Nonce.Decode(d)
		err = m.ServerNonce.Decode(d)
		m.Pq, err = d.String()

		m.ServerPublicKeyFingerprints, err = iface.DecodeInt64List(d)

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// ResPQ <--
type ResPQ = TLResPQ

// ServerDHInnerDataClazz <--
//   - TL_ServerDHInnerData
type ServerDHInnerDataClazz = *TLServerDHInnerData

func DecodeServerDHInnerDataClazz(d *bin.Decoder) (ServerDHInnerDataClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0xb5890dba:
		x := &TLServerDHInnerData{ClazzID: id, ClazzName2: ClazzName_server_DH_inner_data}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeServerDHInnerData - unexpected clazzId: %d", id)
	}

}

// TLServerDHInnerData <--
type TLServerDHInnerData struct {
	ClazzID     uint32     `json:"_id"`
	ClazzName2  string     `json:"_name"`
	Nonce       bin.Int128 `json:"nonce"`
	ServerNonce bin.Int128 `json:"server_nonce"`
	G           int32      `json:"g"`
	DhPrime     string     `json:"dh_prime"`
	GA          string     `json:"g_a"`
	ServerTime  int32      `json:"server_time"`
}

func MakeTLServerDHInnerData(m *TLServerDHInnerData) *TLServerDHInnerData {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_server_DH_inner_data

	return m
}

func (m *TLServerDHInnerData) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "server_DH_inner_data", TLObject: m}
	return wrapper.String()
}

// ServerDHInnerDataClazzName <--
func (m *TLServerDHInnerData) ServerDHInnerDataClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLServerDHInnerData) ClazzName() string {
	return ClazzName_server_DH_inner_data
}

// ToServerDHInnerData <--
func (m *TLServerDHInnerData) ToServerDHInnerData() *ServerDHInnerData {
	if m == nil {
		return nil
	}

	return m

}

// Encode <--
func (m *TLServerDHInnerData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_server_DH_inner_data, int(layer)); clazzId {
	case 0xb5890dba:
		x.PutClazzID(0xb5890dba)

		x.PutInt128(m.Nonce)
		x.PutInt128(m.ServerNonce)
		x.PutInt32(m.G)
		x.PutString(m.DhPrime)
		x.PutString(m.GA)
		x.PutInt32(m.ServerTime)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_server_DH_inner_data, layer)
	}
}

// Decode <--
func (m *TLServerDHInnerData) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xb5890dba:
		err = m.Nonce.Decode(d)
		err = m.ServerNonce.Decode(d)
		m.G, err = d.Int32()
		m.DhPrime, err = d.String()
		m.GA, err = d.String()
		m.ServerTime, err = d.Int32()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// ServerDHInnerData <--
type ServerDHInnerData = TLServerDHInnerData

// ServerDHParamsClazz <--
//   - TL_ServerDHParamsFail
//   - TL_ServerDHParamsOk
type ServerDHParamsClazz interface {
	iface.TLObject
	ServerDHParamsClazzName() string
}

func DecodeServerDHParamsClazz(d *bin.Decoder) (ServerDHParamsClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x79cb045d:
		x := &TLServerDHParamsFail{ClazzID: id, ClazzName2: ClazzName_server_DH_params_fail}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0xd0e8075c:
		x := &TLServerDHParamsOk{ClazzID: id, ClazzName2: ClazzName_server_DH_params_ok}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeServerDHParams - unexpected clazzId: %d", id)
	}

}

// TLServerDHParamsFail <--
type TLServerDHParamsFail struct {
	ClazzID      uint32     `json:"_id"`
	ClazzName2   string     `json:"_name"`
	Nonce        bin.Int128 `json:"nonce"`
	ServerNonce  bin.Int128 `json:"server_nonce"`
	NewNonceHash bin.Int128 `json:"new_nonce_hash"`
}

func MakeTLServerDHParamsFail(m *TLServerDHParamsFail) *TLServerDHParamsFail {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_server_DH_params_fail

	return m
}

func (m *TLServerDHParamsFail) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "server_DH_params_fail", TLObject: m}
	return wrapper.String()
}

// ServerDHParamsClazzName <--
func (m *TLServerDHParamsFail) ServerDHParamsClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLServerDHParamsFail) ClazzName() string {
	return ClazzName_server_DH_params_fail
}

// ToServerDHParams <--
func (m *TLServerDHParamsFail) ToServerDHParams() *ServerDHParams {
	if m == nil {
		return nil
	}

	return &ServerDHParams{Clazz: m}

}

// Encode <--
func (m *TLServerDHParamsFail) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_server_DH_params_fail, int(layer)); clazzId {
	case 0x79cb045d:
		x.PutClazzID(0x79cb045d)

		x.PutInt128(m.Nonce)
		x.PutInt128(m.ServerNonce)
		x.PutInt128(m.NewNonceHash)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_server_DH_params_fail, layer)
	}
}

// Decode <--
func (m *TLServerDHParamsFail) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x79cb045d:
		err = m.Nonce.Decode(d)
		err = m.ServerNonce.Decode(d)
		err = m.NewNonceHash.Decode(d)

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLServerDHParamsOk <--
type TLServerDHParamsOk struct {
	ClazzID         uint32     `json:"_id"`
	ClazzName2      string     `json:"_name"`
	Nonce           bin.Int128 `json:"nonce"`
	ServerNonce     bin.Int128 `json:"server_nonce"`
	EncryptedAnswer string     `json:"encrypted_answer"`
}

func MakeTLServerDHParamsOk(m *TLServerDHParamsOk) *TLServerDHParamsOk {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_server_DH_params_ok

	return m
}

func (m *TLServerDHParamsOk) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "server_DH_params_ok", TLObject: m}
	return wrapper.String()
}

// ServerDHParamsClazzName <--
func (m *TLServerDHParamsOk) ServerDHParamsClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLServerDHParamsOk) ClazzName() string {
	return ClazzName_server_DH_params_ok
}

// ToServerDHParams <--
func (m *TLServerDHParamsOk) ToServerDHParams() *ServerDHParams {
	if m == nil {
		return nil
	}

	return &ServerDHParams{Clazz: m}

}

// Encode <--
func (m *TLServerDHParamsOk) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_server_DH_params_ok, int(layer)); clazzId {
	case 0xd0e8075c:
		x.PutClazzID(0xd0e8075c)

		x.PutInt128(m.Nonce)
		x.PutInt128(m.ServerNonce)
		x.PutString(m.EncryptedAnswer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_server_DH_params_ok, layer)
	}
}

// Decode <--
func (m *TLServerDHParamsOk) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xd0e8075c:
		err = m.Nonce.Decode(d)
		err = m.ServerNonce.Decode(d)
		m.EncryptedAnswer, err = d.String()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// ServerDHParams <--
type ServerDHParams struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz ServerDHParamsClazz `json:"_clazz"`
}

func (m *ServerDHParams) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.ServerDHParamsClazzName()
	}
}

func (m *ServerDHParams) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: m.ClazzName(), TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *ServerDHParams) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("ServerDHParams - invalid Clazz")
}

// Decode <--
func (m *ServerDHParams) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeServerDHParamsClazz(d)
	return
}

// Match <--
func (m *ServerDHParams) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}

	switch c := m.Clazz.(type) {
	case *TLServerDHParamsFail:
		for _, v := range f {
			if f1, ok := v.(func(c *TLServerDHParamsFail) interface{}); ok {
				f1(c)
			}
		}
	case *TLServerDHParamsOk:
		for _, v := range f {
			if f1, ok := v.(func(c *TLServerDHParamsOk) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToServerDHParamsFail <--
func (m *ServerDHParams) ToServerDHParamsFail() (*TLServerDHParamsFail, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLServerDHParamsFail); ok {
		return x, true
	}

	return nil, false
}

// ToServerDHParamsOk <--
func (m *ServerDHParams) ToServerDHParamsOk() (*TLServerDHParamsOk, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLServerDHParamsOk); ok {
		return x, true
	}

	return nil, false
}

// SetClientDHParamsAnswerClazz <--
//   - TL_DhGenOk
//   - TL_DhGenRetry
//   - TL_DhGenFail
type SetClientDHParamsAnswerClazz interface {
	iface.TLObject
	SetClientDHParamsAnswerClazzName() string
}

func DecodeSetClientDHParamsAnswerClazz(d *bin.Decoder) (SetClientDHParamsAnswerClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x3bcbf734:
		x := &TLDhGenOk{ClazzID: id, ClazzName2: ClazzName_dh_gen_ok}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x46dc1fb9:
		x := &TLDhGenRetry{ClazzID: id, ClazzName2: ClazzName_dh_gen_retry}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0xa69dae02:
		x := &TLDhGenFail{ClazzID: id, ClazzName2: ClazzName_dh_gen_fail}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeSetClientDHParamsAnswer - unexpected clazzId: %d", id)
	}

}

// TLDhGenOk <--
type TLDhGenOk struct {
	ClazzID       uint32     `json:"_id"`
	ClazzName2    string     `json:"_name"`
	Nonce         bin.Int128 `json:"nonce"`
	ServerNonce   bin.Int128 `json:"server_nonce"`
	NewNonceHash1 bin.Int128 `json:"new_nonce_hash1"`
}

func MakeTLDhGenOk(m *TLDhGenOk) *TLDhGenOk {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_dh_gen_ok

	return m
}

func (m *TLDhGenOk) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "dh_gen_ok", TLObject: m}
	return wrapper.String()
}

// SetClientDHParamsAnswerClazzName <--
func (m *TLDhGenOk) SetClientDHParamsAnswerClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLDhGenOk) ClazzName() string {
	return ClazzName_dh_gen_ok
}

// ToSetClientDHParamsAnswer <--
func (m *TLDhGenOk) ToSetClientDHParamsAnswer() *SetClientDHParamsAnswer {
	if m == nil {
		return nil
	}

	return &SetClientDHParamsAnswer{Clazz: m}

}

// Encode <--
func (m *TLDhGenOk) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dh_gen_ok, int(layer)); clazzId {
	case 0x3bcbf734:
		x.PutClazzID(0x3bcbf734)

		x.PutInt128(m.Nonce)
		x.PutInt128(m.ServerNonce)
		x.PutInt128(m.NewNonceHash1)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dh_gen_ok, layer)
	}
}

// Decode <--
func (m *TLDhGenOk) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x3bcbf734:
		err = m.Nonce.Decode(d)
		err = m.ServerNonce.Decode(d)
		err = m.NewNonceHash1.Decode(d)

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDhGenRetry <--
type TLDhGenRetry struct {
	ClazzID       uint32     `json:"_id"`
	ClazzName2    string     `json:"_name"`
	Nonce         bin.Int128 `json:"nonce"`
	ServerNonce   bin.Int128 `json:"server_nonce"`
	NewNonceHash2 bin.Int128 `json:"new_nonce_hash2"`
}

func MakeTLDhGenRetry(m *TLDhGenRetry) *TLDhGenRetry {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_dh_gen_retry

	return m
}

func (m *TLDhGenRetry) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "dh_gen_retry", TLObject: m}
	return wrapper.String()
}

// SetClientDHParamsAnswerClazzName <--
func (m *TLDhGenRetry) SetClientDHParamsAnswerClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLDhGenRetry) ClazzName() string {
	return ClazzName_dh_gen_retry
}

// ToSetClientDHParamsAnswer <--
func (m *TLDhGenRetry) ToSetClientDHParamsAnswer() *SetClientDHParamsAnswer {
	if m == nil {
		return nil
	}

	return &SetClientDHParamsAnswer{Clazz: m}

}

// Encode <--
func (m *TLDhGenRetry) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dh_gen_retry, int(layer)); clazzId {
	case 0x46dc1fb9:
		x.PutClazzID(0x46dc1fb9)

		x.PutInt128(m.Nonce)
		x.PutInt128(m.ServerNonce)
		x.PutInt128(m.NewNonceHash2)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dh_gen_retry, layer)
	}
}

// Decode <--
func (m *TLDhGenRetry) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x46dc1fb9:
		err = m.Nonce.Decode(d)
		err = m.ServerNonce.Decode(d)
		err = m.NewNonceHash2.Decode(d)

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDhGenFail <--
type TLDhGenFail struct {
	ClazzID       uint32     `json:"_id"`
	ClazzName2    string     `json:"_name"`
	Nonce         bin.Int128 `json:"nonce"`
	ServerNonce   bin.Int128 `json:"server_nonce"`
	NewNonceHash3 bin.Int128 `json:"new_nonce_hash3"`
}

func MakeTLDhGenFail(m *TLDhGenFail) *TLDhGenFail {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_dh_gen_fail

	return m
}

func (m *TLDhGenFail) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "dh_gen_fail", TLObject: m}
	return wrapper.String()
}

// SetClientDHParamsAnswerClazzName <--
func (m *TLDhGenFail) SetClientDHParamsAnswerClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLDhGenFail) ClazzName() string {
	return ClazzName_dh_gen_fail
}

// ToSetClientDHParamsAnswer <--
func (m *TLDhGenFail) ToSetClientDHParamsAnswer() *SetClientDHParamsAnswer {
	if m == nil {
		return nil
	}

	return &SetClientDHParamsAnswer{Clazz: m}

}

// Encode <--
func (m *TLDhGenFail) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dh_gen_fail, int(layer)); clazzId {
	case 0xa69dae02:
		x.PutClazzID(0xa69dae02)

		x.PutInt128(m.Nonce)
		x.PutInt128(m.ServerNonce)
		x.PutInt128(m.NewNonceHash3)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dh_gen_fail, layer)
	}
}

// Decode <--
func (m *TLDhGenFail) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xa69dae02:
		err = m.Nonce.Decode(d)
		err = m.ServerNonce.Decode(d)
		err = m.NewNonceHash3.Decode(d)

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// SetClientDHParamsAnswer <--
type SetClientDHParamsAnswer struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz SetClientDHParamsAnswerClazz `json:"_clazz"`
}

func (m *SetClientDHParamsAnswer) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.SetClientDHParamsAnswerClazzName()
	}
}

func (m *SetClientDHParamsAnswer) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: m.ClazzName(), TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *SetClientDHParamsAnswer) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("SetClientDHParamsAnswer - invalid Clazz")
}

// Decode <--
func (m *SetClientDHParamsAnswer) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeSetClientDHParamsAnswerClazz(d)
	return
}

// Match <--
func (m *SetClientDHParamsAnswer) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}

	switch c := m.Clazz.(type) {
	case *TLDhGenOk:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDhGenOk) interface{}); ok {
				f1(c)
			}
		}
	case *TLDhGenRetry:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDhGenRetry) interface{}); ok {
				f1(c)
			}
		}
	case *TLDhGenFail:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDhGenFail) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToDhGenOk <--
func (m *SetClientDHParamsAnswer) ToDhGenOk() (*TLDhGenOk, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLDhGenOk); ok {
		return x, true
	}

	return nil, false
}

// ToDhGenRetry <--
func (m *SetClientDHParamsAnswer) ToDhGenRetry() (*TLDhGenRetry, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLDhGenRetry); ok {
		return x, true
	}

	return nil, false
}

// ToDhGenFail <--
func (m *SetClientDHParamsAnswer) ToDhGenFail() (*TLDhGenFail, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLDhGenFail); ok {
		return x, true
	}

	return nil, false
}

// AccessPointRuleClazz <--
//   - TL_AccessPointRule
type AccessPointRuleClazz = *TLAccessPointRule

func DecodeAccessPointRuleClazz(d *bin.Decoder) (AccessPointRuleClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x4679b65f:
		x := &TLAccessPointRule{ClazzID: id, ClazzName2: ClazzName_accessPointRule}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeAccessPointRule - unexpected clazzId: %d", id)
	}

}

// TLAccessPointRule <--
type TLAccessPointRule struct {
	ClazzID          uint32        `json:"_id"`
	ClazzName2       string        `json:"_name"`
	PhonePrefixRules string        `json:"phone_prefix_rules"`
	DcId             int32         `json:"dc_id"`
	Ips              []IpPortClazz `json:"ips"`
}

func MakeTLAccessPointRule(m *TLAccessPointRule) *TLAccessPointRule {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_accessPointRule

	return m
}

func (m *TLAccessPointRule) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "accessPointRule", TLObject: m}
	return wrapper.String()
}

// AccessPointRuleClazzName <--
func (m *TLAccessPointRule) AccessPointRuleClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLAccessPointRule) ClazzName() string {
	return ClazzName_accessPointRule
}

// ToAccessPointRule <--
func (m *TLAccessPointRule) ToAccessPointRule() *AccessPointRule {
	if m == nil {
		return nil
	}

	return m

}

// Encode <--
func (m *TLAccessPointRule) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_accessPointRule, int(layer)); clazzId {
	case 0x4679b65f:
		x.PutClazzID(0x4679b65f)

		x.PutString(m.PhonePrefixRules)
		x.PutInt32(m.DcId)
		// x.PutClazzID(iface.ClazzID_vector)
		x.PutInt(len(m.Ips))
		for _, v := range m.Ips {
			_ = v.Encode(x, layer)
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_accessPointRule, layer)
	}
}

// Decode <--
func (m *TLAccessPointRule) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x4679b65f:
		m.PhonePrefixRules, err = d.String()
		m.DcId, err = d.Int32()
		// c2, err2 := d.ClazzID()
		// if c2 != int32(iface.ClazzID_vector) {
		//     err2 = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 2, c2)
		//     return err2
		// }
		l2, err2 := d.Int()
		if err2 != nil {
			return err2
		}
		v2 := make([]IpPortClazz, l2)
		for i := 0; i < l2; i++ {
			v2[i], err2 = DecodeIpPortClazz(d)
			if err2 != nil {
				return err2
			}
		}
		m.Ips = v2

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// AccessPointRule <--
type AccessPointRule = TLAccessPointRule

// BadMsgNotificationClazz <--
//   - TL_BadMsgNotification
//   - TL_BadServerSalt
type BadMsgNotificationClazz interface {
	iface.TLObject
	BadMsgNotificationClazzName() string
}

func DecodeBadMsgNotificationClazz(d *bin.Decoder) (BadMsgNotificationClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0xa7eff811:
		x := &TLBadMsgNotification{ClazzID: id, ClazzName2: ClazzName_bad_msg_notification}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0xedab447b:
		x := &TLBadServerSalt{ClazzID: id, ClazzName2: ClazzName_bad_server_salt}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeBadMsgNotification - unexpected clazzId: %d", id)
	}

}

// TLBadMsgNotification <--
type TLBadMsgNotification struct {
	ClazzID     uint32 `json:"_id"`
	ClazzName2  string `json:"_name"`
	BadMsgId    int64  `json:"bad_msg_id"`
	BadMsgSeqno int32  `json:"bad_msg_seqno"`
	ErrorCode   int32  `json:"error_code"`
}

func MakeTLBadMsgNotification(m *TLBadMsgNotification) *TLBadMsgNotification {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_bad_msg_notification

	return m
}

func (m *TLBadMsgNotification) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "bad_msg_notification", TLObject: m}
	return wrapper.String()
}

// BadMsgNotificationClazzName <--
func (m *TLBadMsgNotification) BadMsgNotificationClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLBadMsgNotification) ClazzName() string {
	return ClazzName_bad_msg_notification
}

// ToBadMsgNotification <--
func (m *TLBadMsgNotification) ToBadMsgNotification() *BadMsgNotification {
	if m == nil {
		return nil
	}

	return &BadMsgNotification{Clazz: m}

}

// Encode <--
func (m *TLBadMsgNotification) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_bad_msg_notification, int(layer)); clazzId {
	case 0xa7eff811:
		x.PutClazzID(0xa7eff811)

		x.PutInt64(m.BadMsgId)
		x.PutInt32(m.BadMsgSeqno)
		x.PutInt32(m.ErrorCode)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_bad_msg_notification, layer)
	}
}

// Decode <--
func (m *TLBadMsgNotification) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xa7eff811:
		m.BadMsgId, err = d.Int64()
		m.BadMsgSeqno, err = d.Int32()
		m.ErrorCode, err = d.Int32()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLBadServerSalt <--
type TLBadServerSalt struct {
	ClazzID       uint32 `json:"_id"`
	ClazzName2    string `json:"_name"`
	BadMsgId      int64  `json:"bad_msg_id"`
	BadMsgSeqno   int32  `json:"bad_msg_seqno"`
	ErrorCode     int32  `json:"error_code"`
	NewServerSalt int64  `json:"new_server_salt"`
}

func MakeTLBadServerSalt(m *TLBadServerSalt) *TLBadServerSalt {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_bad_server_salt

	return m
}

func (m *TLBadServerSalt) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "bad_server_salt", TLObject: m}
	return wrapper.String()
}

// BadMsgNotificationClazzName <--
func (m *TLBadServerSalt) BadMsgNotificationClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLBadServerSalt) ClazzName() string {
	return ClazzName_bad_server_salt
}

// ToBadMsgNotification <--
func (m *TLBadServerSalt) ToBadMsgNotification() *BadMsgNotification {
	if m == nil {
		return nil
	}

	return &BadMsgNotification{Clazz: m}

}

// Encode <--
func (m *TLBadServerSalt) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_bad_server_salt, int(layer)); clazzId {
	case 0xedab447b:
		x.PutClazzID(0xedab447b)

		x.PutInt64(m.BadMsgId)
		x.PutInt32(m.BadMsgSeqno)
		x.PutInt32(m.ErrorCode)
		x.PutInt64(m.NewServerSalt)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_bad_server_salt, layer)
	}
}

// Decode <--
func (m *TLBadServerSalt) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xedab447b:
		m.BadMsgId, err = d.Int64()
		m.BadMsgSeqno, err = d.Int32()
		m.ErrorCode, err = d.Int32()
		m.NewServerSalt, err = d.Int64()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// BadMsgNotification <--
type BadMsgNotification struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz BadMsgNotificationClazz `json:"_clazz"`
}

func (m *BadMsgNotification) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.BadMsgNotificationClazzName()
	}
}

func (m *BadMsgNotification) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: m.ClazzName(), TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *BadMsgNotification) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("BadMsgNotification - invalid Clazz")
}

// Decode <--
func (m *BadMsgNotification) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeBadMsgNotificationClazz(d)
	return
}

// Match <--
func (m *BadMsgNotification) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}

	switch c := m.Clazz.(type) {
	case *TLBadMsgNotification:
		for _, v := range f {
			if f1, ok := v.(func(c *TLBadMsgNotification) interface{}); ok {
				f1(c)
			}
		}
	case *TLBadServerSalt:
		for _, v := range f {
			if f1, ok := v.(func(c *TLBadServerSalt) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToBadMsgNotification <--
func (m *BadMsgNotification) ToBadMsgNotification() (*TLBadMsgNotification, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLBadMsgNotification); ok {
		return x, true
	}

	return nil, false
}

// ToBadServerSalt <--
func (m *BadMsgNotification) ToBadServerSalt() (*TLBadServerSalt, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLBadServerSalt); ok {
		return x, true
	}

	return nil, false
}

// DestroySessionResClazz <--
//   - TL_DestroySessionOk
//   - TL_DestroySessionNone
type DestroySessionResClazz interface {
	iface.TLObject
	DestroySessionResClazzName() string
}

func DecodeDestroySessionResClazz(d *bin.Decoder) (DestroySessionResClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0xe22045fc:
		x := &TLDestroySessionOk{ClazzID: id, ClazzName2: ClazzName_destroy_session_ok}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x62d350c9:
		x := &TLDestroySessionNone{ClazzID: id, ClazzName2: ClazzName_destroy_session_none}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeDestroySessionRes - unexpected clazzId: %d", id)
	}

}

// TLDestroySessionOk <--
type TLDestroySessionOk struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	SessionId  int64  `json:"session_id"`
}

func MakeTLDestroySessionOk(m *TLDestroySessionOk) *TLDestroySessionOk {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_destroy_session_ok

	return m
}

func (m *TLDestroySessionOk) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "destroy_session_ok", TLObject: m}
	return wrapper.String()
}

// DestroySessionResClazzName <--
func (m *TLDestroySessionOk) DestroySessionResClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLDestroySessionOk) ClazzName() string {
	return ClazzName_destroy_session_ok
}

// ToDestroySessionRes <--
func (m *TLDestroySessionOk) ToDestroySessionRes() *DestroySessionRes {
	if m == nil {
		return nil
	}

	return &DestroySessionRes{Clazz: m}

}

// Encode <--
func (m *TLDestroySessionOk) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_destroy_session_ok, int(layer)); clazzId {
	case 0xe22045fc:
		x.PutClazzID(0xe22045fc)

		x.PutInt64(m.SessionId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_destroy_session_ok, layer)
	}
}

// Decode <--
func (m *TLDestroySessionOk) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xe22045fc:
		m.SessionId, err = d.Int64()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDestroySessionNone <--
type TLDestroySessionNone struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	SessionId  int64  `json:"session_id"`
}

func MakeTLDestroySessionNone(m *TLDestroySessionNone) *TLDestroySessionNone {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_destroy_session_none

	return m
}

func (m *TLDestroySessionNone) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "destroy_session_none", TLObject: m}
	return wrapper.String()
}

// DestroySessionResClazzName <--
func (m *TLDestroySessionNone) DestroySessionResClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLDestroySessionNone) ClazzName() string {
	return ClazzName_destroy_session_none
}

// ToDestroySessionRes <--
func (m *TLDestroySessionNone) ToDestroySessionRes() *DestroySessionRes {
	if m == nil {
		return nil
	}

	return &DestroySessionRes{Clazz: m}

}

// Encode <--
func (m *TLDestroySessionNone) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_destroy_session_none, int(layer)); clazzId {
	case 0x62d350c9:
		x.PutClazzID(0x62d350c9)

		x.PutInt64(m.SessionId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_destroy_session_none, layer)
	}
}

// Decode <--
func (m *TLDestroySessionNone) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x62d350c9:
		m.SessionId, err = d.Int64()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// DestroySessionRes <--
type DestroySessionRes struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz DestroySessionResClazz `json:"_clazz"`
}

func (m *DestroySessionRes) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.DestroySessionResClazzName()
	}
}

func (m *DestroySessionRes) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: m.ClazzName(), TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *DestroySessionRes) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("DestroySessionRes - invalid Clazz")
}

// Decode <--
func (m *DestroySessionRes) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeDestroySessionResClazz(d)
	return
}

// Match <--
func (m *DestroySessionRes) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}

	switch c := m.Clazz.(type) {
	case *TLDestroySessionOk:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDestroySessionOk) interface{}); ok {
				f1(c)
			}
		}
	case *TLDestroySessionNone:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDestroySessionNone) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToDestroySessionOk <--
func (m *DestroySessionRes) ToDestroySessionOk() (*TLDestroySessionOk, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLDestroySessionOk); ok {
		return x, true
	}

	return nil, false
}

// ToDestroySessionNone <--
func (m *DestroySessionRes) ToDestroySessionNone() (*TLDestroySessionNone, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLDestroySessionNone); ok {
		return x, true
	}

	return nil, false
}

// FutureSaltClazz <--
//   - TL_FutureSalt
type FutureSaltClazz = *TLFutureSalt

func DecodeFutureSaltClazz(d *bin.Decoder) (FutureSaltClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x949d9dc:
		x := &TLFutureSalt{ClazzID: id, ClazzName2: ClazzName_future_salt}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeFutureSalt - unexpected clazzId: %d", id)
	}

}

// TLFutureSalt <--
type TLFutureSalt struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	ValidSince int32  `json:"valid_since"`
	ValidUntil int32  `json:"valid_until"`
	Salt       int64  `json:"salt"`
}

func MakeTLFutureSalt(m *TLFutureSalt) *TLFutureSalt {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_future_salt

	return m
}

func (m *TLFutureSalt) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "future_salt", TLObject: m}
	return wrapper.String()
}

// FutureSaltClazzName <--
func (m *TLFutureSalt) FutureSaltClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLFutureSalt) ClazzName() string {
	return ClazzName_future_salt
}

// ToFutureSalt <--
func (m *TLFutureSalt) ToFutureSalt() *FutureSalt {
	if m == nil {
		return nil
	}

	return m

}

// Encode <--
func (m *TLFutureSalt) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_future_salt, int(layer)); clazzId {
	case 0x949d9dc:
		x.PutClazzID(0x949d9dc)

		x.PutInt32(m.ValidSince)
		x.PutInt32(m.ValidUntil)
		x.PutInt64(m.Salt)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_future_salt, layer)
	}
}

// Decode <--
func (m *TLFutureSalt) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x949d9dc:
		m.ValidSince, err = d.Int32()
		m.ValidUntil, err = d.Int32()
		m.Salt, err = d.Int64()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// FutureSalt <--
type FutureSalt = TLFutureSalt

// FutureSaltsClazz <--
//   - TL_FutureSalts
type FutureSaltsClazz = *TLFutureSalts

func DecodeFutureSaltsClazz(d *bin.Decoder) (FutureSaltsClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0xae500895:
		x := &TLFutureSalts{ClazzID: id, ClazzName2: ClazzName_future_salts}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeFutureSalts - unexpected clazzId: %d", id)
	}

}

// TLFutureSalts <--
type TLFutureSalts struct {
	ClazzID    uint32          `json:"_id"`
	ClazzName2 string          `json:"_name"`
	ReqMsgId   int64           `json:"req_msg_id"`
	Now        int32           `json:"now"`
	Salts      []*TLFutureSalt `json:"salts"`
}

func MakeTLFutureSalts(m *TLFutureSalts) *TLFutureSalts {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_future_salts

	return m
}

func (m *TLFutureSalts) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "future_salts", TLObject: m}
	return wrapper.String()
}

// FutureSaltsClazzName <--
func (m *TLFutureSalts) FutureSaltsClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLFutureSalts) ClazzName() string {
	return ClazzName_future_salts
}

// ToFutureSalts <--
func (m *TLFutureSalts) ToFutureSalts() *FutureSalts {
	if m == nil {
		return nil
	}

	return m

}

// Encode <--
func (m *TLFutureSalts) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_future_salts, int(layer)); clazzId {
	case 0xae500895:
		x.PutClazzID(0xae500895)

		x.PutInt64(m.ReqMsgId)
		x.PutInt32(m.Now)
		// x.PutClazzID(iface.ClazzID_vector)
		x.PutInt(len(m.Salts))
		for _, v := range m.Salts {
			_ = v.Encode(x, layer)
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_future_salts, layer)
	}
}

// Decode <--
func (m *TLFutureSalts) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xae500895:
		m.ReqMsgId, err = d.Int64()
		m.Now, err = d.Int32()
		// c2, err2 := d.ClazzID()
		// if c2 != int32(iface.ClazzID_vector) {
		//     err2 = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 2, c2)
		//     return err2
		// }
		l2, err2 := d.Int()
		if err2 != nil {
			return err2
		}
		v2 := make([]*TLFutureSalt, l2)
		for i := 0; i < l2; i++ {
			v2[i] = &TLFutureSalt{ClazzID: ClazzID_future_salts}
			err2 = v2[i].Decode(d)
			if err2 != nil {
				return err2
			}
		}
		m.Salts = v2

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// FutureSalts <--
type FutureSalts = TLFutureSalts

// HelpConfigSimpleClazz <--
//   - TL_HelpConfigSimple
type HelpConfigSimpleClazz = *TLHelpConfigSimple

func DecodeHelpConfigSimpleClazz(d *bin.Decoder) (HelpConfigSimpleClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x5a592a6c:
		x := &TLHelpConfigSimple{ClazzID: id, ClazzName2: ClazzName_help_configSimple}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeHelpConfigSimple - unexpected clazzId: %d", id)
	}

}

// TLHelpConfigSimple <--
type TLHelpConfigSimple struct {
	ClazzID    uint32                 `json:"_id"`
	ClazzName2 string                 `json:"_name"`
	Date       int32                  `json:"date"`
	Expires    int32                  `json:"expires"`
	Rules      []AccessPointRuleClazz `json:"rules"`
}

func MakeTLHelpConfigSimple(m *TLHelpConfigSimple) *TLHelpConfigSimple {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_help_configSimple

	return m
}

func (m *TLHelpConfigSimple) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "help_configSimple", TLObject: m}
	return wrapper.String()
}

// HelpConfigSimpleClazzName <--
func (m *TLHelpConfigSimple) HelpConfigSimpleClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLHelpConfigSimple) ClazzName() string {
	return ClazzName_help_configSimple
}

// ToHelpConfigSimple <--
func (m *TLHelpConfigSimple) ToHelpConfigSimple() *HelpConfigSimple {
	if m == nil {
		return nil
	}

	return m

}

// Encode <--
func (m *TLHelpConfigSimple) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_help_configSimple, int(layer)); clazzId {
	case 0x5a592a6c:
		x.PutClazzID(0x5a592a6c)

		x.PutInt32(m.Date)
		x.PutInt32(m.Expires)
		// x.PutClazzID(iface.ClazzID_vector)
		x.PutInt(len(m.Rules))
		for _, v := range m.Rules {
			_ = v.Encode(x, layer)
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_help_configSimple, layer)
	}
}

// Decode <--
func (m *TLHelpConfigSimple) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x5a592a6c:
		m.Date, err = d.Int32()
		m.Expires, err = d.Int32()
		// c2, err2 := d.ClazzID()
		// if c2 != int32(iface.ClazzID_vector) {
		//     err2 = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 2, c2)
		//     return err2
		// }
		l2, err2 := d.Int()
		if err2 != nil {
			return err2
		}
		v2 := make([]AccessPointRuleClazz, l2)
		for i := 0; i < l2; i++ {
			v2[i], err2 = DecodeAccessPointRuleClazz(d)
			if err2 != nil {
				return err2
			}
		}
		m.Rules = v2

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// HelpConfigSimple <--
type HelpConfigSimple = TLHelpConfigSimple

// HttpWaitClazz <--
//   - TL_HttpWait
type HttpWaitClazz = *TLHttpWait

func DecodeHttpWaitClazz(d *bin.Decoder) (HttpWaitClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x9299359f:
		x := &TLHttpWait{ClazzID: id, ClazzName2: ClazzName_http_wait}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeHttpWait - unexpected clazzId: %d", id)
	}

}

// TLHttpWait <--
type TLHttpWait struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	MaxDelay   int32  `json:"max_delay"`
	WaitAfter  int32  `json:"wait_after"`
	MaxWait    int32  `json:"max_wait"`
}

func MakeTLHttpWait(m *TLHttpWait) *TLHttpWait {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_http_wait

	return m
}

func (m *TLHttpWait) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "http_wait", TLObject: m}
	return wrapper.String()
}

// HttpWaitClazzName <--
func (m *TLHttpWait) HttpWaitClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLHttpWait) ClazzName() string {
	return ClazzName_http_wait
}

// ToHttpWait <--
func (m *TLHttpWait) ToHttpWait() *HttpWait {
	if m == nil {
		return nil
	}

	return m

}

// Encode <--
func (m *TLHttpWait) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_http_wait, int(layer)); clazzId {
	case 0x9299359f:
		x.PutClazzID(0x9299359f)

		x.PutInt32(m.MaxDelay)
		x.PutInt32(m.WaitAfter)
		x.PutInt32(m.MaxWait)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_http_wait, layer)
	}
}

// Decode <--
func (m *TLHttpWait) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x9299359f:
		m.MaxDelay, err = d.Int32()
		m.WaitAfter, err = d.Int32()
		m.MaxWait, err = d.Int32()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// HttpWait <--
type HttpWait = TLHttpWait

// IpPortClazz <--
//   - TL_IpPort
//   - TL_IpPortSecret
type IpPortClazz interface {
	iface.TLObject
	IpPortClazzName() string
}

func DecodeIpPortClazz(d *bin.Decoder) (IpPortClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0xd433ad73:
		x := &TLIpPort{ClazzID: id, ClazzName2: ClazzName_ipPort}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x37982646:
		x := &TLIpPortSecret{ClazzID: id, ClazzName2: ClazzName_ipPortSecret}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeIpPort - unexpected clazzId: %d", id)
	}

}

// TLIpPort <--
type TLIpPort struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Ipv4       int32  `json:"ipv4"`
	Port       int32  `json:"port"`
}

func MakeTLIpPort(m *TLIpPort) *TLIpPort {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_ipPort

	return m
}

func (m *TLIpPort) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "ipPort", TLObject: m}
	return wrapper.String()
}

// IpPortClazzName <--
func (m *TLIpPort) IpPortClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLIpPort) ClazzName() string {
	return ClazzName_ipPort
}

// ToIpPort <--
func (m *TLIpPort) ToIpPort() *IpPort {
	if m == nil {
		return nil
	}

	return &IpPort{Clazz: m}

}

// Encode <--
func (m *TLIpPort) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_ipPort, int(layer)); clazzId {
	case 0xd433ad73:
		x.PutClazzID(0xd433ad73)

		x.PutInt32(m.Ipv4)
		x.PutInt32(m.Port)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_ipPort, layer)
	}
}

// Decode <--
func (m *TLIpPort) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xd433ad73:
		m.Ipv4, err = d.Int32()
		m.Port, err = d.Int32()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLIpPortSecret <--
type TLIpPortSecret struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Ipv4       int32  `json:"ipv4"`
	Port       int32  `json:"port"`
	Secret     []byte `json:"secret"`
}

func MakeTLIpPortSecret(m *TLIpPortSecret) *TLIpPortSecret {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_ipPortSecret

	return m
}

func (m *TLIpPortSecret) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "ipPortSecret", TLObject: m}
	return wrapper.String()
}

// IpPortClazzName <--
func (m *TLIpPortSecret) IpPortClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLIpPortSecret) ClazzName() string {
	return ClazzName_ipPortSecret
}

// ToIpPort <--
func (m *TLIpPortSecret) ToIpPort() *IpPort {
	if m == nil {
		return nil
	}

	return &IpPort{Clazz: m}

}

// Encode <--
func (m *TLIpPortSecret) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_ipPortSecret, int(layer)); clazzId {
	case 0x37982646:
		x.PutClazzID(0x37982646)

		x.PutInt32(m.Ipv4)
		x.PutInt32(m.Port)
		x.PutBytes(m.Secret)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_ipPortSecret, layer)
	}
}

// Decode <--
func (m *TLIpPortSecret) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x37982646:
		m.Ipv4, err = d.Int32()
		m.Port, err = d.Int32()
		m.Secret, err = d.Bytes()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// IpPort <--
type IpPort struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz IpPortClazz `json:"_clazz"`
}

func (m *IpPort) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.IpPortClazzName()
	}
}

func (m *IpPort) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: m.ClazzName(), TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *IpPort) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("IpPort - invalid Clazz")
}

// Decode <--
func (m *IpPort) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeIpPortClazz(d)
	return
}

// Match <--
func (m *IpPort) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}

	switch c := m.Clazz.(type) {
	case *TLIpPort:
		for _, v := range f {
			if f1, ok := v.(func(c *TLIpPort) interface{}); ok {
				f1(c)
			}
		}
	case *TLIpPortSecret:
		for _, v := range f {
			if f1, ok := v.(func(c *TLIpPortSecret) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToIpPort <--
func (m *IpPort) ToIpPort() (*TLIpPort, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLIpPort); ok {
		return x, true
	}

	return nil, false
}

// ToIpPortSecret <--
func (m *IpPort) ToIpPortSecret() (*TLIpPortSecret, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLIpPortSecret); ok {
		return x, true
	}

	return nil, false
}

// MsgDetailedInfoClazz <--
//   - TL_MsgDetailedInfo
//   - TL_MsgNewDetailedInfo
type MsgDetailedInfoClazz interface {
	iface.TLObject
	MsgDetailedInfoClazzName() string
}

func DecodeMsgDetailedInfoClazz(d *bin.Decoder) (MsgDetailedInfoClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x276d3ec6:
		x := &TLMsgDetailedInfo{ClazzID: id, ClazzName2: ClazzName_msg_detailed_info}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x809db6df:
		x := &TLMsgNewDetailedInfo{ClazzID: id, ClazzName2: ClazzName_msg_new_detailed_info}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeMsgDetailedInfo - unexpected clazzId: %d", id)
	}

}

// TLMsgDetailedInfo <--
type TLMsgDetailedInfo struct {
	ClazzID     uint32 `json:"_id"`
	ClazzName2  string `json:"_name"`
	MsgId       int64  `json:"msg_id"`
	AnswerMsgId int64  `json:"answer_msg_id"`
	Bytes       int32  `json:"bytes"`
	Status      int32  `json:"status"`
}

func MakeTLMsgDetailedInfo(m *TLMsgDetailedInfo) *TLMsgDetailedInfo {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_msg_detailed_info

	return m
}

func (m *TLMsgDetailedInfo) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "msg_detailed_info", TLObject: m}
	return wrapper.String()
}

// MsgDetailedInfoClazzName <--
func (m *TLMsgDetailedInfo) MsgDetailedInfoClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLMsgDetailedInfo) ClazzName() string {
	return ClazzName_msg_detailed_info
}

// ToMsgDetailedInfo <--
func (m *TLMsgDetailedInfo) ToMsgDetailedInfo() *MsgDetailedInfo {
	if m == nil {
		return nil
	}

	return &MsgDetailedInfo{Clazz: m}

}

// Encode <--
func (m *TLMsgDetailedInfo) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_msg_detailed_info, int(layer)); clazzId {
	case 0x276d3ec6:
		x.PutClazzID(0x276d3ec6)

		x.PutInt64(m.MsgId)
		x.PutInt64(m.AnswerMsgId)
		x.PutInt32(m.Bytes)
		x.PutInt32(m.Status)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_detailed_info, layer)
	}
}

// Decode <--
func (m *TLMsgDetailedInfo) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x276d3ec6:
		m.MsgId, err = d.Int64()
		m.AnswerMsgId, err = d.Int64()
		m.Bytes, err = d.Int32()
		m.Status, err = d.Int32()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMsgNewDetailedInfo <--
type TLMsgNewDetailedInfo struct {
	ClazzID     uint32 `json:"_id"`
	ClazzName2  string `json:"_name"`
	AnswerMsgId int64  `json:"answer_msg_id"`
	Bytes       int32  `json:"bytes"`
	Status      int32  `json:"status"`
}

func MakeTLMsgNewDetailedInfo(m *TLMsgNewDetailedInfo) *TLMsgNewDetailedInfo {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_msg_new_detailed_info

	return m
}

func (m *TLMsgNewDetailedInfo) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "msg_new_detailed_info", TLObject: m}
	return wrapper.String()
}

// MsgDetailedInfoClazzName <--
func (m *TLMsgNewDetailedInfo) MsgDetailedInfoClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLMsgNewDetailedInfo) ClazzName() string {
	return ClazzName_msg_new_detailed_info
}

// ToMsgDetailedInfo <--
func (m *TLMsgNewDetailedInfo) ToMsgDetailedInfo() *MsgDetailedInfo {
	if m == nil {
		return nil
	}

	return &MsgDetailedInfo{Clazz: m}

}

// Encode <--
func (m *TLMsgNewDetailedInfo) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_msg_new_detailed_info, int(layer)); clazzId {
	case 0x809db6df:
		x.PutClazzID(0x809db6df)

		x.PutInt64(m.AnswerMsgId)
		x.PutInt32(m.Bytes)
		x.PutInt32(m.Status)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_new_detailed_info, layer)
	}
}

// Decode <--
func (m *TLMsgNewDetailedInfo) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x809db6df:
		m.AnswerMsgId, err = d.Int64()
		m.Bytes, err = d.Int32()
		m.Status, err = d.Int32()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// MsgDetailedInfo <--
type MsgDetailedInfo struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz MsgDetailedInfoClazz `json:"_clazz"`
}

func (m *MsgDetailedInfo) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.MsgDetailedInfoClazzName()
	}
}

func (m *MsgDetailedInfo) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: m.ClazzName(), TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *MsgDetailedInfo) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("MsgDetailedInfo - invalid Clazz")
}

// Decode <--
func (m *MsgDetailedInfo) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeMsgDetailedInfoClazz(d)
	return
}

// Match <--
func (m *MsgDetailedInfo) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}

	switch c := m.Clazz.(type) {
	case *TLMsgDetailedInfo:
		for _, v := range f {
			if f1, ok := v.(func(c *TLMsgDetailedInfo) interface{}); ok {
				f1(c)
			}
		}
	case *TLMsgNewDetailedInfo:
		for _, v := range f {
			if f1, ok := v.(func(c *TLMsgNewDetailedInfo) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToMsgDetailedInfo <--
func (m *MsgDetailedInfo) ToMsgDetailedInfo() (*TLMsgDetailedInfo, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLMsgDetailedInfo); ok {
		return x, true
	}

	return nil, false
}

// ToMsgNewDetailedInfo <--
func (m *MsgDetailedInfo) ToMsgNewDetailedInfo() (*TLMsgNewDetailedInfo, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLMsgNewDetailedInfo); ok {
		return x, true
	}

	return nil, false
}

// MsgResendReqClazz <--
//   - TL_MsgResendReq
type MsgResendReqClazz = *TLMsgResendReq

func DecodeMsgResendReqClazz(d *bin.Decoder) (MsgResendReqClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x7d861a08:
		x := &TLMsgResendReq{ClazzID: id, ClazzName2: ClazzName_msg_resend_req}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeMsgResendReq - unexpected clazzId: %d", id)
	}

}

// TLMsgResendReq <--
type TLMsgResendReq struct {
	ClazzID    uint32  `json:"_id"`
	ClazzName2 string  `json:"_name"`
	MsgIds     []int64 `json:"msg_ids"`
}

func MakeTLMsgResendReq(m *TLMsgResendReq) *TLMsgResendReq {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_msg_resend_req

	return m
}

func (m *TLMsgResendReq) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "msg_resend_req", TLObject: m}
	return wrapper.String()
}

// MsgResendReqClazzName <--
func (m *TLMsgResendReq) MsgResendReqClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLMsgResendReq) ClazzName() string {
	return ClazzName_msg_resend_req
}

// ToMsgResendReq <--
func (m *TLMsgResendReq) ToMsgResendReq() *MsgResendReq {
	if m == nil {
		return nil
	}

	return m

}

// Encode <--
func (m *TLMsgResendReq) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_msg_resend_req, int(layer)); clazzId {
	case 0x7d861a08:
		x.PutClazzID(0x7d861a08)

		iface.EncodeInt64List(x, m.MsgIds)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_resend_req, layer)
	}
}

// Decode <--
func (m *TLMsgResendReq) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x7d861a08:

		m.MsgIds, err = iface.DecodeInt64List(d)

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// MsgResendReq <--
type MsgResendReq = TLMsgResendReq

// MsgsAckClazz <--
//   - TL_MsgsAck
type MsgsAckClazz = *TLMsgsAck

func DecodeMsgsAckClazz(d *bin.Decoder) (MsgsAckClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x62d6b459:
		x := &TLMsgsAck{ClazzID: id, ClazzName2: ClazzName_msgs_ack}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeMsgsAck - unexpected clazzId: %d", id)
	}

}

// TLMsgsAck <--
type TLMsgsAck struct {
	ClazzID    uint32  `json:"_id"`
	ClazzName2 string  `json:"_name"`
	MsgIds     []int64 `json:"msg_ids"`
}

func MakeTLMsgsAck(m *TLMsgsAck) *TLMsgsAck {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_msgs_ack

	return m
}

func (m *TLMsgsAck) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "msgs_ack", TLObject: m}
	return wrapper.String()
}

// MsgsAckClazzName <--
func (m *TLMsgsAck) MsgsAckClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLMsgsAck) ClazzName() string {
	return ClazzName_msgs_ack
}

// ToMsgsAck <--
func (m *TLMsgsAck) ToMsgsAck() *MsgsAck {
	if m == nil {
		return nil
	}

	return m

}

// Encode <--
func (m *TLMsgsAck) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_msgs_ack, int(layer)); clazzId {
	case 0x62d6b459:
		x.PutClazzID(0x62d6b459)

		iface.EncodeInt64List(x, m.MsgIds)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msgs_ack, layer)
	}
}

// Decode <--
func (m *TLMsgsAck) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x62d6b459:

		m.MsgIds, err = iface.DecodeInt64List(d)

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// MsgsAck <--
type MsgsAck = TLMsgsAck

// MsgsAllInfoClazz <--
//   - TL_MsgsAllInfo
type MsgsAllInfoClazz = *TLMsgsAllInfo

func DecodeMsgsAllInfoClazz(d *bin.Decoder) (MsgsAllInfoClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x8cc0d131:
		x := &TLMsgsAllInfo{ClazzID: id, ClazzName2: ClazzName_msgs_all_info}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeMsgsAllInfo - unexpected clazzId: %d", id)
	}

}

// TLMsgsAllInfo <--
type TLMsgsAllInfo struct {
	ClazzID    uint32  `json:"_id"`
	ClazzName2 string  `json:"_name"`
	MsgIds     []int64 `json:"msg_ids"`
	Info       string  `json:"info"`
}

func MakeTLMsgsAllInfo(m *TLMsgsAllInfo) *TLMsgsAllInfo {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_msgs_all_info

	return m
}

func (m *TLMsgsAllInfo) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "msgs_all_info", TLObject: m}
	return wrapper.String()
}

// MsgsAllInfoClazzName <--
func (m *TLMsgsAllInfo) MsgsAllInfoClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLMsgsAllInfo) ClazzName() string {
	return ClazzName_msgs_all_info
}

// ToMsgsAllInfo <--
func (m *TLMsgsAllInfo) ToMsgsAllInfo() *MsgsAllInfo {
	if m == nil {
		return nil
	}

	return m

}

// Encode <--
func (m *TLMsgsAllInfo) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_msgs_all_info, int(layer)); clazzId {
	case 0x8cc0d131:
		x.PutClazzID(0x8cc0d131)

		iface.EncodeInt64List(x, m.MsgIds)

		x.PutString(m.Info)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msgs_all_info, layer)
	}
}

// Decode <--
func (m *TLMsgsAllInfo) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x8cc0d131:

		m.MsgIds, err = iface.DecodeInt64List(d)

		m.Info, err = d.String()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// MsgsAllInfo <--
type MsgsAllInfo = TLMsgsAllInfo

// MsgsStateInfoClazz <--
//   - TL_MsgsStateInfo
type MsgsStateInfoClazz = *TLMsgsStateInfo

func DecodeMsgsStateInfoClazz(d *bin.Decoder) (MsgsStateInfoClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x4deb57d:
		x := &TLMsgsStateInfo{ClazzID: id, ClazzName2: ClazzName_msgs_state_info}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeMsgsStateInfo - unexpected clazzId: %d", id)
	}

}

// TLMsgsStateInfo <--
type TLMsgsStateInfo struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	ReqMsgId   int64  `json:"req_msg_id"`
	Info       string `json:"info"`
}

func MakeTLMsgsStateInfo(m *TLMsgsStateInfo) *TLMsgsStateInfo {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_msgs_state_info

	return m
}

func (m *TLMsgsStateInfo) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "msgs_state_info", TLObject: m}
	return wrapper.String()
}

// MsgsStateInfoClazzName <--
func (m *TLMsgsStateInfo) MsgsStateInfoClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLMsgsStateInfo) ClazzName() string {
	return ClazzName_msgs_state_info
}

// ToMsgsStateInfo <--
func (m *TLMsgsStateInfo) ToMsgsStateInfo() *MsgsStateInfo {
	if m == nil {
		return nil
	}

	return m

}

// Encode <--
func (m *TLMsgsStateInfo) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_msgs_state_info, int(layer)); clazzId {
	case 0x4deb57d:
		x.PutClazzID(0x4deb57d)

		x.PutInt64(m.ReqMsgId)
		x.PutString(m.Info)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msgs_state_info, layer)
	}
}

// Decode <--
func (m *TLMsgsStateInfo) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x4deb57d:
		m.ReqMsgId, err = d.Int64()
		m.Info, err = d.String()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// MsgsStateInfo <--
type MsgsStateInfo = TLMsgsStateInfo

// MsgsStateReqClazz <--
//   - TL_MsgsStateReq
type MsgsStateReqClazz = *TLMsgsStateReq

func DecodeMsgsStateReqClazz(d *bin.Decoder) (MsgsStateReqClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0xda69fb52:
		x := &TLMsgsStateReq{ClazzID: id, ClazzName2: ClazzName_msgs_state_req}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeMsgsStateReq - unexpected clazzId: %d", id)
	}

}

// TLMsgsStateReq <--
type TLMsgsStateReq struct {
	ClazzID    uint32  `json:"_id"`
	ClazzName2 string  `json:"_name"`
	MsgIds     []int64 `json:"msg_ids"`
}

func MakeTLMsgsStateReq(m *TLMsgsStateReq) *TLMsgsStateReq {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_msgs_state_req

	return m
}

func (m *TLMsgsStateReq) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "msgs_state_req", TLObject: m}
	return wrapper.String()
}

// MsgsStateReqClazzName <--
func (m *TLMsgsStateReq) MsgsStateReqClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLMsgsStateReq) ClazzName() string {
	return ClazzName_msgs_state_req
}

// ToMsgsStateReq <--
func (m *TLMsgsStateReq) ToMsgsStateReq() *MsgsStateReq {
	if m == nil {
		return nil
	}

	return m

}

// Encode <--
func (m *TLMsgsStateReq) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_msgs_state_req, int(layer)); clazzId {
	case 0xda69fb52:
		x.PutClazzID(0xda69fb52)

		iface.EncodeInt64List(x, m.MsgIds)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msgs_state_req, layer)
	}
}

// Decode <--
func (m *TLMsgsStateReq) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xda69fb52:

		m.MsgIds, err = iface.DecodeInt64List(d)

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// MsgsStateReq <--
type MsgsStateReq = TLMsgsStateReq

// NewSessionClazz <--
//   - TL_NewSessionCreated
type NewSessionClazz = *TLNewSessionCreated

func DecodeNewSessionClazz(d *bin.Decoder) (NewSessionClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x9ec20908:
		x := &TLNewSessionCreated{ClazzID: id, ClazzName2: ClazzName_new_session_created}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeNewSession - unexpected clazzId: %d", id)
	}

}

// TLNewSessionCreated <--
type TLNewSessionCreated struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	FirstMsgId int64  `json:"first_msg_id"`
	UniqueId   int64  `json:"unique_id"`
	ServerSalt int64  `json:"server_salt"`
}

func MakeTLNewSessionCreated(m *TLNewSessionCreated) *TLNewSessionCreated {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_new_session_created

	return m
}

func (m *TLNewSessionCreated) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "new_session_created", TLObject: m}
	return wrapper.String()
}

// NewSessionClazzName <--
func (m *TLNewSessionCreated) NewSessionClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLNewSessionCreated) ClazzName() string {
	return ClazzName_new_session_created
}

// ToNewSession <--
func (m *TLNewSessionCreated) ToNewSession() *NewSession {
	if m == nil {
		return nil
	}

	return m

}

// Encode <--
func (m *TLNewSessionCreated) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_new_session_created, int(layer)); clazzId {
	case 0x9ec20908:
		x.PutClazzID(0x9ec20908)

		x.PutInt64(m.FirstMsgId)
		x.PutInt64(m.UniqueId)
		x.PutInt64(m.ServerSalt)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_new_session_created, layer)
	}
}

// Decode <--
func (m *TLNewSessionCreated) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x9ec20908:
		m.FirstMsgId, err = d.Int64()
		m.UniqueId, err = d.Int64()
		m.ServerSalt, err = d.Int64()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// NewSession <--
type NewSession = TLNewSessionCreated

// PongClazz <--
//   - TL_Pong
type PongClazz = *TLPong

func DecodePongClazz(d *bin.Decoder) (PongClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x347773c5:
		x := &TLPong{ClazzID: id, ClazzName2: ClazzName_pong}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodePong - unexpected clazzId: %d", id)
	}

}

// TLPong <--
type TLPong struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	MsgId      int64  `json:"msg_id"`
	PingId     int64  `json:"ping_id"`
}

func MakeTLPong(m *TLPong) *TLPong {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_pong

	return m
}

func (m *TLPong) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "pong", TLObject: m}
	return wrapper.String()
}

// PongClazzName <--
func (m *TLPong) PongClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLPong) ClazzName() string {
	return ClazzName_pong
}

// ToPong <--
func (m *TLPong) ToPong() *Pong {
	if m == nil {
		return nil
	}

	return m

}

// Encode <--
func (m *TLPong) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_pong, int(layer)); clazzId {
	case 0x347773c5:
		x.PutClazzID(0x347773c5)

		x.PutInt64(m.MsgId)
		x.PutInt64(m.PingId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_pong, layer)
	}
}

// Decode <--
func (m *TLPong) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x347773c5:
		m.MsgId, err = d.Int64()
		m.PingId, err = d.Int64()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// Pong <--
type Pong = TLPong

// RpcDropAnswerClazz <--
//   - TL_RpcAnswerUnknown
//   - TL_RpcAnswerDroppedRunning
//   - TL_RpcAnswerDropped
type RpcDropAnswerClazz interface {
	iface.TLObject
	RpcDropAnswerClazzName() string
}

func DecodeRpcDropAnswerClazz(d *bin.Decoder) (RpcDropAnswerClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x5e2ad36e:
		x := &TLRpcAnswerUnknown{ClazzID: id, ClazzName2: ClazzName_rpc_answer_unknown}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0xcd78e586:
		x := &TLRpcAnswerDroppedRunning{ClazzID: id, ClazzName2: ClazzName_rpc_answer_dropped_running}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0xa43ad8b7:
		x := &TLRpcAnswerDropped{ClazzID: id, ClazzName2: ClazzName_rpc_answer_dropped}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeRpcDropAnswer - unexpected clazzId: %d", id)
	}

}

// TLRpcAnswerUnknown <--
type TLRpcAnswerUnknown struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
}

func MakeTLRpcAnswerUnknown(m *TLRpcAnswerUnknown) *TLRpcAnswerUnknown {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_rpc_answer_unknown

	return m
}

func (m *TLRpcAnswerUnknown) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "rpc_answer_unknown", TLObject: m}
	return wrapper.String()
}

// RpcDropAnswerClazzName <--
func (m *TLRpcAnswerUnknown) RpcDropAnswerClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLRpcAnswerUnknown) ClazzName() string {
	return ClazzName_rpc_answer_unknown
}

// ToRpcDropAnswer <--
func (m *TLRpcAnswerUnknown) ToRpcDropAnswer() *RpcDropAnswer {
	if m == nil {
		return nil
	}

	return &RpcDropAnswer{Clazz: m}

}

// Encode <--
func (m *TLRpcAnswerUnknown) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_rpc_answer_unknown, int(layer)); clazzId {
	case 0x5e2ad36e:
		x.PutClazzID(0x5e2ad36e)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_rpc_answer_unknown, layer)
	}
}

// Decode <--
func (m *TLRpcAnswerUnknown) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x5e2ad36e:

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLRpcAnswerDroppedRunning <--
type TLRpcAnswerDroppedRunning struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
}

func MakeTLRpcAnswerDroppedRunning(m *TLRpcAnswerDroppedRunning) *TLRpcAnswerDroppedRunning {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_rpc_answer_dropped_running

	return m
}

func (m *TLRpcAnswerDroppedRunning) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "rpc_answer_dropped_running", TLObject: m}
	return wrapper.String()
}

// RpcDropAnswerClazzName <--
func (m *TLRpcAnswerDroppedRunning) RpcDropAnswerClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLRpcAnswerDroppedRunning) ClazzName() string {
	return ClazzName_rpc_answer_dropped_running
}

// ToRpcDropAnswer <--
func (m *TLRpcAnswerDroppedRunning) ToRpcDropAnswer() *RpcDropAnswer {
	if m == nil {
		return nil
	}

	return &RpcDropAnswer{Clazz: m}

}

// Encode <--
func (m *TLRpcAnswerDroppedRunning) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_rpc_answer_dropped_running, int(layer)); clazzId {
	case 0xcd78e586:
		x.PutClazzID(0xcd78e586)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_rpc_answer_dropped_running, layer)
	}
}

// Decode <--
func (m *TLRpcAnswerDroppedRunning) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xcd78e586:

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLRpcAnswerDropped <--
type TLRpcAnswerDropped struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	MsgId      int64  `json:"msg_id"`
	SeqNo      int32  `json:"seq_no"`
	Bytes      int32  `json:"bytes"`
}

func MakeTLRpcAnswerDropped(m *TLRpcAnswerDropped) *TLRpcAnswerDropped {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_rpc_answer_dropped

	return m
}

func (m *TLRpcAnswerDropped) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "rpc_answer_dropped", TLObject: m}
	return wrapper.String()
}

// RpcDropAnswerClazzName <--
func (m *TLRpcAnswerDropped) RpcDropAnswerClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLRpcAnswerDropped) ClazzName() string {
	return ClazzName_rpc_answer_dropped
}

// ToRpcDropAnswer <--
func (m *TLRpcAnswerDropped) ToRpcDropAnswer() *RpcDropAnswer {
	if m == nil {
		return nil
	}

	return &RpcDropAnswer{Clazz: m}

}

// Encode <--
func (m *TLRpcAnswerDropped) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_rpc_answer_dropped, int(layer)); clazzId {
	case 0xa43ad8b7:
		x.PutClazzID(0xa43ad8b7)

		x.PutInt64(m.MsgId)
		x.PutInt32(m.SeqNo)
		x.PutInt32(m.Bytes)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_rpc_answer_dropped, layer)
	}
}

// Decode <--
func (m *TLRpcAnswerDropped) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xa43ad8b7:
		m.MsgId, err = d.Int64()
		m.SeqNo, err = d.Int32()
		m.Bytes, err = d.Int32()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// RpcDropAnswer <--
type RpcDropAnswer struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz RpcDropAnswerClazz `json:"_clazz"`
}

func (m *RpcDropAnswer) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.RpcDropAnswerClazzName()
	}
}

func (m *RpcDropAnswer) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: m.ClazzName(), TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *RpcDropAnswer) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("RpcDropAnswer - invalid Clazz")
}

// Decode <--
func (m *RpcDropAnswer) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeRpcDropAnswerClazz(d)
	return
}

// Match <--
func (m *RpcDropAnswer) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}

	switch c := m.Clazz.(type) {
	case *TLRpcAnswerUnknown:
		for _, v := range f {
			if f1, ok := v.(func(c *TLRpcAnswerUnknown) interface{}); ok {
				f1(c)
			}
		}
	case *TLRpcAnswerDroppedRunning:
		for _, v := range f {
			if f1, ok := v.(func(c *TLRpcAnswerDroppedRunning) interface{}); ok {
				f1(c)
			}
		}
	case *TLRpcAnswerDropped:
		for _, v := range f {
			if f1, ok := v.(func(c *TLRpcAnswerDropped) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToRpcAnswerUnknown <--
func (m *RpcDropAnswer) ToRpcAnswerUnknown() (*TLRpcAnswerUnknown, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLRpcAnswerUnknown); ok {
		return x, true
	}

	return nil, false
}

// ToRpcAnswerDroppedRunning <--
func (m *RpcDropAnswer) ToRpcAnswerDroppedRunning() (*TLRpcAnswerDroppedRunning, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLRpcAnswerDroppedRunning); ok {
		return x, true
	}

	return nil, false
}

// ToRpcAnswerDropped <--
func (m *RpcDropAnswer) ToRpcAnswerDropped() (*TLRpcAnswerDropped, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLRpcAnswerDropped); ok {
		return x, true
	}

	return nil, false
}

// RpcErrorClazz <--
//   - TL_RpcError
type RpcErrorClazz = *TLRpcError

func DecodeRpcErrorClazz(d *bin.Decoder) (RpcErrorClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x2144ca19:
		x := &TLRpcError{ClazzID: id, ClazzName2: ClazzName_rpc_error}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeRpcError - unexpected clazzId: %d", id)
	}

}

// TLRpcError <--
type TLRpcError struct {
	ClazzID      uint32 `json:"_id"`
	ClazzName2   string `json:"_name"`
	ErrorCode    int32  `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func MakeTLRpcError(m *TLRpcError) *TLRpcError {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_rpc_error

	return m
}

func (m *TLRpcError) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "rpc_error", TLObject: m}
	return wrapper.String()
}

// RpcErrorClazzName <--
func (m *TLRpcError) RpcErrorClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLRpcError) ClazzName() string {
	return ClazzName_rpc_error
}

// ToRpcError <--
func (m *TLRpcError) ToRpcError() *RpcError {
	if m == nil {
		return nil
	}

	return m

}

// Encode <--
func (m *TLRpcError) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_rpc_error, int(layer)); clazzId {
	case 0x2144ca19:
		x.PutClazzID(0x2144ca19)

		x.PutInt32(m.ErrorCode)
		x.PutString(m.ErrorMessage)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_rpc_error, layer)
	}
}

// Decode <--
func (m *TLRpcError) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x2144ca19:
		m.ErrorCode, err = d.Int32()
		m.ErrorMessage, err = d.String()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// RpcError <--
type RpcError = TLRpcError

// TlsBlockClazz <--
//   - TL_TlsBlockString
//   - TL_TlsBlockRandom
//   - TL_TlsBlockZero
//   - TL_TlsBlockDomain
//   - TL_TlsBlockGrease
//   - TL_TlsBlockPublicKey
//   - TL_TlsBlockScope
type TlsBlockClazz interface {
	iface.TLObject
	TlsBlockClazzName() string
}

func DecodeTlsBlockClazz(d *bin.Decoder) (TlsBlockClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x4218a164:
		x := &TLTlsBlockString{ClazzID: id, ClazzName2: ClazzName_tlsBlockString}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x4d4dc41e:
		x := &TLTlsBlockRandom{ClazzID: id, ClazzName2: ClazzName_tlsBlockRandom}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x9333afb:
		x := &TLTlsBlockZero{ClazzID: id, ClazzName2: ClazzName_tlsBlockZero}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x10e8636f:
		x := &TLTlsBlockDomain{ClazzID: id, ClazzName2: ClazzName_tlsBlockDomain}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0xe675a1c1:
		x := &TLTlsBlockGrease{ClazzID: id, ClazzName2: ClazzName_tlsBlockGrease}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x9eb95b5c:
		x := &TLTlsBlockPublicKey{ClazzID: id, ClazzName2: ClazzName_tlsBlockPublicKey}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0xe725d44f:
		x := &TLTlsBlockScope{ClazzID: id, ClazzName2: ClazzName_tlsBlockScope}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeTlsBlock - unexpected clazzId: %d", id)
	}

}

// TLTlsBlockString <--
type TLTlsBlockString struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Data       string `json:"data"`
}

func MakeTLTlsBlockString(m *TLTlsBlockString) *TLTlsBlockString {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_tlsBlockString

	return m
}

func (m *TLTlsBlockString) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "tlsBlockString", TLObject: m}
	return wrapper.String()
}

// TlsBlockClazzName <--
func (m *TLTlsBlockString) TlsBlockClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLTlsBlockString) ClazzName() string {
	return ClazzName_tlsBlockString
}

// ToTlsBlock <--
func (m *TLTlsBlockString) ToTlsBlock() *TlsBlock {
	if m == nil {
		return nil
	}

	return &TlsBlock{Clazz: m}

}

// Encode <--
func (m *TLTlsBlockString) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_tlsBlockString, int(layer)); clazzId {
	case 0x4218a164:
		x.PutClazzID(0x4218a164)

		x.PutString(m.Data)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_tlsBlockString, layer)
	}
}

// Decode <--
func (m *TLTlsBlockString) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x4218a164:
		m.Data, err = d.String()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLTlsBlockRandom <--
type TLTlsBlockRandom struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Length     int32  `json:"length"`
}

func MakeTLTlsBlockRandom(m *TLTlsBlockRandom) *TLTlsBlockRandom {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_tlsBlockRandom

	return m
}

func (m *TLTlsBlockRandom) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "tlsBlockRandom", TLObject: m}
	return wrapper.String()
}

// TlsBlockClazzName <--
func (m *TLTlsBlockRandom) TlsBlockClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLTlsBlockRandom) ClazzName() string {
	return ClazzName_tlsBlockRandom
}

// ToTlsBlock <--
func (m *TLTlsBlockRandom) ToTlsBlock() *TlsBlock {
	if m == nil {
		return nil
	}

	return &TlsBlock{Clazz: m}

}

// Encode <--
func (m *TLTlsBlockRandom) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_tlsBlockRandom, int(layer)); clazzId {
	case 0x4d4dc41e:
		x.PutClazzID(0x4d4dc41e)

		x.PutInt32(m.Length)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_tlsBlockRandom, layer)
	}
}

// Decode <--
func (m *TLTlsBlockRandom) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x4d4dc41e:
		m.Length, err = d.Int32()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLTlsBlockZero <--
type TLTlsBlockZero struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Length     int32  `json:"length"`
}

func MakeTLTlsBlockZero(m *TLTlsBlockZero) *TLTlsBlockZero {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_tlsBlockZero

	return m
}

func (m *TLTlsBlockZero) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "tlsBlockZero", TLObject: m}
	return wrapper.String()
}

// TlsBlockClazzName <--
func (m *TLTlsBlockZero) TlsBlockClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLTlsBlockZero) ClazzName() string {
	return ClazzName_tlsBlockZero
}

// ToTlsBlock <--
func (m *TLTlsBlockZero) ToTlsBlock() *TlsBlock {
	if m == nil {
		return nil
	}

	return &TlsBlock{Clazz: m}

}

// Encode <--
func (m *TLTlsBlockZero) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_tlsBlockZero, int(layer)); clazzId {
	case 0x9333afb:
		x.PutClazzID(0x9333afb)

		x.PutInt32(m.Length)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_tlsBlockZero, layer)
	}
}

// Decode <--
func (m *TLTlsBlockZero) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x9333afb:
		m.Length, err = d.Int32()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLTlsBlockDomain <--
type TLTlsBlockDomain struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
}

func MakeTLTlsBlockDomain(m *TLTlsBlockDomain) *TLTlsBlockDomain {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_tlsBlockDomain

	return m
}

func (m *TLTlsBlockDomain) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "tlsBlockDomain", TLObject: m}
	return wrapper.String()
}

// TlsBlockClazzName <--
func (m *TLTlsBlockDomain) TlsBlockClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLTlsBlockDomain) ClazzName() string {
	return ClazzName_tlsBlockDomain
}

// ToTlsBlock <--
func (m *TLTlsBlockDomain) ToTlsBlock() *TlsBlock {
	if m == nil {
		return nil
	}

	return &TlsBlock{Clazz: m}

}

// Encode <--
func (m *TLTlsBlockDomain) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_tlsBlockDomain, int(layer)); clazzId {
	case 0x10e8636f:
		x.PutClazzID(0x10e8636f)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_tlsBlockDomain, layer)
	}
}

// Decode <--
func (m *TLTlsBlockDomain) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x10e8636f:

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLTlsBlockGrease <--
type TLTlsBlockGrease struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Seed       int32  `json:"seed"`
}

func MakeTLTlsBlockGrease(m *TLTlsBlockGrease) *TLTlsBlockGrease {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_tlsBlockGrease

	return m
}

func (m *TLTlsBlockGrease) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "tlsBlockGrease", TLObject: m}
	return wrapper.String()
}

// TlsBlockClazzName <--
func (m *TLTlsBlockGrease) TlsBlockClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLTlsBlockGrease) ClazzName() string {
	return ClazzName_tlsBlockGrease
}

// ToTlsBlock <--
func (m *TLTlsBlockGrease) ToTlsBlock() *TlsBlock {
	if m == nil {
		return nil
	}

	return &TlsBlock{Clazz: m}

}

// Encode <--
func (m *TLTlsBlockGrease) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_tlsBlockGrease, int(layer)); clazzId {
	case 0xe675a1c1:
		x.PutClazzID(0xe675a1c1)

		x.PutInt32(m.Seed)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_tlsBlockGrease, layer)
	}
}

// Decode <--
func (m *TLTlsBlockGrease) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xe675a1c1:
		m.Seed, err = d.Int32()

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLTlsBlockPublicKey <--
type TLTlsBlockPublicKey struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
}

func MakeTLTlsBlockPublicKey(m *TLTlsBlockPublicKey) *TLTlsBlockPublicKey {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_tlsBlockPublicKey

	return m
}

func (m *TLTlsBlockPublicKey) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "tlsBlockPublicKey", TLObject: m}
	return wrapper.String()
}

// TlsBlockClazzName <--
func (m *TLTlsBlockPublicKey) TlsBlockClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLTlsBlockPublicKey) ClazzName() string {
	return ClazzName_tlsBlockPublicKey
}

// ToTlsBlock <--
func (m *TLTlsBlockPublicKey) ToTlsBlock() *TlsBlock {
	if m == nil {
		return nil
	}

	return &TlsBlock{Clazz: m}

}

// Encode <--
func (m *TLTlsBlockPublicKey) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_tlsBlockPublicKey, int(layer)); clazzId {
	case 0x9eb95b5c:
		x.PutClazzID(0x9eb95b5c)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_tlsBlockPublicKey, layer)
	}
}

// Decode <--
func (m *TLTlsBlockPublicKey) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x9eb95b5c:

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLTlsBlockScope <--
type TLTlsBlockScope struct {
	ClazzID    uint32          `json:"_id"`
	ClazzName2 string          `json:"_name"`
	Entries    []TlsBlockClazz `json:"entries"`
}

func MakeTLTlsBlockScope(m *TLTlsBlockScope) *TLTlsBlockScope {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_tlsBlockScope

	return m
}

func (m *TLTlsBlockScope) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "tlsBlockScope", TLObject: m}
	return wrapper.String()
}

// TlsBlockClazzName <--
func (m *TLTlsBlockScope) TlsBlockClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLTlsBlockScope) ClazzName() string {
	return ClazzName_tlsBlockScope
}

// ToTlsBlock <--
func (m *TLTlsBlockScope) ToTlsBlock() *TlsBlock {
	if m == nil {
		return nil
	}

	return &TlsBlock{Clazz: m}

}

// Encode <--
func (m *TLTlsBlockScope) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_tlsBlockScope, int(layer)); clazzId {
	case 0xe725d44f:
		x.PutClazzID(0xe725d44f)

		_ = iface.EncodeObjectList(x, m.Entries, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_tlsBlockScope, layer)
	}
}

// Decode <--
func (m *TLTlsBlockScope) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xe725d44f:
		c3, err2 := d.ClazzID()
		if c3 != iface.ClazzID_vector {
			// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 3, c3)
			return err2
		}
		l3, err3 := d.Int()
		if err3 != nil {
			return err3
		}
		v3 := make([]TlsBlockClazz, l3)
		for i := 0; i < l3; i++ {
			v3[i], err3 = DecodeTlsBlockClazz(d)
			if err3 != nil {
				return err3
			}
		}
		m.Entries = v3

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TlsBlock <--
type TlsBlock struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz TlsBlockClazz `json:"_clazz"`
}

func (m *TlsBlock) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.TlsBlockClazzName()
	}
}

func (m *TlsBlock) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: m.ClazzName(), TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TlsBlock) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("TlsBlock - invalid Clazz")
}

// Decode <--
func (m *TlsBlock) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeTlsBlockClazz(d)
	return
}

// Match <--
func (m *TlsBlock) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}

	switch c := m.Clazz.(type) {
	case *TLTlsBlockString:
		for _, v := range f {
			if f1, ok := v.(func(c *TLTlsBlockString) interface{}); ok {
				f1(c)
			}
		}
	case *TLTlsBlockRandom:
		for _, v := range f {
			if f1, ok := v.(func(c *TLTlsBlockRandom) interface{}); ok {
				f1(c)
			}
		}
	case *TLTlsBlockZero:
		for _, v := range f {
			if f1, ok := v.(func(c *TLTlsBlockZero) interface{}); ok {
				f1(c)
			}
		}
	case *TLTlsBlockDomain:
		for _, v := range f {
			if f1, ok := v.(func(c *TLTlsBlockDomain) interface{}); ok {
				f1(c)
			}
		}
	case *TLTlsBlockGrease:
		for _, v := range f {
			if f1, ok := v.(func(c *TLTlsBlockGrease) interface{}); ok {
				f1(c)
			}
		}
	case *TLTlsBlockPublicKey:
		for _, v := range f {
			if f1, ok := v.(func(c *TLTlsBlockPublicKey) interface{}); ok {
				f1(c)
			}
		}
	case *TLTlsBlockScope:
		for _, v := range f {
			if f1, ok := v.(func(c *TLTlsBlockScope) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToTlsBlockString <--
func (m *TlsBlock) ToTlsBlockString() (*TLTlsBlockString, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLTlsBlockString); ok {
		return x, true
	}

	return nil, false
}

// ToTlsBlockRandom <--
func (m *TlsBlock) ToTlsBlockRandom() (*TLTlsBlockRandom, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLTlsBlockRandom); ok {
		return x, true
	}

	return nil, false
}

// ToTlsBlockZero <--
func (m *TlsBlock) ToTlsBlockZero() (*TLTlsBlockZero, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLTlsBlockZero); ok {
		return x, true
	}

	return nil, false
}

// ToTlsBlockDomain <--
func (m *TlsBlock) ToTlsBlockDomain() (*TLTlsBlockDomain, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLTlsBlockDomain); ok {
		return x, true
	}

	return nil, false
}

// ToTlsBlockGrease <--
func (m *TlsBlock) ToTlsBlockGrease() (*TLTlsBlockGrease, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLTlsBlockGrease); ok {
		return x, true
	}

	return nil, false
}

// ToTlsBlockPublicKey <--
func (m *TlsBlock) ToTlsBlockPublicKey() (*TLTlsBlockPublicKey, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLTlsBlockPublicKey); ok {
		return x, true
	}

	return nil, false
}

// ToTlsBlockScope <--
func (m *TlsBlock) ToTlsBlockScope() (*TLTlsBlockScope, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLTlsBlockScope); ok {
		return x, true
	}

	return nil, false
}

// TlsClientHelloClazz <--
//   - TL_TlsClientHello
type TlsClientHelloClazz = *TLTlsClientHello

func DecodeTlsClientHelloClazz(d *bin.Decoder) (TlsClientHelloClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x6c52c484:
		x := &TLTlsClientHello{ClazzID: id, ClazzName2: ClazzName_tlsClientHello}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeTlsClientHello - unexpected clazzId: %d", id)
	}

}

// TLTlsClientHello <--
type TLTlsClientHello struct {
	ClazzID    uint32          `json:"_id"`
	ClazzName2 string          `json:"_name"`
	Blocks     []TlsBlockClazz `json:"blocks"`
}

func MakeTLTlsClientHello(m *TLTlsClientHello) *TLTlsClientHello {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_tlsClientHello

	return m
}

func (m *TLTlsClientHello) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: "tlsClientHello", TLObject: m}
	return wrapper.String()
}

// TlsClientHelloClazzName <--
func (m *TLTlsClientHello) TlsClientHelloClazzName() string {
	return m.ClazzName2
}

// ClazzName <--
func (m *TLTlsClientHello) ClazzName() string {
	return ClazzName_tlsClientHello
}

// ToTlsClientHello <--
func (m *TLTlsClientHello) ToTlsClientHello() *TlsClientHello {
	if m == nil {
		return nil
	}

	return m

}

// Encode <--
func (m *TLTlsClientHello) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_tlsClientHello, int(layer)); clazzId {
	case 0x6c52c484:
		x.PutClazzID(0x6c52c484)

		// x.PutClazzID(iface.ClazzID_vector)
		x.PutInt(len(m.Blocks))
		for _, v := range m.Blocks {
			_ = v.Encode(x, layer)
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_tlsClientHello, layer)
	}
}

// Decode <--
func (m *TLTlsClientHello) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x6c52c484:
		// c0, err2 := d.ClazzID()
		// if c0 != int32(iface.ClazzID_vector) {
		//     err2 = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 0, c0)
		//     return err2
		// }
		l0, err2 := d.Int()
		if err2 != nil {
			return err2
		}
		v0 := make([]TlsBlockClazz, l0)
		for i := 0; i < l0; i++ {
			v0[i], err2 = DecodeTlsBlockClazz(d)
			if err2 != nil {
				return err2
			}
		}
		m.Blocks = v0

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TlsClientHello <--
type TlsClientHello = TLTlsClientHello
