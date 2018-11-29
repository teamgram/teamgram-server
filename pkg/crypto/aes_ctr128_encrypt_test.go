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

package crypto

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAesCTR128Encrypt(t *testing.T) {
	key := []byte("12345678901234561234567890123456")
	iv := []byte("1234567890123456")

	MyString := "This is my string and I want to protect it with encryption"

	fmt.Printf("We start with a plain text: %s \n", MyString)
	MyStringByte := []byte(MyString)
	encryptor, _ := NewAesCTR128Encrypt(key, iv)
	Encrypted := encryptor.Encrypt(MyStringByte)
	fmt.Printf("We encrypted the string this way: %s \n", hex.EncodeToString(Encrypted))

	decryptor, _ := NewAesCTR128Encrypt(key, iv)
	Decrypted := decryptor.Encrypt(MyStringByte)
	fmt.Printf("Than we have the plain text again: %s \n", string(Decrypted))
}
