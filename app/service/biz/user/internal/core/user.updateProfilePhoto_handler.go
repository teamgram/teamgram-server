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

// UserUpdateProfilePhoto
// user.updateProfilePhoto user_id:long id:long = Int64;
func (c *UserCore) UserUpdateProfilePhoto(in *user.TLUserUpdateProfilePhoto) (*mtproto.Int64, error) {
	rV := c.svcCtx.Dao.UpdateProfilePhoto(
		c.ctx,
		in.GetUserId(),
		in.GetId())

	return &mtproto.Int64{
		V: rV,
	}, nil
}
