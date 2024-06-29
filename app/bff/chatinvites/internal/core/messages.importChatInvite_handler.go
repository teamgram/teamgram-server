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
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"math/rand"

	"github.com/teamgram/proto/mtproto"
	msgpb "github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// MessagesImportChatInvite
// messages.importChatInvite#6c50051c hash:string = Updates;
func (c *ChatInvitesCore) MessagesImportChatInvite(in *mtproto.TLMessagesImportChatInvite) (*mtproto.Updates, error) {
	// Code	Type	Description
	// 400	INVITE_HASH_EMPTY	The invite hash is empty.
	// 400	INVITE_HASH_EXPIRED	The invite link has expired.
	// 400	INVITE_HASH_INVALID	The invite hash is invalid.

	if len(in.Hash) == 0 {
		err := mtproto.ErrInviteHashEmpty
		c.Logger.Errorf("messages.importChatInvite - error: %v", err)
		return nil, err
	}
	if len(in.Hash) != 20 {
		err := mtproto.ErrInviteHashInvalid
		c.Logger.Errorf("messages.importChatInvite - error: %v", err)
		return nil, err
	}

	if !chatpb.IsChatInviteHash(in.Hash) {
		err := mtproto.ErrInviteHashInvalid
		c.Logger.Errorf("messages.importChatInvite - error: %v", err)
		return nil, err
	}

	chatInviteImported, err := c.svcCtx.Dao.ChatClient.ChatImportChatInvite2(c.ctx, &chatpb.TLChatImportChatInvite2{
		SelfId: c.MD.UserId,
		Hash:   in.Hash,
	})
	if err != nil {
		c.Logger.Errorf("messages.importChatInvite - error: %v", err)
		return nil, err
	}

	mChat := chatInviteImported.GetChat()

	if chatInviteImported.GetRequesters() != nil {
		//c.Logger.Errorf("error: %v, chat: %s", err, mChat)
		//if nErr, ok := status.FromError(err); ok {
		//	if nErr.Message() == "INVITE_REQUEST_SENT" {
		requesters := chatInviteImported.GetRequesters()

		updatePendingJoinRequests := mtproto.MakeTLUpdatePendingJoinRequests(&mtproto.Update{
			Peer_PEER:        mtproto.MakePeerChat(mChat.Id()),
			RequestsPending:  requesters.GetRequestsPending(),
			RecentRequesters: requesters.GetRecentRequesters(),
		}).To_Update()

		idList := make([]int64, 0)
		mChat.Walk(func(userId int64, participant *mtproto.ImmutableChatParticipant) error {
			idList = append(idList, participant.GetUserId())
			return nil
		})

		pushUserList, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &user.TLUserGetMutableUsers{
			Id: idList,
		})

		mChat.Walk(func(userId int64, participant *mtproto.ImmutableChatParticipant) error {
			c.svcCtx.Dao.SyncClient.SyncPushUpdates(
				c.ctx,
				&sync.TLSyncPushUpdates{
					UserId: participant.UserId,
					Updates: mtproto.MakeUpdatesByUpdatesUsers(
						pushUserList.GetUserListByIdList(participant.UserId, c.MD.UserId),
						updatePendingJoinRequests),
				},
			)

			return nil
		})

		c.Logger.Errorf("messages.importChatInvite - reply: %v", mtproto.ErrInviteRequestSent)
		return nil, mtproto.ErrInviteRequestSent
	} else {
		// TODO: found link
		rUpdates, err := c.svcCtx.Dao.MsgClient.MsgSendMessageV2(
			c.ctx,
			&msgpb.TLMsgSendMessageV2{
				UserId:    c.MD.UserId,
				AuthKeyId: c.MD.PermAuthKeyId,
				PeerType:  mtproto.PEER_CHAT,
				PeerId:    mChat.Id(),
				Message: []*msgpb.OutboxMessage{
					msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
						NoWebpage:    true,
						Background:   false,
						RandomId:     rand.Int63(),
						Message:      mChat.MakeMessageService(c.MD.UserId, mtproto.MakeMessageActionChatJoinByLink(mChat.Creator())),
						ScheduleDate: nil,
					}).To_OutboxMessage(),
				},
			})
		if err != nil {
			c.Logger.Errorf("messages.importChatInvite - error: %v", err)
			return nil, err
		}

		return rUpdates, nil
	}
}
