/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package msg

const (
	Predicate_sender                     = "sender"
	Predicate_outboxMessage              = "outboxMessage"
	Predicate_contentMessage             = "contentMessage"
	Predicate_msg_pushUserMessage        = "msg_pushUserMessage"
	Predicate_msg_readMessageContents    = "msg_readMessageContents"
	Predicate_msg_sendMessageV2          = "msg_sendMessageV2"
	Predicate_msg_editMessage            = "msg_editMessage"
	Predicate_msg_editMessageV2          = "msg_editMessageV2"
	Predicate_msg_deleteMessages         = "msg_deleteMessages"
	Predicate_msg_deleteHistory          = "msg_deleteHistory"
	Predicate_msg_deletePhoneCallHistory = "msg_deletePhoneCallHistory"
	Predicate_msg_deleteChatHistory      = "msg_deleteChatHistory"
	Predicate_msg_readHistory            = "msg_readHistory"
	Predicate_msg_readHistoryV2          = "msg_readHistoryV2"
	Predicate_msg_updatePinnedMessage    = "msg_updatePinnedMessage"
	Predicate_msg_unpinAllMessages       = "msg_unpinAllMessages"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_sender: {
		0: 1513645242, // 0x5a3864ba

	},
	Predicate_outboxMessage: {
		0: 1402283185, // 0x539524b1

	},
	Predicate_contentMessage: {
		0: -1922780877, // 0x8d64b133

	},
	Predicate_msg_pushUserMessage: {
		0: 902887962, // 0x35d0fa1a

	},
	Predicate_msg_readMessageContents: {
		0: 673481940, // 0x282484d4

	},
	Predicate_msg_sendMessageV2: {
		0: -188056380, // 0xf4ca7cc4

	},
	Predicate_msg_editMessage: {
		0: -2129725231, // 0x810ef8d1

	},
	Predicate_msg_editMessageV2: {
		0: 1778278369, // 0x69fe5fe1

	},
	Predicate_msg_deleteMessages: {
		0: 568855069, // 0x21e80a1d

	},
	Predicate_msg_deleteHistory: {
		0: 1975576778, // 0x75c0e8ca

	},
	Predicate_msg_deletePhoneCallHistory: {
		0: 649568574, // 0x26b7a13e

	},
	Predicate_msg_deleteChatHistory: {
		0: -283155749, // 0xef1f62db

	},
	Predicate_msg_readHistory: {
		0: 1510960658, // 0x5a0f6e12

	},
	Predicate_msg_readHistoryV2: {
		0: 263827974, // 0xfb9b206

	},
	Predicate_msg_updatePinnedMessage: {
		0: -441560663, // 0xe5ae51a9

	},
	Predicate_msg_unpinAllMessages: {
		0: -1199153371, // 0xb8865f25

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	1513645242:  Predicate_sender,                     // 0x5a3864ba
	1402283185:  Predicate_outboxMessage,              // 0x539524b1
	-1922780877: Predicate_contentMessage,             // 0x8d64b133
	902887962:   Predicate_msg_pushUserMessage,        // 0x35d0fa1a
	673481940:   Predicate_msg_readMessageContents,    // 0x282484d4
	-188056380:  Predicate_msg_sendMessageV2,          // 0xf4ca7cc4
	-2129725231: Predicate_msg_editMessage,            // 0x810ef8d1
	1778278369:  Predicate_msg_editMessageV2,          // 0x69fe5fe1
	568855069:   Predicate_msg_deleteMessages,         // 0x21e80a1d
	1975576778:  Predicate_msg_deleteHistory,          // 0x75c0e8ca
	649568574:   Predicate_msg_deletePhoneCallHistory, // 0x26b7a13e
	-283155749:  Predicate_msg_deleteChatHistory,      // 0xef1f62db
	1510960658:  Predicate_msg_readHistory,            // 0x5a0f6e12
	263827974:   Predicate_msg_readHistoryV2,          // 0xfb9b206
	-441560663:  Predicate_msg_updatePinnedMessage,    // 0xe5ae51a9
	-1199153371: Predicate_msg_unpinAllMessages,       // 0xb8865f25

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
