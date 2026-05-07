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
	extendBotsModel interface {
		SelectByBotIdList(ctx context.Context, idList []int64) ([]Bots, error)
	}
)

func (m *customBotsModel) SelectByBotIdList(ctx context.Context, idList []int64) ([]Bots, error) {
	if len(idList) == 0 {
		return []Bots{}, nil
	}
	query := fmt.Sprintf("select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder, attach_menu_enabled, bot_attach_menu, bot_business, bot_has_main_app, bot_active_users, has_menu_button, menu_button_text, menu_button_url, bot_can_edit, has_preview_medias, description_photo_id, description_document_id, main_app_url, has_app_settings, placeholder_path, background_color, background_dark_color, header_color, header_dark_color, privacy_policy_url from bots where bot_id in (%s)", sqlx.InInt64List(idList))
	var values []Bots
	if err := m.db.QueryRowsPartial(ctx, &values, query); err != nil {
		return nil, fmt.Errorf("bots.SelectByBotIdList: %w", err)
	}
	return values, nil
}
