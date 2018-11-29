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
	"github.com/nebula-chat/chatengine/pkg/util"
)

// Redis client config.
type RedisConfig struct {
	Name         string // redis name
	Addr         string
	Active       int // pool
	Idle         int // pool
	DialTimeout  util.Duration
	ReadTimeout  util.Duration
	WriteTimeout util.Duration
	IdleTimeout  util.Duration

	DBNum    string // db号
	Password string // 密码
}

func (c *RedisConfig) ToRedisCacheConfig() string {
	// config is like {"key":"collection key","conn":"connection info","dbNum":"0"}
	// rc.key = cf["key"]
	// rc.conninfo = cf["conn"]
	// rc.dbNum, _ = strconv.Atoi(cf["dbNum"])
	// rc.password = cf["password"]
	return fmt.Sprintf(`{"conn":"%s", "dbNum":"%s", "password":"%s"}`,
		c.Addr,
		c.DBNum,
		c.Password)
}
