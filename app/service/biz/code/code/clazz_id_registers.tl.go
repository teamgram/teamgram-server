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

package code

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0x83739698, func() iface.TLObject { return &TLPhoneCodeTransaction{ClazzID: 0x83739698} }) // 0x83739698

	// Method
	iface.RegisterClazzID(0x6023e09e, func() iface.TLObject { return &TLCodeCreatePhoneCode{ClazzID: 0x6023e09e} })     // 0x6023e09e
	iface.RegisterClazzID(0x61a4a0f9, func() iface.TLObject { return &TLCodeGetPhoneCode{ClazzID: 0x61a4a0f9} })        // 0x61a4a0f9
	iface.RegisterClazzID(0xa6b06a50, func() iface.TLObject { return &TLCodeDeletePhoneCode{ClazzID: 0xa6b06a50} })     // 0xa6b06a50
	iface.RegisterClazzID(0xb6950a95, func() iface.TLObject { return &TLCodeUpdatePhoneCodeData{ClazzID: 0xb6950a95} }) // 0xb6950a95
}
