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

// PhoneCodeTransactionClazz <--
//   - TL_PhoneCodeTransaction
type PhoneCodeTransactionClazz interface {
	iface.TLObject
	PhoneCodeTransactionClazzName() string
}

func DecodePhoneCodeTransactionClazz(d *bin.Decoder) (PhoneCodeTransactionClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_phoneCodeTransaction:
		x := &TLPhoneCodeTransaction{ClazzID: id, ClazzName2: ClazzName_phoneCodeTransaction}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodePhoneCodeTransaction - unexpected clazzId: %d", id)
	}
}

// TLPhoneCodeTransaction <--
type TLPhoneCodeTransaction struct {
	ClazzID               uint32 `json:"_id"`
	ClazzName2            string `json:"_name"`
	AuthKeyId             int64  `json:"auth_key_id"`
	SessionId             int64  `json:"session_id"`
	Phone                 string `json:"phone"`
	PhoneNumberRegistered bool   `json:"phone_number_registered"`
	PhoneCode             string `json:"phone_code"`
	PhoneCodeHash         string `json:"phone_code_hash"`
	PhoneCodeExpired      int32  `json:"phone_code_expired"`
	PhoneCodeExtraData    string `json:"phone_code_extra_data"`
	SentCodeType          int32  `json:"sent_code_type"`
	FlashCallPattern      string `json:"flash_call_pattern"`
	NextCodeType          int32  `json:"next_code_type"`
	State                 int32  `json:"state"`
}

func MakeTLPhoneCodeTransaction(m *TLPhoneCodeTransaction) *TLPhoneCodeTransaction {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_phoneCodeTransaction

	return m
}

func (m *TLPhoneCodeTransaction) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// PhoneCodeTransactionClazzName <--
func (m *TLPhoneCodeTransaction) PhoneCodeTransactionClazzName() string {
	return ClazzName_phoneCodeTransaction
}

// ClazzName <--
func (m *TLPhoneCodeTransaction) ClazzName() string {
	return m.ClazzName2
}

// ToPhoneCodeTransaction <--
func (m *TLPhoneCodeTransaction) ToPhoneCodeTransaction() *PhoneCodeTransaction {
	if m == nil {
		return nil
	}

	return &PhoneCodeTransaction{Clazz: m}
}

// Encode <--
func (m *TLPhoneCodeTransaction) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x83739698: func() error {
			x.PutClazzID(0x83739698)

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
			x.PutString(m.PhoneCode)
			x.PutString(m.PhoneCodeHash)
			x.PutInt32(m.PhoneCodeExpired)
			x.PutString(m.PhoneCodeExtraData)
			x.PutInt32(m.SentCodeType)
			x.PutString(m.FlashCallPattern)
			x.PutInt32(m.NextCodeType)
			x.PutInt32(m.State)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_phoneCodeTransaction, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_phoneCodeTransaction, layer)
	}
}

// Decode <--
func (m *TLPhoneCodeTransaction) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x83739698: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.AuthKeyId, err = d.Int64()
			m.SessionId, err = d.Int64()
			m.Phone, err = d.String()
			if (flags & (1 << 0)) != 0 {
				m.PhoneNumberRegistered = true
			}
			m.PhoneCode, err = d.String()
			m.PhoneCodeHash, err = d.String()
			m.PhoneCodeExpired, err = d.Int32()
			m.PhoneCodeExtraData, err = d.String()
			m.SentCodeType, err = d.Int32()
			m.FlashCallPattern, err = d.String()
			m.NextCodeType, err = d.Int32()
			m.State, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// PhoneCodeTransaction <--
type PhoneCodeTransaction struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz PhoneCodeTransactionClazz `json:"_clazz"`
}

func (m *PhoneCodeTransaction) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *PhoneCodeTransaction) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.PhoneCodeTransactionClazzName()
	}
}

// Encode <--
func (m *PhoneCodeTransaction) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("PhoneCodeTransaction - invalid Clazz")
}

// Decode <--
func (m *PhoneCodeTransaction) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodePhoneCodeTransactionClazz(d)
	return
}

// Match <--
func (m *PhoneCodeTransaction) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLPhoneCodeTransaction:
		for _, v := range f {
			if f1, ok := v.(func(c *TLPhoneCodeTransaction) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToPhoneCodeTransaction <--
func (m *PhoneCodeTransaction) ToPhoneCodeTransaction() (*TLPhoneCodeTransaction, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLPhoneCodeTransaction); ok {
		return x, true
	}

	return nil, false
}
