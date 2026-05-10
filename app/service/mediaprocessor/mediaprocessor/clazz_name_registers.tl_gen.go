/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mediaprocessor

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

const (
	ClazzName_processorDerivative         = "processorDerivative"
	ClazzName_processedPhoto              = "processedPhoto"
	ClazzName_processedDocument           = "processedDocument"
	ClazzName_mediaProcessor_processPhoto = "mediaProcessor_processPhoto"
	ClazzName_mediaProcessor_processGif   = "mediaProcessor_processGif"
	ClazzName_mediaProcessor_processMp4   = "mediaProcessor_processMp4"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_processorDerivative, 0, 0x9ef0eecd)         // 9ef0eecd
	iface.RegisterClazzName(ClazzName_processedPhoto, 0, 0x606d445)               // 606d445
	iface.RegisterClazzName(ClazzName_processedDocument, 0, 0xfb5d44f8)           // fb5d44f8
	iface.RegisterClazzName(ClazzName_mediaProcessor_processPhoto, 0, 0x23289b04) // 23289b04
	iface.RegisterClazzName(ClazzName_mediaProcessor_processGif, 0, 0xcaa60c8c)   // caa60c8c
	iface.RegisterClazzName(ClazzName_mediaProcessor_processMp4, 0, 0xac180ca1)   // ac180ca1

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_processorDerivative, 0x9ef0eecd)         // 9ef0eecd
	iface.RegisterClazzIDName(ClazzName_processedPhoto, 0x606d445)               // 606d445
	iface.RegisterClazzIDName(ClazzName_processedDocument, 0xfb5d44f8)           // fb5d44f8
	iface.RegisterClazzIDName(ClazzName_mediaProcessor_processPhoto, 0x23289b04) // 23289b04
	iface.RegisterClazzIDName(ClazzName_mediaProcessor_processGif, 0xcaa60c8c)   // caa60c8c
	iface.RegisterClazzIDName(ClazzName_mediaProcessor_processMp4, 0xac180ca1)   // ac180ca1
}
