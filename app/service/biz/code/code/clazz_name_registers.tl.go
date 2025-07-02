/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package code

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_phoneCodeTransaction     = "phoneCodeTransaction"
	ClazzName_code_createPhoneCode     = "code_createPhoneCode"
	ClazzName_code_getPhoneCode        = "code_getPhoneCode"
	ClazzName_code_deletePhoneCode     = "code_deletePhoneCode"
	ClazzName_code_updatePhoneCodeData = "code_updatePhoneCodeData"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_phoneCodeTransaction, 0, 0x83739698)     // 83739698
	iface.RegisterClazzName(ClazzName_code_createPhoneCode, 0, 0x6023e09e)     // 6023e09e
	iface.RegisterClazzName(ClazzName_code_getPhoneCode, 0, 0x61a4a0f9)        // 61a4a0f9
	iface.RegisterClazzName(ClazzName_code_deletePhoneCode, 0, 0xa6b06a50)     // a6b06a50
	iface.RegisterClazzName(ClazzName_code_updatePhoneCodeData, 0, 0xb6950a95) // b6950a95

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_phoneCodeTransaction, 0x83739698)     // 83739698
	iface.RegisterClazzIDName(ClazzName_code_createPhoneCode, 0x6023e09e)     // 6023e09e
	iface.RegisterClazzIDName(ClazzName_code_getPhoneCode, 0x61a4a0f9)        // 61a4a0f9
	iface.RegisterClazzIDName(ClazzName_code_deletePhoneCode, 0xa6b06a50)     // a6b06a50
	iface.RegisterClazzIDName(ClazzName_code_updatePhoneCodeData, 0xb6950a95) // b6950a95
}
