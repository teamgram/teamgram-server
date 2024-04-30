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
	CRC32_UNKNOWN                        TLConstructor = 0
	CRC32_channelDifference              TLConstructor = -853998774  // 0xcd19034a
	CRC32_differenceEmpty                TLConstructor = -1948526002 // 0x8bdbda4e
	CRC32_difference                     TLConstructor = 1417839403  // 0x5482832b
	CRC32_differenceSlice                TLConstructor = -879338017  // 0xcb965ddf
	CRC32_differenceTooLong              TLConstructor = 896724528   // 0x3572ee30
	CRC32_updates_getStateV2             TLConstructor = 1173671269  // 0x45f4cd65
	CRC32_updates_getDifferenceV2        TLConstructor = -1217698151 // 0xb76b6699
	CRC32_updates_getChannelDifferenceV2 TLConstructor = 1302540682  // 0x4da3318a
)
