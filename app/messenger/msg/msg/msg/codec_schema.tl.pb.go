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

package msg

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
	1402283185: func() mtproto.TLObject { // 0x539524b1
		o := MakeTLOutboxMessage(nil)
		o.Data2.Constructor = 1402283185
		return o
	},
	1513645242: func() mtproto.TLObject { // 0x5a3864ba
		o := MakeTLSender(nil)
		o.Data2.Constructor = 1513645242
		return o
	},

	// Method
	902887962: func() mtproto.TLObject { // 0x35d0fa1a
		return &TLMsgPushUserMessage{
			Constructor: 902887962,
		}
	},
	673481940: func() mtproto.TLObject { // 0x282484d4
		return &TLMsgReadMessageContents{
			Constructor: 673481940,
		}
	},
	-188056380: func() mtproto.TLObject { // 0xf4ca7cc4
		return &TLMsgSendMessageV2{
			Constructor: -188056380,
		}
	},
	-2129725231: func() mtproto.TLObject { // 0x810ef8d1
		return &TLMsgEditMessage{
			Constructor: -2129725231,
		}
	},
	1778278369: func() mtproto.TLObject { // 0x69fe5fe1
		return &TLMsgEditMessageV2{
			Constructor: 1778278369,
		}
	},
	568855069: func() mtproto.TLObject { // 0x21e80a1d
		return &TLMsgDeleteMessages{
			Constructor: 568855069,
		}
	},
	1975576778: func() mtproto.TLObject { // 0x75c0e8ca
		return &TLMsgDeleteHistory{
			Constructor: 1975576778,
		}
	},
	649568574: func() mtproto.TLObject { // 0x26b7a13e
		return &TLMsgDeletePhoneCallHistory{
			Constructor: 649568574,
		}
	},
	-283155749: func() mtproto.TLObject { // 0xef1f62db
		return &TLMsgDeleteChatHistory{
			Constructor: -283155749,
		}
	},
	1510960658: func() mtproto.TLObject { // 0x5a0f6e12
		return &TLMsgReadHistory{
			Constructor: 1510960658,
		}
	},
	263827974: func() mtproto.TLObject { // 0xfb9b206
		return &TLMsgReadHistoryV2{
			Constructor: 263827974,
		}
	},
	-441560663: func() mtproto.TLObject { // 0xe5ae51a9
		return &TLMsgUpdatePinnedMessage{
			Constructor: -441560663,
		}
	},
	-1199153371: func() mtproto.TLObject { // 0xb8865f25
		return &TLMsgUnpinAllMessages{
			Constructor: -1199153371,
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
	case 0x539524b1:
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

func (m *TLOutboxMessage) SetScheduleDate(v *wrapperspb.Int32Value) { m.Data2.ScheduleDate = v }
func (m *TLOutboxMessage) GetScheduleDate() *wrapperspb.Int32Value  { return m.Data2.ScheduleDate }

func (m *TLOutboxMessage) GetPredicateName() string {
	return Predicate_outboxMessage
}

func (m *TLOutboxMessage) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x539524b1: func() error {
			x.UInt(0x539524b1)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.GetNoWebpage() == true {
					flags |= 1 << 0
				}
				if m.GetBackground() == true {
					flags |= 1 << 1
				}

				if m.GetScheduleDate() != nil {
					flags |= 1 << 2
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.UInt(flags)
			x.Long(m.GetRandomId())
			m.GetMessage().Encode(x, layer)
			if m.GetScheduleDate() != nil {
				x.Int(m.GetScheduleDate().Value)
			}

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
		0x539524b1: func() error {
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

			if (flags & (1 << 2)) != 0 {
				m.SetScheduleDate(&wrapperspb.Int32Value{Value: dBuf.Int()})
			}

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
// TLMsgPushUserMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgPushUserMessage) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x35d0fa1a:
		x.UInt(0x35d0fa1a)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetPushType())
		m.GetMessage().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMsgPushUserMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgPushUserMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x35d0fa1a:

		// not has flags

		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.PushType = dBuf.Int()

		m6 := &OutboxMessage{}
		m6.Decode(dBuf)
		m.Message = m6

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMsgReadMessageContents
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgReadMessageContents) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x282484d4:
		x.UInt(0x282484d4)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

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

func (m *TLMsgReadMessageContents) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgReadMessageContents) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x282484d4:

		// not has flags

		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		c5 := dBuf.Int()
		if c5 != int32(mtproto.CRC32_vector) {
			// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 5, c5)
			return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 5, c5)
		}
		l5 := dBuf.Int()
		v5 := make([]*ContentMessage, l5)
		for i := int32(0); i < l5; i++ {
			v5[i] = &ContentMessage{}
			v5[i].Decode(dBuf)
		}
		m.Id = v5

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMsgSendMessageV2
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgSendMessageV2) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xf4ca7cc4:
		x.UInt(0xf4ca7cc4)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

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

func (m *TLMsgSendMessageV2) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgSendMessageV2) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf4ca7cc4:

		// not has flags

		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		c5 := dBuf.Int()
		if c5 != int32(mtproto.CRC32_vector) {
			// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 5, c5)
			return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 5, c5)
		}
		l5 := dBuf.Int()
		v5 := make([]*OutboxMessage, l5)
		for i := int32(0); i < l5; i++ {
			v5[i] = &OutboxMessage{}
			v5[i].Decode(dBuf)
		}
		m.Message = v5

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMsgEditMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgEditMessage) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x810ef8d1:
		x.UInt(0x810ef8d1)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetEditType())
		m.GetMessage().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMsgEditMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgEditMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x810ef8d1:

		// not has flags

		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.EditType = dBuf.Int()

		m6 := &OutboxMessage{}
		m6.Decode(dBuf)
		m.Message = m6

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMsgEditMessageV2
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgEditMessageV2) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x69fe5fe1:
		x.UInt(0x69fe5fe1)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetEditType())
		m.GetNewMessage().Encode(x, layer)
		m.GetDstMessage().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMsgEditMessageV2) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgEditMessageV2) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x69fe5fe1:

		// not has flags

		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.EditType = dBuf.Int()

		m6 := &OutboxMessage{}
		m6.Decode(dBuf)
		m.NewMessage = m6

		m7 := &mtproto.MessageBox{}
		m7.Decode(dBuf)
		m.DstMessage = m7

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMsgDeleteMessages
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgDeleteMessages) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x21e80a1d:
		x.UInt(0x21e80a1d)

		// set flags
		var flags uint32 = 0

		if m.GetRevoke() == true {
			flags |= 1 << 1
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

		x.VectorInt(m.GetId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMsgDeleteMessages) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgDeleteMessages) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x21e80a1d:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		if (flags & (1 << 1)) != 0 {
			m.Revoke = true
		}

		m.Id = dBuf.VectorInt()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMsgDeleteHistory
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgDeleteHistory) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x75c0e8ca:
		x.UInt(0x75c0e8ca)

		// set flags
		var flags uint32 = 0

		if m.GetJustClear() == true {
			flags |= 1 << 0
		}
		if m.GetRevoke() == true {
			flags |= 1 << 1
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetMaxId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMsgDeleteHistory) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgDeleteHistory) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x75c0e8ca:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		if (flags & (1 << 0)) != 0 {
			m.JustClear = true
		}
		if (flags & (1 << 1)) != 0 {
			m.Revoke = true
		}
		m.MaxId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMsgDeletePhoneCallHistory
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgDeletePhoneCallHistory) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x26b7a13e:
		x.UInt(0x26b7a13e)

		// set flags
		var flags uint32 = 0

		if m.GetRevoke() == true {
			flags |= 1 << 1
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMsgDeletePhoneCallHistory) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgDeletePhoneCallHistory) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x26b7a13e:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		if (flags & (1 << 1)) != 0 {
			m.Revoke = true
		}
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMsgDeleteChatHistory
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgDeleteChatHistory) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xef1f62db:
		x.UInt(0xef1f62db)

		// no flags

		x.Long(m.GetChatId())
		x.Long(m.GetDeleteUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMsgDeleteChatHistory) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgDeleteChatHistory) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xef1f62db:

		// not has flags

		m.ChatId = dBuf.Long()
		m.DeleteUserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMsgReadHistory
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgReadHistory) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x5a0f6e12:
		x.UInt(0x5a0f6e12)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetMaxId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMsgReadHistory) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgReadHistory) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x5a0f6e12:

		// not has flags

		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.MaxId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMsgReadHistoryV2
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgReadHistoryV2) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xfb9b206:
		x.UInt(0xfb9b206)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetMaxId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMsgReadHistoryV2) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgReadHistoryV2) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xfb9b206:

		// not has flags

		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.MaxId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMsgUpdatePinnedMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgUpdatePinnedMessage) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xe5ae51a9:
		x.UInt(0xe5ae51a9)

		// set flags
		var flags uint32 = 0

		if m.GetSilent() == true {
			flags |= 1 << 0
		}
		if m.GetUnpin() == true {
			flags |= 1 << 1
		}
		if m.GetPmOneside() == true {
			flags |= 1 << 2
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMsgUpdatePinnedMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgUpdatePinnedMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe5ae51a9:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		if (flags & (1 << 0)) != 0 {
			m.Silent = true
		}
		if (flags & (1 << 1)) != 0 {
			m.Unpin = true
		}
		if (flags & (1 << 2)) != 0 {
			m.PmOneside = true
		}
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.Id = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMsgUnpinAllMessages
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgUnpinAllMessages) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xb8865f25:
		x.UInt(0xb8865f25)

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

func (m *TLMsgUnpinAllMessages) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgUnpinAllMessages) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb8865f25:

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
