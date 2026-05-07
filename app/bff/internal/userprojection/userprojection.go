package userprojection

import (
	"context"
	"errors"
	"fmt"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type UserClient interface {
	UserGetUserProjectionBundle(ctx context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error)
}

type MissingPolicy int

const (
	MissingExplicitInput MissingPolicy = iota
	MissingStoredReference
)

func ProjectUsers(ctx context.Context, client UserClient, viewerUserId int64, targetUserIds []int64, missing MissingPolicy) ([]tg.UserClazz, error) {
	if len(targetUserIds) == 0 {
		return []tg.UserClazz{}, nil
	}
	if client == nil {
		return nil, fmt.Errorf("user projection client is nil")
	}
	bundle, err := client.UserGetUserProjectionBundle(ctx, &userpb.TLUserGetUserProjectionBundle{
		ViewerUserIds: []int64{viewerUserId},
		TargetUserIds: targetUserIds,
	})
	if err != nil {
		if errors.Is(err, userpb.ErrUserInvalidArgument) {
			return nil, fmt.Errorf("user projection invalid request: %w", err)
		}
		return nil, err
	}
	if bundle == nil {
		return nil, fmt.Errorf("user projection bundle is nil")
	}
	if len(bundle.MissingUserIds) > 0 && missing == MissingExplicitInput {
		return nil, tg.ErrUserIdInvalid
	}
	for _, viewer := range bundle.ViewerUsers {
		if viewer != nil && viewer.ViewerUserId == viewerUserId {
			return viewer.Users, nil
		}
	}
	return []tg.UserClazz{}, nil
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
