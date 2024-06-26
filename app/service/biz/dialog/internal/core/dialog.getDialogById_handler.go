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
)

// DialogGetDialogById
// dialog.getDialogById user_id:long peer_type:int peer_id:long = DialogExt;
func (c *DialogCore) DialogGetDialogById(in *dialog.TLDialogGetDialogById) (*dialog.DialogExt, error) {
	dialogDO, err := c.svcCtx.Dao.SelectDialog(c.ctx, in.UserId, in.PeerType, in.PeerId)
	if err != nil {
		c.Logger.Errorf("dialog.getDialogById - error: %v", err)
		return nil, err
	} else if dialogDO == nil {
		c.Logger.Errorf("dialog.getDialogById - error: not found dialog (%s)", in)
		return nil, mtproto.ErrPeerIdInvalid
	}

	return makeDialog(dialogDO), nil
}
