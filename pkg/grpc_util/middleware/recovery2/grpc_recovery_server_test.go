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

package grpc_recovery2

import (
	"context"
	"fmt"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/middleware/examples/zproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"testing"
)

type ChatTestServiceImpl struct {
}

func (s *ChatTestServiceImpl) Connect(request *zproto.ChatSession, stream zproto.ChatTest_ConnectServer) (err error) {
	return
}

func (s *ChatTestServiceImpl) SendChat(ctx context.Context, request *zproto.ChatMessage) (reply *zproto.VoidRsp2, err error) {
	fmt.Printf("%v.SendChat(_) = _, %v\n", ctx, request)

	switch request.MessageData {
	case "panic":
		panic("very bad thing happened")
	case "nil":
		panic("nil thing happened")
	}
	return &zproto.VoidRsp2{}, nil
}

func unaryRecoveryHandler(ctx context.Context, p interface{}) (err error) {
	return status.Errorf(codes.Unknown, "panic triggered: %v", p)
}

func unaryRecoveryHandler2(ctx context.Context, p interface{}) (err error) {
	return status.Errorf(codes.Unknown, "panic triggered: %v", p)
}

func streamRecoveryHandler(stream grpc.ServerStream, p interface{}) (err error) {
	return
}

func TestRecoveryServer(t *testing.T) {
	lis, err := net.Listen("tcp", "0.0.0.0:22345")
	if err != nil {
		panic(err)
		// glog.Fatalf("failed to listen: %v", err)
	}

	server := NewRecoveryServer2(unaryRecoveryHandler, unaryRecoveryHandler2, streamRecoveryHandler)
	zproto.RegisterChatTestServer(server, &ChatTestServiceImpl{})
	server.Serve(lis)
}
