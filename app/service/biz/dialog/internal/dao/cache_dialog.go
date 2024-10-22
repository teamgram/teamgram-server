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

	"github.com/zeromicro/go-zero/core/jsonx"
)

func (d *Dao) GetNoCachePinnedDialogIdList(ctx context.Context, userId int64) ([]int64, error) {
	var (
		dialogIdList []int64
	)

	_, err := d.DialogsDAO.SelectPinnedDialogsWithCB(
		ctx,
		userId,
		func(sz, i int, v *dataobject.DialogsDO) {
			if i == 0 {
				dialogIdList = make([]int64, 0, sz)
			}
			dialogIdList = append(dialogIdList, v.PeerDialogId)
		})
	if err != nil {
		return nil, err
	}

	if dialogIdList == nil {
		dialogIdList = make([]int64, 0)
	}

	return dialogIdList, nil
}

func (d *Dao) GetPinnedDialogIdList(ctx context.Context, userId int64) ([]int64, error) {
	var (
		dialogIdList []int64
	)

	err := d.CachedConn.QueryRow(
		ctx,
		&dialogIdList,
		dialog.GetPinnedDialogIdListCacheKey(userId),
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

func (d *Dao) GetNoCacheNotPinnedDialogIdList(ctx context.Context, userId int64) ([]int64, error) {
	var (
		dialogIdList []int64
	)

	_, err := d.DialogsDAO.SelectExcludePinnedDialogsWithCB(
		ctx,
		userId,
		func(sz, i int, v *dataobject.DialogsDO) {
			if i == 0 {
				dialogIdList = make([]int64, 0, sz)
			}
			dialogIdList = append(dialogIdList, v.PeerDialogId)
		})
	if err != nil {
		return nil, err
	}

	if dialogIdList == nil {
		dialogIdList = make([]int64, 0)
	}

	return dialogIdList, nil
}

func (d *Dao) GetNotPinnedDialogIdList(ctx context.Context, userId int64) ([]int64, error) {
	var (
		dialogIdList []int64
	)

	err := d.CachedConn.QueryRow(
		ctx,
		&dialogIdList,
		dialog.GetNotPinnedDialogIdListCacheKey(userId),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			var (
				idList []int64
			)

			_, err := d.DialogsDAO.SelectExcludePinnedDialogsWithCB(
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

func (d *Dao) GetNoCacheFolderPinnedDialogIdList(ctx context.Context, userId int64) ([]int64, error) {
	var (
		dialogIdList []int64
	)

	_, err := d.DialogsDAO.SelectFolderPinnedDialogsWithCB(
		ctx,
		userId,
		func(sz, i int, v *dataobject.DialogsDO) {
			if i == 0 {
				dialogIdList = make([]int64, 0, sz)
			}
			dialogIdList = append(dialogIdList, v.PeerDialogId)
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
		dialog.GetFolderPinnedDialogIdListCacheKey(userId),
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

func (d *Dao) GetNoCacheFolderNotPinnedDialogIdList(ctx context.Context, userId int64) ([]int64, error) {
	var (
		dialogIdList []int64
	)

	_, err := d.DialogsDAO.SelectExcludeFolderPinnedDialogsWithCB(
		ctx,
		userId,
		func(sz, i int, v *dataobject.DialogsDO) {
			if i == 0 {
				dialogIdList = make([]int64, 0, sz)
			}
			dialogIdList = append(dialogIdList, v.PeerDialogId)
		})
	if err != nil {
		return nil, err
	}

	if dialogIdList == nil {
		dialogIdList = make([]int64, 0)
	}

	return dialogIdList, nil
}

func (d *Dao) GetFolderNotPinnedDialogIdList(ctx context.Context, userId int64) ([]int64, error) {
	var (
		dialogIdList []int64
	)

	err := d.CachedConn.QueryRow(
		ctx,
		&dialogIdList,
		dialog.GetFolderNotPinnedDialogIdListCacheKey(userId),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			var (
				idList []int64
			)

			_, err := d.DialogsDAO.SelectExcludeFolderPinnedDialogsWithCB(
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

func (d *Dao) GetDialogListByIdList(ctx context.Context, userId int64, idList []int64) ([]*dialog.DialogExt, error) {
	var (
		dlgExtList = make(dialog.DialogExtList, 0, len(idList))
		keyList    = make([]string, 0, len(idList))
	)

	for _, id := range idList {
		keyList = append(keyList, dialog.GetDialogCacheKey(userId, id))
	}

	err := d.CachedConn.QueryRows(
		ctx,
		func(ctx context.Context, conn *sqlx.DB, keys ...string) (map[string]interface{}, error) {
			vList := make(map[string]interface{}, len(keys))

			// TODO: mr
			dIdList2 := make([]int64, 0, len(keys))
			for _, key := range keys {
				_, id := dialog.ParseDialogCacheKey(key)
				dIdList2 = append(dIdList2, id)
			}

			_, err2 := d.DialogsDAO.SelectPeerDialogListWithCB(
				ctx,
				userId,
				dIdList2,
				func(sz, i int, v *dataobject.DialogsDO) {
					dlgExt := d.MakeDialog(v)
					vList[dialog.GetDialogCacheKey(userId, v.PeerDialogId)] = dlgExt
					dlgExtList = append(dlgExtList, dlgExt)
				})
			if err2 != nil {
				return nil, err2
			}

			return vList, nil
		},
		func(k, v string) (interface{}, error) {
			var (
				dlgExt *dialog.DialogExt
				err2   error
			)
			err2 = jsonx.UnmarshalFromString(v, &dlgExt)
			if err2 != nil {
				return nil, err2
			}

			dlgExtList = append(dlgExtList, dlgExt)
			return dlgExt, nil
		},
		keyList...)
	if err != nil {
		return nil, err
	}

	return dlgExtList, nil
}
