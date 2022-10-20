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
	"github.com/zeromicro/go-zero/core/mr"
)

// UserCheckBlockUserList
// user.checkBlockUserList user_id:long id:Vector<long> = Vector<long>;
func (c *UserCore) UserCheckBlockUserList(in *user.TLUserCheckBlockUserList) (*user.Vector_Long, error) {
	var (
		rVal = &user.Vector_Long{}
	)

	// c.svcCtx.Dao
	if len(in.Id) == 1 {
		if c.svcCtx.CheckBlocked(c.ctx, in.GetUserId(), in.Id[0]) {
			rVal.Datas = in.Id
		}
	} else if len(in.Id) > 1 {
		idList := make([]int64, len(in.Id))
		mr.ForEach(
			func(source chan<- interface{}) {
				for idx := 0; idx < len(in.Id); idx++ {
					source <- idx
				}
			},
			func(item interface{}) {
				idx := item.(int)
				if c.svcCtx.CheckBlocked(c.ctx, in.GetUserId(), in.Id[idx]) {
					idList[idx] = in.Id[idx]
				}
			})
		for _, id := range idList {
			if id != 0 {
				rVal.Datas = append(rVal.Datas, id)
			}
		}
	}

	if rVal.Datas == nil {
		rVal.Datas = []int64{}
	}

	return rVal, nil
}
