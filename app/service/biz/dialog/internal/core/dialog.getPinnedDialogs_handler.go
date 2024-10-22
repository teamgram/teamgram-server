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
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
)

// DialogGetPinnedDialogs
// dialog.getPinnedDialogs  user_id:long folder_id:int = Vector<DialogExt>;
func (c *DialogCore) DialogGetPinnedDialogs(in *dialog.TLDialogGetPinnedDialogs) (*dialog.Vector_DialogExt, error) {
	var (
		dlgExtList dialog.DialogExtList
		folderId   = in.GetFolderId()
		meId       = in.GetUserId()
		dIdList    []int64
		err        error
	)

	if folderId == 0 {
		// dIdList, err = c.svcCtx.Dao.GetPinnedDialogIdList(c.ctx, meId)
		dIdList, err = c.svcCtx.Dao.GetNoCachePinnedDialogIdList(c.ctx, meId)
		if err != nil {
			c.Logger.Errorf("dialog.getPinnedDialogs - error: %v", err)
			return nil, err
		}
	} else {
		// dIdList, err = c.svcCtx.Dao.GetFolderPinnedDialogIdList(c.ctx, meId)
		dIdList, err = c.svcCtx.Dao.GetNoCacheFolderPinnedDialogIdList(c.ctx, meId)
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

	dlgExtList, err = c.svcCtx.Dao.GetDialogListByIdList(c.ctx, meId, dIdList)
	if err != nil {
		c.Logger.Errorf("dialog.getPinnedDialogs - error: %v", err)
		return nil, err
	}

	return &dialog.Vector_DialogExt{
		Datas: dlgExtList,
	}, nil
}
