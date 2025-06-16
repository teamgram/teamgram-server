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

package sync

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor

	// Method
	iface.RegisterClazzID(0xe57d411f, func() iface.TLObject { return &TLSyncUpdatesMe{ClazzID: 0xe57d411f} })        // 0xe57d411f
	iface.RegisterClazzID(0x97ac5031, func() iface.TLObject { return &TLSyncUpdatesNotMe{ClazzID: 0x97ac5031} })     // 0x97ac5031
	iface.RegisterClazzID(0x8f0ad9be, func() iface.TLObject { return &TLSyncPushUpdates{ClazzID: 0x8f0ad9be} })      // 0x8f0ad9be
	iface.RegisterClazzID(0x40053fe4, func() iface.TLObject { return &TLSyncPushUpdatesIfNot{ClazzID: 0x40053fe4} }) // 0x40053fe4
	iface.RegisterClazzID(0xadc3f000, func() iface.TLObject { return &TLSyncPushBotUpdates{ClazzID: 0xadc3f000} })   // 0xadc3f000
	iface.RegisterClazzID(0x1a9d4b2, func() iface.TLObject { return &TLSyncPushRpcResult{ClazzID: 0x1a9d4b2} })      // 0x1a9d4b2
	iface.RegisterClazzID(0xf5e35cb6, func() iface.TLObject { return &TLSyncBroadcastUpdates{ClazzID: 0xf5e35cb6} }) // 0xf5e35cb6
}
