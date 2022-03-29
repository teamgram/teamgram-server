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
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
)

// MessageGetPeerUserMessage
// message.getPeerUserMessage user_id:long peer_user_id:long msg_id:int = MessageBox;
func (c *MessageCore) MessageGetPeerUserMessage(in *message.TLMessageGetPeerUserMessage) (*mtproto.MessageBox, error) {
	pDO, err := c.svcCtx.Dao.MessagesDAO.SelectPeerUserMessage(c.ctx, in.PeerUserId, in.UserId, in.MsgId)
	if err != nil {
		c.Logger.Errorf("message.getPeerUserMessage - error: %v", err)
		return nil, err
	} else if pDO == nil {
		err = mtproto.ErrMsgIdInvalid
		c.Logger.Errorf("message.getPeerUserMessage - error: %v", err)
		return nil, err
	}

	return c.svcCtx.Dao.MakeMessageBox(c.ctx, in.UserId, pDO), nil
}
