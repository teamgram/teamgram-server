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
		x := &TLUsernameData{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeUsernameData - unexpected clazzId: %d", id)
	}
}

// TLUsernameData <--
type TLUsernameData struct {
	ClazzID  uint32   `json:"_id"`
	Username string   `json:"username"`
	Peer     *tg.Peer `json:"peer"`
}

func (m *TLUsernameData) String() string {
	wrapper := iface.WithNameWrapper{"usernameData", m}
	return wrapper.String()
}

// UsernameDataClazzName <--
func (m *TLUsernameData) UsernameDataClazzName() string {
	return ClazzName_usernameData
}

// ClazzName <--
func (m *TLUsernameData) ClazzName() string {
	return ClazzName_usernameData
}

// ToUsernameData <--
func (m *TLUsernameData) ToUsernameData() *UsernameData {
	return MakeUsernameData(m)
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
				m2 := &tg.Peer{}
				_ = m2.Decode(d)
				m.Peer = m2
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
	UsernameDataClazz `json:"_clazz"`
}

func (m *UsernameData) String() string {
	wrapper := iface.WithNameWrapper{m.UsernameDataClazzName(), m}
	return wrapper.String()
}

// MakeUsernameData <--
func MakeUsernameData(c UsernameDataClazz) *UsernameData {
	return &UsernameData{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		UsernameDataClazz: c,
	}
}

// Encode <--
func (m *UsernameData) Encode(x *bin.Encoder, layer int32) error {
	if m.UsernameDataClazz != nil {
		return m.UsernameDataClazz.Encode(x, layer)
	}

	return fmt.Errorf("UsernameData - invalid Clazz")
}

// Decode <--
func (m *UsernameData) Decode(d *bin.Decoder) (err error) {
	m.UsernameDataClazz, err = DecodeUsernameDataClazz(d)
	return
}

// Match <--
func (m *UsernameData) Match(f ...interface{}) {
	switch c := m.UsernameDataClazz.(type) {
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
	if m.UsernameDataClazz == nil {
		return nil, false
	}

	if x, ok := m.UsernameDataClazz.(*TLUsernameData); ok {
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
		x := &TLUsernameNotExisted{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_usernameExisted:
		x := &TLUsernameExisted{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_usernameExistedNotMe:
		x := &TLUsernameExistedNotMe{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_usernameExistedIsMe:
		x := &TLUsernameExistedIsMe{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeUsernameExist - unexpected clazzId: %d", id)
	}
}

// TLUsernameNotExisted <--
type TLUsernameNotExisted struct {
	ClazzID uint32 `json:"_id"`
}

func (m *TLUsernameNotExisted) String() string {
	wrapper := iface.WithNameWrapper{"usernameNotExisted", m}
	return wrapper.String()
}

// UsernameExistClazzName <--
func (m *TLUsernameNotExisted) UsernameExistClazzName() string {
	return ClazzName_usernameNotExisted
}

// ClazzName <--
func (m *TLUsernameNotExisted) ClazzName() string {
	return ClazzName_usernameNotExisted
}

// ToUsernameExist <--
func (m *TLUsernameNotExisted) ToUsernameExist() *UsernameExist {
	return MakeUsernameExist(m)
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
	ClazzID uint32 `json:"_id"`
}

func (m *TLUsernameExisted) String() string {
	wrapper := iface.WithNameWrapper{"usernameExisted", m}
	return wrapper.String()
}

// UsernameExistClazzName <--
func (m *TLUsernameExisted) UsernameExistClazzName() string {
	return ClazzName_usernameExisted
}

// ClazzName <--
func (m *TLUsernameExisted) ClazzName() string {
	return ClazzName_usernameExisted
}

// ToUsernameExist <--
func (m *TLUsernameExisted) ToUsernameExist() *UsernameExist {
	return MakeUsernameExist(m)
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
	ClazzID uint32 `json:"_id"`
}

func (m *TLUsernameExistedNotMe) String() string {
	wrapper := iface.WithNameWrapper{"usernameExistedNotMe", m}
	return wrapper.String()
}

// UsernameExistClazzName <--
func (m *TLUsernameExistedNotMe) UsernameExistClazzName() string {
	return ClazzName_usernameExistedNotMe
}

// ClazzName <--
func (m *TLUsernameExistedNotMe) ClazzName() string {
	return ClazzName_usernameExistedNotMe
}

// ToUsernameExist <--
func (m *TLUsernameExistedNotMe) ToUsernameExist() *UsernameExist {
	return MakeUsernameExist(m)
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
	ClazzID uint32 `json:"_id"`
}

func (m *TLUsernameExistedIsMe) String() string {
	wrapper := iface.WithNameWrapper{"usernameExistedIsMe", m}
	return wrapper.String()
}

// UsernameExistClazzName <--
func (m *TLUsernameExistedIsMe) UsernameExistClazzName() string {
	return ClazzName_usernameExistedIsMe
}

// ClazzName <--
func (m *TLUsernameExistedIsMe) ClazzName() string {
	return ClazzName_usernameExistedIsMe
}

// ToUsernameExist <--
func (m *TLUsernameExistedIsMe) ToUsernameExist() *UsernameExist {
	return MakeUsernameExist(m)
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
	UsernameExistClazz `json:"_clazz"`
}

func (m *UsernameExist) String() string {
	wrapper := iface.WithNameWrapper{m.UsernameExistClazzName(), m}
	return wrapper.String()
}

// MakeUsernameExist <--
func MakeUsernameExist(c UsernameExistClazz) *UsernameExist {
	return &UsernameExist{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		UsernameExistClazz: c,
	}
}

// Encode <--
func (m *UsernameExist) Encode(x *bin.Encoder, layer int32) error {
	if m.UsernameExistClazz != nil {
		return m.UsernameExistClazz.Encode(x, layer)
	}

	return fmt.Errorf("UsernameExist - invalid Clazz")
}

// Decode <--
func (m *UsernameExist) Decode(d *bin.Decoder) (err error) {
	m.UsernameExistClazz, err = DecodeUsernameExistClazz(d)
	return
}

// Match <--
func (m *UsernameExist) Match(f ...interface{}) {
	switch c := m.UsernameExistClazz.(type) {
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
	if m.UsernameExistClazz == nil {
		return nil, false
	}

	if x, ok := m.UsernameExistClazz.(*TLUsernameNotExisted); ok {
		return x, true
	}

	return nil, false
}

// ToUsernameExisted <--
func (m *UsernameExist) ToUsernameExisted() (*TLUsernameExisted, bool) {
	if m.UsernameExistClazz == nil {
		return nil, false
	}

	if x, ok := m.UsernameExistClazz.(*TLUsernameExisted); ok {
		return x, true
	}

	return nil, false
}

// ToUsernameExistedNotMe <--
func (m *UsernameExist) ToUsernameExistedNotMe() (*TLUsernameExistedNotMe, bool) {
	if m.UsernameExistClazz == nil {
		return nil, false
	}

	if x, ok := m.UsernameExistClazz.(*TLUsernameExistedNotMe); ok {
		return x, true
	}

	return nil, false
}

// ToUsernameExistedIsMe <--
func (m *UsernameExist) ToUsernameExistedIsMe() (*TLUsernameExistedIsMe, bool) {
	if m.UsernameExistClazz == nil {
		return nil, false
	}

	if x, ok := m.UsernameExistClazz.(*TLUsernameExistedIsMe); ok {
		return x, true
	}

	return nil, false
}
