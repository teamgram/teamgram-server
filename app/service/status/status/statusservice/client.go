/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package statusservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/status/status"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

var _ *tg.Bool

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	StatusSetSessionOnline(ctx context.Context, req *status.TLStatusSetSessionOnline, callOptions ...callopt.Option) (r *tg.Bool, err error)
	StatusSetSessionOffline(ctx context.Context, req *status.TLStatusSetSessionOffline, callOptions ...callopt.Option) (r *tg.Bool, err error)
	StatusGetUserOnlineSessions(ctx context.Context, req *status.TLStatusGetUserOnlineSessions, callOptions ...callopt.Option) (r *status.UserSessionEntryList, err error)
	StatusGetUsersOnlineSessionsList(ctx context.Context, req *status.TLStatusGetUsersOnlineSessionsList, callOptions ...callopt.Option) (r *status.VectorUserSessionEntryList, err error)
	StatusGetChannelOnlineUsers(ctx context.Context, req *status.TLStatusGetChannelOnlineUsers, callOptions ...callopt.Option) (r *status.VectorLong, err error)
	StatusSetUserChannelsOnline(ctx context.Context, req *status.TLStatusSetUserChannelsOnline, callOptions ...callopt.Option) (r *tg.Bool, err error)
	StatusSetUserChannelsOffline(ctx context.Context, req *status.TLStatusSetUserChannelsOffline, callOptions ...callopt.Option) (r *tg.Bool, err error)
	StatusSetChannelUserOffline(ctx context.Context, req *status.TLStatusSetChannelUserOffline, callOptions ...callopt.Option) (r *tg.Bool, err error)
	StatusSetChannelUsersOnline(ctx context.Context, req *status.TLStatusSetChannelUsersOnline, callOptions ...callopt.Option) (r *tg.Bool, err error)
	StatusSetChannelOffline(ctx context.Context, req *status.TLStatusSetChannelOffline, callOptions ...callopt.Option) (r *tg.Bool, err error)
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
	return &kStatusClient{
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

type kStatusClient struct {
	*kClient
}

func NewRPCStatusClient(cli client.Client) Client {
	return &kStatusClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kStatusClient) StatusSetSessionOnline(ctx context.Context, req *status.TLStatusSetSessionOnline, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StatusSetSessionOnline(ctx, req)
}

func (p *kStatusClient) StatusSetSessionOffline(ctx context.Context, req *status.TLStatusSetSessionOffline, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StatusSetSessionOffline(ctx, req)
}

func (p *kStatusClient) StatusGetUserOnlineSessions(ctx context.Context, req *status.TLStatusGetUserOnlineSessions, callOptions ...callopt.Option) (r *status.UserSessionEntryList, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StatusGetUserOnlineSessions(ctx, req)
}

func (p *kStatusClient) StatusGetUsersOnlineSessionsList(ctx context.Context, req *status.TLStatusGetUsersOnlineSessionsList, callOptions ...callopt.Option) (r *status.VectorUserSessionEntryList, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StatusGetUsersOnlineSessionsList(ctx, req)
}

func (p *kStatusClient) StatusGetChannelOnlineUsers(ctx context.Context, req *status.TLStatusGetChannelOnlineUsers, callOptions ...callopt.Option) (r *status.VectorLong, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StatusGetChannelOnlineUsers(ctx, req)
}

func (p *kStatusClient) StatusSetUserChannelsOnline(ctx context.Context, req *status.TLStatusSetUserChannelsOnline, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StatusSetUserChannelsOnline(ctx, req)
}

func (p *kStatusClient) StatusSetUserChannelsOffline(ctx context.Context, req *status.TLStatusSetUserChannelsOffline, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StatusSetUserChannelsOffline(ctx, req)
}

func (p *kStatusClient) StatusSetChannelUserOffline(ctx context.Context, req *status.TLStatusSetChannelUserOffline, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StatusSetChannelUserOffline(ctx, req)
}

func (p *kStatusClient) StatusSetChannelUsersOnline(ctx context.Context, req *status.TLStatusSetChannelUsersOnline, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StatusSetChannelUsersOnline(ctx, req)
}

func (p *kStatusClient) StatusSetChannelOffline(ctx context.Context, req *status.TLStatusSetChannelOffline, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StatusSetChannelOffline(ctx, req)
}
