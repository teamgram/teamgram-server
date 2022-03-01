/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teagramio (teagram.io@gmail.com)
 */

// ConstructorList
// RequestList

package updates

import (
	"fmt"

	"github.com/teamgram/proto/mtproto"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/types"
)

//////////////////////////////////////////////////////////////////////////////////////////

var _ *types.Int32Value
var _ *mtproto.Bool
var _ fmt.GoStringer

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
	524332412: func() mtproto.TLObject { // 0x1f40ad7c
		return &TLUpdatesGetState{
			Constructor: 524332412,
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

func (m *ChannelDifference) Encode(layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	var (
		xBuf []byte
	)

	switch predicateName {
	case Predicate_channelDifference:
		t := m.To_ChannelDifference()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
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

func (m *ChannelDifference) DebugString() string {
	switch m.PredicateName {
	case Predicate_channelDifference:
		t := m.To_ChannelDifference()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_ChannelDifference
// channelDifference flags:# final:flags.0?true pts:int new_messages:Vector<Message> other_updates:Vector<Update> = ChannelDifference;
func (m *ChannelDifference) To_ChannelDifference() *TLChannelDifference {
	m.PredicateName = Predicate_channelDifference
	return &TLChannelDifference{
		Data2: m,
	}
}

// MakeTLChannelDifference
// channelDifference flags:# final:flags.0?true pts:int new_messages:Vector<Message> other_updates:Vector<Update> = ChannelDifference;
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

//// flags
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

func (m *TLChannelDifference) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xcd19034a: func() []byte {
			// channelDifference flags:# final:flags.0?true pts:int new_messages:Vector<Message> other_updates:Vector<Update> = ChannelDifference;
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
				x.Bytes((*v).Encode(layer))
			}

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetOtherUpdates())))
			for _, v := range m.GetOtherUpdates() {
				x.Bytes((*v).Encode(layer))
			}

			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_channelDifference, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_channelDifference, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLChannelDifference) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChannelDifference) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xcd19034a: func() error {
			// channelDifference flags:# final:flags.0?true pts:int new_messages:Vector<Message> other_updates:Vector<Update> = ChannelDifference;
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

func (m *TLChannelDifference) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// Difference <--
//  + TL_DifferenceEmpty
//  + TL_Difference
//  + TL_DifferenceSlice
//  + TL_DifferenceTooLong
//

func (m *Difference) Encode(layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	var (
		xBuf []byte
	)

	switch predicateName {
	case Predicate_differenceEmpty:
		t := m.To_DifferenceEmpty()
		xBuf = t.Encode(layer)
	case Predicate_difference:
		t := m.To_Difference()
		xBuf = t.Encode(layer)
	case Predicate_differenceSlice:
		t := m.To_DifferenceSlice()
		xBuf = t.Encode(layer)
	case Predicate_differenceTooLong:
		t := m.To_DifferenceTooLong()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
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

func (m *Difference) DebugString() string {
	switch m.PredicateName {
	case Predicate_differenceEmpty:
		t := m.To_DifferenceEmpty()
		return t.DebugString()
	case Predicate_difference:
		t := m.To_Difference()
		return t.DebugString()
	case Predicate_differenceSlice:
		t := m.To_DifferenceSlice()
		return t.DebugString()
	case Predicate_differenceTooLong:
		t := m.To_DifferenceTooLong()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_DifferenceEmpty
// differenceEmpty state:updates.State = Difference;
func (m *Difference) To_DifferenceEmpty() *TLDifferenceEmpty {
	m.PredicateName = Predicate_differenceEmpty
	return &TLDifferenceEmpty{
		Data2: m,
	}
}

// To_Difference
// difference new_messages:Vector<Message> other_updates:Vector<Update> state:updates.State = Difference;
func (m *Difference) To_Difference() *TLDifference {
	m.PredicateName = Predicate_difference
	return &TLDifference{
		Data2: m,
	}
}

// To_DifferenceSlice
// differenceSlice new_messages:Vector<Message> other_updates:Vector<Update> intermediate_state:updates.State = Difference;
func (m *Difference) To_DifferenceSlice() *TLDifferenceSlice {
	m.PredicateName = Predicate_differenceSlice
	return &TLDifferenceSlice{
		Data2: m,
	}
}

// To_DifferenceTooLong
// differenceTooLong pts:int = Difference;
func (m *Difference) To_DifferenceTooLong() *TLDifferenceTooLong {
	m.PredicateName = Predicate_differenceTooLong
	return &TLDifferenceTooLong{
		Data2: m,
	}
}

// MakeTLDifferenceEmpty
// differenceEmpty state:updates.State = Difference;
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

func (m *TLDifferenceEmpty) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x8bdbda4e: func() []byte {
			// differenceEmpty state:updates.State = Difference;
			x.UInt(0x8bdbda4e)

			x.Bytes(m.GetState().Encode(layer))
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_differenceEmpty, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_differenceEmpty, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLDifferenceEmpty) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDifferenceEmpty) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x8bdbda4e: func() error {
			// differenceEmpty state:updates.State = Difference;

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

func (m *TLDifferenceEmpty) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// MakeTLDifference
// difference new_messages:Vector<Message> other_updates:Vector<Update> state:updates.State = Difference;
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

func (m *TLDifference) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x5482832b: func() []byte {
			// difference new_messages:Vector<Message> other_updates:Vector<Update> state:updates.State = Difference;
			x.UInt(0x5482832b)

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetNewMessages())))
			for _, v := range m.GetNewMessages() {
				x.Bytes((*v).Encode(layer))
			}

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetOtherUpdates())))
			for _, v := range m.GetOtherUpdates() {
				x.Bytes((*v).Encode(layer))
			}

			x.Bytes(m.GetState().Encode(layer))
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_difference, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_difference, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLDifference) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDifference) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x5482832b: func() error {
			// difference new_messages:Vector<Message> other_updates:Vector<Update> state:updates.State = Difference;
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

func (m *TLDifference) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// MakeTLDifferenceSlice
// differenceSlice new_messages:Vector<Message> other_updates:Vector<Update> intermediate_state:updates.State = Difference;
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

func (m *TLDifferenceSlice) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xcb965ddf: func() []byte {
			// differenceSlice new_messages:Vector<Message> other_updates:Vector<Update> intermediate_state:updates.State = Difference;
			x.UInt(0xcb965ddf)

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetNewMessages())))
			for _, v := range m.GetNewMessages() {
				x.Bytes((*v).Encode(layer))
			}

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetOtherUpdates())))
			for _, v := range m.GetOtherUpdates() {
				x.Bytes((*v).Encode(layer))
			}

			x.Bytes(m.GetIntermediateState().Encode(layer))
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_differenceSlice, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_differenceSlice, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLDifferenceSlice) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDifferenceSlice) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xcb965ddf: func() error {
			// differenceSlice new_messages:Vector<Message> other_updates:Vector<Update> intermediate_state:updates.State = Difference;
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

func (m *TLDifferenceSlice) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// MakeTLDifferenceTooLong
// differenceTooLong pts:int = Difference;
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

func (m *TLDifferenceTooLong) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x3572ee30: func() []byte {
			// differenceTooLong pts:int = Difference;
			x.UInt(0x3572ee30)

			x.Int(m.GetPts())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_differenceTooLong, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_differenceTooLong, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLDifferenceTooLong) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDifferenceTooLong) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x3572ee30: func() error {
			// differenceTooLong pts:int = Difference;
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

func (m *TLDifferenceTooLong) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// TLUpdatesGetState
///////////////////////////////////////////////////////////////////////////////

func (m *TLUpdatesGetState) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_updates_getState))

	switch uint32(m.Constructor) {
	case 0x1f40ad7c:
		// updates.getState auth_key_id:long user_id:long = updates.State;
		x.UInt(0x1f40ad7c)

		// no flags

		x.Long(m.GetAuthKeyId())
		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUpdatesGetState) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUpdatesGetState) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x1f40ad7c:
		// updates.getState auth_key_id:long user_id:long = updates.State;

		// not has flags

		m.AuthKeyId = dBuf.Long()
		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUpdatesGetState) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUpdatesGetDifferenceV2
///////////////////////////////////////////////////////////////////////////////

func (m *TLUpdatesGetDifferenceV2) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_updates_getDifferenceV2))

	switch uint32(m.Constructor) {
	case 0xb76b6699:
		// updates.getDifferenceV2 flags:# auth_key_id:long user_id:long pts:int pts_total_limit:flags.0?int date:long = Difference;
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

	return x.GetBuf()
}

func (m *TLUpdatesGetDifferenceV2) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUpdatesGetDifferenceV2) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb76b6699:
		// updates.getDifferenceV2 flags:# auth_key_id:long user_id:long pts:int pts_total_limit:flags.0?int date:long = Difference;

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.AuthKeyId = dBuf.Long()
		m.UserId = dBuf.Long()
		m.Pts = dBuf.Int()
		if (flags & (1 << 0)) != 0 {
			m.PtsTotalLimit = &types.Int32Value{Value: dBuf.Int()}
		}

		m.Date = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUpdatesGetDifferenceV2) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUpdatesGetChannelDifferenceV2
///////////////////////////////////////////////////////////////////////////////

func (m *TLUpdatesGetChannelDifferenceV2) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_updates_getChannelDifferenceV2))

	switch uint32(m.Constructor) {
	case 0x4da3318a:
		// updates.getChannelDifferenceV2 auth_key_id:long user_id:long channel_id:long pts:int limit:int = ChannelDifference;
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

	return x.GetBuf()
}

func (m *TLUpdatesGetChannelDifferenceV2) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUpdatesGetChannelDifferenceV2) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x4da3318a:
		// updates.getChannelDifferenceV2 auth_key_id:long user_id:long channel_id:long pts:int limit:int = ChannelDifference;

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

func (m *TLUpdatesGetChannelDifferenceV2) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
