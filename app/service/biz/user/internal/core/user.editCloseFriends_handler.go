// Copyright 2024 Teamgram Authors
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
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"time"
)

// UserEditCloseFriends
// user.editCloseFriends user_id:long id:Vector<long> = Bool;
func (c *UserCore) UserEditCloseFriends(in *user.TLUserEditCloseFriends) (*mtproto.Bool, error) {
	tR := sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		_, err := c.svcCtx.Dao.CloseFriendsDAO.DeleteTx(tx, in.UserId)
		if err != nil {
			c.Logger.Errorf("user.editCloseFriends - error: %v", err)
			result.Err = err
			return
		}

		doList := make([]*dataobject.CloseFriendsDO, 0, len(in.Id))
		date := time.Now().Unix()
		for _, id := range in.Id {
			doList = append(doList, &dataobject.CloseFriendsDO{
				UserId:        in.UserId,
				CloseFriendId: id,
				Date:          date,
			})
		}
		if len(doList) == 0 {
			return
		}
		_, _, err = c.svcCtx.Dao.CloseFriendsDAO.InsertBulkTx(tx, doList)
		if err != nil {
			c.Logger.Errorf("user.editCloseFriends - error: %v", err)
			result.Err = err
			return
		}
	})
	if tR.Err != nil {
		return nil, tR.Err
	}

	return mtproto.BoolTrue, nil
}
