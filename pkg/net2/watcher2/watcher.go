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

package watcher2

import (
	"go.etcd.io/etcd/clientv3"
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/net2"
)

// see: /baselib/grpc_util/service_discovery/registry.go
type nodeData struct {
	Addr     string
	Metadata map[string]string
}

// TODO(@benqi): grpc_util/serviec_discovery集成
type ClientWatcher struct {
	etcCli      *clientv3.Client
	registryDir string
	serviceName string
	// rootPath    string
	client *net2.TcpClientGroupManager
	nodes  map[string]*nodeData
}

func NewClientWatcher(registryDir, serviceName string, cfg clientv3.Config, client *net2.TcpClientGroupManager) (watcher *ClientWatcher, err error) {
	var etcdClient *clientv3.Client
	if etcdClient, err = clientv3.New(cfg); err != nil {
		glog.Error("Error: cannot connec to etcd:", err)
		return
	}

	watcher = &ClientWatcher{
		etcCli:      etcdClient,
		registryDir: registryDir,
		serviceName: serviceName,
		client:      client,
		nodes:       map[string]*nodeData{},
	}
	return
}

func (m *ClientWatcher) WatchClients(cb func(etype, addr string)) {
	if m == nil {
		return
	}

	rootPath := fmt.Sprintf("%s/%s", m.registryDir, m.serviceName)

	resp, err := m.etcCli.Get(context.Background(), rootPath, clientv3.WithPrefix())
	if err != nil {
		glog.Error(err)
	}
	for _, kv := range resp.Kvs {
		m.addClient(kv, cb)
	}

	rch := m.etcCli.Watch(context.Background(), rootPath, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			if ev.Type.String() == "EXPIRE" {
				// TODO(@benqi): 采用何种策略？？
				// n, ok := m.nodes[string(ev.Kv.Key)]
				// if ok {
				//	 delete(m.nodes, string(ev.Kv.Key))
				// }
				// if cb != nil {
				// 	cb("EXPIRE", string(ev.Kv.Key), string(ev.Kv.Value))
				//}
			} else if ev.Type.String() == "PUT" {
				m.addClient(ev.Kv, cb)
			} else if ev.Type.String() == "DELETE" {
				if n, ok := m.nodes[string(ev.Kv.Key)]; ok {
					m.client.RemoveClient(m.serviceName, n.Addr)
					if cb != nil {
						cb("delete", n.Addr)
					}
					delete(m.nodes, string(ev.Kv.Key))
				}
			}
		}
	}
}
func (m *ClientWatcher) addClient(kv *mvccpb.KeyValue, cb func(etype, addr string)) {
	node := &nodeData{}
	err := json.Unmarshal(kv.Value, node)
	if err != nil {
		glog.Error(err)
	}
	if n, ok := m.nodes[string(kv.Key)]; ok {
		if node.Addr != n.Addr {
			m.client.RemoveClient(m.serviceName, n.Addr)
			m.nodes[string(kv.Key)] = node
			if cb != nil {
				cb("delete", n.Addr)
			}
			if cb != nil {
				cb("add", node.Addr)
			}
		}
		m.client.AddClient(m.serviceName, node.Addr)
	} else {
		m.nodes[string(kv.Key)] = node
		m.client.AddClient(m.serviceName, node.Addr)
		if cb != nil {
			cb("add", node.Addr)
		}
	}
}
