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
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/mtproto/rpc"
)

var (
	confPath string
	Conf     *frontendConfig
)

type frontendConfig struct {
	ServerId   int32 // 服务器ID
	Server80   *mtproto.MTProtoServerConfig
	Server443  *mtproto.MTProtoServerConfig
	Server5222 *mtproto.MTProtoServerConfig
	Clients    *zrpc.ZRpcClientConfig
}

func (c *frontendConfig) String() string {
	return fmt.Sprintf("{server_id: %d, server80: %v. server443: %v, server5222: %v, clients: %v}",
		c.ServerId,
		c.Server80,
		c.Server443,
		c.Server5222,
		c.Clients)
}

func InitializeConfig() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	if err != nil {
		err = fmt.Errorf("decode file %s error: %v", confPath, err)
	}
	return
}

func init() {
	flag.StringVar(&confPath, "conf", "./frontend.toml", "config path")
}
