/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package chat

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_chatInviteAlready                     = "chatInviteAlready"
	ClazzName_chatInvite                            = "chatInvite"
	ClazzName_chatInvitePeek                        = "chatInvitePeek"
	ClazzName_userChatIdList                        = "userChatIdList"
	ClazzName_recentChatInviteRequesters            = "recentChatInviteRequesters"
	ClazzName_chatInviteImported                    = "chatInviteImported"
	ClazzName_chat_getMutableChat                   = "chat_getMutableChat"
	ClazzName_chat_getChatListByIdList              = "chat_getChatListByIdList"
	ClazzName_chat_getChatBySelfId                  = "chat_getChatBySelfId"
	ClazzName_chat_createChat2                      = "chat_createChat2"
	ClazzName_chat_deleteChat                       = "chat_deleteChat"
	ClazzName_chat_deleteChatUser                   = "chat_deleteChatUser"
	ClazzName_chat_editChatTitle                    = "chat_editChatTitle"
	ClazzName_chat_editChatAbout                    = "chat_editChatAbout"
	ClazzName_chat_editChatPhoto                    = "chat_editChatPhoto"
	ClazzName_chat_editChatAdmin                    = "chat_editChatAdmin"
	ClazzName_chat_editChatDefaultBannedRights      = "chat_editChatDefaultBannedRights"
	ClazzName_chat_addChatUser                      = "chat_addChatUser"
	ClazzName_chat_getMutableChatByLink             = "chat_getMutableChatByLink"
	ClazzName_chat_toggleNoForwards                 = "chat_toggleNoForwards"
	ClazzName_chat_migratedToChannel                = "chat_migratedToChannel"
	ClazzName_chat_getChatParticipantIdList         = "chat_getChatParticipantIdList"
	ClazzName_chat_getUsersChatIdList               = "chat_getUsersChatIdList"
	ClazzName_chat_getMyChatList                    = "chat_getMyChatList"
	ClazzName_chat_exportChatInvite                 = "chat_exportChatInvite"
	ClazzName_chat_getAdminsWithInvites             = "chat_getAdminsWithInvites"
	ClazzName_chat_getExportedChatInvite            = "chat_getExportedChatInvite"
	ClazzName_chat_getExportedChatInvites           = "chat_getExportedChatInvites"
	ClazzName_chat_checkChatInvite                  = "chat_checkChatInvite"
	ClazzName_chat_importChatInvite                 = "chat_importChatInvite"
	ClazzName_chat_getChatInviteImporters           = "chat_getChatInviteImporters"
	ClazzName_chat_deleteExportedChatInvite         = "chat_deleteExportedChatInvite"
	ClazzName_chat_deleteRevokedExportedChatInvites = "chat_deleteRevokedExportedChatInvites"
	ClazzName_chat_editExportedChatInvite           = "chat_editExportedChatInvite"
	ClazzName_chat_setChatAvailableReactions        = "chat_setChatAvailableReactions"
	ClazzName_chat_setHistoryTTL                    = "chat_setHistoryTTL"
	ClazzName_chat_search                           = "chat_search"
	ClazzName_chat_getRecentChatInviteRequesters    = "chat_getRecentChatInviteRequesters"
	ClazzName_chat_hideChatJoinRequests             = "chat_hideChatJoinRequests"
	ClazzName_chat_importChatInvite2                = "chat_importChatInvite2"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_chatInviteAlready, 0, 0xa40e7d5e)                     // a40e7d5e
	iface.RegisterClazzName(ClazzName_chatInvite, 0, 0xdb75d1a7)                            // db75d1a7
	iface.RegisterClazzName(ClazzName_chatInvitePeek, 0, 0xace3e26e)                        // ace3e26e
	iface.RegisterClazzName(ClazzName_userChatIdList, 0, 0x50067224)                        // 50067224
	iface.RegisterClazzName(ClazzName_recentChatInviteRequesters, 0, 0x1c6e3c54)            // 1c6e3c54
	iface.RegisterClazzName(ClazzName_chatInviteImported, 0, 0x721051f6)                    // 721051f6
	iface.RegisterClazzName(ClazzName_chat_getMutableChat, 0, 0x2c2c25d2)                   // 2c2c25d2
	iface.RegisterClazzName(ClazzName_chat_getChatListByIdList, 0, 0xe740f539)              // e740f539
	iface.RegisterClazzName(ClazzName_chat_getChatBySelfId, 0, 0x49b71a48)                  // 49b71a48
	iface.RegisterClazzName(ClazzName_chat_createChat2, 0, 0xf77448d2)                      // f77448d2
	iface.RegisterClazzName(ClazzName_chat_deleteChat, 0, 0x6d11ec1e)                       // 6d11ec1e
	iface.RegisterClazzName(ClazzName_chat_deleteChatUser, 0, 0xb270fd5)                    // b270fd5
	iface.RegisterClazzName(ClazzName_chat_editChatTitle, 0, 0x95c59ea7)                    // 95c59ea7
	iface.RegisterClazzName(ClazzName_chat_editChatAbout, 0, 0x5c737c78)                    // 5c737c78
	iface.RegisterClazzName(ClazzName_chat_editChatPhoto, 0, 0x45c2a668)                    // 45c2a668
	iface.RegisterClazzName(ClazzName_chat_editChatAdmin, 0, 0x1905e5ec)                    // 1905e5ec
	iface.RegisterClazzName(ClazzName_chat_editChatDefaultBannedRights, 0, 0x5a34a687)      // 5a34a687
	iface.RegisterClazzName(ClazzName_chat_addChatUser, 0, 0xe5554168)                      // e5554168
	iface.RegisterClazzName(ClazzName_chat_getMutableChatByLink, 0, 0xa266278b)             // a266278b
	iface.RegisterClazzName(ClazzName_chat_toggleNoForwards, 0, 0xd5952af9)                 // d5952af9
	iface.RegisterClazzName(ClazzName_chat_migratedToChannel, 0, 0x83faadf)                 // 83faadf
	iface.RegisterClazzName(ClazzName_chat_getChatParticipantIdList, 0, 0x329622a9)         // 329622a9
	iface.RegisterClazzName(ClazzName_chat_getUsersChatIdList, 0, 0x2f36ab4c)               // 2f36ab4c
	iface.RegisterClazzName(ClazzName_chat_getMyChatList, 0, 0xf3756c88)                    // f3756c88
	iface.RegisterClazzName(ClazzName_chat_exportChatInvite, 0, 0xc5cf804b)                 // c5cf804b
	iface.RegisterClazzName(ClazzName_chat_getAdminsWithInvites, 0, 0xd2ea41d2)             // d2ea41d2
	iface.RegisterClazzName(ClazzName_chat_getExportedChatInvite, 0, 0xddea3250)            // ddea3250
	iface.RegisterClazzName(ClazzName_chat_getExportedChatInvites, 0, 0xb48f18f6)           // b48f18f6
	iface.RegisterClazzName(ClazzName_chat_checkChatInvite, 0, 0x7387f28c)                  // 7387f28c
	iface.RegisterClazzName(ClazzName_chat_importChatInvite, 0, 0x58e660d4)                 // 58e660d4
	iface.RegisterClazzName(ClazzName_chat_getChatInviteImporters, 0, 0x9846557f)           // 9846557f
	iface.RegisterClazzName(ClazzName_chat_deleteExportedChatInvite, 0, 0x562288b8)         // 562288b8
	iface.RegisterClazzName(ClazzName_chat_deleteRevokedExportedChatInvites, 0, 0xd0126269) // d0126269
	iface.RegisterClazzName(ClazzName_chat_editExportedChatInvite, 0, 0xaf994c76)           // af994c76
	iface.RegisterClazzName(ClazzName_chat_setChatAvailableReactions, 0, 0xc4d08972)        // c4d08972
	iface.RegisterClazzName(ClazzName_chat_setHistoryTTL, 0, 0x3cfb6384)                    // 3cfb6384
	iface.RegisterClazzName(ClazzName_chat_search, 0, 0x21e014fb)                           // 21e014fb
	iface.RegisterClazzName(ClazzName_chat_getRecentChatInviteRequesters, 0, 0xfedc1098)    // fedc1098
	iface.RegisterClazzName(ClazzName_chat_hideChatJoinRequests, 0, 0x3ea52cd1)             // 3ea52cd1
	iface.RegisterClazzName(ClazzName_chat_importChatInvite2, 0, 0xdcd93dbf)                // dcd93dbf

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_chatInviteAlready, 0xa40e7d5e)                     // a40e7d5e
	iface.RegisterClazzIDName(ClazzName_chatInvite, 0xdb75d1a7)                            // db75d1a7
	iface.RegisterClazzIDName(ClazzName_chatInvitePeek, 0xace3e26e)                        // ace3e26e
	iface.RegisterClazzIDName(ClazzName_userChatIdList, 0x50067224)                        // 50067224
	iface.RegisterClazzIDName(ClazzName_recentChatInviteRequesters, 0x1c6e3c54)            // 1c6e3c54
	iface.RegisterClazzIDName(ClazzName_chatInviteImported, 0x721051f6)                    // 721051f6
	iface.RegisterClazzIDName(ClazzName_chat_getMutableChat, 0x2c2c25d2)                   // 2c2c25d2
	iface.RegisterClazzIDName(ClazzName_chat_getChatListByIdList, 0xe740f539)              // e740f539
	iface.RegisterClazzIDName(ClazzName_chat_getChatBySelfId, 0x49b71a48)                  // 49b71a48
	iface.RegisterClazzIDName(ClazzName_chat_createChat2, 0xf77448d2)                      // f77448d2
	iface.RegisterClazzIDName(ClazzName_chat_deleteChat, 0x6d11ec1e)                       // 6d11ec1e
	iface.RegisterClazzIDName(ClazzName_chat_deleteChatUser, 0xb270fd5)                    // b270fd5
	iface.RegisterClazzIDName(ClazzName_chat_editChatTitle, 0x95c59ea7)                    // 95c59ea7
	iface.RegisterClazzIDName(ClazzName_chat_editChatAbout, 0x5c737c78)                    // 5c737c78
	iface.RegisterClazzIDName(ClazzName_chat_editChatPhoto, 0x45c2a668)                    // 45c2a668
	iface.RegisterClazzIDName(ClazzName_chat_editChatAdmin, 0x1905e5ec)                    // 1905e5ec
	iface.RegisterClazzIDName(ClazzName_chat_editChatDefaultBannedRights, 0x5a34a687)      // 5a34a687
	iface.RegisterClazzIDName(ClazzName_chat_addChatUser, 0xe5554168)                      // e5554168
	iface.RegisterClazzIDName(ClazzName_chat_getMutableChatByLink, 0xa266278b)             // a266278b
	iface.RegisterClazzIDName(ClazzName_chat_toggleNoForwards, 0xd5952af9)                 // d5952af9
	iface.RegisterClazzIDName(ClazzName_chat_migratedToChannel, 0x83faadf)                 // 83faadf
	iface.RegisterClazzIDName(ClazzName_chat_getChatParticipantIdList, 0x329622a9)         // 329622a9
	iface.RegisterClazzIDName(ClazzName_chat_getUsersChatIdList, 0x2f36ab4c)               // 2f36ab4c
	iface.RegisterClazzIDName(ClazzName_chat_getMyChatList, 0xf3756c88)                    // f3756c88
	iface.RegisterClazzIDName(ClazzName_chat_exportChatInvite, 0xc5cf804b)                 // c5cf804b
	iface.RegisterClazzIDName(ClazzName_chat_getAdminsWithInvites, 0xd2ea41d2)             // d2ea41d2
	iface.RegisterClazzIDName(ClazzName_chat_getExportedChatInvite, 0xddea3250)            // ddea3250
	iface.RegisterClazzIDName(ClazzName_chat_getExportedChatInvites, 0xb48f18f6)           // b48f18f6
	iface.RegisterClazzIDName(ClazzName_chat_checkChatInvite, 0x7387f28c)                  // 7387f28c
	iface.RegisterClazzIDName(ClazzName_chat_importChatInvite, 0x58e660d4)                 // 58e660d4
	iface.RegisterClazzIDName(ClazzName_chat_getChatInviteImporters, 0x9846557f)           // 9846557f
	iface.RegisterClazzIDName(ClazzName_chat_deleteExportedChatInvite, 0x562288b8)         // 562288b8
	iface.RegisterClazzIDName(ClazzName_chat_deleteRevokedExportedChatInvites, 0xd0126269) // d0126269
	iface.RegisterClazzIDName(ClazzName_chat_editExportedChatInvite, 0xaf994c76)           // af994c76
	iface.RegisterClazzIDName(ClazzName_chat_setChatAvailableReactions, 0xc4d08972)        // c4d08972
	iface.RegisterClazzIDName(ClazzName_chat_setHistoryTTL, 0x3cfb6384)                    // 3cfb6384
	iface.RegisterClazzIDName(ClazzName_chat_search, 0x21e014fb)                           // 21e014fb
	iface.RegisterClazzIDName(ClazzName_chat_getRecentChatInviteRequesters, 0xfedc1098)    // fedc1098
	iface.RegisterClazzIDName(ClazzName_chat_hideChatJoinRequests, 0x3ea52cd1)             // 3ea52cd1
	iface.RegisterClazzIDName(ClazzName_chat_importChatInvite2, 0xdcd93dbf)                // dcd93dbf
}
