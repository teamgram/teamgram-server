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
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo1/echo1"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo1/internal/core"
)

var _ *tg.Bool

// Echo1Echo
// echo1.echo message:string = Echo;
func (s *Service) Echo1Echo(ctx context.Context, request *echo1.TLEcho1Echo) (*echo1.Echo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("echo1.echo - metadata: {}, request: %v", request)

	r, err := c.Echo1Echo(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}
