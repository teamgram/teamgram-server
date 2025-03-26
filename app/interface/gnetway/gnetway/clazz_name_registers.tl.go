/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package gnetway

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_gnetway_sendDataToGateway = "gnetway_sendDataToGateway"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_gnetway_sendDataToGateway, 0, 0x722d5ce0) // 722d5ce0

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_gnetway_sendDataToGateway, 0x722d5ce0) // 722d5ce0
}
