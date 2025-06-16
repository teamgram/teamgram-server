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
	"math/rand"
	"strings"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

const (
	authDataPrefix = "auth_data.2"
)

func genAuthDataCacheKey(id int64) string {
	return fmt.Sprintf("%s#%d", authDataPrefix, id)
}

//
//func genAuthUserCacheKey(id int64) string {
//	return fmt.Sprintf("%s_%d", authUsersTablePrefix, id)
//}

type BindUser struct {
	UserId               int64 `json:"user_id"`
	Hash                 int64 `json:"hash"`
	DateCreated          int64 `json:"date_created"`
	DateActivated        int64 `json:"date_activated"`
	AndroidPushSessionId int64 `json:"android_push_sessionId"`
}

type CacheAuthData struct {
	Client   *authsession.TLClientSession `json:"client"`
	BindUser *BindUser                    `json:"bind_user,omitempty"`
}

func (c *CacheAuthData) ToAuthState() (state int) {
	state = tg.AuthStateUnknown

	if c == nil {
		state = tg.AuthStateNew
		return
	}

	if c.Client == nil {
		state = tg.AuthStateWaitInit
	} else if c.BindUser == nil {
		state = tg.AuthStateUnauthorized
	} else {
		// TODO: need session password
		state = tg.AuthStateNormal
	}

	return
}

func (c *CacheAuthData) GetClient() *authsession.TLClientSession {
	if c == nil {
		return nil
	}
	return c.Client
}

func (c *CacheAuthData) AuthKeyId() int64 {
	return c.Client.AuthKeyId
}

func (c *CacheAuthData) Layer() int32 {
	return c.Client.Layer
}

func (c *CacheAuthData) ApiId() int32 {
	return c.Client.ApiId
}

func (c *CacheAuthData) DeviceModel() string {
	return c.Client.DeviceModel
}

func (c *CacheAuthData) SystemVersion() string {
	return c.Client.SystemVersion
}

func (c *CacheAuthData) AppVersion() string {
	return c.GetClient().AppVersion
}

func (c *CacheAuthData) SystemLangCode() string {
	return c.Client.SystemLangCode
}

func (c *CacheAuthData) LangPack() string {
	return c.Client.LangPack
}

func (c *CacheAuthData) LangCode() string {
	return c.Client.LangCode
}

func (c *CacheAuthData) ClientIp() string {
	return c.Client.Ip
}

func (c *CacheAuthData) Proxy() string {
	return c.Client.Proxy
}

func (c *CacheAuthData) Params() string {
	return c.Client.Params
}

func (c *CacheAuthData) UserId() int64 {
	if c == nil || c.BindUser == nil {
		return 0
	}
	return c.BindUser.UserId
}

func (c *CacheAuthData) DateCreated() int64 {
	if c == nil || c.BindUser == nil {
		return 0
	}
	return c.BindUser.DateCreated
}

func (c *CacheAuthData) DateActivated() int64 {
	if c == nil || c.BindUser == nil {
		return 0
	}
	return c.BindUser.DateActivated
}

func (c *CacheAuthData) Hash() int64 {
	if c == nil || c.BindUser == nil {
		return 0
	}
	return c.BindUser.Hash
}

func (c *CacheAuthData) AndroidPushSessionId() int64 {
	if c == nil || c.BindUser == nil {
		return 0
	}
	return c.BindUser.AndroidPushSessionId
}

func (d *Dao) GetApiLayer(ctx context.Context, authKeyId int64) int32 {
	cData, err := d.GetCacheAuthData(ctx, authKeyId)
	if err != nil {
		logx.WithContext(ctx).Errorf("not find layer - keyId = %d", authKeyId)
		return 0
	}
	return cData.Layer()
}

func (d *Dao) GetLangCode(ctx context.Context, authKeyId int64) string {
	cData, err := d.GetCacheAuthData(ctx, authKeyId)
	if err != nil {
		logx.WithContext(ctx).Errorf("not find lang_code - keyId = %d", authKeyId)
		return "en"
	}
	return cData.LangCode()
}

func (d *Dao) GetLangPack(ctx context.Context, authKeyId int64) string {
	cData, err := d.GetCacheAuthData(ctx, authKeyId)
	if err != nil {
		logx.WithContext(ctx).Errorf("not find lang_pack - keyId = %d", authKeyId)
		return ""
	}

	c := cData.LangPack()
	if c == "" {
		if strings.HasSuffix(cData.AppVersion(), " A") {
			c = "weba"
		} else if strings.HasSuffix(cData.AppVersion(), " Z") {
			c = "weba"
		}
	}
	return c
}

func (d *Dao) GetClient(ctx context.Context, authKeyId int64) string {
	cData, err := d.GetCacheAuthData(ctx, authKeyId)
	if err != nil {
		logx.WithContext(ctx).Errorf("not find client - keyId = %d", authKeyId)
		return ""
	}
	c := cData.LangPack()
	if c == "android" {
		if strings.Index(cData.AppVersion(), "TDLib") >= 0 {
			c = "react"
		}
	} else if c == "" {
		if strings.HasSuffix(cData.AppVersion(), " A") {
			c = "weba"
		} else if strings.HasSuffix(cData.AppVersion(), " Z") {
			c = "weba"
		}
	}
	return c
}

func (d *Dao) GetAuthKeyUserId(ctx context.Context, authKeyId int64) int64 {
	cData, _ := d.GetCacheAuthData(ctx, authKeyId)
	// do, _ := d.AuthUsersDAO.Select(ctx, authKeyId)
	if cData == nil {
		logx.WithContext(ctx).Errorf("not find user - keyId = %d", authKeyId)
		return 0
	}

	return cData.UserId()
}

func (d *Dao) BindAuthKeyUser(ctx context.Context, authKeyId int64, userId int64) int64 {
	now := time.Now().Unix()
	authUsersDO := &dataobject.AuthUsersDO{
		AuthKeyId:   authKeyId,
		UserId:      userId,
		Hash:        rand.Int63(),
		DateCreated: now,
		DateActive:  now,
	}

	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			return d.AuthUsersDAO.InsertOrUpdates(ctx, authUsersDO)
		},
		genAuthDataCacheKey(authKeyId))
	if err != nil {
		return 0
	}

	return authUsersDO.Hash
}

func (d *Dao) UnbindAuthUser(ctx context.Context, authKeyId int64, userId int64) bool {
	var (
		err error
	)

	if authKeyId == 0 {
		var (
			idList []string
		)
		d.AuthUsersDAO.SelectAuthKeyIdsWithCB(
			ctx,
			userId,
			func(sz, i int, v *dataobject.AuthUsersDO) {
				idList = append(idList, genAuthDataCacheKey(v.AuthKeyId))
			})
		if len(idList) > 0 {
			_, _, err = d.CachedConn.Exec(
				ctx,
				func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
					_, err2 := d.AuthUsersDAO.DeleteUser(ctx, userId)
					return 0, 0, err2
				},
				idList...)
		}
	} else {
		_, _, err = d.CachedConn.Exec(
			ctx,
			func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
				_, err2 := d.AuthUsersDAO.Delete(ctx, authKeyId, userId)
				return 0, 0, err2
			},
			genAuthDataCacheKey(authKeyId))
	}

	return err == nil
}

func (d *Dao) SetClientSessionInfo(ctx context.Context, session *authsession.TLClientSession) error {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			do := &dataobject.AuthsDO{
				AuthKeyId:      session.AuthKeyId,
				Layer:          session.Layer,
				ApiId:          session.ApiId,
				DeviceModel:    session.DeviceModel,
				SystemVersion:  session.SystemVersion,
				AppVersion:     session.AppVersion,
				SystemLangCode: session.SystemLangCode,
				LangPack:       session.LangPack,
				LangCode:       session.LangCode,
				ClientIp:       session.Ip,
				Proxy:          session.Proxy,
				Params:         session.Params,
				DateActive:     time.Now().Unix(),
			}
			if do.Params == "" {
				do.Params = "null"
			}
			return d.AuthsDAO.InsertOrUpdate(ctx, do)
		},
		genAuthDataCacheKey(session.AuthKeyId))

	return err
}

func (d *Dao) SetLayer(ctx context.Context, in *authsession.TLAuthsessionSetLayer) error {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			return d.AuthsDAO.InsertOrUpdateLayer(
				ctx,
				&dataobject.AuthsDO{
					AuthKeyId:  in.AuthKeyId,
					Layer:      in.Layer,
					ClientIp:   in.Ip,
					DateActive: time.Now().Unix(),
				})
		},
		genAuthDataCacheKey(in.AuthKeyId))

	return err
}

func (d *Dao) SetInitConnection(ctx context.Context, i *authsession.TLAuthsessionSetInitConnection) error {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			do := &dataobject.AuthsDO{
				AuthKeyId:      i.AuthKeyId,
				ApiId:          i.ApiId,
				DeviceModel:    i.DeviceModel,
				SystemVersion:  i.SystemVersion,
				AppVersion:     i.AppVersion,
				SystemLangCode: i.SystemLangCode,
				LangPack:       i.LangPack,
				LangCode:       i.LangCode,
				ClientIp:       i.Ip,
				Proxy:          i.Proxy,
				Params:         i.Params,
				DateActive:     time.Now().Unix(),
			}
			if do.Params == "" {
				do.Params = "null"
			}
			return d.AuthsDAO.InsertOrUpdate(ctx, do)
		},
		genAuthDataCacheKey(i.AuthKeyId))

	return err
}

func (d *Dao) SetAndroidPushSessionId(ctx context.Context, userId, keyId, sessionId int64) error {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			_, err2 := d.AuthUsersDAO.UpdateAndroidPushSessionId(
				ctx,
				sessionId,
				keyId,
				userId)

			return 0, 0, err2
		},
		genAuthDataCacheKey(keyId))

	return err

}

func (d *Dao) GetCacheAuthData(ctx context.Context, authKeyId int64) (*CacheAuthData, error) {
	var (
		cData *CacheAuthData
	)

	err := d.CachedConn.QueryRow(
		ctx,
		&cData,
		genAuthDataCacheKey(authKeyId),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			cacheAuthData := &CacheAuthData{
				Client:   nil,
				BindUser: nil,
			}
			err := mr.Finish(
				func() error {
					do, err := d.AuthsDAO.SelectByAuthKeyId(ctx, authKeyId)
					if err != nil {
						return err
					}
					if do == nil {
						return sqlc.ErrNotFound
					}
					cacheAuthData.Client = &authsession.TLClientSession{
						AuthKeyId:      authKeyId,
						Ip:             do.ClientIp,
						Layer:          do.Layer,
						ApiId:          do.ApiId,
						DeviceModel:    do.DeviceModel,
						SystemVersion:  do.SystemVersion,
						AppVersion:     do.AppVersion,
						SystemLangCode: do.SystemLangCode,
						LangPack:       do.LangPack,
						LangCode:       do.LangCode,
						Proxy:          do.Proxy,
						Params:         do.Params,
					}

					return nil
				},
				func() error {
					do, _ := d.AuthUsersDAO.Select(ctx, authKeyId)
					if do != nil {
						cacheAuthData.BindUser = &BindUser{
							UserId:               do.UserId,
							Hash:                 do.Hash,
							DateCreated:          do.DateCreated,
							DateActivated:        do.DateActive,
							AndroidPushSessionId: do.AndroidPushSessionId,
						}
					}
					return nil
				})

			if err != nil {
				return err
			}
			*v.(**CacheAuthData) = cacheAuthData

			return nil
		})
	if err != nil {
		return nil, err
	}

	return cData, nil
}
