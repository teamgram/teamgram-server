/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

// ConstructorList
// RequestList

package userupdates

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0x96790279, func() iface.TLObject { return &TLAffectedUserOperation{ClazzID: 0x96790279} })   // 0x96790279
	iface.RegisterClazzID(0xdf6c5662, func() iface.TLObject { return &TLDialogProjection{ClazzID: 0xdf6c5662} })        // 0xdf6c5662
	iface.RegisterClazzID(0x4e60f01f, func() iface.TLObject { return &TLDialogProjectionList{ClazzID: 0x4e60f01f} })    // 0x4e60f01f
	iface.RegisterClazzID(0x6d7ec124, func() iface.TLObject { return &TLDialogProjectionPeer{ClazzID: 0x6d7ec124} })    // 0x6d7ec124
	iface.RegisterClazzID(0x3127345e, func() iface.TLObject { return &TLMessageViewList{ClazzID: 0x3127345e} })         // 0x3127345e
	iface.RegisterClazzID(0x8bf3b9a4, func() iface.TLObject { return &TLMessageViewPeerSeq{ClazzID: 0x8bf3b9a4} })      // 0x8bf3b9a4
	iface.RegisterClazzID(0x55994646, func() iface.TLObject { return &TLUserAuthSeqAppendResult{ClazzID: 0x55994646} }) // 0x55994646
	iface.RegisterClazzID(0xb38ac177, func() iface.TLObject { return &TLUserDifferenceEmpty{ClazzID: 0xb38ac177} })     // 0xb38ac177
	iface.RegisterClazzID(0xb15cb08d, func() iface.TLObject { return &TLUserDifference{ClazzID: 0xb15cb08d} })          // 0xb15cb08d
	iface.RegisterClazzID(0x4ef1987f, func() iface.TLObject { return &TLUserDifferenceSlice{ClazzID: 0x4ef1987f} })     // 0x4ef1987f
	iface.RegisterClazzID(0x1d095703, func() iface.TLObject { return &TLUserDifferenceTooLong{ClazzID: 0x1d095703} })   // 0x1d095703
	iface.RegisterClazzID(0x2d4e84d7, func() iface.TLObject { return &TLUserOperation{ClazzID: 0x2d4e84d7} })           // 0x2d4e84d7
	iface.RegisterClazzID(0x7311db72, func() iface.TLObject { return &TLUserOperationResult{ClazzID: 0x7311db72} })     // 0x7311db72
	iface.RegisterClazzID(0xaa3fff4f, func() iface.TLObject { return &TLUserPtsAppendResult{ClazzID: 0xaa3fff4f} })     // 0xaa3fff4f
	iface.RegisterClazzID(0x635f3815, func() iface.TLObject { return &TLUserState{ClazzID: 0x635f3815} })               // 0x635f3815

	// Method
	iface.RegisterClazzID(0xc200ea59, func() iface.TLObject { return &TLUserupdatesProcessUserOperation{ClazzID: 0xc200ea59} })            // 0xc200ea59
	iface.RegisterClazzID(0xbacea5bf, func() iface.TLObject { return &TLUserupdatesProcessUserOperationWithEffects{ClazzID: 0xbacea5bf} }) // 0xbacea5bf
	iface.RegisterClazzID(0x47a995d1, func() iface.TLObject { return &TLUserupdatesGetOperationResult{ClazzID: 0x47a995d1} })              // 0x47a995d1
	iface.RegisterClazzID(0x3bbbad80, func() iface.TLObject { return &TLUserupdatesGetState{ClazzID: 0x3bbbad80} })                        // 0x3bbbad80
	iface.RegisterClazzID(0x38cdd9fc, func() iface.TLObject { return &TLUserupdatesGetDifference{ClazzID: 0x38cdd9fc} })                   // 0x38cdd9fc
	iface.RegisterClazzID(0x53638fcc, func() iface.TLObject { return &TLUserupdatesListDialogs{ClazzID: 0x53638fcc} })                     // 0x53638fcc
	iface.RegisterClazzID(0xc6a9626f, func() iface.TLObject { return &TLUserupdatesGetDialogsByPeers{ClazzID: 0xc6a9626f} })               // 0xc6a9626f
	iface.RegisterClazzID(0x12060b16, func() iface.TLObject { return &TLUserupdatesGetDialogCount{ClazzID: 0x12060b16} })                  // 0x12060b16
	iface.RegisterClazzID(0x528a3e52, func() iface.TLObject { return &TLUserupdatesGetMessageViewsByPeerSeqs{ClazzID: 0x528a3e52} })       // 0x528a3e52
	iface.RegisterClazzID(0x56fb4ad9, func() iface.TLObject { return &TLUserupdatesGetOutboxReadDate{ClazzID: 0x56fb4ad9} })               // 0x56fb4ad9
	iface.RegisterClazzID(0x170844e5, func() iface.TLObject { return &TLUserupdatesAppendDialogAuthSeqSideEffect{ClazzID: 0x170844e5} })   // 0x170844e5
	iface.RegisterClazzID(0xe93427fd, func() iface.TLObject { return &TLUserupdatesAppendDialogPtsSideEffect{ClazzID: 0xe93427fd} })       // 0xe93427fd
}
