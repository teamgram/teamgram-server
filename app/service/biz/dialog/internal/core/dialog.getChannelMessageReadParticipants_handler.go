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

// DialogGetChannelMessageReadParticipants
// dialog.getChannelMessageReadParticipants user_id:long channel_id:long msg_id:int = Vector<long>;
func (c *DialogCore) DialogGetChannelMessageReadParticipants(in *dialog.TLDialogGetChannelMessageReadParticipants) (*dialog.Vector_Long, error) {
	idList, _ := c.svcCtx.Dao.DialogsDAO.SelectDialogsByGTReadInboxMaxId(
		c.ctx,
		mtproto.PEER_CHANNEL,
		in.ChannelId,
		in.MsgId,
		in.UserId)

	if idList == nil {
		idList = []int64{}
	}

	return &dialog.Vector_Long{
		Datas: idList,
	}, nil
}
