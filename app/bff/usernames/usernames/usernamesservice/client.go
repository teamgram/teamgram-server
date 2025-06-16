/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package usernamesservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AccountCheckUsername(ctx context.Context, req *tg.TLAccountCheckUsername, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountUpdateUsername(ctx context.Context, req *tg.TLAccountUpdateUsername, callOptions ...callopt.Option) (r *tg.User, err error)
	ContactsResolveUsername(ctx context.Context, req *tg.TLContactsResolveUsername, callOptions ...callopt.Option) (r *tg.ContactsResolvedPeer, err error)
	ChannelsCheckUsername(ctx context.Context, req *tg.TLChannelsCheckUsername, callOptions ...callopt.Option) (r *tg.Bool, err error)
	ChannelsUpdateUsername(ctx context.Context, req *tg.TLChannelsUpdateUsername, callOptions ...callopt.Option) (r *tg.Bool, err error)
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
	return &kUsernamesClient{
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

type kUsernamesClient struct {
	*kClient
}

func NewRPCUsernamesClient(cli client.Client) Client {
	return &kUsernamesClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kUsernamesClient) AccountCheckUsername(ctx context.Context, req *tg.TLAccountCheckUsername, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountCheckUsername(ctx, req)
}

func (p *kUsernamesClient) AccountUpdateUsername(ctx context.Context, req *tg.TLAccountUpdateUsername, callOptions ...callopt.Option) (r *tg.User, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountUpdateUsername(ctx, req)
}

func (p *kUsernamesClient) ContactsResolveUsername(ctx context.Context, req *tg.TLContactsResolveUsername, callOptions ...callopt.Option) (r *tg.ContactsResolvedPeer, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsResolveUsername(ctx, req)
}

func (p *kUsernamesClient) ChannelsCheckUsername(ctx context.Context, req *tg.TLChannelsCheckUsername, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChannelsCheckUsername(ctx, req)
}

func (p *kUsernamesClient) ChannelsUpdateUsername(ctx context.Context, req *tg.TLChannelsUpdateUsername, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChannelsUpdateUsername(ctx, req)
}
