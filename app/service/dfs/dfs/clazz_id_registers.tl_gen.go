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

package dfs

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0xe83380f0, func() iface.TLObject { return &TLFileFinalizedObject{ClazzID: 0xe83380f0} }) // 0xe83380f0
	iface.RegisterClazzID(0x146aad14, func() iface.TLObject { return &TLFileHashChunk{ClazzID: 0x146aad14} })       // 0x146aad14

	// Method
	iface.RegisterClazzID(0xdddb9d2c, func() iface.TLObject { return &TLDfsCommitUpload{ClazzID: 0xdddb9d2c} })             // 0xdddb9d2c
	iface.RegisterClazzID(0x6e20c3e7, func() iface.TLObject { return &TLDfsPutFile{ClazzID: 0x6e20c3e7} })                  // 0x6e20c3e7
	iface.RegisterClazzID(0x86c7c115, func() iface.TLObject { return &TLDfsGetFileByReadLease{ClazzID: 0x86c7c115} })       // 0x86c7c115
	iface.RegisterClazzID(0xff974b78, func() iface.TLObject { return &TLDfsGetFileHashesByReadLease{ClazzID: 0xff974b78} }) // 0xff974b78
	iface.RegisterClazzID(0x1a484107, func() iface.TLObject { return &TLDfsWriteFilePartData{ClazzID: 0x1a484107} })        // 0x1a484107
	iface.RegisterClazzID(0x2410d1a2, func() iface.TLObject { return &TLDfsUploadPhotoFileV2{ClazzID: 0x2410d1a2} })        // 0x2410d1a2
	iface.RegisterClazzID(0x872313d8, func() iface.TLObject { return &TLDfsUploadProfilePhotoFileV2{ClazzID: 0x872313d8} }) // 0x872313d8
	iface.RegisterClazzID(0x79d3c523, func() iface.TLObject { return &TLDfsUploadEncryptedFileV2{ClazzID: 0x79d3c523} })    // 0x79d3c523
	iface.RegisterClazzID(0xd6bfee3e, func() iface.TLObject { return &TLDfsDownloadFile{ClazzID: 0xd6bfee3e} })             // 0xd6bfee3e
	iface.RegisterClazzID(0x76336db7, func() iface.TLObject { return &TLDfsUploadDocumentFileV2{ClazzID: 0x76336db7} })     // 0x76336db7
	iface.RegisterClazzID(0x41c4cd00, func() iface.TLObject { return &TLDfsUploadGifDocumentMedia{ClazzID: 0x41c4cd00} })   // 0x41c4cd00
	iface.RegisterClazzID(0xa2a4f818, func() iface.TLObject { return &TLDfsUploadMp4DocumentMedia{ClazzID: 0xa2a4f818} })   // 0xa2a4f818
	iface.RegisterClazzID(0xc1a61056, func() iface.TLObject { return &TLDfsUploadWallPaperFile{ClazzID: 0xc1a61056} })      // 0xc1a61056
	iface.RegisterClazzID(0xdea64f97, func() iface.TLObject { return &TLDfsUploadThemeFile{ClazzID: 0xdea64f97} })          // 0xdea64f97
	iface.RegisterClazzID(0x2b3c5b1, func() iface.TLObject { return &TLDfsUploadRingtoneFile{ClazzID: 0x2b3c5b1} })         // 0x2b3c5b1
	iface.RegisterClazzID(0xa3aa2874, func() iface.TLObject { return &TLDfsUploadedProfilePhoto{ClazzID: 0xa3aa2874} })     // 0xa3aa2874
}
