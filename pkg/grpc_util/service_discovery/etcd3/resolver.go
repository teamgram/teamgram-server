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
	"errors"
	"fmt"
	etcd3 "go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/naming"
)

// EtcdResolver is an implementation of grpc.naming.Resolver
type EtcdResolver struct {
	Config      etcd3.Config
	RegistryDir string
	ServiceName string
}

func NewResolver(registryDir, serviceName string, cfg etcd3.Config) naming.Resolver {
	return &EtcdResolver{RegistryDir: registryDir, ServiceName: serviceName, Config: cfg}
}

// Resolve to resolve the service from etcd
func (er *EtcdResolver) Resolve(target string) (naming.Watcher, error) {
	if er.ServiceName == "" {
		return nil, errors.New("no service name provided")
	}
	client, err := etcd3.New(er.Config)
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%s/%s", er.RegistryDir, er.ServiceName)
	return newEtcdWatcher(key, client), nil
}
