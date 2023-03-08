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
	"context"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/zeromicro/go-zero/core/jsonx"
)

// ChatSetChatAvailableReactions
// chat.setChatAvailableReactions self_id:long chat_id:long available_reactions:Vector<string> = Bool;
func (c *ChatCore) ChatSetChatAvailableReactions(in *chat.TLChatSetChatAvailableReactions) (*mtproto.MutableChat, error) {
	var (
		chat2 *mtproto.MutableChat
		me    *mtproto.ImmutableChatParticipant
		err   error
	)

	chat2, err = c.svcCtx.Dao.GetMutableChat(c.ctx, in.ChatId, in.SelfId)
	if err != nil {
		c.Logger.Errorf("chat.setChatAvailableReactions - error: %v")
		return nil, err
	}

	me, _ = chat2.GetImmutableChatParticipant(in.SelfId)
	if me == nil || me.State != mtproto.ChatMemberStateNormal {
		err = mtproto.ErrParticipantIdInvalid
		c.Logger.Errorf("chat.setChatAvailableReactions - error: %v")
		return nil, err
	}

	if !me.CanAdminAddAdmins() {
		err = mtproto.ErrChatAdminRequired
		c.Logger.Errorf("chat.setChatAvailableReactions - error: %v")
		return nil, err
	}

	var (
		availableReactions string
	)

	if len(in.AvailableReactions) > 0 {
		availableReactionsData, _ := jsonx.Marshal(in.AvailableReactions)
		if availableReactionsData != nil {
			availableReactions = string(availableReactionsData)
		}
	}

	_, _, err = c.svcCtx.Dao.CachedConn.Exec(
		c.ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			affected, err2 := c.svcCtx.Dao.ChatsDAO.UpdateAvailableReactions(c.ctx, in.AvailableReactionsType, availableReactions, in.ChatId)
			return 0, affected, err2
		},
		c.svcCtx.Dao.GetChatCacheKey(in.ChatId))
	if err != nil {
		c.Logger.Errorf("chat.setChatAvailableReactions - error: %v")
		return nil, err
	}

	chat2.Chat.AvailableReactions = in.AvailableReactions
	return chat2, nil
}
