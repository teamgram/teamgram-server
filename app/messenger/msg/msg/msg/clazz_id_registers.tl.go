/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

// ConstructorList
// RequestList

package msg

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0x8d64b133, func() iface.TLObject { return &TLContentMessage{ClazzID: 0x8d64b133} }) // 0x8d64b133
	iface.RegisterClazzID(0x539524b1, func() iface.TLObject { return &TLOutboxMessage{ClazzID: 0x539524b1} })  // 0x539524b1
	iface.RegisterClazzID(0x5a3864ba, func() iface.TLObject { return &TLSender{ClazzID: 0x5a3864ba} })         // 0x5a3864ba

	// Method
	iface.RegisterClazzID(0x35d0fa1a, func() iface.TLObject { return &TLMsgPushUserMessage{ClazzID: 0x35d0fa1a} })        // 0x35d0fa1a
	iface.RegisterClazzID(0x282484d4, func() iface.TLObject { return &TLMsgReadMessageContents{ClazzID: 0x282484d4} })    // 0x282484d4
	iface.RegisterClazzID(0xf4ca7cc4, func() iface.TLObject { return &TLMsgSendMessageV2{ClazzID: 0xf4ca7cc4} })          // 0xf4ca7cc4
	iface.RegisterClazzID(0x69fe5fe1, func() iface.TLObject { return &TLMsgEditMessageV2{ClazzID: 0x69fe5fe1} })          // 0x69fe5fe1
	iface.RegisterClazzID(0x21e80a1d, func() iface.TLObject { return &TLMsgDeleteMessages{ClazzID: 0x21e80a1d} })         // 0x21e80a1d
	iface.RegisterClazzID(0x75c0e8ca, func() iface.TLObject { return &TLMsgDeleteHistory{ClazzID: 0x75c0e8ca} })          // 0x75c0e8ca
	iface.RegisterClazzID(0x26b7a13e, func() iface.TLObject { return &TLMsgDeletePhoneCallHistory{ClazzID: 0x26b7a13e} }) // 0x26b7a13e
	iface.RegisterClazzID(0xef1f62db, func() iface.TLObject { return &TLMsgDeleteChatHistory{ClazzID: 0xef1f62db} })      // 0xef1f62db
	iface.RegisterClazzID(0x5a0f6e12, func() iface.TLObject { return &TLMsgReadHistory{ClazzID: 0x5a0f6e12} })            // 0x5a0f6e12
	iface.RegisterClazzID(0xfb9b206, func() iface.TLObject { return &TLMsgReadHistoryV2{ClazzID: 0xfb9b206} })            // 0xfb9b206
	iface.RegisterClazzID(0xe5ae51a9, func() iface.TLObject { return &TLMsgUpdatePinnedMessage{ClazzID: 0xe5ae51a9} })    // 0xe5ae51a9
	iface.RegisterClazzID(0xb8865f25, func() iface.TLObject { return &TLMsgUnpinAllMessages{ClazzID: 0xb8865f25} })       // 0xb8865f25
}
