/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package username

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

// UsernameDataClazz <--
//   - TL_UsernameData
type UsernameDataClazz interface {
	iface.TLObject
	UsernameDataClazzName() string
}

func DecodeUsernameDataClazz(d *bin.Decoder) (UsernameDataClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_usernameData:
		x := &TLUsernameData{ClazzID: id, ClazzName2: ClazzName_usernameData}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeUsernameData - unexpected clazzId: %d", id)
	}
}

// TLUsernameData <--
type TLUsernameData struct {
	ClazzID    uint32       `json:"_id"`
	ClazzName2 string       `json:"_name"`
	Username   string       `json:"username"`
	Peer       tg.PeerClazz `json:"peer"`
}

func MakeTLUsernameData(m *TLUsernameData) *TLUsernameData {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_usernameData

	return m
}

func (m *TLUsernameData) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// UsernameDataClazzName <--
func (m *TLUsernameData) UsernameDataClazzName() string {
	return ClazzName_usernameData
}

// ClazzName <--
func (m *TLUsernameData) ClazzName() string {
	return m.ClazzName2
}

// ToUsernameData <--
func (m *TLUsernameData) ToUsernameData() *UsernameData {
	if m == nil {
		return nil
	}

	return &UsernameData{Clazz: m}
}

// Encode <--
func (m *TLUsernameData) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xaa4000bf: func() error {
			x.PutClazzID(0xaa4000bf)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.Peer != nil {
					flags |= 1 << 0
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutString(m.Username)
			if m.Peer != nil {
				_ = m.Peer.Encode(x, layer)
			}

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_usernameData, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usernameData, layer)
	}
}

// Decode <--
func (m *TLUsernameData) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xaa4000bf: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.Username, err = d.String()
			if (flags & (1 << 0)) != 0 {
				// m2 := &tg.Peer{}
				// _ = m2.Decode(d)
				// m.Peer = m2
				m.Peer, _ = tg.DecodePeerClazz(d)
			}

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// UsernameData <--
type UsernameData struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz UsernameDataClazz `json:"_clazz"`
}

func (m *UsernameData) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *UsernameData) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.UsernameDataClazzName()
	}
}

// Encode <--
func (m *UsernameData) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("UsernameData - invalid Clazz")
}

// Decode <--
func (m *UsernameData) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeUsernameDataClazz(d)
	return
}

// Match <--
func (m *UsernameData) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLUsernameData:
		for _, v := range f {
			if f1, ok := v.(func(c *TLUsernameData) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToUsernameData <--
func (m *UsernameData) ToUsernameData() (*TLUsernameData, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLUsernameData); ok {
		return x, true
	}

	return nil, false
}

// UsernameExistClazz <--
//   - TL_UsernameNotExisted
//   - TL_UsernameExisted
//   - TL_UsernameExistedNotMe
//   - TL_UsernameExistedIsMe
type UsernameExistClazz interface {
	iface.TLObject
	UsernameExistClazzName() string
}

func DecodeUsernameExistClazz(d *bin.Decoder) (UsernameExistClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_usernameNotExisted:
		x := &TLUsernameNotExisted{ClazzID: id, ClazzName2: ClazzName_usernameNotExisted}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_usernameExisted:
		x := &TLUsernameExisted{ClazzID: id, ClazzName2: ClazzName_usernameExisted}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_usernameExistedNotMe:
		x := &TLUsernameExistedNotMe{ClazzID: id, ClazzName2: ClazzName_usernameExistedNotMe}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_usernameExistedIsMe:
		x := &TLUsernameExistedIsMe{ClazzID: id, ClazzName2: ClazzName_usernameExistedIsMe}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeUsernameExist - unexpected clazzId: %d", id)
	}
}

// TLUsernameNotExisted <--
type TLUsernameNotExisted struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
}

func MakeTLUsernameNotExisted(m *TLUsernameNotExisted) *TLUsernameNotExisted {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_usernameNotExisted

	return m
}

func (m *TLUsernameNotExisted) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// UsernameExistClazzName <--
func (m *TLUsernameNotExisted) UsernameExistClazzName() string {
	return ClazzName_usernameNotExisted
}

// ClazzName <--
func (m *TLUsernameNotExisted) ClazzName() string {
	return m.ClazzName2
}

// ToUsernameExist <--
func (m *TLUsernameNotExisted) ToUsernameExist() *UsernameExist {
	if m == nil {
		return nil
	}

	return &UsernameExist{Clazz: m}
}

// Encode <--
func (m *TLUsernameNotExisted) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xcb3cfb6d: func() error {
			x.PutClazzID(0xcb3cfb6d)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_usernameNotExisted, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usernameNotExisted, layer)
	}
}

// Decode <--
func (m *TLUsernameNotExisted) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xcb3cfb6d: func() (err error) {

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUsernameExisted <--
type TLUsernameExisted struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
}

func MakeTLUsernameExisted(m *TLUsernameExisted) *TLUsernameExisted {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_usernameExisted

	return m
}

func (m *TLUsernameExisted) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// UsernameExistClazzName <--
func (m *TLUsernameExisted) UsernameExistClazzName() string {
	return ClazzName_usernameExisted
}

// ClazzName <--
func (m *TLUsernameExisted) ClazzName() string {
	return m.ClazzName2
}

// ToUsernameExist <--
func (m *TLUsernameExisted) ToUsernameExist() *UsernameExist {
	if m == nil {
		return nil
	}

	return &UsernameExist{Clazz: m}
}

// Encode <--
func (m *TLUsernameExisted) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xace7f4cd: func() error {
			x.PutClazzID(0xace7f4cd)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_usernameExisted, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usernameExisted, layer)
	}
}

// Decode <--
func (m *TLUsernameExisted) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xace7f4cd: func() (err error) {

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUsernameExistedNotMe <--
type TLUsernameExistedNotMe struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
}

func MakeTLUsernameExistedNotMe(m *TLUsernameExistedNotMe) *TLUsernameExistedNotMe {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_usernameExistedNotMe

	return m
}

func (m *TLUsernameExistedNotMe) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// UsernameExistClazzName <--
func (m *TLUsernameExistedNotMe) UsernameExistClazzName() string {
	return ClazzName_usernameExistedNotMe
}

// ClazzName <--
func (m *TLUsernameExistedNotMe) ClazzName() string {
	return m.ClazzName2
}

// ToUsernameExist <--
func (m *TLUsernameExistedNotMe) ToUsernameExist() *UsernameExist {
	if m == nil {
		return nil
	}

	return &UsernameExist{Clazz: m}
}

// Encode <--
func (m *TLUsernameExistedNotMe) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xd01f47b1: func() error {
			x.PutClazzID(0xd01f47b1)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_usernameExistedNotMe, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usernameExistedNotMe, layer)
	}
}

// Decode <--
func (m *TLUsernameExistedNotMe) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xd01f47b1: func() (err error) {

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUsernameExistedIsMe <--
type TLUsernameExistedIsMe struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
}

func MakeTLUsernameExistedIsMe(m *TLUsernameExistedIsMe) *TLUsernameExistedIsMe {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_usernameExistedIsMe

	return m
}

func (m *TLUsernameExistedIsMe) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// UsernameExistClazzName <--
func (m *TLUsernameExistedIsMe) UsernameExistClazzName() string {
	return ClazzName_usernameExistedIsMe
}

// ClazzName <--
func (m *TLUsernameExistedIsMe) ClazzName() string {
	return m.ClazzName2
}

// ToUsernameExist <--
func (m *TLUsernameExistedIsMe) ToUsernameExist() *UsernameExist {
	if m == nil {
		return nil
	}

	return &UsernameExist{Clazz: m}
}

// Encode <--
func (m *TLUsernameExistedIsMe) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x874e7771: func() error {
			x.PutClazzID(0x874e7771)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_usernameExistedIsMe, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_usernameExistedIsMe, layer)
	}
}

// Decode <--
func (m *TLUsernameExistedIsMe) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x874e7771: func() (err error) {

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// UsernameExist <--
type UsernameExist struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz UsernameExistClazz `json:"_clazz"`
}

func (m *UsernameExist) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *UsernameExist) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.UsernameExistClazzName()
	}
}

// Encode <--
func (m *UsernameExist) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("UsernameExist - invalid Clazz")
}

// Decode <--
func (m *UsernameExist) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeUsernameExistClazz(d)
	return
}

// Match <--
func (m *UsernameExist) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLUsernameNotExisted:
		for _, v := range f {
			if f1, ok := v.(func(c *TLUsernameNotExisted) interface{}); ok {
				f1(c)
			}
		}
	case *TLUsernameExisted:
		for _, v := range f {
			if f1, ok := v.(func(c *TLUsernameExisted) interface{}); ok {
				f1(c)
			}
		}
	case *TLUsernameExistedNotMe:
		for _, v := range f {
			if f1, ok := v.(func(c *TLUsernameExistedNotMe) interface{}); ok {
				f1(c)
			}
		}
	case *TLUsernameExistedIsMe:
		for _, v := range f {
			if f1, ok := v.(func(c *TLUsernameExistedIsMe) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToUsernameNotExisted <--
func (m *UsernameExist) ToUsernameNotExisted() (*TLUsernameNotExisted, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLUsernameNotExisted); ok {
		return x, true
	}

	return nil, false
}

// ToUsernameExisted <--
func (m *UsernameExist) ToUsernameExisted() (*TLUsernameExisted, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLUsernameExisted); ok {
		return x, true
	}

	return nil, false
}

// ToUsernameExistedNotMe <--
func (m *UsernameExist) ToUsernameExistedNotMe() (*TLUsernameExistedNotMe, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLUsernameExistedNotMe); ok {
		return x, true
	}

	return nil, false
}

// ToUsernameExistedIsMe <--
func (m *UsernameExist) ToUsernameExistedIsMe() (*TLUsernameExistedIsMe, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLUsernameExistedIsMe); ok {
		return x, true
	}

	return nil, false
}
