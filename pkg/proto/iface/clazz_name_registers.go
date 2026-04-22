// Copyright (c) 2026 The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
