// Copyright 2022 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	gnetwayclient "github.com/teamgram/teamgram-server/v2/app/interface/gnetway/client"
	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/gnetway"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"

	"github.com/zeromicro/go-zero/core/logx"
)

// Gateway eGateClient is a gateway.
type Gateway struct {
	serverId     string
	client       gnetwayclient.GnetwayClient
	ctx          context.Context
	cancel       context.CancelFunc
	eGateServers map[string]*Gateway
}

// NewGateway NewSession new a comet.
func NewGateway(c kitex.RpcClientConf) (*Gateway, error) {
	g := &Gateway{
		serverId: c.Endpoints[0],
	}

	//cli, err := zrpc.NewClient(
	//	c,
	//	zrpc.WithDialOption(grpc.WithReadBufferSize(16*1024*1024)),
	//	zrpc.WithDialOption(grpc.WithWriteBufferSize(16*1024*1024)))
	//if err != nil {
	//	logx.Errorf("watchComet NewClient(%+v) error(%v)", c, err)
	//	return nil, err
	//}
	//cli :=
	g.client = gnetwayclient.NewGnetwayClient(gnetwayclient.MustNewKitexClient(c))
	g.ctx, g.cancel = context.WithCancel(context.Background())

	return g, nil
}

func (c *Gateway) Close() (err error) {
	c.cancel()
	return
}

// SendDataToGate
// egate.sendDataToGateway auth_key_id:long session_id:long payload:bytes = Bool;
func (c *Gateway) SendDataToGate(ctx context.Context, authKeyId, sessionId int64, payload []byte) (b bool, err error) {
	var (
		res *tg.Bool
	)

	res, err = c.client.GnetwaySendDataToGateway(ctx, &gnetway.TLGnetwaySendDataToGateway{
		AuthKeyId: authKeyId,
		SessionId: sessionId,
		Payload:   payload,
	})

	if err != nil {
		logx.Errorf("sendDataToGate error: %v", err)
		b = false
		return
	}

	b = tg.FromBool(res)
	return
}
