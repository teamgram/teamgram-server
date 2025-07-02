/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

// ConstructorList
// RequestList

package message

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor

	// Method
	iface.RegisterClazzID(0x7accb1c8, func() iface.TLObject { return &TLMessageGetUserMessage{ClazzID: 0x7accb1c8} })                       // 0x7accb1c8
	iface.RegisterClazzID(0xd3581c26, func() iface.TLObject { return &TLMessageGetUserMessageList{ClazzID: 0xd3581c26} })                   // 0xd3581c26
	iface.RegisterClazzID(0x1155a17b, func() iface.TLObject { return &TLMessageGetUserMessageListByDataIdList{ClazzID: 0x1155a17b} })       // 0x1155a17b
	iface.RegisterClazzID(0x2cb26a31, func() iface.TLObject { return &TLMessageGetUserMessageListByDataIdUserIdList{ClazzID: 0x2cb26a31} }) // 0x2cb26a31
	iface.RegisterClazzID(0x308a340, func() iface.TLObject { return &TLMessageGetHistoryMessages{ClazzID: 0x308a340} })                     // 0x308a340
	iface.RegisterClazzID(0xf507e13, func() iface.TLObject { return &TLMessageGetHistoryMessagesCount{ClazzID: 0xf507e13} })                // 0xf507e13
	iface.RegisterClazzID(0x73aeb71f, func() iface.TLObject { return &TLMessageGetPeerUserMessageId{ClazzID: 0x73aeb71f} })                 // 0x73aeb71f
	iface.RegisterClazzID(0x63129212, func() iface.TLObject { return &TLMessageGetPeerUserMessage{ClazzID: 0x63129212} })                   // 0x63129212
	iface.RegisterClazzID(0x111c2943, func() iface.TLObject { return &TLMessageSearchByMediaType{ClazzID: 0x111c2943} })                    // 0x111c2943
	iface.RegisterClazzID(0x6835b023, func() iface.TLObject { return &TLMessageSearch{ClazzID: 0x6835b023} })                               // 0x6835b023
	iface.RegisterClazzID(0xb3985dc5, func() iface.TLObject { return &TLMessageSearchGlobal{ClazzID: 0xb3985dc5} })                         // 0xb3985dc5
	iface.RegisterClazzID(0x6e735b55, func() iface.TLObject { return &TLMessageSearchByPinned{ClazzID: 0x6e735b55} })                       // 0x6e735b55
	iface.RegisterClazzID(0xe2cbbf46, func() iface.TLObject { return &TLMessageGetSearchCounter{ClazzID: 0xe2cbbf46} })                     // 0xe2cbbf46
	iface.RegisterClazzID(0xa1c62b21, func() iface.TLObject { return &TLMessageSearchV2{ClazzID: 0xa1c62b21} })                             // 0xa1c62b21
	iface.RegisterClazzID(0xaf9a082b, func() iface.TLObject { return &TLMessageGetLastTwoPinnedMessageId{ClazzID: 0xaf9a082b} })            // 0xaf9a082b
	iface.RegisterClazzID(0xf520edd0, func() iface.TLObject { return &TLMessageUpdatePinnedMessageId{ClazzID: 0xf520edd0} })                // 0xf520edd0
	iface.RegisterClazzID(0xda01d0dd, func() iface.TLObject { return &TLMessageGetPinnedMessageIdList{ClazzID: 0xda01d0dd} })               // 0xda01d0dd
	iface.RegisterClazzID(0xea0a2a73, func() iface.TLObject { return &TLMessageUnPinAllMessages{ClazzID: 0xea0a2a73} })                     // 0xea0a2a73
	iface.RegisterClazzID(0x6fe184b4, func() iface.TLObject { return &TLMessageGetUnreadMentions{ClazzID: 0x6fe184b4} })                    // 0x6fe184b4
	iface.RegisterClazzID(0xb5412049, func() iface.TLObject { return &TLMessageGetUnreadMentionsCount{ClazzID: 0xb5412049} })               // 0xb5412049
}
