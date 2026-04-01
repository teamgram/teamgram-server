/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package echo

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

// EchoClazz <--
//   - TL_Echo
//   - TL_Echo2
type EchoClazz interface {
	iface.TLObject
	EchoClazzName() string
}

func DecodeEchoClazz(d *bin.Decoder) (EchoClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x2e3ba51e:
		x := &TLEcho{ClazzID: id, ClazzName2: ClazzName_echo}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x2249c1b:
		x := &TLEcho2{ClazzID: id, ClazzName2: ClazzName_echo2}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeEcho - unexpected clazzId: %d", id)
	}

}

// TLEcho <--
type TLEcho struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Message    string `json:"message"`
}

func MakeTLEcho(m *TLEcho) *TLEcho {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_echo

	return m
}

func (m *TLEcho) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLEcho) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("echo", m)
}

// EchoClazzName <--
func (m *TLEcho) EchoClazzName() string {
	return ClazzName_echo
}

// ClazzName <--
func (m *TLEcho) ClazzName() string {
	return m.ClazzName2
}

// ToEcho <--
func (m *TLEcho) ToEcho() *Echo {
	if m == nil {
		return nil
	}

	return &Echo{Clazz: m}

}

func (m *TLEcho) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_echo, int(layer)); clazzId {
	case 0x2e3ba51e:
		size := 4
		size += iface.CalcStringSize(m.Message)

		return size
	default:
		return 0
	}
}

func (m *TLEcho) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_echo, int(layer)); clazzId {
	case 0x2e3ba51e:
		if err := iface.ValidateRequiredString("message", m.Message); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_echo, layer)
	}
}

// Encode <--
func (m *TLEcho) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_echo, int(layer)); clazzId {
	case 0x2e3ba51e:
		x.PutClazzID(0x2e3ba51e)

		x.PutString(m.Message)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_echo, layer)
	}
}

// Decode <--
func (m *TLEcho) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x2e3ba51e:
		m.Message, err = d.String()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLEcho2 <--
type TLEcho2 struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Message    string `json:"message"`
}

func MakeTLEcho2(m *TLEcho2) *TLEcho2 {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_echo2

	return m
}

func (m *TLEcho2) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLEcho2) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("echo2", m)
}

// EchoClazzName <--
func (m *TLEcho2) EchoClazzName() string {
	return ClazzName_echo2
}

// ClazzName <--
func (m *TLEcho2) ClazzName() string {
	return m.ClazzName2
}

// ToEcho <--
func (m *TLEcho2) ToEcho() *Echo {
	if m == nil {
		return nil
	}

	return &Echo{Clazz: m}

}

func (m *TLEcho2) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_echo2, int(layer)); clazzId {
	case 0x2249c1b:
		size := 4
		size += iface.CalcStringSize(m.Message)

		return size
	default:
		return 0
	}
}

func (m *TLEcho2) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_echo2, int(layer)); clazzId {
	case 0x2249c1b:
		if err := iface.ValidateRequiredString("message", m.Message); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_echo2, layer)
	}
}

// Encode <--
func (m *TLEcho2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_echo2, int(layer)); clazzId {
	case 0x2249c1b:
		x.PutClazzID(0x2249c1b)

		x.PutString(m.Message)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_echo2, layer)
	}
}

// Decode <--
func (m *TLEcho2) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x2249c1b:
		m.Message, err = d.String()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// Echo <--
type Echo struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz EchoClazz `json:"_clazz"`
}

func (m *Echo) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *Echo) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName(m.ClazzName(), m)
}

func (m *Echo) CalcSize(layer int32) int {
	if m == nil || m.Clazz == nil {
		return 0
	}
	return iface.CalcObjectSize(m.Clazz, layer)
}

func (m *Echo) Validate(layer int32) error {
	if m == nil {
		return fmt.Errorf("Echo is required")
	}
	if m.Clazz == nil {
		return fmt.Errorf("Echo.Clazz is required")
	}
	if v, ok := m.Clazz.(iface.TLObjectValidator); ok {
		return v.Validate(layer)
	}
	return nil
}

func (m *Echo) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.EchoClazzName()
	}
}

// Encode <--
func (m *Echo) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("Echo - invalid Clazz")
}

// Decode <--
func (m *Echo) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeEchoClazz(d)
	return
}

// ToEcho <--
func (m *Echo) ToEcho() (*TLEcho, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLEcho); ok {
		return x, true
	}

	return nil, false
}

// ToEcho2 <--
func (m *Echo) ToEcho2() (*TLEcho2, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLEcho2); ok {
		return x, true
	}

	return nil, false
}
