/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package userupdates

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

// UserDifferenceClazz <--
//   - TL_UserDifferenceEmpty
//   - TL_UserDifference
//   - TL_UserDifferenceSlice
//   - TL_UserDifferenceTooLong
type UserDifferenceClazz interface {
	iface.TLObject
	UserDifferenceClazzName() string
}

func DecodeUserDifferenceClazz(d *bin.Decoder) (UserDifferenceClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode UserDifference: constructor: %w", err)
	}

	switch id {
	case 0xb38ac177:
		x := &TLUserDifferenceEmpty{ClazzID: id, ClazzName2: ClazzName_userDifferenceEmpty}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0xb15cb08d:
		x := &TLUserDifference{ClazzID: id, ClazzName2: ClazzName_userDifference}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x4ef1987f:
		x := &TLUserDifferenceSlice{ClazzID: id, ClazzName2: ClazzName_userDifferenceSlice}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x1d095703:
		x := &TLUserDifferenceTooLong{ClazzID: id, ClazzName2: ClazzName_userDifferenceTooLong}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode UserDifference: invalid constructor %x", id)
	}

}

// TLUserDifferenceEmpty <--
type TLUserDifferenceEmpty struct {
	ClazzID    uint32         `json:"_id"`
	ClazzName2 string         `json:"_name"`
	State      UserStateClazz `json:"state"`
}

func MakeTLUserDifferenceEmpty(m *TLUserDifferenceEmpty) *TLUserDifferenceEmpty {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_userDifferenceEmpty

	return m
}

func (m *TLUserDifferenceEmpty) String() string {
	return iface.DebugStringWithName("userDifferenceEmpty", m)
}

func (m *TLUserDifferenceEmpty) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("userDifferenceEmpty", m)
}

// UserDifferenceClazzName <--
func (m *TLUserDifferenceEmpty) UserDifferenceClazzName() string {
	return ClazzName_userDifferenceEmpty
}

// ClazzName <--
func (m *TLUserDifferenceEmpty) ClazzName() string {
	return m.ClazzName2
}

// ToUserDifference <--
func (m *TLUserDifferenceEmpty) ToUserDifference() *UserDifference {
	if m == nil {
		return nil
	}

	return &UserDifference{Clazz: m}

}

func (m *TLUserDifferenceEmpty) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userDifferenceEmpty, int(layer)); clazzId {
	case 0xb38ac177:
		size := 4
		size += iface.CalcObjectSize(m.State, layer)

		return size
	default:
		return 0
	}
}

func (m *TLUserDifferenceEmpty) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userDifferenceEmpty, int(layer)); clazzId {
	case 0xb38ac177:
		if err := iface.ValidateRequiredObject("state", m.State); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode userDifferenceEmpty: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLUserDifferenceEmpty) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userDifferenceEmpty, int(layer)); clazzId {
	case 0xb38ac177:
		x.PutClazzID(0xb38ac177)

		if m.State == nil {
			return fmt.Errorf("unable to encode userDifferenceEmpty#0xb38ac177: field state is nil")
		}
		if err := m.State.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode userDifferenceEmpty#0xb38ac177: field state: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode userDifferenceEmpty: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserDifferenceEmpty) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xb38ac177:

		m.State, err = DecodeUserStateClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode userDifferenceEmpty#0xb38ac177: field state: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode userDifferenceEmpty: invalid constructor %x", m.ClazzID)
	}
}

// TLUserDifference <--
type TLUserDifference struct {
	ClazzID      uint32            `json:"_id"`
	ClazzName2   string            `json:"_name"`
	NewMessages  []tg.MessageClazz `json:"new_messages"`
	OtherUpdates []tg.UpdateClazz  `json:"other_updates"`
	State        UserStateClazz    `json:"state"`
}

func MakeTLUserDifference(m *TLUserDifference) *TLUserDifference {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_userDifference

	return m
}

func (m *TLUserDifference) String() string {
	return iface.DebugStringWithName("userDifference", m)
}

func (m *TLUserDifference) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("userDifference", m)
}

// UserDifferenceClazzName <--
func (m *TLUserDifference) UserDifferenceClazzName() string {
	return ClazzName_userDifference
}

// ClazzName <--
func (m *TLUserDifference) ClazzName() string {
	return m.ClazzName2
}

// ToUserDifference <--
func (m *TLUserDifference) ToUserDifference() *UserDifference {
	if m == nil {
		return nil
	}

	return &UserDifference{Clazz: m}

}

func (m *TLUserDifference) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userDifference, int(layer)); clazzId {
	case 0xb15cb08d:
		size := 4
		size += iface.CalcObjectListSize(m.NewMessages, layer)
		size += iface.CalcObjectListSize(m.OtherUpdates, layer)
		size += iface.CalcObjectSize(m.State, layer)

		return size
	default:
		return 0
	}
}

func (m *TLUserDifference) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userDifference, int(layer)); clazzId {
	case 0xb15cb08d:
		if err := iface.ValidateRequiredSlice("new_messages", m.NewMessages); err != nil {
			return err
		}

		if err := iface.ValidateRequiredSlice("other_updates", m.OtherUpdates); err != nil {
			return err
		}

		if err := iface.ValidateRequiredObject("state", m.State); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode userDifference: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLUserDifference) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userDifference, int(layer)); clazzId {
	case 0xb15cb08d:
		x.PutClazzID(0xb15cb08d)

		if err := iface.EncodeObjectList(x, m.NewMessages, layer); err != nil {
			return fmt.Errorf("unable to encode userDifference#0xb15cb08d: field new_messages: %w", err)
		}

		if err := iface.EncodeObjectList(x, m.OtherUpdates, layer); err != nil {
			return fmt.Errorf("unable to encode userDifference#0xb15cb08d: field other_updates: %w", err)
		}

		if m.State == nil {
			return fmt.Errorf("unable to encode userDifference#0xb15cb08d: field state is nil")
		}
		if err := m.State.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode userDifference#0xb15cb08d: field state: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode userDifference: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserDifference) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xb15cb08d:
		l1, err3 := d.VectorHeader()
		if err3 != nil {
			return fmt.Errorf("unable to decode userDifference#0xb15cb08d: field new_messages: %w", err3)
		}
		if l1 > bin.MaxVectorLen {
			return fmt.Errorf("unable to decode userDifference#0xb15cb08d: field new_messages: %w", &bin.InvalidLengthError{Type: "vector", Length: int(l1)})
		}
		prealloc1 := int(l1)
		if prealloc1 > bin.PreallocateLimit {
			prealloc1 = bin.PreallocateLimit
		}
		v1 := make([]tg.MessageClazz, 0, prealloc1)
		for i := int32(0); i < l1; i++ {
			vv1, err3 := tg.DecodeMessageClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode userDifference#0xb15cb08d: field new_messages: %w", err3)
			}
			v1 = append(v1, vv1)
		}
		m.NewMessages = v1

		l2, err3 := d.VectorHeader()
		if err3 != nil {
			return fmt.Errorf("unable to decode userDifference#0xb15cb08d: field other_updates: %w", err3)
		}
		if l2 > bin.MaxVectorLen {
			return fmt.Errorf("unable to decode userDifference#0xb15cb08d: field other_updates: %w", &bin.InvalidLengthError{Type: "vector", Length: int(l2)})
		}
		prealloc2 := int(l2)
		if prealloc2 > bin.PreallocateLimit {
			prealloc2 = bin.PreallocateLimit
		}
		v2 := make([]tg.UpdateClazz, 0, prealloc2)
		for i := int32(0); i < l2; i++ {
			vv2, err3 := tg.DecodeUpdateClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode userDifference#0xb15cb08d: field other_updates: %w", err3)
			}
			v2 = append(v2, vv2)
		}
		m.OtherUpdates = v2

		m.State, err = DecodeUserStateClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode userDifference#0xb15cb08d: field state: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode userDifference: invalid constructor %x", m.ClazzID)
	}
}

// TLUserDifferenceSlice <--
type TLUserDifferenceSlice struct {
	ClazzID           uint32            `json:"_id"`
	ClazzName2        string            `json:"_name"`
	NewMessages       []tg.MessageClazz `json:"new_messages"`
	OtherUpdates      []tg.UpdateClazz  `json:"other_updates"`
	IntermediateState UserStateClazz    `json:"intermediate_state"`
}

func MakeTLUserDifferenceSlice(m *TLUserDifferenceSlice) *TLUserDifferenceSlice {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_userDifferenceSlice

	return m
}

func (m *TLUserDifferenceSlice) String() string {
	return iface.DebugStringWithName("userDifferenceSlice", m)
}

func (m *TLUserDifferenceSlice) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("userDifferenceSlice", m)
}

// UserDifferenceClazzName <--
func (m *TLUserDifferenceSlice) UserDifferenceClazzName() string {
	return ClazzName_userDifferenceSlice
}

// ClazzName <--
func (m *TLUserDifferenceSlice) ClazzName() string {
	return m.ClazzName2
}

// ToUserDifference <--
func (m *TLUserDifferenceSlice) ToUserDifference() *UserDifference {
	if m == nil {
		return nil
	}

	return &UserDifference{Clazz: m}

}

func (m *TLUserDifferenceSlice) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userDifferenceSlice, int(layer)); clazzId {
	case 0x4ef1987f:
		size := 4
		size += iface.CalcObjectListSize(m.NewMessages, layer)
		size += iface.CalcObjectListSize(m.OtherUpdates, layer)
		size += iface.CalcObjectSize(m.IntermediateState, layer)

		return size
	default:
		return 0
	}
}

func (m *TLUserDifferenceSlice) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userDifferenceSlice, int(layer)); clazzId {
	case 0x4ef1987f:
		if err := iface.ValidateRequiredSlice("new_messages", m.NewMessages); err != nil {
			return err
		}

		if err := iface.ValidateRequiredSlice("other_updates", m.OtherUpdates); err != nil {
			return err
		}

		if err := iface.ValidateRequiredObject("intermediate_state", m.IntermediateState); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode userDifferenceSlice: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLUserDifferenceSlice) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userDifferenceSlice, int(layer)); clazzId {
	case 0x4ef1987f:
		x.PutClazzID(0x4ef1987f)

		if err := iface.EncodeObjectList(x, m.NewMessages, layer); err != nil {
			return fmt.Errorf("unable to encode userDifferenceSlice#0x4ef1987f: field new_messages: %w", err)
		}

		if err := iface.EncodeObjectList(x, m.OtherUpdates, layer); err != nil {
			return fmt.Errorf("unable to encode userDifferenceSlice#0x4ef1987f: field other_updates: %w", err)
		}

		if m.IntermediateState == nil {
			return fmt.Errorf("unable to encode userDifferenceSlice#0x4ef1987f: field intermediate_state is nil")
		}
		if err := m.IntermediateState.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode userDifferenceSlice#0x4ef1987f: field intermediate_state: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode userDifferenceSlice: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserDifferenceSlice) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x4ef1987f:
		l1, err3 := d.VectorHeader()
		if err3 != nil {
			return fmt.Errorf("unable to decode userDifferenceSlice#0x4ef1987f: field new_messages: %w", err3)
		}
		if l1 > bin.MaxVectorLen {
			return fmt.Errorf("unable to decode userDifferenceSlice#0x4ef1987f: field new_messages: %w", &bin.InvalidLengthError{Type: "vector", Length: int(l1)})
		}
		prealloc1 := int(l1)
		if prealloc1 > bin.PreallocateLimit {
			prealloc1 = bin.PreallocateLimit
		}
		v1 := make([]tg.MessageClazz, 0, prealloc1)
		for i := int32(0); i < l1; i++ {
			vv1, err3 := tg.DecodeMessageClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode userDifferenceSlice#0x4ef1987f: field new_messages: %w", err3)
			}
			v1 = append(v1, vv1)
		}
		m.NewMessages = v1

		l2, err3 := d.VectorHeader()
		if err3 != nil {
			return fmt.Errorf("unable to decode userDifferenceSlice#0x4ef1987f: field other_updates: %w", err3)
		}
		if l2 > bin.MaxVectorLen {
			return fmt.Errorf("unable to decode userDifferenceSlice#0x4ef1987f: field other_updates: %w", &bin.InvalidLengthError{Type: "vector", Length: int(l2)})
		}
		prealloc2 := int(l2)
		if prealloc2 > bin.PreallocateLimit {
			prealloc2 = bin.PreallocateLimit
		}
		v2 := make([]tg.UpdateClazz, 0, prealloc2)
		for i := int32(0); i < l2; i++ {
			vv2, err3 := tg.DecodeUpdateClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode userDifferenceSlice#0x4ef1987f: field other_updates: %w", err3)
			}
			v2 = append(v2, vv2)
		}
		m.OtherUpdates = v2

		m.IntermediateState, err = DecodeUserStateClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode userDifferenceSlice#0x4ef1987f: field intermediate_state: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode userDifferenceSlice: invalid constructor %x", m.ClazzID)
	}
}

// TLUserDifferenceTooLong <--
type TLUserDifferenceTooLong struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Pts        int64  `json:"pts"`
}

func MakeTLUserDifferenceTooLong(m *TLUserDifferenceTooLong) *TLUserDifferenceTooLong {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_userDifferenceTooLong

	return m
}

func (m *TLUserDifferenceTooLong) String() string {
	return iface.DebugStringWithName("userDifferenceTooLong", m)
}

func (m *TLUserDifferenceTooLong) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("userDifferenceTooLong", m)
}

// UserDifferenceClazzName <--
func (m *TLUserDifferenceTooLong) UserDifferenceClazzName() string {
	return ClazzName_userDifferenceTooLong
}

// ClazzName <--
func (m *TLUserDifferenceTooLong) ClazzName() string {
	return m.ClazzName2
}

// ToUserDifference <--
func (m *TLUserDifferenceTooLong) ToUserDifference() *UserDifference {
	if m == nil {
		return nil
	}

	return &UserDifference{Clazz: m}

}

func (m *TLUserDifferenceTooLong) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userDifferenceTooLong, int(layer)); clazzId {
	case 0x1d095703:
		size := 4
		size += 8

		return size
	default:
		return 0
	}
}

func (m *TLUserDifferenceTooLong) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userDifferenceTooLong, int(layer)); clazzId {
	case 0x1d095703:

		return nil
	default:
		return fmt.Errorf("unable to encode userDifferenceTooLong: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLUserDifferenceTooLong) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userDifferenceTooLong, int(layer)); clazzId {
	case 0x1d095703:
		x.PutClazzID(0x1d095703)

		x.PutInt64(m.Pts)

		return nil
	default:
		return fmt.Errorf("unable to encode userDifferenceTooLong: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserDifferenceTooLong) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x1d095703:
		m.Pts, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userDifferenceTooLong#0x1d095703: field pts: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode userDifferenceTooLong: invalid constructor %x", m.ClazzID)
	}
}

// UserDifference <--
type UserDifference struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz UserDifferenceClazz `json:"_clazz"`
}

func (m *UserDifference) String() string {
	return iface.DebugStringWithName(m.ClazzName(), m)
}

func (m *UserDifference) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName(m.ClazzName(), m)
}

func (m *UserDifference) CalcSize(layer int32) int {
	if m == nil || m.Clazz == nil {
		return 0
	}
	return iface.CalcObjectSize(m.Clazz, layer)
}

func (m *UserDifference) Validate(layer int32) error {
	if m == nil {
		return fmt.Errorf("UserDifference is required")
	}
	if m.Clazz == nil {
		return fmt.Errorf("UserDifference.Clazz is required")
	}
	if v, ok := m.Clazz.(iface.TLObjectValidator); ok {
		if err := v.Validate(layer); err != nil {
			return fmt.Errorf("unable to validate UserDifference.Clazz: %w", err)
		}
	}
	return nil
}

func (m *UserDifference) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.UserDifferenceClazzName()
	}
}

// Encode <--
func (m *UserDifference) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		if err := m.Clazz.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode UserDifference.Clazz: %w", err)
		}
		return nil
	}

	return fmt.Errorf("UserDifference - invalid Clazz")
}

// Decode <--
func (m *UserDifference) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeUserDifferenceClazz(d)
	return
}

// ToUserDifferenceEmpty <--
func (m *UserDifference) ToUserDifferenceEmpty() (*TLUserDifferenceEmpty, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLUserDifferenceEmpty); ok {
		return x, true
	}

	return nil, false
}

// ToUserDifference <--
func (m *UserDifference) ToUserDifference() (*TLUserDifference, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLUserDifference); ok {
		return x, true
	}

	return nil, false
}

// ToUserDifferenceSlice <--
func (m *UserDifference) ToUserDifferenceSlice() (*TLUserDifferenceSlice, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLUserDifferenceSlice); ok {
		return x, true
	}

	return nil, false
}

// ToUserDifferenceTooLong <--
func (m *UserDifference) ToUserDifferenceTooLong() (*TLUserDifferenceTooLong, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLUserDifferenceTooLong); ok {
		return x, true
	}

	return nil, false
}

// UserOperationClazz <--
//   - TL_UserOperation
type UserOperationClazz = *TLUserOperation

func DecodeUserOperationClazz(d *bin.Decoder) (UserOperationClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode UserOperation: constructor: %w", err)
	}

	switch id {
	case 0x2d4e84d7:
		x := &TLUserOperation{ClazzID: id, ClazzName2: ClazzName_userOperation}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode UserOperation: invalid constructor %x", id)
	}

}

// TLUserOperation <--
type TLUserOperation struct {
	ClazzID              uint32  `json:"_id"`
	ClazzName2           string  `json:"_name"`
	UserId               int64   `json:"user_id"`
	BucketId             int32   `json:"bucket_id"`
	PartitionId          int32   `json:"partition_id"`
	OperationId          string  `json:"operation_id"`
	OpType               int32   `json:"op_type"`
	OpSource             int32   `json:"op_source"`
	ActorUserId          int64   `json:"actor_user_id"`
	AuthKeyId            *int64  `json:"auth_key_id"`
	AuthKeyIdExclude     *int64  `json:"auth_key_id_exclude"`
	PeerType             int32   `json:"peer_type"`
	PeerId               int64   `json:"peer_id"`
	CanonicalMessageId   *int64  `json:"canonical_message_id"`
	CanonicalPeerSeq     *int64  `json:"canonical_peer_seq"`
	CanonicalDate        *int64  `json:"canonical_date"`
	DependencyPts        *int64  `json:"dependency_pts"`
	RequestId            *string `json:"request_id"`
	PayloadSchemaVersion int32   `json:"payload_schema_version"`
	PayloadCodec         int32   `json:"payload_codec"`
	PayloadHash          []byte  `json:"payload_hash"`
	Payload              []byte  `json:"payload"`
}

func MakeTLUserOperation(m *TLUserOperation) *TLUserOperation {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_userOperation

	return m
}

func (m *TLUserOperation) String() string {
	return iface.DebugStringWithName("userOperation", m)
}

func (m *TLUserOperation) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("userOperation", m)
}

// UserOperationClazzName <--
func (m *TLUserOperation) UserOperationClazzName() string {
	return ClazzName_userOperation
}

// ClazzName <--
func (m *TLUserOperation) ClazzName() string {
	return m.ClazzName2
}

// ToUserOperation <--
func (m *TLUserOperation) ToUserOperation() *UserOperation {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLUserOperation) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userOperation, int(layer)); clazzId {
	case 0x2d4e84d7:
		size := 4
		size += 4
		size += 8
		size += 4
		size += 4
		size += iface.CalcStringSize(m.OperationId)
		size += 4
		size += 4
		size += 8
		if m.AuthKeyId != nil {
			size += 8
		}

		if m.AuthKeyIdExclude != nil {
			size += 8
		}

		size += 4
		size += 8
		if m.CanonicalMessageId != nil {
			size += 8
		}

		if m.CanonicalPeerSeq != nil {
			size += 8
		}

		if m.CanonicalDate != nil {
			size += 8
		}

		if m.DependencyPts != nil {
			size += 8
		}

		if m.RequestId != nil {
			size += iface.CalcStringSize(*m.RequestId)
		}

		size += 4
		size += 4
		size += iface.CalcBytesSize(m.PayloadHash)
		size += iface.CalcBytesSize(m.Payload)

		return size
	default:
		return 0
	}
}

func (m *TLUserOperation) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userOperation, int(layer)); clazzId {
	case 0x2d4e84d7:
		if err := iface.ValidateRequiredString("operation_id", m.OperationId); err != nil {
			return err
		}

		if err := iface.ValidateRequiredBytes("payload_hash", m.PayloadHash); err != nil {
			return err
		}

		if err := iface.ValidateRequiredBytes("payload", m.Payload); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode userOperation: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLUserOperation) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userOperation, int(layer)); clazzId {
	case 0x2d4e84d7:
		x.PutClazzID(0x2d4e84d7)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.AuthKeyId != nil {
				flags |= 1 << 0
			}
			if m.AuthKeyIdExclude != nil {
				flags |= 1 << 1
			}

			if m.CanonicalMessageId != nil {
				flags |= 1 << 2
			}
			if m.CanonicalPeerSeq != nil {
				flags |= 1 << 3
			}
			if m.CanonicalDate != nil {
				flags |= 1 << 4
			}
			if m.DependencyPts != nil {
				flags |= 1 << 5
			}
			if m.RequestId != nil {
				flags |= 1 << 6
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.UserId)
		x.PutInt32(m.BucketId)
		x.PutInt32(m.PartitionId)
		x.PutString(m.OperationId)
		x.PutInt32(m.OpType)
		x.PutInt32(m.OpSource)
		x.PutInt64(m.ActorUserId)
		if m.AuthKeyId != nil {
			x.PutInt64(*m.AuthKeyId)
		}

		if m.AuthKeyIdExclude != nil {
			x.PutInt64(*m.AuthKeyIdExclude)
		}

		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		if m.CanonicalMessageId != nil {
			x.PutInt64(*m.CanonicalMessageId)
		}

		if m.CanonicalPeerSeq != nil {
			x.PutInt64(*m.CanonicalPeerSeq)
		}

		if m.CanonicalDate != nil {
			x.PutInt64(*m.CanonicalDate)
		}

		if m.DependencyPts != nil {
			x.PutInt64(*m.DependencyPts)
		}

		if m.RequestId != nil {
			x.PutString(*m.RequestId)
		}

		x.PutInt32(m.PayloadSchemaVersion)
		x.PutInt32(m.PayloadCodec)
		x.PutBytes(m.PayloadHash)
		x.PutBytes(m.Payload)

		return nil
	default:
		return fmt.Errorf("unable to encode userOperation: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserOperation) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x2d4e84d7:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field flags: %w", err)
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field user_id: %w", err)
		}
		m.BucketId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field bucket_id: %w", err)
		}
		m.PartitionId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field partition_id: %w", err)
		}
		m.OperationId, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field operation_id: %w", err)
		}
		m.OpType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field op_type: %w", err)
		}
		m.OpSource, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field op_source: %w", err)
		}
		m.ActorUserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field actor_user_id: %w", err)
		}
		if (flags & (1 << 0)) != 0 {
			m.AuthKeyId = new(int64)
			*m.AuthKeyId, err = d.Int64()
			if err != nil {
				return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field auth_key_id: %w", err)
			}
		}

		if (flags & (1 << 1)) != 0 {
			m.AuthKeyIdExclude = new(int64)
			*m.AuthKeyIdExclude, err = d.Int64()
			if err != nil {
				return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field auth_key_id_exclude: %w", err)
			}
		}

		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field peer_id: %w", err)
		}
		if (flags & (1 << 2)) != 0 {
			m.CanonicalMessageId = new(int64)
			*m.CanonicalMessageId, err = d.Int64()
			if err != nil {
				return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field canonical_message_id: %w", err)
			}
		}

		if (flags & (1 << 3)) != 0 {
			m.CanonicalPeerSeq = new(int64)
			*m.CanonicalPeerSeq, err = d.Int64()
			if err != nil {
				return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field canonical_peer_seq: %w", err)
			}
		}

		if (flags & (1 << 4)) != 0 {
			m.CanonicalDate = new(int64)
			*m.CanonicalDate, err = d.Int64()
			if err != nil {
				return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field canonical_date: %w", err)
			}
		}

		if (flags & (1 << 5)) != 0 {
			m.DependencyPts = new(int64)
			*m.DependencyPts, err = d.Int64()
			if err != nil {
				return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field dependency_pts: %w", err)
			}
		}

		if (flags & (1 << 6)) != 0 {
			m.RequestId = new(string)
			*m.RequestId, err = d.String()
			if err != nil {
				return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field request_id: %w", err)
			}
		}

		m.PayloadSchemaVersion, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field payload_schema_version: %w", err)
		}
		m.PayloadCodec, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field payload_codec: %w", err)
		}
		m.PayloadHash, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field payload_hash: %w", err)
		}
		m.Payload, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode userOperation#0x2d4e84d7: field payload: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode userOperation: invalid constructor %x", m.ClazzID)
	}
}

// UserOperation <--
type UserOperation = TLUserOperation

// UserOperationResultClazz <--
//   - TL_UserOperationResult
type UserOperationResultClazz = *TLUserOperationResult

func DecodeUserOperationResultClazz(d *bin.Decoder) (UserOperationResultClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode UserOperationResult: constructor: %w", err)
	}

	switch id {
	case 0x7311db72:
		x := &TLUserOperationResult{ClazzID: id, ClazzName2: ClazzName_userOperationResult}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode UserOperationResult: invalid constructor %x", id)
	}

}

// TLUserOperationResult <--
type TLUserOperationResult struct {
	ClazzID               uint32 `json:"_id"`
	ClazzName2            string `json:"_name"`
	UserId                int64  `json:"user_id"`
	OperationId           string `json:"operation_id"`
	Status                int32  `json:"status"`
	Pts                   int64  `json:"pts"`
	PtsCount              int32  `json:"pts_count"`
	CurrentPts            int64  `json:"current_pts"`
	ResponseSchemaVersion *int32 `json:"response_schema_version"`
	ResponsePayload       []byte `json:"response_payload"`
	ResponsePayloadHash   []byte `json:"response_payload_hash"`
}

func MakeTLUserOperationResult(m *TLUserOperationResult) *TLUserOperationResult {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_userOperationResult

	return m
}

func (m *TLUserOperationResult) String() string {
	return iface.DebugStringWithName("userOperationResult", m)
}

func (m *TLUserOperationResult) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("userOperationResult", m)
}

// UserOperationResultClazzName <--
func (m *TLUserOperationResult) UserOperationResultClazzName() string {
	return ClazzName_userOperationResult
}

// ClazzName <--
func (m *TLUserOperationResult) ClazzName() string {
	return m.ClazzName2
}

// ToUserOperationResult <--
func (m *TLUserOperationResult) ToUserOperationResult() *UserOperationResult {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLUserOperationResult) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userOperationResult, int(layer)); clazzId {
	case 0x7311db72:
		size := 4
		size += 4
		size += 8
		size += iface.CalcStringSize(m.OperationId)
		size += 4
		size += 8
		size += 4
		size += 8
		if m.ResponseSchemaVersion != nil {
			size += 4
		}

		if m.ResponsePayload != nil {
			size += iface.CalcBytesSize(m.ResponsePayload)
		}

		if m.ResponsePayloadHash != nil {
			size += iface.CalcBytesSize(m.ResponsePayloadHash)
		}

		return size
	default:
		return 0
	}
}

func (m *TLUserOperationResult) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userOperationResult, int(layer)); clazzId {
	case 0x7311db72:
		if err := iface.ValidateRequiredString("operation_id", m.OperationId); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode userOperationResult: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLUserOperationResult) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userOperationResult, int(layer)); clazzId {
	case 0x7311db72:
		x.PutClazzID(0x7311db72)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.ResponseSchemaVersion != nil {
				flags |= 1 << 0
			}
			if m.ResponsePayload != nil {
				flags |= 1 << 1
			}
			if m.ResponsePayloadHash != nil {
				flags |= 1 << 2
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.UserId)
		x.PutString(m.OperationId)
		x.PutInt32(m.Status)
		x.PutInt64(m.Pts)
		x.PutInt32(m.PtsCount)
		x.PutInt64(m.CurrentPts)
		if m.ResponseSchemaVersion != nil {
			x.PutInt32(*m.ResponseSchemaVersion)
		}

		if m.ResponsePayload != nil {
			x.PutBytes(m.ResponsePayload)
		}

		if m.ResponsePayloadHash != nil {
			x.PutBytes(m.ResponsePayloadHash)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode userOperationResult: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserOperationResult) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x7311db72:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode userOperationResult#0x7311db72: field flags: %w", err)
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userOperationResult#0x7311db72: field user_id: %w", err)
		}
		m.OperationId, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode userOperationResult#0x7311db72: field operation_id: %w", err)
		}
		m.Status, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userOperationResult#0x7311db72: field status: %w", err)
		}
		m.Pts, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userOperationResult#0x7311db72: field pts: %w", err)
		}
		m.PtsCount, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userOperationResult#0x7311db72: field pts_count: %w", err)
		}
		m.CurrentPts, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userOperationResult#0x7311db72: field current_pts: %w", err)
		}
		if (flags & (1 << 0)) != 0 {
			m.ResponseSchemaVersion = new(int32)
			*m.ResponseSchemaVersion, err = d.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode userOperationResult#0x7311db72: field response_schema_version: %w", err)
			}
		}
		if (flags & (1 << 1)) != 0 {
			m.ResponsePayload, err = d.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode userOperationResult#0x7311db72: field response_payload: %w", err)
			}
		}

		if (flags & (1 << 2)) != 0 {
			m.ResponsePayloadHash, err = d.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode userOperationResult#0x7311db72: field response_payload_hash: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to decode userOperationResult: invalid constructor %x", m.ClazzID)
	}
}

// UserOperationResult <--
type UserOperationResult = TLUserOperationResult

// UserStateClazz <--
//   - TL_UserState
type UserStateClazz = *TLUserState

func DecodeUserStateClazz(d *bin.Decoder) (UserStateClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode UserState: constructor: %w", err)
	}

	switch id {
	case 0x635f3815:
		x := &TLUserState{ClazzID: id, ClazzName2: ClazzName_userState}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode UserState: invalid constructor %x", id)
	}

}

// TLUserState <--
type TLUserState struct {
	ClazzID     uint32 `json:"_id"`
	ClazzName2  string `json:"_name"`
	Pts         int64  `json:"pts"`
	Qts         int32  `json:"qts"`
	Date        int32  `json:"date"`
	Seq         int32  `json:"seq"`
	UnreadCount int32  `json:"unread_count"`
}

func MakeTLUserState(m *TLUserState) *TLUserState {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_userState

	return m
}

func (m *TLUserState) String() string {
	return iface.DebugStringWithName("userState", m)
}

func (m *TLUserState) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("userState", m)
}

// UserStateClazzName <--
func (m *TLUserState) UserStateClazzName() string {
	return ClazzName_userState
}

// ClazzName <--
func (m *TLUserState) ClazzName() string {
	return m.ClazzName2
}

// ToUserState <--
func (m *TLUserState) ToUserState() *UserState {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLUserState) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userState, int(layer)); clazzId {
	case 0x635f3815:
		size := 4
		size += 8
		size += 4
		size += 4
		size += 4
		size += 4

		return size
	default:
		return 0
	}
}

func (m *TLUserState) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userState, int(layer)); clazzId {
	case 0x635f3815:

		return nil
	default:
		return fmt.Errorf("unable to encode userState: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLUserState) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userState, int(layer)); clazzId {
	case 0x635f3815:
		x.PutClazzID(0x635f3815)

		x.PutInt64(m.Pts)
		x.PutInt32(m.Qts)
		x.PutInt32(m.Date)
		x.PutInt32(m.Seq)
		x.PutInt32(m.UnreadCount)

		return nil
	default:
		return fmt.Errorf("unable to encode userState: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserState) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x635f3815:
		m.Pts, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userState#0x635f3815: field pts: %w", err)
		}
		m.Qts, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userState#0x635f3815: field qts: %w", err)
		}
		m.Date, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userState#0x635f3815: field date: %w", err)
		}
		m.Seq, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userState#0x635f3815: field seq: %w", err)
		}
		m.UnreadCount, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userState#0x635f3815: field unread_count: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode userState: invalid constructor %x", m.ClazzID)
	}
}

// UserState <--
type UserState = TLUserState
