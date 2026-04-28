package repository

import (
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/zeromicro/go-zero/core/jsonx"
)

const chatReactionsTypeSome int32 = 4

func makeImmutableChat(row *model.Chats, photo tg.PhotoClazz) *tg.ImmutableChat {
	if row == nil {
		return nil
	}

	chat := tg.MakeTLImmutableChat(&tg.TLImmutableChat{
		Id:                     row.Id,
		Creator:                row.CreatorUserId,
		Title:                  row.Title,
		Photo:                  photo,
		Deactivated:            row.Deactivated,
		Noforwards:             row.Noforwards,
		ParticipantsCount:      row.ParticipantCount,
		Date:                   row.Date,
		Version:                row.Version,
		DefaultBannedRights:    chatBannedRightsFromStorage(row.DefaultBannedRights),
		About:                  row.About,
		AvailableReactionsType: row.AvailableReactionsType,
		AvailableReactions:     availableReactionsFromStorage(row.AvailableReactionsType, row.AvailableReactions),
		TtlPeriod:              row.TtlPeriod,
	}).ToImmutableChat()

	if row.MigratedToId != 0 {
		chat.MigratedTo = tg.MakeTLInputChannel(&tg.TLInputChannel{
			ChannelId:  row.MigratedToId,
			AccessHash: row.MigratedToAccessHash,
		})
	}

	return chat
}

func makeImmutableChatParticipant(row *model.ChatParticipants) *tg.ImmutableChatParticipant {
	if row == nil {
		return nil
	}

	return tg.MakeTLImmutableChatParticipant(&tg.TLImmutableChatParticipant{
		Id:              row.Id,
		ChatId:          row.ChatId,
		UserId:          row.UserId,
		State:           row.State,
		ParticipantType: row.ParticipantType,
		Link:            row.Link,
		Useage:          row.Usage2,
		InviterUserId:   row.InviterUserId,
		InvitedAt:       row.InvitedAt,
		KickedAt:        row.KickedAt,
		LeftAt:          row.LeftAt,
		AdminRights:     chatAdminRightsFromStorage(row.AdminRights),
		Date:            row.Date2,
		IsBot:           row.IsBot,
	}).ToImmutableChatParticipant()
}

func makeMutableChat(chat *tg.ImmutableChat, participants []*tg.ImmutableChatParticipant) *tg.MutableChat {
	return tg.MakeTLMutableChat(&tg.TLMutableChat{
		Chat:             chat,
		ChatParticipants: participants,
	}).ToMutableChat()
}

func chatBannedRightsFromStorage(mask int64) tg.ChatBannedRightsClazz {
	return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{
		ViewMessages:    hasInt64Bit(mask, 0),
		SendMessages:    hasInt64Bit(mask, 1),
		SendMedia:       hasInt64Bit(mask, 2),
		SendStickers:    hasInt64Bit(mask, 3),
		SendGifs:        hasInt64Bit(mask, 4),
		SendGames:       hasInt64Bit(mask, 5),
		SendInline:      hasInt64Bit(mask, 6),
		EmbedLinks:      hasInt64Bit(mask, 7),
		SendPolls:       hasInt64Bit(mask, 8),
		ChangeInfo:      hasInt64Bit(mask, 10),
		InviteUsers:     hasInt64Bit(mask, 15),
		PinMessages:     hasInt64Bit(mask, 17),
		ManageTopics:    hasInt64Bit(mask, 18),
		SendPhotos:      hasInt64Bit(mask, 19),
		SendVideos:      hasInt64Bit(mask, 20),
		SendRoundvideos: hasInt64Bit(mask, 21),
		SendAudios:      hasInt64Bit(mask, 22),
		SendVoices:      hasInt64Bit(mask, 23),
		SendDocs:        hasInt64Bit(mask, 24),
		SendPlain:       hasInt64Bit(mask, 25),
		EditRank:        hasInt64Bit(mask, 26),
	}).ToChatBannedRights()
}

func chatBannedRightsToStorage(rights tg.ChatBannedRightsClazz) int64 {
	if rights == nil {
		return 0
	}

	var mask int64
	setInt64Bit(&mask, 0, rights.ViewMessages)
	setInt64Bit(&mask, 1, rights.SendMessages)
	setInt64Bit(&mask, 2, rights.SendMedia)
	setInt64Bit(&mask, 3, rights.SendStickers)
	setInt64Bit(&mask, 4, rights.SendGifs)
	setInt64Bit(&mask, 5, rights.SendGames)
	setInt64Bit(&mask, 6, rights.SendInline)
	setInt64Bit(&mask, 7, rights.EmbedLinks)
	setInt64Bit(&mask, 8, rights.SendPolls)
	setInt64Bit(&mask, 10, rights.ChangeInfo)
	setInt64Bit(&mask, 15, rights.InviteUsers)
	setInt64Bit(&mask, 17, rights.PinMessages)
	setInt64Bit(&mask, 18, rights.ManageTopics)
	setInt64Bit(&mask, 19, rights.SendPhotos)
	setInt64Bit(&mask, 20, rights.SendVideos)
	setInt64Bit(&mask, 21, rights.SendRoundvideos)
	setInt64Bit(&mask, 22, rights.SendAudios)
	setInt64Bit(&mask, 23, rights.SendVoices)
	setInt64Bit(&mask, 24, rights.SendDocs)
	setInt64Bit(&mask, 25, rights.SendPlain)
	setInt64Bit(&mask, 26, rights.EditRank)
	return mask
}

func chatAdminRightsFromStorage(mask int32) tg.ChatAdminRightsClazz {
	return tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{
		ChangeInfo:           hasInt32Bit(mask, 0),
		PostMessages:         hasInt32Bit(mask, 1),
		EditMessages:         hasInt32Bit(mask, 2),
		DeleteMessages:       hasInt32Bit(mask, 3),
		BanUsers:             hasInt32Bit(mask, 4),
		InviteUsers:          hasInt32Bit(mask, 5),
		PinMessages:          hasInt32Bit(mask, 7),
		AddAdmins:            hasInt32Bit(mask, 9),
		Anonymous:            hasInt32Bit(mask, 10),
		ManageCall:           hasInt32Bit(mask, 11),
		Other:                hasInt32Bit(mask, 12),
		ManageTopics:         hasInt32Bit(mask, 13),
		PostStories:          hasInt32Bit(mask, 14),
		EditStories:          hasInt32Bit(mask, 15),
		DeleteStories:        hasInt32Bit(mask, 16),
		ManageDirectMessages: hasInt32Bit(mask, 17),
		ManageRanks:          hasInt32Bit(mask, 18),
	}).ToChatAdminRights()
}

func chatAdminRightsToStorage(rights tg.ChatAdminRightsClazz) int32 {
	if rights == nil {
		return 0
	}

	var mask int32
	setInt32Bit(&mask, 0, rights.ChangeInfo)
	setInt32Bit(&mask, 1, rights.PostMessages)
	setInt32Bit(&mask, 2, rights.EditMessages)
	setInt32Bit(&mask, 3, rights.DeleteMessages)
	setInt32Bit(&mask, 4, rights.BanUsers)
	setInt32Bit(&mask, 5, rights.InviteUsers)
	setInt32Bit(&mask, 7, rights.PinMessages)
	setInt32Bit(&mask, 9, rights.AddAdmins)
	setInt32Bit(&mask, 10, rights.Anonymous)
	setInt32Bit(&mask, 11, rights.ManageCall)
	setInt32Bit(&mask, 12, rights.Other)
	setInt32Bit(&mask, 13, rights.ManageTopics)
	setInt32Bit(&mask, 14, rights.PostStories)
	setInt32Bit(&mask, 15, rights.EditStories)
	setInt32Bit(&mask, 16, rights.DeleteStories)
	setInt32Bit(&mask, 17, rights.ManageDirectMessages)
	setInt32Bit(&mask, 18, rights.ManageRanks)
	return mask
}

func availableReactionsFromStorage(kind int32, payload string) []string {
	if kind != chatReactionsTypeSome || payload == "" {
		return []string{}
	}

	var reactions []string
	if err := jsonx.UnmarshalFromString(payload, &reactions); err != nil {
		return []string{}
	}
	if reactions == nil {
		return []string{}
	}
	return reactions
}

func availableReactionsToStorage(reactions []string) string {
	if len(reactions) == 0 {
		return "[]"
	}
	payload, err := jsonx.MarshalToString(reactions)
	if err != nil {
		return "[]"
	}
	return payload
}

func hasInt64Bit(mask int64, bit uint) bool {
	return mask&(int64(1)<<bit) != 0
}

func setInt64Bit(mask *int64, bit uint, ok bool) {
	if ok {
		*mask |= int64(1) << bit
	}
}

func hasInt32Bit(mask int32, bit uint) bool {
	return mask&(int32(1)<<bit) != 0
}

func setInt32Bit(mask *int32, bit uint, ok bool) {
	if ok {
		*mask |= int32(1) << bit
	}
}
