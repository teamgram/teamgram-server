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
)

// StatusSetSessionOffline
// status.setSessionOffline online:SessionEntry = Bool;
func (c *StatusCore) StatusSetSessionOffline(in *status.TLStatusSetSessionOffline) (*mtproto.Bool, error) {
	if in.GetUserId() <= 0 || in.GetAuthKeyId() == 0 {
		return nil, fmt.Errorf("status.setSessionOffline - invalid params: userId=%d, authKeyId=%d", in.GetUserId(), in.GetAuthKeyId())
	}

	_, err := c.svcCtx.Dao.KV.HdelCtx(
		c.ctx,
		getUserKey(in.GetUserId()),
		strconv.FormatInt(in.GetAuthKeyId(), 10))
	if err != nil {
		c.Logger.Errorf("status.setSessionOffline(userId=%d, authKeyId=%d) error: %v", in.GetUserId(), in.GetAuthKeyId(), err)
		return nil, fmt.Errorf("status.setSessionOffline - hdel: %w", err)
	}

	return mtproto.BoolTrue, nil
}
