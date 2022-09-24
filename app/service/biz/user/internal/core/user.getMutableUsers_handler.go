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
	"github.com/zeromicro/go-zero/core/mr"
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
	}

	mUsers := make([]*user.ImmutableUser, 0, len(id))
	mr.ForEach(
		func(source chan<- interface{}) {
			for idx := 0; idx < len(id); idx++ {
				source <- idx
			}
		},
		func(item interface{}) {
			var (
				idx = item.(int)
				rU  *user.ImmutableUser
				err error
			)

			if ok, _ := container2.Contains(id[idx], in.To); ok {
				rU, err = c.svcCtx.Dao.GetImmutableUser(c.ctx, id[idx], true, in.Id...)
				if err != nil {
					c.Logger.Errorf("getImmutableUser - error: %v", err)
				}
			} else {
				if len(in.To) == 0 {
					rU, err = c.svcCtx.Dao.GetImmutableUser(c.ctx, id[idx], true, in.Id...)
					if err != nil {
						c.Logger.Errorf("getImmutableUser - error: %v", err)
					}
				} else {
					rU, err = c.svcCtx.Dao.GetImmutableUser(c.ctx, id[idx], true, in.To...)
					if err != nil {
						c.Logger.Errorf("getImmutableUser - error: %v", err)
					}
				}
			}
			if rU != nil {
				mUsers = append(mUsers, rU)
			}
		})

	for _, v := range mUsers {
		if v != nil {
			vUser.Datas = append(vUser.Datas, v)
		}
	}

	return vUser, nil
}
