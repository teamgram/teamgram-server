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

const (
	Predicate_clientSession                    = "clientSession"
	Predicate_authKeyStateData                 = "authKeyStateData"
	Predicate_authsession_getAuthorizations    = "authsession_getAuthorizations"
	Predicate_authsession_resetAuthorization   = "authsession_resetAuthorization"
	Predicate_authsession_getLayer             = "authsession_getLayer"
	Predicate_authsession_getLangPack          = "authsession_getLangPack"
	Predicate_authsession_getClient            = "authsession_getClient"
	Predicate_authsession_getLangCode          = "authsession_getLangCode"
	Predicate_authsession_getUserId            = "authsession_getUserId"
	Predicate_authsession_getPushSessionId     = "authsession_getPushSessionId"
	Predicate_authsession_getFutureSalts       = "authsession_getFutureSalts"
	Predicate_authsession_queryAuthKey         = "authsession_queryAuthKey"
	Predicate_authsession_setAuthKey           = "authsession_setAuthKey"
	Predicate_authsession_bindAuthKeyUser      = "authsession_bindAuthKeyUser"
	Predicate_authsession_unbindAuthKeyUser    = "authsession_unbindAuthKeyUser"
	Predicate_authsession_getPermAuthKeyId     = "authsession_getPermAuthKeyId"
	Predicate_authsession_bindTempAuthKey      = "authsession_bindTempAuthKey"
	Predicate_authsession_setClientSessionInfo = "authsession_setClientSessionInfo"
	Predicate_authsession_getAuthorization     = "authsession_getAuthorization"
	Predicate_authsession_getAuthStateData     = "authsession_getAuthStateData"
	Predicate_authsession_setLayer             = "authsession_setLayer"
	Predicate_authsession_setInitConnection    = "authsession_setInitConnection"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_clientSession: {
		0: -1701940816, // 0x9a8e71b0

	},
	Predicate_authKeyStateData: {
		0: -532639977, // 0xe0408f17

	},
	Predicate_authsession_getAuthorizations: {
		0: 820122180, // 0x30e21244

	},
	Predicate_authsession_resetAuthorization: {
		0: -1923126106, // 0x8d5f6ca6

	},
	Predicate_authsession_getLayer: {
		0: -1473309015, // 0xa82f16a9

	},
	Predicate_authsession_getLangPack: {
		0: 700170598, // 0x29bbc166

	},
	Predicate_authsession_getClient: {
		0: 1616401854, // 0x605855be

	},
	Predicate_authsession_getLangCode: {
		0: 1486468441, // 0x5899b559

	},
	Predicate_authsession_getUserId: {
		0: 1464409260, // 0x57491cac

	},
	Predicate_authsession_getPushSessionId: {
		0: -1279119039, // 0xb3c23141

	},
	Predicate_authsession_getFutureSalts: {
		0: -1194371051, // 0xb8cf5815

	},
	Predicate_authsession_queryAuthKey: {
		0: 1421293608, // 0x54b73828

	},
	Predicate_authsession_setAuthKey: {
		0: 1049889937, // 0x3e940c91

	},
	Predicate_authsession_bindAuthKeyUser: {
		0: 198050851, // 0xbce0423

	},
	Predicate_authsession_unbindAuthKeyUser: {
		0: 123258440, // 0x758c648

	},
	Predicate_authsession_getPermAuthKeyId: {
		0: -1871420202, // 0x907464d6

	},
	Predicate_authsession_bindTempAuthKey: {
		0: 1620004742, // 0x608f4f86

	},
	Predicate_authsession_setClientSessionInfo: {
		0: 47841172, // 0x2d9ff94

	},
	Predicate_authsession_getAuthorization: {
		0: 1851660579, // 0x6e5e1923

	},
	Predicate_authsession_getAuthStateData: {
		0: 1331573041, // 0x4f5e3131

	},
	Predicate_authsession_setLayer: {
		0: 1147475077, // 0x44651485

	},
	Predicate_authsession_setInitConnection: {
		0: 2095024780, // 0x7cdf8a8c

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	-1701940816: Predicate_clientSession,                    // 0x9a8e71b0
	-532639977:  Predicate_authKeyStateData,                 // 0xe0408f17
	820122180:   Predicate_authsession_getAuthorizations,    // 0x30e21244
	-1923126106: Predicate_authsession_resetAuthorization,   // 0x8d5f6ca6
	-1473309015: Predicate_authsession_getLayer,             // 0xa82f16a9
	700170598:   Predicate_authsession_getLangPack,          // 0x29bbc166
	1616401854:  Predicate_authsession_getClient,            // 0x605855be
	1486468441:  Predicate_authsession_getLangCode,          // 0x5899b559
	1464409260:  Predicate_authsession_getUserId,            // 0x57491cac
	-1279119039: Predicate_authsession_getPushSessionId,     // 0xb3c23141
	-1194371051: Predicate_authsession_getFutureSalts,       // 0xb8cf5815
	1421293608:  Predicate_authsession_queryAuthKey,         // 0x54b73828
	1049889937:  Predicate_authsession_setAuthKey,           // 0x3e940c91
	198050851:   Predicate_authsession_bindAuthKeyUser,      // 0xbce0423
	123258440:   Predicate_authsession_unbindAuthKeyUser,    // 0x758c648
	-1871420202: Predicate_authsession_getPermAuthKeyId,     // 0x907464d6
	1620004742:  Predicate_authsession_bindTempAuthKey,      // 0x608f4f86
	47841172:    Predicate_authsession_setClientSessionInfo, // 0x2d9ff94
	1851660579:  Predicate_authsession_getAuthorization,     // 0x6e5e1923
	1331573041:  Predicate_authsession_getAuthStateData,     // 0x4f5e3131
	1147475077:  Predicate_authsession_setLayer,             // 0x44651485
	2095024780:  Predicate_authsession_setInitConnection,    // 0x7cdf8a8c

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
