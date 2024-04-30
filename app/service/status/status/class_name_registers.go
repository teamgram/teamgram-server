/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package status

const (
	Predicate_sessionEntry                      = "sessionEntry"
	Predicate_userSessionEntryList              = "userSessionEntryList"
	Predicate_status_setSessionOnline           = "status_setSessionOnline"
	Predicate_status_setSessionOffline          = "status_setSessionOffline"
	Predicate_status_getUserOnlineSessions      = "status_getUserOnlineSessions"
	Predicate_status_getUsersOnlineSessionsList = "status_getUsersOnlineSessionsList"
	Predicate_status_getChannelOnlineUsers      = "status_getChannelOnlineUsers"
	Predicate_status_setUserChannelsOnline      = "status_setUserChannelsOnline"
	Predicate_status_setUserChannelsOffline     = "status_setUserChannelsOffline"
	Predicate_status_setChannelUserOffline      = "status_setChannelUserOffline"
	Predicate_status_setChannelUsersOnline      = "status_setChannelUsersOnline"
	Predicate_status_setChannelOffline          = "status_setChannelOffline"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_sessionEntry: {
		0: 392473649, // 0x1764ac31

	},
	Predicate_userSessionEntryList: {
		0: -269700200, // 0xefecb398

	},
	Predicate_status_setSessionOnline: {
		0: 1381075919, // 0x52518bcf

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
	Predicate_status_getChannelOnlineUsers: {
		0: 1166257237, // 0x4583ac55

	},
	Predicate_status_setUserChannelsOnline: {
		0: -851901363, // 0xcd39044d

	},
	Predicate_status_setUserChannelsOffline: {
		0: 1822646698, // 0x6ca361aa

	},
	Predicate_status_setChannelUserOffline: {
		0: -997471364, // 0xc48bcb7c

	},
	Predicate_status_setChannelUsersOnline: {
		0: -1499734793, // 0xa69bdcf7

	},
	Predicate_status_setChannelOffline: {
		0: 1266112245, // 0x4b7756f5

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	392473649:   Predicate_sessionEntry,                      // 0x1764ac31
	-269700200:  Predicate_userSessionEntryList,              // 0xefecb398
	1381075919:  Predicate_status_setSessionOnline,           // 0x52518bcf
	631663196:   Predicate_status_setSessionOffline,          // 0x25a66a5c
	-406788659:  Predicate_status_getUserOnlineSessions,      // 0xe7c0e5cd
	-2009385532: Predicate_status_getUsersOnlineSessionsList, // 0x883b35c4
	1166257237:  Predicate_status_getChannelOnlineUsers,      // 0x4583ac55
	-851901363:  Predicate_status_setUserChannelsOnline,      // 0xcd39044d
	1822646698:  Predicate_status_setUserChannelsOffline,     // 0x6ca361aa
	-997471364:  Predicate_status_setChannelUserOffline,      // 0xc48bcb7c
	-1499734793: Predicate_status_setChannelUsersOnline,      // 0xa69bdcf7
	1266112245:  Predicate_status_setChannelOffline,          // 0x4b7756f5

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
