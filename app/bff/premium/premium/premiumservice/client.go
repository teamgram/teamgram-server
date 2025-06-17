/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package premiumservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	HelpGetPremiumPromo(ctx context.Context, req *tg.TLHelpGetPremiumPromo, callOptions ...callopt.Option) (r *tg.HelpPremiumPromo, err error)
	PaymentsAssignAppStoreTransaction(ctx context.Context, req *tg.TLPaymentsAssignAppStoreTransaction, callOptions ...callopt.Option) (r *tg.Updates, err error)
	PaymentsAssignPlayMarketTransaction(ctx context.Context, req *tg.TLPaymentsAssignPlayMarketTransaction, callOptions ...callopt.Option) (r *tg.Updates, err error)
	PaymentsCanPurchaseStore(ctx context.Context, req *tg.TLPaymentsCanPurchaseStore, callOptions ...callopt.Option) (r *tg.Bool, err error)
	PaymentsCanPurchasePremium(ctx context.Context, req *tg.TLPaymentsCanPurchasePremium, callOptions ...callopt.Option) (r *tg.Bool, err error)
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
	return &kPremiumClient{
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

type kPremiumClient struct {
	*kClient
}

func NewRPCPremiumClient(cli client.Client) Client {
	return &kPremiumClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kPremiumClient) HelpGetPremiumPromo(ctx context.Context, req *tg.TLHelpGetPremiumPromo, callOptions ...callopt.Option) (r *tg.HelpPremiumPromo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.HelpGetPremiumPromo(ctx, req)
}

func (p *kPremiumClient) PaymentsAssignAppStoreTransaction(ctx context.Context, req *tg.TLPaymentsAssignAppStoreTransaction, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PaymentsAssignAppStoreTransaction(ctx, req)
}

func (p *kPremiumClient) PaymentsAssignPlayMarketTransaction(ctx context.Context, req *tg.TLPaymentsAssignPlayMarketTransaction, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PaymentsAssignPlayMarketTransaction(ctx, req)
}

func (p *kPremiumClient) PaymentsCanPurchaseStore(ctx context.Context, req *tg.TLPaymentsCanPurchaseStore, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PaymentsCanPurchaseStore(ctx, req)
}

func (p *kPremiumClient) PaymentsCanPurchasePremium(ctx context.Context, req *tg.TLPaymentsCanPurchasePremium, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PaymentsCanPurchasePremium(ctx, req)
}
