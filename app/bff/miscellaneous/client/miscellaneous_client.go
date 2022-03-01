/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package miscellaneous_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type MiscellaneousClient interface {
	HelpSaveAppLog(ctx context.Context, in *mtproto.TLHelpSaveAppLog) (*mtproto.Bool, error)
	HelpTest(ctx context.Context, in *mtproto.TLHelpTest) (*mtproto.Bool, error)
}

type defaultMiscellaneousClient struct {
	cli zrpc.Client
}

func NewMiscellaneousClient(cli zrpc.Client) MiscellaneousClient {
	return &defaultMiscellaneousClient{
		cli: cli,
	}
}

// HelpSaveAppLog
// help.saveAppLog#6f02f748 events:Vector<InputAppEvent> = Bool;
func (m *defaultMiscellaneousClient) HelpSaveAppLog(ctx context.Context, in *mtproto.TLHelpSaveAppLog) (*mtproto.Bool, error) {
	client := mtproto.NewRPCMiscellaneousClient(m.cli.Conn())
	return client.HelpSaveAppLog(ctx, in)
}

// HelpTest
// help.test#c0e202f7 = Bool;
func (m *defaultMiscellaneousClient) HelpTest(ctx context.Context, in *mtproto.TLHelpTest) (*mtproto.Bool, error) {
	client := mtproto.NewRPCMiscellaneousClient(m.cli.Conn())
	return client.HelpTest(ctx, in)
}
