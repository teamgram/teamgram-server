/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package updates

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

// ChannelDifferenceClazz <--
//   - TL_ChannelDifference
type ChannelDifferenceClazz interface {
	iface.TLObject
	ChannelDifferenceClazzName() string
}

func DecodeChannelDifferenceClazz(d *bin.Decoder) (ChannelDifferenceClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_channelDifference:
		x := &TLChannelDifference{ClazzID: id, ClazzName2: ClazzName_channelDifference}
		_ = x.Decode(d)
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

	return &ChannelDifference{Clazz: m}
}

// Encode <--
func (m *TLChannelDifference) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xcd19034a: func() error {
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

			_ = iface.EncodeObjectList(x, m.NewMessages, layer)

			_ = iface.EncodeObjectList(x, m.OtherUpdates, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_channelDifference, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_channelDifference, layer)
	}
}

// Decode <--
func (m *TLChannelDifference) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xcd19034a: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			if (flags & (1 << 0)) != 0 {
				m.Final = true
			}
			m.Pts, err = d.Int32()
			c3, err2 := d.ClazzID()
			if c3 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 3, c3)
				return err2
			}
			l3, err3 := d.Int()
			v3 := make([]tg.MessageClazz, l3)
			for i := 0; i < l3; i++ {
				// vv := new(Message)
				// err3 = vv.Decode(d)
				// _ = err3
				// v3[i] = vv
				v3[i], err3 = tg.DecodeMessageClazz(d)
				_ = err3
			}
			m.NewMessages = v3

			c4, err2 := d.ClazzID()
			if c4 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 4, c4)
				return err2
			}
			l4, err3 := d.Int()
			v4 := make([]tg.UpdateClazz, l4)
			for i := 0; i < l4; i++ {
				// vv := new(Update)
				// err3 = vv.Decode(d)
				// _ = err3
				// v4[i] = vv
				v4[i], err3 = tg.DecodeUpdateClazz(d)
				_ = err3
			}
			m.OtherUpdates = v4

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// ChannelDifference <--
type ChannelDifference struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz ChannelDifferenceClazz `json:"_clazz"`
}

func (m *ChannelDifference) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *ChannelDifference) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.ChannelDifferenceClazzName()
	}
}

// Encode <--
func (m *ChannelDifference) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("ChannelDifference - invalid Clazz")
}

// Decode <--
func (m *ChannelDifference) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeChannelDifferenceClazz(d)
	return
}

// Match <--
func (m *ChannelDifference) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLChannelDifference:
		for _, v := range f {
			if f1, ok := v.(func(c *TLChannelDifference) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToChannelDifference <--
func (m *ChannelDifference) ToChannelDifference() (*TLChannelDifference, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLChannelDifference); ok {
		return x, true
	}

	return nil, false
}

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

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_differenceEmpty:
		x := &TLDifferenceEmpty{ClazzID: id, ClazzName2: ClazzName_differenceEmpty}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_difference:
		x := &TLDifference{ClazzID: id, ClazzName2: ClazzName_difference}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_differenceSlice:
		x := &TLDifferenceSlice{ClazzID: id, ClazzName2: ClazzName_differenceSlice}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_differenceTooLong:
		x := &TLDifferenceTooLong{ClazzID: id, ClazzName2: ClazzName_differenceTooLong}
		_ = x.Decode(d)
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

// Encode <--
func (m *TLDifferenceEmpty) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x8bdbda4e: func() error {
			x.PutClazzID(0x8bdbda4e)

			_ = m.State.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_differenceEmpty, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_differenceEmpty, layer)
	}
}

// Decode <--
func (m *TLDifferenceEmpty) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x8bdbda4e: func() (err error) {

			// m0 := &tg.UpdatesState{}
			// _ = m0.Decode(d)
			// m.State = m0
			m.State, _ = tg.DecodeUpdatesStateClazz(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
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

// Encode <--
func (m *TLDifference) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x5482832b: func() error {
			x.PutClazzID(0x5482832b)

			_ = iface.EncodeObjectList(x, m.NewMessages, layer)

			_ = iface.EncodeObjectList(x, m.OtherUpdates, layer)

			_ = m.State.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_difference, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_difference, layer)
	}
}

// Decode <--
func (m *TLDifference) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x5482832b: func() (err error) {
			c1, err2 := d.ClazzID()
			if c1 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 1, c1)
				return err2
			}
			l1, err3 := d.Int()
			v1 := make([]tg.MessageClazz, l1)
			for i := 0; i < l1; i++ {
				// vv := new(Message)
				// err3 = vv.Decode(d)
				// _ = err3
				// v1[i] = vv
				v1[i], err3 = tg.DecodeMessageClazz(d)
				_ = err3
			}
			m.NewMessages = v1

			c2, err2 := d.ClazzID()
			if c2 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 2, c2)
				return err2
			}
			l2, err3 := d.Int()
			v2 := make([]tg.UpdateClazz, l2)
			for i := 0; i < l2; i++ {
				// vv := new(Update)
				// err3 = vv.Decode(d)
				// _ = err3
				// v2[i] = vv
				v2[i], err3 = tg.DecodeUpdateClazz(d)
				_ = err3
			}
			m.OtherUpdates = v2

			// m0 := &tg.UpdatesState{}
			// _ = m0.Decode(d)
			// m.State = m0
			m.State, _ = tg.DecodeUpdatesStateClazz(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
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

// Encode <--
func (m *TLDifferenceSlice) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xcb965ddf: func() error {
			x.PutClazzID(0xcb965ddf)

			_ = iface.EncodeObjectList(x, m.NewMessages, layer)

			_ = iface.EncodeObjectList(x, m.OtherUpdates, layer)

			_ = m.IntermediateState.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_differenceSlice, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_differenceSlice, layer)
	}
}

// Decode <--
func (m *TLDifferenceSlice) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xcb965ddf: func() (err error) {
			c1, err2 := d.ClazzID()
			if c1 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 1, c1)
				return err2
			}
			l1, err3 := d.Int()
			v1 := make([]tg.MessageClazz, l1)
			for i := 0; i < l1; i++ {
				// vv := new(Message)
				// err3 = vv.Decode(d)
				// _ = err3
				// v1[i] = vv
				v1[i], err3 = tg.DecodeMessageClazz(d)
				_ = err3
			}
			m.NewMessages = v1

			c2, err2 := d.ClazzID()
			if c2 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 2, c2)
				return err2
			}
			l2, err3 := d.Int()
			v2 := make([]tg.UpdateClazz, l2)
			for i := 0; i < l2; i++ {
				// vv := new(Update)
				// err3 = vv.Decode(d)
				// _ = err3
				// v2[i] = vv
				v2[i], err3 = tg.DecodeUpdateClazz(d)
				_ = err3
			}
			m.OtherUpdates = v2

			// m3 := &tg.UpdatesState{}
			// _ = m3.Decode(d)
			// m.IntermediateState = m3
			m.IntermediateState, _ = tg.DecodeUpdatesStateClazz(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
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

// Encode <--
func (m *TLDifferenceTooLong) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x3572ee30: func() error {
			x.PutClazzID(0x3572ee30)

			x.PutInt32(m.Pts)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_differenceTooLong, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_differenceTooLong, layer)
	}
}

// Decode <--
func (m *TLDifferenceTooLong) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x3572ee30: func() (err error) {
			m.Pts, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
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

// Match <--
func (m *Difference) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLDifferenceEmpty:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDifferenceEmpty) interface{}); ok {
				f1(c)
			}
		}
	case *TLDifference:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDifference) interface{}); ok {
				f1(c)
			}
		}
	case *TLDifferenceSlice:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDifferenceSlice) interface{}); ok {
				f1(c)
			}
		}
	case *TLDifferenceTooLong:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDifferenceTooLong) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
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
