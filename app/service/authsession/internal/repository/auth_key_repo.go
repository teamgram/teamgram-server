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
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// defaultAuthKeyTTL is the lifetime applied to temporary / media-temporary
// auth keys when the caller does not supply an explicit expires_in. Seven
// days matches the upper bound used by official Telegram clients for cached
// temp keys, and it bounds how long an orphaned row can stay reachable.
const defaultAuthKeyTTL = 7 * 24 * 60 * 60

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

	if isTempAuthKeyType(keyInfo.AuthKeyType) {
		if key.ExpiresAt > 0 {
			if time.Now().UTC().Unix() >= key.ExpiresAt {
				if err := r.allowExpiredBoundTempAuthKey(ctx, keyInfo); err != nil {
					return nil, err
				}
				return keyInfo, nil
			}
			return keyInfo, nil
		}
		active, err := r.checkAuthKeyActive(ctx, authKeyId)
		if err != nil {
			return nil, wrapStorage(err)
		}
		if !active {
			return nil, authsession.ErrAuthKeyNotFound
		}
	}
	return keyInfo, nil
}

func (r *Repository) allowExpiredBoundTempAuthKey(ctx context.Context, keyInfo *tg.AuthKeyInfo) error {
	if keyInfo.PermAuthKeyId == 0 {
		return authsession.ErrAuthKeyNotFound
	}

	permKey, err := r.model.AuthKeysModel.FindOneByAuthKeyId(ctx, keyInfo.PermAuthKeyId)
	if err != nil {
		if isNotFound(err) {
			return authsession.ErrAuthKeyNotFound
		}
		return wrapStorage(err)
	}
	if permKey == nil || permKey.AuthKeyType != tg.AuthKeyTypePerm {
		return authsession.ErrAuthKeyNotFound
	}
	return nil
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

// SaveAuthKey persists an auth key and, for temp / media-temp keys, stores a
// durable expires_at timestamp. The lifecycle store remains a compatibility
// fallback for legacy rows whose expires_at is still zero.
//
// expiredIn is interpreted as seconds, matching the MTProto semantics for
// temp keys. A non-positive value falls back to defaultAuthKeyTTL (7 days)
// to bound orphaned rows even when the caller forgets to supply one.
// Permanent keys do not register a TTL; they live until explicit revocation.
func (r *Repository) SaveAuthKey(ctx context.Context, authKey *tg.AuthKeyInfo, expiredIn int32) error {
	key := fromAuthKeyInfo(authKey)
	ttl := int(expiredIn)
	if isTempAuthKeyType(authKey.AuthKeyType) {
		if ttl <= 0 {
			ttl = defaultAuthKeyTTL
		}
		key.ExpiresAt = time.Now().UTC().Unix() + int64(ttl)
	}
	if _, _, err := r.model.AuthKeysModel.InsertIgnore(ctx, key); err != nil {
		return wrapStorage(err)
	}

	if !isTempAuthKeyType(authKey.AuthKeyType) {
		return nil
	}

	if err := r.activateAuthKey(ctx, authKey.AuthKeyId, ttl); err != nil {
		return wrapStorage(err)
	}
	return nil
}

func isTempAuthKeyType(authKeyType int32) bool {
	return authKeyType == tg.AuthKeyTypeTemp || authKeyType == tg.AuthKeyTypeMediaTemp
}

func (r *Repository) activateAuthKey(ctx context.Context, authKeyId int64, ttlSeconds int) error {
	if r.authKeyLifecycleModel == nil {
		return nil
	}
	return r.authKeyLifecycleModel.Activate(ctx, authKeyId, ttlSeconds)
}

func (r *Repository) checkAuthKeyActive(ctx context.Context, authKeyId int64) (bool, error) {
	if r.authKeyLifecycleModel == nil {
		return true, nil
	}
	return r.authKeyLifecycleModel.IsActive(ctx, authKeyId)
}

// BindKeyId records a perm <-> temp / perm <-> media-temp pairing on the
// auth_keys row identified by keyId.
//
// Both keyId and bindKeyId must be non-zero — bindKeyId == 0 used to be
// silently accepted, which would clear the binding column without
// invalidating the corresponding caches. Callers that genuinely want to
// drop a binding should add a dedicated unbind path rather than relying on
// the zero value here.
func (r *Repository) BindKeyId(ctx context.Context, keyId int64, bindType int32, bindKeyId int64) error {
	if keyId == 0 || bindKeyId == 0 {
		return authsession.ErrAuthKeyInvalid
	}

	var (
		rowsAffected int64
		err          error
	)
	switch bindType {
	case tg.AuthKeyTypePerm:
		rowsAffected, err = r.model.AuthKeysModel.UpdatePermBinding(ctx, bindKeyId, keyId)
	case tg.AuthKeyTypeTemp:
		rowsAffected, err = r.model.AuthKeysModel.UpdateTempBinding(ctx, bindKeyId, keyId)
	case tg.AuthKeyTypeMediaTemp:
		rowsAffected, err = r.model.AuthKeysModel.UpdateMediaTempBinding(ctx, bindKeyId, keyId)
	default:
		return authsession.ErrAuthKeyInvalid
	}
	if err != nil {
		if isNotFound(err) {
			return authsession.ErrAuthKeyNotFound
		}
		return wrapStorage(err)
	}
	if rowsAffected == 0 {
		return authsession.ErrAuthKeyNotFound
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
