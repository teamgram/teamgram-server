/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package presenceservice

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/presence/presence"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

var _ *tg.Bool

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	PresenceSetSessionOnline(ctx context.Context, req *presence.TLPresenceSetSessionOnline, callOptions ...callopt.Option) (r *tg.Bool, err error)
	PresenceSetSessionOffline(ctx context.Context, req *presence.TLPresenceSetSessionOffline, callOptions ...callopt.Option) (r *tg.Bool, err error)
	PresenceGetUserOnlineSessions(ctx context.Context, req *presence.TLPresenceGetUserOnlineSessions, callOptions ...callopt.Option) (r *presence.UserOnlineSessions, err error)
	PresenceGetUsersOnlineSessions(ctx context.Context, req *presence.TLPresenceGetUsersOnlineSessions, callOptions ...callopt.Option) (r *presence.VectorUserOnlineSessions, err error)
	PresenceGetGatewaySessions(ctx context.Context, req *presence.TLPresenceGetGatewaySessions, callOptions ...callopt.Option) (r *presence.VectorOnlineSession, err error)
}

// Deprecated: prefer the generated app client helper or pkg/net/kitex.NewClient for TL-aware transport setup.
// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))
	options = append(options, client.WithCodec(codec.NewZRpcCodec(false)))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kPresenceClient{
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

type kPresenceClient struct {
	*kClient
}

func NewRPCPresenceClient(cli client.Client) Client {
	return &kPresenceClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kPresenceClient) PresenceSetSessionOnline(ctx context.Context, req *presence.TLPresenceSetSessionOnline, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PresenceSetSessionOnline(ctx, req)
}

func (p *kPresenceClient) PresenceSetSessionOffline(ctx context.Context, req *presence.TLPresenceSetSessionOffline, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PresenceSetSessionOffline(ctx, req)
}

func (p *kPresenceClient) PresenceGetUserOnlineSessions(ctx context.Context, req *presence.TLPresenceGetUserOnlineSessions, callOptions ...callopt.Option) (r *presence.UserOnlineSessions, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PresenceGetUserOnlineSessions(ctx, req)
}

func (p *kPresenceClient) PresenceGetUsersOnlineSessions(ctx context.Context, req *presence.TLPresenceGetUsersOnlineSessions, callOptions ...callopt.Option) (r *presence.VectorUserOnlineSessions, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PresenceGetUsersOnlineSessions(ctx, req)
}

func (p *kPresenceClient) PresenceGetGatewaySessions(ctx context.Context, req *presence.TLPresenceGetGatewaySessions, callOptions ...callopt.Option) (r *presence.VectorOnlineSession, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PresenceGetGatewaySessions(ctx, req)
}
