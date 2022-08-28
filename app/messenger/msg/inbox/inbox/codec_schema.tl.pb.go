/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

// ConstructorList
// RequestList

package inbox

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
	1002286548: func() mtproto.TLObject { // 0x3bbdadd4
		o := MakeTLInboxMessageData(nil)
		o.Data2.Constructor = 1002286548
		return o
	},
	-963460705: func() mtproto.TLObject { // 0xc692c19f
		o := MakeTLInboxMessageId(nil)
		o.Data2.Constructor = -963460705
		return o
	},

	// Method
	-208741709: func() mtproto.TLObject { // 0xf38edab3
		return &TLInboxSendUserMessageToInbox{
			Constructor: -208741709,
		}
	},
	-1760197438: func() mtproto.TLObject { // 0x971584c2
		return &TLInboxSendChatMessageToInbox{
			Constructor: -1760197438,
		}
	},
	2050486614: func() mtproto.TLObject { // 0x7a37f156
		return &TLInboxSendChannelMessageToInbox{
			Constructor: 2050486614,
		}
	},
	-1782288007: func() mtproto.TLObject { // 0x95c47179
		return &TLInboxSendUserMultiMessageToInbox{
			Constructor: -1782288007,
		}
	},
	-694455924: func() mtproto.TLObject { // 0xd69b718c
		return &TLInboxSendChatMultiMessageToInbox{
			Constructor: -694455924,
		}
	},
	999414081: func() mtproto.TLObject { // 0x3b91d941
		return &TLInboxSendChannelMultiMessageToInbox{
			Constructor: 999414081,
		}
	},
	1559967656: func() mtproto.TLObject { // 0x5cfb37a8
		return &TLInboxEditUserMessageToInbox{
			Constructor: 1559967656,
		}
	},
	2031122959: func() mtproto.TLObject { // 0x79107a0f
		return &TLInboxEditChatMessageToInbox{
			Constructor: 2031122959,
		}
	},
	-231965363: func() mtproto.TLObject { // 0xf22c7d4d
		return &TLInboxEditChannelMessageToInbox{
			Constructor: -231965363,
		}
	},
	-2061734348: func() mtproto.TLObject { // 0x851c6e34
		return &TLInboxDeleteMessagesToInbox{
			Constructor: -2061734348,
		}
	},
	295332038: func() mtproto.TLObject { // 0x119a68c6
		return &TLInboxDeleteChannelMessagesToInbox{
			Constructor: 295332038,
		}
	},
	336232792: func() mtproto.TLObject { // 0x140a8158
		return &TLInboxDeleteUserHistoryToInbox{
			Constructor: 336232792,
		}
	},
	-659905022: func() mtproto.TLObject { // 0xd8aaa602
		return &TLInboxDeleteChatHistoryToInbox{
			Constructor: -659905022,
		}
	},
	364970827: func() mtproto.TLObject { // 0x15c1034b
		return &TLInboxReadUserMediaUnreadToInbox{
			Constructor: 364970827,
		}
	},
	1430347220: func() mtproto.TLObject { // 0x55415dd4
		return &TLInboxReadChatMediaUnreadToInbox{
			Constructor: 1430347220,
		}
	},
	-1476338378: func() mtproto.TLObject { // 0xa800dd36
		return &TLInboxReadChannelMediaUnreadToInbox{
			Constructor: -1476338378,
		}
	},
	-1010283296: func() mtproto.TLObject { // 0xc3c84ce0
		return &TLInboxUpdateHistoryReaded{
			Constructor: -1010283296,
		}
	},
	-1452528908: func() mtproto.TLObject { // 0xa96c2af4
		return &TLInboxUpdatePinnedMessage{
			Constructor: -1452528908,
		}
	},
	589079137: func() mtproto.TLObject { // 0x231ca261
		return &TLInboxUnpinAllMessages{
			Constructor: 589079137,
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
// InboxMessageData <--
//  + TL_InboxMessageData
//

func (m *InboxMessageData) Encode(layer int32) []byte {
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
	case Predicate_inboxMessageData:
		t := m.To_InboxMessageData()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *InboxMessageData) CalcByteSize(layer int32) int {
	return 0
}

func (m *InboxMessageData) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x3bbdadd4:
		m2 := MakeTLInboxMessageData(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *InboxMessageData) DebugString() string {
	switch m.PredicateName {
	case Predicate_inboxMessageData:
		t := m.To_InboxMessageData()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_InboxMessageData
func (m *InboxMessageData) To_InboxMessageData() *TLInboxMessageData {
	m.PredicateName = Predicate_inboxMessageData
	return &TLInboxMessageData{
		Data2: m,
	}
}

// MakeTLInboxMessageData
func MakeTLInboxMessageData(data2 *InboxMessageData) *TLInboxMessageData {
	if data2 == nil {
		return &TLInboxMessageData{Data2: &InboxMessageData{
			PredicateName: Predicate_inboxMessageData,
		}}
	} else {
		data2.PredicateName = Predicate_inboxMessageData
		return &TLInboxMessageData{Data2: data2}
	}
}

func (m *TLInboxMessageData) To_InboxMessageData() *InboxMessageData {
	m.Data2.PredicateName = Predicate_inboxMessageData
	return m.Data2
}

func (m *TLInboxMessageData) SetRandomId(v int64) { m.Data2.RandomId = v }
func (m *TLInboxMessageData) GetRandomId() int64  { return m.Data2.RandomId }

func (m *TLInboxMessageData) SetDialogMessageId(v int64) { m.Data2.DialogMessageId = v }
func (m *TLInboxMessageData) GetDialogMessageId() int64  { return m.Data2.DialogMessageId }

func (m *TLInboxMessageData) SetMessage(v *mtproto.Message) { m.Data2.Message = v }
func (m *TLInboxMessageData) GetMessage() *mtproto.Message  { return m.Data2.Message }

func (m *TLInboxMessageData) GetPredicateName() string {
	return Predicate_inboxMessageData
}

func (m *TLInboxMessageData) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x3bbdadd4: func() []byte {
			x.UInt(0x3bbdadd4)

			x.Long(m.GetRandomId())
			x.Long(m.GetDialogMessageId())
			x.Bytes(m.GetMessage().Encode(layer))
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_inboxMessageData, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_inboxMessageData, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLInboxMessageData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxMessageData) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x3bbdadd4: func() error {
			m.SetRandomId(dBuf.Long())
			m.SetDialogMessageId(dBuf.Long())

			m2 := &mtproto.Message{}
			m2.Decode(dBuf)
			m.SetMessage(m2)

			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLInboxMessageData) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// InboxMessageId <--
//  + TL_InboxMessageId
//

func (m *InboxMessageId) Encode(layer int32) []byte {
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
	case Predicate_inboxMessageId:
		t := m.To_InboxMessageId()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *InboxMessageId) CalcByteSize(layer int32) int {
	return 0
}

func (m *InboxMessageId) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xc692c19f:
		m2 := MakeTLInboxMessageId(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *InboxMessageId) DebugString() string {
	switch m.PredicateName {
	case Predicate_inboxMessageId:
		t := m.To_InboxMessageId()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_InboxMessageId
func (m *InboxMessageId) To_InboxMessageId() *TLInboxMessageId {
	m.PredicateName = Predicate_inboxMessageId
	return &TLInboxMessageId{
		Data2: m,
	}
}

// MakeTLInboxMessageId
func MakeTLInboxMessageId(data2 *InboxMessageId) *TLInboxMessageId {
	if data2 == nil {
		return &TLInboxMessageId{Data2: &InboxMessageId{
			PredicateName: Predicate_inboxMessageId,
		}}
	} else {
		data2.PredicateName = Predicate_inboxMessageId
		return &TLInboxMessageId{Data2: data2}
	}
}

func (m *TLInboxMessageId) To_InboxMessageId() *InboxMessageId {
	m.Data2.PredicateName = Predicate_inboxMessageId
	return m.Data2
}

func (m *TLInboxMessageId) SetId(v int32) { m.Data2.Id = v }
func (m *TLInboxMessageId) GetId() int32  { return m.Data2.Id }

func (m *TLInboxMessageId) SetDialogMessageId(v int64) { m.Data2.DialogMessageId = v }
func (m *TLInboxMessageId) GetDialogMessageId() int64  { return m.Data2.DialogMessageId }

func (m *TLInboxMessageId) GetPredicateName() string {
	return Predicate_inboxMessageId
}

func (m *TLInboxMessageId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xc692c19f: func() []byte {
			x.UInt(0xc692c19f)

			x.Int(m.GetId())
			x.Long(m.GetDialogMessageId())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_inboxMessageId, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_inboxMessageId, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLInboxMessageId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxMessageId) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xc692c19f: func() error {
			m.SetId(dBuf.Int())
			m.SetDialogMessageId(dBuf.Long())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLInboxMessageId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// TLInboxSendUserMessageToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxSendUserMessageToInbox) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_inbox_sendUserMessageToInbox))

	switch uint32(m.Constructor) {
	case 0xf38edab3:
		x.UInt(0xf38edab3)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerUserId())
		x.Bytes(m.GetMessage().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLInboxSendUserMessageToInbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxSendUserMessageToInbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf38edab3:

		// not has flags

		m.FromId = dBuf.Long()
		m.PeerUserId = dBuf.Long()

		m3 := &InboxMessageData{}
		m3.Decode(dBuf)
		m.Message = m3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLInboxSendUserMessageToInbox) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLInboxSendChatMessageToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxSendChatMessageToInbox) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_inbox_sendChatMessageToInbox))

	switch uint32(m.Constructor) {
	case 0x971584c2:
		x.UInt(0x971584c2)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerChatId())
		x.Bytes(m.GetMessage().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLInboxSendChatMessageToInbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxSendChatMessageToInbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x971584c2:

		// not has flags

		m.FromId = dBuf.Long()
		m.PeerChatId = dBuf.Long()

		m3 := &InboxMessageData{}
		m3.Decode(dBuf)
		m.Message = m3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLInboxSendChatMessageToInbox) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLInboxSendChannelMessageToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxSendChannelMessageToInbox) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_inbox_sendChannelMessageToInbox))

	switch uint32(m.Constructor) {
	case 0x7a37f156:
		x.UInt(0x7a37f156)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerChannelId())
		x.Bytes(m.GetMessage().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLInboxSendChannelMessageToInbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxSendChannelMessageToInbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x7a37f156:

		// not has flags

		m.FromId = dBuf.Long()
		m.PeerChannelId = dBuf.Long()

		m3 := &mtproto.MessageBox{}
		m3.Decode(dBuf)
		m.Message = m3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLInboxSendChannelMessageToInbox) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLInboxSendUserMultiMessageToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxSendUserMultiMessageToInbox) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_inbox_sendUserMultiMessageToInbox))

	switch uint32(m.Constructor) {
	case 0x95c47179:
		x.UInt(0x95c47179)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerUserId())

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetMessage())))
		for _, v := range m.GetMessage() {
			x.Bytes((*v).Encode(layer))
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLInboxSendUserMultiMessageToInbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxSendUserMultiMessageToInbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x95c47179:

		// not has flags

		m.FromId = dBuf.Long()
		m.PeerUserId = dBuf.Long()
		c3 := dBuf.Int()
		if c3 != int32(mtproto.CRC32_vector) {
			// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
			return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
		}
		l3 := dBuf.Int()
		v3 := make([]*InboxMessageData, l3)
		for i := int32(0); i < l3; i++ {
			v3[i] = &InboxMessageData{}
			v3[i].Decode(dBuf)
		}
		m.Message = v3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLInboxSendUserMultiMessageToInbox) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLInboxSendChatMultiMessageToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxSendChatMultiMessageToInbox) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_inbox_sendChatMultiMessageToInbox))

	switch uint32(m.Constructor) {
	case 0xd69b718c:
		x.UInt(0xd69b718c)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerChatId())

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetMessage())))
		for _, v := range m.GetMessage() {
			x.Bytes((*v).Encode(layer))
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLInboxSendChatMultiMessageToInbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxSendChatMultiMessageToInbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xd69b718c:

		// not has flags

		m.FromId = dBuf.Long()
		m.PeerChatId = dBuf.Long()
		c3 := dBuf.Int()
		if c3 != int32(mtproto.CRC32_vector) {
			// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
			return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
		}
		l3 := dBuf.Int()
		v3 := make([]*InboxMessageData, l3)
		for i := int32(0); i < l3; i++ {
			v3[i] = &InboxMessageData{}
			v3[i].Decode(dBuf)
		}
		m.Message = v3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLInboxSendChatMultiMessageToInbox) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLInboxSendChannelMultiMessageToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxSendChannelMultiMessageToInbox) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_inbox_sendChannelMultiMessageToInbox))

	switch uint32(m.Constructor) {
	case 0x3b91d941:
		x.UInt(0x3b91d941)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerChannelId())

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetMessage())))
		for _, v := range m.GetMessage() {
			x.Bytes((*v).Encode(layer))
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLInboxSendChannelMultiMessageToInbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxSendChannelMultiMessageToInbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x3b91d941:

		// not has flags

		m.FromId = dBuf.Long()
		m.PeerChannelId = dBuf.Long()
		c3 := dBuf.Int()
		if c3 != int32(mtproto.CRC32_vector) {
			// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
			return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
		}
		l3 := dBuf.Int()
		v3 := make([]*mtproto.MessageBox, l3)
		for i := int32(0); i < l3; i++ {
			v3[i] = &mtproto.MessageBox{}
			v3[i].Decode(dBuf)
		}
		m.Message = v3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLInboxSendChannelMultiMessageToInbox) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLInboxEditUserMessageToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxEditUserMessageToInbox) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_inbox_editUserMessageToInbox))

	switch uint32(m.Constructor) {
	case 0x5cfb37a8:
		x.UInt(0x5cfb37a8)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerUserId())
		x.Bytes(m.GetMessage().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLInboxEditUserMessageToInbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxEditUserMessageToInbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x5cfb37a8:

		// not has flags

		m.FromId = dBuf.Long()
		m.PeerUserId = dBuf.Long()

		m3 := &mtproto.Message{}
		m3.Decode(dBuf)
		m.Message = m3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLInboxEditUserMessageToInbox) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLInboxEditChatMessageToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxEditChatMessageToInbox) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_inbox_editChatMessageToInbox))

	switch uint32(m.Constructor) {
	case 0x79107a0f:
		x.UInt(0x79107a0f)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerChatId())
		x.Bytes(m.GetMessage().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLInboxEditChatMessageToInbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxEditChatMessageToInbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x79107a0f:

		// not has flags

		m.FromId = dBuf.Long()
		m.PeerChatId = dBuf.Long()

		m3 := &mtproto.Message{}
		m3.Decode(dBuf)
		m.Message = m3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLInboxEditChatMessageToInbox) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLInboxEditChannelMessageToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxEditChannelMessageToInbox) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_inbox_editChannelMessageToInbox))

	switch uint32(m.Constructor) {
	case 0xf22c7d4d:
		x.UInt(0xf22c7d4d)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerChannelId())
		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())
		x.Bytes(m.GetMessage().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLInboxEditChannelMessageToInbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxEditChannelMessageToInbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf22c7d4d:

		// not has flags

		m.FromId = dBuf.Long()
		m.PeerChannelId = dBuf.Long()
		m.Pts = dBuf.Int()
		m.PtsCount = dBuf.Int()

		m5 := &mtproto.Message{}
		m5.Decode(dBuf)
		m.Message = m5

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLInboxEditChannelMessageToInbox) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLInboxDeleteMessagesToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxDeleteMessagesToInbox) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_inbox_deleteMessagesToInbox))

	switch uint32(m.Constructor) {
	case 0x851c6e34:
		x.UInt(0x851c6e34)

		// no flags

		x.Long(m.GetFromId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

		x.VectorLong(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLInboxDeleteMessagesToInbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxDeleteMessagesToInbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x851c6e34:

		// not has flags

		m.FromId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()

		m.Id = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLInboxDeleteMessagesToInbox) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLInboxDeleteChannelMessagesToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxDeleteChannelMessagesToInbox) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_inbox_deleteChannelMessagesToInbox))

	switch uint32(m.Constructor) {
	case 0x119a68c6:
		x.UInt(0x119a68c6)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerChannelId())
		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())

		x.VectorInt(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLInboxDeleteChannelMessagesToInbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxDeleteChannelMessagesToInbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x119a68c6:

		// not has flags

		m.FromId = dBuf.Long()
		m.PeerChannelId = dBuf.Long()
		m.Pts = dBuf.Int()
		m.PtsCount = dBuf.Int()

		m.Id = dBuf.VectorInt()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLInboxDeleteChannelMessagesToInbox) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLInboxDeleteUserHistoryToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxDeleteUserHistoryToInbox) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_inbox_deleteUserHistoryToInbox))

	switch uint32(m.Constructor) {
	case 0x140a8158:
		x.UInt(0x140a8158)

		// set flags
		var flags uint32 = 0

		if m.GetJustClear() == true {
			flags |= 1 << 1
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetFromId())
		x.Long(m.GetPeerUserId())
		x.Int(m.GetMaxId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLInboxDeleteUserHistoryToInbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxDeleteUserHistoryToInbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x140a8158:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.FromId = dBuf.Long()
		m.PeerUserId = dBuf.Long()
		if (flags & (1 << 1)) != 0 {
			m.JustClear = true
		}
		m.MaxId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLInboxDeleteUserHistoryToInbox) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLInboxDeleteChatHistoryToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxDeleteChatHistoryToInbox) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_inbox_deleteChatHistoryToInbox))

	switch uint32(m.Constructor) {
	case 0xd8aaa602:
		x.UInt(0xd8aaa602)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerChatId())
		x.Int(m.GetMaxId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLInboxDeleteChatHistoryToInbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxDeleteChatHistoryToInbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xd8aaa602:

		// not has flags

		m.FromId = dBuf.Long()
		m.PeerChatId = dBuf.Long()
		m.MaxId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLInboxDeleteChatHistoryToInbox) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLInboxReadUserMediaUnreadToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxReadUserMediaUnreadToInbox) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_inbox_readUserMediaUnreadToInbox))

	switch uint32(m.Constructor) {
	case 0x15c1034b:
		x.UInt(0x15c1034b)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerUserId())

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetId())))
		for _, v := range m.GetId() {
			x.Bytes((*v).Encode(layer))
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLInboxReadUserMediaUnreadToInbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxReadUserMediaUnreadToInbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x15c1034b:

		// not has flags

		m.FromId = dBuf.Long()
		m.PeerUserId = dBuf.Long()
		c3 := dBuf.Int()
		if c3 != int32(mtproto.CRC32_vector) {
			// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
			return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
		}
		l3 := dBuf.Int()
		v3 := make([]*InboxMessageId, l3)
		for i := int32(0); i < l3; i++ {
			v3[i] = &InboxMessageId{}
			v3[i].Decode(dBuf)
		}
		m.Id = v3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLInboxReadUserMediaUnreadToInbox) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLInboxReadChatMediaUnreadToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxReadChatMediaUnreadToInbox) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_inbox_readChatMediaUnreadToInbox))

	switch uint32(m.Constructor) {
	case 0x55415dd4:
		x.UInt(0x55415dd4)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerChatId())

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetId())))
		for _, v := range m.GetId() {
			x.Bytes((*v).Encode(layer))
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLInboxReadChatMediaUnreadToInbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxReadChatMediaUnreadToInbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x55415dd4:

		// not has flags

		m.FromId = dBuf.Long()
		m.PeerChatId = dBuf.Long()
		c3 := dBuf.Int()
		if c3 != int32(mtproto.CRC32_vector) {
			// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
			return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
		}
		l3 := dBuf.Int()
		v3 := make([]*InboxMessageId, l3)
		for i := int32(0); i < l3; i++ {
			v3[i] = &InboxMessageId{}
			v3[i].Decode(dBuf)
		}
		m.Id = v3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLInboxReadChatMediaUnreadToInbox) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLInboxReadChannelMediaUnreadToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxReadChannelMediaUnreadToInbox) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_inbox_readChannelMediaUnreadToInbox))

	switch uint32(m.Constructor) {
	case 0xa800dd36:
		x.UInt(0xa800dd36)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerChannelId())

		x.VectorInt(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLInboxReadChannelMediaUnreadToInbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxReadChannelMediaUnreadToInbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xa800dd36:

		// not has flags

		m.FromId = dBuf.Long()
		m.PeerChannelId = dBuf.Long()

		m.Id = dBuf.VectorInt()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLInboxReadChannelMediaUnreadToInbox) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLInboxUpdateHistoryReaded
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxUpdateHistoryReaded) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_inbox_updateHistoryReaded))

	switch uint32(m.Constructor) {
	case 0xc3c84ce0:
		x.UInt(0xc3c84ce0)

		// no flags

		x.Long(m.GetFromId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetMaxId())
		x.Long(m.GetSender())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLInboxUpdateHistoryReaded) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxUpdateHistoryReaded) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc3c84ce0:

		// not has flags

		m.FromId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.MaxId = dBuf.Int()
		m.Sender = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLInboxUpdateHistoryReaded) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLInboxUpdatePinnedMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxUpdatePinnedMessage) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_inbox_updatePinnedMessage))

	switch uint32(m.Constructor) {
	case 0xa96c2af4:
		x.UInt(0xa96c2af4)

		// set flags
		var flags uint32 = 0

		if m.GetUnpin() == true {
			flags |= 1 << 1
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetId())
		x.Long(m.GetDialogMessageId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLInboxUpdatePinnedMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxUpdatePinnedMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xa96c2af4:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.UserId = dBuf.Long()
		if (flags & (1 << 1)) != 0 {
			m.Unpin = true
		}
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.Id = dBuf.Int()
		m.DialogMessageId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLInboxUpdatePinnedMessage) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLInboxUnpinAllMessages
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxUnpinAllMessages) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_inbox_unpinAllMessages))

	switch uint32(m.Constructor) {
	case 0x231ca261:
		x.UInt(0x231ca261)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLInboxUnpinAllMessages) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxUnpinAllMessages) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x231ca261:

		// not has flags

		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLInboxUnpinAllMessages) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
