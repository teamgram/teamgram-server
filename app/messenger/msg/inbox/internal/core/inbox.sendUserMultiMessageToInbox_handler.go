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

// InboxSendUserMultiMessageToInbox
// inbox.sendUserMultiMessageToInbox from_id:long peer_user_id:long message:Vector<InboxMessageData> = Void;
func (c *InboxCore) InboxSendUserMultiMessageToInbox(in *inbox.TLInboxSendUserMultiMessageToInbox) (*mtproto.Void, error) {
	if in.FromId == in.PeerUserId {
		c.Logger.Errorf("inbox.sendUserMultiMessageToInbox - error: sendToSelfUser")
		err := mtproto.ErrPeerIdInvalid
		return nil, err
	}

	inBoxList, err := c.svcCtx.Dao.SendUserMultiMessageToInbox(c.ctx,
		in.FromId,
		in.PeerUserId,
		in.Message)
	if err != nil {
		c.Logger.Errorf("inbox.sendUserMultiMessageToInbox - error: %v", err)
		return nil, err
	}

	pushUpdates := c.makeUpdateNewMessageListUpdates(in.PeerUserId, inBoxList...)

	_, err = c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
		UserId:  in.PeerUserId,
		Updates: pushUpdates,
	})
	if err != nil {
		c.Logger.Errorf("inbox.sendUserMultiMessageToInbox - error: %v", err)
		return nil, err
	}

	return mtproto.EmptyVoid, nil
}
