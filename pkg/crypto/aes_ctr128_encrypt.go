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
	"crypto/cipher"
	"strconv"
)

type AesCTR128KeySizeError int

func (k AesCTR128KeySizeError) Error() string {
	return "AesCTR128KeySizeError: invalid key size " + strconv.Itoa(int(k))
}

type AesCTR128Encrypt struct {
	// block cipher.Block
	stream cipher.Stream
}

// key长度必须为16、24或32
func NewAesCTR128Encrypt(key []byte, iv []byte) (*AesCTR128Encrypt, error) {
	block2, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// iv长度为16
	if len(iv) != 16 {
		return nil, AesCTR128KeySizeError(len(iv))
	}

	stream2 := cipher.NewCTR(block2, iv)

	return &AesCTR128Encrypt{
		// block:	block2,
		stream: stream2,
	}, nil
}

func (this *AesCTR128Encrypt) Encrypt(plaintext []byte) []byte {
	this.stream.XORKeyStream(plaintext, plaintext)
	return plaintext
}
