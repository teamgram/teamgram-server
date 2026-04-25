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
	"encoding/hex"
	"flag"
	"log"

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	echo1client "github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo1/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo1/echo1"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "cli.yaml", "the config file")

type Config struct {
	Echo1 kitex.RpcClientConf
}

func main() {
	flag.Parse()

	var c Config
	conf.MustLoad(*configFile, &c)

	rawInvoker := kitex.NewRawInvoker(map[string]kitex.Client{
		c.Echo1.ServiceName: echo1client.MustNewKitexClient(c.Echo1),
	})

	ctx := context.Background()
	ctx = metainfo.WithValue(ctx, "temp", "temp-value")
	ctx = metainfo.WithPersistentValue(ctx, "logid", "12345")

	rpcMetadata := metadata.RpcMetadata{
		ServerId:   "echo1-raw-client",
		ClientAddr: "127.0.0.1",
		Layer:      0,
		Client:     "raw-tl-example",
	}

	reqPayload, err := encodeEchoRequest("raw tl request")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("request payload: %s", hex.EncodeToString(reqPayload))

	rawResp, err := rawInvoker.InvokeRawMethod(
		ctx,
		&rpcMetadata,
		c.Echo1.ServiceName,
		"echo1.echo",
		reqPayload)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("response payload: %s", hex.EncodeToString(rawResp.Payload))

	resp := new(echo1.Echo)
	if err = resp.Decode(bin.NewDecoder(rawResp.Payload)); err != nil {
		log.Fatal(err)
	}
	log.Printf("response object: %s", resp)
}

func encodeEchoRequest(message string) ([]byte, error) {
	req := &echo1.TLEcho1Echo{
		ClazzID: echo1.ClazzID_echo1_echo,
		Message: message,
	}

	x := bin.NewEncoder()
	defer x.End()
	if err := req.Encode(x, 0); err != nil {
		return nil, err
	}
	return append([]byte(nil), x.Bytes()...), nil
}
