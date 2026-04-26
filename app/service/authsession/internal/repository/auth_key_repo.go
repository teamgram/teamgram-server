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

	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
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

func (r *Repository) QueryAuthKey(ctx context.Context, authKeyId int64) (*tg.AuthKeyInfo, error) {
	key, err := r.model.AuthKeysModel.FindOneByAuthKeyId(ctx, authKeyId)
	if err != nil {
		if isNotFound(err) {
			return nil, authsession.ErrAuthKeyNotFound
		}
		return nil, wrapStorage(err)
	}
	if key == nil {
		return nil, authsession.ErrAuthKeyNotFound
	}

	keyInfo, err := toAuthKeyInfo(key)
	if err != nil {
		return nil, authsession.ErrAuthKeyInvalid
	}
	return keyInfo, nil
}

func (r *Repository) ListAuthKeysByIds(ctx context.Context, authKeyIds []int64) (map[int64]*tg.AuthKeyInfo, error) {
	if len(authKeyIds) == 0 {
		return map[int64]*tg.AuthKeyInfo{}, nil
	}

	rows, err := r.model.AuthKeysModel.FindListByAuthKeyIdList(ctx, authKeyIds...)
	if err != nil {
		return nil, wrapStorage(err)
	}

	keys := make(map[int64]*tg.AuthKeyInfo, len(rows))
	for i := range rows {
		keyInfo, err := toAuthKeyInfo(&rows[i])
		if err != nil {
			return nil, authsession.ErrAuthKeyInvalid
		}
		keys[keyInfo.AuthKeyId] = keyInfo
	}
	return keys, nil
}

func (r *Repository) ExpandAuthKeyIds(ctx context.Context, authKeyIds []int64) ([]int64, error) {
	keys, err := r.ListAuthKeysByIds(ctx, authKeyIds)
	if err != nil {
		return nil, err
	}

	expandedKeyIds := make([]int64, 0, len(authKeyIds))
	for _, authKeyId := range authKeyIds {
		keyData := keys[authKeyId]
		if keyData == nil {
			return nil, authsession.ErrAuthKeyNotFound
		}
		if keyData.TempAuthKeyId != 0 {
			expandedKeyIds = append(expandedKeyIds, keyData.TempAuthKeyId)
		} else {
			expandedKeyIds = append(expandedKeyIds, authKeyId)
		}
	}
	return expandedKeyIds, nil
}

func (r *Repository) SaveAuthKey(ctx context.Context, authKey *tg.AuthKeyInfo, expiredIn int32) error {
	// TODO(@benqi): expiredIn
	_ = expiredIn

	key := fromAuthKeyInfo(authKey)
	_, _, err := r.model.AuthKeysModel.InsertIgnore(ctx, key)
	if err != nil {
		return wrapStorage(err)
	}

	return nil
}

func (r *Repository) BindKeyId(ctx context.Context, keyId int64, bindType int32, bindKeyId int64) error {
	var err error
	switch bindType {
	case tg.AuthKeyTypePerm:
		_, err = r.model.AuthKeysModel.UpdatePermBinding(ctx, bindKeyId, keyId)
	case tg.AuthKeyTypeTemp:
		_, err = r.model.AuthKeysModel.UpdateTempBinding(ctx, bindKeyId, keyId)
	case tg.AuthKeyTypeMediaTemp:
		_, err = r.model.AuthKeysModel.UpdateMediaTempBinding(ctx, bindKeyId, keyId)
	default:
		return authsession.ErrAuthKeyInvalid
	}
	if err != nil {
		return wrapStorage(err)
	}
	return nil
}

func (r *Repository) ResolvePermAuthKey(ctx context.Context, authKeyId int64) (*tg.AuthKeyInfo, error) {
	keyData, err := r.QueryAuthKey(ctx, authKeyId)
	if err != nil {
		return nil, err
	}
	if keyData.PermAuthKeyId == 0 {
		return nil, authsession.ErrPermAuthKeyEmpty
	}
	return keyData, nil
}

func (r *Repository) GetPermAuthKeyIdByAuthKeyId(ctx context.Context, authKeyId int64) (int64, error) {
	keyData, err := r.ResolvePermAuthKey(ctx, authKeyId)
	if err != nil {
		return 0, err
	}
	return keyData.PermAuthKeyId, nil
}
