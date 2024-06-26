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

package status

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
	392473649: func() mtproto.TLObject { // 0x1764ac31
		o := MakeTLSessionEntry(nil)
		o.Data2.Constructor = 392473649
		return o
	},
	-269700200: func() mtproto.TLObject { // 0xefecb398
		o := MakeTLUserSessionEntryList(nil)
		o.Data2.Constructor = -269700200
		return o
	},

	// Method
	1381075919: func() mtproto.TLObject { // 0x52518bcf
		return &TLStatusSetSessionOnline{
			Constructor: 1381075919,
		}
	},
	631663196: func() mtproto.TLObject { // 0x25a66a5c
		return &TLStatusSetSessionOffline{
			Constructor: 631663196,
		}
	},
	-406788659: func() mtproto.TLObject { // 0xe7c0e5cd
		return &TLStatusGetUserOnlineSessions{
			Constructor: -406788659,
		}
	},
	-2009385532: func() mtproto.TLObject { // 0x883b35c4
		return &TLStatusGetUsersOnlineSessionsList{
			Constructor: -2009385532,
		}
	},
	1166257237: func() mtproto.TLObject { // 0x4583ac55
		return &TLStatusGetChannelOnlineUsers{
			Constructor: 1166257237,
		}
	},
	-851901363: func() mtproto.TLObject { // 0xcd39044d
		return &TLStatusSetUserChannelsOnline{
			Constructor: -851901363,
		}
	},
	1822646698: func() mtproto.TLObject { // 0x6ca361aa
		return &TLStatusSetUserChannelsOffline{
			Constructor: 1822646698,
		}
	},
	-997471364: func() mtproto.TLObject { // 0xc48bcb7c
		return &TLStatusSetChannelUserOffline{
			Constructor: -997471364,
		}
	},
	-1499734793: func() mtproto.TLObject { // 0xa69bdcf7
		return &TLStatusSetChannelUsersOnline{
			Constructor: -1499734793,
		}
	},
	1266112245: func() mtproto.TLObject { // 0x4b7756f5
		return &TLStatusSetChannelOffline{
			Constructor: 1266112245,
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
// SessionEntry <--
//  + TL_SessionEntry
//

func (m *SessionEntry) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_sessionEntry:
		t := m.To_SessionEntry()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *SessionEntry) CalcByteSize(layer int32) int {
	return 0
}

func (m *SessionEntry) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x1764ac31:
		m2 := MakeTLSessionEntry(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_SessionEntry
func (m *SessionEntry) To_SessionEntry() *TLSessionEntry {
	m.PredicateName = Predicate_sessionEntry
	return &TLSessionEntry{
		Data2: m,
	}
}

// MakeTLSessionEntry
func MakeTLSessionEntry(data2 *SessionEntry) *TLSessionEntry {
	if data2 == nil {
		return &TLSessionEntry{Data2: &SessionEntry{
			PredicateName: Predicate_sessionEntry,
		}}
	} else {
		data2.PredicateName = Predicate_sessionEntry
		return &TLSessionEntry{Data2: data2}
	}
}

func (m *TLSessionEntry) To_SessionEntry() *SessionEntry {
	m.Data2.PredicateName = Predicate_sessionEntry
	return m.Data2
}

func (m *TLSessionEntry) SetUserId(v int64) { m.Data2.UserId = v }
func (m *TLSessionEntry) GetUserId() int64  { return m.Data2.UserId }

func (m *TLSessionEntry) SetAuthKeyId(v int64) { m.Data2.AuthKeyId = v }
func (m *TLSessionEntry) GetAuthKeyId() int64  { return m.Data2.AuthKeyId }

func (m *TLSessionEntry) SetGateway(v string) { m.Data2.Gateway = v }
func (m *TLSessionEntry) GetGateway() string  { return m.Data2.Gateway }

func (m *TLSessionEntry) SetExpired(v int64) { m.Data2.Expired = v }
func (m *TLSessionEntry) GetExpired() int64  { return m.Data2.Expired }

func (m *TLSessionEntry) SetLayer(v int32) { m.Data2.Layer = v }
func (m *TLSessionEntry) GetLayer() int32  { return m.Data2.Layer }

func (m *TLSessionEntry) SetPermAuthKeyId(v int64) { m.Data2.PermAuthKeyId = v }
func (m *TLSessionEntry) GetPermAuthKeyId() int64  { return m.Data2.PermAuthKeyId }

func (m *TLSessionEntry) SetClient(v string) { m.Data2.Client = v }
func (m *TLSessionEntry) GetClient() string  { return m.Data2.Client }

func (m *TLSessionEntry) GetPredicateName() string {
	return Predicate_sessionEntry
}

func (m *TLSessionEntry) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x1764ac31: func() error {
			x.UInt(0x1764ac31)

			x.Long(m.GetUserId())
			x.Long(m.GetAuthKeyId())
			x.String(m.GetGateway())
			x.Long(m.GetExpired())
			x.Int(m.GetLayer())
			x.Long(m.GetPermAuthKeyId())
			x.String(m.GetClient())
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_sessionEntry, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_sessionEntry, layer)
		return nil
	}

	return nil
}

func (m *TLSessionEntry) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionEntry) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x1764ac31: func() error {
			m.SetUserId(dBuf.Long())
			m.SetAuthKeyId(dBuf.Long())
			m.SetGateway(dBuf.String())
			m.SetExpired(dBuf.Long())
			m.SetLayer(dBuf.Int())
			m.SetPermAuthKeyId(dBuf.Long())
			m.SetClient(dBuf.String())
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
// UserSessionEntryList <--
//  + TL_UserSessionEntryList
//

func (m *UserSessionEntryList) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_userSessionEntryList:
		t := m.To_UserSessionEntryList()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *UserSessionEntryList) CalcByteSize(layer int32) int {
	return 0
}

func (m *UserSessionEntryList) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xefecb398:
		m2 := MakeTLUserSessionEntryList(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_UserSessionEntryList
func (m *UserSessionEntryList) To_UserSessionEntryList() *TLUserSessionEntryList {
	m.PredicateName = Predicate_userSessionEntryList
	return &TLUserSessionEntryList{
		Data2: m,
	}
}

// MakeTLUserSessionEntryList
func MakeTLUserSessionEntryList(data2 *UserSessionEntryList) *TLUserSessionEntryList {
	if data2 == nil {
		return &TLUserSessionEntryList{Data2: &UserSessionEntryList{
			PredicateName: Predicate_userSessionEntryList,
		}}
	} else {
		data2.PredicateName = Predicate_userSessionEntryList
		return &TLUserSessionEntryList{Data2: data2}
	}
}

func (m *TLUserSessionEntryList) To_UserSessionEntryList() *UserSessionEntryList {
	m.Data2.PredicateName = Predicate_userSessionEntryList
	return m.Data2
}

func (m *TLUserSessionEntryList) SetUserId(v int64) { m.Data2.UserId = v }
func (m *TLUserSessionEntryList) GetUserId() int64  { return m.Data2.UserId }

func (m *TLUserSessionEntryList) SetUserSessions(v []*SessionEntry) { m.Data2.UserSessions = v }
func (m *TLUserSessionEntryList) GetUserSessions() []*SessionEntry  { return m.Data2.UserSessions }

func (m *TLUserSessionEntryList) GetPredicateName() string {
	return Predicate_userSessionEntryList
}

func (m *TLUserSessionEntryList) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xefecb398: func() error {
			x.UInt(0xefecb398)

			x.Long(m.GetUserId())

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetUserSessions())))
			for _, v := range m.GetUserSessions() {
				v.Encode(x, layer)
			}

			return nil
		},
	}

	clazzId := GetClazzID(Predicate_userSessionEntryList, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_userSessionEntryList, layer)
		return nil
	}

	return nil
}

func (m *TLUserSessionEntryList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSessionEntryList) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xefecb398: func() error {
			m.SetUserId(dBuf.Long())
			c1 := dBuf.Int()
			if c1 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
			}
			l1 := dBuf.Int()
			v1 := make([]*SessionEntry, l1)
			for i := int32(0); i < l1; i++ {
				v1[i] = &SessionEntry{}
				v1[i].Decode(dBuf)
			}
			m.SetUserSessions(v1)

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
// TLStatusSetSessionOnline
///////////////////////////////////////////////////////////////////////////////

func (m *TLStatusSetSessionOnline) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x52518bcf:
		x.UInt(0x52518bcf)

		// no flags

		x.Long(m.GetUserId())
		m.GetSession().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLStatusSetSessionOnline) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLStatusSetSessionOnline) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x52518bcf:

		// not has flags

		m.UserId = dBuf.Long()

		m2 := &SessionEntry{}
		m2.Decode(dBuf)
		m.Session = m2

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLStatusSetSessionOffline
///////////////////////////////////////////////////////////////////////////////

func (m *TLStatusSetSessionOffline) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x25a66a5c:
		x.UInt(0x25a66a5c)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetAuthKeyId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLStatusSetSessionOffline) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLStatusSetSessionOffline) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x25a66a5c:

		// not has flags

		m.UserId = dBuf.Long()
		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLStatusGetUserOnlineSessions
///////////////////////////////////////////////////////////////////////////////

func (m *TLStatusGetUserOnlineSessions) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xe7c0e5cd:
		x.UInt(0xe7c0e5cd)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLStatusGetUserOnlineSessions) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLStatusGetUserOnlineSessions) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe7c0e5cd:

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLStatusGetUsersOnlineSessionsList
///////////////////////////////////////////////////////////////////////////////

func (m *TLStatusGetUsersOnlineSessionsList) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x883b35c4:
		x.UInt(0x883b35c4)

		// no flags

		x.VectorLong(m.GetUsers())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLStatusGetUsersOnlineSessionsList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLStatusGetUsersOnlineSessionsList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x883b35c4:

		// not has flags

		m.Users = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLStatusGetChannelOnlineUsers
///////////////////////////////////////////////////////////////////////////////

func (m *TLStatusGetChannelOnlineUsers) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x4583ac55:
		x.UInt(0x4583ac55)

		// no flags

		x.Long(m.GetChannelId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLStatusGetChannelOnlineUsers) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLStatusGetChannelOnlineUsers) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x4583ac55:

		// not has flags

		m.ChannelId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLStatusSetUserChannelsOnline
///////////////////////////////////////////////////////////////////////////////

func (m *TLStatusSetUserChannelsOnline) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xcd39044d:
		x.UInt(0xcd39044d)

		// no flags

		x.Long(m.GetUserId())

		x.VectorLong(m.GetChannels())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLStatusSetUserChannelsOnline) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLStatusSetUserChannelsOnline) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xcd39044d:

		// not has flags

		m.UserId = dBuf.Long()

		m.Channels = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLStatusSetUserChannelsOffline
///////////////////////////////////////////////////////////////////////////////

func (m *TLStatusSetUserChannelsOffline) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x6ca361aa:
		x.UInt(0x6ca361aa)

		// no flags

		x.Long(m.GetUserId())

		x.VectorLong(m.GetChannels())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLStatusSetUserChannelsOffline) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLStatusSetUserChannelsOffline) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x6ca361aa:

		// not has flags

		m.UserId = dBuf.Long()

		m.Channels = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLStatusSetChannelUserOffline
///////////////////////////////////////////////////////////////////////////////

func (m *TLStatusSetChannelUserOffline) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xc48bcb7c:
		x.UInt(0xc48bcb7c)

		// no flags

		x.Long(m.GetChannelId())
		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLStatusSetChannelUserOffline) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLStatusSetChannelUserOffline) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc48bcb7c:

		// not has flags

		m.ChannelId = dBuf.Long()
		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLStatusSetChannelUsersOnline
///////////////////////////////////////////////////////////////////////////////

func (m *TLStatusSetChannelUsersOnline) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xa69bdcf7:
		x.UInt(0xa69bdcf7)

		// no flags

		x.Long(m.GetChannelId())

		x.VectorLong(m.GetId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLStatusSetChannelUsersOnline) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLStatusSetChannelUsersOnline) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xa69bdcf7:

		// not has flags

		m.ChannelId = dBuf.Long()

		m.Id = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLStatusSetChannelOffline
///////////////////////////////////////////////////////////////////////////////

func (m *TLStatusSetChannelOffline) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x4b7756f5:
		x.UInt(0x4b7756f5)

		// no flags

		x.Long(m.GetChannelId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLStatusSetChannelOffline) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLStatusSetChannelOffline) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x4b7756f5:

		// not has flags

		m.ChannelId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// Vector_UserSessionEntryList
// /////////////////////////////////////////////////////////////////////////////
func (m *Vector_UserSessionEntryList) Encode(x *mtproto.EncodeBuf, layer int32) error {
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		v.Encode(x, layer)
	}

	return nil
}

func (m *Vector_UserSessionEntryList) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*UserSessionEntryList, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(UserSessionEntryList)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_UserSessionEntryList) CalcByteSize(layer int32) int {
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
