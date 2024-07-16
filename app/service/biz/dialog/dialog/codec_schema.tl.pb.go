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

package dialog

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
	-1109809056: func() mtproto.TLObject { // 0xbdd9a860
		o := MakeTLDialogExt(nil)
		o.Data2.Constructor = -1109809056
		return o
	},
	-1496016642: func() mtproto.TLObject { // 0xa6d498fe
		o := MakeTLDialogFilterExt(nil)
		o.Data2.Constructor = -1496016642
		return o
	},
	245834284: func() mtproto.TLObject { // 0xea7222c
		o := MakeTLDialogPinnedExt(nil)
		o.Data2.Constructor = 245834284
		return o
	},
	492418141: func() mtproto.TLObject { // 0x1d59b45d
		o := MakeTLSimpleDialogsData(nil)
		o.Data2.Constructor = 492418141
		return o
	},
	-155335502: func() mtproto.TLObject { // 0xf6bdc4b2
		o := MakeTLUpdateDraftMessage(nil)
		o.Data2.Constructor = -155335502
		return o
	},
	2005919834: func() mtproto.TLObject { // 0x778fe85a
		o := MakeTLSavedDialogList(nil)
		o.Data2.Constructor = 2005919834
		return o
	},

	// Method
	1321916826: func() mtproto.TLObject { // 0x4ecad99a
		return &TLDialogSaveDraftMessage{
			Constructor: 1321916826,
		}
	},
	-76500326: func() mtproto.TLObject { // 0xfb70b29a
		return &TLDialogClearDraftMessage{
			Constructor: -76500326,
		}
	},
	-1394716698: func() mtproto.TLObject { // 0xacde4fe6
		return &TLDialogGetAllDrafts{
			Constructor: -1394716698,
		}
	},
	1102614780: func() mtproto.TLObject { // 0x41b890fc
		return &TLDialogClearAllDrafts{
			Constructor: 1102614780,
		}
	},
	1160941838: func() mtproto.TLObject { // 0x4532910e
		return &TLDialogMarkDialogUnread{
			Constructor: 1160941838,
		}
	},
	-2038504145: func() mtproto.TLObject { // 0x867ee52f
		return &TLDialogToggleDialogPin{
			Constructor: -2038504145,
		}
	},
	-893634316: func() mtproto.TLObject { // 0xcabc38f4
		return &TLDialogGetDialogUnreadMarkList{
			Constructor: -893634316,
		}
	},
	-1652652540: func() mtproto.TLObject { // 0x9d7e8604
		return &TLDialogGetDialogsByOffsetDate{
			Constructor: -1652652540,
		}
	},
	-2046091754: func() mtproto.TLObject { // 0x860b1e16
		return &TLDialogGetDialogs{
			Constructor: -2046091754,
		}
	},
	-1390049167: func() mtproto.TLObject { // 0xad258871
		return &TLDialogGetDialogsByIdList{
			Constructor: -1390049167,
		}
	},
	-533089179: func() mtproto.TLObject { // 0xe039b465
		return &TLDialogGetDialogsCount{
			Constructor: -533089179,
		}
	},
	-1463673931: func() mtproto.TLObject { // 0xa8c21bb5
		return &TLDialogGetPinnedDialogs{
			Constructor: -1463673931,
		}
	},
	-18664089: func() mtproto.TLObject { // 0xfee33567
		return &TLDialogReorderPinnedDialogs{
			Constructor: -18664089,
		}
	},
	-1587594251: func() mtproto.TLObject { // 0xa15f3bf5
		return &TLDialogGetDialogById{
			Constructor: -1587594251,
		}
	},
	-92425614: func() mtproto.TLObject { // 0xfa7db272
		return &TLDialogGetTopMessage{
			Constructor: -92425614,
		}
	},
	489158840: func() mtproto.TLObject { // 0x1d27f8b8
		return &TLDialogUpdateReadInbox{
			Constructor: 489158840,
		}
	},
	1483799934: func() mtproto.TLObject { // 0x5870fd7e
		return &TLDialogUpdateReadOutbox{
			Constructor: 1483799934,
		}
	},
	382601889: func() mtproto.TLObject { // 0x16ce0aa1
		return &TLDialogInsertOrUpdateDialog{
			Constructor: 382601889,
		}
	},
	28515811: func() mtproto.TLObject { // 0x1b31de3
		return &TLDialogDeleteDialog{
			Constructor: 28515811,
		}
	},
	-1885617487: func() mtproto.TLObject { // 0x8f9bc2b1
		return &TLDialogGetUserPinnedMessage{
			Constructor: -1885617487,
		}
	},
	371388970: func() mtproto.TLObject { // 0x1622f22a
		return &TLDialogUpdateUserPinnedMessage{
			Constructor: 371388970,
		}
	},
	178824068: func() mtproto.TLObject { // 0xaa8a384
		return &TLDialogInsertOrUpdateDialogFilter{
			Constructor: 178824068,
		}
	},
	31276695: func() mtproto.TLObject { // 0x1dd3e97
		return &TLDialogDeleteDialogFilter{
			Constructor: 31276695,
		}
	},
	-1321465025: func() mtproto.TLObject { // 0xb13c0b3f
		return &TLDialogUpdateDialogFiltersOrder{
			Constructor: -1321465025,
		}
	},
	1818717244: func() mtproto.TLObject { // 0x6c676c3c
		return &TLDialogGetDialogFilters{
			Constructor: 1818717244,
		}
	},
	1092325045: func() mtproto.TLObject { // 0x411b8eb5
		return &TLDialogGetDialogFolder{
			Constructor: 1092325045,
		}
	},
	608601754: func() mtproto.TLObject { // 0x2446869a
		return &TLDialogEditPeerFolders{
			Constructor: 608601754,
		}
	},
	683494715: func() mtproto.TLObject { // 0x28bd4d3b
		return &TLDialogGetChannelMessageReadParticipants{
			Constructor: 683494715,
		}
	},
	-374431190: func() mtproto.TLObject { // 0xe9aea22a
		return &TLDialogSetChatTheme{
			Constructor: -374431190,
		}
	},
	165263532: func() mtproto.TLObject { // 0x9d9b8ac
		return &TLDialogSetHistoryTTL{
			Constructor: 165263532,
		}
	},
	2128645891: func() mtproto.TLObject { // 0x7ee08f03
		return &TLDialogGetMyDialogsData{
			Constructor: 2128645891,
		}
	},
	952227432: func() mtproto.TLObject { // 0x38c1d668
		return &TLDialogGetSavedDialogs{
			Constructor: 952227432,
		}
	},
	1084471271: func() mtproto.TLObject { // 0x40a3b7e7
		return &TLDialogGetPinnedSavedDialogs{
			Constructor: 1084471271,
		}
	},
	1156782041: func() mtproto.TLObject { // 0x44f317d9
		return &TLDialogToggleSavedDialogPin{
			Constructor: 1156782041,
		}
	},
	-665007150: func() mtproto.TLObject { // 0xd85ccbd2
		return &TLDialogReorderPinnedSavedDialogs{
			Constructor: -665007150,
		}
	},
	-209189348: func() mtproto.TLObject { // 0xf388061c
		return &TLDialogGetDialogFilter{
			Constructor: -209189348,
		}
	},
	1313177583: func() mtproto.TLObject { // 0x4e457fef
		return &TLDialogGetDialogFilterBySlug{
			Constructor: 1313177583,
		}
	},
	-959749265: func() mtproto.TLObject { // 0xc6cb636f
		return &TLDialogCreateDialogFilter{
			Constructor: -959749265,
		}
	},
	732705613: func() mtproto.TLObject { // 0x2bac334d
		return &TLDialogUpdateUnreadCount{
			Constructor: 732705613,
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
// DialogExt <--
//  + TL_DialogExt
//

func (m *DialogExt) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_dialogExt:
		t := m.To_DialogExt()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *DialogExt) CalcByteSize(layer int32) int {
	return 0
}

func (m *DialogExt) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xbdd9a860:
		m2 := MakeTLDialogExt(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_DialogExt
func (m *DialogExt) To_DialogExt() *TLDialogExt {
	m.PredicateName = Predicate_dialogExt
	return &TLDialogExt{
		Data2: m,
	}
}

// MakeTLDialogExt
func MakeTLDialogExt(data2 *DialogExt) *TLDialogExt {
	if data2 == nil {
		return &TLDialogExt{Data2: &DialogExt{
			PredicateName: Predicate_dialogExt,
		}}
	} else {
		data2.PredicateName = Predicate_dialogExt
		return &TLDialogExt{Data2: data2}
	}
}

func (m *TLDialogExt) To_DialogExt() *DialogExt {
	m.Data2.PredicateName = Predicate_dialogExt
	return m.Data2
}

func (m *TLDialogExt) SetOrder(v int64) { m.Data2.Order = v }
func (m *TLDialogExt) GetOrder() int64  { return m.Data2.Order }

func (m *TLDialogExt) SetDialog(v *mtproto.Dialog) { m.Data2.Dialog = v }
func (m *TLDialogExt) GetDialog() *mtproto.Dialog  { return m.Data2.Dialog }

func (m *TLDialogExt) SetAvailableMinId(v int32) { m.Data2.AvailableMinId = v }
func (m *TLDialogExt) GetAvailableMinId() int32  { return m.Data2.AvailableMinId }

func (m *TLDialogExt) SetDate(v int64) { m.Data2.Date = v }
func (m *TLDialogExt) GetDate() int64  { return m.Data2.Date }

func (m *TLDialogExt) SetThemeEmoticon(v string) { m.Data2.ThemeEmoticon = v }
func (m *TLDialogExt) GetThemeEmoticon() string  { return m.Data2.ThemeEmoticon }

func (m *TLDialogExt) SetTtlPeriod(v int32) { m.Data2.TtlPeriod = v }
func (m *TLDialogExt) GetTtlPeriod() int32  { return m.Data2.TtlPeriod }

func (m *TLDialogExt) GetPredicateName() string {
	return Predicate_dialogExt
}

func (m *TLDialogExt) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xbdd9a860: func() error {
			x.UInt(0xbdd9a860)

			x.Long(m.GetOrder())
			m.GetDialog().Encode(x, layer)
			x.Int(m.GetAvailableMinId())
			x.Long(m.GetDate())
			x.String(m.GetThemeEmoticon())
			x.Int(m.GetTtlPeriod())
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_dialogExt, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_dialogExt, layer)
		return nil
	}

	return nil
}

func (m *TLDialogExt) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogExt) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xbdd9a860: func() error {
			m.SetOrder(dBuf.Long())

			m1 := &mtproto.Dialog{}
			m1.Decode(dBuf)
			m.SetDialog(m1)

			m.SetAvailableMinId(dBuf.Int())
			m.SetDate(dBuf.Long())
			m.SetThemeEmoticon(dBuf.String())
			m.SetTtlPeriod(dBuf.Int())
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
// DialogFilterExt <--
//  + TL_DialogFilterExt
//

func (m *DialogFilterExt) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_dialogFilterExt:
		t := m.To_DialogFilterExt()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *DialogFilterExt) CalcByteSize(layer int32) int {
	return 0
}

func (m *DialogFilterExt) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xa6d498fe:
		m2 := MakeTLDialogFilterExt(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_DialogFilterExt
func (m *DialogFilterExt) To_DialogFilterExt() *TLDialogFilterExt {
	m.PredicateName = Predicate_dialogFilterExt
	return &TLDialogFilterExt{
		Data2: m,
	}
}

// MakeTLDialogFilterExt
func MakeTLDialogFilterExt(data2 *DialogFilterExt) *TLDialogFilterExt {
	if data2 == nil {
		return &TLDialogFilterExt{Data2: &DialogFilterExt{
			PredicateName: Predicate_dialogFilterExt,
		}}
	} else {
		data2.PredicateName = Predicate_dialogFilterExt
		return &TLDialogFilterExt{Data2: data2}
	}
}

func (m *TLDialogFilterExt) To_DialogFilterExt() *DialogFilterExt {
	m.Data2.PredicateName = Predicate_dialogFilterExt
	return m.Data2
}

// // flags
func (m *TLDialogFilterExt) SetId(v int32) { m.Data2.Id = v }
func (m *TLDialogFilterExt) GetId() int32  { return m.Data2.Id }

func (m *TLDialogFilterExt) SetJoinedBySlug(v bool) { m.Data2.JoinedBySlug = v }
func (m *TLDialogFilterExt) GetJoinedBySlug() bool  { return m.Data2.JoinedBySlug }

func (m *TLDialogFilterExt) SetSlug(v string) { m.Data2.Slug = v }
func (m *TLDialogFilterExt) GetSlug() string  { return m.Data2.Slug }

func (m *TLDialogFilterExt) SetDialogFilter(v *mtproto.DialogFilter) { m.Data2.DialogFilter = v }
func (m *TLDialogFilterExt) GetDialogFilter() *mtproto.DialogFilter  { return m.Data2.DialogFilter }

func (m *TLDialogFilterExt) SetOrder(v int64) { m.Data2.Order = v }
func (m *TLDialogFilterExt) GetOrder() int64  { return m.Data2.Order }

func (m *TLDialogFilterExt) GetPredicateName() string {
	return Predicate_dialogFilterExt
}

func (m *TLDialogFilterExt) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa6d498fe: func() error {
			x.UInt(0xa6d498fe)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.GetJoinedBySlug() == true {
					flags |= 1 << 0
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.UInt(flags)
			x.Int(m.GetId())
			x.String(m.GetSlug())
			m.GetDialogFilter().Encode(x, layer)
			x.Long(m.GetOrder())
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_dialogFilterExt, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_dialogFilterExt, layer)
		return nil
	}

	return nil
}

func (m *TLDialogFilterExt) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogFilterExt) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xa6d498fe: func() error {
			var flags = dBuf.UInt()
			_ = flags
			m.SetId(dBuf.Int())
			if (flags & (1 << 0)) != 0 {
				m.SetJoinedBySlug(true)
			}
			m.SetSlug(dBuf.String())

			m4 := &mtproto.DialogFilter{}
			m4.Decode(dBuf)
			m.SetDialogFilter(m4)

			m.SetOrder(dBuf.Long())
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
// DialogPinnedExt <--
//  + TL_DialogPinnedExt
//

func (m *DialogPinnedExt) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_dialogPinnedExt:
		t := m.To_DialogPinnedExt()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *DialogPinnedExt) CalcByteSize(layer int32) int {
	return 0
}

func (m *DialogPinnedExt) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xea7222c:
		m2 := MakeTLDialogPinnedExt(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_DialogPinnedExt
func (m *DialogPinnedExt) To_DialogPinnedExt() *TLDialogPinnedExt {
	m.PredicateName = Predicate_dialogPinnedExt
	return &TLDialogPinnedExt{
		Data2: m,
	}
}

// MakeTLDialogPinnedExt
func MakeTLDialogPinnedExt(data2 *DialogPinnedExt) *TLDialogPinnedExt {
	if data2 == nil {
		return &TLDialogPinnedExt{Data2: &DialogPinnedExt{
			PredicateName: Predicate_dialogPinnedExt,
		}}
	} else {
		data2.PredicateName = Predicate_dialogPinnedExt
		return &TLDialogPinnedExt{Data2: data2}
	}
}

func (m *TLDialogPinnedExt) To_DialogPinnedExt() *DialogPinnedExt {
	m.Data2.PredicateName = Predicate_dialogPinnedExt
	return m.Data2
}

func (m *TLDialogPinnedExt) SetOrder(v int64) { m.Data2.Order = v }
func (m *TLDialogPinnedExt) GetOrder() int64  { return m.Data2.Order }

func (m *TLDialogPinnedExt) SetPeerType(v int32) { m.Data2.PeerType = v }
func (m *TLDialogPinnedExt) GetPeerType() int32  { return m.Data2.PeerType }

func (m *TLDialogPinnedExt) SetPeerId(v int64) { m.Data2.PeerId = v }
func (m *TLDialogPinnedExt) GetPeerId() int64  { return m.Data2.PeerId }

func (m *TLDialogPinnedExt) GetPredicateName() string {
	return Predicate_dialogPinnedExt
}

func (m *TLDialogPinnedExt) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xea7222c: func() error {
			x.UInt(0xea7222c)

			x.Long(m.GetOrder())
			x.Int(m.GetPeerType())
			x.Long(m.GetPeerId())
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_dialogPinnedExt, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_dialogPinnedExt, layer)
		return nil
	}

	return nil
}

func (m *TLDialogPinnedExt) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogPinnedExt) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xea7222c: func() error {
			m.SetOrder(dBuf.Long())
			m.SetPeerType(dBuf.Int())
			m.SetPeerId(dBuf.Long())
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
// DialogsData <--
//  + TL_SimpleDialogsData
//

func (m *DialogsData) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_simpleDialogsData:
		t := m.To_SimpleDialogsData()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *DialogsData) CalcByteSize(layer int32) int {
	return 0
}

func (m *DialogsData) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x1d59b45d:
		m2 := MakeTLSimpleDialogsData(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_SimpleDialogsData
func (m *DialogsData) To_SimpleDialogsData() *TLSimpleDialogsData {
	m.PredicateName = Predicate_simpleDialogsData
	return &TLSimpleDialogsData{
		Data2: m,
	}
}

// MakeTLSimpleDialogsData
func MakeTLSimpleDialogsData(data2 *DialogsData) *TLSimpleDialogsData {
	if data2 == nil {
		return &TLSimpleDialogsData{Data2: &DialogsData{
			PredicateName: Predicate_simpleDialogsData,
		}}
	} else {
		data2.PredicateName = Predicate_simpleDialogsData
		return &TLSimpleDialogsData{Data2: data2}
	}
}

func (m *TLSimpleDialogsData) To_DialogsData() *DialogsData {
	m.Data2.PredicateName = Predicate_simpleDialogsData
	return m.Data2
}

func (m *TLSimpleDialogsData) SetUsers(v []int64) { m.Data2.Users = v }
func (m *TLSimpleDialogsData) GetUsers() []int64  { return m.Data2.Users }

func (m *TLSimpleDialogsData) SetChats(v []int64) { m.Data2.Chats = v }
func (m *TLSimpleDialogsData) GetChats() []int64  { return m.Data2.Chats }

func (m *TLSimpleDialogsData) SetChannels(v []int64) { m.Data2.Channels = v }
func (m *TLSimpleDialogsData) GetChannels() []int64  { return m.Data2.Channels }

func (m *TLSimpleDialogsData) GetPredicateName() string {
	return Predicate_simpleDialogsData
}

func (m *TLSimpleDialogsData) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x1d59b45d: func() error {
			x.UInt(0x1d59b45d)

			x.VectorLong(m.GetUsers())

			x.VectorLong(m.GetChats())

			x.VectorLong(m.GetChannels())

			return nil
		},
	}

	clazzId := GetClazzID(Predicate_simpleDialogsData, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_simpleDialogsData, layer)
		return nil
	}

	return nil
}

func (m *TLSimpleDialogsData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSimpleDialogsData) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x1d59b45d: func() error {

			m.SetUsers(dBuf.VectorLong())

			m.SetChats(dBuf.VectorLong())

			m.SetChannels(dBuf.VectorLong())

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
// PeerWithDraftMessage <--
//  + TL_UpdateDraftMessage
//

func (m *PeerWithDraftMessage) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_updateDraftMessage:
		t := m.To_UpdateDraftMessage()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *PeerWithDraftMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *PeerWithDraftMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xf6bdc4b2:
		m2 := MakeTLUpdateDraftMessage(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_UpdateDraftMessage
func (m *PeerWithDraftMessage) To_UpdateDraftMessage() *TLUpdateDraftMessage {
	m.PredicateName = Predicate_updateDraftMessage
	return &TLUpdateDraftMessage{
		Data2: m,
	}
}

// MakeTLUpdateDraftMessage
func MakeTLUpdateDraftMessage(data2 *PeerWithDraftMessage) *TLUpdateDraftMessage {
	if data2 == nil {
		return &TLUpdateDraftMessage{Data2: &PeerWithDraftMessage{
			PredicateName: Predicate_updateDraftMessage,
		}}
	} else {
		data2.PredicateName = Predicate_updateDraftMessage
		return &TLUpdateDraftMessage{Data2: data2}
	}
}

func (m *TLUpdateDraftMessage) To_PeerWithDraftMessage() *PeerWithDraftMessage {
	m.Data2.PredicateName = Predicate_updateDraftMessage
	return m.Data2
}

func (m *TLUpdateDraftMessage) SetPeer(v *mtproto.Peer) { m.Data2.Peer = v }
func (m *TLUpdateDraftMessage) GetPeer() *mtproto.Peer  { return m.Data2.Peer }

func (m *TLUpdateDraftMessage) SetDraft(v *mtproto.DraftMessage) { m.Data2.Draft = v }
func (m *TLUpdateDraftMessage) GetDraft() *mtproto.DraftMessage  { return m.Data2.Draft }

func (m *TLUpdateDraftMessage) GetPredicateName() string {
	return Predicate_updateDraftMessage
}

func (m *TLUpdateDraftMessage) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xf6bdc4b2: func() error {
			x.UInt(0xf6bdc4b2)

			m.GetPeer().Encode(x, layer)
			m.GetDraft().Encode(x, layer)
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_updateDraftMessage, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_updateDraftMessage, layer)
		return nil
	}

	return nil
}

func (m *TLUpdateDraftMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUpdateDraftMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xf6bdc4b2: func() error {

			m0 := &mtproto.Peer{}
			m0.Decode(dBuf)
			m.SetPeer(m0)

			m1 := &mtproto.DraftMessage{}
			m1.Decode(dBuf)
			m.SetDraft(m1)

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
// SavedDialogList <--
//  + TL_SavedDialogList
//

func (m *SavedDialogList) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_savedDialogList:
		t := m.To_SavedDialogList()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *SavedDialogList) CalcByteSize(layer int32) int {
	return 0
}

func (m *SavedDialogList) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x778fe85a:
		m2 := MakeTLSavedDialogList(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_SavedDialogList
func (m *SavedDialogList) To_SavedDialogList() *TLSavedDialogList {
	m.PredicateName = Predicate_savedDialogList
	return &TLSavedDialogList{
		Data2: m,
	}
}

// MakeTLSavedDialogList
func MakeTLSavedDialogList(data2 *SavedDialogList) *TLSavedDialogList {
	if data2 == nil {
		return &TLSavedDialogList{Data2: &SavedDialogList{
			PredicateName: Predicate_savedDialogList,
		}}
	} else {
		data2.PredicateName = Predicate_savedDialogList
		return &TLSavedDialogList{Data2: data2}
	}
}

func (m *TLSavedDialogList) To_SavedDialogList() *SavedDialogList {
	m.Data2.PredicateName = Predicate_savedDialogList
	return m.Data2
}

func (m *TLSavedDialogList) SetCount(v int32) { m.Data2.Count = v }
func (m *TLSavedDialogList) GetCount() int32  { return m.Data2.Count }

func (m *TLSavedDialogList) SetDialogs(v []*mtproto.SavedDialog) { m.Data2.Dialogs = v }
func (m *TLSavedDialogList) GetDialogs() []*mtproto.SavedDialog  { return m.Data2.Dialogs }

func (m *TLSavedDialogList) GetPredicateName() string {
	return Predicate_savedDialogList
}

func (m *TLSavedDialogList) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x778fe85a: func() error {
			x.UInt(0x778fe85a)

			x.Int(m.GetCount())

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetDialogs())))
			for _, v := range m.GetDialogs() {
				v.Encode(x, layer)
			}

			return nil
		},
	}

	clazzId := GetClazzID(Predicate_savedDialogList, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_savedDialogList, layer)
		return nil
	}

	return nil
}

func (m *TLSavedDialogList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSavedDialogList) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x778fe85a: func() error {
			m.SetCount(dBuf.Int())
			c1 := dBuf.Int()
			if c1 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
			}
			l1 := dBuf.Int()
			v1 := make([]*mtproto.SavedDialog, l1)
			for i := int32(0); i < l1; i++ {
				v1[i] = &mtproto.SavedDialog{}
				v1[i].Decode(dBuf)
			}
			m.SetDialogs(v1)

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
// TLDialogSaveDraftMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogSaveDraftMessage) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x4ecad99a:
		x.UInt(0x4ecad99a)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		m.GetMessage().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogSaveDraftMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogSaveDraftMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x4ecad99a:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()

		m4 := &mtproto.DraftMessage{}
		m4.Decode(dBuf)
		m.Message = m4

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogClearDraftMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogClearDraftMessage) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xfb70b29a:
		x.UInt(0xfb70b29a)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogClearDraftMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogClearDraftMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xfb70b29a:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogGetAllDrafts
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetAllDrafts) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xacde4fe6:
		x.UInt(0xacde4fe6)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogGetAllDrafts) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetAllDrafts) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xacde4fe6:

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogClearAllDrafts
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogClearAllDrafts) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x41b890fc:
		x.UInt(0x41b890fc)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogClearAllDrafts) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogClearAllDrafts) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x41b890fc:

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogMarkDialogUnread
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogMarkDialogUnread) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x4532910e:
		x.UInt(0x4532910e)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		m.GetUnreadMark().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogMarkDialogUnread) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogMarkDialogUnread) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x4532910e:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()

		m4 := &mtproto.Bool{}
		m4.Decode(dBuf)
		m.UnreadMark = m4

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogToggleDialogPin
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogToggleDialogPin) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x867ee52f:
		x.UInt(0x867ee52f)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		m.GetPinned().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogToggleDialogPin) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogToggleDialogPin) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x867ee52f:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()

		m4 := &mtproto.Bool{}
		m4.Decode(dBuf)
		m.Pinned = m4

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogGetDialogUnreadMarkList
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetDialogUnreadMarkList) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xcabc38f4:
		x.UInt(0xcabc38f4)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogGetDialogUnreadMarkList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetDialogUnreadMarkList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xcabc38f4:

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogGetDialogsByOffsetDate
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetDialogsByOffsetDate) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x9d7e8604:
		x.UInt(0x9d7e8604)

		// no flags

		x.Long(m.GetUserId())
		m.GetExcludePinned().Encode(x, layer)
		x.Int(m.GetOffsetDate())
		x.Int(m.GetLimit())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogGetDialogsByOffsetDate) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetDialogsByOffsetDate) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x9d7e8604:

		// not has flags

		m.UserId = dBuf.Long()

		m2 := &mtproto.Bool{}
		m2.Decode(dBuf)
		m.ExcludePinned = m2

		m.OffsetDate = dBuf.Int()
		m.Limit = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogGetDialogs
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetDialogs) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x860b1e16:
		x.UInt(0x860b1e16)

		// no flags

		x.Long(m.GetUserId())
		m.GetExcludePinned().Encode(x, layer)
		x.Int(m.GetFolderId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogGetDialogs) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetDialogs) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x860b1e16:

		// not has flags

		m.UserId = dBuf.Long()

		m2 := &mtproto.Bool{}
		m2.Decode(dBuf)
		m.ExcludePinned = m2

		m.FolderId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogGetDialogsByIdList
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetDialogsByIdList) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xad258871:
		x.UInt(0xad258871)

		// no flags

		x.Long(m.GetUserId())

		x.VectorLong(m.GetIdList())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogGetDialogsByIdList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetDialogsByIdList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xad258871:

		// not has flags

		m.UserId = dBuf.Long()

		m.IdList = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogGetDialogsCount
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetDialogsCount) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xe039b465:
		x.UInt(0xe039b465)

		// no flags

		x.Long(m.GetUserId())
		m.GetExcludePinned().Encode(x, layer)
		x.Int(m.GetFolderId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogGetDialogsCount) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetDialogsCount) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe039b465:

		// not has flags

		m.UserId = dBuf.Long()

		m2 := &mtproto.Bool{}
		m2.Decode(dBuf)
		m.ExcludePinned = m2

		m.FolderId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogGetPinnedDialogs
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetPinnedDialogs) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xa8c21bb5:
		x.UInt(0xa8c21bb5)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetFolderId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogGetPinnedDialogs) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetPinnedDialogs) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xa8c21bb5:

		// not has flags

		m.UserId = dBuf.Long()
		m.FolderId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogReorderPinnedDialogs
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogReorderPinnedDialogs) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xfee33567:
		x.UInt(0xfee33567)

		// no flags

		x.Long(m.GetUserId())
		m.GetForce().Encode(x, layer)
		x.Int(m.GetFolderId())

		x.VectorLong(m.GetIdList())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogReorderPinnedDialogs) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogReorderPinnedDialogs) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xfee33567:

		// not has flags

		m.UserId = dBuf.Long()

		m2 := &mtproto.Bool{}
		m2.Decode(dBuf)
		m.Force = m2

		m.FolderId = dBuf.Int()

		m.IdList = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogGetDialogById
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetDialogById) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xa15f3bf5:
		x.UInt(0xa15f3bf5)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogGetDialogById) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetDialogById) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xa15f3bf5:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogGetTopMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetTopMessage) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xfa7db272:
		x.UInt(0xfa7db272)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogGetTopMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetTopMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xfa7db272:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogUpdateReadInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogUpdateReadInbox) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x1d27f8b8:
		x.UInt(0x1d27f8b8)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetReadInboxId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogUpdateReadInbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogUpdateReadInbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x1d27f8b8:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.ReadInboxId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogUpdateReadOutbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogUpdateReadOutbox) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x5870fd7e:
		x.UInt(0x5870fd7e)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetReadOutboxId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogUpdateReadOutbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogUpdateReadOutbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x5870fd7e:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.ReadOutboxId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogInsertOrUpdateDialog
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogInsertOrUpdateDialog) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x16ce0aa1:
		x.UInt(0x16ce0aa1)

		// set flags
		var flags uint32 = 0

		if m.GetTopMessage() != nil {
			flags |= 1 << 0
		}
		if m.GetReadOutboxMaxId() != nil {
			flags |= 1 << 1
		}
		if m.GetReadInboxMaxId() != nil {
			flags |= 1 << 2
		}
		if m.GetUnreadCount() != nil {
			flags |= 1 << 3
		}
		if m.GetUnreadMark() == true {
			flags |= 1 << 4
		}
		if m.GetDate2() != nil {
			flags |= 1 << 5
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		if m.GetTopMessage() != nil {
			x.Int(m.GetTopMessage().Value)
		}

		if m.GetReadOutboxMaxId() != nil {
			x.Int(m.GetReadOutboxMaxId().Value)
		}

		if m.GetReadInboxMaxId() != nil {
			x.Int(m.GetReadInboxMaxId().Value)
		}

		if m.GetUnreadCount() != nil {
			x.Int(m.GetUnreadCount().Value)
		}

		if m.GetDate2() != nil {
			x.Long(m.GetDate2().Value)
		}

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogInsertOrUpdateDialog) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogInsertOrUpdateDialog) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x16ce0aa1:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		if (flags & (1 << 0)) != 0 {
			m.TopMessage = &wrapperspb.Int32Value{Value: dBuf.Int()}
		}

		if (flags & (1 << 1)) != 0 {
			m.ReadOutboxMaxId = &wrapperspb.Int32Value{Value: dBuf.Int()}
		}

		if (flags & (1 << 2)) != 0 {
			m.ReadInboxMaxId = &wrapperspb.Int32Value{Value: dBuf.Int()}
		}

		if (flags & (1 << 3)) != 0 {
			m.UnreadCount = &wrapperspb.Int32Value{Value: dBuf.Int()}
		}

		if (flags & (1 << 4)) != 0 {
			m.UnreadMark = true
		}
		if (flags & (1 << 5)) != 0 {
			m.Date2 = &wrapperspb.Int64Value{Value: dBuf.Long()}
		}

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogDeleteDialog
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogDeleteDialog) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x1b31de3:
		x.UInt(0x1b31de3)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogDeleteDialog) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogDeleteDialog) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x1b31de3:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogGetUserPinnedMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetUserPinnedMessage) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x8f9bc2b1:
		x.UInt(0x8f9bc2b1)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogGetUserPinnedMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetUserPinnedMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x8f9bc2b1:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogUpdateUserPinnedMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogUpdateUserPinnedMessage) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x1622f22a:
		x.UInt(0x1622f22a)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetPinnedMsgId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogUpdateUserPinnedMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogUpdateUserPinnedMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x1622f22a:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.PinnedMsgId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogInsertOrUpdateDialogFilter
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogInsertOrUpdateDialogFilter) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xaa8a384:
		x.UInt(0xaa8a384)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetId())
		m.GetDialogFilter().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogInsertOrUpdateDialogFilter) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogInsertOrUpdateDialogFilter) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xaa8a384:

		// not has flags

		m.UserId = dBuf.Long()
		m.Id = dBuf.Int()

		m3 := &mtproto.DialogFilter{}
		m3.Decode(dBuf)
		m.DialogFilter = m3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogDeleteDialogFilter
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogDeleteDialogFilter) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x1dd3e97:
		x.UInt(0x1dd3e97)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogDeleteDialogFilter) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogDeleteDialogFilter) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x1dd3e97:

		// not has flags

		m.UserId = dBuf.Long()
		m.Id = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogUpdateDialogFiltersOrder
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogUpdateDialogFiltersOrder) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xb13c0b3f:
		x.UInt(0xb13c0b3f)

		// no flags

		x.Long(m.GetUserId())

		x.VectorInt(m.GetOrder())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogUpdateDialogFiltersOrder) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogUpdateDialogFiltersOrder) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb13c0b3f:

		// not has flags

		m.UserId = dBuf.Long()

		m.Order = dBuf.VectorInt()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogGetDialogFilters
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetDialogFilters) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x6c676c3c:
		x.UInt(0x6c676c3c)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogGetDialogFilters) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetDialogFilters) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x6c676c3c:

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogGetDialogFolder
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetDialogFolder) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x411b8eb5:
		x.UInt(0x411b8eb5)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetFolderId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogGetDialogFolder) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetDialogFolder) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x411b8eb5:

		// not has flags

		m.UserId = dBuf.Long()
		m.FolderId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogEditPeerFolders
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogEditPeerFolders) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x2446869a:
		x.UInt(0x2446869a)

		// no flags

		x.Long(m.GetUserId())

		x.VectorLong(m.GetPeerDialogList())

		x.Int(m.GetFolderId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogEditPeerFolders) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogEditPeerFolders) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x2446869a:

		// not has flags

		m.UserId = dBuf.Long()

		m.PeerDialogList = dBuf.VectorLong()

		m.FolderId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogGetChannelMessageReadParticipants
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetChannelMessageReadParticipants) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x28bd4d3b:
		x.UInt(0x28bd4d3b)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetChannelId())
		x.Int(m.GetMsgId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogGetChannelMessageReadParticipants) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetChannelMessageReadParticipants) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x28bd4d3b:

		// not has flags

		m.UserId = dBuf.Long()
		m.ChannelId = dBuf.Long()
		m.MsgId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogSetChatTheme
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogSetChatTheme) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xe9aea22a:
		x.UInt(0xe9aea22a)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.String(m.GetThemeEmoticon())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogSetChatTheme) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogSetChatTheme) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe9aea22a:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.ThemeEmoticon = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogSetHistoryTTL
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogSetHistoryTTL) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x9d9b8ac:
		x.UInt(0x9d9b8ac)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetTtlPeriod())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogSetHistoryTTL) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogSetHistoryTTL) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x9d9b8ac:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.TtlPeriod = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogGetMyDialogsData
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetMyDialogsData) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x7ee08f03:
		x.UInt(0x7ee08f03)

		// set flags
		var flags uint32 = 0

		if m.GetUser() == true {
			flags |= 1 << 0
		}
		if m.GetChat() == true {
			flags |= 1 << 1
		}
		if m.GetChannel() == true {
			flags |= 1 << 2
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogGetMyDialogsData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetMyDialogsData) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x7ee08f03:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.UserId = dBuf.Long()
		if (flags & (1 << 0)) != 0 {
			m.User = true
		}
		if (flags & (1 << 1)) != 0 {
			m.Chat = true
		}
		if (flags & (1 << 2)) != 0 {
			m.Channel = true
		}
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogGetSavedDialogs
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetSavedDialogs) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x38c1d668:
		x.UInt(0x38c1d668)

		// no flags

		x.Long(m.GetUserId())
		m.GetExcludePinned().Encode(x, layer)
		x.Int(m.GetOffsetDate())
		x.Int(m.GetOffsetId())
		m.GetOffsetPeer().Encode(x, layer)
		x.Int(m.GetLimit())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogGetSavedDialogs) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetSavedDialogs) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x38c1d668:

		// not has flags

		m.UserId = dBuf.Long()

		m2 := &mtproto.Bool{}
		m2.Decode(dBuf)
		m.ExcludePinned = m2

		m.OffsetDate = dBuf.Int()
		m.OffsetId = dBuf.Int()

		m5 := &mtproto.PeerUtil{}
		m5.Decode(dBuf)
		m.OffsetPeer = m5

		m.Limit = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogGetPinnedSavedDialogs
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetPinnedSavedDialogs) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x40a3b7e7:
		x.UInt(0x40a3b7e7)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogGetPinnedSavedDialogs) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetPinnedSavedDialogs) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x40a3b7e7:

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogToggleSavedDialogPin
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogToggleSavedDialogPin) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x44f317d9:
		x.UInt(0x44f317d9)

		// no flags

		x.Long(m.GetUserId())
		m.GetPeer().Encode(x, layer)
		m.GetPinned().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogToggleSavedDialogPin) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogToggleSavedDialogPin) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x44f317d9:

		// not has flags

		m.UserId = dBuf.Long()

		m2 := &mtproto.PeerUtil{}
		m2.Decode(dBuf)
		m.Peer = m2

		m3 := &mtproto.Bool{}
		m3.Decode(dBuf)
		m.Pinned = m3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogReorderPinnedSavedDialogs
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogReorderPinnedSavedDialogs) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xd85ccbd2:
		x.UInt(0xd85ccbd2)

		// no flags

		x.Long(m.GetUserId())
		m.GetForce().Encode(x, layer)

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetOrder())))
		for _, v := range m.GetOrder() {
			v.Encode(x, layer)
		}

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogReorderPinnedSavedDialogs) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogReorderPinnedSavedDialogs) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xd85ccbd2:

		// not has flags

		m.UserId = dBuf.Long()

		m2 := &mtproto.Bool{}
		m2.Decode(dBuf)
		m.Force = m2

		c3 := dBuf.Int()
		if c3 != int32(mtproto.CRC32_vector) {
			// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
			return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
		}
		l3 := dBuf.Int()
		v3 := make([]*mtproto.PeerUtil, l3)
		for i := int32(0); i < l3; i++ {
			v3[i] = &mtproto.PeerUtil{}
			v3[i].Decode(dBuf)
		}
		m.Order = v3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogGetDialogFilter
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetDialogFilter) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xf388061c:
		x.UInt(0xf388061c)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogGetDialogFilter) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetDialogFilter) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf388061c:

		// not has flags

		m.UserId = dBuf.Long()
		m.Id = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogGetDialogFilterBySlug
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetDialogFilterBySlug) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x4e457fef:
		x.UInt(0x4e457fef)

		// no flags

		x.Long(m.GetUserId())
		x.String(m.GetSlug())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogGetDialogFilterBySlug) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetDialogFilterBySlug) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x4e457fef:

		// not has flags

		m.UserId = dBuf.Long()
		m.Slug = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogCreateDialogFilter
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogCreateDialogFilter) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xc6cb636f:
		x.UInt(0xc6cb636f)

		// no flags

		x.Long(m.GetUserId())
		m.GetDialogFilter().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogCreateDialogFilter) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogCreateDialogFilter) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc6cb636f:

		// not has flags

		m.UserId = dBuf.Long()

		m2 := &DialogFilterExt{}
		m2.Decode(dBuf)
		m.DialogFilter = m2

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLDialogUpdateUnreadCount
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogUpdateUnreadCount) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x2bac334d:
		x.UInt(0x2bac334d)

		// set flags
		var flags uint32 = 0

		if m.GetUnreadCount() != nil {
			flags |= 1 << 0
		}
		if m.GetUnreadMentionsCount() != nil {
			flags |= 1 << 1
		}
		if m.GetUnreadReactionsCount() != nil {
			flags |= 1 << 2
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		if m.GetUnreadCount() != nil {
			x.Int(m.GetUnreadCount().Value)
		}

		if m.GetUnreadMentionsCount() != nil {
			x.Int(m.GetUnreadMentionsCount().Value)
		}

		if m.GetUnreadReactionsCount() != nil {
			x.Int(m.GetUnreadReactionsCount().Value)
		}

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLDialogUpdateUnreadCount) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogUpdateUnreadCount) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x2bac334d:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		if (flags & (1 << 0)) != 0 {
			m.UnreadCount = &wrapperspb.Int32Value{Value: dBuf.Int()}
		}

		if (flags & (1 << 1)) != 0 {
			m.UnreadMentionsCount = &wrapperspb.Int32Value{Value: dBuf.Int()}
		}

		if (flags & (1 << 2)) != 0 {
			m.UnreadReactionsCount = &wrapperspb.Int32Value{Value: dBuf.Int()}
		}

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// Vector_PeerWithDraftMessage
// /////////////////////////////////////////////////////////////////////////////
func (m *Vector_PeerWithDraftMessage) Encode(x *mtproto.EncodeBuf, layer int32) error {
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		v.Encode(x, layer)
	}

	return nil
}

func (m *Vector_PeerWithDraftMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*PeerWithDraftMessage, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(PeerWithDraftMessage)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_PeerWithDraftMessage) CalcByteSize(layer int32) int {
	return 0
}

// Vector_DialogPeer
// /////////////////////////////////////////////////////////////////////////////
func (m *Vector_DialogPeer) Encode(x *mtproto.EncodeBuf, layer int32) error {
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		v.Encode(x, layer)
	}

	return nil
}

func (m *Vector_DialogPeer) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*mtproto.DialogPeer, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(mtproto.DialogPeer)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_DialogPeer) CalcByteSize(layer int32) int {
	return 0
}

// Vector_DialogExt
// /////////////////////////////////////////////////////////////////////////////
func (m *Vector_DialogExt) Encode(x *mtproto.EncodeBuf, layer int32) error {
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		v.Encode(x, layer)
	}

	return nil
}

func (m *Vector_DialogExt) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*DialogExt, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(DialogExt)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_DialogExt) CalcByteSize(layer int32) int {
	return 0
}

// Vector_DialogFilterExt
// /////////////////////////////////////////////////////////////////////////////
func (m *Vector_DialogFilterExt) Encode(x *mtproto.EncodeBuf, layer int32) error {
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		v.Encode(x, layer)
	}

	return nil
}

func (m *Vector_DialogFilterExt) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*DialogFilterExt, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(DialogFilterExt)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_DialogFilterExt) CalcByteSize(layer int32) int {
	return 0
}

// Vector_DialogPinnedExt
// /////////////////////////////////////////////////////////////////////////////
func (m *Vector_DialogPinnedExt) Encode(x *mtproto.EncodeBuf, layer int32) error {
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		v.Encode(x, layer)
	}

	return nil
}

func (m *Vector_DialogPinnedExt) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*DialogPinnedExt, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(DialogPinnedExt)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_DialogPinnedExt) CalcByteSize(layer int32) int {
	return 0
}

// Vector_Long
// /////////////////////////////////////////////////////////////////////////////
func (m *Vector_Long) Encode(x *mtproto.EncodeBuf, layer int32) error {
	x.VectorLong(m.Datas)

	return nil
}

func (m *Vector_Long) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Datas = dBuf.VectorLong()

	return dBuf.GetError()
}

func (m *Vector_Long) CalcByteSize(layer int32) int {
	return 0
}
