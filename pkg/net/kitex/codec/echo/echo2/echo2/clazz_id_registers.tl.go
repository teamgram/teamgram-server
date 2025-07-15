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

package echo2

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0x2e3ba51e, func() iface.TLObject { return &TLEcho{ClazzID: 0x2e3ba51e} }) // 0x2e3ba51e

	// Method
	iface.RegisterClazzID(0x9ddb01c5, func() iface.TLObject { return &TLEcho2Echo{ClazzID: 0x9ddb01c5} }) // 0x9ddb01c5
}
