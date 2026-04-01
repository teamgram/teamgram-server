/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2026 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/examples/echo/echo"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/examples/echo/internal/core"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ *tg.Bool

// EchoEcho
// echo.echo message:string = Echo;
func (s *Service) EchoEcho(ctx context.Context, request *echo.TLEchoEcho) (*echo.Echo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("echo.echo - request: %s", request)

	r, err := c.EchoEcho(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echo.echo - reply: %s", r)
	return r, err
}
