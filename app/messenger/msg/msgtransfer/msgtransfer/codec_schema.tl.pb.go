/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

// ConstructorList
// RequestList

package msgtransfer

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
	-1922780877: func() mtproto.TLObject { // 0x8d64b133
		o := MakeTLContentMessage(nil)
		o.Data2.Constructor = -1922780877
		return o
	},
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
	1763737728: func() mtproto.TLObject { // 0x69208080
		o := MakeTLOutboxMessage(nil)
		o.Data2.Constructor = 1763737728
		return o
	},
	1513645242: func() mtproto.TLObject { // 0x5a3864ba
		o := MakeTLSender(nil)
		o.Data2.Constructor = 1513645242
		return o
	},

	// Method
	-508367556: func() mtproto.TLObject { // 0xe1b2ed3c
		return &TLMsgtransferSendMessageToOutbox{
			Constructor: -508367556,
		}
	},
	-750661413: func() mtproto.TLObject { // 0xd341d0db
		return &TLMsgtransferSendMessageToInbox{
			Constructor: -750661413,
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
// ContentMessage <--
//  + TL_ContentMessage
//

func (m *ContentMessage) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_contentMessage:
		t := m.To_ContentMessage()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *ContentMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *ContentMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x8d64b133:
		m2 := MakeTLContentMessage(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_ContentMessage
func (m *ContentMessage) To_ContentMessage() *TLContentMessage {
	m.PredicateName = Predicate_contentMessage
	return &TLContentMessage{
		Data2: m,
	}
}

// MakeTLContentMessage
func MakeTLContentMessage(data2 *ContentMessage) *TLContentMessage {
	if data2 == nil {
		return &TLContentMessage{Data2: &ContentMessage{
			PredicateName: Predicate_contentMessage,
		}}
	} else {
		data2.PredicateName = Predicate_contentMessage
		return &TLContentMessage{Data2: data2}
	}
}

func (m *TLContentMessage) To_ContentMessage() *ContentMessage {
	m.Data2.PredicateName = Predicate_contentMessage
	return m.Data2
}

// // flags
func (m *TLContentMessage) SetId(v int32) { m.Data2.Id = v }
func (m *TLContentMessage) GetId() int32  { return m.Data2.Id }

func (m *TLContentMessage) SetDialogMessageId(v int64) { m.Data2.DialogMessageId = v }
func (m *TLContentMessage) GetDialogMessageId() int64  { return m.Data2.DialogMessageId }

func (m *TLContentMessage) SetMentioned(v bool) { m.Data2.Mentioned = v }
func (m *TLContentMessage) GetMentioned() bool  { return m.Data2.Mentioned }

func (m *TLContentMessage) SetMediaUnread(v bool) { m.Data2.MediaUnread = v }
func (m *TLContentMessage) GetMediaUnread() bool  { return m.Data2.MediaUnread }

func (m *TLContentMessage) SetReaction(v bool) { m.Data2.Reaction = v }
func (m *TLContentMessage) GetReaction() bool  { return m.Data2.Reaction }

func (m *TLContentMessage) SetSendUserId(v int64) { m.Data2.SendUserId = v }
func (m *TLContentMessage) GetSendUserId() int64  { return m.Data2.SendUserId }

func (m *TLContentMessage) GetPredicateName() string {
	return Predicate_contentMessage
}

func (m *TLContentMessage) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x8d64b133: func() error {
			x.UInt(0x8d64b133)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.GetMentioned() == true {
					flags |= 1 << 0
				}
				if m.GetMediaUnread() == true {
					flags |= 1 << 1
				}
				if m.GetReaction() == true {
					flags |= 1 << 2
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.UInt(flags)
			x.Int(m.GetId())
			x.Long(m.GetDialogMessageId())
			x.Long(m.GetSendUserId())
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_contentMessage, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_contentMessage, layer)
		return nil
	}

	return nil
}

func (m *TLContentMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLContentMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x8d64b133: func() error {
			var flags = dBuf.UInt()
			_ = flags
			m.SetId(dBuf.Int())
			m.SetDialogMessageId(dBuf.Long())
			if (flags & (1 << 0)) != 0 {
				m.SetMentioned(true)
			}
			if (flags & (1 << 1)) != 0 {
				m.SetMediaUnread(true)
			}
			if (flags & (1 << 2)) != 0 {
				m.SetReaction(true)
			}
			m.SetSendUserId(dBuf.Long())
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

///////////////////////////////////////////////////////////////////////////////
// OutboxMessage <--
//  + TL_OutboxMessage
//

func (m *OutboxMessage) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_outboxMessage:
		t := m.To_OutboxMessage()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *OutboxMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *OutboxMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x69208080:
		m2 := MakeTLOutboxMessage(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_OutboxMessage
func (m *OutboxMessage) To_OutboxMessage() *TLOutboxMessage {
	m.PredicateName = Predicate_outboxMessage
	return &TLOutboxMessage{
		Data2: m,
	}
}

// MakeTLOutboxMessage
func MakeTLOutboxMessage(data2 *OutboxMessage) *TLOutboxMessage {
	if data2 == nil {
		return &TLOutboxMessage{Data2: &OutboxMessage{
			PredicateName: Predicate_outboxMessage,
		}}
	} else {
		data2.PredicateName = Predicate_outboxMessage
		return &TLOutboxMessage{Data2: data2}
	}
}

func (m *TLOutboxMessage) To_OutboxMessage() *OutboxMessage {
	m.Data2.PredicateName = Predicate_outboxMessage
	return m.Data2
}

// // flags
func (m *TLOutboxMessage) SetNoWebpage(v bool) { m.Data2.NoWebpage = v }
func (m *TLOutboxMessage) GetNoWebpage() bool  { return m.Data2.NoWebpage }

func (m *TLOutboxMessage) SetBackground(v bool) { m.Data2.Background = v }
func (m *TLOutboxMessage) GetBackground() bool  { return m.Data2.Background }

func (m *TLOutboxMessage) SetRandomId(v int64) { m.Data2.RandomId = v }
func (m *TLOutboxMessage) GetRandomId() int64  { return m.Data2.RandomId }

func (m *TLOutboxMessage) SetMessage(v *mtproto.Message) { m.Data2.Message = v }
func (m *TLOutboxMessage) GetMessage() *mtproto.Message  { return m.Data2.Message }

func (m *TLOutboxMessage) GetPredicateName() string {
	return Predicate_outboxMessage
}

func (m *TLOutboxMessage) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x69208080: func() error {
			x.UInt(0x69208080)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.GetNoWebpage() == true {
					flags |= 1 << 0
				}
				if m.GetBackground() == true {
					flags |= 1 << 1
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.UInt(flags)
			x.Long(m.GetRandomId())
			m.GetMessage().Encode(x, layer)
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_outboxMessage, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_outboxMessage, layer)
		return nil
	}

	return nil
}

func (m *TLOutboxMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLOutboxMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x69208080: func() error {
			var flags = dBuf.UInt()
			_ = flags
			if (flags & (1 << 0)) != 0 {
				m.SetNoWebpage(true)
			}
			if (flags & (1 << 1)) != 0 {
				m.SetBackground(true)
			}
			m.SetRandomId(dBuf.Long())

			m4 := &mtproto.Message{}
			m4.Decode(dBuf)
			m.SetMessage(m4)

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
// Sender <--
//  + TL_Sender
//

func (m *Sender) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_sender:
		t := m.To_Sender()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *Sender) CalcByteSize(layer int32) int {
	return 0
}

func (m *Sender) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x5a3864ba:
		m2 := MakeTLSender(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_Sender
func (m *Sender) To_Sender() *TLSender {
	m.PredicateName = Predicate_sender
	return &TLSender{
		Data2: m,
	}
}

// MakeTLSender
func MakeTLSender(data2 *Sender) *TLSender {
	if data2 == nil {
		return &TLSender{Data2: &Sender{
			PredicateName: Predicate_sender,
		}}
	} else {
		data2.PredicateName = Predicate_sender
		return &TLSender{Data2: data2}
	}
}

func (m *TLSender) To_Sender() *Sender {
	m.Data2.PredicateName = Predicate_sender
	return m.Data2
}

func (m *TLSender) SetUserId(v int64) { m.Data2.UserId = v }
func (m *TLSender) GetUserId() int64  { return m.Data2.UserId }

func (m *TLSender) SetType(v int32) { m.Data2.Type = v }
func (m *TLSender) GetType() int32  { return m.Data2.Type }

func (m *TLSender) SetAuthKeyId(v int64) { m.Data2.AuthKeyId = v }
func (m *TLSender) GetAuthKeyId() int64  { return m.Data2.AuthKeyId }

func (m *TLSender) GetPredicateName() string {
	return Predicate_sender
}

func (m *TLSender) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x5a3864ba: func() error {
			x.UInt(0x5a3864ba)

			x.Long(m.GetUserId())
			x.Int(m.GetType())
			x.Long(m.GetAuthKeyId())
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_sender, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_sender, layer)
		return nil
	}

	return nil
}

func (m *TLSender) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSender) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x5a3864ba: func() error {
			m.SetUserId(dBuf.Long())
			m.SetType(dBuf.Int())
			m.SetAuthKeyId(dBuf.Long())
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
// TLMsgtransferSendMessageToOutbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgtransferSendMessageToOutbox) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xe1b2ed3c:
		x.UInt(0xe1b2ed3c)

		// set flags
		var flags uint32 = 0

		if m.GetUsers() != nil {
			flags |= 1 << 1
		}
		if m.GetChats() != nil {
			flags |= 1 << 2
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetMessage())))
		for _, v := range m.GetMessage() {
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

func (m *TLMsgtransferSendMessageToOutbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgtransferSendMessageToOutbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe1b2ed3c:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		c6 := dBuf.Int()
		if c6 != int32(mtproto.CRC32_vector) {
			// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 6, c6)
			return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 6, c6)
		}
		l6 := dBuf.Int()
		v6 := make([]*OutboxMessage, l6)
		for i := int32(0); i < l6; i++ {
			v6[i] = &OutboxMessage{}
			v6[i].Decode(dBuf)
		}
		m.Message = v6

		if (flags & (1 << 1)) != 0 {
			c7 := dBuf.Int()
			if c7 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 7, c7)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 7, c7)
			}
			l7 := dBuf.Int()
			v7 := make([]*mtproto.User, l7)
			for i := int32(0); i < l7; i++ {
				v7[i] = &mtproto.User{}
				v7[i].Decode(dBuf)
			}
			m.Users = v7
		}
		if (flags & (1 << 2)) != 0 {
			c8 := dBuf.Int()
			if c8 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 8, c8)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 8, c8)
			}
			l8 := dBuf.Int()
			v8 := make([]*mtproto.Chat, l8)
			for i := int32(0); i < l8; i++ {
				v8[i] = &mtproto.Chat{}
				v8[i].Decode(dBuf)
			}
			m.Chats = v8
		}
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMsgtransferSendMessageToInbox
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgtransferSendMessageToInbox) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xd341d0db:
		x.UInt(0xd341d0db)

		// set flags
		var flags uint32 = 0

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

func (m *TLMsgtransferSendMessageToInbox) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgtransferSendMessageToInbox) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xd341d0db:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.UserId = dBuf.Long()
		m.FromId = dBuf.Long()
		m.FromAuthKeyId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		c7 := dBuf.Int()
		if c7 != int32(mtproto.CRC32_vector) {
			// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 7, c7)
			return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 7, c7)
		}
		l7 := dBuf.Int()
		v7 := make([]*mtproto.MessageBox, l7)
		for i := int32(0); i < l7; i++ {
			v7[i] = &mtproto.MessageBox{}
			v7[i].Decode(dBuf)
		}
		m.BoxList = v7

		if (flags & (1 << 1)) != 0 {
			c8 := dBuf.Int()
			if c8 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 8, c8)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 8, c8)
			}
			l8 := dBuf.Int()
			v8 := make([]*mtproto.User, l8)
			for i := int32(0); i < l8; i++ {
				v8[i] = &mtproto.User{}
				v8[i].Decode(dBuf)
			}
			m.Users = v8
		}
		if (flags & (1 << 2)) != 0 {
			c9 := dBuf.Int()
			if c9 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 9, c9)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 9, c9)
			}
			l9 := dBuf.Int()
			v9 := make([]*mtproto.Chat, l9)
			for i := int32(0); i < l9; i++ {
				v9[i] = &mtproto.Chat{}
				v9[i].Decode(dBuf)
			}
			m.Chats = v9
		}
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}
