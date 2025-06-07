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
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/app/service/biz/username/username"
)

// AccountDeleteAccount
// account.deleteAccount#418d4e0b reason:string = Bool;
func (c *AccountCore) AccountDeleteAccount(in *mtproto.TLAccountDeleteAccount) (*mtproto.Bool, error) {
	me, err := c.svcCtx.Dao.UserClient.UserGetUserDataById(c.ctx, &user.TLUserGetUserDataById{
		UserId: c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("account.deleteAccount - error: %v", err)
		return nil, err
	}

	if me.Username != "" {
		_, err = c.svcCtx.Dao.UsernameClient.UsernameDeleteUsername(c.ctx, &username.TLUsernameDeleteUsername{
			Username: me.Username,
		})
		if err != nil {
			c.Logger.Errorf("account.deleteAccount - error: %v", err)
			return nil, err
		}
	}

	// TODO(@benqi): 1. Clear account data 2. Kickoff other client
	_, err = c.svcCtx.UserClient.UserDeleteUser(c.ctx, &user.TLUserDeleteUser{
		UserId: c.MD.UserId,
		Reason: in.Reason,
		Phone:  me.Phone,
	})
	if err != nil {
		c.Logger.Errorf("account.deleteAccount - error: %v", err)
		return nil, err
	}

	// s.AuthSessionRpcClient
	tKeyIdList, err := c.svcCtx.Dao.AuthsessionClient.AuthsessionResetAuthorization(c.ctx, &authsession.TLAuthsessionResetAuthorization{
		UserId:    c.MD.UserId,
		AuthKeyId: 0,
		Hash:      0,
	})
	if err != nil {
		c.Logger.Errorf("account.resetAuthorization#df77f3bc - error: %v", err)
		return nil, err
	}

	for _, id := range tKeyIdList.Datas {
		// notify kill session
		upds := mtproto.MakeTLUpdateAccountResetAuthorization(&mtproto.Updates{
			UserId:    c.MD.UserId,
			AuthKeyId: id,
		}).To_Updates()
		_, _ = c.svcCtx.Dao.SyncClient.SyncUpdatesMe(
			c.ctx,
			&sync.TLSyncUpdatesMe{
				UserId:        c.MD.UserId,
				PermAuthKeyId: id,
				ServerId:      nil,
				AuthKeyId:     nil,
				SessionId:     nil,
				Updates:       upds,
			})
	}

	_, _ = c.svcCtx.Dao.AuthsessionClient.AuthsessionUnbindAuthKeyUser(c.ctx, &authsession.TLAuthsessionUnbindAuthKeyUser{
		AuthKeyId: 0,
		UserId:    c.MD.UserId,
	})

	return mtproto.BoolTrue, nil
}
