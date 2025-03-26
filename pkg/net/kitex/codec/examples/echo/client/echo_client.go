/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package echoclient

import (
	"context"

    "github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/examples/echo/echo"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/examples/echo/echo/echoservice"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type EchoClient interface {
    EchoEcho(ctx context.Context, in *echo.TLEchoEcho) (*echo.Echo, error)

}

type defaultEchoClient struct {
	cli client.Client
}

func NewEchoClient(cli client.Client) EchoClient {
	return &defaultEchoClient{
		cli: cli,
	}
}


// EchoEcho
// echo.echo message:string = Echo;
func (m *defaultEchoClient) EchoEcho(ctx context.Context, in *echo.TLEchoEcho) (*echo.Echo, error) {
	 cli := echoservice.NewRPCEchoClient(m.cli)
	 return cli.EchoEcho(ctx, in)
}

