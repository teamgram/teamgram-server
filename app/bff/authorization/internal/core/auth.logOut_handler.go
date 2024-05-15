// Copyright 2022 Teamgram Authors
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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/crypto"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"
)

// AuthLogOut
// auth.logOut#3e72ba19 = auth.LoggedOut;
func (c *AuthorizationCore) AuthLogOut(in *mtproto.TLAuthLogOut) (*mtproto.Auth_LoggedOut, error) {
	// unbind auth_key and user_id
	_, err := c.svcCtx.Dao.AuthsessionClient.AuthsessionUnbindAuthKeyUser(c.ctx, &authsession.TLAuthsessionUnbindAuthKeyUser{
		AuthKeyId: c.MD.PermAuthKeyId,
		UserId:    c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("auth.logOut - error: %v", err)
		return nil, err
	} else {
		if c.svcCtx.Plugin != nil {
			c.svcCtx.Plugin.OnAuthLogout(c.ctx, c.MD.UserId, c.MD.PermAuthKeyId)
		}
	}

	futureAuthToken := crypto.RandomBytes(64)
	c.svcCtx.Dao.PutFutureAuthToken(c.ctx, futureAuthToken, c.MD.UserId)

	return mtproto.MakeTLAuthLoggedOut(&mtproto.Auth_LoggedOut{
		FutureAuthToken: futureAuthToken,
	}).To_Auth_LoggedOut(), nil
}
