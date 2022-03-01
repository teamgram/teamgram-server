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
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/app/service/authsession/internal/dal/dataobject"
)

func (d *Dao) QueryAuthKey(ctx context.Context, authKeyId int64) (*mtproto.AuthKeyInfo, error) {
	var keyInfo *mtproto.AuthKeyInfo

	cacheKeyData, err := d.GetAuthKey(ctx, authKeyId)
	if err != nil {
		logx.WithContext(ctx).Errorf("queryAuthKey - error: %v", err)
		return nil, err
	} else if cacheKeyData != nil {
		keyInfo = &mtproto.AuthKeyInfo{
			AuthKeyId:          cacheKeyData.AuthKeyId,
			AuthKey:            cacheKeyData.AuthKey,
			AuthKeyType:        int32(cacheKeyData.AuthKeyType),
			PermAuthKeyId:      cacheKeyData.PermAuthKeyId,
			TempAuthKeyId:      cacheKeyData.TempAuthKeyId,
			MediaTempAuthKeyId: cacheKeyData.MediaTempAuthKeyId,
		}
	} else {
		do, _ := d.AuthKeysDAO.SelectByAuthKeyId(ctx, authKeyId)
		if do == nil {
			err := fmt.Errorf("not find key - keyId = %d", authKeyId)
			return nil, err
		}
		authKey, err := base64.RawStdEncoding.DecodeString(do.Body)
		if err != nil {
			logx.WithContext(ctx).Errorf("read keyData error - keyId = %d, %v", authKeyId, err)
			return nil, err
		}
		keyInfo = &mtproto.AuthKeyInfo{
			AuthKeyId:          authKeyId,
			AuthKey:            authKey,
			AuthKeyType:        mtproto.AuthKeyTypePerm,
			PermAuthKeyId:      authKeyId,
			TempAuthKeyId:      0,
			MediaTempAuthKeyId: 0,
		}
	}

	// TODO(@benqi): get salt
	return keyInfo, nil
}

func (d *Dao) InsertAuthKey(ctx context.Context, authKey *mtproto.AuthKeyInfo, salt *mtproto.TLFutureSalt, expiredIn int32) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("storage auth_key error: auth_key_id = %v", authKey)
		}
	}()

	// TODO: uncomment
	// if authKey.AuthKeyType == mtproto.AuthKeyTypePerm {
	_, _, err = d.AuthKeysDAO.Insert(ctx, &dataobject.AuthKeysDO{
		AuthKeyId: authKey.AuthKeyId,
		Body:      base64.RawStdEncoding.EncodeToString(authKey.AuthKey),
	})
	if err != nil {
		return err
	}
	// }

	if salt != nil {
		// cache salt
		err2 := d.putSaltCache(ctx, authKey.AuthKeyId, salt)
		if err2 != nil {
			// only log.
			logx.WithContext(ctx).Errorf("put cache error: ", err2)
		}
	}

	// TODO(@benqi): set expiredIn
	d.PutAuthKey(ctx, authKey.AuthKeyId, authKey, 0)
	return nil
}

func (d *Dao) GetApiLayer(ctx context.Context, authKeyId int64) int32 {
	layer, err := d.AuthsDAO.SelectLayer(ctx, authKeyId)
	if err != nil {
		logx.WithContext(ctx).Errorf("not find layer - keyId = %d", authKeyId)
		return 0
	}
	return layer
}

func (d *Dao) GetLangCode(ctx context.Context, authKeyId int64) string {
	langCode, err := d.AuthsDAO.SelectLangCode(ctx, authKeyId)
	if err != nil {
		logx.WithContext(ctx).Errorf("not find lang_code - keyId = %d", authKeyId)
		return "en"
	}
	return langCode
}

func (d *Dao) GetLangPack(ctx context.Context, authKeyId int64) string {
	langPack, err := d.AuthsDAO.SelectLangPack(ctx, authKeyId)
	if err != nil {
		logx.WithContext(ctx).Errorf("not find lang_pack - keyId = %d", authKeyId)
		return ""
	}
	return langPack
}

func (d *Dao) GetClient(ctx context.Context, authKeyId int64) string {
	do, _ := d.AuthsDAO.SelectByAuthKeyId(ctx, authKeyId)
	if do == nil {
		logx.WithContext(ctx).Errorf("not find lang_pack - keyId = %d", authKeyId)
		return ""
	}
	c := do.LangPack
	if c == "android" {
		if strings.Index(do.AppVersion, "TDLib") >= 0 {
			c = "react"
		}
	} else if c == "" {
		if do.AppVersion == "dev Z" {
			c = "webz"
		}
	}
	return c
}

func (d *Dao) GetAuthKeyUserId(ctx context.Context, authKeyId int64) int64 {
	do, _ := d.AuthUsersDAO.Select(ctx, authKeyId)
	if do == nil {
		logx.WithContext(ctx).Errorf("not find user - keyId = %d", authKeyId)
		return 0
	}
	return do.UserId
}

func (d *Dao) GetPushSessionId(ctx context.Context, userId int64, authKeyId int64, tokenType int32) int64 {
	do, _ := d.DevicesDAO.Select(ctx, authKeyId, userId, tokenType)
	if do == nil {
		logx.WithContext(ctx).Errorf("not find token - keyId = %d", authKeyId)
		return 0
	}
	sessionId, _ := strconv.ParseInt(do.Token, 10, 64)
	return int64(sessionId)
}

func (d *Dao) BindAuthKeyUser(ctx context.Context, authKeyId int64, userId int64) int64 {
	now := time.Now().Unix()
	authUsersDO := &dataobject.AuthUsersDO{
		AuthKeyId:   authKeyId,
		UserId:      userId,
		Hash:        rand.Int63(),
		DateCreated: now,
		DateActived: now,
	}
	d.AuthUsersDAO.InsertOrUpdates(ctx, authUsersDO)
	return authUsersDO.Hash
}

func (d *Dao) UnbindAuthUser(ctx context.Context, authKeyId int64, userId int64) bool {
	if authKeyId == 0 {
		d.AuthUsersDAO.DeleteUser(ctx, userId)
	} else {
		d.AuthUsersDAO.Delete(ctx, authKeyId, userId)
	}
	return true
}

func (d *Dao) SetClientSessionInfo(ctx context.Context, session *authsession.ClientSession) bool {
	do := &dataobject.AuthsDO{
		AuthKeyId:      session.GetAuthKeyId(),
		Layer:          session.GetLayer(),
		ApiId:          session.GetApiId(),
		DeviceModel:    session.GetDeviceModel(),
		SystemVersion:  session.GetSystemVersion(),
		AppVersion:     session.GetAppVersion(),
		SystemLangCode: session.GetSystemLangCode(),
		LangPack:       session.GetLangPack(),
		LangCode:       session.GetLangCode(),
		ClientIp:       session.GetIp(),
		Proxy:          session.GetProxy(),
		Params:         session.GetParams(),
		DateActive:     time.Now().Unix(),
	}
	if do.Params == "" {
		do.Params = "null"
	}
	d.AuthsDAO.InsertOrUpdate(ctx, do)
	return true
}

func (d *Dao) GetFutureSalts(ctx context.Context, authKeyId int64, num int32) (*mtproto.TLFutureSalts, error) {
	pSalts, err := d.getOrNotInsertSaltList(ctx, authKeyId, num)
	if err != nil {
		return nil, err
	}
	salts := &mtproto.TLFutureSalts{Data2: &mtproto.FutureSalts{
		ReqMsgId: 0,
		Now:      0,
		Salts:    pSalts,
	}}
	return salts, nil
}

func (d *Dao) GetPermAuthKeyId(ctx context.Context, authKeyId int64) int64 {
	if k, err := d.GetAuthKey(ctx, authKeyId); err != nil || k == nil {
		return 0
	} else {
		return k.PermAuthKeyId
	}
}
