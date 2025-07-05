/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package usernameservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/username/username"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

var _ *tg.Bool

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UsernameGetAccountUsername(ctx context.Context, req *username.TLUsernameGetAccountUsername, callOptions ...callopt.Option) (r *username.UsernameData, err error)
	UsernameCheckAccountUsername(ctx context.Context, req *username.TLUsernameCheckAccountUsername, callOptions ...callopt.Option) (r *username.UsernameExist, err error)
	UsernameGetChannelUsername(ctx context.Context, req *username.TLUsernameGetChannelUsername, callOptions ...callopt.Option) (r *username.UsernameData, err error)
	UsernameCheckChannelUsername(ctx context.Context, req *username.TLUsernameCheckChannelUsername, callOptions ...callopt.Option) (r *username.UsernameExist, err error)
	UsernameUpdateUsernameByPeer(ctx context.Context, req *username.TLUsernameUpdateUsernameByPeer, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UsernameCheckUsername(ctx context.Context, req *username.TLUsernameCheckUsername, callOptions ...callopt.Option) (r *username.UsernameExist, err error)
	UsernameUpdateUsername(ctx context.Context, req *username.TLUsernameUpdateUsername, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UsernameDeleteUsername(ctx context.Context, req *username.TLUsernameDeleteUsername, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UsernameResolveUsername(ctx context.Context, req *username.TLUsernameResolveUsername, callOptions ...callopt.Option) (r *tg.Peer, err error)
	UsernameGetListByUsernameList(ctx context.Context, req *username.TLUsernameGetListByUsernameList, callOptions ...callopt.Option) (r *username.VectorUsernameData, err error)
	UsernameDeleteUsernameByPeer(ctx context.Context, req *username.TLUsernameDeleteUsernameByPeer, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UsernameSearch(ctx context.Context, req *username.TLUsernameSearch, callOptions ...callopt.Option) (r *username.VectorUsernameData, err error)
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
	return &kUsernameClient{
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

type kUsernameClient struct {
	*kClient
}

func NewRPCUsernameClient(cli client.Client) Client {
	return &kUsernameClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kUsernameClient) UsernameGetAccountUsername(ctx context.Context, req *username.TLUsernameGetAccountUsername, callOptions ...callopt.Option) (r *username.UsernameData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsernameGetAccountUsername(ctx, req)
}

func (p *kUsernameClient) UsernameCheckAccountUsername(ctx context.Context, req *username.TLUsernameCheckAccountUsername, callOptions ...callopt.Option) (r *username.UsernameExist, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsernameCheckAccountUsername(ctx, req)
}

func (p *kUsernameClient) UsernameGetChannelUsername(ctx context.Context, req *username.TLUsernameGetChannelUsername, callOptions ...callopt.Option) (r *username.UsernameData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsernameGetChannelUsername(ctx, req)
}

func (p *kUsernameClient) UsernameCheckChannelUsername(ctx context.Context, req *username.TLUsernameCheckChannelUsername, callOptions ...callopt.Option) (r *username.UsernameExist, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsernameCheckChannelUsername(ctx, req)
}

func (p *kUsernameClient) UsernameUpdateUsernameByPeer(ctx context.Context, req *username.TLUsernameUpdateUsernameByPeer, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsernameUpdateUsernameByPeer(ctx, req)
}

func (p *kUsernameClient) UsernameCheckUsername(ctx context.Context, req *username.TLUsernameCheckUsername, callOptions ...callopt.Option) (r *username.UsernameExist, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsernameCheckUsername(ctx, req)
}

func (p *kUsernameClient) UsernameUpdateUsername(ctx context.Context, req *username.TLUsernameUpdateUsername, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsernameUpdateUsername(ctx, req)
}

func (p *kUsernameClient) UsernameDeleteUsername(ctx context.Context, req *username.TLUsernameDeleteUsername, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsernameDeleteUsername(ctx, req)
}

func (p *kUsernameClient) UsernameResolveUsername(ctx context.Context, req *username.TLUsernameResolveUsername, callOptions ...callopt.Option) (r *tg.Peer, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsernameResolveUsername(ctx, req)
}

func (p *kUsernameClient) UsernameGetListByUsernameList(ctx context.Context, req *username.TLUsernameGetListByUsernameList, callOptions ...callopt.Option) (r *username.VectorUsernameData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsernameGetListByUsernameList(ctx, req)
}

func (p *kUsernameClient) UsernameDeleteUsernameByPeer(ctx context.Context, req *username.TLUsernameDeleteUsernameByPeer, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsernameDeleteUsernameByPeer(ctx, req)
}

func (p *kUsernameClient) UsernameSearch(ctx context.Context, req *username.TLUsernameSearch, callOptions ...callopt.Option) (r *username.VectorUsernameData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsernameSearch(ctx, req)
}
