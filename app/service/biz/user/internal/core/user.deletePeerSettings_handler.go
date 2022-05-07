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

// UserDeletePeerSettings
// user.deletePeerSettings user_id:int peer_type:int peer_id:int = Bool;
func (c *UserCore) UserDeletePeerSettings(in *user.TLUserDeletePeerSettings) (*mtproto.Bool, error) {
	err := c.svcCtx.Dao.DeleteUserPeerSettings(c.ctx, in.UserId, in.PeerType, in.PeerId)

	if err != nil {
		c.Logger.Errorf("user.deletePeerSettings - error: %v", err)
	}

	return mtproto.BoolTrue, nil
}
