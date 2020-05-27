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

package etcd_util

import (
	"go.etcd.io/etcd/clientv3"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery/etcd3"
	"time"
)

func NewEtcdRegistry(discovery service_discovery.ServiceDiscoveryServerConfig) (*etcd3.EtcdReigistry, error) {
	etcdConfg := clientv3.Config{
		Endpoints: discovery.EtcdAddrs,
	}

	option := etcd3.Option{
		EtcdConfig:  etcdConfg,
		RegistryDir: "/nebulaim",
		ServiceName: discovery.ServiceName,
		NodeID:      discovery.NodeID,
		NData: etcd3.NodeData{
			Addr:     discovery.RPCAddr,
			Metadata: map[string]string{},
		},
		Ttl: time.Duration(discovery.TTL),
	}

	return etcd3.NewRegistry(option)
}
