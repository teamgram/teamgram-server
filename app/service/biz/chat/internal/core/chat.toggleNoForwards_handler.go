// Copyright 2022 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"time"
)

// ChatToggleNoForwards
// chat.toggleNoForwards chat_id:long operator_id:long enabled:Bool = MutableChat;
func (c *ChatCore) ChatToggleNoForwards(in *chat.TLChatToggleNoForwards) (*chat.MutableChat, error) {
	var (
		now   = time.Now().Unix()
		chat2 *chat.MutableChat
		me    *chat.ImmutableChatParticipant
		err   error
	)

	chat2, err = c.svcCtx.Dao.GetMutableChat(c.ctx, in.ChatId, in.OperatorId)
	if err != nil {
		c.Logger.Errorf("chat.toggleNoForwards - error: %v", err)
		return nil, err
	}

	me, _ = chat2.GetImmutableChatParticipant(in.OperatorId)
	if me == nil || me.State != mtproto.ChatMemberStateNormal {
		err = mtproto.ErrInputUserDeactivated
		c.Logger.Errorf("chat.toggleNoForwards - error: %v", err)
		return nil, err
	}

	if !me.IsChatMemberCreator() {
		err = mtproto.ErrChatAdminRequired
		c.Logger.Errorf("chat.toggleNoForwards - error: %v", err)
		return nil, err
	}

	_, err = c.svcCtx.Dao.ChatsDAO.UpdateNoforwards(c.ctx, mtproto.FromBool(in.Enabled), in.ChatId)
	if err != nil {
		c.Logger.Errorf("chat.toggleNoForwards - error: %v", err)
		return nil, err
	}
	chat2.Chat.Version += 1
	chat2.Chat.Date = now
	chat2.Chat.Noforwards = mtproto.FromBool(in.Enabled)

	return chat2, nil
}
