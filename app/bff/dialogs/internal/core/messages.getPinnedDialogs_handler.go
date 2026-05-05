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
	"strconv"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesGetPinnedDialogs
// messages.getPinnedDialogs#d6b94df2 folder_id:int = messages.PeerDialogs;
func (c *DialogsCore) MessagesGetPinnedDialogs(in *tg.TLMessagesGetPinnedDialogs) (*tg.MessagesPeerDialogs, error) {
	if c.MD == nil || c.MD.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	if in.FolderId != 0 && in.FolderId != 1 {
		return nil, tg.ErrFolderIdInvalid
	}

	dialogs, err := c.svcCtx.Repo.DialogClient.DialogGetPinnedDialogs(c.ctx, &dialogpb.TLDialogGetPinnedDialogs{
		UserId:   c.MD.UserId,
		FolderId: in.FolderId,
	})
	if err != nil {
		c.Logger.Errorf("messages.getPinnedDialogs - dialog.getPinnedDialogs failed: user_id: %d, folder_id: %d, err: %v",
			c.MD.UserId, in.FolderId, err)
		return nil, tg.ErrInternalServerError
	}
	dialogExts, err := c.hydratePinnedDialogProjections(vectorDialogExts(dialogs))
	if err != nil {
		return nil, err
	}
	return c.makePeerDialogsFromDialogExts("messages.getPinnedDialogs", dialogExts)
}

func (c *DialogsCore) hydratePinnedDialogProjections(dialogExts []dialogpb.DialogExtClazz) ([]dialogpb.DialogExtClazz, error) {
	if len(dialogExts) == 0 || c.svcCtx.Repo.UserupdatesClient == nil {
		return dialogExts, nil
	}
	peers := make([]userupdates.DialogProjectionPeerClazz, 0, len(dialogExts))
	orderByKey := make(map[string]int64, len(dialogExts))
	folderByKey := make(map[string]int32, len(dialogExts))
	for _, dialogExt := range dialogExts {
		if dialogExt == nil || dialogExt.Dialog == nil {
			continue
		}
		dialog, ok := (&tg.Dialog{Clazz: dialogExt.Dialog}).ToDialog()
		if !ok {
			continue
		}
		peerType, peerID, ok := dialogFacadePeerFromPublicPeer(dialog.Peer)
		if !ok {
			continue
		}
		key := dialogProjectionKey(peerType, peerID)
		orderByKey[key] = dialogExt.Order
		if dialog.FolderId != nil {
			folderByKey[key] = *dialog.FolderId
		}
		peers = append(peers, userupdates.MakeTLDialogProjectionPeer(&userupdates.TLDialogProjectionPeer{
			PeerType: peerType,
			PeerId:   peerID,
		}))
	}
	if len(peers) == 0 {
		return dialogExts, nil
	}
	projections, err := c.svcCtx.Repo.UserupdatesClient.UserupdatesGetDialogsByPeers(c.ctx, &userupdates.TLUserupdatesGetDialogsByPeers{
		UserId: c.MD.UserId,
		Peers:  peers,
	})
	if err != nil {
		c.Logger.Errorf("messages.getPinnedDialogs - userupdates.getDialogsByPeers failed: user_id: %d, err: %v", c.MD.UserId, err)
		return nil, tg.ErrInternalServerError
	}
	if projections == nil {
		return dialogExts, nil
	}
	byKey := make(map[string]userupdates.DialogProjectionClazz, len(projections.Datas))
	for _, projection := range projections.Datas {
		if projection == nil {
			continue
		}
		byKey[dialogProjectionKey(projection.PeerType, projection.PeerId)] = projection
	}
	out := make([]dialogpb.DialogExtClazz, 0, len(dialogExts))
	for _, dialogExt := range dialogExts {
		if dialogExt == nil || dialogExt.Dialog == nil {
			continue
		}
		dialog, ok := (&tg.Dialog{Clazz: dialogExt.Dialog}).ToDialog()
		if !ok {
			continue
		}
		peerType, peerID, ok := dialogFacadePeerFromPublicPeer(dialog.Peer)
		if !ok {
			out = append(out, dialogExt)
			continue
		}
		key := dialogProjectionKey(peerType, peerID)
		projection := byKey[key]
		if projection == nil {
			out = append(out, dialogExt)
			continue
		}
		folderID := folderByKey[key]
		ttlPeriod := int32(0)
		out = append(out, dialogpb.MakeTLDialogExt(&dialogpb.TLDialogExt{
			Order: orderByKey[key],
			Dialog: tg.MakeTLDialog(&tg.TLDialog{
				Pinned:               true,
				UnreadMark:           projection.UnreadMark,
				Peer:                 makePublicPeerFromDialogFacade(projection.PeerType, projection.PeerId),
				TopMessage:           int32(projection.TopCanonicalMessageId),
				ReadInboxMaxId:       int32(projection.ReadInboxMaxPeerSeq),
				ReadOutboxMaxId:      int32(projection.ReadOutboxMaxPeerSeq),
				UnreadCount:          projection.UnreadCount,
				UnreadMentionsCount:  projection.UnreadMentionsCount,
				UnreadReactionsCount: projection.UnreadReactionsCount,
				NotifySettings:       tg.MakeTLPeerNotifySettings(&tg.TLPeerNotifySettings{}),
				FolderId:             &folderID,
				TtlPeriod:            &ttlPeriod,
			}),
			Date: projection.TopMessageDate,
		}))
	}
	return out, nil
}

func (c *DialogsCore) makePeerDialogsFromDialogExts(operation string, dialogExts []dialogpb.DialogExtClazz) (*tg.MessagesPeerDialogs, error) {
	if len(dialogExts) == 0 {
		return emptyPeerDialogs(), nil
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
	users, err := c.fetchDialogUsersWithOperation(operation, userIDs)
	if err != nil {
		return nil, err
	}
	chats, err := c.fetchDialogChats(chatIDs)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessagesPeerDialogs(&tg.TLMessagesPeerDialogs{
		Dialogs:  dialogList,
		Messages: messages,
		Chats:    chats,
		Users:    users,
		State:    emptyUpdatesState(),
	}).ToMessagesPeerDialogs(), nil
}

func dialogFacadePeerFromPublicPeer(peer tg.PeerClazz) (int32, int64, bool) {
	switch p := peer.(type) {
	case *tg.TLPeerUser:
		return dialogPeerTypeUser, p.UserId, p.UserId > 0
	case *tg.TLPeerChat:
		return dialogPeerTypeChat, p.ChatId, p.ChatId > 0
	case *tg.TLPeerChannel:
		return dialogPeerTypeChannel, p.ChannelId, p.ChannelId > 0
	default:
		return 0, 0, false
	}
}

func dialogProjectionKey(peerType int32, peerID int64) string {
	return strconv.FormatInt(int64(peerType), 10) + ":" + strconv.FormatInt(peerID, 10)
}

func vectorDialogExts(dialogs *dialogpb.VectorDialogExt) []dialogpb.DialogExtClazz {
	if dialogs == nil {
		return nil
	}
	return dialogs.Datas
}

func emptyPeerDialogs() *tg.MessagesPeerDialogs {
	return tg.MakeTLMessagesPeerDialogs(&tg.TLMessagesPeerDialogs{
		Dialogs:  []tg.DialogClazz{},
		Messages: []tg.MessageClazz{},
		Chats:    []tg.ChatClazz{},
		Users:    []tg.UserClazz{},
		State:    emptyUpdatesState(),
	}).ToMessagesPeerDialogs()
}

func emptyUpdatesState() tg.UpdatesStateClazz {
	return tg.MakeTLUpdatesState(&tg.TLUpdatesState{
		Pts:         0,
		Qts:         0,
		Date:        0,
		Seq:         0,
		UnreadCount: 0,
	})
}
