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
)

// InboxSendUserMessageToInbox
// inbox.sendUserMessageToInbox from_id:long peer_user_id:long message:InboxMessageData = Void;
func (c *InboxCore) InboxSendUserMessageToInbox(in *inbox.TLInboxSendUserMessageToInbox) (*mtproto.Void, error) {
	if in.FromId == in.PeerUserId {
		c.Logger.Errorf("inbox.sendUserMessageToInbox - error: sendToSelfUser")
		err := mtproto.ErrPeerIdInvalid
		return nil, err
	}

	inBox, err := c.svcCtx.Dao.SendUserMessageToInbox(c.ctx,
		in.FromId,
		in.PeerUserId,
		in.GetMessage().GetDialogMessageId(),
		in.GetMessage().GetRandomId(),
		in.GetMessage().GetMessage())
	if err != nil {
		c.Logger.Errorf("inbox.sendUserMessageToInbox - error: %v", err)
		return nil, err
	}

	if inBox.DialogMessageId == 1 &&
		(in.FromId != 42777 && in.FromId != 424000) {
		//isContact, _ := s.UserFacade.GetContactAndMutual(ctx, toId, fromId)
		//if !isContact {
		//	s.UserFacade.AddPeerSettings(ctx, toId, model.MakeUserPeerUtil(fromId), &mtproto.PeerSettings{
		//		AddContact:   true,
		//		BlockContact: true,
		//	})
		//}
	}

	pushUpdates := c.makeUpdateNewMessageListUpdates(in.PeerUserId, inBox)

	var isBot = false
	for _, u := range pushUpdates.GetUsers() {
		if u.GetId() == in.PeerUserId {
			isBot = u.GetBot()
			break
		}
	}
	if isBot {
		if c.svcCtx.Dao.BotSyncClient != nil {
			_, err = c.svcCtx.Dao.BotSyncClient.SyncPushBotUpdates(c.ctx, &sync.TLSyncPushBotUpdates{
				UserId:  inBox.UserId,
				Updates: pushUpdates,
			})
		} else {
			// TODO: log
		}
	} else {
		_, err = c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
			UserId:  inBox.UserId,
			Updates: pushUpdates,
		})
	}
	if err != nil {
		c.Logger.Errorf("inbox.sendUserMessageToInbox - error: %v", err)
	}

	return mtproto.EmptyVoid, nil
}
