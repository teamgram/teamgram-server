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
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserGetImmutableUser
// user.getImmutableUser id:long = ImmutableUser;
func (c *UserCore) UserGetImmutableUser(in *user.TLUserGetImmutableUser) (*user.ImmutableUser, error) {
	imUser, err := c.svcCtx.Dao.GetImmutableUser(c.ctx, in.GetId(), false)
	if err != nil {
		return nil, err
	}

	return imUser, nil
}
