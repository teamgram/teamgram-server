/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package autodownload_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type AutoDownloadClient interface {
	AccountGetAutoDownloadSettings(ctx context.Context, in *mtproto.TLAccountGetAutoDownloadSettings) (*mtproto.Account_AutoDownloadSettings, error)
	AccountSaveAutoDownloadSettings(ctx context.Context, in *mtproto.TLAccountSaveAutoDownloadSettings) (*mtproto.Bool, error)
}

type defaultAutoDownloadClient struct {
	cli zrpc.Client
}

func NewAutoDownloadClient(cli zrpc.Client) AutoDownloadClient {
	return &defaultAutoDownloadClient{
		cli: cli,
	}
}

// AccountGetAutoDownloadSettings
// account.getAutoDownloadSettings#56da0b3f = account.AutoDownloadSettings;
func (m *defaultAutoDownloadClient) AccountGetAutoDownloadSettings(ctx context.Context, in *mtproto.TLAccountGetAutoDownloadSettings) (*mtproto.Account_AutoDownloadSettings, error) {
	client := mtproto.NewRPCAutoDownloadClient(m.cli.Conn())
	return client.AccountGetAutoDownloadSettings(ctx, in)
}

// AccountSaveAutoDownloadSettings
// account.saveAutoDownloadSettings#76f36233 flags:# low:flags.0?true high:flags.1?true settings:AutoDownloadSettings = Bool;
func (m *defaultAutoDownloadClient) AccountSaveAutoDownloadSettings(ctx context.Context, in *mtproto.TLAccountSaveAutoDownloadSettings) (*mtproto.Bool, error) {
	client := mtproto.NewRPCAutoDownloadClient(m.cli.Conn())
	return client.AccountSaveAutoDownloadSettings(ctx, in)
}
