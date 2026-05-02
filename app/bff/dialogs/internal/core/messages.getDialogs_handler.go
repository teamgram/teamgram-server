// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
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
	"sort"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	messagepb "github.com/teamgram/teamgram-server/v2/app/service/biz/message/message"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesGetDialogs
// messages.getDialogs#a0f4cb4f flags:# exclude_pinned:flags.0?true folder_id:flags.1?int offset_date:int offset_id:int offset_peer:InputPeer limit:int hash:long = messages.Dialogs;
func (c *DialogsCore) MessagesGetDialogs(in *tg.TLMessagesGetDialogs) (*tg.MessagesDialogs, error) {
	folderID := int32(0)
	if in.FolderId != nil {
		folderID = *in.FolderId
	}

	dialogs, err := c.svcCtx.Repo.DialogClient.DialogGetDialogs(c.ctx, &dialogpb.TLDialogGetDialogs{
		UserId:        c.MD.UserId,
		ExcludePinned: tg.ToBoolClazz(in.ExcludePinned),
		FolderId:      folderID,
	})
	if err != nil {
		c.Logger.Errorf("messages.getDialogs - dialog.getDialogs failed: user_id: %d, exclude_pinned: %t, folder_id: %d, err: %v",
			c.MD.UserId, in.ExcludePinned, folderID, err)
		return nil, tg.ErrInternalServerError
	}

	dialogExts := make([]dialogpb.DialogExtClazz, 0)
	if dialogs != nil {
		dialogExts = append(dialogExts, dialogs.Datas...)
	}

	sort.SliceStable(dialogExts, func(i, j int) bool {
		if dialogExts[i].Order == dialogExts[j].Order {
			return dialogExts[i].Date > dialogExts[j].Date
		}
		return dialogExts[i].Order > dialogExts[j].Order
	})

	totalCount := int32(len(dialogExts))
	dialogExts = limitDialogExts(offsetDialogExts(dialogExts, in, c.MD.UserId), in.Limit)

	if len(dialogExts) == 0 {
		return tg.MakeTLMessagesDialogsSlice(&tg.TLMessagesDialogsSlice{
			Count:    totalCount,
			Dialogs:  []tg.DialogClazz{},
			Messages: []tg.MessageClazz{},
			Chats:    []tg.ChatClazz{},
			Users:    []tg.UserClazz{},
		}).ToMessagesDialogs(), nil
	}

	dialogList := make([]tg.DialogClazz, 0, len(dialogExts))
	userIDs := make([]int64, 0, len(dialogExts))
	chatIDs := make([]int64, 0, len(dialogExts))
	topMessageIDs := make([]int32, 0, len(dialogExts))
	for _, dialogExt := range dialogExts {
		if dialogExt == nil || dialogExt.Dialog == nil {
			continue
		}
		dialogList = append(dialogList, dialogExt.Dialog)
		dialog, ok := (&tg.Dialog{Clazz: dialogExt.Dialog}).ToDialog()
		if !ok {
			continue
		}
		if dialog.TopMessage > 0 {
			topMessageIDs = append(topMessageIDs, dialog.TopMessage)
		}
		switch peer := dialog.Peer.(type) {
		case *tg.TLPeerUser:
			userIDs = append(userIDs, peer.UserId)
		case *tg.TLPeerChat:
			chatIDs = append(chatIDs, peer.ChatId)
		}
	}

	messages, err := c.fetchDialogTopMessages(topMessageIDs)
	if err != nil {
		return nil, err
	}
	users, err := c.fetchDialogUsers(userIDs)
	if err != nil {
		return nil, err
	}
	chats, err := c.fetchDialogChats(chatIDs)
	if err != nil {
		return nil, err
	}

	return tg.MakeTLMessagesDialogsSlice(&tg.TLMessagesDialogsSlice{
		Count:    totalCount,
		Dialogs:  dialogList,
		Messages: messages,
		Chats:    chats,
		Users:    users,
	}).ToMessagesDialogs(), nil
}

func offsetDialogExts(dialogs []dialogpb.DialogExtClazz, in *tg.TLMessagesGetDialogs, selfID int64) []dialogpb.DialogExtClazz {
	if in.OffsetDate == 0 && in.OffsetId == 0 && isInputPeerEmpty(in.OffsetPeer) {
		return dialogs
	}

	offsetPeer := tg.FromInputPeer2(selfID, in.OffsetPeer)
	for i, dialogExt := range dialogs {
		if dialogExt == nil || dialogExt.Dialog == nil {
			continue
		}
		dialog, ok := (&tg.Dialog{Clazz: dialogExt.Dialog}).ToDialog()
		if !ok {
			continue
		}
		if in.OffsetId != 0 && dialog.TopMessage == in.OffsetId {
			return dialogs[i+1:]
		}
		if in.OffsetDate != 0 && int32(dialogExt.Date) == in.OffsetDate {
			return dialogs[i+1:]
		}
		if !isInputPeerEmpty(in.OffsetPeer) && samePeer(dialog.Peer, offsetPeer) {
			return dialogs[i+1:]
		}
	}
	return dialogs
}

func limitDialogExts(dialogs []dialogpb.DialogExtClazz, limit int32) []dialogpb.DialogExtClazz {
	if limit <= 0 {
		return []dialogpb.DialogExtClazz{}
	}
	if limit > 500 {
		limit = 500
	}
	if int(limit) >= len(dialogs) {
		return dialogs
	}
	return dialogs[:limit]
}

func isInputPeerEmpty(peer tg.InputPeerClazz) bool {
	_, ok := peer.(*tg.TLInputPeerEmpty)
	return peer == nil || ok
}

func samePeer(peer tg.PeerClazz, util tg.PeerUtilClazz) bool {
	if util == nil {
		return false
	}
	switch p := peer.(type) {
	case *tg.TLPeerUser:
		return (util.PeerType == tg.PEER_USER && p.UserId == util.PeerId) ||
			(util.PeerType == tg.PEER_SELF && p.UserId == util.SelfId)
	case *tg.TLPeerChat:
		return util.PeerType == tg.PEER_CHAT && p.ChatId == util.PeerId
	case *tg.TLPeerChannel:
		return util.PeerType == tg.PEER_CHANNEL && p.ChannelId == util.PeerId
	default:
		return false
	}
}

func (c *DialogsCore) fetchDialogTopMessages(ids []int32) ([]tg.MessageClazz, error) {
	ids = uniqueInt32s(ids)
	if len(ids) == 0 {
		return []tg.MessageClazz{}, nil
	}

	boxes, err := c.svcCtx.Repo.MessageClient.MessageGetUserMessageList(c.ctx, &messagepb.TLMessageGetUserMessageList{
		UserId: c.MD.UserId,
		IdList: ids,
	})
	if err != nil {
		c.Logger.Errorf("messages.getDialogs - message.getUserMessageList failed: user_id: %d, id_list: %v, err: %v", c.MD.UserId, ids, err)
		return nil, tg.ErrInternalServerError
	}

	messages := make([]tg.MessageClazz, 0, len(ids))
	if boxes == nil {
		return messages, nil
	}
	for _, box := range boxes.Datas {
		if box != nil && box.Message != nil {
			messages = append(messages, box.Message)
		}
	}
	return messages, nil
}

func (c *DialogsCore) fetchDialogUsers(ids []int64) ([]tg.UserClazz, error) {
	ids = uniqueInt64s(ids)
	if len(ids) == 0 {
		return []tg.UserClazz{}, nil
	}

	users, err := c.svcCtx.Repo.UserClient.UserGetMutableUsersV2(c.ctx, &userpb.TLUserGetMutableUsersV2{
		Id:      ids,
		Privacy: true,
		HasTo:   true,
		To:      []int64{c.MD.UserId},
	})
	if err != nil {
		c.Logger.Errorf("messages.getDialogs - user.getMutableUsersV2 failed: user_id: %d, id: %v, err: %v", c.MD.UserId, ids, err)
		return nil, tg.ErrInternalServerError
	}

	byID := make(map[int64]tg.UserClazz, len(ids))
	if users != nil {
		for _, immutableUser := range users.Users {
			user := projectImmutableUser(immutableUser)
			if id, ok := userID(user); ok {
				byID[id] = user
			}
		}
	}

	out := make([]tg.UserClazz, 0, len(ids))
	for _, id := range ids {
		if user := byID[id]; user != nil {
			out = append(out, user)
			continue
		}
		out = append(out, tg.MakeTLUserEmpty(&tg.TLUserEmpty{Id: id}))
	}
	return out, nil
}

func (c *DialogsCore) fetchDialogChats(ids []int64) ([]tg.ChatClazz, error) {
	ids = uniqueInt64s(ids)
	if len(ids) == 0 {
		return []tg.ChatClazz{}, nil
	}

	chats, err := c.svcCtx.Repo.ChatClient.ChatGetChatListByIdList(c.ctx, &chatpb.TLChatGetChatListByIdList{
		SelfId: c.MD.UserId,
		IdList: ids,
	})
	if err != nil {
		c.Logger.Errorf("messages.getDialogs - chat.getChatListByIdList failed: user_id: %d, id_list: %v, err: %v", c.MD.UserId, ids, err)
		return nil, tg.ErrInternalServerError
	}

	out := make([]tg.ChatClazz, 0, len(ids))
	if chats == nil {
		return out, nil
	}
	for _, mutableChat := range chats.Datas {
		if chat := projectMutableChat(mutableChat, c.MD.UserId); chat != nil {
			out = append(out, chat)
		}
	}
	return out, nil
}

func uniqueInt32s(ids []int32) []int32 {
	seen := make(map[int32]struct{}, len(ids))
	out := make([]int32, 0, len(ids))
	for _, id := range ids {
		if id == 0 {
			continue
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		out = append(out, id)
	}
	return out
}

func uniqueInt64s(ids []int64) []int64 {
	seen := make(map[int64]struct{}, len(ids))
	out := make([]int64, 0, len(ids))
	for _, id := range ids {
		if id == 0 {
			continue
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		out = append(out, id)
	}
	return out
}
