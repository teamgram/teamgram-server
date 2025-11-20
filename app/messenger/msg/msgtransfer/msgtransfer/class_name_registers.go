/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package msgtransfer

const (
	Predicate_sender                          = "sender"
	Predicate_outboxMessage                   = "outboxMessage"
	Predicate_contentMessage                  = "contentMessage"
	Predicate_inboxMessageData                = "inboxMessageData"
	Predicate_inboxMessageId                  = "inboxMessageId"
	Predicate_msgtransfer_sendMessageToOutbox = "msgtransfer_sendMessageToOutbox"
	Predicate_msgtransfer_sendMessageToInbox  = "msgtransfer_sendMessageToInbox"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_sender: {
		0: 1513645242, // 0x5a3864ba

	},
	Predicate_outboxMessage: {
		0: 1763737728, // 0x69208080

	},
	Predicate_contentMessage: {
		0: -1922780877, // 0x8d64b133

	},
	Predicate_inboxMessageData: {
		0: 1002286548, // 0x3bbdadd4

	},
	Predicate_inboxMessageId: {
		0: -963460705, // 0xc692c19f

	},
	Predicate_msgtransfer_sendMessageToOutbox: {
		0: -508367556, // 0xe1b2ed3c

	},
	Predicate_msgtransfer_sendMessageToInbox: {
		0: -750661413, // 0xd341d0db

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	1513645242:  Predicate_sender,                          // 0x5a3864ba
	1763737728:  Predicate_outboxMessage,                   // 0x69208080
	-1922780877: Predicate_contentMessage,                  // 0x8d64b133
	1002286548:  Predicate_inboxMessageData,                // 0x3bbdadd4
	-963460705:  Predicate_inboxMessageId,                  // 0xc692c19f
	-508367556:  Predicate_msgtransfer_sendMessageToOutbox, // 0xe1b2ed3c
	-750661413:  Predicate_msgtransfer_sendMessageToInbox,  // 0xd341d0db

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
