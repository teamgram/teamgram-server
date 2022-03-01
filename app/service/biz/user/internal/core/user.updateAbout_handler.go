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

// UserUpdateAbout
// user.updateAbout user_id:long about:string = Bool;
func (c *UserCore) UserUpdateAbout(in *user.TLUserUpdateAbout) (*mtproto.Bool, error) {
	rowsAffected, err := c.svcCtx.Dao.UsersDAO.UpdateUser(c.ctx, map[string]interface{}{
		"about": in.About,
	}, in.UserId)

	if err != nil {
		c.Logger.Errorf("user.updateAbout - error: %v", err)
		return mtproto.BoolFalse, nil
	}

	return mtproto.ToBool(rowsAffected != 0), nil
}
