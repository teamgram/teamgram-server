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

	c.svcCtx.Dao.DialogFiltersDAO.SelectListWithCB(
		c.ctx,
		in.UserId,
		func(i int, v *dataobject.DialogFiltersDO) {
			dialogFilter := &dialog.DialogFilterExt{
				Id:           v.DialogFilterId,
				DialogFilter: nil,
				Order:        v.OrderValue,
			}

			if err := jsonx.UnmarshalFromString(v.DialogFilter, &dialogFilter.DialogFilter); err != nil {
				c.Logger.Errorf("json.Unmarshal(%v) - error: %v", v, err)
				// continue
				return
			}

			if dialogFilter.DialogFilter == nil {
				dialogFilter.DialogFilter = mtproto.MakeTLDialogFilter(nil).To_DialogFilter()
			}

			dialogFilterExtList = append(dialogFilterExtList, dialogFilter)
		})

	return &dialog.Vector_DialogFilterExt{
		Datas: dialogFilterExtList,
	}, nil
}
