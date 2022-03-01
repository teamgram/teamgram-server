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
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
)

// MsgReadChannelHistory
// msg.readChannelHistory user_id:long auth_key_id:long channel_id:long max_id:int = Bool;
func (c *MsgCore) MsgReadChannelHistory(in *msg.TLMsgReadChannelHistory) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("account.confirmPhone blocked, License key from https://teamgram.net required to unlock enterprise features.")

	return mtproto.BoolTrue, nil
}
