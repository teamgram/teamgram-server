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

package idgen

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/redis_client"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type RedisSeqClient struct {
	redis *redis_client.RedisPool
}

func NewRedisSeqClient(redis *redis_client.RedisPool) *RedisSeqClient {
	cli := &RedisSeqClient{redis}
	return cli
}

func redisSeqClientInstance() SeqIDGen {
	return &RedisSeqClient{}
}

func (c *RedisSeqClient) Initialize(config string) error {
	c.redis = redis_client.GetRedisClient(config)
	if c.redis == nil {
		return fmt.Errorf("init redisSeqClient error: %s", config)
	}
	return nil
}

func (c *RedisSeqClient) GetCurrentSeqID(key string) (seq int64, err error) {
	conn := c.redis.Get()
	defer conn.Close()

	seq, err = redis.Int64(conn.Do("GET", key))
	if err != nil {
		glog.Errorf("redis_seq_client.GetCurrentSeqID - GET {%s}, error: {%v}", key, err)
	}

	return
}

func (c *RedisSeqClient) GetNextSeqID(key string) (seq int64, err error) {
	conn := c.redis.Get()
	defer conn.Close()

	// 设置键
	seq, err = redis.Int64(conn.Do("INCR", key))
	if err != nil {
		glog.Errorf("redis_seq_client.GetNextSeqID - INCR {%s}, error: {%v}", key, err)
	}

	return
}

func (c *RedisSeqClient) GetNextNSeqID(key string, n int) (seq int64, err error) {
	conn := c.redis.Get()
	defer conn.Close()

	// 设置键
	seq, err = redis.Int64(conn.Do("INCRBY", key, n))
	if err != nil {
		glog.Errorf("redis_seq_client.GetNextNSeqID - INCR {%s}, error: {%v}", key, err)
	}

	return
}

func init() {
	SeqIDGenRegister("redis", redisSeqClientInstance)
}
