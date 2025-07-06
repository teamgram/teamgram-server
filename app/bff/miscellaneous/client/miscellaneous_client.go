/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package miscellaneousclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/miscellaneous/miscellaneous/miscellaneousservice"

	"github.com/cloudwego/kitex/client"
)

type MiscellaneousClient interface {
	HelpSaveAppLog(ctx context.Context, in *tg.TLHelpSaveAppLog) (*tg.Bool, error)
	HelpTest(ctx context.Context, in *tg.TLHelpTest) (*tg.Bool, error)
}

type defaultMiscellaneousClient struct {
	cli client.Client
}

func NewMiscellaneousClient(cli client.Client) MiscellaneousClient {
	return &defaultMiscellaneousClient{
		cli: cli,
	}
}

// HelpSaveAppLog
// help.saveAppLog#6f02f748 events:Vector<InputAppEvent> = Bool;
func (m *defaultMiscellaneousClient) HelpSaveAppLog(ctx context.Context, in *tg.TLHelpSaveAppLog) (*tg.Bool, error) {
	cli := miscellaneousservice.NewRPCMiscellaneousClient(m.cli)
	return cli.HelpSaveAppLog(ctx, in)
}

// HelpTest
// help.test#c0e202f7 = Bool;
func (m *defaultMiscellaneousClient) HelpTest(ctx context.Context, in *tg.TLHelpTest) (*tg.Bool, error) {
	cli := miscellaneousservice.NewRPCMiscellaneousClient(m.cli)
	return cli.HelpTest(ctx, in)
}
