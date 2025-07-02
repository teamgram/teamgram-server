/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package updates

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_channelDifference              = "channelDifference"
	ClazzName_differenceEmpty                = "differenceEmpty"
	ClazzName_difference                     = "difference"
	ClazzName_differenceSlice                = "differenceSlice"
	ClazzName_differenceTooLong              = "differenceTooLong"
	ClazzName_updates_getStateV2             = "updates_getStateV2"
	ClazzName_updates_getDifferenceV2        = "updates_getDifferenceV2"
	ClazzName_updates_getChannelDifferenceV2 = "updates_getChannelDifferenceV2"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_channelDifference, 0, 0xcd19034a)              // cd19034a
	iface.RegisterClazzName(ClazzName_differenceEmpty, 0, 0x8bdbda4e)                // 8bdbda4e
	iface.RegisterClazzName(ClazzName_difference, 0, 0x5482832b)                     // 5482832b
	iface.RegisterClazzName(ClazzName_differenceSlice, 0, 0xcb965ddf)                // cb965ddf
	iface.RegisterClazzName(ClazzName_differenceTooLong, 0, 0x3572ee30)              // 3572ee30
	iface.RegisterClazzName(ClazzName_updates_getStateV2, 0, 0x45f4cd65)             // 45f4cd65
	iface.RegisterClazzName(ClazzName_updates_getDifferenceV2, 0, 0xb76b6699)        // b76b6699
	iface.RegisterClazzName(ClazzName_updates_getChannelDifferenceV2, 0, 0x4da3318a) // 4da3318a

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_channelDifference, 0xcd19034a)              // cd19034a
	iface.RegisterClazzIDName(ClazzName_differenceEmpty, 0x8bdbda4e)                // 8bdbda4e
	iface.RegisterClazzIDName(ClazzName_difference, 0x5482832b)                     // 5482832b
	iface.RegisterClazzIDName(ClazzName_differenceSlice, 0xcb965ddf)                // cb965ddf
	iface.RegisterClazzIDName(ClazzName_differenceTooLong, 0x3572ee30)              // 3572ee30
	iface.RegisterClazzIDName(ClazzName_updates_getStateV2, 0x45f4cd65)             // 45f4cd65
	iface.RegisterClazzIDName(ClazzName_updates_getDifferenceV2, 0xb76b6699)        // b76b6699
	iface.RegisterClazzIDName(ClazzName_updates_getChannelDifferenceV2, 0x4da3318a) // 4da3318a
}
