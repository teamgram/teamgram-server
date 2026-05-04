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
)

// PresenceGetGatewaySessions
// presence.getGatewaySessions gateway_id:string = Vector<OnlineSession>;
func (c *PresenceCore) PresenceGetGatewaySessions(in *presence.TLPresenceGetGatewaySessions) (*presence.VectorOnlineSession, error) {
	const method = "presence.getGatewaySessions"
	caller, err := c.authorizedCaller(method, allowedAdminDebugCallers(c.svcCtx.Config.AdminCallers, c.svcCtx.Config.DebugCallers))
	if err != nil {
		return nil, err
	}
	if err := c.requireQuota(method, caller, c.svcCtx.Config.PresenceGatewayDiagnosticsQPSPerCaller); err != nil {
		return nil, err
	}
	if in == nil {
		return nil, fmt.Errorf("%w: %s request is nil", presence.ErrPresenceInvalidArgument, method)
	}
	if in.GatewayId == "" {
		return nil, fmt.Errorf("%w: %s gateway_id is empty", presence.ErrPresenceInvalidArgument, method)
	}
	return &presence.VectorOnlineSession{}, nil
}
