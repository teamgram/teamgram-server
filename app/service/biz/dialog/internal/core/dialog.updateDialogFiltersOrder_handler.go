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
)

// DialogUpdateDialogFiltersOrder
// dialog.updateDialogFiltersOrder user_id:long order:Vector<long> = Bool;
func (c *DialogCore) DialogUpdateDialogFiltersOrder(in *dialog.TLDialogUpdateDialogFiltersOrder) (*mtproto.Bool, error) {
	var (
		err    error
		orderV = time.Now().Unix() << 32
	)

	c.svcCtx.Dao.CachedConn.Exec(
		c.ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			tR := sqlx.TxWrapper(
				ctx,
				conn,
				func(tx *sqlx.Tx, result *sqlx.StoreResult) {
					for _, id := range in.Order {
						_, err = c.svcCtx.DialogFiltersDAO.UpdateOrder(ctx, orderV, in.UserId, id)
						if err != nil {
							result.Err = err
							return
						}
						orderV--
					}
				})

			return 0, 0, tR.Err
		},
		dialog.GenDialogFilterCacheKey(in.UserId))

	return mtproto.BoolTrue, nil
}
