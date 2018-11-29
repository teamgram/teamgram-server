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

package zrpc

import (
	"fmt"
	"hash/crc32"
	"testing"
)

func TestCRC32(t *testing.T) {
	b := []byte("01234567899876543210")
	b1 := []byte("0123456789")
	b2 := []byte("9876543210")

	fmt.Println(crc32.ChecksumIEEE(b))
	crc32Hash := crc32.NewIEEE()
	crc32Hash.Write(b1)
	crc32Hash.Write(b2)
	fmt.Println(crc32Hash.Sum32())
}
