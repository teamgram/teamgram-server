/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package user

const (
	ClazzID_userImportedContacts               = 0x4adf7bc0 // 4adf7bc0
	ClazzID_usersDataFound                     = 0x3fa3dbc7 // 3fa3dbc7
	ClazzID_usersIdFound                       = 0x80c4adfa // 80c4adfa
	ClazzID_peerPeerNotifySettings             = 0x70ea3fa9 // 70ea3fa9
	ClazzID_lastSeenData                       = 0xb3b1a1df // b3b1a1df
	ClazzID_botInfoData                        = 0x1835d1c  // 1835d1c
	ClazzID_usernameNotExisted                 = 0xcb3cfb6d // cb3cfb6d
	ClazzID_usernameExisted                    = 0xace7f4cd // ace7f4cd
	ClazzID_usernameExistedNotMe               = 0xd01f47b1 // d01f47b1
	ClazzID_usernameExistedIsMe                = 0x874e7771 // 874e7771
	ClazzID_usernameData                       = 0xaa4000bf // aa4000bf
	ClazzID_user_getLastSeens                  = 0x7ca17e01 // 7ca17e01
	ClazzID_user_updateLastSeen                = 0xfd405a2d // fd405a2d
	ClazzID_user_getLastSeen                   = 0x9119c8de // 9119c8de
	ClazzID_user_getImmutableUser              = 0x376a6744 // 376a6744
	ClazzID_user_getMutableUsers               = 0x9d3b23d7 // 9d3b23d7
	ClazzID_user_getImmutableUserByPhone       = 0xe9c36fe4 // e9c36fe4
	ClazzID_user_getImmutableUserByToken       = 0xff3e1373 // ff3e1373
	ClazzID_user_setAccountDaysTTL             = 0xd2550b4c // d2550b4c
	ClazzID_user_getAccountDaysTTL             = 0xb2843ee0 // b2843ee0
	ClazzID_user_getNotifySettings             = 0x40ac3766 // 40ac3766
	ClazzID_user_getNotifySettingsList         = 0xe465159c // e465159c
	ClazzID_user_setNotifySettings             = 0xc9ed65e5 // c9ed65e5
	ClazzID_user_resetNotifySettings           = 0xe079d74  // e079d74
	ClazzID_user_getAllNotifySettings          = 0x55926875 // 55926875
	ClazzID_user_getGlobalPrivacySettings      = 0x77f6f112 // 77f6f112
	ClazzID_user_setGlobalPrivacySettings      = 0x8cb592ae // 8cb592ae
	ClazzID_user_getPrivacy                    = 0x9d40a3b4 // 9d40a3b4
	ClazzID_user_setPrivacy                    = 0x8855ad8f // 8855ad8f
	ClazzID_user_checkPrivacy                  = 0xc56e1eaa // c56e1eaa
	ClazzID_user_addPeerSettings               = 0xcae22763 // cae22763
	ClazzID_user_getPeerSettings               = 0xd02ef67  // d02ef67
	ClazzID_user_deletePeerSettings            = 0x5e891967 // 5e891967
	ClazzID_user_changePhone                   = 0xff7a575b // ff7a575b
	ClazzID_user_createNewUser                 = 0x79e01881 // 79e01881
	ClazzID_user_deleteUser                    = 0x626dbd10 // 626dbd10
	ClazzID_user_blockPeer                     = 0x81062eb0 // 81062eb0
	ClazzID_user_unBlockPeer                   = 0xdee7160d // dee7160d
	ClazzID_user_blockedByUser                 = 0xbba0058e // bba0058e
	ClazzID_user_isBlockedByUser               = 0x8caeb1df // 8caeb1df
	ClazzID_user_checkBlockUserList            = 0xc3fd70f0 // c3fd70f0
	ClazzID_user_getBlockedList                = 0x23ffc348 // 23ffc348
	ClazzID_user_getContactSignUpNotification  = 0xe4d1d3d6 // e4d1d3d6
	ClazzID_user_setContactSignUpNotification  = 0x85a17361 // 85a17361
	ClazzID_user_getContentSettings            = 0x94c3ad9f // 94c3ad9f
	ClazzID_user_setContentSettings            = 0x9d63fe6b // 9d63fe6b
	ClazzID_user_deleteContact                 = 0xc6018219 // c6018219
	ClazzID_user_getContactList                = 0xc74bd161 // c74bd161
	ClazzID_user_getContactIdList              = 0xf1dd983e // f1dd983e
	ClazzID_user_getContact                    = 0xdb728be3 // db728be3
	ClazzID_user_addContact                    = 0x79c4bd0e // 79c4bd0e
	ClazzID_user_checkContact                  = 0x82a758a4 // 82a758a4
	ClazzID_user_getImportersByPhone           = 0x47aa8212 // 47aa8212
	ClazzID_user_deleteImportersByPhone        = 0x174ddc54 // 174ddc54
	ClazzID_user_importContacts                = 0x9a00f792 // 9a00f792
	ClazzID_user_getCountryCode                = 0x12006832 // 12006832
	ClazzID_user_updateAbout                   = 0xdb00f187 // db00f187
	ClazzID_user_updateFirstAndLastName        = 0xcb6685ec // cb6685ec
	ClazzID_user_updateVerified                = 0x24c92963 // 24c92963
	ClazzID_user_updateUsername                = 0xf54d1e71 // f54d1e71
	ClazzID_user_updateProfilePhoto            = 0x3b740f87 // 3b740f87
	ClazzID_user_deleteProfilePhotos           = 0x2be3620e // 2be3620e
	ClazzID_user_getProfilePhotos              = 0xdc66c146 // dc66c146
	ClazzID_user_setBotCommands                = 0x753ba916 // 753ba916
	ClazzID_user_isBot                         = 0xc772c7ee // c772c7ee
	ClazzID_user_getBotInfo                    = 0x34663710 // 34663710
	ClazzID_user_checkBots                     = 0x736500c1 // 736500c1
	ClazzID_user_getFullUser                   = 0xfd10e13a // fd10e13a
	ClazzID_user_updateEmojiStatus             = 0xf8c8bad8 // f8c8bad8
	ClazzID_user_getUserDataById               = 0x3bb7103  // 3bb7103
	ClazzID_user_getUserDataListByIdList       = 0x8191eff9 // 8191eff9
	ClazzID_user_getUserDataByToken            = 0x3f09659e // 3f09659e
	ClazzID_user_search                        = 0x7035b6cd // 7035b6cd
	ClazzID_user_updateBotData                 = 0x60f35d28 // 60f35d28
	ClazzID_user_getImmutableUserV2            = 0x300aba4c // 300aba4c
	ClazzID_user_getMutableUsersV2             = 0x94f98b28 // 94f98b28
	ClazzID_user_createNewTestUser             = 0x4c6eccab // 4c6eccab
	ClazzID_user_editCloseFriends              = 0x86247b05 // 86247b05
	ClazzID_user_setStoriesMaxId               = 0x52f5b670 // 52f5b670
	ClazzID_user_setColor                      = 0x22fa0d77 // 22fa0d77
	ClazzID_user_updateBirthday                = 0x587aab92 // 587aab92
	ClazzID_user_getBirthdays                  = 0xfe8ebfa6 // fe8ebfa6
	ClazzID_user_setStoriesHidden              = 0xf7c61858 // f7c61858
	ClazzID_user_updatePersonalChannel         = 0xc7f7bed0 // c7f7bed0
	ClazzID_user_getUserIdByPhone              = 0xfbab83c2 // fbab83c2
	ClazzID_user_setAuthorizationTTL           = 0xd621f3f0 // d621f3f0
	ClazzID_user_getAuthorizationTTL           = 0xde6e493c // de6e493c
	ClazzID_user_updatePremium                 = 0xba08dc99 // ba08dc99
	ClazzID_user_getBotInfoV2                  = 0xd3fc9ca5 // d3fc9ca5
	ClazzID_user_saveMusic                     = 0xda28349  // da28349
	ClazzID_user_getSavedMusicIdList           = 0x5b4ac25f // 5b4ac25f
	ClazzID_user_setMainProfileTab             = 0x9d48a89c // 9d48a89c
	ClazzID_user_setDefaultHistoryTTL          = 0x8f09517f // 8f09517f
	ClazzID_user_getDefaultHistoryTTL          = 0x4d4c2fe0 // 4d4c2fe0
	ClazzID_user_getAccountUsername            = 0xff8b61cf // ff8b61cf
	ClazzID_user_checkAccountUsername          = 0xefe05198 // efe05198
	ClazzID_user_getChannelUsername            = 0x910ac7b1 // 910ac7b1
	ClazzID_user_checkChannelUsername          = 0xfecbacb6 // fecbacb6
	ClazzID_user_updateUsernameByPeer          = 0xe3a0e9e2 // e3a0e9e2
	ClazzID_user_checkUsername                 = 0x3475e700 // 3475e700
	ClazzID_user_updateUsernameByUsername      = 0x13841a86 // 13841a86
	ClazzID_user_deleteUsername                = 0x85e677a4 // 85e677a4
	ClazzID_user_resolveUsername               = 0x4527d121 // 4527d121
	ClazzID_user_getListByUsernameList         = 0x6e606b62 // 6e606b62
	ClazzID_user_deleteUsernameByPeer          = 0x7cafbc1  // 7cafbc1
	ClazzID_user_searchUsername                = 0x266eb885 // 266eb885
	ClazzID_user_toggleUsername                = 0xdd3b5a14 // dd3b5a14
	ClazzID_user_reorderUsernames              = 0x3a61bdc0 // 3a61bdc0
	ClazzID_user_deactivateAllChannelUsernames = 0x9a5fe53c // 9a5fe53c
)
