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

// InboxDeleteChatHistoryToInbox
// inbox.deleteChatHistoryToInbox from_id:long peer_chat_id:long max_id:int = Void;
func (c *InboxCore) InboxDeleteChatHistoryToInbox(in *inbox.TLInboxDeleteChatHistoryToInbox) (*mtproto.Void, error) {
	var (
		pts, ptsCount int32
		peer          = mtproto.MakePeerUtil(mtproto.PEER_CHAT, in.PeerChatId)
	)

	lastMessage, deleteIds := c.svcCtx.Dao.GetLastMessageAndIdListByDialog(c.ctx, in.FromId, peer)
	if len(deleteIds) == 0 ||
		len(deleteIds) == 1 &&
			lastMessage.GetPredicateName() == mtproto.Predicate_messageService &&
			lastMessage.GetAction().GetPredicateName() == mtproto.Predicate_messageActionHistoryClear {
		return mtproto.EmptyVoid, nil
	}

	// TODO(@benqi): chat
	pts = c.svcCtx.Dao.IDGenClient2.NextNPtsId(c.ctx, in.FromId, len(deleteIds)+1)
	ptsCount = int32(len(deleteIds) + 1)

	if _, err := c.svcCtx.Dao.DeleteByMessageIdList(c.ctx, in.FromId, deleteIds); err != nil {
		return nil, err
	}

	pushUpdates := mtproto.MakeUpdatesByUpdates(
		mtproto.MakeTLUpdateDeleteMessages(&mtproto.Update{
			Messages:  deleteIds,
			Pts_INT32: pts - 2,
			PtsCount:  ptsCount - 2,
		}).To_Update(),
		mtproto.MakeTLUpdateReadHistoryInbox(&mtproto.Update{
			Peer_PEER: peer.ToPeer(),
			MaxId:     lastMessage.Id,
			Pts_INT32: pts - 1,
			PtsCount:  1,
		}).To_Update(),
	)

	_, _ = c.svcCtx.Dao.SyncClient.SyncPushUpdates(
		c.ctx,
		&sync.TLSyncPushUpdates{
			UserId:  in.FromId,
			Updates: pushUpdates,
		})

	return mtproto.EmptyVoid, nil
}
