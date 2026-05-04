/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

// ConstructorList
// RequestList

package presence

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0x8a390a9f, func() iface.TLObject { return &TLOnlineSession{ClazzID: 0x8a390a9f} })      // 0x8a390a9f
	iface.RegisterClazzID(0x11eacb03, func() iface.TLObject { return &TLUserOnlineSessions{ClazzID: 0x11eacb03} }) // 0x11eacb03

	// Method
	iface.RegisterClazzID(0x75df4bac, func() iface.TLObject { return &TLPresenceSetSessionOnline{ClazzID: 0x75df4bac} })       // 0x75df4bac
	iface.RegisterClazzID(0x71ed7afb, func() iface.TLObject { return &TLPresenceSetSessionOffline{ClazzID: 0x71ed7afb} })      // 0x71ed7afb
	iface.RegisterClazzID(0x5aede88d, func() iface.TLObject { return &TLPresenceGetUserOnlineSessions{ClazzID: 0x5aede88d} })  // 0x5aede88d
	iface.RegisterClazzID(0x1e5a09c5, func() iface.TLObject { return &TLPresenceGetUsersOnlineSessions{ClazzID: 0x1e5a09c5} }) // 0x1e5a09c5
	iface.RegisterClazzID(0xe5c814c, func() iface.TLObject { return &TLPresenceGetGatewaySessions{ClazzID: 0xe5c814c} })       // 0xe5c814c
}
