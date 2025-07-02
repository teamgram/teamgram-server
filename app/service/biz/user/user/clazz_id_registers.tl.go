/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

// ConstructorList
// RequestList

package user

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0x1835d1c, func() iface.TLObject { return &TLBotInfoData{ClazzID: 0x1835d1c} })              // 0x1835d1c
	iface.RegisterClazzID(0xb3b1a1df, func() iface.TLObject { return &TLLastSeenData{ClazzID: 0xb3b1a1df} })           // 0xb3b1a1df
	iface.RegisterClazzID(0x70ea3fa9, func() iface.TLObject { return &TLPeerPeerNotifySettings{ClazzID: 0x70ea3fa9} }) // 0x70ea3fa9
	iface.RegisterClazzID(0x4adf7bc0, func() iface.TLObject { return &TLUserImportedContacts{ClazzID: 0x4adf7bc0} })   // 0x4adf7bc0
	iface.RegisterClazzID(0x3fa3dbc7, func() iface.TLObject { return &TLUsersDataFound{ClazzID: 0x3fa3dbc7} })         // 0x3fa3dbc7
	iface.RegisterClazzID(0x80c4adfa, func() iface.TLObject { return &TLUsersIdFound{ClazzID: 0x80c4adfa} })           // 0x80c4adfa

	// Method
	iface.RegisterClazzID(0x7ca17e01, func() iface.TLObject { return &TLUserGetLastSeens{ClazzID: 0x7ca17e01} })                 // 0x7ca17e01
	iface.RegisterClazzID(0xfd405a2d, func() iface.TLObject { return &TLUserUpdateLastSeen{ClazzID: 0xfd405a2d} })               // 0xfd405a2d
	iface.RegisterClazzID(0x9119c8de, func() iface.TLObject { return &TLUserGetLastSeen{ClazzID: 0x9119c8de} })                  // 0x9119c8de
	iface.RegisterClazzID(0x376a6744, func() iface.TLObject { return &TLUserGetImmutableUser{ClazzID: 0x376a6744} })             // 0x376a6744
	iface.RegisterClazzID(0x9d3b23d7, func() iface.TLObject { return &TLUserGetMutableUsers{ClazzID: 0x9d3b23d7} })              // 0x9d3b23d7
	iface.RegisterClazzID(0xe9c36fe4, func() iface.TLObject { return &TLUserGetImmutableUserByPhone{ClazzID: 0xe9c36fe4} })      // 0xe9c36fe4
	iface.RegisterClazzID(0xff3e1373, func() iface.TLObject { return &TLUserGetImmutableUserByToken{ClazzID: 0xff3e1373} })      // 0xff3e1373
	iface.RegisterClazzID(0xd2550b4c, func() iface.TLObject { return &TLUserSetAccountDaysTTL{ClazzID: 0xd2550b4c} })            // 0xd2550b4c
	iface.RegisterClazzID(0xb2843ee0, func() iface.TLObject { return &TLUserGetAccountDaysTTL{ClazzID: 0xb2843ee0} })            // 0xb2843ee0
	iface.RegisterClazzID(0x40ac3766, func() iface.TLObject { return &TLUserGetNotifySettings{ClazzID: 0x40ac3766} })            // 0x40ac3766
	iface.RegisterClazzID(0xe465159c, func() iface.TLObject { return &TLUserGetNotifySettingsList{ClazzID: 0xe465159c} })        // 0xe465159c
	iface.RegisterClazzID(0xc9ed65e5, func() iface.TLObject { return &TLUserSetNotifySettings{ClazzID: 0xc9ed65e5} })            // 0xc9ed65e5
	iface.RegisterClazzID(0xe079d74, func() iface.TLObject { return &TLUserResetNotifySettings{ClazzID: 0xe079d74} })            // 0xe079d74
	iface.RegisterClazzID(0x55926875, func() iface.TLObject { return &TLUserGetAllNotifySettings{ClazzID: 0x55926875} })         // 0x55926875
	iface.RegisterClazzID(0x77f6f112, func() iface.TLObject { return &TLUserGetGlobalPrivacySettings{ClazzID: 0x77f6f112} })     // 0x77f6f112
	iface.RegisterClazzID(0x8cb592ae, func() iface.TLObject { return &TLUserSetGlobalPrivacySettings{ClazzID: 0x8cb592ae} })     // 0x8cb592ae
	iface.RegisterClazzID(0x9d40a3b4, func() iface.TLObject { return &TLUserGetPrivacy{ClazzID: 0x9d40a3b4} })                   // 0x9d40a3b4
	iface.RegisterClazzID(0x8855ad8f, func() iface.TLObject { return &TLUserSetPrivacy{ClazzID: 0x8855ad8f} })                   // 0x8855ad8f
	iface.RegisterClazzID(0xc56e1eaa, func() iface.TLObject { return &TLUserCheckPrivacy{ClazzID: 0xc56e1eaa} })                 // 0xc56e1eaa
	iface.RegisterClazzID(0xcae22763, func() iface.TLObject { return &TLUserAddPeerSettings{ClazzID: 0xcae22763} })              // 0xcae22763
	iface.RegisterClazzID(0xd02ef67, func() iface.TLObject { return &TLUserGetPeerSettings{ClazzID: 0xd02ef67} })                // 0xd02ef67
	iface.RegisterClazzID(0x5e891967, func() iface.TLObject { return &TLUserDeletePeerSettings{ClazzID: 0x5e891967} })           // 0x5e891967
	iface.RegisterClazzID(0xff7a575b, func() iface.TLObject { return &TLUserChangePhone{ClazzID: 0xff7a575b} })                  // 0xff7a575b
	iface.RegisterClazzID(0x79e01881, func() iface.TLObject { return &TLUserCreateNewUser{ClazzID: 0x79e01881} })                // 0x79e01881
	iface.RegisterClazzID(0x626dbd10, func() iface.TLObject { return &TLUserDeleteUser{ClazzID: 0x626dbd10} })                   // 0x626dbd10
	iface.RegisterClazzID(0x81062eb0, func() iface.TLObject { return &TLUserBlockPeer{ClazzID: 0x81062eb0} })                    // 0x81062eb0
	iface.RegisterClazzID(0xdee7160d, func() iface.TLObject { return &TLUserUnBlockPeer{ClazzID: 0xdee7160d} })                  // 0xdee7160d
	iface.RegisterClazzID(0xbba0058e, func() iface.TLObject { return &TLUserBlockedByUser{ClazzID: 0xbba0058e} })                // 0xbba0058e
	iface.RegisterClazzID(0x8caeb1df, func() iface.TLObject { return &TLUserIsBlockedByUser{ClazzID: 0x8caeb1df} })              // 0x8caeb1df
	iface.RegisterClazzID(0xc3fd70f0, func() iface.TLObject { return &TLUserCheckBlockUserList{ClazzID: 0xc3fd70f0} })           // 0xc3fd70f0
	iface.RegisterClazzID(0x23ffc348, func() iface.TLObject { return &TLUserGetBlockedList{ClazzID: 0x23ffc348} })               // 0x23ffc348
	iface.RegisterClazzID(0xe4d1d3d6, func() iface.TLObject { return &TLUserGetContactSignUpNotification{ClazzID: 0xe4d1d3d6} }) // 0xe4d1d3d6
	iface.RegisterClazzID(0x85a17361, func() iface.TLObject { return &TLUserSetContactSignUpNotification{ClazzID: 0x85a17361} }) // 0x85a17361
	iface.RegisterClazzID(0x94c3ad9f, func() iface.TLObject { return &TLUserGetContentSettings{ClazzID: 0x94c3ad9f} })           // 0x94c3ad9f
	iface.RegisterClazzID(0x9d63fe6b, func() iface.TLObject { return &TLUserSetContentSettings{ClazzID: 0x9d63fe6b} })           // 0x9d63fe6b
	iface.RegisterClazzID(0xc6018219, func() iface.TLObject { return &TLUserDeleteContact{ClazzID: 0xc6018219} })                // 0xc6018219
	iface.RegisterClazzID(0xc74bd161, func() iface.TLObject { return &TLUserGetContactList{ClazzID: 0xc74bd161} })               // 0xc74bd161
	iface.RegisterClazzID(0xf1dd983e, func() iface.TLObject { return &TLUserGetContactIdList{ClazzID: 0xf1dd983e} })             // 0xf1dd983e
	iface.RegisterClazzID(0xdb728be3, func() iface.TLObject { return &TLUserGetContact{ClazzID: 0xdb728be3} })                   // 0xdb728be3
	iface.RegisterClazzID(0x79c4bd0e, func() iface.TLObject { return &TLUserAddContact{ClazzID: 0x79c4bd0e} })                   // 0x79c4bd0e
	iface.RegisterClazzID(0x82a758a4, func() iface.TLObject { return &TLUserCheckContact{ClazzID: 0x82a758a4} })                 // 0x82a758a4
	iface.RegisterClazzID(0x47aa8212, func() iface.TLObject { return &TLUserGetImportersByPhone{ClazzID: 0x47aa8212} })          // 0x47aa8212
	iface.RegisterClazzID(0x174ddc54, func() iface.TLObject { return &TLUserDeleteImportersByPhone{ClazzID: 0x174ddc54} })       // 0x174ddc54
	iface.RegisterClazzID(0x9a00f792, func() iface.TLObject { return &TLUserImportContacts{ClazzID: 0x9a00f792} })               // 0x9a00f792
	iface.RegisterClazzID(0x12006832, func() iface.TLObject { return &TLUserGetCountryCode{ClazzID: 0x12006832} })               // 0x12006832
	iface.RegisterClazzID(0xdb00f187, func() iface.TLObject { return &TLUserUpdateAbout{ClazzID: 0xdb00f187} })                  // 0xdb00f187
	iface.RegisterClazzID(0xcb6685ec, func() iface.TLObject { return &TLUserUpdateFirstAndLastName{ClazzID: 0xcb6685ec} })       // 0xcb6685ec
	iface.RegisterClazzID(0x24c92963, func() iface.TLObject { return &TLUserUpdateVerified{ClazzID: 0x24c92963} })               // 0x24c92963
	iface.RegisterClazzID(0xf54d1e71, func() iface.TLObject { return &TLUserUpdateUsername{ClazzID: 0xf54d1e71} })               // 0xf54d1e71
	iface.RegisterClazzID(0x3b740f87, func() iface.TLObject { return &TLUserUpdateProfilePhoto{ClazzID: 0x3b740f87} })           // 0x3b740f87
	iface.RegisterClazzID(0x2be3620e, func() iface.TLObject { return &TLUserDeleteProfilePhotos{ClazzID: 0x2be3620e} })          // 0x2be3620e
	iface.RegisterClazzID(0xdc66c146, func() iface.TLObject { return &TLUserGetProfilePhotos{ClazzID: 0xdc66c146} })             // 0xdc66c146
	iface.RegisterClazzID(0x753ba916, func() iface.TLObject { return &TLUserSetBotCommands{ClazzID: 0x753ba916} })               // 0x753ba916
	iface.RegisterClazzID(0xc772c7ee, func() iface.TLObject { return &TLUserIsBot{ClazzID: 0xc772c7ee} })                        // 0xc772c7ee
	iface.RegisterClazzID(0x34663710, func() iface.TLObject { return &TLUserGetBotInfo{ClazzID: 0x34663710} })                   // 0x34663710
	iface.RegisterClazzID(0x736500c1, func() iface.TLObject { return &TLUserCheckBots{ClazzID: 0x736500c1} })                    // 0x736500c1
	iface.RegisterClazzID(0xfd10e13a, func() iface.TLObject { return &TLUserGetFullUser{ClazzID: 0xfd10e13a} })                  // 0xfd10e13a
	iface.RegisterClazzID(0xf8c8bad8, func() iface.TLObject { return &TLUserUpdateEmojiStatus{ClazzID: 0xf8c8bad8} })            // 0xf8c8bad8
	iface.RegisterClazzID(0x3bb7103, func() iface.TLObject { return &TLUserGetUserDataById{ClazzID: 0x3bb7103} })                // 0x3bb7103
	iface.RegisterClazzID(0x8191eff9, func() iface.TLObject { return &TLUserGetUserDataListByIdList{ClazzID: 0x8191eff9} })      // 0x8191eff9
	iface.RegisterClazzID(0x3f09659e, func() iface.TLObject { return &TLUserGetUserDataByToken{ClazzID: 0x3f09659e} })           // 0x3f09659e
	iface.RegisterClazzID(0x7035b6cd, func() iface.TLObject { return &TLUserSearch{ClazzID: 0x7035b6cd} })                       // 0x7035b6cd
	iface.RegisterClazzID(0x60f35d28, func() iface.TLObject { return &TLUserUpdateBotData{ClazzID: 0x60f35d28} })                // 0x60f35d28
	iface.RegisterClazzID(0x300aba4c, func() iface.TLObject { return &TLUserGetImmutableUserV2{ClazzID: 0x300aba4c} })           // 0x300aba4c
	iface.RegisterClazzID(0x94f98b28, func() iface.TLObject { return &TLUserGetMutableUsersV2{ClazzID: 0x94f98b28} })            // 0x94f98b28
	iface.RegisterClazzID(0x4c6eccab, func() iface.TLObject { return &TLUserCreateNewTestUser{ClazzID: 0x4c6eccab} })            // 0x4c6eccab
	iface.RegisterClazzID(0x86247b05, func() iface.TLObject { return &TLUserEditCloseFriends{ClazzID: 0x86247b05} })             // 0x86247b05
	iface.RegisterClazzID(0x52f5b670, func() iface.TLObject { return &TLUserSetStoriesMaxId{ClazzID: 0x52f5b670} })              // 0x52f5b670
	iface.RegisterClazzID(0x22fa0d77, func() iface.TLObject { return &TLUserSetColor{ClazzID: 0x22fa0d77} })                     // 0x22fa0d77
	iface.RegisterClazzID(0x587aab92, func() iface.TLObject { return &TLUserUpdateBirthday{ClazzID: 0x587aab92} })               // 0x587aab92
	iface.RegisterClazzID(0xfe8ebfa6, func() iface.TLObject { return &TLUserGetBirthdays{ClazzID: 0xfe8ebfa6} })                 // 0xfe8ebfa6
	iface.RegisterClazzID(0xf7c61858, func() iface.TLObject { return &TLUserSetStoriesHidden{ClazzID: 0xf7c61858} })             // 0xf7c61858
	iface.RegisterClazzID(0xc7f7bed0, func() iface.TLObject { return &TLUserUpdatePersonalChannel{ClazzID: 0xc7f7bed0} })        // 0xc7f7bed0
	iface.RegisterClazzID(0xfbab83c2, func() iface.TLObject { return &TLUserGetUserIdByPhone{ClazzID: 0xfbab83c2} })             // 0xfbab83c2
	iface.RegisterClazzID(0xd621f3f0, func() iface.TLObject { return &TLUserSetAuthorizationTTL{ClazzID: 0xd621f3f0} })          // 0xd621f3f0
	iface.RegisterClazzID(0xde6e493c, func() iface.TLObject { return &TLUserGetAuthorizationTTL{ClazzID: 0xde6e493c} })          // 0xde6e493c
	iface.RegisterClazzID(0xba08dc99, func() iface.TLObject { return &TLUserUpdatePremium{ClazzID: 0xba08dc99} })                // 0xba08dc99
	iface.RegisterClazzID(0xd3fc9ca5, func() iface.TLObject { return &TLUserGetBotInfoV2{ClazzID: 0xd3fc9ca5} })                 // 0xd3fc9ca5
}
