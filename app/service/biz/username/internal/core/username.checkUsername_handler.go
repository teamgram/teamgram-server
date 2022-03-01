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
	"github.com/teamgram/teamgram-server/app/service/biz/username/username"
)

// UsernameCheckUsername
// username.checkUsername username:string = UsernameExist;
func (c *UsernameCore) UsernameCheckUsername(in *username.TLUsernameCheckUsername) (*username.UsernameExist, error) {
	var (
		checked = usernameNotExisted
	)

	usernameDO, _ := c.svcCtx.Dao.UsernameDAO.SelectByUsername(c.ctx, in.Username)
	if usernameDO != nil {
		checked = usernameExisted
	}

	return checked, nil
}
