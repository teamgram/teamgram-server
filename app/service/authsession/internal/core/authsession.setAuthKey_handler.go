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

	"github.com/zeromicro/go-zero/core/mr"
)

var _ *tg.Bool

// AuthsessionSetAuthKey
// authsession.setAuthKey auth_key:AuthKeyInfo future_salt:FutureSalt expires_in:int = Bool;
func (c *AuthsessionCore) AuthsessionSetAuthKey(in *authsession.TLAuthsessionSetAuthKey) (*tg.Bool, error) {
	var (
		keyInfo, _ = in.AuthKey.(*tg.TLAuthKeyInfo)
		salt       *tg.TLFutureSalt
		err        error
	)

	if in.FutureSalt != nil {
		salt, _ = in.FutureSalt.(*tg.TLFutureSalt)
	}
	if salt == nil {
		err = c.svcCtx.Dao.SetAuthKeyV2(c.ctx, keyInfo, in.ExpiresIn)
	} else {
		err = mr.Finish(
			func() error {
				return c.svcCtx.Dao.SetAuthKeyV2(c.ctx, keyInfo, in.ExpiresIn)
			},
			func() error {
				return c.svcCtx.Dao.PutSaltCache(c.ctx, keyInfo.AuthKeyId, salt)
			})
	}

	if err != nil {
		c.Logger.Errorf("authsession.setAuthKey - error: %v", err)
		return tg.BoolFalse, nil
	}

	return tg.BoolTrue, nil
}
