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
	"github.com/teamgram/marmota/pkg/container2/sets"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/app/messenger/msg/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// InboxUpdatePinnedMessage
// inbox.updatePinnedMessage flags:# user_id:long auth_key_id:long silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer_type:int peer_id:long id:int = Void;
func (c *InboxCore) InboxUpdatePinnedMessage(in *inbox.TLInboxUpdatePinnedMessage) (*mtproto.Void, error) {
	var (
		peer = mtproto.MakePeerUtil(in.PeerType, in.PeerId)
	)

	doUpdatePinnedMessageF := func(peer *mtproto.PeerUtil, v *dataobject.MessagesDO) {
		if v.UserId == in.UserId {
			return
		}

		var (
			pinnedMsgId int32 = 0
		)
		if in.GetUnpin() {
			idList, _ := c.svcCtx.Dao.MessagesDAO.SelectLastTwoPinnedList(c.ctx, v.UserId, v.DialogId1, v.DialogId2)
			if len(idList) == 2 {
				if v.UserMessageBoxId == idList[0] {
					pinnedMsgId = idList[1]
				} else {
					pinnedMsgId = idList[0]
				}
			}
		} else {
			pinnedMsgId = v.UserMessageBoxId
			// c.svcCtx.Dao.DialogsDAO.UpdatePinnedMsgId(c.ctx, v.UserMessageBoxId, v.UserId, mtproto.MakePeerDialogId(peer.PeerType, peer.PeerId))
		}

		c.svcCtx.Dao.MessagesDAO.UpdatePinned(c.ctx, !in.GetUnpin(), v.UserId, v.UserMessageBoxId)

		if peer.PeerType == mtproto.PEER_USER {
			c.svcCtx.Dao.DialogsDAO.UpdatePinnedMsgId(c.ctx, pinnedMsgId, v.UserId, mtproto.MakePeerDialogId(peer.PeerType, in.UserId))
			// sync
			c.svcCtx.Dao.SyncClient.SyncPushUpdates(
				c.ctx,
				&sync.TLSyncPushUpdates{
					UserId: v.UserId,
					Updates: mtproto.MakeUpdatesByUpdates(
						mtproto.MakeTLUpdatePinnedMessages(&mtproto.Update{
							Pinned:    !in.GetUnpin(),
							Peer_PEER: mtproto.MakePeerUser(in.UserId),
							Messages:  []int32{v.UserMessageBoxId},
							Pts_INT32: c.svcCtx.Dao.IDGenClient2.NextPtsId(c.ctx, v.UserId),
							PtsCount:  1,
						}).To_Update()),
				})
		} else {
			c.svcCtx.Dao.DialogsDAO.UpdatePinnedMsgId(c.ctx, pinnedMsgId, v.UserId, mtproto.MakePeerDialogId(peer.PeerType, peer.PeerId))
			// sync
			c.svcCtx.Dao.SyncClient.SyncPushUpdates(
				c.ctx,
				&sync.TLSyncPushUpdates{
					UserId: v.UserId,
					Updates: mtproto.MakeUpdatesByUpdates(
						mtproto.MakeTLUpdatePinnedMessages(&mtproto.Update{
							Pinned:    !in.GetUnpin(),
							Peer_PEER: mtproto.MakePeerChat(peer.PeerId),
							Messages:  []int32{v.UserMessageBoxId},
							Pts_INT32: c.svcCtx.Dao.IDGenClient2.NextPtsId(c.ctx, v.UserId),
							PtsCount:  1,
						}).To_Update()),
				})
		}
	}

	switch peer.PeerType {
	case mtproto.PEER_USER:
		c.svcCtx.Dao.MessagesDAO.SelectByMessageDataIdListWithCB(
			c.ctx,
			c.svcCtx.Dao.MessagesDAO.CalcTableName(in.PeerId),
			[]int64{in.DialogMessageId},
			func(sz, i int, v *dataobject.MessagesDO) {
				doUpdatePinnedMessageF(peer, v)
			})
	case mtproto.PEER_CHAT:
		var (
			tables = sets.NewWithLength(1)
		)

		pUserIdList, _ := c.svcCtx.ChatClient.ChatGetChatParticipantIdList(c.ctx, &chatpb.TLChatGetChatParticipantIdList{
			ChatId: peer.PeerId,
		})

		for _, uId := range pUserIdList.GetDatas() {
			tables.Insert(c.svcCtx.Dao.MessagesDAO.CalcTableName(uId))
		}

		for tableName, _ := range tables {
			c.svcCtx.Dao.MessagesDAO.SelectByMessageDataIdListWithCB(
				c.ctx,
				tableName,
				[]int64{in.DialogMessageId},
				func(sz, i int, v *dataobject.MessagesDO) {
					doUpdatePinnedMessageF(peer, v)
				})
		}
	case mtproto.PEER_CHANNEL:
	}

	return mtproto.EmptyVoid, nil
}
