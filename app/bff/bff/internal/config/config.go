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
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

type Config struct {
	kitex.RpcServerConf
	//KV                        kv.KvConf
	//Code                      *conf.SmsVerifyCodeConfig
	//BizServiceClient          kitex.RpcClientConf
	//AuthSessionClient         kitex.RpcClientConf
	//MediaClient               kitex.RpcClientConf
	//IdgenClient               kitex.RpcClientConf
	//MsgClient                 kitex.RpcClientConf
	//SyncClient                *kafka.KafkaProducerConf
	//DfsClient                 kitex.RpcClientConf
	//StatusClient              kitex.RpcClientConf
	//SignInServiceNotification []conf.MessageEntityConfig `json:",optional"`
	//SignInMessage             []conf.MessageEntityConfig `json:",optional"`
}
