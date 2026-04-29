/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package userupdates

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

const (
	ClazzName_userOperation                    = "userOperation"
	ClazzName_userOperationResult              = "userOperationResult"
	ClazzName_userState                        = "userState"
	ClazzName_userDifferenceEmpty              = "userDifferenceEmpty"
	ClazzName_userDifference                   = "userDifference"
	ClazzName_userDifferenceSlice              = "userDifferenceSlice"
	ClazzName_userDifferenceTooLong            = "userDifferenceTooLong"
	ClazzName_userupdates_processUserOperation = "userupdates_processUserOperation"
	ClazzName_userupdates_getOperationResult   = "userupdates_getOperationResult"
	ClazzName_userupdates_getState             = "userupdates_getState"
	ClazzName_userupdates_getDifference        = "userupdates_getDifference"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_userOperation, 0, 0x2d4e84d7)                    // 2d4e84d7
	iface.RegisterClazzName(ClazzName_userOperationResult, 0, 0x7311db72)              // 7311db72
	iface.RegisterClazzName(ClazzName_userState, 0, 0x635f3815)                        // 635f3815
	iface.RegisterClazzName(ClazzName_userDifferenceEmpty, 0, 0xb38ac177)              // b38ac177
	iface.RegisterClazzName(ClazzName_userDifference, 0, 0xb15cb08d)                   // b15cb08d
	iface.RegisterClazzName(ClazzName_userDifferenceSlice, 0, 0x4ef1987f)              // 4ef1987f
	iface.RegisterClazzName(ClazzName_userDifferenceTooLong, 0, 0x1d095703)            // 1d095703
	iface.RegisterClazzName(ClazzName_userupdates_processUserOperation, 0, 0xc200ea59) // c200ea59
	iface.RegisterClazzName(ClazzName_userupdates_getOperationResult, 0, 0x47a995d1)   // 47a995d1
	iface.RegisterClazzName(ClazzName_userupdates_getState, 0, 0x3bbbad80)             // 3bbbad80
	iface.RegisterClazzName(ClazzName_userupdates_getDifference, 0, 0x38cdd9fc)        // 38cdd9fc

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_userOperation, 0x2d4e84d7)                    // 2d4e84d7
	iface.RegisterClazzIDName(ClazzName_userOperationResult, 0x7311db72)              // 7311db72
	iface.RegisterClazzIDName(ClazzName_userState, 0x635f3815)                        // 635f3815
	iface.RegisterClazzIDName(ClazzName_userDifferenceEmpty, 0xb38ac177)              // b38ac177
	iface.RegisterClazzIDName(ClazzName_userDifference, 0xb15cb08d)                   // b15cb08d
	iface.RegisterClazzIDName(ClazzName_userDifferenceSlice, 0x4ef1987f)              // 4ef1987f
	iface.RegisterClazzIDName(ClazzName_userDifferenceTooLong, 0x1d095703)            // 1d095703
	iface.RegisterClazzIDName(ClazzName_userupdates_processUserOperation, 0xc200ea59) // c200ea59
	iface.RegisterClazzIDName(ClazzName_userupdates_getOperationResult, 0x47a995d1)   // 47a995d1
	iface.RegisterClazzIDName(ClazzName_userupdates_getState, 0x3bbbad80)             // 3bbbad80
	iface.RegisterClazzIDName(ClazzName_userupdates_getDifference, 0x38cdd9fc)        // 38cdd9fc
}
