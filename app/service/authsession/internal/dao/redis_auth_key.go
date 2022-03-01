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

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	cacheAuthKeyPrefix = "auth_keys"
)

func genCacheAuthKeyKey(id int64) string {
	return fmt.Sprintf("%s_%d", cacheAuthKeyPrefix, id)
}

func (d *Dao) PutAuthKey(ctx context.Context, keyId int64, keyData *mtproto.AuthKeyInfo, expiredIn int32) (err error) {
	var (
		key = genCacheAuthKeyKey(keyId)

		args = map[string]string{
			"auth_key_type":          strconv.Itoa(int(keyData.AuthKeyType)),
			"auth_key_id":            strconv.FormatInt(keyData.AuthKeyId, 10),
			"auth_key":               string(keyData.AuthKey),
			"perm_auth_key_id":       strconv.FormatInt(keyData.PermAuthKeyId, 10),
			"temp_auth_key_id":       strconv.FormatInt(keyData.TempAuthKeyId, 10),
			"media_temp_auth_key_id": strconv.FormatInt(keyData.MediaTempAuthKeyId, 10),
		}
	)

	// TODO(@benqi): args error??
	if err = d.kv.Hmset(key, args); err != nil {
		logx.WithContext(ctx).Errorf("conn.Send(HMSET %s,%v) error(%v)", key, args, err)
		return
	}

	if expiredIn > 0 {
		if err = d.kv.Expire(key, int(expiredIn)); err != nil {
			logx.WithContext(ctx).Errorf("conn.Send(EXPIRE %d,%d) error(%v)", key, expiredIn, err)
		}
	}

	return
}

func (d *Dao) UnsafeBindKeyId(ctx context.Context, keyId int64, bindType int32, bindKeyId int64) (err error) {
	var (
		key = genCacheAuthKeyKey(keyId)
	)

	switch bindType {
	case mtproto.AuthKeyTypePerm:
		if err = d.kv.Hset(key, "perm_auth_key_id", strconv.FormatInt(bindKeyId, 10)); err != nil {
			logx.WithContext(ctx).Errorf("conn.Do(HSET %s,perm_auth_key_id,%d) error(%v)", key, bindKeyId, err)
		}
	case mtproto.AuthKeyTypeTemp:
		if err = d.kv.Hset(key, "temp_auth_key_id", strconv.FormatInt(bindKeyId, 10)); err != nil {
			logx.WithContext(ctx).Errorf("conn.Do(HSET %s,temp_auth_key_id,%d) error(%v)", key, bindKeyId, err)
		}
	case mtproto.AuthKeyTypeMediaTemp:
		if err = d.kv.Hset(key, "media_temp_auth_key_id", strconv.FormatInt(bindKeyId, 10)); err != nil {
			logx.WithContext(ctx).Errorf("conn.Do(HSET %s,media_temp_auth_key_id,%d) error(%v)", key, bindKeyId, err)
		}
	default:
		return
	}

	return
}

func (d *Dao) GetAuthKey(ctx context.Context, keyId int64) (keyData *mtproto.AuthKeyInfo, err error) {
	var (
		key    = genCacheAuthKeyKey(keyId)
		values map[string]string
	)

	values, err = d.kv.Hgetall(key)
	if err != nil {
		logx.WithContext(ctx).Errorf("conn.Do(HGETALL %s) error(%v)", key, err)
		return
	} else if len(values) == 0 {
		err = fmt.Errorf("invalid auth_key")
		return
	}

	// TODO(@benqi): check len(values)

	keyData = mtproto.MakeTLAuthKeyInfo(&mtproto.AuthKeyInfo{}).To_AuthKeyInfo()
	for k, v := range values {
		switch k {
		case "auth_key_type":
			authKeyType, _ := strconv.Atoi(v)
			keyData.AuthKeyType = int32(authKeyType)
		case "auth_key_id":
			keyData.AuthKeyId, _ = strconv.ParseInt(v, 10, 64)
		case "auth_key":
			keyData.AuthKey = []byte(v)
		case "perm_auth_key_id":
			keyData.PermAuthKeyId, _ = strconv.ParseInt(v, 10, 64)
		case "temp_auth_key_id":
			keyData.TempAuthKeyId, _ = strconv.ParseInt(v, 10, 64)
		case "media_temp_auth_key_id":
			keyData.MediaTempAuthKeyId, _ = strconv.ParseInt(v, 10, 64)
		}
	}

	return
}
