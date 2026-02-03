/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgooo Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

// ConstructorList
// RequestList

package mt

import (
	"github.com/teamgooo/teamgooo-server/pkg/proto/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0x75a3f765, func() iface.TLObject { return &TLBindAuthKeyInner{ClazzID: 0x75a3f765} })        // 0x75a3f765
	iface.RegisterClazzID(0x6643b654, func() iface.TLObject { return &TLClientDHInnerData{ClazzID: 0x6643b654} })       // 0x6643b654
	iface.RegisterClazzID(0xf660e1d4, func() iface.TLObject { return &TLDestroyAuthKeyOk{ClazzID: 0xf660e1d4} })        // 0xf660e1d4
	iface.RegisterClazzID(0xa9f2259, func() iface.TLObject { return &TLDestroyAuthKeyNone{ClazzID: 0xa9f2259} })        // 0xa9f2259
	iface.RegisterClazzID(0xea109b13, func() iface.TLObject { return &TLDestroyAuthKeyFail{ClazzID: 0xea109b13} })      // 0xea109b13
	iface.RegisterClazzID(0x83c95aec, func() iface.TLObject { return &TLPQInnerData{ClazzID: 0x83c95aec} })             // 0x83c95aec
	iface.RegisterClazzID(0xa9f55f95, func() iface.TLObject { return &TLPQInnerDataDc{ClazzID: 0xa9f55f95} })           // 0xa9f55f95
	iface.RegisterClazzID(0x3c6a84d4, func() iface.TLObject { return &TLPQInnerDataTemp{ClazzID: 0x3c6a84d4} })         // 0x3c6a84d4
	iface.RegisterClazzID(0x56fddf88, func() iface.TLObject { return &TLPQInnerDataTempDc{ClazzID: 0x56fddf88} })       // 0x56fddf88
	iface.RegisterClazzID(0x5162463, func() iface.TLObject { return &TLResPQ{ClazzID: 0x5162463} })                     // 0x5162463
	iface.RegisterClazzID(0xb5890dba, func() iface.TLObject { return &TLServerDHInnerData{ClazzID: 0xb5890dba} })       // 0xb5890dba
	iface.RegisterClazzID(0x79cb045d, func() iface.TLObject { return &TLServerDHParamsFail{ClazzID: 0x79cb045d} })      // 0x79cb045d
	iface.RegisterClazzID(0xd0e8075c, func() iface.TLObject { return &TLServerDHParamsOk{ClazzID: 0xd0e8075c} })        // 0xd0e8075c
	iface.RegisterClazzID(0x3bcbf734, func() iface.TLObject { return &TLDhGenOk{ClazzID: 0x3bcbf734} })                 // 0x3bcbf734
	iface.RegisterClazzID(0x46dc1fb9, func() iface.TLObject { return &TLDhGenRetry{ClazzID: 0x46dc1fb9} })              // 0x46dc1fb9
	iface.RegisterClazzID(0xa69dae02, func() iface.TLObject { return &TLDhGenFail{ClazzID: 0xa69dae02} })               // 0xa69dae02
	iface.RegisterClazzID(0x4679b65f, func() iface.TLObject { return &TLAccessPointRule{ClazzID: 0x4679b65f} })         // 0x4679b65f
	iface.RegisterClazzID(0xa7eff811, func() iface.TLObject { return &TLBadMsgNotification{ClazzID: 0xa7eff811} })      // 0xa7eff811
	iface.RegisterClazzID(0xedab447b, func() iface.TLObject { return &TLBadServerSalt{ClazzID: 0xedab447b} })           // 0xedab447b
	iface.RegisterClazzID(0xe22045fc, func() iface.TLObject { return &TLDestroySessionOk{ClazzID: 0xe22045fc} })        // 0xe22045fc
	iface.RegisterClazzID(0x62d350c9, func() iface.TLObject { return &TLDestroySessionNone{ClazzID: 0x62d350c9} })      // 0x62d350c9
	iface.RegisterClazzID(0x949d9dc, func() iface.TLObject { return &TLFutureSalt{ClazzID: 0x949d9dc} })                // 0x949d9dc
	iface.RegisterClazzID(0xae500895, func() iface.TLObject { return &TLFutureSalts{ClazzID: 0xae500895} })             // 0xae500895
	iface.RegisterClazzID(0x5a592a6c, func() iface.TLObject { return &TLHelpConfigSimple{ClazzID: 0x5a592a6c} })        // 0x5a592a6c
	iface.RegisterClazzID(0x9299359f, func() iface.TLObject { return &TLHttpWait{ClazzID: 0x9299359f} })                // 0x9299359f
	iface.RegisterClazzID(0xd433ad73, func() iface.TLObject { return &TLIpPort{ClazzID: 0xd433ad73} })                  // 0xd433ad73
	iface.RegisterClazzID(0x37982646, func() iface.TLObject { return &TLIpPortSecret{ClazzID: 0x37982646} })            // 0x37982646
	iface.RegisterClazzID(0x276d3ec6, func() iface.TLObject { return &TLMsgDetailedInfo{ClazzID: 0x276d3ec6} })         // 0x276d3ec6
	iface.RegisterClazzID(0x809db6df, func() iface.TLObject { return &TLMsgNewDetailedInfo{ClazzID: 0x809db6df} })      // 0x809db6df
	iface.RegisterClazzID(0x7d861a08, func() iface.TLObject { return &TLMsgResendReq{ClazzID: 0x7d861a08} })            // 0x7d861a08
	iface.RegisterClazzID(0x62d6b459, func() iface.TLObject { return &TLMsgsAck{ClazzID: 0x62d6b459} })                 // 0x62d6b459
	iface.RegisterClazzID(0x8cc0d131, func() iface.TLObject { return &TLMsgsAllInfo{ClazzID: 0x8cc0d131} })             // 0x8cc0d131
	iface.RegisterClazzID(0x4deb57d, func() iface.TLObject { return &TLMsgsStateInfo{ClazzID: 0x4deb57d} })             // 0x4deb57d
	iface.RegisterClazzID(0xda69fb52, func() iface.TLObject { return &TLMsgsStateReq{ClazzID: 0xda69fb52} })            // 0xda69fb52
	iface.RegisterClazzID(0x9ec20908, func() iface.TLObject { return &TLNewSessionCreated{ClazzID: 0x9ec20908} })       // 0x9ec20908
	iface.RegisterClazzID(0x347773c5, func() iface.TLObject { return &TLPong{ClazzID: 0x347773c5} })                    // 0x347773c5
	iface.RegisterClazzID(0x5e2ad36e, func() iface.TLObject { return &TLRpcAnswerUnknown{ClazzID: 0x5e2ad36e} })        // 0x5e2ad36e
	iface.RegisterClazzID(0xcd78e586, func() iface.TLObject { return &TLRpcAnswerDroppedRunning{ClazzID: 0xcd78e586} }) // 0xcd78e586
	iface.RegisterClazzID(0xa43ad8b7, func() iface.TLObject { return &TLRpcAnswerDropped{ClazzID: 0xa43ad8b7} })        // 0xa43ad8b7
	iface.RegisterClazzID(0x2144ca19, func() iface.TLObject { return &TLRpcError{ClazzID: 0x2144ca19} })                // 0x2144ca19
	iface.RegisterClazzID(0x4218a164, func() iface.TLObject { return &TLTlsBlockString{ClazzID: 0x4218a164} })          // 0x4218a164
	iface.RegisterClazzID(0x4d4dc41e, func() iface.TLObject { return &TLTlsBlockRandom{ClazzID: 0x4d4dc41e} })          // 0x4d4dc41e
	iface.RegisterClazzID(0x9333afb, func() iface.TLObject { return &TLTlsBlockZero{ClazzID: 0x9333afb} })              // 0x9333afb
	iface.RegisterClazzID(0x10e8636f, func() iface.TLObject { return &TLTlsBlockDomain{ClazzID: 0x10e8636f} })          // 0x10e8636f
	iface.RegisterClazzID(0xe675a1c1, func() iface.TLObject { return &TLTlsBlockGrease{ClazzID: 0xe675a1c1} })          // 0xe675a1c1
	iface.RegisterClazzID(0x9eb95b5c, func() iface.TLObject { return &TLTlsBlockPublicKey{ClazzID: 0x9eb95b5c} })       // 0x9eb95b5c
	iface.RegisterClazzID(0xe725d44f, func() iface.TLObject { return &TLTlsBlockScope{ClazzID: 0xe725d44f} })           // 0xe725d44f
	iface.RegisterClazzID(0x6c52c484, func() iface.TLObject { return &TLTlsClientHello{ClazzID: 0x6c52c484} })          // 0x6c52c484

	// Method
	iface.RegisterClazzID(0x60469778, func() iface.TLObject { return &TLReqPq{ClazzID: 0x60469778} })               // 0x60469778
	iface.RegisterClazzID(0xbe7e8ef1, func() iface.TLObject { return &TLReqPqMulti{ClazzID: 0xbe7e8ef1} })          // 0xbe7e8ef1
	iface.RegisterClazzID(0xd712e4be, func() iface.TLObject { return &TLReqDHParams{ClazzID: 0xd712e4be} })         // 0xd712e4be
	iface.RegisterClazzID(0xf5045f1f, func() iface.TLObject { return &TLSetClientDHParams{ClazzID: 0xf5045f1f} })   // 0xf5045f1f
	iface.RegisterClazzID(0xd1435160, func() iface.TLObject { return &TLDestroyAuthKey{ClazzID: 0xd1435160} })      // 0xd1435160
	iface.RegisterClazzID(0xc0e202f7, func() iface.TLObject { return &TLHelpTest{ClazzID: 0xc0e202f7} })            // 0xc0e202f7
	iface.RegisterClazzID(0xee75af01, func() iface.TLObject { return &TLTestUseError{ClazzID: 0xee75af01} })        // 0xee75af01
	iface.RegisterClazzID(0xf9b7b23d, func() iface.TLObject { return &TLTestUseConfigSimple{ClazzID: 0xf9b7b23d} }) // 0xf9b7b23d
	iface.RegisterClazzID(0x58e4a740, func() iface.TLObject { return &TLRpcDropAnswer{ClazzID: 0x58e4a740} })       // 0x58e4a740
	iface.RegisterClazzID(0xb921bd04, func() iface.TLObject { return &TLGetFutureSalts{ClazzID: 0xb921bd04} })      // 0xb921bd04
	iface.RegisterClazzID(0x7abe77ec, func() iface.TLObject { return &TLPing{ClazzID: 0x7abe77ec} })                // 0x7abe77ec
	iface.RegisterClazzID(0xf3427b8c, func() iface.TLObject { return &TLPingDelayDisconnect{ClazzID: 0xf3427b8c} }) // 0xf3427b8c
	iface.RegisterClazzID(0xe7512126, func() iface.TLObject { return &TLDestroySession{ClazzID: 0xe7512126} })      // 0xe7512126
}
