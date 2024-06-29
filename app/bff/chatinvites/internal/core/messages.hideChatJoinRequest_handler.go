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
	msgpb "github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"math/rand"
)

// MessagesHideChatJoinRequest
// messages.hideChatJoinRequest#7fe7e815 flags:# approved:flags.0?true peer:InputPeer user_id:InputUser = Updates;
func (c *ChatInvitesCore) MessagesHideChatJoinRequest(in *mtproto.TLMessagesHideChatJoinRequest) (*mtproto.Updates, error) {
	var (
		peer           = mtproto.FromInputPeer2(c.MD.UserId, in.GetPeer())
		userId         = mtproto.FromInputUser(c.MD.UserId, in.GetUserId())
		pushUserIdList = make([]int64, 0)
		pushUsers      *user.Vector_ImmutableUser
	)

	if userId.IsSelf() {
		c.Logger.Errorf("messages.hideChatJoinRequest - error: method MessagesHideChatJoinRequest not impl")
		return nil, mtproto.ErrUserIdInvalid
	}

	if !peer.IsChat() {
		c.Logger.Errorf("messages.hideChatJoinRequest - error: method MessagesHideChatJoinRequest not impl")
		return nil, mtproto.ErrPeerIdInvalid
	}

	mChat, err := c.svcCtx.Dao.ChatClient.ChatGetMutableChat(c.ctx, &chatpb.TLChatGetMutableChat{
		ChatId: peer.PeerId,
	})
	if err != nil {
		c.Logger.Errorf("messages.hideChatJoinRequest - error: %v", err)
		return nil, mtproto.ErrPeerIdInvalid
	}

	me, _ := mChat.GetImmutableChatParticipant(c.MD.UserId)
	if me == nil || !me.CanInviteUsers() {
		c.Logger.Errorf("messages.hideChatJoinRequest - error: %v", err)
		return nil, mtproto.ErrPeerIdInvalid
	}
	join, _ := mChat.GetImmutableChatParticipant(userId.PeerId)
	if join != nil && join.IsChatMemberStateNormal() {
		c.Logger.Errorf("messages.hideChatJoinRequest - error: %v", err)
		return nil, mtproto.ErrUserIdInvalid
	}

	pendingJoinRequests, err := c.svcCtx.Dao.ChatClient.ChatHideChatJoinRequests(c.ctx, &chatpb.TLChatHideChatJoinRequests{
		SelfId:   c.MD.UserId,
		ChatId:   peer.PeerId,
		Approved: in.GetApproved(),
		Link:     nil,
		UserId:   mtproto.MakeFlagsInt64(userId.PeerId),
	})
	if err != nil {
		return mtproto.MakeEmptyUpdates(), nil
	}

	updatePendingJoinRequests := mtproto.MakeTLUpdatePendingJoinRequests(&mtproto.Update{
		Peer_PEER:        mtproto.MakePeerChat(mChat.Id()),
		RequestsPending:  pendingJoinRequests.RequestsPending,
		RecentRequesters: pendingJoinRequests.RecentRequesters,
	}).To_Update()

	pushUserIdList = append(pushUserIdList, pendingJoinRequests.GetRecentRequesters()...)
	if len(pushUserIdList) > 0 {
		mChat.Walk(func(userId int64, participant *mtproto.ImmutableChatParticipant) error {
			if participant.CanInviteUsers() {
				pushUserIdList = append(pushUserIdList, participant.UserId)
			}
			return nil
		})
		pushUsers, _ = c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &user.TLUserGetMutableUsers{
			Id: pushUserIdList,
		})
	}

	if in.GetApproved() {
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
						Message:      mChat.MakeMessageService(c.MD.UserId, mtproto.MakeMessageActionChatJoinedByRequest()),
						ScheduleDate: nil,
					}).To_OutboxMessage(),
				},
			})
		if err != nil {
			c.Logger.Errorf("messages.importChatInvite - error: %v", err)
			return nil, err
		}

		rUpdates.Updates = append(rUpdates.Updates, updatePendingJoinRequests)
		return rUpdates, nil
	} else {
		mChat.Walk(func(userId int64, participant *mtproto.ImmutableChatParticipant) error {
			if c.MD.UserId == participant.UserId {
				return nil
			}

			if participant.CanInviteUsers() {
				if len(pushUserIdList) > 0 {
					c.svcCtx.Dao.SyncClient.SyncPushUpdates(
						c.ctx,
						&sync.TLSyncPushUpdates{
							UserId: participant.UserId,
							Updates: mtproto.MakeUpdatesByUpdates(
								updatePendingJoinRequests),
						},
					)
				} else {
					c.svcCtx.Dao.SyncClient.SyncPushUpdates(
						c.ctx,
						&sync.TLSyncPushUpdates{
							UserId: participant.UserId,
							Updates: mtproto.MakeUpdatesByUpdatesUsers(
								pushUsers.GetUserListByIdList(participant.UserId, pendingJoinRequests.RecentRequesters...),
								updatePendingJoinRequests),
						},
					)
				}
			}
			return nil
		})

		var (
			rUpdates *mtproto.Updates
		)

		if len(pushUserIdList) > 0 {
			rUpdates = mtproto.MakeUpdatesByUpdatesUsers(
				pushUsers.GetUserListByIdList(c.MD.UserId, pendingJoinRequests.RecentRequesters...),
				updatePendingJoinRequests)
		} else {
			rUpdates = mtproto.MakeUpdatesByUpdates(
				updatePendingJoinRequests)
		}

		return rUpdates, nil
	}
}
