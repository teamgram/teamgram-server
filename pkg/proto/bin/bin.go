// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
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

package bin

// WordLen represents 4-byte sequence.
// Values in TL are generally aligned to Word.
const WordLen = 4

const ClazzID_vector uint32 = 0x1cb5c415

func nearestPaddedValueLength(l int) int {
	n := WordLen * (l / WordLen)
	if n < l {
		n += WordLen
	}
	return n
}
