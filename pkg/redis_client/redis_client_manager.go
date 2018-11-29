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

package redis_client

import (
	"fmt"
	"github.com/golang/glog"
)

type redisClientManager struct {
	// TODO(@benqi): 使用sync.Map，动态添加和卸载数据库
	redisClients map[string]*RedisPool
}

var redisClients = &redisClientManager{make(map[string]*RedisPool)}

func InstallRedisClientManager(configs []RedisConfig) {
	for _, config := range configs {
		client := NewRedisPool(&config)
		if client == nil {
			err := fmt.Errorf("InstallRedisClient - NewRedisPool {%v} error!", config)
			panic(err)
			// continue
		}

		// TODO(@benqi): 检查config数据合法性
		redisClients.redisClients[config.Name] = client
	}
}

func GetRedisClient(redisName string) (client *RedisPool) {
	client, ok := redisClients.redisClients[redisName]
	if !ok {
		glog.Errorf("getRedisClient - Not found client: %s", redisName)
	}
	return
}

func GetRedisClientManager() map[string]*RedisPool {
	return redisClients.redisClients
}
