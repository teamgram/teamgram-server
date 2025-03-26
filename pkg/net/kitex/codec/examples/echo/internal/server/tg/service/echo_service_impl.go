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
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/examples/echo/echo"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/examples/echo/internal/core"
)

var _ *tg.Bool

// EchoEcho
// echo.echo message:string = Echo;
func (s *Service) EchoEcho(ctx context.Context, request *echo.TLEchoEcho) (*echo.Echo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("echo.echo - metadata: {}, request: %v", request)

	r, err := c.EchoEcho(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %s", r)
	return r, err
}
