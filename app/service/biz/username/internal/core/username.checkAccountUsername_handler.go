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

// UsernameCheckAccountUsername
// username.checkAccountUsername user_id:int username:string = UsernameExist;
func (c *UsernameCore) UsernameCheckAccountUsername(in *username.TLUsernameCheckAccountUsername) (*username.UsernameExist, error) {
	var (
		checked = usernameNotExisted
	)

	// TODO(@benqi): check len(username) >= 5
	usernameDO, _ := c.svcCtx.UsernameDAO.SelectByUsername(c.ctx, in.Username)
	if usernameDO != nil {
		if usernameDO.PeerType == mtproto.PEER_USER && usernameDO.PeerId == in.UserId {
			checked = usernameExistedIsMe
		} else {
			checked = usernameExistedNotMe
		}
	}

	return checked, nil
}
