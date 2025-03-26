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

package gnetway

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor

	// Method
	iface.RegisterClazzID(0x722d5ce0, func() iface.TLObject { return &TLGnetwaySendDataToGateway{ClazzID: 0x722d5ce0} }) // 0x722d5ce0
}
