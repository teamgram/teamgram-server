// Copyright 2024 Teamgram Authors
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

// InboxReadMediaUnreadToInboxV2
// inbox.readMediaUnreadToInboxV2 user_id:long peer_type:int peer_id:long dialog_message_id:long = Void;
func (c *InboxCore) InboxReadMediaUnreadToInboxV2(in *inbox.TLInboxReadMediaUnreadToInboxV2) (*mtproto.Void, error) {
	unreadDO, err := c.svcCtx.Dao.MessagesDAO.SelectByMessageDataId(c.ctx, in.UserId, in.DialogMessageId)
	if err != nil {
		c.Logger.Errorf("inbox.readMediaUnreadToInboxV2 - error: %v", err)
		return nil, err
	}

	if !unreadDO.MediaUnread || !unreadDO.MediaUnread {
		return mtproto.EmptyVoid, nil
	}
	_, err = c.svcCtx.Dao.MessagesDAO.UpdateMediaUnread(c.ctx, unreadDO.UserId, unreadDO.UserMessageBoxId)
	if err != nil {
		c.Logger.Errorf("inbox.readMediaUnreadToInboxV2 - error: %v", err)
		return nil, err
	}

	pts := c.svcCtx.Dao.IDGenClient2.NextPtsId(c.ctx, in.UserId)
	if pts == 0 {
		c.Logger.Errorf("inbox.readMediaUnreadToInboxV2 - error: nextPtsId(%d) is 0", in.UserId)
		return nil, mtproto.ErrInternalServerError
	}

	c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
		UserId: in.UserId,
		Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateReadMessagesContents(&mtproto.Update{
			Messages:  []int32{unreadDO.UserMessageBoxId},
			Pts_INT32: pts,
			PtsCount:  1,
		}).To_Update()),
	})

	return mtproto.EmptyVoid, nil
}
