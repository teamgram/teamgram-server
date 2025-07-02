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

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_userImportedContacts              = "userImportedContacts"
	ClazzName_usersDataFound                    = "usersDataFound"
	ClazzName_usersIdFound                      = "usersIdFound"
	ClazzName_peerPeerNotifySettings            = "peerPeerNotifySettings"
	ClazzName_lastSeenData                      = "lastSeenData"
	ClazzName_botInfoData                       = "botInfoData"
	ClazzName_user_getLastSeens                 = "user_getLastSeens"
	ClazzName_user_updateLastSeen               = "user_updateLastSeen"
	ClazzName_user_getLastSeen                  = "user_getLastSeen"
	ClazzName_user_getImmutableUser             = "user_getImmutableUser"
	ClazzName_user_getMutableUsers              = "user_getMutableUsers"
	ClazzName_user_getImmutableUserByPhone      = "user_getImmutableUserByPhone"
	ClazzName_user_getImmutableUserByToken      = "user_getImmutableUserByToken"
	ClazzName_user_setAccountDaysTTL            = "user_setAccountDaysTTL"
	ClazzName_user_getAccountDaysTTL            = "user_getAccountDaysTTL"
	ClazzName_user_getNotifySettings            = "user_getNotifySettings"
	ClazzName_user_getNotifySettingsList        = "user_getNotifySettingsList"
	ClazzName_user_setNotifySettings            = "user_setNotifySettings"
	ClazzName_user_resetNotifySettings          = "user_resetNotifySettings"
	ClazzName_user_getAllNotifySettings         = "user_getAllNotifySettings"
	ClazzName_user_getGlobalPrivacySettings     = "user_getGlobalPrivacySettings"
	ClazzName_user_setGlobalPrivacySettings     = "user_setGlobalPrivacySettings"
	ClazzName_user_getPrivacy                   = "user_getPrivacy"
	ClazzName_user_setPrivacy                   = "user_setPrivacy"
	ClazzName_user_checkPrivacy                 = "user_checkPrivacy"
	ClazzName_user_addPeerSettings              = "user_addPeerSettings"
	ClazzName_user_getPeerSettings              = "user_getPeerSettings"
	ClazzName_user_deletePeerSettings           = "user_deletePeerSettings"
	ClazzName_user_changePhone                  = "user_changePhone"
	ClazzName_user_createNewUser                = "user_createNewUser"
	ClazzName_user_deleteUser                   = "user_deleteUser"
	ClazzName_user_blockPeer                    = "user_blockPeer"
	ClazzName_user_unBlockPeer                  = "user_unBlockPeer"
	ClazzName_user_blockedByUser                = "user_blockedByUser"
	ClazzName_user_isBlockedByUser              = "user_isBlockedByUser"
	ClazzName_user_checkBlockUserList           = "user_checkBlockUserList"
	ClazzName_user_getBlockedList               = "user_getBlockedList"
	ClazzName_user_getContactSignUpNotification = "user_getContactSignUpNotification"
	ClazzName_user_setContactSignUpNotification = "user_setContactSignUpNotification"
	ClazzName_user_getContentSettings           = "user_getContentSettings"
	ClazzName_user_setContentSettings           = "user_setContentSettings"
	ClazzName_user_deleteContact                = "user_deleteContact"
	ClazzName_user_getContactList               = "user_getContactList"
	ClazzName_user_getContactIdList             = "user_getContactIdList"
	ClazzName_user_getContact                   = "user_getContact"
	ClazzName_user_addContact                   = "user_addContact"
	ClazzName_user_checkContact                 = "user_checkContact"
	ClazzName_user_getImportersByPhone          = "user_getImportersByPhone"
	ClazzName_user_deleteImportersByPhone       = "user_deleteImportersByPhone"
	ClazzName_user_importContacts               = "user_importContacts"
	ClazzName_user_getCountryCode               = "user_getCountryCode"
	ClazzName_user_updateAbout                  = "user_updateAbout"
	ClazzName_user_updateFirstAndLastName       = "user_updateFirstAndLastName"
	ClazzName_user_updateVerified               = "user_updateVerified"
	ClazzName_user_updateUsername               = "user_updateUsername"
	ClazzName_user_updateProfilePhoto           = "user_updateProfilePhoto"
	ClazzName_user_deleteProfilePhotos          = "user_deleteProfilePhotos"
	ClazzName_user_getProfilePhotos             = "user_getProfilePhotos"
	ClazzName_user_setBotCommands               = "user_setBotCommands"
	ClazzName_user_isBot                        = "user_isBot"
	ClazzName_user_getBotInfo                   = "user_getBotInfo"
	ClazzName_user_checkBots                    = "user_checkBots"
	ClazzName_user_getFullUser                  = "user_getFullUser"
	ClazzName_user_updateEmojiStatus            = "user_updateEmojiStatus"
	ClazzName_user_getUserDataById              = "user_getUserDataById"
	ClazzName_user_getUserDataListByIdList      = "user_getUserDataListByIdList"
	ClazzName_user_getUserDataByToken           = "user_getUserDataByToken"
	ClazzName_user_search                       = "user_search"
	ClazzName_user_updateBotData                = "user_updateBotData"
	ClazzName_user_getImmutableUserV2           = "user_getImmutableUserV2"
	ClazzName_user_getMutableUsersV2            = "user_getMutableUsersV2"
	ClazzName_user_createNewTestUser            = "user_createNewTestUser"
	ClazzName_user_editCloseFriends             = "user_editCloseFriends"
	ClazzName_user_setStoriesMaxId              = "user_setStoriesMaxId"
	ClazzName_user_setColor                     = "user_setColor"
	ClazzName_user_updateBirthday               = "user_updateBirthday"
	ClazzName_user_getBirthdays                 = "user_getBirthdays"
	ClazzName_user_setStoriesHidden             = "user_setStoriesHidden"
	ClazzName_user_updatePersonalChannel        = "user_updatePersonalChannel"
	ClazzName_user_getUserIdByPhone             = "user_getUserIdByPhone"
	ClazzName_user_setAuthorizationTTL          = "user_setAuthorizationTTL"
	ClazzName_user_getAuthorizationTTL          = "user_getAuthorizationTTL"
	ClazzName_user_updatePremium                = "user_updatePremium"
	ClazzName_user_getBotInfoV2                 = "user_getBotInfoV2"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_userImportedContacts, 0, 0x4adf7bc0)              // 4adf7bc0
	iface.RegisterClazzName(ClazzName_usersDataFound, 0, 0x3fa3dbc7)                    // 3fa3dbc7
	iface.RegisterClazzName(ClazzName_usersIdFound, 0, 0x80c4adfa)                      // 80c4adfa
	iface.RegisterClazzName(ClazzName_peerPeerNotifySettings, 0, 0x70ea3fa9)            // 70ea3fa9
	iface.RegisterClazzName(ClazzName_lastSeenData, 0, 0xb3b1a1df)                      // b3b1a1df
	iface.RegisterClazzName(ClazzName_botInfoData, 0, 0x1835d1c)                        // 1835d1c
	iface.RegisterClazzName(ClazzName_user_getLastSeens, 0, 0x7ca17e01)                 // 7ca17e01
	iface.RegisterClazzName(ClazzName_user_updateLastSeen, 0, 0xfd405a2d)               // fd405a2d
	iface.RegisterClazzName(ClazzName_user_getLastSeen, 0, 0x9119c8de)                  // 9119c8de
	iface.RegisterClazzName(ClazzName_user_getImmutableUser, 0, 0x376a6744)             // 376a6744
	iface.RegisterClazzName(ClazzName_user_getMutableUsers, 0, 0x9d3b23d7)              // 9d3b23d7
	iface.RegisterClazzName(ClazzName_user_getImmutableUserByPhone, 0, 0xe9c36fe4)      // e9c36fe4
	iface.RegisterClazzName(ClazzName_user_getImmutableUserByToken, 0, 0xff3e1373)      // ff3e1373
	iface.RegisterClazzName(ClazzName_user_setAccountDaysTTL, 0, 0xd2550b4c)            // d2550b4c
	iface.RegisterClazzName(ClazzName_user_getAccountDaysTTL, 0, 0xb2843ee0)            // b2843ee0
	iface.RegisterClazzName(ClazzName_user_getNotifySettings, 0, 0x40ac3766)            // 40ac3766
	iface.RegisterClazzName(ClazzName_user_getNotifySettingsList, 0, 0xe465159c)        // e465159c
	iface.RegisterClazzName(ClazzName_user_setNotifySettings, 0, 0xc9ed65e5)            // c9ed65e5
	iface.RegisterClazzName(ClazzName_user_resetNotifySettings, 0, 0xe079d74)           // e079d74
	iface.RegisterClazzName(ClazzName_user_getAllNotifySettings, 0, 0x55926875)         // 55926875
	iface.RegisterClazzName(ClazzName_user_getGlobalPrivacySettings, 0, 0x77f6f112)     // 77f6f112
	iface.RegisterClazzName(ClazzName_user_setGlobalPrivacySettings, 0, 0x8cb592ae)     // 8cb592ae
	iface.RegisterClazzName(ClazzName_user_getPrivacy, 0, 0x9d40a3b4)                   // 9d40a3b4
	iface.RegisterClazzName(ClazzName_user_setPrivacy, 0, 0x8855ad8f)                   // 8855ad8f
	iface.RegisterClazzName(ClazzName_user_checkPrivacy, 0, 0xc56e1eaa)                 // c56e1eaa
	iface.RegisterClazzName(ClazzName_user_addPeerSettings, 0, 0xcae22763)              // cae22763
	iface.RegisterClazzName(ClazzName_user_getPeerSettings, 0, 0xd02ef67)               // d02ef67
	iface.RegisterClazzName(ClazzName_user_deletePeerSettings, 0, 0x5e891967)           // 5e891967
	iface.RegisterClazzName(ClazzName_user_changePhone, 0, 0xff7a575b)                  // ff7a575b
	iface.RegisterClazzName(ClazzName_user_createNewUser, 0, 0x79e01881)                // 79e01881
	iface.RegisterClazzName(ClazzName_user_deleteUser, 0, 0x626dbd10)                   // 626dbd10
	iface.RegisterClazzName(ClazzName_user_blockPeer, 0, 0x81062eb0)                    // 81062eb0
	iface.RegisterClazzName(ClazzName_user_unBlockPeer, 0, 0xdee7160d)                  // dee7160d
	iface.RegisterClazzName(ClazzName_user_blockedByUser, 0, 0xbba0058e)                // bba0058e
	iface.RegisterClazzName(ClazzName_user_isBlockedByUser, 0, 0x8caeb1df)              // 8caeb1df
	iface.RegisterClazzName(ClazzName_user_checkBlockUserList, 0, 0xc3fd70f0)           // c3fd70f0
	iface.RegisterClazzName(ClazzName_user_getBlockedList, 0, 0x23ffc348)               // 23ffc348
	iface.RegisterClazzName(ClazzName_user_getContactSignUpNotification, 0, 0xe4d1d3d6) // e4d1d3d6
	iface.RegisterClazzName(ClazzName_user_setContactSignUpNotification, 0, 0x85a17361) // 85a17361
	iface.RegisterClazzName(ClazzName_user_getContentSettings, 0, 0x94c3ad9f)           // 94c3ad9f
	iface.RegisterClazzName(ClazzName_user_setContentSettings, 0, 0x9d63fe6b)           // 9d63fe6b
	iface.RegisterClazzName(ClazzName_user_deleteContact, 0, 0xc6018219)                // c6018219
	iface.RegisterClazzName(ClazzName_user_getContactList, 0, 0xc74bd161)               // c74bd161
	iface.RegisterClazzName(ClazzName_user_getContactIdList, 0, 0xf1dd983e)             // f1dd983e
	iface.RegisterClazzName(ClazzName_user_getContact, 0, 0xdb728be3)                   // db728be3
	iface.RegisterClazzName(ClazzName_user_addContact, 0, 0x79c4bd0e)                   // 79c4bd0e
	iface.RegisterClazzName(ClazzName_user_checkContact, 0, 0x82a758a4)                 // 82a758a4
	iface.RegisterClazzName(ClazzName_user_getImportersByPhone, 0, 0x47aa8212)          // 47aa8212
	iface.RegisterClazzName(ClazzName_user_deleteImportersByPhone, 0, 0x174ddc54)       // 174ddc54
	iface.RegisterClazzName(ClazzName_user_importContacts, 0, 0x9a00f792)               // 9a00f792
	iface.RegisterClazzName(ClazzName_user_getCountryCode, 0, 0x12006832)               // 12006832
	iface.RegisterClazzName(ClazzName_user_updateAbout, 0, 0xdb00f187)                  // db00f187
	iface.RegisterClazzName(ClazzName_user_updateFirstAndLastName, 0, 0xcb6685ec)       // cb6685ec
	iface.RegisterClazzName(ClazzName_user_updateVerified, 0, 0x24c92963)               // 24c92963
	iface.RegisterClazzName(ClazzName_user_updateUsername, 0, 0xf54d1e71)               // f54d1e71
	iface.RegisterClazzName(ClazzName_user_updateProfilePhoto, 0, 0x3b740f87)           // 3b740f87
	iface.RegisterClazzName(ClazzName_user_deleteProfilePhotos, 0, 0x2be3620e)          // 2be3620e
	iface.RegisterClazzName(ClazzName_user_getProfilePhotos, 0, 0xdc66c146)             // dc66c146
	iface.RegisterClazzName(ClazzName_user_setBotCommands, 0, 0x753ba916)               // 753ba916
	iface.RegisterClazzName(ClazzName_user_isBot, 0, 0xc772c7ee)                        // c772c7ee
	iface.RegisterClazzName(ClazzName_user_getBotInfo, 0, 0x34663710)                   // 34663710
	iface.RegisterClazzName(ClazzName_user_checkBots, 0, 0x736500c1)                    // 736500c1
	iface.RegisterClazzName(ClazzName_user_getFullUser, 0, 0xfd10e13a)                  // fd10e13a
	iface.RegisterClazzName(ClazzName_user_updateEmojiStatus, 0, 0xf8c8bad8)            // f8c8bad8
	iface.RegisterClazzName(ClazzName_user_getUserDataById, 0, 0x3bb7103)               // 3bb7103
	iface.RegisterClazzName(ClazzName_user_getUserDataListByIdList, 0, 0x8191eff9)      // 8191eff9
	iface.RegisterClazzName(ClazzName_user_getUserDataByToken, 0, 0x3f09659e)           // 3f09659e
	iface.RegisterClazzName(ClazzName_user_search, 0, 0x7035b6cd)                       // 7035b6cd
	iface.RegisterClazzName(ClazzName_user_updateBotData, 0, 0x60f35d28)                // 60f35d28
	iface.RegisterClazzName(ClazzName_user_getImmutableUserV2, 0, 0x300aba4c)           // 300aba4c
	iface.RegisterClazzName(ClazzName_user_getMutableUsersV2, 0, 0x94f98b28)            // 94f98b28
	iface.RegisterClazzName(ClazzName_user_createNewTestUser, 0, 0x4c6eccab)            // 4c6eccab
	iface.RegisterClazzName(ClazzName_user_editCloseFriends, 0, 0x86247b05)             // 86247b05
	iface.RegisterClazzName(ClazzName_user_setStoriesMaxId, 0, 0x52f5b670)              // 52f5b670
	iface.RegisterClazzName(ClazzName_user_setColor, 0, 0x22fa0d77)                     // 22fa0d77
	iface.RegisterClazzName(ClazzName_user_updateBirthday, 0, 0x587aab92)               // 587aab92
	iface.RegisterClazzName(ClazzName_user_getBirthdays, 0, 0xfe8ebfa6)                 // fe8ebfa6
	iface.RegisterClazzName(ClazzName_user_setStoriesHidden, 0, 0xf7c61858)             // f7c61858
	iface.RegisterClazzName(ClazzName_user_updatePersonalChannel, 0, 0xc7f7bed0)        // c7f7bed0
	iface.RegisterClazzName(ClazzName_user_getUserIdByPhone, 0, 0xfbab83c2)             // fbab83c2
	iface.RegisterClazzName(ClazzName_user_setAuthorizationTTL, 0, 0xd621f3f0)          // d621f3f0
	iface.RegisterClazzName(ClazzName_user_getAuthorizationTTL, 0, 0xde6e493c)          // de6e493c
	iface.RegisterClazzName(ClazzName_user_updatePremium, 0, 0xba08dc99)                // ba08dc99
	iface.RegisterClazzName(ClazzName_user_getBotInfoV2, 0, 0xd3fc9ca5)                 // d3fc9ca5

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_userImportedContacts, 0x4adf7bc0)              // 4adf7bc0
	iface.RegisterClazzIDName(ClazzName_usersDataFound, 0x3fa3dbc7)                    // 3fa3dbc7
	iface.RegisterClazzIDName(ClazzName_usersIdFound, 0x80c4adfa)                      // 80c4adfa
	iface.RegisterClazzIDName(ClazzName_peerPeerNotifySettings, 0x70ea3fa9)            // 70ea3fa9
	iface.RegisterClazzIDName(ClazzName_lastSeenData, 0xb3b1a1df)                      // b3b1a1df
	iface.RegisterClazzIDName(ClazzName_botInfoData, 0x1835d1c)                        // 1835d1c
	iface.RegisterClazzIDName(ClazzName_user_getLastSeens, 0x7ca17e01)                 // 7ca17e01
	iface.RegisterClazzIDName(ClazzName_user_updateLastSeen, 0xfd405a2d)               // fd405a2d
	iface.RegisterClazzIDName(ClazzName_user_getLastSeen, 0x9119c8de)                  // 9119c8de
	iface.RegisterClazzIDName(ClazzName_user_getImmutableUser, 0x376a6744)             // 376a6744
	iface.RegisterClazzIDName(ClazzName_user_getMutableUsers, 0x9d3b23d7)              // 9d3b23d7
	iface.RegisterClazzIDName(ClazzName_user_getImmutableUserByPhone, 0xe9c36fe4)      // e9c36fe4
	iface.RegisterClazzIDName(ClazzName_user_getImmutableUserByToken, 0xff3e1373)      // ff3e1373
	iface.RegisterClazzIDName(ClazzName_user_setAccountDaysTTL, 0xd2550b4c)            // d2550b4c
	iface.RegisterClazzIDName(ClazzName_user_getAccountDaysTTL, 0xb2843ee0)            // b2843ee0
	iface.RegisterClazzIDName(ClazzName_user_getNotifySettings, 0x40ac3766)            // 40ac3766
	iface.RegisterClazzIDName(ClazzName_user_getNotifySettingsList, 0xe465159c)        // e465159c
	iface.RegisterClazzIDName(ClazzName_user_setNotifySettings, 0xc9ed65e5)            // c9ed65e5
	iface.RegisterClazzIDName(ClazzName_user_resetNotifySettings, 0xe079d74)           // e079d74
	iface.RegisterClazzIDName(ClazzName_user_getAllNotifySettings, 0x55926875)         // 55926875
	iface.RegisterClazzIDName(ClazzName_user_getGlobalPrivacySettings, 0x77f6f112)     // 77f6f112
	iface.RegisterClazzIDName(ClazzName_user_setGlobalPrivacySettings, 0x8cb592ae)     // 8cb592ae
	iface.RegisterClazzIDName(ClazzName_user_getPrivacy, 0x9d40a3b4)                   // 9d40a3b4
	iface.RegisterClazzIDName(ClazzName_user_setPrivacy, 0x8855ad8f)                   // 8855ad8f
	iface.RegisterClazzIDName(ClazzName_user_checkPrivacy, 0xc56e1eaa)                 // c56e1eaa
	iface.RegisterClazzIDName(ClazzName_user_addPeerSettings, 0xcae22763)              // cae22763
	iface.RegisterClazzIDName(ClazzName_user_getPeerSettings, 0xd02ef67)               // d02ef67
	iface.RegisterClazzIDName(ClazzName_user_deletePeerSettings, 0x5e891967)           // 5e891967
	iface.RegisterClazzIDName(ClazzName_user_changePhone, 0xff7a575b)                  // ff7a575b
	iface.RegisterClazzIDName(ClazzName_user_createNewUser, 0x79e01881)                // 79e01881
	iface.RegisterClazzIDName(ClazzName_user_deleteUser, 0x626dbd10)                   // 626dbd10
	iface.RegisterClazzIDName(ClazzName_user_blockPeer, 0x81062eb0)                    // 81062eb0
	iface.RegisterClazzIDName(ClazzName_user_unBlockPeer, 0xdee7160d)                  // dee7160d
	iface.RegisterClazzIDName(ClazzName_user_blockedByUser, 0xbba0058e)                // bba0058e
	iface.RegisterClazzIDName(ClazzName_user_isBlockedByUser, 0x8caeb1df)              // 8caeb1df
	iface.RegisterClazzIDName(ClazzName_user_checkBlockUserList, 0xc3fd70f0)           // c3fd70f0
	iface.RegisterClazzIDName(ClazzName_user_getBlockedList, 0x23ffc348)               // 23ffc348
	iface.RegisterClazzIDName(ClazzName_user_getContactSignUpNotification, 0xe4d1d3d6) // e4d1d3d6
	iface.RegisterClazzIDName(ClazzName_user_setContactSignUpNotification, 0x85a17361) // 85a17361
	iface.RegisterClazzIDName(ClazzName_user_getContentSettings, 0x94c3ad9f)           // 94c3ad9f
	iface.RegisterClazzIDName(ClazzName_user_setContentSettings, 0x9d63fe6b)           // 9d63fe6b
	iface.RegisterClazzIDName(ClazzName_user_deleteContact, 0xc6018219)                // c6018219
	iface.RegisterClazzIDName(ClazzName_user_getContactList, 0xc74bd161)               // c74bd161
	iface.RegisterClazzIDName(ClazzName_user_getContactIdList, 0xf1dd983e)             // f1dd983e
	iface.RegisterClazzIDName(ClazzName_user_getContact, 0xdb728be3)                   // db728be3
	iface.RegisterClazzIDName(ClazzName_user_addContact, 0x79c4bd0e)                   // 79c4bd0e
	iface.RegisterClazzIDName(ClazzName_user_checkContact, 0x82a758a4)                 // 82a758a4
	iface.RegisterClazzIDName(ClazzName_user_getImportersByPhone, 0x47aa8212)          // 47aa8212
	iface.RegisterClazzIDName(ClazzName_user_deleteImportersByPhone, 0x174ddc54)       // 174ddc54
	iface.RegisterClazzIDName(ClazzName_user_importContacts, 0x9a00f792)               // 9a00f792
	iface.RegisterClazzIDName(ClazzName_user_getCountryCode, 0x12006832)               // 12006832
	iface.RegisterClazzIDName(ClazzName_user_updateAbout, 0xdb00f187)                  // db00f187
	iface.RegisterClazzIDName(ClazzName_user_updateFirstAndLastName, 0xcb6685ec)       // cb6685ec
	iface.RegisterClazzIDName(ClazzName_user_updateVerified, 0x24c92963)               // 24c92963
	iface.RegisterClazzIDName(ClazzName_user_updateUsername, 0xf54d1e71)               // f54d1e71
	iface.RegisterClazzIDName(ClazzName_user_updateProfilePhoto, 0x3b740f87)           // 3b740f87
	iface.RegisterClazzIDName(ClazzName_user_deleteProfilePhotos, 0x2be3620e)          // 2be3620e
	iface.RegisterClazzIDName(ClazzName_user_getProfilePhotos, 0xdc66c146)             // dc66c146
	iface.RegisterClazzIDName(ClazzName_user_setBotCommands, 0x753ba916)               // 753ba916
	iface.RegisterClazzIDName(ClazzName_user_isBot, 0xc772c7ee)                        // c772c7ee
	iface.RegisterClazzIDName(ClazzName_user_getBotInfo, 0x34663710)                   // 34663710
	iface.RegisterClazzIDName(ClazzName_user_checkBots, 0x736500c1)                    // 736500c1
	iface.RegisterClazzIDName(ClazzName_user_getFullUser, 0xfd10e13a)                  // fd10e13a
	iface.RegisterClazzIDName(ClazzName_user_updateEmojiStatus, 0xf8c8bad8)            // f8c8bad8
	iface.RegisterClazzIDName(ClazzName_user_getUserDataById, 0x3bb7103)               // 3bb7103
	iface.RegisterClazzIDName(ClazzName_user_getUserDataListByIdList, 0x8191eff9)      // 8191eff9
	iface.RegisterClazzIDName(ClazzName_user_getUserDataByToken, 0x3f09659e)           // 3f09659e
	iface.RegisterClazzIDName(ClazzName_user_search, 0x7035b6cd)                       // 7035b6cd
	iface.RegisterClazzIDName(ClazzName_user_updateBotData, 0x60f35d28)                // 60f35d28
	iface.RegisterClazzIDName(ClazzName_user_getImmutableUserV2, 0x300aba4c)           // 300aba4c
	iface.RegisterClazzIDName(ClazzName_user_getMutableUsersV2, 0x94f98b28)            // 94f98b28
	iface.RegisterClazzIDName(ClazzName_user_createNewTestUser, 0x4c6eccab)            // 4c6eccab
	iface.RegisterClazzIDName(ClazzName_user_editCloseFriends, 0x86247b05)             // 86247b05
	iface.RegisterClazzIDName(ClazzName_user_setStoriesMaxId, 0x52f5b670)              // 52f5b670
	iface.RegisterClazzIDName(ClazzName_user_setColor, 0x22fa0d77)                     // 22fa0d77
	iface.RegisterClazzIDName(ClazzName_user_updateBirthday, 0x587aab92)               // 587aab92
	iface.RegisterClazzIDName(ClazzName_user_getBirthdays, 0xfe8ebfa6)                 // fe8ebfa6
	iface.RegisterClazzIDName(ClazzName_user_setStoriesHidden, 0xf7c61858)             // f7c61858
	iface.RegisterClazzIDName(ClazzName_user_updatePersonalChannel, 0xc7f7bed0)        // c7f7bed0
	iface.RegisterClazzIDName(ClazzName_user_getUserIdByPhone, 0xfbab83c2)             // fbab83c2
	iface.RegisterClazzIDName(ClazzName_user_setAuthorizationTTL, 0xd621f3f0)          // d621f3f0
	iface.RegisterClazzIDName(ClazzName_user_getAuthorizationTTL, 0xde6e493c)          // de6e493c
	iface.RegisterClazzIDName(ClazzName_user_updatePremium, 0xba08dc99)                // ba08dc99
	iface.RegisterClazzIDName(ClazzName_user_getBotInfoV2, 0xd3fc9ca5)                 // d3fc9ca5
}
