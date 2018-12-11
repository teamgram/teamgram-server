// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package server

import (
	"context"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/util"
	"github.com/nebula-chat/chatengine/pkg/cache"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery"
	"github.com/nebula-chat/chatengine/mtproto"
	"time"
	"fmt"
)

type cacheAuthValue struct {
	AuthKey       []byte
	UserId        int32
	Layer         int32
	pushSessionId int64
	SaltList      []*mtproto.TLFutureSalt
}

// Impl cache.Value interface
func (cv *cacheAuthValue) Size() int {
	return 1
}

type cacheAuthManager struct {
	cache  *cache.LRUCache
	client mtproto.RPCSessionClient
}

var _cacheAuthManager *cacheAuthManager

func InitCacheAuthManager(cap int64, discovery *service_discovery.ServiceDiscoveryClientConfig) {
	conn, err := grpc_util.NewRPCClientByServiceDiscovery(discovery)
	if err != nil {
		glog.Error(err)
		panic(err)
	}

	_cacheAuthManager = &cacheAuthManager{
		cache:  cache.NewLRUCache(cap),
		client: mtproto.NewRPCSessionClient(conn),
	}
}

func (c *cacheAuthManager) GetAuthKey(authKeyId int64) ([]byte, bool) {
	var (
		cacheK = util.Int64ToString(authKeyId)
	)

	if v, ok := c.cache.Get(cacheK); !ok {
		keyInfo, err := c.client.SessionQueryAuthKey(context.Background(), &mtproto.TLSessionQueryAuthKey{AuthKeyId: authKeyId})
		if err != nil {
			glog.Error(err)
			return nil, false
		}
		//if r.Result != 0 {
		//	glog.Errorf("queryAuthKey err: {%v}", r)
		//	return nil, false
		//}
		c.cache.Set(cacheK, &cacheAuthValue{AuthKey: keyInfo.GetData2().GetAuthKey()})

		// TODO(@benqi): salt.
		return keyInfo.GetData2().GetAuthKey(), true
	} else {
		return v.(*cacheAuthValue).AuthKey, true
	}
}

func (c *cacheAuthManager) GetUserID(authKeyId int64) (int32, bool) {
	var (
		cacheK = util.Int64ToString(authKeyId)
	)

	if v, ok := c.cache.Peek(cacheK); !ok {
		glog.Error("not found authKeyId, bug???")
		return 0, false
	} else {
		cv, _ := v.(*cacheAuthValue)
		if cv.UserId == 0 {
			id, err := c.client.SessionGetUserId(context.Background(), &mtproto.TLSessionGetUserId{AuthKeyId: authKeyId})
			if err != nil {
				glog.Error(err)
				return 0, false
			}
			//if r.Result != 0 {
			//	glog.Errorf("queryAuthKey err: {%v}", r)
			//	return 0, false
			//}

			// update to cache
			cv.UserId = id.GetData2().GetV()
		}

		return cv.UserId, true
	}
}

func (c *cacheAuthManager) GetPushSessionID(userId int32, authKeyId int64) (int64, bool) {
	var (
		cacheK = util.Int64ToString(authKeyId)
	)

	if v, ok := c.cache.Peek(cacheK); !ok {
		glog.Error("not found authKeyId, bug???")
		return 0, false
	} else {
		cv, _ := v.(*cacheAuthValue)
		if cv.pushSessionId == 0 {
			id, err := c.client.SessionGetPushSessionId(context.Background(), &mtproto.TLSessionGetPushSessionId{
				UserId:    userId,
				AuthKeyId: authKeyId,
				TokenType: 7,
			})
			if err != nil {
				glog.Error(err)
				return 0, false
			}
			cv.pushSessionId = id.GetData2().GetV()
		}

		return cv.pushSessionId, true
	}
}

func (c *cacheAuthManager) GetApiLayer(authKeyId int64) (int32, bool) {
	var (
		cacheK = util.Int64ToString(authKeyId)
	)

	if v, ok := c.cache.Peek(cacheK); !ok {
		glog.Error("not found authKeyId, bug???")
		return 0, false
	} else {
		cv, _ := v.(*cacheAuthValue)
		if cv.Layer == 0 {
			id, err := c.client.SessionGetLayer(context.Background(), &mtproto.TLSessionGetLayer{AuthKeyId: authKeyId})
			if err != nil {
				glog.Error(err)
				return 0, false
			}
			//if r.Result != 0 {
			//	glog.Errorf("queryAuthKey err: {%v}", r)
			//	return 0, false
			//}

			// update to cache
			cv.Layer = id.GetData2().GetV()
		}

		return cv.Layer, true
	}
}

func (c *cacheAuthManager) PutUserID(authKeyId int64, userId int32) {
	var (
		cacheK = util.Int64ToString(authKeyId)
	)

	if v, ok := c.cache.Peek(cacheK); ok {
		v.(*cacheAuthValue).UserId = userId
	} else {
		glog.Error("not found authKeyId, bug???")
	}
}

func (c *cacheAuthManager) PutPushSessionID(authKeyId, sessionId int64) {
	var (
		cacheK = util.Int64ToString(authKeyId)
	)

	if v, ok := c.cache.Peek(cacheK); ok {
		v.(*cacheAuthValue).pushSessionId = sessionId
	} else {
		glog.Error("not found authKeyId, bug???")
	}
}

func (c *cacheAuthManager) GetFutureSaltList(authKeyId int64) ([]*mtproto.TLFutureSalt, bool) {
	var (
		cacheK = util.Int64ToString(authKeyId)
		date = int32(time.Now().Unix())
	)

	v, ok := c.cache.Get(cacheK)
	if ok {
		futureSalts := v.(*cacheAuthValue).SaltList
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

	futureSalts, err := c.client.SessionGetFutureSalts(context.Background(), &mtproto.TLSessionGetFutureSalts{AuthKeyId: authKeyId})
	if err != nil {
		glog.Error(err)
		return nil, false
	}

	saltList := futureSalts.GetData2().GetSalts()
	for i, salt := range saltList {
		if salt.Data2.ValidUntil >= date {
			if i > 0 {
				saltList = saltList[i-1:]
				v.(*cacheAuthValue).SaltList = saltList
				return saltList, true
			} else {
				saltList = saltList[i:]
				v.(*cacheAuthValue).SaltList = saltList
				return saltList, true
			}
		}
	}

	return nil, false
}

func getCachePushSessionID(userId int32, authKeyId int64) int64 {
	if _cacheAuthManager == nil {
		panic("not init cacheAuthManager.")
	}

	sessionId, _ := _cacheAuthManager.GetPushSessionID(userId, authKeyId)
	return sessionId
}

func putCachePushSessionId(authKeyId, sessionId int64) {
	if _cacheAuthManager == nil {
		panic("not init cacheAuthManager.")
	}

	_cacheAuthManager.PutPushSessionID(authKeyId, sessionId)
}

func getCacheUserID(authKeyId int64) int32 {
	if _cacheAuthManager == nil {
		panic("not init cacheAuthManager.")
	}

	userId, _ := _cacheAuthManager.GetUserID(authKeyId)
	return userId
}

func putCacheUserId(authKeyId int64, userId int32) {
	if _cacheAuthManager == nil {
		panic("not init cacheAuthManager.")
	}

	_cacheAuthManager.PutUserID(authKeyId, userId)
}

func getCacheAuthKey(authKeyId int64) []byte {
	if _cacheAuthManager == nil {
		panic("not init cacheAuthManager.")
	}

	key, _ := _cacheAuthManager.GetAuthKey(authKeyId)
	return key
}

func getCacheApiLayer(authKeyId int64) int32 {
	if _cacheAuthManager == nil {
		panic("not init cacheAuthManager.")
	}

	layer, _ := _cacheAuthManager.GetApiLayer(authKeyId)
	return layer
}

func uploadInitConnection(authKeyId int64, layer int32, ip string, initConnection *TLInitConnectionExt) error {
	session := &mtproto.TLClientSessionInfo { Data2: &mtproto.ClientSession_Data{
		AuthKeyId:      authKeyId,
		Ip:             ip,
		Layer:          layer,
		ApiId:          initConnection.ApiId,
		DeviceModel:    initConnection.DeviceMode,
		SystemVersion:  initConnection.SystemVersion,
		AppVersion:     initConnection.AppVersion,
		SystemLangCode: initConnection.SystemLangCode,
		LangPack:       initConnection.LangPack,
		LangCode:       initConnection.LangCode,
	}}

	request := &mtproto.TLSessionSetClientSessionInfo{
		Session: session.To_ClientSession(),
	}

	_, err := _cacheAuthManager.client.SessionSetClientSessionInfo(context.Background(), request)

	if err != nil {
		glog.Error(err)
	}

	return err
}

func getOrFetchNewSalt(authKeyId int64) (salt, lastInvalidSalt *mtproto.TLFutureSalt, err error) {
	if _cacheAuthManager == nil {
		panic("not init cacheAuthManager.")
	}

	cacheSalts, _ := _cacheAuthManager.GetFutureSaltList(authKeyId)
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

func getFutureSalts(authKeyId int64, num int32) ([]*mtproto.TLFutureSalt, error) {
	if _cacheAuthManager == nil {
		panic("not init cacheAuthManager.")
	}

	cacheSalts, _ := _cacheAuthManager.GetFutureSaltList(authKeyId)
	//TODO(@benqi): check len(cacheSalts) > 0

	return cacheSalts, nil
}

