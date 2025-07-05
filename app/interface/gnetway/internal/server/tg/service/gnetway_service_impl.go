/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/gnetway"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *tg.Bool

// GnetwaySendDataToGateway
// gnetway.sendDataToGateway auth_key_id:long session_id:long payload:bytes = Bool;
func (s *Service) GnetwaySendDataToGateway(ctx context.Context, request *gnetway.TLGnetwaySendDataToGateway) (*tg.Bool, error) {
	// c := core.New(ctx, s.svcCtx)
	logx.WithContext(ctx).Debugf("gnetway.sendDataToGateway - metadata: {}, request: %s", request)

	r, err := s.RPCGnetway.GnetwaySendDataToGateway(ctx, request)
	if err != nil {
		return nil, err
	}

	logx.WithContext(ctx).Debugf("gnetway.sendDataToGateway - reply: %s", r)
	return r, err
}
