// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package repository

import (
	"context"
	"encoding/base64"

	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func toAuthKeyInfo(do *model.AuthKeys) (*tg.AuthKeyInfo, error) {
	keyData, err := base64.RawStdEncoding.DecodeString(do.Body)
	if err != nil {
		return nil, err
	}

	return tg.MakeTLAuthKeyInfo(&tg.TLAuthKeyInfo{
		AuthKeyId:          do.AuthKeyId,
		AuthKey:            keyData,
		AuthKeyType:        do.AuthKeyType,
		PermAuthKeyId:      do.PermAuthKeyId,
		TempAuthKeyId:      do.TempAuthKeyId,
		MediaTempAuthKeyId: do.MediaTempAuthKeyId,
	}), nil
}

func fromAuthKeyInfo(v *tg.AuthKeyInfo) *model.AuthKeys {
	return &model.AuthKeys{
		AuthKeyId:          v.AuthKeyId,
		Body:               base64.RawStdEncoding.EncodeToString(v.AuthKey),
		AuthKeyType:        v.AuthKeyType,
		PermAuthKeyId:      v.PermAuthKeyId,
		TempAuthKeyId:      v.TempAuthKeyId,
		MediaTempAuthKeyId: v.MediaTempAuthKeyId,
		Deleted:            false,
	}
}

func (r *Repository) QueryAuthKeyV2(ctx context.Context, authKeyId int64) (*tg.AuthKeyInfo, error) {
	key, err := r.AuthKeysModel.FindOneByAuthKeyId(ctx, authKeyId)

	if err != nil {
		return nil, err
	}

	return toAuthKeyInfo(key)
}

func (r *Repository) SetAuthKeyV2(ctx context.Context, authKey *tg.AuthKeyInfo, expiredIn int32) (err error) {
	// TODO(@benqi): expiredIn
	_ = expiredIn

	key := fromAuthKeyInfo(authKey)
	_, err = r.AuthKeysModel.Insert2(ctx, key)

	return
}
