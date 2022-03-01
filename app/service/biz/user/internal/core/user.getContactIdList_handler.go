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
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserGetContactIdList
// user.getContactIdList user_id:long = Vector<long>;
func (c *UserCore) UserGetContactIdList(in *user.TLUserGetContactIdList) (*user.Vector_Long, error) {
	rValList := &user.Vector_Long{
		Datas: []int64{},
	}

	c.svcCtx.UserContactsDAO.SelectUserContactsWithCB(
		c.ctx,
		in.UserId,
		func(i int, v *dataobject.UserContactsDO) {
			rValList.Datas = append(rValList.Datas, v.Id)
		})

	return rValList, nil
}
