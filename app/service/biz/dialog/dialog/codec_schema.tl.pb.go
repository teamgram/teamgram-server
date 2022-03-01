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

package dialog

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
	-155335502: func() mtproto.TLObject { // 0xf6bdc4b2
		o := MakeTLUpdateDraftMessage(nil)
		o.Data2.Constructor = -155335502
		return o
	},
	722420824: func() mtproto.TLObject { // 0x2b0f4458
		o := MakeTLDialogExt(nil)
		o.Data2.Constructor = 722420824
		return o
	},
	245834284: func() mtproto.TLObject { // 0xea7222c
		o := MakeTLDialogPinnedExt(nil)
		o.Data2.Constructor = 245834284
		return o
	},
	-1891683854: func() mtproto.TLObject { // 0x8f3f31f2
		o := MakeTLDialogFilterExt(nil)
		o.Data2.Constructor = -1891683854
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
	-317723281: func() mtproto.TLObject { // 0xed0fed6f
		return &TLDialogInsertOrUpdateDialog{
			Constructor: -317723281,
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
// PeerWithDraftMessage <--
//  + TL_UpdateDraftMessage
//

func (m *PeerWithDraftMessage) Encode(layer int32) []byte {
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
	case Predicate_updateDraftMessage:
		t := m.To_UpdateDraftMessage()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
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

func (m *PeerWithDraftMessage) DebugString() string {
	switch m.PredicateName {
	case Predicate_updateDraftMessage:
		t := m.To_UpdateDraftMessage()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_UpdateDraftMessage
// updateDraftMessage peer:Peer draft:DraftMessage = PeerWithDraftMessage;
func (m *PeerWithDraftMessage) To_UpdateDraftMessage() *TLUpdateDraftMessage {
	m.PredicateName = Predicate_updateDraftMessage
	return &TLUpdateDraftMessage{
		Data2: m,
	}
}

// MakeTLUpdateDraftMessage
// updateDraftMessage peer:Peer draft:DraftMessage = PeerWithDraftMessage;
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

func (m *TLUpdateDraftMessage) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xf6bdc4b2: func() []byte {
			// updateDraftMessage peer:Peer draft:DraftMessage = PeerWithDraftMessage;
			x.UInt(0xf6bdc4b2)

			x.Bytes(m.GetPeer().Encode(layer))
			x.Bytes(m.GetDraft().Encode(layer))
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_updateDraftMessage, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_updateDraftMessage, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLUpdateDraftMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUpdateDraftMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xf6bdc4b2: func() error {
			// updateDraftMessage peer:Peer draft:DraftMessage = PeerWithDraftMessage;

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

func (m *TLUpdateDraftMessage) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// DialogExt <--
//  + TL_DialogExt
//

func (m *DialogExt) Encode(layer int32) []byte {
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
	case Predicate_dialogExt:
		t := m.To_DialogExt()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *DialogExt) CalcByteSize(layer int32) int {
	return 0
}

func (m *DialogExt) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x2b0f4458:
		m2 := MakeTLDialogExt(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *DialogExt) DebugString() string {
	switch m.PredicateName {
	case Predicate_dialogExt:
		t := m.To_DialogExt()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_DialogExt
// dialogExt order:long dialog:Dialog available_min_id:int date:long = DialogExt;
func (m *DialogExt) To_DialogExt() *TLDialogExt {
	m.PredicateName = Predicate_dialogExt
	return &TLDialogExt{
		Data2: m,
	}
}

// MakeTLDialogExt
// dialogExt order:long dialog:Dialog available_min_id:int date:long = DialogExt;
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

func (m *TLDialogExt) GetPredicateName() string {
	return Predicate_dialogExt
}

func (m *TLDialogExt) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x2b0f4458: func() []byte {
			// dialogExt order:long dialog:Dialog available_min_id:int date:long = DialogExt;
			x.UInt(0x2b0f4458)

			x.Long(m.GetOrder())
			x.Bytes(m.GetDialog().Encode(layer))
			x.Int(m.GetAvailableMinId())
			x.Long(m.GetDate())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_dialogExt, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_dialogExt, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLDialogExt) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogExt) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x2b0f4458: func() error {
			// dialogExt order:long dialog:Dialog available_min_id:int date:long = DialogExt;
			m.SetOrder(dBuf.Long())

			m1 := &mtproto.Dialog{}
			m1.Decode(dBuf)
			m.SetDialog(m1)

			m.SetAvailableMinId(dBuf.Int())
			m.SetDate(dBuf.Long())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLDialogExt) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// DialogPinnedExt <--
//  + TL_DialogPinnedExt
//

func (m *DialogPinnedExt) Encode(layer int32) []byte {
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
	case Predicate_dialogPinnedExt:
		t := m.To_DialogPinnedExt()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
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

func (m *DialogPinnedExt) DebugString() string {
	switch m.PredicateName {
	case Predicate_dialogPinnedExt:
		t := m.To_DialogPinnedExt()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_DialogPinnedExt
// dialogPinnedExt order:long peer_type:int peer_id:long = DialogPinnedExt;
func (m *DialogPinnedExt) To_DialogPinnedExt() *TLDialogPinnedExt {
	m.PredicateName = Predicate_dialogPinnedExt
	return &TLDialogPinnedExt{
		Data2: m,
	}
}

// MakeTLDialogPinnedExt
// dialogPinnedExt order:long peer_type:int peer_id:long = DialogPinnedExt;
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

func (m *TLDialogPinnedExt) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xea7222c: func() []byte {
			// dialogPinnedExt order:long peer_type:int peer_id:long = DialogPinnedExt;
			x.UInt(0xea7222c)

			x.Long(m.GetOrder())
			x.Int(m.GetPeerType())
			x.Long(m.GetPeerId())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_dialogPinnedExt, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_dialogPinnedExt, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLDialogPinnedExt) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogPinnedExt) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xea7222c: func() error {
			// dialogPinnedExt order:long peer_type:int peer_id:long = DialogPinnedExt;
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

func (m *TLDialogPinnedExt) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// DialogFilterExt <--
//  + TL_DialogFilterExt
//

func (m *DialogFilterExt) Encode(layer int32) []byte {
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
	case Predicate_dialogFilterExt:
		t := m.To_DialogFilterExt()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *DialogFilterExt) CalcByteSize(layer int32) int {
	return 0
}

func (m *DialogFilterExt) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x8f3f31f2:
		m2 := MakeTLDialogFilterExt(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *DialogFilterExt) DebugString() string {
	switch m.PredicateName {
	case Predicate_dialogFilterExt:
		t := m.To_DialogFilterExt()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_DialogFilterExt
// dialogFilterExt id:int dialog_filter:DialogFilter order:long = DialogFilterExt;
func (m *DialogFilterExt) To_DialogFilterExt() *TLDialogFilterExt {
	m.PredicateName = Predicate_dialogFilterExt
	return &TLDialogFilterExt{
		Data2: m,
	}
}

// MakeTLDialogFilterExt
// dialogFilterExt id:int dialog_filter:DialogFilter order:long = DialogFilterExt;
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

func (m *TLDialogFilterExt) SetId(v int32) { m.Data2.Id = v }
func (m *TLDialogFilterExt) GetId() int32  { return m.Data2.Id }

func (m *TLDialogFilterExt) SetDialogFilter(v *mtproto.DialogFilter) { m.Data2.DialogFilter = v }
func (m *TLDialogFilterExt) GetDialogFilter() *mtproto.DialogFilter  { return m.Data2.DialogFilter }

func (m *TLDialogFilterExt) SetOrder(v int64) { m.Data2.Order = v }
func (m *TLDialogFilterExt) GetOrder() int64  { return m.Data2.Order }

func (m *TLDialogFilterExt) GetPredicateName() string {
	return Predicate_dialogFilterExt
}

func (m *TLDialogFilterExt) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x8f3f31f2: func() []byte {
			// dialogFilterExt id:int dialog_filter:DialogFilter order:long = DialogFilterExt;
			x.UInt(0x8f3f31f2)

			x.Int(m.GetId())
			x.Bytes(m.GetDialogFilter().Encode(layer))
			x.Long(m.GetOrder())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_dialogFilterExt, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_dialogFilterExt, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLDialogFilterExt) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogFilterExt) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x8f3f31f2: func() error {
			// dialogFilterExt id:int dialog_filter:DialogFilter order:long = DialogFilterExt;
			m.SetId(dBuf.Int())

			m1 := &mtproto.DialogFilter{}
			m1.Decode(dBuf)
			m.SetDialogFilter(m1)

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

func (m *TLDialogFilterExt) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// TLDialogSaveDraftMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogSaveDraftMessage) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_saveDraftMessage))

	switch uint32(m.Constructor) {
	case 0x4ecad99a:
		// dialog.saveDraftMessage user_id:long peer_type:int peer_id:long message:DraftMessage = Bool;
		x.UInt(0x4ecad99a)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Bytes(m.GetMessage().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogSaveDraftMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogSaveDraftMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x4ecad99a:
		// dialog.saveDraftMessage user_id:long peer_type:int peer_id:long message:DraftMessage = Bool;

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

func (m *TLDialogSaveDraftMessage) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogClearDraftMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogClearDraftMessage) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_clearDraftMessage))

	switch uint32(m.Constructor) {
	case 0xfb70b29a:
		// dialog.clearDraftMessage user_id:long peer_type:int peer_id:long = Bool;
		x.UInt(0xfb70b29a)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogClearDraftMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogClearDraftMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xfb70b29a:
		// dialog.clearDraftMessage user_id:long peer_type:int peer_id:long = Bool;

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

func (m *TLDialogClearDraftMessage) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogGetAllDrafts
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetAllDrafts) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_getAllDrafts))

	switch uint32(m.Constructor) {
	case 0xacde4fe6:
		// dialog.getAllDrafts user_id:long = Vector<PeerWithDraftMessage>;
		x.UInt(0xacde4fe6)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogGetAllDrafts) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetAllDrafts) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xacde4fe6:
		// dialog.getAllDrafts user_id:long = Vector<PeerWithDraftMessage>;

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDialogGetAllDrafts) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogClearAllDrafts
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogClearAllDrafts) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_clearAllDrafts))

	switch uint32(m.Constructor) {
	case 0x41b890fc:
		// dialog.clearAllDrafts user_id:long = Vector<PeerWithDraftMessage>;
		x.UInt(0x41b890fc)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogClearAllDrafts) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogClearAllDrafts) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x41b890fc:
		// dialog.clearAllDrafts user_id:long = Vector<PeerWithDraftMessage>;

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDialogClearAllDrafts) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogMarkDialogUnread
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogMarkDialogUnread) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_markDialogUnread))

	switch uint32(m.Constructor) {
	case 0x4532910e:
		// dialog.markDialogUnread user_id:long peer_type:int peer_id:long unread_mark:Bool = Bool;
		x.UInt(0x4532910e)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Bytes(m.GetUnreadMark().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogMarkDialogUnread) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogMarkDialogUnread) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x4532910e:
		// dialog.markDialogUnread user_id:long peer_type:int peer_id:long unread_mark:Bool = Bool;

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

func (m *TLDialogMarkDialogUnread) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogToggleDialogPin
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogToggleDialogPin) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_toggleDialogPin))

	switch uint32(m.Constructor) {
	case 0x867ee52f:
		// dialog.toggleDialogPin user_id:long peer_type:int peer_id:long pinned:Bool = Int32;
		x.UInt(0x867ee52f)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Bytes(m.GetPinned().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogToggleDialogPin) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogToggleDialogPin) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x867ee52f:
		// dialog.toggleDialogPin user_id:long peer_type:int peer_id:long pinned:Bool = Int32;

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

func (m *TLDialogToggleDialogPin) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogGetDialogUnreadMarkList
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetDialogUnreadMarkList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_getDialogUnreadMarkList))

	switch uint32(m.Constructor) {
	case 0xcabc38f4:
		// dialog.getDialogUnreadMarkList user_id:long = Vector<DialogPeer>;
		x.UInt(0xcabc38f4)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogGetDialogUnreadMarkList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetDialogUnreadMarkList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xcabc38f4:
		// dialog.getDialogUnreadMarkList user_id:long = Vector<DialogPeer>;

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDialogGetDialogUnreadMarkList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogGetDialogsByOffsetDate
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetDialogsByOffsetDate) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_getDialogsByOffsetDate))

	switch uint32(m.Constructor) {
	case 0x9d7e8604:
		// dialog.getDialogsByOffsetDate user_id:long exclude_pinned:Bool offset_date:int limit:int = Vector<DialogExt>;
		x.UInt(0x9d7e8604)

		// no flags

		x.Long(m.GetUserId())
		x.Bytes(m.GetExcludePinned().Encode(layer))
		x.Int(m.GetOffsetDate())
		x.Int(m.GetLimit())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogGetDialogsByOffsetDate) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetDialogsByOffsetDate) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x9d7e8604:
		// dialog.getDialogsByOffsetDate user_id:long exclude_pinned:Bool offset_date:int limit:int = Vector<DialogExt>;

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

func (m *TLDialogGetDialogsByOffsetDate) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogGetDialogs
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetDialogs) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_getDialogs))

	switch uint32(m.Constructor) {
	case 0x860b1e16:
		// dialog.getDialogs user_id:long exclude_pinned:Bool folder_id:int = Vector<DialogExt>;
		x.UInt(0x860b1e16)

		// no flags

		x.Long(m.GetUserId())
		x.Bytes(m.GetExcludePinned().Encode(layer))
		x.Int(m.GetFolderId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogGetDialogs) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetDialogs) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x860b1e16:
		// dialog.getDialogs user_id:long exclude_pinned:Bool folder_id:int = Vector<DialogExt>;

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

func (m *TLDialogGetDialogs) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogGetDialogsByIdList
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetDialogsByIdList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_getDialogsByIdList))

	switch uint32(m.Constructor) {
	case 0xad258871:
		// dialog.getDialogsByIdList user_id:long id_list:Vector<long> = Vector<DialogExt>;
		x.UInt(0xad258871)

		// no flags

		x.Long(m.GetUserId())

		x.VectorLong(m.GetIdList())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogGetDialogsByIdList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetDialogsByIdList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xad258871:
		// dialog.getDialogsByIdList user_id:long id_list:Vector<long> = Vector<DialogExt>;

		// not has flags

		m.UserId = dBuf.Long()

		m.IdList = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDialogGetDialogsByIdList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogGetDialogsCount
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetDialogsCount) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_getDialogsCount))

	switch uint32(m.Constructor) {
	case 0xe039b465:
		// dialog.getDialogsCount user_id:long exclude_pinned:Bool folder_id:int = Int32;
		x.UInt(0xe039b465)

		// no flags

		x.Long(m.GetUserId())
		x.Bytes(m.GetExcludePinned().Encode(layer))
		x.Int(m.GetFolderId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogGetDialogsCount) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetDialogsCount) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe039b465:
		// dialog.getDialogsCount user_id:long exclude_pinned:Bool folder_id:int = Int32;

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

func (m *TLDialogGetDialogsCount) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogGetPinnedDialogs
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetPinnedDialogs) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_getPinnedDialogs))

	switch uint32(m.Constructor) {
	case 0xa8c21bb5:
		// dialog.getPinnedDialogs  user_id:long folder_id:int = Vector<DialogExt>;
		x.UInt(0xa8c21bb5)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetFolderId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogGetPinnedDialogs) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetPinnedDialogs) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xa8c21bb5:
		// dialog.getPinnedDialogs  user_id:long folder_id:int = Vector<DialogExt>;

		// not has flags

		m.UserId = dBuf.Long()
		m.FolderId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDialogGetPinnedDialogs) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogReorderPinnedDialogs
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogReorderPinnedDialogs) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_reorderPinnedDialogs))

	switch uint32(m.Constructor) {
	case 0xfee33567:
		// dialog.reorderPinnedDialogs user_id:long force:Bool folder_id:int id_list:Vector<long> = Bool;
		x.UInt(0xfee33567)

		// no flags

		x.Long(m.GetUserId())
		x.Bytes(m.GetForce().Encode(layer))
		x.Int(m.GetFolderId())

		x.VectorLong(m.GetIdList())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogReorderPinnedDialogs) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogReorderPinnedDialogs) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xfee33567:
		// dialog.reorderPinnedDialogs user_id:long force:Bool folder_id:int id_list:Vector<long> = Bool;

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

func (m *TLDialogReorderPinnedDialogs) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogGetDialogById
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetDialogById) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_getDialogById))

	switch uint32(m.Constructor) {
	case 0xa15f3bf5:
		// dialog.getDialogById user_id:long peer_type:int peer_id:long = DialogExt;
		x.UInt(0xa15f3bf5)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogGetDialogById) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetDialogById) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xa15f3bf5:
		// dialog.getDialogById user_id:long peer_type:int peer_id:long = DialogExt;

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

func (m *TLDialogGetDialogById) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogGetTopMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetTopMessage) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_getTopMessage))

	switch uint32(m.Constructor) {
	case 0xfa7db272:
		// dialog.getTopMessage user_id:long peer_type:int peer_id:long = Int32;
		x.UInt(0xfa7db272)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogGetTopMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetTopMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xfa7db272:
		// dialog.getTopMessage user_id:long peer_type:int peer_id:long = Int32;

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

func (m *TLDialogGetTopMessage) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogUpdateReadInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogUpdateReadInbox) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_updateReadInbox))

	switch uint32(m.Constructor) {
	case 0x1d27f8b8:
		// dialog.updateReadInbox user_id:long peer_type:int peer_id:long read_inbox_id:int = Bool;
		x.UInt(0x1d27f8b8)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetReadInboxId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogUpdateReadInbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogUpdateReadInbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x1d27f8b8:
		// dialog.updateReadInbox user_id:long peer_type:int peer_id:long read_inbox_id:int = Bool;

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

func (m *TLDialogUpdateReadInbox) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogUpdateReadOutbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogUpdateReadOutbox) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_updateReadOutbox))

	switch uint32(m.Constructor) {
	case 0x5870fd7e:
		// dialog.updateReadOutbox user_id:long peer_type:int peer_id:long read_outbox_id:int = Bool;
		x.UInt(0x5870fd7e)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetReadOutboxId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogUpdateReadOutbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogUpdateReadOutbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x5870fd7e:
		// dialog.updateReadOutbox user_id:long peer_type:int peer_id:long read_outbox_id:int = Bool;

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

func (m *TLDialogUpdateReadOutbox) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogInsertOrUpdateDialog
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogInsertOrUpdateDialog) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_insertOrUpdateDialog))

	switch uint32(m.Constructor) {
	case 0xed0fed6f:
		// dialog.insertOrUpdateDialog flags:# user_id:long peer_type:int peer_id:long top_message:flags.0?int read_outbox_max_id:flags.1?int read_inbox_max_id:flags.2?int unread_count:flags.3?int unread_mark:flags.4?true = Bool;
		x.UInt(0xed0fed6f)

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

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogInsertOrUpdateDialog) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogInsertOrUpdateDialog) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xed0fed6f:
		// dialog.insertOrUpdateDialog flags:# user_id:long peer_type:int peer_id:long top_message:flags.0?int read_outbox_max_id:flags.1?int read_inbox_max_id:flags.2?int unread_count:flags.3?int unread_mark:flags.4?true = Bool;

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		if (flags & (1 << 0)) != 0 {
			m.TopMessage = &types.Int32Value{Value: dBuf.Int()}
		}

		if (flags & (1 << 1)) != 0 {
			m.ReadOutboxMaxId = &types.Int32Value{Value: dBuf.Int()}
		}

		if (flags & (1 << 2)) != 0 {
			m.ReadInboxMaxId = &types.Int32Value{Value: dBuf.Int()}
		}

		if (flags & (1 << 3)) != 0 {
			m.UnreadCount = &types.Int32Value{Value: dBuf.Int()}
		}

		if (flags & (1 << 4)) != 0 {
			m.UnreadMark = true
		}
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDialogInsertOrUpdateDialog) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogDeleteDialog
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogDeleteDialog) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_deleteDialog))

	switch uint32(m.Constructor) {
	case 0x1b31de3:
		// dialog.deleteDialog user_id:long peer_type:int peer_id:long = Bool;
		x.UInt(0x1b31de3)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogDeleteDialog) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogDeleteDialog) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x1b31de3:
		// dialog.deleteDialog user_id:long peer_type:int peer_id:long = Bool;

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

func (m *TLDialogDeleteDialog) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogGetUserPinnedMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetUserPinnedMessage) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_getUserPinnedMessage))

	switch uint32(m.Constructor) {
	case 0x8f9bc2b1:
		// dialog.getUserPinnedMessage user_id:long peer_type:int peer_id:long = Int32;
		x.UInt(0x8f9bc2b1)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogGetUserPinnedMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetUserPinnedMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x8f9bc2b1:
		// dialog.getUserPinnedMessage user_id:long peer_type:int peer_id:long = Int32;

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

func (m *TLDialogGetUserPinnedMessage) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogUpdateUserPinnedMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogUpdateUserPinnedMessage) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_updateUserPinnedMessage))

	switch uint32(m.Constructor) {
	case 0x1622f22a:
		// dialog.updateUserPinnedMessage user_id:long peer_type:int peer_id:long pinned_msg_id:int = Bool;
		x.UInt(0x1622f22a)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetPinnedMsgId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogUpdateUserPinnedMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogUpdateUserPinnedMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x1622f22a:
		// dialog.updateUserPinnedMessage user_id:long peer_type:int peer_id:long pinned_msg_id:int = Bool;

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

func (m *TLDialogUpdateUserPinnedMessage) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogInsertOrUpdateDialogFilter
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogInsertOrUpdateDialogFilter) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_insertOrUpdateDialogFilter))

	switch uint32(m.Constructor) {
	case 0xaa8a384:
		// dialog.insertOrUpdateDialogFilter user_id:long id:int dialog_filter:DialogFilter = Bool;
		x.UInt(0xaa8a384)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetId())
		x.Bytes(m.GetDialogFilter().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogInsertOrUpdateDialogFilter) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogInsertOrUpdateDialogFilter) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xaa8a384:
		// dialog.insertOrUpdateDialogFilter user_id:long id:int dialog_filter:DialogFilter = Bool;

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

func (m *TLDialogInsertOrUpdateDialogFilter) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogDeleteDialogFilter
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogDeleteDialogFilter) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_deleteDialogFilter))

	switch uint32(m.Constructor) {
	case 0x1dd3e97:
		// dialog.deleteDialogFilter user_id:long id:int = Bool;
		x.UInt(0x1dd3e97)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogDeleteDialogFilter) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogDeleteDialogFilter) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x1dd3e97:
		// dialog.deleteDialogFilter user_id:long id:int = Bool;

		// not has flags

		m.UserId = dBuf.Long()
		m.Id = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDialogDeleteDialogFilter) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogUpdateDialogFiltersOrder
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogUpdateDialogFiltersOrder) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_updateDialogFiltersOrder))

	switch uint32(m.Constructor) {
	case 0xb13c0b3f:
		// dialog.updateDialogFiltersOrder user_id:long order:Vector<int> = Bool;
		x.UInt(0xb13c0b3f)

		// no flags

		x.Long(m.GetUserId())

		x.VectorInt(m.GetOrder())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogUpdateDialogFiltersOrder) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogUpdateDialogFiltersOrder) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb13c0b3f:
		// dialog.updateDialogFiltersOrder user_id:long order:Vector<int> = Bool;

		// not has flags

		m.UserId = dBuf.Long()

		m.Order = dBuf.VectorInt()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDialogUpdateDialogFiltersOrder) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogGetDialogFilters
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetDialogFilters) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_getDialogFilters))

	switch uint32(m.Constructor) {
	case 0x6c676c3c:
		// dialog.getDialogFilters user_id:long = Vector<DialogFilterExt>;
		x.UInt(0x6c676c3c)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogGetDialogFilters) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetDialogFilters) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x6c676c3c:
		// dialog.getDialogFilters user_id:long = Vector<DialogFilterExt>;

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDialogGetDialogFilters) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogGetDialogFolder
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetDialogFolder) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_getDialogFolder))

	switch uint32(m.Constructor) {
	case 0x411b8eb5:
		// dialog.getDialogFolder user_id:long folder_id:int = Vector<DialogExt>;
		x.UInt(0x411b8eb5)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetFolderId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogGetDialogFolder) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetDialogFolder) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x411b8eb5:
		// dialog.getDialogFolder user_id:long folder_id:int = Vector<DialogExt>;

		// not has flags

		m.UserId = dBuf.Long()
		m.FolderId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLDialogGetDialogFolder) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogEditPeerFolders
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogEditPeerFolders) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_editPeerFolders))

	switch uint32(m.Constructor) {
	case 0x2446869a:
		// dialog.editPeerFolders user_id:long peer_dialog_list:Vector<long> folder_id:int = Vector<DialogPinnedExt>;
		x.UInt(0x2446869a)

		// no flags

		x.Long(m.GetUserId())

		x.VectorLong(m.GetPeerDialogList())

		x.Int(m.GetFolderId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogEditPeerFolders) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogEditPeerFolders) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x2446869a:
		// dialog.editPeerFolders user_id:long peer_dialog_list:Vector<long> folder_id:int = Vector<DialogPinnedExt>;

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

func (m *TLDialogEditPeerFolders) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLDialogGetChannelMessageReadParticipants
///////////////////////////////////////////////////////////////////////////////

func (m *TLDialogGetChannelMessageReadParticipants) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_dialog_getChannelMessageReadParticipants))

	switch uint32(m.Constructor) {
	case 0x28bd4d3b:
		// dialog.getChannelMessageReadParticipants user_id:long channel_id:long msg_id:int = Vector<long>;
		x.UInt(0x28bd4d3b)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetChannelId())
		x.Int(m.GetMsgId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLDialogGetChannelMessageReadParticipants) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLDialogGetChannelMessageReadParticipants) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x28bd4d3b:
		// dialog.getChannelMessageReadParticipants user_id:long channel_id:long msg_id:int = Vector<long>;

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

func (m *TLDialogGetChannelMessageReadParticipants) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// Vector_PeerWithDraftMessage
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_PeerWithDraftMessage) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
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

func (m *Vector_PeerWithDraftMessage) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// Vector_DialogPeer
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_DialogPeer) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
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

func (m *Vector_DialogPeer) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// Vector_DialogExt
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_DialogExt) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
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

func (m *Vector_DialogExt) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// Vector_DialogFilterExt
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_DialogFilterExt) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
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

func (m *Vector_DialogFilterExt) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// Vector_DialogPinnedExt
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_DialogPinnedExt) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
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

func (m *Vector_DialogPinnedExt) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// Vector_Long
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_Long) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.VectorLong(m.Datas)

	return x.GetBuf()
}

func (m *Vector_Long) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Datas = dBuf.VectorLong()

	return dBuf.GetError()
}

func (m *Vector_Long) CalcByteSize(layer int32) int {
	return 0
}

func (m *Vector_Long) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}
