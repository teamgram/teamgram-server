/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

// ConstructorList
// RequestList

package updates

import (
	"fmt"

	"github.com/teamgram/proto/mtproto"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

//////////////////////////////////////////////////////////////////////////////////////////

var _ *wrapperspb.Int32Value
var _ *mtproto.Bool
var _ fmt.Stringer

var clazzIdRegisters2 = map[int32]func() mtproto.TLObject{
	// Constructor
	-853998774: func() mtproto.TLObject { // 0xcd19034a
		o := MakeTLChannelDifference(nil)
		o.Data2.Constructor = -853998774
		return o
	},
	-1948526002: func() mtproto.TLObject { // 0x8bdbda4e
		o := MakeTLDifferenceEmpty(nil)
		o.Data2.Constructor = -1948526002
		return o
	},
	1417839403: func() mtproto.TLObject { // 0x5482832b
		o := MakeTLDifference(nil)
		o.Data2.Constructor = 1417839403
		return o
	},
	-879338017: func() mtproto.TLObject { // 0xcb965ddf
		o := MakeTLDifferenceSlice(nil)
		o.Data2.Constructor = -879338017
		return o
	},
	896724528: func() mtproto.TLObject { // 0x3572ee30
		o := MakeTLDifferenceTooLong(nil)
		o.Data2.Constructor = 896724528
		return o
	},

	// Method
	1173671269: func() mtproto.TLObject { // 0x45f4cd65
		return &TLUpdatesGetStateV2{
			Constructor: 1173671269,
		}
	},
	-1217698151: func() mtproto.TLObject { // 0xb76b6699
		return &TLUpdatesGetDifferenceV2{
			Constructor: -1217698151,
		}
	},
	1302540682: func() mtproto.TLObject { // 0x4da3318a
		return &TLUpdatesGetChannelDifferenceV2{
			Constructor: 1302540682,
		}
	},
}

func NewTLObjectByClassID(classId int32) mtproto.TLObject {
	f, ok := clazzIdRegisters2[classId]
	if !ok {
		return nil
	}
	return f()
}

func CheckClassID(classId int32) (ok bool) {
	_, ok = clazzIdRegisters2[classId]
	return
}

//----------------------------------------------------------------------------------------------------------------

///////////////////////////////////////////////////////////////////////////////
// ChannelDifference <--
//  + TL_ChannelDifference
//

func (m *ChannelDifference) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_channelDifference:
		t := m.To_ChannelDifference()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *ChannelDifference) CalcByteSize(layer int32) int {
	return 0
}

func (m *ChannelDifference) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xcd19034a:
		m2 := MakeTLChannelDifference(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_ChannelDifference
func (m *ChannelDifference) To_ChannelDifference() *TLChannelDifference {
	m.PredicateName = Predicate_channelDifference
	return &TLChannelDifference{
		Data2: m,
	}
}

// MakeTLChannelDifference
func MakeTLChannelDifference(data2 *ChannelDifference) *TLChannelDifference {
	if data2 == nil {
		return &TLChannelDifference{Data2: &ChannelDifference{
			PredicateName: Predicate_channelDifference,
		}}
	} else {
		data2.PredicateName = Predicate_channelDifference
		return &TLChannelDifference{Data2: data2}
	}
}

func (m *TLChannelDifference) To_ChannelDifference() *ChannelDifference {
	m.Data2.PredicateName = Predicate_channelDifference
	return m.Data2
}

// // flags
func (m *TLChannelDifference) SetFinal(v bool) { m.Data2.Final = v }
func (m *TLChannelDifference) GetFinal() bool  { return m.Data2.Final }

func (m *TLChannelDifference) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLChannelDifference) GetPts() int32  { return m.Data2.Pts }

func (m *TLChannelDifference) SetNewMessages(v []*mtproto.Message) { m.Data2.NewMessages = v }
func (m *TLChannelDifference) GetNewMessages() []*mtproto.Message  { return m.Data2.NewMessages }

func (m *TLChannelDifference) SetOtherUpdates(v []*mtproto.Update) { m.Data2.OtherUpdates = v }
func (m *TLChannelDifference) GetOtherUpdates() []*mtproto.Update  { return m.Data2.OtherUpdates }

func (m *TLChannelDifference) GetPredicateName() string {
	return Predicate_channelDifference
}

func (m *TLChannelDifference) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xcd19034a: func() error {
			x.UInt(0xcd19034a)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.GetFinal() == true {
					flags |= 1 << 0
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.UInt(flags)
			x.Int(m.GetPts())

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetNewMessages())))
			for _, v := range m.GetNewMessages() {
				v.Encode(x, layer)
			}

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetOtherUpdates())))
			for _, v := range m.GetOtherUpdates() {
				v.Encode(x, layer)
			}

			return nil
		},
	}

	clazzId := GetClazzID(Predicate_channelDifference, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_channelDifference, layer)
		return nil
	}

	return nil
}

func (m *TLChannelDifference) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChannelDifference) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xcd19034a: func() error {
			var flags = dBuf.UInt()
			_ = flags
			if (flags & (1 << 0)) != 0 {
				m.SetFinal(true)
			}
			m.SetPts(dBuf.Int())
			c3 := dBuf.Int()
			if c3 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
			}
			l3 := dBuf.Int()
			v3 := make([]*mtproto.Message, l3)
			for i := int32(0); i < l3; i++ {
				v3[i] = &mtproto.Message{}
				v3[i].Decode(dBuf)
			}
			m.SetNewMessages(v3)

			c4 := dBuf.Int()
			if c4 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 4, c4)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 4, c4)
			}
			l4 := dBuf.Int()
			v4 := make([]*mtproto.Update, l4)
			for i := int32(0); i < l4; i++ {
				v4[i] = &mtproto.Update{}
				v4[i].Decode(dBuf)
			}
			m.SetOtherUpdates(v4)

			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

///////////////////////////////////////////////////////////////////////////////
// Difference <--
//  + TL_DifferenceEmpty
//  + TL_Difference
//  + TL_DifferenceSlice
//  + TL_DifferenceTooLong
//

func (m *Difference) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_differenceEmpty:
		t := m.To_DifferenceEmpty()
		t.Encode(x, layer)
	case Predicate_difference:
		t := m.To_Difference()
		t.Encode(x, layer)
	case Predicate_differenceSlice:
		t := m.To_DifferenceSlice()
		t.Encode(x, layer)
	case Predicate_differenceTooLong:
		t := m.To_DifferenceTooLong()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *Difference) CalcByteSize(layer int32) int {
	return 0
}

func (m *Difference) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x8bdbda4e:
		m2 := MakeTLDifferenceEmpty(m)
		m2.Decode(dBuf)
	case 0x5482832b:
		m2 := MakeTLDifference(m)
		m2.Decode(dBuf)
	case 0xcb965ddf:
		m2 := MakeTLDifferenceSlice(m)
		m2.Decode(dBuf)
	case 0x3572ee30:
		m2 := MakeTLDifferenceTooLong(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_DifferenceEmpty
func (m *Difference) To_DifferenceEmpty() *TLDifferenceEmpty {
	m.PredicateName = Predicate_differenceEmpty
	return &TLDifferenceEmpty{
		Data2: m,
	}
}

// To_Difference
func (m *Difference) To_Difference() *TLDifference {
	m.PredicateName = Predicate_difference
	return &TLDifference{
		Data2: m,
	}
}

// To_DifferenceSlice
func (m *Difference) To_DifferenceSlice() *TLDifferenceSlice {
	m.PredicateName = Predicate_differenceSlice
	return &TLDifferenceSlice{
		Data2: m,
	}
}

// To_DifferenceTooLong
func (m *Difference) To_DifferenceTooLong() *TLDifferenceTooLong {
	m.PredicateName = Predicate_differenceTooLong
	return &TLDifferenceTooLong{
		Data2: m,
	}
}

// MakeTLDifferenceEmpty
func MakeTLDifferenceEmpty(data2 *Difference) *TLDifferenceEmpty {
	if data2 == nil {
		return &TLDifferenceEmpty{Data2: &Difference{
			PredicateName: Predicate_differenceEmpty,
		}}
	} else {
		data2.PredicateName = Predicate_differenceEmpty
		return &TLDifferenceEmpty{Data2: data2}
	}
}

func (m *TLDifferenceEmpty) To_Difference() *Difference {
	m.Data2.PredicateName = Predicate_differenceEmpty
	return m.Data2
}

func (m *TLDifferenceEmpty) SetState(v *mtproto.Updates_State) { m.Data2.State = v }
func (m *TLDifferenceEmpty) GetState() *mtproto.Updates_State  { return m.Data2.State }

func (m *TLDifferenceEmpty) GetPredicateName() string {
	return Predicate_differenceEmpty
}

func (m *TLDifferenceEmpty) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x8bdbda4e: func() error {
			x.UInt(0x8bdbda4e)

			m.GetState().Encode(x, layer)
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_differenceEmpty, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_differenceEmpty, layer)
		return nil
	}

	return nil
}

func (m *TLDifferenceEmpty) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDifferenceEmpty) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x8bdbda4e: func() error {

			m0 := &mtproto.Updates_State{}
			m0.Decode(dBuf)
			m.SetState(m0)

			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

// MakeTLDifference
func MakeTLDifference(data2 *Difference) *TLDifference {
	if data2 == nil {
		return &TLDifference{Data2: &Difference{
			PredicateName: Predicate_difference,
		}}
	} else {
		data2.PredicateName = Predicate_difference
		return &TLDifference{Data2: data2}
	}
}

func (m *TLDifference) To_Difference() *Difference {
	m.Data2.PredicateName = Predicate_difference
	return m.Data2
}

func (m *TLDifference) SetNewMessages(v []*mtproto.Message) { m.Data2.NewMessages = v }
func (m *TLDifference) GetNewMessages() []*mtproto.Message  { return m.Data2.NewMessages }

func (m *TLDifference) SetOtherUpdates(v []*mtproto.Update) { m.Data2.OtherUpdates = v }
func (m *TLDifference) GetOtherUpdates() []*mtproto.Update  { return m.Data2.OtherUpdates }

func (m *TLDifference) SetState(v *mtproto.Updates_State) { m.Data2.State = v }
func (m *TLDifference) GetState() *mtproto.Updates_State  { return m.Data2.State }

func (m *TLDifference) GetPredicateName() string {
	return Predicate_difference
}

func (m *TLDifference) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x5482832b: func() error {
			x.UInt(0x5482832b)

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetNewMessages())))
			for _, v := range m.GetNewMessages() {
				v.Encode(x, layer)
			}

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetOtherUpdates())))
			for _, v := range m.GetOtherUpdates() {
				v.Encode(x, layer)
			}

			m.GetState().Encode(x, layer)
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_difference, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_difference, layer)
		return nil
	}

	return nil
}

func (m *TLDifference) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDifference) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x5482832b: func() error {
			c1 := dBuf.Int()
			if c1 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
			}
			l1 := dBuf.Int()
			v1 := make([]*mtproto.Message, l1)
			for i := int32(0); i < l1; i++ {
				v1[i] = &mtproto.Message{}
				v1[i].Decode(dBuf)
			}
			m.SetNewMessages(v1)

			c2 := dBuf.Int()
			if c2 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 2, c2)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 2, c2)
			}
			l2 := dBuf.Int()
			v2 := make([]*mtproto.Update, l2)
			for i := int32(0); i < l2; i++ {
				v2[i] = &mtproto.Update{}
				v2[i].Decode(dBuf)
			}
			m.SetOtherUpdates(v2)

			m0 := &mtproto.Updates_State{}
			m0.Decode(dBuf)
			m.SetState(m0)

			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

// MakeTLDifferenceSlice
func MakeTLDifferenceSlice(data2 *Difference) *TLDifferenceSlice {
	if data2 == nil {
		return &TLDifferenceSlice{Data2: &Difference{
			PredicateName: Predicate_differenceSlice,
		}}
	} else {
		data2.PredicateName = Predicate_differenceSlice
		return &TLDifferenceSlice{Data2: data2}
	}
}

func (m *TLDifferenceSlice) To_Difference() *Difference {
	m.Data2.PredicateName = Predicate_differenceSlice
	return m.Data2
}

func (m *TLDifferenceSlice) SetNewMessages(v []*mtproto.Message) { m.Data2.NewMessages = v }
func (m *TLDifferenceSlice) GetNewMessages() []*mtproto.Message  { return m.Data2.NewMessages }

func (m *TLDifferenceSlice) SetOtherUpdates(v []*mtproto.Update) { m.Data2.OtherUpdates = v }
func (m *TLDifferenceSlice) GetOtherUpdates() []*mtproto.Update  { return m.Data2.OtherUpdates }

func (m *TLDifferenceSlice) SetIntermediateState(v *mtproto.Updates_State) {
	m.Data2.IntermediateState = v
}
func (m *TLDifferenceSlice) GetIntermediateState() *mtproto.Updates_State {
	return m.Data2.IntermediateState
}

func (m *TLDifferenceSlice) GetPredicateName() string {
	return Predicate_differenceSlice
}

func (m *TLDifferenceSlice) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xcb965ddf: func() error {
			x.UInt(0xcb965ddf)

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetNewMessages())))
			for _, v := range m.GetNewMessages() {
				v.Encode(x, layer)
			}

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetOtherUpdates())))
			for _, v := range m.GetOtherUpdates() {
				v.Encode(x, layer)
			}

			m.GetIntermediateState().Encode(x, layer)
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_differenceSlice, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_differenceSlice, layer)
		return nil
	}

	return nil
}

func (m *TLDifferenceSlice) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDifferenceSlice) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xcb965ddf: func() error {
			c1 := dBuf.Int()
			if c1 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
			}
			l1 := dBuf.Int()
			v1 := make([]*mtproto.Message, l1)
			for i := int32(0); i < l1; i++ {
				v1[i] = &mtproto.Message{}
				v1[i].Decode(dBuf)
			}
			m.SetNewMessages(v1)

			c2 := dBuf.Int()
			if c2 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 2, c2)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 2, c2)
			}
			l2 := dBuf.Int()
			v2 := make([]*mtproto.Update, l2)
			for i := int32(0); i < l2; i++ {
				v2[i] = &mtproto.Update{}
				v2[i].Decode(dBuf)
			}
			m.SetOtherUpdates(v2)

			m3 := &mtproto.Updates_State{}
			m3.Decode(dBuf)
			m.SetIntermediateState(m3)

			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

// MakeTLDifferenceTooLong
func MakeTLDifferenceTooLong(data2 *Difference) *TLDifferenceTooLong {
	if data2 == nil {
		return &TLDifferenceTooLong{Data2: &Difference{
			PredicateName: Predicate_differenceTooLong,
		}}
	} else {
		data2.PredicateName = Predicate_differenceTooLong
		return &TLDifferenceTooLong{Data2: data2}
	}
}

func (m *TLDifferenceTooLong) To_Difference() *Difference {
	m.Data2.PredicateName = Predicate_differenceTooLong
	return m.Data2
}

func (m *TLDifferenceTooLong) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLDifferenceTooLong) GetPts() int32  { return m.Data2.Pts }

func (m *TLDifferenceTooLong) GetPredicateName() string {
	return Predicate_differenceTooLong
}

func (m *TLDifferenceTooLong) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x3572ee30: func() error {
			x.UInt(0x3572ee30)

			x.Int(m.GetPts())
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_differenceTooLong, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_differenceTooLong, layer)
		return nil
	}

	return nil
}

func (m *TLDifferenceTooLong) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDifferenceTooLong) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x3572ee30: func() error {
			m.SetPts(dBuf.Int())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

//----------------------------------------------------------------------------------------------------------------
// TLUpdatesGetStateV2
///////////////////////////////////////////////////////////////////////////////

func (m *TLUpdatesGetStateV2) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x45f4cd65:
		x.UInt(0x45f4cd65)

		// no flags

		x.Long(m.GetAuthKeyId())
		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUpdatesGetStateV2) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUpdatesGetStateV2) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x45f4cd65:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUpdatesGetDifferenceV2
///////////////////////////////////////////////////////////////////////////////

func (m *TLUpdatesGetDifferenceV2) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xb76b6699:
		x.UInt(0xb76b6699)

		// set flags
		var flags uint32 = 0

		if m.GetPtsTotalLimit() != nil {
			flags |= 1 << 0
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetAuthKeyId())
		x.Long(m.GetUserId())
		x.Int(m.GetPts())
		if m.GetPtsTotalLimit() != nil {
			x.Int(m.GetPtsTotalLimit().Value)
		}

		x.Long(m.GetDate())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUpdatesGetDifferenceV2) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUpdatesGetDifferenceV2) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb76b6699:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.AuthKeyId = dBuf.Long()
		m.UserId = dBuf.Long()
		m.Pts = dBuf.Int()
		if (flags & (1 << 0)) != 0 {
			m.PtsTotalLimit = &wrapperspb.Int32Value{Value: dBuf.Int()}
		}

		m.Date = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUpdatesGetChannelDifferenceV2
///////////////////////////////////////////////////////////////////////////////

func (m *TLUpdatesGetChannelDifferenceV2) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x4da3318a:
		x.UInt(0x4da3318a)

		// no flags

		x.Long(m.GetAuthKeyId())
		x.Long(m.GetUserId())
		x.Long(m.GetChannelId())
		x.Int(m.GetPts())
		x.Int(m.GetLimit())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUpdatesGetChannelDifferenceV2) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUpdatesGetChannelDifferenceV2) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x4da3318a:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		m.UserId = dBuf.Long()
		m.ChannelId = dBuf.Long()
		m.Pts = dBuf.Int()
		m.Limit = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}
