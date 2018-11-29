// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

package dataobject

type BotsDO struct {
	Id                   int32  `db:"id"`
	BotId                int32  `db:"bot_id"`
	BotType              int8   `db:"bot_type"`
	Description          string `db:"description"`
	BotChatHistory       int8   `db:"bot_chat_history"`
	BotNochats           int8   `db:"bot_nochats"`
	Verified             int8   `db:"verified"`
	BotInlineGeo         int8   `db:"bot_inline_geo"`
	BotInfoVersion       int32  `db:"bot_info_version"`
	BotInlinePlaceholder string `db:"bot_inline_placeholder"`
	CreatedAt            string `db:"created_at"`
	UpdatedAt            string `db:"updated_at"`
}
