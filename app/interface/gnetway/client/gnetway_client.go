/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2026 Teamgram Authors.
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
	Close() error
}

type defaultGnetwayClient struct {
	cli client.Client
}

func NewGnetwayClient(cli client.Client) GnetwayClient {
	return &defaultGnetwayClient{
		cli: cli,
	}
}

func (m *defaultGnetwayClient) Close() error {
	if closer, ok := any(m.cli).(interface{ Close() error }); ok {
		return closer.Close()
	}
	return nil
}

// GnetwaySendDataToGateway
// gnetway.sendDataToGateway auth_key_id:long session_id:long payload:bytes = Bool;
func (m *defaultGnetwayClient) GnetwaySendDataToGateway(ctx context.Context, in *gnetway.TLGnetwaySendDataToGateway) (*tg.Bool, error) {
	cli := gnetwayservice.NewRPCGnetwayClient(m.cli)
	return cli.GnetwaySendDataToGateway(ctx, in)
}
