/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

// ConstructorList
// RequestList

package code

import (
	"fmt"

	"github.com/teamgram/proto/mtproto"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

//////////////////////////////////////////////////////////////////////////////////////////

var _ *wrapperspb.Int32Value
var _ *mtproto.Bool
var _ fmt.Stringer

var clazzIdRegisters2 = map[int32]func() mtproto.TLObject{
	// Constructor
	-2089576808: func() mtproto.TLObject { // 0x83739698
		o := MakeTLPhoneCodeTransaction(nil)
		o.Data2.Constructor = -2089576808
		return o
	},

	// Method
	1612963998: func() mtproto.TLObject { // 0x6023e09e
		return &TLCodeCreatePhoneCode{
			Constructor: 1612963998,
		}
	},
	1638179065: func() mtproto.TLObject { // 0x61a4a0f9
		return &TLCodeGetPhoneCode{
			Constructor: 1638179065,
		}
	},
	-1498387888: func() mtproto.TLObject { // 0xa6b06a50
		return &TLCodeDeletePhoneCode{
			Constructor: -1498387888,
		}
	},
	-1231746411: func() mtproto.TLObject { // 0xb6950a95
		return &TLCodeUpdatePhoneCodeData{
			Constructor: -1231746411,
		}
	},
}

func NewTLObjectByClassID(classId int32) mtproto.TLObject {
	f, ok := clazzIdRegisters2[classId]
	if !ok {
		return nil
	}
	return f()
}

func CheckClassID(classId int32) (ok bool) {
	_, ok = clazzIdRegisters2[classId]
	return
}

//----------------------------------------------------------------------------------------------------------------

///////////////////////////////////////////////////////////////////////////////
// PhoneCodeTransaction <--
//  + TL_PhoneCodeTransaction
//

func (m *PhoneCodeTransaction) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_phoneCodeTransaction:
		t := m.To_PhoneCodeTransaction()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *PhoneCodeTransaction) CalcByteSize(layer int32) int {
	return 0
}

func (m *PhoneCodeTransaction) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x83739698:
		m2 := MakeTLPhoneCodeTransaction(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_PhoneCodeTransaction
func (m *PhoneCodeTransaction) To_PhoneCodeTransaction() *TLPhoneCodeTransaction {
	m.PredicateName = Predicate_phoneCodeTransaction
	return &TLPhoneCodeTransaction{
		Data2: m,
	}
}

// MakeTLPhoneCodeTransaction
func MakeTLPhoneCodeTransaction(data2 *PhoneCodeTransaction) *TLPhoneCodeTransaction {
	if data2 == nil {
		return &TLPhoneCodeTransaction{Data2: &PhoneCodeTransaction{
			PredicateName: Predicate_phoneCodeTransaction,
		}}
	} else {
		data2.PredicateName = Predicate_phoneCodeTransaction
		return &TLPhoneCodeTransaction{Data2: data2}
	}
}

func (m *TLPhoneCodeTransaction) To_PhoneCodeTransaction() *PhoneCodeTransaction {
	m.Data2.PredicateName = Predicate_phoneCodeTransaction
	return m.Data2
}

// // flags
func (m *TLPhoneCodeTransaction) SetAuthKeyId(v int64) { m.Data2.AuthKeyId = v }
func (m *TLPhoneCodeTransaction) GetAuthKeyId() int64  { return m.Data2.AuthKeyId }

func (m *TLPhoneCodeTransaction) SetSessionId(v int64) { m.Data2.SessionId = v }
func (m *TLPhoneCodeTransaction) GetSessionId() int64  { return m.Data2.SessionId }

func (m *TLPhoneCodeTransaction) SetPhone(v string) { m.Data2.Phone = v }
func (m *TLPhoneCodeTransaction) GetPhone() string  { return m.Data2.Phone }

func (m *TLPhoneCodeTransaction) SetPhoneNumberRegistered(v bool) { m.Data2.PhoneNumberRegistered = v }
func (m *TLPhoneCodeTransaction) GetPhoneNumberRegistered() bool {
	return m.Data2.PhoneNumberRegistered
}

func (m *TLPhoneCodeTransaction) SetPhoneCode(v string) { m.Data2.PhoneCode = v }
func (m *TLPhoneCodeTransaction) GetPhoneCode() string  { return m.Data2.PhoneCode }

func (m *TLPhoneCodeTransaction) SetPhoneCodeHash(v string) { m.Data2.PhoneCodeHash = v }
func (m *TLPhoneCodeTransaction) GetPhoneCodeHash() string  { return m.Data2.PhoneCodeHash }

func (m *TLPhoneCodeTransaction) SetPhoneCodeExpired(v int32) { m.Data2.PhoneCodeExpired = v }
func (m *TLPhoneCodeTransaction) GetPhoneCodeExpired() int32  { return m.Data2.PhoneCodeExpired }

func (m *TLPhoneCodeTransaction) SetPhoneCodeExtraData(v string) { m.Data2.PhoneCodeExtraData = v }
func (m *TLPhoneCodeTransaction) GetPhoneCodeExtraData() string  { return m.Data2.PhoneCodeExtraData }

func (m *TLPhoneCodeTransaction) SetSentCodeType(v int32) { m.Data2.SentCodeType = v }
func (m *TLPhoneCodeTransaction) GetSentCodeType() int32  { return m.Data2.SentCodeType }

func (m *TLPhoneCodeTransaction) SetFlashCallPattern(v string) { m.Data2.FlashCallPattern = v }
func (m *TLPhoneCodeTransaction) GetFlashCallPattern() string  { return m.Data2.FlashCallPattern }

func (m *TLPhoneCodeTransaction) SetNextCodeType(v int32) { m.Data2.NextCodeType = v }
func (m *TLPhoneCodeTransaction) GetNextCodeType() int32  { return m.Data2.NextCodeType }

func (m *TLPhoneCodeTransaction) SetState(v int32) { m.Data2.State = v }
func (m *TLPhoneCodeTransaction) GetState() int32  { return m.Data2.State }

func (m *TLPhoneCodeTransaction) GetPredicateName() string {
	return Predicate_phoneCodeTransaction
}

func (m *TLPhoneCodeTransaction) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x83739698: func() error {
			x.UInt(0x83739698)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.GetPhoneNumberRegistered() == true {
					flags |= 1 << 0
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.UInt(flags)
			x.Long(m.GetAuthKeyId())
			x.Long(m.GetSessionId())
			x.String(m.GetPhone())
			x.String(m.GetPhoneCode())
			x.String(m.GetPhoneCodeHash())
			x.Int(m.GetPhoneCodeExpired())
			x.String(m.GetPhoneCodeExtraData())
			x.Int(m.GetSentCodeType())
			x.String(m.GetFlashCallPattern())
			x.Int(m.GetNextCodeType())
			x.Int(m.GetState())
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_phoneCodeTransaction, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_phoneCodeTransaction, layer)
		return nil
	}

	return nil
}

func (m *TLPhoneCodeTransaction) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLPhoneCodeTransaction) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x83739698: func() error {
			var flags = dBuf.UInt()
			_ = flags
			m.SetAuthKeyId(dBuf.Long())
			m.SetSessionId(dBuf.Long())
			m.SetPhone(dBuf.String())
			if (flags & (1 << 0)) != 0 {
				m.SetPhoneNumberRegistered(true)
			}
			m.SetPhoneCode(dBuf.String())
			m.SetPhoneCodeHash(dBuf.String())
			m.SetPhoneCodeExpired(dBuf.Int())
			m.SetPhoneCodeExtraData(dBuf.String())
			m.SetSentCodeType(dBuf.Int())
			m.SetFlashCallPattern(dBuf.String())
			m.SetNextCodeType(dBuf.Int())
			m.SetState(dBuf.Int())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

//----------------------------------------------------------------------------------------------------------------
// TLCodeCreatePhoneCode
///////////////////////////////////////////////////////////////////////////////

func (m *TLCodeCreatePhoneCode) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x6023e09e:
		x.UInt(0x6023e09e)

		// set flags
		var flags uint32 = 0

		if m.GetPhoneNumberRegistered() == true {
			flags |= 1 << 0
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetAuthKeyId())
		x.Long(m.GetSessionId())
		x.String(m.GetPhone())
		x.Int(m.GetSentCodeType())
		x.Int(m.GetNextCodeType())
		x.Int(m.GetState())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLCodeCreatePhoneCode) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLCodeCreatePhoneCode) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x6023e09e:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.AuthKeyId = dBuf.Long()
		m.SessionId = dBuf.Long()
		m.Phone = dBuf.String()
		if (flags & (1 << 0)) != 0 {
			m.PhoneNumberRegistered = true
		}
		m.SentCodeType = dBuf.Int()
		m.NextCodeType = dBuf.Int()
		m.State = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLCodeGetPhoneCode
///////////////////////////////////////////////////////////////////////////////

func (m *TLCodeGetPhoneCode) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x61a4a0f9:
		x.UInt(0x61a4a0f9)

		// no flags

		x.Long(m.GetAuthKeyId())
		x.String(m.GetPhone())
		x.String(m.GetPhoneCodeHash())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLCodeGetPhoneCode) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLCodeGetPhoneCode) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x61a4a0f9:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		m.Phone = dBuf.String()
		m.PhoneCodeHash = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLCodeDeletePhoneCode
///////////////////////////////////////////////////////////////////////////////

func (m *TLCodeDeletePhoneCode) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xa6b06a50:
		x.UInt(0xa6b06a50)

		// no flags

		x.Long(m.GetAuthKeyId())
		x.String(m.GetPhone())
		x.String(m.GetPhoneCodeHash())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLCodeDeletePhoneCode) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLCodeDeletePhoneCode) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xa6b06a50:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		m.Phone = dBuf.String()
		m.PhoneCodeHash = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLCodeUpdatePhoneCodeData
///////////////////////////////////////////////////////////////////////////////

func (m *TLCodeUpdatePhoneCodeData) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xb6950a95:
		x.UInt(0xb6950a95)

		// no flags

		x.Long(m.GetAuthKeyId())
		x.String(m.GetPhone())
		x.String(m.GetPhoneCodeHash())
		m.GetCodeData().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLCodeUpdatePhoneCodeData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLCodeUpdatePhoneCodeData) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb6950a95:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		m.Phone = dBuf.String()
		m.PhoneCodeHash = dBuf.String()

		m4 := &PhoneCodeTransaction{}
		m4.Decode(dBuf)
		m.CodeData = m4

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}
