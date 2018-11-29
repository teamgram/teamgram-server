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

func TestAES256IGECryptor(t *testing.T) {
	key := []byte("11111111111111111111111111111111")
	iv := []byte("1111111111111111")

	cryptor := NewAES256IGECryptor(key, iv)

	MyString := "This is my string and I want to protect it with encryption"
	fmt.Printf("We start with a plain text: %s \n", MyString)

	MyStringByte := []byte(MyString)
	i := len(MyStringByte)
	if i/16 != 0 {
		MyStringByte = append(MyStringByte, make([]byte, 16-i%16)...)
	}

	Encrypted, err1 := cryptor.Encrypt(MyStringByte)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Printf("We encrypted the string this way: %s \n", hex.EncodeToString(Encrypted))

	Decrypted, err2 := cryptor.Decrypt(Encrypted)
	if err2 != nil {
		fmt.Println(err1)
	}

	fmt.Printf("Than we have the plain text again: %s \n", string(Decrypted))
}
