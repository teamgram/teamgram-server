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

package user

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
	-1280204321: func() mtproto.TLObject { // 0xb3b1a1df
		o := MakeTLLastSeenData(nil)
		o.Data2.Constructor = -1280204321
		return o
	},
	1894399913: func() mtproto.TLObject { // 0x70ea3fa9
		o := MakeTLPeerPeerNotifySettings(nil)
		o.Data2.Constructor = 1894399913
		return o
	},
	1256160192: func() mtproto.TLObject { // 0x4adf7bc0
		o := MakeTLUserImportedContacts(nil)
		o.Data2.Constructor = 1256160192
		return o
	},
	1067703239: func() mtproto.TLObject { // 0x3fa3dbc7
		o := MakeTLUsersDataFound(nil)
		o.Data2.Constructor = 1067703239
		return o
	},
	-2134594054: func() mtproto.TLObject { // 0x80c4adfa
		o := MakeTLUsersIdFound(nil)
		o.Data2.Constructor = -2134594054
		return o
	},

	// Method
	2090958337: func() mtproto.TLObject { // 0x7ca17e01
		return &TLUserGetLastSeens{
			Constructor: 2090958337,
		}
	},
	-46114259: func() mtproto.TLObject { // 0xfd405a2d
		return &TLUserUpdateLastSeen{
			Constructor: -46114259,
		}
	},
	-1860581154: func() mtproto.TLObject { // 0x9119c8de
		return &TLUserGetLastSeen{
			Constructor: -1860581154,
		}
	},
	929720132: func() mtproto.TLObject { // 0x376a6744
		return &TLUserGetImmutableUser{
			Constructor: 929720132,
		}
	},
	-1657068585: func() mtproto.TLObject { // 0x9d3b23d7
		return &TLUserGetMutableUsers{
			Constructor: -1657068585,
		}
	},
	-373067804: func() mtproto.TLObject { // 0xe9c36fe4
		return &TLUserGetImmutableUserByPhone{
			Constructor: -373067804,
		}
	},
	-12709005: func() mtproto.TLObject { // 0xff3e1373
		return &TLUserGetImmutableUserByToken{
			Constructor: -12709005,
		}
	},
	-766178484: func() mtproto.TLObject { // 0xd2550b4c
		return &TLUserSetAccountDaysTTL{
			Constructor: -766178484,
		}
	},
	-1299956000: func() mtproto.TLObject { // 0xb2843ee0
		return &TLUserGetAccountDaysTTL{
			Constructor: -1299956000,
		}
	},
	1085028198: func() mtproto.TLObject { // 0x40ac3766
		return &TLUserGetNotifySettings{
			Constructor: 1085028198,
		}
	},
	-463137380: func() mtproto.TLObject { // 0xe465159c
		return &TLUserGetNotifySettingsList{
			Constructor: -463137380,
		}
	},
	-907188763: func() mtproto.TLObject { // 0xc9ed65e5
		return &TLUserSetNotifySettings{
			Constructor: -907188763,
		}
	},
	235380084: func() mtproto.TLObject { // 0xe079d74
		return &TLUserResetNotifySettings{
			Constructor: 235380084,
		}
	},
	1435658357: func() mtproto.TLObject { // 0x55926875
		return &TLUserGetAllNotifySettings{
			Constructor: 1435658357,
		}
	},
	2012672274: func() mtproto.TLObject { // 0x77f6f112
		return &TLUserGetGlobalPrivacySettings{
			Constructor: 2012672274,
		}
	},
	-1934257490: func() mtproto.TLObject { // 0x8cb592ae
		return &TLUserSetGlobalPrivacySettings{
			Constructor: -1934257490,
		}
	},
	-1656708172: func() mtproto.TLObject { // 0x9d40a3b4
		return &TLUserGetPrivacy{
			Constructor: -1656708172,
		}
	},
	-2007650929: func() mtproto.TLObject { // 0x8855ad8f
		return &TLUserSetPrivacy{
			Constructor: -2007650929,
		}
	},
	-982638934: func() mtproto.TLObject { // 0xc56e1eaa
		return &TLUserCheckPrivacy{
			Constructor: -982638934,
		}
	},
	-891148445: func() mtproto.TLObject { // 0xcae22763
		return &TLUserAddPeerSettings{
			Constructor: -891148445,
		}
	},
	218296167: func() mtproto.TLObject { // 0xd02ef67
		return &TLUserGetPeerSettings{
			Constructor: 218296167,
		}
	},
	1586043239: func() mtproto.TLObject { // 0x5e891967
		return &TLUserDeletePeerSettings{
			Constructor: 1586043239,
		}
	},
	-8759461: func() mtproto.TLObject { // 0xff7a575b
		return &TLUserChangePhone{
			Constructor: -8759461,
		}
	},
	2044729473: func() mtproto.TLObject { // 0x79e01881
		return &TLUserCreateNewUser{
			Constructor: 2044729473,
		}
	},
	2132777160: func() mtproto.TLObject { // 0x7f1f98c8
		return &TLUserDeleteUser{
			Constructor: 2132777160,
		}
	},
	-2130301264: func() mtproto.TLObject { // 0x81062eb0
		return &TLUserBlockPeer{
			Constructor: -2130301264,
		}
	},
	-555280883: func() mtproto.TLObject { // 0xdee7160d
		return &TLUserUnBlockPeer{
			Constructor: -555280883,
		}
	},
	-1147140722: func() mtproto.TLObject { // 0xbba0058e
		return &TLUserBlockedByUser{
			Constructor: -1147140722,
		}
	},
	-1934708257: func() mtproto.TLObject { // 0x8caeb1df
		return &TLUserIsBlockedByUser{
			Constructor: -1934708257,
		}
	},
	-1006800656: func() mtproto.TLObject { // 0xc3fd70f0
		return &TLUserCheckBlockUserList{
			Constructor: -1006800656,
		}
	},
	603964232: func() mtproto.TLObject { // 0x23ffc348
		return &TLUserGetBlockedList{
			Constructor: 603964232,
		}
	},
	-456010794: func() mtproto.TLObject { // 0xe4d1d3d6
		return &TLUserGetContactSignUpNotification{
			Constructor: -456010794,
		}
	},
	-2053016735: func() mtproto.TLObject { // 0x85a17361
		return &TLUserSetContactSignUpNotification{
			Constructor: -2053016735,
		}
	},
	-1799115361: func() mtproto.TLObject { // 0x94c3ad9f
		return &TLUserGetContentSettings{
			Constructor: -1799115361,
		}
	},
	-1654391189: func() mtproto.TLObject { // 0x9d63fe6b
		return &TLUserSetContentSettings{
			Constructor: -1654391189,
		}
	},
	-972979687: func() mtproto.TLObject { // 0xc6018219
		return &TLUserDeleteContact{
			Constructor: -972979687,
		}
	},
	-951332511: func() mtproto.TLObject { // 0xc74bd161
		return &TLUserGetContactList{
			Constructor: -951332511,
		}
	},
	-237135810: func() mtproto.TLObject { // 0xf1dd983e
		return &TLUserGetContactIdList{
			Constructor: -237135810,
		}
	},
	-613250077: func() mtproto.TLObject { // 0xdb728be3
		return &TLUserGetContact{
			Constructor: -613250077,
		}
	},
	2042936590: func() mtproto.TLObject { // 0x79c4bd0e
		return &TLUserAddContact{
			Constructor: 2042936590,
		}
	},
	-2102962012: func() mtproto.TLObject { // 0x82a758a4
		return &TLUserCheckContact{
			Constructor: -2102962012,
		}
	},
	1202356754: func() mtproto.TLObject { // 0x47aa8212
		return &TLUserGetImportersByPhone{
			Constructor: 1202356754,
		}
	},
	390978644: func() mtproto.TLObject { // 0x174ddc54
		return &TLUserDeleteImportersByPhone{
			Constructor: 390978644,
		}
	},
	-1711212654: func() mtproto.TLObject { // 0x9a00f792
		return &TLUserImportContacts{
			Constructor: -1711212654,
		}
	},
	302016562: func() mtproto.TLObject { // 0x12006832
		return &TLUserGetCountryCode{
			Constructor: 302016562,
		}
	},
	-620695161: func() mtproto.TLObject { // 0xdb00f187
		return &TLUserUpdateAbout{
			Constructor: -620695161,
		}
	},
	-882473492: func() mtproto.TLObject { // 0xcb6685ec
		return &TLUserUpdateFirstAndLastName{
			Constructor: -882473492,
		}
	},
	617163107: func() mtproto.TLObject { // 0x24c92963
		return &TLUserUpdateVerified{
			Constructor: 617163107,
		}
	},
	-179495311: func() mtproto.TLObject { // 0xf54d1e71
		return &TLUserUpdateUsername{
			Constructor: -179495311,
		}
	},
	997461895: func() mtproto.TLObject { // 0x3b740f87
		return &TLUserUpdateProfilePhoto{
			Constructor: 997461895,
		}
	},
	736322062: func() mtproto.TLObject { // 0x2be3620e
		return &TLUserDeleteProfilePhotos{
			Constructor: 736322062,
		}
	},
	-597245626: func() mtproto.TLObject { // 0xdc66c146
		return &TLUserGetProfilePhotos{
			Constructor: -597245626,
		}
	},
	1966844182: func() mtproto.TLObject { // 0x753ba916
		return &TLUserSetBotCommands{
			Constructor: 1966844182,
		}
	},
	-948779026: func() mtproto.TLObject { // 0xc772c7ee
		return &TLUserIsBot{
			Constructor: -948779026,
		}
	},
	879114000: func() mtproto.TLObject { // 0x34663710
		return &TLUserGetBotInfo{
			Constructor: 879114000,
		}
	},
	1935999169: func() mtproto.TLObject { // 0x736500c1
		return &TLUserCheckBots{
			Constructor: 1935999169,
		}
	},
	-49225414: func() mtproto.TLObject { // 0xfd10e13a
		return &TLUserGetFullUser{
			Constructor: -49225414,
		}
	},
	-121062696: func() mtproto.TLObject { // 0xf8c8bad8
		return &TLUserUpdateEmojiStatus{
			Constructor: -121062696,
		}
	},
	62615811: func() mtproto.TLObject { // 0x3bb7103
		return &TLUserGetUserDataById{
			Constructor: 62615811,
		}
	},
	-2121142279: func() mtproto.TLObject { // 0x8191eff9
		return &TLUserGetUserDataListByIdList{
			Constructor: -2121142279,
		}
	},
	1057580446: func() mtproto.TLObject { // 0x3f09659e
		return &TLUserGetUserDataByToken{
			Constructor: 1057580446,
		}
	},
	1882568397: func() mtproto.TLObject { // 0x7035b6cd
		return &TLUserSearch{
			Constructor: 1882568397,
		}
	},
	-1174586898: func() mtproto.TLObject { // 0xb9fd39ee
		return &TLUserUpdateBotData{
			Constructor: -1174586898,
		}
	},
	806009420: func() mtproto.TLObject { // 0x300aba4c
		return &TLUserGetImmutableUserV2{
			Constructor: 806009420,
		}
	},
	-1795585240: func() mtproto.TLObject { // 0x94f98b28
		return &TLUserGetMutableUsersV2{
			Constructor: -1795585240,
		}
	},
	1282329771: func() mtproto.TLObject { // 0x4c6eccab
		return &TLUserCreateNewTestUser{
			Constructor: 1282329771,
		}
	},
	-2044429563: func() mtproto.TLObject { // 0x86247b05
		return &TLUserEditCloseFriends{
			Constructor: -2044429563,
		}
	},
	1391834736: func() mtproto.TLObject { // 0x52f5b670
		return &TLUserSetStoriesMaxId{
			Constructor: 1391834736,
		}
	},
	586812791: func() mtproto.TLObject { // 0x22fa0d77
		return &TLUserSetColor{
			Constructor: 586812791,
		}
	},
	1484434322: func() mtproto.TLObject { // 0x587aab92
		return &TLUserUpdateBirthday{
			Constructor: 1484434322,
		}
	},
	-24199258: func() mtproto.TLObject { // 0xfe8ebfa6
		return &TLUserGetBirthdays{
			Constructor: -24199258,
		}
	},
	-138012584: func() mtproto.TLObject { // 0xf7c61858
		return &TLUserSetStoriesHidden{
			Constructor: -138012584,
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
// LastSeenData <--
//  + TL_LastSeenData
//

func (m *LastSeenData) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_lastSeenData:
		t := m.To_LastSeenData()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *LastSeenData) CalcByteSize(layer int32) int {
	return 0
}

func (m *LastSeenData) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xb3b1a1df:
		m2 := MakeTLLastSeenData(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_LastSeenData
func (m *LastSeenData) To_LastSeenData() *TLLastSeenData {
	m.PredicateName = Predicate_lastSeenData
	return &TLLastSeenData{
		Data2: m,
	}
}

// MakeTLLastSeenData
func MakeTLLastSeenData(data2 *LastSeenData) *TLLastSeenData {
	if data2 == nil {
		return &TLLastSeenData{Data2: &LastSeenData{
			PredicateName: Predicate_lastSeenData,
		}}
	} else {
		data2.PredicateName = Predicate_lastSeenData
		return &TLLastSeenData{Data2: data2}
	}
}

func (m *TLLastSeenData) To_LastSeenData() *LastSeenData {
	m.Data2.PredicateName = Predicate_lastSeenData
	return m.Data2
}

func (m *TLLastSeenData) SetUserId(v int64) { m.Data2.UserId = v }
func (m *TLLastSeenData) GetUserId() int64  { return m.Data2.UserId }

func (m *TLLastSeenData) SetLastSeenAt(v int64) { m.Data2.LastSeenAt = v }
func (m *TLLastSeenData) GetLastSeenAt() int64  { return m.Data2.LastSeenAt }

func (m *TLLastSeenData) SetExpires(v int32) { m.Data2.Expires = v }
func (m *TLLastSeenData) GetExpires() int32  { return m.Data2.Expires }

func (m *TLLastSeenData) GetPredicateName() string {
	return Predicate_lastSeenData
}

func (m *TLLastSeenData) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xb3b1a1df: func() error {
			x.UInt(0xb3b1a1df)

			x.Long(m.GetUserId())
			x.Long(m.GetLastSeenAt())
			x.Int(m.GetExpires())
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_lastSeenData, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_lastSeenData, layer)
		return nil
	}

	return nil
}

func (m *TLLastSeenData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLLastSeenData) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xb3b1a1df: func() error {
			m.SetUserId(dBuf.Long())
			m.SetLastSeenAt(dBuf.Long())
			m.SetExpires(dBuf.Int())
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
// PeerPeerNotifySettings <--
//  + TL_PeerPeerNotifySettings
//

func (m *PeerPeerNotifySettings) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_peerPeerNotifySettings:
		t := m.To_PeerPeerNotifySettings()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *PeerPeerNotifySettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *PeerPeerNotifySettings) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x70ea3fa9:
		m2 := MakeTLPeerPeerNotifySettings(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_PeerPeerNotifySettings
func (m *PeerPeerNotifySettings) To_PeerPeerNotifySettings() *TLPeerPeerNotifySettings {
	m.PredicateName = Predicate_peerPeerNotifySettings
	return &TLPeerPeerNotifySettings{
		Data2: m,
	}
}

// MakeTLPeerPeerNotifySettings
func MakeTLPeerPeerNotifySettings(data2 *PeerPeerNotifySettings) *TLPeerPeerNotifySettings {
	if data2 == nil {
		return &TLPeerPeerNotifySettings{Data2: &PeerPeerNotifySettings{
			PredicateName: Predicate_peerPeerNotifySettings,
		}}
	} else {
		data2.PredicateName = Predicate_peerPeerNotifySettings
		return &TLPeerPeerNotifySettings{Data2: data2}
	}
}

func (m *TLPeerPeerNotifySettings) To_PeerPeerNotifySettings() *PeerPeerNotifySettings {
	m.Data2.PredicateName = Predicate_peerPeerNotifySettings
	return m.Data2
}

func (m *TLPeerPeerNotifySettings) SetPeerType(v int32) { m.Data2.PeerType = v }
func (m *TLPeerPeerNotifySettings) GetPeerType() int32  { return m.Data2.PeerType }

func (m *TLPeerPeerNotifySettings) SetPeerId(v int64) { m.Data2.PeerId = v }
func (m *TLPeerPeerNotifySettings) GetPeerId() int64  { return m.Data2.PeerId }

func (m *TLPeerPeerNotifySettings) SetSettings(v *mtproto.PeerNotifySettings) { m.Data2.Settings = v }
func (m *TLPeerPeerNotifySettings) GetSettings() *mtproto.PeerNotifySettings  { return m.Data2.Settings }

func (m *TLPeerPeerNotifySettings) GetPredicateName() string {
	return Predicate_peerPeerNotifySettings
}

func (m *TLPeerPeerNotifySettings) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x70ea3fa9: func() error {
			x.UInt(0x70ea3fa9)

			x.Int(m.GetPeerType())
			x.Long(m.GetPeerId())
			m.GetSettings().Encode(x, layer)
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_peerPeerNotifySettings, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_peerPeerNotifySettings, layer)
		return nil
	}

	return nil
}

func (m *TLPeerPeerNotifySettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLPeerPeerNotifySettings) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x70ea3fa9: func() error {
			m.SetPeerType(dBuf.Int())
			m.SetPeerId(dBuf.Long())

			m2 := &mtproto.PeerNotifySettings{}
			m2.Decode(dBuf)
			m.SetSettings(m2)

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
// UserImportedContacts <--
//  + TL_UserImportedContacts
//

func (m *UserImportedContacts) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_userImportedContacts:
		t := m.To_UserImportedContacts()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *UserImportedContacts) CalcByteSize(layer int32) int {
	return 0
}

func (m *UserImportedContacts) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x4adf7bc0:
		m2 := MakeTLUserImportedContacts(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_UserImportedContacts
func (m *UserImportedContacts) To_UserImportedContacts() *TLUserImportedContacts {
	m.PredicateName = Predicate_userImportedContacts
	return &TLUserImportedContacts{
		Data2: m,
	}
}

// MakeTLUserImportedContacts
func MakeTLUserImportedContacts(data2 *UserImportedContacts) *TLUserImportedContacts {
	if data2 == nil {
		return &TLUserImportedContacts{Data2: &UserImportedContacts{
			PredicateName: Predicate_userImportedContacts,
		}}
	} else {
		data2.PredicateName = Predicate_userImportedContacts
		return &TLUserImportedContacts{Data2: data2}
	}
}

func (m *TLUserImportedContacts) To_UserImportedContacts() *UserImportedContacts {
	m.Data2.PredicateName = Predicate_userImportedContacts
	return m.Data2
}

func (m *TLUserImportedContacts) SetImported(v []*mtproto.ImportedContact) { m.Data2.Imported = v }
func (m *TLUserImportedContacts) GetImported() []*mtproto.ImportedContact  { return m.Data2.Imported }

func (m *TLUserImportedContacts) SetPopularInvites(v []*mtproto.PopularContact) {
	m.Data2.PopularInvites = v
}
func (m *TLUserImportedContacts) GetPopularInvites() []*mtproto.PopularContact {
	return m.Data2.PopularInvites
}

func (m *TLUserImportedContacts) SetRetryContacts(v []int64) { m.Data2.RetryContacts = v }
func (m *TLUserImportedContacts) GetRetryContacts() []int64  { return m.Data2.RetryContacts }

func (m *TLUserImportedContacts) SetUsers(v []*mtproto.User) { m.Data2.Users = v }
func (m *TLUserImportedContacts) GetUsers() []*mtproto.User  { return m.Data2.Users }

func (m *TLUserImportedContacts) SetUpdateIdList(v []int64) { m.Data2.UpdateIdList = v }
func (m *TLUserImportedContacts) GetUpdateIdList() []int64  { return m.Data2.UpdateIdList }

func (m *TLUserImportedContacts) GetPredicateName() string {
	return Predicate_userImportedContacts
}

func (m *TLUserImportedContacts) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x4adf7bc0: func() error {
			x.UInt(0x4adf7bc0)

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetImported())))
			for _, v := range m.GetImported() {
				v.Encode(x, layer)
			}

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetPopularInvites())))
			for _, v := range m.GetPopularInvites() {
				v.Encode(x, layer)
			}

			x.VectorLong(m.GetRetryContacts())

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetUsers())))
			for _, v := range m.GetUsers() {
				v.Encode(x, layer)
			}

			x.VectorLong(m.GetUpdateIdList())

			return nil
		},
	}

	clazzId := GetClazzID(Predicate_userImportedContacts, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_userImportedContacts, layer)
		return nil
	}

	return nil
}

func (m *TLUserImportedContacts) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserImportedContacts) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x4adf7bc0: func() error {
			c0 := dBuf.Int()
			if c0 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 0, c0)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 0, c0)
			}
			l0 := dBuf.Int()
			v0 := make([]*mtproto.ImportedContact, l0)
			for i := int32(0); i < l0; i++ {
				v0[i] = &mtproto.ImportedContact{}
				v0[i].Decode(dBuf)
			}
			m.SetImported(v0)

			c1 := dBuf.Int()
			if c1 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
			}
			l1 := dBuf.Int()
			v1 := make([]*mtproto.PopularContact, l1)
			for i := int32(0); i < l1; i++ {
				v1[i] = &mtproto.PopularContact{}
				v1[i].Decode(dBuf)
			}
			m.SetPopularInvites(v1)

			m.SetRetryContacts(dBuf.VectorLong())

			c3 := dBuf.Int()
			if c3 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
			}
			l3 := dBuf.Int()
			v3 := make([]*mtproto.User, l3)
			for i := int32(0); i < l3; i++ {
				v3[i] = &mtproto.User{}
				v3[i].Decode(dBuf)
			}
			m.SetUsers(v3)

			m.SetUpdateIdList(dBuf.VectorLong())

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
// UsersFound <--
//  + TL_UsersDataFound
//  + TL_UsersIdFound
//

func (m *UsersFound) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_usersDataFound:
		t := m.To_UsersDataFound()
		t.Encode(x, layer)
	case Predicate_usersIdFound:
		t := m.To_UsersIdFound()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *UsersFound) CalcByteSize(layer int32) int {
	return 0
}

func (m *UsersFound) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x3fa3dbc7:
		m2 := MakeTLUsersDataFound(m)
		m2.Decode(dBuf)
	case 0x80c4adfa:
		m2 := MakeTLUsersIdFound(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_UsersDataFound
func (m *UsersFound) To_UsersDataFound() *TLUsersDataFound {
	m.PredicateName = Predicate_usersDataFound
	return &TLUsersDataFound{
		Data2: m,
	}
}

// To_UsersIdFound
func (m *UsersFound) To_UsersIdFound() *TLUsersIdFound {
	m.PredicateName = Predicate_usersIdFound
	return &TLUsersIdFound{
		Data2: m,
	}
}

// MakeTLUsersDataFound
func MakeTLUsersDataFound(data2 *UsersFound) *TLUsersDataFound {
	if data2 == nil {
		return &TLUsersDataFound{Data2: &UsersFound{
			PredicateName: Predicate_usersDataFound,
		}}
	} else {
		data2.PredicateName = Predicate_usersDataFound
		return &TLUsersDataFound{Data2: data2}
	}
}

func (m *TLUsersDataFound) To_UsersFound() *UsersFound {
	m.Data2.PredicateName = Predicate_usersDataFound
	return m.Data2
}

func (m *TLUsersDataFound) SetCount(v int32) { m.Data2.Count = v }
func (m *TLUsersDataFound) GetCount() int32  { return m.Data2.Count }

func (m *TLUsersDataFound) SetUsers(v []*mtproto.UserData) { m.Data2.Users = v }
func (m *TLUsersDataFound) GetUsers() []*mtproto.UserData  { return m.Data2.Users }

func (m *TLUsersDataFound) SetNextOffset(v string) { m.Data2.NextOffset = v }
func (m *TLUsersDataFound) GetNextOffset() string  { return m.Data2.NextOffset }

func (m *TLUsersDataFound) GetPredicateName() string {
	return Predicate_usersDataFound
}

func (m *TLUsersDataFound) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x3fa3dbc7: func() error {
			x.UInt(0x3fa3dbc7)

			x.Int(m.GetCount())

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetUsers())))
			for _, v := range m.GetUsers() {
				v.Encode(x, layer)
			}

			x.String(m.GetNextOffset())
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_usersDataFound, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_usersDataFound, layer)
		return nil
	}

	return nil
}

func (m *TLUsersDataFound) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUsersDataFound) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x3fa3dbc7: func() error {
			m.SetCount(dBuf.Int())
			c1 := dBuf.Int()
			if c1 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
			}
			l1 := dBuf.Int()
			v1 := make([]*mtproto.UserData, l1)
			for i := int32(0); i < l1; i++ {
				v1[i] = &mtproto.UserData{}
				v1[i].Decode(dBuf)
			}
			m.SetUsers(v1)

			m.SetNextOffset(dBuf.String())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

// MakeTLUsersIdFound
func MakeTLUsersIdFound(data2 *UsersFound) *TLUsersIdFound {
	if data2 == nil {
		return &TLUsersIdFound{Data2: &UsersFound{
			PredicateName: Predicate_usersIdFound,
		}}
	} else {
		data2.PredicateName = Predicate_usersIdFound
		return &TLUsersIdFound{Data2: data2}
	}
}

func (m *TLUsersIdFound) To_UsersFound() *UsersFound {
	m.Data2.PredicateName = Predicate_usersIdFound
	return m.Data2
}

func (m *TLUsersIdFound) SetIdList(v []int64) { m.Data2.IdList = v }
func (m *TLUsersIdFound) GetIdList() []int64  { return m.Data2.IdList }

func (m *TLUsersIdFound) GetPredicateName() string {
	return Predicate_usersIdFound
}

func (m *TLUsersIdFound) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x80c4adfa: func() error {
			x.UInt(0x80c4adfa)

			x.VectorLong(m.GetIdList())

			return nil
		},
	}

	clazzId := GetClazzID(Predicate_usersIdFound, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_usersIdFound, layer)
		return nil
	}

	return nil
}

func (m *TLUsersIdFound) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUsersIdFound) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x80c4adfa: func() error {

			m.SetIdList(dBuf.VectorLong())

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
// TLUserGetLastSeens
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetLastSeens) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x7ca17e01:
		x.UInt(0x7ca17e01)

		// no flags

		x.VectorLong(m.GetId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetLastSeens) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetLastSeens) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x7ca17e01:

		// not has flags

		m.Id = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserUpdateLastSeen
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUpdateLastSeen) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xfd405a2d:
		x.UInt(0xfd405a2d)

		// no flags

		x.Long(m.GetId())
		x.Long(m.GetLastSeenAt())
		x.Int(m.GetExpires())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserUpdateLastSeen) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUpdateLastSeen) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xfd405a2d:

		// not has flags

		m.Id = dBuf.Long()
		m.LastSeenAt = dBuf.Long()
		m.Expires = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetLastSeen
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetLastSeen) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x9119c8de:
		x.UInt(0x9119c8de)

		// no flags

		x.Long(m.GetId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetLastSeen) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetLastSeen) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x9119c8de:

		// not has flags

		m.Id = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetImmutableUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetImmutableUser) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x376a6744:
		x.UInt(0x376a6744)

		// set flags
		var flags uint32 = 0

		if m.GetPrivacy() == true {
			flags |= 1 << 1
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetId())

		x.VectorLong(m.GetContacts())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetImmutableUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetImmutableUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x376a6744:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.Id = dBuf.Long()
		if (flags & (1 << 1)) != 0 {
			m.Privacy = true
		}

		m.Contacts = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetMutableUsers
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetMutableUsers) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x9d3b23d7:
		x.UInt(0x9d3b23d7)

		// no flags

		x.VectorLong(m.GetId())

		x.VectorLong(m.GetTo())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetMutableUsers) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetMutableUsers) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x9d3b23d7:

		// not has flags

		m.Id = dBuf.VectorLong()

		m.To = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetImmutableUserByPhone
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetImmutableUserByPhone) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xe9c36fe4:
		x.UInt(0xe9c36fe4)

		// no flags

		x.String(m.GetPhone())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetImmutableUserByPhone) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetImmutableUserByPhone) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe9c36fe4:

		// not has flags

		m.Phone = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetImmutableUserByToken
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetImmutableUserByToken) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xff3e1373:
		x.UInt(0xff3e1373)

		// no flags

		x.String(m.GetToken())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetImmutableUserByToken) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetImmutableUserByToken) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xff3e1373:

		// not has flags

		m.Token = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserSetAccountDaysTTL
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserSetAccountDaysTTL) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xd2550b4c:
		x.UInt(0xd2550b4c)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetTtl())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserSetAccountDaysTTL) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSetAccountDaysTTL) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xd2550b4c:

		// not has flags

		m.UserId = dBuf.Long()
		m.Ttl = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetAccountDaysTTL
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetAccountDaysTTL) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xb2843ee0:
		x.UInt(0xb2843ee0)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetAccountDaysTTL) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetAccountDaysTTL) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb2843ee0:

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetNotifySettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetNotifySettings) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x40ac3766:
		x.UInt(0x40ac3766)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetNotifySettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetNotifySettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x40ac3766:

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

// TLUserGetNotifySettingsList
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetNotifySettingsList) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xe465159c:
		x.UInt(0xe465159c)

		// no flags

		x.Long(m.GetUserId())

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetPeers())))
		for _, v := range m.GetPeers() {
			v.Encode(x, layer)
		}

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetNotifySettingsList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetNotifySettingsList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe465159c:

		// not has flags

		m.UserId = dBuf.Long()
		c2 := dBuf.Int()
		if c2 != int32(mtproto.CRC32_vector) {
			// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 2, c2)
			return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 2, c2)
		}
		l2 := dBuf.Int()
		v2 := make([]*mtproto.PeerUtil, l2)
		for i := int32(0); i < l2; i++ {
			v2[i] = &mtproto.PeerUtil{}
			v2[i].Decode(dBuf)
		}
		m.Peers = v2

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserSetNotifySettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserSetNotifySettings) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xc9ed65e5:
		x.UInt(0xc9ed65e5)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		m.GetSettings().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserSetNotifySettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSetNotifySettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc9ed65e5:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()

		m4 := &mtproto.PeerNotifySettings{}
		m4.Decode(dBuf)
		m.Settings = m4

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserResetNotifySettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserResetNotifySettings) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xe079d74:
		x.UInt(0xe079d74)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserResetNotifySettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserResetNotifySettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe079d74:

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetAllNotifySettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetAllNotifySettings) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x55926875:
		x.UInt(0x55926875)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetAllNotifySettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetAllNotifySettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x55926875:

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetGlobalPrivacySettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetGlobalPrivacySettings) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x77f6f112:
		x.UInt(0x77f6f112)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetGlobalPrivacySettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetGlobalPrivacySettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x77f6f112:

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserSetGlobalPrivacySettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserSetGlobalPrivacySettings) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x8cb592ae:
		x.UInt(0x8cb592ae)

		// no flags

		x.Long(m.GetUserId())
		m.GetSettings().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserSetGlobalPrivacySettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSetGlobalPrivacySettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x8cb592ae:

		// not has flags

		m.UserId = dBuf.Long()

		m2 := &mtproto.GlobalPrivacySettings{}
		m2.Decode(dBuf)
		m.Settings = m2

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetPrivacy
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetPrivacy) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x9d40a3b4:
		x.UInt(0x9d40a3b4)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetKeyType())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetPrivacy) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetPrivacy) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x9d40a3b4:

		// not has flags

		m.UserId = dBuf.Long()
		m.KeyType = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserSetPrivacy
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserSetPrivacy) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x8855ad8f:
		x.UInt(0x8855ad8f)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetKeyType())

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetRules())))
		for _, v := range m.GetRules() {
			v.Encode(x, layer)
		}

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserSetPrivacy) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSetPrivacy) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x8855ad8f:

		// not has flags

		m.UserId = dBuf.Long()
		m.KeyType = dBuf.Int()
		c3 := dBuf.Int()
		if c3 != int32(mtproto.CRC32_vector) {
			// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
			return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
		}
		l3 := dBuf.Int()
		v3 := make([]*mtproto.PrivacyRule, l3)
		for i := int32(0); i < l3; i++ {
			v3[i] = &mtproto.PrivacyRule{}
			v3[i].Decode(dBuf)
		}
		m.Rules = v3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserCheckPrivacy
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserCheckPrivacy) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xc56e1eaa:
		x.UInt(0xc56e1eaa)

		// set flags
		var flags uint32 = 0

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetUserId())
		x.Int(m.GetKeyType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserCheckPrivacy) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserCheckPrivacy) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc56e1eaa:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.UserId = dBuf.Long()
		m.KeyType = dBuf.Int()
		m.PeerId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserAddPeerSettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserAddPeerSettings) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xcae22763:
		x.UInt(0xcae22763)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		m.GetSettings().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserAddPeerSettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserAddPeerSettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xcae22763:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()

		m4 := &mtproto.PeerSettings{}
		m4.Decode(dBuf)
		m.Settings = m4

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetPeerSettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetPeerSettings) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xd02ef67:
		x.UInt(0xd02ef67)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetPeerSettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetPeerSettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xd02ef67:

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

// TLUserDeletePeerSettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserDeletePeerSettings) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x5e891967:
		x.UInt(0x5e891967)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserDeletePeerSettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserDeletePeerSettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x5e891967:

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

// TLUserChangePhone
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserChangePhone) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xff7a575b:
		x.UInt(0xff7a575b)

		// no flags

		x.Long(m.GetUserId())
		x.String(m.GetPhone())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserChangePhone) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserChangePhone) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xff7a575b:

		// not has flags

		m.UserId = dBuf.Long()
		m.Phone = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserCreateNewUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserCreateNewUser) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x79e01881:
		x.UInt(0x79e01881)

		// no flags

		x.Long(m.GetSecretKeyId())
		x.String(m.GetPhone())
		x.String(m.GetCountryCode())
		x.String(m.GetFirstName())
		x.String(m.GetLastName())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserCreateNewUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserCreateNewUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x79e01881:

		// not has flags

		m.SecretKeyId = dBuf.Long()
		m.Phone = dBuf.String()
		m.CountryCode = dBuf.String()
		m.FirstName = dBuf.String()
		m.LastName = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserDeleteUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserDeleteUser) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x7f1f98c8:
		x.UInt(0x7f1f98c8)

		// no flags

		x.Long(m.GetUserId())
		x.String(m.GetReason())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserDeleteUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserDeleteUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x7f1f98c8:

		// not has flags

		m.UserId = dBuf.Long()
		m.Reason = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserBlockPeer
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserBlockPeer) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x81062eb0:
		x.UInt(0x81062eb0)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserBlockPeer) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserBlockPeer) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x81062eb0:

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

// TLUserUnBlockPeer
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUnBlockPeer) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xdee7160d:
		x.UInt(0xdee7160d)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserUnBlockPeer) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUnBlockPeer) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xdee7160d:

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

// TLUserBlockedByUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserBlockedByUser) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xbba0058e:
		x.UInt(0xbba0058e)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetPeerUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserBlockedByUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserBlockedByUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xbba0058e:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerUserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserIsBlockedByUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserIsBlockedByUser) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x8caeb1df:
		x.UInt(0x8caeb1df)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetPeerUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserIsBlockedByUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserIsBlockedByUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x8caeb1df:

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerUserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserCheckBlockUserList
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserCheckBlockUserList) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xc3fd70f0:
		x.UInt(0xc3fd70f0)

		// no flags

		x.Long(m.GetUserId())

		x.VectorLong(m.GetId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserCheckBlockUserList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserCheckBlockUserList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc3fd70f0:

		// not has flags

		m.UserId = dBuf.Long()

		m.Id = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetBlockedList
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetBlockedList) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x23ffc348:
		x.UInt(0x23ffc348)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetOffset())
		x.Int(m.GetLimit())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetBlockedList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetBlockedList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x23ffc348:

		// not has flags

		m.UserId = dBuf.Long()
		m.Offset = dBuf.Int()
		m.Limit = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetContactSignUpNotification
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetContactSignUpNotification) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xe4d1d3d6:
		x.UInt(0xe4d1d3d6)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetContactSignUpNotification) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetContactSignUpNotification) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe4d1d3d6:

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserSetContactSignUpNotification
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserSetContactSignUpNotification) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x85a17361:
		x.UInt(0x85a17361)

		// no flags

		x.Long(m.GetUserId())
		m.GetSilent().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserSetContactSignUpNotification) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSetContactSignUpNotification) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x85a17361:

		// not has flags

		m.UserId = dBuf.Long()

		m2 := &mtproto.Bool{}
		m2.Decode(dBuf)
		m.Silent = m2

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetContentSettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetContentSettings) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x94c3ad9f:
		x.UInt(0x94c3ad9f)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetContentSettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetContentSettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x94c3ad9f:

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserSetContentSettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserSetContentSettings) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x9d63fe6b:
		x.UInt(0x9d63fe6b)

		// set flags
		var flags uint32 = 0

		if m.GetSensitiveEnabled() == true {
			flags |= 1 << 0
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserSetContentSettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSetContentSettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x9d63fe6b:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.UserId = dBuf.Long()
		if (flags & (1 << 0)) != 0 {
			m.SensitiveEnabled = true
		}
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserDeleteContact
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserDeleteContact) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xc6018219:
		x.UInt(0xc6018219)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserDeleteContact) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserDeleteContact) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc6018219:

		// not has flags

		m.UserId = dBuf.Long()
		m.Id = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetContactList
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetContactList) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xc74bd161:
		x.UInt(0xc74bd161)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetContactList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetContactList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc74bd161:

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetContactIdList
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetContactIdList) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xf1dd983e:
		x.UInt(0xf1dd983e)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetContactIdList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetContactIdList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf1dd983e:

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetContact
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetContact) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xdb728be3:
		x.UInt(0xdb728be3)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetContact) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetContact) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xdb728be3:

		// not has flags

		m.UserId = dBuf.Long()
		m.Id = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserAddContact
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserAddContact) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x79c4bd0e:
		x.UInt(0x79c4bd0e)

		// no flags

		x.Long(m.GetUserId())
		m.GetAddPhonePrivacyException().Encode(x, layer)
		x.Long(m.GetId())
		x.String(m.GetFirstName())
		x.String(m.GetLastName())
		x.String(m.GetPhone())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserAddContact) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserAddContact) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x79c4bd0e:

		// not has flags

		m.UserId = dBuf.Long()

		m2 := &mtproto.Bool{}
		m2.Decode(dBuf)
		m.AddPhonePrivacyException = m2

		m.Id = dBuf.Long()
		m.FirstName = dBuf.String()
		m.LastName = dBuf.String()
		m.Phone = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserCheckContact
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserCheckContact) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x82a758a4:
		x.UInt(0x82a758a4)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserCheckContact) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserCheckContact) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x82a758a4:

		// not has flags

		m.UserId = dBuf.Long()
		m.Id = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetImportersByPhone
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetImportersByPhone) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x47aa8212:
		x.UInt(0x47aa8212)

		// no flags

		x.String(m.GetPhone())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetImportersByPhone) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetImportersByPhone) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x47aa8212:

		// not has flags

		m.Phone = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserDeleteImportersByPhone
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserDeleteImportersByPhone) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x174ddc54:
		x.UInt(0x174ddc54)

		// no flags

		x.String(m.GetPhone())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserDeleteImportersByPhone) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserDeleteImportersByPhone) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x174ddc54:

		// not has flags

		m.Phone = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserImportContacts
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserImportContacts) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x9a00f792:
		x.UInt(0x9a00f792)

		// no flags

		x.Long(m.GetUserId())

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetContacts())))
		for _, v := range m.GetContacts() {
			v.Encode(x, layer)
		}

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserImportContacts) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserImportContacts) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x9a00f792:

		// not has flags

		m.UserId = dBuf.Long()
		c2 := dBuf.Int()
		if c2 != int32(mtproto.CRC32_vector) {
			// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 2, c2)
			return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 2, c2)
		}
		l2 := dBuf.Int()
		v2 := make([]*mtproto.InputContact, l2)
		for i := int32(0); i < l2; i++ {
			v2[i] = &mtproto.InputContact{}
			v2[i].Decode(dBuf)
		}
		m.Contacts = v2

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetCountryCode
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetCountryCode) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x12006832:
		x.UInt(0x12006832)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetCountryCode) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetCountryCode) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x12006832:

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserUpdateAbout
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUpdateAbout) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xdb00f187:
		x.UInt(0xdb00f187)

		// no flags

		x.Long(m.GetUserId())
		x.String(m.GetAbout())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserUpdateAbout) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUpdateAbout) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xdb00f187:

		// not has flags

		m.UserId = dBuf.Long()
		m.About = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserUpdateFirstAndLastName
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUpdateFirstAndLastName) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xcb6685ec:
		x.UInt(0xcb6685ec)

		// no flags

		x.Long(m.GetUserId())
		x.String(m.GetFirstName())
		x.String(m.GetLastName())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserUpdateFirstAndLastName) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUpdateFirstAndLastName) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xcb6685ec:

		// not has flags

		m.UserId = dBuf.Long()
		m.FirstName = dBuf.String()
		m.LastName = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserUpdateVerified
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUpdateVerified) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x24c92963:
		x.UInt(0x24c92963)

		// no flags

		x.Long(m.GetUserId())
		m.GetVerified().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserUpdateVerified) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUpdateVerified) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x24c92963:

		// not has flags

		m.UserId = dBuf.Long()

		m2 := &mtproto.Bool{}
		m2.Decode(dBuf)
		m.Verified = m2

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserUpdateUsername
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUpdateUsername) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xf54d1e71:
		x.UInt(0xf54d1e71)

		// no flags

		x.Long(m.GetUserId())
		x.String(m.GetUsername())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserUpdateUsername) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUpdateUsername) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf54d1e71:

		// not has flags

		m.UserId = dBuf.Long()
		m.Username = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserUpdateProfilePhoto
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUpdateProfilePhoto) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x3b740f87:
		x.UInt(0x3b740f87)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserUpdateProfilePhoto) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUpdateProfilePhoto) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x3b740f87:

		// not has flags

		m.UserId = dBuf.Long()
		m.Id = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserDeleteProfilePhotos
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserDeleteProfilePhotos) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x2be3620e:
		x.UInt(0x2be3620e)

		// no flags

		x.Long(m.GetUserId())

		x.VectorLong(m.GetId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserDeleteProfilePhotos) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserDeleteProfilePhotos) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x2be3620e:

		// not has flags

		m.UserId = dBuf.Long()

		m.Id = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetProfilePhotos
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetProfilePhotos) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xdc66c146:
		x.UInt(0xdc66c146)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetProfilePhotos) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetProfilePhotos) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xdc66c146:

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserSetBotCommands
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserSetBotCommands) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x753ba916:
		x.UInt(0x753ba916)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetBotId())

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetCommands())))
		for _, v := range m.GetCommands() {
			v.Encode(x, layer)
		}

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserSetBotCommands) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSetBotCommands) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x753ba916:

		// not has flags

		m.UserId = dBuf.Long()
		m.BotId = dBuf.Long()
		c3 := dBuf.Int()
		if c3 != int32(mtproto.CRC32_vector) {
			// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
			return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
		}
		l3 := dBuf.Int()
		v3 := make([]*mtproto.BotCommand, l3)
		for i := int32(0); i < l3; i++ {
			v3[i] = &mtproto.BotCommand{}
			v3[i].Decode(dBuf)
		}
		m.Commands = v3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserIsBot
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserIsBot) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xc772c7ee:
		x.UInt(0xc772c7ee)

		// no flags

		x.Long(m.GetId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserIsBot) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserIsBot) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc772c7ee:

		// not has flags

		m.Id = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetBotInfo
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetBotInfo) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x34663710:
		x.UInt(0x34663710)

		// no flags

		x.Long(m.GetBotId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetBotInfo) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetBotInfo) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x34663710:

		// not has flags

		m.BotId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserCheckBots
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserCheckBots) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x736500c1:
		x.UInt(0x736500c1)

		// no flags

		x.VectorLong(m.GetId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserCheckBots) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserCheckBots) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x736500c1:

		// not has flags

		m.Id = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetFullUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetFullUser) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xfd10e13a:
		x.UInt(0xfd10e13a)

		// no flags

		x.Long(m.GetSelfUserId())
		x.Long(m.GetId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetFullUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetFullUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xfd10e13a:

		// not has flags

		m.SelfUserId = dBuf.Long()
		m.Id = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserUpdateEmojiStatus
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUpdateEmojiStatus) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xf8c8bad8:
		x.UInt(0xf8c8bad8)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetEmojiStatusDocumentId())
		x.Int(m.GetEmojiStatusUntil())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserUpdateEmojiStatus) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUpdateEmojiStatus) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf8c8bad8:

		// not has flags

		m.UserId = dBuf.Long()
		m.EmojiStatusDocumentId = dBuf.Long()
		m.EmojiStatusUntil = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetUserDataById
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetUserDataById) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x3bb7103:
		x.UInt(0x3bb7103)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetUserDataById) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetUserDataById) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x3bb7103:

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetUserDataListByIdList
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetUserDataListByIdList) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x8191eff9:
		x.UInt(0x8191eff9)

		// no flags

		x.VectorLong(m.GetUserIdList())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetUserDataListByIdList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetUserDataListByIdList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x8191eff9:

		// not has flags

		m.UserIdList = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetUserDataByToken
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetUserDataByToken) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x3f09659e:
		x.UInt(0x3f09659e)

		// no flags

		x.String(m.GetToken())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetUserDataByToken) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetUserDataByToken) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x3f09659e:

		// not has flags

		m.Token = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserSearch
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserSearch) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x7035b6cd:
		x.UInt(0x7035b6cd)

		// no flags

		x.String(m.GetQ())

		x.VectorLong(m.GetExcludedContacts())

		x.Long(m.GetOffset())
		x.Int(m.GetLimit())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserSearch) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSearch) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x7035b6cd:

		// not has flags

		m.Q = dBuf.String()

		m.ExcludedContacts = dBuf.VectorLong()

		m.Offset = dBuf.Long()
		m.Limit = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserUpdateBotData
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUpdateBotData) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xb9fd39ee:
		x.UInt(0xb9fd39ee)

		// set flags
		var flags uint32 = 0

		if m.GetBotChatHistory() != nil {
			flags |= 1 << 15
		}
		if m.GetBotNochats() != nil {
			flags |= 1 << 16
		}
		if m.GetBotInlineGeo() != nil {
			flags |= 1 << 21
		}
		if m.GetBotAttachMenu() != nil {
			flags |= 1 << 27
		}
		if m.GetBotInlinePlaceholder() != nil {
			flags |= 1 << 19
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetBotId())
		if m.GetBotChatHistory() != nil {
			m.GetBotChatHistory().Encode(x, layer)
		}

		if m.GetBotNochats() != nil {
			m.GetBotNochats().Encode(x, layer)
		}

		if m.GetBotInlineGeo() != nil {
			m.GetBotInlineGeo().Encode(x, layer)
		}

		if m.GetBotAttachMenu() != nil {
			m.GetBotAttachMenu().Encode(x, layer)
		}

		if m.GetBotInlinePlaceholder() != nil {
			m.GetBotInlinePlaceholder().Encode(x, layer)
		}

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserUpdateBotData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUpdateBotData) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb9fd39ee:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.BotId = dBuf.Long()
		if (flags & (1 << 15)) != 0 {
			m3 := &mtproto.Bool{}
			m3.Decode(dBuf)
			m.BotChatHistory = m3
		}
		if (flags & (1 << 16)) != 0 {
			m4 := &mtproto.Bool{}
			m4.Decode(dBuf)
			m.BotNochats = m4
		}
		if (flags & (1 << 21)) != 0 {
			m5 := &mtproto.Bool{}
			m5.Decode(dBuf)
			m.BotInlineGeo = m5
		}
		if (flags & (1 << 27)) != 0 {
			m6 := &mtproto.Bool{}
			m6.Decode(dBuf)
			m.BotAttachMenu = m6
		}
		if (flags & (1 << 19)) != 0 {
			m7 := &mtproto.Bool{}
			m7.Decode(dBuf)
			m.BotInlinePlaceholder = m7
		}
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetImmutableUserV2
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetImmutableUserV2) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x300aba4c:
		x.UInt(0x300aba4c)

		// set flags
		var flags uint32 = 0

		if m.GetPrivacy() == true {
			flags |= 1 << 0
		}
		if m.GetHasTo() == true {
			flags |= 1 << 2
		}
		if m.GetTo() != nil {
			flags |= 1 << 2
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetId())
		if m.GetTo() != nil {
			x.VectorLong(m.GetTo())
		}

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetImmutableUserV2) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetImmutableUserV2) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x300aba4c:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.Id = dBuf.Long()
		if (flags & (1 << 0)) != 0 {
			m.Privacy = true
		}
		if (flags & (1 << 2)) != 0 {
			m.HasTo = true
		}
		if (flags & (1 << 2)) != 0 {
			m.To = dBuf.VectorLong()
		}
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetMutableUsersV2
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetMutableUsersV2) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x94f98b28:
		x.UInt(0x94f98b28)

		// set flags
		var flags uint32 = 0

		if m.GetPrivacy() == true {
			flags |= 1 << 0
		}
		if m.GetHasTo() == true {
			flags |= 1 << 2
		}
		if m.GetTo() != nil {
			flags |= 1 << 2
		}

		x.UInt(flags)

		// flags Debug by @benqi

		x.VectorLong(m.GetId())

		if m.GetTo() != nil {
			x.VectorLong(m.GetTo())
		}

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetMutableUsersV2) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetMutableUsersV2) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x94f98b28:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi

		m.Id = dBuf.VectorLong()

		if (flags & (1 << 0)) != 0 {
			m.Privacy = true
		}
		if (flags & (1 << 2)) != 0 {
			m.HasTo = true
		}
		if (flags & (1 << 2)) != 0 {
			m.To = dBuf.VectorLong()
		}
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserCreateNewTestUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserCreateNewTestUser) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x4c6eccab:
		x.UInt(0x4c6eccab)

		// no flags

		x.Long(m.GetSecretKeyId())
		x.Long(m.GetMinId())
		x.Long(m.GetMaxId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserCreateNewTestUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserCreateNewTestUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x4c6eccab:

		// not has flags

		m.SecretKeyId = dBuf.Long()
		m.MinId = dBuf.Long()
		m.MaxId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserEditCloseFriends
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserEditCloseFriends) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x86247b05:
		x.UInt(0x86247b05)

		// no flags

		x.Long(m.GetUserId())

		x.VectorLong(m.GetId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserEditCloseFriends) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserEditCloseFriends) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x86247b05:

		// not has flags

		m.UserId = dBuf.Long()

		m.Id = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserSetStoriesMaxId
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserSetStoriesMaxId) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x52f5b670:
		x.UInt(0x52f5b670)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserSetStoriesMaxId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSetStoriesMaxId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x52f5b670:

		// not has flags

		m.UserId = dBuf.Long()
		m.Id = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserSetColor
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserSetColor) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x22fa0d77:
		x.UInt(0x22fa0d77)

		// set flags
		var flags uint32 = 0

		if m.GetForProfile() == true {
			flags |= 1 << 1
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetUserId())
		x.Int(m.GetColor())
		x.Long(m.GetBackgroundEmojiId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserSetColor) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSetColor) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x22fa0d77:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.UserId = dBuf.Long()
		if (flags & (1 << 1)) != 0 {
			m.ForProfile = true
		}
		m.Color = dBuf.Int()
		m.BackgroundEmojiId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserUpdateBirthday
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUpdateBirthday) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x587aab92:
		x.UInt(0x587aab92)

		// set flags
		var flags uint32 = 0

		if m.GetBirthday() != nil {
			flags |= 1 << 1
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetUserId())
		if m.GetBirthday() != nil {
			m.GetBirthday().Encode(x, layer)
		}

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserUpdateBirthday) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUpdateBirthday) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x587aab92:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.UserId = dBuf.Long()
		if (flags & (1 << 1)) != 0 {
			m3 := &mtproto.Birthday{}
			m3.Decode(dBuf)
			m.Birthday = m3
		}
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserGetBirthdays
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetBirthdays) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xfe8ebfa6:
		x.UInt(0xfe8ebfa6)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserGetBirthdays) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetBirthdays) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xfe8ebfa6:

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLUserSetStoriesHidden
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserSetStoriesHidden) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xf7c61858:
		x.UInt(0xf7c61858)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetId())
		m.GetHidden().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLUserSetStoriesHidden) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSetStoriesHidden) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf7c61858:

		// not has flags

		m.UserId = dBuf.Long()
		m.Id = dBuf.Long()

		m3 := &mtproto.Bool{}
		m3.Decode(dBuf)
		m.Hidden = m3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// Vector_LastSeenData
// /////////////////////////////////////////////////////////////////////////////
func (m *Vector_LastSeenData) Encode(x *mtproto.EncodeBuf, layer int32) error {
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		v.Encode(x, layer)
	}

	return nil
}

func (m *Vector_LastSeenData) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*LastSeenData, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(LastSeenData)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_LastSeenData) CalcByteSize(layer int32) int {
	return 0
}

// Vector_ImmutableUser
// /////////////////////////////////////////////////////////////////////////////
func (m *Vector_ImmutableUser) Encode(x *mtproto.EncodeBuf, layer int32) error {
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		v.Encode(x, layer)
	}

	return nil
}

func (m *Vector_ImmutableUser) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*mtproto.ImmutableUser, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(mtproto.ImmutableUser)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_ImmutableUser) CalcByteSize(layer int32) int {
	return 0
}

// Vector_PeerPeerNotifySettings
// /////////////////////////////////////////////////////////////////////////////
func (m *Vector_PeerPeerNotifySettings) Encode(x *mtproto.EncodeBuf, layer int32) error {
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		v.Encode(x, layer)
	}

	return nil
}

func (m *Vector_PeerPeerNotifySettings) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*PeerPeerNotifySettings, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(PeerPeerNotifySettings)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_PeerPeerNotifySettings) CalcByteSize(layer int32) int {
	return 0
}

// Vector_PrivacyRule
// /////////////////////////////////////////////////////////////////////////////
func (m *Vector_PrivacyRule) Encode(x *mtproto.EncodeBuf, layer int32) error {
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		v.Encode(x, layer)
	}

	return nil
}

func (m *Vector_PrivacyRule) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*mtproto.PrivacyRule, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(mtproto.PrivacyRule)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_PrivacyRule) CalcByteSize(layer int32) int {
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

// Vector_PeerBlocked
// /////////////////////////////////////////////////////////////////////////////
func (m *Vector_PeerBlocked) Encode(x *mtproto.EncodeBuf, layer int32) error {
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		v.Encode(x, layer)
	}

	return nil
}

func (m *Vector_PeerBlocked) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*mtproto.PeerBlocked, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(mtproto.PeerBlocked)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_PeerBlocked) CalcByteSize(layer int32) int {
	return 0
}

// Vector_ContactData
// /////////////////////////////////////////////////////////////////////////////
func (m *Vector_ContactData) Encode(x *mtproto.EncodeBuf, layer int32) error {
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		v.Encode(x, layer)
	}

	return nil
}

func (m *Vector_ContactData) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*mtproto.ContactData, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(mtproto.ContactData)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_ContactData) CalcByteSize(layer int32) int {
	return 0
}

// Vector_InputContact
// /////////////////////////////////////////////////////////////////////////////
func (m *Vector_InputContact) Encode(x *mtproto.EncodeBuf, layer int32) error {
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		v.Encode(x, layer)
	}

	return nil
}

func (m *Vector_InputContact) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*mtproto.InputContact, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(mtproto.InputContact)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_InputContact) CalcByteSize(layer int32) int {
	return 0
}

// Vector_UserData
// /////////////////////////////////////////////////////////////////////////////
func (m *Vector_UserData) Encode(x *mtproto.EncodeBuf, layer int32) error {
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		v.Encode(x, layer)
	}

	return nil
}

func (m *Vector_UserData) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*mtproto.UserData, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(mtproto.UserData)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_UserData) CalcByteSize(layer int32) int {
	return 0
}

// Vector_ContactBirthday
// /////////////////////////////////////////////////////////////////////////////
func (m *Vector_ContactBirthday) Encode(x *mtproto.EncodeBuf, layer int32) error {
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		v.Encode(x, layer)
	}

	return nil
}

func (m *Vector_ContactBirthday) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*mtproto.ContactBirthday, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(mtproto.ContactBirthday)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_ContactBirthday) CalcByteSize(layer int32) int {
	return 0
}
