// Copyright (c) 2026 The Teamgram Authors (https://teamgram.net).
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

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// MessagesEditChatParticipantRank
// messages.editChatParticipantRank#a00f32b0 peer:InputPeer participant:InputPeer rank:string = Updates;
func (c *ChatsCore) MessagesEditChatParticipantRank(in *mtproto.TLMessagesEditChatParticipantRank) (*mtproto.Updates, error) {
	var (
		peer        = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		participant = mtproto.FromInputPeer2(c.MD.UserId, in.Participant)
	)

	if !peer.IsChat() {
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.editChatParticipantRank - invalid peer, err: %v", err)
		return nil, err
	}

	if !participant.IsUser() {
		err := mtproto.ErrUserIdInvalid
		c.Logger.Errorf("messages.editChatParticipantRank - invalid participant, err: %v", err)
		return nil, err
	}

	chat, err := c.svcCtx.Dao.ChatClient.Client().ChatEditChatParticipantRank(c.ctx, &chatpb.TLChatEditChatParticipantRank{
		SelfId:      c.MD.UserId,
		ChatId:      peer.PeerId,
		Participant: participant.PeerId,
		Rank:        in.Rank,
	})
	if err != nil {
		c.Logger.Errorf("messages.editChatParticipantRank - error: %v", err)
		return nil, err
	}

	var (
		idList []int64
	)

	updateChatParticipants := mtproto.MakeTLUpdateChatParticipants(&mtproto.Update{
		Participants_CHATPARTICIPANTS: chat.ToChatParticipants(0),
	}).To_Update()

	chat.Walk(func(userId int64, participant *mtproto.ImmutableChatParticipant) error {
		if participant.IsChatMemberStateNormal() {
			idList = append(idList, userId)
		}
		return nil
	})

	mUsers, err := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: idList,
	})
	if err != nil {
		c.Logger.Errorf("messages.editChatParticipantRank - error: %v", err)
		return nil, err
	}

	chat.Walk(func(userId int64, participant *mtproto.ImmutableChatParticipant) error {
		if !participant.IsChatMemberStateNormal() {
			return nil
		}

		c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx,
			&sync.TLSyncPushUpdates{
				UserId: userId,
				Updates: mtproto.MakeUpdatesByUpdatesUsersChats(
					mUsers.GetUserListByIdList(userId, idList...),
					[]*mtproto.Chat{chat.ToUnsafeChat(userId)},
					updateChatParticipants),
			})

		return nil
	})

	return mtproto.MakeUpdatesByUpdatesUsersChats(
		mUsers.GetUserListByIdList(c.MD.UserId, idList...),
		[]*mtproto.Chat{chat.ToUnsafeChat(c.MD.UserId)},
		updateChatParticipants), nil
}
