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
	"github.com/teamgram/teamgram-server/app/service/biz/username/username"
)

// UsernameResolveUsername
// username.resolveUsername username:string = Peer;
func (c *UsernameCore) UsernameResolveUsername(in *username.TLUsernameResolveUsername) (*mtproto.Peer, error) {
	var (
		peer     *mtproto.Peer
		err      error
		username = in.Username
	)

	switch username {
	case "gif":
	case "pic":
	case "bing":
	default:
		if len(username) < 5 {
			err = mtproto.ErrUsernameInvalid
			return nil, err
		}
	}

	usernameDO, _ := c.svcCtx.Dao.UsernameDAO.SelectByUsername(c.ctx, username)
	if usernameDO == nil {
		c.Logger.Errorf("username.resolveUsername - error: %v", err)
		err = mtproto.ErrUsernameNotOccupied
		return nil, err
	}

	switch usernameDO.PeerType {
	case mtproto.PEER_USER:
		peer = mtproto.MakePeerUser(usernameDO.PeerId)
	case mtproto.PEER_CHANNEL:
		peer = mtproto.MakePeerChannel(usernameDO.PeerId)
	default:
		err = mtproto.ErrUsernameNotOccupied
		c.Logger.Errorf("username.resolveUsername - error: %v", err)
		return nil, err
	}

	return peer, nil
}
