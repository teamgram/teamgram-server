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

package echo

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0x2e3ba51e, func() iface.TLObject { return &TLEcho{ClazzID: 0x2e3ba51e} }) // 0x2e3ba51e
	iface.RegisterClazzID(0x2249c1b, func() iface.TLObject { return &TLEcho2{ClazzID: 0x2249c1b} })  // 0x2249c1b

	// Method
	iface.RegisterClazzID(0xf653b67d, func() iface.TLObject { return &TLEchoEcho{ClazzID: 0xf653b67d} }) // 0xf653b67d
}
