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
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"time"
)

// DialogReorderPinnedDialogs
// dialog.reorderPinnedDialogs user_id:long force:Bool folder_id:int id_list:Vector<long> = Bool;
func (c *DialogCore) DialogReorderPinnedDialogs(in *dialog.TLDialogReorderPinnedDialogs) (*mtproto.Bool, error) {
	var (
		userId      = in.GetUserId()
		force       = mtproto.FromBool(in.GetForce())
		folderId    = in.GetFolderId()
		idList      = in.GetIdList()
		orderPinned = time.Now().Unix()
	)

	sqlx.TxWrapper(
		c.ctx,
		c.svcCtx.Dao.DB,
		func(tx *sqlx.Tx, result *sqlx.StoreResult) {
			if folderId == 0 {
				if force {
					_, result.Err = c.svcCtx.Dao.DialogsDAO.UpdateUnPinnedNotIdListTx(
						tx,
						userId,
						idList)
					if result.Err != nil {
						return
					}
				}

				for _, id := range idList {
					_, result.Err = c.svcCtx.Dao.DialogsDAO.UpdatePeerDialogListPinnedTx(
						tx,
						orderPinned<<32,
						in.UserId, []int64{id})
					if result.Err != nil {
						return
					}
					orderPinned -= 1
				}
			} else {
				if force {
					_, result.Err = c.svcCtx.DialogsDAO.UpdateFolderUnPinnedNotIdListTx(
						tx,
						userId,
						idList)
					if result.Err != nil {
						return
					}
				}

				for _, id := range idList {
					_, result.Err = c.svcCtx.Dao.DialogsDAO.UpdateFolderPeerDialogListPinnedTx(
						tx,
						orderPinned<<32,
						in.UserId,
						[]int64{id})
					orderPinned -= 1
				}
			}
		})

	return mtproto.BoolTrue, nil
}
