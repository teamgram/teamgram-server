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

	"github.com/teamgram/proto/v2/rpc/codec"
	"github.com/teamgram/teamgram-server/v2/app/interface/session/client"
	"github.com/teamgram/teamgram-server/v2/app/interface/session/session"
	"github.com/teamgram/teamgram-server/v2/app/interface/session/session/sessionservice"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"

	"github.com/cloudwego/kitex/client"
	"github.com/zeromicro/go-zero/core/discov"
)

func main() {
	c := kitex.RpcClientConf{
		Etcd:          discov.EtcdConf{Hosts: []string{"127.0.0.1:2379"}, Key: "interface.session"},
		Endpoints:     []string{"0.0.0.0:20120"},
		Target:        "",
		App:           "",
		Token:         "",
		NonBlock:      false,
		Timeout:       0,
		KeepaliveTime: 0,
	}
	_ = c
	cli := kitex.MustNewClient(
		c,
		func(opts ...client.Option) (client.Client, error) {
			return client.NewClient(sessionservice.NewServiceInfoForClient(), opts...)
		},
		client.WithCodec(codec.NewZRpcCodec(true)))

	cli2 := sessionclient.NewSessionClient(cli)
	for {
		req := &session.TLSessionQueryAuthKey{
			ClazzID:   session.ClazzID_session_queryAuthKey,
			AuthKeyId: 1234567890,
		}
		resp, err := cli2.SessionQueryAuthKey(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%s\n", resp)
		time.Sleep(time.Second)
	}
}
