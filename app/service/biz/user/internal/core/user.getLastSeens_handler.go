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

// UserGetLastSeens
// user.getLastSeens id:Vector<long> = Vector<LastSeenData>;
func (c *UserCore) UserGetLastSeens(in *user.TLUserGetLastSeens) (*user.Vector_LastSeenData, error) {
	doList, _ := c.svcCtx.Dao.UserPresencesDAO.SelectList(c.ctx, in.Id)

	rValues := &user.Vector_LastSeenData{
		Datas: make([]*user.LastSeenData, 0, len(doList)),
	}

	for i := 0; i < len(doList); i++ {
		rValues.Datas = append(rValues.Datas, user.MakeTLLastSeenData(&user.LastSeenData{
			UserId:     doList[i].UserId,
			LastSeenAt: doList[i].LastSeenAt,
			Expries:    doList[i].Expires,
		}).To_LastSeenData())
	}

	return rValues, nil
}
