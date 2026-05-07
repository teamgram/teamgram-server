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
	"errors"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UsersGetMe
// users.getMe id:long token:string = User;
func (c *UsersCore) UsersGetMe(in *tg.TLUsersGetMe) (*tg.User, error) {
	if in == nil || in.Token == "" || in.Id <= 0 {
		return nil, tg.ErrInputRequestInvalid
	}
	if c.svcCtx == nil || c.svcCtx.Repo == nil || c.svcCtx.Repo.UserClient == nil {
		return nil, fmt.Errorf("users.getMe: user client is nil")
	}

	immutableUser, err := c.svcCtx.Repo.UserClient.UserGetImmutableUserByToken(c.ctx, &userpb.TLUserGetImmutableUserByToken{
		Token: in.Token,
	})
	if err != nil {
		if errors.Is(err, userpb.ErrBotNotFound) {
			return nil, tg.ErrTokenInvalid
		}
		return nil, err
	}
	if immutableUser == nil || immutableUser.User == nil {
		return nil, tg.ErrTokenInvalid
	}
	if immutableUser.User.Id != in.Id {
		return nil, tg.ErrTokenInvalid
	}

	users, err := userprojection.ProjectUsers(c.ctx, c.svcCtx.Repo.UserClient, in.Id, []int64{in.Id}, userprojection.MissingExplicitInput)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, tg.ErrUserIdInvalid
	}

	return &tg.User{Clazz: users[0]}, nil
}
