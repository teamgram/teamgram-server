/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teagramio (teagram.io@gmail.com)
 */

package sync

const (
	Predicate_sync_updatesMe        = "sync_updatesMe"
	Predicate_sync_updatesNotMe     = "sync_updatesNotMe"
	Predicate_sync_pushUpdates      = "sync_pushUpdates"
	Predicate_sync_pushUpdatesIfNot = "sync_pushUpdatesIfNot"
	Predicate_sync_pushBotUpdates   = "sync_pushBotUpdates"
	Predicate_sync_pushRpcResult    = "sync_pushRpcResult"
	Predicate_sync_broadcastUpdates = "sync_broadcastUpdates"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_sync_updatesMe: {
		0: 1614568688, // 0x603c5cf0

	},
	Predicate_sync_updatesNotMe: {
		0: 16458447, // 0xfb22cf

	},
	Predicate_sync_pushUpdates: {
		0: -1895114306, // 0x8f0ad9be

	},
	Predicate_sync_pushUpdatesIfNot: {
		0: 1074085860, // 0x40053fe4

	},
	Predicate_sync_pushBotUpdates: {
		0: -1379667968, // 0xadc3f000

	},
	Predicate_sync_pushRpcResult: {
		0: -1874085983, // 0x904bb7a1

	},
	Predicate_sync_broadcastUpdates: {
		0: -169648970, // 0xf5e35cb6

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	1614568688:  Predicate_sync_updatesMe,        // 0x603c5cf0
	16458447:    Predicate_sync_updatesNotMe,     // 0xfb22cf
	-1895114306: Predicate_sync_pushUpdates,      // 0x8f0ad9be
	1074085860:  Predicate_sync_pushUpdatesIfNot, // 0x40053fe4
	-1379667968: Predicate_sync_pushBotUpdates,   // 0xadc3f000
	-1874085983: Predicate_sync_pushRpcResult,    // 0x904bb7a1
	-169648970:  Predicate_sync_broadcastUpdates, // 0xf5e35cb6

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
