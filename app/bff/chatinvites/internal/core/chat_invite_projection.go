package core

import (
	"regexp"

	chatprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/chatprojection"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var chatInviteHashRE = regexp.MustCompile(`^[A-Za-z0-9_-]+$`)

func validChatInviteHash(hash string) bool {
	return len(hash) == 20 && chatInviteHashRE.MatchString(hash)
}

func (c *ChatInvitesCore) projectChatInviteExt(invite *chatpb.ChatInviteExt, selfID int64) (*tg.ChatInvite, error) {
	if already, ok := invite.ToChatInviteAlready(); ok {
		chat, err := chatprojection.ProjectMutableChat(already.Chat, selfID)
		if err != nil {
			return nil, tg.ErrInternalServerError
		}
		return tg.MakeTLChatInviteAlready(&tg.TLChatInviteAlready{
			Chat: chat,
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
		chat, err := chatprojection.ProjectMutableChat(peek.Chat, selfID)
		if err != nil {
			return nil, tg.ErrInternalServerError
		}
		return tg.MakeTLChatInvitePeek(&tg.TLChatInvitePeek{
			Chat:    chat,
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
