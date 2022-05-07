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

// UserBlockedByUser
// user.blockedByUser userId:long peer_user_id:long = Bool;
func (c *UserCore) UserBlockedByUser(in *user.TLUserBlockedByUser) (*mtproto.Bool, error) {
	blocked := c.svcCtx.Dao.CheckBlocked(c.ctx, in.GetUserId(), in.GetPeerUserId())

	return mtproto.ToBool(blocked), nil
}
