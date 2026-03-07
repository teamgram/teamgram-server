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
	"fmt"
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
	)

	if in.GetUserId() <= 0 || sess == nil || sess.GetAuthKeyId() == 0 {
		return nil, fmt.Errorf("status.setSessionOnline - invalid params: userId=%d, session=%v", in.GetUserId(), sess)
	}

	sessData, err := jsonx.Marshal(sess)
	if err != nil {
		c.Logger.Errorf("status.setSessionOnline - marshal session error: %v", err)
		return nil, fmt.Errorf("status.setSessionOnline - marshal session: %w", err)
	}

	err = c.svcCtx.Dao.KV.HsetCtx(
		c.ctx,
		userK,
		strconv.FormatInt(sess.GetAuthKeyId(), 10),
		string(sessData))
	if err != nil {
		c.Logger.Errorf("status.setSessionOnline - hset(userId=%d) error: %v", in.GetUserId(), err)
		return nil, fmt.Errorf("status.setSessionOnline - hset: %w", err)
	}

	_, err = c.svcCtx.Dao.KV.ExpireCtx(
		c.ctx,
		userK,
		c.svcCtx.Config.StatusExpire)
	if err != nil {
		c.Logger.Errorf("status.setSessionOnline - expire(userId=%d) error: %v", in.GetUserId(), err)
		return nil, fmt.Errorf("status.setSessionOnline - expire: %w", err)
	}

	return mtproto.BoolTrue, nil
}
