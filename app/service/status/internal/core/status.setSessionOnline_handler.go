// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"strconv"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/status/status"

	"github.com/zeromicro/go-zero/core/jsonx"
)

var _ *tg.Bool

// StatusSetSessionOnline
// status.setSessionOnline user_id:long session:SessionEntry = Bool;
func (c *StatusCore) StatusSetSessionOnline(in *status.TLStatusSetSessionOnline) (*tg.Bool, error) {
	var (
		userK   = getUserKey(in.UserId)
		sess, _ = in.Session.ToSessionEntry()
	)

	sessData, _ := jsonx.Marshal(sess)
	err := c.svcCtx.Dao.KV.HsetCtx(
		c.ctx,
		userK,
		strconv.FormatInt(sess.AuthKeyId, 10),
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

	return tg.BoolTrue, nil
}
