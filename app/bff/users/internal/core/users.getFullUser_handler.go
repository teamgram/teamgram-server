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
	"fmt"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UsersGetFullUser
// users.getFullUser#b60f5918 id:InputUser = users.UserFull;
func (c *UsersCore) UsersGetFullUser(in *tg.TLUsersGetFullUser) (*tg.UsersUserFull, error) {
	selfID, err := requireSelfID(c)
	if err != nil {
		return nil, err
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	if c.svcCtx == nil || c.svcCtx.Repo == nil || c.svcCtx.Repo.UserClient == nil {
		return nil, fmt.Errorf("users.getFullUser: user client is nil")
	}

	peerID, err := userIDFromInputUser(selfID, in.Id)
	if err != nil {
		return nil, err
	}

	return c.svcCtx.Repo.UserClient.UserGetFullUser(c.ctx, &userpb.TLUserGetFullUser{
		SelfUserId: selfID,
		Id:         peerID,
	})
}
