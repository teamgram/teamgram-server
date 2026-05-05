/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package userupdates

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

const (
	ClazzName_userOperation                             = "userOperation"
	ClazzName_userOperationResult                       = "userOperationResult"
	ClazzName_userState                                 = "userState"
	ClazzName_userDifferenceEmpty                       = "userDifferenceEmpty"
	ClazzName_userDifference                            = "userDifference"
	ClazzName_userDifferenceSlice                       = "userDifferenceSlice"
	ClazzName_userDifferenceTooLong                     = "userDifferenceTooLong"
	ClazzName_dialogProjectionPeer                      = "dialogProjectionPeer"
	ClazzName_dialogProjection                          = "dialogProjection"
	ClazzName_dialogProjectionList                      = "dialogProjectionList"
	ClazzName_messageViewPeerSeq                        = "messageViewPeerSeq"
	ClazzName_messageViewList                           = "messageViewList"
	ClazzName_userAuthSeqAppendResult                   = "userAuthSeqAppendResult"
	ClazzName_userPtsAppendResult                       = "userPtsAppendResult"
	ClazzName_userupdates_processUserOperation          = "userupdates_processUserOperation"
	ClazzName_userupdates_getOperationResult            = "userupdates_getOperationResult"
	ClazzName_userupdates_getState                      = "userupdates_getState"
	ClazzName_userupdates_getDifference                 = "userupdates_getDifference"
	ClazzName_userupdates_listDialogs                   = "userupdates_listDialogs"
	ClazzName_userupdates_getDialogsByPeers             = "userupdates_getDialogsByPeers"
	ClazzName_userupdates_getDialogCount                = "userupdates_getDialogCount"
	ClazzName_userupdates_getMessageViewsByPeerSeqs     = "userupdates_getMessageViewsByPeerSeqs"
	ClazzName_userupdates_appendDialogAuthSeqSideEffect = "userupdates_appendDialogAuthSeqSideEffect"
	ClazzName_userupdates_appendDialogPtsSideEffect     = "userupdates_appendDialogPtsSideEffect"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_userOperation, 0, 0x2d4e84d7)                             // 2d4e84d7
	iface.RegisterClazzName(ClazzName_userOperationResult, 0, 0x7311db72)                       // 7311db72
	iface.RegisterClazzName(ClazzName_userState, 0, 0x635f3815)                                 // 635f3815
	iface.RegisterClazzName(ClazzName_userDifferenceEmpty, 0, 0xb38ac177)                       // b38ac177
	iface.RegisterClazzName(ClazzName_userDifference, 0, 0xb15cb08d)                            // b15cb08d
	iface.RegisterClazzName(ClazzName_userDifferenceSlice, 0, 0x4ef1987f)                       // 4ef1987f
	iface.RegisterClazzName(ClazzName_userDifferenceTooLong, 0, 0x1d095703)                     // 1d095703
	iface.RegisterClazzName(ClazzName_dialogProjectionPeer, 0, 0x6d7ec124)                      // 6d7ec124
	iface.RegisterClazzName(ClazzName_dialogProjection, 0, 0xb9bc23fd)                          // b9bc23fd
	iface.RegisterClazzName(ClazzName_dialogProjectionList, 0, 0x4e60f01f)                      // 4e60f01f
	iface.RegisterClazzName(ClazzName_messageViewPeerSeq, 0, 0x8bf3b9a4)                        // 8bf3b9a4
	iface.RegisterClazzName(ClazzName_messageViewList, 0, 0x3127345e)                           // 3127345e
	iface.RegisterClazzName(ClazzName_userAuthSeqAppendResult, 0, 0x55994646)                   // 55994646
	iface.RegisterClazzName(ClazzName_userPtsAppendResult, 0, 0xaa3fff4f)                       // aa3fff4f
	iface.RegisterClazzName(ClazzName_userupdates_processUserOperation, 0, 0xc200ea59)          // c200ea59
	iface.RegisterClazzName(ClazzName_userupdates_getOperationResult, 0, 0x47a995d1)            // 47a995d1
	iface.RegisterClazzName(ClazzName_userupdates_getState, 0, 0x3bbbad80)                      // 3bbbad80
	iface.RegisterClazzName(ClazzName_userupdates_getDifference, 0, 0x38cdd9fc)                 // 38cdd9fc
	iface.RegisterClazzName(ClazzName_userupdates_listDialogs, 0, 0x53638fcc)                   // 53638fcc
	iface.RegisterClazzName(ClazzName_userupdates_getDialogsByPeers, 0, 0xc6a9626f)             // c6a9626f
	iface.RegisterClazzName(ClazzName_userupdates_getDialogCount, 0, 0x12060b16)                // 12060b16
	iface.RegisterClazzName(ClazzName_userupdates_getMessageViewsByPeerSeqs, 0, 0x528a3e52)     // 528a3e52
	iface.RegisterClazzName(ClazzName_userupdates_appendDialogAuthSeqSideEffect, 0, 0x170844e5) // 170844e5
	iface.RegisterClazzName(ClazzName_userupdates_appendDialogPtsSideEffect, 0, 0xe93427fd)     // e93427fd

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_userOperation, 0x2d4e84d7)                             // 2d4e84d7
	iface.RegisterClazzIDName(ClazzName_userOperationResult, 0x7311db72)                       // 7311db72
	iface.RegisterClazzIDName(ClazzName_userState, 0x635f3815)                                 // 635f3815
	iface.RegisterClazzIDName(ClazzName_userDifferenceEmpty, 0xb38ac177)                       // b38ac177
	iface.RegisterClazzIDName(ClazzName_userDifference, 0xb15cb08d)                            // b15cb08d
	iface.RegisterClazzIDName(ClazzName_userDifferenceSlice, 0x4ef1987f)                       // 4ef1987f
	iface.RegisterClazzIDName(ClazzName_userDifferenceTooLong, 0x1d095703)                     // 1d095703
	iface.RegisterClazzIDName(ClazzName_dialogProjectionPeer, 0x6d7ec124)                      // 6d7ec124
	iface.RegisterClazzIDName(ClazzName_dialogProjection, 0xb9bc23fd)                          // b9bc23fd
	iface.RegisterClazzIDName(ClazzName_dialogProjectionList, 0x4e60f01f)                      // 4e60f01f
	iface.RegisterClazzIDName(ClazzName_messageViewPeerSeq, 0x8bf3b9a4)                        // 8bf3b9a4
	iface.RegisterClazzIDName(ClazzName_messageViewList, 0x3127345e)                           // 3127345e
	iface.RegisterClazzIDName(ClazzName_userAuthSeqAppendResult, 0x55994646)                   // 55994646
	iface.RegisterClazzIDName(ClazzName_userPtsAppendResult, 0xaa3fff4f)                       // aa3fff4f
	iface.RegisterClazzIDName(ClazzName_userupdates_processUserOperation, 0xc200ea59)          // c200ea59
	iface.RegisterClazzIDName(ClazzName_userupdates_getOperationResult, 0x47a995d1)            // 47a995d1
	iface.RegisterClazzIDName(ClazzName_userupdates_getState, 0x3bbbad80)                      // 3bbbad80
	iface.RegisterClazzIDName(ClazzName_userupdates_getDifference, 0x38cdd9fc)                 // 38cdd9fc
	iface.RegisterClazzIDName(ClazzName_userupdates_listDialogs, 0x53638fcc)                   // 53638fcc
	iface.RegisterClazzIDName(ClazzName_userupdates_getDialogsByPeers, 0xc6a9626f)             // c6a9626f
	iface.RegisterClazzIDName(ClazzName_userupdates_getDialogCount, 0x12060b16)                // 12060b16
	iface.RegisterClazzIDName(ClazzName_userupdates_getMessageViewsByPeerSeqs, 0x528a3e52)     // 528a3e52
	iface.RegisterClazzIDName(ClazzName_userupdates_appendDialogAuthSeqSideEffect, 0x170844e5) // 170844e5
	iface.RegisterClazzIDName(ClazzName_userupdates_appendDialogPtsSideEffect, 0xe93427fd)     // e93427fd
}
