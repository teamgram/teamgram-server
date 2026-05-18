package userprojection

import (
	"context"

	public "github.com/teamgram/teamgram-server/v2/app/bff/projection/userprojection"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type UserClient = public.UserClient

type MissingPolicy = public.MissingPolicy

const (
	MissingExplicitInput   = public.MissingExplicitInput
	MissingStoredReference = public.MissingStoredReference
)

func ProjectUsers(ctx context.Context, client UserClient, viewerUserId int64, targetUserIds []int64, missing MissingPolicy) ([]tg.UserClazz, error) {
	return public.ProjectUsers(ctx, client, viewerUserId, targetUserIds, missing)
}

func FillUpdatesUsers(ctx context.Context, client UserClient, viewerUserId int64, updates *tg.Updates, missing MissingPolicy) error {
	return public.FillUpdatesUsers(ctx, client, viewerUserId, updates, missing)
}

func FillDifferenceUsers(ctx context.Context, client UserClient, viewerUserId int64, diff *tg.UpdatesDifference, missing MissingPolicy) error {
	return public.FillDifferenceUsers(ctx, client, viewerUserId, diff, missing)
}

func FillMessagesMessagesUsers(ctx context.Context, client UserClient, viewerUserId int64, messages *tg.MessagesMessages, missing MissingPolicy) error {
	return public.FillMessagesMessagesUsers(ctx, client, viewerUserId, messages, missing)
}
