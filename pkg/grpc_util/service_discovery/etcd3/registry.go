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
	"fmt"
	etcd3 "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc/grpclog"
	"time"
)

type EtcdReigistry struct {
	etcd3Client *etcd3.Client
	key         string
	value       string
	ttl         time.Duration
	ctx         context.Context
	cancel      context.CancelFunc
}

type Option struct {
	EtcdConfig  etcd3.Config
	RegistryDir string
	ServiceName string
	NodeID      string
	NData       NodeData
	Ttl         time.Duration
}

type NodeData struct {
	Addr     string
	Metadata map[string]string
}

func NewRegistry(option Option) (*EtcdReigistry, error) {
	client, err := etcd3.New(option.EtcdConfig)
	if err != nil {
		return nil, err
	}

	val, err := json.Marshal(option.NData)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	registry := &EtcdReigistry{
		etcd3Client: client,
		key:         option.RegistryDir + "/" + option.ServiceName + "/" + option.NodeID,
		value:       string(val),
		ttl:         option.Ttl,
		ctx:         ctx,
		cancel:      cancel,
	}

	fmt.Println(registry)
	return registry, nil
}

func (e *EtcdReigistry) Register() error {

	insertFunc := func() error {
		// fmt.Println("Grant: ", e)
		resp, err := e.etcd3Client.Grant(e.ctx, 10) // int64(e.ttl))
		if err != nil {
			fmt.Println("Grant error: ", err)
			return err
		}
		_, err = e.etcd3Client.Get(e.ctx, e.key)
		if err != nil {
			if err == rpctypes.ErrKeyNotFound {
				if _, err := e.etcd3Client.Put(e.ctx, e.key, e.value, etcd3.WithLease(resp.ID)); err != nil {
					grpclog.Printf("grpclb: set key '%s' with ttl to etcd3 failed: %s", e.key, err.Error())
				}
			} else {
				grpclog.Printf("grpclb: key '%s' connect to etcd3 failed: %s", e.key, err.Error())
			}
			return err
		} else {
			// refresh set to true for not notifying the watcher
			// fmt.Println("refresh: ", resp)
			if _, err := e.etcd3Client.Put(e.ctx, e.key, e.value, etcd3.WithLease(resp.ID)); err != nil {
				grpclog.Printf("grpclb: refresh key '%s' with ttl to etcd3 failed: %s", e.key, err.Error())
				return err
			}
		}
		return nil
	}

	err := insertFunc()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(e.ttl / 5)
	for {
		select {
		case <-ticker.C:
			insertFunc()
		case <-e.ctx.Done():
			ticker.Stop()
			if _, err := e.etcd3Client.Delete(context.Background(), e.key); err != nil {
				grpclog.Printf("grpclb: deregister '%s' failed: %s", e.key, err.Error())
			}
			glog.Info("Delete...")
			return nil
		}
	}

	return nil
}

func (e *EtcdReigistry) Deregister() error {
	e.cancel()
	return nil
}
