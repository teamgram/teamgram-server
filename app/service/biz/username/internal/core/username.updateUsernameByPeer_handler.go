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
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/username/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/username/username"
)

// UsernameUpdateUsernameByPeer
// username.updateUsernameByPeer peer_type:int peer_id:int username:string = Bool;
func (c *UsernameCore) UsernameUpdateUsernameByPeer(in *username.TLUsernameUpdateUsernameByPeer) (*mtproto.Bool, error) {
	_, _, err := c.svcCtx.Dao.UsernameDAO.Insert(c.ctx, &dataobject.UsernameDO{
		Username: in.Username,
		PeerType: in.PeerType,
		PeerId:   in.PeerId,
	})

	if err != nil {
		if sqlx.IsDuplicate(err) {
			return mtproto.BoolFalse, nil
		} else {
			c.Logger.Errorf("username.updateUsername - error: %v", err)
			return nil, err
		}
	}

	return mtproto.BoolTrue, nil
}
