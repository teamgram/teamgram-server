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

// UserGetContactIdList
// user.getContactIdList user_id:long = Vector<long>;
func (c *UserCore) UserGetContactIdList(in *user.TLUserGetContactIdList) (*user.Vector_Long, error) {
	rValList := &user.Vector_Long{
		Datas: []int64{},
	}

	cacheUserData := c.svcCtx.Dao.GetCacheUserData(c.ctx, in.GetUserId())
	if len(cacheUserData.GetContactIdList()) > 0 {
		rValList.Datas = cacheUserData.GetContactIdList()
	}

	return rValList, nil
}
