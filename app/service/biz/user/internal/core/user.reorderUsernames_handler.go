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
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserReorderUsernames
// user.reorderUsernames peer_type:int peer_id:long username_list:Vector<string> = Bool;
func (c *UserCore) UserReorderUsernames(in *user.TLUserReorderUsernames) (*mtproto.Bool, error) {
	var (
		order2 = time.Now().Unix() << 32
	)

	_ = sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		for i, username := range in.GetUsernameList() {
			_, err := c.svcCtx.Dao.UsernameDAO.UpdateTx(tx,
				map[string]interface{}{
					"order2": order2 + int64(i),
				},
				username)
			if err != nil {
				result.Err = err
				return
			}
		}
	})

	return mtproto.BoolTrue, nil
}
