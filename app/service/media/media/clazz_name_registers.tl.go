/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package media

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_photoSizeList                = "photoSizeList"
	ClazzName_videoSizeList                = "videoSizeList"
	ClazzName_media_uploadPhotoFile        = "media_uploadPhotoFile"
	ClazzName_media_uploadProfilePhotoFile = "media_uploadProfilePhotoFile"
	ClazzName_media_getPhoto               = "media_getPhoto"
	ClazzName_media_getPhotoSizeList       = "media_getPhotoSizeList"
	ClazzName_media_getPhotoSizeListList   = "media_getPhotoSizeListList"
	ClazzName_media_getVideoSizeList       = "media_getVideoSizeList"
	ClazzName_media_uploadedDocumentMedia  = "media_uploadedDocumentMedia"
	ClazzName_media_getDocument            = "media_getDocument"
	ClazzName_media_getDocumentList        = "media_getDocumentList"
	ClazzName_media_uploadEncryptedFile    = "media_uploadEncryptedFile"
	ClazzName_media_getEncryptedFile       = "media_getEncryptedFile"
	ClazzName_media_uploadWallPaperFile    = "media_uploadWallPaperFile"
	ClazzName_media_uploadThemeFile        = "media_uploadThemeFile"
	ClazzName_media_uploadStickerFile      = "media_uploadStickerFile"
	ClazzName_media_uploadRingtoneFile     = "media_uploadRingtoneFile"
	ClazzName_media_uploadedProfilePhoto   = "media_uploadedProfilePhoto"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_photoSizeList, 0, 0x67139b3)                 // 67139b3
	iface.RegisterClazzName(ClazzName_videoSizeList, 0, 0x38d19bf2)                // 38d19bf2
	iface.RegisterClazzName(ClazzName_media_uploadPhotoFile, 0, 0x3c2b0b17)        // 3c2b0b17
	iface.RegisterClazzName(ClazzName_media_uploadProfilePhotoFile, 0, 0x973f2f24) // 973f2f24
	iface.RegisterClazzName(ClazzName_media_getPhoto, 0, 0x657eb86b)               // 657eb86b
	iface.RegisterClazzName(ClazzName_media_getPhotoSizeList, 0, 0xa1eb7f45)       // a1eb7f45
	iface.RegisterClazzName(ClazzName_media_getPhotoSizeListList, 0, 0xfb5c80e0)   // fb5c80e0
	iface.RegisterClazzName(ClazzName_media_getVideoSizeList, 0, 0xc47692ea)       // c47692ea
	iface.RegisterClazzName(ClazzName_media_uploadedDocumentMedia, 0, 0x4f5fb06c)  // 4f5fb06c
	iface.RegisterClazzName(ClazzName_media_getDocument, 0, 0x3fe5974d)            // 3fe5974d
	iface.RegisterClazzName(ClazzName_media_getDocumentList, 0, 0xc52fd26f)        // c52fd26f
	iface.RegisterClazzName(ClazzName_media_uploadEncryptedFile, 0, 0xab00c69b)    // ab00c69b
	iface.RegisterClazzName(ClazzName_media_getEncryptedFile, 0, 0xfc6080d1)       // fc6080d1
	iface.RegisterClazzName(ClazzName_media_uploadWallPaperFile, 0, 0x9cfaadfe)    // 9cfaadfe
	iface.RegisterClazzName(ClazzName_media_uploadThemeFile, 0, 0x42e6b860)        // 42e6b860
	iface.RegisterClazzName(ClazzName_media_uploadStickerFile, 0, 0xacb624ed)      // acb624ed
	iface.RegisterClazzName(ClazzName_media_uploadRingtoneFile, 0, 0x3dbab209)     // 3dbab209
	iface.RegisterClazzName(ClazzName_media_uploadedProfilePhoto, 0, 0x89d159d2)   // 89d159d2

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_photoSizeList, 0x67139b3)                 // 67139b3
	iface.RegisterClazzIDName(ClazzName_videoSizeList, 0x38d19bf2)                // 38d19bf2
	iface.RegisterClazzIDName(ClazzName_media_uploadPhotoFile, 0x3c2b0b17)        // 3c2b0b17
	iface.RegisterClazzIDName(ClazzName_media_uploadProfilePhotoFile, 0x973f2f24) // 973f2f24
	iface.RegisterClazzIDName(ClazzName_media_getPhoto, 0x657eb86b)               // 657eb86b
	iface.RegisterClazzIDName(ClazzName_media_getPhotoSizeList, 0xa1eb7f45)       // a1eb7f45
	iface.RegisterClazzIDName(ClazzName_media_getPhotoSizeListList, 0xfb5c80e0)   // fb5c80e0
	iface.RegisterClazzIDName(ClazzName_media_getVideoSizeList, 0xc47692ea)       // c47692ea
	iface.RegisterClazzIDName(ClazzName_media_uploadedDocumentMedia, 0x4f5fb06c)  // 4f5fb06c
	iface.RegisterClazzIDName(ClazzName_media_getDocument, 0x3fe5974d)            // 3fe5974d
	iface.RegisterClazzIDName(ClazzName_media_getDocumentList, 0xc52fd26f)        // c52fd26f
	iface.RegisterClazzIDName(ClazzName_media_uploadEncryptedFile, 0xab00c69b)    // ab00c69b
	iface.RegisterClazzIDName(ClazzName_media_getEncryptedFile, 0xfc6080d1)       // fc6080d1
	iface.RegisterClazzIDName(ClazzName_media_uploadWallPaperFile, 0x9cfaadfe)    // 9cfaadfe
	iface.RegisterClazzIDName(ClazzName_media_uploadThemeFile, 0x42e6b860)        // 42e6b860
	iface.RegisterClazzIDName(ClazzName_media_uploadStickerFile, 0xacb624ed)      // acb624ed
	iface.RegisterClazzIDName(ClazzName_media_uploadRingtoneFile, 0x3dbab209)     // 3dbab209
	iface.RegisterClazzIDName(ClazzName_media_uploadedProfilePhoto, 0x89d159d2)   // 89d159d2
}
