package chatprojection

import (
	"context"
	"errors"
	"fmt"

	bizchatproj "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chatprojection"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type Client = bizchatproj.Client

type MissingPolicy int

const (
	MissingExplicitInput MissingPolicy = iota
	MissingStoredReference
)

func publicMissingPolicy(missing MissingPolicy) bizchatproj.MissingPolicy {
	switch missing {
	case MissingExplicitInput:
		return bizchatproj.MissingExplicitInput
	case MissingStoredReference:
		return bizchatproj.MissingStoredReference
	default:
		return bizchatproj.MissingStoredReference
	}
}

func ProjectChats(ctx context.Context, client Client, viewerUserId int64, targetChatIds []int64, missing MissingPolicy) ([]tg.ChatClazz, error) {
	chats, err := bizchatproj.ProjectChats(ctx, client, viewerUserId, targetChatIds, bizchatproj.Options{
		Missing:         publicMissingPolicy(missing),
		RequireNonEmpty: false,
	})
	if err != nil {
		switch {
		case errors.Is(err, bizchatproj.ErrExplicitChatMissing):
			return nil, tg.ErrChatIdInvalid
		case errors.Is(err, bizchatproj.ErrInvalidRequest):
			return nil, fmt.Errorf("chat projection invalid request: %w", err)
		case errors.Is(err, bizchatproj.ErrViewerProjectionMissing):
			return []tg.ChatClazz{}, nil
		default:
			return nil, err
		}
	}
	if chats == nil {
		return []tg.ChatClazz{}, nil
	}
	return chats, nil
}

func ProjectMutableChat(chat tg.MutableChatClazz, viewerUserId int64) (tg.ChatClazz, error) {
	return bizchatproj.ProjectMutableChat(chat, viewerUserId)
}

func FillDifferenceChats(ctx context.Context, client Client, viewerUserId int64, diff *tg.UpdatesDifference, missing MissingPolicy) error {
	if diff == nil {
		return nil
	}
	ids := tg.CollectChatIDsFromDifference(diff)
	chats, err := ProjectChats(ctx, client, viewerUserId, ids, missing)
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
