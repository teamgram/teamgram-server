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

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_inputId                   = "inputId"
	ClazzName_inputIds                  = "inputIds"
	ClazzName_inputSeqId                = "inputSeqId"
	ClazzName_inputNSeqId               = "inputNSeqId"
	ClazzName_idVal                     = "idVal"
	ClazzName_idVals                    = "idVals"
	ClazzName_seqIdVal                  = "seqIdVal"
	ClazzName_idgen_nextId              = "idgen_nextId"
	ClazzName_idgen_nextIds             = "idgen_nextIds"
	ClazzName_idgen_getCurrentSeqId     = "idgen_getCurrentSeqId"
	ClazzName_idgen_setCurrentSeqId     = "idgen_setCurrentSeqId"
	ClazzName_idgen_getNextSeqId        = "idgen_getNextSeqId"
	ClazzName_idgen_getNextNSeqId       = "idgen_getNextNSeqId"
	ClazzName_idgen_getNextIdValList    = "idgen_getNextIdValList"
	ClazzName_idgen_getCurrentSeqIdList = "idgen_getCurrentSeqIdList"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_inputId, 0, 0x8af2196c)                   // 8af2196c
	iface.RegisterClazzName(ClazzName_inputIds, 0, 0x7f285fbc)                  // 7f285fbc
	iface.RegisterClazzName(ClazzName_inputSeqId, 0, 0xcd52bbcd)                // cd52bbcd
	iface.RegisterClazzName(ClazzName_inputNSeqId, 0, 0x7ab16d81)               // 7ab16d81
	iface.RegisterClazzName(ClazzName_idVal, 0, 0xc07844cb)                     // c07844cb
	iface.RegisterClazzName(ClazzName_idVals, 0, 0x1c3baa66)                    // 1c3baa66
	iface.RegisterClazzName(ClazzName_seqIdVal, 0, 0x2a047d08)                  // 2a047d08
	iface.RegisterClazzName(ClazzName_idgen_nextId, 0, 0xbe711020)              // be711020
	iface.RegisterClazzName(ClazzName_idgen_nextIds, 0, 0x47c56fae)             // 47c56fae
	iface.RegisterClazzName(ClazzName_idgen_getCurrentSeqId, 0, 0x9d5bab80)     // 9d5bab80
	iface.RegisterClazzName(ClazzName_idgen_setCurrentSeqId, 0, 0xcd2c196d)     // cd2c196d
	iface.RegisterClazzName(ClazzName_idgen_getNextSeqId, 0, 0xf6716968)        // f6716968
	iface.RegisterClazzName(ClazzName_idgen_getNextNSeqId, 0, 0xa7d4cc6e)       // a7d4cc6e
	iface.RegisterClazzName(ClazzName_idgen_getNextIdValList, 0, 0xaa85f137)    // aa85f137
	iface.RegisterClazzName(ClazzName_idgen_getCurrentSeqIdList, 0, 0xd229ae43) // d229ae43

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_inputId, 0x8af2196c)                   // 8af2196c
	iface.RegisterClazzIDName(ClazzName_inputIds, 0x7f285fbc)                  // 7f285fbc
	iface.RegisterClazzIDName(ClazzName_inputSeqId, 0xcd52bbcd)                // cd52bbcd
	iface.RegisterClazzIDName(ClazzName_inputNSeqId, 0x7ab16d81)               // 7ab16d81
	iface.RegisterClazzIDName(ClazzName_idVal, 0xc07844cb)                     // c07844cb
	iface.RegisterClazzIDName(ClazzName_idVals, 0x1c3baa66)                    // 1c3baa66
	iface.RegisterClazzIDName(ClazzName_seqIdVal, 0x2a047d08)                  // 2a047d08
	iface.RegisterClazzIDName(ClazzName_idgen_nextId, 0xbe711020)              // be711020
	iface.RegisterClazzIDName(ClazzName_idgen_nextIds, 0x47c56fae)             // 47c56fae
	iface.RegisterClazzIDName(ClazzName_idgen_getCurrentSeqId, 0x9d5bab80)     // 9d5bab80
	iface.RegisterClazzIDName(ClazzName_idgen_setCurrentSeqId, 0xcd2c196d)     // cd2c196d
	iface.RegisterClazzIDName(ClazzName_idgen_getNextSeqId, 0xf6716968)        // f6716968
	iface.RegisterClazzIDName(ClazzName_idgen_getNextNSeqId, 0xa7d4cc6e)       // a7d4cc6e
	iface.RegisterClazzIDName(ClazzName_idgen_getNextIdValList, 0xaa85f137)    // aa85f137
	iface.RegisterClazzIDName(ClazzName_idgen_getCurrentSeqIdList, 0xd229ae43) // d229ae43
}
