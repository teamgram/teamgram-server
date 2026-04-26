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
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AuthsessionGetUserId
// authsession.getUserId auth_key_id:long = Int64;
func (c *AuthsessionCore) AuthsessionGetUserId(in *authsession.TLAuthsessionGetUserId) (*tg.Int64, error) {
	keyData, err := c.svcCtx.Repo.ResolvePermAuthKey(c.ctx, in.AuthKeyId)
	if err != nil {
		return nil, err
	}
	userId, err := c.svcCtx.Repo.GetAuthKeyUserId(c.ctx, keyData.PermAuthKeyId)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLInt64(&tg.TLInt64{V: userId}).ToInt64(), nil
}
