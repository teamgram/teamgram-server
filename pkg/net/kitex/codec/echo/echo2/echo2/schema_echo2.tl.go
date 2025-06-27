/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package echo2

import (
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
)

var _ iface.TLObject
var _ fmt.Stringer
var _ *tg.Bool
var _ bin.Fields

// EchoClazz <--
//   - TL_Echo
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

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_echo:
		x := &TLEcho{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeEcho - unexpected clazzId: %d", id)
	}
}

// TLEcho <--
type TLEcho struct {
	ClazzID uint32 `json:"_id"`
	Message string `json:"message"`
}

// EchoClazzName <--
func (m *TLEcho) EchoClazzName() string {
	return ClazzName_echo
}

// ClazzName <--
func (m *TLEcho) ClazzName() string {
	return ClazzName_echo
}

// ToEcho <--
func (m *TLEcho) ToEcho() *Echo {
	return MakeEcho(m)
}

// Encode <--
func (m *TLEcho) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x2e3ba51e: func() error {
			x.PutClazzID(0x2e3ba51e)

			x.PutString(m.Message)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_echo, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_echo, layer)
	}
}

// Decode <--
func (m *TLEcho) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x2e3ba51e: func() (err error) {
			m.Message, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// Echo <--
type Echo struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	EchoClazz
}

// MakeEcho <--
func MakeEcho(c EchoClazz) *Echo {
	return &Echo{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		EchoClazz: c,
	}
}

// Encode <--
func (m *Echo) Encode(x *bin.Encoder, layer int32) error {
	if m.EchoClazz != nil {
		return m.EchoClazz.Encode(x, layer)
	}

	return fmt.Errorf("Echo - invalid Clazz")
}

// Decode <--
func (m *Echo) Decode(d *bin.Decoder) (err error) {
	m.EchoClazz, err = DecodeEchoClazz(d)
	return
}

// Match <--
func (m *Echo) Match(f ...interface{}) {
	switch c := m.EchoClazz.(type) {
	case *TLEcho:
		for _, v := range f {
			if f1, ok := v.(func(c *TLEcho) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToEcho <--
func (m *Echo) ToEcho() (*TLEcho, bool) {
	if m.EchoClazz == nil {
		return nil, false
	}

	if x, ok := m.EchoClazz.(*TLEcho); ok {
		return x, true
	}

	return nil, false
}
