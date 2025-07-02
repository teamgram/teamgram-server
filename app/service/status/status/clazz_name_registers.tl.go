/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package status

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_sessionEntry                      = "sessionEntry"
	ClazzName_userSessionEntryList              = "userSessionEntryList"
	ClazzName_status_setSessionOnline           = "status_setSessionOnline"
	ClazzName_status_setSessionOffline          = "status_setSessionOffline"
	ClazzName_status_getUserOnlineSessions      = "status_getUserOnlineSessions"
	ClazzName_status_getUsersOnlineSessionsList = "status_getUsersOnlineSessionsList"
	ClazzName_status_getChannelOnlineUsers      = "status_getChannelOnlineUsers"
	ClazzName_status_setUserChannelsOnline      = "status_setUserChannelsOnline"
	ClazzName_status_setUserChannelsOffline     = "status_setUserChannelsOffline"
	ClazzName_status_setChannelUserOffline      = "status_setChannelUserOffline"
	ClazzName_status_setChannelUsersOnline      = "status_setChannelUsersOnline"
	ClazzName_status_setChannelOffline          = "status_setChannelOffline"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_sessionEntry, 0, 0x1764ac31)                      // 1764ac31
	iface.RegisterClazzName(ClazzName_userSessionEntryList, 0, 0xefecb398)              // efecb398
	iface.RegisterClazzName(ClazzName_status_setSessionOnline, 0, 0x52518bcf)           // 52518bcf
	iface.RegisterClazzName(ClazzName_status_setSessionOffline, 0, 0x25a66a5c)          // 25a66a5c
	iface.RegisterClazzName(ClazzName_status_getUserOnlineSessions, 0, 0xe7c0e5cd)      // e7c0e5cd
	iface.RegisterClazzName(ClazzName_status_getUsersOnlineSessionsList, 0, 0x883b35c4) // 883b35c4
	iface.RegisterClazzName(ClazzName_status_getChannelOnlineUsers, 0, 0x4583ac55)      // 4583ac55
	iface.RegisterClazzName(ClazzName_status_setUserChannelsOnline, 0, 0xcd39044d)      // cd39044d
	iface.RegisterClazzName(ClazzName_status_setUserChannelsOffline, 0, 0x6ca361aa)     // 6ca361aa
	iface.RegisterClazzName(ClazzName_status_setChannelUserOffline, 0, 0xc48bcb7c)      // c48bcb7c
	iface.RegisterClazzName(ClazzName_status_setChannelUsersOnline, 0, 0xa69bdcf7)      // a69bdcf7
	iface.RegisterClazzName(ClazzName_status_setChannelOffline, 0, 0x4b7756f5)          // 4b7756f5

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_sessionEntry, 0x1764ac31)                      // 1764ac31
	iface.RegisterClazzIDName(ClazzName_userSessionEntryList, 0xefecb398)              // efecb398
	iface.RegisterClazzIDName(ClazzName_status_setSessionOnline, 0x52518bcf)           // 52518bcf
	iface.RegisterClazzIDName(ClazzName_status_setSessionOffline, 0x25a66a5c)          // 25a66a5c
	iface.RegisterClazzIDName(ClazzName_status_getUserOnlineSessions, 0xe7c0e5cd)      // e7c0e5cd
	iface.RegisterClazzIDName(ClazzName_status_getUsersOnlineSessionsList, 0x883b35c4) // 883b35c4
	iface.RegisterClazzIDName(ClazzName_status_getChannelOnlineUsers, 0x4583ac55)      // 4583ac55
	iface.RegisterClazzIDName(ClazzName_status_setUserChannelsOnline, 0xcd39044d)      // cd39044d
	iface.RegisterClazzIDName(ClazzName_status_setUserChannelsOffline, 0x6ca361aa)     // 6ca361aa
	iface.RegisterClazzIDName(ClazzName_status_setChannelUserOffline, 0xc48bcb7c)      // c48bcb7c
	iface.RegisterClazzIDName(ClazzName_status_setChannelUsersOnline, 0xa69bdcf7)      // a69bdcf7
	iface.RegisterClazzIDName(ClazzName_status_setChannelOffline, 0x4b7756f5)          // 4b7756f5
}
