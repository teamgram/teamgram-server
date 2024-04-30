/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package username

const (
	Predicate_usernameNotExisted             = "usernameNotExisted"
	Predicate_usernameExisted                = "usernameExisted"
	Predicate_usernameExistedNotMe           = "usernameExistedNotMe"
	Predicate_usernameExistedIsMe            = "usernameExistedIsMe"
	Predicate_usernameData                   = "usernameData"
	Predicate_username_getAccountUsername    = "username_getAccountUsername"
	Predicate_username_checkAccountUsername  = "username_checkAccountUsername"
	Predicate_username_getChannelUsername    = "username_getChannelUsername"
	Predicate_username_checkChannelUsername  = "username_checkChannelUsername"
	Predicate_username_updateUsernameByPeer  = "username_updateUsernameByPeer"
	Predicate_username_checkUsername         = "username_checkUsername"
	Predicate_username_updateUsername        = "username_updateUsername"
	Predicate_username_deleteUsername        = "username_deleteUsername"
	Predicate_username_resolveUsername       = "username_resolveUsername"
	Predicate_username_getListByUsernameList = "username_getListByUsernameList"
	Predicate_username_deleteUsernameByPeer  = "username_deleteUsernameByPeer"
	Predicate_username_search                = "username_search"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_usernameNotExisted: {
		0: -885195923, // 0xcb3cfb6d

	},
	Predicate_usernameExisted: {
		0: -1394084659, // 0xace7f4cd

	},
	Predicate_usernameExistedNotMe: {
		0: -803256399, // 0xd01f47b1

	},
	Predicate_usernameExistedIsMe: {
		0: -2024900751, // 0x874e7771

	},
	Predicate_usernameData: {
		0: -1438646081, // 0xaa4000bf

	},
	Predicate_username_getAccountUsername: {
		0: 154073301, // 0x92ef8d5

	},
	Predicate_username_checkAccountUsername: {
		0: 1240985861, // 0x49f7f105

	},
	Predicate_username_getChannelUsername: {
		0: -2038134827, // 0x868487d5

	},
	Predicate_username_checkChannelUsername: {
		0: 651476637, // 0x26d4be9d

	},
	Predicate_username_updateUsernameByPeer: {
		0: 1718205916, // 0x6669bddc

	},
	Predicate_username_checkUsername: {
		0: 684369621, // 0x28caa6d5

	},
	Predicate_username_updateUsername: {
		0: 1389777971, // 0x52d65433

	},
	Predicate_username_deleteUsername: {
		0: -1065913464, // 0xc0777388

	},
	Predicate_username_resolveUsername: {
		0: 2008689862, // 0x77ba2cc6

	},
	Predicate_username_getListByUsernameList: {
		0: 1218942797, // 0x48a7974d

	},
	Predicate_username_deleteUsernameByPeer: {
		0: 507822189, // 0x1e44c06d

	},
	Predicate_username_search: {
		0: -391798010, // 0xe8a5a306

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	-885195923:  Predicate_usernameNotExisted,             // 0xcb3cfb6d
	-1394084659: Predicate_usernameExisted,                // 0xace7f4cd
	-803256399:  Predicate_usernameExistedNotMe,           // 0xd01f47b1
	-2024900751: Predicate_usernameExistedIsMe,            // 0x874e7771
	-1438646081: Predicate_usernameData,                   // 0xaa4000bf
	154073301:   Predicate_username_getAccountUsername,    // 0x92ef8d5
	1240985861:  Predicate_username_checkAccountUsername,  // 0x49f7f105
	-2038134827: Predicate_username_getChannelUsername,    // 0x868487d5
	651476637:   Predicate_username_checkChannelUsername,  // 0x26d4be9d
	1718205916:  Predicate_username_updateUsernameByPeer,  // 0x6669bddc
	684369621:   Predicate_username_checkUsername,         // 0x28caa6d5
	1389777971:  Predicate_username_updateUsername,        // 0x52d65433
	-1065913464: Predicate_username_deleteUsername,        // 0xc0777388
	2008689862:  Predicate_username_resolveUsername,       // 0x77ba2cc6
	1218942797:  Predicate_username_getListByUsernameList, // 0x48a7974d
	507822189:   Predicate_username_deleteUsernameByPeer,  // 0x1e44c06d
	-391798010:  Predicate_username_search,                // 0xe8a5a306

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
