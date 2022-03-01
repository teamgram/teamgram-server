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

package msg

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
	1402283185: func() mtproto.TLObject { // 0x539524b1
		o := MakeTLOutboxMessage(nil)
		o.Data2.Constructor = 1402283185
		return o
	},
	-1316261635: func() mtproto.TLObject { // 0xb18b70fd
		o := MakeTLContentMessage(nil)
		o.Data2.Constructor = -1316261635
		return o
	},
	1513645242: func() mtproto.TLObject { // 0x5a3864ba
		o := MakeTLSender(nil)
		o.Data2.Constructor = 1513645242
		return o
	},

	// Method
	1218652155: func() mtproto.TLObject { // 0x48a327fb
		return &TLMsgSendMessage{
			Constructor: 1218652155,
		}
	},
	-1727589428: func() mtproto.TLObject { // 0x990713cc
		return &TLMsgSendMultiMessage{
			Constructor: -1727589428,
		}
	},
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
	770211174: func() mtproto.TLObject { // 0x2de87d66
		return &TLMsgSendMessageV2{
			Constructor: 770211174,
		}
	},
	-1770495214: func() mtproto.TLObject { // 0x96786312
		return &TLMsgEditMessage{
			Constructor: -1770495214,
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
	893376573: func() mtproto.TLObject { // 0x353fd83d
		return &TLMsgDeleteChannelHistory{
			Constructor: 893376573,
		}
	},
	1611457677: func() mtproto.TLObject { // 0x600ce48d
		return &TLMsgDeleteChannelUserHistory{
			Constructor: 1611457677,
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
	-806710520: func() mtproto.TLObject { // 0xcfea9308
		return &TLMsgReadChannelHistory{
			Constructor: -806710520,
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
// OutboxMessage <--
//  + TL_OutboxMessage
//

func (m *OutboxMessage) Encode(layer int32) []byte {
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
	case Predicate_outboxMessage:
		t := m.To_OutboxMessage()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
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

func (m *OutboxMessage) DebugString() string {
	switch m.PredicateName {
	case Predicate_outboxMessage:
		t := m.To_OutboxMessage()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_OutboxMessage
// outboxMessage flags:# no_webpage:flags.0?true background:flags.1?true random_id:long message:Message schedule_date:flags.2?int = OutboxMessage;
func (m *OutboxMessage) To_OutboxMessage() *TLOutboxMessage {
	m.PredicateName = Predicate_outboxMessage
	return &TLOutboxMessage{
		Data2: m,
	}
}

// MakeTLOutboxMessage
// outboxMessage flags:# no_webpage:flags.0?true background:flags.1?true random_id:long message:Message schedule_date:flags.2?int = OutboxMessage;
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

//// flags
func (m *TLOutboxMessage) SetNoWebpage(v bool) { m.Data2.NoWebpage = v }
func (m *TLOutboxMessage) GetNoWebpage() bool  { return m.Data2.NoWebpage }

func (m *TLOutboxMessage) SetBackground(v bool) { m.Data2.Background = v }
func (m *TLOutboxMessage) GetBackground() bool  { return m.Data2.Background }

func (m *TLOutboxMessage) SetRandomId(v int64) { m.Data2.RandomId = v }
func (m *TLOutboxMessage) GetRandomId() int64  { return m.Data2.RandomId }

func (m *TLOutboxMessage) SetMessage(v *mtproto.Message) { m.Data2.Message = v }
func (m *TLOutboxMessage) GetMessage() *mtproto.Message  { return m.Data2.Message }

func (m *TLOutboxMessage) SetScheduleDate(v *types.Int32Value) { m.Data2.ScheduleDate = v }
func (m *TLOutboxMessage) GetScheduleDate() *types.Int32Value  { return m.Data2.ScheduleDate }

func (m *TLOutboxMessage) GetPredicateName() string {
	return Predicate_outboxMessage
}

func (m *TLOutboxMessage) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x539524b1: func() []byte {
			// outboxMessage flags:# no_webpage:flags.0?true background:flags.1?true random_id:long message:Message schedule_date:flags.2?int = OutboxMessage;
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
			x.Bytes(m.GetMessage().Encode(layer))
			if m.GetScheduleDate() != nil {
				x.Int(m.GetScheduleDate().Value)
			}

			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_outboxMessage, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_outboxMessage, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLOutboxMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLOutboxMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x539524b1: func() error {
			// outboxMessage flags:# no_webpage:flags.0?true background:flags.1?true random_id:long message:Message schedule_date:flags.2?int = OutboxMessage;
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
				m.SetScheduleDate(&types.Int32Value{Value: dBuf.Int()})
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

func (m *TLOutboxMessage) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// ContentMessage <--
//  + TL_ContentMessage
//

func (m *ContentMessage) Encode(layer int32) []byte {
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
	case Predicate_contentMessage:
		t := m.To_ContentMessage()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *ContentMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *ContentMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xb18b70fd:
		m2 := MakeTLContentMessage(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *ContentMessage) DebugString() string {
	switch m.PredicateName {
	case Predicate_contentMessage:
		t := m.To_ContentMessage()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_ContentMessage
// contentMessage id:int is_mentioned:Bool = ContentMessage;
func (m *ContentMessage) To_ContentMessage() *TLContentMessage {
	m.PredicateName = Predicate_contentMessage
	return &TLContentMessage{
		Data2: m,
	}
}

// MakeTLContentMessage
// contentMessage id:int is_mentioned:Bool = ContentMessage;
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

func (m *TLContentMessage) SetId(v int32) { m.Data2.Id = v }
func (m *TLContentMessage) GetId() int32  { return m.Data2.Id }

func (m *TLContentMessage) SetIsMentioned(v *mtproto.Bool) { m.Data2.IsMentioned = v }
func (m *TLContentMessage) GetIsMentioned() *mtproto.Bool  { return m.Data2.IsMentioned }

func (m *TLContentMessage) GetPredicateName() string {
	return Predicate_contentMessage
}

func (m *TLContentMessage) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xb18b70fd: func() []byte {
			// contentMessage id:int is_mentioned:Bool = ContentMessage;
			x.UInt(0xb18b70fd)

			x.Int(m.GetId())
			x.Bytes(m.GetIsMentioned().Encode(layer))
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_contentMessage, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_contentMessage, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLContentMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLContentMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xb18b70fd: func() error {
			// contentMessage id:int is_mentioned:Bool = ContentMessage;
			m.SetId(dBuf.Int())

			m1 := &mtproto.Bool{}
			m1.Decode(dBuf)
			m.SetIsMentioned(m1)

			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLContentMessage) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// Sender <--
//  + TL_Sender
//

func (m *Sender) Encode(layer int32) []byte {
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
	case Predicate_sender:
		t := m.To_Sender()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
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

func (m *Sender) DebugString() string {
	switch m.PredicateName {
	case Predicate_sender:
		t := m.To_Sender()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_Sender
// sender user_id:long type:int auth_key_id:long = Sender;
func (m *Sender) To_Sender() *TLSender {
	m.PredicateName = Predicate_sender
	return &TLSender{
		Data2: m,
	}
}

// MakeTLSender
// sender user_id:long type:int auth_key_id:long = Sender;
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

func (m *TLSender) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x5a3864ba: func() []byte {
			// sender user_id:long type:int auth_key_id:long = Sender;
			x.UInt(0x5a3864ba)

			x.Long(m.GetUserId())
			x.Int(m.GetType())
			x.Long(m.GetAuthKeyId())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_sender, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_sender, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLSender) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSender) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x5a3864ba: func() error {
			// sender user_id:long type:int auth_key_id:long = Sender;
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

func (m *TLSender) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// TLMsgSendMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgSendMessage) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_msg_sendMessage))

	switch uint32(m.Constructor) {
	case 0x48a327fb:
		// msg.sendMessage user_id:long auth_key_id:long peer_type:int peer_id:long message:OutboxMessage = Updates;
		x.UInt(0x48a327fb)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Bytes(m.GetMessage().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMsgSendMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgSendMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x48a327fb:
		// msg.sendMessage user_id:long auth_key_id:long peer_type:int peer_id:long message:OutboxMessage = Updates;

		// not has flags

		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()

		m5 := &OutboxMessage{}
		m5.Decode(dBuf)
		m.Message = m5

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMsgSendMessage) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMsgSendMultiMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgSendMultiMessage) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_msg_sendMultiMessage))

	switch uint32(m.Constructor) {
	case 0x990713cc:
		// msg.sendMultiMessage user_id:long auth_key_id:long peer_type:int peer_id:long message:Vector<OutboxMessage> = Updates;
		x.UInt(0x990713cc)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

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

func (m *TLMsgSendMultiMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgSendMultiMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x990713cc:
		// msg.sendMultiMessage user_id:long auth_key_id:long peer_type:int peer_id:long message:Vector<OutboxMessage> = Updates;

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

func (m *TLMsgSendMultiMessage) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMsgPushUserMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgPushUserMessage) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_msg_pushUserMessage))

	switch uint32(m.Constructor) {
	case 0x35d0fa1a:
		// msg.pushUserMessage user_id:long auth_key_id:long peer_type:int peer_id:long push_type:int message:OutboxMessage = Bool;
		x.UInt(0x35d0fa1a)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetPushType())
		x.Bytes(m.GetMessage().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMsgPushUserMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgPushUserMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x35d0fa1a:
		// msg.pushUserMessage user_id:long auth_key_id:long peer_type:int peer_id:long push_type:int message:OutboxMessage = Bool;

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

func (m *TLMsgPushUserMessage) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMsgReadMessageContents
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgReadMessageContents) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_msg_readMessageContents))

	switch uint32(m.Constructor) {
	case 0x282484d4:
		// msg.readMessageContents user_id:long auth_key_id:long peer_type:int peer_id:long id:Vector<ContentMessage> = messages.AffectedMessages;
		x.UInt(0x282484d4)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

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

func (m *TLMsgReadMessageContents) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgReadMessageContents) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x282484d4:
		// msg.readMessageContents user_id:long auth_key_id:long peer_type:int peer_id:long id:Vector<ContentMessage> = messages.AffectedMessages;

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

func (m *TLMsgReadMessageContents) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMsgSendMessageV2
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgSendMessageV2) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_msg_sendMessageV2))

	switch uint32(m.Constructor) {
	case 0x2de87d66:
		// msg.sendMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long message:Vector<OutboxMessage> = UpdateList;
		x.UInt(0x2de87d66)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

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

func (m *TLMsgSendMessageV2) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgSendMessageV2) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x2de87d66:
		// msg.sendMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long message:Vector<OutboxMessage> = UpdateList;

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

func (m *TLMsgSendMessageV2) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMsgEditMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgEditMessage) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_msg_editMessage))

	switch uint32(m.Constructor) {
	case 0x96786312:
		// msg.editMessage user_id:long auth_key_id:long peer_type:int peer_id:long message:OutboxMessage = Updates;
		x.UInt(0x96786312)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Bytes(m.GetMessage().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMsgEditMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgEditMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x96786312:
		// msg.editMessage user_id:long auth_key_id:long peer_type:int peer_id:long message:OutboxMessage = Updates;

		// not has flags

		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()

		m5 := &OutboxMessage{}
		m5.Decode(dBuf)
		m.Message = m5

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMsgEditMessage) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMsgDeleteMessages
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgDeleteMessages) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_msg_deleteMessages))

	switch uint32(m.Constructor) {
	case 0x21e80a1d:
		// msg.deleteMessages flags:# user_id:long auth_key_id:long peer_type:int peer_id:long revoke:flags.1?true id:Vector<int> = messages.AffectedMessages;
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

	return x.GetBuf()
}

func (m *TLMsgDeleteMessages) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgDeleteMessages) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x21e80a1d:
		// msg.deleteMessages flags:# user_id:long auth_key_id:long peer_type:int peer_id:long revoke:flags.1?true id:Vector<int> = messages.AffectedMessages;

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

func (m *TLMsgDeleteMessages) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMsgDeleteHistory
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgDeleteHistory) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_msg_deleteHistory))

	switch uint32(m.Constructor) {
	case 0x75c0e8ca:
		// msg.deleteHistory flags:# user_id:long auth_key_id:long peer_type:int peer_id:long just_clear:flags.0?true revoke:flags.1?true max_id:int = messages.AffectedHistory;
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

	return x.GetBuf()
}

func (m *TLMsgDeleteHistory) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgDeleteHistory) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x75c0e8ca:
		// msg.deleteHistory flags:# user_id:long auth_key_id:long peer_type:int peer_id:long just_clear:flags.0?true revoke:flags.1?true max_id:int = messages.AffectedHistory;

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

func (m *TLMsgDeleteHistory) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMsgDeleteChannelHistory
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgDeleteChannelHistory) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_msg_deleteChannelHistory))

	switch uint32(m.Constructor) {
	case 0x353fd83d:
		// msg.deleteChannelHistory user_id:long auth_key_id:long peer_type:int peer_id:long max_id:int = messages.AffectedHistory;
		x.UInt(0x353fd83d)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Int(m.GetMaxId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMsgDeleteChannelHistory) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgDeleteChannelHistory) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x353fd83d:
		// msg.deleteChannelHistory user_id:long auth_key_id:long peer_type:int peer_id:long max_id:int = messages.AffectedHistory;

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

func (m *TLMsgDeleteChannelHistory) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMsgDeleteChannelUserHistory
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgDeleteChannelUserHistory) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_msg_deleteChannelUserHistory))

	switch uint32(m.Constructor) {
	case 0x600ce48d:
		// msg.deleteChannelUserHistory user_id:long auth_key_id:long channel_id:long peer_type:int peer_id:long = messages.AffectedHistory;
		x.UInt(0x600ce48d)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Long(m.GetChannelId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMsgDeleteChannelUserHistory) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgDeleteChannelUserHistory) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x600ce48d:
		// msg.deleteChannelUserHistory user_id:long auth_key_id:long channel_id:long peer_type:int peer_id:long = messages.AffectedHistory;

		// not has flags

		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		m.ChannelId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMsgDeleteChannelUserHistory) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMsgDeletePhoneCallHistory
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgDeletePhoneCallHistory) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_msg_deletePhoneCallHistory))

	switch uint32(m.Constructor) {
	case 0x26b7a13e:
		// msg.deletePhoneCallHistory flags:# user_id:long auth_key_id:long revoke:flags.1?true = messages.AffectedFoundMessages;
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

	return x.GetBuf()
}

func (m *TLMsgDeletePhoneCallHistory) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgDeletePhoneCallHistory) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x26b7a13e:
		// msg.deletePhoneCallHistory flags:# user_id:long auth_key_id:long revoke:flags.1?true = messages.AffectedFoundMessages;

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

func (m *TLMsgDeletePhoneCallHistory) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMsgDeleteChatHistory
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgDeleteChatHistory) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_msg_deleteChatHistory))

	switch uint32(m.Constructor) {
	case 0xef1f62db:
		// msg.deleteChatHistory chat_id:long delete_user_id:long = Bool;
		x.UInt(0xef1f62db)

		// no flags

		x.Long(m.GetChatId())
		x.Long(m.GetDeleteUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMsgDeleteChatHistory) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgDeleteChatHistory) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xef1f62db:
		// msg.deleteChatHistory chat_id:long delete_user_id:long = Bool;

		// not has flags

		m.ChatId = dBuf.Long()
		m.DeleteUserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMsgDeleteChatHistory) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMsgReadHistory
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgReadHistory) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_msg_readHistory))

	switch uint32(m.Constructor) {
	case 0x5a0f6e12:
		// msg.readHistory user_id:long auth_key_id:long peer_type:int peer_id:long max_id:int = messages.AffectedMessages;
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

	return x.GetBuf()
}

func (m *TLMsgReadHistory) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgReadHistory) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x5a0f6e12:
		// msg.readHistory user_id:long auth_key_id:long peer_type:int peer_id:long max_id:int = messages.AffectedMessages;

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

func (m *TLMsgReadHistory) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMsgReadChannelHistory
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgReadChannelHistory) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_msg_readChannelHistory))

	switch uint32(m.Constructor) {
	case 0xcfea9308:
		// msg.readChannelHistory user_id:long auth_key_id:long channel_id:long max_id:int = Bool;
		x.UInt(0xcfea9308)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Long(m.GetChannelId())
		x.Int(m.GetMaxId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMsgReadChannelHistory) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgReadChannelHistory) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xcfea9308:
		// msg.readChannelHistory user_id:long auth_key_id:long channel_id:long max_id:int = Bool;

		// not has flags

		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		m.ChannelId = dBuf.Long()
		m.MaxId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMsgReadChannelHistory) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMsgUpdatePinnedMessage
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgUpdatePinnedMessage) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_msg_updatePinnedMessage))

	switch uint32(m.Constructor) {
	case 0xe5ae51a9:
		// msg.updatePinnedMessage flags:# user_id:long auth_key_id:long silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer_type:int peer_id:long id:int = Updates;
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

	return x.GetBuf()
}

func (m *TLMsgUpdatePinnedMessage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgUpdatePinnedMessage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe5ae51a9:
		// msg.updatePinnedMessage flags:# user_id:long auth_key_id:long silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer_type:int peer_id:long id:int = Updates;

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

func (m *TLMsgUpdatePinnedMessage) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMsgUnpinAllMessages
///////////////////////////////////////////////////////////////////////////////

func (m *TLMsgUnpinAllMessages) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_msg_unpinAllMessages))

	switch uint32(m.Constructor) {
	case 0xb8865f25:
		// msg.unpinAllMessages user_id:long auth_key_id:long peer_type:int peer_id:long = messages.AffectedHistory;
		x.UInt(0xb8865f25)

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

func (m *TLMsgUnpinAllMessages) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMsgUnpinAllMessages) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb8865f25:
		// msg.unpinAllMessages user_id:long auth_key_id:long peer_type:int peer_id:long = messages.AffectedHistory;

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

func (m *TLMsgUnpinAllMessages) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
