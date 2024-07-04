/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package inbox

const (
	Predicate_inboxMessageData                 = "inboxMessageData"
	Predicate_inboxMessageId                   = "inboxMessageId"
	Predicate_inbox_editUserMessageToInbox     = "inbox_editUserMessageToInbox"
	Predicate_inbox_editChatMessageToInbox     = "inbox_editChatMessageToInbox"
	Predicate_inbox_deleteMessagesToInbox      = "inbox_deleteMessagesToInbox"
	Predicate_inbox_deleteUserHistoryToInbox   = "inbox_deleteUserHistoryToInbox"
	Predicate_inbox_deleteChatHistoryToInbox   = "inbox_deleteChatHistoryToInbox"
	Predicate_inbox_readUserMediaUnreadToInbox = "inbox_readUserMediaUnreadToInbox"
	Predicate_inbox_readChatMediaUnreadToInbox = "inbox_readChatMediaUnreadToInbox"
	Predicate_inbox_updateHistoryReaded        = "inbox_updateHistoryReaded"
	Predicate_inbox_updatePinnedMessage        = "inbox_updatePinnedMessage"
	Predicate_inbox_unpinAllMessages           = "inbox_unpinAllMessages"
	Predicate_inbox_sendUserMessageToInboxV2   = "inbox_sendUserMessageToInboxV2"
	Predicate_inbox_editMessageToInboxV2       = "inbox_editMessageToInboxV2"
	Predicate_inbox_readInboxHistory           = "inbox_readInboxHistory"
	Predicate_inbox_readOutboxHistory          = "inbox_readOutboxHistory"
	Predicate_inbox_readMediaUnreadToInboxV2   = "inbox_readMediaUnreadToInboxV2"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_inboxMessageData: {
		0: 1002286548, // 0x3bbdadd4

	},
	Predicate_inboxMessageId: {
		0: -963460705, // 0xc692c19f

	},
	Predicate_inbox_editUserMessageToInbox: {
		0: 1559967656, // 0x5cfb37a8

	},
	Predicate_inbox_editChatMessageToInbox: {
		0: 2031122959, // 0x79107a0f

	},
	Predicate_inbox_deleteMessagesToInbox: {
		0: -2061734348, // 0x851c6e34

	},
	Predicate_inbox_deleteUserHistoryToInbox: {
		0: 336232792, // 0x140a8158

	},
	Predicate_inbox_deleteChatHistoryToInbox: {
		0: -659905022, // 0xd8aaa602

	},
	Predicate_inbox_readUserMediaUnreadToInbox: {
		0: 364970827, // 0x15c1034b

	},
	Predicate_inbox_readChatMediaUnreadToInbox: {
		0: 1430347220, // 0x55415dd4

	},
	Predicate_inbox_updateHistoryReaded: {
		0: -1010283296, // 0xc3c84ce0

	},
	Predicate_inbox_updatePinnedMessage: {
		0: -1452528908, // 0xa96c2af4

	},
	Predicate_inbox_unpinAllMessages: {
		0: 589079137, // 0x231ca261

	},
	Predicate_inbox_sendUserMessageToInboxV2: {
		0: 2043341160, // 0x79cae968

	},
	Predicate_inbox_editMessageToInboxV2: {
		0: -625238423, // 0xdabb9e69

	},
	Predicate_inbox_readInboxHistory: {
		0: -465427029, // 0xe44225ab

	},
	Predicate_inbox_readOutboxHistory: {
		0: 477116106, // 0x1c7036ca

	},
	Predicate_inbox_readMediaUnreadToInboxV2: {
		0: -356170942, // 0xeac54342

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	1002286548:  Predicate_inboxMessageData,                 // 0x3bbdadd4
	-963460705:  Predicate_inboxMessageId,                   // 0xc692c19f
	1559967656:  Predicate_inbox_editUserMessageToInbox,     // 0x5cfb37a8
	2031122959:  Predicate_inbox_editChatMessageToInbox,     // 0x79107a0f
	-2061734348: Predicate_inbox_deleteMessagesToInbox,      // 0x851c6e34
	336232792:   Predicate_inbox_deleteUserHistoryToInbox,   // 0x140a8158
	-659905022:  Predicate_inbox_deleteChatHistoryToInbox,   // 0xd8aaa602
	364970827:   Predicate_inbox_readUserMediaUnreadToInbox, // 0x15c1034b
	1430347220:  Predicate_inbox_readChatMediaUnreadToInbox, // 0x55415dd4
	-1010283296: Predicate_inbox_updateHistoryReaded,        // 0xc3c84ce0
	-1452528908: Predicate_inbox_updatePinnedMessage,        // 0xa96c2af4
	589079137:   Predicate_inbox_unpinAllMessages,           // 0x231ca261
	2043341160:  Predicate_inbox_sendUserMessageToInboxV2,   // 0x79cae968
	-625238423:  Predicate_inbox_editMessageToInboxV2,       // 0xdabb9e69
	-465427029:  Predicate_inbox_readInboxHistory,           // 0xe44225ab
	477116106:   Predicate_inbox_readOutboxHistory,          // 0x1c7036ca
	-356170942:  Predicate_inbox_readMediaUnreadToInboxV2,   // 0xeac54342

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
