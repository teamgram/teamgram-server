/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package inbox

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_inboxMessageData                 = "inboxMessageData"
	ClazzName_inboxMessageId                   = "inboxMessageId"
	ClazzName_inbox_editUserMessageToInbox     = "inbox_editUserMessageToInbox"
	ClazzName_inbox_editChatMessageToInbox     = "inbox_editChatMessageToInbox"
	ClazzName_inbox_deleteMessagesToInbox      = "inbox_deleteMessagesToInbox"
	ClazzName_inbox_deleteUserHistoryToInbox   = "inbox_deleteUserHistoryToInbox"
	ClazzName_inbox_deleteChatHistoryToInbox   = "inbox_deleteChatHistoryToInbox"
	ClazzName_inbox_readUserMediaUnreadToInbox = "inbox_readUserMediaUnreadToInbox"
	ClazzName_inbox_readChatMediaUnreadToInbox = "inbox_readChatMediaUnreadToInbox"
	ClazzName_inbox_updateHistoryReaded        = "inbox_updateHistoryReaded"
	ClazzName_inbox_updatePinnedMessage        = "inbox_updatePinnedMessage"
	ClazzName_inbox_unpinAllMessages           = "inbox_unpinAllMessages"
	ClazzName_inbox_sendUserMessageToInboxV2   = "inbox_sendUserMessageToInboxV2"
	ClazzName_inbox_editMessageToInboxV2       = "inbox_editMessageToInboxV2"
	ClazzName_inbox_readInboxHistory           = "inbox_readInboxHistory"
	ClazzName_inbox_readOutboxHistory          = "inbox_readOutboxHistory"
	ClazzName_inbox_readMediaUnreadToInboxV2   = "inbox_readMediaUnreadToInboxV2"
	ClazzName_inbox_updatePinnedMessageV2      = "inbox_updatePinnedMessageV2"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_inboxMessageData, 0, 0x3bbdadd4)                 // 3bbdadd4
	iface.RegisterClazzName(ClazzName_inboxMessageId, 0, 0xc692c19f)                   // c692c19f
	iface.RegisterClazzName(ClazzName_inbox_editUserMessageToInbox, 0, 0x5cfb37a8)     // 5cfb37a8
	iface.RegisterClazzName(ClazzName_inbox_editChatMessageToInbox, 0, 0x79107a0f)     // 79107a0f
	iface.RegisterClazzName(ClazzName_inbox_deleteMessagesToInbox, 0, 0x851c6e34)      // 851c6e34
	iface.RegisterClazzName(ClazzName_inbox_deleteUserHistoryToInbox, 0, 0x140a8158)   // 140a8158
	iface.RegisterClazzName(ClazzName_inbox_deleteChatHistoryToInbox, 0, 0xd8aaa602)   // d8aaa602
	iface.RegisterClazzName(ClazzName_inbox_readUserMediaUnreadToInbox, 0, 0x15c1034b) // 15c1034b
	iface.RegisterClazzName(ClazzName_inbox_readChatMediaUnreadToInbox, 0, 0x55415dd4) // 55415dd4
	iface.RegisterClazzName(ClazzName_inbox_updateHistoryReaded, 0, 0xc3c84ce0)        // c3c84ce0
	iface.RegisterClazzName(ClazzName_inbox_updatePinnedMessage, 0, 0xa96c2af4)        // a96c2af4
	iface.RegisterClazzName(ClazzName_inbox_unpinAllMessages, 0, 0x231ca261)           // 231ca261
	iface.RegisterClazzName(ClazzName_inbox_sendUserMessageToInboxV2, 0, 0x5bd7522)    // 5bd7522
	iface.RegisterClazzName(ClazzName_inbox_editMessageToInboxV2, 0, 0xdabb9e69)       // dabb9e69
	iface.RegisterClazzName(ClazzName_inbox_readInboxHistory, 0, 0x1f73675)            // 1f73675
	iface.RegisterClazzName(ClazzName_inbox_readOutboxHistory, 0, 0x1c7036ca)          // 1c7036ca
	iface.RegisterClazzName(ClazzName_inbox_readMediaUnreadToInboxV2, 0, 0xeac54342)   // eac54342
	iface.RegisterClazzName(ClazzName_inbox_updatePinnedMessageV2, 0, 0x56b79e7c)      // 56b79e7c

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_inboxMessageData, 0x3bbdadd4)                 // 3bbdadd4
	iface.RegisterClazzIDName(ClazzName_inboxMessageId, 0xc692c19f)                   // c692c19f
	iface.RegisterClazzIDName(ClazzName_inbox_editUserMessageToInbox, 0x5cfb37a8)     // 5cfb37a8
	iface.RegisterClazzIDName(ClazzName_inbox_editChatMessageToInbox, 0x79107a0f)     // 79107a0f
	iface.RegisterClazzIDName(ClazzName_inbox_deleteMessagesToInbox, 0x851c6e34)      // 851c6e34
	iface.RegisterClazzIDName(ClazzName_inbox_deleteUserHistoryToInbox, 0x140a8158)   // 140a8158
	iface.RegisterClazzIDName(ClazzName_inbox_deleteChatHistoryToInbox, 0xd8aaa602)   // d8aaa602
	iface.RegisterClazzIDName(ClazzName_inbox_readUserMediaUnreadToInbox, 0x15c1034b) // 15c1034b
	iface.RegisterClazzIDName(ClazzName_inbox_readChatMediaUnreadToInbox, 0x55415dd4) // 55415dd4
	iface.RegisterClazzIDName(ClazzName_inbox_updateHistoryReaded, 0xc3c84ce0)        // c3c84ce0
	iface.RegisterClazzIDName(ClazzName_inbox_updatePinnedMessage, 0xa96c2af4)        // a96c2af4
	iface.RegisterClazzIDName(ClazzName_inbox_unpinAllMessages, 0x231ca261)           // 231ca261
	iface.RegisterClazzIDName(ClazzName_inbox_sendUserMessageToInboxV2, 0x5bd7522)    // 5bd7522
	iface.RegisterClazzIDName(ClazzName_inbox_editMessageToInboxV2, 0xdabb9e69)       // dabb9e69
	iface.RegisterClazzIDName(ClazzName_inbox_readInboxHistory, 0x1f73675)            // 1f73675
	iface.RegisterClazzIDName(ClazzName_inbox_readOutboxHistory, 0x1c7036ca)          // 1c7036ca
	iface.RegisterClazzIDName(ClazzName_inbox_readMediaUnreadToInboxV2, 0xeac54342)   // eac54342
	iface.RegisterClazzIDName(ClazzName_inbox_updatePinnedMessageV2, 0x56b79e7c)      // 56b79e7c
}
