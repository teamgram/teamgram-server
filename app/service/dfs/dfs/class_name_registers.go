/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package dfs

const (
	Predicate_dfs_writeFilePartData        = "dfs_writeFilePartData"
	Predicate_dfs_uploadPhotoFileV2        = "dfs_uploadPhotoFileV2"
	Predicate_dfs_uploadProfilePhotoFileV2 = "dfs_uploadProfilePhotoFileV2"
	Predicate_dfs_uploadEncryptedFileV2    = "dfs_uploadEncryptedFileV2"
	Predicate_dfs_downloadFile             = "dfs_downloadFile"
	Predicate_dfs_uploadDocumentFileV2     = "dfs_uploadDocumentFileV2"
	Predicate_dfs_uploadGifDocumentMedia   = "dfs_uploadGifDocumentMedia"
	Predicate_dfs_uploadMp4DocumentMedia   = "dfs_uploadMp4DocumentMedia"
	Predicate_dfs_uploadWallPaperFile      = "dfs_uploadWallPaperFile"
	Predicate_dfs_uploadThemeFile          = "dfs_uploadThemeFile"
	Predicate_dfs_uploadRingtoneFile       = "dfs_uploadRingtoneFile"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_dfs_writeFilePartData: {
		0: 440942855, // 0x1a484107

	},
	Predicate_dfs_uploadPhotoFileV2: {
		0: 605082018, // 0x2410d1a2

	},
	Predicate_dfs_uploadProfilePhotoFileV2: {
		0: -870473038, // 0xcc1da2b2

	},
	Predicate_dfs_uploadEncryptedFileV2: {
		0: 2043921699, // 0x79d3c523

	},
	Predicate_dfs_downloadFile: {
		0: -692064706, // 0xd6bfee3e

	},
	Predicate_dfs_uploadDocumentFileV2: {
		0: 1983081911, // 0x76336db7

	},
	Predicate_dfs_uploadGifDocumentMedia: {
		0: 1103416576, // 0x41c4cd00

	},
	Predicate_dfs_uploadMp4DocumentMedia: {
		0: -1566246888, // 0xa2a4f818

	},
	Predicate_dfs_uploadWallPaperFile: {
		0: -1046081450, // 0xc1a61056

	},
	Predicate_dfs_uploadThemeFile: {
		0: -559525993, // 0xdea64f97

	},
	Predicate_dfs_uploadRingtoneFile: {
		0: 45335985, // 0x2b3c5b1

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	440942855:   Predicate_dfs_writeFilePartData,        // 0x1a484107
	605082018:   Predicate_dfs_uploadPhotoFileV2,        // 0x2410d1a2
	-870473038:  Predicate_dfs_uploadProfilePhotoFileV2, // 0xcc1da2b2
	2043921699:  Predicate_dfs_uploadEncryptedFileV2,    // 0x79d3c523
	-692064706:  Predicate_dfs_downloadFile,             // 0xd6bfee3e
	1983081911:  Predicate_dfs_uploadDocumentFileV2,     // 0x76336db7
	1103416576:  Predicate_dfs_uploadGifDocumentMedia,   // 0x41c4cd00
	-1566246888: Predicate_dfs_uploadMp4DocumentMedia,   // 0xa2a4f818
	-1046081450: Predicate_dfs_uploadWallPaperFile,      // 0xc1a61056
	-559525993:  Predicate_dfs_uploadThemeFile,          // 0xdea64f97
	45335985:    Predicate_dfs_uploadRingtoneFile,       // 0x2b3c5b1

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
