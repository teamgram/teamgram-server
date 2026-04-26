/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package miscellaneousclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/bff/miscellaneous/miscellaneous/miscellaneousservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

type MiscellaneousClient interface {
	HelpSaveAppLog(ctx context.Context, in *tg.TLHelpSaveAppLog) (*tg.Bool, error)
	HelpTest(ctx context.Context, in *tg.TLHelpTest) (*tg.Bool, error)
	Close() error
}

type defaultMiscellaneousClient struct {
	cli client.Client
	rpc miscellaneousservice.Client
}

func NewMiscellaneousClient(cli client.Client) MiscellaneousClient {
	return &defaultMiscellaneousClient{
		cli: cli,
		rpc: miscellaneousservice.NewRPCMiscellaneousClient(cli),
	}
}

func (m *defaultMiscellaneousClient) Close() error {
	if closer, ok := any(m.cli).(interface{ Close() error }); ok {
		return closer.Close()
	}
	return nil
}

// HelpSaveAppLog
// help.saveAppLog#6f02f748 events:Vector<InputAppEvent> = Bool;
func (m *defaultMiscellaneousClient) HelpSaveAppLog(ctx context.Context, in *tg.TLHelpSaveAppLog) (*tg.Bool, error) {
	return m.rpc.HelpSaveAppLog(ctx, in)
}

// HelpTest
// help.test#c0e202f7 = Bool;
func (m *defaultMiscellaneousClient) HelpTest(ctx context.Context, in *tg.TLHelpTest) (*tg.Bool, error) {
	return m.rpc.HelpTest(ctx, in)
}
