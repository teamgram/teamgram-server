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

package authsession

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
	-646863312: func() mtproto.TLObject { // 0xd971a630
		o := MakeTLAuthKeyStateData(nil)
		o.Data2.Constructor = -646863312
		return o
	},
	-1701940816: func() mtproto.TLObject { // 0x9a8e71b0
		o := MakeTLClientSession(nil)
		o.Data2.Constructor = -1701940816
		return o
	},

	// Method
	820122180: func() mtproto.TLObject { // 0x30e21244
		return &TLAuthsessionGetAuthorizations{
			Constructor: 820122180,
		}
	},
	-1923126106: func() mtproto.TLObject { // 0x8d5f6ca6
		return &TLAuthsessionResetAuthorization{
			Constructor: -1923126106,
		}
	},
	-1473309015: func() mtproto.TLObject { // 0xa82f16a9
		return &TLAuthsessionGetLayer{
			Constructor: -1473309015,
		}
	},
	700170598: func() mtproto.TLObject { // 0x29bbc166
		return &TLAuthsessionGetLangPack{
			Constructor: 700170598,
		}
	},
	1616401854: func() mtproto.TLObject { // 0x605855be
		return &TLAuthsessionGetClient{
			Constructor: 1616401854,
		}
	},
	1486468441: func() mtproto.TLObject { // 0x5899b559
		return &TLAuthsessionGetLangCode{
			Constructor: 1486468441,
		}
	},
	1464409260: func() mtproto.TLObject { // 0x57491cac
		return &TLAuthsessionGetUserId{
			Constructor: 1464409260,
		}
	},
	-1279119039: func() mtproto.TLObject { // 0xb3c23141
		return &TLAuthsessionGetPushSessionId{
			Constructor: -1279119039,
		}
	},
	-1194371051: func() mtproto.TLObject { // 0xb8cf5815
		return &TLAuthsessionGetFutureSalts{
			Constructor: -1194371051,
		}
	},
	1421293608: func() mtproto.TLObject { // 0x54b73828
		return &TLAuthsessionQueryAuthKey{
			Constructor: 1421293608,
		}
	},
	1049889937: func() mtproto.TLObject { // 0x3e940c91
		return &TLAuthsessionSetAuthKey{
			Constructor: 1049889937,
		}
	},
	198050851: func() mtproto.TLObject { // 0xbce0423
		return &TLAuthsessionBindAuthKeyUser{
			Constructor: 198050851,
		}
	},
	123258440: func() mtproto.TLObject { // 0x758c648
		return &TLAuthsessionUnbindAuthKeyUser{
			Constructor: 123258440,
		}
	},
	-1871420202: func() mtproto.TLObject { // 0x907464d6
		return &TLAuthsessionGetPermAuthKeyId{
			Constructor: -1871420202,
		}
	},
	1620004742: func() mtproto.TLObject { // 0x608f4f86
		return &TLAuthsessionBindTempAuthKey{
			Constructor: 1620004742,
		}
	},
	47841172: func() mtproto.TLObject { // 0x2d9ff94
		return &TLAuthsessionSetClientSessionInfo{
			Constructor: 47841172,
		}
	},
	1851660579: func() mtproto.TLObject { // 0x6e5e1923
		return &TLAuthsessionGetAuthorization{
			Constructor: 1851660579,
		}
	},
	1331573041: func() mtproto.TLObject { // 0x4f5e3131
		return &TLAuthsessionGetAuthStateData{
			Constructor: 1331573041,
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
// AuthKeyStateData <--
//  + TL_AuthKeyStateData
//

func (m *AuthKeyStateData) Encode(layer int32) []byte {
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
	case Predicate_authKeyStateData:
		t := m.To_AuthKeyStateData()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *AuthKeyStateData) CalcByteSize(layer int32) int {
	return 0
}

func (m *AuthKeyStateData) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xd971a630:
		m2 := MakeTLAuthKeyStateData(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *AuthKeyStateData) DebugString() string {
	switch m.PredicateName {
	case Predicate_authKeyStateData:
		t := m.To_AuthKeyStateData()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_AuthKeyStateData
// authKeyStateData auth_key_id:long user_id:long key_state:int layer:int client_type:int android_push_session_id:long = AuthKeyStateData;
func (m *AuthKeyStateData) To_AuthKeyStateData() *TLAuthKeyStateData {
	m.PredicateName = Predicate_authKeyStateData
	return &TLAuthKeyStateData{
		Data2: m,
	}
}

// MakeTLAuthKeyStateData
// authKeyStateData auth_key_id:long user_id:long key_state:int layer:int client_type:int android_push_session_id:long = AuthKeyStateData;
func MakeTLAuthKeyStateData(data2 *AuthKeyStateData) *TLAuthKeyStateData {
	if data2 == nil {
		return &TLAuthKeyStateData{Data2: &AuthKeyStateData{
			PredicateName: Predicate_authKeyStateData,
		}}
	} else {
		data2.PredicateName = Predicate_authKeyStateData
		return &TLAuthKeyStateData{Data2: data2}
	}
}

func (m *TLAuthKeyStateData) To_AuthKeyStateData() *AuthKeyStateData {
	m.Data2.PredicateName = Predicate_authKeyStateData
	return m.Data2
}

func (m *TLAuthKeyStateData) SetAuthKeyId(v int64) { m.Data2.AuthKeyId = v }
func (m *TLAuthKeyStateData) GetAuthKeyId() int64  { return m.Data2.AuthKeyId }

func (m *TLAuthKeyStateData) SetUserId(v int64) { m.Data2.UserId = v }
func (m *TLAuthKeyStateData) GetUserId() int64  { return m.Data2.UserId }

func (m *TLAuthKeyStateData) SetKeyState(v int32) { m.Data2.KeyState = v }
func (m *TLAuthKeyStateData) GetKeyState() int32  { return m.Data2.KeyState }

func (m *TLAuthKeyStateData) SetLayer(v int32) { m.Data2.Layer = v }
func (m *TLAuthKeyStateData) GetLayer() int32  { return m.Data2.Layer }

func (m *TLAuthKeyStateData) SetClientType(v int32) { m.Data2.ClientType = v }
func (m *TLAuthKeyStateData) GetClientType() int32  { return m.Data2.ClientType }

func (m *TLAuthKeyStateData) SetAndroidPushSessionId(v int64) { m.Data2.AndroidPushSessionId = v }
func (m *TLAuthKeyStateData) GetAndroidPushSessionId() int64  { return m.Data2.AndroidPushSessionId }

func (m *TLAuthKeyStateData) GetPredicateName() string {
	return Predicate_authKeyStateData
}

func (m *TLAuthKeyStateData) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xd971a630: func() []byte {
			// authKeyStateData auth_key_id:long user_id:long key_state:int layer:int client_type:int android_push_session_id:long = AuthKeyStateData;
			x.UInt(0xd971a630)

			x.Long(m.GetAuthKeyId())
			x.Long(m.GetUserId())
			x.Int(m.GetKeyState())
			x.Int(m.GetLayer())
			x.Int(m.GetClientType())
			x.Long(m.GetAndroidPushSessionId())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_authKeyStateData, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_authKeyStateData, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLAuthKeyStateData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthKeyStateData) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xd971a630: func() error {
			// authKeyStateData auth_key_id:long user_id:long key_state:int layer:int client_type:int android_push_session_id:long = AuthKeyStateData;
			m.SetAuthKeyId(dBuf.Long())
			m.SetUserId(dBuf.Long())
			m.SetKeyState(dBuf.Int())
			m.SetLayer(dBuf.Int())
			m.SetClientType(dBuf.Int())
			m.SetAndroidPushSessionId(dBuf.Long())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLAuthKeyStateData) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// ClientSession <--
//  + TL_ClientSession
//

func (m *ClientSession) Encode(layer int32) []byte {
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
	case Predicate_clientSession:
		t := m.To_ClientSession()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *ClientSession) CalcByteSize(layer int32) int {
	return 0
}

func (m *ClientSession) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x9a8e71b0:
		m2 := MakeTLClientSession(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *ClientSession) DebugString() string {
	switch m.PredicateName {
	case Predicate_clientSession:
		t := m.To_ClientSession()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_ClientSession
// clientSession auth_key_id:long ip:string layer:int api_id:int device_model:string system_version:string app_version:string system_lang_code:string lang_pack:string lang_code:string proxy:string params:string = ClientSession;
func (m *ClientSession) To_ClientSession() *TLClientSession {
	m.PredicateName = Predicate_clientSession
	return &TLClientSession{
		Data2: m,
	}
}

// MakeTLClientSession
// clientSession auth_key_id:long ip:string layer:int api_id:int device_model:string system_version:string app_version:string system_lang_code:string lang_pack:string lang_code:string proxy:string params:string = ClientSession;
func MakeTLClientSession(data2 *ClientSession) *TLClientSession {
	if data2 == nil {
		return &TLClientSession{Data2: &ClientSession{
			PredicateName: Predicate_clientSession,
		}}
	} else {
		data2.PredicateName = Predicate_clientSession
		return &TLClientSession{Data2: data2}
	}
}

func (m *TLClientSession) To_ClientSession() *ClientSession {
	m.Data2.PredicateName = Predicate_clientSession
	return m.Data2
}

func (m *TLClientSession) SetAuthKeyId(v int64) { m.Data2.AuthKeyId = v }
func (m *TLClientSession) GetAuthKeyId() int64  { return m.Data2.AuthKeyId }

func (m *TLClientSession) SetIp(v string) { m.Data2.Ip = v }
func (m *TLClientSession) GetIp() string  { return m.Data2.Ip }

func (m *TLClientSession) SetLayer(v int32) { m.Data2.Layer = v }
func (m *TLClientSession) GetLayer() int32  { return m.Data2.Layer }

func (m *TLClientSession) SetApiId(v int32) { m.Data2.ApiId = v }
func (m *TLClientSession) GetApiId() int32  { return m.Data2.ApiId }

func (m *TLClientSession) SetDeviceModel(v string) { m.Data2.DeviceModel = v }
func (m *TLClientSession) GetDeviceModel() string  { return m.Data2.DeviceModel }

func (m *TLClientSession) SetSystemVersion(v string) { m.Data2.SystemVersion = v }
func (m *TLClientSession) GetSystemVersion() string  { return m.Data2.SystemVersion }

func (m *TLClientSession) SetAppVersion(v string) { m.Data2.AppVersion = v }
func (m *TLClientSession) GetAppVersion() string  { return m.Data2.AppVersion }

func (m *TLClientSession) SetSystemLangCode(v string) { m.Data2.SystemLangCode = v }
func (m *TLClientSession) GetSystemLangCode() string  { return m.Data2.SystemLangCode }

func (m *TLClientSession) SetLangPack(v string) { m.Data2.LangPack = v }
func (m *TLClientSession) GetLangPack() string  { return m.Data2.LangPack }

func (m *TLClientSession) SetLangCode(v string) { m.Data2.LangCode = v }
func (m *TLClientSession) GetLangCode() string  { return m.Data2.LangCode }

func (m *TLClientSession) SetProxy(v string) { m.Data2.Proxy = v }
func (m *TLClientSession) GetProxy() string  { return m.Data2.Proxy }

func (m *TLClientSession) SetParams(v string) { m.Data2.Params = v }
func (m *TLClientSession) GetParams() string  { return m.Data2.Params }

func (m *TLClientSession) GetPredicateName() string {
	return Predicate_clientSession
}

func (m *TLClientSession) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x9a8e71b0: func() []byte {
			// clientSession auth_key_id:long ip:string layer:int api_id:int device_model:string system_version:string app_version:string system_lang_code:string lang_pack:string lang_code:string proxy:string params:string = ClientSession;
			x.UInt(0x9a8e71b0)

			x.Long(m.GetAuthKeyId())
			x.String(m.GetIp())
			x.Int(m.GetLayer())
			x.Int(m.GetApiId())
			x.String(m.GetDeviceModel())
			x.String(m.GetSystemVersion())
			x.String(m.GetAppVersion())
			x.String(m.GetSystemLangCode())
			x.String(m.GetLangPack())
			x.String(m.GetLangCode())
			x.String(m.GetProxy())
			x.String(m.GetParams())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_clientSession, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_clientSession, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLClientSession) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLClientSession) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x9a8e71b0: func() error {
			// clientSession auth_key_id:long ip:string layer:int api_id:int device_model:string system_version:string app_version:string system_lang_code:string lang_pack:string lang_code:string proxy:string params:string = ClientSession;
			m.SetAuthKeyId(dBuf.Long())
			m.SetIp(dBuf.String())
			m.SetLayer(dBuf.Int())
			m.SetApiId(dBuf.Int())
			m.SetDeviceModel(dBuf.String())
			m.SetSystemVersion(dBuf.String())
			m.SetAppVersion(dBuf.String())
			m.SetSystemLangCode(dBuf.String())
			m.SetLangPack(dBuf.String())
			m.SetLangCode(dBuf.String())
			m.SetProxy(dBuf.String())
			m.SetParams(dBuf.String())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLClientSession) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// TLAuthsessionGetAuthorizations
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetAuthorizations) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_authsession_getAuthorizations))

	switch uint32(m.Constructor) {
	case 0x30e21244:
		// authsession.getAuthorizations user_id:long exclude_auth_keyId:long = account.Authorizations;
		x.UInt(0x30e21244)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetExcludeAuthKeyId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthsessionGetAuthorizations) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetAuthorizations) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x30e21244:
		// authsession.getAuthorizations user_id:long exclude_auth_keyId:long = account.Authorizations;

		// not has flags

		m.UserId = dBuf.Long()
		m.ExcludeAuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthsessionGetAuthorizations) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLAuthsessionResetAuthorization
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionResetAuthorization) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_authsession_resetAuthorization))

	switch uint32(m.Constructor) {
	case 0x8d5f6ca6:
		// authsession.resetAuthorization user_id:long auth_key_id:long hash:long = Vector<long>;
		x.UInt(0x8d5f6ca6)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Long(m.GetHash())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthsessionResetAuthorization) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionResetAuthorization) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x8d5f6ca6:
		// authsession.resetAuthorization user_id:long auth_key_id:long hash:long = Vector<long>;

		// not has flags

		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		m.Hash = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthsessionResetAuthorization) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLAuthsessionGetLayer
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetLayer) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_authsession_getLayer))

	switch uint32(m.Constructor) {
	case 0xa82f16a9:
		// authsession.getLayer auth_key_id:long = Int32;
		x.UInt(0xa82f16a9)

		// no flags

		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthsessionGetLayer) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetLayer) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xa82f16a9:
		// authsession.getLayer auth_key_id:long = Int32;

		// not has flags

		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthsessionGetLayer) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLAuthsessionGetLangPack
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetLangPack) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_authsession_getLangPack))

	switch uint32(m.Constructor) {
	case 0x29bbc166:
		// authsession.getLangPack auth_key_id:long = String;
		x.UInt(0x29bbc166)

		// no flags

		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthsessionGetLangPack) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetLangPack) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x29bbc166:
		// authsession.getLangPack auth_key_id:long = String;

		// not has flags

		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthsessionGetLangPack) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLAuthsessionGetClient
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetClient) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_authsession_getClient))

	switch uint32(m.Constructor) {
	case 0x605855be:
		// authsession.getClient auth_key_id:long = String;
		x.UInt(0x605855be)

		// no flags

		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthsessionGetClient) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetClient) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x605855be:
		// authsession.getClient auth_key_id:long = String;

		// not has flags

		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthsessionGetClient) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLAuthsessionGetLangCode
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetLangCode) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_authsession_getLangCode))

	switch uint32(m.Constructor) {
	case 0x5899b559:
		// authsession.getLangCode auth_key_id:long = String;
		x.UInt(0x5899b559)

		// no flags

		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthsessionGetLangCode) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetLangCode) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x5899b559:
		// authsession.getLangCode auth_key_id:long = String;

		// not has flags

		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthsessionGetLangCode) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLAuthsessionGetUserId
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetUserId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_authsession_getUserId))

	switch uint32(m.Constructor) {
	case 0x57491cac:
		// authsession.getUserId auth_key_id:long = Int64;
		x.UInt(0x57491cac)

		// no flags

		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthsessionGetUserId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetUserId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x57491cac:
		// authsession.getUserId auth_key_id:long = Int64;

		// not has flags

		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthsessionGetUserId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLAuthsessionGetPushSessionId
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetPushSessionId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_authsession_getPushSessionId))

	switch uint32(m.Constructor) {
	case 0xb3c23141:
		// authsession.getPushSessionId user_id:long auth_key_id:long token_type:int = Int64;
		x.UInt(0xb3c23141)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetTokenType())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthsessionGetPushSessionId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetPushSessionId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb3c23141:
		// authsession.getPushSessionId user_id:long auth_key_id:long token_type:int = Int64;

		// not has flags

		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		m.TokenType = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthsessionGetPushSessionId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLAuthsessionGetFutureSalts
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetFutureSalts) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_authsession_getFutureSalts))

	switch uint32(m.Constructor) {
	case 0xb8cf5815:
		// authsession.getFutureSalts auth_key_id:long num:int = FutureSalts;
		x.UInt(0xb8cf5815)

		// no flags

		x.Long(m.GetAuthKeyId())
		x.Int(m.GetNum())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthsessionGetFutureSalts) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetFutureSalts) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb8cf5815:
		// authsession.getFutureSalts auth_key_id:long num:int = FutureSalts;

		// not has flags

		m.AuthKeyId = dBuf.Long()
		m.Num = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthsessionGetFutureSalts) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLAuthsessionQueryAuthKey
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionQueryAuthKey) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_authsession_queryAuthKey))

	switch uint32(m.Constructor) {
	case 0x54b73828:
		// authsession.queryAuthKey auth_key_id:long = AuthKeyInfo;
		x.UInt(0x54b73828)

		// no flags

		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthsessionQueryAuthKey) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionQueryAuthKey) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x54b73828:
		// authsession.queryAuthKey auth_key_id:long = AuthKeyInfo;

		// not has flags

		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthsessionQueryAuthKey) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLAuthsessionSetAuthKey
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionSetAuthKey) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_authsession_setAuthKey))

	switch uint32(m.Constructor) {
	case 0x3e940c91:
		// authsession.setAuthKey auth_key:AuthKeyInfo future_salt:FutureSalt expires_in:int = Bool;
		x.UInt(0x3e940c91)

		// no flags

		x.Bytes(m.GetAuthKey().Encode(layer))
		x.Bytes(m.GetFutureSalt().Encode(layer))
		x.Int(m.GetExpiresIn())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthsessionSetAuthKey) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionSetAuthKey) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x3e940c91:
		// authsession.setAuthKey auth_key:AuthKeyInfo future_salt:FutureSalt expires_in:int = Bool;

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

func (m *TLAuthsessionSetAuthKey) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLAuthsessionBindAuthKeyUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionBindAuthKeyUser) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_authsession_bindAuthKeyUser))

	switch uint32(m.Constructor) {
	case 0xbce0423:
		// authsession.bindAuthKeyUser auth_key_id:long user_id:long = Int64;
		x.UInt(0xbce0423)

		// no flags

		x.Long(m.GetAuthKeyId())
		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthsessionBindAuthKeyUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionBindAuthKeyUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xbce0423:
		// authsession.bindAuthKeyUser auth_key_id:long user_id:long = Int64;

		// not has flags

		m.AuthKeyId = dBuf.Long()
		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthsessionBindAuthKeyUser) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLAuthsessionUnbindAuthKeyUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionUnbindAuthKeyUser) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_authsession_unbindAuthKeyUser))

	switch uint32(m.Constructor) {
	case 0x758c648:
		// authsession.unbindAuthKeyUser auth_key_id:long user_id:long = Bool;
		x.UInt(0x758c648)

		// no flags

		x.Long(m.GetAuthKeyId())
		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthsessionUnbindAuthKeyUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionUnbindAuthKeyUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x758c648:
		// authsession.unbindAuthKeyUser auth_key_id:long user_id:long = Bool;

		// not has flags

		m.AuthKeyId = dBuf.Long()
		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthsessionUnbindAuthKeyUser) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLAuthsessionGetPermAuthKeyId
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetPermAuthKeyId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_authsession_getPermAuthKeyId))

	switch uint32(m.Constructor) {
	case 0x907464d6:
		// authsession.getPermAuthKeyId auth_key_id:long= Int64;
		x.UInt(0x907464d6)

		// no flags

		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthsessionGetPermAuthKeyId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetPermAuthKeyId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x907464d6:
		// authsession.getPermAuthKeyId auth_key_id:long= Int64;

		// not has flags

		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthsessionGetPermAuthKeyId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLAuthsessionBindTempAuthKey
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionBindTempAuthKey) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_authsession_bindTempAuthKey))

	switch uint32(m.Constructor) {
	case 0x608f4f86:
		// authsession.bindTempAuthKey perm_auth_key_id:long nonce:long expires_at:int encrypted_message:bytes = Bool;
		x.UInt(0x608f4f86)

		// no flags

		x.Long(m.GetPermAuthKeyId())
		x.Long(m.GetNonce())
		x.Int(m.GetExpiresAt())
		x.StringBytes(m.GetEncryptedMessage())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthsessionBindTempAuthKey) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionBindTempAuthKey) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x608f4f86:
		// authsession.bindTempAuthKey perm_auth_key_id:long nonce:long expires_at:int encrypted_message:bytes = Bool;

		// not has flags

		m.PermAuthKeyId = dBuf.Long()
		m.Nonce = dBuf.Long()
		m.ExpiresAt = dBuf.Int()
		m.EncryptedMessage = dBuf.StringBytes()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthsessionBindTempAuthKey) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLAuthsessionSetClientSessionInfo
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionSetClientSessionInfo) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_authsession_setClientSessionInfo))

	switch uint32(m.Constructor) {
	case 0x2d9ff94:
		// authsession.setClientSessionInfo data:ClientSession = Bool;
		x.UInt(0x2d9ff94)

		// no flags

		x.Bytes(m.GetData().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthsessionSetClientSessionInfo) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionSetClientSessionInfo) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x2d9ff94:
		// authsession.setClientSessionInfo data:ClientSession = Bool;

		// not has flags

		m1 := &ClientSession{}
		m1.Decode(dBuf)
		m.Data = m1

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthsessionSetClientSessionInfo) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLAuthsessionGetAuthorization
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetAuthorization) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_authsession_getAuthorization))

	switch uint32(m.Constructor) {
	case 0x6e5e1923:
		// authsession.getAuthorization auth_key_id:long = Authorization;
		x.UInt(0x6e5e1923)

		// no flags

		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthsessionGetAuthorization) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetAuthorization) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x6e5e1923:
		// authsession.getAuthorization auth_key_id:long = Authorization;

		// not has flags

		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthsessionGetAuthorization) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLAuthsessionGetAuthStateData
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetAuthStateData) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_authsession_getAuthStateData))

	switch uint32(m.Constructor) {
	case 0x4f5e3131:
		// authsession.getAuthStateData auth_key_id:long = AuthKeyStateData;
		x.UInt(0x4f5e3131)

		// no flags

		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLAuthsessionGetAuthStateData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetAuthStateData) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x4f5e3131:
		// authsession.getAuthStateData auth_key_id:long = AuthKeyStateData;

		// not has flags

		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLAuthsessionGetAuthStateData) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
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
