/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package username

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_usernameNotExisted             = "usernameNotExisted"
	ClazzName_usernameExisted                = "usernameExisted"
	ClazzName_usernameExistedNotMe           = "usernameExistedNotMe"
	ClazzName_usernameExistedIsMe            = "usernameExistedIsMe"
	ClazzName_usernameData                   = "usernameData"
	ClazzName_username_getAccountUsername    = "username_getAccountUsername"
	ClazzName_username_checkAccountUsername  = "username_checkAccountUsername"
	ClazzName_username_getChannelUsername    = "username_getChannelUsername"
	ClazzName_username_checkChannelUsername  = "username_checkChannelUsername"
	ClazzName_username_updateUsernameByPeer  = "username_updateUsernameByPeer"
	ClazzName_username_checkUsername         = "username_checkUsername"
	ClazzName_username_updateUsername        = "username_updateUsername"
	ClazzName_username_deleteUsername        = "username_deleteUsername"
	ClazzName_username_resolveUsername       = "username_resolveUsername"
	ClazzName_username_getListByUsernameList = "username_getListByUsernameList"
	ClazzName_username_deleteUsernameByPeer  = "username_deleteUsernameByPeer"
	ClazzName_username_search                = "username_search"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_usernameNotExisted, 0, 0xcb3cfb6d)             // cb3cfb6d
	iface.RegisterClazzName(ClazzName_usernameExisted, 0, 0xace7f4cd)                // ace7f4cd
	iface.RegisterClazzName(ClazzName_usernameExistedNotMe, 0, 0xd01f47b1)           // d01f47b1
	iface.RegisterClazzName(ClazzName_usernameExistedIsMe, 0, 0x874e7771)            // 874e7771
	iface.RegisterClazzName(ClazzName_usernameData, 0, 0xaa4000bf)                   // aa4000bf
	iface.RegisterClazzName(ClazzName_username_getAccountUsername, 0, 0x92ef8d5)     // 92ef8d5
	iface.RegisterClazzName(ClazzName_username_checkAccountUsername, 0, 0x49f7f105)  // 49f7f105
	iface.RegisterClazzName(ClazzName_username_getChannelUsername, 0, 0x868487d5)    // 868487d5
	iface.RegisterClazzName(ClazzName_username_checkChannelUsername, 0, 0x26d4be9d)  // 26d4be9d
	iface.RegisterClazzName(ClazzName_username_updateUsernameByPeer, 0, 0x6669bddc)  // 6669bddc
	iface.RegisterClazzName(ClazzName_username_checkUsername, 0, 0x28caa6d5)         // 28caa6d5
	iface.RegisterClazzName(ClazzName_username_updateUsername, 0, 0x52d65433)        // 52d65433
	iface.RegisterClazzName(ClazzName_username_deleteUsername, 0, 0xc0777388)        // c0777388
	iface.RegisterClazzName(ClazzName_username_resolveUsername, 0, 0x77ba2cc6)       // 77ba2cc6
	iface.RegisterClazzName(ClazzName_username_getListByUsernameList, 0, 0x48a7974d) // 48a7974d
	iface.RegisterClazzName(ClazzName_username_deleteUsernameByPeer, 0, 0x1e44c06d)  // 1e44c06d
	iface.RegisterClazzName(ClazzName_username_search, 0, 0xe8a5a306)                // e8a5a306

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_usernameNotExisted, 0xcb3cfb6d)             // cb3cfb6d
	iface.RegisterClazzIDName(ClazzName_usernameExisted, 0xace7f4cd)                // ace7f4cd
	iface.RegisterClazzIDName(ClazzName_usernameExistedNotMe, 0xd01f47b1)           // d01f47b1
	iface.RegisterClazzIDName(ClazzName_usernameExistedIsMe, 0x874e7771)            // 874e7771
	iface.RegisterClazzIDName(ClazzName_usernameData, 0xaa4000bf)                   // aa4000bf
	iface.RegisterClazzIDName(ClazzName_username_getAccountUsername, 0x92ef8d5)     // 92ef8d5
	iface.RegisterClazzIDName(ClazzName_username_checkAccountUsername, 0x49f7f105)  // 49f7f105
	iface.RegisterClazzIDName(ClazzName_username_getChannelUsername, 0x868487d5)    // 868487d5
	iface.RegisterClazzIDName(ClazzName_username_checkChannelUsername, 0x26d4be9d)  // 26d4be9d
	iface.RegisterClazzIDName(ClazzName_username_updateUsernameByPeer, 0x6669bddc)  // 6669bddc
	iface.RegisterClazzIDName(ClazzName_username_checkUsername, 0x28caa6d5)         // 28caa6d5
	iface.RegisterClazzIDName(ClazzName_username_updateUsername, 0x52d65433)        // 52d65433
	iface.RegisterClazzIDName(ClazzName_username_deleteUsername, 0xc0777388)        // c0777388
	iface.RegisterClazzIDName(ClazzName_username_resolveUsername, 0x77ba2cc6)       // 77ba2cc6
	iface.RegisterClazzIDName(ClazzName_username_getListByUsernameList, 0x48a7974d) // 48a7974d
	iface.RegisterClazzIDName(ClazzName_username_deleteUsernameByPeer, 0x1e44c06d)  // 1e44c06d
	iface.RegisterClazzIDName(ClazzName_username_search, 0xe8a5a306)                // e8a5a306
}
