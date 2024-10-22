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

// DialogGetDialogById
// dialog.getDialogById user_id:long peer_type:int peer_id:long = DialogExt;
func (c *DialogCore) DialogGetDialogById(in *dialog.TLDialogGetDialogById) (*dialog.DialogExt, error) {
	dlgExt, err := c.svcCtx.Dao.GetDialog(c.ctx, in.GetUserId(), in.GetPeerType(), in.GetPeerId())
	if err != nil {
		c.Logger.Errorf("dialog.getDialogById - error: %v", err)
		return nil, err
	}

	return dlgExt, nil
}
