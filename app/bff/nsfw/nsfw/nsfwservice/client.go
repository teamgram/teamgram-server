/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package nsfwservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AccountSetContentSettings(ctx context.Context, req *tg.TLAccountSetContentSettings, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountGetContentSettings(ctx context.Context, req *tg.TLAccountGetContentSettings, callOptions ...callopt.Option) (r *tg.AccountContentSettings, err error)
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
	return &kNsfwClient{
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

type kNsfwClient struct {
	*kClient
}

func NewRPCNsfwClient(cli client.Client) Client {
	return &kNsfwClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kNsfwClient) AccountSetContentSettings(ctx context.Context, req *tg.TLAccountSetContentSettings, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountSetContentSettings(ctx, req)
}

func (p *kNsfwClient) AccountGetContentSettings(ctx context.Context, req *tg.TLAccountGetContentSettings, callOptions ...callopt.Option) (r *tg.AccountContentSettings, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountGetContentSettings(ctx, req)
}
