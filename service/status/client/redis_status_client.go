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
	"github.com/nebula-chat/chatengine/pkg/util"
	"github.com/nebula-chat/chatengine/pkg/redis_client"
	"github.com/nebula-chat/chatengine/service/status/proto"
	"strings"
	"time"
)

const (
	onlineKeyPrefix = "online" //
)

func makeSessionEntry(userId int32, k, v string) (sess *status.SessionEntry, err error) {
	sess = &status.SessionEntry{
		UserId: userId,
	}

	if sess.AuthKeyId, err = util.StringToInt64(k); err != nil {
		glog.Errorf("makeSessionEntry(%d, %s, %s) error: %v", userId, k, v, err)
		return
	}

	vals := strings.Split(v, "@")
	if len(vals) != 3 {
		err = fmt.Errorf("makeSessionEntry(%d, %s, %s) - Invalid value: %s", userId, k, v, v)
		return
	}

	if sess.ServerId, err = util.StringToInt32(vals[0]); err != nil {
		return
	}
	if sess.Expired, err = util.StringToInt64(vals[1]); err != nil {
		return
	}
	if sess.Layer, err = util.StringToInt32(vals[2]); err != nil {
		return
	}

	return
}

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

func (c *redisStatusClient) SetSessionOnline(userId int32, authKeyId int64, serverId, layer int32) (err error) {
	conn := c.redis.Get()
	defer conn.Close()

	id := fmt.Sprintf("%s_%d", onlineKeyPrefix, userId)
	k := util.Int64ToString(authKeyId)
	v := fmt.Sprintf("%d@%d@%d", serverId, time.Now().Unix(), layer)
	if _, err = conn.Do("HSET", id, k, v); err != nil {
		glog.Errorf("setOnline - HSET {%s, %s, %s}, error: %s", id, k, v, err)
		return
	}

	if _, err = conn.Do("EXPIRE", id, ONLINE_TIMEOUT); err != nil {
		glog.Errorf("setOnline - EXPIRE {%s, %s, %s}, error: %s", id, k, v, err)
		return
	}
	return
}

func (c *redisStatusClient) SetSessionOffline(userId int32, serverId int32, authKeyId int64) (err error) {
	conn := c.redis.Get()
	defer conn.Close()

	id := fmt.Sprintf("%s_%d", onlineKeyPrefix, userId)
	k := util.Int64ToString(authKeyId)

	if _, err = conn.Do("HDEL", id, k); err != nil {
		glog.Errorf("setOffline - HDEL {%s, %s}, error: %s", id, k, err)
		return
	}
	return
}

func (c *redisStatusClient) getOnlineSession(conn redis.Conn, userId int32) (sessList []*status.SessionEntry, err error) {
	fmt.Printf("%s_%d\n", onlineKeyPrefix, userId)
	m, err := redis.StringMap(conn.Do("HGETALL", fmt.Sprintf("%s_%d", onlineKeyPrefix, userId)))
	if err != nil {
		glog.Errorf("getOnlineSession - HGETALL {online_%d}, error: %s", userId, err)
		return
	}

	fmt.Println(m)
	for k, v := range m {
		sess, err2 := makeSessionEntry(userId, k, v)
		if err2 != nil {
			glog.Errorf("getOnlineSession - makeSessionEntry {online_%d} error: {k: %s, v: %s}, error: %s", userId, k, v)
			continue
		}

		if time.Now().Unix() < sess.Expired+CHECK_ONLINE_TIMEOUT {
			sessList = append(sessList, sess)
			fmt.Println("getOnlineSession - ", sess)
		}
	}
	return
}

//func getOnlineByUserId(userId int32) ([]*SessionStatus, error) {
//	redis_client := redis_client.GetRedisClient(dao.CACHE)
//
//	conn := redis_client.Get()
//	defer conn.Close()
//
//	return getOnline(conn, userId)
//}
//
func (c *redisStatusClient) GetUserOnlineSessions(userId int32) (*status.SessionEntryList, error) {
	conn := c.redis.Get()
	defer conn.Close()

	sessList, err := c.getOnlineSession(conn, userId)
	if err != nil {
		return nil, err
	}

	return &status.SessionEntryList{Sessions: sessList}, nil
}

// TODO(@benqi): 优化
// 取多个用户的状态信息，可以使用lua脚本，避免多次请求
// eval "local rst={}; for i,v in pairs(KEYS) do rst[i]=redis.call('hgetall', v) end; return rst" 2 user:1 user:2
func (c *redisStatusClient) GetUsersOnlineSessionsList(userIdList []int32) (usersSessList *status.UsersSessionEntryList, err error) {
	conn := c.redis.Get()
	defer conn.Close()

	usersSessions := make(map[int32]*status.SessionEntryList)
	for _, userId := range userIdList {
		ss, err := c.getOnlineSession(conn, userId)
		if err != nil {
			glog.Errorf("getUsersOnlineSessionsList error: %s", userId, err)
			return nil, err
		}

		usersSessions[userId] = &status.SessionEntryList{
			Sessions: ss,
		}
	}

	return
}

func init() {
	Register("redis", redisStatusClientInstance)
}
