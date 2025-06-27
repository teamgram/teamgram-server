/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package echo2client

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo2/echo2"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo2/echo2/echo2service"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type Echo2Client interface {
	Echo2Echo(ctx context.Context, in *echo2.TLEcho2Echo) (*echo2.Echo, error)
}

type defaultEcho2Client struct {
	cli client.Client
}

func NewEcho2Client(cli client.Client) Echo2Client {
	return &defaultEcho2Client{
		cli: cli,
	}
}

// Echo2Echo
// echo2.echo message:string = Echo;
func (m *defaultEcho2Client) Echo2Echo(ctx context.Context, in *echo2.TLEcho2Echo) (*echo2.Echo, error) {
	cli := echo2service.NewRPCEcho2Client(m.cli)
	return cli.Echo2Echo(ctx, in)
}
