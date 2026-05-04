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

package sync

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

func init() {
	// Constructor

	// Method
	iface.RegisterClazzID(0x6d993b09, func() iface.TLObject { return &TLSyncUpdatesMe{ClazzID: 0x6d993b09} })        // 0x6d993b09
	iface.RegisterClazzID(0x97ac5031, func() iface.TLObject { return &TLSyncUpdatesNotMe{ClazzID: 0x97ac5031} })     // 0x97ac5031
	iface.RegisterClazzID(0x8f0ad9be, func() iface.TLObject { return &TLSyncPushUpdates{ClazzID: 0x8f0ad9be} })      // 0x8f0ad9be
	iface.RegisterClazzID(0x2d3778bc, func() iface.TLObject { return &TLSyncPushUpdatesIfNot{ClazzID: 0x2d3778bc} }) // 0x2d3778bc
	iface.RegisterClazzID(0x3fd7da47, func() iface.TLObject { return &TLSyncPushRpcResult{ClazzID: 0x3fd7da47} })    // 0x3fd7da47
}
