/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package echo

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_echo      = "echo"
	ClazzName_echo2     = "echo2"
	ClazzName_echo_echo = "echo_echo"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_echo, 0, 0x2e3ba51e)      // 2e3ba51e
	iface.RegisterClazzName(ClazzName_echo2, 0, 0x2249c1b)      // 2249c1b
	iface.RegisterClazzName(ClazzName_echo_echo, 0, 0xf653b67d) // f653b67d

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_echo, 0x2e3ba51e)      // 2e3ba51e
	iface.RegisterClazzIDName(ClazzName_echo2, 0x2249c1b)      // 2249c1b
	iface.RegisterClazzIDName(ClazzName_echo_echo, 0xf653b67d) // f653b67d
}
