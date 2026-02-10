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
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UsersGetMe
// users.getMe id:long token:string = User;
func (c *UsersCore) UsersGetMe(in *mtproto.TLUsersGetMe) (*mtproto.User, error) {
	// 通过 in.token 获取user
	user, err := c.svcCtx.Dao.UserClient.UserGetImmutableUserByToken(c.ctx, &user.TLUserGetImmutableUserByToken{
		Token: in.GetToken(),
	})
	if err != nil || user == nil {
		c.Logger.Errorf("users.getMe - error: %v", err)
		return nil, err
	} else if user.Id() != in.Id {
		err = mtproto.ErrTokenInvalid
		c.Logger.Errorf("users.getMe - error: %v", err)
		return nil, err
	}

	return user.ToSelfUser(), nil
}
