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

const (
	Predicate_sessionClientEvent             = "sessionClientEvent"
	Predicate_sessionClientData              = "sessionClientData"
	Predicate_httpSessionData                = "httpSessionData"
	Predicate_session_queryAuthKey           = "session_queryAuthKey"
	Predicate_session_setAuthKey             = "session_setAuthKey"
	Predicate_session_createSession          = "session_createSession"
	Predicate_session_sendDataToSession      = "session_sendDataToSession"
	Predicate_session_sendHttpDataToSession  = "session_sendHttpDataToSession"
	Predicate_session_closeSession           = "session_closeSession"
	Predicate_session_pushUpdatesData        = "session_pushUpdatesData"
	Predicate_session_pushSessionUpdatesData = "session_pushSessionUpdatesData"
	Predicate_session_pushRpcResultData      = "session_pushRpcResultData"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_sessionClientEvent: {
		0: -243320993, // 0xf17f375f

	},
	Predicate_sessionClientData: {
		0: 1101139022, // 0x41a20c4e

	},
	Predicate_httpSessionData: {
		0: -606579889, // 0xdbd8534f

	},
	Predicate_session_queryAuthKey: {
		0: 1798174801, // 0x6b2df851

	},
	Predicate_session_setAuthKey: {
		0: 487672075, // 0x1d11490b

	},
	Predicate_session_createSession: {
		0: 1091351053, // 0x410cb20d

	},
	Predicate_session_sendDataToSession: {
		0: -2023019028, // 0x876b2dec

	},
	Predicate_session_sendHttpDataToSession: {
		0: -1142152274, // 0xbbec23ae

	},
	Predicate_session_closeSession: {
		0: 393200211, // 0x176fc253

	},
	Predicate_session_pushUpdatesData: {
		0: -1519069143, // 0xa574d829

	},
	Predicate_session_pushSessionUpdatesData: {
		0: 1173618080, // 0x45f3fda0

	},
	Predicate_session_pushRpcResultData: {
		0: 1262947465, // 0x4b470c89

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	-243320993:  Predicate_sessionClientEvent,             // 0xf17f375f
	1101139022:  Predicate_sessionClientData,              // 0x41a20c4e
	-606579889:  Predicate_httpSessionData,                // 0xdbd8534f
	1798174801:  Predicate_session_queryAuthKey,           // 0x6b2df851
	487672075:   Predicate_session_setAuthKey,             // 0x1d11490b
	1091351053:  Predicate_session_createSession,          // 0x410cb20d
	-2023019028: Predicate_session_sendDataToSession,      // 0x876b2dec
	-1142152274: Predicate_session_sendHttpDataToSession,  // 0xbbec23ae
	393200211:   Predicate_session_closeSession,           // 0x176fc253
	-1519069143: Predicate_session_pushUpdatesData,        // 0xa574d829
	1173618080:  Predicate_session_pushSessionUpdatesData, // 0x45f3fda0
	1262947465:  Predicate_session_pushRpcResultData,      // 0x4b470c89

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
