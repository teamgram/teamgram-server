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

// UserSearch
// user.search q:string offset_id:int limit:int hash:long = UsersFound;
func (c *UserCore) UserSearch(in *user.TLUserSearch) (*user.UsersFound, error) {
	var (
		founds *user.UsersFound
		isData = len(in.ExcludedContacts) == 0
	)

	if isData {
		founds = user.MakeTLUsersDataFound(&user.UsersFound{
			Count:      0,
			Users:      []*mtproto.UserData{},
			NextOffset: "",
		}).To_UsersFound()
	} else {
		founds = user.MakeTLUsersIdFound(&user.UsersFound{
			IdList: []int64{},
		}).To_UsersFound()
	}

	// Check query string and limit
	if len(in.Q) < 3 || in.Limit <= 0 {
		return founds, nil
	}

	if in.Limit > 50 {
		in.Limit = 50
	}

	// 构造模糊查询字符串
	q := in.Q + "%"
	q2 := "%" + in.Q + "%"

	rList, _ := c.svcCtx.Dao.UsersDAO.SearchByQueryString(
		c.ctx,
		q,
		q2,
		[]int64{0},
		in.Limit)

	if isData {
		userDataList, _ := c.UserGetUserDataListByIdList(&user.TLUserGetUserDataListByIdList{
			UserIdList: rList,
		})

		if len(userDataList.GetDatas()) > 0 {
			founds.Users = userDataList.Datas
			founds.Count = int32(len(rList))
		}
	} else {
		founds.IdList = rList
	}

	return founds, nil
}
