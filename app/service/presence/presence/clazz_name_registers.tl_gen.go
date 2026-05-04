/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package presence

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

const (
	ClazzName_onlineSession                   = "onlineSession"
	ClazzName_userOnlineSessions              = "userOnlineSessions"
	ClazzName_presence_setSessionOnline       = "presence_setSessionOnline"
	ClazzName_presence_setSessionOffline      = "presence_setSessionOffline"
	ClazzName_presence_getUserOnlineSessions  = "presence_getUserOnlineSessions"
	ClazzName_presence_getUsersOnlineSessions = "presence_getUsersOnlineSessions"
	ClazzName_presence_getGatewaySessions     = "presence_getGatewaySessions"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_onlineSession, 0, 0x8a390a9f)                   // 8a390a9f
	iface.RegisterClazzName(ClazzName_userOnlineSessions, 0, 0x11eacb03)              // 11eacb03
	iface.RegisterClazzName(ClazzName_presence_setSessionOnline, 0, 0x75df4bac)       // 75df4bac
	iface.RegisterClazzName(ClazzName_presence_setSessionOffline, 0, 0x71ed7afb)      // 71ed7afb
	iface.RegisterClazzName(ClazzName_presence_getUserOnlineSessions, 0, 0x5aede88d)  // 5aede88d
	iface.RegisterClazzName(ClazzName_presence_getUsersOnlineSessions, 0, 0x1e5a09c5) // 1e5a09c5
	iface.RegisterClazzName(ClazzName_presence_getGatewaySessions, 0, 0xe5c814c)      // e5c814c

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_onlineSession, 0x8a390a9f)                   // 8a390a9f
	iface.RegisterClazzIDName(ClazzName_userOnlineSessions, 0x11eacb03)              // 11eacb03
	iface.RegisterClazzIDName(ClazzName_presence_setSessionOnline, 0x75df4bac)       // 75df4bac
	iface.RegisterClazzIDName(ClazzName_presence_setSessionOffline, 0x71ed7afb)      // 71ed7afb
	iface.RegisterClazzIDName(ClazzName_presence_getUserOnlineSessions, 0x5aede88d)  // 5aede88d
	iface.RegisterClazzIDName(ClazzName_presence_getUsersOnlineSessions, 0x1e5a09c5) // 1e5a09c5
	iface.RegisterClazzIDName(ClazzName_presence_getGatewaySessions, 0xe5c814c)      // e5c814c
}
