package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/envelope"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	userprojection "github.com/teamgram/teamgram-server/v2/app/service/biz/user/userprojection"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type UserProjectionClient = userprojection.Client

type ChatProjectionClient interface {
	ChatGetChatListByIdList(ctx context.Context, in *chatpb.TLChatGetChatListByIdList) (*chatpb.VectorMutableChat, error)
}

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
	if len(ids) == 0 {
		return nil, nil
	}
	if r == nil || r.chatProjector == nil {
		return nil, fmt.Errorf("%w: project chats: chat client is not configured", userupdates.ErrUserupdatesStorage)
	}
	mutableChats, err := r.chatProjector.ChatGetChatListByIdList(ctx, &chatpb.TLChatGetChatListByIdList{
		SelfId: viewerUserID,
		IdList: ids,
	})
	if err != nil {
		return nil, storageError("project chats", err)
	}
	if mutableChats == nil {
		return nil, storageError("project chats", errors.New("nil mutable chat list"))
	}
	chats := make([]tg.ChatClazz, 0, len(mutableChats.Datas))
	for _, mutableChat := range mutableChats.Datas {
		chat, err := projectMutableChatForEnvelope(mutableChat, viewerUserID)
		if err != nil {
			return nil, storageError("project chats", err)
		}
		if chat != nil {
			chats = append(chats, chat)
		}
	}
	if len(chats) == 0 {
		return nil, storageError("project chats", fmt.Errorf("empty chat projection for ids %v", ids))
	}
	return chats, nil
}

func projectMutableChatForEnvelope(mutableChat tg.MutableChatClazz, viewerUserID int64) (tg.ChatClazz, error) {
	if mutableChat == nil || mutableChat.Chat == nil {
		return nil, nil
	}
	chat := mutableChat.Chat
	if chat.Date < minInt32 || chat.Date > maxInt32 {
		return nil, fmt.Errorf("chat date overflows int32: chat_id %d date %d", chat.Id, chat.Date)
	}
	return tg.MakeTLChat(&tg.TLChat{
		Creator:             chat.Creator == viewerUserID,
		Deactivated:         chat.Deactivated,
		CallActive:          chat.CallActive,
		CallNotEmpty:        chat.CallNotEmpty,
		Noforwards:          chat.Noforwards,
		Id:                  chat.Id,
		Title:               chat.Title,
		Photo:               projectChatPhotoForEnvelope(chat.Photo),
		ParticipantsCount:   chat.ParticipantsCount,
		Date:                int32(chat.Date),
		Version:             chat.Version,
		MigratedTo:          chat.MigratedTo,
		DefaultBannedRights: chat.DefaultBannedRights,
	}), nil
}

func projectChatPhotoForEnvelope(photo tg.PhotoClazz) tg.ChatPhotoClazz {
	if p, ok := photo.(*tg.TLPhoto); ok {
		return tg.MakeTLChatPhoto(&tg.TLChatPhoto{
			PhotoId: p.Id,
			DcId:    p.DcId,
		})
	}
	return tg.MakeTLChatPhotoEmpty(&tg.TLChatPhotoEmpty{})
}

const (
	minInt32 = -1 << 31
	maxInt32 = 1<<31 - 1
)

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
