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

// CombineInt64Hash2
// ?????
func CombineInt64Hash2(acc, id int64) int64 {
	acc ^= acc >> 21
	acc ^= acc << 35
	acc ^= acc >> 4
	return acc + id
}

// CombineInt64Hash
// ?????
func CombineInt64Hash(acc, id int64) int64 {
	acc ^= id >> 21
	acc ^= id << 35
	acc ^= id >> 4
	return acc + id
}

func HashInt64(acc int64) int32 {
	return int32(acc&0xffffffff ^ acc>>32)
}
