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
	"github.com/teamgram/marmota/pkg/container2"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"

	"github.com/zeromicro/go-zero/core/timex"
)

// UserGetMutableUsers
// user.getMutableUsers id:Vector<int> = Vector<ImmutableUser>;
func (c *UserCore) UserGetMutableUsers(in *user.TLUserGetMutableUsers) (*user.Vector_ImmutableUser, error) {
	id := make([]int64, 0, len(in.Id)+len(in.To))
	for _, v := range in.Id {
		if ok, _ := container2.Contains(v, id); !ok {
			id = append(id, v)
		}
	}
	for _, v := range in.To {
		if ok, _ := container2.Contains(v, id); !ok {
			id = append(id, v)
		}
	}

	vUser := &user.Vector_ImmutableUser{
		Datas: make([]*user.ImmutableUser, 0, len(id)),
	}

	if len(id) == 0 {
		return vUser, nil
	} else if len(id) == 1 {
		immutableUser, _ := c.svcCtx.Dao.GetImmutableUser(c.ctx, id[0], false)
		if immutableUser != nil {
			vUser.Datas = append(vUser.Datas, immutableUser)
		}

		return vUser, nil
	} else {
		since2 := timex.Now()
		vUser.Datas = c.svcCtx.Dao.GetCacheImmutableUserList(c.ctx, in.Id, in.To)
		c.Logger.WithDuration(timex.Since(since2)).Infof("V1: len(%d)", len(in.Id))
	}

	return vUser, nil
}
