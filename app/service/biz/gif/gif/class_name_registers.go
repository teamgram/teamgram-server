/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teagramio (teagram.io@gmail.com)
 */

package gif

const (
	Predicate_gif_saveGif        = "gif_saveGif"
	Predicate_gif_getSavedGifs   = "gif_getSavedGifs"
	Predicate_gif_deleteSavedGif = "gif_deleteSavedGif"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_gif_saveGif: {
		0: 556825867, // 0x21307d0b

	},
	Predicate_gif_getSavedGifs: {
		0: 926787430, // 0x373da766

	},
	Predicate_gif_deleteSavedGif: {
		0: 523645139, // 0x1f3630d3

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	556825867: Predicate_gif_saveGif,        // 0x21307d0b
	926787430: Predicate_gif_getSavedGifs,   // 0x373da766
	523645139: Predicate_gif_deleteSavedGif, // 0x1f3630d3

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
