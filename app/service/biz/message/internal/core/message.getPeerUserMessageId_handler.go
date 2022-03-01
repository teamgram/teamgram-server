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

// MessageGetPeerUserMessageId
// message.getPeerUserMessageId user_id:long peer_user_id:long msg_id:int = Int32;
func (c *MessageCore) MessageGetPeerUserMessageId(in *message.TLMessageGetPeerUserMessageId) (*mtproto.Int32, error) {
	// TODO: not impl
	c.Logger.Errorf("message.getPeerUserMessageId - error: method MessageGetPeerUserMessageId not impl")

	return nil, mtproto.ErrMethodNotImpl
}
