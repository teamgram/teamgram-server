/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package gnetway

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
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
