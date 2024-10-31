// Copyright © 2024 Teamgram Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"math/rand"
	"os"

	"github.com/teamgram/teamgram-server/pkg/net2/brpc"
)

func main() {
	msg := &brpc.BaiduRpcMessage{
		Meta: &brpc.RpcMeta{
			Request: &brpc.RpcRequestMeta{
				ServiceName:  "Test",
				MethodName:   "test",
				LogId:        rand.Int63(),
				TraceId:      rand.Int63(),
				SpanId:       rand.Int63(),
				ParentSpanId: rand.Int63(),
			},
			Response: &brpc.RpcResponseMeta{
				ErrorCode: 0,
				ErrorText: "",
			},
			CompressType:   0,
			CorrelationId:  rand.Int63(),
			AttachmentSize: 0,
			MtprotoMeta:    nil,
		},
		Payload:    make([]byte, rand.Int()%2048),
		Attachment: make([]byte, rand.Int()%1024),
	}

	data, _ := msg.Encode()
	os.WriteFile("test.data", bytes.Join(data, []byte{}), 0644)
}
