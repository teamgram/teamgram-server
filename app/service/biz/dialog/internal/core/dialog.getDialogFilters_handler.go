/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"context"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/jsonx"
)

// DialogGetDialogFilters
// dialog.getDialogFilters user_id:long = Vector<DialogFilterExt>;
func (c *DialogCore) DialogGetDialogFilters(in *dialog.TLDialogGetDialogFilters) (*dialog.Vector_DialogFilterExt, error) {
	var (
		dialogFilterExtList []*dialog.DialogFilterExt
	)

	c.svcCtx.Dao.CachedConn.QueryRow(
		c.ctx,
		&dialogFilterExtList,
		dialog.GenDialogFilterCacheKey(in.UserId),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			// vList :=
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

	return &dialog.Vector_DialogFilterExt{
		Datas: dialogFilterExtList,
	}, nil
}
