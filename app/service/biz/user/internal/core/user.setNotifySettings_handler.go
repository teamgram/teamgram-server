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

// UserSetNotifySettings
// user.setNotifySettings user_id:int peer_type:int peer_id:int settings:PeerNotifySettings = Bool;
func (c *UserCore) UserSetNotifySettings(in *user.TLUserSetNotifySettings) (*mtproto.Bool, error) {
	err := c.svcCtx.Dao.SetUserPeerNotifySettings(
		c.ctx,
		in.GetUserId(),
		in.GetPeerType(),
		in.GetPeerId(),
		in.GetSettings())

	if err != nil {
		c.Logger.Errorf("user.setNotifySettings - error: %v", err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
