// Copyright (c) 2024-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

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
