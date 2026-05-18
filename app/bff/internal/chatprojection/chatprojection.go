package chatprojection

import (
	"context"

	public "github.com/teamgram/teamgram-server/v2/app/bff/projection/chatprojection"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type Client = public.Client

type MissingPolicy = public.MissingPolicy

const (
	MissingExplicitInput   = public.MissingExplicitInput
	MissingStoredReference = public.MissingStoredReference
)

func ProjectChats(ctx context.Context, client Client, viewerUserId int64, targetChatIds []int64, missing MissingPolicy) ([]tg.ChatClazz, error) {
	return public.ProjectChats(ctx, client, viewerUserId, targetChatIds, missing)
}

func ProjectMutableChat(chat tg.MutableChatClazz, viewerUserId int64) (tg.ChatClazz, error) {
	return public.ProjectMutableChat(chat, viewerUserId)
}

func FillUpdatesChats(ctx context.Context, client Client, viewerUserId int64, updates *tg.Updates, missing MissingPolicy) error {
	return public.FillUpdatesChats(ctx, client, viewerUserId, updates, missing)
}

func FillDifferenceChats(ctx context.Context, client Client, viewerUserId int64, diff *tg.UpdatesDifference, missing MissingPolicy) error {
	return public.FillDifferenceChats(ctx, client, viewerUserId, diff, missing)
}
