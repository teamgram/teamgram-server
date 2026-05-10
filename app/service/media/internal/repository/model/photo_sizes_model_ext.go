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
	extendPhotoSizesModel interface {
		DeleteByPhotoSizeId(ctx context.Context, photoSizeId int64) error
	}
)

func (m *customPhotoSizesModel) DeleteByPhotoSizeId(ctx context.Context, photoSizeId int64) error {
	_, err := m.db.Exec(ctx, "delete from photo_sizes where photo_size_id = ?", photoSizeId)
	if err != nil {
		return fmt.Errorf("photo_sizes.DeleteByPhotoSizeId exec: %w", err)
	}
	return nil
}
