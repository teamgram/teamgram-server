// Copyright (c) 2021-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package message

const (
	Predicate_peerMessageId                          = "peerMessageId"
	Predicate_message_getUserMessage                 = "message_getUserMessage"
	Predicate_message_getUserMessageList             = "message_getUserMessageList"
	Predicate_message_getUserMessageListByDataIdList = "message_getUserMessageListByDataIdList"
	Predicate_message_getHistoryMessages             = "message_getHistoryMessages"
	Predicate_message_getHistoryMessagesCount        = "message_getHistoryMessagesCount"
	Predicate_message_getPeerUserMessageId           = "message_getPeerUserMessageId"
	Predicate_message_getPeerUserMessage             = "message_getPeerUserMessage"
	Predicate_message_getPeerChatMessageIdList       = "message_getPeerChatMessageIdList"
	Predicate_message_getPeerChatMessageList         = "message_getPeerChatMessageList"
	Predicate_message_searchByMediaType              = "message_searchByMediaType"
	Predicate_message_search                         = "message_search"
	Predicate_message_searchGlobal                   = "message_searchGlobal"
	Predicate_message_searchByPinned                 = "message_searchByPinned"
	Predicate_message_getSearchCounter               = "message_getSearchCounter"
	Predicate_message_searchV2                       = "message_searchV2"
	Predicate_message_getLastTwoPinnedMessageId      = "message_getLastTwoPinnedMessageId"
	Predicate_message_updatePinnedMessageId          = "message_updatePinnedMessageId"
	Predicate_message_getPinnedMessageIdList         = "message_getPinnedMessageIdList"
	Predicate_message_unPinAllMessages               = "message_unPinAllMessages"
	Predicate_message_getUnreadMentions              = "message_getUnreadMentions"
	Predicate_message_getUnreadMentionsCount         = "message_getUnreadMentionsCount"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_peerMessageId: {
		0: 1988948676, // 0x768cf2c4

	},
	Predicate_message_getUserMessage: {
		0: 2060235208, // 0x7accb1c8

	},
	Predicate_message_getUserMessageList: {
		0: -749200346, // 0xd3581c26

	},
	Predicate_message_getUserMessageListByDataIdList: {
		0: 290824571, // 0x1155a17b

	},
	Predicate_message_getHistoryMessages: {
		0: 50897728, // 0x308a340

	},
	Predicate_message_getHistoryMessagesCount: {
		0: 256933395, // 0xf507e13

	},
	Predicate_message_getPeerUserMessageId: {
		0: 1940829983, // 0x73aeb71f

	},
	Predicate_message_getPeerUserMessage: {
		0: 1662161426, // 0x63129212

	},
	Predicate_message_getPeerChatMessageIdList: {
		0: -917982612, // 0xc948b26c

	},
	Predicate_message_getPeerChatMessageList: {
		0: -1442816248, // 0xaa005f08

	},
	Predicate_message_searchByMediaType: {
		0: 287058243, // 0x111c2943

	},
	Predicate_message_search: {
		0: 1748348963, // 0x6835b023

	},
	Predicate_message_searchGlobal: {
		0: -1281860155, // 0xb3985dc5

	},
	Predicate_message_searchByPinned: {
		0: 1853053781, // 0x6e735b55

	},
	Predicate_message_getSearchCounter: {
		0: -489963706, // 0xe2cbbf46

	},
	Predicate_message_searchV2: {
		0: -1580848351, // 0xa1c62b21

	},
	Predicate_message_getLastTwoPinnedMessageId: {
		0: -1348859861, // 0xaf9a082b

	},
	Predicate_message_updatePinnedMessageId: {
		0: -182391344, // 0xf520edd0

	},
	Predicate_message_getPinnedMessageIdList: {
		0: -637415203, // 0xda01d0dd

	},
	Predicate_message_unPinAllMessages: {
		0: -368432525, // 0xea0a2a73

	},
	Predicate_message_getUnreadMentions: {
		0: 1877050548, // 0x6fe184b4

	},
	Predicate_message_getUnreadMentionsCount: {
		0: -1254023095, // 0xb5412049

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	1988948676:  Predicate_peerMessageId,                          // 0x768cf2c4
	2060235208:  Predicate_message_getUserMessage,                 // 0x7accb1c8
	-749200346:  Predicate_message_getUserMessageList,             // 0xd3581c26
	290824571:   Predicate_message_getUserMessageListByDataIdList, // 0x1155a17b
	50897728:    Predicate_message_getHistoryMessages,             // 0x308a340
	256933395:   Predicate_message_getHistoryMessagesCount,        // 0xf507e13
	1940829983:  Predicate_message_getPeerUserMessageId,           // 0x73aeb71f
	1662161426:  Predicate_message_getPeerUserMessage,             // 0x63129212
	-917982612:  Predicate_message_getPeerChatMessageIdList,       // 0xc948b26c
	-1442816248: Predicate_message_getPeerChatMessageList,         // 0xaa005f08
	287058243:   Predicate_message_searchByMediaType,              // 0x111c2943
	1748348963:  Predicate_message_search,                         // 0x6835b023
	-1281860155: Predicate_message_searchGlobal,                   // 0xb3985dc5
	1853053781:  Predicate_message_searchByPinned,                 // 0x6e735b55
	-489963706:  Predicate_message_getSearchCounter,               // 0xe2cbbf46
	-1580848351: Predicate_message_searchV2,                       // 0xa1c62b21
	-1348859861: Predicate_message_getLastTwoPinnedMessageId,      // 0xaf9a082b
	-182391344:  Predicate_message_updatePinnedMessageId,          // 0xf520edd0
	-637415203:  Predicate_message_getPinnedMessageIdList,         // 0xda01d0dd
	-368432525:  Predicate_message_unPinAllMessages,               // 0xea0a2a73
	1877050548:  Predicate_message_getUnreadMentions,              // 0x6fe184b4
	-1254023095: Predicate_message_getUnreadMentionsCount,         // 0xb5412049

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
