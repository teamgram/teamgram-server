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
	CRC32_UNKNOWN                           TLConstructor = 0
	CRC32_userImportedContacts              TLConstructor = 1256160192  // 0x4adf7bc0
	CRC32_usersDataFound                    TLConstructor = 1067703239  // 0x3fa3dbc7
	CRC32_usersIdFound                      TLConstructor = -2134594054 // 0x80c4adfa
	CRC32_peerPeerNotifySettings            TLConstructor = 1894399913  // 0x70ea3fa9
	CRC32_lastSeenData                      TLConstructor = -1280204321 // 0xb3b1a1df
	CRC32_user_getLastSeens                 TLConstructor = 2090958337  // 0x7ca17e01
	CRC32_user_updateLastSeen               TLConstructor = -46114259   // 0xfd405a2d
	CRC32_user_getLastSeen                  TLConstructor = -1860581154 // 0x9119c8de
	CRC32_user_getImmutableUser             TLConstructor = 929720132   // 0x376a6744
	CRC32_user_getMutableUsers              TLConstructor = -1657068585 // 0x9d3b23d7
	CRC32_user_getImmutableUserByPhone      TLConstructor = -373067804  // 0xe9c36fe4
	CRC32_user_getImmutableUserByToken      TLConstructor = -12709005   // 0xff3e1373
	CRC32_user_setAccountDaysTTL            TLConstructor = -766178484  // 0xd2550b4c
	CRC32_user_getAccountDaysTTL            TLConstructor = -1299956000 // 0xb2843ee0
	CRC32_user_getNotifySettings            TLConstructor = 1085028198  // 0x40ac3766
	CRC32_user_getNotifySettingsList        TLConstructor = -463137380  // 0xe465159c
	CRC32_user_setNotifySettings            TLConstructor = -907188763  // 0xc9ed65e5
	CRC32_user_resetNotifySettings          TLConstructor = 235380084   // 0xe079d74
	CRC32_user_getAllNotifySettings         TLConstructor = 1435658357  // 0x55926875
	CRC32_user_getGlobalPrivacySettings     TLConstructor = 2012672274  // 0x77f6f112
	CRC32_user_setGlobalPrivacySettings     TLConstructor = -1934257490 // 0x8cb592ae
	CRC32_user_getPrivacy                   TLConstructor = -1656708172 // 0x9d40a3b4
	CRC32_user_setPrivacy                   TLConstructor = -2007650929 // 0x8855ad8f
	CRC32_user_checkPrivacy                 TLConstructor = -982638934  // 0xc56e1eaa
	CRC32_user_addPeerSettings              TLConstructor = -891148445  // 0xcae22763
	CRC32_user_getPeerSettings              TLConstructor = 218296167   // 0xd02ef67
	CRC32_user_deletePeerSettings           TLConstructor = 1586043239  // 0x5e891967
	CRC32_user_changePhone                  TLConstructor = -8759461    // 0xff7a575b
	CRC32_user_createNewUser                TLConstructor = 2044729473  // 0x79e01881
	CRC32_user_deleteUser                   TLConstructor = 2132777160  // 0x7f1f98c8
	CRC32_user_blockPeer                    TLConstructor = -2130301264 // 0x81062eb0
	CRC32_user_unBlockPeer                  TLConstructor = -555280883  // 0xdee7160d
	CRC32_user_blockedByUser                TLConstructor = -1147140722 // 0xbba0058e
	CRC32_user_isBlockedByUser              TLConstructor = -1934708257 // 0x8caeb1df
	CRC32_user_checkBlockUserList           TLConstructor = -1006800656 // 0xc3fd70f0
	CRC32_user_getBlockedList               TLConstructor = 603964232   // 0x23ffc348
	CRC32_user_getContactSignUpNotification TLConstructor = -456010794  // 0xe4d1d3d6
	CRC32_user_setContactSignUpNotification TLConstructor = -2053016735 // 0x85a17361
	CRC32_user_getContentSettings           TLConstructor = -1799115361 // 0x94c3ad9f
	CRC32_user_setContentSettings           TLConstructor = -1654391189 // 0x9d63fe6b
	CRC32_user_deleteContact                TLConstructor = -972979687  // 0xc6018219
	CRC32_user_getContactList               TLConstructor = -951332511  // 0xc74bd161
	CRC32_user_getContactIdList             TLConstructor = -237135810  // 0xf1dd983e
	CRC32_user_getContact                   TLConstructor = -613250077  // 0xdb728be3
	CRC32_user_addContact                   TLConstructor = 2042936590  // 0x79c4bd0e
	CRC32_user_checkContact                 TLConstructor = -2102962012 // 0x82a758a4
	CRC32_user_getImportersByPhone          TLConstructor = 1202356754  // 0x47aa8212
	CRC32_user_deleteImportersByPhone       TLConstructor = 390978644   // 0x174ddc54
	CRC32_user_importContacts               TLConstructor = -1711212654 // 0x9a00f792
	CRC32_user_getCountryCode               TLConstructor = 302016562   // 0x12006832
	CRC32_user_updateAbout                  TLConstructor = -620695161  // 0xdb00f187
	CRC32_user_updateFirstAndLastName       TLConstructor = -882473492  // 0xcb6685ec
	CRC32_user_updateVerified               TLConstructor = 617163107   // 0x24c92963
	CRC32_user_updateUsername               TLConstructor = -179495311  // 0xf54d1e71
	CRC32_user_updateProfilePhoto           TLConstructor = 997461895   // 0x3b740f87
	CRC32_user_deleteProfilePhotos          TLConstructor = 736322062   // 0x2be3620e
	CRC32_user_getProfilePhotos             TLConstructor = -597245626  // 0xdc66c146
	CRC32_user_setBotCommands               TLConstructor = 1966844182  // 0x753ba916
	CRC32_user_isBot                        TLConstructor = -948779026  // 0xc772c7ee
	CRC32_user_getBotInfo                   TLConstructor = 879114000   // 0x34663710
	CRC32_user_checkBots                    TLConstructor = 1935999169  // 0x736500c1
	CRC32_user_getFullUser                  TLConstructor = -49225414   // 0xfd10e13a
	CRC32_user_updateEmojiStatus            TLConstructor = -121062696  // 0xf8c8bad8
	CRC32_user_getUserDataById              TLConstructor = 62615811    // 0x3bb7103
	CRC32_user_getUserDataListByIdList      TLConstructor = -2121142279 // 0x8191eff9
	CRC32_user_getUserDataByToken           TLConstructor = 1057580446  // 0x3f09659e
	CRC32_user_search                       TLConstructor = 1882568397  // 0x7035b6cd
	CRC32_user_updateBotData                TLConstructor = -1174586898 // 0xb9fd39ee
	CRC32_user_getImmutableUserV2           TLConstructor = 806009420   // 0x300aba4c
	CRC32_user_getMutableUsersV2            TLConstructor = -1795585240 // 0x94f98b28
	CRC32_user_createNewTestUser            TLConstructor = 1282329771  // 0x4c6eccab
	CRC32_user_editCloseFriends             TLConstructor = -2044429563 // 0x86247b05
	CRC32_user_setStoriesMaxId              TLConstructor = 1391834736  // 0x52f5b670
	CRC32_user_setColor                     TLConstructor = 586812791   // 0x22fa0d77
	CRC32_user_updateBirthday               TLConstructor = 1484434322  // 0x587aab92
	CRC32_user_getBirthdays                 TLConstructor = -24199258   // 0xfe8ebfa6
	CRC32_user_setStoriesHidden             TLConstructor = -138012584  // 0xf7c61858
)
