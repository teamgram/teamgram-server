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

// UserAddPeerSettings
// user.addPeerSettings user_id:int peer_type:int peer_id:int settings:PeerSettings = Bool;
func (c *UserCore) UserAddPeerSettings(in *user.TLUserAddPeerSettings) (*mtproto.Bool, error) {
	err := c.svcCtx.Dao.SetUserPeerSettings(
		c.ctx,
		in.GetUserId(),
		in.GetPeerType(),
		in.GetPeerId(),
		in.GetSettings())

	if err != nil {
		c.Logger.Errorf("user.addPeerSettings - error: %v", err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
