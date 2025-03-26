// Copyright 2021 CloudWeGo Authors
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
//

package main

import (
	"context"
	"log"
	"time"

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	echoclient "github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/examples/echo/client"
	api "github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/examples/echo/echo"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
)

func main() {
	// zCodec := codec.NewZRpcCodec(true)
	c := kitex.RpcClientConf{
		ServiceName:   "examples.echo",
		Codec:         "zrpc",
		Etcd:          discov.EtcdConf{},
		Endpoints:     []string{"127.0.0.1:8888"},
		Target:        "",
		App:           "",
		Token:         "",
		NonBlock:      false,
		Timeout:       0,
		KeepaliveTime: 0,
	}

	cli2 := echoclient.NewEchoClient(echoclient.MustNewKitexClient(c))
	for {
		req := &api.TLEchoEcho{
			ClazzID: api.ClazzID_echo_echo,
			Message: "my request",
		}
		resp, err := cli2.EchoEcho(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		logx.Debugf("resp: %s", resp)
		time.Sleep(time.Second)
	}
}
