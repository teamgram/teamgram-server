/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package sync

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

const (
	ClazzName_sync_updatesMe        = "sync_updatesMe"
	ClazzName_sync_updatesNotMe     = "sync_updatesNotMe"
	ClazzName_sync_pushUpdates      = "sync_pushUpdates"
	ClazzName_sync_pushUpdatesIfNot = "sync_pushUpdatesIfNot"
	ClazzName_sync_pushRpcResult    = "sync_pushRpcResult"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_sync_updatesMe, 0, 0x6d993b09)        // 6d993b09
	iface.RegisterClazzName(ClazzName_sync_updatesNotMe, 0, 0x97ac5031)     // 97ac5031
	iface.RegisterClazzName(ClazzName_sync_pushUpdates, 0, 0x8f0ad9be)      // 8f0ad9be
	iface.RegisterClazzName(ClazzName_sync_pushUpdatesIfNot, 0, 0x2d3778bc) // 2d3778bc
	iface.RegisterClazzName(ClazzName_sync_pushRpcResult, 0, 0x3fd7da47)    // 3fd7da47

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_sync_updatesMe, 0x6d993b09)        // 6d993b09
	iface.RegisterClazzIDName(ClazzName_sync_updatesNotMe, 0x97ac5031)     // 97ac5031
	iface.RegisterClazzIDName(ClazzName_sync_pushUpdates, 0x8f0ad9be)      // 8f0ad9be
	iface.RegisterClazzIDName(ClazzName_sync_pushUpdatesIfNot, 0x2d3778bc) // 2d3778bc
	iface.RegisterClazzIDName(ClazzName_sync_pushRpcResult, 0x3fd7da47)    // 3fd7da47
}
