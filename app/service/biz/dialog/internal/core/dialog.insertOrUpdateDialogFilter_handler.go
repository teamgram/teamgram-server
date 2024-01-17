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
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/jsonx"
)

// DialogInsertOrUpdateDialogFilter
// dialog.insertOrUpdateDialogFilter user_id:long id:int dialog_filter:DialogFilter = Bool;
func (c *DialogCore) DialogInsertOrUpdateDialogFilter(in *dialog.TLDialogInsertOrUpdateDialogFilter) (*mtproto.Bool, error) {
	dialogFilterData, err := jsonx.Marshal(in.GetDialogFilter())
	isChatlist := in.GetDialogFilter().GetPredicateName() == mtproto.Predicate_dialogFilterChatlist

	if err != nil {
		c.Logger.Errorf("dialog.insertOrUpdateDialogFilter - error: %v", err)
		return nil, err
	}

	c.svcCtx.Dao.CachedConn.Exec(
		c.ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			_, _, err2 := c.svcCtx.Dao.DialogFiltersDAO.InsertOrUpdate(
				ctx,
				&dataobject.DialogFiltersDO{
					UserId:         in.UserId,
					DialogFilterId: in.Id,
					IsChatlist:     isChatlist,
					DialogFilter:   string(dialogFilterData),
					OrderValue:     time.Now().Unix() << 32,
					FromSuggested:  -1,
					Deleted:        false,
				})

			return 0, 0, err2
		},
		dialog.GenDialogFilterCacheKey(in.UserId))

	return mtproto.BoolTrue, nil
}
