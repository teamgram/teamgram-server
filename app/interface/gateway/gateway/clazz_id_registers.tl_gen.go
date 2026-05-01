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

package gateway

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

func init() {
	// Constructor

	// Method
	iface.RegisterClazzID(0x10dcca87, func() iface.TLObject { return &TLGatewayPushUpdatesData{ClazzID: 0x10dcca87} })        // 0x10dcca87
	iface.RegisterClazzID(0x794c7ded, func() iface.TLObject { return &TLGatewayPushSessionUpdatesData{ClazzID: 0x794c7ded} }) // 0x794c7ded
	iface.RegisterClazzID(0xfc960f5, func() iface.TLObject { return &TLGatewayPushRpcResultData{ClazzID: 0xfc960f5} })        // 0xfc960f5
}
