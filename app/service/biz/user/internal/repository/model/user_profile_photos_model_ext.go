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
	"errors"
	"fmt"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

type (
	extendUserProfilePhotosModel interface {
		SelectNextTx(tx *sqlx.Tx, userID int64, idList []int64) (int64, error)
	}
)

func (m *defaultUserProfilePhotosModel) SelectNextTx(tx *sqlx.Tx, userID int64, idList []int64) (int64, error) {
	if len(idList) == 0 {
		return 0, nil
	}

	var photoID int64
	query := fmt.Sprintf("select photo_id from user_profile_photos where user_id = ? and photo_id not in (%s) and deleted = 0 order by date2 desc limit 1", sqlx.InInt64List(idList))
	err := tx.QueryRowPartial(&photoID, query, userID)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return 0, nil
		}
		return 0, fmt.Errorf("user_profile_photos.SelectNextTx: %w", err)
	}
	return photoID, nil
}
