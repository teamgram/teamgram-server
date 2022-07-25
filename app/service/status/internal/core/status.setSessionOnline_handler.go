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

	"github.com/zeromicro/go-zero/core/jsonx"
)

// StatusSetSessionOnline
// status.setSessionOnline session:SessionEntry = Bool;
func (c *StatusCore) StatusSetSessionOnline(in *status.TLStatusSetSessionOnline) (*mtproto.Bool, error) {
	var (
		userK = getUserKey(in.GetUserId())
		sess  = in.GetSession()
		authK = getAuthKeyIdKey(sess.GetAuthKeyId())
	)

	sessData, _ := jsonx.Marshal(sess)

	if err := c.svcCtx.KV.Hset(
		userK,
		strconv.FormatInt(sess.GetAuthKeyId(), 10),
		string(sessData)); err != nil {
		c.Logger.Errorf("status.setSessionOnline(%s) error(%v)", in.DebugString(), err)
		return nil, err
	}

	if err := c.svcCtx.KV.Expire(userK, c.svcCtx.Config.StatusExpire); err != nil {
		c.Logger.Errorf("status.setSessionOnline(%s) error(%v)", in.DebugString(), err)
		return nil, err
	}

	if err := c.svcCtx.KV.Setex(authK, sess.GetGateway(), c.svcCtx.Config.StatusExpire); err != nil {
		c.Logger.Errorf("status.setSessionOnline(%s) error(%v)", in.DebugString(), err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
