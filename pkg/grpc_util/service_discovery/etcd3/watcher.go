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

package etcd3

import (
	"encoding/json"

	etcd3 "go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/naming"
)

// EtcdWatcher is the implementation of grpc.naming.Watcher
type EtcdWatcher struct {
	key           string
	client        *etcd3.Client
	ctx           context.Context
	cancel        context.CancelFunc
	isInitialized bool
}

func (w *EtcdWatcher) Close() {
	w.cancel()
}

func newEtcdWatcher(key string, cli *etcd3.Client) naming.Watcher {
	ctx, cancel := context.WithCancel(context.Background())
	w := &EtcdWatcher{
		key:    key,
		client: cli,
		ctx:    ctx,
		cancel: cancel,
	}
	return w
}

func (w *EtcdWatcher) Next() ([]*naming.Update, error) {
	updates := make([]*naming.Update, 0)

	if !w.isInitialized {
		// query addresses from etcd
		resp, err := w.client.Get(w.ctx, w.key, etcd3.WithPrefix())
		w.isInitialized = true
		if err == nil {
			addrs := extractAddrs(resp)
			if len(addrs) > 0 {
				for _, v := range addrs {
					updates = append(updates, &naming.Update{Op: naming.Add, Addr: v.Addr, Metadata: &v.Metadata})
				}
				return updates, nil
			}
		} else {
			grpclog.Println("Etcd Watcher Get key error:", err)
		}
	}

	// generate etcd Watcher
	rch := w.client.Watch(w.ctx, w.key, etcd3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				nodeData := NodeData{}
				err := json.Unmarshal([]byte(ev.Kv.Value), &nodeData)
				if err != nil {
					grpclog.Println("Parse node data error:", err)
					continue
				}
				updates = append(updates, &naming.Update{Op: naming.Add, Addr: nodeData.Addr, Metadata: &nodeData.Metadata})
			case mvccpb.DELETE:
				nodeData := NodeData{}
				err := json.Unmarshal([]byte(ev.Kv.Value), &nodeData)
				if err != nil {
					grpclog.Println("Parse node data error:", err)
					continue
				}
				updates = append(updates, &naming.Update{Op: naming.Delete, Addr: nodeData.Addr, Metadata: &nodeData.Metadata})
			}
		}
		return updates, nil
	}
	return nil, nil
}

func extractAddrs(resp *etcd3.GetResponse) []NodeData {
	addrs := []NodeData{}

	if resp == nil || resp.Kvs == nil {
		return addrs
	}

	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			nodeData := NodeData{}
			err := json.Unmarshal(v, &nodeData)
			if err != nil {
				grpclog.Println("Parse node data error:", err)
				continue
			}
			addrs = append(addrs, nodeData)
		}
	}

	return addrs
}
