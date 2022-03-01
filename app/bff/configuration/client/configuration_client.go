/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package configuration_client

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
	HelpGetAppChangelog(ctx context.Context, in *mtproto.TLHelpGetAppChangelog) (*mtproto.Updates, error)
	HelpGetAppConfig(ctx context.Context, in *mtproto.TLHelpGetAppConfig) (*mtproto.JSONValue, error)
	HelpGetSupportName(ctx context.Context, in *mtproto.TLHelpGetSupportName) (*mtproto.Help_SupportName, error)
	HelpDismissSuggestion(ctx context.Context, in *mtproto.TLHelpDismissSuggestion) (*mtproto.Bool, error)
	HelpGetCountriesList(ctx context.Context, in *mtproto.TLHelpGetCountriesList) (*mtproto.Help_CountriesList, error)
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

// HelpGetAppChangelog
// help.getAppChangelog#9010ef6f prev_app_version:string = Updates;
func (m *defaultConfigurationClient) HelpGetAppChangelog(ctx context.Context, in *mtproto.TLHelpGetAppChangelog) (*mtproto.Updates, error) {
	client := mtproto.NewRPCConfigurationClient(m.cli.Conn())
	return client.HelpGetAppChangelog(ctx, in)
}

// HelpGetAppConfig
// help.getAppConfig#98914110 = JSONValue;
func (m *defaultConfigurationClient) HelpGetAppConfig(ctx context.Context, in *mtproto.TLHelpGetAppConfig) (*mtproto.JSONValue, error) {
	client := mtproto.NewRPCConfigurationClient(m.cli.Conn())
	return client.HelpGetAppConfig(ctx, in)
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
