/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package nsfwclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/nsfw/nsfw/nsfwservice"

	"github.com/cloudwego/kitex/client"
)

type NsfwClient interface {
	AccountSetContentSettings(ctx context.Context, in *tg.TLAccountSetContentSettings) (*tg.Bool, error)
	AccountGetContentSettings(ctx context.Context, in *tg.TLAccountGetContentSettings) (*tg.AccountContentSettings, error)
}

type defaultNsfwClient struct {
	cli client.Client
}

func NewNsfwClient(cli client.Client) NsfwClient {
	return &defaultNsfwClient{
		cli: cli,
	}
}

// AccountSetContentSettings
// account.setContentSettings#b574b16b flags:# sensitive_enabled:flags.0?true = Bool;
func (m *defaultNsfwClient) AccountSetContentSettings(ctx context.Context, in *tg.TLAccountSetContentSettings) (*tg.Bool, error) {
	cli := nsfwservice.NewRPCNsfwClient(m.cli)
	return cli.AccountSetContentSettings(ctx, in)
}

// AccountGetContentSettings
// account.getContentSettings#8b9b4dae = account.ContentSettings;
func (m *defaultNsfwClient) AccountGetContentSettings(ctx context.Context, in *tg.TLAccountGetContentSettings) (*tg.AccountContentSettings, error) {
	cli := nsfwservice.NewRPCNsfwClient(m.cli)
	return cli.AccountGetContentSettings(ctx, in)
}
