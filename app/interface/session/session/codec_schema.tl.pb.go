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

package session

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
	-739769057: func() mtproto.TLObject { // 0xd3e8051f
		o := MakeTLSessionClientEvent(nil)
		o.Data2.Constructor = -739769057
		return o
	},
	825806990: func() mtproto.TLObject { // 0x3138d08e
		o := MakeTLSessionClientData(nil)
		o.Data2.Constructor = 825806990
		return o
	},
	-606579889: func() mtproto.TLObject { // 0xdbd8534f
		o := MakeTLHttpSessionData(nil)
		o.Data2.Constructor = -606579889
		return o
	},

	// Method
	1798174801: func() mtproto.TLObject { // 0x6b2df851
		return &TLSessionQueryAuthKey{
			Constructor: 1798174801,
		}
	},
	487672075: func() mtproto.TLObject { // 0x1d11490b
		return &TLSessionSetAuthKey{
			Constructor: 487672075,
		}
	},
	1091351053: func() mtproto.TLObject { // 0x410cb20d
		return &TLSessionCreateSession{
			Constructor: 1091351053,
		}
	},
	-2023019028: func() mtproto.TLObject { // 0x876b2dec
		return &TLSessionSendDataToSession{
			Constructor: -2023019028,
		}
	},
	-1142152274: func() mtproto.TLObject { // 0xbbec23ae
		return &TLSessionSendHttpDataToSession{
			Constructor: -1142152274,
		}
	},
	393200211: func() mtproto.TLObject { // 0x176fc253
		return &TLSessionCloseSession{
			Constructor: 393200211,
		}
	},
	1075152191: func() mtproto.TLObject { // 0x4015853f
		return &TLSessionPushUpdatesData{
			Constructor: 1075152191,
		}
	},
	106898165: func() mtproto.TLObject { // 0x65f22f5
		return &TLSessionPushSessionUpdatesData{
			Constructor: 106898165,
		}
	},
	556344000: func() mtproto.TLObject { // 0x212922c0
		return &TLSessionPushRpcResultData{
			Constructor: 556344000,
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
// SessionClientEvent <--
//  + TL_SessionClientEvent
//

func (m *SessionClientEvent) Encode(layer int32) []byte {
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
	case Predicate_sessionClientEvent:
		t := m.To_SessionClientEvent()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *SessionClientEvent) CalcByteSize(layer int32) int {
	return 0
}

func (m *SessionClientEvent) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xd3e8051f:
		m2 := MakeTLSessionClientEvent(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *SessionClientEvent) DebugString() string {
	switch m.PredicateName {
	case Predicate_sessionClientEvent:
		t := m.To_SessionClientEvent()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_SessionClientEvent
// sessionClientEvent server_id:string conn_type:int auth_key_id:long session_id:long client_ip:string = SessionClientEvent;
func (m *SessionClientEvent) To_SessionClientEvent() *TLSessionClientEvent {
	m.PredicateName = Predicate_sessionClientEvent
	return &TLSessionClientEvent{
		Data2: m,
	}
}

// MakeTLSessionClientEvent
// sessionClientEvent server_id:string conn_type:int auth_key_id:long session_id:long client_ip:string = SessionClientEvent;
func MakeTLSessionClientEvent(data2 *SessionClientEvent) *TLSessionClientEvent {
	if data2 == nil {
		return &TLSessionClientEvent{Data2: &SessionClientEvent{
			PredicateName: Predicate_sessionClientEvent,
		}}
	} else {
		data2.PredicateName = Predicate_sessionClientEvent
		return &TLSessionClientEvent{Data2: data2}
	}
}

func (m *TLSessionClientEvent) To_SessionClientEvent() *SessionClientEvent {
	m.Data2.PredicateName = Predicate_sessionClientEvent
	return m.Data2
}

func (m *TLSessionClientEvent) SetServerId(v string) { m.Data2.ServerId = v }
func (m *TLSessionClientEvent) GetServerId() string  { return m.Data2.ServerId }

func (m *TLSessionClientEvent) SetConnType(v int32) { m.Data2.ConnType = v }
func (m *TLSessionClientEvent) GetConnType() int32  { return m.Data2.ConnType }

func (m *TLSessionClientEvent) SetAuthKeyId(v int64) { m.Data2.AuthKeyId = v }
func (m *TLSessionClientEvent) GetAuthKeyId() int64  { return m.Data2.AuthKeyId }

func (m *TLSessionClientEvent) SetSessionId(v int64) { m.Data2.SessionId = v }
func (m *TLSessionClientEvent) GetSessionId() int64  { return m.Data2.SessionId }

func (m *TLSessionClientEvent) SetClientIp(v string) { m.Data2.ClientIp = v }
func (m *TLSessionClientEvent) GetClientIp() string  { return m.Data2.ClientIp }

func (m *TLSessionClientEvent) GetPredicateName() string {
	return Predicate_sessionClientEvent
}

func (m *TLSessionClientEvent) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xd3e8051f: func() []byte {
			// sessionClientEvent server_id:string conn_type:int auth_key_id:long session_id:long client_ip:string = SessionClientEvent;
			x.UInt(0xd3e8051f)

			x.String(m.GetServerId())
			x.Int(m.GetConnType())
			x.Long(m.GetAuthKeyId())
			x.Long(m.GetSessionId())
			x.String(m.GetClientIp())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_sessionClientEvent, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_sessionClientEvent, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLSessionClientEvent) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionClientEvent) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xd3e8051f: func() error {
			// sessionClientEvent server_id:string conn_type:int auth_key_id:long session_id:long client_ip:string = SessionClientEvent;
			m.SetServerId(dBuf.String())
			m.SetConnType(dBuf.Int())
			m.SetAuthKeyId(dBuf.Long())
			m.SetSessionId(dBuf.Long())
			m.SetClientIp(dBuf.String())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLSessionClientEvent) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// SessionClientData <--
//  + TL_SessionClientData
//

func (m *SessionClientData) Encode(layer int32) []byte {
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
	case Predicate_sessionClientData:
		t := m.To_SessionClientData()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *SessionClientData) CalcByteSize(layer int32) int {
	return 0
}

func (m *SessionClientData) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x3138d08e:
		m2 := MakeTLSessionClientData(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *SessionClientData) DebugString() string {
	switch m.PredicateName {
	case Predicate_sessionClientData:
		t := m.To_SessionClientData()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_SessionClientData
// sessionClientData  server_id:string conn_type:int auth_key_id:long session_id:long client_ip:string quick_ack:int salt:long payload:bytes = SessionClientData;
func (m *SessionClientData) To_SessionClientData() *TLSessionClientData {
	m.PredicateName = Predicate_sessionClientData
	return &TLSessionClientData{
		Data2: m,
	}
}

// MakeTLSessionClientData
// sessionClientData  server_id:string conn_type:int auth_key_id:long session_id:long client_ip:string quick_ack:int salt:long payload:bytes = SessionClientData;
func MakeTLSessionClientData(data2 *SessionClientData) *TLSessionClientData {
	if data2 == nil {
		return &TLSessionClientData{Data2: &SessionClientData{
			PredicateName: Predicate_sessionClientData,
		}}
	} else {
		data2.PredicateName = Predicate_sessionClientData
		return &TLSessionClientData{Data2: data2}
	}
}

func (m *TLSessionClientData) To_SessionClientData() *SessionClientData {
	m.Data2.PredicateName = Predicate_sessionClientData
	return m.Data2
}

func (m *TLSessionClientData) SetServerId(v string) { m.Data2.ServerId = v }
func (m *TLSessionClientData) GetServerId() string  { return m.Data2.ServerId }

func (m *TLSessionClientData) SetConnType(v int32) { m.Data2.ConnType = v }
func (m *TLSessionClientData) GetConnType() int32  { return m.Data2.ConnType }

func (m *TLSessionClientData) SetAuthKeyId(v int64) { m.Data2.AuthKeyId = v }
func (m *TLSessionClientData) GetAuthKeyId() int64  { return m.Data2.AuthKeyId }

func (m *TLSessionClientData) SetSessionId(v int64) { m.Data2.SessionId = v }
func (m *TLSessionClientData) GetSessionId() int64  { return m.Data2.SessionId }

func (m *TLSessionClientData) SetClientIp(v string) { m.Data2.ClientIp = v }
func (m *TLSessionClientData) GetClientIp() string  { return m.Data2.ClientIp }

func (m *TLSessionClientData) SetQuickAck(v int32) { m.Data2.QuickAck = v }
func (m *TLSessionClientData) GetQuickAck() int32  { return m.Data2.QuickAck }

func (m *TLSessionClientData) SetSalt(v int64) { m.Data2.Salt = v }
func (m *TLSessionClientData) GetSalt() int64  { return m.Data2.Salt }

func (m *TLSessionClientData) SetPayload(v []byte) { m.Data2.Payload = v }
func (m *TLSessionClientData) GetPayload() []byte  { return m.Data2.Payload }

func (m *TLSessionClientData) GetPredicateName() string {
	return Predicate_sessionClientData
}

func (m *TLSessionClientData) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x3138d08e: func() []byte {
			// sessionClientData  server_id:string conn_type:int auth_key_id:long session_id:long client_ip:string quick_ack:int salt:long payload:bytes = SessionClientData;
			x.UInt(0x3138d08e)

			x.String(m.GetServerId())
			x.Int(m.GetConnType())
			x.Long(m.GetAuthKeyId())
			x.Long(m.GetSessionId())
			x.String(m.GetClientIp())
			x.Int(m.GetQuickAck())
			x.Long(m.GetSalt())
			x.StringBytes(m.GetPayload())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_sessionClientData, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_sessionClientData, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLSessionClientData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionClientData) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x3138d08e: func() error {
			// sessionClientData  server_id:string conn_type:int auth_key_id:long session_id:long client_ip:string quick_ack:int salt:long payload:bytes = SessionClientData;
			m.SetServerId(dBuf.String())
			m.SetConnType(dBuf.Int())
			m.SetAuthKeyId(dBuf.Long())
			m.SetSessionId(dBuf.Long())
			m.SetClientIp(dBuf.String())
			m.SetQuickAck(dBuf.Int())
			m.SetSalt(dBuf.Long())
			m.SetPayload(dBuf.StringBytes())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLSessionClientData) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// HttpSessionData <--
//  + TL_HttpSessionData
//

func (m *HttpSessionData) Encode(layer int32) []byte {
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
	case Predicate_httpSessionData:
		t := m.To_HttpSessionData()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *HttpSessionData) CalcByteSize(layer int32) int {
	return 0
}

func (m *HttpSessionData) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xdbd8534f:
		m2 := MakeTLHttpSessionData(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *HttpSessionData) DebugString() string {
	switch m.PredicateName {
	case Predicate_httpSessionData:
		t := m.To_HttpSessionData()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_HttpSessionData
// httpSessionData payload:bytes = HttpSessionData;
func (m *HttpSessionData) To_HttpSessionData() *TLHttpSessionData {
	m.PredicateName = Predicate_httpSessionData
	return &TLHttpSessionData{
		Data2: m,
	}
}

// MakeTLHttpSessionData
// httpSessionData payload:bytes = HttpSessionData;
func MakeTLHttpSessionData(data2 *HttpSessionData) *TLHttpSessionData {
	if data2 == nil {
		return &TLHttpSessionData{Data2: &HttpSessionData{
			PredicateName: Predicate_httpSessionData,
		}}
	} else {
		data2.PredicateName = Predicate_httpSessionData
		return &TLHttpSessionData{Data2: data2}
	}
}

func (m *TLHttpSessionData) To_HttpSessionData() *HttpSessionData {
	m.Data2.PredicateName = Predicate_httpSessionData
	return m.Data2
}

func (m *TLHttpSessionData) SetPayload(v []byte) { m.Data2.Payload = v }
func (m *TLHttpSessionData) GetPayload() []byte  { return m.Data2.Payload }

func (m *TLHttpSessionData) GetPredicateName() string {
	return Predicate_httpSessionData
}

func (m *TLHttpSessionData) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xdbd8534f: func() []byte {
			// httpSessionData payload:bytes = HttpSessionData;
			x.UInt(0xdbd8534f)

			x.StringBytes(m.GetPayload())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_httpSessionData, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_httpSessionData, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLHttpSessionData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLHttpSessionData) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xdbd8534f: func() error {
			// httpSessionData payload:bytes = HttpSessionData;
			m.SetPayload(dBuf.StringBytes())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLHttpSessionData) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// TLSessionQueryAuthKey
///////////////////////////////////////////////////////////////////////////////

func (m *TLSessionQueryAuthKey) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_session_queryAuthKey))

	switch uint32(m.Constructor) {
	case 0x6b2df851:
		// session.queryAuthKey auth_key_id:long = AuthKeyInfo;
		x.UInt(0x6b2df851)

		// no flags

		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLSessionQueryAuthKey) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionQueryAuthKey) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x6b2df851:
		// session.queryAuthKey auth_key_id:long = AuthKeyInfo;

		// not has flags

		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLSessionQueryAuthKey) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLSessionSetAuthKey
///////////////////////////////////////////////////////////////////////////////

func (m *TLSessionSetAuthKey) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_session_setAuthKey))

	switch uint32(m.Constructor) {
	case 0x1d11490b:
		// session.setAuthKey auth_key:AuthKeyInfo future_salt:FutureSalt expires_in:int = Bool;
		x.UInt(0x1d11490b)

		// no flags

		x.Bytes(m.GetAuthKey().Encode(layer))
		x.Bytes(m.GetFutureSalt().Encode(layer))
		x.Int(m.GetExpiresIn())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLSessionSetAuthKey) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionSetAuthKey) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x1d11490b:
		// session.setAuthKey auth_key:AuthKeyInfo future_salt:FutureSalt expires_in:int = Bool;

		// not has flags

		m1 := &mtproto.AuthKeyInfo{}
		m1.Decode(dBuf)
		m.AuthKey = m1

		m2 := &mtproto.FutureSalt{}
		m2.Decode(dBuf)
		m.FutureSalt = m2

		m.ExpiresIn = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLSessionSetAuthKey) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLSessionCreateSession
///////////////////////////////////////////////////////////////////////////////

func (m *TLSessionCreateSession) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_session_createSession))

	switch uint32(m.Constructor) {
	case 0x410cb20d:
		// session.createSession client:SessionClientEvent = Bool;
		x.UInt(0x410cb20d)

		// no flags

		x.Bytes(m.GetClient().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLSessionCreateSession) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionCreateSession) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x410cb20d:
		// session.createSession client:SessionClientEvent = Bool;

		// not has flags

		m1 := &SessionClientEvent{}
		m1.Decode(dBuf)
		m.Client = m1

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLSessionCreateSession) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLSessionSendDataToSession
///////////////////////////////////////////////////////////////////////////////

func (m *TLSessionSendDataToSession) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_session_sendDataToSession))

	switch uint32(m.Constructor) {
	case 0x876b2dec:
		// session.sendDataToSession data:SessionClientData = Bool;
		x.UInt(0x876b2dec)

		// no flags

		x.Bytes(m.GetData().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLSessionSendDataToSession) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionSendDataToSession) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x876b2dec:
		// session.sendDataToSession data:SessionClientData = Bool;

		// not has flags

		m1 := &SessionClientData{}
		m1.Decode(dBuf)
		m.Data = m1

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLSessionSendDataToSession) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLSessionSendHttpDataToSession
///////////////////////////////////////////////////////////////////////////////

func (m *TLSessionSendHttpDataToSession) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_session_sendHttpDataToSession))

	switch uint32(m.Constructor) {
	case 0xbbec23ae:
		// session.sendHttpDataToSession client:SessionClientData = HttpSessionData;
		x.UInt(0xbbec23ae)

		// no flags

		x.Bytes(m.GetClient().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLSessionSendHttpDataToSession) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionSendHttpDataToSession) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xbbec23ae:
		// session.sendHttpDataToSession client:SessionClientData = HttpSessionData;

		// not has flags

		m1 := &SessionClientData{}
		m1.Decode(dBuf)
		m.Client = m1

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLSessionSendHttpDataToSession) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLSessionCloseSession
///////////////////////////////////////////////////////////////////////////////

func (m *TLSessionCloseSession) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_session_closeSession))

	switch uint32(m.Constructor) {
	case 0x176fc253:
		// session.closeSession client:SessionClientEvent = Bool;
		x.UInt(0x176fc253)

		// no flags

		x.Bytes(m.GetClient().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLSessionCloseSession) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionCloseSession) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x176fc253:
		// session.closeSession client:SessionClientEvent = Bool;

		// not has flags

		m1 := &SessionClientEvent{}
		m1.Decode(dBuf)
		m.Client = m1

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLSessionCloseSession) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLSessionPushUpdatesData
///////////////////////////////////////////////////////////////////////////////

func (m *TLSessionPushUpdatesData) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_session_pushUpdatesData))

	switch uint32(m.Constructor) {
	case 0x4015853f:
		// session.pushUpdatesData flags:# auth_key_id:long notification:flags.0?true updates:Updates = Bool;
		x.UInt(0x4015853f)

		// set flags
		var flags uint32 = 0

		if m.GetNotification() == true {
			flags |= 1 << 0
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetAuthKeyId())
		x.Bytes(m.GetUpdates().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLSessionPushUpdatesData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionPushUpdatesData) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x4015853f:
		// session.pushUpdatesData flags:# auth_key_id:long notification:flags.0?true updates:Updates = Bool;

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.AuthKeyId = dBuf.Long()
		if (flags & (1 << 0)) != 0 {
			m.Notification = true
		}

		m4 := &mtproto.Updates{}
		m4.Decode(dBuf)
		m.Updates = m4

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLSessionPushUpdatesData) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLSessionPushSessionUpdatesData
///////////////////////////////////////////////////////////////////////////////

func (m *TLSessionPushSessionUpdatesData) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_session_pushSessionUpdatesData))

	switch uint32(m.Constructor) {
	case 0x65f22f5:
		// session.pushSessionUpdatesData auth_key_id:long session_id:long updates:Updates = Bool;
		x.UInt(0x65f22f5)

		// no flags

		x.Long(m.GetAuthKeyId())
		x.Long(m.GetSessionId())
		x.Bytes(m.GetUpdates().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLSessionPushSessionUpdatesData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionPushSessionUpdatesData) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x65f22f5:
		// session.pushSessionUpdatesData auth_key_id:long session_id:long updates:Updates = Bool;

		// not has flags

		m.AuthKeyId = dBuf.Long()
		m.SessionId = dBuf.Long()

		m3 := &mtproto.Updates{}
		m3.Decode(dBuf)
		m.Updates = m3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLSessionPushSessionUpdatesData) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLSessionPushRpcResultData
///////////////////////////////////////////////////////////////////////////////

func (m *TLSessionPushRpcResultData) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_session_pushRpcResultData))

	switch uint32(m.Constructor) {
	case 0x212922c0:
		// session.pushRpcResultData auth_key_id:long session_id:long client_req_msg_id:long rpc_result_data:bytes = Bool;
		x.UInt(0x212922c0)

		// no flags

		x.Long(m.GetAuthKeyId())
		x.Long(m.GetSessionId())
		x.Long(m.GetClientReqMsgId())
		x.StringBytes(m.GetRpcResultData())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLSessionPushRpcResultData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionPushRpcResultData) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x212922c0:
		// session.pushRpcResultData auth_key_id:long session_id:long client_req_msg_id:long rpc_result_data:bytes = Bool;

		// not has flags

		m.AuthKeyId = dBuf.Long()
		m.SessionId = dBuf.Long()
		m.ClientReqMsgId = dBuf.Long()
		m.RpcResultData = dBuf.StringBytes()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLSessionPushRpcResultData) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
