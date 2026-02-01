// Copyright (c) 2024 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package tg

const (
	AuthStateUnknown      = 0
	AuthStateNew          = 1
	AuthStatePermBound    = 2
	AuthStateWaitInit     = 3
	AuthStateUnauthorized = 4
	AuthStateNeedPassword = 5
	AuthStateNormal       = 6
	AuthStateLogout       = 7
	AuthStateDeleted      = 8
)

const (
	AuthKeyTypeUnknown   = -1
	AuthKeyTypePerm      = 0
	AuthKeyTypeTemp      = 1
	AuthKeyTypeMediaTemp = 2
)

func NewAuthKeyInfo(keyId int64, key []byte, keyType int) *TLAuthKeyInfo {
	keyData := &TLAuthKeyInfo{
		AuthKeyId:          keyId,
		AuthKey:            key,
		AuthKeyType:        int32(keyType),
		PermAuthKeyId:      0,
		TempAuthKeyId:      0,
		MediaTempAuthKeyId: 0,
	}

	switch keyType {
	case AuthKeyTypePerm:
		keyData.PermAuthKeyId = keyId
	case AuthKeyTypeTemp:
		keyData.TempAuthKeyId = keyId
	case AuthKeyTypeMediaTemp:
		keyData.MediaTempAuthKeyId = keyId
	}

	return keyData
}
