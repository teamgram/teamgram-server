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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserGetAllPredefinedUser
// user.getAllPredefinedUser = Vector<PredefinedUser>;
func (c *UserCore) UserGetAllPredefinedUser(in *user.TLUserGetAllPredefinedUser) (*user.Vector_PredefinedUser, error) {
	// TODO: not impl
	c.Logger.Errorf("user.getAllPredefinedUser - error: method UserGetAllPredefinedUser not impl")

	return &user.Vector_PredefinedUser{
		Datas: []*mtproto.PredefinedUser{},
	}, nil
}
