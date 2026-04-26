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

package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/teamgram/teamgram-server/v2/app/service/status/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/status/status"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/kv"
)

// Repository is the dependency container for repository instances.
type Repository struct {
	kv kv.Store
}

// NewRepository creates a new Repository.
func NewRepository(c config.Config) *Repository {
	return &Repository{
		kv: kv.NewStore(c.KV),
	}
}

// Close releases repository-owned clients.
func (r *Repository) Close() error {
	if r == nil {
		return nil
	}
	return nil
}

func getUserKey(userID int64) string {
	return fmt.Sprintf("%s#%d", userKeyPrefix, userID)
}

// SetSessionOnline atomically sets the session entry in the user's online hash
// and refreshes the key-level TTL.
func (r *Repository) SetSessionOnline(ctx context.Context, userID int64, session *status.SessionEntry, expireSeconds int) error {
	userKey := getUserKey(userID)
	field := strconv.FormatInt(session.AuthKeyId, 10)

	// Bypass the custom MarshalJSON on TLSessionEntry (which wraps output
	// in {"_name":"...","_object":...}) so we store plain JSON.
	type plainSessionEntry status.TLSessionEntry
	sessData, err := json.Marshal((*plainSessionEntry)(session))
	if err != nil {
		return fmt.Errorf("marshal session entry: %w", err)
	}

	_, err = r.kv.EvalCtx(ctx, hsetAndExpireScript, userKey, field, string(sessData), strconv.Itoa(expireSeconds))
	if err != nil {
		return fmt.Errorf("set session online: %w", err)
	}
	return nil
}

// SetSessionOffline removes the session for the given auth key from the user's online hash.
// It is idempotent: deleting a non-existent field is a no-op.
func (r *Repository) SetSessionOffline(ctx context.Context, userID, authKeyID int64) error {
	userKey := getUserKey(userID)
	field := strconv.FormatInt(authKeyID, 10)

	_, err := r.kv.HdelCtx(ctx, userKey, field)
	if err != nil {
		return fmt.Errorf("set session offline: %w", err)
	}
	return nil
}

// GetUserOnlineSessions returns all online sessions for a user.
// Bad JSON entries are logged and skipped, matching master behavior.
func (r *Repository) GetUserOnlineSessions(ctx context.Context, userID int64) (*status.UserSessionEntryList, error) {
	userKey := getUserKey(userID)

	rMap, err := r.kv.HgetallCtx(ctx, userKey)
	if err != nil {
		return nil, fmt.Errorf("get user online sessions: %w", err)
	}

	rValues := &status.TLUserSessionEntryList{
		UserSessions: make([]*status.TLSessionEntry, 0, len(rMap)),
	}

	for field, rawValue := range rMap {
		sess := new(status.TLSessionEntry)
		if err := json.Unmarshal([]byte(rawValue), sess); err != nil {
			authKeyID, _ := strconv.ParseInt(field, 10, 64)
			preview := rawValue
			if len(preview) > 100 {
				preview = preview[:100]
			}
			logx.WithContext(ctx).Infof(
				"status: skip bad session JSON: user_id=%d auth_key_id=%d payload_preview=%s err=%v",
				userID, authKeyID, preview, err,
			)
			continue
		}
		rValues.UserSessions = append(rValues.UserSessions, sess)
	}

	return rValues, nil
}

// GetUsersOnlineSessionsList returns online sessions for multiple users.
// Uses sequential HGETALL calls since kv.Store does not expose pipeline.
// Users with no online sessions return empty lists.
func (r *Repository) GetUsersOnlineSessionsList(ctx context.Context, userIDs []int64) (*status.VectorUserSessionEntryList, error) {
	if len(userIDs) == 0 {
		return &status.VectorUserSessionEntryList{}, nil
	}

	result := &status.VectorUserSessionEntryList{
		Datas: make([]*status.TLUserSessionEntryList, 0, len(userIDs)),
	}

	for _, userID := range userIDs {
		rMap, err := r.kv.HgetallCtx(ctx, getUserKey(userID))
		if err != nil {
			return nil, fmt.Errorf("get users online sessions list for user %d: %w", userID, err)
		}

		entry := &status.TLUserSessionEntryList{
			UserId:       userID,
			UserSessions: make([]*status.TLSessionEntry, 0, len(rMap)),
		}

		for field, rawValue := range rMap {
			sess := new(status.TLSessionEntry)
			if err := json.Unmarshal([]byte(rawValue), sess); err != nil {
				authKeyID, _ := strconv.ParseInt(field, 10, 64)
				preview := rawValue
				if len(preview) > 100 {
					preview = preview[:100]
				}
				logx.WithContext(ctx).Infof(
					"status: skip bad session JSON: user_id=%d auth_key_id=%d payload_preview=%s err=%v",
					userID, authKeyID, preview, err,
				)
				continue
			}
			entry.UserSessions = append(entry.UserSessions, sess)
		}

		result.Datas = append(result.Datas, entry)
	}

	return result, nil
}
