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
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// SyncPushUpdatesIfNot
// sync.pushUpdatesIfNot flags:# user_id:long includes:flags.0?Vector<long> excludes:flags.1?Vector<long> updates:Updates = Void;
func (c *SyncCore) SyncPushUpdatesIfNot(in *sync.TLSyncPushUpdatesIfNot) (*sync.Void, error) {
	const method = "sync.pushUpdatesIfNot"
	if in == nil {
		return nil, sync.ErrSyncInvalidArgument
	}
	if err := c.requireCaller(method); err != nil {
		return nil, err
	}
	if err := validateUserID(method, in.UserId); err != nil {
		return nil, err
	}
	if err := validatePushUpdatesIfNot(method, in.Includes, in.Excludes); err != nil {
		return nil, err
	}
	if err := validateUpdates(method, in.Updates); err != nil {
		return nil, err
	}
	if err := c.svcCtx.Repo.PushUpdatesIfNot(c.ctx, in.UserId, in.Includes != nil, in.Includes, in.Excludes != nil, in.Excludes, in.Updates); err != nil {
		return nil, err
	}
	return tg.VoidValue, nil
}
