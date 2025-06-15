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
	"errors"
	"flag"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	echoclient "github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/examples/echo/client"
	api "github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/examples/echo/echo"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/hash"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	ErrEchoNotFound = errors.New("not found echo")
)

var configFile = flag.String("f", "t.yaml", "the config file")

type Config struct {
	EchoClient kitex.RpcClientConf
}

type EchoClient struct {
	dispatcher *hash.ConsistentHash
	cList      map[string]echoclient.EchoClient
}

func NewEchoClient(c Config) *EchoClient {
	cli := &EchoClient{
		dispatcher: hash.NewConsistentHash(),
		cList:      make(map[string]echoclient.EchoClient),
	}
	cli.watch(c.EchoClient)

	return cli
}

func (cli *EchoClient) watch(c kitex.RpcClientConf) {
	sub, err := discov.NewSubscriber(c.Etcd.Hosts, c.Etcd.Key)
	if err != nil {
		log.Fatalf("NewSubscriber(%+v) error(%v)", c.Etcd, err)
	}

	update := func() {
		var (
			addCliList    []string
			removeCliList []string
		)

		values := sub.Values()
		sessions := map[string]echoclient.EchoClient{}
		for _, v := range values {
			if old, ok := cli.cList[v]; ok {
				sessions[v] = old
				continue
			}
			c.Endpoints = []string{v}

			//cli, err := zrpc.NewClient(c)
			//if err != nil {
			//	logx.Error("watchComet NewClient(%+v) error(%v)", values, err)
			//	return
			//}
			//sessionCli := sessionclient.NewSessionClient(cli)
			sessions[v] = echoclient.NewEchoClient(echoclient.MustNewKitexClient(c))

			addCliList = append(addCliList, v)
		}

		for key, _ := range cli.cList {
			if !stringx.Contains(values, key) {
				removeCliList = append(removeCliList, key)
			}
		}

		for _, n := range addCliList {
			cli.dispatcher.Add(n)
		}

		for _, n := range removeCliList {
			cli.dispatcher.Remove(n)
		}

		cli.cList = sessions
	}

	sub.AddListener(update)
	update()
}

func (cli *EchoClient) InvokeByKey(key string, cb func(client echoclient.EchoClient) (err error)) error {
	val, ok := cli.dispatcher.Get(key)
	if !ok {
		return ErrEchoNotFound
	}

	cli2, ok := cli.cList[val.(string)]
	if !ok {
		return ErrEchoNotFound
	}

	if cb == nil {
		return nil
	}

	return cb(cli2)
}

func main() {
	flag.Parse()

	var (
		c Config
	)

	conf.MustLoad(*configFile, &c)

	cli1 := echoclient.NewEchoClient(echoclient.MustNewKitexClient(c.EchoClient))
	cli2 := NewEchoClient(c)
	for {
		req := &api.TLEchoEcho{
			ClazzID: api.ClazzID_echo_echo,
			Message: "my request",
		}

		v := rand.Int63()
		if v%2 == 0 {
			resp, err := cli1.EchoEcho(context.Background(), req)
			logx.Debugf("resp: %s", resp)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			_ = cli2.InvokeByKey(strconv.FormatInt(v, 10), func(client echoclient.EchoClient) (err error) {
				resp, err := client.EchoEcho(context.Background(), req)
				logx.Debugf("resp: %s", resp)
				if err != nil {
					log.Fatal(err)
				}

				return err
			})
		}

		//resp, err := cli2.EchoEcho(context.Background(), req)
		time.Sleep(time.Millisecond * 100)
	}
}
