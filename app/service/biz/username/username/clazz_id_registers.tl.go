/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

// ConstructorList
// RequestList

package username

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0xaa4000bf, func() iface.TLObject { return &TLUsernameData{ClazzID: 0xaa4000bf} })         // 0xaa4000bf
	iface.RegisterClazzID(0xcb3cfb6d, func() iface.TLObject { return &TLUsernameNotExisted{ClazzID: 0xcb3cfb6d} })   // 0xcb3cfb6d
	iface.RegisterClazzID(0xace7f4cd, func() iface.TLObject { return &TLUsernameExisted{ClazzID: 0xace7f4cd} })      // 0xace7f4cd
	iface.RegisterClazzID(0xd01f47b1, func() iface.TLObject { return &TLUsernameExistedNotMe{ClazzID: 0xd01f47b1} }) // 0xd01f47b1
	iface.RegisterClazzID(0x874e7771, func() iface.TLObject { return &TLUsernameExistedIsMe{ClazzID: 0x874e7771} })  // 0x874e7771

	// Method
	iface.RegisterClazzID(0x92ef8d5, func() iface.TLObject { return &TLUsernameGetAccountUsername{ClazzID: 0x92ef8d5} })      // 0x92ef8d5
	iface.RegisterClazzID(0x49f7f105, func() iface.TLObject { return &TLUsernameCheckAccountUsername{ClazzID: 0x49f7f105} })  // 0x49f7f105
	iface.RegisterClazzID(0x868487d5, func() iface.TLObject { return &TLUsernameGetChannelUsername{ClazzID: 0x868487d5} })    // 0x868487d5
	iface.RegisterClazzID(0x26d4be9d, func() iface.TLObject { return &TLUsernameCheckChannelUsername{ClazzID: 0x26d4be9d} })  // 0x26d4be9d
	iface.RegisterClazzID(0x6669bddc, func() iface.TLObject { return &TLUsernameUpdateUsernameByPeer{ClazzID: 0x6669bddc} })  // 0x6669bddc
	iface.RegisterClazzID(0x28caa6d5, func() iface.TLObject { return &TLUsernameCheckUsername{ClazzID: 0x28caa6d5} })         // 0x28caa6d5
	iface.RegisterClazzID(0x52d65433, func() iface.TLObject { return &TLUsernameUpdateUsername{ClazzID: 0x52d65433} })        // 0x52d65433
	iface.RegisterClazzID(0xc0777388, func() iface.TLObject { return &TLUsernameDeleteUsername{ClazzID: 0xc0777388} })        // 0xc0777388
	iface.RegisterClazzID(0x77ba2cc6, func() iface.TLObject { return &TLUsernameResolveUsername{ClazzID: 0x77ba2cc6} })       // 0x77ba2cc6
	iface.RegisterClazzID(0x48a7974d, func() iface.TLObject { return &TLUsernameGetListByUsernameList{ClazzID: 0x48a7974d} }) // 0x48a7974d
	iface.RegisterClazzID(0x1e44c06d, func() iface.TLObject { return &TLUsernameDeleteUsernameByPeer{ClazzID: 0x1e44c06d} })  // 0x1e44c06d
	iface.RegisterClazzID(0xe8a5a306, func() iface.TLObject { return &TLUsernameSearch{ClazzID: 0xe8a5a306} })                // 0xe8a5a306
}
