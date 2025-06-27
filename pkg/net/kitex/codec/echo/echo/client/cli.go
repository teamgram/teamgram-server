// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: Benqi (wubenqi@gmail.com)

package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	echo1client "github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo1/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo1/echo1"
	echo2client "github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo2/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo2/echo2"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
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

	cli1 := echo1client.MustNewKitexClient(c.Echo1) // echo1client.NewEcho1Client(echo1client.MustNewKitexClient(c.Echo1))
	cli2 := echo2client.MustNewKitexClient(c.Echo2)

	for {
		req1 := &echo1.TLEcho1Echo{
			ClazzID: echo1.ClazzID_echo1_echo,
			Message: "my reqeuest1",
		}

		resp1 := &echo1.Echo{}
		err := cli1.Call(context.Background(), "echo1.echo", req1, resp1)
		if err != nil {
			log.Fatal(err)
			return
		}

		// resp1, err := cli1.Echo1Echo(context.Background(), req1)
		logx.Debugf("resp1: %s", resp1)
		if err != nil {
			log.Fatal(err)
		}

		req2 := &echo2.TLEcho2Echo{
			ClazzID: echo2.ClazzID_echo2_echo,
			Message: "my reqeuest2",
		}

		resp2 := &echo2.Echo{}
		err = cli2.Call(context.Background(), "echo2.echo", req2, resp2)
		logx.Debugf("resp2: %s", resp2)
		if err != nil {
			log.Fatal(err)
		}

		//resp, err := cli2.EchoEcho(context.Background(), req)
		time.Sleep(time.Millisecond * 1000)
	}
}
