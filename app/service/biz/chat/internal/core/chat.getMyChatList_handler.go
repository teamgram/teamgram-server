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
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// ChatGetMyChatList
// chat.getMyChatList user_id:long is_creator:Bool = Vector<MutableChat>;
func (c *ChatCore) ChatGetMyChatList(in *chat.TLChatGetMyChatList) (*chat.Vector_MutableChat, error) {
	var (
		chatList = make([]*mtproto.MutableChat, 0)
	)

	//
	if mtproto.FromBool(in.IsCreator) {
		c.svcCtx.Dao.ChatParticipantsDAO.SelectMyAdminListWithCB(
			c.ctx,
			in.UserId,
			func(sz, i int, v int64) {
				chat, err := c.svcCtx.Dao.GetMutableChat(c.ctx, v, in.UserId)
				if err != nil {
					c.Logger.Errorf("chat.getMyChatList - error: %v", err)
				} else if chat != nil {
					chatList = append(chatList, chat)
				}
			})
	} else {
		// TODO(@benqi):
	}

	return &chat.Vector_MutableChat{
		Datas: chatList,
	}, nil
}
