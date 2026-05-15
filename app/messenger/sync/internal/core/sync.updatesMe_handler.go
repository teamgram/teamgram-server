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

	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// SyncUpdatesMe
// sync.updatesMe flags:# user_id:long perm_auth_key_id:long auth_key_id:flags.0?long session_id:flags.1?long updates:Updates = Void;
func (c *SyncCore) SyncUpdatesMe(in *sync.TLSyncUpdatesMe) (*sync.Void, error) {
	const method = "sync.updatesMe"
	if in == nil {
		return nil, sync.ErrSyncInvalidArgument
	}
	if err := c.requireCaller(method); err != nil {
		return nil, err
	}
	if err := validateUserID(method, in.UserId); err != nil {
		return nil, err
	}
	if err := validateNonZeroID(method, "perm_auth_key_id", in.PermAuthKeyId); err != nil {
		return nil, err
	}
	if (in.AuthKeyId == nil) != (in.SessionId == nil) {
		return nil, fmt.Errorf("%w: %s auth_key_id and session_id must be both set or both absent", sync.ErrSyncInvalidArgument, method)
	}
	if err := validateUpdates(method, in.Updates); err != nil {
		return nil, err
	}
	var authKeyID, sessionID int64
	precise := in.AuthKeyId != nil
	if precise {
		authKeyID = *in.AuthKeyId
		sessionID = *in.SessionId
		if err := validateNonZeroID(method, "auth_key_id", authKeyID); err != nil {
			return nil, err
		}
		if err := validateNonZeroID(method, "session_id", sessionID); err != nil {
			return nil, err
		}
	}
	if err := c.svcCtx.Repo.UpdatesMe(c.ctx, in.UserId, in.PermAuthKeyId, authKeyID, sessionID, precise, in.Updates); err != nil {
		return nil, err
	}
	return tg.VoidValue, nil
}
