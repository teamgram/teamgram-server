/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package media

const (
	Predicate_photoSizeList                = "photoSizeList"
	Predicate_videoSizeList                = "videoSizeList"
	Predicate_media_uploadPhotoFile        = "media_uploadPhotoFile"
	Predicate_media_uploadProfilePhotoFile = "media_uploadProfilePhotoFile"
	Predicate_media_getPhoto               = "media_getPhoto"
	Predicate_media_getPhotoSizeList       = "media_getPhotoSizeList"
	Predicate_media_getPhotoSizeListList   = "media_getPhotoSizeListList"
	Predicate_media_getVideoSizeList       = "media_getVideoSizeList"
	Predicate_media_uploadedDocumentMedia  = "media_uploadedDocumentMedia"
	Predicate_media_getDocument            = "media_getDocument"
	Predicate_media_getDocumentList        = "media_getDocumentList"
	Predicate_media_uploadEncryptedFile    = "media_uploadEncryptedFile"
	Predicate_media_getEncryptedFile       = "media_getEncryptedFile"
	Predicate_media_uploadWallPaperFile    = "media_uploadWallPaperFile"
	Predicate_media_uploadThemeFile        = "media_uploadThemeFile"
	Predicate_media_uploadStickerFile      = "media_uploadStickerFile"
	Predicate_media_uploadRingtoneFile     = "media_uploadRingtoneFile"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_photoSizeList: {
		0: 108083635, // 0x67139b3

	},
	Predicate_videoSizeList: {
		0: 953261042, // 0x38d19bf2

	},
	Predicate_media_uploadPhotoFile: {
		0: 1009453847, // 0x3c2b0b17

	},
	Predicate_media_uploadProfilePhotoFile: {
		0: -1757466844, // 0x973f2f24

	},
	Predicate_media_getPhoto: {
		0: 1702803563, // 0x657eb86b

	},
	Predicate_media_getPhotoSizeList: {
		0: -1578401979, // 0xa1eb7f45

	},
	Predicate_media_getPhotoSizeListList: {
		0: -77823776, // 0xfb5c80e0

	},
	Predicate_media_getVideoSizeList: {
		0: -998862102, // 0xc47692ea

	},
	Predicate_media_uploadedDocumentMedia: {
		0: 1331671148, // 0x4f5fb06c

	},
	Predicate_media_getDocument: {
		0: 1072011085, // 0x3fe5974d

	},
	Predicate_media_getDocumentList: {
		0: -986721681, // 0xc52fd26f

	},
	Predicate_media_uploadEncryptedFile: {
		0: -1426012517, // 0xab00c69b

	},
	Predicate_media_getEncryptedFile: {
		0: -60784431, // 0xfc6080d1

	},
	Predicate_media_uploadWallPaperFile: {
		0: -1661293058, // 0x9cfaadfe

	},
	Predicate_media_uploadThemeFile: {
		0: 1122416736, // 0x42e6b860

	},
	Predicate_media_uploadStickerFile: {
		0: -1397349139, // 0xacb624ed

	},
	Predicate_media_uploadRingtoneFile: {
		0: 1035645449, // 0x3dbab209

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	108083635:   Predicate_photoSizeList,                // 0x67139b3
	953261042:   Predicate_videoSizeList,                // 0x38d19bf2
	1009453847:  Predicate_media_uploadPhotoFile,        // 0x3c2b0b17
	-1757466844: Predicate_media_uploadProfilePhotoFile, // 0x973f2f24
	1702803563:  Predicate_media_getPhoto,               // 0x657eb86b
	-1578401979: Predicate_media_getPhotoSizeList,       // 0xa1eb7f45
	-77823776:   Predicate_media_getPhotoSizeListList,   // 0xfb5c80e0
	-998862102:  Predicate_media_getVideoSizeList,       // 0xc47692ea
	1331671148:  Predicate_media_uploadedDocumentMedia,  // 0x4f5fb06c
	1072011085:  Predicate_media_getDocument,            // 0x3fe5974d
	-986721681:  Predicate_media_getDocumentList,        // 0xc52fd26f
	-1426012517: Predicate_media_uploadEncryptedFile,    // 0xab00c69b
	-60784431:   Predicate_media_getEncryptedFile,       // 0xfc6080d1
	-1661293058: Predicate_media_uploadWallPaperFile,    // 0x9cfaadfe
	1122416736:  Predicate_media_uploadThemeFile,        // 0x42e6b860
	-1397349139: Predicate_media_uploadStickerFile,      // 0xacb624ed
	1035645449:  Predicate_media_uploadRingtoneFile,     // 0x3dbab209

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
