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
	do, _ := c.svcCtx.Dao.UserPeerBlocksDAO.Select(c.ctx, in.UserId, mtproto.PEER_USER, in.PeerUserId)

	return mtproto.ToBool(do != nil), nil
}
