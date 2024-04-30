/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package sync

const (
	CRC32_UNKNOWN               TLConstructor = 0
	CRC32_sync_updatesMe        TLConstructor = 1614568688  // 0x603c5cf0
	CRC32_sync_updatesNotMe     TLConstructor = 16458447    // 0xfb22cf
	CRC32_sync_pushUpdates      TLConstructor = -1895114306 // 0x8f0ad9be
	CRC32_sync_pushUpdatesIfNot TLConstructor = 1074085860  // 0x40053fe4
	CRC32_sync_pushBotUpdates   TLConstructor = -1379667968 // 0xadc3f000
	CRC32_sync_pushRpcResult    TLConstructor = -1874085983 // 0x904bb7a1
	CRC32_sync_broadcastUpdates TLConstructor = -169648970  // 0xf5e35cb6
)
