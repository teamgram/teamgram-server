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
	userprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesGetFullChat
// messages.getFullChat#aeb00b34 chat_id:long = messages.ChatFull;
func (c *ChatsCore) MessagesGetFullChat(in *tg.TLMessagesGetFullChat) (*tg.MessagesChatFull, error) {
	selfID := selfID(c.MD)
	mutableChat, err := c.svcCtx.Repo.ChatClient.ChatGetMutableChat(c.ctx, &chatpb.TLChatGetMutableChat{
		ChatId: in.ChatId,
	})
	if err != nil {
		return nil, mapChatError(err)
	}
	if mutableChat == nil || mutableChat.Chat == nil {
		return nil, tg.ErrChatIdInvalid
	}

	me, ok := chatpb.GetImmutableChatParticipant(mutableChat, selfID)
	if !ok || !chatpb.IsChatMemberStateNormal(me) {
		return nil, tg.Err400PeerIdInvalid
	}

	chatFull := tg.MakeTLChatFull(&tg.TLChatFull{
		CanSetUsername:     mutableChat.Chat.CanSetUsername,
		Id:                 mutableChat.Chat.Id,
		About:              mutableChat.Chat.About,
		Participants:       projectChatParticipants(mutableChat),
		ChatPhoto:          mutableChat.Chat.Photo,
		NotifySettings:     tg.MakeTLPeerNotifySettings(&tg.TLPeerNotifySettings{}),
		ExportedInvite:     mutableChat.Chat.ExportedInvite,
		BotInfo:            append([]tg.BotInfoClazz(nil), mutableChat.Chat.BotInfo...),
		FolderId:           c.getChatFolderID(in.ChatId, selfID),
		Call:               mutableChat.Chat.Call,
		TtlPeriod:          tg.MakeFlagsInt32(mutableChat.Chat.TtlPeriod),
		AvailableReactions: projectAvailableReactions(mutableChat.Chat),
	})

	if c.svcCtx.Repo.UserClient != nil {
		if settings, _ := c.svcCtx.Repo.UserClient.UserGetNotifySettings(c.ctx, &userpb.TLUserGetNotifySettings{
			UserId:   selfID,
			PeerType: tg.PEER_CHAT,
			PeerId:   in.ChatId,
		}); settings != nil {
			chatFull.NotifySettings = settings
		}
	}

	if chatpb.CanInviteUsers(me) {
		if me.Link != "" {
			if invite, _ := c.svcCtx.Repo.ChatClient.ChatGetExportedChatInvite(c.ctx, &chatpb.TLChatGetExportedChatInvite{
				ChatId: in.ChatId,
				Link:   me.Link,
			}); invite != nil && invite.Clazz != nil {
				chatFull.ExportedInvite = invite.Clazz
			}
		}
		if requesters, _ := c.svcCtx.Repo.ChatClient.ChatGetRecentChatInviteRequesters(c.ctx, &chatpb.TLChatGetRecentChatInviteRequesters{
			SelfId: selfID,
			ChatId: in.ChatId,
		}); requesters != nil && len(requesters.RecentRequesters) > 0 {
			chatFull.RequestsPending = &requesters.RequestsPending
			chatFull.RecentRequesters = requesters.RecentRequesters
		}
	}

	userIDs := collectFullChatUserIDs(mutableChat)
	if c.svcCtx.Repo.UserClient != nil {
		for _, p := range mutableChat.ChatParticipants {
			if p == nil || !p.IsBot || !chatpb.IsChatMemberStateNormal(p) {
				continue
			}
			if botInfo, _ := c.svcCtx.Repo.UserClient.UserGetBotInfo(c.ctx, &userpb.TLUserGetBotInfo{
				BotId: p.UserId,
			}); botInfo != nil {
				chatFull.BotInfo = append(chatFull.BotInfo, botInfo)
			}
		}
	}

	users := []tg.UserClazz{}
	if c.svcCtx.Repo.UserClient != nil {
		users, err = userprojection.ProjectUsers(c.ctx, c.svcCtx.Repo.UserClient, selfID, userIDs, userprojection.MissingStoredReference)
		if err != nil {
			return nil, err
		}
	}

	return tg.MakeTLMessagesChatFull(&tg.TLMessagesChatFull{
		FullChat: chatFull,
		Chats: []tg.ChatClazz{
			projectMutableChat(mutableChat, selfID),
		},
		Users: users,
	}).ToMessagesChatFull(), nil
}

func (c *ChatsCore) getChatFolderID(chatID, selfID int64) *int32 {
	if c.svcCtx.Repo.DialogClient == nil {
		return nil
	}
	dialogs, err := c.svcCtx.Repo.DialogClient.DialogGetPeerDialogsV2(c.ctx, &dialogpb.TLDialogGetPeerDialogsV2{
		UserId: selfID,
		Peers: []dialogpb.DialogPeerClazz{
			dialogpb.MakeTLDialogPeer(&dialogpb.TLDialogPeer{
				PeerType: dialogPeerTypeChat,
				PeerId:   chatID,
			}),
		},
	})
	if err != nil || dialogs == nil || len(dialogs.Datas) == 0 {
		return nil
	}
	dialogExt := dialogs.Datas[0]
	if dialogExt == nil {
		return nil
	}
	if dialogExt.FolderId == 0 {
		return nil
	}
	return &dialogExt.FolderId
}

const dialogPeerTypeChat int32 = 2

func projectChatParticipants(chat *tg.MutableChat) tg.ChatParticipantsClazz {
	if chat == nil || chat.Chat == nil {
		return tg.MakeTLChatParticipants(&tg.TLChatParticipants{})
	}

	participants := make([]tg.ChatParticipantClazz, 0, len(chat.ChatParticipants))
	for _, p := range chat.ChatParticipants {
		if !chatpb.IsChatMemberStateNormal(p) {
			continue
		}
		participants = append(participants, projectChatParticipant(p))
	}

	return tg.MakeTLChatParticipants(&tg.TLChatParticipants{
		ChatId:       chat.Chat.Id,
		Participants: participants,
		Version:      chat.Chat.Version,
	})
}

func projectChatParticipant(p *tg.ImmutableChatParticipant) tg.ChatParticipantClazz {
	switch p.ParticipantType {
	case chatpb.ChatMemberCreator:
		return tg.MakeTLChatParticipantCreator(&tg.TLChatParticipantCreator{
			UserId: p.UserId,
		})
	case chatpb.ChatMemberAdmin:
		return tg.MakeTLChatParticipantAdmin(&tg.TLChatParticipantAdmin{
			UserId:    p.UserId,
			InviterId: p.InviterUserId,
			Date:      int32(p.Date),
		})
	default:
		return tg.MakeTLChatParticipant(&tg.TLChatParticipant{
			UserId:    p.UserId,
			InviterId: p.InviterUserId,
			Date:      int32(p.Date),
		})
	}
}

func projectAvailableReactions(chat *tg.ImmutableChat) tg.ChatReactionsClazz {
	if chat == nil {
		return nil
	}
	switch {
	case chat.AvailableReactionsType == 0:
		return tg.MakeTLChatReactionsNone(&tg.TLChatReactionsNone{})
	case chat.AvailableReactionsType == 4:
		reactions := make([]tg.ReactionClazz, 0, len(chat.AvailableReactions))
		for _, emoticon := range chat.AvailableReactions {
			reactions = append(reactions, tg.MakeTLReactionEmoji(&tg.TLReactionEmoji{Emoticon: emoticon}))
		}
		return tg.MakeTLChatReactionsSome(&tg.TLChatReactionsSome{Reactions: reactions})
	default:
		return tg.MakeTLChatReactionsAll(&tg.TLChatReactionsAll{})
	}
}

func collectFullChatUserIDs(chat *tg.MutableChat) []int64 {
	if chat == nil {
		return nil
	}
	seen := make(map[int64]struct{}, len(chat.ChatParticipants))
	ids := make([]int64, 0, len(chat.ChatParticipants))
	for _, p := range chat.ChatParticipants {
		if !chatpb.IsChatMemberStateNormal(p) {
			continue
		}
		if _, ok := seen[p.UserId]; ok {
			continue
		}
		seen[p.UserId] = struct{}{}
		ids = append(ids, p.UserId)
	}
	return ids
}
