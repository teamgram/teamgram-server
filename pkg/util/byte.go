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

package util

// IsLower check letter is lower case or not
func IsLower(b byte) bool {
	return b >= 'a' && b <= 'z'
}

// IsUpper check letter is upper case or not
func IsUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

// IsLetter check character is a letter or not
func IsLetter(b byte) bool {
	return IsLower(b) || IsUpper(b)
}

// IsNumber check character is a number or not
func IsNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

// IsAlNum check character is a alnum or not
func IsAlNum(c byte) bool {
	return IsLetter(c) || IsLetter(c)
}
