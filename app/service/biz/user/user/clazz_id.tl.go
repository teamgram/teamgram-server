/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package user

const (
	ClazzID_userImportedContacts              = 0x4adf7bc0 // 4adf7bc0
	ClazzID_usersDataFound                    = 0x3fa3dbc7 // 3fa3dbc7
	ClazzID_usersIdFound                      = 0x80c4adfa // 80c4adfa
	ClazzID_peerPeerNotifySettings            = 0x70ea3fa9 // 70ea3fa9
	ClazzID_lastSeenData                      = 0xb3b1a1df // b3b1a1df
	ClazzID_botInfoData                       = 0x1835d1c  // 1835d1c
	ClazzID_user_getLastSeens                 = 0x7ca17e01 // 7ca17e01
	ClazzID_user_updateLastSeen               = 0xfd405a2d // fd405a2d
	ClazzID_user_getLastSeen                  = 0x9119c8de // 9119c8de
	ClazzID_user_getImmutableUser             = 0x376a6744 // 376a6744
	ClazzID_user_getMutableUsers              = 0x9d3b23d7 // 9d3b23d7
	ClazzID_user_getImmutableUserByPhone      = 0xe9c36fe4 // e9c36fe4
	ClazzID_user_getImmutableUserByToken      = 0xff3e1373 // ff3e1373
	ClazzID_user_setAccountDaysTTL            = 0xd2550b4c // d2550b4c
	ClazzID_user_getAccountDaysTTL            = 0xb2843ee0 // b2843ee0
	ClazzID_user_getNotifySettings            = 0x40ac3766 // 40ac3766
	ClazzID_user_getNotifySettingsList        = 0xe465159c // e465159c
	ClazzID_user_setNotifySettings            = 0xc9ed65e5 // c9ed65e5
	ClazzID_user_resetNotifySettings          = 0xe079d74  // e079d74
	ClazzID_user_getAllNotifySettings         = 0x55926875 // 55926875
	ClazzID_user_getGlobalPrivacySettings     = 0x77f6f112 // 77f6f112
	ClazzID_user_setGlobalPrivacySettings     = 0x8cb592ae // 8cb592ae
	ClazzID_user_getPrivacy                   = 0x9d40a3b4 // 9d40a3b4
	ClazzID_user_setPrivacy                   = 0x8855ad8f // 8855ad8f
	ClazzID_user_checkPrivacy                 = 0xc56e1eaa // c56e1eaa
	ClazzID_user_addPeerSettings              = 0xcae22763 // cae22763
	ClazzID_user_getPeerSettings              = 0xd02ef67  // d02ef67
	ClazzID_user_deletePeerSettings           = 0x5e891967 // 5e891967
	ClazzID_user_changePhone                  = 0xff7a575b // ff7a575b
	ClazzID_user_createNewUser                = 0x79e01881 // 79e01881
	ClazzID_user_deleteUser                   = 0x626dbd10 // 626dbd10
	ClazzID_user_blockPeer                    = 0x81062eb0 // 81062eb0
	ClazzID_user_unBlockPeer                  = 0xdee7160d // dee7160d
	ClazzID_user_blockedByUser                = 0xbba0058e // bba0058e
	ClazzID_user_isBlockedByUser              = 0x8caeb1df // 8caeb1df
	ClazzID_user_checkBlockUserList           = 0xc3fd70f0 // c3fd70f0
	ClazzID_user_getBlockedList               = 0x23ffc348 // 23ffc348
	ClazzID_user_getContactSignUpNotification = 0xe4d1d3d6 // e4d1d3d6
	ClazzID_user_setContactSignUpNotification = 0x85a17361 // 85a17361
	ClazzID_user_getContentSettings           = 0x94c3ad9f // 94c3ad9f
	ClazzID_user_setContentSettings           = 0x9d63fe6b // 9d63fe6b
	ClazzID_user_deleteContact                = 0xc6018219 // c6018219
	ClazzID_user_getContactList               = 0xc74bd161 // c74bd161
	ClazzID_user_getContactIdList             = 0xf1dd983e // f1dd983e
	ClazzID_user_getContact                   = 0xdb728be3 // db728be3
	ClazzID_user_addContact                   = 0x79c4bd0e // 79c4bd0e
	ClazzID_user_checkContact                 = 0x82a758a4 // 82a758a4
	ClazzID_user_getImportersByPhone          = 0x47aa8212 // 47aa8212
	ClazzID_user_deleteImportersByPhone       = 0x174ddc54 // 174ddc54
	ClazzID_user_importContacts               = 0x9a00f792 // 9a00f792
	ClazzID_user_getCountryCode               = 0x12006832 // 12006832
	ClazzID_user_updateAbout                  = 0xdb00f187 // db00f187
	ClazzID_user_updateFirstAndLastName       = 0xcb6685ec // cb6685ec
	ClazzID_user_updateVerified               = 0x24c92963 // 24c92963
	ClazzID_user_updateUsername               = 0xf54d1e71 // f54d1e71
	ClazzID_user_updateProfilePhoto           = 0x3b740f87 // 3b740f87
	ClazzID_user_deleteProfilePhotos          = 0x2be3620e // 2be3620e
	ClazzID_user_getProfilePhotos             = 0xdc66c146 // dc66c146
	ClazzID_user_setBotCommands               = 0x753ba916 // 753ba916
	ClazzID_user_isBot                        = 0xc772c7ee // c772c7ee
	ClazzID_user_getBotInfo                   = 0x34663710 // 34663710
	ClazzID_user_checkBots                    = 0x736500c1 // 736500c1
	ClazzID_user_getFullUser                  = 0xfd10e13a // fd10e13a
	ClazzID_user_updateEmojiStatus            = 0xf8c8bad8 // f8c8bad8
	ClazzID_user_getUserDataById              = 0x3bb7103  // 3bb7103
	ClazzID_user_getUserDataListByIdList      = 0x8191eff9 // 8191eff9
	ClazzID_user_getUserDataByToken           = 0x3f09659e // 3f09659e
	ClazzID_user_search                       = 0x7035b6cd // 7035b6cd
	ClazzID_user_updateBotData                = 0x60f35d28 // 60f35d28
	ClazzID_user_getImmutableUserV2           = 0x300aba4c // 300aba4c
	ClazzID_user_getMutableUsersV2            = 0x94f98b28 // 94f98b28
	ClazzID_user_createNewTestUser            = 0x4c6eccab // 4c6eccab
	ClazzID_user_editCloseFriends             = 0x86247b05 // 86247b05
	ClazzID_user_setStoriesMaxId              = 0x52f5b670 // 52f5b670
	ClazzID_user_setColor                     = 0x22fa0d77 // 22fa0d77
	ClazzID_user_updateBirthday               = 0x587aab92 // 587aab92
	ClazzID_user_getBirthdays                 = 0xfe8ebfa6 // fe8ebfa6
	ClazzID_user_setStoriesHidden             = 0xf7c61858 // f7c61858
	ClazzID_user_updatePersonalChannel        = 0xc7f7bed0 // c7f7bed0
	ClazzID_user_getUserIdByPhone             = 0xfbab83c2 // fbab83c2
	ClazzID_user_setAuthorizationTTL          = 0xd621f3f0 // d621f3f0
	ClazzID_user_getAuthorizationTTL          = 0xde6e493c // de6e493c
	ClazzID_user_updatePremium                = 0xba08dc99 // ba08dc99
	ClazzID_user_getBotInfoV2                 = 0xd3fc9ca5 // d3fc9ca5
)
