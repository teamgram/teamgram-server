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

package redis_dao

import (
	"github.com/nebula-chat/chatengine/pkg/redis_client"
	"testing"
)

func TestNextID(t *testing.T) {
	redisConfig := &redis_client.RedisConfig{
		Name:         "test",
		Addr:         "127.0.0.1:6379",
		Idle:         100,
		Active:       100,
		DialTimeout:  1000000,
		ReadTimeout:  1000000,
		WriteTimeout: 1000000,
		IdleTimeout:  15000000,
		DBNum:        "0",
		Password:     "",
	}

	redisPool := redis_client.NewRedisPool(redisConfig)

	_ := NewSequenceDAO(redisPool)
	//seq.NextID("1")
	//seq.NextID("1")
	//seq.NextID("1")
	//seq.NextID("1")
	//seq.NextID("2")
	//seq.NextID("2")
	//seq.NextID("2")
	//seq.NextID("2")
}
