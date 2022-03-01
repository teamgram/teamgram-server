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
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserSetBotCommands
// user.setBotCommands user_id:long bot_id:long commands:Vector<BotCommand> = Bool;
func (c *UserCore) UserSetBotCommands(in *user.TLUserSetBotCommands) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("user.setBotCommands - error: method UserSetBotCommands not impl")

	return nil, mtproto.ErrMethodNotImpl
}
