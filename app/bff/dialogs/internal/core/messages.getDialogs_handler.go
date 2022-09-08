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
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
	"sort"

	"github.com/teamgram/proto/mtproto"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"

	"github.com/zeromicro/go-zero/core/mr"
)

// MessagesGetDialogs
// messages.getDialogs#a0f4cb4f flags:# exclude_pinned:flags.0?true folder_id:flags.1?int offset_date:int offset_id:int offset_peer:InputPeer limit:int hash:long = messages.Dialogs;
func (c *DialogsCore) MessagesGetDialogs(in *mtproto.TLMessagesGetDialogs) (*mtproto.Messages_Dialogs, error) {
	var (
		offsetPeer         = mtproto.FromInputPeer2(c.MD.UserId, in.OffsetPeer)
		folderId           = in.GetFolderId().GetValue()
		limit              = in.Limit
		dialogExtList      dialog.DialogExtList
		notifySettingsList []*userpb.PeerPeerNotifySettings
	)

	if limit > 500 {
		limit = 500
	}

	mr.FinishVoid(
		func() {
			dialogs, err := c.svcCtx.Dao.DialogClient.DialogGetDialogs(c.ctx, &dialog.TLDialogGetDialogs{
				UserId:        c.MD.UserId,
				ExcludePinned: mtproto.ToBool(in.ExcludePinned),
				FolderId:      folderId,
			})
			if err != nil {
				c.Logger.Errorf("messages.getDialogs - error: %v", err)
			} else {
				dialogExtList = dialogs.GetDatas()
			}
		},
		func() {
			settingsList, err := c.svcCtx.Dao.UserClient.UserGetAllNotifySettings(c.ctx, &userpb.TLUserGetAllNotifySettings{
				UserId: c.MD.UserId,
			})
			if err != nil {
				c.Logger.Errorf("messages.getDialogs - error: %v", err)
			} else {
				notifySettingsList = settingsList.GetDatas()
			}
		})

	if len(dialogExtList) == 0 {
		return mtproto.MakeTLMessagesDialogsSlice(&mtproto.Messages_Dialogs{
			Dialogs:  []*mtproto.Dialog{},
			Messages: []*mtproto.Message{},
			Chats:    []*mtproto.Chat{},
			Users:    []*mtproto.User{},
			Count:    0,
		}).To_Messages_Dialogs(), nil
	}

	var (
		dialogCount = int32(dialogExtList.Len())
	)

	foundF := func(settingsList []*userpb.PeerPeerNotifySettings, peer *mtproto.PeerUtil) *mtproto.PeerNotifySettings {
		for _, s := range settingsList {
			if s.PeerType == peer.PeerType && s.PeerId == peer.PeerId {
				return s.Settings
			}
		}

		return mtproto.MakeTLPeerNotifySettings(nil).To_PeerNotifySettings()
	}

	for _, dialogEx := range dialogExtList {
		peer2 := mtproto.FromPeer(dialogEx.GetDialog().GetPeer())

		if peer2.IsChannel() {
			if c.svcCtx.Plugin != nil {
				dialog, _ := c.svcCtx.Plugin.GetChannelDialogById(c.ctx, c.MD.UserId, peer2.PeerId)
				if dialog != nil {
					dialogEx.Dialog.TopMessage = dialog.Dialog.TopMessage
					dialogEx.Dialog.Pts = dialog.Dialog.Pts
					// dialogEx.Dialog.ReadOutboxMaxId = channel2.Megagroup.ReadInboxMaxId
					dialogEx.Date = dialog.Date
					dialogEx.Order = dialog.Order
					// TODO:
					// dialogEx.AvailableMinId = megagroup2.GetParticipants()[0].AvailableMinId
					dialogEx.Dialog.UnreadCount = dialog.Dialog.TopMessage - dialogEx.Dialog.ReadInboxMaxId
					if dialog.Dialog.UnreadCount < 0 {
						dialog.Dialog.UnreadCount = 0
					}
				}
			} else {
				c.Logger.Errorf("messages.getDialogs blocked, License key from https://teamgram.net required to unlock enterprise features.")
			}
		}

		dialogEx.Dialog.NotifySettings = foundF(notifySettingsList, peer2)
	}

	r2 := sort.Reverse(dialogExtList)
	sort.Sort(r2)

	dialogExtList = dialogExtList.GetDialogsByOffsetLimit(
		in.OffsetDate,
		in.OffsetId,
		offsetPeer,
		in.Limit)

	messageDialogs := dialogExtList.DoGetMessagesDialogs(
		c.ctx,
		c.MD.UserId,
		func(ctx context.Context, selfUserId int64, id ...dialog.TopMessageId) []*mtproto.Message {
			var (
				msgList   = make([]*mtproto.Message, 0, len(id))
				msgIdList = make([]int32, 0, len(id))
			)
			for _, id2 := range id {
				if id2.Peer.IsChannel() {
					box, _ := c.svcCtx.Plugin.GetChannelMessage(c.ctx, c.MD.UserId, id2.Peer.PeerId, id2.TopMessage)
					if box != nil {
						msgList = append(msgList, box.ToMessage(c.MD.UserId))
					}
				} else {
					msgIdList = append(msgIdList, id2.TopMessage)
				}
			}
			if len(msgIdList) > 0 {
				boxList, _ := c.svcCtx.Dao.MessageClient.MessageGetUserMessageList(c.ctx, &message.TLMessageGetUserMessageList{
					UserId: c.MD.UserId,
					IdList: msgIdList,
				})
				boxList.Walk(func(idx int, v *mtproto.MessageBox) {
					msgList = append(msgList, v.ToMessage(c.MD.UserId))
				})
			}

			return msgList
		},
		func(ctx context.Context, selfUserId int64, id ...int64) []*mtproto.User {
			users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: id,
				})

			return users.GetUserListByIdList(c.MD.UserId, id...)
		},
		func(ctx context.Context, selfUserId int64, id ...int64) []*mtproto.Chat {
			chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: id,
				})

			return chats.GetChatListByIdList(c.MD.UserId, id...)
		},
		func(ctx context.Context, selfUserId int64, id ...int64) []*mtproto.Chat {
			var (
				chats []*mtproto.Chat
			)
			if c.svcCtx.Plugin != nil {
				chats = c.svcCtx.Plugin.GetChannelListByIdList(c.ctx, c.MD.UserId, id...)
			} else {
				c.Logger.Errorf("blocked, License key from https://teamgram.net required to unlock enterprise features.")
			}

			return chats
		})

	return messageDialogs.ToMessagesDialogs(dialogCount), nil
}
