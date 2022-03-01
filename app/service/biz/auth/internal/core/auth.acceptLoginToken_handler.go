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
	"github.com/teamgram/teamgram-server/app/service/biz/auth/auth"
)

// AuthAcceptLoginToken
// auth.acceptLoginToken token:bytes = Authorization;
func (c *AuthCore) AuthAcceptLoginToken(in *auth.TLAuthAcceptLoginToken) (*mtproto.Authorization, error) {
	// TODO: not impl
	c.Logger.Errorf("auth.acceptLoginToken - error: method AuthAcceptLoginToken not impl")

	return nil, mtproto.ErrMethodNotImpl
}
