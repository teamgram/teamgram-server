/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package nsfw_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type NsfwClient interface {
	AccountSetContentSettings(ctx context.Context, in *mtproto.TLAccountSetContentSettings) (*mtproto.Bool, error)
	AccountGetContentSettings(ctx context.Context, in *mtproto.TLAccountGetContentSettings) (*mtproto.Account_ContentSettings, error)
}

type defaultNsfwClient struct {
	cli zrpc.Client
}

func NewNsfwClient(cli zrpc.Client) NsfwClient {
	return &defaultNsfwClient{
		cli: cli,
	}
}

// AccountSetContentSettings
// account.setContentSettings#b574b16b flags:# sensitive_enabled:flags.0?true = Bool;
func (m *defaultNsfwClient) AccountSetContentSettings(ctx context.Context, in *mtproto.TLAccountSetContentSettings) (*mtproto.Bool, error) {
	client := mtproto.NewRPCNsfwClient(m.cli.Conn())
	return client.AccountSetContentSettings(ctx, in)
}

// AccountGetContentSettings
// account.getContentSettings#8b9b4dae = account.ContentSettings;
func (m *defaultNsfwClient) AccountGetContentSettings(ctx context.Context, in *mtproto.TLAccountGetContentSettings) (*mtproto.Account_ContentSettings, error) {
	client := mtproto.NewRPCNsfwClient(m.cli.Conn())
	return client.AccountGetContentSettings(ctx, in)
}
