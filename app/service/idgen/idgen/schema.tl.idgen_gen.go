/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package idgen

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

	switch id {
	case 0xc07844cb:
		x := &TLIdVal{ClazzID: id, ClazzName2: ClazzName_idVal}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x1c3baa66:
		x := &TLIdVals{ClazzID: id, ClazzName2: ClazzName_idVals}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x2a047d08:
		x := &TLSeqIdVal{ClazzID: id, ClazzName2: ClazzName_seqIdVal}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeIdVal - unexpected clazzId: %d", id)
	}

}

// TLIdVal <--
type TLIdVal struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Id         int64  `json:"id"`
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

func (m *TLIdVal) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("idVal", m)
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

func (m *TLIdVal) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_idVal, int(layer)); clazzId {
	case 0xc07844cb:
		size := 4
		size += 8

		return size
	default:
		return 0
	}
}

func (m *TLIdVal) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_idVal, int(layer)); clazzId {
	case 0xc07844cb:

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idVal, layer)
	}
}

// Encode <--
func (m *TLIdVal) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_idVal, int(layer)); clazzId {
	case 0xc07844cb:
		x.PutClazzID(0xc07844cb)

		x.PutInt64(m.Id)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idVal, layer)
	}
}

// Decode <--
func (m *TLIdVal) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xc07844cb:
		m.Id, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLIdVals <--
type TLIdVals struct {
	ClazzID    uint32  `json:"_id"`
	ClazzName2 string  `json:"_name"`
	Id         []int64 `json:"id"`
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

func (m *TLIdVals) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("idVals", m)
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

func (m *TLIdVals) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_idVals, int(layer)); clazzId {
	case 0x1c3baa66:
		size := 4
		size += iface.CalcInt64ListSize(m.Id)

		return size
	default:
		return 0
	}
}

func (m *TLIdVals) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_idVals, int(layer)); clazzId {
	case 0x1c3baa66:
		if err := iface.ValidateRequiredSlice("id", m.Id); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idVals, layer)
	}
}

// Encode <--
func (m *TLIdVals) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_idVals, int(layer)); clazzId {
	case 0x1c3baa66:
		x.PutClazzID(0x1c3baa66)

		iface.EncodeInt64List(x, m.Id)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idVals, layer)
	}
}

// Decode <--
func (m *TLIdVals) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x1c3baa66:

		m.Id, err = iface.DecodeInt64List(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLSeqIdVal <--
type TLSeqIdVal struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Id         int64  `json:"id"`
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

func (m *TLSeqIdVal) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("seqIdVal", m)
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

func (m *TLSeqIdVal) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_seqIdVal, int(layer)); clazzId {
	case 0x2a047d08:
		size := 4
		size += 8

		return size
	default:
		return 0
	}
}

func (m *TLSeqIdVal) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_seqIdVal, int(layer)); clazzId {
	case 0x2a047d08:

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_seqIdVal, layer)
	}
}

// Encode <--
func (m *TLSeqIdVal) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_seqIdVal, int(layer)); clazzId {
	case 0x2a047d08:
		x.PutClazzID(0x2a047d08)

		x.PutInt64(m.Id)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_seqIdVal, layer)
	}
}

// Decode <--
func (m *TLSeqIdVal) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x2a047d08:
		m.Id, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
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

func (m *IdVal) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName(m.ClazzName(), m)
}

func (m *IdVal) CalcSize(layer int32) int {
	if m == nil || m.Clazz == nil {
		return 0
	}
	return iface.CalcObjectSize(m.Clazz, layer)
}

func (m *IdVal) Validate(layer int32) error {
	if m == nil {
		return fmt.Errorf("IdVal is required")
	}
	if m.Clazz == nil {
		return fmt.Errorf("IdVal.Clazz is required")
	}
	if v, ok := m.Clazz.(iface.TLObjectValidator); ok {
		return v.Validate(layer)
	}
	return nil
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

	switch id {
	case 0x8af2196c:
		x := &TLInputId{ClazzID: id, ClazzName2: ClazzName_inputId}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x7f285fbc:
		x := &TLInputIds{ClazzID: id, ClazzName2: ClazzName_inputIds}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0xcd52bbcd:
		x := &TLInputSeqId{ClazzID: id, ClazzName2: ClazzName_inputSeqId}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x7ab16d81:
		x := &TLInputNSeqId{ClazzID: id, ClazzName2: ClazzName_inputNSeqId}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
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

func (m *TLInputId) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("inputId", m)
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

func (m *TLInputId) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inputId, int(layer)); clazzId {
	case 0x8af2196c:
		size := 4

		return size
	default:
		return 0
	}
}

func (m *TLInputId) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inputId, int(layer)); clazzId {
	case 0x8af2196c:

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inputId, layer)
	}
}

// Encode <--
func (m *TLInputId) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inputId, int(layer)); clazzId {
	case 0x8af2196c:
		x.PutClazzID(0x8af2196c)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inputId, layer)
	}
}

// Decode <--
func (m *TLInputId) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x8af2196c:

		return nil
	default:
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

func (m *TLInputIds) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("inputIds", m)
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

func (m *TLInputIds) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inputIds, int(layer)); clazzId {
	case 0x7f285fbc:
		size := 4
		size += 4

		return size
	default:
		return 0
	}
}

func (m *TLInputIds) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inputIds, int(layer)); clazzId {
	case 0x7f285fbc:

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inputIds, layer)
	}
}

// Encode <--
func (m *TLInputIds) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inputIds, int(layer)); clazzId {
	case 0x7f285fbc:
		x.PutClazzID(0x7f285fbc)

		x.PutInt32(m.Num)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inputIds, layer)
	}
}

// Decode <--
func (m *TLInputIds) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x7f285fbc:
		m.Num, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
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

func (m *TLInputSeqId) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("inputSeqId", m)
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

func (m *TLInputSeqId) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inputSeqId, int(layer)); clazzId {
	case 0xcd52bbcd:
		size := 4
		size += iface.CalcStringSize(m.Key)

		return size
	default:
		return 0
	}
}

func (m *TLInputSeqId) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inputSeqId, int(layer)); clazzId {
	case 0xcd52bbcd:
		if err := iface.ValidateRequiredString("key", m.Key); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inputSeqId, layer)
	}
}

// Encode <--
func (m *TLInputSeqId) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inputSeqId, int(layer)); clazzId {
	case 0xcd52bbcd:
		x.PutClazzID(0xcd52bbcd)

		x.PutString(m.Key)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inputSeqId, layer)
	}
}

// Decode <--
func (m *TLInputSeqId) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xcd52bbcd:
		m.Key, err = d.String()
		if err != nil {
			return err
		}

		return nil
	default:
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

func (m *TLInputNSeqId) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("inputNSeqId", m)
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

func (m *TLInputNSeqId) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inputNSeqId, int(layer)); clazzId {
	case 0x7ab16d81:
		size := 4
		size += iface.CalcStringSize(m.Key)
		size += 4

		return size
	default:
		return 0
	}
}

func (m *TLInputNSeqId) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inputNSeqId, int(layer)); clazzId {
	case 0x7ab16d81:
		if err := iface.ValidateRequiredString("key", m.Key); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inputNSeqId, layer)
	}
}

// Encode <--
func (m *TLInputNSeqId) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inputNSeqId, int(layer)); clazzId {
	case 0x7ab16d81:
		x.PutClazzID(0x7ab16d81)

		x.PutString(m.Key)
		x.PutInt32(m.N)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inputNSeqId, layer)
	}
}

// Decode <--
func (m *TLInputNSeqId) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x7ab16d81:
		m.Key, err = d.String()
		if err != nil {
			return err
		}
		m.N, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
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

func (m *InputId) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName(m.ClazzName(), m)
}

func (m *InputId) CalcSize(layer int32) int {
	if m == nil || m.Clazz == nil {
		return 0
	}
	return iface.CalcObjectSize(m.Clazz, layer)
}

func (m *InputId) Validate(layer int32) error {
	if m == nil {
		return fmt.Errorf("InputId is required")
	}
	if m.Clazz == nil {
		return fmt.Errorf("InputId.Clazz is required")
	}
	if v, ok := m.Clazz.(iface.TLObjectValidator); ok {
		return v.Validate(layer)
	}
	return nil
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
