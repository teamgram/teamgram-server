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
	"context"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dao"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserSetMainProfileTab
// user.setMainProfileTab user_id:long tab:ProfileTab = Bool;
func (c *UserCore) UserSetMainProfileTab(in *user.TLUserSetMainProfileTab) (*mtproto.Bool, error) {
	_, _, err := c.svcCtx.Dao.CachedConn.Exec(
		c.ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			_, err := c.svcCtx.Dao.UsersDAO.UpdateMainTab(c.ctx, mtproto.FromProfileTabToType(in.GetTab()), in.GetUserId())
			if err != nil {
				c.Logger.Errorf("user.setMainProfileTab - error: %v", err)
			}
			return 0, 0, err
		},
		dao.GenCacheUserDataCacheKey(in.UserId))
	if err != nil {
		c.Logger.Errorf("user.setMainProfileTab - error: %v", err)
	}

	return mtproto.BoolTrue, nil
}
