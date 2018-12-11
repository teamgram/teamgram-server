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

package status_client

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/gomodule/redigo/redis"
	// "github.com/nebula-chat/chatengine/pkg/util"
	"github.com/nebula-chat/chatengine/pkg/redis_client"
	"github.com/nebula-chat/chatengine/service/status/proto"
	"strings"
	"time"
	"github.com/nebula-chat/chatengine/pkg/util"
)

const (
	onlineKeyPrefix = "online" //
	userKeyIdsPrefix = "user_keys" //
)

//func makeSessionEntry(userId int32, k, v string) (sess *status.SessionEntry, err error) {
//	sess = &status.SessionEntry{
//		UserId: userId,
//	}
//
//	if sess.AuthKeyId, err = util.StringToInt64(k); err != nil {
//		glog.Errorf("makeSessionEntry(%d, %s, %s) error: %v", userId, k, v, err)
//		return
//	}
//
//	vals := strings.Split(v, "@")
//	if len(vals) != 3 {
//		err = fmt.Errorf("makeSessionEntry(%d, %s, %s) - Invalid value: %s", userId, k, v, v)
//		return
//	}
//
//	if sess.ServerId, err = util.StringToInt32(vals[0]); err != nil {
//		return
//	}
//	if sess.Expired, err = util.StringToInt64(vals[1]); err != nil {
//		return
//	}
//	if sess.Layer, err = util.StringToInt32(vals[2]); err != nil {
//		return
//	}
//
//	return
//}

//////////////////////////////////////////////////////////////////////////
type redisStatusClient struct {
	redis *redis_client.RedisPool
}

func redisStatusClientInstance() StatusClient {
	return &redisStatusClient{}
}

func NewRedisStatusClient(redis *redis_client.RedisPool) *redisStatusClient {
	cli := &redisStatusClient{redis}
	return cli
}

func (c *redisStatusClient) Initialize(config string) error {
	c.redis = redis_client.GetRedisClient(config)
	if c.redis == nil {
		return fmt.Errorf("init redisStatusClient error: %s", config)
	}
	return nil
}

func (c *redisStatusClient) SetSessionOnlineTTL(userId int32, authKeyId int64, serverId, layer int32, ttl int32) (err error) {
	conn := c.redis.Get()
	defer conn.Close()

	// TODO(@benqi): expired
	if _, err = conn.Do("SADD", fmt.Sprintf("%s_%d", userKeyIdsPrefix, userId), fmt.Sprintf("%d", authKeyId)); err != nil {
		glog.Errorf("setSessionOnlineTTL - SADD {%d, %d}, error: %s", userId, authKeyId, err)
		return
	}

	id := fmt.Sprintf("%s_%d", onlineKeyPrefix, authKeyId)
	v := fmt.Sprintf("%d@%d@%d@%d", userId, serverId, time.Now().Unix() + int64(ttl), layer)

	glog.Info("setSessionOnlineTTL: ", id, " -- ", v)
	if _, err = conn.Do("SETEX", id, int64(ttl), v); err != nil {
		glog.Errorf("setSessionOnlineTTL - SETEX {%s, %s, %s}, error: %v", id, v, ttl, err)
		return
	}

	return
}

func (c *redisStatusClient) SetSessionOnline(userId int32, authKeyId int64, serverId, layer int32) (err error) {
	//conn := c.redis.Get()
	//defer conn.Close()
	//
	//// TODO(@benqi): expired
	//if _, err = conn.Do("SADD", fmt.Sprintf("%s_%d", userKeyIdsPrefix, userId), fmt.Sprintf("update_%d", authKeyId)); err != nil {
	//	glog.Errorf("setOnline - SADD {%d, %d}, error: %s", userId, authKeyId, err)
	//	return
	//}
	//
	//id := fmt.Sprintf("%s_update_%d", onlineKeyPrefix, authKeyId)
	//v := fmt.Sprintf("%d@%d@%d@%d", userId, serverId, time.Now().Unix() + int64(ONLINE_TIMEOUT), layer)
	//
	//glog.Info("setOnline: ", id, " -- ", v)
	//if _, err = conn.Do("SETEX", id, int64(ONLINE_TIMEOUT), v); err != nil {
	//	glog.Errorf("setOnline - SETEX {%s, %s, %s}, error: %v", id, v, ONLINE_TIMEOUT, err)
	//	return
	//}

	return
}

func (c *redisStatusClient) SetSessionOffline(userId int32, serverId int32, authKeyId int64) (err error) {
	//conn := c.redis.Get()
	//defer conn.Close()
	//
	//if _, err = conn.Do("SREM", fmt.Sprintf("%s_%d", userKeyIdsPrefix, userId), fmt.Sprintf("update_%d", authKeyId)); err != nil {
	//	glog.Errorf("setOnline - SREM {%d, %d}, error: %s", userId, authKeyId, err)
	//	return
	//}
	//
	//id := fmt.Sprintf("%s_update_%d", onlineKeyPrefix, authKeyId)
	//if _, err = conn.Do("DEL", id); err != nil {
	//	glog.Errorf("setOffline - DEL {%s, %s}, error: %s", id, err)
	//	return
	//}

	return
}

func (c *redisStatusClient) SetSessionOfflineTTL(userId int32, serverId int32, authKeyId int64) (err error) {
	conn := c.redis.Get()
	defer conn.Close()

	if _, err = conn.Do("SREM", fmt.Sprintf("%s_%d", userKeyIdsPrefix, userId), fmt.Sprintf("%d", authKeyId)); err != nil {
		glog.Errorf("setSessionOfflineTTL - SREM {%d, %d}, error: %s", userId, authKeyId, err)
		return
	}

	id := fmt.Sprintf("%s_%d", onlineKeyPrefix, authKeyId)
	if _, err = conn.Do("DEL", id); err != nil {
		glog.Errorf("setSessionOfflineTTL - DEL {%s, %s}, error: %s", id, err)
		return
	}

	return
}

func (c *redisStatusClient) GetUserOnlineSessions(userId int32) (*status.SessionEntryList, error) {
	conn := c.redis.Get()
	defer conn.Close()

	sesses, err := c.GetUsersOnlineSessionsList([]int32{userId})
	if err != nil{
		return nil, err
	}

	if _, ok := sesses.UsersSessions[userId]; !ok {
		return &status.SessionEntryList{
			Sessions:     []*status.SessionEntry{},
			// PushSessions: []*status.SessionEntry{},
		}, nil
	}
	return sesses.UsersSessions[userId], nil
}

func (c *redisStatusClient) GetUsersOnlineSessionsList(userIdList []int32) (usersSessList *status.UsersSessionEntryList, err error) {
	conn := c.redis.Get()
	defer conn.Close()

	usersSessList = &status.UsersSessionEntryList{
		UsersSessions: map[int32]*status.SessionEntryList{},
	}

	var idList = make([]string, 0, len(userIdList))
	for _, userId := range userIdList {
		idList = append(idList, fmt.Sprintf("%s_%d", userKeyIdsPrefix, userId))
	}

	keys, err := redis.Strings(conn.Do("SUNION", strings.Join(idList, " ")))
	if err != nil {
		glog.Errorf("getUsersOnlineSessionsList - SUNION {%s}, error: %s", strings.Join(idList, " "), err)
		return
	}

	if len(keys) == 0 {
		return
	}

	var sessionKeys = make([]interface{}, 0, 2*len(keys))
	for i := 0; i < len(keys); i++ {
		sessionKeys = append(sessionKeys, fmt.Sprintf("%s_%s", onlineKeyPrefix, keys[i]))
		// sessionKeys = append(sessionKeys, fmt.Sprintf("%s_update_%s", onlineKeyPrefix, keys[i]))
		// keys[i] = fmt.Sprintf("%s_push_%d", onlineKeyPrefix, authKeyId)
	}

	// redisSessionKeys := strings.Join(sessionKeys, " ")
	glog.Info("sessionKeys: ", sessionKeys)
	// redis.Int64Map()
	onlineSessions, err := redis.Strings(conn.Do("MGET", sessionKeys...))
	if err != nil {
		glog.Errorf("getUsersOnlineSessionsList - MGET {%v}, error: %s", sessionKeys, err)
		return
	}

	// values, err := redis.Values(c.Do("MGET", args...))
	glog.Info("getUsersOnlineSessionsList - onlineSessions: ", onlineSessions, ", length = ", len(onlineSessions))

	if len(onlineSessions) == 0 {
		return
	}

	usersSessions := make(map[int32]*status.SessionEntryList)
	for i := 0; i < len(onlineSessions); i++ {
		if onlineSessions[i] == "" {
			continue
		}

		// for k, v := range onlineSessions {
		k := sessionKeys[i].(string)
		v := onlineSessions[i]

		ks := strings.Split(k, "_")
		if len(ks) != 2 {
			continue
		}
		// isPush := ks[1] == "push"
		authKeyId, _ := util.StringToInt64(ks[1])

		vs := strings.Split(v, "@")
		if len(vs) != 4 {
			continue
		}
		userId, _ := util.StringToInt32(vs[0])
		serverId, _ := util.StringToInt32(vs[1])
		expired, _ := util.StringToInt64(vs[2])
		layer, _ := util.StringToInt32(vs[3])

		entry := &status.SessionEntry{
			UserId:    userId,
			ServerId:  serverId,
			AuthKeyId: authKeyId,
			Expired:   expired,
			Layer:     layer,
		}

		if _, ok := usersSessions[userId]; !ok {
			usersSessions[userId] = &status.SessionEntryList{
				Sessions:     []*status.SessionEntry{},
				// PushSessions: []*status.SessionEntry{},
			}
		}
		entryList := usersSessions[userId]

		entryList.Sessions = append(entryList.Sessions, entry)

		//if isPush {
		//	entryList.PushSessions = append(entryList.PushSessions, entry)
		//} else {
		//	entryList.Sessions = append(entryList.Sessions, entry)
		//}
	}
 	// */
	usersSessList.UsersSessions = usersSessions
	glog.Info("result - ", usersSessList)
	//= &status.UsersSessionEntryList{
	//	UsersSessions: usersSessions,
	//}

	return
}

func init() {
	Register("redis", redisStatusClientInstance)
}
