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
	CRC32_UNKNOWN                                      TLConstructor = 0
	CRC32_message_getUserMessage                       TLConstructor = 2060235208  // 0x7accb1c8
	CRC32_message_getUserMessageList                   TLConstructor = -749200346  // 0xd3581c26
	CRC32_message_getUserMessageListByDataIdList       TLConstructor = 290824571   // 0x1155a17b
	CRC32_message_getUserMessageListByDataIdUserIdList TLConstructor = 749890097   // 0x2cb26a31
	CRC32_message_getHistoryMessages                   TLConstructor = 50897728    // 0x308a340
	CRC32_message_getHistoryMessagesCount              TLConstructor = 256933395   // 0xf507e13
	CRC32_message_getPeerUserMessageId                 TLConstructor = 1940829983  // 0x73aeb71f
	CRC32_message_getPeerUserMessage                   TLConstructor = 1662161426  // 0x63129212
	CRC32_message_searchByMediaType                    TLConstructor = -1152381832 // 0xbb500c78
	CRC32_message_search                               TLConstructor = 251910661   // 0xf03da05
	CRC32_message_searchGlobal                         TLConstructor = 1113214626  // 0x425a4ea2
	CRC32_message_searchByPinned                       TLConstructor = 721580084   // 0x2b027034
	CRC32_message_getSearchCounter                     TLConstructor = -489963706  // 0xe2cbbf46
	CRC32_message_searchV2                             TLConstructor = -356633351  // 0xeabe34f9
	CRC32_message_getLastTwoPinnedMessageId            TLConstructor = -1348859861 // 0xaf9a082b
	CRC32_message_updatePinnedMessageId                TLConstructor = -182391344  // 0xf520edd0
	CRC32_message_getPinnedMessageIdList               TLConstructor = -637415203  // 0xda01d0dd
	CRC32_message_unPinAllMessages                     TLConstructor = -368432525  // 0xea0a2a73
	CRC32_message_getUnreadMentions                    TLConstructor = 1877050548  // 0x6fe184b4
	CRC32_message_getUnreadMentionsCount               TLConstructor = -1254023095 // 0xb5412049
	CRC32_message_getSavedHistoryMessages              TLConstructor = -60243377   // 0xfc68c24f
)
