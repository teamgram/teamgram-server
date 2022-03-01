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

// MessageGetPeerChatMessageList
// message.getPeerChatMessageList user_id:long peer_chat_id:long msg_id:int = Vector<MessageBox>;
func (c *MessageCore) MessageGetPeerChatMessageList(in *message.TLMessageGetPeerChatMessageList) (*message.Vector_MessageBox, error) {
	// TODO: not impl
	c.Logger.Errorf("message.getPeerChatMessageList - error: method MessageGetPeerChatMessageList not impl")

	return nil, mtproto.ErrMethodNotImpl
}
