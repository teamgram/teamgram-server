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
