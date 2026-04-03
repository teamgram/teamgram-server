/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package updates

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

// ChannelDifferenceClazz <--
//   - TL_ChannelDifference
type ChannelDifferenceClazz = *TLChannelDifference

func DecodeChannelDifferenceClazz(d *bin.Decoder) (ChannelDifferenceClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0xcd19034a:
		x := &TLChannelDifference{ClazzID: id, ClazzName2: ClazzName_channelDifference}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeChannelDifference - unexpected clazzId: %d", id)
	}

}

// TLChannelDifference <--
type TLChannelDifference struct {
	ClazzID      uint32            `json:"_id"`
	ClazzName2   string            `json:"_name"`
	Final        bool              `json:"final"`
	Pts          int32             `json:"pts"`
	NewMessages  []tg.MessageClazz `json:"new_messages"`
	OtherUpdates []tg.UpdateClazz  `json:"other_updates"`
}

func MakeTLChannelDifference(m *TLChannelDifference) *TLChannelDifference {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_channelDifference

	return m
}

func (m *TLChannelDifference) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLChannelDifference) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("channelDifference", m)
}

// ChannelDifferenceClazzName <--
func (m *TLChannelDifference) ChannelDifferenceClazzName() string {
	return ClazzName_channelDifference
}

// ClazzName <--
func (m *TLChannelDifference) ClazzName() string {
	return m.ClazzName2
}

// ToChannelDifference <--
func (m *TLChannelDifference) ToChannelDifference() *ChannelDifference {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLChannelDifference) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_channelDifference, int(layer)); clazzId {
	case 0xcd19034a:
		size := 4
		size += 4
		size += 4
		size += iface.CalcObjectListSize(m.NewMessages, layer)
		size += iface.CalcObjectListSize(m.OtherUpdates, layer)

		return size
	default:
		return 0
	}
}

func (m *TLChannelDifference) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_channelDifference, int(layer)); clazzId {
	case 0xcd19034a:
		if err := iface.ValidateRequiredSlice("new_messages", m.NewMessages); err != nil {
			return err
		}

		if err := iface.ValidateRequiredSlice("other_updates", m.OtherUpdates); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_channelDifference, layer)
	}
}

// Encode <--
func (m *TLChannelDifference) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_channelDifference, int(layer)); clazzId {
	case 0xcd19034a:
		x.PutClazzID(0xcd19034a)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.Final == true {
				flags |= 1 << 0
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt32(m.Pts)

		if err := iface.EncodeObjectList(x, m.NewMessages, layer); err != nil {
			return err
		}

		if err := iface.EncodeObjectList(x, m.OtherUpdates, layer); err != nil {
			return err
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_channelDifference, layer)
	}
}

// Decode <--
func (m *TLChannelDifference) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xcd19034a:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		if (flags & (1 << 0)) != 0 {
			m.Final = true
		}
		m.Pts, err = d.Int32()
		if err != nil {
			return err
		}
		c3, err2 := d.ClazzID()
		if err2 != nil {
			return err2
		}
		if c3 != iface.ClazzID_vector {
			return fmt.Errorf("invalid ClazzID_vector, c%d: %d", 3, c3)
		}
		l3, err3 := d.Int()
		if err3 != nil {
			return err3
		}
		v3 := make([]tg.MessageClazz, l3)
		for i := 0; i < l3; i++ {
			v3[i], err3 = tg.DecodeMessageClazz(d)
			if err3 != nil {
				return err3
			}
		}
		m.NewMessages = v3

		c4, err2 := d.ClazzID()
		if err2 != nil {
			return err2
		}
		if c4 != iface.ClazzID_vector {
			return fmt.Errorf("invalid ClazzID_vector, c%d: %d", 4, c4)
		}
		l4, err3 := d.Int()
		if err3 != nil {
			return err3
		}
		v4 := make([]tg.UpdateClazz, l4)
		for i := 0; i < l4; i++ {
			v4[i], err3 = tg.DecodeUpdateClazz(d)
			if err3 != nil {
				return err3
			}
		}
		m.OtherUpdates = v4

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// ChannelDifference <--
type ChannelDifference = TLChannelDifference

// DifferenceClazz <--
//   - TL_DifferenceEmpty
//   - TL_Difference
//   - TL_DifferenceSlice
//   - TL_DifferenceTooLong
type DifferenceClazz interface {
	iface.TLObject
	DifferenceClazzName() string
}

func DecodeDifferenceClazz(d *bin.Decoder) (DifferenceClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x8bdbda4e:
		x := &TLDifferenceEmpty{ClazzID: id, ClazzName2: ClazzName_differenceEmpty}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x5482832b:
		x := &TLDifference{ClazzID: id, ClazzName2: ClazzName_difference}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0xcb965ddf:
		x := &TLDifferenceSlice{ClazzID: id, ClazzName2: ClazzName_differenceSlice}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	case 0x3572ee30:
		x := &TLDifferenceTooLong{ClazzID: id, ClazzName2: ClazzName_differenceTooLong}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeDifference - unexpected clazzId: %d", id)
	}

}

// TLDifferenceEmpty <--
type TLDifferenceEmpty struct {
	ClazzID    uint32               `json:"_id"`
	ClazzName2 string               `json:"_name"`
	State      tg.UpdatesStateClazz `json:"state"`
}

func MakeTLDifferenceEmpty(m *TLDifferenceEmpty) *TLDifferenceEmpty {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_differenceEmpty

	return m
}

func (m *TLDifferenceEmpty) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLDifferenceEmpty) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("differenceEmpty", m)
}

// DifferenceClazzName <--
func (m *TLDifferenceEmpty) DifferenceClazzName() string {
	return ClazzName_differenceEmpty
}

// ClazzName <--
func (m *TLDifferenceEmpty) ClazzName() string {
	return m.ClazzName2
}

// ToDifference <--
func (m *TLDifferenceEmpty) ToDifference() *Difference {
	if m == nil {
		return nil
	}

	return &Difference{Clazz: m}

}

func (m *TLDifferenceEmpty) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_differenceEmpty, int(layer)); clazzId {
	case 0x8bdbda4e:
		size := 4
		size += iface.CalcObjectSize(m.State, layer)

		return size
	default:
		return 0
	}
}

func (m *TLDifferenceEmpty) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_differenceEmpty, int(layer)); clazzId {
	case 0x8bdbda4e:
		if err := iface.ValidateRequiredObject("state", m.State); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_differenceEmpty, layer)
	}
}

// Encode <--
func (m *TLDifferenceEmpty) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_differenceEmpty, int(layer)); clazzId {
	case 0x8bdbda4e:
		x.PutClazzID(0x8bdbda4e)

		_ = m.State.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_differenceEmpty, layer)
	}
}

// Decode <--
func (m *TLDifferenceEmpty) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x8bdbda4e:

		m.State, err = tg.DecodeUpdatesStateClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDifference <--
type TLDifference struct {
	ClazzID      uint32               `json:"_id"`
	ClazzName2   string               `json:"_name"`
	NewMessages  []tg.MessageClazz    `json:"new_messages"`
	OtherUpdates []tg.UpdateClazz     `json:"other_updates"`
	State        tg.UpdatesStateClazz `json:"state"`
}

func MakeTLDifference(m *TLDifference) *TLDifference {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_difference

	return m
}

func (m *TLDifference) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLDifference) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("difference", m)
}

// DifferenceClazzName <--
func (m *TLDifference) DifferenceClazzName() string {
	return ClazzName_difference
}

// ClazzName <--
func (m *TLDifference) ClazzName() string {
	return m.ClazzName2
}

// ToDifference <--
func (m *TLDifference) ToDifference() *Difference {
	if m == nil {
		return nil
	}

	return &Difference{Clazz: m}

}

func (m *TLDifference) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_difference, int(layer)); clazzId {
	case 0x5482832b:
		size := 4
		size += iface.CalcObjectListSize(m.NewMessages, layer)
		size += iface.CalcObjectListSize(m.OtherUpdates, layer)
		size += iface.CalcObjectSize(m.State, layer)

		return size
	default:
		return 0
	}
}

func (m *TLDifference) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_difference, int(layer)); clazzId {
	case 0x5482832b:
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
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_difference, layer)
	}
}

// Encode <--
func (m *TLDifference) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_difference, int(layer)); clazzId {
	case 0x5482832b:
		x.PutClazzID(0x5482832b)

		if err := iface.EncodeObjectList(x, m.NewMessages, layer); err != nil {
			return err
		}

		if err := iface.EncodeObjectList(x, m.OtherUpdates, layer); err != nil {
			return err
		}

		_ = m.State.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_difference, layer)
	}
}

// Decode <--
func (m *TLDifference) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x5482832b:
		c1, err2 := d.ClazzID()
		if err2 != nil {
			return err2
		}
		if c1 != iface.ClazzID_vector {
			return fmt.Errorf("invalid ClazzID_vector, c%d: %d", 1, c1)
		}
		l1, err3 := d.Int()
		if err3 != nil {
			return err3
		}
		v1 := make([]tg.MessageClazz, l1)
		for i := 0; i < l1; i++ {
			v1[i], err3 = tg.DecodeMessageClazz(d)
			if err3 != nil {
				return err3
			}
		}
		m.NewMessages = v1

		c2, err2 := d.ClazzID()
		if err2 != nil {
			return err2
		}
		if c2 != iface.ClazzID_vector {
			return fmt.Errorf("invalid ClazzID_vector, c%d: %d", 2, c2)
		}
		l2, err3 := d.Int()
		if err3 != nil {
			return err3
		}
		v2 := make([]tg.UpdateClazz, l2)
		for i := 0; i < l2; i++ {
			v2[i], err3 = tg.DecodeUpdateClazz(d)
			if err3 != nil {
				return err3
			}
		}
		m.OtherUpdates = v2

		m.State, err = tg.DecodeUpdatesStateClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDifferenceSlice <--
type TLDifferenceSlice struct {
	ClazzID           uint32               `json:"_id"`
	ClazzName2        string               `json:"_name"`
	NewMessages       []tg.MessageClazz    `json:"new_messages"`
	OtherUpdates      []tg.UpdateClazz     `json:"other_updates"`
	IntermediateState tg.UpdatesStateClazz `json:"intermediate_state"`
}

func MakeTLDifferenceSlice(m *TLDifferenceSlice) *TLDifferenceSlice {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_differenceSlice

	return m
}

func (m *TLDifferenceSlice) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLDifferenceSlice) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("differenceSlice", m)
}

// DifferenceClazzName <--
func (m *TLDifferenceSlice) DifferenceClazzName() string {
	return ClazzName_differenceSlice
}

// ClazzName <--
func (m *TLDifferenceSlice) ClazzName() string {
	return m.ClazzName2
}

// ToDifference <--
func (m *TLDifferenceSlice) ToDifference() *Difference {
	if m == nil {
		return nil
	}

	return &Difference{Clazz: m}

}

func (m *TLDifferenceSlice) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_differenceSlice, int(layer)); clazzId {
	case 0xcb965ddf:
		size := 4
		size += iface.CalcObjectListSize(m.NewMessages, layer)
		size += iface.CalcObjectListSize(m.OtherUpdates, layer)
		size += iface.CalcObjectSize(m.IntermediateState, layer)

		return size
	default:
		return 0
	}
}

func (m *TLDifferenceSlice) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_differenceSlice, int(layer)); clazzId {
	case 0xcb965ddf:
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
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_differenceSlice, layer)
	}
}

// Encode <--
func (m *TLDifferenceSlice) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_differenceSlice, int(layer)); clazzId {
	case 0xcb965ddf:
		x.PutClazzID(0xcb965ddf)

		if err := iface.EncodeObjectList(x, m.NewMessages, layer); err != nil {
			return err
		}

		if err := iface.EncodeObjectList(x, m.OtherUpdates, layer); err != nil {
			return err
		}

		_ = m.IntermediateState.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_differenceSlice, layer)
	}
}

// Decode <--
func (m *TLDifferenceSlice) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xcb965ddf:
		c1, err2 := d.ClazzID()
		if err2 != nil {
			return err2
		}
		if c1 != iface.ClazzID_vector {
			return fmt.Errorf("invalid ClazzID_vector, c%d: %d", 1, c1)
		}
		l1, err3 := d.Int()
		if err3 != nil {
			return err3
		}
		v1 := make([]tg.MessageClazz, l1)
		for i := 0; i < l1; i++ {
			v1[i], err3 = tg.DecodeMessageClazz(d)
			if err3 != nil {
				return err3
			}
		}
		m.NewMessages = v1

		c2, err2 := d.ClazzID()
		if err2 != nil {
			return err2
		}
		if c2 != iface.ClazzID_vector {
			return fmt.Errorf("invalid ClazzID_vector, c%d: %d", 2, c2)
		}
		l2, err3 := d.Int()
		if err3 != nil {
			return err3
		}
		v2 := make([]tg.UpdateClazz, l2)
		for i := 0; i < l2; i++ {
			v2[i], err3 = tg.DecodeUpdateClazz(d)
			if err3 != nil {
				return err3
			}
		}
		m.OtherUpdates = v2

		m.IntermediateState, err = tg.DecodeUpdatesStateClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDifferenceTooLong <--
type TLDifferenceTooLong struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Pts        int32  `json:"pts"`
}

func MakeTLDifferenceTooLong(m *TLDifferenceTooLong) *TLDifferenceTooLong {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_differenceTooLong

	return m
}

func (m *TLDifferenceTooLong) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLDifferenceTooLong) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("differenceTooLong", m)
}

// DifferenceClazzName <--
func (m *TLDifferenceTooLong) DifferenceClazzName() string {
	return ClazzName_differenceTooLong
}

// ClazzName <--
func (m *TLDifferenceTooLong) ClazzName() string {
	return m.ClazzName2
}

// ToDifference <--
func (m *TLDifferenceTooLong) ToDifference() *Difference {
	if m == nil {
		return nil
	}

	return &Difference{Clazz: m}

}

func (m *TLDifferenceTooLong) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_differenceTooLong, int(layer)); clazzId {
	case 0x3572ee30:
		size := 4
		size += 4

		return size
	default:
		return 0
	}
}

func (m *TLDifferenceTooLong) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_differenceTooLong, int(layer)); clazzId {
	case 0x3572ee30:

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_differenceTooLong, layer)
	}
}

// Encode <--
func (m *TLDifferenceTooLong) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_differenceTooLong, int(layer)); clazzId {
	case 0x3572ee30:
		x.PutClazzID(0x3572ee30)

		x.PutInt32(m.Pts)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_differenceTooLong, layer)
	}
}

// Decode <--
func (m *TLDifferenceTooLong) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x3572ee30:
		m.Pts, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// Difference <--
type Difference struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz DifferenceClazz `json:"_clazz"`
}

func (m *Difference) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *Difference) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName(m.ClazzName(), m)
}

func (m *Difference) CalcSize(layer int32) int {
	if m == nil || m.Clazz == nil {
		return 0
	}
	return iface.CalcObjectSize(m.Clazz, layer)
}

func (m *Difference) Validate(layer int32) error {
	if m == nil {
		return fmt.Errorf("Difference is required")
	}
	if m.Clazz == nil {
		return fmt.Errorf("Difference.Clazz is required")
	}
	if v, ok := m.Clazz.(iface.TLObjectValidator); ok {
		return v.Validate(layer)
	}
	return nil
}

func (m *Difference) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.DifferenceClazzName()
	}
}

// Encode <--
func (m *Difference) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("Difference - invalid Clazz")
}

// Decode <--
func (m *Difference) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeDifferenceClazz(d)
	return
}

// ToDifferenceEmpty <--
func (m *Difference) ToDifferenceEmpty() (*TLDifferenceEmpty, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLDifferenceEmpty); ok {
		return x, true
	}

	return nil, false
}

// ToDifference <--
func (m *Difference) ToDifference() (*TLDifference, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLDifference); ok {
		return x, true
	}

	return nil, false
}

// ToDifferenceSlice <--
func (m *Difference) ToDifferenceSlice() (*TLDifferenceSlice, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLDifferenceSlice); ok {
		return x, true
	}

	return nil, false
}

// ToDifferenceTooLong <--
func (m *Difference) ToDifferenceTooLong() (*TLDifferenceTooLong, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLDifferenceTooLong); ok {
		return x, true
	}

	return nil, false
}
