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

// UserGetImmutableUser
// user.getImmutableUser flags:# id:long privacy:flags.1?true contacts:Vector<long> = ImmutableUser;
func (c *UserCore) UserGetImmutableUser(in *user.TLUserGetImmutableUser) (*mtproto.ImmutableUser, error) {
	imUser, err := c.svcCtx.Dao.GetImmutableUser(c.ctx, in.GetId(), in.GetPrivacy(), in.Contacts...)
	if err != nil {
		return nil, err
	}

	return imUser, nil
}
