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
	"time"

	"github.com/teamgram/marmota/pkg/hack"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"

	"github.com/zeromicro/go-zero/core/logx"
)

type cacheAuthValue struct {
	UserId        int64
	Layer         int32
	pushSessionId int64
	client        string
	langpack      string
	SaltList      []*mtproto.TLFutureSalt
}

// Impl cache.Value interface
func (cv *cacheAuthValue) Size() int {
	return 1
}

func (d *Dao) getCacheValue(authKeyId int64) *cacheAuthValue {
	var (
		cacheK = strconv.FormatInt(authKeyId, 10)
	)

	if v, ok := d.cache.Get(cacheK); !ok {
		cv := new(cacheAuthValue)
		d.cache.Set(cacheK, cv)
		return cv
	} else {
		return v.(*cacheAuthValue)
	}
}

func (d *Dao) GetCacheUserID(ctx context.Context, authKeyId int64) (int64, bool) {
	cv := d.getCacheValue(authKeyId)
	if cv.UserId == 0 {
		id, err := d.AuthsessionClient.AuthsessionGetUserId(ctx, &authsession.TLAuthsessionGetUserId{
			AuthKeyId: authKeyId,
		})
		if err != nil {
			logx.WithContext(ctx).Error(err.Error())
			return 0, false
		}

		// update to cache
		cv.UserId = id.GetV()
	}
	return cv.UserId, true
}

func (d *Dao) GetCachePushSessionID(ctx context.Context, userId int64, authKeyId int64) (int64, bool) {
	cv := d.getCacheValue(authKeyId)
	if cv.pushSessionId == 0 {
		id, err := d.AuthsessionClient.AuthsessionGetPushSessionId(ctx, &authsession.TLAuthsessionGetPushSessionId{
			UserId:    userId,
			AuthKeyId: authKeyId,
			TokenType: 7,
		})
		if err != nil {
			logx.WithContext(ctx).Error(err.Error())
			return 0, false
		}
		cv.pushSessionId = id.GetV()
	}

	return cv.pushSessionId, true
}

func (d *Dao) GetCacheApiLayer(ctx context.Context, authKeyId int64) (int32, bool) {
	cv := d.getCacheValue(authKeyId)
	if cv.Layer == 0 {
		id, err := d.AuthsessionClient.AuthsessionGetLayer(ctx, &authsession.TLAuthsessionGetLayer{
			AuthKeyId: authKeyId,
		})
		if err != nil {
			logx.WithContext(ctx).Errorf(err.Error())
			return 0, false
		}
		//if r.Result != 0 {
		//	log.Errorf("queryAuthKey err: {%v}", r)
		//	return 0, false
		//}

		// update to cache
		cv.Layer = id.GetV()
	}

	return cv.Layer, true
}

func (d *Dao) GetCacheClient(ctx context.Context, authKeyId int64) string {
	cv := d.getCacheValue(authKeyId)
	if cv.client == "" {
		r, err := d.AuthsessionClient.AuthsessionGetClient(ctx, &authsession.TLAuthsessionGetClient{
			AuthKeyId: authKeyId,
		})
		if err != nil {
			logx.WithContext(ctx).Error(err.Error())
			return ""
		}

		// update to cache
		cv.client = r.GetV()
	}

	return cv.client
}

func (d *Dao) GetCacheLangpack(ctx context.Context, authKeyId int64) string {
	cv := d.getCacheValue(authKeyId)
	if cv.langpack == "" {
		r, err := d.AuthsessionClient.AuthsessionGetLangPack(ctx, &authsession.TLAuthsessionGetLangPack{
			AuthKeyId: authKeyId,
		})
		if err != nil {
			logx.WithContext(ctx).Errorf(err.Error())
			return ""
		}

		// update to cache
		cv.langpack = r.GetV()
	}

	return cv.langpack
}

func (d *Dao) PutCacheApiLayer(ctx context.Context, authKeyId int64, layer int32) {
	cv := d.getCacheValue(authKeyId)
	cv.Layer = layer
}

func (d *Dao) PutCacheClient(ctx context.Context, authKeyId int64, v string) {
	cv := d.getCacheValue(authKeyId)
	cv.client = v
}

func (d *Dao) PutCacheLangpack(ctx context.Context, authKeyId int64, v string) {
	cv := d.getCacheValue(authKeyId)
	cv.langpack = v
}

func (d *Dao) PutCacheUserId(ctx context.Context, authKeyId int64, userId int64) {
	cv := d.getCacheValue(authKeyId)
	cv.UserId = userId
}

func (d *Dao) PutCachePushSessionId(ctx context.Context, authKeyId, sessionId int64) {
	cv := d.getCacheValue(authKeyId)
	cv.pushSessionId = sessionId
}

func (d *Dao) getFutureSaltList(ctx context.Context, authKeyId int64) ([]*mtproto.TLFutureSalt, bool) {
	var (
		cv   = d.getCacheValue(authKeyId)
		date = int32(time.Now().Unix())
	)

	if len(cv.SaltList) > 0 {
		futureSalts := cv.SaltList
		for i, salt := range futureSalts {
			if salt.Data2.ValidUntil >= date {
				if i > 0 {
					return futureSalts[i-1:], true
				} else {
					return futureSalts[i:], true
				}
			}
		}
	}

	futureSalts, err := d.AuthsessionClient.AuthsessionGetFutureSalts(ctx, &authsession.TLAuthsessionGetFutureSalts{
		AuthKeyId: authKeyId,
	})
	if err != nil {
		logx.WithContext(ctx).Error(err.Error())
		return nil, false
	}

	saltList := futureSalts.GetSalts()
	for i, salt := range saltList {
		if salt.Data2.ValidUntil >= date {
			if i > 0 {
				saltList = saltList[i-1:]
				cv.SaltList = saltList
				return saltList, true
			} else {
				saltList = saltList[i:]
				cv.SaltList = saltList
				return saltList, true
			}
		}
	}

	return nil, false
}

func (d *Dao) PutUploadInitConnection(ctx context.Context, authKeyId int64, layer int32, ip string, initConnection *mtproto.TLInitConnection) error {
	session := authsession.MakeTLClientSession(&authsession.ClientSession{
		AuthKeyId:      authKeyId,
		Ip:             ip,
		Layer:          layer,
		ApiId:          initConnection.GetApiId(),
		DeviceModel:    initConnection.GetDeviceModel(),
		SystemVersion:  initConnection.GetSystemVersion(),
		AppVersion:     initConnection.GetAppVersion(),
		SystemLangCode: initConnection.GetSystemLangCode(),
		LangPack:       initConnection.GetLangPack(),
		LangCode:       initConnection.GetLangCode(),
		Proxy:          "",
		Params:         "",
	}).To_ClientSession()

	if initConnection.GetProxy() != nil {
		session.Proxy = hack.String(mtproto.TLObjectToJson(initConnection.Proxy))
	}
	if initConnection.GetParams() != nil {
		session.Params = hack.String(mtproto.TLObjectToJson(initConnection.Params))
	}

	_, err := d.AuthsessionClient.AuthsessionSetClientSessionInfo(ctx, &authsession.TLAuthsessionSetClientSessionInfo{
		Data: session,
	})

	if err != nil {
		logx.WithContext(ctx).Error(err.Error())
	}

	return err
}

func (d *Dao) GetOrFetchNewSalt(ctx context.Context, authKeyId int64) (salt, lastInvalidSalt *mtproto.TLFutureSalt, err error) {
	cacheSalts, _ := d.getFutureSaltList(ctx, authKeyId)
	//TODO(@benqi): check len(cacheSalts) > 0
	if len(cacheSalts) < 2 {
		return nil, nil, fmt.Errorf("get salt error")
	} else {
		if cacheSalts[0].GetValidUntil() >= int32(time.Now().Unix()) {
			return cacheSalts[0], nil, nil
		} else {
			return cacheSalts[1], cacheSalts[0], nil
		}
	}
}

func (d *Dao) GetFutureSalts(ctx context.Context, authKeyId int64, num int32) ([]*mtproto.TLFutureSalt, error) {
	cacheSalts, _ := d.getFutureSaltList(ctx, authKeyId)
	//TODO(@benqi): check len(cacheSalts) > 0

	return cacheSalts, nil
}

func (d *Dao) GetKeyStateData(ctx context.Context, authKeyId int64) (*authsession.AuthKeyStateData, error) {
	return d.AuthsessionClient.AuthsessionGetAuthStateData(ctx, &authsession.TLAuthsessionGetAuthStateData{
		AuthKeyId: authKeyId,
	})
}
