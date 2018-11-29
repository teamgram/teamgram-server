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
	"crypto/aes"
	"errors"
)

type AES256IGECryptor struct {
	aesKey []byte
	aesIV  []byte
}

func NewAES256IGECryptor(aesKey, aesIV []byte) *AES256IGECryptor {
	// guard conditions
	if (len(aesIV)) < aes.BlockSize {
		return nil
	}

	k := len(aesKey)
	switch k {
	default:
		return nil
	case 16, 24, 32:
		break
	}
	return &AES256IGECryptor{
		aesKey: aesKey,
		aesIV:  aesIV,
	}
}

// data长度必须是aes.BlockSize(16)的倍数，如果不是请调用者补齐
func (c *AES256IGECryptor) Encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.aesKey)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, errors.New("AES256IGE: data too small to encrypt")
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("AES256IGE: data not divisible by block size")
	}

	t := make([]byte, aes.BlockSize)
	x := make([]byte, aes.BlockSize)
	y := make([]byte, aes.BlockSize)
	copy(x, c.aesIV[:aes.BlockSize])
	copy(y, c.aesIV[aes.BlockSize:])
	encrypted := make([]byte, len(data))

	i := 0
	for i < len(data) {
		xor(x, data[i:i+aes.BlockSize])
		block.Encrypt(t, x)
		xor(t, y)
		x, y = t, data[i:i+aes.BlockSize]
		copy(encrypted[i:], t)
		i += aes.BlockSize
	}

	return encrypted, nil
}

func (c *AES256IGECryptor) Decrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.aesKey)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, errors.New("AES256IGE: data too small to decrypt")
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("AES256IGE: data not divisible by block size")
	}

	t := make([]byte, aes.BlockSize)
	x := make([]byte, aes.BlockSize)
	y := make([]byte, aes.BlockSize)
	copy(x, c.aesIV[:aes.BlockSize])
	copy(y, c.aesIV[aes.BlockSize:])
	decrypted := make([]byte, len(data))

	i := 0
	for i < len(data) {
		xor(y, data[i:i+aes.BlockSize])
		block.Decrypt(t, y)
		xor(t, x)
		y, x = t, data[i:i+aes.BlockSize]
		copy(decrypted[i:], t)
		i += aes.BlockSize
	}

	return decrypted, nil
}

func xor(dst, src []byte) {
	for i := range dst {
		dst[i] = dst[i] ^ src[i]
	}
}
