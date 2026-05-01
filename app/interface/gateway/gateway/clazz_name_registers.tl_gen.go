/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package gateway

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

const (
	ClazzName_gateway_pushUpdatesData        = "gateway_pushUpdatesData"
	ClazzName_gateway_pushSessionUpdatesData = "gateway_pushSessionUpdatesData"
	ClazzName_gateway_pushRpcResultData      = "gateway_pushRpcResultData"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_gateway_pushUpdatesData, 0, 0x10dcca87)        // 10dcca87
	iface.RegisterClazzName(ClazzName_gateway_pushSessionUpdatesData, 0, 0x794c7ded) // 794c7ded
	iface.RegisterClazzName(ClazzName_gateway_pushRpcResultData, 0, 0xfc960f5)       // fc960f5

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_gateway_pushUpdatesData, 0x10dcca87)        // 10dcca87
	iface.RegisterClazzIDName(ClazzName_gateway_pushSessionUpdatesData, 0x794c7ded) // 794c7ded
	iface.RegisterClazzIDName(ClazzName_gateway_pushRpcResultData, 0xfc960f5)       // fc960f5
}
