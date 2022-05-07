// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"context"
	"fmt"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/zeromicro/go-zero/core/mr"
)

const (
	userPresencesKeyPrefix = "user_presences"
)

func genUserPresencesKey(userId int64) string {
	return fmt.Sprintf("%s_%d", userPresencesKeyPrefix, userId)
}

func (d *Dao) GetLastSeenAt(ctx context.Context, id int64) (*dataobject.UserPresencesDO, error) {
	var (
		do = &dataobject.UserPresencesDO{}
	)

	err := d.CachedConn.QueryRow(
		ctx,
		do,
		genUserPresencesKey(id),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			do2, err := d.UserPresencesDAO.Select(ctx, id)
			if err != nil {
				return err
			}
			*v.(*dataobject.UserPresencesDO) = *do2
			return nil
		})
	if err != nil {
		return nil, err
	}

	return do, nil
}

func (d *Dao) PutLastSeenAt(ctx context.Context, lastSeenAt int64, expires int32, userId int64) {
	do := &dataobject.UserPresencesDO{
		UserId:     userId,
		LastSeenAt: lastSeenAt,
		Expires:    expires,
	}

	mr.FinishVoid(
		func() {
			d.UserPresencesDAO.UpdateLastSeenAt(ctx, lastSeenAt, expires, userId)
		},
		func() {
			d.CachedConn.SetCache(ctx, genUserPresencesKey(userId), do)
		})
}
