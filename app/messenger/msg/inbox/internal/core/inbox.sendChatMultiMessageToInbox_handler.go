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
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// InboxSendChatMultiMessageToInbox
// inbox.sendChatMultiMessageToInbox from_id:long peer_chat_id:long message:Vector<InboxMessageData> = Void;
func (c *InboxCore) InboxSendChatMultiMessageToInbox(in *inbox.TLInboxSendChatMultiMessageToInbox) (*mtproto.Void, error) {
	chat, err := c.svcCtx.Dao.ChatClient.ChatGetMutableChat(c.ctx, &chatpb.TLChatGetMutableChat{
		ChatId: in.PeerChatId,
	})
	if err != nil {
		c.Logger.Errorf("inbox.sendChatMultiMessageToInbox - error: %v", err)
		return nil, err
	}

	chat.Walk(func(userId int64, participant *chatpb.ImmutableChatParticipant) error {
		if in.FromId == userId {
			return nil
		}

		if !participant.IsChatMemberStateNormal() {
			return nil
		}

		inBoxList, err2 := c.svcCtx.Dao.SendChatMultiMessageToInbox(
			c.ctx,
			in.FromId,
			in.PeerChatId,
			userId,
			in.Message)
		if err2 != nil {
			c.Logger.Errorf("inbox.sendChatMultiMessageToInbox - error: %v", err2)
			return nil
		}

		_, err = c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
			UserId:  userId,
			Updates: c.makeUpdateNewMessageListUpdates(userId, inBoxList...),
		})
		if err != nil {
			c.Logger.Errorf("inbox.sendChatMultiMessageToInbox - error: %v", err)
		}

		return nil
	})

	return mtproto.EmptyVoid, nil
}
