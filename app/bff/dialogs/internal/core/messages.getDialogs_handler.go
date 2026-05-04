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
	"strconv"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
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
		if totalCount == 0 && c.canFallbackToCanonicalDialogs(in, folderID) {
			fallback, err := c.fetchCanonicalDialogsFromDifference("messages.getDialogs", in.Limit)
			if err != nil {
				return nil, err
			}
			if fallback != nil && len(fallback.Dialogs) > 0 {
				return tg.MakeTLMessagesDialogsSlice(&tg.TLMessagesDialogsSlice{
					Count:    int32(len(fallback.Dialogs)),
					Dialogs:  fallback.Dialogs,
					Messages: fallback.Messages,
					Chats:    []tg.ChatClazz{},
					Users:    fallback.Users,
				}).ToMessagesDialogs(), nil
			}
		}
		if totalCount == 0 && c.canFallbackToCanonicalSelfDialog(in, folderID) {
			fallback, err := c.fetchCanonicalUserDialog("messages.getDialogs", c.MD.UserId)
			if err != nil {
				return nil, err
			}
			if fallback != nil {
				return tg.MakeTLMessagesDialogsSlice(&tg.TLMessagesDialogsSlice{
					Count:    1,
					Dialogs:  []tg.DialogClazz{fallback.Dialog},
					Messages: fallback.Messages,
					Chats:    []tg.ChatClazz{},
					Users:    fallback.Users,
				}).ToMessagesDialogs(), nil
			}
		}
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

	if c.canFallbackToCanonicalDialogs(in, folderID) {
		fallback, err := c.fetchCanonicalDialogsFromDifference("messages.getDialogs", in.Limit)
		if err != nil {
			return nil, err
		}
		dialogList, messages, chats, users, totalCount = mergeCanonicalDialogs(
			dialogList,
			messages,
			chats,
			users,
			fallback,
			totalCount,
			in.Limit,
		)
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
	return c.fetchDialogUsersWithOperation("messages.getDialogs", ids)
}

func (c *DialogsCore) fetchDialogUsersWithOperation(operation string, ids []int64) ([]tg.UserClazz, error) {
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
		c.Logger.Errorf("%s - user.getMutableUsersV2 failed: user_id: %d, id: %v, err: %v", operation, c.MD.UserId, ids, err)
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

type canonicalDialogResult struct {
	Dialog   tg.DialogClazz
	Messages []tg.MessageClazz
	Users    []tg.UserClazz
}

type canonicalDialogsResult struct {
	Dialogs  []tg.DialogClazz
	Messages []tg.MessageClazz
	Users    []tg.UserClazz
}

func (c *DialogsCore) canFallbackToCanonicalDialogs(in *tg.TLMessagesGetDialogs, folderID int32) bool {
	return in != nil &&
		in.Limit > 0 &&
		folderID == 0 &&
		in.OffsetDate == 0 &&
		in.OffsetId == 0 &&
		isInputPeerEmpty(in.OffsetPeer)
}

func (c *DialogsCore) canFallbackToCanonicalSelfDialog(in *tg.TLMessagesGetDialogs, folderID int32) bool {
	return in != nil &&
		in.Limit > 0 &&
		folderID == 0 &&
		in.OffsetDate == 0 &&
		in.OffsetId == 0 &&
		isInputPeerEmpty(in.OffsetPeer)
}

func (c *DialogsCore) fetchCanonicalDialogsFromDifference(operation string, limit int32) (*canonicalDialogsResult, error) {
	if c.svcCtx == nil || c.svcCtx.Repo == nil || c.svcCtx.Repo.UserupdatesClient == nil {
		return nil, nil
	}
	if limit <= 0 {
		return nil, nil
	}
	if limit > 100 {
		limit = 100
	}

	diff, err := c.svcCtx.Repo.UserupdatesClient.UserupdatesGetDifference(c.ctx, &userupdates.TLUserupdatesGetDifference{
		UserId:        c.MD.UserId,
		AuthKeyId:     c.MD.PermAuthKeyId,
		Pts:           0,
		PtsTotalLimit: &limit,
	})
	if err != nil {
		return nil, tg.ErrInternalServerError
	}

	messages := messagesFromUserDifference(diff)
	if len(messages) == 0 {
		return nil, nil
	}

	type dialogProjection struct {
		peerUserID  int64
		topMessage  tg.MessageClazz
		topID       int32
		topDate     int32
		unreadCount int32
	}
	byPeer := make(map[int64]*dialogProjection)
	for _, message := range messages {
		peerUserID, ok := messagePeerUserID(message)
		if !ok || peerUserID <= 0 {
			continue
		}
		id, ok := messageID(message)
		if !ok {
			continue
		}
		date := messageDate(message)
		projection := byPeer[peerUserID]
		if projection == nil {
			projection = &dialogProjection{peerUserID: peerUserID}
			byPeer[peerUserID] = projection
		}
		if !messageOutgoing(message) {
			projection.unreadCount++
		}
		if projection.topMessage == nil || date > projection.topDate || (date == projection.topDate && id > projection.topID) {
			projection.topMessage = message
			projection.topID = id
			projection.topDate = date
		}
	}
	if len(byPeer) == 0 {
		return nil, nil
	}

	dialogs := make([]*dialogProjection, 0, len(byPeer))
	for _, projection := range byPeer {
		dialogs = append(dialogs, projection)
	}
	sort.SliceStable(dialogs, func(i, j int) bool {
		if dialogs[i].topDate == dialogs[j].topDate {
			return dialogs[i].topID > dialogs[j].topID
		}
		return dialogs[i].topDate > dialogs[j].topDate
	})

	if int(limit) < len(dialogs) {
		dialogs = dialogs[:limit]
	}

	dialogList := make([]tg.DialogClazz, 0, len(dialogs))
	topMessages := make([]tg.MessageClazz, 0, len(dialogs))
	userIDs := make([]int64, 0, len(dialogs))
	for _, projection := range dialogs {
		dialogList = append(dialogList, tg.MakeTLDialog(&tg.TLDialog{
			Peer:           tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: projection.peerUserID}),
			TopMessage:     projection.topID,
			UnreadCount:    projection.unreadCount,
			NotifySettings: tg.MakeTLPeerNotifySettings(&tg.TLPeerNotifySettings{}),
		}))
		topMessages = append(topMessages, projection.topMessage)
		userIDs = append(userIDs, projection.peerUserID)
	}
	users, err := c.fetchDialogUsersWithOperation(operation, userIDs)
	if err != nil {
		return nil, err
	}

	return &canonicalDialogsResult{
		Dialogs:  dialogList,
		Messages: topMessages,
		Users:    users,
	}, nil
}

func (c *DialogsCore) fetchCanonicalUserDialog(operation string, peerUserID int64) (*canonicalDialogResult, error) {
	if c.svcCtx == nil || c.svcCtx.Repo == nil || c.svcCtx.Repo.MsgClient == nil {
		return nil, nil
	}

	r, err := c.svcCtx.Repo.MsgClient.MsgGetHistory(c.ctx, &msg.TLMsgGetHistory{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.PermAuthKeyId,
		PeerType:  payload.PeerTypeUser,
		PeerId:    peerUserID,
		Limit:     1,
	})
	if err != nil {
		c.Logger.Errorf("%s - msg.getHistory failed: user_id: %d, peer_id: %d, err: %v", operation, c.MD.UserId, peerUserID, err)
		return nil, tg.ErrInternalServerError
	}

	messages := messagesFromHistoryResult(r)
	if len(messages) == 0 {
		return nil, nil
	}
	topMessageID, ok := messageID(messages[0])
	if !ok {
		return nil, nil
	}

	users, err := c.fetchDialogUsersWithOperation(operation, []int64{peerUserID})
	if err != nil {
		return nil, err
	}

	return &canonicalDialogResult{
		Dialog: tg.MakeTLDialog(&tg.TLDialog{
			Peer:            tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: peerUserID}),
			TopMessage:      topMessageID,
			ReadInboxMaxId:  0,
			ReadOutboxMaxId: 0,
			NotifySettings:  tg.MakeTLPeerNotifySettings(&tg.TLPeerNotifySettings{}),
		}),
		Messages: messages,
		Users:    users,
	}, nil
}

func messagesFromHistoryResult(r *tg.MessagesMessages) []tg.MessageClazz {
	if r == nil {
		return []tg.MessageClazz{}
	}
	if messages, ok := r.ToMessagesMessages(); ok {
		return messages.Messages
	}
	if messages, ok := r.ToMessagesMessagesSlice(); ok {
		return messages.Messages
	}
	return []tg.MessageClazz{}
}

func messagesFromUserDifference(diff *userupdates.UserDifference) []tg.MessageClazz {
	if diff == nil {
		return []tg.MessageClazz{}
	}
	if d, ok := diff.ToUserDifference(); ok {
		return d.NewMessages
	}
	if d, ok := diff.ToUserDifferenceSlice(); ok {
		return d.NewMessages
	}
	return []tg.MessageClazz{}
}

func mergeCanonicalDialogs(
	dialogs []tg.DialogClazz,
	messages []tg.MessageClazz,
	chats []tg.ChatClazz,
	users []tg.UserClazz,
	fallback *canonicalDialogsResult,
	totalCount int32,
	limit int32,
) ([]tg.DialogClazz, []tg.MessageClazz, []tg.ChatClazz, []tg.UserClazz, int32) {
	if fallback == nil || len(fallback.Dialogs) == 0 {
		return dialogs, messages, chats, users, totalCount
	}

	peerSeen := make(map[string]struct{}, len(dialogs)+len(fallback.Dialogs))
	for _, dialogClazz := range dialogs {
		if dialog, ok := (&tg.Dialog{Clazz: dialogClazz}).ToDialog(); ok {
			if key, ok := peerKey(dialog.Peer); ok {
				peerSeen[key] = struct{}{}
			}
		}
	}

	missingTopIDs := make(map[int32]struct{}, len(fallback.Dialogs))
	for i, dialogClazz := range fallback.Dialogs {
		dialog, ok := (&tg.Dialog{Clazz: dialogClazz}).ToDialog()
		if !ok {
			continue
		}
		key, ok := peerKey(dialog.Peer)
		if !ok {
			continue
		}
		if _, ok := peerSeen[key]; ok {
			continue
		}
		peerSeen[key] = struct{}{}
		dialogs = append(dialogs, dialogClazz)
		if dialog.TopMessage > 0 {
			missingTopIDs[dialog.TopMessage] = struct{}{}
		}
		if i < len(fallback.Messages) {
			if id, ok := messageID(fallback.Messages[i]); ok && id > 0 {
				missingTopIDs[id] = struct{}{}
			}
		}
		totalCount++
	}
	if len(missingTopIDs) == 0 {
		return limitMergedDialogs(dialogs, messages, chats, users, totalCount, limit)
	}

	messageSeen := make(map[int32]struct{}, len(messages)+len(fallback.Messages))
	for _, message := range messages {
		if id, ok := messageID(message); ok && id > 0 {
			messageSeen[id] = struct{}{}
		}
	}
	for _, message := range fallback.Messages {
		id, ok := messageID(message)
		if !ok || id <= 0 {
			continue
		}
		if _, wanted := missingTopIDs[id]; !wanted {
			continue
		}
		if _, ok := messageSeen[id]; ok {
			continue
		}
		messageSeen[id] = struct{}{}
		messages = append(messages, message)
	}

	userSeen := make(map[int64]struct{}, len(users)+len(fallback.Users))
	for _, user := range users {
		if id, ok := userID(user); ok && id > 0 {
			userSeen[id] = struct{}{}
		}
	}
	for _, user := range fallback.Users {
		id, ok := userID(user)
		if !ok || id <= 0 {
			continue
		}
		if _, ok := userSeen[id]; ok {
			continue
		}
		userSeen[id] = struct{}{}
		users = append(users, user)
	}

	return limitMergedDialogs(dialogs, messages, chats, users, totalCount, limit)
}

func limitMergedDialogs(
	dialogs []tg.DialogClazz,
	messages []tg.MessageClazz,
	chats []tg.ChatClazz,
	users []tg.UserClazz,
	totalCount int32,
	limit int32,
) ([]tg.DialogClazz, []tg.MessageClazz, []tg.ChatClazz, []tg.UserClazz, int32) {
	messageByID := make(map[int32]tg.MessageClazz, len(messages))
	for _, message := range messages {
		if id, ok := messageID(message); ok && id > 0 {
			messageByID[id] = message
		}
	}
	sort.SliceStable(dialogs, func(i, j int) bool {
		leftID, leftDate := dialogTop(dialogs[i], messageByID)
		rightID, rightDate := dialogTop(dialogs[j], messageByID)
		if leftDate == rightDate {
			return leftID > rightID
		}
		return leftDate > rightDate
	})
	if limit <= 0 {
		return []tg.DialogClazz{}, []tg.MessageClazz{}, []tg.ChatClazz{}, []tg.UserClazz{}, totalCount
	}
	if int(limit) < len(dialogs) {
		dialogs = dialogs[:limit]
	}

	topIDs := make(map[int32]struct{}, len(dialogs))
	userIDs := make(map[int64]struct{}, len(dialogs))
	chatIDs := make(map[int64]struct{}, len(dialogs))
	for _, dialogClazz := range dialogs {
		dialog, ok := (&tg.Dialog{Clazz: dialogClazz}).ToDialog()
		if !ok {
			continue
		}
		if dialog.TopMessage > 0 {
			topIDs[dialog.TopMessage] = struct{}{}
		}
		switch peer := dialog.Peer.(type) {
		case *tg.TLPeerUser:
			userIDs[peer.UserId] = struct{}{}
		case *tg.TLPeerChat:
			chatIDs[peer.ChatId] = struct{}{}
		}
	}

	filteredMessages := make([]tg.MessageClazz, 0, len(messages))
	for _, message := range messages {
		if id, ok := messageID(message); ok {
			if _, wanted := topIDs[id]; wanted {
				filteredMessages = append(filteredMessages, message)
			}
		}
	}
	filteredUsers := make([]tg.UserClazz, 0, len(users))
	for _, user := range users {
		if id, ok := userID(user); ok {
			if _, wanted := userIDs[id]; wanted {
				filteredUsers = append(filteredUsers, user)
			}
		}
	}
	filteredChats := make([]tg.ChatClazz, 0, len(chats))
	for _, chat := range chats {
		if id, ok := chatID(chat); ok {
			if _, wanted := chatIDs[id]; wanted {
				filteredChats = append(filteredChats, chat)
			}
		}
	}

	return dialogs, filteredMessages, filteredChats, filteredUsers, totalCount
}

func dialogTop(dialogClazz tg.DialogClazz, messageByID map[int32]tg.MessageClazz) (int32, int32) {
	dialog, ok := (&tg.Dialog{Clazz: dialogClazz}).ToDialog()
	if !ok {
		return 0, 0
	}
	if message, ok := messageByID[dialog.TopMessage]; ok {
		return dialog.TopMessage, messageDate(message)
	}
	return dialog.TopMessage, 0
}

func peerKey(peer tg.PeerClazz) (string, bool) {
	switch p := peer.(type) {
	case *tg.TLPeerUser:
		return "u:" + strconv.FormatInt(p.UserId, 10), p.UserId > 0
	case *tg.TLPeerChat:
		return "c:" + strconv.FormatInt(p.ChatId, 10), p.ChatId > 0
	case *tg.TLPeerChannel:
		return "n:" + strconv.FormatInt(p.ChannelId, 10), p.ChannelId > 0
	default:
		return "", false
	}
}

func chatID(chat tg.ChatClazz) (int64, bool) {
	switch c := chat.(type) {
	case *tg.TLChat:
		return c.Id, c.Id > 0
	case *tg.TLChatForbidden:
		return c.Id, c.Id > 0
	case *tg.TLChannel:
		return c.Id, c.Id > 0
	case *tg.TLChannelForbidden:
		return c.Id, c.Id > 0
	default:
		return 0, false
	}
}

func messageID(message tg.MessageClazz) (int32, bool) {
	switch m := message.(type) {
	case *tg.TLMessage:
		return m.Id, m.Id > 0
	case *tg.TLMessageService:
		return m.Id, m.Id > 0
	default:
		return 0, false
	}
}

func messageDate(message tg.MessageClazz) int32 {
	switch m := message.(type) {
	case *tg.TLMessage:
		return m.Date
	case *tg.TLMessageService:
		return m.Date
	default:
		return 0
	}
}

func messageOutgoing(message tg.MessageClazz) bool {
	switch m := message.(type) {
	case *tg.TLMessage:
		return m.Out
	case *tg.TLMessageService:
		return m.Out
	default:
		return false
	}
}

func messagePeerUserID(message tg.MessageClazz) (int64, bool) {
	switch m := message.(type) {
	case *tg.TLMessage:
		if peer, ok := m.PeerId.(*tg.TLPeerUser); ok {
			return peer.UserId, peer.UserId > 0
		}
	case *tg.TLMessageService:
		if peer, ok := m.PeerId.(*tg.TLPeerUser); ok {
			return peer.UserId, peer.UserId > 0
		}
	}
	return 0, false
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
