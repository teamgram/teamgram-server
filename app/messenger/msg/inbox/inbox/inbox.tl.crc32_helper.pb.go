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
	CRC32_UNKNOWN                          TLConstructor = 0
	CRC32_inboxMessageData                 TLConstructor = 1002286548  // 0x3bbdadd4
	CRC32_inboxMessageId                   TLConstructor = -963460705  // 0xc692c19f
	CRC32_inbox_editUserMessageToInbox     TLConstructor = 1559967656  // 0x5cfb37a8
	CRC32_inbox_editChatMessageToInbox     TLConstructor = 2031122959  // 0x79107a0f
	CRC32_inbox_deleteMessagesToInbox      TLConstructor = -2061734348 // 0x851c6e34
	CRC32_inbox_deleteUserHistoryToInbox   TLConstructor = 336232792   // 0x140a8158
	CRC32_inbox_deleteChatHistoryToInbox   TLConstructor = -659905022  // 0xd8aaa602
	CRC32_inbox_readUserMediaUnreadToInbox TLConstructor = 364970827   // 0x15c1034b
	CRC32_inbox_readChatMediaUnreadToInbox TLConstructor = 1430347220  // 0x55415dd4
	CRC32_inbox_updateHistoryReaded        TLConstructor = -1010283296 // 0xc3c84ce0
	CRC32_inbox_updatePinnedMessage        TLConstructor = -1452528908 // 0xa96c2af4
	CRC32_inbox_unpinAllMessages           TLConstructor = 589079137   // 0x231ca261
	CRC32_inbox_sendUserMessageToInboxV2   TLConstructor = 2043341160  // 0x79cae968
	CRC32_inbox_editMessageToInboxV2       TLConstructor = -625238423  // 0xdabb9e69
	CRC32_inbox_readInboxHistory           TLConstructor = -465427029  // 0xe44225ab
	CRC32_inbox_readOutboxHistory          TLConstructor = 477116106   // 0x1c7036ca
	CRC32_inbox_readMediaUnreadToInboxV2   TLConstructor = -356170942  // 0xeac54342
)
