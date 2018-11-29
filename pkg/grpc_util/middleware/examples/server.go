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
	"net"
	"log"
	"golang.org/x/net/context"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/middleware/examples/helloworld"
	"google.golang.org/grpc"
)

// GreeterServer is the server API for Greeter service.
type GreeterServerImpl struct {
}

func (s *GreeterServerImpl) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	r := &helloworld.HelloReply{
		Message: request.Name,
	}
	return r, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return
	}
	log.Printf("rpc listening on 0.0.0.0:8100")

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &GreeterServerImpl{})
	// proto.RegisterEchoServiceServer(s.s, s)
	s.Serve(listener)
}
