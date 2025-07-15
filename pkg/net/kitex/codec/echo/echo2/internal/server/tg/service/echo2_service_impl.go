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
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo2/echo2"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo2/internal/core"
)

var _ *tg.Bool

// Echo2Echo
// echo2.echo message:string = Echo;
func (s *Service) Echo2Echo(ctx context.Context, request *echo2.TLEcho2Echo) (*echo2.Echo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("echo2.echo - metadata: %s, request: %s", c.MD, request)

	r, err := c.Echo2Echo(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}
