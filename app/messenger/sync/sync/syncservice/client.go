/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package syncservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SyncUpdatesMe(ctx context.Context, req *sync.TLSyncUpdatesMe, callOptions ...callopt.Option) (r *tg.Void, err error)
	SyncUpdatesNotMe(ctx context.Context, req *sync.TLSyncUpdatesNotMe, callOptions ...callopt.Option) (r *tg.Void, err error)
	SyncPushUpdates(ctx context.Context, req *sync.TLSyncPushUpdates, callOptions ...callopt.Option) (r *tg.Void, err error)
	SyncPushUpdatesIfNot(ctx context.Context, req *sync.TLSyncPushUpdatesIfNot, callOptions ...callopt.Option) (r *tg.Void, err error)
	SyncPushBotUpdates(ctx context.Context, req *sync.TLSyncPushBotUpdates, callOptions ...callopt.Option) (r *tg.Void, err error)
	SyncPushRpcResult(ctx context.Context, req *sync.TLSyncPushRpcResult, callOptions ...callopt.Option) (r *tg.Void, err error)
	SyncBroadcastUpdates(ctx context.Context, req *sync.TLSyncBroadcastUpdates, callOptions ...callopt.Option) (r *tg.Void, err error)
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
	return &kSyncClient{
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

type kSyncClient struct {
	*kClient
}

func NewRPCSyncClient(cli client.Client) Client {
	return &kSyncClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kSyncClient) SyncUpdatesMe(ctx context.Context, req *sync.TLSyncUpdatesMe, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SyncUpdatesMe(ctx, req)
}

func (p *kSyncClient) SyncUpdatesNotMe(ctx context.Context, req *sync.TLSyncUpdatesNotMe, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SyncUpdatesNotMe(ctx, req)
}

func (p *kSyncClient) SyncPushUpdates(ctx context.Context, req *sync.TLSyncPushUpdates, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SyncPushUpdates(ctx, req)
}

func (p *kSyncClient) SyncPushUpdatesIfNot(ctx context.Context, req *sync.TLSyncPushUpdatesIfNot, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SyncPushUpdatesIfNot(ctx, req)
}

func (p *kSyncClient) SyncPushBotUpdates(ctx context.Context, req *sync.TLSyncPushBotUpdates, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SyncPushBotUpdates(ctx, req)
}

func (p *kSyncClient) SyncPushRpcResult(ctx context.Context, req *sync.TLSyncPushRpcResult, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SyncPushRpcResult(ctx, req)
}

func (p *kSyncClient) SyncBroadcastUpdates(ctx context.Context, req *sync.TLSyncBroadcastUpdates, callOptions ...callopt.Option) (r *tg.Void, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SyncBroadcastUpdates(ctx, req)
}
