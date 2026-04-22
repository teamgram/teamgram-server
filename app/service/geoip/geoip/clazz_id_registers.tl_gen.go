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

package geoip

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0xca964a2f, func() iface.TLObject { return &TLRegion{ClazzID: 0xca964a2f} }) // 0xca964a2f

	// Method
	iface.RegisterClazzID(0x676b04a4, func() iface.TLObject { return &TLGeoipGetCountryAndRegionByIp{ClazzID: 0x676b04a4} }) // 0x676b04a4
}
