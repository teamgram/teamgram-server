package repository

import (
	"context"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
)

func (r *Repository) GetUserProjectionBundle(ctx context.Context, viewerUserIds []int64, targetUserIds []int64, withFacts bool) (*UserProjectionBundle, error) {
	return nil, userpb.ErrUserInvalidArgument
}
