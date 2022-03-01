/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teagramio (teagram.io@gmail.com)
 */

package banned

const (
	Predicate_banned_checkPhoneNumberBanned = "banned_checkPhoneNumberBanned"
	Predicate_banned_getBannedByPhoneList   = "banned_getBannedByPhoneList"
	Predicate_banned_ban                    = "banned_ban"
	Predicate_banned_unBan                  = "banned_unBan"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_banned_checkPhoneNumberBanned: {
		0: -515891261, // 0xe1401fc3

	},
	Predicate_banned_getBannedByPhoneList: {
		0: -453047268, // 0xe4ff0c1c

	},
	Predicate_banned_ban: {
		0: 1037800024, // 0x3ddb9258

	},
	Predicate_banned_unBan: {
		0: 1912613348, // 0x720029e4

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	-515891261: Predicate_banned_checkPhoneNumberBanned, // 0xe1401fc3
	-453047268: Predicate_banned_getBannedByPhoneList,   // 0xe4ff0c1c
	1037800024: Predicate_banned_ban,                    // 0x3ddb9258
	1912613348: Predicate_banned_unBan,                  // 0x720029e4

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
