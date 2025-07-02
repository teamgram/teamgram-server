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

package updates

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0xcd19034a, func() iface.TLObject { return &TLChannelDifference{ClazzID: 0xcd19034a} }) // 0xcd19034a
	iface.RegisterClazzID(0x8bdbda4e, func() iface.TLObject { return &TLDifferenceEmpty{ClazzID: 0x8bdbda4e} })   // 0x8bdbda4e
	iface.RegisterClazzID(0x5482832b, func() iface.TLObject { return &TLDifference{ClazzID: 0x5482832b} })        // 0x5482832b
	iface.RegisterClazzID(0xcb965ddf, func() iface.TLObject { return &TLDifferenceSlice{ClazzID: 0xcb965ddf} })   // 0xcb965ddf
	iface.RegisterClazzID(0x3572ee30, func() iface.TLObject { return &TLDifferenceTooLong{ClazzID: 0x3572ee30} }) // 0x3572ee30

	// Method
	iface.RegisterClazzID(0x45f4cd65, func() iface.TLObject { return &TLUpdatesGetStateV2{ClazzID: 0x45f4cd65} })             // 0x45f4cd65
	iface.RegisterClazzID(0xb76b6699, func() iface.TLObject { return &TLUpdatesGetDifferenceV2{ClazzID: 0xb76b6699} })        // 0xb76b6699
	iface.RegisterClazzID(0x4da3318a, func() iface.TLObject { return &TLUpdatesGetChannelDifferenceV2{ClazzID: 0x4da3318a} }) // 0x4da3318a
}
