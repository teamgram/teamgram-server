/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package configurationservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	HelpGetConfig(ctx context.Context, req *tg.TLHelpGetConfig, callOptions ...callopt.Option) (r *tg.Config, err error)
	HelpGetNearestDc(ctx context.Context, req *tg.TLHelpGetNearestDc, callOptions ...callopt.Option) (r *tg.NearestDc, err error)
	HelpGetAppUpdate(ctx context.Context, req *tg.TLHelpGetAppUpdate, callOptions ...callopt.Option) (r *tg.HelpAppUpdate, err error)
	HelpGetInviteText(ctx context.Context, req *tg.TLHelpGetInviteText, callOptions ...callopt.Option) (r *tg.HelpInviteText, err error)
	HelpGetSupport(ctx context.Context, req *tg.TLHelpGetSupport, callOptions ...callopt.Option) (r *tg.HelpSupport, err error)
	HelpGetAppConfig(ctx context.Context, req *tg.TLHelpGetAppConfig, callOptions ...callopt.Option) (r *tg.HelpAppConfig, err error)
	HelpGetSupportName(ctx context.Context, req *tg.TLHelpGetSupportName, callOptions ...callopt.Option) (r *tg.HelpSupportName, err error)
	HelpDismissSuggestion(ctx context.Context, req *tg.TLHelpDismissSuggestion, callOptions ...callopt.Option) (r *tg.Bool, err error)
	HelpGetCountriesList(ctx context.Context, req *tg.TLHelpGetCountriesList, callOptions ...callopt.Option) (r *tg.HelpCountriesList, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kConfigurationClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kConfigurationClient struct {
	*kClient
}

func NewRPCConfigurationClient(cli client.Client) Client {
	return &kConfigurationClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kConfigurationClient) HelpGetConfig(ctx context.Context, req *tg.TLHelpGetConfig, callOptions ...callopt.Option) (r *tg.Config, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.HelpGetConfig(ctx, req)
}

func (p *kConfigurationClient) HelpGetNearestDc(ctx context.Context, req *tg.TLHelpGetNearestDc, callOptions ...callopt.Option) (r *tg.NearestDc, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.HelpGetNearestDc(ctx, req)
}

func (p *kConfigurationClient) HelpGetAppUpdate(ctx context.Context, req *tg.TLHelpGetAppUpdate, callOptions ...callopt.Option) (r *tg.HelpAppUpdate, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.HelpGetAppUpdate(ctx, req)
}

func (p *kConfigurationClient) HelpGetInviteText(ctx context.Context, req *tg.TLHelpGetInviteText, callOptions ...callopt.Option) (r *tg.HelpInviteText, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.HelpGetInviteText(ctx, req)
}

func (p *kConfigurationClient) HelpGetSupport(ctx context.Context, req *tg.TLHelpGetSupport, callOptions ...callopt.Option) (r *tg.HelpSupport, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.HelpGetSupport(ctx, req)
}

func (p *kConfigurationClient) HelpGetAppConfig(ctx context.Context, req *tg.TLHelpGetAppConfig, callOptions ...callopt.Option) (r *tg.HelpAppConfig, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.HelpGetAppConfig(ctx, req)
}

func (p *kConfigurationClient) HelpGetSupportName(ctx context.Context, req *tg.TLHelpGetSupportName, callOptions ...callopt.Option) (r *tg.HelpSupportName, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.HelpGetSupportName(ctx, req)
}

func (p *kConfigurationClient) HelpDismissSuggestion(ctx context.Context, req *tg.TLHelpDismissSuggestion, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.HelpDismissSuggestion(ctx, req)
}

func (p *kConfigurationClient) HelpGetCountriesList(ctx context.Context, req *tg.TLHelpGetCountriesList, callOptions ...callopt.Option) (r *tg.HelpCountriesList, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.HelpGetCountriesList(ctx, req)
}
