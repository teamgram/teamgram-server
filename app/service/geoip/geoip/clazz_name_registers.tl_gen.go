/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package geoip

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

const (
	ClazzName_region                        = "region"
	ClazzName_geoip_getCountryAndRegionByIp = "geoip_getCountryAndRegionByIp"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_region, 0, 0xca964a2f)                        // ca964a2f
	iface.RegisterClazzName(ClazzName_geoip_getCountryAndRegionByIp, 0, 0x676b04a4) // 676b04a4

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_region, 0xca964a2f)                        // ca964a2f
	iface.RegisterClazzIDName(ClazzName_geoip_getCountryAndRegionByIp, 0x676b04a4) // 676b04a4
}
