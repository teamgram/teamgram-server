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
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
)

// MessageGetPeerChatMessageIdList
// message.getPeerChatMessageIdList user_id:long peer_chat_id:long msg_id:int = Vector<PeerMessageId>;
func (c *MessageCore) MessageGetPeerChatMessageIdList(in *message.TLMessageGetPeerChatMessageIdList) (*message.Vector_PeerMessageId, error) {
	var (
		idList = make([]*message.PeerMessageId, 0)
	)

	c.svcCtx.Dao.MessagesDAO.SelectDialogMessageListByMessageIdWithCB(
		c.ctx,
		in.UserId,
		in.MsgId,
		func(i int, v *dataobject.MessagesDO) {
			if v.UserId != in.UserId {
				idList = append(idList, &message.PeerMessageId{
					UserId: v.UserId,
					MsgId:  v.UserMessageBoxId,
				})
			}
		})

	return &message.Vector_PeerMessageId{
		Datas: idList,
	}, nil
}
