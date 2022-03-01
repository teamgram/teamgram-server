/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teagramio (teagram.io@gmail.com)
 */

package user

const (
	Predicate_peerPeerNotifySettings                = "peerPeerNotifySettings"
	Predicate_privacyKeyRules                       = "privacyKeyRules"
	Predicate_lastSeenData                          = "lastSeenData"
	Predicate_contactData                           = "contactData"
	Predicate_botData                               = "botData"
	Predicate_userData                              = "userData"
	Predicate_immutableUser                         = "immutableUser"
	Predicate_userImportedContacts                  = "userImportedContacts"
	Predicate_user_getLastSeens                     = "user_getLastSeens"
	Predicate_user_updateLastSeen                   = "user_updateLastSeen"
	Predicate_user_getLastSeen                      = "user_getLastSeen"
	Predicate_user_getImmutableUser                 = "user_getImmutableUser"
	Predicate_user_getMutableUsers                  = "user_getMutableUsers"
	Predicate_user_getImmutableUserByPhone          = "user_getImmutableUserByPhone"
	Predicate_user_getImmutableUserByToken          = "user_getImmutableUserByToken"
	Predicate_user_setAccountDaysTTL                = "user_setAccountDaysTTL"
	Predicate_user_getAccountDaysTTL                = "user_getAccountDaysTTL"
	Predicate_user_getNotifySettings                = "user_getNotifySettings"
	Predicate_user_setNotifySettings                = "user_setNotifySettings"
	Predicate_user_resetNotifySettings              = "user_resetNotifySettings"
	Predicate_user_getAllNotifySettings             = "user_getAllNotifySettings"
	Predicate_user_getGlobalPrivacySettings         = "user_getGlobalPrivacySettings"
	Predicate_user_setGlobalPrivacySettings         = "user_setGlobalPrivacySettings"
	Predicate_user_getPrivacy                       = "user_getPrivacy"
	Predicate_user_setPrivacy                       = "user_setPrivacy"
	Predicate_user_checkPrivacy                     = "user_checkPrivacy"
	Predicate_user_addPeerSettings                  = "user_addPeerSettings"
	Predicate_user_getPeerSettings                  = "user_getPeerSettings"
	Predicate_user_deletePeerSettings               = "user_deletePeerSettings"
	Predicate_user_changePhone                      = "user_changePhone"
	Predicate_user_createNewPredefinedUser          = "user_createNewPredefinedUser"
	Predicate_user_getPredefinedUser                = "user_getPredefinedUser"
	Predicate_user_getAllPredefinedUser             = "user_getAllPredefinedUser"
	Predicate_user_updatePredefinedFirstAndLastName = "user_updatePredefinedFirstAndLastName"
	Predicate_user_updatePredefinedVerified         = "user_updatePredefinedVerified"
	Predicate_user_updatePredefinedUsername         = "user_updatePredefinedUsername"
	Predicate_user_updatePredefinedCode             = "user_updatePredefinedCode"
	Predicate_user_predefinedBindRegisteredUserId   = "user_predefinedBindRegisteredUserId"
	Predicate_user_createNewUser                    = "user_createNewUser"
	Predicate_user_blockPeer                        = "user_blockPeer"
	Predicate_user_unBlockPeer                      = "user_unBlockPeer"
	Predicate_user_blockedByUser                    = "user_blockedByUser"
	Predicate_user_isBlockedByUser                  = "user_isBlockedByUser"
	Predicate_user_checkBlockUserList               = "user_checkBlockUserList"
	Predicate_user_getBlockedList                   = "user_getBlockedList"
	Predicate_user_getContactSignUpNotification     = "user_getContactSignUpNotification"
	Predicate_user_setContactSignUpNotification     = "user_setContactSignUpNotification"
	Predicate_user_getContentSettings               = "user_getContentSettings"
	Predicate_user_setContentSettings               = "user_setContentSettings"
	Predicate_user_deleteContact                    = "user_deleteContact"
	Predicate_user_getContactList                   = "user_getContactList"
	Predicate_user_getContactIdList                 = "user_getContactIdList"
	Predicate_user_getContact                       = "user_getContact"
	Predicate_user_addContact                       = "user_addContact"
	Predicate_user_checkContact                     = "user_checkContact"
	Predicate_user_importContacts                   = "user_importContacts"
	Predicate_user_getCountryCode                   = "user_getCountryCode"
	Predicate_user_updateAbout                      = "user_updateAbout"
	Predicate_user_updateFirstAndLastName           = "user_updateFirstAndLastName"
	Predicate_user_updateVerified                   = "user_updateVerified"
	Predicate_user_updateUsername                   = "user_updateUsername"
	Predicate_user_updateProfilePhoto               = "user_updateProfilePhoto"
	Predicate_user_deleteProfilePhotos              = "user_deleteProfilePhotos"
	Predicate_user_getProfilePhotos                 = "user_getProfilePhotos"
	Predicate_user_setBotCommands                   = "user_setBotCommands"
	Predicate_user_isBot                            = "user_isBot"
	Predicate_user_getBotInfo                       = "user_getBotInfo"
	Predicate_user_getFullUser                      = "user_getFullUser"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_peerPeerNotifySettings: {
		0: 1894399913, // 0x70ea3fa9

	},
	Predicate_privacyKeyRules: {
		0: -1810715178, // 0x9412add6

	},
	Predicate_lastSeenData: {
		0: -313287543, // 0xed539c89

	},
	Predicate_contactData: {
		0: 722018346, // 0x2b09202a

	},
	Predicate_botData: {
		0: 23110840, // 0x160a4b8

	},
	Predicate_userData: {
		0: 2138633749, // 0x7f78f615

	},
	Predicate_immutableUser: {
		0: 361114766, // 0x15862c8e

	},
	Predicate_userImportedContacts: {
		0: 1256160192, // 0x4adf7bc0

	},
	Predicate_user_getLastSeens: {
		0: 2090958337, // 0x7ca17e01

	},
	Predicate_user_updateLastSeen: {
		0: 1314677789, // 0x4e5c641d

	},
	Predicate_user_getLastSeen: {
		0: -1860581154, // 0x9119c8de

	},
	Predicate_user_getImmutableUser: {
		0: -47047585, // 0xfd321c5f

	},
	Predicate_user_getMutableUsers: {
		0: 187684863, // 0xb2fd7ff

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
	Predicate_user_createNewPredefinedUser: {
		0: 1464414785, // 0x57493241

	},
	Predicate_user_getPredefinedUser: {
		0: 876047141, // 0x34376b25

	},
	Predicate_user_getAllPredefinedUser: {
		0: -1053474843, // 0xc1353fe5

	},
	Predicate_user_updatePredefinedFirstAndLastName: {
		0: 976922006, // 0x3a3aa596

	},
	Predicate_user_updatePredefinedVerified: {
		0: -1158303159, // 0xbaf5b249

	},
	Predicate_user_updatePredefinedUsername: {
		0: 1269284562, // 0x4ba7bed2

	},
	Predicate_user_updatePredefinedCode: {
		0: 1626771303, // 0x60f68f67

	},
	Predicate_user_predefinedBindRegisteredUserId: {
		0: 68106153, // 0x40f37a9

	},
	Predicate_user_createNewUser: {
		0: 2044729473, // 0x79e01881

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
	Predicate_user_getFullUser: {
		0: -49225414, // 0xfd10e13a

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	1894399913:  Predicate_peerPeerNotifySettings,                // 0x70ea3fa9
	-1810715178: Predicate_privacyKeyRules,                       // 0x9412add6
	-313287543:  Predicate_lastSeenData,                          // 0xed539c89
	722018346:   Predicate_contactData,                           // 0x2b09202a
	23110840:    Predicate_botData,                               // 0x160a4b8
	2138633749:  Predicate_userData,                              // 0x7f78f615
	361114766:   Predicate_immutableUser,                         // 0x15862c8e
	1256160192:  Predicate_userImportedContacts,                  // 0x4adf7bc0
	2090958337:  Predicate_user_getLastSeens,                     // 0x7ca17e01
	1314677789:  Predicate_user_updateLastSeen,                   // 0x4e5c641d
	-1860581154: Predicate_user_getLastSeen,                      // 0x9119c8de
	-47047585:   Predicate_user_getImmutableUser,                 // 0xfd321c5f
	187684863:   Predicate_user_getMutableUsers,                  // 0xb2fd7ff
	-373067804:  Predicate_user_getImmutableUserByPhone,          // 0xe9c36fe4
	-12709005:   Predicate_user_getImmutableUserByToken,          // 0xff3e1373
	-766178484:  Predicate_user_setAccountDaysTTL,                // 0xd2550b4c
	-1299956000: Predicate_user_getAccountDaysTTL,                // 0xb2843ee0
	1085028198:  Predicate_user_getNotifySettings,                // 0x40ac3766
	-907188763:  Predicate_user_setNotifySettings,                // 0xc9ed65e5
	235380084:   Predicate_user_resetNotifySettings,              // 0xe079d74
	1435658357:  Predicate_user_getAllNotifySettings,             // 0x55926875
	2012672274:  Predicate_user_getGlobalPrivacySettings,         // 0x77f6f112
	-1934257490: Predicate_user_setGlobalPrivacySettings,         // 0x8cb592ae
	-1656708172: Predicate_user_getPrivacy,                       // 0x9d40a3b4
	-2007650929: Predicate_user_setPrivacy,                       // 0x8855ad8f
	-982638934:  Predicate_user_checkPrivacy,                     // 0xc56e1eaa
	-891148445:  Predicate_user_addPeerSettings,                  // 0xcae22763
	218296167:   Predicate_user_getPeerSettings,                  // 0xd02ef67
	1586043239:  Predicate_user_deletePeerSettings,               // 0x5e891967
	-8759461:    Predicate_user_changePhone,                      // 0xff7a575b
	1464414785:  Predicate_user_createNewPredefinedUser,          // 0x57493241
	876047141:   Predicate_user_getPredefinedUser,                // 0x34376b25
	-1053474843: Predicate_user_getAllPredefinedUser,             // 0xc1353fe5
	976922006:   Predicate_user_updatePredefinedFirstAndLastName, // 0x3a3aa596
	-1158303159: Predicate_user_updatePredefinedVerified,         // 0xbaf5b249
	1269284562:  Predicate_user_updatePredefinedUsername,         // 0x4ba7bed2
	1626771303:  Predicate_user_updatePredefinedCode,             // 0x60f68f67
	68106153:    Predicate_user_predefinedBindRegisteredUserId,   // 0x40f37a9
	2044729473:  Predicate_user_createNewUser,                    // 0x79e01881
	-2130301264: Predicate_user_blockPeer,                        // 0x81062eb0
	-555280883:  Predicate_user_unBlockPeer,                      // 0xdee7160d
	-1147140722: Predicate_user_blockedByUser,                    // 0xbba0058e
	-1934708257: Predicate_user_isBlockedByUser,                  // 0x8caeb1df
	-1006800656: Predicate_user_checkBlockUserList,               // 0xc3fd70f0
	603964232:   Predicate_user_getBlockedList,                   // 0x23ffc348
	-456010794:  Predicate_user_getContactSignUpNotification,     // 0xe4d1d3d6
	-2053016735: Predicate_user_setContactSignUpNotification,     // 0x85a17361
	-1799115361: Predicate_user_getContentSettings,               // 0x94c3ad9f
	-1654391189: Predicate_user_setContentSettings,               // 0x9d63fe6b
	-972979687:  Predicate_user_deleteContact,                    // 0xc6018219
	-951332511:  Predicate_user_getContactList,                   // 0xc74bd161
	-237135810:  Predicate_user_getContactIdList,                 // 0xf1dd983e
	-613250077:  Predicate_user_getContact,                       // 0xdb728be3
	2042936590:  Predicate_user_addContact,                       // 0x79c4bd0e
	-2102962012: Predicate_user_checkContact,                     // 0x82a758a4
	-1711212654: Predicate_user_importContacts,                   // 0x9a00f792
	302016562:   Predicate_user_getCountryCode,                   // 0x12006832
	-620695161:  Predicate_user_updateAbout,                      // 0xdb00f187
	-882473492:  Predicate_user_updateFirstAndLastName,           // 0xcb6685ec
	617163107:   Predicate_user_updateVerified,                   // 0x24c92963
	-179495311:  Predicate_user_updateUsername,                   // 0xf54d1e71
	997461895:   Predicate_user_updateProfilePhoto,               // 0x3b740f87
	736322062:   Predicate_user_deleteProfilePhotos,              // 0x2be3620e
	-597245626:  Predicate_user_getProfilePhotos,                 // 0xdc66c146
	1966844182:  Predicate_user_setBotCommands,                   // 0x753ba916
	-948779026:  Predicate_user_isBot,                            // 0xc772c7ee
	879114000:   Predicate_user_getBotInfo,                       // 0x34663710
	-49225414:   Predicate_user_getFullUser,                      // 0xfd10e13a

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
