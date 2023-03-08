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

// UserGetImmutableUserByPhone
// user.getImmutableUserByPhone phone:string = ImmutableUser;
func (c *UserCore) UserGetImmutableUserByPhone(in *user.TLUserGetImmutableUserByPhone) (*mtproto.ImmutableUser, error) {
	// TODO: performance optimization
	do, err := c.svcCtx.Dao.UsersDAO.SelectByPhoneNumber(c.ctx, in.Phone)
	if err != nil {
		c.Logger.Errorf("user.getImmutableUserByPhone - error: %v", err)
		return nil, err
	} else if do == nil {
		err = mtproto.ErrPhoneNumberUnoccupied
		c.Logger.Errorf("user.getImmutableUserByPhone - error: %v", err)
		return nil, err
	}

	return c.UserGetImmutableUser(&user.TLUserGetImmutableUser{
		Id: do.Id,
	})
}
