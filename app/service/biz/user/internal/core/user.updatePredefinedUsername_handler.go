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

// UserUpdatePredefinedUsername
// user.updatePredefinedUsername flags:# phone:string username:flags.1?string = PredefinedUser;
func (c *UserCore) UserUpdatePredefinedUsername(in *user.TLUserUpdatePredefinedUsername) (*mtproto.PredefinedUser, error) {
	// TODO: not impl
	c.Logger.Errorf("user.updatePredefinedUsername - error: method UserUpdatePredefinedUsername not impl")

	return nil, mtproto.ErrMethodNotImpl
}
