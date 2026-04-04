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
	"github.com/teamgram/teamgram-server/v2/app/interface/session/session"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/app/service/status/status"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ *tg.Bool

// SyncUpdatesNotMe
// sync.updatesNotMe user_id:long perm_auth_key_id:long updates:Updates = Void;
func (c *SyncCore) SyncUpdatesNotMe(in *sync.TLSyncUpdatesNotMe) (*tg.Void, error) {
	if c.svcCtx == nil || c.svcCtx.SessionClient == nil {
		return tg.MakeTLVoid(&tg.TLVoid{}).ToVoid(), nil
	}

	// When StatusClient is available, fan out to all online sessions except
	// the one identified by the sender's PermAuthKeyId — matching v1 behavior.
	if c.svcCtx.StatusClient != nil {
		sessionList, err := c.svcCtx.StatusClient.StatusGetUserOnlineSessions(c.ctx, &status.TLStatusGetUserOnlineSessions{
			UserId: in.UserId,
		})
		if err != nil {
			c.Logger.Errorf("sync.updatesNotMe - StatusGetUserOnlineSessions(%d) error: %v", in.UserId, err)
		} else {
			for _, sess := range sessionList.UserSessions {
				if sess.AuthKeyId == in.PermAuthKeyId {
					continue
				}
				_, pushErr := c.svcCtx.SessionClient.SessionPushUpdatesData(c.ctx, &session.TLSessionPushUpdatesData{
					PermAuthKeyId: sess.PermAuthKeyId,
					Updates:       in.Updates,
				})
				if pushErr != nil {
					c.Logger.Errorf("sync.updatesNotMe - push to session (permAuthKeyId=%d) error: %v", sess.PermAuthKeyId, pushErr)
				}
			}
			return tg.MakeTLVoid(&tg.TLVoid{}).ToVoid(), nil
		}
	}

	// Fallback: push directly using the sender's PermAuthKeyId when
	// StatusClient is not available.
	_, err := c.svcCtx.SessionClient.SessionPushUpdatesData(c.ctx, &session.TLSessionPushUpdatesData{
		PermAuthKeyId: in.PermAuthKeyId,
		Updates:       in.Updates,
	})
	if err != nil {
		return nil, err
	}

	return tg.MakeTLVoid(&tg.TLVoid{}).ToVoid(), nil
}
