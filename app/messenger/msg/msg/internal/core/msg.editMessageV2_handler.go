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
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/plugin"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"github.com/zeromicro/go-zero/core/mr"
)

// MsgEditMessageV2
// msg.editMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long edit_type:int new_message:OutboxMessage dst_message:Message = Updates;
func (c *MsgCore) MsgEditMessageV2(in *msg.TLMsgEditMessageV2) (*mtproto.Updates, error) {
	var (
		err        error
		rUpdates   *mtproto.Updates
		newMessage = in.NewMessage
		dstMessage = in.DstMessage
	)

	if dstMessage == nil {
		err = mtproto.ErrInputRequestInvalid
		c.Logger.Errorf("msg.editMessage - error: request(%s) error - %v", in, err)
		return nil, err
	}

	switch in.PeerType {
	case mtproto.PEER_USER:
		rUpdates, err = c.editUserOutgoingMessageV2(in.UserId, in.AuthKeyId, in.PeerId, newMessage, dstMessage)
	case mtproto.PEER_CHAT:
		rUpdates, err = c.editChatOutgoingMessageV2(in.UserId, in.AuthKeyId, in.PeerId, newMessage, dstMessage)
	case mtproto.PEER_CHANNEL:
		c.Logger.Errorf("msg.sendMessageV2 blocked, License key from https://teamgram.net required to unlock enterprise features.")
		return nil, mtproto.ErrEnterpriseIsBlocked
	default:
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("msg.editMessage - error: %v", err)
		return nil, err
	}

	if err != nil {
		c.Logger.Errorf("msg.editMessage - error: %v", err)
		return nil, err
	}

	return rUpdates, nil
}

func (c *MsgCore) editUserOutgoingMessageV2(fromUserId, fromAuthKeyId, toUserId int64, editBox *msg.OutboxMessage, dstMessage *mtproto.MessageBox) (*mtproto.Updates, error) {
	var (
		idHelper = mtproto.NewIDListHelper(fromUserId, toUserId)
	)

	idHelper.PickByMessage(editBox.GetMessage())

	users, err := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: idHelper.UserIdList,
		To: []int64{fromUserId, toUserId},
	})
	if err != nil {
		c.Logger.Errorf("msg.sendUserOutgoingMessageV2 - error: %v", err)
		return nil, err
	}

	sender, _ := users.GetImmutableUser(fromUserId)
	if sender == nil || sender.Deleted() {
		err = mtproto.ErrInputUserDeactivated
		c.Logger.Errorf("msg.sendUserOutgoingMessageV2 - error: %v", err)
		return nil, err
	}

	// TODO(@benqi): check
	// if sender.Restricted() {
	//	err = mtproto.ErrUserRestricted
	//	return
	// }

	peerUser, _ := users.GetImmutableUser(toUserId)
	if peerUser == nil || peerUser.Deleted() {
		err = mtproto.ErrInputUserDeactivated
		c.Logger.Errorf("msg.sendUserOutgoingMessage - error: %v", err)
		return nil, err
	}

	editBox.Message = plugin.RemakeMessage(
		c.ctx,
		c.svcCtx.MsgPlugin,
		editBox.Message,
		fromUserId,
		editBox.NoWebpage,
		func() bool {
			hasBot := false
			users.Visit(func(it *mtproto.ImmutableUser) {
				if it.IsBot() {
					hasBot = true
				}
			})

			return hasBot
		})

	outBox, err := c.svcCtx.Dao.EditUserOutboxMessageV2(c.ctx, fromUserId, toUserId, editBox, dstMessage)
	if err != nil {
		c.Logger.Errorf("msg.editMessage - error: %v", err)
		return nil, err
	}

	_ = outBox
	_, err2 := c.svcCtx.Dao.InboxClient.InboxEditMessageToInboxV2(
		c.ctx,
		&inbox.TLInboxEditMessageToInboxV2{
			UserId:        fromUserId,
			Out:           true,
			FromId:        fromUserId,
			FromAuthKeyId: fromAuthKeyId,
			PeerType:      mtproto.PEER_USER,
			PeerId:        toUserId,
			NewMessage:    outBox,
			DstMessage:    dstMessage,
			Users:         users.GetUserListByIdList(fromUserId, idHelper.UserIdList...),
			Chats:         nil,
		})
	if err2 != nil {
		return nil, err2
	}

	if fromUserId != toUserId {
		blocked, _ := c.svcCtx.Dao.UserClient.UserBlockedByUser(c.ctx, &userpb.TLUserBlockedByUser{
			UserId:     toUserId,
			PeerUserId: fromUserId,
		})

		if !mtproto.FromBool(blocked) {
			_, err2 = c.svcCtx.Dao.InboxClient.InboxEditMessageToInboxV2(
				c.ctx,
				&inbox.TLInboxEditMessageToInboxV2{
					UserId:        toUserId,
					Out:           false,
					FromId:        fromUserId,
					FromAuthKeyId: fromAuthKeyId,
					PeerType:      mtproto.PEER_USER,
					PeerId:        toUserId,
					NewMessage:    outBox,
					DstMessage:    nil,
					Users:         users.GetUserListByIdList(toUserId, idHelper.UserIdList...),
					Chats:         nil,
				})
			if err2 != nil {
				return nil, err2
			}
		}
	}

	return mtproto.MakeReplyUpdates(
		func(idList []int64) []*mtproto.User {
			return users.GetUserListByIdList(fromUserId, idList...)
		},
		func(idList []int64) []*mtproto.Chat {
			return []*mtproto.Chat{}
		},
		func(idList []int64) []*mtproto.Chat {
			return []*mtproto.Chat{}
		},
		mtproto.MakeTLUpdateEditMessage(&mtproto.Update{
			Pts_INT32:       outBox.Pts,
			PtsCount:        outBox.PtsCount,
			Message_MESSAGE: outBox.Message,
		}).To_Update()), nil
}

func (c *MsgCore) editChatOutgoingMessageV2(fromUserId, fromAuthKeyId, peerChatId int64, editBox *msg.OutboxMessage, dstMessage *mtproto.MessageBox) (*mtproto.Updates, error) {
	var (
		chat      *mtproto.MutableChat
		sUserList *mtproto.MutableUsers
		idHelper  = mtproto.NewIDListHelper(fromUserId)
	)
	idHelper.PickByMessage(editBox.GetMessage())

	err := mr.Finish(
		func() error {
			var (
				err error
			)
			chat, err = c.svcCtx.Dao.ChatClient.ChatGetMutableChat(
				c.ctx,
				&chatpb.TLChatGetMutableChat{
					ChatId: peerChatId,
				})
			if err != nil {
				c.Logger.Errorf("inbox.sendChatMessageToInbox - error: %v", err)
			}
			return err
		},
		func() error {
			var (
				err error
			)
			sUserList, err = c.svcCtx.Dao.UserClient.UserGetMutableUsersV2(
				c.ctx,
				&userpb.TLUserGetMutableUsersV2{
					Id:      idHelper.UserIdList,
					Privacy: true,
					HasTo:   true,
					To:      nil,
				})
			if err != nil {
				c.Logger.Errorf("inbox.sendChatMessageToInbox - error: %v", err)
			}

			return err
		})
	if err != nil {
		// c.Logger.Errorf("inbox.sendChatMessageToInbox - error: %v", err)
		return nil, err
	}

	if _, ok := chat.GetImmutableChatParticipant(fromUserId); !ok {
		c.Logger.Errorf("msg.sendChatOutgoingMessageV2 - error: ErrChatParticipantNotExists")
		err = mtproto.ErrChatWriteForbidden
		return nil, err
	}

	editBox.Message = plugin.RemakeMessage(
		c.ctx,
		c.svcCtx.MsgPlugin,
		editBox.Message,
		fromUserId,
		editBox.NoWebpage,
		func() bool {
			hasBot := false
			chat.Walk(func(userId int64, participant *mtproto.ImmutableChatParticipant) error {
				if participant.IsBot {
					hasBot = true
				}
				return nil
			})

			return hasBot
		})

	outBox, err2 := c.svcCtx.Dao.EditChatOutboxMessageV2(c.ctx, fromUserId, peerChatId, editBox, dstMessage)
	if err != nil {
		c.Logger.Errorf("msg.editMessage - error: %v", err)
		return nil, err
	}

	chat.Walk(func(userId int64, participant *mtproto.ImmutableChatParticipant) error {
		if !participant.IsChatMemberStateNormal() {
			return nil
		}
		if err2 != nil {
			return nil
		}

		out := participant.UserId == fromUserId
		if out {
			_, err2 = c.svcCtx.Dao.InboxClient.InboxEditMessageToInboxV2(
				c.ctx,
				&inbox.TLInboxEditMessageToInboxV2{
					UserId:        participant.UserId,
					Out:           true,
					FromId:        fromUserId,
					FromAuthKeyId: fromAuthKeyId,
					PeerType:      mtproto.PEER_CHAT,
					PeerId:        peerChatId,
					NewMessage:    outBox,
					DstMessage:    dstMessage,
					Users:         sUserList.GetUserListByIdList(participant.UserId, idHelper.UserIdList...),
					Chats:         []*mtproto.Chat{chat.ToUnsafeChat(participant.UserId)},
				})
		} else {
			toUsers := make([]*mtproto.User, 0, sUserList.Length())
			sUserList.Visit(func(it *mtproto.ImmutableUser) {
				toUsers = append(toUsers, it.ToUser(participant.UserId))
			})
			_, err2 = c.svcCtx.Dao.InboxClient.InboxEditMessageToInboxV2(
				c.ctx,
				&inbox.TLInboxEditMessageToInboxV2{
					UserId:        participant.UserId,
					Out:           false,
					FromId:        fromUserId,
					FromAuthKeyId: fromAuthKeyId,
					PeerType:      mtproto.PEER_CHAT,
					PeerId:        peerChatId,
					NewMessage:    outBox,
					DstMessage:    nil,
					Users:         toUsers,
					Chats:         []*mtproto.Chat{chat.ToUnsafeChat(participant.UserId)},
				})
		}
		return nil
	})

	if err2 != nil {
		c.Logger.Error(err2.Error())
		return nil, err
	}

	return mtproto.MakeReplyUpdates(
		func(idList []int64) []*mtproto.User {
			return sUserList.GetUserListByIdList(fromUserId, idList...)
		},
		func(idList []int64) []*mtproto.Chat {
			return []*mtproto.Chat{chat.ToUnsafeChat(fromUserId)}
		},
		func(idList []int64) []*mtproto.Chat {
			// TODO
			return nil
		},
		mtproto.MakeTLUpdateEditMessage(&mtproto.Update{
			Pts_INT32:       outBox.Pts,
			PtsCount:        outBox.PtsCount,
			Message_MESSAGE: outBox.Message,
		}).To_Update()), nil
}
