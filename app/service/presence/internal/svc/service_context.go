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

package svc

import (
	"context"
	"sync"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/presence/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/presence/internal/repository"
	presencepb "github.com/teamgram/teamgram-server/v2/app/service/presence/presence"
)

type Repository interface {
	SetSessionOnline(ctx context.Context, session *presencepb.OnlineSession, now int64, expireSeconds int, hashTTLSeconds int, cleanupIntervalSeconds int) error
	SetSessionOffline(ctx context.Context, userID, authKeyID, sessionID int64) error
	GetUserOnlineSessions(ctx context.Context, userID int64, now int64) (*presencepb.TLUserOnlineSessions, error)
	GetUsersOnlineSessions(ctx context.Context, userIDs []int64, now int64) (*presencepb.VectorUserOnlineSessions, error)
	CleanupExpiredForUser(ctx context.Context, userID int64, now int64) error
}

type ServiceContext struct {
	Config  config.Config
	Repo    Repository
	Limiter *CallerLimiter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Repo:    repository.NewRepository(c),
		Limiter: NewCallerLimiter(),
	}
}

func (s *ServiceContext) Close() error {
	if s == nil || s.Repo == nil {
		return nil
	}
	if closer, ok := s.Repo.(interface{ Close() error }); ok {
		return closer.Close()
	}
	return nil
}

func (s *ServiceContext) AllowPresenceCall(method, caller string, limit int) bool {
	if limit <= 0 {
		return true
	}
	if s.Limiter == nil {
		s.Limiter = NewCallerLimiter()
	}
	return s.Limiter.Allow(method, caller, limit, time.Now())
}

type CallerLimiter struct {
	mu      sync.Mutex
	windows map[callerLimitKey]callerLimitWindow
}

type callerLimitKey struct {
	method string
	caller string
}

type callerLimitWindow struct {
	second int64
	count  int
}

func NewCallerLimiter() *CallerLimiter {
	return &CallerLimiter{
		windows: make(map[callerLimitKey]callerLimitWindow),
	}
}

func (l *CallerLimiter) Allow(method, caller string, limit int, now time.Time) bool {
	if limit <= 0 {
		return true
	}
	l.mu.Lock()
	defer l.mu.Unlock()

	key := callerLimitKey{method: method, caller: caller}
	second := now.Unix()
	window := l.windows[key]
	if window.second != second {
		window = callerLimitWindow{second: second}
	}
	if window.count >= limit {
		l.windows[key] = window
		return false
	}
	window.count++
	l.windows[key] = window
	return true
}
