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
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/jsonx"
)

// DialogGetPinnedDialogs
// dialog.getPinnedDialogs  user_id:long folder_id:int = Vector<DialogExt>;
func (c *DialogCore) DialogGetPinnedDialogs(in *dialog.TLDialogGetPinnedDialogs) (*dialog.Vector_DialogExt, error) {
	var (
		dList    dialog.DialogExtList
		folderId = in.GetFolderId()
		meId     = in.GetUserId()
		dIdList  []int64
		err      error
	)

	if folderId == 0 {
		dIdList, err = c.svcCtx.Dao.GetPinnedDialogIdList(c.ctx, meId)
		if err != nil {
			c.Logger.Errorf("dialog.getPinnedDialogs - error: %v", err)
			return nil, err
		}
	} else {
		dIdList, err = c.svcCtx.Dao.GetFolderPinnedDialogIdList(c.ctx, meId)
		if err != nil {
			c.Logger.Errorf("dialog.getPinnedDialogs - error: %v", err)
			return nil, err
		}
	}
	if len(dIdList) == 0 {
		return &dialog.Vector_DialogExt{
			Datas: make(dialog.DialogExtList, 0),
		}, nil
	}

	keyList := make([]string, 0, len(dIdList))
	for _, id := range dIdList {
		keyList = append(keyList, dialog.GetDialogCacheKey(meId, id))
	}

	c.svcCtx.Dao.CachedConn.QueryRows(
		c.ctx,
		func(ctx context.Context, conn *sqlx.DB, keys ...string) (map[string]interface{}, error) {
			vList := make(map[string]interface{}, len(keys))

			// TODO: mr
			dIdList2 := make([]int64, 0, len(keys))
			for _, key := range keys {
				_, id := dialog.ParseDialogCacheKey(key)
				dIdList2 = append(dIdList2, id)
			}

			_, err2 := c.svcCtx.Dao.DialogsDAO.SelectPeerDialogListWithCB(
				ctx,
				meId,
				dIdList2,
				func(sz, i int, v *dataobject.DialogsDO) {
					dlgExt := makeDialog(v)
					vList[dialog.GetDialogCacheKey(meId, v.PeerDialogId)] = dlgExt
					dList = append(dList, dlgExt)
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
				return nil, err
			}

			dList = append(dList, dlgExt)
			return dlgExt, nil
		},
		keyList...)

	return &dialog.Vector_DialogExt{
		Datas: dList,
	}, nil
}
