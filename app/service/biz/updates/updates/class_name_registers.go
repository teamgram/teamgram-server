/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package updates

const (
	Predicate_channelDifference              = "channelDifference"
	Predicate_differenceEmpty                = "differenceEmpty"
	Predicate_difference                     = "difference"
	Predicate_differenceSlice                = "differenceSlice"
	Predicate_differenceTooLong              = "differenceTooLong"
	Predicate_updates_getStateV2             = "updates_getStateV2"
	Predicate_updates_getDifferenceV2        = "updates_getDifferenceV2"
	Predicate_updates_getChannelDifferenceV2 = "updates_getChannelDifferenceV2"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_channelDifference: {
		0: -853998774, // 0xcd19034a

	},
	Predicate_differenceEmpty: {
		0: -1948526002, // 0x8bdbda4e

	},
	Predicate_difference: {
		0: 1417839403, // 0x5482832b

	},
	Predicate_differenceSlice: {
		0: -879338017, // 0xcb965ddf

	},
	Predicate_differenceTooLong: {
		0: 896724528, // 0x3572ee30

	},
	Predicate_updates_getStateV2: {
		0: 1173671269, // 0x45f4cd65

	},
	Predicate_updates_getDifferenceV2: {
		0: -1217698151, // 0xb76b6699

	},
	Predicate_updates_getChannelDifferenceV2: {
		0: 1302540682, // 0x4da3318a

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	-853998774:  Predicate_channelDifference,              // 0xcd19034a
	-1948526002: Predicate_differenceEmpty,                // 0x8bdbda4e
	1417839403:  Predicate_difference,                     // 0x5482832b
	-879338017:  Predicate_differenceSlice,                // 0xcb965ddf
	896724528:   Predicate_differenceTooLong,              // 0x3572ee30
	1173671269:  Predicate_updates_getStateV2,             // 0x45f4cd65
	-1217698151: Predicate_updates_getDifferenceV2,        // 0xb76b6699
	1302540682:  Predicate_updates_getChannelDifferenceV2, // 0x4da3318a

}

func GetClazzID(clazzName string, layer int) int32 {
	if m, ok := clazzNameRegisters2[clazzName]; ok {
		m2, ok2 := m[layer]
		if ok2 {
			return m2
		}
		m2, ok2 = m[0]
		if ok2 {
			return m2
		}
	}
	return 0
}
