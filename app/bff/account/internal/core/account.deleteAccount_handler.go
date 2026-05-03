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
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AccountDeleteAccount
// account.deleteAccount#418d4e0b reason:string = Bool;
func (c *AccountCore) AccountDeleteAccount(in *tg.TLAccountDeleteAccount) (*tg.Bool, error) {
	selfID, err := requireSelfID(c)
	if err != nil {
		return nil, err
	}
	if err := requireUserClient(c); err != nil {
		return nil, err
	}
	if err := requireAuthsessionClient(c); err != nil {
		return nil, err
	}

	reason := ""
	if in != nil {
		reason = in.Reason
	}

	me, err := c.svcCtx.Repo.UserClient.UserGetUserDataById(c.ctx, &userpb.TLUserGetUserDataById{
		UserId: selfID,
	})
	if err != nil {
		return nil, err
	}
	if me == nil {
		return nil, tg.ErrUserIdInvalid
	}

	if me.Username != "" {
		if _, err = c.svcCtx.Repo.UserClient.UserDeleteUsername(c.ctx, &userpb.TLUserDeleteUsername{
			Username: me.Username,
		}); err != nil {
			return nil, err
		}
	}

	if _, err = c.svcCtx.Repo.UserClient.UserDeleteUser(c.ctx, &userpb.TLUserDeleteUser{
		UserId: selfID,
		Reason: reason,
		Phone:  me.Phone,
	}); err != nil {
		return nil, err
	}

	if _, err = c.svcCtx.Repo.AuthsessionClient.AuthsessionResetAuthorization(c.ctx, &authsession.TLAuthsessionResetAuthorization{
		UserId:    selfID,
		AuthKeyId: 0,
		Hash:      0,
	}); err != nil {
		return nil, err
	}

	// TODO(v2 account): master notified killed sessions through sync; do not migrate sync calls until v2 userupdates/gateway delivery is defined.
	if _, err = c.svcCtx.Repo.AuthsessionClient.AuthsessionUnbindAuthKeyUser(c.ctx, &authsession.TLAuthsessionUnbindAuthKeyUser{
		AuthKeyId: 0,
		UserId:    selfID,
	}); err != nil {
		c.Logger.Errorf("account.deleteAccount - unbind auth key user failed: user_id: %d, err: %v", selfID, err)
	}

	return tg.BoolTrue, nil
}
