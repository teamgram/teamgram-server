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
	"time"

	"github.com/zeromicro/go-zero/core/jsonx"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
)

func getEmptyDraftMessage() string {
	draftData, _ := jsonx.Marshal(
		mtproto.MakeTLDraftMessageEmpty(&mtproto.DraftMessage{
			Date_FLAGINT32: mtproto.MakeFlagsInt32(int32(time.Now().Unix())),
		}).To_DraftMessage())
	return string(draftData)
}

// DialogClearDraftMessage
// dialog.clearDraftMessage user_id:long peer_type:int peer_id:long = Bool;
func (c *DialogCore) DialogClearDraftMessage(in *dialog.TLDialogClearDraftMessage) (*mtproto.Bool, error) {
	_, err := c.svcCtx.Dao.DialogsDAO.SaveDraft(c.ctx,
		1,
		getEmptyDraftMessage(),
		in.UserId,
		in.PeerType,
		in.PeerId)
	if err != nil {
		c.Logger.Errorf("dialog.clearDraftMessage - error: %v", err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
