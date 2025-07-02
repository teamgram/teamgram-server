/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package dfs

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_dfs_writeFilePartData        = "dfs_writeFilePartData"
	ClazzName_dfs_uploadPhotoFileV2        = "dfs_uploadPhotoFileV2"
	ClazzName_dfs_uploadProfilePhotoFileV2 = "dfs_uploadProfilePhotoFileV2"
	ClazzName_dfs_uploadEncryptedFileV2    = "dfs_uploadEncryptedFileV2"
	ClazzName_dfs_downloadFile             = "dfs_downloadFile"
	ClazzName_dfs_uploadDocumentFileV2     = "dfs_uploadDocumentFileV2"
	ClazzName_dfs_uploadGifDocumentMedia   = "dfs_uploadGifDocumentMedia"
	ClazzName_dfs_uploadMp4DocumentMedia   = "dfs_uploadMp4DocumentMedia"
	ClazzName_dfs_uploadWallPaperFile      = "dfs_uploadWallPaperFile"
	ClazzName_dfs_uploadThemeFile          = "dfs_uploadThemeFile"
	ClazzName_dfs_uploadRingtoneFile       = "dfs_uploadRingtoneFile"
	ClazzName_dfs_uploadedProfilePhoto     = "dfs_uploadedProfilePhoto"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_dfs_writeFilePartData, 0, 0x1a484107)        // 1a484107
	iface.RegisterClazzName(ClazzName_dfs_uploadPhotoFileV2, 0, 0x2410d1a2)        // 2410d1a2
	iface.RegisterClazzName(ClazzName_dfs_uploadProfilePhotoFileV2, 0, 0xcc1da2b2) // cc1da2b2
	iface.RegisterClazzName(ClazzName_dfs_uploadEncryptedFileV2, 0, 0x79d3c523)    // 79d3c523
	iface.RegisterClazzName(ClazzName_dfs_downloadFile, 0, 0xd6bfee3e)             // d6bfee3e
	iface.RegisterClazzName(ClazzName_dfs_uploadDocumentFileV2, 0, 0x76336db7)     // 76336db7
	iface.RegisterClazzName(ClazzName_dfs_uploadGifDocumentMedia, 0, 0x41c4cd00)   // 41c4cd00
	iface.RegisterClazzName(ClazzName_dfs_uploadMp4DocumentMedia, 0, 0xa2a4f818)   // a2a4f818
	iface.RegisterClazzName(ClazzName_dfs_uploadWallPaperFile, 0, 0xc1a61056)      // c1a61056
	iface.RegisterClazzName(ClazzName_dfs_uploadThemeFile, 0, 0xdea64f97)          // dea64f97
	iface.RegisterClazzName(ClazzName_dfs_uploadRingtoneFile, 0, 0x2b3c5b1)        // 2b3c5b1
	iface.RegisterClazzName(ClazzName_dfs_uploadedProfilePhoto, 0, 0xa3aa2874)     // a3aa2874

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_dfs_writeFilePartData, 0x1a484107)        // 1a484107
	iface.RegisterClazzIDName(ClazzName_dfs_uploadPhotoFileV2, 0x2410d1a2)        // 2410d1a2
	iface.RegisterClazzIDName(ClazzName_dfs_uploadProfilePhotoFileV2, 0xcc1da2b2) // cc1da2b2
	iface.RegisterClazzIDName(ClazzName_dfs_uploadEncryptedFileV2, 0x79d3c523)    // 79d3c523
	iface.RegisterClazzIDName(ClazzName_dfs_downloadFile, 0xd6bfee3e)             // d6bfee3e
	iface.RegisterClazzIDName(ClazzName_dfs_uploadDocumentFileV2, 0x76336db7)     // 76336db7
	iface.RegisterClazzIDName(ClazzName_dfs_uploadGifDocumentMedia, 0x41c4cd00)   // 41c4cd00
	iface.RegisterClazzIDName(ClazzName_dfs_uploadMp4DocumentMedia, 0xa2a4f818)   // a2a4f818
	iface.RegisterClazzIDName(ClazzName_dfs_uploadWallPaperFile, 0xc1a61056)      // c1a61056
	iface.RegisterClazzIDName(ClazzName_dfs_uploadThemeFile, 0xdea64f97)          // dea64f97
	iface.RegisterClazzIDName(ClazzName_dfs_uploadRingtoneFile, 0x2b3c5b1)        // 2b3c5b1
	iface.RegisterClazzIDName(ClazzName_dfs_uploadedProfilePhoto, 0xa3aa2874)     // a3aa2874
}
