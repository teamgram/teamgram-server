/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package autodownloadclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/autodownload/autodownload/autodownloadservice"

	"github.com/cloudwego/kitex/client"
)

type AutoDownloadClient interface {
	AccountGetAutoDownloadSettings(ctx context.Context, in *tg.TLAccountGetAutoDownloadSettings) (*tg.AccountAutoDownloadSettings, error)
	AccountSaveAutoDownloadSettings(ctx context.Context, in *tg.TLAccountSaveAutoDownloadSettings) (*tg.Bool, error)
}

type defaultAutoDownloadClient struct {
	cli client.Client
}

func NewAutoDownloadClient(cli client.Client) AutoDownloadClient {
	return &defaultAutoDownloadClient{
		cli: cli,
	}
}

// AccountGetAutoDownloadSettings
// account.getAutoDownloadSettings#56da0b3f = account.AutoDownloadSettings;
func (m *defaultAutoDownloadClient) AccountGetAutoDownloadSettings(ctx context.Context, in *tg.TLAccountGetAutoDownloadSettings) (*tg.AccountAutoDownloadSettings, error) {
	cli := autodownloadservice.NewRPCAutoDownloadClient(m.cli)
	return cli.AccountGetAutoDownloadSettings(ctx, in)
}

// AccountSaveAutoDownloadSettings
// account.saveAutoDownloadSettings#76f36233 flags:# low:flags.0?true high:flags.1?true settings:AutoDownloadSettings = Bool;
func (m *defaultAutoDownloadClient) AccountSaveAutoDownloadSettings(ctx context.Context, in *tg.TLAccountSaveAutoDownloadSettings) (*tg.Bool, error) {
	cli := autodownloadservice.NewRPCAutoDownloadClient(m.cli)
	return cli.AccountSaveAutoDownloadSettings(ctx, in)
}
