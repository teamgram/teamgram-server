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

package model

import (
	"context"
	"fmt"
)

type (
	extendChatInviteParticipantsModel interface {
		SelectCountByLink(ctx context.Context, link string, requested int32) (int32, error)
	}
)

func (m *defaultChatInviteParticipantsModel) SelectCountByLink(ctx context.Context, link string, requested int32) (int32, error) {
	query := "select count(*) from chat_invite_participants where link = ? and requested = ?"
	var count int32
	if err := m.db.QueryRow(ctx, &count, query, link, requested); err != nil {
		return 0, fmt.Errorf("chat_invite_participants.SelectCountByLink: %w", err)
	}
	return count, nil
}
