// Copyright 2024 Teamgram Authors
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
	"strconv"

	"github.com/teamgram/proto/mtproto"
	sessionpb "github.com/teamgram/teamgram-server/app/interface/session/session"

	"github.com/zeromicro/go-zero/core/logx"
)

type CacheV struct {
	V *mtproto.AuthKeyInfo
}

func (c CacheV) Size() int {
	return 1
}

func (d *Dao) GetCacheAuthKey(ctx context.Context, authKeyId int64) (*mtproto.AuthKeyInfo, error) {
	var (
		cacheK = strconv.Itoa(int(authKeyId))
		value  *mtproto.AuthKeyInfo
	)

	if v, ok := d.cache.Get(cacheK); ok {
		value = v.(*CacheV).V
	} else {
		sessClient, err := d.session.getSessionClient(strconv.FormatInt(authKeyId, 10))
		if err != nil {
			logx.WithContext(ctx).Errorf("getSessionClient error: %v, {authKeyId: %d}", err, authKeyId)
			return nil, err
		} else {
			value, err = sessClient.SessionQueryAuthKey(ctx, &sessionpb.TLSessionQueryAuthKey{
				AuthKeyId: authKeyId,
			})
			if err != nil {
				logx.WithContext(ctx).Errorf("sessionQueryAuthKey - error: %v", err)
				return nil, err
			}
			d.PutAuthKey(value)
		}
	}

	return value, nil
}

func (d *Dao) PutAuthKey(keyInfo *mtproto.AuthKeyInfo) {
	var (
		cacheK = strconv.Itoa(int(keyInfo.AuthKeyId))
	)

	// TODO: expires_in
	d.cache.Set(cacheK, &CacheV{V: keyInfo})
}
