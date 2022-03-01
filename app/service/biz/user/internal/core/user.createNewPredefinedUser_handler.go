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

// UserCreateNewPredefinedUser
// user.createNewPredefinedUser flags:# phone:string first_name:string last_name:flags.0?string username:string code:string verified:flags.1?true = PredefinedUser;
func (c *UserCore) UserCreateNewPredefinedUser(in *user.TLUserCreateNewPredefinedUser) (*mtproto.PredefinedUser, error) {
	// TODO: not impl
	c.Logger.Errorf("user.createNewPredefinedUser - error: method UserCreateNewPredefinedUser not impl")

	return nil, mtproto.ErrMethodNotImpl
}
