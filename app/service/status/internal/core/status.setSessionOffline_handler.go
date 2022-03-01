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

// StatusSetSessionOffline
// status.setSessionOffline online:SessionEntry = Bool;
func (c *StatusCore) StatusSetSessionOffline(in *status.TLStatusSetSessionOffline) (*mtproto.Bool, error) {
	var (
		userK = getUserKey(in.GetUserId())
		authK = getAuthKeyIdKey(in.GetAuthKeyId())
	)

	if _, err := c.svcCtx.KV.Hdel(userK, strconv.FormatInt(in.GetAuthKeyId(), 10)); err != nil {
		c.Logger.Errorf("status.setSessionOffline(%s) error(%v)", in.DebugString(), err)
		return nil, err
	}

	if _, err := c.svcCtx.KV.Del(authK); err != nil {
		c.Logger.Errorf("status.setSessionOffline(%s) error(%v)", in.DebugString(), err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
