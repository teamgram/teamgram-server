/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package idgen

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

// IdValClazz <--
//   - TL_IdVal
//   - TL_IdVals
//   - TL_SeqIdVal
type IdValClazz interface {
	iface.TLObject
	IdValClazzName() string
}

func DecodeIdValClazz(d *bin.Decoder) (IdValClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_idVal:
		x := &TLIdVal{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_idVals:
		x := &TLIdVals{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_seqIdVal:
		x := &TLSeqIdVal{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeIdVal - unexpected clazzId: %d", id)
	}
}

// TLIdVal <--
type TLIdVal struct {
	ClazzID  uint32 `json:"_id"`
	Id_INT64 int64  `json:"id_INT64"`
}

func (m *TLIdVal) String() string {
	wrapper := iface.WithNameWrapper{"idVal", m}
	return wrapper.String()
}

// IdValClazzName <--
func (m *TLIdVal) IdValClazzName() string {
	return ClazzName_idVal
}

// ClazzName <--
func (m *TLIdVal) ClazzName() string {
	return ClazzName_idVal
}

// ToIdVal <--
func (m *TLIdVal) ToIdVal() *IdVal {
	if m == nil {
		return nil
	}

	return MakeIdVal(m)
}

// Encode <--
func (m *TLIdVal) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc07844cb: func() error {
			x.PutClazzID(0xc07844cb)

			x.PutInt64(m.Id_INT64)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_idVal, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idVal, layer)
	}
}

// Decode <--
func (m *TLIdVal) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xc07844cb: func() (err error) {
			m.Id_INT64, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLIdVals <--
type TLIdVals struct {
	ClazzID        uint32  `json:"_id"`
	Id_VECTORINT64 []int64 `json:"id_VECTORINT64"`
}

func (m *TLIdVals) String() string {
	wrapper := iface.WithNameWrapper{"idVals", m}
	return wrapper.String()
}

// IdValClazzName <--
func (m *TLIdVals) IdValClazzName() string {
	return ClazzName_idVals
}

// ClazzName <--
func (m *TLIdVals) ClazzName() string {
	return ClazzName_idVals
}

// ToIdVal <--
func (m *TLIdVals) ToIdVal() *IdVal {
	if m == nil {
		return nil
	}

	return MakeIdVal(m)
}

// Encode <--
func (m *TLIdVals) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x1c3baa66: func() error {
			x.PutClazzID(0x1c3baa66)

			iface.EncodeInt64List(x, m.Id_VECTORINT64)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_idVals, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idVals, layer)
	}
}

// Decode <--
func (m *TLIdVals) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x1c3baa66: func() (err error) {

			m.Id_VECTORINT64, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLSeqIdVal <--
type TLSeqIdVal struct {
	ClazzID  uint32 `json:"_id"`
	Id_INT64 int64  `json:"id_INT64"`
}

func (m *TLSeqIdVal) String() string {
	wrapper := iface.WithNameWrapper{"seqIdVal", m}
	return wrapper.String()
}

// IdValClazzName <--
func (m *TLSeqIdVal) IdValClazzName() string {
	return ClazzName_seqIdVal
}

// ClazzName <--
func (m *TLSeqIdVal) ClazzName() string {
	return ClazzName_seqIdVal
}

// ToIdVal <--
func (m *TLSeqIdVal) ToIdVal() *IdVal {
	if m == nil {
		return nil
	}

	return MakeIdVal(m)
}

// Encode <--
func (m *TLSeqIdVal) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x2a047d08: func() error {
			x.PutClazzID(0x2a047d08)

			x.PutInt64(m.Id_INT64)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_seqIdVal, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_seqIdVal, layer)
	}
}

// Decode <--
func (m *TLSeqIdVal) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x2a047d08: func() (err error) {
			m.Id_INT64, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// IdVal <--
type IdVal struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	IdValClazz `json:"_clazz"`
}

func (m *IdVal) String() string {
	wrapper := iface.WithNameWrapper{m.IdValClazzName(), m}
	return wrapper.String()
}

// MakeIdVal <--
func MakeIdVal(c IdValClazz) *IdVal {
	return &IdVal{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		IdValClazz: c,
	}
}

// Encode <--
func (m *IdVal) Encode(x *bin.Encoder, layer int32) error {
	if m.IdValClazz != nil {
		return m.IdValClazz.Encode(x, layer)
	}

	return fmt.Errorf("IdVal - invalid Clazz")
}

// Decode <--
func (m *IdVal) Decode(d *bin.Decoder) (err error) {
	m.IdValClazz, err = DecodeIdValClazz(d)
	return
}

// Match <--
func (m *IdVal) Match(f ...interface{}) {
	switch c := m.IdValClazz.(type) {
	case *TLIdVal:
		for _, v := range f {
			if f1, ok := v.(func(c *TLIdVal) interface{}); ok {
				f1(c)
			}
		}
	case *TLIdVals:
		for _, v := range f {
			if f1, ok := v.(func(c *TLIdVals) interface{}); ok {
				f1(c)
			}
		}
	case *TLSeqIdVal:
		for _, v := range f {
			if f1, ok := v.(func(c *TLSeqIdVal) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToIdVal <--
func (m *IdVal) ToIdVal() (*TLIdVal, bool) {
	if m == nil {
		return nil, false
	}

	if m.IdValClazz == nil {
		return nil, false
	}

	if x, ok := m.IdValClazz.(*TLIdVal); ok {
		return x, true
	}

	return nil, false
}

// ToIdVals <--
func (m *IdVal) ToIdVals() (*TLIdVals, bool) {
	if m == nil {
		return nil, false
	}

	if m.IdValClazz == nil {
		return nil, false
	}

	if x, ok := m.IdValClazz.(*TLIdVals); ok {
		return x, true
	}

	return nil, false
}

// ToSeqIdVal <--
func (m *IdVal) ToSeqIdVal() (*TLSeqIdVal, bool) {
	if m == nil {
		return nil, false
	}

	if m.IdValClazz == nil {
		return nil, false
	}

	if x, ok := m.IdValClazz.(*TLSeqIdVal); ok {
		return x, true
	}

	return nil, false
}

// InputIdClazz <--
//   - TL_InputId
//   - TL_InputIds
//   - TL_InputSeqId
//   - TL_InputNSeqId
type InputIdClazz interface {
	iface.TLObject
	InputIdClazzName() string
}

func DecodeInputIdClazz(d *bin.Decoder) (InputIdClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_inputId:
		x := &TLInputId{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_inputIds:
		x := &TLInputIds{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_inputSeqId:
		x := &TLInputSeqId{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_inputNSeqId:
		x := &TLInputNSeqId{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeInputId - unexpected clazzId: %d", id)
	}
}

// TLInputId <--
type TLInputId struct {
	ClazzID uint32 `json:"_id"`
}

func (m *TLInputId) String() string {
	wrapper := iface.WithNameWrapper{"inputId", m}
	return wrapper.String()
}

// InputIdClazzName <--
func (m *TLInputId) InputIdClazzName() string {
	return ClazzName_inputId
}

// ClazzName <--
func (m *TLInputId) ClazzName() string {
	return ClazzName_inputId
}

// ToInputId <--
func (m *TLInputId) ToInputId() *InputId {
	if m == nil {
		return nil
	}

	return MakeInputId(m)
}

// Encode <--
func (m *TLInputId) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x8af2196c: func() error {
			x.PutClazzID(0x8af2196c)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inputId, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inputId, layer)
	}
}

// Decode <--
func (m *TLInputId) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x8af2196c: func() (err error) {

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLInputIds <--
type TLInputIds struct {
	ClazzID uint32 `json:"_id"`
	Num     int32  `json:"num"`
}

func (m *TLInputIds) String() string {
	wrapper := iface.WithNameWrapper{"inputIds", m}
	return wrapper.String()
}

// InputIdClazzName <--
func (m *TLInputIds) InputIdClazzName() string {
	return ClazzName_inputIds
}

// ClazzName <--
func (m *TLInputIds) ClazzName() string {
	return ClazzName_inputIds
}

// ToInputId <--
func (m *TLInputIds) ToInputId() *InputId {
	if m == nil {
		return nil
	}

	return MakeInputId(m)
}

// Encode <--
func (m *TLInputIds) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x7f285fbc: func() error {
			x.PutClazzID(0x7f285fbc)

			x.PutInt32(m.Num)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inputIds, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inputIds, layer)
	}
}

// Decode <--
func (m *TLInputIds) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x7f285fbc: func() (err error) {
			m.Num, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLInputSeqId <--
type TLInputSeqId struct {
	ClazzID uint32 `json:"_id"`
	Key     string `json:"key"`
}

func (m *TLInputSeqId) String() string {
	wrapper := iface.WithNameWrapper{"inputSeqId", m}
	return wrapper.String()
}

// InputIdClazzName <--
func (m *TLInputSeqId) InputIdClazzName() string {
	return ClazzName_inputSeqId
}

// ClazzName <--
func (m *TLInputSeqId) ClazzName() string {
	return ClazzName_inputSeqId
}

// ToInputId <--
func (m *TLInputSeqId) ToInputId() *InputId {
	if m == nil {
		return nil
	}

	return MakeInputId(m)
}

// Encode <--
func (m *TLInputSeqId) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xcd52bbcd: func() error {
			x.PutClazzID(0xcd52bbcd)

			x.PutString(m.Key)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inputSeqId, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inputSeqId, layer)
	}
}

// Decode <--
func (m *TLInputSeqId) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xcd52bbcd: func() (err error) {
			m.Key, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLInputNSeqId <--
type TLInputNSeqId struct {
	ClazzID uint32 `json:"_id"`
	Key     string `json:"key"`
	N       int32  `json:"n"`
}

func (m *TLInputNSeqId) String() string {
	wrapper := iface.WithNameWrapper{"inputNSeqId", m}
	return wrapper.String()
}

// InputIdClazzName <--
func (m *TLInputNSeqId) InputIdClazzName() string {
	return ClazzName_inputNSeqId
}

// ClazzName <--
func (m *TLInputNSeqId) ClazzName() string {
	return ClazzName_inputNSeqId
}

// ToInputId <--
func (m *TLInputNSeqId) ToInputId() *InputId {
	if m == nil {
		return nil
	}

	return MakeInputId(m)
}

// Encode <--
func (m *TLInputNSeqId) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x7ab16d81: func() error {
			x.PutClazzID(0x7ab16d81)

			x.PutString(m.Key)
			x.PutInt32(m.N)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inputNSeqId, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inputNSeqId, layer)
	}
}

// Decode <--
func (m *TLInputNSeqId) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x7ab16d81: func() (err error) {
			m.Key, err = d.String()
			m.N, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// InputId <--
type InputId struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	InputIdClazz `json:"_clazz"`
}

func (m *InputId) String() string {
	wrapper := iface.WithNameWrapper{m.InputIdClazzName(), m}
	return wrapper.String()
}

// MakeInputId <--
func MakeInputId(c InputIdClazz) *InputId {
	return &InputId{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		InputIdClazz: c,
	}
}

// Encode <--
func (m *InputId) Encode(x *bin.Encoder, layer int32) error {
	if m.InputIdClazz != nil {
		return m.InputIdClazz.Encode(x, layer)
	}

	return fmt.Errorf("InputId - invalid Clazz")
}

// Decode <--
func (m *InputId) Decode(d *bin.Decoder) (err error) {
	m.InputIdClazz, err = DecodeInputIdClazz(d)
	return
}

// Match <--
func (m *InputId) Match(f ...interface{}) {
	switch c := m.InputIdClazz.(type) {
	case *TLInputId:
		for _, v := range f {
			if f1, ok := v.(func(c *TLInputId) interface{}); ok {
				f1(c)
			}
		}
	case *TLInputIds:
		for _, v := range f {
			if f1, ok := v.(func(c *TLInputIds) interface{}); ok {
				f1(c)
			}
		}
	case *TLInputSeqId:
		for _, v := range f {
			if f1, ok := v.(func(c *TLInputSeqId) interface{}); ok {
				f1(c)
			}
		}
	case *TLInputNSeqId:
		for _, v := range f {
			if f1, ok := v.(func(c *TLInputNSeqId) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToInputId <--
func (m *InputId) ToInputId() (*TLInputId, bool) {
	if m == nil {
		return nil, false
	}

	if m.InputIdClazz == nil {
		return nil, false
	}

	if x, ok := m.InputIdClazz.(*TLInputId); ok {
		return x, true
	}

	return nil, false
}

// ToInputIds <--
func (m *InputId) ToInputIds() (*TLInputIds, bool) {
	if m == nil {
		return nil, false
	}

	if m.InputIdClazz == nil {
		return nil, false
	}

	if x, ok := m.InputIdClazz.(*TLInputIds); ok {
		return x, true
	}

	return nil, false
}

// ToInputSeqId <--
func (m *InputId) ToInputSeqId() (*TLInputSeqId, bool) {
	if m == nil {
		return nil, false
	}

	if m.InputIdClazz == nil {
		return nil, false
	}

	if x, ok := m.InputIdClazz.(*TLInputSeqId); ok {
		return x, true
	}

	return nil, false
}

// ToInputNSeqId <--
func (m *InputId) ToInputNSeqId() (*TLInputNSeqId, bool) {
	if m == nil {
		return nil, false
	}

	if m.InputIdClazz == nil {
		return nil, false
	}

	if x, ok := m.InputIdClazz.(*TLInputNSeqId); ok {
		return x, true
	}

	return nil, false
}
