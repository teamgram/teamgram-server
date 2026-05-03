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

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UpdatesGetState
// updates.getState#edd4882a = updates.State;
func (c *UpdatesCore) UpdatesGetState(in *tg.TLUpdatesGetState) (*tg.UpdatesState, error) {
	if c.MD == nil || c.MD.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	client := c.svcCtx.Repo.UserupdatesClient
	if client == nil {
		return nil, fmt.Errorf("updates.getState: userupdates client is nil")
	}
	state, err := client.UserupdatesGetState(c.ctx, &userupdates.TLUserupdatesGetState{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.PermAuthKeyId,
	})
	if err != nil {
		return nil, err
	}

	return userStateToUpdatesState(state).ToUpdatesState(), nil
}
