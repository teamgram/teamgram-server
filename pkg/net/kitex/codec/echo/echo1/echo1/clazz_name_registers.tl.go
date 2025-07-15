/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package echo1

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_echo       = "echo"
	ClazzName_echo1_echo = "echo1_echo"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_echo, 0, 0x2e3ba51e)       // 2e3ba51e
	iface.RegisterClazzName(ClazzName_echo1_echo, 0, 0x9f0506e2) // 9f0506e2

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_echo, 0x2e3ba51e)       // 2e3ba51e
	iface.RegisterClazzIDName(ClazzName_echo1_echo, 0x9f0506e2) // 9f0506e2
}
