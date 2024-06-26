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

package authsession

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
	-532639977: func() mtproto.TLObject { // 0xe0408f17
		o := MakeTLAuthKeyStateData(nil)
		o.Data2.Constructor = -532639977
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
	1147475077: func() mtproto.TLObject { // 0x44651485
		return &TLAuthsessionSetLayer{
			Constructor: 1147475077,
		}
	},
	2095024780: func() mtproto.TLObject { // 0x7cdf8a8c
		return &TLAuthsessionSetInitConnection{
			Constructor: 2095024780,
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

func (m *AuthKeyStateData) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_authKeyStateData:
		t := m.To_AuthKeyStateData()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *AuthKeyStateData) CalcByteSize(layer int32) int {
	return 0
}

func (m *AuthKeyStateData) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xe0408f17:
		m2 := MakeTLAuthKeyStateData(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_AuthKeyStateData
func (m *AuthKeyStateData) To_AuthKeyStateData() *TLAuthKeyStateData {
	m.PredicateName = Predicate_authKeyStateData
	return &TLAuthKeyStateData{
		Data2: m,
	}
}

// MakeTLAuthKeyStateData
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

// // flags
func (m *TLAuthKeyStateData) SetAuthKeyId(v int64) { m.Data2.AuthKeyId = v }
func (m *TLAuthKeyStateData) GetAuthKeyId() int64  { return m.Data2.AuthKeyId }

func (m *TLAuthKeyStateData) SetKeyState(v int32) { m.Data2.KeyState = v }
func (m *TLAuthKeyStateData) GetKeyState() int32  { return m.Data2.KeyState }

func (m *TLAuthKeyStateData) SetUserId(v int64) { m.Data2.UserId = v }
func (m *TLAuthKeyStateData) GetUserId() int64  { return m.Data2.UserId }

func (m *TLAuthKeyStateData) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLAuthKeyStateData) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLAuthKeyStateData) SetClient(v *ClientSession) { m.Data2.Client = v }
func (m *TLAuthKeyStateData) GetClient() *ClientSession  { return m.Data2.Client }

func (m *TLAuthKeyStateData) SetAndroidPushSessionId(v *wrapperspb.Int64Value) {
	m.Data2.AndroidPushSessionId = v
}
func (m *TLAuthKeyStateData) GetAndroidPushSessionId() *wrapperspb.Int64Value {
	return m.Data2.AndroidPushSessionId
}

func (m *TLAuthKeyStateData) GetPredicateName() string {
	return Predicate_authKeyStateData
}

func (m *TLAuthKeyStateData) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xe0408f17: func() error {
			x.UInt(0xe0408f17)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.GetClient() != nil {
					flags |= 1 << 0
				}
				if m.GetAndroidPushSessionId() != nil {
					flags |= 1 << 1
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.UInt(flags)
			x.Long(m.GetAuthKeyId())
			x.Int(m.GetKeyState())
			x.Long(m.GetUserId())
			x.Long(m.GetAccessHash())
			if m.GetClient() != nil {
				m.GetClient().Encode(x, layer)
			}

			if m.GetAndroidPushSessionId() != nil {
				x.Long(m.GetAndroidPushSessionId().Value)
			}

			return nil
		},
	}

	clazzId := GetClazzID(Predicate_authKeyStateData, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_authKeyStateData, layer)
		return nil
	}

	return nil
}

func (m *TLAuthKeyStateData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthKeyStateData) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xe0408f17: func() error {
			var flags = dBuf.UInt()
			_ = flags
			m.SetAuthKeyId(dBuf.Long())
			m.SetKeyState(dBuf.Int())
			m.SetUserId(dBuf.Long())
			m.SetAccessHash(dBuf.Long())
			if (flags & (1 << 0)) != 0 {
				m5 := &ClientSession{}
				m5.Decode(dBuf)
				m.SetClient(m5)
			}
			if (flags & (1 << 1)) != 0 {
				m.SetAndroidPushSessionId(&wrapperspb.Int64Value{Value: dBuf.Long()})
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
// ClientSession <--
//  + TL_ClientSession
//

func (m *ClientSession) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_clientSession:
		t := m.To_ClientSession()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
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

// To_ClientSession
func (m *ClientSession) To_ClientSession() *TLClientSession {
	m.PredicateName = Predicate_clientSession
	return &TLClientSession{
		Data2: m,
	}
}

// MakeTLClientSession
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

func (m *TLClientSession) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x9a8e71b0: func() error {
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
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_clientSession, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_clientSession, layer)
		return nil
	}

	return nil
}

func (m *TLClientSession) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLClientSession) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x9a8e71b0: func() error {
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

//----------------------------------------------------------------------------------------------------------------
// TLAuthsessionGetAuthorizations
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetAuthorizations) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x30e21244:
		x.UInt(0x30e21244)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetExcludeAuthKeyId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionGetAuthorizations) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetAuthorizations) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x30e21244:

		// not has flags

		m.UserId = dBuf.Long()
		m.ExcludeAuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLAuthsessionResetAuthorization
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionResetAuthorization) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x8d5f6ca6:
		x.UInt(0x8d5f6ca6)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Long(m.GetHash())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionResetAuthorization) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionResetAuthorization) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x8d5f6ca6:

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

// TLAuthsessionGetLayer
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetLayer) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xa82f16a9:
		x.UInt(0xa82f16a9)

		// no flags

		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionGetLayer) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetLayer) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xa82f16a9:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLAuthsessionGetLangPack
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetLangPack) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x29bbc166:
		x.UInt(0x29bbc166)

		// no flags

		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionGetLangPack) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetLangPack) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x29bbc166:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLAuthsessionGetClient
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetClient) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x605855be:
		x.UInt(0x605855be)

		// no flags

		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionGetClient) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetClient) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x605855be:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLAuthsessionGetLangCode
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetLangCode) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x5899b559:
		x.UInt(0x5899b559)

		// no flags

		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionGetLangCode) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetLangCode) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x5899b559:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLAuthsessionGetUserId
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetUserId) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x57491cac:
		x.UInt(0x57491cac)

		// no flags

		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionGetUserId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetUserId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x57491cac:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLAuthsessionGetPushSessionId
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetPushSessionId) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xb3c23141:
		x.UInt(0xb3c23141)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetTokenType())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionGetPushSessionId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetPushSessionId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb3c23141:

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

// TLAuthsessionGetFutureSalts
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetFutureSalts) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xb8cf5815:
		x.UInt(0xb8cf5815)

		// no flags

		x.Long(m.GetAuthKeyId())
		x.Int(m.GetNum())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionGetFutureSalts) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetFutureSalts) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb8cf5815:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		m.Num = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLAuthsessionQueryAuthKey
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionQueryAuthKey) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x54b73828:
		x.UInt(0x54b73828)

		// no flags

		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionQueryAuthKey) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionQueryAuthKey) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x54b73828:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLAuthsessionSetAuthKey
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionSetAuthKey) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x3e940c91:
		x.UInt(0x3e940c91)

		// no flags

		m.GetAuthKey().Encode(x, layer)
		m.GetFutureSalt().Encode(x, layer)
		x.Int(m.GetExpiresIn())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionSetAuthKey) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionSetAuthKey) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x3e940c91:

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

// TLAuthsessionBindAuthKeyUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionBindAuthKeyUser) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xbce0423:
		x.UInt(0xbce0423)

		// no flags

		x.Long(m.GetAuthKeyId())
		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionBindAuthKeyUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionBindAuthKeyUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xbce0423:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLAuthsessionUnbindAuthKeyUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionUnbindAuthKeyUser) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x758c648:
		x.UInt(0x758c648)

		// no flags

		x.Long(m.GetAuthKeyId())
		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionUnbindAuthKeyUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionUnbindAuthKeyUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x758c648:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLAuthsessionGetPermAuthKeyId
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetPermAuthKeyId) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x907464d6:
		x.UInt(0x907464d6)

		// no flags

		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionGetPermAuthKeyId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetPermAuthKeyId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x907464d6:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLAuthsessionBindTempAuthKey
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionBindTempAuthKey) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x608f4f86:
		x.UInt(0x608f4f86)

		// no flags

		x.Long(m.GetPermAuthKeyId())
		x.Long(m.GetNonce())
		x.Int(m.GetExpiresAt())
		x.StringBytes(m.GetEncryptedMessage())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionBindTempAuthKey) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionBindTempAuthKey) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x608f4f86:

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

// TLAuthsessionSetClientSessionInfo
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionSetClientSessionInfo) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x2d9ff94:
		x.UInt(0x2d9ff94)

		// no flags

		m.GetData().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionSetClientSessionInfo) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionSetClientSessionInfo) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x2d9ff94:

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

// TLAuthsessionGetAuthorization
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetAuthorization) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x6e5e1923:
		x.UInt(0x6e5e1923)

		// no flags

		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionGetAuthorization) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetAuthorization) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x6e5e1923:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLAuthsessionGetAuthStateData
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionGetAuthStateData) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x4f5e3131:
		x.UInt(0x4f5e3131)

		// no flags

		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionGetAuthStateData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionGetAuthStateData) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x4f5e3131:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLAuthsessionSetLayer
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionSetLayer) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x44651485:
		x.UInt(0x44651485)

		// no flags

		x.Long(m.GetAuthKeyId())
		x.String(m.GetIp())
		x.Int(m.GetLayer())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionSetLayer) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionSetLayer) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x44651485:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		m.Ip = dBuf.String()
		m.Layer = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLAuthsessionSetInitConnection
///////////////////////////////////////////////////////////////////////////////

func (m *TLAuthsessionSetInitConnection) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x7cdf8a8c:
		x.UInt(0x7cdf8a8c)

		// no flags

		x.Long(m.GetAuthKeyId())
		x.String(m.GetIp())
		x.Int(m.GetApiId())
		x.String(m.GetDeviceModel())
		x.String(m.GetSystemVersion())
		x.String(m.GetAppVersion())
		x.String(m.GetSystemLangCode())
		x.String(m.GetLangPack())
		x.String(m.GetLangCode())
		x.String(m.GetProxy())
		x.String(m.GetParams())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLAuthsessionSetInitConnection) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthsessionSetInitConnection) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x7cdf8a8c:

		// not has flags

		m.AuthKeyId = dBuf.Long()
		m.Ip = dBuf.String()
		m.ApiId = dBuf.Int()
		m.DeviceModel = dBuf.String()
		m.SystemVersion = dBuf.String()
		m.AppVersion = dBuf.String()
		m.SystemLangCode = dBuf.String()
		m.LangPack = dBuf.String()
		m.LangCode = dBuf.String()
		m.Proxy = dBuf.String()
		m.Params = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
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
