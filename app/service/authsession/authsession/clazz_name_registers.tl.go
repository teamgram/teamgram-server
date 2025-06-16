/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package authsession

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_clientSession                       = "clientSession"
	ClazzName_authKeyStateData                    = "authKeyStateData"
	ClazzName_authsession_getAuthorizations       = "authsession_getAuthorizations"
	ClazzName_authsession_resetAuthorization      = "authsession_resetAuthorization"
	ClazzName_authsession_getLayer                = "authsession_getLayer"
	ClazzName_authsession_getLangPack             = "authsession_getLangPack"
	ClazzName_authsession_getClient               = "authsession_getClient"
	ClazzName_authsession_getLangCode             = "authsession_getLangCode"
	ClazzName_authsession_getUserId               = "authsession_getUserId"
	ClazzName_authsession_getPushSessionId        = "authsession_getPushSessionId"
	ClazzName_authsession_getFutureSalts          = "authsession_getFutureSalts"
	ClazzName_authsession_queryAuthKey            = "authsession_queryAuthKey"
	ClazzName_authsession_setAuthKey              = "authsession_setAuthKey"
	ClazzName_authsession_bindAuthKeyUser         = "authsession_bindAuthKeyUser"
	ClazzName_authsession_unbindAuthKeyUser       = "authsession_unbindAuthKeyUser"
	ClazzName_authsession_getPermAuthKeyId        = "authsession_getPermAuthKeyId"
	ClazzName_authsession_bindTempAuthKey         = "authsession_bindTempAuthKey"
	ClazzName_authsession_setClientSessionInfo    = "authsession_setClientSessionInfo"
	ClazzName_authsession_getAuthorization        = "authsession_getAuthorization"
	ClazzName_authsession_getAuthStateData        = "authsession_getAuthStateData"
	ClazzName_authsession_setLayer                = "authsession_setLayer"
	ClazzName_authsession_setInitConnection       = "authsession_setInitConnection"
	ClazzName_authsession_setAndroidPushSessionId = "authsession_setAndroidPushSessionId"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_clientSession, 0, 0x9a8e71b0)                       // 9a8e71b0
	iface.RegisterClazzName(ClazzName_authKeyStateData, 0, 0xe0408f17)                    // e0408f17
	iface.RegisterClazzName(ClazzName_authsession_getAuthorizations, 0, 0x30e21244)       // 30e21244
	iface.RegisterClazzName(ClazzName_authsession_resetAuthorization, 0, 0x8d5f6ca6)      // 8d5f6ca6
	iface.RegisterClazzName(ClazzName_authsession_getLayer, 0, 0xa82f16a9)                // a82f16a9
	iface.RegisterClazzName(ClazzName_authsession_getLangPack, 0, 0x29bbc166)             // 29bbc166
	iface.RegisterClazzName(ClazzName_authsession_getClient, 0, 0x605855be)               // 605855be
	iface.RegisterClazzName(ClazzName_authsession_getLangCode, 0, 0x5899b559)             // 5899b559
	iface.RegisterClazzName(ClazzName_authsession_getUserId, 0, 0x57491cac)               // 57491cac
	iface.RegisterClazzName(ClazzName_authsession_getPushSessionId, 0, 0xb3c23141)        // b3c23141
	iface.RegisterClazzName(ClazzName_authsession_getFutureSalts, 0, 0xb8cf5815)          // b8cf5815
	iface.RegisterClazzName(ClazzName_authsession_queryAuthKey, 0, 0x54b73828)            // 54b73828
	iface.RegisterClazzName(ClazzName_authsession_setAuthKey, 0, 0x3e940c91)              // 3e940c91
	iface.RegisterClazzName(ClazzName_authsession_bindAuthKeyUser, 0, 0xbce0423)          // bce0423
	iface.RegisterClazzName(ClazzName_authsession_unbindAuthKeyUser, 0, 0x758c648)        // 758c648
	iface.RegisterClazzName(ClazzName_authsession_getPermAuthKeyId, 0, 0x907464d6)        // 907464d6
	iface.RegisterClazzName(ClazzName_authsession_bindTempAuthKey, 0, 0x608f4f86)         // 608f4f86
	iface.RegisterClazzName(ClazzName_authsession_setClientSessionInfo, 0, 0x2d9ff94)     // 2d9ff94
	iface.RegisterClazzName(ClazzName_authsession_getAuthorization, 0, 0x6e5e1923)        // 6e5e1923
	iface.RegisterClazzName(ClazzName_authsession_getAuthStateData, 0, 0x4f5e3131)        // 4f5e3131
	iface.RegisterClazzName(ClazzName_authsession_setLayer, 0, 0x44651485)                // 44651485
	iface.RegisterClazzName(ClazzName_authsession_setInitConnection, 0, 0x7cdf8a8c)       // 7cdf8a8c
	iface.RegisterClazzName(ClazzName_authsession_setAndroidPushSessionId, 0, 0x92a8233c) // 92a8233c

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_clientSession, 0x9a8e71b0)                       // 9a8e71b0
	iface.RegisterClazzIDName(ClazzName_authKeyStateData, 0xe0408f17)                    // e0408f17
	iface.RegisterClazzIDName(ClazzName_authsession_getAuthorizations, 0x30e21244)       // 30e21244
	iface.RegisterClazzIDName(ClazzName_authsession_resetAuthorization, 0x8d5f6ca6)      // 8d5f6ca6
	iface.RegisterClazzIDName(ClazzName_authsession_getLayer, 0xa82f16a9)                // a82f16a9
	iface.RegisterClazzIDName(ClazzName_authsession_getLangPack, 0x29bbc166)             // 29bbc166
	iface.RegisterClazzIDName(ClazzName_authsession_getClient, 0x605855be)               // 605855be
	iface.RegisterClazzIDName(ClazzName_authsession_getLangCode, 0x5899b559)             // 5899b559
	iface.RegisterClazzIDName(ClazzName_authsession_getUserId, 0x57491cac)               // 57491cac
	iface.RegisterClazzIDName(ClazzName_authsession_getPushSessionId, 0xb3c23141)        // b3c23141
	iface.RegisterClazzIDName(ClazzName_authsession_getFutureSalts, 0xb8cf5815)          // b8cf5815
	iface.RegisterClazzIDName(ClazzName_authsession_queryAuthKey, 0x54b73828)            // 54b73828
	iface.RegisterClazzIDName(ClazzName_authsession_setAuthKey, 0x3e940c91)              // 3e940c91
	iface.RegisterClazzIDName(ClazzName_authsession_bindAuthKeyUser, 0xbce0423)          // bce0423
	iface.RegisterClazzIDName(ClazzName_authsession_unbindAuthKeyUser, 0x758c648)        // 758c648
	iface.RegisterClazzIDName(ClazzName_authsession_getPermAuthKeyId, 0x907464d6)        // 907464d6
	iface.RegisterClazzIDName(ClazzName_authsession_bindTempAuthKey, 0x608f4f86)         // 608f4f86
	iface.RegisterClazzIDName(ClazzName_authsession_setClientSessionInfo, 0x2d9ff94)     // 2d9ff94
	iface.RegisterClazzIDName(ClazzName_authsession_getAuthorization, 0x6e5e1923)        // 6e5e1923
	iface.RegisterClazzIDName(ClazzName_authsession_getAuthStateData, 0x4f5e3131)        // 4f5e3131
	iface.RegisterClazzIDName(ClazzName_authsession_setLayer, 0x44651485)                // 44651485
	iface.RegisterClazzIDName(ClazzName_authsession_setInitConnection, 0x7cdf8a8c)       // 7cdf8a8c
	iface.RegisterClazzIDName(ClazzName_authsession_setAndroidPushSessionId, 0x92a8233c) // 92a8233c
}
