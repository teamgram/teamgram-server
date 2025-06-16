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

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_message_getUserMessage                       = "message_getUserMessage"
	ClazzName_message_getUserMessageList                   = "message_getUserMessageList"
	ClazzName_message_getUserMessageListByDataIdList       = "message_getUserMessageListByDataIdList"
	ClazzName_message_getUserMessageListByDataIdUserIdList = "message_getUserMessageListByDataIdUserIdList"
	ClazzName_message_getHistoryMessages                   = "message_getHistoryMessages"
	ClazzName_message_getHistoryMessagesCount              = "message_getHistoryMessagesCount"
	ClazzName_message_getPeerUserMessageId                 = "message_getPeerUserMessageId"
	ClazzName_message_getPeerUserMessage                   = "message_getPeerUserMessage"
	ClazzName_message_searchByMediaType                    = "message_searchByMediaType"
	ClazzName_message_search                               = "message_search"
	ClazzName_message_searchGlobal                         = "message_searchGlobal"
	ClazzName_message_searchByPinned                       = "message_searchByPinned"
	ClazzName_message_getSearchCounter                     = "message_getSearchCounter"
	ClazzName_message_searchV2                             = "message_searchV2"
	ClazzName_message_getLastTwoPinnedMessageId            = "message_getLastTwoPinnedMessageId"
	ClazzName_message_updatePinnedMessageId                = "message_updatePinnedMessageId"
	ClazzName_message_getPinnedMessageIdList               = "message_getPinnedMessageIdList"
	ClazzName_message_unPinAllMessages                     = "message_unPinAllMessages"
	ClazzName_message_getUnreadMentions                    = "message_getUnreadMentions"
	ClazzName_message_getUnreadMentionsCount               = "message_getUnreadMentionsCount"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_message_getUserMessage, 0, 0x7accb1c8)                       // 7accb1c8
	iface.RegisterClazzName(ClazzName_message_getUserMessageList, 0, 0xd3581c26)                   // d3581c26
	iface.RegisterClazzName(ClazzName_message_getUserMessageListByDataIdList, 0, 0x1155a17b)       // 1155a17b
	iface.RegisterClazzName(ClazzName_message_getUserMessageListByDataIdUserIdList, 0, 0x2cb26a31) // 2cb26a31
	iface.RegisterClazzName(ClazzName_message_getHistoryMessages, 0, 0x308a340)                    // 308a340
	iface.RegisterClazzName(ClazzName_message_getHistoryMessagesCount, 0, 0xf507e13)               // f507e13
	iface.RegisterClazzName(ClazzName_message_getPeerUserMessageId, 0, 0x73aeb71f)                 // 73aeb71f
	iface.RegisterClazzName(ClazzName_message_getPeerUserMessage, 0, 0x63129212)                   // 63129212
	iface.RegisterClazzName(ClazzName_message_searchByMediaType, 0, 0x111c2943)                    // 111c2943
	iface.RegisterClazzName(ClazzName_message_search, 0, 0x6835b023)                               // 6835b023
	iface.RegisterClazzName(ClazzName_message_searchGlobal, 0, 0xb3985dc5)                         // b3985dc5
	iface.RegisterClazzName(ClazzName_message_searchByPinned, 0, 0x6e735b55)                       // 6e735b55
	iface.RegisterClazzName(ClazzName_message_getSearchCounter, 0, 0xe2cbbf46)                     // e2cbbf46
	iface.RegisterClazzName(ClazzName_message_searchV2, 0, 0xa1c62b21)                             // a1c62b21
	iface.RegisterClazzName(ClazzName_message_getLastTwoPinnedMessageId, 0, 0xaf9a082b)            // af9a082b
	iface.RegisterClazzName(ClazzName_message_updatePinnedMessageId, 0, 0xf520edd0)                // f520edd0
	iface.RegisterClazzName(ClazzName_message_getPinnedMessageIdList, 0, 0xda01d0dd)               // da01d0dd
	iface.RegisterClazzName(ClazzName_message_unPinAllMessages, 0, 0xea0a2a73)                     // ea0a2a73
	iface.RegisterClazzName(ClazzName_message_getUnreadMentions, 0, 0x6fe184b4)                    // 6fe184b4
	iface.RegisterClazzName(ClazzName_message_getUnreadMentionsCount, 0, 0xb5412049)               // b5412049

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_message_getUserMessage, 0x7accb1c8)                       // 7accb1c8
	iface.RegisterClazzIDName(ClazzName_message_getUserMessageList, 0xd3581c26)                   // d3581c26
	iface.RegisterClazzIDName(ClazzName_message_getUserMessageListByDataIdList, 0x1155a17b)       // 1155a17b
	iface.RegisterClazzIDName(ClazzName_message_getUserMessageListByDataIdUserIdList, 0x2cb26a31) // 2cb26a31
	iface.RegisterClazzIDName(ClazzName_message_getHistoryMessages, 0x308a340)                    // 308a340
	iface.RegisterClazzIDName(ClazzName_message_getHistoryMessagesCount, 0xf507e13)               // f507e13
	iface.RegisterClazzIDName(ClazzName_message_getPeerUserMessageId, 0x73aeb71f)                 // 73aeb71f
	iface.RegisterClazzIDName(ClazzName_message_getPeerUserMessage, 0x63129212)                   // 63129212
	iface.RegisterClazzIDName(ClazzName_message_searchByMediaType, 0x111c2943)                    // 111c2943
	iface.RegisterClazzIDName(ClazzName_message_search, 0x6835b023)                               // 6835b023
	iface.RegisterClazzIDName(ClazzName_message_searchGlobal, 0xb3985dc5)                         // b3985dc5
	iface.RegisterClazzIDName(ClazzName_message_searchByPinned, 0x6e735b55)                       // 6e735b55
	iface.RegisterClazzIDName(ClazzName_message_getSearchCounter, 0xe2cbbf46)                     // e2cbbf46
	iface.RegisterClazzIDName(ClazzName_message_searchV2, 0xa1c62b21)                             // a1c62b21
	iface.RegisterClazzIDName(ClazzName_message_getLastTwoPinnedMessageId, 0xaf9a082b)            // af9a082b
	iface.RegisterClazzIDName(ClazzName_message_updatePinnedMessageId, 0xf520edd0)                // f520edd0
	iface.RegisterClazzIDName(ClazzName_message_getPinnedMessageIdList, 0xda01d0dd)               // da01d0dd
	iface.RegisterClazzIDName(ClazzName_message_unPinAllMessages, 0xea0a2a73)                     // ea0a2a73
	iface.RegisterClazzIDName(ClazzName_message_getUnreadMentions, 0x6fe184b4)                    // 6fe184b4
	iface.RegisterClazzIDName(ClazzName_message_getUnreadMentionsCount, 0xb5412049)               // b5412049
}
