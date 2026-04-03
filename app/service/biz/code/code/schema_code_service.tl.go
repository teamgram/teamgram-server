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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_code_createPhoneCode, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLCodeCreatePhoneCode) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_code_createPhoneCode, int(layer)); clazzId {
	case 0x6023e09e:
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
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_code_createPhoneCode, layer)
	}
}

// Decode <--
func (m *TLCodeCreatePhoneCode) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x6023e09e:
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
		m.SentCodeType, err = d.Int32()
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

// TLCodeGetPhoneCode <--
type TLCodeGetPhoneCode struct {
	ClazzID       uint32 `json:"_id"`
	AuthKeyId     int64  `json:"auth_key_id"`
	Phone         string `json:"phone"`
	PhoneCodeHash string `json:"phone_code_hash"`
}

func (m *TLCodeGetPhoneCode) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_code_getPhoneCode, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLCodeGetPhoneCode) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_code_getPhoneCode, int(layer)); clazzId {
	case 0x61a4a0f9:
		x.PutClazzID(0x61a4a0f9)

		x.PutInt64(m.AuthKeyId)
		x.PutString(m.Phone)
		x.PutString(m.PhoneCodeHash)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_code_getPhoneCode, layer)
	}
}

// Decode <--
func (m *TLCodeGetPhoneCode) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x61a4a0f9:
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.Phone, err = d.String()
		if err != nil {
			return err
		}
		m.PhoneCodeHash, err = d.String()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_code_deletePhoneCode, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLCodeDeletePhoneCode) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_code_deletePhoneCode, int(layer)); clazzId {
	case 0xa6b06a50:
		x.PutClazzID(0xa6b06a50)

		x.PutInt64(m.AuthKeyId)
		x.PutString(m.Phone)
		x.PutString(m.PhoneCodeHash)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_code_deletePhoneCode, layer)
	}
}

// Decode <--
func (m *TLCodeDeletePhoneCode) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xa6b06a50:
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.Phone, err = d.String()
		if err != nil {
			return err
		}
		m.PhoneCodeHash, err = d.String()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLCodeUpdatePhoneCodeData <--
type TLCodeUpdatePhoneCodeData struct {
	ClazzID       uint32                    `json:"_id"`
	AuthKeyId     int64                     `json:"auth_key_id"`
	Phone         string                    `json:"phone"`
	PhoneCodeHash string                    `json:"phone_code_hash"`
	CodeData      PhoneCodeTransactionClazz `json:"code_data"`
}

func (m *TLCodeUpdatePhoneCodeData) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_code_updatePhoneCodeData, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLCodeUpdatePhoneCodeData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_code_updatePhoneCodeData, int(layer)); clazzId {
	case 0xb6950a95:
		x.PutClazzID(0xb6950a95)

		x.PutInt64(m.AuthKeyId)
		x.PutString(m.Phone)
		x.PutString(m.PhoneCodeHash)
		_ = m.CodeData.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_code_updatePhoneCodeData, layer)
	}
}

// Decode <--
func (m *TLCodeUpdatePhoneCodeData) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xb6950a95:
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.Phone, err = d.String()
		if err != nil {
			return err
		}
		m.PhoneCodeHash, err = d.String()
		if err != nil {
			return err
		}

		m.CodeData, err = DecodePhoneCodeTransactionClazz(d)
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

// ----------------------------------------------------------------------------
// rpc

type RPCCode interface {
	CodeCreatePhoneCode(ctx context.Context, in *TLCodeCreatePhoneCode) (*PhoneCodeTransaction, error)
	CodeGetPhoneCode(ctx context.Context, in *TLCodeGetPhoneCode) (*PhoneCodeTransaction, error)
	CodeDeletePhoneCode(ctx context.Context, in *TLCodeDeletePhoneCode) (*tg.Bool, error)
	CodeUpdatePhoneCodeData(ctx context.Context, in *TLCodeUpdatePhoneCodeData) (*tg.Bool, error)
}
