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
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dao"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserSaveMusic
// user.saveMusic flags:# unsave:flags.0?true user_id:long id:long after_id:flags.15?long = Bool;
func (c *UserCore) UserSaveMusic(in *user.TLUserSaveMusic) (*mtproto.Bool, error) {
	if in.GetUnsave() {
		unsaveIdx := -1
		nextId := int64(-1)

		doList, _ := c.svcCtx.Dao.UserSavedMusicDAO.SelectListWithCB(
			c.ctx,
			in.GetUserId(),
			func(sz int, i int, v *dataobject.UserSavedMusicDO) {
				if v.SavedMusicId == in.GetUserId() {
					_, _ = c.svcCtx.Dao.UserSavedMusicDAO.Delete(c.ctx, in.GetUserId(), in.GetId())
					unsaveIdx = i
				}
			},
		)

		// if unsaveIdx >= 0 {
		if unsaveIdx == 0 {
			if len(doList) > 1 {
				nextId = doList[1].SavedMusicId
			} else {
				nextId = 0
			}
		} else if unsaveIdx > 0 {
			nextId = doList[0].SavedMusicId
		}

		if nextId >= 0 {
			_, _, _ = c.svcCtx.Dao.CachedConn.Exec(
				c.ctx,
				func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
					_, err := c.svcCtx.Dao.UsersDAO.UpdateSavedMusicId(c.ctx, nextId, in.GetUserId())
					return 0, 0, err
				},
				dao.GenCacheUserDataCacheKey(in.GetUserId()))
		}
	} else {
		_, _, err := c.svcCtx.Dao.UserSavedMusicDAO.InsertOrUpdate(c.ctx, &dataobject.UserSavedMusicDO{
			UserId:       in.GetUserId(),
			SavedMusicId: in.GetId(),
			Order2:       time.Now().Unix() << 32,
			Deleted:      false,
		})
		if err != nil {
			c.Logger.Errorf("user.saveMusic - error: %v", err)
			return mtproto.BoolFalse, nil
		}

		_, _, _ = c.svcCtx.Dao.CachedConn.Exec(
			c.ctx,
			func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
				_, err = c.svcCtx.Dao.UsersDAO.UpdateSavedMusicId(c.ctx, in.GetId(), in.GetUserId())
				return 0, 0, err
			},
			dao.GenCacheUserDataCacheKey(in.GetUserId()))
	}

	return mtproto.BoolTrue, nil
}
