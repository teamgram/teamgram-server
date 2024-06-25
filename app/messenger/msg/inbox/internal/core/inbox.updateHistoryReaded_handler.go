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
	"github.com/teamgram/teamgram-server/app/messenger/msg/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
)

// InboxUpdateHistoryReaded
// inbox.updateHistoryReaded from_id:long peer_type:int peer_id:long max_id:int = Void;
func (c *InboxCore) InboxUpdateHistoryReaded(in *inbox.TLInboxUpdateHistoryReaded) (*mtproto.Void, error) {
	switch in.PeerType {
	case mtproto.PEER_USER:
		replyId, err := c.svcCtx.Dao.MessagesDAO.SelectPeerUserMessage(
			c.ctx,
			in.PeerId,
			in.FromId,
			in.MaxId)
		if err != nil {
			c.Logger.Errorf("inbox.updateHistoryReaded - error: %v", err)
			return nil, err
		} else if replyId == nil {
			err = mtproto.ErrPeerIdInvalid
			c.Logger.Errorf("inbox.updateHistoryReaded - error: %v", err)
			return nil, err
		}
		c.Logger.Infof("inbox.updateHistoryReaded: %v", replyId)

		c.svcCtx.Dao.DialogsDAO.UpdateReadOutboxMaxId(
			c.ctx,
			replyId.UserMessageBoxId,
			in.PeerId,
			mtproto.MakePeerDialogId(in.PeerType, in.FromId))
		c.Logger.Infof("inbox.updateHistoryReaded: (%d, %d, %d)",
			replyId.UserMessageBoxId,
			in.PeerId,
			mtproto.MakePeerDialogId(in.PeerType, in.FromId))

		c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
			UserId: in.PeerId,
			Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateReadHistoryOutbox(&mtproto.Update{
				Peer_PEER: mtproto.MakePeerUser(in.FromId),
				MaxId:     replyId.UserMessageBoxId,
				Pts_INT32: c.svcCtx.Dao.IDGenClient2.NextPtsId(c.ctx, in.PeerId),
				PtsCount:  1,
			}).To_Update()),
		})
	case mtproto.PEER_CHAT:
		replyId, err := c.svcCtx.Dao.MessagesDAO.SelectPeerUserMessage(
			c.ctx,
			in.Sender,
			in.FromId,
			in.MaxId)
		if err != nil {
			c.Logger.Errorf("inbox.updateHistoryReaded - error: %v", err)
			return nil, err
		} else if replyId == nil {
			err = mtproto.ErrPeerIdInvalid
			c.Logger.Errorf("inbox.updateHistoryReaded - error: %v", err)
			return nil, err
		}
		c.Logger.Infof("inbox.updateHistoryReaded: %v", replyId)

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
						UserId: replyId.UserId,
						Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateReadHistoryOutbox(&mtproto.Update{
							Peer_PEER: mtproto.MakePeerChat(in.PeerId),
							MaxId:     replyId.UserMessageBoxId,
							Pts_INT32: c.svcCtx.Dao.IDGenClient2.NextPtsId(c.ctx, replyId.UserId),
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
