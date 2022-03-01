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

// UserCheckBlockUserList
// user.checkBlockUserList user_id:long id:long = Vector<long>;
func (c *UserCore) UserCheckBlockUserList(in *user.TLUserCheckBlockUserList) (*user.Vector_Long, error) {
	var (
		rVal = &user.Vector_Long{
			Datas: []int64{},
		}
	)

	if len(in.Id) > 0 {
		rVal.Datas, _ = c.svcCtx.Dao.UserPeerBlocksDAO.SelectListByIdList(c.ctx, in.UserId, in.Id)
	}

	return rVal, nil
}
