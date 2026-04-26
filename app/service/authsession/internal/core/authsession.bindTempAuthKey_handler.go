// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
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

package core

import (
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/mt"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AuthsessionBindTempAuthKey
// authsession.bindTempAuthKey perm_auth_key_id:long nonce:long expires_at:int encrypted_message:bytes = Bool;
func (c *AuthsessionCore) AuthsessionBindTempAuthKey(in *authsession.TLAuthsessionBindTempAuthKey) (*tg.Bool, error) {
	permKeyData, err := c.svcCtx.Repo.QueryAuthKey(c.ctx, in.PermAuthKeyId)
	if err != nil {
		return nil, authsession.ErrEncryptedMessageInvalid
	}
	if len(in.EncryptedMessage) < 24 {
		return nil, authsession.ErrEncryptedMessageInvalid
	}

	permAuthKey := crypto.NewAuthKey(in.PermAuthKeyId, permKeyData.AuthKey)
	innerData, err := permAuthKey.AesIgeDecryptV1(in.EncryptedMessage[8:24], in.EncryptedMessage[24:])
	if err != nil || len(innerData) < 32 {
		return nil, authsession.ErrEncryptedMessageInvalid
	}

	inner, err := mt.DecodeBindAuthKeyInnerClazz(bin.NewDecoder(innerData[32:]))
	if err != nil || inner == nil {
		return nil, authsession.ErrEncryptedMessageInvalid
	}
	if inner.PermAuthKeyId != in.PermAuthKeyId || inner.Nonce != in.Nonce || inner.ExpiresAt != in.ExpiresAt {
		return nil, authsession.ErrEncryptedMessageInvalid
	}

	tempKeyData, err := c.svcCtx.Repo.QueryAuthKey(c.ctx, inner.TempAuthKeyId)
	if err != nil {
		return nil, authsession.ErrEncryptedMessageInvalid
	}
	if err := c.svcCtx.Repo.BindKeyId(c.ctx, inner.PermAuthKeyId, tempKeyData.AuthKeyType, inner.TempAuthKeyId); err != nil {
		return nil, err
	}
	if err := c.svcCtx.Repo.BindKeyId(c.ctx, inner.TempAuthKeyId, tg.AuthKeyTypePerm, inner.PermAuthKeyId); err != nil {
		return nil, err
	}

	return tg.BoolTrue, nil
}
