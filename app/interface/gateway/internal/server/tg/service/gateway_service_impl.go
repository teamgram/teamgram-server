/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/gateway"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/core"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ *tg.Bool

// GatewayPushUpdatesData
// gateway.pushUpdatesData flags:# perm_auth_key_id:long notification:flags.0?true updates:Updates = Bool;
func (s *Service) GatewayPushUpdatesData(ctx context.Context, request *gateway.TLGatewayPushUpdatesData) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("gateway.pushUpdatesData - request: %s", request)

	r, err := c.GatewayPushUpdatesData(request)
	if err != nil {
		c.Logger.Errorf("gateway.pushUpdatesData - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("gateway.pushUpdatesData - reply: %s", r)
	return r, err
}

// GatewayPushSessionUpdatesData
// gateway.pushSessionUpdatesData flags:# perm_auth_key_id:long auth_key_id:long session_id:long updates:Updates = Bool;
func (s *Service) GatewayPushSessionUpdatesData(ctx context.Context, request *gateway.TLGatewayPushSessionUpdatesData) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("gateway.pushSessionUpdatesData - request: %s", request)

	r, err := c.GatewayPushSessionUpdatesData(request)
	if err != nil {
		c.Logger.Errorf("gateway.pushSessionUpdatesData - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("gateway.pushSessionUpdatesData - reply: %s", r)
	return r, err
}

// GatewayPushRpcResultData
// gateway.pushRpcResultData perm_auth_key_id:long auth_key_id:long session_id:long client_req_msg_id:long rpc_result_data:bytes = Bool;
func (s *Service) GatewayPushRpcResultData(ctx context.Context, request *gateway.TLGatewayPushRpcResultData) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("gateway.pushRpcResultData - request: %s", request)

	r, err := c.GatewayPushRpcResultData(request)
	if err != nil {
		c.Logger.Errorf("gateway.pushRpcResultData - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("gateway.pushRpcResultData - reply: %s", r)
	return r, err
}
