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

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

type (
	extendUsernameModel interface {
		SelectListByUserIdList(ctx context.Context, idList []int64) ([]Username, error)
	}
)

func (m *customUsernameModel) SelectListByUserIdList(ctx context.Context, idList []int64) ([]Username, error) {
	if len(idList) == 0 {
		return []Username{}, nil
	}
	query := fmt.Sprintf("select peer_type, peer_id, username, editable, active, order2 from username where peer_type = 2 and peer_id in (%s)", sqlx.InInt64List(idList))
	var values []Username
	if err := m.db.QueryRowsPartial(ctx, &values, query); err != nil {
		return nil, fmt.Errorf("username.SelectListByUserIdList: %w", err)
	}
	return values, nil
}
