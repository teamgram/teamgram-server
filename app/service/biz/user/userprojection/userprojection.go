package userprojection

import (
	"context"
	"errors"
	"fmt"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type Client interface {
	UserGetUserProjectionBundle(ctx context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error)
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
	ErrClientNotConfigured     = errors.New("user projection: client not configured")
	ErrInvalidRequest          = errors.New("user projection: invalid request")
	ErrNilBundle               = errors.New("user projection: nil bundle")
	ErrExplicitUserMissing     = errors.New("user projection: explicit user missing")
	ErrViewerProjectionMissing = errors.New("user projection: viewer projection missing")
	ErrViewerProjectionEmpty   = errors.New("user projection: viewer projection empty")
)

func ProjectUsers(ctx context.Context, client Client, viewerUserId int64, targetUserIds []int64, opts Options) ([]tg.UserClazz, error) {
	if len(targetUserIds) == 0 {
		return nil, nil
	}
	if client == nil {
		return nil, ErrClientNotConfigured
	}

	bundle, err := client.UserGetUserProjectionBundle(ctx, &userpb.TLUserGetUserProjectionBundle{
		ViewerUserIds: []int64{viewerUserId},
		TargetUserIds: targetUserIds,
	})
	if err != nil {
		if errors.Is(err, userpb.ErrUserInvalidArgument) {
			return nil, fmt.Errorf("%w: %w", ErrInvalidRequest, err)
		}
		return nil, err
	}
	if bundle == nil {
		return nil, ErrNilBundle
	}
	if opts.Missing == MissingExplicitInput && len(bundle.MissingUserIds) > 0 {
		return nil, fmt.Errorf("%w: %v", ErrExplicitUserMissing, bundle.MissingUserIds)
	}

	for _, viewer := range bundle.ViewerUsers {
		if viewer == nil || viewer.ViewerUserId != viewerUserId {
			continue
		}
		if opts.RequireNonEmpty && len(viewer.Users) == 0 {
			return nil, ErrViewerProjectionEmpty
		}
		return viewer.Users, nil
	}

	return nil, ErrViewerProjectionMissing
}

func FillUpdatesUsers(ctx context.Context, client Client, viewerUserId int64, updates *tg.Updates, opts Options) error {
	if updates == nil {
		return nil
	}
	ids := tg.CollectUserIDsFromUpdates(updates)
	users, err := ProjectUsers(ctx, client, viewerUserId, ids, opts)
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

func FillDifferenceUsers(ctx context.Context, client Client, viewerUserId int64, diff *tg.UpdatesDifference, opts Options) error {
	if diff == nil {
		return nil
	}
	ids := tg.CollectUserIDsFromDifference(diff)
	users, err := ProjectUsers(ctx, client, viewerUserId, ids, opts)
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

func FillMessagesMessagesUsers(ctx context.Context, client Client, viewerUserId int64, messages *tg.MessagesMessages, opts Options) error {
	if messages == nil {
		return nil
	}
	ids := tg.CollectUserIDsFromMessagesMessages(messages)
	users, err := ProjectUsers(ctx, client, viewerUserId, ids, opts)
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
