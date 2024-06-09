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

	"github.com/zeromicro/go-zero/core/timex"
)

var (
	useV2 = true
)

// UserGetMutableUsers
// user.getMutableUsers id:Vector<int> = Vector<ImmutableUser>;
func (c *UserCore) UserGetMutableUsers(in *user.TLUserGetMutableUsers) (*user.Vector_ImmutableUser, error) {
	vUser := &user.Vector_ImmutableUser{
		Datas: nil,
	}

	since2 := timex.Now()
	if useV2 {
		vUser.Datas = c.svcCtx.Dao.GetCacheImmutableUserListV2(c.ctx, in.Id, in.To)
		if len(vUser.Datas) == 0 {
			c.Logger.Errorf("user.getMutableUsersV2 - not found users: {id: %v, to: %v}", in.To, in.To)
		}
		c.Logger.WithDuration(timex.Since(since2)).Infof("V2: len(%d)", len(in.Id))
	} else {
		vUser.Datas = c.svcCtx.Dao.GetCacheImmutableUserList(c.ctx, in.Id, in.To)
		if len(vUser.Datas) == 0 {
			c.Logger.Errorf("user.getMutableUsers - not found users: {id: %v, to: %v}", in.To, in.To)
		}
		c.Logger.WithDuration(timex.Since(since2)).Infof("V1: len(%d)", len(in.Id))
	}

	return vUser, nil
}
