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
	"github.com/teamgram/teamgram-server/v2/app/bff/drafts/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesClearAllDrafts
// messages.clearAllDrafts#7e58ee9c = Bool;
func (c *DraftsCore) MessagesClearAllDrafts(in *tg.TLMessagesClearAllDrafts) (*tg.Bool, error) {
	if c.svcCtx == nil || c.svcCtx.Repo == nil || c.svcCtx.Repo.DialogClient == nil {
		return tg.BoolTrue, nil
	}

	rValues, err := c.svcCtx.Repo.DialogClient.DialogClearAllDrafts(c.ctx, &repository.DialogClearAll{
		UserId: c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("messages.clearAllDrafts: %v", err)
		return nil, err
	}

	if len(rValues.Datas) == 0 {
		return tg.BoolTrue, nil
	}

	// TODO: for each cleared draft, build syncUpdates with user/chat
	// resolution and call SyncUpdatesNotMe. PEER_CHANNEL case requires
	// plugin (enterprise feature).

	return tg.BoolTrue, nil
}
