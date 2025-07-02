/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package msg

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_sender                     = "sender"
	ClazzName_outboxMessage              = "outboxMessage"
	ClazzName_contentMessage             = "contentMessage"
	ClazzName_msg_pushUserMessage        = "msg_pushUserMessage"
	ClazzName_msg_readMessageContents    = "msg_readMessageContents"
	ClazzName_msg_sendMessageV2          = "msg_sendMessageV2"
	ClazzName_msg_editMessageV2          = "msg_editMessageV2"
	ClazzName_msg_deleteMessages         = "msg_deleteMessages"
	ClazzName_msg_deleteHistory          = "msg_deleteHistory"
	ClazzName_msg_deletePhoneCallHistory = "msg_deletePhoneCallHistory"
	ClazzName_msg_deleteChatHistory      = "msg_deleteChatHistory"
	ClazzName_msg_readHistory            = "msg_readHistory"
	ClazzName_msg_readHistoryV2          = "msg_readHistoryV2"
	ClazzName_msg_updatePinnedMessage    = "msg_updatePinnedMessage"
	ClazzName_msg_unpinAllMessages       = "msg_unpinAllMessages"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_sender, 0, 0x5a3864ba)                     // 5a3864ba
	iface.RegisterClazzName(ClazzName_outboxMessage, 0, 0x539524b1)              // 539524b1
	iface.RegisterClazzName(ClazzName_contentMessage, 0, 0x8d64b133)             // 8d64b133
	iface.RegisterClazzName(ClazzName_msg_pushUserMessage, 0, 0x35d0fa1a)        // 35d0fa1a
	iface.RegisterClazzName(ClazzName_msg_readMessageContents, 0, 0x282484d4)    // 282484d4
	iface.RegisterClazzName(ClazzName_msg_sendMessageV2, 0, 0xf4ca7cc4)          // f4ca7cc4
	iface.RegisterClazzName(ClazzName_msg_editMessageV2, 0, 0x69fe5fe1)          // 69fe5fe1
	iface.RegisterClazzName(ClazzName_msg_deleteMessages, 0, 0x21e80a1d)         // 21e80a1d
	iface.RegisterClazzName(ClazzName_msg_deleteHistory, 0, 0x75c0e8ca)          // 75c0e8ca
	iface.RegisterClazzName(ClazzName_msg_deletePhoneCallHistory, 0, 0x26b7a13e) // 26b7a13e
	iface.RegisterClazzName(ClazzName_msg_deleteChatHistory, 0, 0xef1f62db)      // ef1f62db
	iface.RegisterClazzName(ClazzName_msg_readHistory, 0, 0x5a0f6e12)            // 5a0f6e12
	iface.RegisterClazzName(ClazzName_msg_readHistoryV2, 0, 0xfb9b206)           // fb9b206
	iface.RegisterClazzName(ClazzName_msg_updatePinnedMessage, 0, 0xe5ae51a9)    // e5ae51a9
	iface.RegisterClazzName(ClazzName_msg_unpinAllMessages, 0, 0xb8865f25)       // b8865f25

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_sender, 0x5a3864ba)                     // 5a3864ba
	iface.RegisterClazzIDName(ClazzName_outboxMessage, 0x539524b1)              // 539524b1
	iface.RegisterClazzIDName(ClazzName_contentMessage, 0x8d64b133)             // 8d64b133
	iface.RegisterClazzIDName(ClazzName_msg_pushUserMessage, 0x35d0fa1a)        // 35d0fa1a
	iface.RegisterClazzIDName(ClazzName_msg_readMessageContents, 0x282484d4)    // 282484d4
	iface.RegisterClazzIDName(ClazzName_msg_sendMessageV2, 0xf4ca7cc4)          // f4ca7cc4
	iface.RegisterClazzIDName(ClazzName_msg_editMessageV2, 0x69fe5fe1)          // 69fe5fe1
	iface.RegisterClazzIDName(ClazzName_msg_deleteMessages, 0x21e80a1d)         // 21e80a1d
	iface.RegisterClazzIDName(ClazzName_msg_deleteHistory, 0x75c0e8ca)          // 75c0e8ca
	iface.RegisterClazzIDName(ClazzName_msg_deletePhoneCallHistory, 0x26b7a13e) // 26b7a13e
	iface.RegisterClazzIDName(ClazzName_msg_deleteChatHistory, 0xef1f62db)      // ef1f62db
	iface.RegisterClazzIDName(ClazzName_msg_readHistory, 0x5a0f6e12)            // 5a0f6e12
	iface.RegisterClazzIDName(ClazzName_msg_readHistoryV2, 0xfb9b206)           // fb9b206
	iface.RegisterClazzIDName(ClazzName_msg_updatePinnedMessage, 0xe5ae51a9)    // e5ae51a9
	iface.RegisterClazzIDName(ClazzName_msg_unpinAllMessages, 0xb8865f25)       // b8865f25
}
