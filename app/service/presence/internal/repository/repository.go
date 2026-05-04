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
	"errors"
	"fmt"
	"strconv"

	"github.com/teamgram/teamgram-server/v2/app/service/presence/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/presence/internal/metrics"
	presencepb "github.com/teamgram/teamgram-server/v2/app/service/presence/presence"

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

type sessionCacheData struct {
	UserId            int64  `json:"user_id"`
	PermAuthKeyId     int64  `json:"perm_auth_key_id"`
	AuthKeyId         int64  `json:"auth_key_id"`
	AuthKeyType       int32  `json:"auth_key_type"`
	SessionId         int64  `json:"session_id"`
	GatewayId         string `json:"gateway_id"`
	GatewayGeneration string `json:"gateway_generation"`
	GatewayRpcAddr    string `json:"gateway_rpc_addr"`
	Layer             int32  `json:"layer"`
	Client            string `json:"client"`
	UpdatedAt         int64  `json:"updated_at"`
	ExpiresAt         int64  `json:"expires_at"`
}

type userOnlineBatchEntry struct {
	index  int
	key    string
	userID int64
}

type cleanupField struct {
	field      string
	errorClass string
}

var (
	errNilRepository = errors.New("repository is nil")
	errNilKV         = errors.New("repository kv is nil")
)

func sessionCacheDataFromTL(session *presencepb.OnlineSession, now int64, expireSeconds int) *sessionCacheData {
	if session == nil {
		return nil
	}
	return &sessionCacheData{
		UserId:            session.UserId,
		PermAuthKeyId:     session.PermAuthKeyId,
		AuthKeyId:         session.AuthKeyId,
		AuthKeyType:       session.AuthKeyType,
		SessionId:         session.SessionId,
		GatewayId:         session.GatewayId,
		GatewayGeneration: session.GatewayGeneration,
		GatewayRpcAddr:    session.GatewayRpcAddr,
		Layer:             session.Layer,
		Client:            session.Client,
		UpdatedAt:         now,
		ExpiresAt:         now + int64(expireSeconds),
	}
}

func (s *sessionCacheData) toTL() *presencepb.TLOnlineSession {
	if s == nil {
		return nil
	}
	return presencepb.MakeTLOnlineSession(&presencepb.TLOnlineSession{
		UserId:            s.UserId,
		PermAuthKeyId:     s.PermAuthKeyId,
		AuthKeyId:         s.AuthKeyId,
		AuthKeyType:       s.AuthKeyType,
		SessionId:         s.SessionId,
		GatewayId:         s.GatewayId,
		GatewayGeneration: s.GatewayGeneration,
		GatewayRpcAddr:    s.GatewayRpcAddr,
		Layer:             s.Layer,
		Client:            s.Client,
		UpdatedAt:         s.UpdatedAt,
		ExpiresAt:         s.ExpiresAt,
	})
}

func buildOnlineSessions(ctx context.Context, userID int64, raw map[string]string, now int64) ([]*presencepb.TLOnlineSession, []string) {
	sessions, cleanup := buildOnlineSessionsWithCleanup(ctx, userID, raw, now)
	fields := make([]string, 0, len(cleanup))
	for _, entry := range cleanup {
		fields = append(fields, entry.field)
	}
	return sessions, fields
}

func buildOnlineSessionsWithCleanup(ctx context.Context, userID int64, raw map[string]string, now int64) ([]*presencepb.TLOnlineSession, []cleanupField) {
	sessions := make([]*presencepb.TLOnlineSession, 0, len(raw))
	cleanup := make([]cleanupField, 0)
	for field, value := range raw {
		var data sessionCacheData
		if err := json.Unmarshal([]byte(value), &data); err != nil {
			logx.WithContext(ctx).Errorf("presence: corrupt session entry: user_id=%d field=%s err=%v", userID, field, err)
			metrics.CorruptEntry("decode")
			cleanup = append(cleanup, cleanupField{field: field, errorClass: "decode"})
			continue
		}
		if data.UserId != userID || data.AuthKeyId <= 0 || data.SessionId == 0 || data.PermAuthKeyId <= 0 || data.GatewayId == "" || data.GatewayGeneration == "" || data.GatewayRpcAddr == "" || field != sessionField(data.AuthKeyId, data.SessionId) {
			logx.WithContext(ctx).Errorf("presence: inconsistent session entry: user_id=%d field=%s", userID, field)
			metrics.CorruptEntry("inconsistent")
			cleanup = append(cleanup, cleanupField{field: field, errorClass: "inconsistent"})
			continue
		}
		if data.ExpiresAt <= now {
			cleanup = append(cleanup, cleanupField{field: field, errorClass: ""})
			continue
		}
		sessions = append(sessions, data.toTL())
	}
	return sessions, cleanup
}

func (r *Repository) SetSessionOnline(ctx context.Context, session *presencepb.OnlineSession, now int64, expireSeconds int, hashTTLSeconds int, cleanupIntervalSeconds int) error {
	if err := validateRepositoryOnlineSession("set session online", session); err != nil {
		return err
	}
	if err := r.ensureStore("set session online"); err != nil {
		return err
	}
	data, err := json.Marshal(sessionCacheDataFromTL(session, now, expireSeconds))
	if err != nil {
		return wrapStorageError("marshal online session", err)
	}
	_, err = r.kv.EvalCtx(
		ctx,
		hsetAndExpireScript,
		userKey(session.UserId),
		sessionField(session.AuthKeyId, session.SessionId),
		string(data),
		strconv.Itoa(hashTTLSeconds),
	)
	if err != nil {
		return wrapStorageError("set session online", err)
	}
	r.cleanupOnWrite(ctx, session.UserId, now, cleanupIntervalSeconds)
	return nil
}

func (r *Repository) SetSessionOffline(ctx context.Context, userID, authKeyID, sessionID int64) error {
	if err := r.ensureStore("set session offline"); err != nil {
		return err
	}
	_, err := r.kv.HdelCtx(ctx, userKey(userID), sessionField(authKeyID, sessionID))
	if err != nil {
		return wrapStorageError("set session offline", err)
	}
	return nil
}

func (r *Repository) GetUserOnlineSessions(ctx context.Context, userID int64, now int64) (*presencepb.TLUserOnlineSessions, error) {
	if err := r.ensureStore("get user online sessions"); err != nil {
		return nil, err
	}
	raw, err := r.kv.HgetallCtx(ctx, userKey(userID))
	if err != nil {
		return nil, wrapStorageError("get user online sessions", err)
	}
	sessions, cleanup := buildOnlineSessionsWithCleanup(ctx, userID, raw, now)
	r.cleanupFields(ctx, userID, cleanup)
	return presencepb.MakeTLUserOnlineSessions(&presencepb.TLUserOnlineSessions{
		UserId:   userID,
		Sessions: sessions,
	}), nil
}

func (r *Repository) GetUsersOnlineSessions(ctx context.Context, userIDs []int64, now int64) (*presencepb.VectorUserOnlineSessions, error) {
	if err := r.ensureStore("get users online sessions"); err != nil {
		return nil, err
	}
	if len(userIDs) == 0 {
		return &presencepb.VectorUserOnlineSessions{}, nil
	}

	groups := make(map[kv.Pipeline][]userOnlineBatchEntry)
	for idx, userID := range userIDs {
		key := userKey(userID)
		pipe, err := r.kv.GetPipeline(key)
		if err != nil {
			return nil, wrapStorageError(fmt.Sprintf("get pipeline for user %d", userID), err)
		}
		groups[pipe] = append(groups[pipe], userOnlineBatchEntry{index: idx, key: key, userID: userID})
	}

	result := &presencepb.VectorUserOnlineSessions{
		Datas: make([]*presencepb.TLUserOnlineSessions, len(userIDs)),
	}

	for pipe, entries := range groups {
		cmds := make([]*kv.MapStringStringCmd, len(entries))
		err := pipe.PipelinedCtx(ctx, func(pipeliner kv.Pipeliner) error {
			for i, entry := range entries {
				cmds[i] = pipeliner.HGetAll(ctx, entry.key)
			}
			return nil
		})
		if err != nil {
			return nil, wrapStorageError("pipeline execute", err)
		}
		for i, cmd := range cmds {
			raw, cmdErr := cmd.Result()
			if cmdErr != nil {
				return nil, wrapStorageError(fmt.Sprintf("hgetall for user %d", entries[i].userID), cmdErr)
			}
			sessions, cleanup := buildOnlineSessionsWithCleanup(ctx, entries[i].userID, raw, now)
			r.cleanupFields(ctx, entries[i].userID, cleanup)
			result.Datas[entries[i].index] = presencepb.MakeTLUserOnlineSessions(&presencepb.TLUserOnlineSessions{
				UserId:   entries[i].userID,
				Sessions: sessions,
			})
		}
	}

	return result, nil
}

func (r *Repository) CleanupExpiredForUser(ctx context.Context, userID int64, now int64) error {
	if err := r.ensureStore("cleanup expired for user"); err != nil {
		return err
	}
	raw, err := r.kv.HgetallCtx(ctx, userKey(userID))
	if err != nil {
		return wrapStorageError("cleanup expired for user", err)
	}
	_, cleanup := buildOnlineSessionsWithCleanup(ctx, userID, raw, now)
	r.cleanupFields(ctx, userID, cleanup)
	return nil
}

func (r *Repository) cleanupFields(ctx context.Context, userID int64, cleanup []cleanupField) {
	for _, entry := range cleanup {
		if _, err := r.kv.HdelCtx(ctx, userKey(userID), entry.field); err != nil && entry.errorClass != "" {
			metrics.CorruptEntryCleanupFailure(entry.errorClass)
			logx.WithContext(ctx).Errorf("presence: cleanup corrupt session entry failed: user_id=%d field=%s error_class=%s err=%v", userID, entry.field, entry.errorClass, err)
		}
	}
}

func (r *Repository) ensureStore(op string) error {
	if r == nil {
		return wrapStorageError(op, errNilRepository)
	}
	if r.kv == nil {
		return wrapStorageError(op, errNilKV)
	}
	return nil
}

func validateRepositoryOnlineSession(op string, session *presencepb.OnlineSession) error {
	if session == nil {
		return fmt.Errorf("%w: %s: session is nil", presencepb.ErrPresenceInvalidArgument, op)
	}
	if session.UserId <= 0 || session.PermAuthKeyId <= 0 || session.AuthKeyId <= 0 || session.SessionId == 0 || session.GatewayId == "" || session.GatewayGeneration == "" || session.GatewayRpcAddr == "" {
		return fmt.Errorf("%w: %s: invalid online session", presencepb.ErrPresenceInvalidArgument, op)
	}
	return nil
}

func shouldCleanupOnWrite(cleanupIntervalSeconds int) bool {
	return cleanupIntervalSeconds > 0
}

func (r *Repository) cleanupOnWrite(ctx context.Context, userID int64, now int64, cleanupIntervalSeconds int) {
	if !shouldCleanupOnWrite(cleanupIntervalSeconds) {
		return
	}
	claimed, err := r.kv.SetnxExCtx(ctx, cleanupKey(userID), strconv.FormatInt(now, 10), cleanupIntervalSeconds)
	if err != nil {
		metrics.CorruptEntryCleanupFailure("cleanup")
		logx.WithContext(ctx).Errorf("presence: cleanup interval claim failed: user_id=%d err=%v", userID, err)
		return
	}
	if !claimed {
		return
	}
	if err := r.CleanupExpiredForUser(ctx, userID, now); err != nil {
		metrics.CorruptEntryCleanupFailure("cleanup")
		logx.WithContext(ctx).Errorf("presence: cleanup on write failed: user_id=%d err=%v", userID, err)
	}
}
