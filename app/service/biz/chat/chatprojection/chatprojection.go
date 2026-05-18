package chatprojection

import (
	"context"
	"errors"
	"fmt"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type Client interface {
	ChatGetChatProjectionBundle(ctx context.Context, in *chatpb.TLChatGetChatProjectionBundle) (*chatpb.ChatProjectionBundle, error)
}

type MissingPolicy int

const (
	MissingExplicitInput MissingPolicy = iota
	MissingStoredReference
)

type Options struct {
	Missing         MissingPolicy
	RequireNonEmpty bool
}

var (
	ErrClientNotConfigured     = errors.New("chat projection: client not configured")
	ErrInvalidRequest          = errors.New("chat projection: invalid request")
	ErrNilBundle               = errors.New("chat projection: nil bundle")
	ErrExplicitChatMissing     = errors.New("chat projection: explicit chat missing")
	ErrViewerProjectionMissing = errors.New("chat projection: viewer projection missing")
	ErrViewerProjectionEmpty   = errors.New("chat projection: viewer projection empty")
	ErrChatIDOverflow          = errors.New("chat projection: chat id overflow")
	ErrChatDateOverflow        = errors.New("chat projection: chat date overflow")
)

const (
	minInt32 = -1 << 31
	maxInt32 = 1<<31 - 1
)

func Int32ChatID(chatID int64) (int32, error) {
	if chatID < minInt32 || chatID > maxInt32 {
		return 0, fmt.Errorf("%w: chat_id=%d", ErrChatIDOverflow, chatID)
	}
	return int32(chatID), nil
}

func ProjectChats(ctx context.Context, client Client, viewerUserID int64, targetChatIDs []int64, opts Options) ([]tg.ChatClazz, error) {
	if len(targetChatIDs) == 0 {
		return nil, nil
	}
	if client == nil {
		return nil, ErrClientNotConfigured
	}

	bundle, err := client.ChatGetChatProjectionBundle(ctx, &chatpb.TLChatGetChatProjectionBundle{
		ViewerUserIds: []int64{viewerUserID},
		TargetChatIds: targetChatIDs,
	})
	if err != nil {
		if errors.Is(err, chatpb.ErrChatInvalidArgument) {
			return nil, fmt.Errorf("%w: %w", ErrInvalidRequest, err)
		}
		return nil, err
	}
	if bundle == nil {
		return nil, ErrNilBundle
	}
	if opts.Missing == MissingExplicitInput && len(bundle.MissingChatIds) > 0 {
		return nil, fmt.Errorf("%w: %v", ErrExplicitChatMissing, bundle.MissingChatIds)
	}

	for _, viewer := range bundle.ViewerChats {
		if viewer == nil || viewer.ViewerUserId != viewerUserID {
			continue
		}
		if opts.RequireNonEmpty && len(viewer.Chats) == 0 {
			return nil, ErrViewerProjectionEmpty
		}
		return viewer.Chats, nil
	}

	return nil, ErrViewerProjectionMissing
}

func ProjectMutableChat(mutableChat tg.MutableChatClazz, viewerUserID int64) (tg.ChatClazz, error) {
	if mutableChat == nil || mutableChat.Chat == nil {
		return nil, nil
	}
	chat := mutableChat.Chat
	date, err := int32ChatDate(chat.Id, chat.Date)
	if err != nil {
		return nil, err
	}

	return tg.MakeTLChat(&tg.TLChat{
		Creator:             chat.Creator == viewerUserID,
		Deactivated:         chat.Deactivated,
		CallActive:          chat.CallActive,
		CallNotEmpty:        chat.CallNotEmpty,
		Noforwards:          chat.Noforwards,
		Id:                  chat.Id,
		Title:               chat.Title,
		Photo:               ProjectChatPhoto(chat.Photo),
		ParticipantsCount:   chat.ParticipantsCount,
		Date:                date,
		Version:             chat.Version,
		MigratedTo:          chat.MigratedTo,
		DefaultBannedRights: chat.DefaultBannedRights,
	}), nil
}

func ProjectMutableChatList(chats []tg.MutableChatClazz, viewerUserID int64) ([]tg.ChatClazz, error) {
	var projected []tg.ChatClazz
	for _, mutableChat := range chats {
		chat, err := ProjectMutableChat(mutableChat, viewerUserID)
		if err != nil {
			return nil, err
		}
		if chat != nil {
			projected = append(projected, chat)
		}
	}
	return projected, nil
}

func ProjectMutableChatForViewers(chats []tg.MutableChatClazz, viewerUserIDs []int64) ([]chatpb.ViewerChatsClazz, error) {
	if len(viewerUserIDs) == 0 {
		return nil, nil
	}

	viewerChats := make([]chatpb.ViewerChatsClazz, 0, len(viewerUserIDs))
	for _, viewerUserID := range viewerUserIDs {
		projected, err := ProjectMutableChatList(chats, viewerUserID)
		if err != nil {
			return nil, err
		}
		viewerChats = append(viewerChats, chatpb.MakeTLViewerChats(&chatpb.TLViewerChats{
			ViewerUserId: viewerUserID,
			Chats:        projected,
		}))
	}
	return viewerChats, nil
}

func ProjectChatPhoto(photo tg.PhotoClazz) tg.ChatPhotoClazz {
	if p, ok := photo.(*tg.TLPhoto); ok && p != nil {
		return tg.MakeTLChatPhoto(&tg.TLChatPhoto{
			PhotoId: p.Id,
			DcId:    p.DcId,
		})
	}
	return tg.MakeTLChatPhotoEmpty(&tg.TLChatPhotoEmpty{})
}

func FillDifferenceChats(ctx context.Context, client Client, viewerUserID int64, diff *tg.UpdatesDifference, opts Options) error {
	if diff == nil {
		return nil
	}
	ids := tg.CollectChatIDsFromDifference(diff)
	chats, err := ProjectChats(ctx, client, viewerUserID, ids, opts)
	if err != nil {
		return err
	}
	if full, ok := diff.ToUpdatesDifference(); ok {
		full.Chats = chats
	}
	if slice, ok := diff.ToUpdatesDifferenceSlice(); ok {
		slice.Chats = chats
	}
	return nil
}

func int32ChatDate(chatID, date int64) (int32, error) {
	if date < minInt32 || date > maxInt32 {
		return 0, fmt.Errorf("%w: chat_id=%d date=%d", ErrChatDateOverflow, chatID, date)
	}
	return int32(date), nil
}
