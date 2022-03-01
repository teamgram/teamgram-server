/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teagramio (teagram.io@gmail.com)
 */

package webpage

const (
	Predicate_webpage_getPendingWebPagePreview = "webpage_getPendingWebPagePreview"
	Predicate_webpage_getWebPagePreview        = "webpage_getWebPagePreview"
	Predicate_webpage_getWebPage               = "webpage_getWebPage"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_webpage_getPendingWebPagePreview: {
		0: 1074946247, // 0x401260c7

	},
	Predicate_webpage_getWebPagePreview: {
		0: -2059356164, // 0x8540b7fc

	},
	Predicate_webpage_getWebPage: {
		0: -142855528, // 0xf77c3298

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	1074946247:  Predicate_webpage_getPendingWebPagePreview, // 0x401260c7
	-2059356164: Predicate_webpage_getWebPagePreview,        // 0x8540b7fc
	-142855528:  Predicate_webpage_getWebPage,               // 0xf77c3298

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
