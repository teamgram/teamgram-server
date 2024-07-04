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
	"github.com/teamgram/teamgram-server/app/messenger/msg/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
)

// InboxReadOutboxHistory
// inbox.readOutboxHistory user_id:long peer_type:int peer_id:long max_dialog_message_id:int64 = Void;
func (c *InboxCore) InboxReadOutboxHistory(in *inbox.TLInboxReadOutboxHistory) (*mtproto.Void, error) {
	switch in.PeerType {
	case mtproto.PEER_USER:
		replyId, err := c.svcCtx.Dao.MessagesDAO.SelectByMessageDataId(
			c.ctx,
			in.UserId,
			in.MaxDialogMessageId)
		if err != nil {
			c.Logger.Errorf("inbox.readOutboxHistory - error: %v", err)
			return nil, err
		} else if replyId == nil {
			err = mtproto.ErrPeerIdInvalid
			c.Logger.Errorf("inbox.readOutboxHistory - error: %v", err)
			return nil, err
		}
		c.Logger.Infof("inbox.readOutboxHistory: %v", replyId)

		c.svcCtx.Dao.DialogsDAO.UpdateReadOutboxMaxId(
			c.ctx,
			replyId.UserMessageBoxId,
			in.UserId,
			mtproto.MakePeerDialogId(in.PeerType, in.PeerId))
		c.Logger.Infof("inbox.readOutboxHistory: (%d, %d, %d)",
			replyId.UserMessageBoxId,
			in.UserId,
			mtproto.MakePeerDialogId(in.PeerType, in.PeerId))

		c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
			UserId: in.UserId,
			Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateReadHistoryOutbox(&mtproto.Update{
				Peer_PEER: mtproto.MakePeerUser(in.PeerId),
				MaxId:     replyId.UserMessageBoxId,
				Pts_INT32: c.svcCtx.Dao.IDGenClient2.NextPtsId(c.ctx, in.UserId),
				PtsCount:  1,
			}).To_Update()),
		})
	case mtproto.PEER_CHAT:
		replyId, err := c.svcCtx.Dao.MessagesDAO.SelectByMessageDataId(
			c.ctx,
			in.UserId,
			in.MaxDialogMessageId)
		if err != nil {
			c.Logger.Errorf("inbox.readOutboxHistory - error: %v", err)
			return nil, err
		} else if replyId == nil {
			err = mtproto.ErrPeerIdInvalid
			c.Logger.Errorf("inbox.readOutboxHistory - error: %v", err)
			return nil, err
		}
		c.Logger.Infof("inbox.readOutboxHistory: %v", replyId)

		c.svcCtx.Dao.DialogsDAO.SelectPeerDialogListWithCB(
			c.ctx,
			replyId.UserId,
			[]int64{mtproto.MakePeerDialogId(in.PeerType, in.PeerId)},
			func(sz, i int, v *dataobject.DialogsDO) {
				if v.ReadOutboxMaxId < replyId.UserMessageBoxId {
					c.svcCtx.Dao.DialogsDAO.UpdateReadOutboxMaxId(
						c.ctx,
						replyId.UserMessageBoxId,
						replyId.UserId,
						mtproto.MakePeerDialogId(in.PeerType, in.PeerId))
					c.Logger.Infof("inbox.updateHistoryReaded: (%d, %d, %d)",
						replyId.UserMessageBoxId,
						replyId.PeerId,
						mtproto.MakePeerDialogId(in.PeerType, in.PeerId))

					// topMessage := c.svcCtx.Dao.DialogsDAO.Select
					c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
						UserId: in.UserId,
						Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateReadHistoryOutbox(&mtproto.Update{
							Peer_PEER: mtproto.MakePeerChat(in.PeerId),
							MaxId:     replyId.UserMessageBoxId,
							Pts_INT32: c.svcCtx.Dao.IDGenClient2.NextPtsId(c.ctx, in.UserId),
							PtsCount:  1,
						}).To_Update()),
					})
				}
			},
		)
	case mtproto.PEER_CHANNEL:
		// TODO
	}

	return mtproto.EmptyVoid, nil
}
