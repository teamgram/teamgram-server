/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package gateway

const (
	Predicate_gateway_sendDataToGateway = "gateway_sendDataToGateway"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_gateway_sendDataToGateway: {
		0: 645953552, // 0x26807810

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	645953552: Predicate_gateway_sendDataToGateway, // 0x26807810

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
