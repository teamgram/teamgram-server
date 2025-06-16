/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package session

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_sessionClientEvent             = "sessionClientEvent"
	ClazzName_sessionClientData              = "sessionClientData"
	ClazzName_httpSessionData                = "httpSessionData"
	ClazzName_session_queryAuthKey           = "session_queryAuthKey"
	ClazzName_session_setAuthKey             = "session_setAuthKey"
	ClazzName_session_createSession          = "session_createSession"
	ClazzName_session_sendDataToSession      = "session_sendDataToSession"
	ClazzName_session_sendHttpDataToSession  = "session_sendHttpDataToSession"
	ClazzName_session_closeSession           = "session_closeSession"
	ClazzName_session_pushUpdatesData        = "session_pushUpdatesData"
	ClazzName_session_pushSessionUpdatesData = "session_pushSessionUpdatesData"
	ClazzName_session_pushRpcResultData      = "session_pushRpcResultData"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_sessionClientEvent, 0, 0xf17f375f)             // f17f375f
	iface.RegisterClazzName(ClazzName_sessionClientData, 0, 0x41a20c4e)              // 41a20c4e
	iface.RegisterClazzName(ClazzName_httpSessionData, 0, 0xdbd8534f)                // dbd8534f
	iface.RegisterClazzName(ClazzName_session_queryAuthKey, 0, 0x6b2df851)           // 6b2df851
	iface.RegisterClazzName(ClazzName_session_setAuthKey, 0, 0x1d11490b)             // 1d11490b
	iface.RegisterClazzName(ClazzName_session_createSession, 0, 0x410cb20d)          // 410cb20d
	iface.RegisterClazzName(ClazzName_session_sendDataToSession, 0, 0x876b2dec)      // 876b2dec
	iface.RegisterClazzName(ClazzName_session_sendHttpDataToSession, 0, 0xbbec23ae)  // bbec23ae
	iface.RegisterClazzName(ClazzName_session_closeSession, 0, 0x176fc253)           // 176fc253
	iface.RegisterClazzName(ClazzName_session_pushUpdatesData, 0, 0xa574d829)        // a574d829
	iface.RegisterClazzName(ClazzName_session_pushSessionUpdatesData, 0, 0x45f3fda0) // 45f3fda0
	iface.RegisterClazzName(ClazzName_session_pushRpcResultData, 0, 0x4b470c89)      // 4b470c89

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_sessionClientEvent, 0xf17f375f)             // f17f375f
	iface.RegisterClazzIDName(ClazzName_sessionClientData, 0x41a20c4e)              // 41a20c4e
	iface.RegisterClazzIDName(ClazzName_httpSessionData, 0xdbd8534f)                // dbd8534f
	iface.RegisterClazzIDName(ClazzName_session_queryAuthKey, 0x6b2df851)           // 6b2df851
	iface.RegisterClazzIDName(ClazzName_session_setAuthKey, 0x1d11490b)             // 1d11490b
	iface.RegisterClazzIDName(ClazzName_session_createSession, 0x410cb20d)          // 410cb20d
	iface.RegisterClazzIDName(ClazzName_session_sendDataToSession, 0x876b2dec)      // 876b2dec
	iface.RegisterClazzIDName(ClazzName_session_sendHttpDataToSession, 0xbbec23ae)  // bbec23ae
	iface.RegisterClazzIDName(ClazzName_session_closeSession, 0x176fc253)           // 176fc253
	iface.RegisterClazzIDName(ClazzName_session_pushUpdatesData, 0xa574d829)        // a574d829
	iface.RegisterClazzIDName(ClazzName_session_pushSessionUpdatesData, 0x45f3fda0) // 45f3fda0
	iface.RegisterClazzIDName(ClazzName_session_pushRpcResultData, 0x4b470c89)      // 4b470c89
}
