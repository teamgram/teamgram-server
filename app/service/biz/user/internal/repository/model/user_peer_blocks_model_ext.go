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
	extendUserPeerBlocksModel interface {
		SelectListOffset(ctx context.Context, userID int64, offset, limit int32) ([]UserPeerBlocks, error)
	}
)

func (m *defaultUserPeerBlocksModel) SelectListOffset(ctx context.Context, userID int64, offset, limit int32) ([]UserPeerBlocks, error) {
	if limit <= 0 {
		return []UserPeerBlocks{}, nil
	}
	if offset < 0 {
		offset = 0
	}

	var values []UserPeerBlocks
	err := m.db.QueryRowsPartial(ctx, &values,
		"select user_id, peer_type, peer_id, `date` from user_peer_blocks where user_id = ? and deleted = 0 order by id asc limit ? offset ?",
		userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("user_peer_blocks.SelectListOffset: %w", err)
	}
	return values, nil
}
