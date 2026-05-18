package userprojection

import (
	"context"
	"errors"
	"fmt"

	bizuserproj "github.com/teamgram/teamgram-server/v2/app/service/biz/user/userprojection"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UserClient is the BFF compatibility alias for the public user projection client.
//
// New cross-service callers should import
// app/service/biz/user/userprojection directly instead of this package.
type UserClient = bizuserproj.Client

type MissingPolicy int

const (
	MissingExplicitInput MissingPolicy = iota
	MissingStoredReference
)

func publicMissingPolicy(missing MissingPolicy) bizuserproj.MissingPolicy {
	switch missing {
	case MissingExplicitInput:
		return bizuserproj.MissingExplicitInput
	case MissingStoredReference:
		return bizuserproj.MissingStoredReference
	default:
		return bizuserproj.MissingStoredReference
	}
}

// ProjectUsers adapts the public user projection helper to BFF legacy semantics.
//
// New cross-service callers should import
// app/service/biz/user/userprojection directly instead of this package.
func ProjectUsers(ctx context.Context, client UserClient, viewerUserId int64, targetUserIds []int64, missing MissingPolicy) ([]tg.UserClazz, error) {
	users, err := bizuserproj.ProjectUsers(ctx, client, viewerUserId, targetUserIds, bizuserproj.Options{
		Missing:         publicMissingPolicy(missing),
		RequireNonEmpty: false,
	})
	if err != nil {
		switch {
		case errors.Is(err, bizuserproj.ErrExplicitUserMissing):
			return nil, tg.ErrUserIdInvalid
		case errors.Is(err, bizuserproj.ErrInvalidRequest):
			return nil, fmt.Errorf("user projection invalid request: %w", err)
		case errors.Is(err, bizuserproj.ErrViewerProjectionMissing):
			return []tg.UserClazz{}, nil
		default:
			return nil, err
		}
	}
	if users == nil {
		return []tg.UserClazz{}, nil
	}
	return users, nil
}

func FillUpdatesUsers(ctx context.Context, client UserClient, viewerUserId int64, updates *tg.Updates, missing MissingPolicy) error {
	if updates == nil {
		return nil
	}
	ids := tg.CollectUserIDsFromUpdates(updates)
	users, err := ProjectUsers(ctx, client, viewerUserId, ids, missing)
	if err != nil {
		return err
	}
	if full, ok := updates.ToUpdates(); ok {
		full.Users = users
	}
	if combined, ok := updates.ToUpdatesCombined(); ok {
		combined.Users = users
	}
	return nil
}

func FillDifferenceUsers(ctx context.Context, client UserClient, viewerUserId int64, diff *tg.UpdatesDifference, missing MissingPolicy) error {
	if diff == nil {
		return nil
	}
	ids := tg.CollectUserIDsFromDifference(diff)
	users, err := ProjectUsers(ctx, client, viewerUserId, ids, missing)
	if err != nil {
		return err
	}
	if full, ok := diff.ToUpdatesDifference(); ok {
		full.Users = users
	}
	if slice, ok := diff.ToUpdatesDifferenceSlice(); ok {
		slice.Users = users
	}
	return nil
}

func FillMessagesMessagesUsers(ctx context.Context, client UserClient, viewerUserId int64, messages *tg.MessagesMessages, missing MissingPolicy) error {
	if messages == nil {
		return nil
	}
	ids := tg.CollectUserIDsFromMessagesMessages(messages)
	users, err := ProjectUsers(ctx, client, viewerUserId, ids, missing)
	if err != nil {
		return err
	}
	if full, ok := messages.ToMessagesMessages(); ok {
		full.Users = users
	}
	if slice, ok := messages.ToMessagesMessagesSlice(); ok {
		slice.Users = users
	}
	if channel, ok := messages.ToMessagesChannelMessages(); ok {
		channel.Users = users
	}
	return nil
}
