/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2026 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package tosclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/bff/tos/tos/tosservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

type TosClient interface {
	HelpGetTermsOfServiceUpdate(ctx context.Context, in *tg.TLHelpGetTermsOfServiceUpdate) (*tg.HelpTermsOfServiceUpdate, error)
	HelpAcceptTermsOfService(ctx context.Context, in *tg.TLHelpAcceptTermsOfService) (*tg.Bool, error)
	Close() error
}

type defaultTosClient struct {
	cli client.Client
}

func NewTosClient(cli client.Client) TosClient {
	return &defaultTosClient{
		cli: cli,
	}
}

func (m *defaultTosClient) Close() error {
	if closer, ok := any(m.cli).(interface{ Close() error }); ok {
		return closer.Close()
	}
	return nil
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
