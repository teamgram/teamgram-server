/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package message

const (
	Predicate_message_getUserMessage                       = "message_getUserMessage"
	Predicate_message_getUserMessageList                   = "message_getUserMessageList"
	Predicate_message_getUserMessageListByDataIdList       = "message_getUserMessageListByDataIdList"
	Predicate_message_getUserMessageListByDataIdUserIdList = "message_getUserMessageListByDataIdUserIdList"
	Predicate_message_getHistoryMessages                   = "message_getHistoryMessages"
	Predicate_message_getHistoryMessagesCount              = "message_getHistoryMessagesCount"
	Predicate_message_getPeerUserMessageId                 = "message_getPeerUserMessageId"
	Predicate_message_getPeerUserMessage                   = "message_getPeerUserMessage"
	Predicate_message_searchByMediaType                    = "message_searchByMediaType"
	Predicate_message_search                               = "message_search"
	Predicate_message_searchGlobal                         = "message_searchGlobal"
	Predicate_message_searchByPinned                       = "message_searchByPinned"
	Predicate_message_getSearchCounter                     = "message_getSearchCounter"
	Predicate_message_searchV2                             = "message_searchV2"
	Predicate_message_getLastTwoPinnedMessageId            = "message_getLastTwoPinnedMessageId"
	Predicate_message_updatePinnedMessageId                = "message_updatePinnedMessageId"
	Predicate_message_getPinnedMessageIdList               = "message_getPinnedMessageIdList"
	Predicate_message_unPinAllMessages                     = "message_unPinAllMessages"
	Predicate_message_getUnreadMentions                    = "message_getUnreadMentions"
	Predicate_message_getUnreadMentionsCount               = "message_getUnreadMentionsCount"
	Predicate_message_getSavedHistoryMessages              = "message_getSavedHistoryMessages"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_message_getUserMessage: {
		0: 2060235208, // 0x7accb1c8

	},
	Predicate_message_getUserMessageList: {
		0: -749200346, // 0xd3581c26

	},
	Predicate_message_getUserMessageListByDataIdList: {
		0: 290824571, // 0x1155a17b

	},
	Predicate_message_getUserMessageListByDataIdUserIdList: {
		0: 749890097, // 0x2cb26a31

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
	Predicate_message_searchByMediaType: {
		0: -1152381832, // 0xbb500c78

	},
	Predicate_message_search: {
		0: 251910661, // 0xf03da05

	},
	Predicate_message_searchGlobal: {
		0: 1113214626, // 0x425a4ea2

	},
	Predicate_message_searchByPinned: {
		0: 721580084, // 0x2b027034

	},
	Predicate_message_getSearchCounter: {
		0: -489963706, // 0xe2cbbf46

	},
	Predicate_message_searchV2: {
		0: -356633351, // 0xeabe34f9

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
	Predicate_message_getSavedHistoryMessages: {
		0: -60243377, // 0xfc68c24f

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	2060235208:  Predicate_message_getUserMessage,                       // 0x7accb1c8
	-749200346:  Predicate_message_getUserMessageList,                   // 0xd3581c26
	290824571:   Predicate_message_getUserMessageListByDataIdList,       // 0x1155a17b
	749890097:   Predicate_message_getUserMessageListByDataIdUserIdList, // 0x2cb26a31
	50897728:    Predicate_message_getHistoryMessages,                   // 0x308a340
	256933395:   Predicate_message_getHistoryMessagesCount,              // 0xf507e13
	1940829983:  Predicate_message_getPeerUserMessageId,                 // 0x73aeb71f
	1662161426:  Predicate_message_getPeerUserMessage,                   // 0x63129212
	-1152381832: Predicate_message_searchByMediaType,                    // 0xbb500c78
	251910661:   Predicate_message_search,                               // 0xf03da05
	1113214626:  Predicate_message_searchGlobal,                         // 0x425a4ea2
	721580084:   Predicate_message_searchByPinned,                       // 0x2b027034
	-489963706:  Predicate_message_getSearchCounter,                     // 0xe2cbbf46
	-356633351:  Predicate_message_searchV2,                             // 0xeabe34f9
	-1348859861: Predicate_message_getLastTwoPinnedMessageId,            // 0xaf9a082b
	-182391344:  Predicate_message_updatePinnedMessageId,                // 0xf520edd0
	-637415203:  Predicate_message_getPinnedMessageIdList,               // 0xda01d0dd
	-368432525:  Predicate_message_unPinAllMessages,                     // 0xea0a2a73
	1877050548:  Predicate_message_getUnreadMentions,                    // 0x6fe184b4
	-1254023095: Predicate_message_getUnreadMentionsCount,               // 0xb5412049
	-60243377:   Predicate_message_getSavedHistoryMessages,              // 0xfc68c24f

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
