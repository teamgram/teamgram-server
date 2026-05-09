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

package mediaprocessor

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0xfb5d44f8, func() iface.TLObject { return &TLProcessedDocument{ClazzID: 0xfb5d44f8} })   // 0xfb5d44f8
	iface.RegisterClazzID(0x606d445, func() iface.TLObject { return &TLProcessedPhoto{ClazzID: 0x606d445} })        // 0x606d445
	iface.RegisterClazzID(0x2af751de, func() iface.TLObject { return &TLProcessorDerivative{ClazzID: 0x2af751de} }) // 0x2af751de

	// Method
	iface.RegisterClazzID(0x23289b04, func() iface.TLObject { return &TLMediaProcessorProcessPhoto{ClazzID: 0x23289b04} }) // 0x23289b04
	iface.RegisterClazzID(0xcaa60c8c, func() iface.TLObject { return &TLMediaProcessorProcessGif{ClazzID: 0xcaa60c8c} })   // 0xcaa60c8c
	iface.RegisterClazzID(0xac180ca1, func() iface.TLObject { return &TLMediaProcessorProcessMp4{ClazzID: 0xac180ca1} })   // 0xac180ca1
}
