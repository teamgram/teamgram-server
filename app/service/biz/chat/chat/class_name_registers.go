// Copyright (c) 2021-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package chat

const (
	Predicate_immutableChatParticipant              = "immutableChatParticipant"
	Predicate_immutableChat                         = "immutableChat"
	Predicate_mutableChat                           = "mutableChat"
	Predicate_chatInviteAlready                     = "chatInviteAlready"
	Predicate_chatInvite                            = "chatInvite"
	Predicate_chatInvitePeek                        = "chatInvitePeek"
	Predicate_userChatIdList                        = "userChatIdList"
	Predicate_chat_getMutableChat                   = "chat_getMutableChat"
	Predicate_chat_getChatListByIdList              = "chat_getChatListByIdList"
	Predicate_chat_getChatBySelfId                  = "chat_getChatBySelfId"
	Predicate_chat_createChat2                      = "chat_createChat2"
	Predicate_chat_deleteChat                       = "chat_deleteChat"
	Predicate_chat_deleteChatUser                   = "chat_deleteChatUser"
	Predicate_chat_editChatTitle                    = "chat_editChatTitle"
	Predicate_chat_editChatAbout                    = "chat_editChatAbout"
	Predicate_chat_editChatPhoto                    = "chat_editChatPhoto"
	Predicate_chat_editChatAdmin                    = "chat_editChatAdmin"
	Predicate_chat_editChatDefaultBannedRights      = "chat_editChatDefaultBannedRights"
	Predicate_chat_addChatUser                      = "chat_addChatUser"
	Predicate_chat_getMutableChatByLink             = "chat_getMutableChatByLink"
	Predicate_chat_toggleNoForwards                 = "chat_toggleNoForwards"
	Predicate_chat_migratedToChannel                = "chat_migratedToChannel"
	Predicate_chat_getChatParticipantIdList         = "chat_getChatParticipantIdList"
	Predicate_chat_getUsersChatIdList               = "chat_getUsersChatIdList"
	Predicate_chat_getMyChatList                    = "chat_getMyChatList"
	Predicate_chat_exportChatInvite                 = "chat_exportChatInvite"
	Predicate_chat_getAdminsWithInvites             = "chat_getAdminsWithInvites"
	Predicate_chat_getExportedChatInvite            = "chat_getExportedChatInvite"
	Predicate_chat_getExportedChatInvites           = "chat_getExportedChatInvites"
	Predicate_chat_checkChatInvite                  = "chat_checkChatInvite"
	Predicate_chat_importChatInvite                 = "chat_importChatInvite"
	Predicate_chat_getChatInviteImporters           = "chat_getChatInviteImporters"
	Predicate_chat_deleteExportedChatInvite         = "chat_deleteExportedChatInvite"
	Predicate_chat_deleteRevokedExportedChatInvites = "chat_deleteRevokedExportedChatInvites"
	Predicate_chat_editExportedChatInvite           = "chat_editExportedChatInvite"
	Predicate_chat_setChatAvailableReactions        = "chat_setChatAvailableReactions"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_immutableChatParticipant: {
		0: 650553001, // 0x26c6a6a9

	},
	Predicate_immutableChat: {
		0: -771834191, // 0xd1febeb1

	},
	Predicate_mutableChat: {
		0: -34609042, // 0xfdefe86e

	},
	Predicate_chatInviteAlready: {
		0: -1542554274, // 0xa40e7d5e

	},
	Predicate_chatInvite: {
		0: -613035609, // 0xdb75d1a7

	},
	Predicate_chatInvitePeek: {
		0: -1394351506, // 0xace3e26e

	},
	Predicate_userChatIdList: {
		0: 1342599716, // 0x50067224

	},
	Predicate_chat_getMutableChat: {
		0: 741090770, // 0x2c2c25d2

	},
	Predicate_chat_getChatListByIdList: {
		0: -415173319, // 0xe740f539

	},
	Predicate_chat_getChatBySelfId: {
		0: 1236736584, // 0x49b71a48

	},
	Predicate_chat_createChat2: {
		0: -465608273, // 0xe43f61af

	},
	Predicate_chat_deleteChat: {
		0: 1829891102, // 0x6d11ec1e

	},
	Predicate_chat_deleteChatUser: {
		0: 187109333, // 0xb270fd5

	},
	Predicate_chat_editChatTitle: {
		0: -1782210905, // 0x95c59ea7

	},
	Predicate_chat_editChatAbout: {
		0: 1551072376, // 0x5c737c78

	},
	Predicate_chat_editChatPhoto: {
		0: 1170384488, // 0x45c2a668

	},
	Predicate_chat_editChatAdmin: {
		0: 419816940, // 0x1905e5ec

	},
	Predicate_chat_editChatDefaultBannedRights: {
		0: 1513399943, // 0x5a34a687

	},
	Predicate_chat_addChatUser: {
		0: 2086511757, // 0x7c5da48d

	},
	Predicate_chat_getMutableChatByLink: {
		0: -1570363509, // 0xa266278b

	},
	Predicate_chat_toggleNoForwards: {
		0: -711644423, // 0xd5952af9

	},
	Predicate_chat_migratedToChannel: {
		0: 138390239, // 0x83faadf

	},
	Predicate_chat_getChatParticipantIdList: {
		0: 848700073, // 0x329622a9

	},
	Predicate_chat_getUsersChatIdList: {
		0: 792111948, // 0x2f36ab4c

	},
	Predicate_chat_getMyChatList: {
		0: -210408312, // 0xf3756c88

	},
	Predicate_chat_exportChatInvite: {
		0: -976256949, // 0xc5cf804b

	},
	Predicate_chat_getAdminsWithInvites: {
		0: -756399662, // 0xd2ea41d2

	},
	Predicate_chat_getExportedChatInvite: {
		0: -571854256, // 0xddea3250

	},
	Predicate_chat_getExportedChatInvites: {
		0: -1265690378, // 0xb48f18f6

	},
	Predicate_chat_checkChatInvite: {
		0: 1938289292, // 0x7387f28c

	},
	Predicate_chat_importChatInvite: {
		0: 1491493076, // 0x58e660d4

	},
	Predicate_chat_getChatInviteImporters: {
		0: -1740221057, // 0x9846557f

	},
	Predicate_chat_deleteExportedChatInvite: {
		0: 1445103800, // 0x562288b8

	},
	Predicate_chat_deleteRevokedExportedChatInvites: {
		0: -804101527, // 0xd0126269

	},
	Predicate_chat_editExportedChatInvite: {
		0: -1348907914, // 0xaf994c76

	},
	Predicate_chat_setChatAvailableReactions: {
		0: 1372233637, // 0x51ca9fa5

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	650553001:   Predicate_immutableChatParticipant,              // 0x26c6a6a9
	-771834191:  Predicate_immutableChat,                         // 0xd1febeb1
	-34609042:   Predicate_mutableChat,                           // 0xfdefe86e
	-1542554274: Predicate_chatInviteAlready,                     // 0xa40e7d5e
	-613035609:  Predicate_chatInvite,                            // 0xdb75d1a7
	-1394351506: Predicate_chatInvitePeek,                        // 0xace3e26e
	1342599716:  Predicate_userChatIdList,                        // 0x50067224
	741090770:   Predicate_chat_getMutableChat,                   // 0x2c2c25d2
	-415173319:  Predicate_chat_getChatListByIdList,              // 0xe740f539
	1236736584:  Predicate_chat_getChatBySelfId,                  // 0x49b71a48
	-465608273:  Predicate_chat_createChat2,                      // 0xe43f61af
	1829891102:  Predicate_chat_deleteChat,                       // 0x6d11ec1e
	187109333:   Predicate_chat_deleteChatUser,                   // 0xb270fd5
	-1782210905: Predicate_chat_editChatTitle,                    // 0x95c59ea7
	1551072376:  Predicate_chat_editChatAbout,                    // 0x5c737c78
	1170384488:  Predicate_chat_editChatPhoto,                    // 0x45c2a668
	419816940:   Predicate_chat_editChatAdmin,                    // 0x1905e5ec
	1513399943:  Predicate_chat_editChatDefaultBannedRights,      // 0x5a34a687
	2086511757:  Predicate_chat_addChatUser,                      // 0x7c5da48d
	-1570363509: Predicate_chat_getMutableChatByLink,             // 0xa266278b
	-711644423:  Predicate_chat_toggleNoForwards,                 // 0xd5952af9
	138390239:   Predicate_chat_migratedToChannel,                // 0x83faadf
	848700073:   Predicate_chat_getChatParticipantIdList,         // 0x329622a9
	792111948:   Predicate_chat_getUsersChatIdList,               // 0x2f36ab4c
	-210408312:  Predicate_chat_getMyChatList,                    // 0xf3756c88
	-976256949:  Predicate_chat_exportChatInvite,                 // 0xc5cf804b
	-756399662:  Predicate_chat_getAdminsWithInvites,             // 0xd2ea41d2
	-571854256:  Predicate_chat_getExportedChatInvite,            // 0xddea3250
	-1265690378: Predicate_chat_getExportedChatInvites,           // 0xb48f18f6
	1938289292:  Predicate_chat_checkChatInvite,                  // 0x7387f28c
	1491493076:  Predicate_chat_importChatInvite,                 // 0x58e660d4
	-1740221057: Predicate_chat_getChatInviteImporters,           // 0x9846557f
	1445103800:  Predicate_chat_deleteExportedChatInvite,         // 0x562288b8
	-804101527:  Predicate_chat_deleteRevokedExportedChatInvites, // 0xd0126269
	-1348907914: Predicate_chat_editExportedChatInvite,           // 0xaf994c76
	1372233637:  Predicate_chat_setChatAvailableReactions,        // 0x51ca9fa5

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
