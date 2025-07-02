/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package sync

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_sync_updatesMe        = "sync_updatesMe"
	ClazzName_sync_updatesNotMe     = "sync_updatesNotMe"
	ClazzName_sync_pushUpdates      = "sync_pushUpdates"
	ClazzName_sync_pushUpdatesIfNot = "sync_pushUpdatesIfNot"
	ClazzName_sync_pushBotUpdates   = "sync_pushBotUpdates"
	ClazzName_sync_pushRpcResult    = "sync_pushRpcResult"
	ClazzName_sync_broadcastUpdates = "sync_broadcastUpdates"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_sync_updatesMe, 0, 0xe57d411f)        // e57d411f
	iface.RegisterClazzName(ClazzName_sync_updatesNotMe, 0, 0x97ac5031)     // 97ac5031
	iface.RegisterClazzName(ClazzName_sync_pushUpdates, 0, 0x8f0ad9be)      // 8f0ad9be
	iface.RegisterClazzName(ClazzName_sync_pushUpdatesIfNot, 0, 0x2d3778bc) // 2d3778bc
	iface.RegisterClazzName(ClazzName_sync_pushBotUpdates, 0, 0xadc3f000)   // adc3f000
	iface.RegisterClazzName(ClazzName_sync_pushRpcResult, 0, 0x1a9d4b2)     // 1a9d4b2
	iface.RegisterClazzName(ClazzName_sync_broadcastUpdates, 0, 0xf5e35cb6) // f5e35cb6

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_sync_updatesMe, 0xe57d411f)        // e57d411f
	iface.RegisterClazzIDName(ClazzName_sync_updatesNotMe, 0x97ac5031)     // 97ac5031
	iface.RegisterClazzIDName(ClazzName_sync_pushUpdates, 0x8f0ad9be)      // 8f0ad9be
	iface.RegisterClazzIDName(ClazzName_sync_pushUpdatesIfNot, 0x2d3778bc) // 2d3778bc
	iface.RegisterClazzIDName(ClazzName_sync_pushBotUpdates, 0xadc3f000)   // adc3f000
	iface.RegisterClazzIDName(ClazzName_sync_pushRpcResult, 0x1a9d4b2)     // 1a9d4b2
	iface.RegisterClazzIDName(ClazzName_sync_broadcastUpdates, 0xf5e35cb6) // f5e35cb6
}
