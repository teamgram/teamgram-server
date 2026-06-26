/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026 The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package webbrowserclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type WebBrowserClient interface {
	AccountGetWebBrowserSettings(ctx context.Context, in *mtproto.TLAccountGetWebBrowserSettings) (*mtproto.Account_WebBrowserSettings, error)
	AccountUpdateWebBrowserSettings(ctx context.Context, in *mtproto.TLAccountUpdateWebBrowserSettings) (*mtproto.Account_WebBrowserSettings, error)
	AccountToggleWebBrowserSettingsException60ED4229(ctx context.Context, in *mtproto.TLAccountToggleWebBrowserSettingsException60ED4229) (*mtproto.Updates, error)
	AccountDeleteWebBrowserSettingsExceptions(ctx context.Context, in *mtproto.TLAccountDeleteWebBrowserSettingsExceptions) (*mtproto.Account_WebBrowserSettings, error)
	AccountToggleWebBrowserSettingsException2D0A0571(ctx context.Context, in *mtproto.TLAccountToggleWebBrowserSettingsException2D0A0571) (*mtproto.Bool, error)
}

type defaultWebBrowserClient struct {
	cli zrpc.Client
}

func NewWebBrowserClient(cli zrpc.Client) WebBrowserClient {
	return &defaultWebBrowserClient{
		cli: cli,
	}
}

// AccountGetWebBrowserSettings
// account.getWebBrowserSettings#56655768 hash:long = account.WebBrowserSettings;
func (m *defaultWebBrowserClient) AccountGetWebBrowserSettings(ctx context.Context, in *mtproto.TLAccountGetWebBrowserSettings) (*mtproto.Account_WebBrowserSettings, error) {
	client := mtproto.NewRPCWebBrowserClient(m.cli.Conn())
	return client.AccountGetWebBrowserSettings(ctx, in)
}

// AccountUpdateWebBrowserSettings
// account.updateWebBrowserSettings#9adf82fe flags:# open_external_browser:flags.0?true display_close_button:flags.1?true = account.WebBrowserSettings;
func (m *defaultWebBrowserClient) AccountUpdateWebBrowserSettings(ctx context.Context, in *mtproto.TLAccountUpdateWebBrowserSettings) (*mtproto.Account_WebBrowserSettings, error) {
	client := mtproto.NewRPCWebBrowserClient(m.cli.Conn())
	return client.AccountUpdateWebBrowserSettings(ctx, in)
}

// AccountToggleWebBrowserSettingsException60ED4229
// account.toggleWebBrowserSettingsException#60ed4229 flags:# delete:flags.1?true open_external_browser:flags.0?Bool url:string = Updates;
func (m *defaultWebBrowserClient) AccountToggleWebBrowserSettingsException60ED4229(ctx context.Context, in *mtproto.TLAccountToggleWebBrowserSettingsException60ED4229) (*mtproto.Updates, error) {
	client := mtproto.NewRPCWebBrowserClient(m.cli.Conn())
	return client.AccountToggleWebBrowserSettingsException60ED4229(ctx, in)
}

// AccountDeleteWebBrowserSettingsExceptions
// account.deleteWebBrowserSettingsExceptions#86a0765d = account.WebBrowserSettings;
func (m *defaultWebBrowserClient) AccountDeleteWebBrowserSettingsExceptions(ctx context.Context, in *mtproto.TLAccountDeleteWebBrowserSettingsExceptions) (*mtproto.Account_WebBrowserSettings, error) {
	client := mtproto.NewRPCWebBrowserClient(m.cli.Conn())
	return client.AccountDeleteWebBrowserSettingsExceptions(ctx, in)
}

// AccountToggleWebBrowserSettingsException2D0A0571
// account.toggleWebBrowserSettingsException#2d0a0571 flags:# delete:flags.1?true open_external_browser:flags.0?Bool url:string = Bool;
func (m *defaultWebBrowserClient) AccountToggleWebBrowserSettingsException2D0A0571(ctx context.Context, in *mtproto.TLAccountToggleWebBrowserSettingsException2D0A0571) (*mtproto.Bool, error) {
	client := mtproto.NewRPCWebBrowserClient(m.cli.Conn())
	return client.AccountToggleWebBrowserSettingsException2D0A0571(ctx, in)
}
