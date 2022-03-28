// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"context"
	"github.com/teamgram/proto/mtproto"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/media/media"
	"github.com/zeromicro/go-zero/core/jsonx"
)

func (d *Dao) MakeImmutableChatByDO(ctx context.Context, chatsDO *dataobject.ChatsDO) (chat *chatpb.ImmutableChat) {
	chat = &chatpb.ImmutableChat{
		Id:                  chatsDO.Id,
		Creator:             chatsDO.CreatorUserId,
		Title:               chatsDO.Title,
		Photo:               nil,
		Deactivated:         chatsDO.Deactivated,
		ParticipantsCount:   chatsDO.ParticipantCount,
		Date:                chatsDO.Date,
		Version:             chatsDO.Version,
		MigratedTo:          nil,
		DefaultBannedRights: mtproto.BannedRights(chatsDO.DefaultBannedRights).ToChatBannedRights(),
		CanSetUsername:      false,
		About:               chatsDO.About,
		ExportedInvite:      nil,
		BotInfo:             nil,
	}

	if chatsDO.MigratedToId != 0 && chatsDO.MigratedToAccessHash != 0 {
		chat.MigratedTo = mtproto.MakeTLInputChannel(&mtproto.InputChannel{
			ChannelId:  chatsDO.MigratedToId,
			AccessHash: chatsDO.MigratedToAccessHash,
		}).To_InputChannel()
	}

	//// chat_photo && photo
	if chatsDO.PhotoId != 0 {
		chat.Photo, _ = d.MediaClient.MediaGetPhoto(ctx, &media.TLMediaGetPhoto{
			PhotoId: chatsDO.PhotoId,
		})
	}
	if chat.Photo == nil {
		chat.Photo = mtproto.MakeTLPhotoEmpty(nil).To_Photo()
	}

	//if chatsDO.Link != "" {
	//	// chatInviteExported#6e24fc9d flags:#
	//	//	revoked:flags.0?true
	//	//	permanent:flags.5?true
	//	//	link:string
	//	//	admin_id:int
	//	//	date:int
	//	//	start_date:flags.4?int
	//	//	expire_date:flags.1?int
	//	//	usage_limit:flags.2?int
	//	//	usage:flags.3?int = ExportedChatInvite;
	//	//
	//	chat.ExportedInvite = mtproto.MakeTLChatInviteExported(&mtproto.ExportedChatInvite{
	//		Revoked:    false,
	//		Permanent:  true,
	//		Link:       chatsDO.Link,
	//		AdminId:    chat.Creator,
	//		Date:       int32(chat.Date),
	//		StartDate:  nil,
	//		ExpireDate: nil,
	//		UsageLimit: nil,
	//		Usage:      nil,
	//	}).To_ExportedChatInvite()
	//} else {
	chat.ExportedInvite = nil // model.ExportedChatInviteEmpty
	//}

	if chatsDO.AvailableReactions != "" {
		jsonx.UnmarshalFromString(chatsDO.AvailableReactions, &chat.AvailableReactions)
	}

	return
}

func (d *Dao) MakeImmutableChatParticipant(chatParticipantsDO *dataobject.ChatParticipantsDO) (participant *chatpb.ImmutableChatParticipant) {
	participant = &chatpb.ImmutableChatParticipant{
		Id:              chatParticipantsDO.Id,
		ChatId:          chatParticipantsDO.ChatId,
		UserId:          chatParticipantsDO.UserId,
		State:           chatParticipantsDO.State,
		ParticipantType: chatParticipantsDO.ParticipantType,
		Link:            chatParticipantsDO.Link,
		InviterUserId:   chatParticipantsDO.InviterUserId,
		InvitedAt:       chatParticipantsDO.InvitedAt,
		KickedAt:        chatParticipantsDO.KickedAt,
		LeftAt:          chatParticipantsDO.LeftAt,
		AdminRights:     nil,
		Date:            chatParticipantsDO.Date2,
	}

	if participant.ParticipantType == mtproto.ChatMemberAdmin {
		participant.AdminRights = mtproto.MakeDefaultChatAdminRights()
	}

	return
}

func (d *Dao) GetMutableChat(ctx context.Context, chatId int64, id ...int64) (chat *chatpb.MutableChat, err error) {
	var (
		immutableChat *chatpb.ImmutableChat
		participants  []*chatpb.ImmutableChatParticipant
	)

	immutableChat, err = d.getImmutableChat(ctx, chatId)
	if err != nil {
		return
	}
	if d.Plugin != nil {
		immutableChat.CallActive, immutableChat.CallNotEmpty = d.Plugin.GetChatCallActiveAndNotEmpty(ctx, 0, chatId)
		immutableChat.Call = d.Plugin.GetChatGroupCall(ctx, 0, chatId)
	}
	participants, err = d.getImmutableChatParticipants(ctx, immutableChat, id...)
	if err != nil {
		return
	}

	chat = chatpb.MakeTLMutableChat(&chatpb.MutableChat{
		Chat:             immutableChat,
		ChatParticipants: participants,
	}).To_MutableChat()

	return
}

func (d *Dao) getImmutableChat(ctx context.Context, chatId int64) (chat *chatpb.ImmutableChat, err error) {
	var (
		chatsDO *dataobject.ChatsDO
	)

	chatsDO, err = d.ChatsDAO.Select(ctx, chatId)
	if err != nil {
		return
	} else if chatsDO == nil {
		err = mtproto.ErrChatIdInvalid
		return
	}
	// logx.Errorf("chatsDO: %#v", chatsDO)
	chat = d.MakeImmutableChatByDO(ctx, chatsDO)

	return
}

func (d *Dao) getImmutableChatParticipants(ctx context.Context, chat *chatpb.ImmutableChat, id ...int64) (participants []*chatpb.ImmutableChatParticipant, err error) {
	if len(id) == 0 {
		_, err = d.ChatParticipantsDAO.SelectListWithCB(
			ctx,
			chat.Id,
			func(i int, v *dataobject.ChatParticipantsDO) {
				participants = append(participants, d.MakeImmutableChatParticipant(v))
			})
	} else {
		_, err = d.ChatParticipantsDAO.SelectListByParticipantIdListWithCB(
			ctx,
			chat.Id,
			id,
			func(i int, v *dataobject.ChatParticipantsDO) {
				participants = append(participants, d.MakeImmutableChatParticipant(v))
			})
	}
	return
}
