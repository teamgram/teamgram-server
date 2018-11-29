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
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"math/big"
	"io/ioutil"
)

type RSACryptor struct {
	key *rsa.PrivateKey
}

func NewRSACryptor(keyFile string) *RSACryptor {
	pkcs1PemPrivateKey, err := ioutil.ReadFile(keyFile)
	if err != nil {
		panic("Invalid pemsKeyFile: " + keyFile)
	}
	block, _ := pem.Decode(pkcs1PemPrivateKey)
	if block == nil {
		panic("Invalid pemsKeyData: " + string(pkcs1PemPrivateKey))
		return nil
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic("Failed to parse private key: " + err.Error())
	}

	return &RSACryptor{
		key: key,
	}
}

func (m *RSACryptor) Encrypt(b []byte) []byte {
	c := new(big.Int)
	c.Exp(new(big.Int).SetBytes(b), big.NewInt(int64(m.key.E)), m.key.N)

	e := c.Bytes()
	var size = len(e)
	if size < 256 {
		size = 256
	}

	res := make([]byte, size)
	copy(res, c.Bytes())

	return res
}

func (m *RSACryptor) Decrypt(b []byte) []byte {
	c := new(big.Int)
	c.Exp(new(big.Int).SetBytes(b), m.key.D, m.key.N)
	return c.Bytes()
}
