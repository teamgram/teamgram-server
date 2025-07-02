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

package authsession

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0xe0408f17, func() iface.TLObject { return &TLAuthKeyStateData{ClazzID: 0xe0408f17} }) // 0xe0408f17
	iface.RegisterClazzID(0x9a8e71b0, func() iface.TLObject { return &TLClientSession{ClazzID: 0x9a8e71b0} })    // 0x9a8e71b0

	// Method
	iface.RegisterClazzID(0x30e21244, func() iface.TLObject { return &TLAuthsessionGetAuthorizations{ClazzID: 0x30e21244} })       // 0x30e21244
	iface.RegisterClazzID(0x8d5f6ca6, func() iface.TLObject { return &TLAuthsessionResetAuthorization{ClazzID: 0x8d5f6ca6} })      // 0x8d5f6ca6
	iface.RegisterClazzID(0xa82f16a9, func() iface.TLObject { return &TLAuthsessionGetLayer{ClazzID: 0xa82f16a9} })                // 0xa82f16a9
	iface.RegisterClazzID(0x29bbc166, func() iface.TLObject { return &TLAuthsessionGetLangPack{ClazzID: 0x29bbc166} })             // 0x29bbc166
	iface.RegisterClazzID(0x605855be, func() iface.TLObject { return &TLAuthsessionGetClient{ClazzID: 0x605855be} })               // 0x605855be
	iface.RegisterClazzID(0x5899b559, func() iface.TLObject { return &TLAuthsessionGetLangCode{ClazzID: 0x5899b559} })             // 0x5899b559
	iface.RegisterClazzID(0x57491cac, func() iface.TLObject { return &TLAuthsessionGetUserId{ClazzID: 0x57491cac} })               // 0x57491cac
	iface.RegisterClazzID(0xb3c23141, func() iface.TLObject { return &TLAuthsessionGetPushSessionId{ClazzID: 0xb3c23141} })        // 0xb3c23141
	iface.RegisterClazzID(0xb8cf5815, func() iface.TLObject { return &TLAuthsessionGetFutureSalts{ClazzID: 0xb8cf5815} })          // 0xb8cf5815
	iface.RegisterClazzID(0x54b73828, func() iface.TLObject { return &TLAuthsessionQueryAuthKey{ClazzID: 0x54b73828} })            // 0x54b73828
	iface.RegisterClazzID(0x3e940c91, func() iface.TLObject { return &TLAuthsessionSetAuthKey{ClazzID: 0x3e940c91} })              // 0x3e940c91
	iface.RegisterClazzID(0xbce0423, func() iface.TLObject { return &TLAuthsessionBindAuthKeyUser{ClazzID: 0xbce0423} })           // 0xbce0423
	iface.RegisterClazzID(0x758c648, func() iface.TLObject { return &TLAuthsessionUnbindAuthKeyUser{ClazzID: 0x758c648} })         // 0x758c648
	iface.RegisterClazzID(0x907464d6, func() iface.TLObject { return &TLAuthsessionGetPermAuthKeyId{ClazzID: 0x907464d6} })        // 0x907464d6
	iface.RegisterClazzID(0x608f4f86, func() iface.TLObject { return &TLAuthsessionBindTempAuthKey{ClazzID: 0x608f4f86} })         // 0x608f4f86
	iface.RegisterClazzID(0x2d9ff94, func() iface.TLObject { return &TLAuthsessionSetClientSessionInfo{ClazzID: 0x2d9ff94} })      // 0x2d9ff94
	iface.RegisterClazzID(0x6e5e1923, func() iface.TLObject { return &TLAuthsessionGetAuthorization{ClazzID: 0x6e5e1923} })        // 0x6e5e1923
	iface.RegisterClazzID(0x4f5e3131, func() iface.TLObject { return &TLAuthsessionGetAuthStateData{ClazzID: 0x4f5e3131} })        // 0x4f5e3131
	iface.RegisterClazzID(0x44651485, func() iface.TLObject { return &TLAuthsessionSetLayer{ClazzID: 0x44651485} })                // 0x44651485
	iface.RegisterClazzID(0x7cdf8a8c, func() iface.TLObject { return &TLAuthsessionSetInitConnection{ClazzID: 0x7cdf8a8c} })       // 0x7cdf8a8c
	iface.RegisterClazzID(0x92a8233c, func() iface.TLObject { return &TLAuthsessionSetAndroidPushSessionId{ClazzID: 0x92a8233c} }) // 0x92a8233c
}
