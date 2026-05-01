/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package gatewayclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/gateway"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/gateway/gatewayservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type GatewayClient interface {
	GatewayPushUpdatesData(ctx context.Context, in *gateway.TLGatewayPushUpdatesData) (*tg.Bool, error)
	GatewayPushSessionUpdatesData(ctx context.Context, in *gateway.TLGatewayPushSessionUpdatesData) (*tg.Bool, error)
	GatewayPushRpcResultData(ctx context.Context, in *gateway.TLGatewayPushRpcResultData) (*tg.Bool, error)
}

type defaultGatewayClient struct {
	cli client.Client
	rpc gatewayservice.Client
}

func NewGatewayClient(cli client.Client) GatewayClient {
	return &defaultGatewayClient{
		cli: cli,
		rpc: gatewayservice.NewRPCGatewayClient(cli),
	}
}

// GatewayPushUpdatesData
// gateway.pushUpdatesData flags:# perm_auth_key_id:long notification:flags.0?true updates:Updates = Bool;
func (m *defaultGatewayClient) GatewayPushUpdatesData(ctx context.Context, in *gateway.TLGatewayPushUpdatesData) (*tg.Bool, error) {
	return m.rpc.GatewayPushUpdatesData(ctx, in)
}

// GatewayPushSessionUpdatesData
// gateway.pushSessionUpdatesData flags:# perm_auth_key_id:long auth_key_id:long session_id:long updates:Updates = Bool;
func (m *defaultGatewayClient) GatewayPushSessionUpdatesData(ctx context.Context, in *gateway.TLGatewayPushSessionUpdatesData) (*tg.Bool, error) {
	return m.rpc.GatewayPushSessionUpdatesData(ctx, in)
}

// GatewayPushRpcResultData
// gateway.pushRpcResultData perm_auth_key_id:long auth_key_id:long session_id:long client_req_msg_id:long rpc_result_data:bytes = Bool;
func (m *defaultGatewayClient) GatewayPushRpcResultData(ctx context.Context, in *gateway.TLGatewayPushRpcResultData) (*tg.Bool, error) {
	return m.rpc.GatewayPushRpcResultData(ctx, in)
}
