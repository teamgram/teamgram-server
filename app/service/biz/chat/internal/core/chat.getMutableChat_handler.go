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

// ChatGetMutableChat
// chat.getMutableChat chat_id:long = MutableChat;
func (c *ChatCore) ChatGetMutableChat(in *chat.TLChatGetMutableChat) (*mtproto.MutableChat, error) {
	chat2, err := c.svcCtx.Dao.GetMutableChat(c.ctx, in.ChatId)
	if err != nil {
		c.Logger.Errorf("chat.getMutableChat - error: %v", err)
		return nil, err
	}

	return chat2, nil
}
