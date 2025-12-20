// Copyright 2025 Teamgram Authors
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
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserSearchUsername
// user.searchUsername q:string excluded_contacts:Vector<long> limit:int = Vector<UsernameData>;
func (c *UserCore) UserSearchUsername(in *user.TLUserSearchUsername) (*user.Vector_UsernameData, error) {
	var (
		rValList = &user.Vector_UsernameData{}
	)

	// Check query string and limit
	if len(in.Q) < 3 || in.Limit <= 0 {
		return rValList, nil
	}

	if in.Limit > 50 {
		in.Limit = 50
	}

	// 构造模糊查询字符串
	q2 := in.Q + "%"
	doList, _ := c.svcCtx.Dao.UsernameDAO.SearchByQueryNotIdListWithCB(
		c.ctx,
		q2,
		in.ExcludedContacts,
		in.Limit,
		func(sz, i int, v *dataobject.UsernameDO) {
			switch v.PeerType {
			case mtproto.PEER_USER:
				rValList.Datas = append(rValList.Datas, user.MakeTLUsernameData(&user.UsernameData{
					Username: v.Username,
					Peer:     mtproto.MakePeerUser(v.PeerId),
				}).To_UsernameData())
			case mtproto.PEER_CHANNEL:
				rValList.Datas = append(rValList.Datas, user.MakeTLUsernameData(&user.UsernameData{
					Username: v.Username,
					Peer:     mtproto.MakePeerChannel(v.PeerId),
				}).To_UsernameData())
			}
		})

	c.Logger.Infof("username.search - doList: %v", doList)
	return rValList, nil
}
