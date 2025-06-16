/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

// ConstructorList
// RequestList

package status

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0x1764ac31, func() iface.TLObject { return &TLSessionEntry{ClazzID: 0x1764ac31} })         // 0x1764ac31
	iface.RegisterClazzID(0xefecb398, func() iface.TLObject { return &TLUserSessionEntryList{ClazzID: 0xefecb398} }) // 0xefecb398

	// Method
	iface.RegisterClazzID(0x52518bcf, func() iface.TLObject { return &TLStatusSetSessionOnline{ClazzID: 0x52518bcf} })           // 0x52518bcf
	iface.RegisterClazzID(0x25a66a5c, func() iface.TLObject { return &TLStatusSetSessionOffline{ClazzID: 0x25a66a5c} })          // 0x25a66a5c
	iface.RegisterClazzID(0xe7c0e5cd, func() iface.TLObject { return &TLStatusGetUserOnlineSessions{ClazzID: 0xe7c0e5cd} })      // 0xe7c0e5cd
	iface.RegisterClazzID(0x883b35c4, func() iface.TLObject { return &TLStatusGetUsersOnlineSessionsList{ClazzID: 0x883b35c4} }) // 0x883b35c4
	iface.RegisterClazzID(0x4583ac55, func() iface.TLObject { return &TLStatusGetChannelOnlineUsers{ClazzID: 0x4583ac55} })      // 0x4583ac55
	iface.RegisterClazzID(0xcd39044d, func() iface.TLObject { return &TLStatusSetUserChannelsOnline{ClazzID: 0xcd39044d} })      // 0xcd39044d
	iface.RegisterClazzID(0x6ca361aa, func() iface.TLObject { return &TLStatusSetUserChannelsOffline{ClazzID: 0x6ca361aa} })     // 0x6ca361aa
	iface.RegisterClazzID(0xc48bcb7c, func() iface.TLObject { return &TLStatusSetChannelUserOffline{ClazzID: 0xc48bcb7c} })      // 0xc48bcb7c
	iface.RegisterClazzID(0xa69bdcf7, func() iface.TLObject { return &TLStatusSetChannelUsersOnline{ClazzID: 0xa69bdcf7} })      // 0xa69bdcf7
	iface.RegisterClazzID(0x4b7756f5, func() iface.TLObject { return &TLStatusSetChannelOffline{ClazzID: 0x4b7756f5} })          // 0x4b7756f5
}
