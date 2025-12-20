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

// UserGetListByUsernameList
// user.getListByUsernameList names:Vector<string> = Vector<UsernameData>;
func (c *UserCore) UserGetListByUsernameList(in *user.TLUserGetListByUsernameList) (*user.Vector_UsernameData, error) {
	var (
		rValues = &user.Vector_UsernameData{
			Datas: []*user.UsernameData{},
		}
	)

	if _, err := c.svcCtx.Dao.UsernameDAO.SelectListWithCB(c.ctx, in.Names, func(sz, i int, v *dataobject.UsernameDO) {
		var (
			peer *mtproto.Peer
		)

		switch v.PeerType {
		case mtproto.PEER_USER:
			peer = mtproto.MakePeerUser(v.PeerId)
		case mtproto.PEER_CHANNEL:
			peer = mtproto.MakePeerChannel(v.PeerId)
		default:
			return
		}

		rValues.Datas = append(rValues.Datas, user.MakeTLUsernameData(&user.UsernameData{
			Username: v.Username,
			Peer:     peer,
			Editable: v.Editable,
			Active:   v.Active,
		}).To_UsernameData())
	}); err != nil {
		c.Logger.Errorf("username.getListByUsernameList - error: %v", err)
	}

	return rValues, nil
}
