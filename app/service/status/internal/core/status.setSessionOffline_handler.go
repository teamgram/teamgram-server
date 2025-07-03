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
)

var _ *tg.Bool

// StatusSetSessionOffline
// status.setSessionOffline user_id:long auth_key_id:long = Bool;
func (c *StatusCore) StatusSetSessionOffline(in *status.TLStatusSetSessionOffline) (*tg.Bool, error) {
	_, err := c.svcCtx.Dao.KV.HdelCtx(
		c.ctx,
		getUserKey(in.UserId),
		strconv.FormatInt(in.AuthKeyId, 10))
	if err != nil {
		c.Logger.Errorf("status.setSessionOffline(%s) error(%v)", in, err)
		return nil, err
	}

	return tg.BoolTrue, nil
}
