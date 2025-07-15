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
		x := &TLIdVal{ClazzID: id, ClazzName2: ClazzName_idVal}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_idVals:
		x := &TLIdVals{ClazzID: id, ClazzName2: ClazzName_idVals}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_seqIdVal:
		x := &TLSeqIdVal{ClazzID: id, ClazzName2: ClazzName_seqIdVal}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeIdVal - unexpected clazzId: %d", id)
	}
}

// TLIdVal <--
type TLIdVal struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Id_INT64   int64  `json:"id_INT64"`
}

func MakeTLIdVal(m *TLIdVal) *TLIdVal {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_idVal

	return m
}

func (m *TLIdVal) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// IdValClazzName <--
func (m *TLIdVal) IdValClazzName() string {
	return ClazzName_idVal
}

// ClazzName <--
func (m *TLIdVal) ClazzName() string {
	return m.ClazzName2
}

// ToIdVal <--
func (m *TLIdVal) ToIdVal() *IdVal {
	if m == nil {
		return nil
	}

	return &IdVal{Clazz: m}
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
	ClazzName2     string  `json:"_name"`
	Id_VECTORINT64 []int64 `json:"id_VECTORINT64"`
}

func MakeTLIdVals(m *TLIdVals) *TLIdVals {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_idVals

	return m
}

func (m *TLIdVals) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// IdValClazzName <--
func (m *TLIdVals) IdValClazzName() string {
	return ClazzName_idVals
}

// ClazzName <--
func (m *TLIdVals) ClazzName() string {
	return m.ClazzName2
}

// ToIdVal <--
func (m *TLIdVals) ToIdVal() *IdVal {
	if m == nil {
		return nil
	}

	return &IdVal{Clazz: m}
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
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Id_INT64   int64  `json:"id_INT64"`
}

func MakeTLSeqIdVal(m *TLSeqIdVal) *TLSeqIdVal {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_seqIdVal

	return m
}

func (m *TLSeqIdVal) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// IdValClazzName <--
func (m *TLSeqIdVal) IdValClazzName() string {
	return ClazzName_seqIdVal
}

// ClazzName <--
func (m *TLSeqIdVal) ClazzName() string {
	return m.ClazzName2
}

// ToIdVal <--
func (m *TLSeqIdVal) ToIdVal() *IdVal {
	if m == nil {
		return nil
	}

	return &IdVal{Clazz: m}
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
	Clazz IdValClazz `json:"_clazz"`
}

func (m *IdVal) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *IdVal) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.IdValClazzName()
	}
}

// Encode <--
func (m *IdVal) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("IdVal - invalid Clazz")
}

// Decode <--
func (m *IdVal) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeIdValClazz(d)
	return
}

// Match <--
func (m *IdVal) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
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

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLIdVal); ok {
		return x, true
	}

	return nil, false
}

// ToIdVals <--
func (m *IdVal) ToIdVals() (*TLIdVals, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLIdVals); ok {
		return x, true
	}

	return nil, false
}

// ToSeqIdVal <--
func (m *IdVal) ToSeqIdVal() (*TLSeqIdVal, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLSeqIdVal); ok {
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
		x := &TLInputId{ClazzID: id, ClazzName2: ClazzName_inputId}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_inputIds:
		x := &TLInputIds{ClazzID: id, ClazzName2: ClazzName_inputIds}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_inputSeqId:
		x := &TLInputSeqId{ClazzID: id, ClazzName2: ClazzName_inputSeqId}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_inputNSeqId:
		x := &TLInputNSeqId{ClazzID: id, ClazzName2: ClazzName_inputNSeqId}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeInputId - unexpected clazzId: %d", id)
	}
}

// TLInputId <--
type TLInputId struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
}

func MakeTLInputId(m *TLInputId) *TLInputId {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_inputId

	return m
}

func (m *TLInputId) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// InputIdClazzName <--
func (m *TLInputId) InputIdClazzName() string {
	return ClazzName_inputId
}

// ClazzName <--
func (m *TLInputId) ClazzName() string {
	return m.ClazzName2
}

// ToInputId <--
func (m *TLInputId) ToInputId() *InputId {
	if m == nil {
		return nil
	}

	return &InputId{Clazz: m}
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
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Num        int32  `json:"num"`
}

func MakeTLInputIds(m *TLInputIds) *TLInputIds {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_inputIds

	return m
}

func (m *TLInputIds) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// InputIdClazzName <--
func (m *TLInputIds) InputIdClazzName() string {
	return ClazzName_inputIds
}

// ClazzName <--
func (m *TLInputIds) ClazzName() string {
	return m.ClazzName2
}

// ToInputId <--
func (m *TLInputIds) ToInputId() *InputId {
	if m == nil {
		return nil
	}

	return &InputId{Clazz: m}
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
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Key        string `json:"key"`
}

func MakeTLInputSeqId(m *TLInputSeqId) *TLInputSeqId {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_inputSeqId

	return m
}

func (m *TLInputSeqId) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// InputIdClazzName <--
func (m *TLInputSeqId) InputIdClazzName() string {
	return ClazzName_inputSeqId
}

// ClazzName <--
func (m *TLInputSeqId) ClazzName() string {
	return m.ClazzName2
}

// ToInputId <--
func (m *TLInputSeqId) ToInputId() *InputId {
	if m == nil {
		return nil
	}

	return &InputId{Clazz: m}
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
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Key        string `json:"key"`
	N          int32  `json:"n"`
}

func MakeTLInputNSeqId(m *TLInputNSeqId) *TLInputNSeqId {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_inputNSeqId

	return m
}

func (m *TLInputNSeqId) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// InputIdClazzName <--
func (m *TLInputNSeqId) InputIdClazzName() string {
	return ClazzName_inputNSeqId
}

// ClazzName <--
func (m *TLInputNSeqId) ClazzName() string {
	return m.ClazzName2
}

// ToInputId <--
func (m *TLInputNSeqId) ToInputId() *InputId {
	if m == nil {
		return nil
	}

	return &InputId{Clazz: m}
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
	Clazz InputIdClazz `json:"_clazz"`
}

func (m *InputId) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *InputId) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.InputIdClazzName()
	}
}

// Encode <--
func (m *InputId) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("InputId - invalid Clazz")
}

// Decode <--
func (m *InputId) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeInputIdClazz(d)
	return
}

// Match <--
func (m *InputId) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
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

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLInputId); ok {
		return x, true
	}

	return nil, false
}

// ToInputIds <--
func (m *InputId) ToInputIds() (*TLInputIds, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLInputIds); ok {
		return x, true
	}

	return nil, false
}

// ToInputSeqId <--
func (m *InputId) ToInputSeqId() (*TLInputSeqId, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLInputSeqId); ok {
		return x, true
	}

	return nil, false
}

// ToInputNSeqId <--
func (m *InputId) ToInputNSeqId() (*TLInputNSeqId, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLInputNSeqId); ok {
		return x, true
	}

	return nil, false
}
