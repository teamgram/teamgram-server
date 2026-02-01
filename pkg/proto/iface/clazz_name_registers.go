// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package iface

var (
	clazzNameRegisters2   = make(map[string]map[int]uint32)
	clazzIdNameRegisters2 = make(map[uint32]string)
)

func RegisterClazzName(clazzName string, layer int, clazzId uint32) {
	if _, ok := clazzNameRegisters2[clazzName]; !ok {
		clazzNameRegisters2[clazzName] = make(map[int]uint32)
	}
	clazzNameRegisters2[clazzName][layer] = clazzId
	clazzIdNameRegisters2[clazzId] = clazzName
}

func GetClazzIDByName(clazzName string, layer int) uint32 {
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

func RegisterClazzIDName(clazzName string, clazzId uint32) {
	clazzIdNameRegisters2[clazzId] = clazzName
}

func GetClazzNameByID(clazzId uint32) string {
	if clazzName, ok := clazzIdNameRegisters2[clazzId]; ok {
		return clazzName
	}
	return ""
}
