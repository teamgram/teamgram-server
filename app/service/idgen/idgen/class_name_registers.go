/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package idgen

const (
	Predicate_inputId                   = "inputId"
	Predicate_inputIds                  = "inputIds"
	Predicate_inputSeqId                = "inputSeqId"
	Predicate_inputNSeqId               = "inputNSeqId"
	Predicate_idVal                     = "idVal"
	Predicate_idVals                    = "idVals"
	Predicate_seqIdVal                  = "seqIdVal"
	Predicate_idgen_nextId              = "idgen_nextId"
	Predicate_idgen_nextIds             = "idgen_nextIds"
	Predicate_idgen_getCurrentSeqId     = "idgen_getCurrentSeqId"
	Predicate_idgen_setCurrentSeqId     = "idgen_setCurrentSeqId"
	Predicate_idgen_getNextSeqId        = "idgen_getNextSeqId"
	Predicate_idgen_getNextNSeqId       = "idgen_getNextNSeqId"
	Predicate_idgen_getNextIdValList    = "idgen_getNextIdValList"
	Predicate_idgen_getCurrentSeqIdList = "idgen_getCurrentSeqIdList"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_inputId: {
		0: -1963845268, // 0x8af2196c

	},
	Predicate_inputIds: {
		0: 2133352380, // 0x7f285fbc

	},
	Predicate_inputSeqId: {
		0: -850215987, // 0xcd52bbcd

	},
	Predicate_inputNSeqId: {
		0: 2058448257, // 0x7ab16d81

	},
	Predicate_idVal: {
		0: -1065859893, // 0xc07844cb

	},
	Predicate_idVals: {
		0: 473672294, // 0x1c3baa66

	},
	Predicate_seqIdVal: {
		0: 704937224, // 0x2a047d08

	},
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
	Predicate_idgen_getNextIdValList: {
		0: -1434062537, // 0xaa85f137

	},
	Predicate_idgen_getCurrentSeqIdList: {
		0: -769020349, // 0xd229ae43

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	-1963845268: Predicate_inputId,                   // 0x8af2196c
	2133352380:  Predicate_inputIds,                  // 0x7f285fbc
	-850215987:  Predicate_inputSeqId,                // 0xcd52bbcd
	2058448257:  Predicate_inputNSeqId,               // 0x7ab16d81
	-1065859893: Predicate_idVal,                     // 0xc07844cb
	473672294:   Predicate_idVals,                    // 0x1c3baa66
	704937224:   Predicate_seqIdVal,                  // 0x2a047d08
	-1099886560: Predicate_idgen_nextId,              // 0xbe711020
	1204121518:  Predicate_idgen_nextIds,             // 0x47c56fae
	-1654936704: Predicate_idgen_getCurrentSeqId,     // 0x9d5bab80
	-852747923:  Predicate_idgen_setCurrentSeqId,     // 0xcd2c196d
	-160339608:  Predicate_idgen_getNextSeqId,        // 0xf6716968
	-1479226258: Predicate_idgen_getNextNSeqId,       // 0xa7d4cc6e
	-1434062537: Predicate_idgen_getNextIdValList,    // 0xaa85f137
	-769020349:  Predicate_idgen_getCurrentSeqIdList, // 0xd229ae43

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
