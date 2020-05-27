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

package service

import (
	"go.etcd.io/etcd/clientv3"
	"github.com/golang/glog"
	"time"
)

type Etcd struct {
	EtcCli   *clientv3.Client
	rootPath string
}

func NewEtcd(c *etcdConf) (etcd *Etcd, err error) {
	var etcdClient *clientv3.Client
	cfg := clientv3.Config{
		Endpoints:   c.Addrs,
		DialTimeout: time.Duration(c.Timeout),
	}
	if etcdClient, err = clientv3.New(cfg); err != nil {
		glog.Error("error: cannot connec to etcd:", err)
		return
	}
	etcd = &Etcd{
		EtcCli:   etcdClient,
		rootPath: c.Root,
	}
	return
}
