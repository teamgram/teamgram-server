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
	"encoding/json"
	"fmt"

	"github.com/teamgram/teamgram-server/app/bff/authorization/internal/model"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	phoneCodeTimeout int64 = 90 // salt timeout
	cachePhonePrefix       = "phone_codes"
)

func genCachePhoneCodeKey(authKeyId int64, phoneNumber string) string {
	return fmt.Sprintf("%s_%d_%s", cachePhonePrefix, authKeyId, phoneNumber)
}

func (d *Dao) GetCachePhoneCode(ctx context.Context, authKeyId int64, phoneNumber string) (*model.PhoneCodeTransaction, error) {
	cacheKey := genCachePhoneCodeKey(authKeyId, phoneNumber)

	v, err := d.kv.GetCtx(ctx, cacheKey)
	if err != nil {
		logx.WithContext(ctx).Errorf("conn.GET(%s) error(%v)", cacheKey, err)
		return nil, err
	} else if v == "" {
		return nil, nil
	}

	codeData := &model.PhoneCodeTransaction{}
	err = jsonx.UnmarshalFromString(v, codeData)
	return codeData, err
}

func (d *Dao) PutCachePhoneCode(ctx context.Context, authKeyId int64, phoneNumber string, codeData *model.PhoneCodeTransaction) (err error) {
	cacheKey := genCachePhoneCodeKey(authKeyId, phoneNumber)
	b, _ := json.Marshal(codeData)

	if err = d.kv.SetexCtx(ctx, cacheKey, string(b), int(phoneCodeTimeout)); err != nil {
		logx.WithContext(ctx).Errorf("conn.SETEX(%s) error(%v)", cacheKey, err)
	}
	return
}

func (d *Dao) DeleteCachePhoneCode(ctx context.Context, authKeyId int64, phoneNumber string) (err error) {
	cacheKey := genCachePhoneCodeKey(authKeyId, phoneNumber)

	if _, err = d.kv.DelCtx(ctx, cacheKey); err != nil {
		logx.WithContext(ctx).Errorf("conn.DEL(%s) error(%v)", cacheKey, err)
	}

	return
}
