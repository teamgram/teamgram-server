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

package inbox

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
	-2061734348: func() mtproto.TLObject { // 0x851c6e34
		return &TLInboxDeleteMessagesToInbox{
			Constructor: -2061734348,
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
	-209599207: func() mtproto.TLObject { // 0xf381c519
		return &TLInboxSendUserMessageToInboxV2{
			Constructor: -209599207,
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

func (m *InboxMessageData) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_inboxMessageData:
		t := m.To_InboxMessageData()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
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

func (m *TLInboxMessageData) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x3bbdadd4: func() error {
			x.UInt(0x3bbdadd4)

			x.Long(m.GetRandomId())
			x.Long(m.GetDialogMessageId())
			m.GetMessage().Encode(x, layer)
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_inboxMessageData, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_inboxMessageData, layer)
		return nil
	}

	return nil
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

///////////////////////////////////////////////////////////////////////////////
// InboxMessageId <--
//  + TL_InboxMessageId
//

func (m *InboxMessageId) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_inboxMessageId:
		t := m.To_InboxMessageId()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
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

func (m *TLInboxMessageId) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc692c19f: func() error {
			x.UInt(0xc692c19f)

			x.Int(m.GetId())
			x.Long(m.GetDialogMessageId())
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_inboxMessageId, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_inboxMessageId, layer)
		return nil
	}

	return nil
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

//----------------------------------------------------------------------------------------------------------------
// TLInboxSendUserMessageToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxSendUserMessageToInbox) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xf38edab3:
		x.UInt(0xf38edab3)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerUserId())
		m.GetMessage().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
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

// TLInboxSendChatMessageToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxSendChatMessageToInbox) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x971584c2:
		x.UInt(0x971584c2)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerChatId())
		m.GetMessage().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
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

// TLInboxSendUserMultiMessageToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxSendUserMultiMessageToInbox) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x95c47179:
		x.UInt(0x95c47179)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerUserId())

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetMessage())))
		for _, v := range m.GetMessage() {
			v.Encode(x, layer)
		}

	default:
		// log.Errorf("")
	}

	return nil
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

// TLInboxSendChatMultiMessageToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxSendChatMultiMessageToInbox) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xd69b718c:
		x.UInt(0xd69b718c)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerChatId())

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetMessage())))
		for _, v := range m.GetMessage() {
			v.Encode(x, layer)
		}

	default:
		// log.Errorf("")
	}

	return nil
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

// TLInboxEditUserMessageToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxEditUserMessageToInbox) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x5cfb37a8:
		x.UInt(0x5cfb37a8)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerUserId())
		m.GetMessage().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
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

// TLInboxEditChatMessageToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxEditChatMessageToInbox) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x79107a0f:
		x.UInt(0x79107a0f)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerChatId())
		m.GetMessage().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
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

// TLInboxDeleteMessagesToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxDeleteMessagesToInbox) Encode(x *mtproto.EncodeBuf, layer int32) error {
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

	return nil
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

// TLInboxDeleteUserHistoryToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxDeleteUserHistoryToInbox) Encode(x *mtproto.EncodeBuf, layer int32) error {
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

	return nil
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

// TLInboxDeleteChatHistoryToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxDeleteChatHistoryToInbox) Encode(x *mtproto.EncodeBuf, layer int32) error {
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

	return nil
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

// TLInboxReadUserMediaUnreadToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxReadUserMediaUnreadToInbox) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x15c1034b:
		x.UInt(0x15c1034b)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerUserId())

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetId())))
		for _, v := range m.GetId() {
			v.Encode(x, layer)
		}

	default:
		// log.Errorf("")
	}

	return nil
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

// TLInboxReadChatMediaUnreadToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxReadChatMediaUnreadToInbox) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x55415dd4:
		x.UInt(0x55415dd4)

		// no flags

		x.Long(m.GetFromId())
		x.Long(m.GetPeerChatId())

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetId())))
		for _, v := range m.GetId() {
			v.Encode(x, layer)
		}

	default:
		// log.Errorf("")
	}

	return nil
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

// TLInboxUpdateHistoryReaded
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxUpdateHistoryReaded) Encode(x *mtproto.EncodeBuf, layer int32) error {
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

	return nil
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

// TLInboxUpdatePinnedMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxUpdatePinnedMessage) Encode(x *mtproto.EncodeBuf, layer int32) error {
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

	return nil
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

// TLInboxUnpinAllMessages
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxUnpinAllMessages) Encode(x *mtproto.EncodeBuf, layer int32) error {
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

	return nil
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

// TLInboxSendUserMessageToInboxV2
///////////////////////////////////////////////////////////////////////////////

func (m *TLInboxSendUserMessageToInboxV2) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xf381c519:
		x.UInt(0xf381c519)

		// set flags
		var flags uint32 = 0

		if m.GetOut() == true {
			flags |= 1 << 0
		}

		if m.GetUsers() != nil {
			flags |= 1 << 1
		}
		if m.GetChats() != nil {
			flags |= 1 << 2
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetUserId())
		x.Long(m.GetFromId())
		x.Long(m.GetFromAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetBoxList())))
		for _, v := range m.GetBoxList() {
			v.Encode(x, layer)
		}

		if m.GetUsers() != nil {
			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetUsers())))
			for _, v := range m.GetUsers() {
				v.Encode(x, layer)
			}
		}
		if m.GetChats() != nil {
			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetChats())))
			for _, v := range m.GetChats() {
				v.Encode(x, layer)
			}
		}

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLInboxSendUserMessageToInboxV2) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInboxSendUserMessageToInboxV2) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf381c519:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.UserId = dBuf.Long()
		if (flags & (1 << 0)) != 0 {
			m.Out = true
		}
		m.FromId = dBuf.Long()
		m.FromAuthKeyId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		c8 := dBuf.Int()
		if c8 != int32(mtproto.CRC32_vector) {
			// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 8, c8)
			return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 8, c8)
		}
		l8 := dBuf.Int()
		v8 := make([]*mtproto.MessageBox, l8)
		for i := int32(0); i < l8; i++ {
			v8[i] = &mtproto.MessageBox{}
			v8[i].Decode(dBuf)
		}
		m.BoxList = v8

		if (flags & (1 << 1)) != 0 {
			c9 := dBuf.Int()
			if c9 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 9, c9)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 9, c9)
			}
			l9 := dBuf.Int()
			v9 := make([]*mtproto.User, l9)
			for i := int32(0); i < l9; i++ {
				v9[i] = &mtproto.User{}
				v9[i].Decode(dBuf)
			}
			m.Users = v9
		}
		if (flags & (1 << 2)) != 0 {
			c10 := dBuf.Int()
			if c10 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 10, c10)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 10, c10)
			}
			l10 := dBuf.Int()
			v10 := make([]*mtproto.Chat, l10)
			for i := int32(0); i < l10; i++ {
				v10[i] = &mtproto.Chat{}
				v10[i].Decode(dBuf)
			}
			m.Chats = v10
		}
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}
