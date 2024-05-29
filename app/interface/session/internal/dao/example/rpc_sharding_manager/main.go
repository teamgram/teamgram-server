// Copyright © 2024 Teamgram Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/teamgram/marmota/pkg/net/ip"
	"github.com/teamgram/teamgram-server/app/interface/session/internal/config"
	"github.com/teamgram/teamgram-server/app/interface/session/internal/dao"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/session.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	pubListenOn := ip.FigureOutListenOn(c.ListenOn)
	sharding := dao.NewRpcShardingManager(pubListenOn, c.Etcd, func(sharding *dao.RpcShardingManager, oldList, addList, removeList []string) {
		for i := 0; i < 100; i++ {
			k := fmt.Sprintf("127.0.0.%d:8080", i)
			if sharding.ShardingVIsListenOn(k) {
				fmt.Println(k, "is listen on", pubListenOn)
			} else {
				fmt.Println(k, "not listen on", pubListenOn)
			}
		}
	})

	_ = sharding
	sharding.Start()

	time.Sleep(time.Hour)
}
