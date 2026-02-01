// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package iface

var (
	clazzIdRegisters2 = make(map[uint32]func() TLObject)
)

func RegisterClazzID(clazzId uint32, f func() TLObject) {
	clazzIdRegisters2[clazzId] = f
}

func NewTLObjectByClazzID(clazzId uint32) TLObject {
	f, ok := clazzIdRegisters2[clazzId]
	if !ok {
		return nil
	}
	return f()
}

func CheckClazzID(clazzId uint32) (ok bool) {
	_, ok = clazzIdRegisters2[clazzId]
	return
}
