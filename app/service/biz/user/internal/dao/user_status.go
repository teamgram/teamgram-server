// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/marmota/pkg/threading2"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
)

const (
	userPresencesKeyPrefix = "user_presences"
)

func genUserPresencesKey(userId int64) string {
	return fmt.Sprintf("%s_%d", userPresencesKeyPrefix, userId)
}

func isUserPresencesKey(k string) bool {
	return strings.HasPrefix(k, userPresencesKeyPrefix)
}

func parseUserPresencesKey(k string) int64 {
	if strings.HasPrefix(k, userPresencesKeyPrefix+"_") {
		v, _ := strconv.ParseInt(k[len(userPresencesKeyPrefix)+1:], 10, 64)
		return v
	}

	return 0
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
			if do2 != nil {
				*v.(*dataobject.UserPresencesDO) = *do2
			} else {
				return sqlc.ErrNotFound
			}
			return nil
		})
	if err != nil {
		return nil, err
	}

	return do, nil
}

func (d *Dao) PutLastSeenAt(ctx context.Context, userId int64, lastSeenAt int64, expires int32) {
	do := &dataobject.UserPresencesDO{
		UserId:     userId,
		LastSeenAt: lastSeenAt,
		Expires:    expires,
	}

	d.CachedConn.SetCache(ctx, genUserPresencesKey(userId), do)
	threading2.WrapperGoFunc(ctx, nil, func(ctx context.Context) {
		d.UserPresencesDAO.InsertOrUpdate(ctx, do)
	})
}
