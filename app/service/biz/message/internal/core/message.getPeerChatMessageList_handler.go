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
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
)

// MessageGetPeerChatMessageList
// message.getPeerChatMessageList user_id:long peer_chat_id:long msg_id:int = Vector<MessageBox>;
func (c *MessageCore) MessageGetPeerChatMessageList(in *message.TLMessageGetPeerChatMessageList) (*message.Vector_MessageBox, error) {
	rValueList := &message.Vector_MessageBox{
		Datas: make([]*mtproto.MessageBox, 0),
	}
	//
	//peerMsgs = make(map[int32]*model.MessageBox)
	myDO, err := c.svcCtx.Dao.MessagesDAO.SelectByMessageId(c.ctx, in.UserId, in.MsgId)
	if err != nil || myDO == nil {
		err = mtproto.ErrMsgIdInvalid
		c.Logger.Errorf("message.getPeerUserMessage - error: %v", err)
		return nil, err
	}

	_, err = c.svcCtx.Dao.MessagesDAO.SelectByMessageDataIdListWithCB(
		c.ctx,
		[]int64{myDO.DialogMessageId},
		func(i int, v *dataobject.MessagesDO) {
			rValueList.Datas = append(rValueList.Datas, c.svcCtx.Dao.MakeMessageBox(c.ctx, in.UserId, v))
		})
	if err != nil {
		c.Logger.Errorf("message.getPeerUserMessage - error: %v", err)
		return nil, err
	}

	return rValueList, nil
}
