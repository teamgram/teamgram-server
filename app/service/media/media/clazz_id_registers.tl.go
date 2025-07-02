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

package media

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0x67139b3, func() iface.TLObject { return &TLPhotoSizeList{ClazzID: 0x67139b3} })   // 0x67139b3
	iface.RegisterClazzID(0x38d19bf2, func() iface.TLObject { return &TLVideoSizeList{ClazzID: 0x38d19bf2} }) // 0x38d19bf2

	// Method
	iface.RegisterClazzID(0x3c2b0b17, func() iface.TLObject { return &TLMediaUploadPhotoFile{ClazzID: 0x3c2b0b17} })        // 0x3c2b0b17
	iface.RegisterClazzID(0x973f2f24, func() iface.TLObject { return &TLMediaUploadProfilePhotoFile{ClazzID: 0x973f2f24} }) // 0x973f2f24
	iface.RegisterClazzID(0x657eb86b, func() iface.TLObject { return &TLMediaGetPhoto{ClazzID: 0x657eb86b} })               // 0x657eb86b
	iface.RegisterClazzID(0xa1eb7f45, func() iface.TLObject { return &TLMediaGetPhotoSizeList{ClazzID: 0xa1eb7f45} })       // 0xa1eb7f45
	iface.RegisterClazzID(0xfb5c80e0, func() iface.TLObject { return &TLMediaGetPhotoSizeListList{ClazzID: 0xfb5c80e0} })   // 0xfb5c80e0
	iface.RegisterClazzID(0xc47692ea, func() iface.TLObject { return &TLMediaGetVideoSizeList{ClazzID: 0xc47692ea} })       // 0xc47692ea
	iface.RegisterClazzID(0x4f5fb06c, func() iface.TLObject { return &TLMediaUploadedDocumentMedia{ClazzID: 0x4f5fb06c} })  // 0x4f5fb06c
	iface.RegisterClazzID(0x3fe5974d, func() iface.TLObject { return &TLMediaGetDocument{ClazzID: 0x3fe5974d} })            // 0x3fe5974d
	iface.RegisterClazzID(0xc52fd26f, func() iface.TLObject { return &TLMediaGetDocumentList{ClazzID: 0xc52fd26f} })        // 0xc52fd26f
	iface.RegisterClazzID(0xab00c69b, func() iface.TLObject { return &TLMediaUploadEncryptedFile{ClazzID: 0xab00c69b} })    // 0xab00c69b
	iface.RegisterClazzID(0xfc6080d1, func() iface.TLObject { return &TLMediaGetEncryptedFile{ClazzID: 0xfc6080d1} })       // 0xfc6080d1
	iface.RegisterClazzID(0x9cfaadfe, func() iface.TLObject { return &TLMediaUploadWallPaperFile{ClazzID: 0x9cfaadfe} })    // 0x9cfaadfe
	iface.RegisterClazzID(0x42e6b860, func() iface.TLObject { return &TLMediaUploadThemeFile{ClazzID: 0x42e6b860} })        // 0x42e6b860
	iface.RegisterClazzID(0xacb624ed, func() iface.TLObject { return &TLMediaUploadStickerFile{ClazzID: 0xacb624ed} })      // 0xacb624ed
	iface.RegisterClazzID(0x3dbab209, func() iface.TLObject { return &TLMediaUploadRingtoneFile{ClazzID: 0x3dbab209} })     // 0x3dbab209
	iface.RegisterClazzID(0x89d159d2, func() iface.TLObject { return &TLMediaUploadedProfilePhoto{ClazzID: 0x89d159d2} })   // 0x89d159d2
}
