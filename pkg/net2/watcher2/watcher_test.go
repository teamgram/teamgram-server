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
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"github.com/nebula-chat/chatengine/pkg/net2"
	"testing"
	"time"
)

const etcdAddr = "http://127.0.0.1:2379"

func AddService(namespace string, serviceName string, nodeID string, addr string) error {
	etcdConfig := clientv3.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
	}
	cli, err := clientv3.New(etcdConfig)
	if err != nil {
		return err
	}
	defer cli.Close()

	nodeData := &nodeData{Addr: addr}
	val, err := json.Marshal(nodeData)
	if err != nil {
		return err
	}

	_, err = cli.Put(context.TODO(), fmt.Sprintf("/%s/%s/%s", namespace, serviceName, nodeID), string(val))
	if err != nil {
		return err
	}
	return nil
}

func TestRegisterBeforeWatcher(t *testing.T) {
	etcdConfig := clientv3.Config{
		Endpoints: []string{etcdAddr},
	}

	services := make(map[string][]string)
	client := net2.NewTcpClientGroupManager("TestProto", services, nil)
	w, _ := NewClientWatcher("/nebulaim", "test_before", etcdConfig, client)

	AddService("nebulaim", "test_before", "node1", "0.0.0.0:12345")
	AddService("nebulaim", "test_before", "node2", "0.0.0.0:98765")

	receivedChan := make(chan interface{})

	go w.WatchClients(func(etype, addr string) {
		receivedChan <- fmt.Sprintf(`watcher_action_type:%s, watcher_action_value:%s`, etype, addr)
	})

	<-receivedChan
	<-receivedChan
}

func TestRegisterAfterWatcher(t *testing.T) {
	etcdConfig := clientv3.Config{
		Endpoints: []string{etcdAddr},
	}

	services := make(map[string][]string)
	client := net2.NewTcpClientGroupManager("TestProto", services, nil)
	w, _ := NewClientWatcher("/nebulaim", "test_after", etcdConfig, client)

	receivedChan := make(chan interface{})

	go w.WatchClients(func(etype, addr string) {
		//fmt.Printf(`watcher_action_type:%s, watcher_action_value:%s\n`, etype, addr)
		receivedChan <- fmt.Sprintf(`watcher_action_type:%s, watcher_action_value:%s`, etype, addr)
	})

	time.Sleep(time.Second)

	go AddService("nebulaim", "test_after", "test_node1", "0.0.0.0:98765")
	go AddService("nebulaim", "test_after", "test_node2", "0.0.0.0:12345")

	<-receivedChan
	<-receivedChan
}
