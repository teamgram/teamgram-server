package core

import (
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func selfID(md *metadata.RpcMetadata) int64 {
	if md == nil {
		return 0
	}
	return md.UserId
}

func projectMutableChat(chat *tg.MutableChat, selfID int64) tg.ChatClazz {
	if chat == nil || chat.Chat == nil {
		return nil
	}

	return tg.MakeTLChat(&tg.TLChat{
		Creator:             chat.Chat.Creator == selfID,
		Deactivated:         chat.Chat.Deactivated,
		CallActive:          chat.Chat.CallActive,
		CallNotEmpty:        chat.Chat.CallNotEmpty,
		Noforwards:          chat.Chat.Noforwards,
		Id:                  chat.Chat.Id,
		Title:               chat.Chat.Title,
		Photo:               projectChatPhoto(chat.Chat.Photo),
		ParticipantsCount:   chat.Chat.ParticipantsCount,
		Date:                int32(chat.Chat.Date),
		Version:             chat.Chat.Version,
		MigratedTo:          chat.Chat.MigratedTo,
		DefaultBannedRights: chat.Chat.DefaultBannedRights,
	})
}

func projectChatPhoto(photo tg.PhotoClazz) tg.ChatPhotoClazz {
	if p, ok := photo.(*tg.TLPhoto); ok {
		return tg.MakeTLChatPhoto(&tg.TLChatPhoto{
			PhotoId: p.Id,
			DcId:    p.DcId,
		})
	}

	return tg.MakeTLChatPhotoEmpty(&tg.TLChatPhotoEmpty{})
}

func updatesWithChat(chat *tg.MutableChat, selfID int64) *tg.Updates {
	chatID := int64(0)
	if chat != nil && chat.Chat != nil {
		chatID = chat.Chat.Id
	}

	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{
			tg.MakeTLUpdateChat(&tg.TLUpdateChat{ChatId: chatID}),
		},
		Users: []tg.UserClazz{},
		Chats: []tg.ChatClazz{
			projectMutableChat(chat, selfID),
		},
	}).ToUpdates()
}

func invitedUsersWithChat(chat *tg.MutableChat, selfID int64) *tg.MessagesInvitedUsers {
	return tg.MakeTLMessagesInvitedUsers(&tg.TLMessagesInvitedUsers{
		Updates:         updatesWithChat(chat, selfID).Clazz,
		MissingInvitees: []tg.MissingInviteeClazz{},
	}).ToMessagesInvitedUsers()
}
