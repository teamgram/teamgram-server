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
	"encoding/json"
	"math"
	"strconv"

	userprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	msgpb "github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
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

	cursor, err := c.dialogCursorFromGetDialogs(in, folderID)
	if err != nil {
		return nil, err
	}
	if in.OffsetId > 0 {
		resolved, err := c.resolveDialogCursorTopMessage(in.OffsetId)
		if err != nil {
			return nil, err
		}
		if resolved == nil {
			return emptyMessagesDialogsSlice(0), nil
		}
		if cursorData := cursor.ToDialogCursor(); cursorData != nil {
			cursorData.TopPeerSeq = resolved.PeerSeq
			cursorData.TopMessageDate = resolved.MessageDate
			cursorData.PeerType = resolved.PeerType
			cursorData.PeerId = resolved.PeerId
		}
	}
	limit := normalizeDialogLimit(in.Limit)
	page, err := c.svcCtx.Repo.DialogClient.DialogGetDialogsV2(c.ctx, &dialogpb.TLDialogGetDialogsV2{
		UserId:        c.MD.UserId,
		Cursor:        cursor,
		ExcludePinned: tg.ToBoolClazz(in.ExcludePinned),
		Limit:         limit,
	})
	if err != nil {
		c.Logger.Errorf("messages.getDialogs - dialog.getDialogsV2 failed: user_id: %d, exclude_pinned: %t, folder_id: %d, err: %v",
			c.MD.UserId, in.ExcludePinned, folderID, err)
		return nil, tg.ErrInternalServerError
	}

	dialogExts := vectorDialogExtV2s(page)
	totalCount := int32(len(dialogExts))
	hydrated, err := c.hydrateDialogExtV2s("messages.getDialogs", dialogExts)
	if err != nil {
		return nil, err
	}
	if len(hydrated.Dialogs) == 0 {
		return emptyMessagesDialogsSlice(totalCount), nil
	}

	return tg.MakeTLMessagesDialogsSlice(&tg.TLMessagesDialogsSlice{
		Count:    totalCount,
		Dialogs:  hydrated.Dialogs,
		Messages: hydrated.Messages,
		Chats:    hydrated.Chats,
		Users:    hydrated.Users,
	}).ToMessagesDialogs(), nil
}

func emptyMessagesDialogsSlice(count int32) *tg.MessagesDialogs {
	return tg.MakeTLMessagesDialogsSlice(&tg.TLMessagesDialogsSlice{
		Count:    count,
		Dialogs:  []tg.DialogClazz{},
		Messages: []tg.MessageClazz{},
		Chats:    []tg.ChatClazz{},
		Users:    []tg.UserClazz{},
	}).ToMessagesDialogs()
}

func normalizeDialogLimit(limit int32) int32 {
	if limit <= 0 {
		return 0
	}
	if limit > 500 {
		return 500
	}
	return limit
}

func isInputPeerEmpty(peer tg.InputPeerClazz) bool {
	_, ok := peer.(*tg.TLInputPeerEmpty)
	return peer == nil || ok
}

type hydratedDialogV2Result struct {
	Dialogs  []tg.DialogClazz
	Messages []tg.MessageClazz
	Chats    []tg.ChatClazz
	Users    []tg.UserClazz
}

type dialogTopMessageRef struct {
	PeerType int32
	PeerID   int64
	PeerSeq  int64
}

type dialogReadState struct {
	ReadInboxMaxUserMessageID  int64
	ReadOutboxMaxUserMessageID int64
}

func vectorDialogExtV2s(page *dialogpb.DialogPage) []dialogpb.DialogExtV2Clazz {
	if page == nil {
		return nil
	}
	return page.Dialogs
}

func vectorDialogExtV2Datas(dialogs *dialogpb.VectorDialogExtV2) []dialogpb.DialogExtV2Clazz {
	if dialogs == nil {
		return nil
	}
	return dialogs.Datas
}

func (c *DialogsCore) dialogCursorFromGetDialogs(in *tg.TLMessagesGetDialogs, folderID int32) (dialogpb.DialogCursorClazz, error) {
	cursor := &dialogpb.TLDialogCursor{
		FolderId:       folderID,
		Section:        "",
		TopMessageDate: int64(in.OffsetDate),
		TopPeerSeq:     int64(in.OffsetId),
	}
	if !isInputPeerEmpty(in.OffsetPeer) {
		resolved := tg.FromInputPeer2(c.MD.UserId, in.OffsetPeer)
		peerType, err := dialogFacadePeerType(resolved.PeerType)
		if err != nil {
			return nil, err
		}
		if resolved.PeerId <= 0 {
			return nil, tg.Err400PeerIdInvalid
		}
		cursor.PeerType = peerType
		cursor.PeerId = resolved.PeerId
	}
	return dialogpb.MakeTLDialogCursor(cursor), nil
}

func (c *DialogsCore) hydrateDialogExtV2s(operation string, dialogExts []dialogpb.DialogExtV2Clazz) (*hydratedDialogV2Result, error) {
	result := &hydratedDialogV2Result{
		Dialogs:  []tg.DialogClazz{},
		Messages: []tg.MessageClazz{},
		Chats:    []tg.ChatClazz{},
		Users:    []tg.UserClazz{},
	}
	if len(dialogExts) == 0 {
		return result, nil
	}

	dialogs := make([]tg.DialogClazz, 0, len(dialogExts))
	userIDs := make([]int64, 0, len(dialogExts))
	chatIDs := make([]int64, 0, len(dialogExts))
	topMessageRefs := make([]dialogTopMessageRef, 0, len(dialogExts))
	notifyPeers := make([]tg.PeerUtilClazz, 0, len(dialogExts))
	notifyKeys := make([]string, 0, len(dialogExts))
	projectionPeers := make([]userupdates.DialogProjectionPeerClazz, 0, len(dialogExts))
	projectionKeys := make([]string, 0, len(dialogExts))

	for _, dialogExt := range dialogExts {
		if dialogExt == nil {
			continue
		}
		topMessageID, err := int64ToDialogMessageID(dialogExt.TopUserMessageId)
		if err != nil {
			c.Logger.Errorf("%s - invalid top user message id: user_id: %d, peer_type: %d, peer_id: %d, top_user_message_id: %d, top_peer_seq: %d, err: %v",
				operation, c.MD.UserId, dialogExt.PeerType, dialogExt.PeerId, dialogExt.TopUserMessageId, dialogExt.TopPeerSeq, err)
			return nil, tg.ErrInternalServerError
		}
		notifyPeer, ok := notifyPeerFromDialogFacade(c.MD.UserId, dialogExt.PeerType, dialogExt.PeerId)
		if !ok {
			return nil, tg.Err400PeerIdInvalid
		}
		notifyPeers = append(notifyPeers, notifyPeer)
		notifyKeys = append(notifyKeys, notifySettingsKey(notifyPeer.PeerType, notifyPeer.PeerId))
		projectionPeers = append(projectionPeers, userupdates.MakeTLDialogProjectionPeer(&userupdates.TLDialogProjectionPeer{
			PeerType: dialogExt.PeerType,
			PeerId:   dialogExt.PeerId,
		}))
		projectionKeys = append(projectionKeys, dialogProjectionKey(dialogExt.PeerType, dialogExt.PeerId))

		dialog, err := makePublicDialogFromExtV2(dialogExt, topMessageID, tg.MakeTLPeerNotifySettings(&tg.TLPeerNotifySettings{}))
		if err != nil {
			c.Logger.Errorf("%s - invalid dialog draft payload: user_id: %d, peer_type: %d, peer_id: %d, err: %v",
				operation, c.MD.UserId, dialogExt.PeerType, dialogExt.PeerId, err)
			return nil, tg.ErrInternalServerError
		}
		dialogs = append(dialogs, dialog)
		if topMessageID > 0 {
			topMessageRefs = append(topMessageRefs, dialogTopMessageRef{
				PeerType: dialogExt.PeerType,
				PeerID:   dialogExt.PeerId,
				PeerSeq:  dialogExt.TopPeerSeq,
			})
		}
		switch dialogExt.PeerType {
		case dialogPeerTypeUser:
			userIDs = append(userIDs, dialogExt.PeerId)
		case dialogPeerTypeChat, dialogPeerTypeChannel:
			chatIDs = append(chatIDs, dialogExt.PeerId)
		}
	}

	readStates, err := c.fetchDialogReadStates(operation, projectionPeers)
	if err != nil {
		return nil, err
	}
	for i, dialogClazz := range dialogs {
		dialog, ok := (&tg.Dialog{Clazz: dialogClazz}).ToDialog()
		if !ok {
			continue
		}
		if state, ok := readStates[projectionKeys[i]]; ok {
			dialog.ReadInboxMaxId = int64ToInt32Saturating(state.ReadInboxMaxUserMessageID)
			dialog.ReadOutboxMaxId = int64ToInt32Saturating(state.ReadOutboxMaxUserMessageID)
		}
	}

	notifySettings, err := c.fetchDialogNotifySettings(operation, notifyPeers)
	if err != nil {
		return nil, err
	}
	for i, dialogClazz := range dialogs {
		dialog, ok := (&tg.Dialog{Clazz: dialogClazz}).ToDialog()
		if !ok {
			continue
		}
		if settings := notifySettings[notifyKeys[i]]; settings != nil {
			dialog.NotifySettings = settings
		}
	}

	messages, err := c.fetchDialogTopMessages(operation, topMessageRefs)
	if err != nil {
		return nil, err
	}
	users, err := c.fetchDialogUsersWithOperation(operation, userIDs)
	if err != nil {
		return nil, err
	}
	chats, err := c.fetchDialogChats(chatIDs)
	if err != nil {
		return nil, err
	}

	result.Dialogs = dialogs
	result.Messages = messages
	result.Users = users
	result.Chats = chats
	return result, nil
}

func makePublicDialogFromExtV2(dialogExt *dialogpb.TLDialogExtV2, topMessageID int32, notifySettings tg.PeerNotifySettingsClazz) (tg.DialogClazz, error) {
	folderID := dialogExt.FolderId
	ttlPeriod := int32(0)
	var draft tg.DraftMessageClazz
	if extras := dialogExt.Extras; extras != nil && extras.PrivateTtlPeriod != nil {
		ttlPeriod = *extras.PrivateTtlPeriod
	}
	if extras := dialogExt.Extras; extras != nil {
		parsedDraft, err := draftMessageFromPayload(extras.DraftPayload)
		if err != nil {
			return nil, err
		}
		draft = parsedDraft
	}
	return tg.MakeTLDialog(&tg.TLDialog{
		Pinned:               dialogExt.MainPinnedOrder > 0 || dialogExt.FolderPinnedOrder > 0,
		UnreadMark:           dialogExt.UnreadMark,
		Peer:                 makePublicPeerFromDialogFacade(dialogExt.PeerType, dialogExt.PeerId),
		TopMessage:           topMessageID,
		ReadInboxMaxId:       int64ToInt32Saturating(dialogExt.ReadInboxMaxUserMessageId),
		ReadOutboxMaxId:      int64ToInt32Saturating(dialogExt.ReadOutboxMaxUserMessageId),
		UnreadCount:          dialogExt.UnreadCount,
		UnreadMentionsCount:  dialogExt.UnreadMentionsCount,
		UnreadReactionsCount: dialogExt.UnreadReactionsCount,
		NotifySettings:       notifySettings,
		Draft:                draft,
		FolderId:             &folderID,
		TtlPeriod:            &ttlPeriod,
	}), nil
}

func (c *DialogsCore) resolveDialogCursorTopMessage(topMessageID int32) (*msgpb.ResolvedDialogCursor, error) {
	if c.svcCtx.Repo.MsgClient == nil {
		return nil, tg.ErrInternalServerError
	}
	resolved, err := c.svcCtx.Repo.MsgClient.MsgResolveDialogCursorTopMessage(c.ctx, &msgpb.TLMsgResolveDialogCursorTopMessage{
		UserId:       c.MD.UserId,
		TopMessageId: topMessageID,
	})
	if err != nil {
		c.Logger.Errorf("messages.getDialogs - msg.resolveDialogCursorTopMessage failed: user_id: %d, top_message_id: %d, err: %v",
			c.MD.UserId, topMessageID, err)
		return nil, tg.ErrInternalServerError
	}
	if resolved == nil || !tg.FromBoolClazz(resolved.Found) {
		return nil, nil
	}
	return resolved, nil
}

type draftPayloadEnvelope struct {
	Name   string             `json:"_name"`
	Object draftPayloadObject `json:"_object"`
}

type draftPayloadObject struct {
	NoWebpage   bool   `json:"no_webpage"`
	InvertMedia bool   `json:"invert_media"`
	Message     string `json:"message"`
	Date        int32  `json:"date"`
	Effect      *int64 `json:"effect"`
}

func draftMessageFromPayload(raw []byte) (tg.DraftMessageClazz, error) {
	if len(raw) == 0 {
		return nil, nil
	}
	var env draftPayloadEnvelope
	if err := json.Unmarshal(raw, &env); err != nil {
		return nil, err
	}
	switch env.Name {
	case "", tg.ClazzName_draftMessage:
		obj := env.Object
		if env.Name == "" {
			if err := json.Unmarshal(raw, &obj); err != nil {
				return nil, err
			}
		}
		return tg.MakeTLDraftMessage(&tg.TLDraftMessage{
			NoWebpage:   obj.NoWebpage,
			InvertMedia: obj.InvertMedia,
			Message:     obj.Message,
			Date:        obj.Date,
			Effect:      obj.Effect,
		}), nil
	case tg.ClazzName_draftMessageEmpty:
		return tg.MakeTLDraftMessageEmpty(&tg.TLDraftMessageEmpty{}), nil
	default:
		return nil, tg.ErrInternalServerError
	}
}

func int64ToDialogMessageID(v int64) (int32, error) {
	if v <= 0 {
		return 0, nil
	}
	if v > math.MaxInt32 {
		return 0, tg.ErrInternalServerError
	}
	return int32(v), nil
}

func int64ToInt32Saturating(v int64) int32 {
	if v <= 0 {
		return 0
	}
	if v > math.MaxInt32 {
		return math.MaxInt32
	}
	return int32(v)
}

func notifyPeerFromDialogFacade(selfID int64, peerType int32, peerID int64) (tg.PeerUtilClazz, bool) {
	if peerID <= 0 {
		return nil, false
	}
	out := &tg.TLPeerUtil{SelfId: selfID, PeerId: peerID}
	switch peerType {
	case dialogPeerTypeUser:
		out.PeerType = tg.PEER_USER
	case dialogPeerTypeChat:
		out.PeerType = tg.PEER_CHAT
	case dialogPeerTypeChannel:
		out.PeerType = tg.PEER_CHANNEL
	default:
		return nil, false
	}
	return tg.MakeTLPeerUtil(out), true
}

func notifySettingsKey(peerType int32, peerID int64) string {
	return strconv.FormatInt(int64(peerType), 10) + ":" + strconv.FormatInt(peerID, 10)
}

func dialogProjectionKey(peerType int32, peerID int64) string {
	return strconv.FormatInt(int64(peerType), 10) + ":" + strconv.FormatInt(peerID, 10)
}

func (c *DialogsCore) fetchDialogReadStates(operation string, peers []userupdates.DialogProjectionPeerClazz) (map[string]dialogReadState, error) {
	out := make(map[string]dialogReadState, len(peers))
	peers = uniqueDialogProjectionPeers(peers)
	if len(peers) == 0 {
		return out, nil
	}
	projections, err := c.svcCtx.Repo.UserupdatesClient.UserupdatesGetDialogsByPeers(c.ctx, &userupdates.TLUserupdatesGetDialogsByPeers{
		UserId: c.MD.UserId,
		Peers:  peers,
	})
	if err != nil {
		c.Logger.Errorf("%s - userupdates.getDialogsByPeers failed: user_id: %d, peer_count: %d, err: %v", operation, c.MD.UserId, len(peers), err)
		return nil, tg.ErrInternalServerError
	}
	if projections == nil {
		return out, nil
	}
	for _, projection := range projections.Datas {
		if projection == nil {
			continue
		}
		out[dialogProjectionKey(projection.PeerType, projection.PeerId)] = dialogReadState{
			ReadInboxMaxUserMessageID:  projection.ReadInboxMaxUserMessageId,
			ReadOutboxMaxUserMessageID: projection.ReadOutboxMaxUserMessageId,
		}
	}
	return out, nil
}

func uniqueDialogProjectionPeers(peers []userupdates.DialogProjectionPeerClazz) []userupdates.DialogProjectionPeerClazz {
	seen := make(map[string]struct{}, len(peers))
	out := make([]userupdates.DialogProjectionPeerClazz, 0, len(peers))
	for _, peer := range peers {
		if peer == nil || peer.PeerId <= 0 {
			continue
		}
		key := dialogProjectionKey(peer.PeerType, peer.PeerId)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, peer)
	}
	return out
}

func (c *DialogsCore) fetchDialogNotifySettings(operation string, peers []tg.PeerUtilClazz) (map[string]tg.PeerNotifySettingsClazz, error) {
	out := make(map[string]tg.PeerNotifySettingsClazz, len(peers))
	peers = uniquePeerUtils(peers)
	if len(peers) == 0 {
		return out, nil
	}
	if c.svcCtx.Repo.UserClient == nil {
		return out, nil
	}
	settings, err := c.svcCtx.Repo.UserClient.UserGetNotifySettingsList(c.ctx, &userpb.TLUserGetNotifySettingsList{
		UserId: c.MD.UserId,
		Peers:  peers,
	})
	if err != nil {
		c.Logger.Errorf("%s - user.getNotifySettingsList failed: user_id: %d, peer_count: %d, err: %v", operation, c.MD.UserId, len(peers), err)
		return nil, tg.ErrInternalServerError
	}
	if settings == nil {
		return out, nil
	}
	for _, item := range settings.Datas {
		if item == nil || item.Settings == nil {
			continue
		}
		out[notifySettingsKey(item.PeerType, item.PeerId)] = item.Settings
	}
	return out, nil
}

func uniquePeerUtils(peers []tg.PeerUtilClazz) []tg.PeerUtilClazz {
	seen := make(map[string]struct{}, len(peers))
	out := make([]tg.PeerUtilClazz, 0, len(peers))
	for _, peer := range peers {
		if peer == nil || peer.PeerId <= 0 {
			continue
		}
		key := notifySettingsKey(peer.PeerType, peer.PeerId)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, peer)
	}
	return out
}

func (c *DialogsCore) fetchDialogTopMessages(operation string, refs []dialogTopMessageRef) ([]tg.MessageClazz, error) {
	refs = uniqueDialogTopMessageRefs(refs)
	if len(refs) == 0 {
		return []tg.MessageClazz{}, nil
	}

	peers := make([]userupdates.MessageViewPeerSeqClazz, 0, len(refs))
	for _, ref := range refs {
		peers = append(peers, userupdates.MakeTLMessageViewPeerSeq(&userupdates.TLMessageViewPeerSeq{
			PeerType: ref.PeerType,
			PeerId:   ref.PeerID,
			PeerSeq:  ref.PeerSeq,
		}))
	}
	views, err := c.svcCtx.Repo.UserupdatesClient.UserupdatesGetMessageViewsByPeerSeqs(c.ctx, &userupdates.TLUserupdatesGetMessageViewsByPeerSeqs{
		UserId: c.MD.UserId,
		Peers:  peers,
	})
	if err != nil {
		c.Logger.Errorf("%s - userupdates.getMessageViewsByPeerSeqs failed: user_id: %d, refs: %+v, err: %v", operation, c.MD.UserId, refs, err)
		return nil, tg.ErrInternalServerError
	}
	if views == nil {
		return []tg.MessageClazz{}, nil
	}
	return views.Messages, nil
}

func (c *DialogsCore) fetchDialogUsers(ids []int64) ([]tg.UserClazz, error) {
	return c.fetchDialogUsersWithOperation("messages.getDialogs", ids)
}

func (c *DialogsCore) fetchDialogUsersWithOperation(operation string, ids []int64) ([]tg.UserClazz, error) {
	ids = uniqueInt64s(ids)
	if len(ids) == 0 {
		return []tg.UserClazz{}, nil
	}

	users, err := userprojection.ProjectUsers(c.ctx, c.svcCtx.Repo.UserClient, c.MD.UserId, ids, userprojection.MissingStoredReference)
	if err != nil {
		c.Logger.Errorf("%s - user.getUserProjectionBundle failed: user_id: %d, id: %v, err: %v", operation, c.MD.UserId, ids, err)
		return nil, tg.ErrInternalServerError
	}
	return users, nil
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

func uniqueDialogTopMessageRefs(refs []dialogTopMessageRef) []dialogTopMessageRef {
	seen := make(map[dialogTopMessageRef]struct{}, len(refs))
	out := make([]dialogTopMessageRef, 0, len(refs))
	for _, ref := range refs {
		if ref.PeerType == 0 || ref.PeerID == 0 || ref.PeerSeq == 0 {
			continue
		}
		if _, ok := seen[ref]; ok {
			continue
		}
		seen[ref] = struct{}{}
		out = append(out, ref)
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
