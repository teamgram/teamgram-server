/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package code

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

// PhoneCodeTransactionClazz <--
//   - TL_PhoneCodeTransaction
type PhoneCodeTransactionClazz = *TLPhoneCodeTransaction

func DecodePhoneCodeTransactionClazz(d *bin.Decoder) (PhoneCodeTransactionClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x83739698:
		x := &TLPhoneCodeTransaction{ClazzID: id, ClazzName2: ClazzName_phoneCodeTransaction}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
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

func (m *TLPhoneCodeTransaction) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("phoneCodeTransaction", m)
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

	return m

}

func (m *TLPhoneCodeTransaction) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_phoneCodeTransaction, int(layer)); clazzId {
	case 0x83739698:
		size := 4
		size += 4
		size += 8
		size += 8
		size += iface.CalcStringSize(m.Phone)
		size += iface.CalcStringSize(m.PhoneCode)
		size += iface.CalcStringSize(m.PhoneCodeHash)
		size += 4
		size += iface.CalcStringSize(m.PhoneCodeExtraData)
		size += 4
		size += iface.CalcStringSize(m.FlashCallPattern)
		size += 4
		size += 4

		return size
	default:
		return 0
	}
}

func (m *TLPhoneCodeTransaction) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_phoneCodeTransaction, int(layer)); clazzId {
	case 0x83739698:
		if err := iface.ValidateRequiredString("phone", m.Phone); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("phone_code", m.PhoneCode); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("phone_code_hash", m.PhoneCodeHash); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("phone_code_extra_data", m.PhoneCodeExtraData); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("flash_call_pattern", m.FlashCallPattern); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_phoneCodeTransaction, layer)
	}
}

// Encode <--
func (m *TLPhoneCodeTransaction) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_phoneCodeTransaction, int(layer)); clazzId {
	case 0x83739698:
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
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_phoneCodeTransaction, layer)
	}
}

// Decode <--
func (m *TLPhoneCodeTransaction) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x83739698:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.SessionId, err = d.Int64()
		if err != nil {
			return err
		}
		m.Phone, err = d.String()
		if err != nil {
			return err
		}
		if (flags & (1 << 0)) != 0 {
			m.PhoneNumberRegistered = true
		}
		m.PhoneCode, err = d.String()
		if err != nil {
			return err
		}
		m.PhoneCodeHash, err = d.String()
		if err != nil {
			return err
		}
		m.PhoneCodeExpired, err = d.Int32()
		if err != nil {
			return err
		}
		m.PhoneCodeExtraData, err = d.String()
		if err != nil {
			return err
		}
		m.SentCodeType, err = d.Int32()
		if err != nil {
			return err
		}
		m.FlashCallPattern, err = d.String()
		if err != nil {
			return err
		}
		m.NextCodeType, err = d.Int32()
		if err != nil {
			return err
		}
		m.State, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// PhoneCodeTransaction <--
type PhoneCodeTransaction = TLPhoneCodeTransaction
