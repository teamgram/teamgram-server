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

// UserUpdatePredefinedVerified
// user.updatePredefinedVerified flags:# phone:string verified:flags.1?true = PredefinedUser;
func (c *UserCore) UserUpdatePredefinedVerified(in *user.TLUserUpdatePredefinedVerified) (*mtproto.PredefinedUser, error) {
	// TODO: not impl
	c.Logger.Errorf("user.updatePredefinedVerified - error: method UserUpdatePredefinedVerified not impl")

	return nil, mtproto.ErrMethodNotImpl
}
