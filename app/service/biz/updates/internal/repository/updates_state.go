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

package repository

import (
	"context"
	"sync"
	"time"
)

// UpdatesStateDO holds the per-user updates state (pts, qts, date, seq).
type UpdatesStateDO struct {
	UserId      int64
	Pts         int32
	Qts         int32
	Date        int32
	Seq         int32
	UnreadCount int32
}

func defaultState(userId int64) *UpdatesStateDO {
	return &UpdatesStateDO{
		UserId: userId,
		Pts:    1,
		Qts:    0,
		Date:   int32(time.Now().Unix()),
		Seq:    0,
	}
}

// UpdatesStateModel manages per-user updates state.
type UpdatesStateModel interface {
	GetUserUpdatesState(ctx context.Context, userId int64) (*UpdatesStateDO, error)
	SetUserUpdatesState(ctx context.Context, state *UpdatesStateDO) error
	IncrementPts(ctx context.Context, userId int64) (int32, error)
}

type inMemoryUpdatesStateModel struct {
	mu     sync.RWMutex
	states map[int64]*UpdatesStateDO
}

func NewUpdatesStateModel() UpdatesStateModel {
	return &inMemoryUpdatesStateModel{
		states: make(map[int64]*UpdatesStateDO),
	}
}

func (m *inMemoryUpdatesStateModel) GetUserUpdatesState(ctx context.Context, userId int64) (*UpdatesStateDO, error) {
	m.mu.RLock()
	state, ok := m.states[userId]
	m.mu.RUnlock()
	if ok {
		return state, nil
	}
	// Return a fresh default state for unknown users.
	return defaultState(userId), nil
}

func (m *inMemoryUpdatesStateModel) SetUserUpdatesState(ctx context.Context, state *UpdatesStateDO) error {
	m.mu.Lock()
	m.states[state.UserId] = state
	m.mu.Unlock()
	return nil
}

func (m *inMemoryUpdatesStateModel) IncrementPts(ctx context.Context, userId int64) (int32, error) {
	m.mu.Lock()
	state, ok := m.states[userId]
	if !ok {
		state = defaultState(userId)
	}
	state.Pts++
	state.Date = int32(time.Now().Unix())
	m.states[userId] = state
	m.mu.Unlock()
	return state.Pts, nil
}
