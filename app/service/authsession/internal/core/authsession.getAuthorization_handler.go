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

package core

import (
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
)

var _ *tg.Bool

// AuthsessionGetAuthorization
// authsession.getAuthorization auth_key_id:long = Authorization;
func (c *AuthsessionCore) AuthsessionGetAuthorization(in *authsession.TLAuthsessionGetAuthorization) (*tg.Authorization, error) {
	var (
		inKeyId = in.AuthKeyId
	)

	keyData, err := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, inKeyId)
	if err != nil {
		c.Logger.Errorf("queryAuthKeyV2(%d) is error: %v", inKeyId, err)
		return nil, err
	} else if keyData.PermAuthKeyId == 0 {
		c.Logger.Errorf("queryAuthKeyV2(%d) - PermAuthKeyId is empty", inKeyId)
		return nil, tg.ErrAuthKeyPermEmpty
	}

	authorization, err := c.svcCtx.Dao.GetAuthorization(c.ctx, keyData.PermAuthKeyId)
	if err != nil {
		c.Logger.Errorf("session.getAuthorization - error: %v", err)
		return nil, err
	}

	return authorization.ToAuthorization(), nil
}
