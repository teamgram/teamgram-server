// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"context"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/dal/dataobject"
)

func (d *Dao) GetPinnedDialogIdList(ctx context.Context, userId int64) ([]int64, error) {
	var (
		dialogIdList []int64
	)

	err := d.CachedConn.QueryRow(
		ctx,
		&dialogIdList,
		dialog.GetPinnedDialogListCacheKey(userId),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			var (
				idList []int64
			)

			_, err := d.DialogsDAO.SelectPinnedDialogsWithCB(
				ctx,
				userId,
				func(sz, i int, v *dataobject.DialogsDO) {
					if i == 0 {
						idList = make([]int64, 0, sz)
					}
					idList = append(idList, v.PeerDialogId)
				})
			if err != nil {
				return err
			}

			*v.(*[]int64) = idList

			return nil
		})
	if err != nil {
		return nil, err
	}

	if dialogIdList == nil {
		dialogIdList = make([]int64, 0)
	}

	return dialogIdList, nil
}

func (d *Dao) GetFolderPinnedDialogIdList(ctx context.Context, userId int64) ([]int64, error) {
	var (
		dialogIdList []int64
	)

	err := d.CachedConn.QueryRow(
		ctx,
		&dialogIdList,
		dialog.GetFolderPinnedDialogListCacheKey(userId),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			var (
				idList []int64
			)

			_, err := d.DialogsDAO.SelectFolderPinnedDialogsWithCB(
				ctx,
				userId,
				func(sz, i int, v *dataobject.DialogsDO) {
					if i == 0 {
						idList = make([]int64, 0, sz)
					}
					idList = append(idList, v.PeerDialogId)
				})
			if err != nil {
				return err
			}

			*v.(*[]int64) = idList

			return nil
		})
	if err != nil {
		return nil, err
	}

	if dialogIdList == nil {
		dialogIdList = make([]int64, 0)
	}

	return dialogIdList, nil
}
