package core

import (
	"context"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
)

type activePermAuthKeySource interface {
	AuthsessionGetPermAuthKeyIds(ctx context.Context, in *authsession.TLAuthsessionGetPermAuthKeyIds) (*authsession.VectorLong, error)
}

func resolveAuthSeqNotMeTargets(ctx context.Context, source activePermAuthKeySource, userID int64, sourcePermAuthKeyID int64) ([]int64, error) {
	if source == nil {
		return nil, fmt.Errorf("%w: target_lookup: authsession source is nil", userupdates.ErrUserupdatesStorage)
	}
	keys, err := source.AuthsessionGetPermAuthKeyIds(ctx, &authsession.TLAuthsessionGetPermAuthKeyIds{UserId: userID})
	if err != nil {
		return nil, fmt.Errorf("%w: target_lookup: %w", userupdates.ErrUserupdatesStorage, err)
	}
	if keys == nil || len(keys.Datas) == 0 {
		return []int64{}, nil
	}
	seen := make(map[int64]struct{}, len(keys.Datas))
	targets := make([]int64, 0, len(keys.Datas))
	for _, keyID := range keys.Datas {
		if keyID == 0 || keyID == sourcePermAuthKeyID {
			continue
		}
		if _, ok := seen[keyID]; ok {
			continue
		}
		seen[keyID] = struct{}{}
		targets = append(targets, keyID)
	}
	return targets, nil
}
