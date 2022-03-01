/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teagramio (teagram.io@gmail.com)
 */

package status

const (
	Predicate_sessionEntry                      = "sessionEntry"
	Predicate_userSessionEntryList              = "userSessionEntryList"
	Predicate_status_setSessionOnline           = "status_setSessionOnline"
	Predicate_status_setSessionOffline          = "status_setSessionOffline"
	Predicate_status_getUserOnlineSessions      = "status_getUserOnlineSessions"
	Predicate_status_getUsersOnlineSessionsList = "status_getUsersOnlineSessionsList"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_sessionEntry: {
		0: -1409734405, // 0xabf928fb

	},
	Predicate_userSessionEntryList: {
		0: -269700200, // 0xefecb398

	},
	Predicate_status_setSessionOnline: {
		0: -535445567, // 0xe015bfc1

	},
	Predicate_status_setSessionOffline: {
		0: 631663196, // 0x25a66a5c

	},
	Predicate_status_getUserOnlineSessions: {
		0: -406788659, // 0xe7c0e5cd

	},
	Predicate_status_getUsersOnlineSessionsList: {
		0: -2009385532, // 0x883b35c4

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	-1409734405: Predicate_sessionEntry,                      // 0xabf928fb
	-269700200:  Predicate_userSessionEntryList,              // 0xefecb398
	-535445567:  Predicate_status_setSessionOnline,           // 0xe015bfc1
	631663196:   Predicate_status_setSessionOffline,          // 0x25a66a5c
	-406788659:  Predicate_status_getUserOnlineSessions,      // 0xe7c0e5cd
	-2009385532: Predicate_status_getUsersOnlineSessionsList, // 0x883b35c4

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
