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
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery"
	"fmt"
	"github.com/nebula-chat/chatengine/service/status/client"
	"math/rand"
)

type SessionID struct {
	userId    int32
	serverId  int32
	authKeyId int64
}

func genRandomSession(n int32) SessionID {
	return SessionID{rand.Int31n(n) + 1, 1, rand.Int63()}
}

func main() {
	discovery := &service_discovery.ServiceDiscoveryClientConfig{
		ServiceName: "status",
		EtcdAddrs:   []string{"http://127.0.0.1:2379"},
		Balancer:    "round_robin",
	}
	cli := status_client.NewRpcStatusClient(discovery)

	for i := 0; i < 100; i++ {
		// s := genRandomSession(100)
		// cli.SetSessionOnline(s.userId, s.serverId, s.authKeyId)
	}

	for i := 0; i < 100; i++ {
		userId := rand.Int31n(100)
		slist, _ := cli.GetUserOnlineSessions(userId)
		fmt.Println("user: ", userId, ", status: ", slist)
	}

	//for i := 0; i < 1000; i++ {
	//	userId := rand.Int31n(1000)
	//	slist, _ := cli.GetUserOnlineSessions(userId)
	//	if len(slist.Sessions) == 0 {
	//		fmt.Println("offline: ", userId)
	//	}
	//}
}
