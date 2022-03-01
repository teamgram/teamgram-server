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

// UserUpdateUsername
// user.updateUsername user_id:long username:string = Bool;
func (c *UserCore) UserUpdateUsername(in *user.TLUserUpdateUsername) (*mtproto.Bool, error) {
	rowsAffected, err := c.svcCtx.Dao.UsersDAO.UpdateUser(c.ctx, map[string]interface{}{
		"username": in.Username,
	}, in.GetUserId())
	if err != nil {
		c.Logger.Errorf("user.updateUsername - error: %v", err)
		return mtproto.BoolFalse, nil
	}

	return mtproto.ToBool(rowsAffected != 0), nil
}
