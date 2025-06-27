/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package echo1client

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo1/echo1"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo1/echo1/echo1service"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type Echo1Client interface {
	Echo1Echo(ctx context.Context, in *echo1.TLEcho1Echo) (*echo1.Echo, error)
}

type defaultEcho1Client struct {
	cli client.Client
}

func NewEcho1Client(cli client.Client) Echo1Client {
	return &defaultEcho1Client{
		cli: cli,
	}
}

// Echo1Echo
// echo1.echo message:string = Echo;
func (m *defaultEcho1Client) Echo1Echo(ctx context.Context, in *echo1.TLEcho1Echo) (*echo1.Echo, error) {
	cli := echo1service.NewRPCEcho1Client(m.cli)
	return cli.Echo1Echo(ctx, in)
}
