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
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// InboxSendChatMessageToInbox
// inbox.sendChatMessageToInbox from_id:long peer_chat_id:long message:InboxMessageData = Void;
func (c *InboxCore) InboxSendChatMessageToInbox(in *inbox.TLInboxSendChatMessageToInbox) (*mtproto.Void, error) {
	_, err := c.svcCtx.Dao.ChatParticipantsDAO.SelectListWithCB(
		c.ctx,
		in.PeerChatId,
		func(i int, v *dataobject.ChatParticipantsDO) {
			if v.UserId == in.FromId {
				return
			}

			switch v.State {
			case mtproto.ChatMemberStateNormal:
			case mtproto.ChatMemberStateLeft:
				return
			case mtproto.ChatMemberStateKicked:
				if in.Message.GetMessage().GetAction().GetPredicateName() == mtproto.Predicate_messageActionChatDeleteUser &&
					in.Message.GetMessage().GetAction().GetUserId() == v.UserId {
					// send messageActionChatDeleteUser message to deleteUser
				} else {
					return
				}
			case mtproto.ChatMemberStateMigrated:
				if in.Message.GetMessage().GetAction().GetPredicateName() == mtproto.Predicate_messageActionChatMigrateTo {
					// messageActionChatMigrateTo to user
				} else {
					return
				}
			default:
				return
			}

			inBox, err := c.svcCtx.Dao.SendChatMessageToInbox(
				c.ctx,
				in.FromId,
				in.PeerChatId,
				v.UserId,
				in.Message.DialogMessageId,
				in.Message.RandomId,
				in.Message.GetMessage())
			if err != nil {
				c.Logger.Errorf(err.Error())
				return
			}

			var (
				updates = []*mtproto.Update{mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
					Message_MESSAGE: inBox.Message,
					Pts_INT32:       inBox.Pts,
					PtsCount:        inBox.PtsCount,
				}).To_Update()}
			)

			if inBox.GetMessage().GetAction().GetPredicateName() == mtproto.Predicate_messageActionChatMigrateTo {
				c.svcCtx.Dao.DialogsDAO.UpdateReadInboxMaxId(
					c.ctx,
					inBox.MessageId,
					v.UserId,
					mtproto.MakePeerDialogId(mtproto.PEER_CHAT, in.PeerChatId))

				updates = append(updates, mtproto.MakeTLUpdateReadHistoryInbox(&mtproto.Update{
					FolderId:         nil,
					Peer_PEER:        mtproto.MakePeerChat(in.PeerChatId),
					MaxId:            inBox.MessageId,
					StillUnreadCount: 0,
					Pts_INT32:        c.svcCtx.Dao.NextPtsId(c.ctx, v.UserId),
					PtsCount:         1,
				}).To_Update())
			}

			pushUpdates := mtproto.MakePushUpdates(
				func(idList []int64) []*mtproto.User {
					users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
						&userpb.TLUserGetMutableUsers{
							Id: idList,
						})
					return users.GetUserListByIdList(v.UserId, idList...)
				},
				func(idList []int64) []*mtproto.Chat {
					chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
						&chatpb.TLChatGetChatListByIdList{
							IdList: idList,
						})
					return chats.GetChatListByIdList(v.UserId, idList...)
				},
				func(idList []int64) []*mtproto.Chat {
					// TODO
					return nil
				},
				updates...)

			c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
				UserId:  v.UserId,
				Updates: pushUpdates,
			})
		})
	if err != nil {
		c.Logger.Errorf("inbox.sendUserMessageToInbox - error: %v", err)
	}

	return mtproto.EmptyVoid, nil
}
