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
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// AuthCheckPassword
// auth.checkPassword#d18b4d16 password:InputCheckPasswordSRP = auth.Authorization;
func (c *AuthorizationCore) AuthCheckPassword(in *mtproto.TLAuthCheckPassword) (*mtproto.Auth_Authorization, error) {
	// TODO: check password
	c.Logger.Errorf("auth.checkPassword blocked, License key from https://teamgram.net required to unlock enterprise features.")

	user, err := c.svcCtx.UserClient.UserGetImmutableUser(c.ctx, &userpb.TLUserGetImmutableUser{
		Id: c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("auth.checkPassword - error: %v", err)
		return nil, err
	}

	return mtproto.MakeTLAuthAuthorization(&mtproto.Auth_Authorization{
		SetupPasswordRequired: false,
		OtherwiseReloginDays:  nil,
		TmpSessions:           nil,
		FutureAuthToken:       nil,
		User:                  user.ToSelfUser(),
	}).To_Auth_Authorization(), nil
}
