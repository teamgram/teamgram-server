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

	"github.com/teamgram/teamgram-server/v2/app/service/presence/presence"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// PresenceSetSessionOffline
// presence.setSessionOffline user_id:long auth_key_id:long session_id:long = Bool;
func (c *PresenceCore) PresenceSetSessionOffline(in *presence.TLPresenceSetSessionOffline) (*presence.Bool, error) {
	const method = "presence.setSessionOffline"
	if err := c.requireCaller(method, c.svcCtx.Config.GatewayCallers); err != nil {
		return nil, err
	}
	if in == nil {
		return nil, fmt.Errorf("%w: %s request is nil", presence.ErrPresenceInvalidArgument, method)
	}
	if in.UserId <= 0 {
		return nil, fmt.Errorf("%w: %s invalid user_id %d", presence.ErrPresenceInvalidArgument, method, in.UserId)
	}
	if in.AuthKeyId <= 0 {
		return nil, fmt.Errorf("%w: %s invalid auth_key_id %d", presence.ErrPresenceInvalidArgument, method, in.AuthKeyId)
	}
	if in.SessionId == 0 {
		return nil, fmt.Errorf("%w: %s invalid session_id", presence.ErrPresenceInvalidArgument, method)
	}
	if err := c.svcCtx.Repo.SetSessionOffline(c.ctx, in.UserId, in.AuthKeyId, in.SessionId); err != nil {
		return nil, err
	}
	return tg.BoolTrue, nil
}
