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

package gnet

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/crypto"
)

type AuthKeyUtil = authKeyUtil

var NewAuthKeyUtil = newAuthKeyUtil

type authKeyUtil struct {
	keyData *mtproto.AuthKeyInfo
	key     *crypto.AuthKey
}

func newAuthKeyUtil(k *mtproto.AuthKeyInfo) *authKeyUtil {
	return &authKeyUtil{
		keyData: k,
		key:     crypto.NewAuthKey(k.AuthKeyId, k.AuthKey),
	}
}

func (k *authKeyUtil) Equal(o *authKeyUtil) bool {
	return k.keyData.AuthKeyId == o.keyData.AuthKeyId
}

func (k *authKeyUtil) AuthKeyId() int64 {
	return k.keyData.AuthKeyId
}

func (k *authKeyUtil) AuthKeyType() int {
	return int(k.keyData.AuthKeyType)
}

func (k *authKeyUtil) PermAuthKeyId() int64 {
	return k.keyData.PermAuthKeyId
}

func (k *authKeyUtil) TempAuthKeyId() int64 {
	return k.keyData.TempAuthKeyId
}

func (k *authKeyUtil) MediaTempAuthKeyId() int64 {
	return k.keyData.MediaTempAuthKeyId
}

func (k *authKeyUtil) AesIgeEncrypt(rawData []byte) ([]byte, []byte, error) {
	return k.key.AesIgeEncrypt(rawData)
}

func (k *authKeyUtil) AesIgeDecrypt(msgKey, rawData []byte) ([]byte, error) {
	return k.key.AesIgeDecrypt(msgKey, rawData)
}
