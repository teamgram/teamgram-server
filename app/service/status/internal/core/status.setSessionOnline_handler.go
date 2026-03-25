// Copyright 2024 Teamgooo Authors
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
	"fmt"
	"strconv"

	"github.com/teamgram/teamgram-server/v2/app/service/status/status"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/zeromicro/go-zero/core/jsonx"
)

var _ *tg.Bool

// StatusSetSessionOnline
// status.setSessionOnline user_id:long session:SessionEntry = Bool;
func (c *StatusCore) StatusSetSessionOnline(in *status.TLStatusSetSessionOnline) (*tg.Bool, error) {
	var (
		userK   = getUserKey(in.UserId)
		sess, _ = in.Session.(*status.TLSessionEntry)
	)

	if in.UserId <= 0 || sess == nil || sess.AuthKeyId == 0 {
		return nil, fmt.Errorf("status.setSessionOnline - invalid params: userId=%d, session=%v", in.UserId, sess)
	}

	sessData, err := jsonx.Marshal(sess)
	if err != nil {
		c.Logger.Errorf("status.setSessionOnline - marshal session error: %v", err)
		return nil, fmt.Errorf("status.setSessionOnline - marshal session: %w", err)
	}

	_, err = c.svcCtx.Dao.KV.EvalCtx(
		c.ctx,
		hsetAndExpireScript,
		userK,
		strconv.FormatInt(sess.AuthKeyId, 10),
		string(sessData),
		c.svcCtx.Config.StatusExpire)
	if err != nil {
		c.Logger.Errorf("status.setSessionOnline - eval(userId=%d) error: %v", in.UserId, err)
		return nil, fmt.Errorf("status.setSessionOnline - eval: %w", err)
	}

	return tg.BoolTrue, nil
}
