// Copyright 2022 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package hashx

import (
	"fmt"
	"testing"
)

func TestCombineInt64Hash(t *testing.T) {
	var acc int64 = 0
	for i := 1; i < 2; i++ {
		acc = CombineInt64Hash(acc, int64(i))
	}
	fmt.Println(acc)

	var acc2 int64 = 0
	for i := 1; i < 2; i++ {
		acc2 = CombineInt64Hash2(acc2, int64(i))
	}
	fmt.Println(acc2)
}
