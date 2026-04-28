package core

import (
	"regexp"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var chatInviteHashRE = regexp.MustCompile(`^[A-Za-z0-9_-]+$`)

func validChatInviteHash(hash string) bool {
	return len(hash) == 20 && chatInviteHashRE.MatchString(hash)
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

func (c *ChatInvitesCore) projectChatInviteExt(invite *chatpb.ChatInviteExt, selfID int64) (*tg.ChatInvite, error) {
	if already, ok := invite.ToChatInviteAlready(); ok {
		return tg.MakeTLChatInviteAlready(&tg.TLChatInviteAlready{
			Chat: projectMutableChat(already.Chat, selfID),
		}).ToChatInvite(), nil
	}

	if normal, ok := invite.ToChatInvite(); ok {
		users, err := c.fetchUserClazzes(normal.Participants, selfID)
		if err != nil {
			return nil, err
		}
		return tg.MakeTLChatInvite(&tg.TLChatInvite{
			RequestNeeded:     normal.RequestNeeded,
			Title:             normal.Title,
			About:             normal.About,
			Photo:             normal.Photo,
			ParticipantsCount: normal.ParticipantsCount,
			Participants:      users,
		}).ToChatInvite(), nil
	}

	if peek, ok := invite.ToChatInvitePeek(); ok {
		return tg.MakeTLChatInvitePeek(&tg.TLChatInvitePeek{
			Chat:    projectMutableChat(peek.Chat, selfID),
			Expires: peek.Expires,
		}).ToChatInvite(), nil
	}

	return nil, tg.ErrInternalServerError
}

func exportedInviteClazz(invite *tg.ExportedChatInvite) tg.ExportedChatInviteClazz {
	if invite == nil {
		return nil
	}
	return invite.Clazz
}

func adminIDsFromInvites(invites []tg.ExportedChatInviteClazz) []int64 {
	ids := make([]int64, 0, len(invites))
	for _, invite := range invites {
		if x, ok := invite.(*tg.TLChatInviteExported); ok {
			ids = append(ids, x.AdminId)
		}
	}
	return ids
}

func adminIDsFromAdmins(admins []tg.ChatAdminWithInvitesClazz) []int64 {
	ids := make([]int64, 0, len(admins))
	for _, admin := range admins {
		if admin != nil {
			ids = append(ids, admin.AdminId)
		}
	}
	return ids
}

func userIDsFromImporters(importers []tg.ChatInviteImporterClazz) []int64 {
	ids := make([]int64, 0, len(importers))
	for _, importer := range importers {
		if importer != nil {
			ids = append(ids, importer.UserId)
		}
	}
	return ids
}
