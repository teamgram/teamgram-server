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
	"flag"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery/etcd3"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery/examples/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

var nodeID = flag.String("node", "node1", "node ID")
var port = flag.Int("port", 8080, "listening port")

type RpcServer struct {
	addr string
	s    *grpc.Server
}

func NewRpcServer(addr string) *RpcServer {
	s := grpc.NewServer()
	rs := &RpcServer{
		addr: addr,
		s:    s,
	}
	return rs
}

func (s *RpcServer) Run() {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return
	}
	log.Printf("rpc listening on:%s", s.addr)

	proto.RegisterEchoServiceServer(s.s, s)
	s.s.Serve(listener)
}

func (s *RpcServer) Stop() {
	s.s.GracefulStop()
}

func (s *RpcServer) Echo(ctx context.Context, req *proto.EchoReq) (*proto.EchoRsp, error) {
	text := "Hello " + req.EchoData + ", I am " + *nodeID
	log.Println(text)

	return &proto.EchoRsp{EchoData: text}, nil
}

func StartService() {
	etcdConfg := clientv3.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
	}

	registry, err := etcd3.NewRegistry(
		etcd3.Option{
			EtcdConfig:  etcdConfg,
			RegistryDir: "/grpc-lb",
			ServiceName: "test",
			NodeID:      *nodeID,
			NData: etcd3.NodeData{
				Addr: fmt.Sprintf("127.0.0.1:%d", *port),
				//Metadata: map[string]string{"weight": "1"},
			},
			Ttl: 10, // * time.Second,
		})
	if err != nil {
		log.Panic(err)
		return
	}
	server := NewRpcServer(fmt.Sprintf("0.0.0.0:%d", *port))
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		server.Run()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		registry.Register()
		wg.Done()
	}()

	//stop the server after one minute
	//go func() {
	//	time.Sleep(time.Minute)
	//	server.Stop()
	//	registry.Deregister()
	//}()

	wg.Wait()
}

//go run server.go -node node1 -port 28544
//go run server.go -node node2 -port 18562
//go run server.go -node node3 -port 27772
func main() {
	flag.Parse()
	StartService()
}
