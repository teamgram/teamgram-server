/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package echo1

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
		x := &TLEcho{ClazzID: id, ClazzName2: ClazzName_echo}
		_ = x.Decode(d)
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
	Clazz EchoClazz `json:"_clazz"`
}

func (m *Echo) String() string {
	data, _ := json.Marshal(m)
	return string(data)
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

// Match <--
func (m *Echo) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
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
