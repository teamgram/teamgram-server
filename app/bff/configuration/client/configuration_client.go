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

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/configuration/configuration/configurationservice"

	"github.com/cloudwego/kitex/client"
)

type ConfigurationClient interface {
	HelpGetConfig(ctx context.Context, in *tg.TLHelpGetConfig) (*tg.Config, error)
	HelpGetNearestDc(ctx context.Context, in *tg.TLHelpGetNearestDc) (*tg.NearestDc, error)
	HelpGetAppUpdate(ctx context.Context, in *tg.TLHelpGetAppUpdate) (*tg.HelpAppUpdate, error)
	HelpGetInviteText(ctx context.Context, in *tg.TLHelpGetInviteText) (*tg.HelpInviteText, error)
	HelpGetSupport(ctx context.Context, in *tg.TLHelpGetSupport) (*tg.HelpSupport, error)
	HelpGetAppConfig(ctx context.Context, in *tg.TLHelpGetAppConfig) (*tg.HelpAppConfig, error)
	HelpGetSupportName(ctx context.Context, in *tg.TLHelpGetSupportName) (*tg.HelpSupportName, error)
	HelpDismissSuggestion(ctx context.Context, in *tg.TLHelpDismissSuggestion) (*tg.Bool, error)
	HelpGetCountriesList(ctx context.Context, in *tg.TLHelpGetCountriesList) (*tg.HelpCountriesList, error)
}

type defaultConfigurationClient struct {
	cli client.Client
}

func NewConfigurationClient(cli client.Client) ConfigurationClient {
	return &defaultConfigurationClient{
		cli: cli,
	}
}

// HelpGetConfig
// help.getConfig#c4f9186b = Config;
func (m *defaultConfigurationClient) HelpGetConfig(ctx context.Context, in *tg.TLHelpGetConfig) (*tg.Config, error) {
	cli := configurationservice.NewRPCConfigurationClient(m.cli)
	return cli.HelpGetConfig(ctx, in)
}

// HelpGetNearestDc
// help.getNearestDc#1fb33026 = NearestDc;
func (m *defaultConfigurationClient) HelpGetNearestDc(ctx context.Context, in *tg.TLHelpGetNearestDc) (*tg.NearestDc, error) {
	cli := configurationservice.NewRPCConfigurationClient(m.cli)
	return cli.HelpGetNearestDc(ctx, in)
}

// HelpGetAppUpdate
// help.getAppUpdate#522d5a7d source:string = help.AppUpdate;
func (m *defaultConfigurationClient) HelpGetAppUpdate(ctx context.Context, in *tg.TLHelpGetAppUpdate) (*tg.HelpAppUpdate, error) {
	cli := configurationservice.NewRPCConfigurationClient(m.cli)
	return cli.HelpGetAppUpdate(ctx, in)
}

// HelpGetInviteText
// help.getInviteText#4d392343 = help.InviteText;
func (m *defaultConfigurationClient) HelpGetInviteText(ctx context.Context, in *tg.TLHelpGetInviteText) (*tg.HelpInviteText, error) {
	cli := configurationservice.NewRPCConfigurationClient(m.cli)
	return cli.HelpGetInviteText(ctx, in)
}

// HelpGetSupport
// help.getSupport#9cdf08cd = help.Support;
func (m *defaultConfigurationClient) HelpGetSupport(ctx context.Context, in *tg.TLHelpGetSupport) (*tg.HelpSupport, error) {
	cli := configurationservice.NewRPCConfigurationClient(m.cli)
	return cli.HelpGetSupport(ctx, in)
}

// HelpGetAppConfig
// help.getAppConfig#61e3f854 hash:int = help.AppConfig;
func (m *defaultConfigurationClient) HelpGetAppConfig(ctx context.Context, in *tg.TLHelpGetAppConfig) (*tg.HelpAppConfig, error) {
	cli := configurationservice.NewRPCConfigurationClient(m.cli)
	return cli.HelpGetAppConfig(ctx, in)
}

// HelpGetSupportName
// help.getSupportName#d360e72c = help.SupportName;
func (m *defaultConfigurationClient) HelpGetSupportName(ctx context.Context, in *tg.TLHelpGetSupportName) (*tg.HelpSupportName, error) {
	cli := configurationservice.NewRPCConfigurationClient(m.cli)
	return cli.HelpGetSupportName(ctx, in)
}

// HelpDismissSuggestion
// help.dismissSuggestion#f50dbaa1 peer:InputPeer suggestion:string = Bool;
func (m *defaultConfigurationClient) HelpDismissSuggestion(ctx context.Context, in *tg.TLHelpDismissSuggestion) (*tg.Bool, error) {
	cli := configurationservice.NewRPCConfigurationClient(m.cli)
	return cli.HelpDismissSuggestion(ctx, in)
}

// HelpGetCountriesList
// help.getCountriesList#735787a8 lang_code:string hash:int = help.CountriesList;
func (m *defaultConfigurationClient) HelpGetCountriesList(ctx context.Context, in *tg.TLHelpGetCountriesList) (*tg.HelpCountriesList, error) {
	cli := configurationservice.NewRPCConfigurationClient(m.cli)
	return cli.HelpGetCountriesList(ctx, in)
}
