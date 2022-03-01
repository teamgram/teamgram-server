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
	// TODO: not impl
	c.Logger.Errorf("message.getPeerUserMessage - error: method MessageGetPeerUserMessage not impl")

	return nil, mtproto.ErrMethodNotImpl
}
