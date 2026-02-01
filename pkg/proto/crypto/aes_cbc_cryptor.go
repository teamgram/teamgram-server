// Copyright 2024 Teamgram Authors
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

package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

type aesCbcCryptor struct {
	aesKey []byte
	aesIV  []byte
}

func NewAesCBCEncrypt(aesKey, aesIV []byte) (cipher.BlockMode, error) {
	if len(aesIV) != aes.BlockSize {
		return nil, fmt.Errorf("invalid iv")
	}

	c, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	return cipher.NewCBCEncrypter(c, aesIV), nil
}

func NewAesCBCDecrypt(aesKey, aesIV []byte) (cipher.BlockMode, error) {
	if len(aesIV) != aes.BlockSize {
		return nil, fmt.Errorf("invalid iv")
	}

	c, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	return cipher.NewCBCDecrypter(c, aesIV), nil
}
