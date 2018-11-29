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

package main

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/load_balancer"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery/etcd3"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery/examples/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	etcdConfg := clientv3.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
	}
	r := etcd3.NewResolver("/grpc-lb", "test", etcdConfg)
	b := load_balancer.NewBalancer(r, load_balancer.NewRoundRobinSelector())
	c, err := grpc.Dial("", grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithTimeout(time.Second*5))
	if err != nil {
		log.Printf("grpc dial: %s", err)
		return
	}
	defer c.Close()

	client := proto.NewEchoServiceClient(c)

	for i := 0; i < 1000; i++ {
		resp, err := client.Echo(context.Background(), &proto.EchoReq{EchoData: "round robin"})
		if err != nil {
			log.Println("aa:", err)
			time.Sleep(time.Second)
			continue
		}
		log.Printf(resp.EchoData)
		time.Sleep(time.Second)
	}
}
