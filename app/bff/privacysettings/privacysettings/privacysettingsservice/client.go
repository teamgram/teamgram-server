/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package privacysettingsservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AccountGetPrivacy(ctx context.Context, req *tg.TLAccountGetPrivacy, callOptions ...callopt.Option) (r *tg.AccountPrivacyRules, err error)
	AccountSetPrivacy(ctx context.Context, req *tg.TLAccountSetPrivacy, callOptions ...callopt.Option) (r *tg.AccountPrivacyRules, err error)
	AccountGetGlobalPrivacySettings(ctx context.Context, req *tg.TLAccountGetGlobalPrivacySettings, callOptions ...callopt.Option) (r *tg.GlobalPrivacySettings, err error)
	AccountSetGlobalPrivacySettings(ctx context.Context, req *tg.TLAccountSetGlobalPrivacySettings, callOptions ...callopt.Option) (r *tg.GlobalPrivacySettings, err error)
	UsersGetIsPremiumRequiredToContact(ctx context.Context, req *tg.TLUsersGetIsPremiumRequiredToContact, callOptions ...callopt.Option) (r *tg.VectorBool, err error)
	MessagesSetDefaultHistoryTTL(ctx context.Context, req *tg.TLMessagesSetDefaultHistoryTTL, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MessagesGetDefaultHistoryTTL(ctx context.Context, req *tg.TLMessagesGetDefaultHistoryTTL, callOptions ...callopt.Option) (r *tg.DefaultHistoryTTL, err error)
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
	return &kPrivacySettingsClient{
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

type kPrivacySettingsClient struct {
	*kClient
}

func NewRPCPrivacySettingsClient(cli client.Client) Client {
	return &kPrivacySettingsClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kPrivacySettingsClient) AccountGetPrivacy(ctx context.Context, req *tg.TLAccountGetPrivacy, callOptions ...callopt.Option) (r *tg.AccountPrivacyRules, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountGetPrivacy(ctx, req)
}

func (p *kPrivacySettingsClient) AccountSetPrivacy(ctx context.Context, req *tg.TLAccountSetPrivacy, callOptions ...callopt.Option) (r *tg.AccountPrivacyRules, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountSetPrivacy(ctx, req)
}

func (p *kPrivacySettingsClient) AccountGetGlobalPrivacySettings(ctx context.Context, req *tg.TLAccountGetGlobalPrivacySettings, callOptions ...callopt.Option) (r *tg.GlobalPrivacySettings, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountGetGlobalPrivacySettings(ctx, req)
}

func (p *kPrivacySettingsClient) AccountSetGlobalPrivacySettings(ctx context.Context, req *tg.TLAccountSetGlobalPrivacySettings, callOptions ...callopt.Option) (r *tg.GlobalPrivacySettings, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountSetGlobalPrivacySettings(ctx, req)
}

func (p *kPrivacySettingsClient) UsersGetIsPremiumRequiredToContact(ctx context.Context, req *tg.TLUsersGetIsPremiumRequiredToContact, callOptions ...callopt.Option) (r *tg.VectorBool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsersGetIsPremiumRequiredToContact(ctx, req)
}

func (p *kPrivacySettingsClient) MessagesSetDefaultHistoryTTL(ctx context.Context, req *tg.TLMessagesSetDefaultHistoryTTL, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesSetDefaultHistoryTTL(ctx, req)
}

func (p *kPrivacySettingsClient) MessagesGetDefaultHistoryTTL(ctx context.Context, req *tg.TLMessagesGetDefaultHistoryTTL, callOptions ...callopt.Option) (r *tg.DefaultHistoryTTL, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetDefaultHistoryTTL(ctx, req)
}
