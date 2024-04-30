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
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

// MessagesGetFullChat
// messages.getFullChat#aeb00b34 chat_id:long = messages.ChatFull;
func (c *ChatsCore) MessagesGetFullChat(in *mtproto.TLMessagesGetFullChat) (*mtproto.Messages_ChatFull, error) {
	chat, err := c.svcCtx.Dao.ChatClient.Client().ChatGetMutableChat(c.ctx, &chatpb.TLChatGetMutableChat{
		ChatId: in.ChatId,
	})
	if err != nil {
		c.Logger.Errorf("messages.getFullChat - error: %v", err)
		return nil, err
	}

	me, ok := chat.GetImmutableChatParticipant(c.MD.UserId)
	if !ok {
		c.Logger.Errorf("messages.getFullChat - error: not participant{chat_id: %d, chat_participant_id: %d}", in.ChatId, c.MD.UserId)
		err = mtproto.ErrPeerIdInvalid
		return nil, err
	}

	dialog2, err := c.svcCtx.Dao.DialogClient.DialogGetDialogsByIdList(c.ctx, &dialog.TLDialogGetDialogsByIdList{
		UserId: c.MD.UserId,
		IdList: []int64{mtproto.MakePeerDialogId(mtproto.PEER_CHAT, in.ChatId)},
	})
	if err != nil {
		c.Logger.Errorf("messages.getFullChat - error: %v", err)
		err = mtproto.ErrPeerIdInvalid
		return nil, err
	} else if len(dialog2.Datas) == 0 {
		c.Logger.Errorf("messages.getFullChat - error: not found dialog")
		err = mtproto.ErrPeerIdInvalid
		return nil, err
	}

	dlg := dialog2.Datas[0].GetDialog()
	chatFull := mtproto.MakeTLChatFull(&mtproto.ChatFull{
		CanSetUsername:                       true,
		HasScheduled:                         false, // TODO
		Id:                                   chat.Id(),
		About:                                chat.About(),
		Participants:                         chat.ToChatParticipants(c.MD.UserId),
		ChatPhoto:                            chat.Photo(),
		NotifySettings:                       nil,
		ExportedInvite:                       nil, // TODO
		BotInfo:                              nil, // TODO
		PinnedMsgId:                          nil, // TODO
		FolderId:                             dlg.FolderId,
		Call:                                 chat.Call(),
		TtlPeriod:                            mtproto.MakeFlagsInt32(chat.TTLPeriod()), // TODO
		GroupcallDefaultJoinAs:               nil,                                      // TODO
		ThemeEmoticon:                        nil,                                      // TODO
		RequestsPending:                      nil,                                      // TODO
		RecentRequesters:                     nil,                                      // TODO
		AvailableReactions_FLAGVECTORSTRING:  chat.GetChat().GetAvailableReactions(),
		AvailableReactions_FLAGCHATREACTIONS: chat.AvailableReactions(),
	}).To_ChatFull()

	var (
		idList []int64
	)

	// NotifySettings
	if settings, _ := c.svcCtx.Dao.UserClient.UserGetNotifySettings(c.ctx, &userpb.TLUserGetNotifySettings{
		UserId:   c.MD.UserId,
		PeerType: mtproto.PEER_CHAT,
		PeerId:   in.ChatId,
	}); settings != nil {
		chatFull.NotifySettings = settings
	}

	if me.CanInviteUsers() {
		if me.Link != "" {
			chatFull.ExportedInvite, _ = c.svcCtx.Dao.ChatClient.Client().ChatGetExportedChatInvite(c.ctx, &chatpb.TLChatGetExportedChatInvite{
				ChatId: in.ChatId,
				Link:   me.Link,
			})
		}

		requesters, _ := c.svcCtx.Dao.ChatClient.Client().ChatGetRecentChatInviteRequesters(c.ctx, &chatpb.TLChatGetRecentChatInviteRequesters{
			SelfId: c.MD.UserId,
			ChatId: in.ChatId,
		})

		if len(requesters.GetRecentRequesters()) > 0 {
			chatFull.RequestsPending = &wrapperspb.Int32Value{Value: requesters.GetRequestsPending()}
			chatFull.RecentRequesters = requesters.GetRecentRequesters()
		}
	}

	rValue := mtproto.MakeTLMessagesChatFull(&mtproto.Messages_ChatFull{
		FullChat: chatFull,
		Chats:    []*mtproto.Chat{chat.ToUnsafeChat(c.MD.UserId)},
		Users:    nil,
	}).To_Messages_ChatFull()

	chat.Walk(func(userId int64, participant *mtproto.ImmutableChatParticipant) error {
		if participant.IsChatMemberStateNormal() {
			idList = append(idList, participant.UserId)
			if participant.IsBot {
				// TODO: 优化
				botInfo, _ := c.svcCtx.Dao.UserClient.UserGetBotInfo(c.ctx, &userpb.TLUserGetBotInfo{
					BotId: participant.UserId,
				})
				if botInfo != nil {
					chatFull.BotInfo = append(chatFull.BotInfo, botInfo)
				}
			}
		}
		return nil
	})

	mUsers, err := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: idList,
	})
	if err != nil {
		c.Logger.Errorf("messages.getFullChat - error: not found dialog")
	}
	rValue.Users = mUsers.GetUserListByIdList(c.MD.UserId, idList...)

	return rValue, nil
}
