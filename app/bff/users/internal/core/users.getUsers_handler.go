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

	userprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UsersGetUsers
// users.getUsers#d91a548 id:Vector<InputUser> = Vector<User>;
func (c *UsersCore) UsersGetUsers(in *tg.TLUsersGetUsers) (*tg.VectorUser, error) {
	selfID, err := requireSelfID(c)
	if err != nil {
		return nil, err
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	if c.svcCtx == nil || c.svcCtx.Repo == nil || c.svcCtx.Repo.UserClient == nil {
		return nil, fmt.Errorf("users.getUsers: user client is nil")
	}

	requestedIDs := make([]int64, 0, len(in.Id))
	for _, inputUser := range in.Id {
		id, err := userIDFromInputUser(selfID, inputUser)
		if err != nil {
			return nil, err
		}
		requestedIDs = append(requestedIDs, id)
	}

	users, err := userprojection.ProjectUsers(c.ctx, c.svcCtx.Repo.UserClient, selfID, requestedIDs, userprojection.MissingExplicitInput)
	if err != nil {
		return nil, err
	}

	return &tg.VectorUser{Datas: users}, nil
}
