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

// UserGetMutableUsers
// user.getMutableUsers id:Vector<int> = Vector<ImmutableUser>;
func (c *UserCore) UserGetMutableUsers(in *user.TLUserGetMutableUsers) (*user.Vector_ImmutableUser, error) {
	vDatas := c.svcCtx.Dao.GetCacheImmutableUserList(c.ctx, in.Id, in.To)
	if len(vDatas) == 0 {
		c.Logger.Errorf("user.getMutableUsers - not found users: {id: %v, to: %v}", in.To, in.To)
	}

	return &user.Vector_ImmutableUser{
		Datas: vDatas,
	}, nil
}
