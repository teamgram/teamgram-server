/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package msg

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

const (
	ClazzName_sender                            = "sender"
	ClazzName_outboxMessage                     = "outboxMessage"
	ClazzName_contentMessage                    = "contentMessage"
	ClazzName_resolvedDialogCursor              = "resolvedDialogCursor"
	ClazzName_updateFact                        = "updateFact"
	ClazzName_msg_pushUserMessage               = "msg_pushUserMessage"
	ClazzName_msg_readMessageContents           = "msg_readMessageContents"
	ClazzName_msg_sendMessage                   = "msg_sendMessage"
	ClazzName_msg_editMessage                   = "msg_editMessage"
	ClazzName_msg_deleteMessages                = "msg_deleteMessages"
	ClazzName_msg_deleteHistory                 = "msg_deleteHistory"
	ClazzName_msg_deletePhoneCallHistory        = "msg_deletePhoneCallHistory"
	ClazzName_msg_deleteChatHistory             = "msg_deleteChatHistory"
	ClazzName_msg_readHistory                   = "msg_readHistory"
	ClazzName_msg_readHistoryV2                 = "msg_readHistoryV2"
	ClazzName_msg_getHistory                    = "msg_getHistory"
	ClazzName_msg_getUserMessage                = "msg_getUserMessage"
	ClazzName_msg_getUserMessageList            = "msg_getUserMessageList"
	ClazzName_msg_searchHashtag                 = "msg_searchHashtag"
	ClazzName_msg_resolveDialogCursorTopMessage = "msg_resolveDialogCursorTopMessage"
	ClazzName_msg_updatePinnedMessage           = "msg_updatePinnedMessage"
	ClazzName_msg_unpinAllMessages              = "msg_unpinAllMessages"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_sender, 0, 0x5a3864ba)                            // 5a3864ba
	iface.RegisterClazzName(ClazzName_outboxMessage, 0, 0x625d8b25)                     // 625d8b25
	iface.RegisterClazzName(ClazzName_contentMessage, 0, 0x8d64b133)                    // 8d64b133
	iface.RegisterClazzName(ClazzName_resolvedDialogCursor, 0, 0x7debda91)              // 7debda91
	iface.RegisterClazzName(ClazzName_updateFact, 0, 0x4561a083)                        // 4561a083
	iface.RegisterClazzName(ClazzName_msg_pushUserMessage, 0, 0x35d0fa1a)               // 35d0fa1a
	iface.RegisterClazzName(ClazzName_msg_readMessageContents, 0, 0x282484d4)           // 282484d4
	iface.RegisterClazzName(ClazzName_msg_sendMessage, 0, 0x93e882df)                   // 93e882df
	iface.RegisterClazzName(ClazzName_msg_editMessage, 0, 0x1ddc94)                     // 1ddc94
	iface.RegisterClazzName(ClazzName_msg_deleteMessages, 0, 0x21e80a1d)                // 21e80a1d
	iface.RegisterClazzName(ClazzName_msg_deleteHistory, 0, 0x75c0e8ca)                 // 75c0e8ca
	iface.RegisterClazzName(ClazzName_msg_deletePhoneCallHistory, 0, 0x26b7a13e)        // 26b7a13e
	iface.RegisterClazzName(ClazzName_msg_deleteChatHistory, 0, 0xef1f62db)             // ef1f62db
	iface.RegisterClazzName(ClazzName_msg_readHistory, 0, 0x5a0f6e12)                   // 5a0f6e12
	iface.RegisterClazzName(ClazzName_msg_readHistoryV2, 0, 0xfb9b206)                  // fb9b206
	iface.RegisterClazzName(ClazzName_msg_getHistory, 0, 0x7f4083df)                    // 7f4083df
	iface.RegisterClazzName(ClazzName_msg_getUserMessage, 0, 0x385f5e90)                // 385f5e90
	iface.RegisterClazzName(ClazzName_msg_getUserMessageList, 0, 0xfb80f3c1)            // fb80f3c1
	iface.RegisterClazzName(ClazzName_msg_searchHashtag, 0, 0x7e39bca9)                 // 7e39bca9
	iface.RegisterClazzName(ClazzName_msg_resolveDialogCursorTopMessage, 0, 0xc5d16bc5) // c5d16bc5
	iface.RegisterClazzName(ClazzName_msg_updatePinnedMessage, 0, 0xe5ae51a9)           // e5ae51a9
	iface.RegisterClazzName(ClazzName_msg_unpinAllMessages, 0, 0xb8865f25)              // b8865f25

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_sender, 0x5a3864ba)                            // 5a3864ba
	iface.RegisterClazzIDName(ClazzName_outboxMessage, 0x625d8b25)                     // 625d8b25
	iface.RegisterClazzIDName(ClazzName_contentMessage, 0x8d64b133)                    // 8d64b133
	iface.RegisterClazzIDName(ClazzName_resolvedDialogCursor, 0x7debda91)              // 7debda91
	iface.RegisterClazzIDName(ClazzName_updateFact, 0x4561a083)                        // 4561a083
	iface.RegisterClazzIDName(ClazzName_msg_pushUserMessage, 0x35d0fa1a)               // 35d0fa1a
	iface.RegisterClazzIDName(ClazzName_msg_readMessageContents, 0x282484d4)           // 282484d4
	iface.RegisterClazzIDName(ClazzName_msg_sendMessage, 0x93e882df)                   // 93e882df
	iface.RegisterClazzIDName(ClazzName_msg_editMessage, 0x1ddc94)                     // 1ddc94
	iface.RegisterClazzIDName(ClazzName_msg_deleteMessages, 0x21e80a1d)                // 21e80a1d
	iface.RegisterClazzIDName(ClazzName_msg_deleteHistory, 0x75c0e8ca)                 // 75c0e8ca
	iface.RegisterClazzIDName(ClazzName_msg_deletePhoneCallHistory, 0x26b7a13e)        // 26b7a13e
	iface.RegisterClazzIDName(ClazzName_msg_deleteChatHistory, 0xef1f62db)             // ef1f62db
	iface.RegisterClazzIDName(ClazzName_msg_readHistory, 0x5a0f6e12)                   // 5a0f6e12
	iface.RegisterClazzIDName(ClazzName_msg_readHistoryV2, 0xfb9b206)                  // fb9b206
	iface.RegisterClazzIDName(ClazzName_msg_getHistory, 0x7f4083df)                    // 7f4083df
	iface.RegisterClazzIDName(ClazzName_msg_getUserMessage, 0x385f5e90)                // 385f5e90
	iface.RegisterClazzIDName(ClazzName_msg_getUserMessageList, 0xfb80f3c1)            // fb80f3c1
	iface.RegisterClazzIDName(ClazzName_msg_searchHashtag, 0x7e39bca9)                 // 7e39bca9
	iface.RegisterClazzIDName(ClazzName_msg_resolveDialogCursorTopMessage, 0xc5d16bc5) // c5d16bc5
	iface.RegisterClazzIDName(ClazzName_msg_updatePinnedMessage, 0xe5ae51a9)           // e5ae51a9
	iface.RegisterClazzIDName(ClazzName_msg_unpinAllMessages, 0xb8865f25)              // b8865f25
}
