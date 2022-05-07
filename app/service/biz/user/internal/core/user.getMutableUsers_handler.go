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
	"github.com/zeromicro/go-zero/core/mr"

	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserGetMutableUsers
// user.getMutableUsers id:Vector<int> = Vector<ImmutableUser>;
func (c *UserCore) UserGetMutableUsers(in *user.TLUserGetMutableUsers) (*user.Vector_ImmutableUser, error) {
	vUser := &user.Vector_ImmutableUser{
		Datas: make([]*user.ImmutableUser, 0, len(in.Id)),
	}

	if len(in.Id) == 0 {
		return vUser, nil
	} else if len(in.Id) == 1 {
		immutableUser, _ := c.getImmutableUser(in.Id[0], false, false)
		if immutableUser != nil {
			vUser.Datas = append(vUser.Datas, immutableUser)
		}

		return vUser, nil
	}

	mutableUsers := make([]*user.ImmutableUser, len(in.Id))
	mr.ForEach(
		func(source chan<- interface{}) {
			for idx := 0; idx < len(in.Id); idx++ {
				source <- idx
			}
		},
		func(item interface{}) {
			idx := item.(int)
			immutableUser, _ := c.getImmutableUser(in.Id[idx], true, true)
			mutableUsers[idx] = immutableUser
		})

	for _, v := range mutableUsers {
		if v != nil {
			vUser.Datas = append(vUser.Datas, v)
		}
	}

	return vUser, nil
}
