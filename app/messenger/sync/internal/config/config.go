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

package config

import (
	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

// Routine routine.
type Routine struct {
	Size uint64
	Chan uint64
}

type Config struct {
	zrpc.RpcServerConf
	Mysql         sqlx.Config
	Cache         cache.CacheConf
	KV            kv.KvConf
	Routine       Routine
	SyncConsumer  kafka.KafkaConsumerConf
	SessionClient zrpc.RpcClientConf
	IdgenClient   zrpc.RpcClientConf
	StatusClient  zrpc.RpcClientConf
	ChatClient    zrpc.RpcClientConf
	PushClient    *kafka.KafkaProducerConf `json:",optional"`
}
