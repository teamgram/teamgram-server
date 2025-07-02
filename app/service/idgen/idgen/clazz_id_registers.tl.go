/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

// ConstructorList
// RequestList

package idgen

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0xc07844cb, func() iface.TLObject { return &TLIdVal{ClazzID: 0xc07844cb} })       // 0xc07844cb
	iface.RegisterClazzID(0x1c3baa66, func() iface.TLObject { return &TLIdVals{ClazzID: 0x1c3baa66} })      // 0x1c3baa66
	iface.RegisterClazzID(0x2a047d08, func() iface.TLObject { return &TLSeqIdVal{ClazzID: 0x2a047d08} })    // 0x2a047d08
	iface.RegisterClazzID(0x8af2196c, func() iface.TLObject { return &TLInputId{ClazzID: 0x8af2196c} })     // 0x8af2196c
	iface.RegisterClazzID(0x7f285fbc, func() iface.TLObject { return &TLInputIds{ClazzID: 0x7f285fbc} })    // 0x7f285fbc
	iface.RegisterClazzID(0xcd52bbcd, func() iface.TLObject { return &TLInputSeqId{ClazzID: 0xcd52bbcd} })  // 0xcd52bbcd
	iface.RegisterClazzID(0x7ab16d81, func() iface.TLObject { return &TLInputNSeqId{ClazzID: 0x7ab16d81} }) // 0x7ab16d81

	// Method
	iface.RegisterClazzID(0xbe711020, func() iface.TLObject { return &TLIdgenNextId{ClazzID: 0xbe711020} })              // 0xbe711020
	iface.RegisterClazzID(0x47c56fae, func() iface.TLObject { return &TLIdgenNextIds{ClazzID: 0x47c56fae} })             // 0x47c56fae
	iface.RegisterClazzID(0x9d5bab80, func() iface.TLObject { return &TLIdgenGetCurrentSeqId{ClazzID: 0x9d5bab80} })     // 0x9d5bab80
	iface.RegisterClazzID(0xcd2c196d, func() iface.TLObject { return &TLIdgenSetCurrentSeqId{ClazzID: 0xcd2c196d} })     // 0xcd2c196d
	iface.RegisterClazzID(0xf6716968, func() iface.TLObject { return &TLIdgenGetNextSeqId{ClazzID: 0xf6716968} })        // 0xf6716968
	iface.RegisterClazzID(0xa7d4cc6e, func() iface.TLObject { return &TLIdgenGetNextNSeqId{ClazzID: 0xa7d4cc6e} })       // 0xa7d4cc6e
	iface.RegisterClazzID(0xaa85f137, func() iface.TLObject { return &TLIdgenGetNextIdValList{ClazzID: 0xaa85f137} })    // 0xaa85f137
	iface.RegisterClazzID(0xd229ae43, func() iface.TLObject { return &TLIdgenGetCurrentSeqIdList{ClazzID: 0xd229ae43} }) // 0xd229ae43
}
