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
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"

	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/contextx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/threading"
)

const (
	cacheAuthKeyPrefix   = "auth_keys"
	cacheAuthKeyV2Prefix = "auth_keys2"
)

func genCacheAuthKeyKey(id int64) string {
	return fmt.Sprintf("%s_%d", cacheAuthKeyPrefix, id)
}

func genCacheAuthKeyV2Key(id int64) string {
	return fmt.Sprintf("%s_%d", cacheAuthKeyV2Prefix, id)
}

func (d *Dao) getAuthKey(ctx context.Context, keyId int64) (keyData *tg.AuthKeyInfo, err error) {
	var (
		key    = genCacheAuthKeyKey(keyId)
		values map[string]string
	)

	values, err = d.kv.HgetallCtx(ctx, key)
	if err != nil {
		logx.WithContext(ctx).Errorf("conn.Do(HGETALL %s) error(%v)", key, err)
		return
	} else if len(values) == 0 {
		err = sqlc.ErrNotFound
		return
	}

	// TODO(@benqi): check len(values)

	keyData2 := &tg.TLAuthKeyInfo{}
	// keyData2, _ := keyData.ToAuthKeyInfo()
	for k, v := range values {
		switch k {
		case "auth_key_type":
			authKeyType, _ := strconv.Atoi(v)
			keyData2.AuthKeyType = int32(authKeyType)
		case "auth_key_id":
			keyData2.AuthKeyId, _ = strconv.ParseInt(v, 10, 64)
		case "auth_key":
			keyData2.AuthKey = []byte(v)
		case "perm_auth_key_id":
			keyData2.PermAuthKeyId, _ = strconv.ParseInt(v, 10, 64)
		case "temp_auth_key_id":
			keyData2.TempAuthKeyId, _ = strconv.ParseInt(v, 10, 64)
		case "media_temp_auth_key_id":
			keyData2.MediaTempAuthKeyId, _ = strconv.ParseInt(v, 10, 64)
		}
	}

	keyData = keyData2.ToAuthKeyInfo()
	return
}

func (d *Dao) QueryAuthKeyV2(ctx context.Context, authKeyId int64) (*tg.TLAuthKeyInfo, error) {
	var (
		keyInfo = &tg.TLAuthKeyInfo{
			AuthKeyId:          authKeyId,
			AuthKey:            nil,
			AuthKeyType:        0,
			PermAuthKeyId:      0,
			TempAuthKeyId:      0,
			MediaTempAuthKeyId: 0,
		}
	)

	err := d.CachedConn.QueryRow(
		ctx,
		keyInfo,
		genCacheAuthKeyV2Key(authKeyId),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			kInfo := v.(*tg.TLAuthKeyInfo)
			err := mr.Finish(
				func() error {
					do1, err2 := d.AuthKeyInfosDAO.SelectByAuthKeyId(ctx, authKeyId)
					if err2 != nil {
						return err2
					} else if do1 == nil {
						return sqlc.ErrNotFound
					}
					kInfo.AuthKeyType = do1.AuthKeyType
					kInfo.PermAuthKeyId = do1.PermAuthKeyId
					kInfo.TempAuthKeyId = do1.TempAuthKeyId
					kInfo.MediaTempAuthKeyId = do1.MediaTempAuthKeyId

					return nil
				},
				func() error {
					do2, err2 := d.AuthKeysDAO.SelectByAuthKeyId(ctx, authKeyId)
					if err2 != nil {
						return err2
					} else if do2 == nil {
						return sqlc.ErrNotFound
					}

					kInfo.AuthKey, err2 = base64.RawStdEncoding.DecodeString(do2.Body)
					if err2 != nil {
						return err2
					}

					return nil
				})
			if err != nil && errors.Is(err, sqlc.ErrNotFound) {
				kInfo2, _ := d.getAuthKey(ctx, authKeyId)
				kInfo2.Match(
					func(c *tg.TLAuthKeyInfo) interface{} {
						kInfo.AuthKeyType = c.AuthKeyType
						kInfo.AuthKey = c.AuthKey
						kInfo.TempAuthKeyId = c.TempAuthKeyId
						kInfo.PermAuthKeyId = c.PermAuthKeyId
						kInfo.MediaTempAuthKeyId = c.MediaTempAuthKeyId

						threading.GoSafe(func() {
							_, _, _ = d.AuthKeyInfosDAO.Insert(
								contextx.ValueOnlyFrom(ctx),
								&dataobject.AuthKeyInfosDO{
									AuthKeyId:          keyInfo.AuthKeyId,
									AuthKeyType:        keyInfo.AuthKeyType,
									PermAuthKeyId:      keyInfo.PermAuthKeyId,
									TempAuthKeyId:      keyInfo.TempAuthKeyId,
									MediaTempAuthKeyId: keyInfo.MediaTempAuthKeyId,
									Deleted:            false,
								})
						})

						err = nil
						return nil
					},
				)
			}
			return err
		})
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			err = tg.ErrAuthKeyUnregistered
		} else {
			err = tg.ErrInternalServerError
		}
		return nil, err
	}

	return keyInfo, nil
}

func (d *Dao) SetAuthKeyV2(ctx context.Context, authKey *tg.TLAuthKeyInfo, expiredIn int32) (err error) {
	// TODO(@benqi): expiredIn
	_ = expiredIn

	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		_, _, err = d.AuthKeysDAO.InsertTx(tx, &dataobject.AuthKeysDO{
			AuthKeyId: authKey.AuthKeyId,
			Body:      base64.RawStdEncoding.EncodeToString(authKey.AuthKey),
		})
		if err != nil {
			result.Err = err
			return
		}
		_, _, err = d.AuthKeyInfosDAO.InsertTx(tx, &dataobject.AuthKeyInfosDO{
			AuthKeyId:          authKey.AuthKeyId,
			AuthKeyType:        authKey.AuthKeyType,
			PermAuthKeyId:      authKey.PermAuthKeyId,
			TempAuthKeyId:      authKey.TempAuthKeyId,
			MediaTempAuthKeyId: authKey.MediaTempAuthKeyId,
			Deleted:            false,
		})
		if err != nil {
			result.Err = err
		}
	})

	return tR.Err
}

func (d *Dao) UnsafeBindKeyIdV2(ctx context.Context, keyId int64, bindType int32, bindKeyId int64) (err error) {
	var (
		key = genCacheAuthKeyV2Key(keyId)
	)

	_, _, err = d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			//
			var (
				err2 error
			)

			switch bindType {
			case tg.AuthKeyTypePerm:
				_, err2 = d.AuthKeyInfosDAO.UpdateCustomMap(
					ctx,
					map[string]interface{}{
						"perm_auth_key_id": bindKeyId,
					},
					keyId)
			case tg.AuthKeyTypeTemp:
				_, err2 = d.AuthKeyInfosDAO.UpdateCustomMap(
					ctx,
					map[string]interface{}{
						"temp_auth_key_id": bindKeyId,
					},
					keyId)
			case tg.AuthKeyTypeMediaTemp:
				_, err2 = d.AuthKeyInfosDAO.UpdateCustomMap(
					ctx,
					map[string]interface{}{
						"media_temp_auth_key_id": bindKeyId,
					},
					keyId)
			}

			return 0, 0, err2
		},
		key)

	return
}

func (d *Dao) GetPermAuthKeyId(ctx context.Context, authKeyId int64) int64 {
	if k, err := d.QueryAuthKeyV2(ctx, authKeyId); err != nil || k == nil {
		return 0
	} else {
		return k.PermAuthKeyId
	}
}
