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

package echo1

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0x2e3ba51e, func() iface.TLObject { return &TLEcho{ClazzID: 0x2e3ba51e} }) // 0x2e3ba51e

	// Method
	iface.RegisterClazzID(0x9f0506e2, func() iface.TLObject { return &TLEcho1Echo{ClazzID: 0x9f0506e2} }) // 0x9f0506e2
}
