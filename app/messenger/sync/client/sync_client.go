/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package sync_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type SyncClient interface {
	SyncUpdatesMe(ctx context.Context, in *sync.TLSyncUpdatesMe) (*mtproto.Void, error)
	SyncUpdatesNotMe(ctx context.Context, in *sync.TLSyncUpdatesNotMe) (*mtproto.Void, error)
	SyncPushUpdates(ctx context.Context, in *sync.TLSyncPushUpdates) (*mtproto.Void, error)
	SyncPushUpdatesIfNot(ctx context.Context, in *sync.TLSyncPushUpdatesIfNot) (*mtproto.Void, error)
	SyncPushBotUpdates(ctx context.Context, in *sync.TLSyncPushBotUpdates) (*mtproto.Void, error)
	SyncPushRpcResult(ctx context.Context, in *sync.TLSyncPushRpcResult) (*mtproto.Void, error)
	SyncBroadcastUpdates(ctx context.Context, in *sync.TLSyncBroadcastUpdates) (*mtproto.Void, error)
}

type defaultSyncClient struct {
	cli zrpc.Client
}

func NewSyncClient(cli zrpc.Client) SyncClient {
	return &defaultSyncClient{
		cli: cli,
	}
}

// SyncUpdatesMe
// sync.updatesMe flags:# user_id:long auth_key_id:long server_id:string session_id:flags.0?long updates:Updates = Void;
func (m *defaultSyncClient) SyncUpdatesMe(ctx context.Context, in *sync.TLSyncUpdatesMe) (*mtproto.Void, error) {
	client := sync.NewRPCSyncClient(m.cli.Conn())
	return client.SyncUpdatesMe(ctx, in)
}

// SyncUpdatesNotMe
// sync.updatesNotMe user_id:long auth_key_id:long updates:Updates = Void;
func (m *defaultSyncClient) SyncUpdatesNotMe(ctx context.Context, in *sync.TLSyncUpdatesNotMe) (*mtproto.Void, error) {
	client := sync.NewRPCSyncClient(m.cli.Conn())
	return client.SyncUpdatesNotMe(ctx, in)
}

// SyncPushUpdates
// sync.pushUpdates user_id:long updates:Updates = Void;
func (m *defaultSyncClient) SyncPushUpdates(ctx context.Context, in *sync.TLSyncPushUpdates) (*mtproto.Void, error) {
	client := sync.NewRPCSyncClient(m.cli.Conn())
	return client.SyncPushUpdates(ctx, in)
}

// SyncPushUpdatesIfNot
// sync.pushUpdatesIfNot user_id:long excludes:Vector<long> updates:Updates = Void;
func (m *defaultSyncClient) SyncPushUpdatesIfNot(ctx context.Context, in *sync.TLSyncPushUpdatesIfNot) (*mtproto.Void, error) {
	client := sync.NewRPCSyncClient(m.cli.Conn())
	return client.SyncPushUpdatesIfNot(ctx, in)
}

// SyncPushBotUpdates
// sync.pushBotUpdates user_id:long updates:Updates = Void;
func (m *defaultSyncClient) SyncPushBotUpdates(ctx context.Context, in *sync.TLSyncPushBotUpdates) (*mtproto.Void, error) {
	client := sync.NewRPCSyncClient(m.cli.Conn())
	return client.SyncPushBotUpdates(ctx, in)
}

// SyncPushRpcResult
// sync.pushRpcResult auth_key_id:long server_id:string session_id:long client_req_msg_id:long rpc_result:bytes = Void;
func (m *defaultSyncClient) SyncPushRpcResult(ctx context.Context, in *sync.TLSyncPushRpcResult) (*mtproto.Void, error) {
	client := sync.NewRPCSyncClient(m.cli.Conn())
	return client.SyncPushRpcResult(ctx, in)
}

// SyncBroadcastUpdates
// sync.broadcastUpdates broadcast_type:int chat_id:long exclude_id_list:Vector<long> updates:Updates = Void;
func (m *defaultSyncClient) SyncBroadcastUpdates(ctx context.Context, in *sync.TLSyncBroadcastUpdates) (*mtproto.Void, error) {
	client := sync.NewRPCSyncClient(m.cli.Conn())
	return client.SyncBroadcastUpdates(ctx, in)
}
