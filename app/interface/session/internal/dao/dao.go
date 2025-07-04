// Copyright 2024 Teamgram Authors
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

package dao

import (
	"github.com/teamgram/marmota/pkg/cache"
	"github.com/teamgram/marmota/pkg/net/ip"
	bff_proxy_client "github.com/teamgram/teamgram-server/v2/app/bff/bff/client"
	"github.com/teamgram/teamgram-server/v2/app/interface/session/internal/config"
	authsessionclient "github.com/teamgram/teamgram-server/v2/app/service/authsession/client"
	statusclient "github.com/teamgram/teamgram-server/v2/app/service/status/client"
)

type Dao struct {
	cache *cache.LRUCache
	authsessionclient.AuthsessionClient
	statusclient.StatusClient
	*bff_proxy_client.BFFProxyClient2
	eGateServers map[string]*Gateway
	MyServerId   string
	*RpcShardingManager
}

func New(c config.Config) *Dao {
	myServerId := ip.FigureOutListenOn(c.ListenOn)
	d := &Dao{
		cache:              cache.NewLRUCache(1024 * 1024 * 1024),
		AuthsessionClient:  authsessionclient.NewAuthsessionClient(authsessionclient.MustNewKitexClient(c.AuthSession)),
		BFFProxyClient2:    bff_proxy_client.NewBFFProxyClient2(c.BFFProxyClients),
		StatusClient:       statusclient.NewStatusClient(statusclient.MustNewKitexClient(c.StatusClient)),
		eGateServers:       make(map[string]*Gateway),
		MyServerId:         myServerId,
		RpcShardingManager: NewRpcShardingManager(myServerId, c.Etcd.EtcdConf),
	}

	d.watchGateway(c.GatewayClient)

	return d
}
