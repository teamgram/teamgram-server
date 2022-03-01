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
	"strconv"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/status/status"
)

// StatusSetSessionOnline
// status.setSessionOnline online:SessionEntry = Bool;
func (c *StatusCore) StatusSetSessionOnline(in *status.TLStatusSetSessionOnline) (*mtproto.Bool, error) {
	var (
		userK = getUserKey(in.GetUserId())
		authK = getAuthKeyIdKey(in.GetAuthKeyId())
	)

	if err := c.svcCtx.KV.Hset(
		userK,
		strconv.FormatInt(in.GetAuthKeyId(), 10),
		in.GetGateway()); err != nil {
		c.Logger.Errorf("status.setSessionOnline(%s) error(%v)", in.DebugString(), err)
		return nil, err
	}

	if err := c.svcCtx.KV.Expire(userK, c.svcCtx.Config.StatusExpire); err != nil {
		c.Logger.Errorf("status.setSessionOnline(%s) error(%v)", in.DebugString(), err)
		return nil, err
	}

	if err := c.svcCtx.KV.Setex(authK, in.GetGateway(), c.svcCtx.Config.StatusExpire); err != nil {
		c.Logger.Errorf("status.setSessionOnline(%s) error(%v)", in.DebugString(), err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
