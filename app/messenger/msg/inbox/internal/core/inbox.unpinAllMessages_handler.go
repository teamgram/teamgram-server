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

// InboxUnpinAllMessages
// inbox.unpinAllMessages user_id:long auth_key_id:long peer_type:int peer_id:long = Void;
func (c *InboxCore) InboxUnpinAllMessages(in *inbox.TLInboxUnpinAllMessages) (*mtproto.Void, error) {
	var (
		peer     = mtproto.MakePeerUtil(in.PeerType, in.PeerId)
		idList   = make([]int32, 0)
		pts      int32
		ptsCount int32
	)

	switch peer.PeerType {
	case mtproto.PEER_USER:
		dialogId := mtproto.MakeDialogId(peer.PeerId, peer.PeerType, in.UserId)
		c.svcCtx.Dao.MessagesDAO.SelectPinnedListWithCB(
			c.ctx,
			peer.PeerId,
			dialogId.A,
			dialogId.B,
			func(sz, i int, v *dataobject.MessagesDO) {
				idList = append(idList, v.UserMessageBoxId)
			})
		if len(idList) == 0 {
			break
		}

		pts = c.svcCtx.Dao.IDGenClient2.NextNPtsId(c.ctx, peer.PeerId, len(idList))
		ptsCount = int32(len(idList))
		c.svcCtx.Dao.SyncClient.SyncPushUpdates(
			c.ctx,
			&sync.TLSyncPushUpdates{
				UserId: peer.PeerId,
				Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdatePinnedMessages(&mtproto.Update{
					Pinned:    false,
					Peer_PEER: mtproto.MakePeerUser(in.UserId),
					Messages:  idList,
					Pts_INT32: pts,
					PtsCount:  ptsCount,
				}).To_Update()),
			})
	case mtproto.PEER_CHAT:
		// TODO: 性能优化
		dialogId := mtproto.MakeDialogId(0, peer.PeerType, in.PeerId)
		c.svcCtx.Dao.ChatParticipantsDAO.SelectListWithCB(
			c.ctx,
			peer.PeerId,
			func(sz, i int, v *dataobject.ChatParticipantsDO) {
				if v.UserId == in.UserId {
					return
				}
				if v.State != mtproto.ChatMemberStateNormal {
					return
				}

				c.svcCtx.Dao.MessagesDAO.SelectPinnedListWithCB(
					c.ctx,
					v.UserId,
					dialogId.A,
					dialogId.B,
					func(sz, i int, v *dataobject.MessagesDO) {
						idList = append(idList, v.UserMessageBoxId)
					})

				if len(idList) == 0 {
					return
				}

				pts = c.svcCtx.Dao.IDGenClient2.NextNPtsId(c.ctx, v.UserId, len(idList))
				ptsCount = int32(len(idList))
				c.svcCtx.Dao.SyncClient.SyncPushUpdates(
					c.ctx,
					&sync.TLSyncPushUpdates{
						UserId: v.UserId,
						Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdatePinnedMessages(&mtproto.Update{
							Pinned:    false,
							Peer_PEER: mtproto.MakePeerChat(peer.PeerId),
							Messages:  idList,
							Pts_INT32: pts,
							PtsCount:  ptsCount,
						}).To_Update()),
					})
			},
		)
	case mtproto.PEER_CHANNEL:
	}

	return mtproto.EmptyVoid, nil
}
