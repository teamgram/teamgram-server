// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

package auth_session

import (
	//"math/rand"
	//"testing"
)

// GetOrInsertSaltList
//func TestGetOrInsertSaltList(t *testing.T) {
//	id := rand.Int63()
//	salts, err := GetOrInsertSaltList(id, 32)
//	if err != nil {
//		t.Error(err)
//	} else {
//		t.Log(salts)
//	}
//
//	var salt int64 = 0
//	salt, err = GetOrInsertSalt(id)
//	t.Log(salt)
//
//	if CheckBySalt(id, salt) {
//		t.Logf("CheckBySalt(%d, %d) = true", id, salt)
//	}
//
//	if !CheckBySalt(id, 123) {
//		t.Logf("CheckBySalt(%d, 123) = false", id)
//	}
//}
//
//func BenchmarkGetSalt(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		val, _ := GetOrInsertSalt(rand.Int63())
//		_ = val
//	}
//}
