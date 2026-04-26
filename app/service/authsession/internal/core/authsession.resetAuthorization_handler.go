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
)

// AuthsessionResetAuthorization
// authsession.resetAuthorization user_id:long auth_key_id:long hash:long = Vector<long>;
func (c *AuthsessionCore) AuthsessionResetAuthorization(in *authsession.TLAuthsessionResetAuthorization) (*authsession.VectorLong, error) {
	excludePermAuthKeyId := int64(0)
	if in.AuthKeyId != 0 {
		keyData, err := c.svcCtx.Repo.ResolvePermAuthKey(c.ctx, in.AuthKeyId)
		if err != nil {
			return nil, err
		}
		excludePermAuthKeyId = keyData.PermAuthKeyId
	}

	keyIds, err := c.svcCtx.Repo.ResetAuthorization(c.ctx, in.UserId, excludePermAuthKeyId, in.Hash)
	if err != nil {
		return nil, err
	}

	expandedKeyIds := make([]int64, 0, len(keyIds))
	for _, keyId := range keyIds {
		keyData, err := c.svcCtx.Repo.QueryAuthKey(c.ctx, keyId)
		if err != nil {
			return nil, err
		}
		if keyData.TempAuthKeyId != 0 {
			expandedKeyIds = append(expandedKeyIds, keyData.TempAuthKeyId)
		} else {
			expandedKeyIds = append(expandedKeyIds, keyId)
		}
	}

	return &authsession.VectorLong{Datas: expandedKeyIds}, nil
}
