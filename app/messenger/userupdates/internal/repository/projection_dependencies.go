package repository

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/envelope"
	chatprojection "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chatprojection"
	userprojection "github.com/teamgram/teamgram-server/v2/app/service/biz/user/userprojection"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type UserProjectionClient = userprojection.Client

type ChatProjectionClient = chatprojection.Client

func (r *Repository) SetPeerProjectionClients(user UserProjectionClient, chat ChatProjectionClient) {
	if r == nil {
		return
	}
	r.userProjector = user
	r.chatProjector = chat
}

func (r *Repository) BuildUpdatesWithDependencies(ctx context.Context, viewerUserID int64, in envelope.Input) (*tg.Updates, error) {
	return BuildUpdatesWithDependencies(ctx, r, viewerUserID, in)
}

func BuildUpdatesWithDependencies(ctx context.Context, projector envelope.PeerObjectProjector, viewerUserID int64, in envelope.Input) (*tg.Updates, error) {
	return envelope.BuildUpdatesWithDependencies(ctx, projector, viewerUserID, in)
}

func (r *Repository) ProjectUsers(ctx context.Context, viewerUserID int64, ids []int64) ([]tg.UserClazz, error) {
	var client UserProjectionClient
	if r != nil {
		client = r.userProjector
	}
	users, err := userprojection.ProjectUsers(ctx, client, viewerUserID, ids, userprojection.Options{
		Missing:         userprojection.MissingStoredReference,
		RequireNonEmpty: true,
	})
	if err != nil {
		return nil, storageError("project users", err)
	}
	return users, nil
}

func (r *Repository) ProjectChats(ctx context.Context, viewerUserID int64, ids []int64) ([]tg.ChatClazz, error) {
	var client ChatProjectionClient
	if r != nil {
		client = r.chatProjector
	}
	chats, err := chatprojection.ProjectChats(ctx, client, viewerUserID, ids, chatprojection.Options{
		Missing:         chatprojection.MissingStoredReference,
		RequireNonEmpty: true,
	})
	if err != nil {
		return nil, storageError("project chats", err)
	}
	return chats, nil
}

func userID(user tg.UserClazz) int64 {
	switch u := user.(type) {
	case *tg.TLUser:
		if u == nil {
			return 0
		}
		return u.Id
	case *tg.TLUserEmpty:
		if u == nil {
			return 0
		}
		return u.Id
	default:
		return 0
	}
}

func chatID(chat tg.ChatClazz) int64 {
	switch c := chat.(type) {
	case *tg.TLChat:
		if c == nil {
			return 0
		}
		return c.Id
	case *tg.TLChatEmpty:
		if c == nil {
			return 0
		}
		return c.Id
	case *tg.TLChatForbidden:
		if c == nil {
			return 0
		}
		return c.Id
	default:
		return 0
	}
}
