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

// UsersGetUsers
// users.getUsers#d91a548 id:Vector<InputUser> = Vector<User>;
func (c *UsersCore) UsersGetUsers(in *mtproto.TLUsersGetUsers) (*mtproto.Vector_User, error) {
	var (
		idList []int64
	)

	for _, inputUser := range in.Id {
		peer := mtproto.FromInputUser(c.MD.UserId, inputUser)
		switch peer.PeerType {
		case mtproto.PEER_SELF, mtproto.PEER_USER:
			idList = append(idList, peer.PeerId)
		default:
			c.Logger.Errorf("invalid userId")
		}
	}

	mUsers, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: idList,
		To: []int64{c.MD.UserId},
	})

	return &mtproto.Vector_User{
		Datas: mUsers.GetUserListByIdList(c.MD.UserId, idList...),
	}, nil
}
