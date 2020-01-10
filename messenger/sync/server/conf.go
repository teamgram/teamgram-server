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
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/mysql_client"
	"github.com/nebula-chat/chatengine/pkg/redis_client"
	"github.com/nebula-chat/chatengine/mtproto/rpc"
	"github.com/nebula-chat/chatengine/pkg/util"
)

var (
	confPath string
	Conf     *syncConfig
)

type syncConfig struct {
	ServerId      int32 // 服务器ID
	Redis         []redis_client.RedisConfig
	Mysql         []mysql_client.MySQLConfig
	Server        *grpc_util.RPCServerConfig
	SessionClient *zrpc.ZRpcClientConfig
}

func (c *syncConfig) String() string {
	return fmt.Sprintf("{server_id: %d, redis: %v. mysql: %v, server: %v, sessionClient: %v}",
		c.ServerId,
		c.Redis,
		c.Mysql,
		c.Server,
		c.SessionClient)
}

func init() {
	tomlPath := util.GetWorkingDirectory() + "/sync.toml"
	flag.StringVar(&confPath, "conf", tomlPath, "config path")
}

func InitializeConfig() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	if err != nil {
		err = fmt.Errorf("decode file %s error: %v", confPath, err)
	}
	return
}
