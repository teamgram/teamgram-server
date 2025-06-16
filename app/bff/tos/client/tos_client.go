/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package tosclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/tos/tos/tosservice"

	"github.com/cloudwego/kitex/client"
)

type TosClient interface {
	HelpGetTermsOfServiceUpdate(ctx context.Context, in *tg.TLHelpGetTermsOfServiceUpdate) (*tg.HelpTermsOfServiceUpdate, error)
	HelpAcceptTermsOfService(ctx context.Context, in *tg.TLHelpAcceptTermsOfService) (*tg.Bool, error)
}

type defaultTosClient struct {
	cli client.Client
}

func NewTosClient(cli client.Client) TosClient {
	return &defaultTosClient{
		cli: cli,
	}
}

// HelpGetTermsOfServiceUpdate
// help.getTermsOfServiceUpdate#2ca51fd1 = help.TermsOfServiceUpdate;
func (m *defaultTosClient) HelpGetTermsOfServiceUpdate(ctx context.Context, in *tg.TLHelpGetTermsOfServiceUpdate) (*tg.HelpTermsOfServiceUpdate, error) {
	cli := tosservice.NewRPCTosClient(m.cli)
	return cli.HelpGetTermsOfServiceUpdate(ctx, in)
}

// HelpAcceptTermsOfService
// help.acceptTermsOfService#ee72f79a id:DataJSON = Bool;
func (m *defaultTosClient) HelpAcceptTermsOfService(ctx context.Context, in *tg.TLHelpAcceptTermsOfService) (*tg.Bool, error) {
	cli := tosservice.NewRPCTosClient(m.cli)
	return cli.HelpAcceptTermsOfService(ctx, in)
}
