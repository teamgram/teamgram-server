/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package configurationclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type ConfigurationClient interface {
	HelpGetConfig(ctx context.Context, in *mtproto.TLHelpGetConfig) (*mtproto.Config, error)
	HelpGetNearestDc(ctx context.Context, in *mtproto.TLHelpGetNearestDc) (*mtproto.NearestDc, error)
	HelpGetAppUpdate(ctx context.Context, in *mtproto.TLHelpGetAppUpdate) (*mtproto.Help_AppUpdate, error)
	HelpGetInviteText(ctx context.Context, in *mtproto.TLHelpGetInviteText) (*mtproto.Help_InviteText, error)
	HelpGetSupport(ctx context.Context, in *mtproto.TLHelpGetSupport) (*mtproto.Help_Support, error)
	HelpGetAppConfig61E3F854(ctx context.Context, in *mtproto.TLHelpGetAppConfig61E3F854) (*mtproto.Help_AppConfig, error)
	HelpGetSupportName(ctx context.Context, in *mtproto.TLHelpGetSupportName) (*mtproto.Help_SupportName, error)
	HelpDismissSuggestion(ctx context.Context, in *mtproto.TLHelpDismissSuggestion) (*mtproto.Bool, error)
	HelpGetCountriesList(ctx context.Context, in *mtproto.TLHelpGetCountriesList) (*mtproto.Help_CountriesList, error)
	HelpGetAppChangelog(ctx context.Context, in *mtproto.TLHelpGetAppChangelog) (*mtproto.Updates, error)
	HelpGetAppConfig98914110(ctx context.Context, in *mtproto.TLHelpGetAppConfig98914110) (*mtproto.JSONValue, error)
}

type defaultConfigurationClient struct {
	cli zrpc.Client
}

func NewConfigurationClient(cli zrpc.Client) ConfigurationClient {
	return &defaultConfigurationClient{
		cli: cli,
	}
}

// HelpGetConfig
// help.getConfig#c4f9186b = Config;
func (m *defaultConfigurationClient) HelpGetConfig(ctx context.Context, in *mtproto.TLHelpGetConfig) (*mtproto.Config, error) {
	client := mtproto.NewRPCConfigurationClient(m.cli.Conn())
	return client.HelpGetConfig(ctx, in)
}

// HelpGetNearestDc
// help.getNearestDc#1fb33026 = NearestDc;
func (m *defaultConfigurationClient) HelpGetNearestDc(ctx context.Context, in *mtproto.TLHelpGetNearestDc) (*mtproto.NearestDc, error) {
	client := mtproto.NewRPCConfigurationClient(m.cli.Conn())
	return client.HelpGetNearestDc(ctx, in)
}

// HelpGetAppUpdate
// help.getAppUpdate#522d5a7d source:string = help.AppUpdate;
func (m *defaultConfigurationClient) HelpGetAppUpdate(ctx context.Context, in *mtproto.TLHelpGetAppUpdate) (*mtproto.Help_AppUpdate, error) {
	client := mtproto.NewRPCConfigurationClient(m.cli.Conn())
	return client.HelpGetAppUpdate(ctx, in)
}

// HelpGetInviteText
// help.getInviteText#4d392343 = help.InviteText;
func (m *defaultConfigurationClient) HelpGetInviteText(ctx context.Context, in *mtproto.TLHelpGetInviteText) (*mtproto.Help_InviteText, error) {
	client := mtproto.NewRPCConfigurationClient(m.cli.Conn())
	return client.HelpGetInviteText(ctx, in)
}

// HelpGetSupport
// help.getSupport#9cdf08cd = help.Support;
func (m *defaultConfigurationClient) HelpGetSupport(ctx context.Context, in *mtproto.TLHelpGetSupport) (*mtproto.Help_Support, error) {
	client := mtproto.NewRPCConfigurationClient(m.cli.Conn())
	return client.HelpGetSupport(ctx, in)
}

// HelpGetAppConfig61E3F854
// help.getAppConfig#61e3f854 hash:int = help.AppConfig;
func (m *defaultConfigurationClient) HelpGetAppConfig61E3F854(ctx context.Context, in *mtproto.TLHelpGetAppConfig61E3F854) (*mtproto.Help_AppConfig, error) {
	client := mtproto.NewRPCConfigurationClient(m.cli.Conn())
	return client.HelpGetAppConfig61E3F854(ctx, in)
}

// HelpGetSupportName
// help.getSupportName#d360e72c = help.SupportName;
func (m *defaultConfigurationClient) HelpGetSupportName(ctx context.Context, in *mtproto.TLHelpGetSupportName) (*mtproto.Help_SupportName, error) {
	client := mtproto.NewRPCConfigurationClient(m.cli.Conn())
	return client.HelpGetSupportName(ctx, in)
}

// HelpDismissSuggestion
// help.dismissSuggestion#f50dbaa1 peer:InputPeer suggestion:string = Bool;
func (m *defaultConfigurationClient) HelpDismissSuggestion(ctx context.Context, in *mtproto.TLHelpDismissSuggestion) (*mtproto.Bool, error) {
	client := mtproto.NewRPCConfigurationClient(m.cli.Conn())
	return client.HelpDismissSuggestion(ctx, in)
}

// HelpGetCountriesList
// help.getCountriesList#735787a8 lang_code:string hash:int = help.CountriesList;
func (m *defaultConfigurationClient) HelpGetCountriesList(ctx context.Context, in *mtproto.TLHelpGetCountriesList) (*mtproto.Help_CountriesList, error) {
	client := mtproto.NewRPCConfigurationClient(m.cli.Conn())
	return client.HelpGetCountriesList(ctx, in)
}

// HelpGetAppChangelog
// help.getAppChangelog#9010ef6f prev_app_version:string = Updates;
func (m *defaultConfigurationClient) HelpGetAppChangelog(ctx context.Context, in *mtproto.TLHelpGetAppChangelog) (*mtproto.Updates, error) {
	client := mtproto.NewRPCConfigurationClient(m.cli.Conn())
	return client.HelpGetAppChangelog(ctx, in)
}

// HelpGetAppConfig98914110
// help.getAppConfig#98914110 = JSONValue;
func (m *defaultConfigurationClient) HelpGetAppConfig98914110(ctx context.Context, in *mtproto.TLHelpGetAppConfig98914110) (*mtproto.JSONValue, error) {
	client := mtproto.NewRPCConfigurationClient(m.cli.Conn())
	return client.HelpGetAppConfig98914110(ctx, in)
}
