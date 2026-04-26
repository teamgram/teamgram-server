// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
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

	"github.com/teamgram/teamgram-server/v2/app/service/status/status"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// StatusSetSessionOnline
// status.setSessionOnline user_id:long session:SessionEntry = Bool;
func (c *StatusCore) StatusSetSessionOnline(in *status.TLStatusSetSessionOnline) (*tg.Bool, error) {
	if in.UserId <= 0 {
		return nil, fmt.Errorf("%w: setSessionOnline: invalid user_id %d", status.ErrStatusInvalidArgument, in.UserId)
	}
	if in.Session == nil {
		return nil, fmt.Errorf("%w: setSessionOnline: session is nil", status.ErrStatusInvalidArgument)
	}
	if in.Session.AuthKeyId == 0 {
		return nil, fmt.Errorf("%w: setSessionOnline: invalid auth_key_id", status.ErrStatusInvalidArgument)
	}

	err := c.svcCtx.Repo.SetSessionOnline(c.ctx, in.UserId, in.Session, c.svcCtx.Config.StatusExpire)
	if err != nil {
		c.Logger.Errorf("status.setSessionOnline - error: %v", err)
		return nil, err
	}

	return tg.MakeTLBoolTrue(&tg.TLBoolTrue{}).ToBool(), nil
}
