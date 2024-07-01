/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package user

const (
	Predicate_userImportedContacts              = "userImportedContacts"
	Predicate_usersDataFound                    = "usersDataFound"
	Predicate_usersIdFound                      = "usersIdFound"
	Predicate_peerPeerNotifySettings            = "peerPeerNotifySettings"
	Predicate_lastSeenData                      = "lastSeenData"
	Predicate_user_getLastSeens                 = "user_getLastSeens"
	Predicate_user_updateLastSeen               = "user_updateLastSeen"
	Predicate_user_getLastSeen                  = "user_getLastSeen"
	Predicate_user_getImmutableUser             = "user_getImmutableUser"
	Predicate_user_getMutableUsers              = "user_getMutableUsers"
	Predicate_user_getImmutableUserByPhone      = "user_getImmutableUserByPhone"
	Predicate_user_getImmutableUserByToken      = "user_getImmutableUserByToken"
	Predicate_user_setAccountDaysTTL            = "user_setAccountDaysTTL"
	Predicate_user_getAccountDaysTTL            = "user_getAccountDaysTTL"
	Predicate_user_getNotifySettings            = "user_getNotifySettings"
	Predicate_user_getNotifySettingsList        = "user_getNotifySettingsList"
	Predicate_user_setNotifySettings            = "user_setNotifySettings"
	Predicate_user_resetNotifySettings          = "user_resetNotifySettings"
	Predicate_user_getAllNotifySettings         = "user_getAllNotifySettings"
	Predicate_user_getGlobalPrivacySettings     = "user_getGlobalPrivacySettings"
	Predicate_user_setGlobalPrivacySettings     = "user_setGlobalPrivacySettings"
	Predicate_user_getPrivacy                   = "user_getPrivacy"
	Predicate_user_setPrivacy                   = "user_setPrivacy"
	Predicate_user_checkPrivacy                 = "user_checkPrivacy"
	Predicate_user_addPeerSettings              = "user_addPeerSettings"
	Predicate_user_getPeerSettings              = "user_getPeerSettings"
	Predicate_user_deletePeerSettings           = "user_deletePeerSettings"
	Predicate_user_changePhone                  = "user_changePhone"
	Predicate_user_createNewUser                = "user_createNewUser"
	Predicate_user_deleteUser                   = "user_deleteUser"
	Predicate_user_blockPeer                    = "user_blockPeer"
	Predicate_user_unBlockPeer                  = "user_unBlockPeer"
	Predicate_user_blockedByUser                = "user_blockedByUser"
	Predicate_user_isBlockedByUser              = "user_isBlockedByUser"
	Predicate_user_checkBlockUserList           = "user_checkBlockUserList"
	Predicate_user_getBlockedList               = "user_getBlockedList"
	Predicate_user_getContactSignUpNotification = "user_getContactSignUpNotification"
	Predicate_user_setContactSignUpNotification = "user_setContactSignUpNotification"
	Predicate_user_getContentSettings           = "user_getContentSettings"
	Predicate_user_setContentSettings           = "user_setContentSettings"
	Predicate_user_deleteContact                = "user_deleteContact"
	Predicate_user_getContactList               = "user_getContactList"
	Predicate_user_getContactIdList             = "user_getContactIdList"
	Predicate_user_getContact                   = "user_getContact"
	Predicate_user_addContact                   = "user_addContact"
	Predicate_user_checkContact                 = "user_checkContact"
	Predicate_user_getImportersByPhone          = "user_getImportersByPhone"
	Predicate_user_deleteImportersByPhone       = "user_deleteImportersByPhone"
	Predicate_user_importContacts               = "user_importContacts"
	Predicate_user_getCountryCode               = "user_getCountryCode"
	Predicate_user_updateAbout                  = "user_updateAbout"
	Predicate_user_updateFirstAndLastName       = "user_updateFirstAndLastName"
	Predicate_user_updateVerified               = "user_updateVerified"
	Predicate_user_updateUsername               = "user_updateUsername"
	Predicate_user_updateProfilePhoto           = "user_updateProfilePhoto"
	Predicate_user_deleteProfilePhotos          = "user_deleteProfilePhotos"
	Predicate_user_getProfilePhotos             = "user_getProfilePhotos"
	Predicate_user_setBotCommands               = "user_setBotCommands"
	Predicate_user_isBot                        = "user_isBot"
	Predicate_user_getBotInfo                   = "user_getBotInfo"
	Predicate_user_checkBots                    = "user_checkBots"
	Predicate_user_getFullUser                  = "user_getFullUser"
	Predicate_user_updateEmojiStatus            = "user_updateEmojiStatus"
	Predicate_user_getUserDataById              = "user_getUserDataById"
	Predicate_user_getUserDataListByIdList      = "user_getUserDataListByIdList"
	Predicate_user_getUserDataByToken           = "user_getUserDataByToken"
	Predicate_user_search                       = "user_search"
	Predicate_user_updateBotData                = "user_updateBotData"
	Predicate_user_getImmutableUserV2           = "user_getImmutableUserV2"
	Predicate_user_getMutableUsersV2            = "user_getMutableUsersV2"
	Predicate_user_createNewTestUser            = "user_createNewTestUser"
	Predicate_user_editCloseFriends             = "user_editCloseFriends"
	Predicate_user_setStoriesMaxId              = "user_setStoriesMaxId"
	Predicate_user_setColor                     = "user_setColor"
	Predicate_user_updateBirthday               = "user_updateBirthday"
	Predicate_user_getBirthdays                 = "user_getBirthdays"
	Predicate_user_setStoriesHidden             = "user_setStoriesHidden"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_userImportedContacts: {
		0: 1256160192, // 0x4adf7bc0

	},
	Predicate_usersDataFound: {
		0: 1067703239, // 0x3fa3dbc7

	},
	Predicate_usersIdFound: {
		0: -2134594054, // 0x80c4adfa

	},
	Predicate_peerPeerNotifySettings: {
		0: 1894399913, // 0x70ea3fa9

	},
	Predicate_lastSeenData: {
		0: -1280204321, // 0xb3b1a1df

	},
	Predicate_user_getLastSeens: {
		0: 2090958337, // 0x7ca17e01

	},
	Predicate_user_updateLastSeen: {
		0: -46114259, // 0xfd405a2d

	},
	Predicate_user_getLastSeen: {
		0: -1860581154, // 0x9119c8de

	},
	Predicate_user_getImmutableUser: {
		0: 929720132, // 0x376a6744

	},
	Predicate_user_getMutableUsers: {
		0: -1657068585, // 0x9d3b23d7

	},
	Predicate_user_getImmutableUserByPhone: {
		0: -373067804, // 0xe9c36fe4

	},
	Predicate_user_getImmutableUserByToken: {
		0: -12709005, // 0xff3e1373

	},
	Predicate_user_setAccountDaysTTL: {
		0: -766178484, // 0xd2550b4c

	},
	Predicate_user_getAccountDaysTTL: {
		0: -1299956000, // 0xb2843ee0

	},
	Predicate_user_getNotifySettings: {
		0: 1085028198, // 0x40ac3766

	},
	Predicate_user_getNotifySettingsList: {
		0: -463137380, // 0xe465159c

	},
	Predicate_user_setNotifySettings: {
		0: -907188763, // 0xc9ed65e5

	},
	Predicate_user_resetNotifySettings: {
		0: 235380084, // 0xe079d74

	},
	Predicate_user_getAllNotifySettings: {
		0: 1435658357, // 0x55926875

	},
	Predicate_user_getGlobalPrivacySettings: {
		0: 2012672274, // 0x77f6f112

	},
	Predicate_user_setGlobalPrivacySettings: {
		0: -1934257490, // 0x8cb592ae

	},
	Predicate_user_getPrivacy: {
		0: -1656708172, // 0x9d40a3b4

	},
	Predicate_user_setPrivacy: {
		0: -2007650929, // 0x8855ad8f

	},
	Predicate_user_checkPrivacy: {
		0: -982638934, // 0xc56e1eaa

	},
	Predicate_user_addPeerSettings: {
		0: -891148445, // 0xcae22763

	},
	Predicate_user_getPeerSettings: {
		0: 218296167, // 0xd02ef67

	},
	Predicate_user_deletePeerSettings: {
		0: 1586043239, // 0x5e891967

	},
	Predicate_user_changePhone: {
		0: -8759461, // 0xff7a575b

	},
	Predicate_user_createNewUser: {
		0: 2044729473, // 0x79e01881

	},
	Predicate_user_deleteUser: {
		0: 2132777160, // 0x7f1f98c8

	},
	Predicate_user_blockPeer: {
		0: -2130301264, // 0x81062eb0

	},
	Predicate_user_unBlockPeer: {
		0: -555280883, // 0xdee7160d

	},
	Predicate_user_blockedByUser: {
		0: -1147140722, // 0xbba0058e

	},
	Predicate_user_isBlockedByUser: {
		0: -1934708257, // 0x8caeb1df

	},
	Predicate_user_checkBlockUserList: {
		0: -1006800656, // 0xc3fd70f0

	},
	Predicate_user_getBlockedList: {
		0: 603964232, // 0x23ffc348

	},
	Predicate_user_getContactSignUpNotification: {
		0: -456010794, // 0xe4d1d3d6

	},
	Predicate_user_setContactSignUpNotification: {
		0: -2053016735, // 0x85a17361

	},
	Predicate_user_getContentSettings: {
		0: -1799115361, // 0x94c3ad9f

	},
	Predicate_user_setContentSettings: {
		0: -1654391189, // 0x9d63fe6b

	},
	Predicate_user_deleteContact: {
		0: -972979687, // 0xc6018219

	},
	Predicate_user_getContactList: {
		0: -951332511, // 0xc74bd161

	},
	Predicate_user_getContactIdList: {
		0: -237135810, // 0xf1dd983e

	},
	Predicate_user_getContact: {
		0: -613250077, // 0xdb728be3

	},
	Predicate_user_addContact: {
		0: 2042936590, // 0x79c4bd0e

	},
	Predicate_user_checkContact: {
		0: -2102962012, // 0x82a758a4

	},
	Predicate_user_getImportersByPhone: {
		0: 1202356754, // 0x47aa8212

	},
	Predicate_user_deleteImportersByPhone: {
		0: 390978644, // 0x174ddc54

	},
	Predicate_user_importContacts: {
		0: -1711212654, // 0x9a00f792

	},
	Predicate_user_getCountryCode: {
		0: 302016562, // 0x12006832

	},
	Predicate_user_updateAbout: {
		0: -620695161, // 0xdb00f187

	},
	Predicate_user_updateFirstAndLastName: {
		0: -882473492, // 0xcb6685ec

	},
	Predicate_user_updateVerified: {
		0: 617163107, // 0x24c92963

	},
	Predicate_user_updateUsername: {
		0: -179495311, // 0xf54d1e71

	},
	Predicate_user_updateProfilePhoto: {
		0: 997461895, // 0x3b740f87

	},
	Predicate_user_deleteProfilePhotos: {
		0: 736322062, // 0x2be3620e

	},
	Predicate_user_getProfilePhotos: {
		0: -597245626, // 0xdc66c146

	},
	Predicate_user_setBotCommands: {
		0: 1966844182, // 0x753ba916

	},
	Predicate_user_isBot: {
		0: -948779026, // 0xc772c7ee

	},
	Predicate_user_getBotInfo: {
		0: 879114000, // 0x34663710

	},
	Predicate_user_checkBots: {
		0: 1935999169, // 0x736500c1

	},
	Predicate_user_getFullUser: {
		0: -49225414, // 0xfd10e13a

	},
	Predicate_user_updateEmojiStatus: {
		0: -121062696, // 0xf8c8bad8

	},
	Predicate_user_getUserDataById: {
		0: 62615811, // 0x3bb7103

	},
	Predicate_user_getUserDataListByIdList: {
		0: -2121142279, // 0x8191eff9

	},
	Predicate_user_getUserDataByToken: {
		0: 1057580446, // 0x3f09659e

	},
	Predicate_user_search: {
		0: 1882568397, // 0x7035b6cd

	},
	Predicate_user_updateBotData: {
		0: -1174586898, // 0xb9fd39ee

	},
	Predicate_user_getImmutableUserV2: {
		0: 806009420, // 0x300aba4c

	},
	Predicate_user_getMutableUsersV2: {
		0: -1795585240, // 0x94f98b28

	},
	Predicate_user_createNewTestUser: {
		0: 1282329771, // 0x4c6eccab

	},
	Predicate_user_editCloseFriends: {
		0: -2044429563, // 0x86247b05

	},
	Predicate_user_setStoriesMaxId: {
		0: 1391834736, // 0x52f5b670

	},
	Predicate_user_setColor: {
		0: 586812791, // 0x22fa0d77

	},
	Predicate_user_updateBirthday: {
		0: 1484434322, // 0x587aab92

	},
	Predicate_user_getBirthdays: {
		0: -24199258, // 0xfe8ebfa6

	},
	Predicate_user_setStoriesHidden: {
		0: -138012584, // 0xf7c61858

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	1256160192:  Predicate_userImportedContacts,              // 0x4adf7bc0
	1067703239:  Predicate_usersDataFound,                    // 0x3fa3dbc7
	-2134594054: Predicate_usersIdFound,                      // 0x80c4adfa
	1894399913:  Predicate_peerPeerNotifySettings,            // 0x70ea3fa9
	-1280204321: Predicate_lastSeenData,                      // 0xb3b1a1df
	2090958337:  Predicate_user_getLastSeens,                 // 0x7ca17e01
	-46114259:   Predicate_user_updateLastSeen,               // 0xfd405a2d
	-1860581154: Predicate_user_getLastSeen,                  // 0x9119c8de
	929720132:   Predicate_user_getImmutableUser,             // 0x376a6744
	-1657068585: Predicate_user_getMutableUsers,              // 0x9d3b23d7
	-373067804:  Predicate_user_getImmutableUserByPhone,      // 0xe9c36fe4
	-12709005:   Predicate_user_getImmutableUserByToken,      // 0xff3e1373
	-766178484:  Predicate_user_setAccountDaysTTL,            // 0xd2550b4c
	-1299956000: Predicate_user_getAccountDaysTTL,            // 0xb2843ee0
	1085028198:  Predicate_user_getNotifySettings,            // 0x40ac3766
	-463137380:  Predicate_user_getNotifySettingsList,        // 0xe465159c
	-907188763:  Predicate_user_setNotifySettings,            // 0xc9ed65e5
	235380084:   Predicate_user_resetNotifySettings,          // 0xe079d74
	1435658357:  Predicate_user_getAllNotifySettings,         // 0x55926875
	2012672274:  Predicate_user_getGlobalPrivacySettings,     // 0x77f6f112
	-1934257490: Predicate_user_setGlobalPrivacySettings,     // 0x8cb592ae
	-1656708172: Predicate_user_getPrivacy,                   // 0x9d40a3b4
	-2007650929: Predicate_user_setPrivacy,                   // 0x8855ad8f
	-982638934:  Predicate_user_checkPrivacy,                 // 0xc56e1eaa
	-891148445:  Predicate_user_addPeerSettings,              // 0xcae22763
	218296167:   Predicate_user_getPeerSettings,              // 0xd02ef67
	1586043239:  Predicate_user_deletePeerSettings,           // 0x5e891967
	-8759461:    Predicate_user_changePhone,                  // 0xff7a575b
	2044729473:  Predicate_user_createNewUser,                // 0x79e01881
	2132777160:  Predicate_user_deleteUser,                   // 0x7f1f98c8
	-2130301264: Predicate_user_blockPeer,                    // 0x81062eb0
	-555280883:  Predicate_user_unBlockPeer,                  // 0xdee7160d
	-1147140722: Predicate_user_blockedByUser,                // 0xbba0058e
	-1934708257: Predicate_user_isBlockedByUser,              // 0x8caeb1df
	-1006800656: Predicate_user_checkBlockUserList,           // 0xc3fd70f0
	603964232:   Predicate_user_getBlockedList,               // 0x23ffc348
	-456010794:  Predicate_user_getContactSignUpNotification, // 0xe4d1d3d6
	-2053016735: Predicate_user_setContactSignUpNotification, // 0x85a17361
	-1799115361: Predicate_user_getContentSettings,           // 0x94c3ad9f
	-1654391189: Predicate_user_setContentSettings,           // 0x9d63fe6b
	-972979687:  Predicate_user_deleteContact,                // 0xc6018219
	-951332511:  Predicate_user_getContactList,               // 0xc74bd161
	-237135810:  Predicate_user_getContactIdList,             // 0xf1dd983e
	-613250077:  Predicate_user_getContact,                   // 0xdb728be3
	2042936590:  Predicate_user_addContact,                   // 0x79c4bd0e
	-2102962012: Predicate_user_checkContact,                 // 0x82a758a4
	1202356754:  Predicate_user_getImportersByPhone,          // 0x47aa8212
	390978644:   Predicate_user_deleteImportersByPhone,       // 0x174ddc54
	-1711212654: Predicate_user_importContacts,               // 0x9a00f792
	302016562:   Predicate_user_getCountryCode,               // 0x12006832
	-620695161:  Predicate_user_updateAbout,                  // 0xdb00f187
	-882473492:  Predicate_user_updateFirstAndLastName,       // 0xcb6685ec
	617163107:   Predicate_user_updateVerified,               // 0x24c92963
	-179495311:  Predicate_user_updateUsername,               // 0xf54d1e71
	997461895:   Predicate_user_updateProfilePhoto,           // 0x3b740f87
	736322062:   Predicate_user_deleteProfilePhotos,          // 0x2be3620e
	-597245626:  Predicate_user_getProfilePhotos,             // 0xdc66c146
	1966844182:  Predicate_user_setBotCommands,               // 0x753ba916
	-948779026:  Predicate_user_isBot,                        // 0xc772c7ee
	879114000:   Predicate_user_getBotInfo,                   // 0x34663710
	1935999169:  Predicate_user_checkBots,                    // 0x736500c1
	-49225414:   Predicate_user_getFullUser,                  // 0xfd10e13a
	-121062696:  Predicate_user_updateEmojiStatus,            // 0xf8c8bad8
	62615811:    Predicate_user_getUserDataById,              // 0x3bb7103
	-2121142279: Predicate_user_getUserDataListByIdList,      // 0x8191eff9
	1057580446:  Predicate_user_getUserDataByToken,           // 0x3f09659e
	1882568397:  Predicate_user_search,                       // 0x7035b6cd
	-1174586898: Predicate_user_updateBotData,                // 0xb9fd39ee
	806009420:   Predicate_user_getImmutableUserV2,           // 0x300aba4c
	-1795585240: Predicate_user_getMutableUsersV2,            // 0x94f98b28
	1282329771:  Predicate_user_createNewTestUser,            // 0x4c6eccab
	-2044429563: Predicate_user_editCloseFriends,             // 0x86247b05
	1391834736:  Predicate_user_setStoriesMaxId,              // 0x52f5b670
	586812791:   Predicate_user_setColor,                     // 0x22fa0d77
	1484434322:  Predicate_user_updateBirthday,               // 0x587aab92
	-24199258:   Predicate_user_getBirthdays,                 // 0xfe8ebfa6
	-138012584:  Predicate_user_setStoriesHidden,             // 0xf7c61858

}

func GetClazzID(clazzName string, layer int) int32 {
	if m, ok := clazzNameRegisters2[clazzName]; ok {
		m2, ok2 := m[layer]
		if ok2 {
			return m2
		}
		m2, ok2 = m[0]
		if ok2 {
			return m2
		}
	}
	return 0
}
