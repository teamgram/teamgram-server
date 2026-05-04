// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"fmt"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/presence/presence"
)

// PresenceGetUserOnlineSessions
// presence.getUserOnlineSessions user_id:long = UserOnlineSessions;
func (c *PresenceCore) PresenceGetUserOnlineSessions(in *presence.TLPresenceGetUserOnlineSessions) (*presence.UserOnlineSessions, error) {
	const method = "presence.getUserOnlineSessions"
	caller, err := c.authorizedCaller(method, allowedQueryCallers(c.svcCtx.Config.SyncCallers, c.svcCtx.Config.AdminCallers, c.svcCtx.Config.DebugCallers))
	if err != nil {
		return nil, err
	}
	if err := c.requireQuota(method, caller, c.svcCtx.Config.PresenceQueryDefaultQPSPerCaller); err != nil {
		return nil, err
	}
	if in == nil {
		return nil, fmt.Errorf("%w: %s request is nil", presence.ErrPresenceInvalidArgument, method)
	}
	if in.UserId <= 0 {
		return nil, fmt.Errorf("%w: %s invalid user_id %d", presence.ErrPresenceInvalidArgument, method, in.UserId)
	}
	return c.svcCtx.Repo.GetUserOnlineSessions(c.ctx, in.UserId, time.Now().Unix())
}
