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
	msgpb "github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
)

// MessagesDeleteChat
// messages.deleteChat#5bd0ee50 chat_id:long = Bool;
func (c *ChatsCore) MessagesDeleteChat(in *mtproto.TLMessagesDeleteChat) (*mtproto.Bool, error) {
	operatorId := c.MD.UserId
	if c.MD.IsAdmin {
		operatorId = 0
	}

	// 2. delete chat
	chat, err := c.svcCtx.Dao.ChatClient.Client().ChatDeleteChat(c.ctx, &chatpb.TLChatDeleteChat{
		ChatId:     in.ChatId,
		OperatorId: operatorId,
	})
	if err != nil {
		c.Logger.Errorf("messages.deleteChat - error: %v", err)
		return nil, err
	}

	pushUpdates := mtproto.MakeUpdatesByUpdatesChats(
		[]*mtproto.Chat{chat.ToChatForbidden()},
		mtproto.MakeUpdateChat(chat.Id()))

	// 1. kicked all
	chat.Walk(func(userId int64, participant *mtproto.ImmutableChatParticipant) error {
		c.svcCtx.Dao.DialogClient.DialogDeleteDialog(c.ctx, &dialog.TLDialogDeleteDialog{
			UserId:   userId,
			PeerType: mtproto.PEER_CHAT,
			PeerId:   chat.Id(),
		})

		if userId == c.MD.UserId || participant.IsChatMemberStateNormal() {
			c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
				UserId:  userId,
				Updates: pushUpdates,
			})
		}

		c.svcCtx.Dao.MsgClient.MsgDeleteHistory(c.ctx, &msgpb.TLMsgDeleteHistory{
			UserId:    userId,
			AuthKeyId: 0,
			PeerType:  mtproto.PEER_CHAT,
			PeerId:    chat.Id(),
			JustClear: false,
			Revoke:    false,
		})
		return nil
	})

	return mtproto.BoolTrue, nil
}
