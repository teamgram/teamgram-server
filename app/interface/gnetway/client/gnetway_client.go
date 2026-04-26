/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package gnetwayclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/gnetway"
	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/gnetway/gnetwayservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type GnetwayClient interface {
	GnetwaySendDataToGateway(ctx context.Context, in *gnetway.TLGnetwaySendDataToGateway) (*tg.Bool, error)
}

type defaultGnetwayClient struct {
	cli client.Client
	rpc gnetwayservice.Client
}

func NewGnetwayClient(cli client.Client) GnetwayClient {
	return &defaultGnetwayClient{
		cli: cli,
		rpc: gnetwayservice.NewRPCGnetwayClient(cli),
	}
}

// GnetwaySendDataToGateway
// gnetway.sendDataToGateway auth_key_id:long session_id:long payload:bytes = Bool;
func (m *defaultGnetwayClient) GnetwaySendDataToGateway(ctx context.Context, in *gnetway.TLGnetwaySendDataToGateway) (*tg.Bool, error) {
	return m.rpc.GnetwaySendDataToGateway(ctx, in)
}
