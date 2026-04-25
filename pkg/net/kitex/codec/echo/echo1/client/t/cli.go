// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
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
	"context"
	"flag"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	echo1client "github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo1/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo1/echo1"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/zeromicro/go-zero/core/conf"
	"log"
	"time"

	_ "github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo2/echo2/echo2service"
)

var configFile = flag.String("f", "cli.yaml", "the config file")

type Config struct {
	Echo1 kitex.RpcClientConf
	Echo2 kitex.RpcClientConf
}

func main() {
	flag.Parse()

	var (
		c Config
	)

	conf.MustLoad(*configFile, &c)

	cli := echo1client.NewEcho1Client(kitex.GetCachedKitexClient(c.Echo1))
	for {
		ctx := context.Background()
		ctx = metainfo.WithValue(ctx, "temp", "temp-value")       // only present in next service
		ctx = metainfo.WithPersistentValue(ctx, "logid", "12345") // will present in the next service and its successors
		md2 := metadata.RpcMetadata{
			ServerId:      "12345",
			ClientAddr:    "12345",
			AuthId:        0,
			SessionId:     0,
			ReceiveTime:   0,
			UserId:        0,
			ClientMsgId:   0,
			IsBot:         false,
			Layer:         0,
			Client:        "",
			IsAdmin:       false,
			Takeout:       nil,
			Langpack:      "",
			PermAuthKeyId: 0,
			LangCode:      "",
		}
		ctx, _ = metadata.RpcMetadataToOutgoing(ctx, &md2)

		req1 := &echo1.TLEcho1Echo{
			ClazzID: echo1.ClazzID_echo1_echo,
			Message: "my reqeuest1",
		}
		resp, err := cli.Echo1Echo(context.Background(), req1)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
