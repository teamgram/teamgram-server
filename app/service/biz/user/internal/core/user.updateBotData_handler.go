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

package core

import (
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UserUpdateBotData
// user.updateBotData flags:# bot_id:long bot_chat_history:flags.15?Bool bot_nochats:flags.16?Bool bot_inline_geo:flags.21?Bool bot_attach_menu:flags.27?Bool bot_inline_placeholder:flags.19?string bot_has_main_app:flags.13?Bool = Bool;
func (c *UserCore) UserUpdateBotData(in *user.TLUserUpdateBotData) (*tg.Bool, error) {
	if err := c.svcCtx.Repo.UpdateBotData(
		c.ctx,
		in.BotId,
		in.BotChatHistory,
		in.BotNochats,
		in.BotInlineGeo,
		in.BotAttachMenu,
		in.BotHasMainApp,
		in.BotInlinePlaceholder); err != nil {
		return nil, err
	}
	return tg.BoolTrue, nil
}
