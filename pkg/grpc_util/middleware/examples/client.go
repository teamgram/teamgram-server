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
	"context"
	"fmt"
	"github.com/nebula-chat/chatengine/pkg/util"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/middleware/examples/zproto"
	"google.golang.org/grpc"
	"io"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	conn, err := grpc.Dial("127.0.0.1:22345", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("fail to dial: %v\n", err)
	}
	defer conn.Close()
	client := zproto.NewChatTestClient(conn)
	sess := &zproto.ChatSession{util.Int64ToString(rand.Int63())}
	fmt.Println("sessionId : ", sess.SessionId)

	var message string
	for {
		fmt.Print("> ")
		if n, err := fmt.Scanln(&message); err == io.EOF {
			return
		} else if n > 0 {
			if message == "quit" {
				return
			} else {
				_, err := client.SendChat(context.Background(), &zproto.ChatMessage{SenderSessionId: sess.SessionId, MessageData: message})
				if err != nil {
					fmt.Printf("%v.SendChat(_) = _, %v\n", client, err)
				}
			}
		}
	}
}
