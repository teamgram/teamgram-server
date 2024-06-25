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
	"context"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/jsonx"
)

// DialogCreateDialogFilter
// dialog.createDialogFilter user_id:long dialog_filter:DialogFilterExt = DialogFilterExt;
func (c *DialogCore) DialogCreateDialogFilter(in *dialog.TLDialogCreateDialogFilter) (*dialog.DialogFilterExt, error) {
	var (
		dialogFilterExtList []*dialog.DialogFilterExt
		dialogExt           = in.GetDialogFilter()
		created             = false
	)

	c.svcCtx.Dao.CachedConn.QueryRow(
		c.ctx,
		&dialogFilterExtList,
		dialog.GenDialogFilterCacheKey(in.UserId),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			var (
				vList []*dialog.DialogFilterExt
			)
			_, err := c.svcCtx.Dao.DialogFiltersDAO.SelectListWithCB(
				c.ctx,
				in.UserId,
				func(sz, i int, v *dataobject.DialogFiltersDO) {
					dialogFilter := &dialog.DialogFilterExt{
						Id:           v.DialogFilterId,
						DialogFilter: nil,
						Order:        v.OrderValue,
					}

					if err := jsonx.UnmarshalFromString(v.DialogFilter, &dialogFilter.DialogFilter); err != nil {
						c.Logger.Errorf("jsonx.UnmarshalFromString(%v) - error: %v", v, err)
						// continue
						return
					}

					if dialogFilter.DialogFilter == nil {
						dialogFilter.DialogFilter = mtproto.MakeTLDialogFilter(nil).To_DialogFilter()
					}

					vList = append(vList, dialogFilter)
				})
			if err != nil {
				return err
			}

			*v.(*[]*dialog.DialogFilterExt) = vList
			return err
		},
	)

	for _, v := range dialogFilterExtList {
		if v.Slug == dialogExt.Slug {
			created = true
			dialogExt.Id = v.Id
		}
		if created {
			break
		}
		if v.Id > dialogExt.Id {
			dialogExt.Id = v.Id
		}
	}

	if created {
		return dialogExt, nil
	}

	dialogExt.Id++
	c.svcCtx.Dao.CachedConn.Exec(
		c.ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			dialogExt.DialogFilter.Id = dialogExt.Id
			dData, _ := jsonx.MarshalToString(dialogExt.DialogFilter)
			return c.svcCtx.Dao.DialogFiltersDAO.InsertOrUpdate(c.ctx, &dataobject.DialogFiltersDO{
				UserId:         in.UserId,
				DialogFilterId: dialogExt.Id,
				IsChatlist:     true,
				JoinedBySlug:   true,
				Slug:           dialogExt.Slug,
				HasMyInvites:   0,
				DialogFilter:   dData,
				OrderValue:     time.Now().Unix() << 32,
				FromSuggested:  -1,
				Deleted:        false,
			})
		},
		dialog.GenDialogFilterCacheKey(in.UserId))

	return dialogExt, nil
}
