/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package code

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
)

var (
	_ iface.TLObject
	_ fmt.Stringer
	_ *tg.Bool
	_ bin.Fields
	_ json.Marshaler
)

// TLCodeCreatePhoneCode <--
type TLCodeCreatePhoneCode struct {
	ClazzID               uint32 `json:"_id"`
	AuthKeyId             int64  `json:"auth_key_id"`
	SessionId             int64  `json:"session_id"`
	Phone                 string `json:"phone"`
	PhoneNumberRegistered bool   `json:"phone_number_registered"`
	SentCodeType          int32  `json:"sent_code_type"`
	NextCodeType          int32  `json:"next_code_type"`
	State                 int32  `json:"state"`
}

func (m *TLCodeCreatePhoneCode) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLCodeCreatePhoneCode) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x6023e09e: func() error {
			x.PutClazzID(0x6023e09e)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.PhoneNumberRegistered == true {
					flags |= 1 << 0
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.AuthKeyId)
			x.PutInt64(m.SessionId)
			x.PutString(m.Phone)
			x.PutInt32(m.SentCodeType)
			x.PutInt32(m.NextCodeType)
			x.PutInt32(m.State)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_code_createPhoneCode, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_code_createPhoneCode, layer)
	}
}

// Decode <--
func (m *TLCodeCreatePhoneCode) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x6023e09e: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.AuthKeyId, err = d.Int64()
			m.SessionId, err = d.Int64()
			m.Phone, err = d.String()
			if (flags & (1 << 0)) != 0 {
				m.PhoneNumberRegistered = true
			}
			m.SentCodeType, err = d.Int32()
			m.NextCodeType, err = d.Int32()
			m.State, err = d.Int32()

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

// TLCodeGetPhoneCode <--
type TLCodeGetPhoneCode struct {
	ClazzID       uint32 `json:"_id"`
	AuthKeyId     int64  `json:"auth_key_id"`
	Phone         string `json:"phone"`
	PhoneCodeHash string `json:"phone_code_hash"`
}

func (m *TLCodeGetPhoneCode) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLCodeGetPhoneCode) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x61a4a0f9: func() error {
			x.PutClazzID(0x61a4a0f9)

			x.PutInt64(m.AuthKeyId)
			x.PutString(m.Phone)
			x.PutString(m.PhoneCodeHash)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_code_getPhoneCode, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_code_getPhoneCode, layer)
	}
}

// Decode <--
func (m *TLCodeGetPhoneCode) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x61a4a0f9: func() (err error) {
			m.AuthKeyId, err = d.Int64()
			m.Phone, err = d.String()
			m.PhoneCodeHash, err = d.String()

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

// TLCodeDeletePhoneCode <--
type TLCodeDeletePhoneCode struct {
	ClazzID       uint32 `json:"_id"`
	AuthKeyId     int64  `json:"auth_key_id"`
	Phone         string `json:"phone"`
	PhoneCodeHash string `json:"phone_code_hash"`
}

func (m *TLCodeDeletePhoneCode) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLCodeDeletePhoneCode) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa6b06a50: func() error {
			x.PutClazzID(0xa6b06a50)

			x.PutInt64(m.AuthKeyId)
			x.PutString(m.Phone)
			x.PutString(m.PhoneCodeHash)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_code_deletePhoneCode, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_code_deletePhoneCode, layer)
	}
}

// Decode <--
func (m *TLCodeDeletePhoneCode) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa6b06a50: func() (err error) {
			m.AuthKeyId, err = d.Int64()
			m.Phone, err = d.String()
			m.PhoneCodeHash, err = d.String()

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

// TLCodeUpdatePhoneCodeData <--
type TLCodeUpdatePhoneCodeData struct {
	ClazzID       uint32                `json:"_id"`
	AuthKeyId     int64                 `json:"auth_key_id"`
	Phone         string                `json:"phone"`
	PhoneCodeHash string                `json:"phone_code_hash"`
	CodeData      *PhoneCodeTransaction `json:"code_data"`
}

func (m *TLCodeUpdatePhoneCodeData) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLCodeUpdatePhoneCodeData) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xb6950a95: func() error {
			x.PutClazzID(0xb6950a95)

			x.PutInt64(m.AuthKeyId)
			x.PutString(m.Phone)
			x.PutString(m.PhoneCodeHash)
			_ = m.CodeData.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_code_updatePhoneCodeData, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_code_updatePhoneCodeData, layer)
	}
}

// Decode <--
func (m *TLCodeUpdatePhoneCodeData) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xb6950a95: func() (err error) {
			m.AuthKeyId, err = d.Int64()
			m.Phone, err = d.String()
			m.PhoneCodeHash, err = d.String()

			m4 := &PhoneCodeTransaction{}
			_ = m4.Decode(d)
			m.CodeData = m4

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

// Vector api result type
// ----------------------------------------------------------------------------
// VectorResList <--

// ----------------------------------------------------------------------------
// rpc

type RPCCode interface {
	CodeCreatePhoneCode(ctx context.Context, in *TLCodeCreatePhoneCode) (*PhoneCodeTransaction, error)
	CodeGetPhoneCode(ctx context.Context, in *TLCodeGetPhoneCode) (*PhoneCodeTransaction, error)
	CodeDeletePhoneCode(ctx context.Context, in *TLCodeDeletePhoneCode) (*tg.Bool, error)
	CodeUpdatePhoneCodeData(ctx context.Context, in *TLCodeUpdatePhoneCodeData) (*tg.Bool, error)
}
