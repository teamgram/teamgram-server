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

package inbox

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0x3bbdadd4, func() iface.TLObject { return &TLInboxMessageData{ClazzID: 0x3bbdadd4} }) // 0x3bbdadd4
	iface.RegisterClazzID(0xc692c19f, func() iface.TLObject { return &TLInboxMessageId{ClazzID: 0xc692c19f} })   // 0xc692c19f

	// Method
	iface.RegisterClazzID(0x5cfb37a8, func() iface.TLObject { return &TLInboxEditUserMessageToInbox{ClazzID: 0x5cfb37a8} })     // 0x5cfb37a8
	iface.RegisterClazzID(0x79107a0f, func() iface.TLObject { return &TLInboxEditChatMessageToInbox{ClazzID: 0x79107a0f} })     // 0x79107a0f
	iface.RegisterClazzID(0x851c6e34, func() iface.TLObject { return &TLInboxDeleteMessagesToInbox{ClazzID: 0x851c6e34} })      // 0x851c6e34
	iface.RegisterClazzID(0x140a8158, func() iface.TLObject { return &TLInboxDeleteUserHistoryToInbox{ClazzID: 0x140a8158} })   // 0x140a8158
	iface.RegisterClazzID(0xd8aaa602, func() iface.TLObject { return &TLInboxDeleteChatHistoryToInbox{ClazzID: 0xd8aaa602} })   // 0xd8aaa602
	iface.RegisterClazzID(0x15c1034b, func() iface.TLObject { return &TLInboxReadUserMediaUnreadToInbox{ClazzID: 0x15c1034b} }) // 0x15c1034b
	iface.RegisterClazzID(0x55415dd4, func() iface.TLObject { return &TLInboxReadChatMediaUnreadToInbox{ClazzID: 0x55415dd4} }) // 0x55415dd4
	iface.RegisterClazzID(0xc3c84ce0, func() iface.TLObject { return &TLInboxUpdateHistoryReaded{ClazzID: 0xc3c84ce0} })        // 0xc3c84ce0
	iface.RegisterClazzID(0xa96c2af4, func() iface.TLObject { return &TLInboxUpdatePinnedMessage{ClazzID: 0xa96c2af4} })        // 0xa96c2af4
	iface.RegisterClazzID(0x231ca261, func() iface.TLObject { return &TLInboxUnpinAllMessages{ClazzID: 0x231ca261} })           // 0x231ca261
	iface.RegisterClazzID(0x5bd7522, func() iface.TLObject { return &TLInboxSendUserMessageToInboxV2{ClazzID: 0x5bd7522} })     // 0x5bd7522
	iface.RegisterClazzID(0xdabb9e69, func() iface.TLObject { return &TLInboxEditMessageToInboxV2{ClazzID: 0xdabb9e69} })       // 0xdabb9e69
	iface.RegisterClazzID(0x1f73675, func() iface.TLObject { return &TLInboxReadInboxHistory{ClazzID: 0x1f73675} })             // 0x1f73675
	iface.RegisterClazzID(0x1c7036ca, func() iface.TLObject { return &TLInboxReadOutboxHistory{ClazzID: 0x1c7036ca} })          // 0x1c7036ca
	iface.RegisterClazzID(0xeac54342, func() iface.TLObject { return &TLInboxReadMediaUnreadToInboxV2{ClazzID: 0xeac54342} })   // 0xeac54342
}
