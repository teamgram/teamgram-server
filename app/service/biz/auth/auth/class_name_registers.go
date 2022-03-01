/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teagramio (teagram.io@gmail.com)
 */

package auth

const (
	Predicate_auth_exportLoginToken = "auth_exportLoginToken"
	Predicate_auth_importLoginToken = "auth_importLoginToken"
	Predicate_auth_acceptLoginToken = "auth_acceptLoginToken"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_auth_exportLoginToken: {
		0: -1210022402, // 0xb7e085fe

	},
	Predicate_auth_importLoginToken: {
		0: -1783866140, // 0x95ac5ce4

	},
	Predicate_auth_acceptLoginToken: {
		0: -392909491, // 0xe894ad4d

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	-1210022402: Predicate_auth_exportLoginToken, // 0xb7e085fe
	-1783866140: Predicate_auth_importLoginToken, // 0x95ac5ce4
	-392909491:  Predicate_auth_acceptLoginToken, // 0xe894ad4d

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
