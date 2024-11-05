// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
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

package codec

import (
	"github.com/teamgram/proto/mtproto/crypto"
)

// AesCTR128Crypto AesCTR128Crypto
type AesCTR128Crypto struct {
	decrypt *crypto.AesCTR128Encrypt
	encrypt *crypto.AesCTR128Encrypt
}

func newAesCTR128Crypto(d, e *crypto.AesCTR128Encrypt) *AesCTR128Crypto {
	return &AesCTR128Crypto{
		decrypt: d,
		encrypt: e,
	}
}
func (e *AesCTR128Crypto) Encrypt(plaintext []byte) []byte {
	if e == nil {
		return plaintext
	} else {
		return e.encrypt.Encrypt(plaintext)
	}
}

func (e *AesCTR128Crypto) Decrypt(plaintext []byte) []byte {
	if e == nil {
		return plaintext
	} else {
		return e.decrypt.Encrypt(plaintext)
	}
}
