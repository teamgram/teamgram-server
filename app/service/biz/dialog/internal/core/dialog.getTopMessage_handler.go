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

// DialogGetTopMessage
// dialog.getTopMessage user_id:long peer_type:int peer_id:long = Int32;
func (c *DialogCore) DialogGetTopMessage(in *dialog.TLDialogGetTopMessage) (*mtproto.Int32, error) {
	var (
		rValue = &mtproto.Int32{
			V: 0,
		}
	)

	topMessage, err := c.svcCtx.Dao.DialogsDAO.SelectDialog(c.ctx, in.UserId, in.PeerType, in.PeerId)
	if err != nil {
		c.Logger.Errorf("dialog.getTopMessage - error: %v", err)
	} else if topMessage == nil {
		c.Logger.Errorf("dialog.getTopMessage - error: not found dialog")
	} else {
		rValue.V = topMessage.TopMessage
	}

	return rValue, nil
}
