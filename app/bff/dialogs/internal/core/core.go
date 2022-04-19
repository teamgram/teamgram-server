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
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/rpc/metadata"
	"github.com/teamgram/teamgram-server/app/bff/dialogs/internal/svc"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DialogsCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *DialogsCore {
	return &DialogsCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}

type dialogsDataHelper struct {
	Dialogs  []*mtproto.Dialog
	Messages []*mtproto.Message
	Chats    []*mtproto.Chat
	Users    []*mtproto.User
}

func (m *dialogsDataHelper) ToMessagesDialogs(count int32) *mtproto.Messages_Dialogs {
	return mtproto.MakeTLMessagesDialogsSlice(&mtproto.Messages_Dialogs{
		Dialogs:  m.Dialogs,
		Messages: m.Messages,
		Chats:    m.Chats,
		Users:    m.Users,
		Count:    count,
	}).To_Messages_Dialogs()
}

func (m *dialogsDataHelper) ToMessagesPeerDialogs(state *mtproto.Updates_State) *mtproto.Messages_PeerDialogs {
	return mtproto.MakeTLMessagesPeerDialogs(&mtproto.Messages_PeerDialogs{
		Dialogs:  m.Dialogs,
		Messages: m.Messages,
		Users:    m.Users,
		Chats:    m.Chats,
		State:    state,
	}).To_Messages_PeerDialogs()
}

func (c *DialogsCore) makeMessagesDialogs(dialogExtList dialog.DialogExtList) *dialogsDataHelper {
	dialogsData := &dialogsDataHelper{
		Dialogs:  []*mtproto.Dialog{},
		Messages: []*mtproto.Message{},
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	}

	if len(dialogExtList) == 0 {
		return dialogsData
	}

	idHelper := mtproto.NewIDListHelper(c.MD.UserId)

	for _, dialogExt := range dialogExtList {
		// c.Logger.Infof("dialogEx: %v", dialogExt)
		peer2 := mtproto.FromPeer(dialogExt.Dialog.Peer)
		// idHelper.PickByPeer(peer2)
		switch peer2.PeerType {
		case mtproto.PEER_CHANNEL:
			if c.svcCtx.Plugin != nil {
				dialog, _ := c.svcCtx.Plugin.GetChannelDialogById(c.ctx, c.MD.UserId, peer2.PeerId)
				if dialog != nil {
					dialogExt.Dialog.TopMessage = dialog.Dialog.TopMessage
					dialogExt.Dialog.Pts = dialog.Dialog.Pts
					dialogExt.Dialog.UnreadCount = dialog.Dialog.TopMessage - dialogExt.Dialog.ReadInboxMaxId
					if dialog.Dialog.UnreadCount < 0 {
						dialog.Dialog.UnreadCount = 0
					}
					msgBox, _ := c.svcCtx.Plugin.GetChannelMessage(c.ctx, c.MD.UserId, peer2.PeerId, dialogExt.Dialog.TopMessage)
					if msgBox != nil {
						m := msgBox.ToMessage(c.MD.UserId)
						idHelper.PickByMessage(m)
						dialogsData.Messages = append(dialogsData.Messages, m)
						mentionsCount, _ := c.svcCtx.Dao.MessageClient.MessageGetUnreadMentionsCount(c.ctx, &message.TLMessageGetUnreadMentionsCount{
							UserId:   c.MD.UserId,
							PeerType: mtproto.PEER_CHANNEL,
							PeerId:   peer2.PeerId,
						})
						dialogExt.Dialog.UnreadMentionsCount = mentionsCount.GetV()
					}
				}
			} else {
				c.Logger.Errorf("blocked, License key from https://teamgram.net required to unlock enterprise features.")
			}
		default:
			msgBox, _ := c.svcCtx.Dao.MessageClient.MessageGetUserMessage(c.ctx, &message.TLMessageGetUserMessage{
				UserId: c.MD.UserId,
				Id:     dialogExt.Dialog.TopMessage,
			})
			if msgBox != nil {
				m := msgBox.ToMessage(c.MD.UserId)
				idHelper.PickByMessage(m)
				dialogsData.Messages = append(dialogsData.Messages, m)
				if peer2.PeerType == mtproto.PEER_CHAT {
					mentionsCount, _ := c.svcCtx.Dao.MessageClient.MessageGetUnreadMentionsCount(c.ctx, &message.TLMessageGetUnreadMentionsCount{
						UserId:   c.MD.UserId,
						PeerType: mtproto.PEER_CHAT,
						PeerId:   peer2.PeerId,
					})
					dialogExt.Dialog.UnreadMentionsCount = mentionsCount.GetV()
				}
			}
		}
		dialogExt.Dialog.NotifySettings, _ = c.svcCtx.Dao.UserClient.UserGetNotifySettings(c.ctx, &userpb.TLUserGetNotifySettings{
			UserId:   c.MD.UserId,
			PeerType: peer2.PeerType,
			PeerId:   peer2.PeerId,
		})
		dialogsData.Dialogs = append(dialogsData.Dialogs, dialogExt.Dialog)
	}

	idHelper.Visit(
		func(userIdList []int64) {
			users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: userIdList,
				})
			dialogsData.Users = append(dialogsData.Users, users.GetUserListByIdList(c.MD.UserId, userIdList...)...)
		},
		func(chatIdList []int64) {
			chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: chatIdList,
				})
			dialogsData.Chats = append(dialogsData.Chats, chats.GetChatListByIdList(c.MD.UserId, chatIdList...)...)
		},
		func(channelIdList []int64) {
			if c.svcCtx.Plugin != nil {
				chats := c.svcCtx.Plugin.GetChannelListByIdList(c.ctx, c.MD.UserId, channelIdList...)
				dialogsData.Chats = append(dialogsData.Chats, chats...)
			} else {
				c.Logger.Errorf("blocked, License key from https://teamgram.net required to unlock enterprise features.")
			}
		})

	return dialogsData
}
