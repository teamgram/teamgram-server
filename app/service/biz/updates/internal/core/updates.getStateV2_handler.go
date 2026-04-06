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

var _ *tg.Bool

// UpdatesGetStateV2
// updates.getStateV2 auth_key_id:long user_id:long = updates.State;
func (c *UpdatesCore) UpdatesGetStateV2(in *updates.TLUpdatesGetStateV2) (*tg.UpdatesState, error) {
	state, err := c.getUserUpdatesState(in.UserId)
	if err != nil {
		return nil, err
	}
	return state.ToUpdatesState(), nil
}

func (c *UpdatesCore) getUserUpdatesState(userId int64) (*tg.TLUpdatesState, error) {
	if c.svcCtx != nil && c.svcCtx.Repository != nil {
		state, err := c.svcCtx.Repository.UpdatesState.GetUserUpdatesState(c.ctx, userId)
		if err == nil && state != nil {
			return &tg.TLUpdatesState{
				Pts:         state.Pts,
				Qts:         state.Qts,
				Date:        state.Date,
				Seq:         state.Seq,
				UnreadCount: state.UnreadCount,
			}, nil
		}
	}
	// Fallback: return default placeholder state.
	return &tg.TLUpdatesState{
		Pts:         1,
		Qts:         0,
		Date:        10,
		Seq:         0,
		UnreadCount: 0,
	}, nil
}

// getUserUpdatesStateWithDefaults returns the stored state if available,
// otherwise returns a state built from the provided defaults.
func (c *UpdatesCore) getUserUpdatesStateWithDefaults(userId int64, defaultPts int32, defaultDate int32) (*tg.TLUpdatesState, error) {
	if c.svcCtx != nil && c.svcCtx.Repository != nil {
		state, err := c.svcCtx.Repository.UpdatesState.GetUserUpdatesState(c.ctx, userId)
		if err == nil && state != nil {
			return &tg.TLUpdatesState{
				Pts:         state.Pts,
				Qts:         state.Qts,
				Date:        state.Date,
				Seq:         state.Seq,
				UnreadCount: state.UnreadCount,
			}, nil
		}
	}
	// Fallback: use the provided defaults.
	return &tg.TLUpdatesState{
		Pts:         defaultPts,
		Qts:         0,
		Date:        defaultDate,
		Seq:         0,
		UnreadCount: 0,
	}, nil
}

// getUserUpdatesStateForDifference returns stored state if available,
// otherwise returns a fallback based on normalized client values.
func (c *UpdatesCore) getUserUpdatesStateForDifference(userId int64, clientPts int32, clientDate int32) (*tg.TLUpdatesState, error) {
	if c.svcCtx != nil && c.svcCtx.Repository != nil {
		state, err := c.svcCtx.Repository.UpdatesState.GetUserUpdatesState(c.ctx, userId)
		if err == nil && state != nil {
			return &tg.TLUpdatesState{
				Pts:         state.Pts,
				Qts:         state.Qts,
				Date:        state.Date,
				Seq:         state.Seq,
				UnreadCount: state.UnreadCount,
			}, nil
		}
	}
	// Fallback: use normalized client values (matches old behavior).
	if clientPts <= 0 {
		clientPts = 1
	}
	if clientDate <= 0 {
		clientDate = 10
	}
	return &tg.TLUpdatesState{
		Pts:         clientPts,
		Qts:         0,
		Date:        clientDate,
		Seq:         0,
		UnreadCount: 0,
	}, nil
}
