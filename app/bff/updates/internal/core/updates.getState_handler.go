// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
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
	"github.com/teamgram/teamgram-server/v2/app/service/biz/updates/updates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UpdatesGetState
// updates.getState#edd4882a = updates.State;
func (c *UpdatesCore) UpdatesGetState(in *tg.TLUpdatesGetState) (*tg.UpdatesState, error) {
	if c.svcCtx != nil && c.svcCtx.UpdatesClient != nil {
		var authKeyId, userId int64
		if c.MD != nil {
			authKeyId = c.MD.AuthId
			userId = c.MD.UserId
		}

		state, err := c.svcCtx.UpdatesClient.UpdatesGetStateV2(c.ctx, &updates.TLUpdatesGetStateV2{
			AuthKeyId: authKeyId,
			UserId:    userId,
		})
		if err != nil {
			c.Logger.Errorf("updates.getState - UpdatesGetStateV2 error: %v", err)
			return nil, err
		}
		return state, nil
	}

	return makePlaceholderUpdatesState(1, 10), nil
}
