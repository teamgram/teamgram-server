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

package messages

import (
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	// "github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/channel"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/bot"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/channel"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/chat"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/dialog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/message"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/sticker"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/user"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/username"
)

type MessagesServiceImpl struct {
	*user.UserModel
	*message.MessageModel
	*chat.ChatModel
	*channel.ChannelModel
	*sticker.StickerModel
	*dialog.DialogModel
	*username.UsernameModel
	*bot.BotModel
}

func NewMessagesServiceImpl(models []core.CoreModel) *MessagesServiceImpl {
	impl := &MessagesServiceImpl{}

	for _, m := range models {
		switch m.(type) {
		case *user.UserModel:
			impl.UserModel = m.(*user.UserModel)
		case *message.MessageModel:
			impl.MessageModel = m.(*message.MessageModel)
		case *chat.ChatModel:
			impl.ChatModel = m.(*chat.ChatModel)
		case *channel.ChannelModel:
			impl.ChannelModel = m.(*channel.ChannelModel)
		case *sticker.StickerModel:
			impl.StickerModel = m.(*sticker.StickerModel)
		case *dialog.DialogModel:
			impl.DialogModel = m.(*dialog.DialogModel)
		case *username.UsernameModel:
			impl.UsernameModel = m.(*username.UsernameModel)
		case *bot.BotModel:
			impl.BotModel = m.(*bot.BotModel)
		}
	}

	return impl
}
