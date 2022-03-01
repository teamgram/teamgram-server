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

// UserGetPredefinedUser
// user.getPredefinedUser phone:string = PredefinedUser;
func (c *UserCore) UserGetPredefinedUser(in *user.TLUserGetPredefinedUser) (*mtproto.PredefinedUser, error) {
	// TODO: not impl
	c.Logger.Errorf("user.getPredefinedUser - error: method UserGetPredefinedUser not impl")

	return nil, mtproto.ErrMethodNotImpl
}
