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

package session

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0xdbd8534f, func() iface.TLObject { return &TLHttpSessionData{ClazzID: 0xdbd8534f} })    // 0xdbd8534f
	iface.RegisterClazzID(0x41a20c4e, func() iface.TLObject { return &TLSessionClientData{ClazzID: 0x41a20c4e} })  // 0x41a20c4e
	iface.RegisterClazzID(0xf17f375f, func() iface.TLObject { return &TLSessionClientEvent{ClazzID: 0xf17f375f} }) // 0xf17f375f

	// Method
	iface.RegisterClazzID(0x6b2df851, func() iface.TLObject { return &TLSessionQueryAuthKey{ClazzID: 0x6b2df851} })           // 0x6b2df851
	iface.RegisterClazzID(0x1d11490b, func() iface.TLObject { return &TLSessionSetAuthKey{ClazzID: 0x1d11490b} })             // 0x1d11490b
	iface.RegisterClazzID(0x410cb20d, func() iface.TLObject { return &TLSessionCreateSession{ClazzID: 0x410cb20d} })          // 0x410cb20d
	iface.RegisterClazzID(0x876b2dec, func() iface.TLObject { return &TLSessionSendDataToSession{ClazzID: 0x876b2dec} })      // 0x876b2dec
	iface.RegisterClazzID(0xbbec23ae, func() iface.TLObject { return &TLSessionSendHttpDataToSession{ClazzID: 0xbbec23ae} })  // 0xbbec23ae
	iface.RegisterClazzID(0x176fc253, func() iface.TLObject { return &TLSessionCloseSession{ClazzID: 0x176fc253} })           // 0x176fc253
	iface.RegisterClazzID(0xa574d829, func() iface.TLObject { return &TLSessionPushUpdatesData{ClazzID: 0xa574d829} })        // 0xa574d829
	iface.RegisterClazzID(0x45f3fda0, func() iface.TLObject { return &TLSessionPushSessionUpdatesData{ClazzID: 0x45f3fda0} }) // 0x45f3fda0
	iface.RegisterClazzID(0x4b470c89, func() iface.TLObject { return &TLSessionPushRpcResultData{ClazzID: 0x4b470c89} })      // 0x4b470c89
}
