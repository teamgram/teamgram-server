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

package userupdates

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0xb38ac177, func() iface.TLObject { return &TLUserDifferenceEmpty{ClazzID: 0xb38ac177} })   // 0xb38ac177
	iface.RegisterClazzID(0xb15cb08d, func() iface.TLObject { return &TLUserDifference{ClazzID: 0xb15cb08d} })        // 0xb15cb08d
	iface.RegisterClazzID(0x4ef1987f, func() iface.TLObject { return &TLUserDifferenceSlice{ClazzID: 0x4ef1987f} })   // 0x4ef1987f
	iface.RegisterClazzID(0x1d095703, func() iface.TLObject { return &TLUserDifferenceTooLong{ClazzID: 0x1d095703} }) // 0x1d095703
	iface.RegisterClazzID(0x2d4e84d7, func() iface.TLObject { return &TLUserOperation{ClazzID: 0x2d4e84d7} })         // 0x2d4e84d7
	iface.RegisterClazzID(0x7311db72, func() iface.TLObject { return &TLUserOperationResult{ClazzID: 0x7311db72} })   // 0x7311db72
	iface.RegisterClazzID(0x635f3815, func() iface.TLObject { return &TLUserState{ClazzID: 0x635f3815} })             // 0x635f3815

	// Method
	iface.RegisterClazzID(0xc200ea59, func() iface.TLObject { return &TLUserupdatesProcessUserOperation{ClazzID: 0xc200ea59} }) // 0xc200ea59
	iface.RegisterClazzID(0x47a995d1, func() iface.TLObject { return &TLUserupdatesGetOperationResult{ClazzID: 0x47a995d1} })   // 0x47a995d1
	iface.RegisterClazzID(0x3bbbad80, func() iface.TLObject { return &TLUserupdatesGetState{ClazzID: 0x3bbbad80} })             // 0x3bbbad80
	iface.RegisterClazzID(0x38cdd9fc, func() iface.TLObject { return &TLUserupdatesGetDifference{ClazzID: 0x38cdd9fc} })        // 0x38cdd9fc
}
