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
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserGetUserDataListByIdList
// user.getUserDataListByIdList user_id_list:Vector<long> = Vector<UserData>;
func (c *UserCore) UserGetUserDataListByIdList(in *user.TLUserGetUserDataListByIdList) (*user.Vector_UserData, error) {
	users := &user.Vector_UserData{
		Datas: []*user.UserData{},
	}

	for _, id := range in.UserIdList {
		cacheData := c.svcCtx.Dao.GetCacheUserData(c.ctx, id)
		if cacheData == nil {
			c.Logger.Errorf("user.getUserDataById - error: not found userId(%d)", id)
			continue
		}
		users.Datas = append(users.Datas, cacheData.GetUserData())
	}

	return users, nil

}
