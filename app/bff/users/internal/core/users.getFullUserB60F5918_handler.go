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
)

// UsersGetFullUserB60F5918
// users.getFullUser#b60f5918 id:InputUser = users.UserFull;
func (c *UsersCore) UsersGetFullUserB60F5918(in *mtproto.TLUsersGetFullUserB60F5918) (*mtproto.Users_UserFull, error) {
	userFull, err := c.UsersGetFullUserCA30A5B1(&mtproto.TLUsersGetFullUserCA30A5B1{
		Id: in.Id,
	})

	if err != nil {
		c.Logger.Errorf("users.getFullUser#b60f5918 - error: %v", err)
		return nil, err
	}

	return mtproto.MakeTLUsersUserFull(&mtproto.Users_UserFull{
		FullUser: userFull,
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{userFull.GetUser()},
	}).To_Users_UserFull(), nil
}
