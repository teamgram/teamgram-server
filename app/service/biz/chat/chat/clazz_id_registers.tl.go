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

package chat

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0xa40e7d5e, func() iface.TLObject { return &TLChatInviteAlready{ClazzID: 0xa40e7d5e} })          // 0xa40e7d5e
	iface.RegisterClazzID(0xdb75d1a7, func() iface.TLObject { return &TLChatInvite{ClazzID: 0xdb75d1a7} })                 // 0xdb75d1a7
	iface.RegisterClazzID(0xace3e26e, func() iface.TLObject { return &TLChatInvitePeek{ClazzID: 0xace3e26e} })             // 0xace3e26e
	iface.RegisterClazzID(0x721051f6, func() iface.TLObject { return &TLChatInviteImported{ClazzID: 0x721051f6} })         // 0x721051f6
	iface.RegisterClazzID(0x1c6e3c54, func() iface.TLObject { return &TLRecentChatInviteRequesters{ClazzID: 0x1c6e3c54} }) // 0x1c6e3c54
	iface.RegisterClazzID(0x50067224, func() iface.TLObject { return &TLUserChatIdList{ClazzID: 0x50067224} })             // 0x50067224

	// Method
	iface.RegisterClazzID(0x2c2c25d2, func() iface.TLObject { return &TLChatGetMutableChat{ClazzID: 0x2c2c25d2} })                   // 0x2c2c25d2
	iface.RegisterClazzID(0xe740f539, func() iface.TLObject { return &TLChatGetChatListByIdList{ClazzID: 0xe740f539} })              // 0xe740f539
	iface.RegisterClazzID(0x49b71a48, func() iface.TLObject { return &TLChatGetChatBySelfId{ClazzID: 0x49b71a48} })                  // 0x49b71a48
	iface.RegisterClazzID(0xf77448d2, func() iface.TLObject { return &TLChatCreateChat2{ClazzID: 0xf77448d2} })                      // 0xf77448d2
	iface.RegisterClazzID(0x6d11ec1e, func() iface.TLObject { return &TLChatDeleteChat{ClazzID: 0x6d11ec1e} })                       // 0x6d11ec1e
	iface.RegisterClazzID(0xb270fd5, func() iface.TLObject { return &TLChatDeleteChatUser{ClazzID: 0xb270fd5} })                     // 0xb270fd5
	iface.RegisterClazzID(0x95c59ea7, func() iface.TLObject { return &TLChatEditChatTitle{ClazzID: 0x95c59ea7} })                    // 0x95c59ea7
	iface.RegisterClazzID(0x5c737c78, func() iface.TLObject { return &TLChatEditChatAbout{ClazzID: 0x5c737c78} })                    // 0x5c737c78
	iface.RegisterClazzID(0x45c2a668, func() iface.TLObject { return &TLChatEditChatPhoto{ClazzID: 0x45c2a668} })                    // 0x45c2a668
	iface.RegisterClazzID(0x1905e5ec, func() iface.TLObject { return &TLChatEditChatAdmin{ClazzID: 0x1905e5ec} })                    // 0x1905e5ec
	iface.RegisterClazzID(0x5a34a687, func() iface.TLObject { return &TLChatEditChatDefaultBannedRights{ClazzID: 0x5a34a687} })      // 0x5a34a687
	iface.RegisterClazzID(0xe5554168, func() iface.TLObject { return &TLChatAddChatUser{ClazzID: 0xe5554168} })                      // 0xe5554168
	iface.RegisterClazzID(0xa266278b, func() iface.TLObject { return &TLChatGetMutableChatByLink{ClazzID: 0xa266278b} })             // 0xa266278b
	iface.RegisterClazzID(0xd5952af9, func() iface.TLObject { return &TLChatToggleNoForwards{ClazzID: 0xd5952af9} })                 // 0xd5952af9
	iface.RegisterClazzID(0x83faadf, func() iface.TLObject { return &TLChatMigratedToChannel{ClazzID: 0x83faadf} })                  // 0x83faadf
	iface.RegisterClazzID(0x329622a9, func() iface.TLObject { return &TLChatGetChatParticipantIdList{ClazzID: 0x329622a9} })         // 0x329622a9
	iface.RegisterClazzID(0x2f36ab4c, func() iface.TLObject { return &TLChatGetUsersChatIdList{ClazzID: 0x2f36ab4c} })               // 0x2f36ab4c
	iface.RegisterClazzID(0xf3756c88, func() iface.TLObject { return &TLChatGetMyChatList{ClazzID: 0xf3756c88} })                    // 0xf3756c88
	iface.RegisterClazzID(0xc5cf804b, func() iface.TLObject { return &TLChatExportChatInvite{ClazzID: 0xc5cf804b} })                 // 0xc5cf804b
	iface.RegisterClazzID(0xd2ea41d2, func() iface.TLObject { return &TLChatGetAdminsWithInvites{ClazzID: 0xd2ea41d2} })             // 0xd2ea41d2
	iface.RegisterClazzID(0xddea3250, func() iface.TLObject { return &TLChatGetExportedChatInvite{ClazzID: 0xddea3250} })            // 0xddea3250
	iface.RegisterClazzID(0xb48f18f6, func() iface.TLObject { return &TLChatGetExportedChatInvites{ClazzID: 0xb48f18f6} })           // 0xb48f18f6
	iface.RegisterClazzID(0x7387f28c, func() iface.TLObject { return &TLChatCheckChatInvite{ClazzID: 0x7387f28c} })                  // 0x7387f28c
	iface.RegisterClazzID(0x58e660d4, func() iface.TLObject { return &TLChatImportChatInvite{ClazzID: 0x58e660d4} })                 // 0x58e660d4
	iface.RegisterClazzID(0x9846557f, func() iface.TLObject { return &TLChatGetChatInviteImporters{ClazzID: 0x9846557f} })           // 0x9846557f
	iface.RegisterClazzID(0x562288b8, func() iface.TLObject { return &TLChatDeleteExportedChatInvite{ClazzID: 0x562288b8} })         // 0x562288b8
	iface.RegisterClazzID(0xd0126269, func() iface.TLObject { return &TLChatDeleteRevokedExportedChatInvites{ClazzID: 0xd0126269} }) // 0xd0126269
	iface.RegisterClazzID(0xaf994c76, func() iface.TLObject { return &TLChatEditExportedChatInvite{ClazzID: 0xaf994c76} })           // 0xaf994c76
	iface.RegisterClazzID(0xc4d08972, func() iface.TLObject { return &TLChatSetChatAvailableReactions{ClazzID: 0xc4d08972} })        // 0xc4d08972
	iface.RegisterClazzID(0x3cfb6384, func() iface.TLObject { return &TLChatSetHistoryTTL{ClazzID: 0x3cfb6384} })                    // 0x3cfb6384
	iface.RegisterClazzID(0x21e014fb, func() iface.TLObject { return &TLChatSearch{ClazzID: 0x21e014fb} })                           // 0x21e014fb
	iface.RegisterClazzID(0xfedc1098, func() iface.TLObject { return &TLChatGetRecentChatInviteRequesters{ClazzID: 0xfedc1098} })    // 0xfedc1098
	iface.RegisterClazzID(0x3ea52cd1, func() iface.TLObject { return &TLChatHideChatJoinRequests{ClazzID: 0x3ea52cd1} })             // 0x3ea52cd1
	iface.RegisterClazzID(0xdcd93dbf, func() iface.TLObject { return &TLChatImportChatInvite2{ClazzID: 0xdcd93dbf} })                // 0xdcd93dbf
}
