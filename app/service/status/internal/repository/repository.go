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

	"github.com/teamgram/marmota/pkg/stores/kv"
	"github.com/zeromicro/go-zero/core/logx"
)

// Repository is the dependency container for repository instances.
type Repository struct {
	kv kv.ExtStore
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

type userSessionBatchEntry struct {
	index  int
	key    string
	userID int64
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
		return wrapStorageError("marshal session entry", err)
	}

	_, err = r.kv.EvalCtx(ctx, hsetAndExpireScript, userKey, field, string(sessData), strconv.Itoa(expireSeconds))
	if err != nil {
		return wrapStorageError("set session online", err)
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
		return wrapStorageError("set session offline", err)
	}
	return nil
}

// GetUserOnlineSessions returns all online sessions for a user.
// Bad JSON entries are logged and skipped, matching master behavior.
func (r *Repository) GetUserOnlineSessions(ctx context.Context, userID int64) (*status.UserSessionEntryList, error) {
	userKey := getUserKey(userID)

	rMap, err := r.kv.HgetallCtx(ctx, userKey)
	if err != nil {
		return nil, wrapStorageError("get user online sessions", err)
	}

	return buildUserSessionEntryList(ctx, userID, rMap), nil
}

func buildUserSessionEntryList(ctx context.Context, userID int64, rMap map[string]string) *status.TLUserSessionEntryList {
	rValues := &status.TLUserSessionEntryList{
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
		rValues.UserSessions = append(rValues.UserSessions, sess)
	}

	return rValues
}

func assignUserSessionBatchResult(ctx context.Context, result *status.VectorUserSessionEntryList, entries []userSessionBatchEntry, maps []map[string]string) {
	for i, rMap := range maps {
		result.Datas[entries[i].index] = buildUserSessionEntryList(ctx, entries[i].userID, rMap)
	}
}

// GetUsersOnlineSessionsList returns online sessions for multiple users.
// Groups user keys by pipeline node and executes per-node batched HGETALL.
// Users with no online sessions return empty lists.
func (r *Repository) GetUsersOnlineSessionsList(ctx context.Context, userIDs []int64) (*status.VectorUserSessionEntryList, error) {
	if len(userIDs) == 0 {
		return &status.VectorUserSessionEntryList{}, nil
	}

	// Group keys by pipeline node (consistent-hash routing for Redis cluster).
	groups := make(map[kv.Pipeline][]userSessionBatchEntry)
	for idx, id := range userIDs {
		k := getUserKey(id)
		rawPipe, err := r.kv.GetPipeline(k)
		if err != nil {
			return nil, wrapStorageError(fmt.Sprintf("get pipeline for user %d", id), err)
		}
		pipe, ok := rawPipe.(kv.Pipeline)
		if !ok {
			return nil, wrapStorageError(fmt.Sprintf("unexpected pipeline type for user %d", id), fmt.Errorf("%T", rawPipe))
		}
		groups[pipe] = append(groups[pipe], userSessionBatchEntry{index: idx, key: k, userID: id})
	}

	result := &status.VectorUserSessionEntryList{
		Datas: make([]*status.TLUserSessionEntryList, len(userIDs)),
	}

	for pipe, entries := range groups {
		cmds := make([]*kv.MapStringStringCmd, len(entries))
		err := pipe.PipelinedCtx(ctx, func(pipeliner kv.Pipeliner) error {
			for i, e := range entries {
				cmds[i] = pipeliner.HGetAll(ctx, e.key)
			}
			return nil
		})
		if err != nil {
			return nil, wrapStorageError("pipeline execute", err)
		}

		maps := make([]map[string]string, len(cmds))
		for i, cmd := range cmds {
			rMap, cmdErr := cmd.Result()
			if cmdErr != nil {
				return nil, wrapStorageError(fmt.Sprintf("hgetall for user %d", entries[i].userID), cmdErr)
			}
			maps[i] = rMap
		}
		assignUserSessionBatchResult(ctx, result, entries, maps)
	}

	return result, nil
}
