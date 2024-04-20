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
	"github.com/teamgram/marmota/pkg/hack"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"

	"github.com/zeromicro/go-zero/core/jsonx"
)

// DialogSaveDraftMessage
// dialog.saveDraftMessage user_id:long peer_type:int peer_id:long message:DraftMessage = Bool;
func (c *DialogCore) DialogSaveDraftMessage(in *dialog.TLDialogSaveDraftMessage) (*mtproto.Bool, error) {
	draft, _ := jsonx.Marshal(in.Message)

	_, err := c.svcCtx.Dao.DialogsDAO.SaveDraft(c.ctx,
		2,
		hack.String(draft),
		in.UserId,
		in.PeerType,
		in.PeerId)
	if err != nil {
		c.Logger.Errorf("dialog.saveDraftMessage - error: %v", err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
