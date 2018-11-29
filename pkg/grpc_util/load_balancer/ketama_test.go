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

package load_balancer

import (
	"fmt"
	"testing"
)

func TestKetama(t *testing.T) {
	k := NewKetama(10, nil)
	k.Add("127.0.0.1:10000")
	k.Add("127.0.0.1:10001")
	k.Add("127.0.0.1:10002")
	k.Add("127.0.0.1:10003")

	fmt.Println(k.Get("123"))
	fmt.Println(k.Get("123"))
	fmt.Println(k.Get("234"))
	fmt.Println(k.Get("345"))
}
