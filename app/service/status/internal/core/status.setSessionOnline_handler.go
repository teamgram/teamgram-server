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
	)

	sessData, _ := jsonx.Marshal(sess)
	err := c.svcCtx.Dao.KV.HsetCtx(
		c.ctx,
		userK,
		strconv.FormatInt(sess.GetAuthKeyId(), 10),
		string(sessData))
	if err != nil {
		c.Logger.Errorf("status.setSessionOnline(%s) error(%v)", in, err)
		return nil, err
	}

	_, err = c.svcCtx.Dao.KV.ExpireCtx(
		c.ctx,
		userK,
		c.svcCtx.Config.StatusExpire)
	if err != nil {
		c.Logger.Errorf("status.setSessionOnline(%s) error(%v)", in, err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
