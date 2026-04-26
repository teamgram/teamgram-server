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

	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/gnetway"
	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/internal/core"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ *tg.Bool

// GnetwaySendDataToGateway
// gnetway.sendDataToGateway auth_key_id:long session_id:long payload:bytes = Bool;
func (s *Service) GnetwaySendDataToGateway(ctx context.Context, request *gnetway.TLGnetwaySendDataToGateway) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("gnetway.sendDataToGateway - request: %s", request)

	r, err := c.GnetwaySendDataToGateway(request)
	if err != nil {
		c.Logger.Errorf("gnetway.sendDataToGateway - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("gnetway.sendDataToGateway - reply: %s", r)
	return r, err
}
