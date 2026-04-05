// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/message/message"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesGetDialogs
// messages.getDialogs#a0f4cb4f flags:# exclude_pinned:flags.0?true folder_id:flags.1?int offset_date:int offset_id:int offset_peer:InputPeer limit:int hash:long = messages.Dialogs;
func (c *DialogsCore) MessagesGetDialogs(in *tg.TLMessagesGetDialogs) (*tg.MessagesDialogs, error) {
	var userId int64
	if c.MD != nil {
		userId = c.MD.UserId
	}

	// When DialogClient is wired, delegate to dialog service.
	if c.svcCtx != nil && c.svcCtx.DialogClient != nil && userId != 0 {
		folderId := int32(0)
		if in.FolderId != nil {
			folderId = *in.FolderId
		}

		var excludePinned tg.BoolClazz
		if in.ExcludePinned {
			excludePinned = tg.MakeTLBoolTrue(&tg.TLBoolTrue{})
		} else {
			excludePinned = tg.MakeTLBoolFalse(&tg.TLBoolFalse{})
		}

		dialogExts, err := c.svcCtx.DialogClient.DialogGetDialogs(c.ctx, &dialog.TLDialogGetDialogs{
			UserId:        userId,
			ExcludePinned: excludePinned,
			FolderId:      folderId,
		})
		if err != nil {
			c.Logger.Errorf("messages.getDialogs - DialogGetDialogs error: %v", err)
			return nil, err
		}

		dialogs := make([]tg.DialogClazz, 0, len(dialogExts.Datas))
		messages := make([]tg.MessageClazz, 0, len(dialogExts.Datas))
		for _, ext := range dialogExts.Datas {
			if ext != nil && ext.Dialog != nil {
				dialogs = append(dialogs, ext.Dialog)
				if fetched := c.fetchDialogTopMessage(userId, ext.Dialog); fetched != nil {
					messages = append(messages, fetched)
				}
			}
		}

		dialogData, err := c.svcCtx.DialogClient.DialogGetMyDialogsData(c.ctx, &dialog.TLDialogGetMyDialogsData{
			UserId:  userId,
			User:    true,
			Chat:    true,
			Channel: true,
		})
		if err != nil {
			c.Logger.Errorf("messages.getDialogs - DialogGetMyDialogsData error: %v", err)
			return nil, err
		}

		return tg.MakeTLMessagesDialogsSlice(&tg.TLMessagesDialogsSlice{
			Count:    int32(len(dialogs)),
			Dialogs:  dialogs,
			Messages: messages,
			Chats:    makeDialogChatsFromData(dialogData),
			Users:    makeDialogUsersFromData(dialogData),
		}).ToMessagesDialogs(), nil
	}

	// Fallback placeholder when DialogClient is not available.
	if in != nil && in.Limit > 0 {
		peer := tg.FromInputPeer2(userId, in.OffsetPeer)
		if peer.PeerType == tg.PEER_SELF || peer.PeerType == tg.PEER_USER {
			return tg.MakeTLMessagesDialogsSlice(&tg.TLMessagesDialogsSlice{
				Count: 1,
				Dialogs: []tg.DialogClazz{
					makePlaceholderDialog(peer.PeerId, 10),
				},
				Messages: []tg.MessageClazz{
					makePlaceholderDialogMessage(peer.PeerId, 10),
				},
				Chats: []tg.ChatClazz{},
				Users: []tg.UserClazz{
					makePlaceholderUser(peer.PeerId),
				},
			}).ToMessagesDialogs(), nil
		}
	}

	return tg.MakeTLMessagesDialogsSlice(&tg.TLMessagesDialogsSlice{
		Count:    0,
		Dialogs:  []tg.DialogClazz{},
		Messages: []tg.MessageClazz{},
		Chats:    []tg.ChatClazz{},
		Users:    []tg.UserClazz{},
	}).ToMessagesDialogs(), nil
}

func makeDialogUsersFromData(data *dialog.DialogsData) []tg.UserClazz {
	if data == nil || len(data.Users) == 0 {
		return []tg.UserClazz{}
	}

	users := make([]tg.UserClazz, 0, len(data.Users))
	for _, userID := range data.Users {
		users = append(users, makePlaceholderUser(userID))
	}
	return users
}

func makeDialogChatsFromData(data *dialog.DialogsData) []tg.ChatClazz {
	if data == nil || (len(data.Chats) == 0 && len(data.Channels) == 0) {
		return []tg.ChatClazz{}
	}

	chats := make([]tg.ChatClazz, 0, len(data.Chats)+len(data.Channels))
	for _, chatID := range data.Chats {
		chats = append(chats, tg.MakeTLChatEmpty(&tg.TLChatEmpty{Id: chatID}))
	}
	for _, channelID := range data.Channels {
		chats = append(chats, tg.MakeTLChannelForbidden(&tg.TLChannelForbidden{
			Id:         channelID,
			AccessHash: 0,
			Title:      "",
		}))
	}
	return chats
}

func (c *DialogsCore) fetchDialogTopMessage(userID int64, dialogItem tg.DialogClazz) tg.MessageClazz {
	if c == nil || c.svcCtx == nil || c.svcCtx.MessageClient == nil || userID == 0 || dialogItem == nil {
		return nil
	}

	peerType, peerID, ok := dialogPeerInfo(dialogItem)
	if !ok {
		return nil
	}
	topMessage := extractDialogTopMessage(dialogItem)
	if topMessage <= 0 {
		return nil
	}

	boxes, err := c.svcCtx.MessageClient.MessageGetHistoryMessages(c.ctx, &message.TLMessageGetHistoryMessages{
		UserId:   userID,
		PeerType: peerType,
		PeerId:   peerID,
		MaxId:    topMessage,
		Limit:    1,
	})
	if err != nil {
		c.Logger.Errorf("messages.getDialogs - MessageGetHistoryMessages error: %v", err)
		return nil
	}
	if boxes == nil || len(boxes.Datas) == 0 {
		return nil
	}
	if box := boxes.Datas[0]; box != nil && box.Message != nil {
		return box.Message
	}
	return nil
}

func dialogPeerInfo(dialogItem tg.DialogClazz) (int32, int64, bool) {
	dialogValue, ok := dialogItem.(*tg.TLDialog)
	if !ok || dialogValue.Peer == nil {
		return 0, 0, false
	}

	switch peer := dialogValue.Peer.(type) {
	case *tg.TLPeerUser:
		return tg.PEER_USER, peer.UserId, true
	case *tg.TLPeerChat:
		return tg.PEER_CHAT, peer.ChatId, true
	case *tg.TLPeerChannel:
		return tg.PEER_CHANNEL, peer.ChannelId, true
	default:
		return 0, 0, false
	}
}
