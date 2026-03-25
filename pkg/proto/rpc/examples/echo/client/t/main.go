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

	"github.com/teamgram/teamgram-server/v2/pkg/proto/rpc/codec"
	echoclient "github.com/teamgram/teamgram-server/v2/pkg/proto/rpc/examples/echo/client"
	api "github.com/teamgram/teamgram-server/v2/pkg/proto/rpc/examples/echo/echo"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/rpc/examples/echo/echo/echo"

	"github.com/cloudwego/kitex/client"
)

func main() {
	zCodec := codec.NewZRpcCodec(true)
	cli, err := echo.NewClient("echo", client.WithHostPorts("0.0.0.0:8888"), client.WithCodec(zCodec))
	if err != nil {
		log.Fatal(err)
	}

	cli2 := echoclient.NewEchoClient(cli)
	for {
		req := &api.TLEchoEcho{
			ClazzID: api.ClazzID_echo_echo,
			Message: "my request",
		}
		resp, err := cli2.EchosEcho(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
