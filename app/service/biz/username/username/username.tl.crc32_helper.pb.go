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
	CRC32_UNKNOWN                        TLConstructor = 0
	CRC32_usernameNotExisted             TLConstructor = -885195923  // 0xcb3cfb6d
	CRC32_usernameExisted                TLConstructor = -1394084659 // 0xace7f4cd
	CRC32_usernameExistedNotMe           TLConstructor = -803256399  // 0xd01f47b1
	CRC32_usernameExistedIsMe            TLConstructor = -2024900751 // 0x874e7771
	CRC32_usernameData                   TLConstructor = -1438646081 // 0xaa4000bf
	CRC32_username_getAccountUsername    TLConstructor = 154073301   // 0x92ef8d5
	CRC32_username_checkAccountUsername  TLConstructor = 1240985861  // 0x49f7f105
	CRC32_username_getChannelUsername    TLConstructor = -2038134827 // 0x868487d5
	CRC32_username_checkChannelUsername  TLConstructor = 651476637   // 0x26d4be9d
	CRC32_username_updateUsernameByPeer  TLConstructor = 1718205916  // 0x6669bddc
	CRC32_username_checkUsername         TLConstructor = 684369621   // 0x28caa6d5
	CRC32_username_updateUsername        TLConstructor = 1389777971  // 0x52d65433
	CRC32_username_deleteUsername        TLConstructor = -1065913464 // 0xc0777388
	CRC32_username_resolveUsername       TLConstructor = 2008689862  // 0x77ba2cc6
	CRC32_username_getListByUsernameList TLConstructor = 1218942797  // 0x48a7974d
	CRC32_username_deleteUsernameByPeer  TLConstructor = 507822189   // 0x1e44c06d
	CRC32_username_search                TLConstructor = -391798010  // 0xe8a5a306
)
