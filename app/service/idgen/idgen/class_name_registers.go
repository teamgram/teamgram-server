/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teagramio (teagram.io@gmail.com)
 */

package idgen

const (
	Predicate_idgen_nextId          = "idgen_nextId"
	Predicate_idgen_nextIds         = "idgen_nextIds"
	Predicate_idgen_getCurrentSeqId = "idgen_getCurrentSeqId"
	Predicate_idgen_setCurrentSeqId = "idgen_setCurrentSeqId"
	Predicate_idgen_getNextSeqId    = "idgen_getNextSeqId"
	Predicate_idgen_getNextNSeqId   = "idgen_getNextNSeqId"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_idgen_nextId: {
		0: -1099886560, // 0xbe711020

	},
	Predicate_idgen_nextIds: {
		0: 1204121518, // 0x47c56fae

	},
	Predicate_idgen_getCurrentSeqId: {
		0: -1654936704, // 0x9d5bab80

	},
	Predicate_idgen_setCurrentSeqId: {
		0: -852747923, // 0xcd2c196d

	},
	Predicate_idgen_getNextSeqId: {
		0: -160339608, // 0xf6716968

	},
	Predicate_idgen_getNextNSeqId: {
		0: -1479226258, // 0xa7d4cc6e

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	-1099886560: Predicate_idgen_nextId,          // 0xbe711020
	1204121518:  Predicate_idgen_nextIds,         // 0x47c56fae
	-1654936704: Predicate_idgen_getCurrentSeqId, // 0x9d5bab80
	-852747923:  Predicate_idgen_setCurrentSeqId, // 0xcd2c196d
	-160339608:  Predicate_idgen_getNextSeqId,    // 0xf6716968
	-1479226258: Predicate_idgen_getNextNSeqId,   // 0xa7d4cc6e

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
