/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package code

const (
	Predicate_phoneCodeTransaction     = "phoneCodeTransaction"
	Predicate_code_createPhoneCode     = "code_createPhoneCode"
	Predicate_code_getPhoneCode        = "code_getPhoneCode"
	Predicate_code_deletePhoneCode     = "code_deletePhoneCode"
	Predicate_code_updatePhoneCodeData = "code_updatePhoneCodeData"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_phoneCodeTransaction: {
		0: -2089576808, // 0x83739698

	},
	Predicate_code_createPhoneCode: {
		0: 1612963998, // 0x6023e09e

	},
	Predicate_code_getPhoneCode: {
		0: 1638179065, // 0x61a4a0f9

	},
	Predicate_code_deletePhoneCode: {
		0: -1498387888, // 0xa6b06a50

	},
	Predicate_code_updatePhoneCodeData: {
		0: -1231746411, // 0xb6950a95

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	-2089576808: Predicate_phoneCodeTransaction,     // 0x83739698
	1612963998:  Predicate_code_createPhoneCode,     // 0x6023e09e
	1638179065:  Predicate_code_getPhoneCode,        // 0x61a4a0f9
	-1498387888: Predicate_code_deletePhoneCode,     // 0xa6b06a50
	-1231746411: Predicate_code_updatePhoneCodeData, // 0xb6950a95

}

func GetClazzID(clazzName string, layer int) int32 {
	if m, ok := clazzNameRegisters2[clazzName]; ok {
		m2, ok2 := m[layer]
		if ok2 {
			return m2
		}
		m2, ok2 = m[0]
		if ok2 {
			return m2
		}
	}
	return 0
}
