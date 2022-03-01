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

package user

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
	2138633749: func() mtproto.TLObject { // 0x7f78f615
		o := MakeTLUserData(nil)
		o.Data2.Constructor = 2138633749
		return o
	},
	361114766: func() mtproto.TLObject { // 0x15862c8e
		o := MakeTLImmutableUser(nil)
		o.Data2.Constructor = 361114766
		return o
	},
	1256160192: func() mtproto.TLObject { // 0x4adf7bc0
		o := MakeTLUserImportedContacts(nil)
		o.Data2.Constructor = 1256160192
		return o
	},
	1894399913: func() mtproto.TLObject { // 0x70ea3fa9
		o := MakeTLPeerPeerNotifySettings(nil)
		o.Data2.Constructor = 1894399913
		return o
	},
	-1810715178: func() mtproto.TLObject { // 0x9412add6
		o := MakeTLPrivacyKeyRules(nil)
		o.Data2.Constructor = -1810715178
		return o
	},
	-313287543: func() mtproto.TLObject { // 0xed539c89
		o := MakeTLLastSeenData(nil)
		o.Data2.Constructor = -313287543
		return o
	},
	722018346: func() mtproto.TLObject { // 0x2b09202a
		o := MakeTLContactData(nil)
		o.Data2.Constructor = 722018346
		return o
	},
	23110840: func() mtproto.TLObject { // 0x160a4b8
		o := MakeTLBotData(nil)
		o.Data2.Constructor = 23110840
		return o
	},

	// Method
	2090958337: func() mtproto.TLObject { // 0x7ca17e01
		return &TLUserGetLastSeens{
			Constructor: 2090958337,
		}
	},
	1314677789: func() mtproto.TLObject { // 0x4e5c641d
		return &TLUserUpdateLastSeen{
			Constructor: 1314677789,
		}
	},
	-1860581154: func() mtproto.TLObject { // 0x9119c8de
		return &TLUserGetLastSeen{
			Constructor: -1860581154,
		}
	},
	-47047585: func() mtproto.TLObject { // 0xfd321c5f
		return &TLUserGetImmutableUser{
			Constructor: -47047585,
		}
	},
	187684863: func() mtproto.TLObject { // 0xb2fd7ff
		return &TLUserGetMutableUsers{
			Constructor: 187684863,
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
	1464414785: func() mtproto.TLObject { // 0x57493241
		return &TLUserCreateNewPredefinedUser{
			Constructor: 1464414785,
		}
	},
	876047141: func() mtproto.TLObject { // 0x34376b25
		return &TLUserGetPredefinedUser{
			Constructor: 876047141,
		}
	},
	-1053474843: func() mtproto.TLObject { // 0xc1353fe5
		return &TLUserGetAllPredefinedUser{
			Constructor: -1053474843,
		}
	},
	976922006: func() mtproto.TLObject { // 0x3a3aa596
		return &TLUserUpdatePredefinedFirstAndLastName{
			Constructor: 976922006,
		}
	},
	-1158303159: func() mtproto.TLObject { // 0xbaf5b249
		return &TLUserUpdatePredefinedVerified{
			Constructor: -1158303159,
		}
	},
	1269284562: func() mtproto.TLObject { // 0x4ba7bed2
		return &TLUserUpdatePredefinedUsername{
			Constructor: 1269284562,
		}
	},
	1626771303: func() mtproto.TLObject { // 0x60f68f67
		return &TLUserUpdatePredefinedCode{
			Constructor: 1626771303,
		}
	},
	68106153: func() mtproto.TLObject { // 0x40f37a9
		return &TLUserPredefinedBindRegisteredUserId{
			Constructor: 68106153,
		}
	},
	2044729473: func() mtproto.TLObject { // 0x79e01881
		return &TLUserCreateNewUser{
			Constructor: 2044729473,
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
	-49225414: func() mtproto.TLObject { // 0xfd10e13a
		return &TLUserGetFullUser{
			Constructor: -49225414,
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
// UserData <--
//  + TL_UserData
//

func (m *UserData) Encode(layer int32) []byte {
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
	case Predicate_userData:
		t := m.To_UserData()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *UserData) CalcByteSize(layer int32) int {
	return 0
}

func (m *UserData) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x7f78f615:
		m2 := MakeTLUserData(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *UserData) DebugString() string {
	switch m.PredicateName {
	case Predicate_userData:
		t := m.To_UserData()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_UserData
// userData flags:# id:long access_hash:long user_type:int sceret_key_id:long first_name:string last_name:string username:string phone:string profile_photo:flags.0?Photo bot:flags.8?BotData country_code:string verified:flags.1?true support:flags.2?true scam:flags.3?true fake:flags.4?true about:flags.5?string restricted:flags.7?true restriction_reason:flags.7?Vector<RestrictionReason> contacts_version:int privacies_version:int deleted:flags.9?true = UserData;
func (m *UserData) To_UserData() *TLUserData {
	m.PredicateName = Predicate_userData
	return &TLUserData{
		Data2: m,
	}
}

// MakeTLUserData
// userData flags:# id:long access_hash:long user_type:int sceret_key_id:long first_name:string last_name:string username:string phone:string profile_photo:flags.0?Photo bot:flags.8?BotData country_code:string verified:flags.1?true support:flags.2?true scam:flags.3?true fake:flags.4?true about:flags.5?string restricted:flags.7?true restriction_reason:flags.7?Vector<RestrictionReason> contacts_version:int privacies_version:int deleted:flags.9?true = UserData;
func MakeTLUserData(data2 *UserData) *TLUserData {
	if data2 == nil {
		return &TLUserData{Data2: &UserData{
			PredicateName: Predicate_userData,
		}}
	} else {
		data2.PredicateName = Predicate_userData
		return &TLUserData{Data2: data2}
	}
}

func (m *TLUserData) To_UserData() *UserData {
	m.Data2.PredicateName = Predicate_userData
	return m.Data2
}

//// flags
func (m *TLUserData) SetId(v int64) { m.Data2.Id = v }
func (m *TLUserData) GetId() int64  { return m.Data2.Id }

func (m *TLUserData) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLUserData) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLUserData) SetUserType(v int32) { m.Data2.UserType = v }
func (m *TLUserData) GetUserType() int32  { return m.Data2.UserType }

func (m *TLUserData) SetSceretKeyId(v int64) { m.Data2.SceretKeyId = v }
func (m *TLUserData) GetSceretKeyId() int64  { return m.Data2.SceretKeyId }

func (m *TLUserData) SetFirstName(v string) { m.Data2.FirstName = v }
func (m *TLUserData) GetFirstName() string  { return m.Data2.FirstName }

func (m *TLUserData) SetLastName(v string) { m.Data2.LastName = v }
func (m *TLUserData) GetLastName() string  { return m.Data2.LastName }

func (m *TLUserData) SetUsername(v string) { m.Data2.Username = v }
func (m *TLUserData) GetUsername() string  { return m.Data2.Username }

func (m *TLUserData) SetPhone(v string) { m.Data2.Phone = v }
func (m *TLUserData) GetPhone() string  { return m.Data2.Phone }

func (m *TLUserData) SetProfilePhoto(v *mtproto.Photo) { m.Data2.ProfilePhoto = v }
func (m *TLUserData) GetProfilePhoto() *mtproto.Photo  { return m.Data2.ProfilePhoto }

func (m *TLUserData) SetBot(v *BotData) { m.Data2.Bot = v }
func (m *TLUserData) GetBot() *BotData  { return m.Data2.Bot }

func (m *TLUserData) SetCountryCode(v string) { m.Data2.CountryCode = v }
func (m *TLUserData) GetCountryCode() string  { return m.Data2.CountryCode }

func (m *TLUserData) SetVerified(v bool) { m.Data2.Verified = v }
func (m *TLUserData) GetVerified() bool  { return m.Data2.Verified }

func (m *TLUserData) SetSupport(v bool) { m.Data2.Support = v }
func (m *TLUserData) GetSupport() bool  { return m.Data2.Support }

func (m *TLUserData) SetScam(v bool) { m.Data2.Scam = v }
func (m *TLUserData) GetScam() bool  { return m.Data2.Scam }

func (m *TLUserData) SetFake(v bool) { m.Data2.Fake = v }
func (m *TLUserData) GetFake() bool  { return m.Data2.Fake }

func (m *TLUserData) SetAbout(v *types.StringValue) { m.Data2.About = v }
func (m *TLUserData) GetAbout() *types.StringValue  { return m.Data2.About }

func (m *TLUserData) SetRestricted(v bool) { m.Data2.Restricted = v }
func (m *TLUserData) GetRestricted() bool  { return m.Data2.Restricted }

func (m *TLUserData) SetRestrictionReason(v []*mtproto.RestrictionReason) {
	m.Data2.RestrictionReason = v
}
func (m *TLUserData) GetRestrictionReason() []*mtproto.RestrictionReason {
	return m.Data2.RestrictionReason
}

func (m *TLUserData) SetContactsVersion(v int32) { m.Data2.ContactsVersion = v }
func (m *TLUserData) GetContactsVersion() int32  { return m.Data2.ContactsVersion }

func (m *TLUserData) SetPrivaciesVersion(v int32) { m.Data2.PrivaciesVersion = v }
func (m *TLUserData) GetPrivaciesVersion() int32  { return m.Data2.PrivaciesVersion }

func (m *TLUserData) SetDeleted(v bool) { m.Data2.Deleted = v }
func (m *TLUserData) GetDeleted() bool  { return m.Data2.Deleted }

func (m *TLUserData) GetPredicateName() string {
	return Predicate_userData
}

func (m *TLUserData) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x7f78f615: func() []byte {
			// userData flags:# id:long access_hash:long user_type:int sceret_key_id:long first_name:string last_name:string username:string phone:string profile_photo:flags.0?Photo bot:flags.8?BotData country_code:string verified:flags.1?true support:flags.2?true scam:flags.3?true fake:flags.4?true about:flags.5?string restricted:flags.7?true restriction_reason:flags.7?Vector<RestrictionReason> contacts_version:int privacies_version:int deleted:flags.9?true = UserData;
			x.UInt(0x7f78f615)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.GetProfilePhoto() != nil {
					flags |= 1 << 0
				}
				if m.GetBot() != nil {
					flags |= 1 << 8
				}

				if m.GetVerified() == true {
					flags |= 1 << 1
				}
				if m.GetSupport() == true {
					flags |= 1 << 2
				}
				if m.GetScam() == true {
					flags |= 1 << 3
				}
				if m.GetFake() == true {
					flags |= 1 << 4
				}
				if m.GetAbout() != nil {
					flags |= 1 << 5
				}
				if m.GetRestricted() == true {
					flags |= 1 << 7
				}
				if m.GetRestrictionReason() != nil {
					flags |= 1 << 7
				}

				if m.GetDeleted() == true {
					flags |= 1 << 9
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.UInt(flags)
			x.Long(m.GetId())
			x.Long(m.GetAccessHash())
			x.Int(m.GetUserType())
			x.Long(m.GetSceretKeyId())
			x.String(m.GetFirstName())
			x.String(m.GetLastName())
			x.String(m.GetUsername())
			x.String(m.GetPhone())
			if m.GetProfilePhoto() != nil {
				x.Bytes(m.GetProfilePhoto().Encode(layer))
			}

			if m.GetBot() != nil {
				x.Bytes(m.GetBot().Encode(layer))
			}

			x.String(m.GetCountryCode())
			if m.GetAbout() != nil {
				x.String(m.GetAbout().Value)
			}

			if m.GetRestrictionReason() != nil {
				x.Int(int32(mtproto.CRC32_vector))
				x.Int(int32(len(m.GetRestrictionReason())))
				for _, v := range m.GetRestrictionReason() {
					x.Bytes((*v).Encode(layer))
				}
			}
			x.Int(m.GetContactsVersion())
			x.Int(m.GetPrivaciesVersion())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_userData, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_userData, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLUserData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserData) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x7f78f615: func() error {
			// userData flags:# id:long access_hash:long user_type:int sceret_key_id:long first_name:string last_name:string username:string phone:string profile_photo:flags.0?Photo bot:flags.8?BotData country_code:string verified:flags.1?true support:flags.2?true scam:flags.3?true fake:flags.4?true about:flags.5?string restricted:flags.7?true restriction_reason:flags.7?Vector<RestrictionReason> contacts_version:int privacies_version:int deleted:flags.9?true = UserData;
			var flags = dBuf.UInt()
			_ = flags
			m.SetId(dBuf.Long())
			m.SetAccessHash(dBuf.Long())
			m.SetUserType(dBuf.Int())
			m.SetSceretKeyId(dBuf.Long())
			m.SetFirstName(dBuf.String())
			m.SetLastName(dBuf.String())
			m.SetUsername(dBuf.String())
			m.SetPhone(dBuf.String())
			if (flags & (1 << 0)) != 0 {
				m9 := &mtproto.Photo{}
				m9.Decode(dBuf)
				m.SetProfilePhoto(m9)
			}
			if (flags & (1 << 8)) != 0 {
				m10 := &BotData{}
				m10.Decode(dBuf)
				m.SetBot(m10)
			}
			m.SetCountryCode(dBuf.String())
			if (flags & (1 << 1)) != 0 {
				m.SetVerified(true)
			}
			if (flags & (1 << 2)) != 0 {
				m.SetSupport(true)
			}
			if (flags & (1 << 3)) != 0 {
				m.SetScam(true)
			}
			if (flags & (1 << 4)) != 0 {
				m.SetFake(true)
			}
			if (flags & (1 << 5)) != 0 {
				m.SetAbout(&types.StringValue{Value: dBuf.String()})
			}

			if (flags & (1 << 7)) != 0 {
				m.SetRestricted(true)
			}
			if (flags & (1 << 7)) != 0 {
				c18 := dBuf.Int()
				if c18 != int32(mtproto.CRC32_vector) {
					// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 18, c18)
					return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 18, c18)
				}
				l18 := dBuf.Int()
				v18 := make([]*mtproto.RestrictionReason, l18)
				for i := int32(0); i < l18; i++ {
					v18[i] = &mtproto.RestrictionReason{}
					v18[i].Decode(dBuf)
				}
				m.SetRestrictionReason(v18)
			}
			m.SetContactsVersion(dBuf.Int())
			m.SetPrivaciesVersion(dBuf.Int())
			if (flags & (1 << 9)) != 0 {
				m.SetDeleted(true)
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

func (m *TLUserData) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// ImmutableUser <--
//  + TL_ImmutableUser
//

func (m *ImmutableUser) Encode(layer int32) []byte {
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
	case Predicate_immutableUser:
		t := m.To_ImmutableUser()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *ImmutableUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *ImmutableUser) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x15862c8e:
		m2 := MakeTLImmutableUser(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *ImmutableUser) DebugString() string {
	switch m.PredicateName {
	case Predicate_immutableUser:
		t := m.To_ImmutableUser()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_ImmutableUser
// immutableUser flags:# user:UserData last_seen_at:long contacts:flags.0?Vector<ContactData> keys_privacy_rules:Vector<PrivacyKeyRules> = ImmutableUser;
func (m *ImmutableUser) To_ImmutableUser() *TLImmutableUser {
	m.PredicateName = Predicate_immutableUser
	return &TLImmutableUser{
		Data2: m,
	}
}

// MakeTLImmutableUser
// immutableUser flags:# user:UserData last_seen_at:long contacts:flags.0?Vector<ContactData> keys_privacy_rules:Vector<PrivacyKeyRules> = ImmutableUser;
func MakeTLImmutableUser(data2 *ImmutableUser) *TLImmutableUser {
	if data2 == nil {
		return &TLImmutableUser{Data2: &ImmutableUser{
			PredicateName: Predicate_immutableUser,
		}}
	} else {
		data2.PredicateName = Predicate_immutableUser
		return &TLImmutableUser{Data2: data2}
	}
}

func (m *TLImmutableUser) To_ImmutableUser() *ImmutableUser {
	m.Data2.PredicateName = Predicate_immutableUser
	return m.Data2
}

//// flags
func (m *TLImmutableUser) SetUser(v *UserData) { m.Data2.User = v }
func (m *TLImmutableUser) GetUser() *UserData  { return m.Data2.User }

func (m *TLImmutableUser) SetLastSeenAt(v int64) { m.Data2.LastSeenAt = v }
func (m *TLImmutableUser) GetLastSeenAt() int64  { return m.Data2.LastSeenAt }

func (m *TLImmutableUser) SetContacts(v []*ContactData) { m.Data2.Contacts = v }
func (m *TLImmutableUser) GetContacts() []*ContactData  { return m.Data2.Contacts }

func (m *TLImmutableUser) SetKeysPrivacyRules(v []*PrivacyKeyRules) { m.Data2.KeysPrivacyRules = v }
func (m *TLImmutableUser) GetKeysPrivacyRules() []*PrivacyKeyRules  { return m.Data2.KeysPrivacyRules }

func (m *TLImmutableUser) GetPredicateName() string {
	return Predicate_immutableUser
}

func (m *TLImmutableUser) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x15862c8e: func() []byte {
			// immutableUser flags:# user:UserData last_seen_at:long contacts:flags.0?Vector<ContactData> keys_privacy_rules:Vector<PrivacyKeyRules> = ImmutableUser;
			x.UInt(0x15862c8e)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.GetContacts() != nil {
					flags |= 1 << 0
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.UInt(flags)
			x.Bytes(m.GetUser().Encode(layer))
			x.Long(m.GetLastSeenAt())
			if m.GetContacts() != nil {
				x.Int(int32(mtproto.CRC32_vector))
				x.Int(int32(len(m.GetContacts())))
				for _, v := range m.GetContacts() {
					x.Bytes((*v).Encode(layer))
				}
			}

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetKeysPrivacyRules())))
			for _, v := range m.GetKeysPrivacyRules() {
				x.Bytes((*v).Encode(layer))
			}

			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_immutableUser, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_immutableUser, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLImmutableUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLImmutableUser) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x15862c8e: func() error {
			// immutableUser flags:# user:UserData last_seen_at:long contacts:flags.0?Vector<ContactData> keys_privacy_rules:Vector<PrivacyKeyRules> = ImmutableUser;
			var flags = dBuf.UInt()
			_ = flags

			m1 := &UserData{}
			m1.Decode(dBuf)
			m.SetUser(m1)

			m.SetLastSeenAt(dBuf.Long())
			if (flags & (1 << 0)) != 0 {
				c3 := dBuf.Int()
				if c3 != int32(mtproto.CRC32_vector) {
					// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
					return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
				}
				l3 := dBuf.Int()
				v3 := make([]*ContactData, l3)
				for i := int32(0); i < l3; i++ {
					v3[i] = &ContactData{}
					v3[i].Decode(dBuf)
				}
				m.SetContacts(v3)
			}
			c4 := dBuf.Int()
			if c4 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 4, c4)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 4, c4)
			}
			l4 := dBuf.Int()
			v4 := make([]*PrivacyKeyRules, l4)
			for i := int32(0); i < l4; i++ {
				v4[i] = &PrivacyKeyRules{}
				v4[i].Decode(dBuf)
			}
			m.SetKeysPrivacyRules(v4)

			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLImmutableUser) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// UserImportedContacts <--
//  + TL_UserImportedContacts
//

func (m *UserImportedContacts) Encode(layer int32) []byte {
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
	case Predicate_userImportedContacts:
		t := m.To_UserImportedContacts()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
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

func (m *UserImportedContacts) DebugString() string {
	switch m.PredicateName {
	case Predicate_userImportedContacts:
		t := m.To_UserImportedContacts()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_UserImportedContacts
// userImportedContacts imported:Vector<ImportedContact> popular_invites:Vector<PopularContact> retry_contacts:Vector<long> users:Vector<User> update_id_list:Vector<long> = UserImportedContacts;
func (m *UserImportedContacts) To_UserImportedContacts() *TLUserImportedContacts {
	m.PredicateName = Predicate_userImportedContacts
	return &TLUserImportedContacts{
		Data2: m,
	}
}

// MakeTLUserImportedContacts
// userImportedContacts imported:Vector<ImportedContact> popular_invites:Vector<PopularContact> retry_contacts:Vector<long> users:Vector<User> update_id_list:Vector<long> = UserImportedContacts;
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

func (m *TLUserImportedContacts) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x4adf7bc0: func() []byte {
			// userImportedContacts imported:Vector<ImportedContact> popular_invites:Vector<PopularContact> retry_contacts:Vector<long> users:Vector<User> update_id_list:Vector<long> = UserImportedContacts;
			x.UInt(0x4adf7bc0)

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetImported())))
			for _, v := range m.GetImported() {
				x.Bytes((*v).Encode(layer))
			}

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetPopularInvites())))
			for _, v := range m.GetPopularInvites() {
				x.Bytes((*v).Encode(layer))
			}

			x.VectorLong(m.GetRetryContacts())

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetUsers())))
			for _, v := range m.GetUsers() {
				x.Bytes((*v).Encode(layer))
			}

			x.VectorLong(m.GetUpdateIdList())

			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_userImportedContacts, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_userImportedContacts, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLUserImportedContacts) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserImportedContacts) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x4adf7bc0: func() error {
			// userImportedContacts imported:Vector<ImportedContact> popular_invites:Vector<PopularContact> retry_contacts:Vector<long> users:Vector<User> update_id_list:Vector<long> = UserImportedContacts;
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

func (m *TLUserImportedContacts) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// PeerPeerNotifySettings <--
//  + TL_PeerPeerNotifySettings
//

func (m *PeerPeerNotifySettings) Encode(layer int32) []byte {
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
	case Predicate_peerPeerNotifySettings:
		t := m.To_PeerPeerNotifySettings()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
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

func (m *PeerPeerNotifySettings) DebugString() string {
	switch m.PredicateName {
	case Predicate_peerPeerNotifySettings:
		t := m.To_PeerPeerNotifySettings()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_PeerPeerNotifySettings
// peerPeerNotifySettings peer_type:int peer_id:long settings:PeerNotifySettings = PeerPeerNotifySettings;
func (m *PeerPeerNotifySettings) To_PeerPeerNotifySettings() *TLPeerPeerNotifySettings {
	m.PredicateName = Predicate_peerPeerNotifySettings
	return &TLPeerPeerNotifySettings{
		Data2: m,
	}
}

// MakeTLPeerPeerNotifySettings
// peerPeerNotifySettings peer_type:int peer_id:long settings:PeerNotifySettings = PeerPeerNotifySettings;
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

func (m *TLPeerPeerNotifySettings) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x70ea3fa9: func() []byte {
			// peerPeerNotifySettings peer_type:int peer_id:long settings:PeerNotifySettings = PeerPeerNotifySettings;
			x.UInt(0x70ea3fa9)

			x.Int(m.GetPeerType())
			x.Long(m.GetPeerId())
			x.Bytes(m.GetSettings().Encode(layer))
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_peerPeerNotifySettings, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_peerPeerNotifySettings, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLPeerPeerNotifySettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLPeerPeerNotifySettings) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x70ea3fa9: func() error {
			// peerPeerNotifySettings peer_type:int peer_id:long settings:PeerNotifySettings = PeerPeerNotifySettings;
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

func (m *TLPeerPeerNotifySettings) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// PrivacyKeyRules <--
//  + TL_PrivacyKeyRules
//

func (m *PrivacyKeyRules) Encode(layer int32) []byte {
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
	case Predicate_privacyKeyRules:
		t := m.To_PrivacyKeyRules()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *PrivacyKeyRules) CalcByteSize(layer int32) int {
	return 0
}

func (m *PrivacyKeyRules) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x9412add6:
		m2 := MakeTLPrivacyKeyRules(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *PrivacyKeyRules) DebugString() string {
	switch m.PredicateName {
	case Predicate_privacyKeyRules:
		t := m.To_PrivacyKeyRules()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_PrivacyKeyRules
// privacyKeyRules key:int rules:Vector<PrivacyRule> = PrivacyKeyRules;
func (m *PrivacyKeyRules) To_PrivacyKeyRules() *TLPrivacyKeyRules {
	m.PredicateName = Predicate_privacyKeyRules
	return &TLPrivacyKeyRules{
		Data2: m,
	}
}

// MakeTLPrivacyKeyRules
// privacyKeyRules key:int rules:Vector<PrivacyRule> = PrivacyKeyRules;
func MakeTLPrivacyKeyRules(data2 *PrivacyKeyRules) *TLPrivacyKeyRules {
	if data2 == nil {
		return &TLPrivacyKeyRules{Data2: &PrivacyKeyRules{
			PredicateName: Predicate_privacyKeyRules,
		}}
	} else {
		data2.PredicateName = Predicate_privacyKeyRules
		return &TLPrivacyKeyRules{Data2: data2}
	}
}

func (m *TLPrivacyKeyRules) To_PrivacyKeyRules() *PrivacyKeyRules {
	m.Data2.PredicateName = Predicate_privacyKeyRules
	return m.Data2
}

func (m *TLPrivacyKeyRules) SetKey(v int32) { m.Data2.Key = v }
func (m *TLPrivacyKeyRules) GetKey() int32  { return m.Data2.Key }

func (m *TLPrivacyKeyRules) SetRules(v []*mtproto.PrivacyRule) { m.Data2.Rules = v }
func (m *TLPrivacyKeyRules) GetRules() []*mtproto.PrivacyRule  { return m.Data2.Rules }

func (m *TLPrivacyKeyRules) GetPredicateName() string {
	return Predicate_privacyKeyRules
}

func (m *TLPrivacyKeyRules) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x9412add6: func() []byte {
			// privacyKeyRules key:int rules:Vector<PrivacyRule> = PrivacyKeyRules;
			x.UInt(0x9412add6)

			x.Int(m.GetKey())

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetRules())))
			for _, v := range m.GetRules() {
				x.Bytes((*v).Encode(layer))
			}

			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_privacyKeyRules, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_privacyKeyRules, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLPrivacyKeyRules) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLPrivacyKeyRules) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x9412add6: func() error {
			// privacyKeyRules key:int rules:Vector<PrivacyRule> = PrivacyKeyRules;
			m.SetKey(dBuf.Int())
			c1 := dBuf.Int()
			if c1 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
			}
			l1 := dBuf.Int()
			v1 := make([]*mtproto.PrivacyRule, l1)
			for i := int32(0); i < l1; i++ {
				v1[i] = &mtproto.PrivacyRule{}
				v1[i].Decode(dBuf)
			}
			m.SetRules(v1)

			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLPrivacyKeyRules) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// LastSeenData <--
//  + TL_LastSeenData
//

func (m *LastSeenData) Encode(layer int32) []byte {
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
	case Predicate_lastSeenData:
		t := m.To_LastSeenData()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *LastSeenData) CalcByteSize(layer int32) int {
	return 0
}

func (m *LastSeenData) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xed539c89:
		m2 := MakeTLLastSeenData(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *LastSeenData) DebugString() string {
	switch m.PredicateName {
	case Predicate_lastSeenData:
		t := m.To_LastSeenData()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_LastSeenData
// lastSeenData user_id:long last_seen_at:long expries:int = LastSeenData;
func (m *LastSeenData) To_LastSeenData() *TLLastSeenData {
	m.PredicateName = Predicate_lastSeenData
	return &TLLastSeenData{
		Data2: m,
	}
}

// MakeTLLastSeenData
// lastSeenData user_id:long last_seen_at:long expries:int = LastSeenData;
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

func (m *TLLastSeenData) SetExpries(v int32) { m.Data2.Expries = v }
func (m *TLLastSeenData) GetExpries() int32  { return m.Data2.Expries }

func (m *TLLastSeenData) GetPredicateName() string {
	return Predicate_lastSeenData
}

func (m *TLLastSeenData) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xed539c89: func() []byte {
			// lastSeenData user_id:long last_seen_at:long expries:int = LastSeenData;
			x.UInt(0xed539c89)

			x.Long(m.GetUserId())
			x.Long(m.GetLastSeenAt())
			x.Int(m.GetExpries())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_lastSeenData, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_lastSeenData, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLLastSeenData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLLastSeenData) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xed539c89: func() error {
			// lastSeenData user_id:long last_seen_at:long expries:int = LastSeenData;
			m.SetUserId(dBuf.Long())
			m.SetLastSeenAt(dBuf.Long())
			m.SetExpries(dBuf.Int())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLLastSeenData) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// ContactData <--
//  + TL_ContactData
//

func (m *ContactData) Encode(layer int32) []byte {
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
	case Predicate_contactData:
		t := m.To_ContactData()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *ContactData) CalcByteSize(layer int32) int {
	return 0
}

func (m *ContactData) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x2b09202a:
		m2 := MakeTLContactData(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *ContactData) DebugString() string {
	switch m.PredicateName {
	case Predicate_contactData:
		t := m.To_ContactData()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_ContactData
// contactData flags:# user_id:long contact_user_id:long first_name:flags.0?string last_name:flags.1?string mutual_contact:flags.2?true = ContactData;
func (m *ContactData) To_ContactData() *TLContactData {
	m.PredicateName = Predicate_contactData
	return &TLContactData{
		Data2: m,
	}
}

// MakeTLContactData
// contactData flags:# user_id:long contact_user_id:long first_name:flags.0?string last_name:flags.1?string mutual_contact:flags.2?true = ContactData;
func MakeTLContactData(data2 *ContactData) *TLContactData {
	if data2 == nil {
		return &TLContactData{Data2: &ContactData{
			PredicateName: Predicate_contactData,
		}}
	} else {
		data2.PredicateName = Predicate_contactData
		return &TLContactData{Data2: data2}
	}
}

func (m *TLContactData) To_ContactData() *ContactData {
	m.Data2.PredicateName = Predicate_contactData
	return m.Data2
}

//// flags
func (m *TLContactData) SetUserId(v int64) { m.Data2.UserId = v }
func (m *TLContactData) GetUserId() int64  { return m.Data2.UserId }

func (m *TLContactData) SetContactUserId(v int64) { m.Data2.ContactUserId = v }
func (m *TLContactData) GetContactUserId() int64  { return m.Data2.ContactUserId }

func (m *TLContactData) SetFirstName(v *types.StringValue) { m.Data2.FirstName = v }
func (m *TLContactData) GetFirstName() *types.StringValue  { return m.Data2.FirstName }

func (m *TLContactData) SetLastName(v *types.StringValue) { m.Data2.LastName = v }
func (m *TLContactData) GetLastName() *types.StringValue  { return m.Data2.LastName }

func (m *TLContactData) SetMutualContact(v bool) { m.Data2.MutualContact = v }
func (m *TLContactData) GetMutualContact() bool  { return m.Data2.MutualContact }

func (m *TLContactData) GetPredicateName() string {
	return Predicate_contactData
}

func (m *TLContactData) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x2b09202a: func() []byte {
			// contactData flags:# user_id:long contact_user_id:long first_name:flags.0?string last_name:flags.1?string mutual_contact:flags.2?true = ContactData;
			x.UInt(0x2b09202a)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.GetFirstName() != nil {
					flags |= 1 << 0
				}
				if m.GetLastName() != nil {
					flags |= 1 << 1
				}
				if m.GetMutualContact() == true {
					flags |= 1 << 2
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.UInt(flags)
			x.Long(m.GetUserId())
			x.Long(m.GetContactUserId())
			if m.GetFirstName() != nil {
				x.String(m.GetFirstName().Value)
			}

			if m.GetLastName() != nil {
				x.String(m.GetLastName().Value)
			}

			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_contactData, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_contactData, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLContactData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLContactData) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x2b09202a: func() error {
			// contactData flags:# user_id:long contact_user_id:long first_name:flags.0?string last_name:flags.1?string mutual_contact:flags.2?true = ContactData;
			var flags = dBuf.UInt()
			_ = flags
			m.SetUserId(dBuf.Long())
			m.SetContactUserId(dBuf.Long())
			if (flags & (1 << 0)) != 0 {
				m.SetFirstName(&types.StringValue{Value: dBuf.String()})
			}

			if (flags & (1 << 1)) != 0 {
				m.SetLastName(&types.StringValue{Value: dBuf.String()})
			}

			if (flags & (1 << 2)) != 0 {
				m.SetMutualContact(true)
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

func (m *TLContactData) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// BotData <--
//  + TL_BotData
//

func (m *BotData) Encode(layer int32) []byte {
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
	case Predicate_botData:
		t := m.To_BotData()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *BotData) CalcByteSize(layer int32) int {
	return 0
}

func (m *BotData) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x160a4b8:
		m2 := MakeTLBotData(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *BotData) DebugString() string {
	switch m.PredicateName {
	case Predicate_botData:
		t := m.To_BotData()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_BotData
// botData flags:# id:long bot_type:int creator:long token:string description:string bot_chat_history:flags.1?true bot_nochats:flags.2?true bot_inline_geo:flags.3?true bot_info_version:int bot_inline_placeholder:flags.4?string= BotData;
func (m *BotData) To_BotData() *TLBotData {
	m.PredicateName = Predicate_botData
	return &TLBotData{
		Data2: m,
	}
}

// MakeTLBotData
// botData flags:# id:long bot_type:int creator:long token:string description:string bot_chat_history:flags.1?true bot_nochats:flags.2?true bot_inline_geo:flags.3?true bot_info_version:int bot_inline_placeholder:flags.4?string= BotData;
func MakeTLBotData(data2 *BotData) *TLBotData {
	if data2 == nil {
		return &TLBotData{Data2: &BotData{
			PredicateName: Predicate_botData,
		}}
	} else {
		data2.PredicateName = Predicate_botData
		return &TLBotData{Data2: data2}
	}
}

func (m *TLBotData) To_BotData() *BotData {
	m.Data2.PredicateName = Predicate_botData
	return m.Data2
}

//// flags
func (m *TLBotData) SetId(v int64) { m.Data2.Id = v }
func (m *TLBotData) GetId() int64  { return m.Data2.Id }

func (m *TLBotData) SetBotType(v int32) { m.Data2.BotType = v }
func (m *TLBotData) GetBotType() int32  { return m.Data2.BotType }

func (m *TLBotData) SetCreator(v int64) { m.Data2.Creator = v }
func (m *TLBotData) GetCreator() int64  { return m.Data2.Creator }

func (m *TLBotData) SetToken(v string) { m.Data2.Token = v }
func (m *TLBotData) GetToken() string  { return m.Data2.Token }

func (m *TLBotData) SetDescription(v string) { m.Data2.Description = v }
func (m *TLBotData) GetDescription() string  { return m.Data2.Description }

func (m *TLBotData) SetBotChatHistory(v bool) { m.Data2.BotChatHistory = v }
func (m *TLBotData) GetBotChatHistory() bool  { return m.Data2.BotChatHistory }

func (m *TLBotData) SetBotNochats(v bool) { m.Data2.BotNochats = v }
func (m *TLBotData) GetBotNochats() bool  { return m.Data2.BotNochats }

func (m *TLBotData) SetBotInlineGeo(v bool) { m.Data2.BotInlineGeo = v }
func (m *TLBotData) GetBotInlineGeo() bool  { return m.Data2.BotInlineGeo }

func (m *TLBotData) SetBotInfoVersion(v int32) { m.Data2.BotInfoVersion = v }
func (m *TLBotData) GetBotInfoVersion() int32  { return m.Data2.BotInfoVersion }

func (m *TLBotData) SetBotInlinePlaceholder(v *types.StringValue) { m.Data2.BotInlinePlaceholder = v }
func (m *TLBotData) GetBotInlinePlaceholder() *types.StringValue  { return m.Data2.BotInlinePlaceholder }

func (m *TLBotData) GetPredicateName() string {
	return Predicate_botData
}

func (m *TLBotData) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x160a4b8: func() []byte {
			// botData flags:# id:long bot_type:int creator:long token:string description:string bot_chat_history:flags.1?true bot_nochats:flags.2?true bot_inline_geo:flags.3?true bot_info_version:int bot_inline_placeholder:flags.4?string= BotData;
			x.UInt(0x160a4b8)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.GetBotChatHistory() == true {
					flags |= 1 << 1
				}
				if m.GetBotNochats() == true {
					flags |= 1 << 2
				}
				if m.GetBotInlineGeo() == true {
					flags |= 1 << 3
				}

				if m.GetBotInlinePlaceholder() != nil {
					flags |= 1 << 4
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.UInt(flags)
			x.Long(m.GetId())
			x.Int(m.GetBotType())
			x.Long(m.GetCreator())
			x.String(m.GetToken())
			x.String(m.GetDescription())
			x.Int(m.GetBotInfoVersion())
			if m.GetBotInlinePlaceholder() != nil {
				x.String(m.GetBotInlinePlaceholder().Value)
			}

			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_botData, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_botData, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLBotData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLBotData) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x160a4b8: func() error {
			// botData flags:# id:long bot_type:int creator:long token:string description:string bot_chat_history:flags.1?true bot_nochats:flags.2?true bot_inline_geo:flags.3?true bot_info_version:int bot_inline_placeholder:flags.4?string= BotData;
			var flags = dBuf.UInt()
			_ = flags
			m.SetId(dBuf.Long())
			m.SetBotType(dBuf.Int())
			m.SetCreator(dBuf.Long())
			m.SetToken(dBuf.String())
			m.SetDescription(dBuf.String())
			if (flags & (1 << 1)) != 0 {
				m.SetBotChatHistory(true)
			}
			if (flags & (1 << 2)) != 0 {
				m.SetBotNochats(true)
			}
			if (flags & (1 << 3)) != 0 {
				m.SetBotInlineGeo(true)
			}
			m.SetBotInfoVersion(dBuf.Int())
			if (flags & (1 << 4)) != 0 {
				m.SetBotInlinePlaceholder(&types.StringValue{Value: dBuf.String()})
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

func (m *TLBotData) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// TLUserGetLastSeens
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetLastSeens) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getLastSeens))

	switch uint32(m.Constructor) {
	case 0x7ca17e01:
		// user.getLastSeens id:Vector<long> = Vector<LastSeenData>;
		x.UInt(0x7ca17e01)

		// no flags

		x.VectorLong(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetLastSeens) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetLastSeens) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x7ca17e01:
		// user.getLastSeens id:Vector<long> = Vector<LastSeenData>;

		// not has flags

		m.Id = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetLastSeens) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserUpdateLastSeen
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUpdateLastSeen) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_updateLastSeen))

	switch uint32(m.Constructor) {
	case 0x4e5c641d:
		// user.updateLastSeen id:long last_seen_at:long expries:int = Bool;
		x.UInt(0x4e5c641d)

		// no flags

		x.Long(m.GetId())
		x.Long(m.GetLastSeenAt())
		x.Int(m.GetExpries())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserUpdateLastSeen) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUpdateLastSeen) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x4e5c641d:
		// user.updateLastSeen id:long last_seen_at:long expries:int = Bool;

		// not has flags

		m.Id = dBuf.Long()
		m.LastSeenAt = dBuf.Long()
		m.Expries = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserUpdateLastSeen) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetLastSeen
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetLastSeen) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getLastSeen))

	switch uint32(m.Constructor) {
	case 0x9119c8de:
		// user.getLastSeen id:long = LastSeenData;
		x.UInt(0x9119c8de)

		// no flags

		x.Long(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetLastSeen) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetLastSeen) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x9119c8de:
		// user.getLastSeen id:long = LastSeenData;

		// not has flags

		m.Id = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetLastSeen) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetImmutableUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetImmutableUser) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getImmutableUser))

	switch uint32(m.Constructor) {
	case 0xfd321c5f:
		// user.getImmutableUser id:long = ImmutableUser;
		x.UInt(0xfd321c5f)

		// no flags

		x.Long(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetImmutableUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetImmutableUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xfd321c5f:
		// user.getImmutableUser id:long = ImmutableUser;

		// not has flags

		m.Id = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetImmutableUser) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetMutableUsers
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetMutableUsers) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getMutableUsers))

	switch uint32(m.Constructor) {
	case 0xb2fd7ff:
		// user.getMutableUsers id:Vector<long> = Vector<ImmutableUser>;
		x.UInt(0xb2fd7ff)

		// no flags

		x.VectorLong(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetMutableUsers) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetMutableUsers) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb2fd7ff:
		// user.getMutableUsers id:Vector<long> = Vector<ImmutableUser>;

		// not has flags

		m.Id = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetMutableUsers) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetImmutableUserByPhone
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetImmutableUserByPhone) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getImmutableUserByPhone))

	switch uint32(m.Constructor) {
	case 0xe9c36fe4:
		// user.getImmutableUserByPhone phone:string = ImmutableUser;
		x.UInt(0xe9c36fe4)

		// no flags

		x.String(m.GetPhone())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetImmutableUserByPhone) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetImmutableUserByPhone) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe9c36fe4:
		// user.getImmutableUserByPhone phone:string = ImmutableUser;

		// not has flags

		m.Phone = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetImmutableUserByPhone) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetImmutableUserByToken
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetImmutableUserByToken) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getImmutableUserByToken))

	switch uint32(m.Constructor) {
	case 0xff3e1373:
		// user.getImmutableUserByToken token:string = ImmutableUser;
		x.UInt(0xff3e1373)

		// no flags

		x.String(m.GetToken())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetImmutableUserByToken) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetImmutableUserByToken) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xff3e1373:
		// user.getImmutableUserByToken token:string = ImmutableUser;

		// not has flags

		m.Token = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetImmutableUserByToken) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserSetAccountDaysTTL
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserSetAccountDaysTTL) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_setAccountDaysTTL))

	switch uint32(m.Constructor) {
	case 0xd2550b4c:
		// user.setAccountDaysTTL user_id:long ttl:int = Bool;
		x.UInt(0xd2550b4c)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetTtl())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserSetAccountDaysTTL) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSetAccountDaysTTL) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xd2550b4c:
		// user.setAccountDaysTTL user_id:long ttl:int = Bool;

		// not has flags

		m.UserId = dBuf.Long()
		m.Ttl = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserSetAccountDaysTTL) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetAccountDaysTTL
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetAccountDaysTTL) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getAccountDaysTTL))

	switch uint32(m.Constructor) {
	case 0xb2843ee0:
		// user.getAccountDaysTTL user_id:long = AccountDaysTTL;
		x.UInt(0xb2843ee0)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetAccountDaysTTL) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetAccountDaysTTL) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb2843ee0:
		// user.getAccountDaysTTL user_id:long = AccountDaysTTL;

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetAccountDaysTTL) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetNotifySettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetNotifySettings) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getNotifySettings))

	switch uint32(m.Constructor) {
	case 0x40ac3766:
		// user.getNotifySettings user_id:long peer_type:int peer_id:long = PeerNotifySettings;
		x.UInt(0x40ac3766)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetNotifySettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetNotifySettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x40ac3766:
		// user.getNotifySettings user_id:long peer_type:int peer_id:long = PeerNotifySettings;

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

func (m *TLUserGetNotifySettings) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserSetNotifySettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserSetNotifySettings) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_setNotifySettings))

	switch uint32(m.Constructor) {
	case 0xc9ed65e5:
		// user.setNotifySettings user_id:long peer_type:int peer_id:long settings:PeerNotifySettings = Bool;
		x.UInt(0xc9ed65e5)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Bytes(m.GetSettings().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserSetNotifySettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSetNotifySettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc9ed65e5:
		// user.setNotifySettings user_id:long peer_type:int peer_id:long settings:PeerNotifySettings = Bool;

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

func (m *TLUserSetNotifySettings) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserResetNotifySettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserResetNotifySettings) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_resetNotifySettings))

	switch uint32(m.Constructor) {
	case 0xe079d74:
		// user.resetNotifySettings user_id:long = Bool;
		x.UInt(0xe079d74)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserResetNotifySettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserResetNotifySettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe079d74:
		// user.resetNotifySettings user_id:long = Bool;

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserResetNotifySettings) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetAllNotifySettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetAllNotifySettings) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getAllNotifySettings))

	switch uint32(m.Constructor) {
	case 0x55926875:
		// user.getAllNotifySettings user_id:long = Vector<PeerPeerNotifySettings>;
		x.UInt(0x55926875)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetAllNotifySettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetAllNotifySettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x55926875:
		// user.getAllNotifySettings user_id:long = Vector<PeerPeerNotifySettings>;

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetAllNotifySettings) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetGlobalPrivacySettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetGlobalPrivacySettings) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getGlobalPrivacySettings))

	switch uint32(m.Constructor) {
	case 0x77f6f112:
		// user.getGlobalPrivacySettings user_id:long = GlobalPrivacySettings;
		x.UInt(0x77f6f112)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetGlobalPrivacySettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetGlobalPrivacySettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x77f6f112:
		// user.getGlobalPrivacySettings user_id:long = GlobalPrivacySettings;

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetGlobalPrivacySettings) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserSetGlobalPrivacySettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserSetGlobalPrivacySettings) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_setGlobalPrivacySettings))

	switch uint32(m.Constructor) {
	case 0x8cb592ae:
		// user.setGlobalPrivacySettings user_id:long settings:GlobalPrivacySettings = Bool;
		x.UInt(0x8cb592ae)

		// no flags

		x.Long(m.GetUserId())
		x.Bytes(m.GetSettings().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserSetGlobalPrivacySettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSetGlobalPrivacySettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x8cb592ae:
		// user.setGlobalPrivacySettings user_id:long settings:GlobalPrivacySettings = Bool;

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

func (m *TLUserSetGlobalPrivacySettings) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetPrivacy
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetPrivacy) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getPrivacy))

	switch uint32(m.Constructor) {
	case 0x9d40a3b4:
		// user.getPrivacy user_id:long key_type:int = Vector<PrivacyRule>;
		x.UInt(0x9d40a3b4)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetKeyType())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetPrivacy) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetPrivacy) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x9d40a3b4:
		// user.getPrivacy user_id:long key_type:int = Vector<PrivacyRule>;

		// not has flags

		m.UserId = dBuf.Long()
		m.KeyType = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetPrivacy) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserSetPrivacy
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserSetPrivacy) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_setPrivacy))

	switch uint32(m.Constructor) {
	case 0x8855ad8f:
		// user.setPrivacy user_id:long key_type:int rules:Vector<PrivacyRule> = Bool;
		x.UInt(0x8855ad8f)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetKeyType())

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetRules())))
		for _, v := range m.GetRules() {
			x.Bytes((*v).Encode(layer))
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserSetPrivacy) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSetPrivacy) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x8855ad8f:
		// user.setPrivacy user_id:long key_type:int rules:Vector<PrivacyRule> = Bool;

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

func (m *TLUserSetPrivacy) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserCheckPrivacy
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserCheckPrivacy) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_checkPrivacy))

	switch uint32(m.Constructor) {
	case 0xc56e1eaa:
		// user.checkPrivacy flags:# user_id:long key_type:int peer_id:long = Bool;
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

	return x.GetBuf()
}

func (m *TLUserCheckPrivacy) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserCheckPrivacy) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc56e1eaa:
		// user.checkPrivacy flags:# user_id:long key_type:int peer_id:long = Bool;

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

func (m *TLUserCheckPrivacy) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserAddPeerSettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserAddPeerSettings) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_addPeerSettings))

	switch uint32(m.Constructor) {
	case 0xcae22763:
		// user.addPeerSettings user_id:long peer_type:int peer_id:long settings:PeerSettings = Bool;
		x.UInt(0xcae22763)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Bytes(m.GetSettings().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserAddPeerSettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserAddPeerSettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xcae22763:
		// user.addPeerSettings user_id:long peer_type:int peer_id:long settings:PeerSettings = Bool;

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

func (m *TLUserAddPeerSettings) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetPeerSettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetPeerSettings) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getPeerSettings))

	switch uint32(m.Constructor) {
	case 0xd02ef67:
		// user.getPeerSettings user_id:long peer_type:int peer_id:long = PeerSettings;
		x.UInt(0xd02ef67)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetPeerSettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetPeerSettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xd02ef67:
		// user.getPeerSettings user_id:long peer_type:int peer_id:long = PeerSettings;

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

func (m *TLUserGetPeerSettings) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserDeletePeerSettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserDeletePeerSettings) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_deletePeerSettings))

	switch uint32(m.Constructor) {
	case 0x5e891967:
		// user.deletePeerSettings user_id:long peer_type:int peer_id:long = Bool;
		x.UInt(0x5e891967)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserDeletePeerSettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserDeletePeerSettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x5e891967:
		// user.deletePeerSettings user_id:long peer_type:int peer_id:long = Bool;

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

func (m *TLUserDeletePeerSettings) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserChangePhone
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserChangePhone) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_changePhone))

	switch uint32(m.Constructor) {
	case 0xff7a575b:
		// user.changePhone user_id:long phone:string = Bool;
		x.UInt(0xff7a575b)

		// no flags

		x.Long(m.GetUserId())
		x.String(m.GetPhone())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserChangePhone) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserChangePhone) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xff7a575b:
		// user.changePhone user_id:long phone:string = Bool;

		// not has flags

		m.UserId = dBuf.Long()
		m.Phone = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserChangePhone) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserCreateNewPredefinedUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserCreateNewPredefinedUser) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_createNewPredefinedUser))

	switch uint32(m.Constructor) {
	case 0x57493241:
		// user.createNewPredefinedUser flags:# phone:string first_name:string last_name:flags.0?string username:string code:string verified:flags.1?true = PredefinedUser;
		x.UInt(0x57493241)

		// set flags
		var flags uint32 = 0

		if m.GetLastName() != nil {
			flags |= 1 << 0
		}

		if m.GetVerified() == true {
			flags |= 1 << 1
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.String(m.GetPhone())
		x.String(m.GetFirstName())
		if m.GetLastName() != nil {
			x.String(m.GetLastName().Value)
		}

		x.String(m.GetUsername())
		x.String(m.GetCode())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserCreateNewPredefinedUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserCreateNewPredefinedUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x57493241:
		// user.createNewPredefinedUser flags:# phone:string first_name:string last_name:flags.0?string username:string code:string verified:flags.1?true = PredefinedUser;

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.Phone = dBuf.String()
		m.FirstName = dBuf.String()
		if (flags & (1 << 0)) != 0 {
			m.LastName = &types.StringValue{Value: dBuf.String()}
		}

		m.Username = dBuf.String()
		m.Code = dBuf.String()
		if (flags & (1 << 1)) != 0 {
			m.Verified = true
		}
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserCreateNewPredefinedUser) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetPredefinedUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetPredefinedUser) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getPredefinedUser))

	switch uint32(m.Constructor) {
	case 0x34376b25:
		// user.getPredefinedUser phone:string = PredefinedUser;
		x.UInt(0x34376b25)

		// no flags

		x.String(m.GetPhone())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetPredefinedUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetPredefinedUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x34376b25:
		// user.getPredefinedUser phone:string = PredefinedUser;

		// not has flags

		m.Phone = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetPredefinedUser) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetAllPredefinedUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetAllPredefinedUser) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getAllPredefinedUser))

	switch uint32(m.Constructor) {
	case 0xc1353fe5:
		// user.getAllPredefinedUser = Vector<PredefinedUser>;
		x.UInt(0xc1353fe5)

		// no flags

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetAllPredefinedUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetAllPredefinedUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc1353fe5:
		// user.getAllPredefinedUser = Vector<PredefinedUser>;

		// not has flags

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetAllPredefinedUser) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserUpdatePredefinedFirstAndLastName
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUpdatePredefinedFirstAndLastName) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_updatePredefinedFirstAndLastName))

	switch uint32(m.Constructor) {
	case 0x3a3aa596:
		// user.updatePredefinedFirstAndLastName flags:# phone:string first_name:string last_name:flags.0?string = PredefinedUser;
		x.UInt(0x3a3aa596)

		// set flags
		var flags uint32 = 0

		if m.GetLastName() != nil {
			flags |= 1 << 0
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.String(m.GetPhone())
		x.String(m.GetFirstName())
		if m.GetLastName() != nil {
			x.String(m.GetLastName().Value)
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserUpdatePredefinedFirstAndLastName) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUpdatePredefinedFirstAndLastName) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x3a3aa596:
		// user.updatePredefinedFirstAndLastName flags:# phone:string first_name:string last_name:flags.0?string = PredefinedUser;

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.Phone = dBuf.String()
		m.FirstName = dBuf.String()
		if (flags & (1 << 0)) != 0 {
			m.LastName = &types.StringValue{Value: dBuf.String()}
		}

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserUpdatePredefinedFirstAndLastName) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserUpdatePredefinedVerified
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUpdatePredefinedVerified) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_updatePredefinedVerified))

	switch uint32(m.Constructor) {
	case 0xbaf5b249:
		// user.updatePredefinedVerified flags:# phone:string verified:flags.1?true = PredefinedUser;
		x.UInt(0xbaf5b249)

		// set flags
		var flags uint32 = 0

		if m.GetVerified() == true {
			flags |= 1 << 1
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.String(m.GetPhone())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserUpdatePredefinedVerified) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUpdatePredefinedVerified) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xbaf5b249:
		// user.updatePredefinedVerified flags:# phone:string verified:flags.1?true = PredefinedUser;

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.Phone = dBuf.String()
		if (flags & (1 << 1)) != 0 {
			m.Verified = true
		}
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserUpdatePredefinedVerified) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserUpdatePredefinedUsername
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUpdatePredefinedUsername) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_updatePredefinedUsername))

	switch uint32(m.Constructor) {
	case 0x4ba7bed2:
		// user.updatePredefinedUsername flags:# phone:string username:flags.1?string = PredefinedUser;
		x.UInt(0x4ba7bed2)

		// set flags
		var flags uint32 = 0

		if m.GetUsername() != nil {
			flags |= 1 << 1
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.String(m.GetPhone())
		if m.GetUsername() != nil {
			x.String(m.GetUsername().Value)
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserUpdatePredefinedUsername) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUpdatePredefinedUsername) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x4ba7bed2:
		// user.updatePredefinedUsername flags:# phone:string username:flags.1?string = PredefinedUser;

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.Phone = dBuf.String()
		if (flags & (1 << 1)) != 0 {
			m.Username = &types.StringValue{Value: dBuf.String()}
		}

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserUpdatePredefinedUsername) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserUpdatePredefinedCode
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUpdatePredefinedCode) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_updatePredefinedCode))

	switch uint32(m.Constructor) {
	case 0x60f68f67:
		// user.updatePredefinedCode phone:string code:string = PredefinedUser;
		x.UInt(0x60f68f67)

		// no flags

		x.String(m.GetPhone())
		x.String(m.GetCode())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserUpdatePredefinedCode) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUpdatePredefinedCode) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x60f68f67:
		// user.updatePredefinedCode phone:string code:string = PredefinedUser;

		// not has flags

		m.Phone = dBuf.String()
		m.Code = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserUpdatePredefinedCode) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserPredefinedBindRegisteredUserId
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserPredefinedBindRegisteredUserId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_predefinedBindRegisteredUserId))

	switch uint32(m.Constructor) {
	case 0x40f37a9:
		// user.predefinedBindRegisteredUserId phone:string registered_userId:long = Bool;
		x.UInt(0x40f37a9)

		// no flags

		x.String(m.GetPhone())
		x.Long(m.GetRegisteredUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserPredefinedBindRegisteredUserId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserPredefinedBindRegisteredUserId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x40f37a9:
		// user.predefinedBindRegisteredUserId phone:string registered_userId:long = Bool;

		// not has flags

		m.Phone = dBuf.String()
		m.RegisteredUserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserPredefinedBindRegisteredUserId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserCreateNewUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserCreateNewUser) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_createNewUser))

	switch uint32(m.Constructor) {
	case 0x79e01881:
		// user.createNewUser secret_key_id:long phone:string country_code:string first_name:string last_name:string = ImmutableUser;
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

	return x.GetBuf()
}

func (m *TLUserCreateNewUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserCreateNewUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x79e01881:
		// user.createNewUser secret_key_id:long phone:string country_code:string first_name:string last_name:string = ImmutableUser;

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

func (m *TLUserCreateNewUser) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserBlockPeer
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserBlockPeer) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_blockPeer))

	switch uint32(m.Constructor) {
	case 0x81062eb0:
		// user.blockPeer user_id:long peer_type:int peer_id:long = Bool;
		x.UInt(0x81062eb0)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserBlockPeer) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserBlockPeer) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x81062eb0:
		// user.blockPeer user_id:long peer_type:int peer_id:long = Bool;

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

func (m *TLUserBlockPeer) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserUnBlockPeer
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUnBlockPeer) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_unBlockPeer))

	switch uint32(m.Constructor) {
	case 0xdee7160d:
		// user.unBlockPeer user_id:long peer_type:int peer_id:long = Bool;
		x.UInt(0xdee7160d)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserUnBlockPeer) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUnBlockPeer) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xdee7160d:
		// user.unBlockPeer user_id:long peer_type:int peer_id:long = Bool;

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

func (m *TLUserUnBlockPeer) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserBlockedByUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserBlockedByUser) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_blockedByUser))

	switch uint32(m.Constructor) {
	case 0xbba0058e:
		// user.blockedByUser user_id:long peer_user_id:long = Bool;
		x.UInt(0xbba0058e)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetPeerUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserBlockedByUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserBlockedByUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xbba0058e:
		// user.blockedByUser user_id:long peer_user_id:long = Bool;

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerUserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserBlockedByUser) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserIsBlockedByUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserIsBlockedByUser) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_isBlockedByUser))

	switch uint32(m.Constructor) {
	case 0x8caeb1df:
		// user.isBlockedByUser user_id:long peer_user_id:long = Bool;
		x.UInt(0x8caeb1df)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetPeerUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserIsBlockedByUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserIsBlockedByUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x8caeb1df:
		// user.isBlockedByUser user_id:long peer_user_id:long = Bool;

		// not has flags

		m.UserId = dBuf.Long()
		m.PeerUserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserIsBlockedByUser) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserCheckBlockUserList
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserCheckBlockUserList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_checkBlockUserList))

	switch uint32(m.Constructor) {
	case 0xc3fd70f0:
		// user.checkBlockUserList user_id:long id:Vector<long> = Vector<long>;
		x.UInt(0xc3fd70f0)

		// no flags

		x.Long(m.GetUserId())

		x.VectorLong(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserCheckBlockUserList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserCheckBlockUserList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc3fd70f0:
		// user.checkBlockUserList user_id:long id:Vector<long> = Vector<long>;

		// not has flags

		m.UserId = dBuf.Long()

		m.Id = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserCheckBlockUserList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetBlockedList
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetBlockedList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getBlockedList))

	switch uint32(m.Constructor) {
	case 0x23ffc348:
		// user.getBlockedList user_id:long offset:int limit:int = Vector<PeerBlocked>;
		x.UInt(0x23ffc348)

		// no flags

		x.Long(m.GetUserId())
		x.Int(m.GetOffset())
		x.Int(m.GetLimit())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetBlockedList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetBlockedList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x23ffc348:
		// user.getBlockedList user_id:long offset:int limit:int = Vector<PeerBlocked>;

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

func (m *TLUserGetBlockedList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetContactSignUpNotification
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetContactSignUpNotification) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getContactSignUpNotification))

	switch uint32(m.Constructor) {
	case 0xe4d1d3d6:
		// user.getContactSignUpNotification user_id:long = Bool;
		x.UInt(0xe4d1d3d6)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetContactSignUpNotification) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetContactSignUpNotification) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe4d1d3d6:
		// user.getContactSignUpNotification user_id:long = Bool;

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetContactSignUpNotification) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserSetContactSignUpNotification
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserSetContactSignUpNotification) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_setContactSignUpNotification))

	switch uint32(m.Constructor) {
	case 0x85a17361:
		// user.setContactSignUpNotification user_id:long silent:Bool = Bool;
		x.UInt(0x85a17361)

		// no flags

		x.Long(m.GetUserId())
		x.Bytes(m.GetSilent().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserSetContactSignUpNotification) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSetContactSignUpNotification) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x85a17361:
		// user.setContactSignUpNotification user_id:long silent:Bool = Bool;

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

func (m *TLUserSetContactSignUpNotification) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetContentSettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetContentSettings) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getContentSettings))

	switch uint32(m.Constructor) {
	case 0x94c3ad9f:
		// user.getContentSettings user_id:long = account.ContentSettings;
		x.UInt(0x94c3ad9f)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetContentSettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetContentSettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x94c3ad9f:
		// user.getContentSettings user_id:long = account.ContentSettings;

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetContentSettings) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserSetContentSettings
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserSetContentSettings) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_setContentSettings))

	switch uint32(m.Constructor) {
	case 0x9d63fe6b:
		// user.setContentSettings flags:# user_id:long sensitive_enabled:flags.0?true = Bool;
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

	return x.GetBuf()
}

func (m *TLUserSetContentSettings) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSetContentSettings) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x9d63fe6b:
		// user.setContentSettings flags:# user_id:long sensitive_enabled:flags.0?true = Bool;

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

func (m *TLUserSetContentSettings) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserDeleteContact
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserDeleteContact) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_deleteContact))

	switch uint32(m.Constructor) {
	case 0xc6018219:
		// user.deleteContact user_id:long id:long = Bool;
		x.UInt(0xc6018219)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserDeleteContact) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserDeleteContact) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc6018219:
		// user.deleteContact user_id:long id:long = Bool;

		// not has flags

		m.UserId = dBuf.Long()
		m.Id = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserDeleteContact) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetContactList
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetContactList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getContactList))

	switch uint32(m.Constructor) {
	case 0xc74bd161:
		// user.getContactList user_id:long = Vector<ContactData>;
		x.UInt(0xc74bd161)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetContactList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetContactList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc74bd161:
		// user.getContactList user_id:long = Vector<ContactData>;

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetContactList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetContactIdList
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetContactIdList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getContactIdList))

	switch uint32(m.Constructor) {
	case 0xf1dd983e:
		// user.getContactIdList user_id:long = Vector<long>;
		x.UInt(0xf1dd983e)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetContactIdList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetContactIdList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf1dd983e:
		// user.getContactIdList user_id:long = Vector<long>;

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetContactIdList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetContact
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetContact) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getContact))

	switch uint32(m.Constructor) {
	case 0xdb728be3:
		// user.getContact user_id:long id:long = ContactData;
		x.UInt(0xdb728be3)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetContact) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetContact) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xdb728be3:
		// user.getContact user_id:long id:long = ContactData;

		// not has flags

		m.UserId = dBuf.Long()
		m.Id = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetContact) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserAddContact
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserAddContact) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_addContact))

	switch uint32(m.Constructor) {
	case 0x79c4bd0e:
		// user.addContact user_id:long add_phone_privacy_exception:Bool id:long first_name:string last_name:string phone:string = Bool;
		x.UInt(0x79c4bd0e)

		// no flags

		x.Long(m.GetUserId())
		x.Bytes(m.GetAddPhonePrivacyException().Encode(layer))
		x.Long(m.GetId())
		x.String(m.GetFirstName())
		x.String(m.GetLastName())
		x.String(m.GetPhone())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserAddContact) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserAddContact) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x79c4bd0e:
		// user.addContact user_id:long add_phone_privacy_exception:Bool id:long first_name:string last_name:string phone:string = Bool;

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

func (m *TLUserAddContact) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserCheckContact
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserCheckContact) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_checkContact))

	switch uint32(m.Constructor) {
	case 0x82a758a4:
		// user.checkContact user_id:long id:long = Bool;
		x.UInt(0x82a758a4)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserCheckContact) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserCheckContact) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x82a758a4:
		// user.checkContact user_id:long id:long = Bool;

		// not has flags

		m.UserId = dBuf.Long()
		m.Id = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserCheckContact) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserImportContacts
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserImportContacts) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_importContacts))

	switch uint32(m.Constructor) {
	case 0x9a00f792:
		// user.importContacts user_id:long contacts:Vector<InputContact> = UserImportedContacts;
		x.UInt(0x9a00f792)

		// no flags

		x.Long(m.GetUserId())

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetContacts())))
		for _, v := range m.GetContacts() {
			x.Bytes((*v).Encode(layer))
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserImportContacts) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserImportContacts) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x9a00f792:
		// user.importContacts user_id:long contacts:Vector<InputContact> = UserImportedContacts;

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

func (m *TLUserImportContacts) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetCountryCode
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetCountryCode) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getCountryCode))

	switch uint32(m.Constructor) {
	case 0x12006832:
		// user.getCountryCode user_id:long = String;
		x.UInt(0x12006832)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetCountryCode) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetCountryCode) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x12006832:
		// user.getCountryCode user_id:long = String;

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetCountryCode) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserUpdateAbout
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUpdateAbout) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_updateAbout))

	switch uint32(m.Constructor) {
	case 0xdb00f187:
		// user.updateAbout user_id:long about:string = Bool;
		x.UInt(0xdb00f187)

		// no flags

		x.Long(m.GetUserId())
		x.String(m.GetAbout())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserUpdateAbout) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUpdateAbout) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xdb00f187:
		// user.updateAbout user_id:long about:string = Bool;

		// not has flags

		m.UserId = dBuf.Long()
		m.About = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserUpdateAbout) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserUpdateFirstAndLastName
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUpdateFirstAndLastName) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_updateFirstAndLastName))

	switch uint32(m.Constructor) {
	case 0xcb6685ec:
		// user.updateFirstAndLastName user_id:long first_name:string last_name:string = Bool;
		x.UInt(0xcb6685ec)

		// no flags

		x.Long(m.GetUserId())
		x.String(m.GetFirstName())
		x.String(m.GetLastName())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserUpdateFirstAndLastName) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUpdateFirstAndLastName) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xcb6685ec:
		// user.updateFirstAndLastName user_id:long first_name:string last_name:string = Bool;

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

func (m *TLUserUpdateFirstAndLastName) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserUpdateVerified
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUpdateVerified) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_updateVerified))

	switch uint32(m.Constructor) {
	case 0x24c92963:
		// user.updateVerified user_id:long verified:Bool = Bool;
		x.UInt(0x24c92963)

		// no flags

		x.Long(m.GetUserId())
		x.Bytes(m.GetVerified().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserUpdateVerified) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUpdateVerified) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x24c92963:
		// user.updateVerified user_id:long verified:Bool = Bool;

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

func (m *TLUserUpdateVerified) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserUpdateUsername
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUpdateUsername) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_updateUsername))

	switch uint32(m.Constructor) {
	case 0xf54d1e71:
		// user.updateUsername user_id:long username:string = Bool;
		x.UInt(0xf54d1e71)

		// no flags

		x.Long(m.GetUserId())
		x.String(m.GetUsername())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserUpdateUsername) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUpdateUsername) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf54d1e71:
		// user.updateUsername user_id:long username:string = Bool;

		// not has flags

		m.UserId = dBuf.Long()
		m.Username = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserUpdateUsername) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserUpdateProfilePhoto
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserUpdateProfilePhoto) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_updateProfilePhoto))

	switch uint32(m.Constructor) {
	case 0x3b740f87:
		// user.updateProfilePhoto user_id:long id:long = Int64;
		x.UInt(0x3b740f87)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserUpdateProfilePhoto) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserUpdateProfilePhoto) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x3b740f87:
		// user.updateProfilePhoto user_id:long id:long = Int64;

		// not has flags

		m.UserId = dBuf.Long()
		m.Id = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserUpdateProfilePhoto) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserDeleteProfilePhotos
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserDeleteProfilePhotos) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_deleteProfilePhotos))

	switch uint32(m.Constructor) {
	case 0x2be3620e:
		// user.deleteProfilePhotos user_id:long id:Vector<long> = Int64;
		x.UInt(0x2be3620e)

		// no flags

		x.Long(m.GetUserId())

		x.VectorLong(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserDeleteProfilePhotos) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserDeleteProfilePhotos) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x2be3620e:
		// user.deleteProfilePhotos user_id:long id:Vector<long> = Int64;

		// not has flags

		m.UserId = dBuf.Long()

		m.Id = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserDeleteProfilePhotos) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetProfilePhotos
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetProfilePhotos) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getProfilePhotos))

	switch uint32(m.Constructor) {
	case 0xdc66c146:
		// user.getProfilePhotos user_id:long = Vector<long>;
		x.UInt(0xdc66c146)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetProfilePhotos) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetProfilePhotos) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xdc66c146:
		// user.getProfilePhotos user_id:long = Vector<long>;

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetProfilePhotos) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserSetBotCommands
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserSetBotCommands) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_setBotCommands))

	switch uint32(m.Constructor) {
	case 0x753ba916:
		// user.setBotCommands user_id:long bot_id:long commands:Vector<BotCommand> = Bool;
		x.UInt(0x753ba916)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetBotId())

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetCommands())))
		for _, v := range m.GetCommands() {
			x.Bytes((*v).Encode(layer))
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserSetBotCommands) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserSetBotCommands) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x753ba916:
		// user.setBotCommands user_id:long bot_id:long commands:Vector<BotCommand> = Bool;

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

func (m *TLUserSetBotCommands) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserIsBot
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserIsBot) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_isBot))

	switch uint32(m.Constructor) {
	case 0xc772c7ee:
		// user.isBot id:long = Bool;
		x.UInt(0xc772c7ee)

		// no flags

		x.Long(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserIsBot) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserIsBot) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc772c7ee:
		// user.isBot id:long = Bool;

		// not has flags

		m.Id = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserIsBot) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetBotInfo
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetBotInfo) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getBotInfo))

	switch uint32(m.Constructor) {
	case 0x34663710:
		// user.getBotInfo bot_id:long = BotInfo;
		x.UInt(0x34663710)

		// no flags

		x.Long(m.GetBotId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetBotInfo) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetBotInfo) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x34663710:
		// user.getBotInfo bot_id:long = BotInfo;

		// not has flags

		m.BotId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetBotInfo) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUserGetFullUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLUserGetFullUser) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_user_getFullUser))

	switch uint32(m.Constructor) {
	case 0xfd10e13a:
		// user.getFullUser self_user_id:long id:long = users.UserFull;
		x.UInt(0xfd10e13a)

		// no flags

		x.Long(m.GetSelfUserId())
		x.Long(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUserGetFullUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserGetFullUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xfd10e13a:
		// user.getFullUser self_user_id:long id:long = users.UserFull;

		// not has flags

		m.SelfUserId = dBuf.Long()
		m.Id = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUserGetFullUser) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// Vector_LastSeenData
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_LastSeenData) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
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

func (m *Vector_LastSeenData) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// Vector_ImmutableUser
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_ImmutableUser) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
}

func (m *Vector_ImmutableUser) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*ImmutableUser, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(ImmutableUser)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_ImmutableUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *Vector_ImmutableUser) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// Vector_PeerPeerNotifySettings
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_PeerPeerNotifySettings) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
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

func (m *Vector_PeerPeerNotifySettings) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// Vector_PrivacyRule
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_PrivacyRule) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
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

func (m *Vector_PrivacyRule) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// Vector_PredefinedUser
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_PredefinedUser) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
}

func (m *Vector_PredefinedUser) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*mtproto.PredefinedUser, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(mtproto.PredefinedUser)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_PredefinedUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *Vector_PredefinedUser) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

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

// Vector_PeerBlocked
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_PeerBlocked) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
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

func (m *Vector_PeerBlocked) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// Vector_ContactData
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_ContactData) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
}

func (m *Vector_ContactData) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*ContactData, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(ContactData)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_ContactData) CalcByteSize(layer int32) int {
	return 0
}

func (m *Vector_ContactData) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}
